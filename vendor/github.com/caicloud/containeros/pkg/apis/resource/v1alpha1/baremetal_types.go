package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// BaremetalProvider is the name of controller
	BaremetalProvider = "BaremetalProvider"
)

const (
	// BaremetalMachineKind is the name of the BaremetalMachine resource kind.
	BaremetalMachineKind = "BaremetalMachine"

	// BaremetalMachineName is the name of the BaremetalMachine resource (plural).
	BaremetalMachineName = "baremetalmachines"

	// BaremetalMachineKindKey is used as the key when mapping to the BaremetalMachine resource.
	BaremetalMachineKindKey = "barematalmachine"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="ProviderID",type="string",JSONPath=".spec.providerID",description="The Cloud Provider ID of Machine"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready",description="Is Ready"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// BaremetalMachine is the Schema for the barematalmachines API
// NOTE: this type is used only in baremetal-provider
type BaremetalMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of BaremetalMachine
	Spec BaremetalMachineSpec `json:"spec"`

	// Status defines the observed state of BaremetalMachine
	// +optional
	Status BaremetalMachineStatus `json:"status,omitempty"`
}

// BaremetalMachineSpec defines the desired state of BaremetalMachine
type BaremetalMachineSpec struct {
	// ProviderID is the identification ID of the machine provided by the provider.
	// This field must match the provider ID as seen on the node object corresponding to this machine.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`
}

// BaremetalMachineStatus defines the observed state of BaremetalMachine
type BaremetalMachineStatus struct {
	// Conditions represents the observations of baremetalmachine's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// Ready defines the machine is ready
	// +optional
	Ready bool `json:"ready"`
}

const (
	// BaremetalMachineConditionReady defines the baremetal machine result
	BaremetalMachineConditionReady = "Ready"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaremetalMachineList is a list of BaremetalMachine resources
type BaremetalMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of BaremetalMachines.
	Items []BaremetalMachine `json:"items"`
}

const (
	// BaremetalClusterKind is the name of the BaremetalCluster resource kind.
	BaremetalClusterKind = "BaremetalCluster"

	// BaremetalClusterName is the name of the BaremetalCluster resource (plural).
	BaremetalClusterName = "baremetalclusters"

	// BaremetalClusterKindKey is used as the key when mapping to the BaremetalCluster resource.
	BaremetalClusterKindKey = "baremetalcluster"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready",description="Is Ready"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// BaremetalCluster is the Schema for the baremetalclusters API
// NOTE: this type is used only in baremetal-provider
type BaremetalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of BaremetalCluster
	Spec BaremetalClusterSpec `json:"spec"`

	// Status defines the observed state of BaremetalCluster
	// +optional
	Status BaremetalClusterStatus `json:"status,omitempty"`
}

// BaremetalClusterSpec defines the desired state of BaremetalCluster
type BaremetalClusterSpec struct {
	// Provider is the name of infrastructure provider.
	Provider CloudProvider `json:"provider"`
}

// BaremetalClusterStatus defines the observed state of BaremetalCluster
type BaremetalClusterStatus struct {
	// Conditions represents the observations of baremetalcluster's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// Ready defines the cluster is ready
	// +optional
	Ready bool `json:"ready,omitempty"`
}

const (
	// BaremetalClusterConditionReady defines the baremetal cluster result
	BaremetalClusterConditionReady = "Ready"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaremetalClusterList is a list of BaremetalCluster resources
type BaremetalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of BaremetalClusters.
	Items []BaremetalCluster `json:"items"`
}

const (
	// CloudProviderBaremetal is the cloud provider name for baremetal machine and cluster
	CloudProviderBaremetal CloudProvider = "containeros.dev/baremetal"
)
