/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	custom_metav1 "github.com/caicloud/api/meta/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// PrometheusClusterSetsKind is the name of the PrometheusClusterSet resource kind.
	PrometheusClusterSetsKind = "PrometheusClusterSet"
	// PrometheusClusterSetName is the name of the PrometheusClusterSet resource (plural).
	PrometheusClusterSetName = "prometheusclustersets"
	// PrometheusClusterSetKindKey is used as the key when mapping to the PrometheusClusterSet resource.
	PrometheusClusterSetKindKey = "prometheusclusterset"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName="pcs"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.prometheus.spec.version",description="The version of Prometheus"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.prometheus.spec.replicas",description="The desired replicas number of Prometheuses per cluster"
// +kubebuilder:printcolumn:name="Clusters",type="integer",JSONPath=".status.desiredClusters",description="The desired number of clusters to deploy Prometheuses"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// PrometheusClusterSet describes a multi-cluster deployment of Prometheus instances. Note that
// PrometheusClusterSet only controls the Prometheus resource from prometheus-operator; it does
// not see or care about the resources created by prometheus-operator.
type PrometheusClusterSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec defines the specifications of the desired Prometheus resource and how to deploy it
	// across clusters.
	Spec PrometheusClusterSetSpec `json:"spec"`
	// Status is the most recent observation of the deployment status of Prometheus resource
	// across the clusters.
	Status *PrometheusClusterSetStatus `json:"status,omitempty"`
}

// PrometheusClusterSetSpec defines the specifications of a PrometheusClusterSet object.
type PrometheusClusterSetSpec struct {
	// ClusterSelector selects Cluster resources; Prometheus resources will be deployed onto
	// the clusters corresponding to these Cluster resources. If left empty, all clusters will
	// be selected.
	// +optional
	ClusterSelector metav1.LabelSelector `json:"clusterSelector"`
	// Prometheus defines the Prometheus resources that will be created across clusters.
	Prometheus PrometheusTemplate `json:"prometheus"`
	// Service defines the Service resource that will be created for the Prometheus deployment
	// across clusters.
	// +optional
	Service *PrometheusService `json:"service,omitempty"`
	// Paused can be used to stop the controller from reacting to this PrometheusClusterSet.
	// +optional
	Paused bool `json:"paused,omitempty"`
}

// PrometheusTemplate describes the Prometheus resources that will be created.
type PrometheusTemplate struct {
	// Metadata describes the metadata of the Prometheus resources that will be created.
	// +optional
	Metadata *custom_metav1.EmbeddedObjectMeta `json:"metadata,omitempty"`
	// Spec describes the specification of the Prometheus resources that will be created.
	Spec monitoringv1.PrometheusSpec `json:"spec"`
}

// PrometheusService describes the Service that will be created.
type PrometheusService struct {
	// Type determines how the Service is exposed. It corresponds to .spec.type field
	// from corev1.Service.
	Type corev1.ServiceType `json:"type,omitempty"`
	// ClusterIP corresponds to .spec.clusterIP field from corev1.Service
	ClusterIP string `json:"clusterIP,omitempty"`
	// Prometheus describes the port specification for the Prometheus HTTP port.
	// +optional
	Prometheus *ServicePort `json:"prometheus,omitempty"`
	// ThanosGRPC describes the port specification for the ThanosGRPC gRPC port.
	// +optional
	ThanosGRPC *ServicePort `json:"thanosGRPC,omitempty"`
	// ThanosHTTP describes the port specification for the ThanosHTTP gRPC port.
	// +optional
	ThanosHTTP *ServicePort `json:"thanosHTTP,omitempty"`
}

// ServicePort is similar to corev1.ServicePort but only with the necessary fields
// to create a Prometheus service (other fields are automatically generated).
type ServicePort struct {
	// Port is the port that will be exposed by this service.
	Port int32 `json:"port,omitempty"`
	// NodePort is the port on each node on which this service is exposed when
	// type=NodePort or LoadBalancer.
	// +optional
	NodePort int32 `json:"nodePort,omitempty"`
}

// PrometheusClusterSetStatus describes the status of a PrometheusClusterSet object.
type PrometheusClusterSetStatus struct {
	// DesiredClusters denotes the number of clusters that are within the scope of this PrometheusClusterSet.
	DesiredClusters int `json:"desiredClusters"`
	// AvailableClusters denotes the number of clusters that have reconciled successfully.
	AvailableClusters int `json:"availableClusters"`
	// Clusters is a list of ClusterStatuses of every cluster that have been affected by this PrometheusClusterSet.
	Clusters []ClusterStatus `json:"clusterStatus,omitempty"`
}

// +kubebuilder:validation:Enum=OK;Pending;Error;GarbageCollection

// State describes the deployment status of the desired resource.
type State string

const (
	// StateOK means that the resource has been successfully deployed without errors.
	StateOK State = "OK"
	// StatePending means that the resource cannot be deployed due to some missing
	// prerequisites.
	StatePending State = "Pending"
	// StateError means that the resource cannot be deployed due to an error.
	StateError State = "Error"
	// StateGC means that the existing resource is being reclaimed.
	StateGC State = "GarbageCollection"
)

// ClusterStatus describes the deployment status of the desired resources on a particular
// cluster.
type ClusterStatus struct {
	// Cluster is the name of the Cluster resource corresponding to the cluster.
	Cluster string `json:"cluster,omitempty"`
	// State describes the current state of the Prometheus resource on the cluster.
	State State `json:"state,omitempty"`
	// Message gives a full explanation on why the Prometheus resource is in its current
	// state.
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrometheusClusterSetList is a list of PrometheusClusterSet resources
type PrometheusClusterSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PrometheusClusterSet `json:"items"`
}
