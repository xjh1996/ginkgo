package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// StorageTypeKind is the name of the StorageType resource kind.
	StorageTypeKind = "StorageType"

	// StorageTypeName is the name of the StorageType resource (plural).
	StorageTypeName = "storagetypes"

	// StorageTypeKindKey is used as the key when mapping to the StorageType resource.
	StorageTypeKindKey = "storagetype"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Provisioner",type="string",JSONPath=".provisioner",description="The provisioner name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// StorageType describes the parameters for a class of storage for
// which StorageService can be created.
type StorageType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Provisioner indicates the name of the provisioner.
	Provisioner string `json:"provisioner" protobuf:"bytes,2,opt,name=provisioner"`

	// RequiredParameters holds the parameters for the provisioner that should
	// create volumes of this storage type.
	// Required ones for create storage service.
	// +optional
	RequiredParameters map[string]string `json:"requiredParameters,omitempty" protobuf:"bytes,3,rep,name=requiredParameters"`

	// OptionalParameters holds the parameters for the provisioner that should
	// create volumes of this storage type.
	// Required ones for create storage class.
	// +optional
	OptionalParameters map[string]string `json:"optionalParameters,omitempty" protobuf:"bytes,3,rep,name=classOptionalParameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageTypeList is a list of StorageType resources
type StorageTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of StorageTypes.
	Items []StorageType `json:"items"`
}
