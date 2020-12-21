package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterKind is the name of the Cluster resource kind.
	ClusterKind = "Cluster"

	// ClusterName is the name of the Cluster resource (plural).
	ClusterName = "clusters"

	// ClusterKindKey is used as the key when mapping to the Cluster resource.
	ClusterKindKey = "cluster"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="CloudProvider",type="string",JSONPath=".spec.provider",description="The infrastructure provider name of Cluster"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Current phase of Cluster actuation"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version.kubernetesVersion",description="Current kubernetes version of Cluster"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of Cluster
	Spec ClusterSpec `json:"spec"`

	// Status defines the observed state of Cluster
	// +optional
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// InfrastructureRef is a reference to a provider-specific resource that holds the details
	// for provisioning infrastructure for a cluster in said provider.
	InfrastructureRef *corev1.ObjectReference `json:"infrastructureRef,omitempty"`

	// Provider is the name of infrastructure provider.
	// NOTE: if this field is empty, this will be synced from infrastructure .spec.provider field
	// +optional
	Provider CloudProvider `json:"provider,omitempty"`

	// ControlPlaneRef is an optional reference to a provider-specific resource that holds
	// the details for provisioning the Control Plane for a Cluster.
	// NOTE: if this is not empty, the k8s will provisioned by ref object.
	// +optional
	ControlPlaneRef *corev1.ObjectReference `json:"controlPlaneRef,omitempty"`

	// ClusterNetwork is the network configuration.
	// +optional
	ClusterNetwork *ClusterNetwork `json:"clusterNetwork,omitempty"`

	// ControlPlaneEndpoints represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoints []APIEndpoint `json:"controlPlaneEndpoints,omitempty"`

	// ServiceEndpoints represents the endpoint used to communicate with cluster-scope components service.
	// +optional
	ServiceEndpoints []APIEndpoint `json:"serviceEndpoints,omitempty"`

	// IsControlCluster represents the cluster role for a Cluster.
	IsControlCluster bool `json:"isControlCluster"`

	// Version represents the version for a Cluster.
	Version ClusterVersion `json:"version"`

	// Auth represents the auth information for a Cluster.
	// NOTE: the controlplane provider **MUST** filled this field after the kubernetes is provisioned.
	// +optional
	Auth ClusterAuth `json:"auth,omitempty"`
}

