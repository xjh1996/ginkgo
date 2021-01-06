package app

import (
	v1 "github.com/caicloud/api/meta/v1"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
)

type DpModifier func(deployment *types.Deployment)

// NewDeployment returns a deployment.
func NewDeployment(name, namespace string, rpNum int32, f DpModifier) *types.Deployment {
	dp := &types.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Kind:       "v1",
			APIVersion: "Deployment",
			Name:       name,
			Namespace:  namespace,
		},
		Spec: types.DeploymentSpec{
			Replicas: &rpNum,
			Template: types.TemplateSpec{
				Spec: types.PodSpec{
					Containers: []types.Container{
						{
							Name:  "c0",
							Image: "cargo.dev.caicloud.xyz/qatest/testtools:nginx",
							Resources: types.ResourceRequirements{
								Limits: []types.KV{
									{
										Key:   "cpu",
										Value: "20m",
									},
									{
										Key:   "memory",
										Value: "20Mi",
									},
								},
								Requests: []types.KV{
									{
										Key:   "cpu",
										Value: "10m",
									},
									{
										Key:   "memory",
										Value: "10Mi",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	f(dp)
	return dp
}

func NewDpGetOption(clusterName, namespace, name string) types.DeploymentGetOption {
	return types.DeploymentGetOption{
		Cluster: types.Cluster{
			ClusterName: clusterName,
			Namespace:   NS4Auth(clusterName, namespace),
			Name:        name,
		},
	}
}

func NewDpDeleteOption(clusterName, namespace, name string) types.DeploymentDeleteOption {
	return types.DeploymentDeleteOption{
		Cluster: types.Cluster{
			ClusterName: clusterName,
			Namespace:   NS4Auth(clusterName, namespace),
			Name:        name,
		},
	}
}

func NewDpListOption(clusterName, namespace string) types.DeploymentListOption {
	return types.DeploymentListOption{
		Cluster: types.Cluster{
			ClusterName: clusterName,
			Namespace:   NS4Auth(clusterName, namespace),
		},
	}
}
