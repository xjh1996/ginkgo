package app

import (
	"context"
	"time"

	appClient "github.com/caicloud/app/pkg/server/client"
	"k8s.io/apimachinery/pkg/util/wait"

	v1 "github.com/caicloud/api/meta/v1"
	types "github.com/caicloud/app/pkg/server/client/v20201010"

	"k8s.io/apimachinery/pkg/util/uuid"
)

// NewConfigMap returns a configmap.
func NewConfigMap(name, namespace, key, value string) *types.ConfigMap {
	return &types.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "v1",
			APIVersion:        "ConfigMap",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
		},
		Type: "KV",
		Data: []types.ConfigMapData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

func NewUpdateConfigMap(name, namespace, key, value string) *types.ConfigMap {
	return &types.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "v1",
			APIVersion:        "ConfigMap",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
		},
		Type: "KV",
		Data: []types.ConfigMapData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

func CreateCfmAndWait(appAPI appClient.Interface, ConfigName, namespace, key, value, clusterID string) (*types.ConfigMap, error) {

	configmapData := NewConfigMap(ConfigName, namespace, key, value)
	configmapGetOption := NewClusterOption(clusterID, namespace, ConfigName)
	_, err := appAPI.V20201010().CreateConfigMap(context.TODO(), configmapGetOption, configmapData)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		configmapData, err = appAPI.V20201010().GetConfigMap(context.TODO(), configmapGetOption)
		if err != nil {
			return false, err
		}
		if configmapData != nil {
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return configmapData, nil
}

func GetCfm(appAPI appClient.Interface, ConfigName, namespace, clusterID string) (*types.ConfigMap, error) {
	clusterOption := NewClusterOption(clusterID, namespace, ConfigName)
	return appAPI.V20201010().GetConfigMap(context.TODO(), clusterOption)
}

func ListCfm(appAPI appClient.Interface, namespace, clusterID string) (*types.ConfigMapList, error) {
	clusterOption := NewListOption(clusterID, namespace)
	return appAPI.V20201010().ListConfigMaps(context.TODO(), clusterOption, NewPageNation())
}

func DeleteCfm(appAPI appClient.Interface, ConfigName, namespace, clusterID string) error {
	clusterOption := NewClusterOption(clusterID, namespace, ConfigName)
	return appAPI.V20201010().DeleteConfigMap(context.TODO(), clusterOption)
}

func UpdateCfmAndWait(appAPI appClient.Interface, ConfigName, namespace, key, value, clusterID string) (*types.ConfigMap, error) {

	configmapData := NewConfigMap(ConfigName, namespace, key, value)
	clusterOption := NewClusterOption(clusterID, namespace, ConfigName)
	_, err := appAPI.V20201010().UpdateConfigMap(context.TODO(), clusterOption, configmapData)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		configmapData, err = appAPI.V20201010().GetConfigMap(context.TODO(), clusterOption)
		if err != nil {
			return false, err
		}
		if configmapData.Data[0].Key == key {
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return configmapData, nil
}