// ClusterNetwork specifies the different networking parameters for a cluster.
type ClusterNetwork struct {
	// Services defines the network range from which service VIPs are allocated.
	// +optional
	Services *NetworkRanges `json:"services,omitempty"`

	// Pods defines the network range from which Pod networks are allocated.
	// +optional
	Pods *NetworkRanges `json:"pods,omitempty"`

	// DNSServiceIP defines the ip of dns service, and it must within the service network range.
	// +optional
	DNSServiceIP string `json:"dnsServiceIP,omitempty"`

	// NetworkRef is a reference to network advance configuration object.
	// +optional
	NetworkRef *corev1.ObjectReference `json:"networkRef,omitempty"`

	// Parameters holds parameters for the network.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

// NetworkRanges represents ranges of network addresses.
type NetworkRanges struct {
	// CIDRBlocks is the list of cidr addresses.
	// +optional
	CIDRBlocks []string `json:"cidrBlocks,omitempty"`
}

// APIEndpoint represents the apiserver endpoint of cluster.
type APIEndpoint struct {
	// Address is the network address of endpoint.
	Address string `json:"address"`

	// Type is the address type
	Type corev1.NodeAddressType `json:"type"`

	// Port is the apiserver access port.
	Port int32 `json:"port"`
}

// ClusterVersion represents the version for a Cluster.
type ClusterVersion struct {
	// KubernetesVersion is the desired kubernetes version of a Cluster.
	KubernetesVersion string `json:"kubernetesVersion"`

	// Parameters holds parameters for the Bootstrap Provider.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

// ClusterAuth represents the auth information for a Cluster.
type ClusterAuth struct {
	// KubeConfig is the cluster access information about clusters, users, namespaces, and authentication mechanisms.
	// This will been filled by Control Plane Provider after the Control Plane is provisioned.
	// +optional
	KubeConfig []byte `json:"kubeConfig,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// Phase represents the current phase of cluster actuation.
	// +optional
	Phase ClusterPhase `json:"phase,omitempty"`

	// Conditions represents the observations of cluster's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

const (
	// ClusterConditionControlPlaneReady defines the controlplane result, if the status is False the cluster will be consider as Failed
	// NOTE: this field is sync from controlplane ref .status.conditions[Ready]
	ClusterConditionControlPlaneReady = "ControlPlaneReady"

	// ClusterConditionControlPlaneDelete defines the controlplane delete result, if the status is True the cluster will be consider deleteable.
	// NOTE: this field **MUST** be set when there is no controlplane ref
	ClusterConditionControlPlaneDelete = "ControlPlaneDelete"

	// ClusterConditionInfrastructureReady defines the infrastructure result, if the status is False the cluster will be consider as Failed
	// NOTE: this field is sync from infrastrcture ref .status.conditions[Ready]
	ClusterConditionInfrastructureReady = "InfrastrctureReady"

	// ClusterConditionReady defines the cluster create result, if the status is True the cluster will be consider as Running.
	// NOTE: this field will summary from all condition witch type has Ready suffix,
	//    1. if all *Ready condition is True, this condition is True
	//    2. if any *Ready condition is False, this condition is False
	//    3. otherwise, this condition is Unknown
	ClusterConditionReady = "Ready"

	// ClusterConditionHealth defines the cluster health result, if the stauts is True the cluster will consider as available.
	// NOTE: this field will summary from all condition witch type has Health suffix,
	//    1. if all *Health condition is True, this condition is True
	//    2. if any *Health condition is False, this condition is False
	//    3. otherwise, this condition is Unknown
	ClusterConditionHealth = "Health"

	// ClusterConditionAPIServerHealth defines the cluster apiserver health result
	ClusterConditionAPIServerHealth = "APIServerHealth"

	// ClusterConditionMasterHealth defines the cluster master's health result
	// If the quorum master is not ready, this condition is False
	ClusterConditionMasterHealth = "MasterHealth"

	// ClusterConditionEtcdHealth defines the cluster etcd's health result
	// If the is quorum etcd is not ready, this condition is False
	ClusterConditionEtcdHealth = "EtcdHealth"

	// ClusterConditionAddonHealth defines the cluster addon's health result
	// NOTE: this field is managed by other's component, we just define it at there as well known condition type
	ClusterConditionAddonHealth = "AddonHealth"
)

// +kubebuilder:validation:Enum=Pending;Provisioning;Provisioned;Bootstraping;Running;Deleting;Failed;Unknown

// ClusterPhase is a string representation of a Cluster Phase.
//
// This type is a high-level indicator of the status of the Cluster as it is provisioned,
// from the API user’s perspective.
//
// The value should not be interpreted by any software components as a reliable indication
// of the actual state of the Cluster, and controllers should not use the Cluster Phase field
// value when making decisions about what action to take.
//
// Controllers should always look at the actual state of the Cluster’s fields to make those decisions.
type ClusterPhase string

const (
	// ClusterPhasePending is the first state a Cluster is assigned by cluster operator after being created.
	ClusterPhasePending ClusterPhase = "Pending"

	// ClusterPhaseProvisioning is the state when the Cluster has a provider infrastructure
	// object associated and can start provisioning.
	ClusterPhaseProvisioning ClusterPhase = "Provisioning"

	// ClusterPhaseProvisioned is the state when it's infrastructure has been created and configured.
	ClusterPhaseProvisioned ClusterPhase = "Provisioned"

	// ClusterPhaseBootstraping is the state when the Cluster has a ControlPlane provider object
	// associated and can start kubernetes provisioning.
	ClusterPhaseBootstraping ClusterPhase = "Bootstraping"

	// ClusterPhaseRunning is the state when the Cluster has been provisioned(infrastructure and controlplane).
	ClusterPhaseRunning ClusterPhase = "Running"

	// ClusterPhaseDeleting is the Cluster state when a delete request has been sent to the API Server,
	// but its controlplane and infrastructure has not yet been fully deleted.
	ClusterPhaseDeleting ClusterPhase = "Deleting"

	// ClusterPhaseFailed is the Cluster state when the system might require user intervention.
	ClusterPhaseFailed ClusterPhase = "Failed"

	// ClusterPhaseUnknown is returned if the Cluster state cannot be determined.
	ClusterPhaseUnknown ClusterPhase = "Unknown"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList is a list of Cluster resources
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of Clusters.
	Items []Cluster `json:"items"`
}
