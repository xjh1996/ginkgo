/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1beta1

import (
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	// LoadBalancerKind is the name of the LoadBalancer resource kind.
	LoadBalancerKind = "LoadBalancer"
	// LoadBalancerName is the name of the LoadBalancer resource (plural).
	LoadBalancerName = "loadbalancers"
	// LoadBalancerKindKey is used as the key when mapping to the LoadBalancer resource.
	LoadBalancerKindKey = "loadbalancer"
	// NginxProxyName is the type name of ingress-nginx controller.
	NginxProxyName = "nginx"
	// IngressNginxController is the name of ingress-nginx controller.
	IngressNginxController = "k8s.io/ingress-nginx"
	// IPVSDRProviderName is the name of ipvsdr provider.
	IPVSDRProviderName = "ipvs"
	// ExternalProviderName is the name of external provider.
	ExternalProviderName = "external"
	// DefaultIngressController is the name of default ingress controller.
	DefaultIngressController = IngressNginxController
	// DefaultProxyName is the default type name of proxy.
	DefaultProxyName = NginxProxyName
	// DefaultProviderName is the name of default provider.
	DefaultProviderName = IPVSDRProviderName
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName="lb"
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="VIPS",type="string",JSONPath=".spec.provider.listener.addresses",description="Entrance Virtual IPs of LoadBalancer"
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.provider.class",description="Provider class type of this LoadBalancer"
// +kubebuilder:printcolumn:name="Proxy",type="string",JSONPath=".spec.proxy.class",description="Proxy class type of this LoadBalancer"
// +kubebuilder:printcolumn:name="Nodes",type="string",JSONPath=".spec.resources.nodes.matches",description="Node list of LoadBalancer hosted"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Current phase of LoadBalancer"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// LoadBalancer describes a LoadBalancer which provides Load Balancing for applications
// LoadBalancer contains a proxy and multiple providers to load balance
// either internal or external traffic.
//
// A proxy is an ingress controller watching ingress resource to provide access that
// allow inbound connections to reach the cluster services
//
// A provider is the entrance of the cluster providing high availability for connections
// to proxy (ingress controller)
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the LoadBalancer.
	Spec LoadBalancerSpec `json:"spec"`
	// Most recently observed status of the loadbalancer.
	// This data may not be up to date.
	// Populated by the system.
	// Read-only.
	// +optional
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoadBalancerList is a collection of LoadBalancer.
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []LoadBalancer `json:"items"`
}

// LoadBalancerSpec is a description of a LoadBalancer.
type LoadBalancerSpec struct {
	// Resources specify resources loadbalancer can use.
	// +optional
	Resources Resources `json:"resources,omitempty"`
	// Proxy defines parameters and resources for ingress controller.
	// +optional
	Proxy Proxy `json:"proxy,omitempty"`
	// Provider defines parameters for entrance loadblancer.
	// +optional
	Provider Provider `json:"provider,omitempty"`
}

// Resources defines what resources a loadbalancer can use.
type Resources struct {
	// Nodes defines which nodes should be chosed by loadbalancer.
	// +optional
	Nodes Nodes `json:"nodes,omitempty"`
	// Ports defines port ranges the loadbalancer can use.
	// +optional
	Ports []PortRange `json:"ports,omitempty"`
}

// Nodes defines nodes can be used by loadbalancer.
type Nodes struct {
	// Selector is a label selector for loadbalancer to choose nodes.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// Matches specify a list of hostnames of nodes.
	// +optional
	Matches []string `json:"matches,omitempty"`
}

// PortRange defines a range type of port. Start must be less or equal to End.
type PortRange struct {
	Start int32 `json:"start"`
	End   int32 `json:"end"`
}

