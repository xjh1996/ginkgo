package resource

import (
	"context"
	"encoding/json"

	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/config"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	namespaces         = "az1"
	storageService     = "nfs-20201118135156-f9a5e4"
	storageServiceNum  = "300"
	storageServiceCap  = "300Gi"
	storageServiceType = "nfs"
	storageClass       = "nfs-20201118135210-aee099"
	storageClassNum    = "300"
	storageClassCap    = "300Gi"
	pvc                = "pvc"
	pvc1               = "pvcdelete"
	storageClass1      = "scdelete"
	storageService1    = "ssdelete"
	namespaces1        = "az2"
)

type NamespceMetadate struct {
	nsName          string
	limitCPU        string
	limitMem        string
	requestCPU      string
	requestMem      string
	storageClassNum string
	storageClassCap string
	requestGPU      string
}

var nsQuota = NamespceMetadate{namespaces, "4", "10Gi", "2", "6Gi", "200", "200Gi", "0"}

var _ = SIGDescribe("存储管理", func() {
	f := framework.NewFramework("volume", false, true)
	ginkgo.BeforeEach(func() {
	})
	ginkgo.Context("不同存储类型", func() {
		ginkgo.It("NFS存储可用", func() {
			testVolumeAvailable(f)
		})
	})
})

func testVolumeAvailable(f *framework.Framework) {
	var k8scl clientset.Interface
	k8scl = f.ClientSet.K8S
	storageServiceCreate()
	storageClasseCreate(k8scl)
	nameSpacesCreate(k8scl, nsQuota)
	pvcCreate(k8scl)
	// DeploymentCreate()
	// DeploymentDelete()
	pvcDelete(k8scl)
	nameSpacesDelete(k8scl, namespaces1)
	storageClasseDelete(k8scl)
	storageServiceDelete()
}

func storageServiceCreate() {
	// TODO 业务完成创建存储服务操作
	// check 1： TODO 业务完成查询操作，存在此存储服务，且存储服务配置正确
	// check 2： 检测 k8s storageclass 参数一致性
	config, err := clientcmd.BuildConfigFromFlags("", config.Context.KubeConfig)
	expect.NoError(err)
	dynamicClient, err := dynamic.NewForConfig(config)
	expect.NoError(err)
	storageServiceResource := schema.GroupVersionResource{Group: "resource.caicloud.io", Version: "v1beta1", Resource: "storageservices"}
	k8sStorageService, err := dynamicClient.Resource(storageServiceResource).Get(context.TODO(), storageService, metav1.GetOptions{})
	expect.NoError(err)
	gomega.Expect(storageServiceType).To(gomega.BeEquivalentTo(k8sStorageService.Object["typeName"]))
	i, _, err := unstructured.NestedMap(k8sStorageService.Object, "allocated")
	expect.NoError(err)
	gomega.Expect(storageServiceNum).To(gomega.BeEquivalentTo(i["persistentvolumeclaims"]))
	gomega.Expect(storageServiceCap).To(gomega.BeEquivalentTo(i["requests.storage"]))
	logger.Infof("Storageservice k8s verified successfully")
}

func storageClasseCreate(k8scl clientset.Interface) {
	// TODO 业务完成创建存储方案操作
	// check 1： TODO 业务完成查询操作
	// check 2： 检测 k8s storageclass 参数一致性
	k8sSc, err := k8scl.StorageV1().StorageClasses().Get(context.TODO(), storageClass, metav1.GetOptions{})
	expect.NoError(err)
	scJson := (k8sSc.Annotations["storage.resource.caicloud.io/storageclass-quota"])
	type scInfo struct {
		Persistentvolumeclaims string `json:"persistentvolumeclaims"`
		Requestsstorage        string `json:"requests.storage"`
	}
	scinfo := scInfo{}
	err = json.Unmarshal([]byte(scJson), &scinfo)
	if err != nil {
		logger.Failf("Deserialization failed", err)
	}
	gomega.Expect(storageClassNum).Should(gomega.BeEquivalentTo(scinfo.Persistentvolumeclaims))
	gomega.Expect(storageClassCap).Should(gomega.BeEquivalentTo(scinfo.Requestsstorage))
	logger.Infof("Storageclass k8s verified successfully")
}

