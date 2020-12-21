package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineSetKind is the name of the MachineSet resource kind.
	MachineSetKind = "MachineSet"

	// MachineSetName is the name of the MachineSet resource (plural).
	MachineSetName = "machinesets"

	// MachineSetKindKey is used as the key when mapping to the MachineSet resource.
	MachineSetKindKey = "machineset"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.clusterName",description="The cluster id of MachineSet"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".status.replicas",description="Total number of non-terminated machines targeted by this machineset"
// +kubebuilder:printcolumn:name="Available",type="integer",JSONPath=".status.availableReplicas",description="Total number of available machines (ready for at least minReadySeconds)"
// +kubebuilder:printcolumn:name="Ready",type="integer",JSONPath=".status.readyReplicas",description="Total number of ready machines targeted by this machineset."
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// MachineSet is the Schema for the machinesets API
type MachineSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the desired state of MachineSet
	Spec MachineSetSpec `json:"spec"`

	// Status is the observed state of MachineSet
	// +optional
	Status MachineSetStatus `json:"status,omitempty"`
}

// MachineSetSpec defines the desired state of MachineSet
type MachineSetSpec struct {
	// ClusterName is the name of the Cluster this object belongs to.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// Replicas is the number of desired replicas.
	// This is a pointer to distinguish between explicit zero and unspecified.
	// Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// MinReadySeconds is the minimum number of seconds for which a newly created machine should be ready.
	// Defaults to 0 (machine will be considered available as soon as it is ready)
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// DeletePolicy defines the policy used to identify machines to delete when downscaling.
	// +optional
	DeletePolicy MachineSetDeletePolicy `json:"deletePolicy,omitempty"`

	// Selector is a label query over machines that should match the replica count.
	// Label keys and values that must match in order to be controlled by this MachineSet.
	// It must match the machine template's labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector metav1.LabelSelector `json:"selector"`

	// Template is the object that describes the machine that will be created if
	// insufficient replicas are detected.
	// Object references to custom resources resources are treated as templates.
	Template MachineTemplateSpec `json:"template"`
}

// MachineSetDeletePolicy defines the policy used to delete machine when downscaling.
type MachineSetDeletePolicy struct {
	// Type defines the policy used to identify nodes to delete when downscaling.
	// Defaults to "Random".  Valid values are "Random, "Newest", "Oldest"
	Type MachineSetDeletePolicyType `json:"type"`
}

// +kubebuilder:validation:Enum=Random;Newest;Oldest

// MachineSetDeletePolicyType defines how priority is assigned to nodes to delete when
// downscaling a MachineSet. Defaults to "Random".
type MachineSetDeletePolicyType string

const (
	// RandomMachineSetDeletePolicy prioritizes both Machines that have the annotation
	// "cluster.x-k8s.io/delete-machine=yes" and Machines that are unhealthy.
	// Finally, it picks Machines at random to delete.
	RandomMachineSetDeletePolicy MachineSetDeletePolicyType = "Random"

	// NewestMachineSetDeletePolicy prioritizes both Machines that have the annotation
	// "cluster.x-k8s.io/delete-machine=yes" and Machines that are unhealthy.
	// It then prioritizes the newest Machines for deletion based on the Machine's CreationTimestamp.
	NewestMachineSetDeletePolicy MachineSetDeletePolicyType = "Newest"

	// OldestMachineSetDeletePolicy prioritizes both Machines that have the annotation
	// "cluster.x-k8s.io/delete-machine=yes" and Machines that are unhealthy.
	// It then prioritizes the oldest Machines for deletion based on the Machine's CreationTimestamp.
	OldestMachineSetDeletePolicy MachineSetDeletePolicyType = "Oldest"
)

// MachineTemplateSpec describes the data needed to create a Machine from a template
type MachineTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata"`

	// Spec of the desired behavior of the machine.
	Spec MachineSpec `json:"spec"`
}

// MachineSetStatus defines the observed state of MachineSet
type MachineSetStatus struct {
	// Replicas is the most recently observed number of replicas.
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// FullyLabeledReplicas is the number of replicas that have labels matching the labels of the machine template of the MachineSet.
	// +optional
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty"`

	// ReadyReplicas is the number of ready replicas for this MachineSet. A machine is considered ready when the object has been created and is "Ready".
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// AvailableReplicas is the number of available replicas (ready for at least minReadySeconds) for this MachineSet.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// Conditions represents the observations of nodeclaim's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineSetList is a list of MachineSet resources
type MachineSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of MachineSets.
	Items []MachineSet `json:"items"`
}
