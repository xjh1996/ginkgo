package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// NodeAutoscalerConfigurationKind is the name of the NodeAutoscalerConfiguration resource kind.
	NodeAutoscalerConfigurationKind = "NodeAutoscalerConfiguration"

	// NodeAutoscalerConfigurationName is the name of the NodeAutoscalerConfiguration resource (plural).
	NodeAutoscalerConfigurationName = "nodeautoscalerconfigrations"

	// NodeAutoscalerConfigurationKindKey is used as the key when mapping to the NodeAutoscalerConfiguration resource.
	NodeAutoscalerConfigurationKindKey = "nodeautoscalerconfigration"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Algorithm",type="string",JSONPath=".spec.algorithm",description="The algorithm of scale up"
// +kubebuilder:printcolumn:name="IsQuotaUpdateEnabled",type="boolean",JSONPath=".spec.isQuotaUpdateEnabled",description="Auto quota update option"
// +kubebuilder:printcolumn:name="IsScaleDownEnabled",type="boolean",JSONPath=".spec.isScaleDownEnabled",description="Auto scale down option"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// NodeAutoscalerConfiguration is the Schema for the nodeautoscalerconfigrations API
type NodeAutoscalerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of NodeAutoscalerConfiguration
	Spec NodeAutoscalerConfigurationSpec `json:"spec"`

	// Status defines the observed state of NodeAutoscalerConfiguration
	// +optional
	Status NodeAutoscalerConfigurationStatus `json:"status,omitempty"`
}

// NodeAutoscalerConfigurationSpec defines the desired state of NodeAutoscalerConfiguration
type NodeAutoscalerConfigurationSpec struct {
	// Algorithm is the type of scale up
	Algorithm HorizontalNodeAutoscalerAlgorithmType `json:"algorithm"`

	// IsQuotaUpdateEnabled defines auto quota update option
	IsQuotaUpdateEnabled bool `json:"isQuotaUpdateEnabled"`

	// IsScaleDownEnabled defines auto scale down option
	IsScaleDownEnabled bool `json:"isScaleDownEnabled"`

	// CoolDown defines the scale down duration from last scale up action
	CoolDown metav1.Duration `json:"coolDown"`

	// IdleTime defines the idle duration when machine scale down
	IdleTime metav1.Duration `json:"idleTime"`

	// IdleThreshold defines the threshold when machine consider as idle
	IdleThreshold int `json:"idleThreshold"`

	// NotifySetting defines the notify information
	// +optional
	NotifySetting NodeAutoscalerNotifySetting `json:"notifySetting,omitempty"`
}

// +kubebuilder:validation:Enum=Random;MostPods;LeastWaste

// HorizontalNodeAutoscalerAlgorithmType defines the algorithm type when scale up
type HorizontalNodeAutoscalerAlgorithmType string

const (
	// HorizontalNodeAutoscalerUpAlgorithmRandom pickup hna at random to scale up.
	// ex:
	// 1. There is three hna (hna1、hna2 and hna3) enabled.
	// 2. There is ten pods Pending
	// 3. hna1 doesn't match any Pending pods
	// 4. hna2 match 1 pods
	// 5. hna3 match 6 pods
	// results:
	// Random algorithm will random pickup one hna from (hna2, hna3), not from hna1, because there must be at
	// one pods will consider as schedulable when scale up
	HorizontalNodeAutoscalerUpAlgorithmRandom HorizontalNodeAutoscalerAlgorithmType = "Random"

	// HorizontalNodeAutoscalerUpAlgorithmMostPods pickup hna witch can scale up most pods.
	// ex:
	// 1. There is three hna (hna1、hna2 and hna3) enabled.
	// 2. There is ten pods Pending
	// 3. hna1 doesn't match any Pending pods
	// 4. hna2 match 5 pods with 2 node scale up
	// 5. hna3 match 6 pods with 1 node scale up
	// results:
	// MostPods algorithm will pickup hna3 because it will make most pods as schedulable when scale up
	HorizontalNodeAutoscalerUpAlgorithmMostPods HorizontalNodeAutoscalerAlgorithmType = "MostPods"

	// HorizontalNodeAutoscalerUpAlgorithmLeastWaste pickup hna witch can waste node resource least.
	// ex:
	// 1. There is three hna (hna1、hna2 and hna3) enabled.
	// 2. There is ten pods Pending
	// 3. hna1 doesn't match any Pending pods
	// 4. hna2 match 1 pods (node: 2c, pods: 2c)
	// 5. hna3 match 6 pods (node: 4c, pods: 1c)
	// results:
	// LeastWaste algorithm will pickup hna2 because it will waste least node resource.
	//   1. hna2 used all node resource without waste
	//   2. hna3 will scale up 2 node, with 2c waste
	HorizontalNodeAutoscalerUpAlgorithmLeastWaste HorizontalNodeAutoscalerAlgorithmType = "LeastWaste"
)

// NodeAutoscalerConfigurationStatus defines the observed state of NodeAutoscalerConfiguration
type NodeAutoscalerConfigurationStatus struct {
	// LastScaleUpTime is the timestamp when last scale up
	// +optional
	LastScaleUpTime *metav1.Time `json:"lastScaleUpTime,omitempty"`

	// LastScaleDownTime is the timestamp when last scale down
	// +optional
	LastScaleDownTime *metav1.Time `json:"lastScaleDownTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeAutoscalerConfigurationList is a list of NodeAutoscalerConfiguration resources
type NodeAutoscalerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of NodeAutoscalerConfigurations.
	Items []NodeAutoscalerConfiguration `json:"items"`
}
