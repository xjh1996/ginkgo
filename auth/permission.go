package auth

import (
	authclient "github.com/caicloud/auth/pkg/server/client"
	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/auth"
	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/rand"
)

var _ = SIGDescribe("命名空间权限管理[permission]", func() {
	// TODO: 添加一个判断用户是否为系统用户操作，若为系统用户，则跳过。
	// 框架会提供一个租户，该租户已经分配好资源，同时提供该租户的管理员用户。目前框架先写死了。
	f := framework.NewDefaultFramework("abc")
	var (
		authAPI              authclient.Interface
		nsName               string
		baseInfo             *auth.BaseInfo
		normalUserAuthAPI    authclient.Interface
		permission, resource []string
		err                  error
	)
	// 指定namespace新建更新配额值
	nsQuotaOld := auth.GenerateNSQuotaString(auth.NamespceMetadate{
		LimitCPU:         "0.25",
		LimitMem:         "0.5Gi",
		RequestCPU:       "0.1",
		RequestMem:       "0.2Gi",
		StorageClassName: "aaa",
		StorageSize:      "0",
		PVCSize:          "0",
		GPU:              "0",
	})
	nsQuotaNew := auth.GenerateNSQuotaString(auth.NamespceMetadate{
		LimitCPU:         "0.2",
		LimitMem:         "0.5Gi",
		RequestCPU:       "0.1",
		RequestMem:       "0.2Gi",
		StorageClassName: "aaa",
		StorageSize:      "0",
		PVCSize:          "0",
		GPU:              "0",
	})
	ginkgo.Describe("管理权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("管理权限", func() {
			permission = []string{"ManageNamespace"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID} // 格式trn:cps:::resourceType/resourceValue,具体Type，Value和开发沟通，或参考https://bytedance.feishu.cn/docs/doccnUdvIc3bCQ724C87idUQWIe#
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{true, true, true, true, true}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("新建权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.DeleteNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
			expect.NoError(err)
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("新建权限", func() {
			permission = []string{"CreateNamespace"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID} // 格式trn:cps:::resourceType/resourceValue,具体Type，Value和开发沟通，或参考https://bytedance.feishu.cn/docs/doccnUdvIc3bCQ724C87idUQWIe#
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{true, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("删除权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			_, err = auth.CreateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, nsQuotaOld, baseInfo.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("删除权限", func() {
			permission = []string{"DeleteNamespace"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + nsName} // 格式trn:cps:::resourceType/resourceValue,具体Type，Value和开发沟通，或参考https://bytedance.feishu.cn/docs/doccnUdvIc3bCQ724C87idUQWIe#
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{false, true, true, false, true}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("更新权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			_, err = auth.CreateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, nsQuotaOld, baseInfo.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.DeleteNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
			expect.NoError(err)
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("更新权限", func() {
			permission = []string{"UpdateNamespace"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + nsName} // 格式trn:cps:::resourceType/resourceValue,具体Type，Value和开发沟通，或参考https://bytedance.feishu.cn/docs/doccnUdvIc3bCQ724C87idUQWIe#
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{false, true, true, true, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("查看权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			_, err = auth.CreateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, nsQuotaOld, baseInfo.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.DeleteNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
			expect.NoError(err)
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("查看权限", func() {
			permission = []string{"VisitNamespace"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/cluster/" + f.ClusterID + "/" + nsName} // 格式trn:cps:::resourceType/resourceValue,具体Type，Value和开发沟通，或参考https://bytedance.feishu.cn/docs/doccnUdvIc3bCQ724C87idUQWIe#
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{false, true, true, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
	ginkgo.Describe("无权限", func() {
		ginkgo.BeforeEach(func() {
			// 创建基础变量并赋值
			nsName = rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			_, err = auth.CreateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, nsQuotaOld, baseInfo.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = auth.DeleteNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
			expect.NoError(err)
			err = auth.PostsetOperation(authAPI, baseInfo)
			expect.NoError(err)
		})
		ginkgo.It("无权限", func() {
			permission = []string{""}
			normalUserAuthAPI = auth.GetNormalUserAuthAPI(authAPI, baseInfo, permission, resource)
			errs := crudNamespace(normalUserAuthAPI, baseInfo, nsName, nsQuotaOld, nsQuotaNew)
			auth.CheckResult(errs, []bool{false, false, false, false, false}) // 顺序create, get, list, update, delete权限
		})
	})
})

func crudNamespace(authAPI authclient.Interface, baseInfo *auth.BaseInfo, nsName, quota, newquota string) []error {
	var errs []error
	var err error
	// 验证create权限
	_, err = auth.CreateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, quota, baseInfo.ClusterID)
	errs = append(errs, err)
	// 验证get权限
	_, err = auth.GetNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
	errs = append(errs, err)
	// 验证list权限
	_, err = auth.ListNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID)
	errs = append(errs, err)
	// 验证update权限
	_, err = auth.UpdateNamespaceAndWait(authAPI, baseInfo.TenantID, nsName, newquota, baseInfo.ClusterID)
	errs = append(errs, err)
	// 验证delete权限
	err = auth.DeleteNamespace(authAPI, baseInfo.TenantID, baseInfo.ClusterID, nsName)
	errs = append(errs, err)
	return errs
}
