package v1alpha1

import (
	"fmt"
	"strings"

	custom_metav1 "github.com/caicloud/api/meta/v1"
)

// MetaNamespaceKey provides a way to generates meta namespace key without returning error
func (pcs *PrometheusClusterSet) MetaNamespaceKey() string {
	return fmt.Sprintf("%s/%s", pcs.Namespace, pcs.Name)
}

// Mutate attempts to mutate the receiver to make it valid and fill its default values.
func (pcs *PrometheusClusterSet) Mutate() *PrometheusClusterSet {
	if pcs == nil {
		return nil
	}
	ret := pcs.DeepCopy()
	if ret.Spec.Prometheus.Metadata == nil {
		ret.Spec.Prometheus.Metadata = &custom_metav1.EmbeddedObjectMeta{}
	}
	metaTmpl := ret.Spec.Prometheus.Metadata
	if len(metaTmpl.Namespace) == 0 {
		metaTmpl.Namespace = pcs.Namespace
	}
	if len(metaTmpl.Name) == 0 {
		metaTmpl.Name = fmt.Sprintf("pcs-%s", pcs.Name)
	}
	return ret
}

// GetClusterStatus fetch the deployment status of the given cluster.
func (pcs *PrometheusClusterSet) GetClusterStatus(cluster string) (*ClusterStatus, bool) {
	if pcs.Status != nil {
		for i := range pcs.Status.Clusters {
			if pcs.Status.Clusters[i].Cluster == cluster {
				return &pcs.Status.Clusters[i], true
			}
		}
	}
	return nil, false
}

// ClusterStatuses is a alias of []ClusterStatus that implements sort.Interface
type ClusterStatuses []ClusterStatus

// Len implements sort.Interface.
func (cs ClusterStatuses) Len() int {
	return len(cs)
}

// Swap implements sort.Interface.
func (cs ClusterStatuses) Swap(i, j int) {
	tmp := cs[i].DeepCopy()
	cs[j].DeepCopyInto(&cs[i])
	tmp.DeepCopyInto(&cs[i])
}

// Less implements sort.Interface.
func (cs ClusterStatuses) Less(i, j int) bool {
	return strings.Compare(cs[i].Cluster, cs[j].Cluster) < 0
}
