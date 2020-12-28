package app

import (
	"context"

	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"

	"github.com/caicloud/nubela/expect"

	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/rand"
)

var _ = SIGDescribe("服务", func() {
	f := framework.NewDefaultFramework("Service-basic")
	ginkgo.Context("基础部署", func() {
		ginkgo.It("创建+查询+删除", func() {
			testCRDService(f)
		})
	})
})

func testCRDService(f *framework.Framework) {

	// 随机生成配置名称
	ServiceName := "svc" + rand.String(20)
	clusterID := f.ClusterID

	a, err := f.APIClient.App()
	expect.NoError(err, "App Client Build Failed")
	service := a.V20201010()

	//新建Service 传入名称和NameSpace
	serviceData := app.NewService(ServiceName, "default")
	serviceCreateOption := app.NewServiceGetOptions(clusterID, namespace, ServiceName)
	_, err = service.CreateService(context.TODO(), serviceCreateOption, serviceData)
	expect.NoError(err, "Create Service Failed")

	serviceGetOption := app.NewServiceGetOptions(clusterID, namespace, ServiceName)
	serviceSpec := app.NewServiceSpec("ClusterIP", "TCP", 80, 0)
	serviceData, err = service.GetService(context.TODO(), serviceGetOption)
	expect.NoError(err, "Get NewService Failed")

	expect.Equal(serviceData.Spec.Ports, serviceSpec.Ports, "Information is not applied to the service")

	//删除Service
	serviceDeleteOption := app.NewServiceDeleteOptions(clusterID, namespace, ServiceName)
	err = service.DeleteService(context.TODO(), serviceDeleteOption)
	expect.NoError(err, "Del Service Failed")

	//验证删除成功
	_, err = service.GetService(context.TODO(), serviceGetOption)
	expect.Error(err, "Del Service Failed")

}
