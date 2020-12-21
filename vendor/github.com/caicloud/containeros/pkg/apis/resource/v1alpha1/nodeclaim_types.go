package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// NodeClaimKind is the name of the NodeClaim resource kind.
	NodeClaimKind = "NodeClaim"

	// NodeClaimName is the name of the NodeClaim resource (plural).
	NodeClaimName = "nodeclaims"

	// NodeClaimKindKey is used as the key when mapping to the NodeClaim resource.
	NodeClaimKindKey = "nodeclaim"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="CloudProvider",type="string",JSONPath=".spec.provider",description="The infrastructure provider name of Cluster"
// +kubebuilder:printcolumn:name="ProviderID",type="string",JSONPath=".spec.providerID",description="The Cloud Provider ID of Machine"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// NodeClaim is the Schema for the nodeclaims API
type NodeClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of NodeClaim
	Spec NodeClaimSpec `json:"spec"`

	// Status defines the observed state of NodeClaim
	// +optional
	Status NodeClaimStatus `json:"status,omitempty"`
}

// NodeClaimSpec defines the desired state of NodeClaim
type NodeClaimSpec struct {
	// Provider defines the CloudProvider name of infrastructure provider.
	// +optional
	Provider CloudProvider `json:"provider,omitempty"`

	// ProviderID is the identification ID of the machine provided by the provider.
	// This field must match the provider ID as seen on the node object corresponding to this machine.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// Auth defines the SSH information of a machine.
	Auth MachineAuth `json:"auth"`
}

// NodeClaimStatus defines the observed state of NodeClaim
type NodeClaimStatus struct {
	// NodeRef will point to the corresponding Node if it exists.
	// NOTE: the installer should set this field when node is ready
	// +optional
	NodeRef *corev1.ObjectReference `json:"nodeRef,omitempty"`

	// Conditions represents the observations of nodeclaim's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

const (
	// NodeClaimConditionReady defines the nodeclaim result
	// NOTE: the installer should set the condition to True/False if progress done
	NodeClaimConditionReady = "Ready"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeClaimList is a list of NodeClaim resources
type NodeClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of NodeClaims.
	Items []NodeClaim `json:"items"`
}
