package v1

import (
	"time"
)

// ObjectMeta contains generic metadata that all API objects must have. Unlike K8s v1.ObjectMeta,
// ObjectMeta is intended to be used inlined at the root level of a full API object. This is because
// it has to be compatible with MongoDB and MongoDB requires `_id` at the root level.
type ObjectMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Cannot be updated.
	// In CamelCase.
	// eg: StorageType
	Kind string `json:"Kind" yaml:"Kind" bson:"Kind"`
	// APIVersion defines the versioned schema of this representation of an object.
	// eg: v1alpha1 / v1 / v2beta1
	APIVersion string `json:"ApiVersion" yaml:"ApiVersion" bson:"ApiVersion"`
	// UID is the unique id for this object, It can be UUID/GUID/Auto-incrementing Key.
	// eg: 123456 / 6f16d5e7ff2272c9cd45b24fde186681 / 668b6e0c-1deb-11e8-9023-1c1b0dffb81f
	UID string `json:"Uid" yaml:"Uid" bson:"_id"`
	// CreationTimestamp represents the object's creation time, it is represented in RFC3339 form and is in UTC.
	// eg: 2020-05-20T13:00:55Z
	CreationTimestamp time.Time `json:"CreationTimestamp" yaml:"CreationTimestamp" bson:"CreationTimestamp"`
	// Name is the unique name for this object.
	// For non-Kubernetes objects, this field might be empty.
	// +optional
	Name string `json:"Name,omitempty" yaml:"Name,omitempty" bson:"Name,omitempty"`
	// Namespace defines the space within each name must be unique.
	// +optional
	Namespace string `json:"Namespace,omitempty" yaml:"Namespace,omitempty" bson:"Namespace,omitempty"`
	// Tenant is the name of the tenant which the object belongs to.
	// +optional
	Tenant string `json:"Tenant,omitempty" yaml:"Tenant,omitempty" bson:"Tenant,omitempty"`
	// ClusterName is the name of the cluster which the object belongs to.
	// +optional
	ClusterName string `json:"ClusterName,omitempty" yaml:"ClusterName,omitempty" bson:"ClusterName,omitempty"`
	// Alias is the alias for this object.
	// +optional
	Alias string `json:"Alias,omitempty" yaml:"Alias,omitempty" bson:"Alias,omitempty"`
	// Description is the description for this object.
	// +optional
	Description string `json:"Description,omitempty" yaml:"Description,omitempty" bson:"Description,omitempty"`
	// DeletionTimestamp represents the object's deletion time, it is represented in RFC3339 form and is in UTC.
	// eg: 2020-05-20T13:00:55Z
	// +optional
	DeletionTimestamp *time.Time `json:"DeletionTimestamp,omitempty" yaml:"DeletionTimestamp,omitempty" bson:"DeletionTimestamp,omitempty"`
	// Labels is a map of string keys and values that can be used to organize and categorize
	// (scope and select) objects.
	// For Kubernetes objects, only return the contents of this field when really necessary.
	// The underlying value is map[string]string.
	Labels string `json:"Labels,omitempty" yaml:"Labels,omitempty" bson:"Labels,omitempty"`
	// Annotations is a map of string keys and values that can be used to attach arbitrary non-identifying
	// metadata to objects.
	// For Kubernetes objects, only return the contents of this field when really necessary.
	// The underlying value is map[string]string.
	Annotations string `json:"Annotations,omitempty" yaml:"Annotations,omitempty" bson:"Annotations,omitempty"`
}

// ListMeta contains generic metadata that all List objects must have. A List object is a list of API
// objects sharing the same TypeMeta.
type ListMeta struct {
	Kind       string `json:"Kind" yaml:"Kind"`
	APIVersion string `json:"ApiVersion" yaml:"ApiVersion"`
	Total      int    `json:"Total" yaml:"Total"`
}
