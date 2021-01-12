package app

import (
	"context"

	cpsmetav1 "github.com/caicloud/api/meta/v1"

	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"

	types "github.com/caicloud/app/pkg/server/client/v20201010"
	"github.com/caicloud/nubela/expect"

	"github.com/ghodss/yaml"
	"github.com/onsi/ginkgo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

// CustomValues is values struct for custom application
type CustomValues struct {
	Deployments  []types.Deployment  `json:"deployments"`
	StatefulSets []types.StatefulSet `json:"statefulSets"`
	Services     []types.Service     `json:"services"`
}

var _ = SIGDescribe("无状态工作负载", func() {
	f := framework.NewDefaultFramework("Deployment-basic")
	ginkgo.Context("基础部署", func() {
		ginkgo.It("创建+查询+删除", func() {
			testCRDDeployment(f)
		})
		ginkgo.FIt("应用测试", func() {
			testCRDApplication(f)
		})
		ginkgo.It("基本信息", func() {
			testListDeployment(f)
		})
	})

})

func testCRDDeployment(f *framework.Framework) {
	dpName := "dp-" + rand.String(20)
	clusterID := f.ClusterID
	rpNum := int32(rand.Intn(10))
	k8sctl := f.ClientSet.K8S

	a, err := f.AdminAPIClient.App()
	expect.NoError(err, "App Client Build Failed")
	client := a.V20201010()

	deployment := app.NewDeployment(dpName, namespace, rpNum, func(deployment *types.Deployment) {})
	_, err = app.CreateDeployment(a, deployment, clusterID, namespace, dpName)
	expect.NoError(err, "Create Deployment Failed")

	k8sDeployment, err := k8sctl.AppsV1().Deployments(namespace).Get(context.TODO(), dpName, metav1.GetOptions{})
	expect.NoError(err, "k8s Client Build Failed")
	//K8s 校验实例数是否正常下发
	expect.Equal(*k8sDeployment.Spec.Replicas, rpNum)

	_, err = client.GetDeployment(context.TODO(), app.NewClusterOption(clusterID, namespace, dpName))
	expect.NoError(err, "Get Deployment Failed")

	err = client.DeleteDeployment(context.TODO(), app.NewClusterOption(clusterID, namespace, dpName))
	expect.NoError(err, "Del Deployment Failed")

	_, err = client.GetDeployment(context.TODO(), app.NewClusterOption(clusterID, namespace, dpName))
	expect.Error(err, "Deployment Deleted")

}

func testCRDApplication(f *framework.Framework) {
	dpName := "dp-" + rand.String(20)
	clusterID := f.ClusterID
	rpNum := int32(rand.Intn(10))
	// k8sctl := f.ClientSet.K8S

	a, err := f.AdminAPIClient.App()
	expect.NoError(err, "App Client Build Failed")
	// client := a.V20201010()

	deployment := app.NewDeployment(dpName, namespace, rpNum, func(deployment *types.Deployment) {})
	values := CustomValues{
		Deployments: []types.Deployment{*deployment},
	}
	valuesYaml, _ := yaml.Marshal(values)
	application := &types.HelmApp{
		ObjectMeta: cpsmetav1.ObjectMeta{
			Alias:       "test-app",
			Namespace:   namespace,
			ClusterName: clusterID,
		},
		Spec: types.HelmAppSpec{
			IsCustomChart: true,
			Values:        string(valuesYaml),
			Network:       "",
		},
		Status: types.HelmAppStatus{},
	}

	_, err = app.CreateApplication(a, application, clusterID, namespace, dpName)
	expect.NoError(err, "Create Deployment Failed")

}

func testListDeployment(f *framework.Framework) {
	clusterID := f.ClusterID
	var dpName [20]string
	num := rand.Intn(20)
	a, err := f.AdminAPIClient.App()
	expect.NoError(err, "App Client Build Failed")
	client := a.V20201010()

	for i := 0; i < num; i++ {
		dpName[i] = "dp-" + rand.String(20)
		deployment := app.NewDeployment(dpName[i], namespace, 1, func(deployment *types.Deployment) {})
		_, err = app.CreateDeployment(a, deployment, clusterID, namespace, dpName[i])
		expect.NoError(err, "Create Deployment Failed")
	}

	res, err := client.ListDeployments(context.TODO(), app.NewListOption(clusterID, namespace), app.NewPageNation())
	expect.NoError(err, "List Deployment Failed")
	// 目前分区还是手动定制，因此list只有第一次才正确
	expect.Equal(res.Total, num, "List Num Right")

	//删除所有创建的 Deployment
	for i := 0; i < num; i++ {
		err = client.DeleteDeployment(context.TODO(), app.NewClusterOption(clusterID, namespace, dpName[i]))
		expect.NoError(err, "Del Secret Failed")
	}

}
