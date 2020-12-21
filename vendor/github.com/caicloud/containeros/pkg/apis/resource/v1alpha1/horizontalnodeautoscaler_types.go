package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HorizontalNodeAutoscalerKind is the name of the HorizontalNodeAutoscaler resource kind.
	HorizontalNodeAutoscalerKind = "HorizontalNodeAutoscaler"

	// HorizontalNodeAutoscalerName is the name of the HorizontalNodeAutoscaler resource (plural).
	HorizontalNodeAutoscalerName = "horizontalnodeautoscalers"

	// HorizontalNodeAutoscalerKindKey is used as the key when mapping to the HorizontalNodeAutoscaler resource.
	HorizontalNodeAutoscalerKindKey = "horizontalnodeautoscaler"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:shortName="hna"
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="IsEnabled",type="boolean",JSONPath=".spec.isEnabled",description="Enable option"
// +kubebuilder:printcolumn:name="Min",type="integer",JSONPath=".spec.minReplicas",description="The min replicas count"
// +kubebuilder:printcolumn:name="Max",type="integer",JSONPath=".spec.maxReplicas",description="The max replicas count"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".status.replicas",description="The current replicas count"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// HorizontalNodeAutoscaler is the Schema for the horizontalnodeautoscalers API
type HorizontalNodeAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of HorizontalNodeAutoscaler
	Spec HorizontalNodeAutoscalerSpec `json:"spec"`

	// Status defines the observed state of HorizontalNodeAutoscaler
	// +optional
	Status HorizontalNodeAutoscalerStatus `json:"status,omitempty"`
}

// HorizontalNodeAutoscalerSpec defines the desired state of HorizontalNodeAutoscaler
type HorizontalNodeAutoscalerSpec struct {
	// IsEnabled defines the HorizontalNodeAutoscaler enable option
	IsEnabled bool `json:"isEnabled"`

	// MinReplicas is the min replicas count
	// Defaults to 0.
	// +kubebuilder:validation:Minimum=0
	MinReplicas int `json:"minReplicas"`

	// MinReplicas is the max replicas count
	// Defaults to 0.
	// +kubebuilder:validation:Minimum=0
	MaxReplicas int `json:"maxReplicas"`

	// NotifySetting defines the notify information
	// +optional
	NotifySetting NodeAutoscalerNotifySetting `json:"notifySetting,omitempty"`

	// Capacity defines the template node capacity, used for scheduler simulate
	Capacity corev1.ResourceList `json:"capacity"`

	// Tags defines the node labels, used for scheduler simulate
	// +optional
	Tags MachineTags `json:"tags,omitempty"`

	// Taints defines the node taints, used for scheduler simulate
	// +optional
	Taints []corev1.Taint `json:"taints,omitempty"`

	// MachineSetRef is the reference to MachineSet Object offered for scale
	MachineSetRef *corev1.ObjectReference `json:"machineSetRef"`
}

// HorizontalNodeAutoscalerStatus defines the observed state of HorizontalNodeAutoscaler
type HorizontalNodeAutoscalerStatus struct {
	// Conditions represents the observations of machine's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// Replicas is the current replicas
	// Defaults to 0.
	// +kubebuilder:validation:Minimum=0
	Replicas int32 `json:"replicas"`

	// ReadyReplicas is the current replicas witch has ready
	// Defaults to 0.
	// +kubebuilder:validation:Minimum=0
	ReadyReplicas int32 `json:"readyReplicas"`

	// LastScaleUpTime is the timestamp when last scale up
	// +optional
	LastScaleUpTime *metav1.Time `json:"lastScaleUpTime,omitempty"`

	// LastScaleDownTime is the timestamp when last scale down
	// +optional
	LastScaleDownTime *metav1.Time `json:"lastScaleDownTime,omitempty"`

	// Nodes is the list of Node which belongs to this object
	// +optional
	Nodes []HorizontalNodeAutoscalerItem `json:"nodes,omitempty"`
}

// HorizontalNodeAutoscalerItem defines the Node object information
type HorizontalNodeAutoscalerItem struct {
	// Name is the Node name
	Name string `json:"name"`

	// LastBusyTime is the timestamp when last seen the node is busy
	// +optional
	LastBusyTime *metav1.Time `json:"lastBusyTime,omitempty"`

	// CreateTimestamp is the timestamp when the item is created
	CreateTimestamp *metav1.Time `json:"createTimestamp"`

	// DeleteTimestamp is the timestamp when the item is scale down
	// +optional
	DeleteTimestamp *metav1.Time `json:"deleteTimestamp,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorizontalNodeAutoscalerList is a list of HorizontalNodeAutoscaler resources
type HorizontalNodeAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of HorizontalNodeAutoscalers.
	Items []HorizontalNodeAutoscaler `json:"items"`
}
