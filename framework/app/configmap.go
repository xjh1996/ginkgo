package app

import (
	"time"

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
