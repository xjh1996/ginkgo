package resource

import (
	"context"
	"time"

	resourceclient "github.com/caicloud/resource/pkg/server/client"
	v20201010 "github.com/caicloud/resource/pkg/server/client/v20201010"

	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	interval = time.Second * 2
	timeout  = time.Second * 300
)

// Describe resource metadate for a namespace
func CreatePVCAndWait(resourceAPI resourceclient.Interface, nsName, PVCName, storageClass, clusterID string) (*v20201010.PVCObject, error) {
	createPVCReq := &v20201010.CreatePVCRequest{
		Cluster:      clusterID,
		Namespace:    nsName,
		Name:         PVCName,
		StorageClass: storageClass,
		Size:         "1Gi",
	}
	pvc, err := resourceAPI.V20201010().CreatePersistentVolumeClaim(context.TODO(), createPVCReq)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		pvc, err = resourceAPI.V20201010().GetPersistentVolumeClaim(context.TODO(), clusterID, nsName, PVCName)
		if err != nil {
			return false, err
		}
		if pvc.Status.Phase == "Bound" {
			return true, nil
		} else {
			return false, nil
		}
	})
	return pvc, err
}

