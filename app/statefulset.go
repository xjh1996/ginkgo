package app

import (
	"context"

	"github.com/caicloud/nubela/expect"

	"github.com/caicloud/zeus/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSet struct {
	nsName            string
	sfsName           string
	kind              string
	replicanum        int64
	pvcreadonlystatus bool
	configpvcname     string
	pvcmount          string
	pvcname           string
	pvcCapacity       string
	image             string
	svcName           string
	headlesssvcport   int64
}

//以下为手动创建的Statefulset参数(临时使用)
var statefulSetData = StatefulSet{
	"az1",
	"sfs2",
	"statefulset",
	1,
	true,
	"pvcsfs",
	"/pvcsfslujing",
	"pvcsfs-sfs2-0",
	"20Gi",
	"cargo30.dev.caicloud.xyz/library/nginx:1.13.8-alpine",
	"hdls",
	65531,
}

var _ = SIGDescribe("有状态工作负载", func() {

	f := framework.NewDefaultFramework("Statefulset-basic")

	ginkgo.Context("基础部署", func() {
		ginkgo.It("创建", func() {
			testCreateStatefulset(f)
		})

	})

	ginkgo.Context("服务管理", func() {
		ginkgo.It("修改副本和镜像", func() {
			testUpdateStatefulset(f)
		})
		ginkgo.It("停止工作负载", func() {
			testStopStatefulset(f)
		})
		ginkgo.It("回滚工作负载", func() {
			testRollout(f)
		})
		ginkgo.It("修改存储挂载路径&读写路径", func() {
			testUpdatePvcParameter(f)
		})
		ginkgo.It("修改Headless Svc端口", func() {
			testUpdateHeadlesssvc(f)
		})
	})

})

func testCreateStatefulset(f *framework.Framework) {
	//TODO 业务创建Statefulset
	k8scl := f.ClientSet.K8S

	// check 2： 检测 k8s 一致性
	k8sStatefulset, err := k8scl.AppsV1().StatefulSets(statefulSetData.nsName).Get(context.TODO(), statefulSetData.sfsName, metav1.GetOptions{})
	expect.NoError(err, "K8sStatefulset Cmd Build Failed")

	gomega.Expect(k8sStatefulset.ObjectMeta.Labels["controller.caicloud.io/kind"]).To(gomega.BeEquivalentTo(statefulSetData.kind), "服务类型错误")
	gomega.Expect(statefulSetData.replicanum).To(gomega.BeNumerically("==", *k8sStatefulset.Spec.Replicas), "副本数量错误")
	gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].Image).To(gomega.Equal(statefulSetData.image), "绑定的镜像错误")

	gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].ReadOnly).To(gomega.BeEquivalentTo(statefulSetData.pvcreadonlystatus), "绑定PVC的读写权限错误")
	gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].MountPath).To(gomega.BeEquivalentTo(statefulSetData.pvcmount), "绑定PVC的路径错误")
	gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].Name).To(gomega.BeEquivalentTo(statefulSetData.configpvcname), "绑定PVC的名称错误")

	k8sSvc, err := k8scl.CoreV1().Services(statefulSetData.nsName).Get(context.TODO(), statefulSetData.svcName, metav1.GetOptions{})
	expect.NoError(err, "K8s Svc Cmd Build Failed")
	gomega.Expect(k8sSvc.Spec.Selector["controller.caicloud.io/name"]).To(gomega.BeEquivalentTo(statefulSetData.sfsName), "Headless svc名称错误")
	gomega.Expect(k8sSvc.Spec.Ports[0].Port).To(gomega.BeNumerically("==", statefulSetData.headlesssvcport), "Headless svc port错误")

	k8sStorage, err := k8scl.CoreV1().PersistentVolumeClaims(statefulSetData.nsName).Get(context.TODO(), statefulSetData.pvcname, metav1.GetOptions{})
	expect.NoError(err, "K8s Storage Cmd Build Failed")
	gomega.Expect(statefulSetData.pvcCapacity).To(gomega.BeEquivalentTo(k8sStorage.Spec.Resources.Requests.Storage().String()), "有状态工作负载绑定的Pvc容量错误")

	// TODO 访问、创建文件、查看日志、使用终端 ...
	// check 3： 业务层面检查

}

func testUpdateStatefulset(f *framework.Framework) {
	// TODO 基础服务
	//k8scl := f.K8sClientSet
	//nsName := "az1"
	// TODO 业务层面修改副本数和镜像

	//gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].Image).To(gomega.Equal(image),"镜像修改失败")
	//gomega.Expect(*k8sStatefulset.Spec.Replicas).To(gomega.BeEquivalentTo(replicas),"副本数修改失败")

}

func testStopStatefulset(f *framework.Framework) {
	//TODO 业务停止操作
	nsName := "az1"
	pvcname := "pvcsfs-sfs2-0"
	pvcCapacity := "20Gi"
	//检测服务已经停止
	//检测数据卷仍在挂载
	k8scl := f.ClientSet.K8S
	k8sStorage, err := k8scl.CoreV1().PersistentVolumeClaims(nsName).Get(context.TODO(), pvcname, metav1.GetOptions{})
	expect.NoError(err)
	gomega.Expect(pvcCapacity).To(gomega.BeEquivalentTo(k8sStorage.Spec.Resources.Requests.Storage().String()), "有状态工作负载绑定的Pvc容量错误")

}

func testRollout(f *framework.Framework) {

}

func testUpdatePvcParameter(f *framework.Framework) {
	//TODO 对业务的存储进行修改挂载路径和读写权限的操作
	//TODO 业务更改pvc状态从只读改为读写
	//TODO 业务更改挂载路径为新路径

	//gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].ReadOnly).To(gomega.BeEquivalentTo(pvcreadonlystatus),"绑定PVC的读写权限修改失败")
	//gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].MountPath).To(gomega.BeEquivalentTo(pvcmount),"绑定PVC的路径修改失败")
	//gomega.Expect(k8sStatefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].Name).To(gomega.BeEquivalentTo(configpvcname),"绑定PVC的名称错误")

}

func testUpdateHeadlesssvc(f *framework.Framework) {
	// TODO 对业务的Headless svc端口进行修改

	//newHeadlessSvcport
	//gomega.Expect(k8sSvc.Spec.Selector["controller.caicloud.io/name"]).To(gomega.BeEquivalentTo(sfsName),"Headless svc名称错误")
	//gomega.Expect(k8sSvc.Spec.Ports[0].Port).To(gomega.BeNumerically("==",newHeadlessSvcport),"Headless svc port修改错误")

}
