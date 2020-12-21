package v1

// +k8s:deepcopy-gen=true

// EmbeddedObjectMeta contains those fields from k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta that might
// be relevant when creating new Kubernetes resource objects. It is useful when defining a template-like
// structure in a parent resource that creates other resources.
type EmbeddedObjectMeta struct {
	// Name is the unique name given to the created resources. Some resources might support automatic generation
	// for this field, in which case you might leave it empty.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace defines the namespace of the created resources. This field is only intended for the rare
	// occasion when you want your parent resource and created resources to be in different namespaces. In most
	// cases, you might want to leave this empty and just use the namespace of the parent resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Labels defines the additional labels that should be added to the created resources.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// Labels defines the additional annotations that should be added to the created resources.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}