func nameSpacesCreate(k8scl clientset.Interface, ns NamespceMetadate) {
	// TODO 业务完成创建分区操作
	// check 1： TODO 业务完成查询操作
	// check 2： 检测 k8s namespaces 参数一致性
	storageClassNum := v1.ResourceName(storageClass + ".storageclass.storage.k8s.io/persistentvolumeclaims")
	storageClassCap := v1.ResourceName(storageClass + ".storageclass.storage.k8s.io/requests.storage")
	resourceQuotas, err := k8scl.CoreV1().ResourceQuotas(ns.nsName).Get(context.TODO(), ns.nsName, metav1.GetOptions{})
	expect.NoError(err)
	gomega.Expect(resourceQuotas.ObjectMeta.Namespace).To(gomega.BeEquivalentTo(ns.nsName))

	rLimitsCpu := resourceQuotas.Spec.Hard["limits.cpu"]
	rLimitsMemory := resourceQuotas.Spec.Hard["limits.memory"]
	rRequestsCpu := resourceQuotas.Spec.Hard["requests.cpu"]
	rRequestsMemory := resourceQuotas.Spec.Hard["requests.memory"]
	rGpu := resourceQuotas.Spec.Hard["requests.nvidia.com/gpu"]
	rstorageClassNum := resourceQuotas.Spec.Hard[storageClassNum]
	rstorageClassCap := resourceQuotas.Spec.Hard[storageClassCap]

	gomega.Expect(rLimitsCpu.String()).To(gomega.BeEquivalentTo(ns.limitCPU))
	gomega.Expect(rLimitsMemory.String()).To(gomega.BeEquivalentTo(ns.limitMem))
	gomega.Expect(rRequestsCpu.String()).To(gomega.BeEquivalentTo(ns.requestCPU))
	gomega.Expect(rRequestsMemory.String()).To(gomega.BeEquivalentTo(ns.requestMem))
	gomega.Expect(rstorageClassNum.String()).To(gomega.BeEquivalentTo(ns.storageClassNum))
	gomega.Expect(rstorageClassCap.String()).To(gomega.BeEquivalentTo(ns.storageClassCap))
	gomega.Expect(rGpu.String()).To(gomega.BeEquivalentTo(ns.requestGPU))
	logger.Infof("Namespace k8s verified successfully")
}

func pvcCreate(k8scl clientset.Interface) {
	//TODO 创建数据卷
	//check 1： TODO  业务完成查询操作，存在数据卷，且数据卷配置正确
	//check 2： 检测 k8s PVC 参数一致性
	k8sPVC, err := k8scl.CoreV1().PersistentVolumeClaims(namespaces).Get(context.TODO(), pvc, metav1.GetOptions{})
	expect.NoError(err)
	gomega.Expect(*k8sPVC.Spec.StorageClassName).Should(gomega.BeEquivalentTo(storageClass))
	gomega.Expect(k8sPVC.Spec.AccessModes[0]).Should(gomega.BeEquivalentTo("ReadWriteMany"))
	gomega.Expect(k8sPVC.Spec.Resources.Requests.Storage().String()).Should(gomega.BeEquivalentTo("1Gi"))
	gomega.Expect(k8sPVC.Status.Phase).Should(gomega.BeEquivalentTo("Bound"))
	logger.Infof("PVC k8s verified successfully")
}

func pvcDelete(k8scl clientset.Interface) {
	//TODO 创建数据卷
	//check 1： TODO  务完成查询操作，不存在PVC
	//check 2： 检测 k8s PVC 被删除
	_, err := k8scl.CoreV1().PersistentVolumeClaims(namespaces).Get(context.TODO(), pvc1, metav1.GetOptions{})
	gomega.Expect(err).To(gomega.HaveOccurred())
	logger.Infof("PVC k8s was deleted")
}

func nameSpacesDelete(k8scl clientset.Interface, ns string) {
	// TODO 业务完成删除分区操作
	// check 1： TODO 业务完成查询操作，不存在分区
	// check 2： 检测 k8s namespace 被删除
	_, err := k8scl.CoreV1().Namespaces().Get(context.TODO(), ns, metav1.GetOptions{})
	gomega.Expect(err).To(gomega.HaveOccurred())
	logger.Infof("Namespace k8s was deleted")
}

func storageClasseDelete(k8scl clientset.Interface) {
	//TODO 业务完成删除存储方案
	//check 1： TODO  业务完成查询操作，不存在存储方案
	//check 2： 检测 k8s storageclass 被删除
	_, err := k8scl.StorageV1().StorageClasses().Get(context.TODO(), storageClass1, metav1.GetOptions{})
	gomega.Expect(err).To(gomega.HaveOccurred())
	logger.Infof("Storageclass k8s was deleted")
}

func storageServiceDelete() {
	// TODO 业务完成删除存储服务操作
	// check 1： TODO 业务完成查询操作，不存在存储服务
	// check 2： 检测 k8s  storageservice 被删除
	config, err := clientcmd.BuildConfigFromFlags("", config.Context.KubeConfig)
	expect.NoError(err)
	dynamicClient, err := dynamic.NewForConfig(config)
	expect.NoError(err)
	storageServiceResource := schema.GroupVersionResource{Group: "resource.caicloud.io", Version: "v1beta1", Resource: "storageservices"}
	_, err = dynamicClient.Resource(storageServiceResource).Get(context.TODO(), storageService1, metav1.GetOptions{})
	gomega.Expect(err).To(gomega.HaveOccurred())
	logger.Infof("Storageservice k8s was deleted")
}
