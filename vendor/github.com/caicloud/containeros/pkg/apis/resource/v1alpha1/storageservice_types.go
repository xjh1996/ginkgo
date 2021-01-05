package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// StorageClassKind is the name of the StorageClass resource kind.
	// NOTE: this is only used for display
	StorageClassKind = "StorageClass"

	// StorageClassName is the name of the StorageClass resource (plural).
	StorageClassName = "storageclasses"

	// StorageClassKindKey is used as the key when mapping to the StorageClass resource.
	StorageClassKindKey = "storageclass"
)

const (
	// StorageServiceKind is the name of the StorageService resource kind.
	StorageServiceKind = "StorageService"

	// StorageServiceName is the name of the StorageService resource (plural).
	StorageServiceName = "storageservices"

	// StorageServiceKindKey is used as the key when mapping to the StorageService resource.
	StorageServiceKindKey = "storageservice"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.typeName",description="The storage type"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="The storage phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// StorageService describes the parameters for a class of storage for
// which PersistentVolumes can be dynamically provisioned.
type StorageService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of StorageService
	Spec StorageServiceSpec `json:"spec"`

	// Status defines the observed state of StorageService
	// +optional
	Status StorageServiceStatus `json:"status,omitempty"`
}

// StorageServiceSpec defines the desired state of StorageService
type StorageServiceSpec struct {
	// TypeName indicates the name of the storage type that this service belongs to.
	TypeName string `json:"typeName" protobuf:"bytes,2,opt,name=typeName"`

	// Parameters holds the parameters for the provisioner that should
	// create volumes of this storage class.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`

	// Hard is the set of desired hard limits for each named resource.
	// +optional
	Hard corev1.ResourceList `json:"hard,omitempty"`
}

// StorageServiceStatus defines the observed state of StorageService
type StorageServiceStatus struct {
	// Phase is a string representation of a StorageService Phase.
	// +optional
	Phase StorageServicePhase `json:"phase,omitempty"`

	// StorageMetaData represents the current metadata for each storage backend.
	// NOTE: this is something being hacking, as this field is only used by some spectify storage.
	// +optional
	StorageMetaData StorageMetaData `json:"storageMetaData,omitempty"`

	// Allocated is the amount of resources that have been allocated.
	// +optional
	Allocated corev1.ResourceList `json:"allocated,omitempty"`

	// Conditions represents the observations of baremetalcluster's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:validation:Enum=Active;Terminating

// StorageServicePhase is a string representation of a StorageService Phase.
type StorageServicePhase string

const (
	// StorageServicePhaseActive means the storage service is available for use in the system
	StorageServicePhaseActive StorageServicePhase = "Active"

	// StorageServicePhaseTerminating means the storage service is undergoing graceful termination
	StorageServicePhaseTerminating StorageServicePhase = "Terminating"
)

const (
	// StorageServiceConditionExceedQuota defines the storage service is out of quota
	StorageServiceConditionExceedQuota = "ExceedQuota"
)

// StorageMetaData is the data structure for each storage backend metadata.
type StorageMetaData struct {
	// Ceph describe the ceph storage backend metadata
	// +optional
	Ceph *CephMetaData `json:"ceph,omitempty"`
}

// CephMetaData is the data structure for Ceph metadata.
type CephMetaData struct {
	// Pools is the list of ceph poll attr
	// +optional
	Pools []CephPool `json:"pools,omitempty"`
}

// CephPool is the data structure for single Ceph storage pool.
type CephPool struct {
	// Name is the pool name
	Name string `json:"name"`

	// ReplicaSize is the replica count
	ReplicaSize int `json:"replicaSize"`

	// MaxAvailable capacity, kb
	MaxAvailable int `json:"maxAvailable"`

	// PercentUsed, 0 - 100
	PercentUsed string `json:"percentUsed"`

	// Used capacity, kb
	Used int `json:"used"`

	// Objects number in the pool
	Objects int `json:"objects"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageServiceList is a list of StorageService resources
type StorageServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of StorageServices.
	Items []StorageService `json:"items"`
}
