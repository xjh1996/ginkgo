package app

import (
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"
	"github.com/caicloud/zeus/framework/auth"

	appClient "github.com/caicloud/app/pkg/server/client"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
	authClient "github.com/caicloud/auth/pkg/server/client"
	"github.com/caicloud/nubela/expect"

	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/rand"
)

var _ = SIGDescribe("无状态工作负载权限管理[permission]abc", func() {
	// TODO: 添加一个判断用户是否为系统用户操作，若为系统用户，则跳过。
	// 框架会提供一个租户，该租户已经分配好资源，同时提供该租户的管理员用户。目前框架先写死了。
	f := framework.NewDefaultFramework("abc")
	var (
		authAPI              authClient.Interface
		baseInfo             *auth.BaseInfo
		deploymentName       string
		rpNum                int32
		permission, resource []string
		err                  error
	)
	ginkgo.Describe("查看权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			deploymentName = "dp-" + rand.String(20)
			rpNum = int32(rand.Intn(10))
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			deployment := app.NewDeployment(deploymentName, namespace, rpNum, func(deployment *types.Deployment) {})
			_, err = app.CreateDeployment(appAPI, deployment, f.ClusterID, namespace, deploymentName)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			// TODO: 框架暂时有问题 所有AfterEach均注释 后续等待上游修改再修改
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("查看权限", func() {
			permission = []string{"VisitDeployment"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudDeployment(normalUserAppAPI, deploymentName, namespace, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("新建权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			deploymentName = "dp-" + rand.String(20)
			rpNum = int32(rand.Intn(10))
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("新建权限", func() {
			permission = []string{"CreateDeployment"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudDeployment(normalUserAppAPI, deploymentName, namespace, f.ClusterID)
			auth.CheckResult(errs, []bool{true, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("更新权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			deploymentName = "dp-" + rand.String(20)
			rpNum = int32(rand.Intn(10))
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			deployment := app.NewDeployment(deploymentName, namespace, rpNum, func(deployment *types.Deployment) {})
			_, err = app.CreateDeployment(appAPI, deployment, f.ClusterID, namespace, deploymentName)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("更新权限", func() {
			permission = []string{"UpdateDeployment"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudDeployment(normalUserAppAPI, deploymentName, namespace, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, true, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("删除权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			deploymentName = "dp-" + rand.String(20)
			rpNum = int32(rand.Intn(10))
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			appAPI, err := f.AdminAPIClient.App()
			expect.NoError(err)
			deployment := app.NewDeployment(deploymentName, namespace, rpNum, func(deployment *types.Deployment) {})
			_, err = app.CreateDeployment(appAPI, deployment, f.ClusterID, namespace, deploymentName)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("删除权限", func() {
			permission = []string{"DeleteDeployment"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + namespace}
			normalUserAppAPI := app.GetNormalUserAppAPI(authAPI, baseInfo, permission, resource)
			errs := crudDeployment(normalUserAppAPI, deploymentName, namespace, f.ClusterID)
			auth.CheckResult(errs, []bool{false, true, true, false, true}) // 顺序create, get, list, update, delete权限
		})
	})
})

func crudDeployment(appAPI appClient.Interface, deploymentName, namespace, clusterID string) []error {
	var errs []error
	var err error
	rpNum := int32(rand.Intn(10))
	deployment := app.NewDeployment(deploymentName, namespace, rpNum, func(deployment *types.Deployment) {})
	// 验证create权限
	_, err = app.CreateDeployment(appAPI, deployment, clusterID, namespace, deploymentName)
	errs = append(errs, err)
	// 验证get权限
	_, err = app.GetDeployment(appAPI, deploymentName, namespace, clusterID)
	errs = append(errs, err)
	// 验证list权限
	_, err = app.ListDeployment(appAPI, namespace, clusterID)
	errs = append(errs, err)
	// 验证update权限
	_, err = app.UpdateDeployment(appAPI, deployment, clusterID, namespace, deploymentName)
	errs = append(errs, err)
	// 验证delete权限
	err = app.DeleteDeployment(appAPI, deploymentName, namespace, clusterID)
	errs = append(errs, err)
	return errs
}
