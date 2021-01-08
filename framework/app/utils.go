package app

import (
	types "github.com/caicloud/app/pkg/server/client/v20201010"
)

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
