/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

// FindClusterStatus finds the clusterName in ClusterStatus.
func FindClusterStatus(clusterStatus []ClusterStatus, clusterName string) *ClusterStatus {
	for idx := range clusterStatus {
		if clusterStatus[idx].Cluster == clusterName {
			return &clusterStatus[idx]
		}
	}
	return nil
}

// SetClusterStatus sets the corresponding status in status to newStatus.
func SetClusterStatus(clusterStatus *[]ClusterStatus, newStatus ClusterStatus) {
	if clusterStatus == nil {
		return
	}
	if fs := FindClusterStatus(*clusterStatus, newStatus.Cluster); fs != nil {
		newStatus.DeepCopyInto(fs)
		return
	}

	*clusterStatus = append(*clusterStatus, newStatus)
}

// RemoveClusterStatus removes the corresponding clusterName from ClusterStatus.
func RemoveClusterStatus(clusterStatus *[]ClusterStatus, clusterName string) {
	if clusterStatus == nil {
		return
	}
	newClusterStatus := make([]ClusterStatus, 0, len(*clusterStatus)-1)
	for _, status := range *clusterStatus {
		if status.Cluster != clusterName {
			newClusterStatus = append(newClusterStatus, status)
		}
	}

	*clusterStatus = newClusterStatus
}
