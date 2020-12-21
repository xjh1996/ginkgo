/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// TenantKind is the name of the tenant resource kind.
	TenantKind = "Tenant"

	// TenantName is the name of the tenant resource (plural).
	TenantName = "tenants"

	// TenantKindKey is used as the key when mapping to the tenant resource.
	TenantKindKey = "tenant"

	// SystemTenantName defines name of system tenant
	SystemTenantName = "system-tenant"
)

// TenantPhase is a string representation of a Tenant Phase.
type TenantPhase string

const (
	// TenantPhaseActive means the tenant is available for use in the system
	TenantPhaseActive TenantPhase = "Active"

	// TenantPhaseTerminating means the tenant is undergoing graceful termination
	TenantPhaseTerminating TenantPhase = "Terminating"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Current phase of Tenant actuation"
// +kubebuilder:printcolumn:name="NamespaceCount",type="integer",JSONPath=".status.namespaceCount",description="Namespace count hold by tenant"
// +kubebuilder:subresource:status

// Tenant is the Schema for the tenants API.
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired of Tenant
	Spec TenantSpec `json:"spec"`

	// Status defines the observed state of Tenant
	// +optional
	Status TenantStatus `json:"status,omitempty"`
}

// TenantSpec defines the desired state of Tenant.
type TenantSpec struct {
	// Quota is the resources allocated.
	// If the tenant is system-tenant, the value is equal to clusterquota, because system-tenant can use all quota in cluster.
	// If the tenant is non system-tenant, the value is user assigned.
	Quota corev1.ResourceList `json:"quota"`
}

// TenantStatus defines the observed state of Tenant.
type TenantStatus struct {
	// Phase is a string representation of a Tenant Phase.
	// +kubebuilder:validation:Enum=Active;Terminating
	// +optional
	Phase TenantPhase `json:"phase,omitempty"`

	// Conditions represents the observations of tenant's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// Allocated is the resources witch have been allocated to namespace.
	// +optional
	Allocated corev1.ResourceList `json:"allocated,omitempty"`

	// Used is the resources witch have been used in namespace.
	// +optional
	Used corev1.ResourceList `json:"used,omitempty"`

	// NamespaceCount is the count of namespace witch is hold by the tenant.
	// If the tenant is system-tenant, the count contains the system namespace.
	// +optional
	NamespaceCount int `json:"namespaceCount"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TenantList contains a list of Tenant
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of Tenants.
	Items []Tenant `json:"items"`
}
