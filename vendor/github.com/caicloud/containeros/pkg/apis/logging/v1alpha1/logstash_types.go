/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	custom_metav1 "github.com/caicloud/api/meta/v1"
	commonv1 "github.com/elastic/cloud-on-k8s/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/caicloud/containeros/pkg/apis/logging"
	crdutils "github.com/caicloud/containeros/pkg/utils/crd"
)

const (
	// LogstashClusterSetsKind is the name of the LogstashClusterSetsKind resource kind.
	LogstashClusterSetsKind = "LogstashClusterSet"
	// LogstashClusterSetName is the name of the LogstashClusterSetsKind resource (plural).
	LogstashClusterSetName = "logstashclustersets"
	// LogstashClusterSetKindKey is used as the key when mapping to the LogstashClusterSet resource.
	LogstashClusterSetKindKey = "logstashclusterset"

	// LogstashsKind is the name of the LogstashsKind resource kind.
	LogstashsKind = "Logstash"
	// LogstashName is the name of the LogstashsKind resource (plural).
	LogstashName = "logstashs"
	// LogstashKindKey is used as the key when mapping to the Logstash resource.
	LogstashKindKey = "logstash"
)

// LogstashClusterSetCRDChartFilename is the filename of the LogstashClusterSetName CRD chart.
var LogstashClusterSetCRDChartFilename = crdutils.GenerateCRDChartFilename(logging.GroupName, LogstashClusterSetName)

// LogstashCRDChartFilename is the filename of the LogstashName CRD chart.
var LogstashCRDChartFilename = crdutils.GenerateCRDChartFilename(logging.GroupName, LogstashName)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName="lcs"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.template.spec.version",description="The version of Logstash"
// +kubebuilder:printcolumn:name="Clusters",type="integer",JSONPath=".status.expectedClusters",description="The expected number of clusters to deploy Logstashes"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// LogstashClusterSet describes a multi-cluster deployment of Logstash instances. Note that
// LogstashClusterSet only controls the Logstash resource with some Secret from logstash-operator; it does
// not see or care about the resources created by logstash-operator.
type LogstashClusterSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec defines the specifications of the desired Logstash resource and how to deploy it
	// across clusters.
	Spec LogstashClusterSetSpec `json:"spec"`
	// Status is the most recent observation of the deployment status of Logstash resource
	// across the clusters.
	Status *LogstashClusterSetStatus `json:"status,omitempty"`
}

// LogstashClusterSetSpec defines the specifications of a LogstashClusterSet object.
type LogstashClusterSetSpec struct {
	// ClusterSelector selects Cluster resources; Logstash resources will be deployed onto
	// the clusters corresponding to these Cluster resources. If left empty, all clusters will
	// be selected.
	// +optional
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector,omitempty"`
	// ElasticsearchEndpointSelector selects ElasticsearchEndpoint which is a reference to an existing Elasticsearch Cluster.
	// It allows automatioc setup of Index Template and ILM.
	// +optional
	ElasticsearchEndpointSelector *metav1.LabelSelector `json:"elasticsearchEndpointSelector,omitempty"`
	// PipelineRef contains a reference to an existing Kubernetes Secret holding the Logstash Pipeline configuration.
	// Pipeline settings must be specified as yaml, under a single "pipelines.yml" entry.
	PipelineRef *ConfigSource `json:"pipelineRef,omitempty"`
	// Selector is provided to select Logstash resources that were created on account of this
	// LogstashClusterSet. If left empty, it will be generated automatically according to the
	// Logstash template.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// Template defines the Logstash resources that will be created across clusters.
	Template LogstashTemplate `json:"template"`
}

// LogstashTemplate describes the Logstash resources that will be created.
type LogstashTemplate struct {
	// Metadata describes the metadata of the Logstash resources that will be created.
	// +optional
	Metadata *custom_metav1.EmbeddedObjectMeta `json:"metadata,omitempty"`
	// Spec describes the specification of the Logstash resources that will be created.
	Spec LogstashSpec `json:"spec"`
}

// LogstashClusterSetStatus describes the status of a LogstashClusterSet object.
type LogstashClusterSetStatus struct {
	Conditions       []LogstashClusterSetCondition `json:"conditions,omitempty"`
	ExpectedClusters int                           `json:"expectedClusters"`
	UpdatedClusters  int                           `json:"updatedClusters"`
	LogstashStatus   []ClusterStatus               `json:"logstashStatus,omitempty"`
}

// +kubebuilder:validation:Enum=IndexSettingsReady;LogstashsUpdated

// LogstashClusterSetConditionType is the type of the condition.
type LogstashClusterSetConditionType string

const (
	// LogstashClusterSetIndexSettingsReady indicates whether all Index`s Settings are Ready
	LogstashClusterSetIndexSettingsReady LogstashClusterSetConditionType = "IndexSettingsReady"
	// LogstashClusterSetLogstashsUpdated indicates whether all Logstashs on the Cluster are Updated
	LogstashClusterSetLogstashsUpdated LogstashClusterSetConditionType = "LogstashsUpdated"
)

// LogstashClusterSetCondition contains details for the current condition of this LogstashClusterSet.
type LogstashClusterSetCondition struct {
	// Type is the type of the condition.
	Type LogstashClusterSetConditionType `json:"type"`

	// Status is the status of the condition.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogstashClusterSetList is a list of LogstashClusterSet resources
type LogstashClusterSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []LogstashClusterSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version",description="The version of Logstash"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".status.expectedReplicas",description="The expected number of replicas to deploy Logstashes"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// Logstash defines a Logstash deployment.
type Logstash struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogstashSpec    `json:"spec"`
	Status *LogstashStatus `json:"status,omitempty"`
}

// LogstashSpec defines the desired state of a Logstash.
type LogstashSpec struct {
	// Version of Logstash.
	Version string `json:"version"`

	// Image is the Logstash image to deploy.
	Image string `json:"image,omitempty"`

	// HTTP holds HTTP layer settings for Logstash.
	HTTP commonv1.HTTPConfig `json:"http,omitempty"`

	// PipelineRef contains a reference to an existing Kubernetes Secret holding the Logstash Pipeline configuration.
	// Pipeline settings must be specified as yaml, under a single "pipelines.yml" entry.
	PipelineRef *ConfigSource `json:"pipelineRef,omitempty"`

	// Deployment provides customisation options (replicas, podTemplate) for the Deployment belonging to this Logstash Cluster.
	Deployment *DeploymentSpec `json:"deployment,omitempty"`
}

// LogstashStatus defines the observed state of a Logstash.
type LogstashStatus struct {
	ExpectedReplicas  int32 `json:"expectedReplicas"`
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogstashList is a list of Logstash resources
type LogstashList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Logstash `json:"items"`
}
