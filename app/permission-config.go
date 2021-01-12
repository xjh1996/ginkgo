package app

import (
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"
	"github.com/caicloud/zeus/framework/auth"

	appClient "github.com/caicloud/app/pkg/server/client"
	authClient "github.com/caicloud/auth/pkg/server/client"
	"github.com/caicloud/nubela/expect"

	"github.com/onsi/ginkgo"
)

var _ = SIGDescribe("配置项权限管理[permission]", func() {
	// TODO: 添加一个判断用户是否为系统用户操作，若为系统用户，则跳过。
	// 框架会提供一个租户，该租户已经分配好资源，同时提供该租户的管理员用户。目前框架先写死了。
	f := framework.NewDefaultFramework("abc")
	var (
		authAPI              authClient.Interface
		baseInfo             *auth.BaseInfo
		configBaseInfo       *app.ConfigBaseInfo
		permission, resource []string
		err                  error
	)
	ginkgo.Describe("查看权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			configBaseInfo = app.CreateConfigInfo()
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			_, err = app.CreateCfmAndWait(appAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			// TODO: 框架暂时有问题 所有AfterEach均注释 后续等待上游修改再修改
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("查看权限", func() {
			permission = []string{"VisitConfigMap"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudConfig(normalUserAppAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("新建权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			configBaseInfo = app.CreateConfigInfo()
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("新建权限", func() {
			permission = []string{"CreateConfigMap"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudConfig(normalUserAppAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			auth.CheckResult(errs, []bool{true, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("更新权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			configBaseInfo = app.CreateConfigInfo()
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			_, err = app.CreateCfmAndWait(appAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("更新权限", func() {
			permission = []string{"UpdateConfigMap"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudConfig(normalUserAppAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, true, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("删除权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			configBaseInfo = app.CreateConfigInfo()
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			_, err = app.CreateCfmAndWait(appAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("删除权限", func() {
			permission = []string{"DeleteConfigMap"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudConfig(normalUserAppAPI, configBaseInfo.ConfigName, namespace, configBaseInfo.Key, configBaseInfo.Value, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, false, true}) // 顺序create, get, list, update, delete权限
		})
	})
})

func crudConfig(appAPI appClient.Interface, configName, namespace, key, value, clusterID string) []error {
	var errs []error
	var err error
	// 验证create权限
	_, err = app.CreateCfmAndWait(appAPI, configName, namespace, key, value, clusterID)
	errs = append(errs, err)
	// 验证get权限
	_, err = app.GetCfm(appAPI, configName, namespace, clusterID)
	errs = append(errs, err)
	// 验证list权限
	_, err = app.ListCfm(appAPI, namespace, clusterID)
	errs = append(errs, err)
	// 验证update权限
	_, err = app.UpdateCfmAndWait(appAPI, configName, namespace, "new"+key, "new"+value, clusterID)
	errs = append(errs, err)
	// 验证delete权限
	err = app.DeleteCfm(appAPI, configName, namespace, clusterID)
	errs = append(errs, err)
	return errs
}