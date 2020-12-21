/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterQuotaKind is the name of the ClusterQuota resource kind.
	ClusterQuotaKind = "ClusterQuota"

	// ClusterQuotaName is the name of the ClusterQuota resource (plural).
	ClusterQuotaName = "clusterquotas"

	// ClusterQuotaKindKey is used as the key when mapping to the ClusterQuota resource.
	ClusterQuotaKindKey = "clusterquota"

	// SystemClusterQuotaName defines cluster quota name
	SystemClusterQuotaName = "system"

	// ResourceNvidiaGPU defines nvidia gpu
	ResourceNvidiaGPU = "nvidia.com/gpu"

	// ResourceRequestsNvidiaGPU  defines nvidia gpu requests
	ResourceRequestsNvidiaGPU = "requests.nvidia.com/gpu"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// ClusterQuota is the Schema for the cluster quota API.
type ClusterQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired of ClusterQuota
	Spec ClusterQuotaSpec `json:"spec"`

	// Status defines the observed state of ClusterQuota
	// +optional
	Status ClusterQuotaStatus `json:"status,omitempty"`
}

// ClusterQuotaSpec defines the desired spec of ClusterQuota.
type ClusterQuotaSpec struct {
	// Ratio is the resources scale rate.
	Ratio map[corev1.ResourceName]int64 `json:"ratio"`
}

// ClusterQuotaStatus defines the observed state of Tenant.
type ClusterQuotaStatus struct {
	// Logical cluster logical resource quota
	// +optional
	Logical LogicalQuota `json:"logical,omitempty"`
	// Physical cluster physical resource quota
	// +optional
	Physical PhysicalQuota `json:"physical,omitempty"`
}

// PhysicalQuota defines the physical resource quota of cluster
type PhysicalQuota struct {
	// Capacity cluster total resource quota
	// +optional
	Capacity corev1.ResourceList `json:"capacity,omitempty"`
	// Allocatable application can alloc resource quota
	// +optional
	Allocatable corev1.ResourceList `json:"allocatable,omitempty"`
	// Unavailable not ready or annotation quota.auth.containeros.dev/type =unavailable node resource
	// +optional
	Unavailable corev1.ResourceList `json:"unavailable,omitempty"`
	// Unschedulable annotation quota.auth.containeros.dev/type = unschedulable node resource
	// +optional
	Unschedulable corev1.ResourceList `json:"unschedulable,omitempty"`
}

// LogicalQuota defines the logical resource quota of cluster
type LogicalQuota struct {
	// Total all of logical resource quota
	// +optional
	Total corev1.ResourceList `json:"total,omitempty"`
	// Allocated already alloc resource quota
	// +optional
	Allocated corev1.ResourceList `json:"allocated,omitempty"`
	// SystemUsed system used resource quota
	// +optional
	SystemUsed corev1.ResourceList `json:"systemUsed,omitempty"`
	// Used application used resource quota
	// +optional
	Used corev1.ResourceList `json:"used,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterQuotaList contains a list of ClusterQuota
type ClusterQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of ClusterQuotas.
	Items []ClusterQuota `json:"items"`
}
