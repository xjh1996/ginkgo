package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineHealthCheckKind is the name of the MachineHealthCheck resource kind.
	MachineHealthCheckKind = "MachineHealthCheck"

	// MachineHealthCheckName is the name of the MachineHealthCheck resource (plural).
	MachineHealthCheckName = "machinehealthchecks"

	// MachineHealthCheckKindKey is used as the key when mapping to the MachineHealthCheck resource.
	MachineHealthCheckKindKey = "machinehealthcheck"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,shortName="mhc"
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.clusterName",description="The cluster id of MachineHealthCheck"
// +kubebuilder:printcolumn:name="ExpectedMachines",type="integer",JSONPath=".status.expectedMachines",description="Total number of machines counted by this machine health check"
// +kubebuilder:printcolumn:name="CurrentHealthy",type="integer",JSONPath=".status.currentHealthy",description="Total number of healthy machines counted by this machine health check"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// MachineHealthCheck is the Schema for the machinehealthchecks API
type MachineHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec of machine health check policy
	Spec MachineHealthCheckSpec `json:"spec"`

	// Status is the observed status of MachineHealthCheck resource
	// +optional
	Status MachineHealthCheckStatus `json:"status,omitempty"`
}

// MachineHealthCheckSpec defines the desired state of MachineHealthCheck
type MachineHealthCheckSpec struct {
	// ClusterName is the name of the Cluster this object belongs to.
	// When this field is empty, it will be applyed on all clusters.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// Label selector to match machines whose health will be exercised
	// +optional
	Selector metav1.LabelSelector `json:"selector,omitempty"`

	// UnhealthyConditions contains a list of the conditions that determine
	// whether a machine is considered unhealthy.  The conditions are combined in a
	// logical OR, i.e. if any of the conditions is met, the machine is unhealthy.
	//
	// +kubebuilder:validation:MinItems=1
	UnhealthyConditions []UnhealthyCondition `json:"unhealthyConditions"`
}

// UnhealthyCondition represents a Machine condition type and value with a timeout
// specified as a duration.  When the named condition has been in the given
// status for at least the timeout value, a machine is considered unhealthy.
type UnhealthyCondition struct {
	// Type is the condition type value
	Type string `json:"type"`

	// Status is the condition status value
	Status corev1.ConditionStatus `json:"status"`

	// Timeout is the duration when the condition treated as unhealthy
	Timeout metav1.Duration `json:"timeout"`
}

// MachineHealthCheckStatus defines the observed state of MachineHealthCheck
type MachineHealthCheckStatus struct {
	// ExpectedMachines is the total number of machines counted by this machine health check
	// +kubebuilder:validation:Minimum=0
	ExpectedMachines int32 `json:"expectedMachines,omitempty"`

	// CurrentHealthy is the total number of healthy machines counted by this machine health check
	// +kubebuilder:validation:Minimum=0
	CurrentHealthy int32 `json:"currentHealthy,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineHealthCheckList is a list of MachineHealthCheck resources
type MachineHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of MachineHealthChecks.
	Items []MachineHealthCheck `json:"items"`
}
