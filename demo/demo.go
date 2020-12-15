package demo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/config"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = SIGDescribe("无状态服务基础部署", func() {
	//var k8scl clientset.Interface
	//var ns string

	f := framework.NewDefaultFramework("deployment-basic")
	ginkgo.BeforeEach(func() {
		//k8scl = f.K8sClientSet
		//ns = f.Namespace.Name
		//ns = "default"
	})

	ginkgo.Context("使用不同镜像", func() {
		ginkgo.It("自定义镜像，容器数量和存储数量为1，部署能够成功", func() {
			// 建议用例内容比较少的情况下直接编写，如果用例内容较多，写一个 testxxx 函数较为清爽简洁
			testBasicDeployment(f)
		})
		ginkgo.It("选择集群内镜像，容器数量和存储数量为1，部署能够成功", func() {
		})
	})
	ginkgo.Context("会话保持", func() {
		ginkgo.It("开关状态修改，配置生效", func() {
		})
	})
})

var _ = SIGDescribe("无状态服务高级配置", func() {})

func testBasicDeployment(f *framework.Framework) {
	c := f.ClientSet.K8S
	//ns = f.Namespace.Name
	ns := "default"

	// TODO 业务完成创建服务操作
	// check 1： 业务完成查询操作
	// check 2： 检测 k8s 一致性

	k8sDeployment, err := c.AppsV1().Deployments(ns).Get(context.TODO(), "app-admin-admin-v1-0", metav1.GetOptions{})
	expect.NoError(err)
	gomega.Expect(1).To(gomega.BeNumerically("==", *k8sDeployment.Spec.Replicas), "服务副本数量不正确")
	logger.Infof("deployment status: %v", k8sDeployment.Status.Conditions)

	//check 3： 业务层面检查
	url := "http://baidu.com" //  从业务资源获取
	gomega.Expect(testConnection(url)).Should(gomega.HaveHTTPStatus(http.StatusOK), "服务 nodeport 类型 service %s 无法通信", url)

	/*
		获取 crd 的临时解决方案
		如下暂时可以获取 2.x 版本的 crd 资源，但需要自己传入 crd group version 等资源（目前框架不统一提供该功能），后续会对接效能团队提供的 client 更加便捷的获取 crd 资源
	*/
	config, err := clientcmd.BuildConfigFromFlags("", config.Context.KubeConfig)
	dynamicClient, err := dynamic.NewForConfig(config)
	deploymentResource := schema.GroupVersionResource{Group: "release.caicloud.io", Version: "v1alpha1", Resource: "releases"}
	k8sRelease, err := dynamicClient.Resource(deploymentResource).Namespace(ns).Get(context.TODO(), "user-web", metav1.GetOptions{})
	expect.NoError(err)
	name, _, err := unstructured.NestedString(k8sRelease.Object, "metadata", "name")
	expect.NoError(err)
	gomega.Expect("user-web").To(gomega.BeEquivalentTo(name), "未找到 user-web")

	/*
		cos 层使用
	*/
	//f.K8sCRDClient.Cos.NetworkingV1beta1().LoadBalancers(ns).Get(context.TODO(), "test", metav1.GetOptions{})
}

func testConnection(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request is fail ")
	}
	//http client
	client := &http.Client{}
	logger.Infof("Get %s URL : %s \n", http.MethodGet, req.URL.String())
	return client.Do(req)
}
