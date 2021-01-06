package auth

import (
	"context"
	"strings"

	authclient "github.com/caicloud/auth/pkg/server/client"
	v20201010 "github.com/caicloud/auth/pkg/server/client/v20201010"
	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework/client"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/wait"
)

var operation = []string{"create", "get", "list", "update", "delete"}

const (
	passwd = "Pwd123456"
)

type BaseInfo struct {
	RoleName  string
	UserName  string
	Email     string
	TenantID  string
	ClusterID string
}

func AddTenantMembers(authAPI authclient.Interface, tenantID, userName string, roleNames []string) error {
	addTenantMemReq := &v20201010.AddTenantMemberReq{
		TenantID: tenantID,
		Items:    []v20201010.OperateTenantMember{{UserName: userName, RoleNames: roleNames}},
	}
	_, err := authAPI.V20201010().AddTenantMembers(context.TODO(), addTenantMemReq)
	return err
}

func CreateSingleUserAndWait(authAPI authclient.Interface, name, email, passwd string) error {
	userReq := v20201010.UserReq{
		Name:     name,
		Nick:     name,
		Email:    email,
		Password: passwd,
	}
	createUserReq := &v20201010.CreateUserReq{Items: []v20201010.UserReq{userReq}}
	if _, err := authAPI.V20201010().CreateUsers(context.TODO(), createUserReq); err != nil {
		return err
	}
	return wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		user, err := GetUser(authAPI, name)
		if err != nil {
			return false, err
		}
		if user.State == "Normal" { // FIXME: this field value need verify.
			return true, nil
		} else {
			logger.Infof("tenant is not ready, %q, retrying...", user.State)
			return false, nil
		}
	})
}

func CreateRole(authAPI authclient.Interface, name string, policy, resource []string) (*v20201010.Role, error) {
	createRole := &v20201010.CreateRoleReq{
		Name:        name,
		PolicyNames: policy,
		Resource:    resource,
	}
	return authAPI.V20201010().CreateRole(context.TODO(), createRole)
}

func CreateRoleBinding(authAPI authclient.Interface, roleID string, userIDs []string) error {
	roleBindingReq := &v20201010.RoleBindReq{
		UID:   roleID,
		Items: userIDs,
	}
	_, err := authAPI.V20201010().RoleBindUsers(context.TODO(), roleBindingReq)
	return err
}

func GetUser(authAPI authclient.Interface, name string) (*v20201010.UserResp, error) {
	getUserReq := &v20201010.GetUserReq{Name: name}
	return authAPI.V20201010().GetUser(context.TODO(), getUserReq)
}

func PresetOperation(authAPI authclient.Interface, baseInfo *BaseInfo, permission, resource []string) authclient.Interface {
	// 创建普通用户
	var err error
	if err = CreateSingleUserAndWait(authAPI, baseInfo.UserName, baseInfo.Email, passwd); err != nil {
		logger.Failf("create normal user failed, %v", err)
	}
	// 将用户添加到租户
	err = AddTenantMembers(authAPI, baseInfo.TenantID, baseInfo.UserName, []string{})
	//if err != nil {
	//	logger.Failf("add user to tenant failed, %v", err)   // FIXME: 一直报409错误（名字冲突），待与开发对齐
	//}
	// 创建角色，并将普通用户和角色绑定
	role, err := CreateRole(authAPI, baseInfo.RoleName, permission, resource)
	if err != nil {
		logger.Failf("create role failed, %v", err)
	}
	normalUser, err := GetUser(authAPI, baseInfo.UserName)
	if err != nil {
		logger.Failf("get normal user failed, %v", err)
	}
	if err = CreateRoleBinding(authAPI, role.UID, []string{normalUser.Name}); err != nil {
		logger.Failf("create role binding failed, %v", err)
	}
	// 获取普通用户API
	user := client.User{
		Tenant:   baseInfo.TenantID,
		Username: baseInfo.UserName,
		Password: passwd,
	}
	normalUserAuthAPI, err := user.Auth()
	if err != nil {
		logger.Failf("get normal user api failed, %v", err)
	}
	// check绑定成功
	roles, err := normalUserAuthAPI.V20201010().ListUserRole(context.TODO())
	if roles.Items[0].Name != role.Name { // 测试中一个用户只绑定一个角色
		logger.Failf("bindding role failed, expected %q, binded %q, all roles %q, userName %q", role.Name, roles.Items[0].Name, roles.Items, user.Username)
	}
	return normalUserAuthAPI
}

func CreateBaseInfo(tenantID, clusterID string) *BaseInfo {
	return &BaseInfo{
		RoleName:  rand.String(16),
		UserName:  rand.String(16),
		Email:     rand.String(6) + "@cai.com",
		TenantID:  tenantID,
		ClusterID: clusterID,
	}
}

func DeleteUserAndWait() error {
	// TODO: 业务组还未提供删除User的API
	return nil
}

func DeleteRole(authAPI authclient.Interface, name string) error {
	delRole := &v20201010.DeleteRoleReq{
		UID: name,
	}
	_, err := authAPI.V20201010().DeleteRole(context.TODO(), delRole)
	return err
}
func PostsetOperation(authAPI authclient.Interface, baseInfo *BaseInfo) error {
	if err := DeleteRole(authAPI, baseInfo.RoleName); err != nil {
		return err
	}
	if err := DeleteUserAndWait(); err != nil {
		return err
	}
	return nil
}

func CheckResult(errs []error, expects []bool) {
	for i := 0; i < len(errs); i++ {
		if errs[i] == nil {
			expect.Equal(expects[i], true, operation[i]+" operation check failed")
		} else {
			if strings.Contains(errs[i].Error(), "authorization not allowed") { // FIXME: 返回的err无错误码，需要再次和效能团队沟通，
				// 根据文档https://bytedance.feishu.cn/docs/doccnuBNBNBgRQ2lNsQheCpZObf#WU7vWG， authorization not allowed表示无权限操作
				expect.Equal(expects[i], false, operation[i]+" operation check failed"+errs[i].Error())
			} else {
				logger.Failf("operation %q failed, %v", operation[i], errs[i].Error())
			}
		}
	}
}
