package app

import (
	"context"

	v1 "github.com/caicloud/api/meta/v1"
	appClient "github.com/caicloud/app/pkg/server/client"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
)

type Modifier func(service *types.Service)

func NewService(name, namespace string, f Modifier) *types.Service {
	svc := &types.Service{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: types.ServiceSpec{
			Type: "ClusterIP",
			Ports: []types.Port{
				{
					Protocol: "TCP",
					Port:     80,
				},
			},
		},
	}
	f(svc)
	return svc
}

func CreateServiceAndWait(appAPI appClient.Interface, ServiceName, namespace, clusterID string) (*types.Service, error) {
	svc := NewService(ServiceName, namespace, func(service *types.Service) {})
	cluster := NewClusterOption(clusterID, namespace, ServiceName)
	service, err := appAPI.V20201010().CreateService(context.TODO(), cluster, svc)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func GetService(appAPI appClient.Interface, ServiceName, namespace, clusterID string) (*types.Service, error) {
	clusterOption := NewClusterOption(clusterID, namespace, ServiceName)
	return appAPI.V20201010().GetService(context.TODO(), clusterOption)
}

func ListService(appAPI appClient.Interface, namespace, clusterID string) (*types.ServiceList, error) {
	clusterOption := NewListOption(clusterID, namespace)
	return appAPI.V20201010().ListServices(context.TODO(), clusterOption, NewPageNation())
}

func DeleteService(appAPI appClient.Interface, ServiceName, namespace, clusterID string) error {
	clusterOption := NewClusterOption(clusterID, namespace, ServiceName)
	return appAPI.V20201010().DeleteService(context.TODO(), clusterOption)
}

func UpdateServiceAndWait(appAPI appClient.Interface, ServiceName, namespace, clusterID string) (*types.Service, error) {
	res, _ := GetService(appAPI, ServiceName, namespace, clusterID)
	svc := NewService(ServiceName, namespace, func(service *types.Service) {
		service.Spec.ClusterIP = res.Spec.ClusterIP
	})
	clusterOption := NewClusterOption(clusterID, namespace, ServiceName)
	_, err := appAPI.V20201010().UpdateService(context.TODO(), clusterOption, svc)
	if err != nil {
		return nil, err
	}
	return svc, nil
}
