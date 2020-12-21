/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// CargoKind is the name of the Cargo resource kind.
	CargoKind = "Cargo"
	// CargoName is the name of the Cargo resource (plural).
	CargoName = "cargos"
	// CargoKindKey is used as the key when mapping to the Cargo resource.
	CargoKindKey = "cargo"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="domain",type=string,JSONPath=`.spec.domain`
// +kubebuilder:subresource:status

// Cargo describes an instance of Cargo registry.
type Cargo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the specification of a cargo registry
	Spec CargoSpec `json:"spec"`

	// Status of cargo registry
	// +optional
	Status CargoStatus `json:"status,omitempty"`
}

// CargoSpec describes specification of a cargo registry.
type CargoSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:MinLength=1
	// Domain of the cargo registry, for example: test.cargo.io
	Domain string `json:"domain"`

	// Host of the cargo registry
	Host string `json:"host"`
	// AccountRef is the reference of a cargo account which used to manage the cargo registry
	AccountRef string `json:"accountRef"`
}

// CargoStatus describes status of a cargo registry
type CargoStatus struct {
	// +kubebuilder:default:=true
	// Healthy indicates whether the cargo registry can work normally
	Healthy bool `json:"healthy"`
	// LastUpdateTime records the time of last updating the object
	// +optional
	LastUpdateTime metav1.Time `json:"LastUpdateTime"`
	// ComponentsStatus of cargo registry components
	// +optional
	ComponentsStatus []ComponentStatus `json:"ComponentsStatus,omitempty"`
}

// ComponentStatus describe status of a cargo registry component
type ComponentStatus struct {
	// Name of the component
	Name string `json:"name"`

	// +kubebuilder:validation:Enum:=healthy;unhealthy;unknown
	// +kubebuilder:default:=unknown
	// Status of the component
	Status string `json:"status"`
	// Error of the component
	// +optional
	Error string `json:"error,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CargoList describes an array of Cargo instances.
type CargoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cargo `json:"items"`
}