// Proxy defines the specifications of Proxy.
type Proxy struct {
	// Class is the proxy class type.
	Class string `json:"class,omitempty"`
	// Replica is the replica number of instances of ingress controller.
	// +optional
	Replica *int32 `json:"replica,omitempty"`
	// IngressClass defines parameters of ingress controller.
	IngressClass IngressClass `json:"ingressClass,omitempty"`
	// Service defines configurations of the service created for ingress controller.
	Service IngressService `json:"service,omitempty"`
	// Resources defines resources requirements for ingress controller.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

// IngressClass defines parameters for ingress controller.
type IngressClass struct {
	// Name is the name of ingress class, used by ingress controller.
	Name string `json:"name,omitempty"`
	// Controller is the name of ingress controller.
	// This should be specified as a domain-prefixed path,e.g. "acme.io/ingress-controller"
	Controller string `json:"controller,omitempty"`
	// Parameters defines additional configuration for the controller.
	// It is a TypedLocalObject.
	Parameters *Parameters `json:"parameters,omitempty"`
}

// IngressService defines configuration of service for ingress controller.
type IngressService struct {
	// Host networking requested for ingress controller.
	// Default to true.
	// +k8s:conversion-gen=true
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`
	// Name is the name of the service.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace is the namespace of the service.
	Namespace string `json:"namespace,omitempty"`
	// Ports is list of service port ingress controller exposes.
	Ports []corev1.ServicePort `json:"ports,omitempty"`
}

// Provider defines parameters for entrance loadbalancer provider.
type Provider struct {
	// Class is the class type of provider.
	Class string `json:"class,omitempty"`
	// Replica is the replica number of provider instances.
	// Only work when nodes are not defined.
	// +optional
	Replica *int32 `json:"replica,omitempty"`
	// Listener defines listner for entrance loadbalancer.
	Listener Listener `json:"listener,omitempty"`
	// Parameters defines additional configuration for provider.
	// It is a TypedLocalObject.
	Parameters *Parameters `json:"parameters,omitempty"`
}

// Listener defines specifications of listener.
type Listener struct {
	// Addresses contains a list of ip addresses.
	// The list could be different version of ip address or some ip addresses of same version.
	Addresses []string `json:"addresses,omitempty"`
	// Ports defines ports listener should listen and ports listener should forward to.
	Ports []corev1.ServicePort `json:"ports,omitempty"`
}

// Parameters specifies parameters that should be passed to a loadbalancer at the time of initialization.
type Parameters struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:EmbeddedResource
	// +kubebuilder:validation:nullable
	runtime.RawExtension `json:",inline"`
}

// DecodeInto decode raw data to runtime object.
func (in *Parameters) DecodeInto(obj runtime.Object) error {
	// TODO: if DecodeNestedObjects could fix this.
	// TODO: in.Object is not nil
	// only support internal version
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(schema.GroupVersion{Group: "", Version: runtime.APIVersionInternal}, obj)
	codec := serializer.NewCodecFactory(scheme, serializer.DisableStrict)
	out, _, err := codec.UniversalDecoder().Decode(in.Raw, nil, obj)
	if err != nil || out == obj {
		return errors.Wrapf(err, "universal decoder error")
	}
	return nil
}

// LoadBalancerStatus represents the current status of a LoadBalancer.
type LoadBalancerStatus struct {
	// Ready indicates ready status of loadbalancer.
	// Default to false.
	// +k8s:conversion-gen=false
	Ready bool `json:"ready,omitempty"`
	// Phase shows the different phase state of loadbalancer.
	// Default to Pending.
	// +k8s:conversion-gen=Pending
	Phase LoadBalancerPhase `json:"phase,omitempty"`
	// Resources Statuses contain the details of resources.
	Resources ResourcesStatuses `json:"resources,omitempty"`
	// Proxy Status represents the current status of proxy controller.
	Proxy ControllerStatus `json:"proxy,omitempty"`
	// Provider Status represents the current status of provider controller.
	Provider ControllerStatus `json:"provider,omitempty"`
}

// ResourcesStatuses contain details of resources.
type ResourcesStatuses struct {
	NodesStatuses []NodeStatus `json:"nodesStatuses,omitempty"`
}

// NodeStatus contains details of nodes.
type NodeStatus struct {
	Name         string          `json:"name,omitempty"`
	IfaceNetList []*InterfaceNet `json:"ifaces,omitempty"`
}

// InterfaceNet represents the current status of an interface.
type InterfaceNet struct {
	Name string   `json:"name,omitempty"`
	Mac  string   `json:"mac,omitempty"`
	IPs  []string `json:"ips,omitempty"`
}

// LoadBalancerPhase type
type LoadBalancerPhase string

const (
	// LoadBalancerPending means LoadBalancer is waiting to be accepted by controller.
	LoadBalancerPending LoadBalancerPhase = "Pending"
	// LoadBalancerCreating means controller is deploying this loadbalancer.
	LoadBalancerCreating LoadBalancerPhase = "Creating"
	// LoadBalancerRunning means proxy and provider are deployed successfully.
	LoadBalancerRunning LoadBalancerPhase = "Running"
	// LoadBalancerFailed means LoadBalancer failed, check conditions.
	LoadBalancerFailed LoadBalancerPhase = "Failed"
)

// ControllerStatus represent current status of proxy or provider controller.
type ControllerStatus struct {
	Pods       PodStatuses        `json:"pods,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// PodStatuses represents the current statuses of a list of pods.
type PodStatuses struct {
	ExpectedReplicas int32       `json:"expectedReplicas,omitempty"`
	Replicas         int32       `json:"replicas,omitempty"`
	Statuses         []PodStatus `json:"podStatuses,omitempty"`
}

// PodStatus represents the current status of pods.
type PodStatus struct {
	Name     string `json:"name,omitempty"`
	Ready    bool   `json:"ready,omitempty"`
	Phase    string `json:"phase,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Message  string `json:"message,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
}
