package app

import (
	"time"

	"github.com/caicloud/zeus/framework/auth"

	appClient "github.com/caicloud/app/pkg/server/client"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
	authClient "github.com/caicloud/auth/pkg/server/client"
	"github.com/caicloud/nubela/logger"

	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	interval = time.Second * 2
	timeout  = time.Second * 10
)

type ConfigBaseInfo struct {
	ConfigName string
	Key        string
	Value      string
}

func NewClusterOption(clusterName, namespace, name string) types.Cluster {
	return types.Cluster{
		ClusterName: clusterName,
		Namespace:   NS4Auth(clusterName, namespace),
		Name:        name,
	}
}

func NewPageNation() types.Pagination {
	return types.Pagination{}
}

func NewListOption(clusterName, namespace string) types.Cluster {
	return types.Cluster{
		ClusterName: clusterName,
		Namespace:   NS4Auth(clusterName, namespace),
	}
}

// NS4Auth the namespace for auth rpc style
func NS4Auth(cluster, namespace string) string {
	NS4Auth := "cluster/" + cluster + "/" + namespace
	return NS4Auth
}

func CreateConfigInfo() *ConfigBaseInfo {
	return &ConfigBaseInfo{
		ConfigName: "config" + rand.String(8),
		Key:        "key" + rand.String(5),
		Value:      "value" + rand.String(5),
	}
}

func GetNormalUserAppAPI(authAPI authClient.Interface, baseInfo *auth.BaseInfo, permission, resource []string) appClient.Interface {
	user := auth.PresetOperation(authAPI, baseInfo, permission, resource)
	normalUserAppAPI, err := user.App()
	if err != nil {
		logger.Failf("get normal user failed, %v", err)
	}
	return normalUserAppAPI
}
