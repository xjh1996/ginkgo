package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterInfoKind is the name of the ClusterInfo resource kind.
	ClusterInfoKind = "ClusterInfo"

	// ClusterInfoName is the name of the ClusterInfo resource (plural).
	ClusterInfoName = "clusterinfoes"

	// ClusterInfoKindKey is used as the key when mapping to the ClusterInfo resource.
	ClusterInfoKindKey = "clusterinfo"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:path=clusterinfoes
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="CloudProvider",type="string",JSONPath=".provider",description="The infrastructure provider name of Cluster"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterInfo is the Schema for the clusterinfos API
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// IsControlCluster represents the cluster role for a Cluster.
	IsControlCluster bool `json:"isControlCluster"`

	// Provider is the name of infrastructure provider.
	// +optional
	Provider CloudProvider `json:"provider,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterInfoList is a list of ClusterInfo resources
type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of ClusterInfos.
	Items []ClusterInfo `json:"items"`
}
