package app

import (
	"time"

	v1 "github.com/caicloud/api/meta/v1"
	types "github.com/caicloud/app/pkg/server/client/v20201010"

	"k8s.io/apimachinery/pkg/util/uuid"
)

type Modifier func(service *types.Service)

func FakeService(name, namespace string, f Modifier) *types.Service {
	svc := &types.Service{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "Service",
			APIVersion:        "v1",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
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

func NewService(name, namespace string) *types.Service {
	return &types.Service{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "Service",
			APIVersion:        "v1",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
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
}

func NewServiceSpec(specType, protocol string, port, nodePort int32) types.ServiceSpec {
	return types.ServiceSpec{
		Type: specType,
		Ports: []types.Port{
			{
				Protocol: protocol,
				Port:     port,
				NodePort: nodePort,
			},
		},
	}

}
