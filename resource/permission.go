package resource

import (
	"context"

	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/auth"
	client "github.com/caicloud/zeus/framework/config"

	authclient "github.com/caicloud/auth/pkg/server/client"
	"github.com/caicloud/nubela/expect"
	resourceclient "github.com/caicloud/resource/pkg/server/client"
	v20201010 "github.com/caicloud/resource/pkg/server/client/v20201010"

	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/rand"
)

var _ = SIGDescribe("数据卷权限测试[permission]", func() {
	// TODO: 添加一个判断用户是否为系统用户操作，若为系统用户，则跳过。
	// 框架会提供一个租户，该租户已经分配好资源，同时提供该租户的管理员用户。目前框架先写死了。
	f := framework.NewFramework("PVC Permission", false, true)
	var (
		authAPI              authclient.Interface
		resourceAPI          resourceclient.Interface
		PVCName              string
		nsName               = "cluster/kubernetes-stack/p2"
		storageClass         = "glusterfs-20210106090307-c67496"
		baseInfo             *auth.BaseInfo
		normalUser           client.User
		permission, resource []string
		err                  error
	)

	ginkgo.Describe("数据卷新建权限", func() {
		ginkgo.BeforeEach(func() {
			PVCName = "pvc-" + rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			resourceAPI, err = f.APIClient.Resource()
			expect.NoError(err)
			//TODO 创建含有存储方案的分区

		})
		ginkgo.AfterEach(func() {
			//TODO 删除分区，暂时先删除数据卷
			err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), f.ClusterID, nsName, PVCName)
			expect.NoError(err)
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("数据卷新建权限", func() {
			permission = []string{"CreateStorageVolume"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/" + nsName}
			normalUser = auth.PresetOperation(authAPI, baseInfo, permission, resource)
			resource, err := normalUser.Resource()
			expect.NoError(err)
			errs := crudPVC(resource, f.ClusterID, nsName, PVCName, storageClass)
			auth.CheckResult(errs, []bool{true, true, true, false, false}) // 顺序create, get, list, update, delete权限, 暂时没有做upload权限的测试
		})
	})

	ginkgo.Describe("删除权限", func() {
		ginkgo.BeforeEach(func() {
			PVCName = "pvc-" + rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			resourceAPI, err = f.APIClient.Resource()
			expect.NoError(err)
			//TODO 创建含有存储方案的分区
			_, err = CreatePVCAndWait(resourceAPI, nsName, PVCName, storageClass, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//TODO 删除分区，暂时先删除数据卷
			err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), f.ClusterID, nsName, PVCName)
			expect.NoError(err)
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("删除权限", func() {
			permission = []string{"DeleteStorageVolume"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/" + nsName}
			normalUser = auth.PresetOperation(authAPI, baseInfo, permission, resource)
			resource, err := normalUser.Resource()
			expect.NoError(err)
			errs := crudPVC(resource, f.ClusterID, nsName, PVCName, storageClass)
			auth.CheckResult(errs, []bool{false, true, true, false, true})
		})
	})

	ginkgo.Describe("更新权限", func() {
		ginkgo.BeforeEach(func() {
			PVCName = "pvc-" + rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			resourceAPI, err = f.APIClient.Resource()
			expect.NoError(err)
			//TODO 创建含有存储方案的分区
			_, err = CreatePVCAndWait(resourceAPI, nsName, PVCName, storageClass, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//TODO 删除分区，暂时先删除数据卷
			err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), f.ClusterID, nsName, PVCName)
			expect.NoError(err)
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("更新权限", func() {
			permission = []string{"UpdateStorageVolume"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/" + nsName}
			normalUser = auth.PresetOperation(authAPI, baseInfo, permission, resource)
			resource, err := normalUser.Resource()
			expect.NoError(err)
			errs := crudPVC(resource, f.ClusterID, nsName, PVCName, storageClass)
			auth.CheckResult(errs, []bool{false, true, true, true, false}) // 顺序create, get, list, update, delete权限, 暂时没有做upload权限的测试
		})
	})

	ginkgo.Describe("数据卷上传权限", func() {
		//TODO 暂时开发还没实现
	})

	ginkgo.Describe("查看权限", func() {
		ginkgo.BeforeEach(func() {
			PVCName = "pvc-" + rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			resourceAPI, err = f.APIClient.Resource()
			expect.NoError(err)
			//TODO 创建含有存储方案的分区
			_, err = CreatePVCAndWait(resourceAPI, nsName, PVCName, storageClass, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			//TODO 删除分区，暂时先删除数据卷
			err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), f.ClusterID, nsName, PVCName)
			expect.NoError(err)
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("查看权限", func() {
			permission = []string{"VisitStorageVolume"}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace/" + nsName}
			normalUser = auth.PresetOperation(authAPI, baseInfo, permission, resource)
			resource, err := normalUser.Resource()
			expect.NoError(err)
			errs := crudPVC(resource, f.ClusterID, nsName, PVCName, storageClass)
			auth.CheckResult(errs, []bool{false, true, true, false, false}) // 顺序create, get, list, update, delete权限, 暂时没有做upload权限的测试
		})
	})

	ginkgo.Describe("无权限", func() {
		ginkgo.BeforeEach(func() {
			PVCName = "pvc-" + rand.String(8)
			baseInfo = auth.CreateBaseInfo(f.PresetResource.Auth.TenantID, f.ClusterID)
			authAPI, err = f.APIClient.Auth()
			expect.NoError(err)
			resourceAPI, err = f.APIClient.Resource()
			expect.NoError(err)
			//TODO 创建含有存储方案的分区
			_, err = CreatePVCAndWait(resourceAPI, nsName, PVCName, storageClass, f.ClusterID)
			expect.NoError(err)
		})
		ginkgo.AfterEach(func() {
			err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), f.ClusterID, nsName, PVCName)
			expect.NoError(err)
			//err = auth.PostsetOperation(authAPI, baseInfo)
			//expect.NoError(err)
		})
		ginkgo.It("无权限", func() {
			permission = []string{""}
			resource = []string{"trn:cps:::cluster/" + f.ClusterID, "trn:cps:::namespace" + nsName}
			normalUser = auth.PresetOperation(authAPI, baseInfo, permission, resource)
			resource, err := normalUser.Resource()
			expect.NoError(err)
			errs := crudPVC(resource, f.ClusterID, nsName, PVCName, storageClass)
			auth.CheckResult(errs, []bool{false, false, false, false, false}) // 顺序create, get, list, update, delete权限, 暂时没有做upload权限的测试
		})
	})
})

func crudPVC(resourceAPI resourceclient.Interface, clusterid, nsName, PVCName, storageClass string) []error {
	var errs []error
	var err error
	// 验证create权限
	_, err = CreatePVCAndWait(resourceAPI, nsName, PVCName, storageClass, clusterid)
	errs = append(errs, err)
	// 验证get权限
	_, err = resourceAPI.V20201010().GetPersistentVolumeClaim(context.TODO(), clusterid, nsName, PVCName)
	errs = append(errs, err)
	// 验证list权限
	f := false
	_, err = resourceAPI.V20201010().ListPersistentVolumeClaim(context.TODO(), clusterid, nsName, "", "", &f, 1, 10)
	errs = append(errs, err)
	// 验证update权限
	updatePVCReq := &v20201010.UpdatePVCRequest{
		Cluster:   clusterid,
		Namespace: nsName,
		Name:      PVCName,
		Size:      "2Gi",
	}
	_, err = resourceAPI.V20201010().UpdatePersistentVolumeClaim(context.TODO(), updatePVCReq)
	errs = append(errs, err)
	//TODO 验证数据卷上传权限

	// 验证delete权限
	err = resourceAPI.V20201010().DeletePersistentVolumeClaim(context.TODO(), clusterid, nsName, PVCName)
	errs = append(errs, err)
	return errs
}

