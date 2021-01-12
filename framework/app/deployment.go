package app

import (
	"context"

	v1 "github.com/caicloud/api/meta/v1"
	appClient "github.com/caicloud/app/pkg/server/client"
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

func CreateDeployment(appAPI appClient.Interface, deployment *types.Deployment, clusterID, namespace, DpName string) (deployment1 *types.Deployment, err error) {
	cluster := NewClusterOption(clusterID, namespace, DpName)
	deployment1, err = appAPI.V20201010().CreateDeployment(context.TODO(), cluster, deployment)
	return deployment1, err
}

func UpdateDeployment(appAPI appClient.Interface, deployment *types.Deployment, clusterID, namespace, DpName string) (deployment1 *types.Deployment, err error) {
	cluster := NewClusterOption(clusterID, namespace, DpName)
	deployment1, err = appAPI.V20201010().UpdateDeployment(context.TODO(), cluster, deployment)
	return deployment1, err
}

func GetDeployment(appAPI appClient.Interface, deploymentName, namespace, clusterID string) (*types.Deployment, error) {
	clusterOption := NewClusterOption(clusterID, namespace, deploymentName)
	return appAPI.V20201010().GetDeployment(context.TODO(), clusterOption)
}

func ListDeployment(appAPI appClient.Interface, namespace, clusterID string) (*types.DeploymentList, error) {
	clusterOption := NewListOption(clusterID, namespace)
	return appAPI.V20201010().ListDeployments(context.TODO(), clusterOption, NewPageNation())
}

func DeleteDeployment(appAPI appClient.Interface, ServiceName, namespace, clusterID string) error {
	clusterOption := NewClusterOption(clusterID, namespace, ServiceName)
	return appAPI.V20201010().DeleteDeployment(context.TODO(), clusterOption)
}

func CreateApplication(appAPI appClient.Interface, app *types.HelmApp, clusterID, namespace, appName string) (app1 *types.HelmApp, err error) {
	cluster := NewClusterOption(clusterID, namespace, appName)
	app1, err = appAPI.V20201010().CreateHelmApp(context.TODO(), cluster, app)
	return app1, err
}
