package app

import (
	"context"

	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"

	types "github.com/caicloud/app/pkg/server/client/v20201010"
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
	ginkgo.Context("更新", func() {
		ginkgo.It("服务类型，协议，端口，会话保持", func() {
			testUpdateService(f)
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
	serviceData := app.NewService(ServiceName, namespace)
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

func testUpdateService(f *framework.Framework) {
	a, err := f.APIClient.App()
	expect.NoError(err, "App Client Build Failed")
	client := a.V20201010()

	serviceName := "svc-" + rand.String(20)
	clusterID := f.ClusterID

	service := app.FakeService(serviceName, namespace, func(service *types.Service) {
		// todo 数据和代码分离, 表单驱动
		service.Spec.Ports = []types.Port{
			{
				Name:     "udp-53",
				Protocol: "UDP",
				Port:     53,
			},
			{
				Name:     "tcp-8080",
				Protocol: "TCP",
				Port:     8080,
			},
		}
	})

	res, err := client.CreateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.NoError(err, "Failed to create service")
	for i, port := range service.Spec.Ports {
		expect.Equal(res.Spec.Ports[i].Protocol, port.Protocol)
		expect.Equal(res.Spec.Ports[i].Port, port.Port)
	}

	// ClusterIP 在 Update 时必须要填，且不能修改
	service.Spec.ClusterIP = res.Spec.ClusterIP

	// 修改协议和端口
	service.Spec.Ports[1].Protocol = "UDP"
	service.Spec.Ports[0].Port = int32(rand.IntnRange(1, 65535))
	resUpdate01, err := client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.NoError(err, "Failed to update service")
	for i, port := range service.Spec.Ports {
		expect.Equal(resUpdate01.Spec.Ports[i].Protocol, port.Protocol, "Failed to update service Protocol:", port.Protocol)
		expect.Equal(resUpdate01.Spec.Ports[i].Port, port.Port, "Failed to update service Port:", port.Port)
	}
	// 非法协议
	service.Spec.Ports[1].Protocol = "HTTP"
	_, err = client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.Error(err, "Unexpect update successfully, invalid Protocol:", service.Spec.Ports[1].Protocol)
	service.Spec.Ports[1].Protocol = "UDP"

	// 非法端口
	service.Spec.Ports[1].Port = 65538
	_, err = client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.Error(err, "Unexpect update successfully, invalid Port:", service.Spec.Ports[1].Port)
	service.Spec.Ports[1].Port = 8080

	// 修改 service type ClusterIP => NodePort  NodePort 端口不能是被使用的
	service.Spec.Type = "NodePort"
	service.Spec.Ports[0].NodePort = int32(rand.IntnRange(30000, 32767))
	service.Spec.Ports[1].NodePort = int32(rand.IntnRange(30000, 32767))
	resUpdate02, err := client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.NoError(err, "Failed to update service type:", service.Spec)
	expect.Equal(resUpdate02.Spec.Type, service.Spec.Type)
	for i, port := range service.Spec.Ports {
		expect.Equal(resUpdate02.Spec.Ports[i].NodePort, port.NodePort)
	}

	// 非法 NodePort 端口
	service.Spec.Ports[0].NodePort = int32(rand.IntnRange(0, 29999))
	_, err = client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.Error(err, "Unexpect update successfully, invalid NodePort:", service.Spec.Ports[0].NodePort)
	service.Spec.Ports[0].NodePort = int32(rand.IntnRange(32768, 65536))
	_, err = client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.Error(err, "Unexpect update successfully, invalid NodePort:", service.Spec.Ports[0].NodePort)

	service.Spec.Ports[0].NodePort = int32(rand.IntnRange(30000, 32767))
	service.Spec.Ports[1].NodePort = int32(rand.IntnRange(30000, 32767))

	// 修改会话保持
	timeout := int32(10086)
	service.Spec.SessionAffinity = &types.SessionAffinity{
		TimeoutSeconds: &timeout,
	}
	resUpdate03, err := client.UpdateService(context.TODO(), app.NewServiceGetOptions(clusterID, namespace, serviceName), service)
	expect.NoError(err, "Failed to update service, SessionAffinity:", service.Spec.SessionAffinity)
	expect.Equal(resUpdate03.Spec.SessionAffinity, service.Spec.SessionAffinity)

	// 删除
	serviceDeleteOption := app.NewServiceDeleteOptions(clusterID, namespace, serviceName)
	err = client.DeleteService(context.TODO(), serviceDeleteOption)
	expect.NoError(err, "Failed to delete service")
}
