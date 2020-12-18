package app

import (
	"time"

	v1 "github.com/caicloud/api/meta/v1"
	types "github.com/caicloud/app/pkg/server/client/v20201010"

	"k8s.io/apimachinery/pkg/util/uuid"
)

func NewSecret(name, namespace, key, value string) *types.Secret {
	return &types.Secret{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "v1",
			APIVersion:        "Secret",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
		},
		Data: []types.SecretData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

func NewSecretGetOptions(clustername, namespace, name string) types.SecretGetOption {
	return types.SecretGetOption{
		Cluster: types.Cluster{
			ClusterName: clustername,
			Namespace:   namespace,
			Name:        name,
		},
	}
}

func NewSecretDeleteOptions(clustername, namespace, name string) types.SecretDeleteOption {
	return types.SecretDeleteOption{
		Cluster: types.Cluster{
			ClusterName: clustername,
			Namespace:   namespace,
			Name:        name,
		},
	}
}

func NewUpdateSecret(name, namespace, key, value string) *types.Secret {
	return &types.Secret{
		ObjectMeta: v1.ObjectMeta{
			Kind:              "v1",
			APIVersion:        "ConfigMap",
			UID:               string(uuid.NewUUID()),
			CreationTimestamp: time.Now(),
			Name:              name,
			Namespace:         namespace,
		},
		Data: []types.SecretData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}
