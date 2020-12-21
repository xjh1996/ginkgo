package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineKind is the name of the Machine resource kind.
	MachineKind = "Machine"

	// MachineName is the name of the Machine resource (plural).
	MachineName = "machines"

	// MachineKindKey is used as the key when mapping to the Machine resource.
	MachineKindKey = "machine"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Current phase of Machine"
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.clusterName",description="The cluster id of Machine"
// +kubebuilder:printcolumn:name="CloudProvider",type="string",JSONPath=".spec.provider",description="The infrastructure provider name of Machine"
// +kubebuilder:printcolumn:name="ProviderID",type="string",JSONPath=".spec.providerID",description="The Cloud Provider ID of Machine"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// Machine is the Schema for the machines API
type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of Machine
	Spec MachineSpec `json:"spec"`

	// Status defines the observed state of Machine
	// +optional
	Status MachineStatus `json:"status,omitempty"`
}

// MachineSpec defines the desired state of Machine
type MachineSpec struct {
	// InfrastructureRef is a required reference to a custom resource
	// offered by an infrastructure provider.
	InfrastructureRef *corev1.ObjectReference `json:"infrastructureRef"`

	// Provider defines the CloudProvider name of infrastructure provider.
	// NOTE: if this field is empty, this will be synced from infrastructure .spec.provider field
	// +optional
	Provider CloudProvider `json:"provider,omitempty"`

	// ProviderID is the identification ID of the machine provided by the provider.
	// This field must match the provider ID as seen on the node object corresponding to this machine.
	// NOTE: if this field is empty, this will be sync from infrastructure .spec.providerID field
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// Bootstrap is a reference to a local struct which encapsulates
	// fields to configure the Machine’s bootstrapping mechanism.
	// +optional
	Bootstrap Bootstrap `json:"bootstrap,omitempty"`

	// ClusterName is the name of the Cluster this object belongs to.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// Version defines the desired machine component version.
	// This field is meant to be optionally used by bootstrap providers.
	// +optional
	Version *MachineVersion `json:"version,omitempty"`

	// Tags defines the desired node labels should be appled to corresponding Node object.
	// +optional
	Tags MachineTags `json:"tags,omitempty"`

	// Taints is a list of taints of a machine.
	// All taints here will be synced to a corresponding Node object when adding a machine to a cluster, and will be cleared when the machine is removed from a cluster.
	// This field is used internally via machine controller to persistently represent taints when a machine is added to a cluster (when Node object doesn't exist); otherwise the information will be lost across component restart.
	// NOTE: If the value is "-" (witch is not allowed by labels syntax), the taint will be removed from node.
	// +optional
	Taints []corev1.Taint `json:"taints,omitempty"`

	// Auth defines the SSH information of a machine.
	// NOTE: if this field is empty, the value will sync from infrastructure .status.auth field with same strcut type.
	// +optional
	Auth *MachineAuth `json:"auth,omitempty"`
}

// Bootstrap capsulates fields to configure the Machine’s bootstrapping mechanism.
type Bootstrap struct {
	// ConfigRef is a reference to a bootstrap provider-specific resource
	// that holds configuration details. The reference is optional to
	// allow users/operators to specify Bootstrap.Data without
	// the need of a controller.
	// +optional
	ConfigRef *corev1.ObjectReference `json:"configRef,omitempty"`

	// DataSecretName is the name of the secret that stores the bootstrap data script.
	// +optional
	DataSecretName *string `json:"dataSecretName,omitempty"`
}

// MachineVersion defines the desired machine component version.
type MachineVersion struct {
	// KubernetesVersion defines the desired Kubernetes version.
	// +optional
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
}

// MachineTags defines the desired labels.
// More info: http://kubernetes.io/docs/user-guide/labels
// NOTE: If the value is "-" (witch is not allowed by labels syntax), the label will be removed from node.
type MachineTags map[string]string

const (
	// MachineTagValueForRemoved defines the placeholder value for remove tags and taints
	MachineTagValueForRemoved = "-"
)

// MachineAuth defines the machine SSH information.
type MachineAuth struct {
	// User defines the SSH username.
	User string `json:"user"`

	// Password defines the password of SSH connection.
	// This maybe empty when machine should connect via PublicKey.
	// +optional
	Password string `json:"password,omitempty"`

	// Key defines the secret name witch contains the PublicKey/PrivateKey in cos-system namespace.
	// +optional
	Key string `json:"key,omitempty"`

	// Addresses is List of addresses reachable to the machine.
	// This will be appled by infrastructure provider.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	Addresses []NodeAddress `json:"addresses,omitempty"`

	// Port defines the SSH port of machine.
	Port int32 `json:"port"`
}

