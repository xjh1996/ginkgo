package app

import (
	"context"

	appClient "github.com/caicloud/app/pkg/server/client"
	"k8s.io/apimachinery/pkg/util/wait"

	v1 "github.com/caicloud/api/meta/v1"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
)

func NewSecret(name, namespace, key, value string) *types.Secret {
	return &types.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: []types.SecretData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

func NewUpdateSecret(name, namespace, key, value string) *types.Secret {
	return &types.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: []types.SecretData{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

func CreateSecretAndWait(appAPI appClient.Interface, SecretName, namespace, key, value, clusterID string) (*types.Secret, error) {
	secret := NewSecret(SecretName, namespace, key, value)
	cluster := NewClusterOption(clusterID, namespace, SecretName)
	_, err := appAPI.V20201010().CreateSecret(context.TODO(), cluster, secret)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		res, err := appAPI.V20201010().GetSecret(context.TODO(), cluster)
		if err != nil {
			return false, err
		}
		if res != nil {
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func GetSecret(appAPI appClient.Interface, SecretName, namespace, clusterID string) (*types.Secret, error) {
	clusterOption := NewClusterOption(clusterID, namespace, SecretName)
	return appAPI.V20201010().GetSecret(context.TODO(), clusterOption)
}

func ListSecret(appAPI appClient.Interface, namespace, clusterID string) (*types.SecretList, error) {
	clusterOption := NewListOption(clusterID, namespace)
	return appAPI.V20201010().ListSecrets(context.TODO(), clusterOption, NewPageNation())
}

func DeleteSecret(appAPI appClient.Interface, SecretName, namespace, clusterID string) error {
	clusterOption := NewClusterOption(clusterID, namespace, SecretName)
	return appAPI.V20201010().DeleteSecret(context.TODO(), clusterOption)
}

func UpdateSecretAndWait(appAPI appClient.Interface, SecretName, namespace, key, value, clusterID string) (*types.Secret, error) {
	secret := NewSecret(SecretName, namespace, key, value)
	clusterOption := NewClusterOption(clusterID, namespace, SecretName)
	_, err := appAPI.V20201010().UpdateSecret(context.TODO(), clusterOption, secret)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		secret, err = appAPI.V20201010().GetSecret(context.TODO(), clusterOption)
		if err != nil {
			return false, err
		}
		if secret.Data[0].Key == key {
			return true, nil
		} else {
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return secret, nil
}
