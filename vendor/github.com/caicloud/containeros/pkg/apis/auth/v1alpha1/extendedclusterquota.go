/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ExtendedClusterQuotaKind is the name of the ClusterQuota resource kind.
	ExtendedClusterQuotaKind = "ExtendedClusterQuota"

	// ExtendedClusterQuotaName is the name of the ClusterQuota resource (plural).
	ExtendedClusterQuotaName = "extendedclusterquotas"

	// ExtendedClusterQuotaKindKey is used as the key when mapping to the ClusterQuota resource.
	ExtendedClusterQuotaKindKey = "extendedclusterquota"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="RawResourceName",type="string",JSONPath=".spec.rawResourceName",description="raw resource name"
// +kubebuilder:subresource:status

// ExtendedClusterQuota extended cluster quota
type ExtendedClusterQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired of ExtendedClusterQuota
	Spec ExtendedClusterQuotaSpec `json:"spec"`

	// Status defines the observed state of ExtendedClusterQuota
	// +optional
	Status ExtendedClusterQuotaStatus `json:"status,omitempty"`
}

// ExtendedClusterQuotaSpec defines spec of ExtendedClusterQuota
type ExtendedClusterQuotaSpec struct {
	RawResourceName string `json:"rawResourceName,omitempty"`
}

// ExtendedClusterQuotaStatus defines status of ExtendedClusterQuota
type ExtendedClusterQuotaStatus struct {
	// Logical cluster logic extended resource
	// +optional
	Logical ExtendedLogical `json:"logical,omitempty"`
	// Physical cluster physical extended resource
	// +optional
	Physical ExtendedPhysical `json:"physical,omitempty"`
}

// ExtendedPhysical defines extended Physical
type ExtendedPhysical struct {
	// Capacity cluster total physical resource
	// +optional
	Capacity corev1.ResourceList `json:"capacity,omitempty"`
	// Allocatable cluster allocatable physical resource
	// +optional
	Allocatable corev1.ResourceList `json:"allocatable,omitempty"`
	// Unavailable cluster can't allocatable physical resource
	// +optional
	Unavailable corev1.ResourceList `json:"unavailable,omitempty"`
	// Unschedulable cluster can't schedulabel physical resource
	// +optional
	Unschedulable corev1.ResourceList `json:"unschedulable,omitempty"`
}

// ExtendedLogical defines extended Logical
type ExtendedLogical struct {
	Total corev1.ResourceList `json:"total,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExtendedClusterQuotaList contains a list of ExtendedClusterQuota
type ExtendedClusterQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of ExtendedClusterQuota.
	Items []ExtendedClusterQuota `json:"items"`
}
