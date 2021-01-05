package resource

import (
	"context"
	"fmt"

	resourceclient "github.com/caicloud/resource/pkg/server/client"
	v20201010 "github.com/caicloud/resource/pkg/server/client/v20201010"
)

func GetRandomCluster(rsClient resourceclient.Interface, status string) (*v20201010.Cluster, error) {
	listClusterReq := &v20201010.ListClusterRequest{
		Status: status,
	}
	cls, err := rsClient.V20201010().ListCluster(context.TODO(), listClusterReq)
	if err != nil {
		return nil, err
	}
	for _, c := range cls.Items {
		return &c, nil
	}
	return nil, fmt.Errorf("no %q cluster found", status)
}