// NodeAddress defines the address of machine.
type NodeAddress struct {
	// Type is Node address type, one of Hostname, ExternalIP or InternalIP.
	Type corev1.NodeAddressType `json:"type"`

	// Address is the reachable address.
	Address string `json:"address"`

	// IPv6Address is the reachable ipv6 address with the same inc.
	// +optional
	IPv6Address string `json:"ipv6Address,omitempty"`
}

// MachineStatus defines the observed state of Machine
type MachineStatus struct {
	// Phase represents the current phase of machine actuation.
	// +optional
	Phase MachinePhase `json:"phase,omitempty"`

	// ClusterName is the name of the Cluster this object belongs to.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// NodeRef will point to the corresponding Node if it exists.
	// This field will be populated from NodeClaim object.
	// +optional
	NodeRef *corev1.ObjectReference `json:"nodeRef,omitempty"`

	// NodeClaimRef will point to the corresponding NodeClaim if it exists.
	// +optional
	NodeClaimRef *corev1.ObjectReference `json:"nodeClaimRef,omitempty"`

	// MachineInfoRef will point to the corresponding MachineInfo if it exists.
	// +optional
	MachineInfoRef *corev1.ObjectReference `json:"machineInfoRef,omitempty"`

	// Conditions represents the observations of machine's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

const (
	// MachineConditionPreReadyHooksReady defines the pre-ready hooks execute result, it will be False if any pre-ready hooks failed
	MachineConditionPreReadyHooksReady = "PreReadyHooksReady"

	// MachineConditionBootstrapReady defines the bootstrap config result, if the status is False the machine will be consider as Failed
	// NOTE: this field is sync from bootstrap config ref .status.conditions[Ready]
	MachineConditionBootstrapReady = "BootstrapReady"

	// MachineConditionInfrastructureReady defines the infrastructure result, if the status is False the machine will be consider as Failed
	// NOTE: this field is sync from infrastrcture ref .status.conditions[Ready]
	MachineConditionInfrastructureReady = "InfrastructureReady"

	// MachineConditionReady defines the machine create result
	// NOTE: this field will summary from all condition witch type has Ready suffix,
	//    1. if all *Ready condition is True, this condition is True
	//    2. if any *Ready condition is False, this condition is False
	//    3. otherwise, this condition is Unknown
	MachineConditionReady = "Ready"

	// MachineConditionHealth defines the machine health result, if the stauts is True the machine will consider as available.
	// NOTE: this field will summary from all condition witch type has Health suffix,
	//    1. if all *Health condition is True, this condition is True
	//    2. if any *Health condition is False, this condition is False
	//    3. otherwise, this condition is Unknown
	MachineConditionHealth = "Health"
)

// +kubebuilder:validation:Enum=Pending;Provisioning;Provisioned;Ready;Binding;Allocated;Unbinding;Deleting;Failed;Unknown

// MachinePhase is a string representation of a Machine Phase.
//
// This type is a high-level indicator of the status of the Machine as it is provisioned,
// from the API user’s perspective.
//
// The value should not be interpreted by any software components as a reliable indication
// of the actual state of the Machine, and controllers should not use the Machine Phase field
// value when making decisions about what action to take.
//
// Controllers should always look at the actual state of the Machine’s fields to make those decisions.
type MachinePhase string

const (
	// MachinePhasePending is the first state a Machine is assigned by machine operator after being created.
	MachinePhasePending MachinePhase = "Pending"

	// MachinePhaseProvisioning is the state when the Machine has a provider infrastructure
	// object associated and can start provisioning.
	MachinePhaseProvisioning MachinePhase = "Provisioning"

	// MachinePhaseProvisioned is the state when it's infrastrcture has been created and configured.
	MachinePhaseProvisioned MachinePhase = "Provisioned"

	// MachinePhaseReady is the state when the Machine is ready to allocated.
	MachinePhaseReady MachinePhase = "Ready"

	// MachinePhaseBinding is the state when the Machine will be allocated to Cluster.
	MachinePhaseBinding MachinePhase = "Binding"

	// MachinePhaseAllocated is the state when the Machine has join the cluster and it's node have been created.
	MachinePhaseAllocated MachinePhase = "Allocated"

	// MachinePhaseUnbinding is the state when the Machine will be removed from Cluster.
	MachinePhaseUnbinding MachinePhase = "Unbinding"

	// MachinePhaseDeleting is the Machine state when a delete request has been send to the API Server,
	// but iss infrastructure has not yet been fully deleted.
	MachinePhaseDeleting MachinePhase = "Deleting"

	// MachinePhaseFailed is the Machine state when the system might require user intervention.
	MachinePhaseFailed MachinePhase = "Failed"

	// MachinePhaseUnknown is returned if the Machine state cannot be determined.
	MachinePhaseUnknown MachinePhase = "Unknown"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineList is a list of Machine resources
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of Machines.
	Items []Machine `json:"items"`
}
