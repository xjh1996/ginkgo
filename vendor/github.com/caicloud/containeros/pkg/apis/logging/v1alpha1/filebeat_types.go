/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	custom_metav1 "github.com/caicloud/api/meta/v1"
	beatv1b1 "github.com/elastic/cloud-on-k8s/pkg/apis/beat/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/caicloud/containeros/pkg/apis/logging"
	crdutils "github.com/caicloud/containeros/pkg/utils/crd"
)

const (
	// FilebeatClusterSetsKind is the name of the FilebeatClusterSetsKind resource kind.
	FilebeatClusterSetsKind = "FilebeatClusterSet"
	// FilebeatClusterSetName is the name of the FilebeatClusterSetsKind resource (plural).
	FilebeatClusterSetName = "filebeatclustersets"
	// FilebeatClusterSetKindKey is used as the key when mapping to the FilebeatClusterSet resource.
	FilebeatClusterSetKindKey = "filebeatclusterset"
)

// FilebeatClusterSetCRDChartFilename is the filename of the FilebeatClusterSetName CRD chart.
var FilebeatClusterSetCRDChartFilename = crdutils.GenerateCRDChartFilename(logging.GroupName, FilebeatClusterSetName)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName="fbcs"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.template.spec.version",description="The version of Filebeat"
// +kubebuilder:printcolumn:name="Clusters",type="integer",JSONPath=".status.expectedClusters",description="The expected number of clusters to deploy Filebeates"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// FilebeatClusterSet describes a multi-cluster deployment of Filebeat instances. Note that
// FilebeatClusterSet only controls the Filebeat resource from elastic-operator; it does
// not see or care about the resources created by elastic-operator.
type FilebeatClusterSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec defines the specifications of the desired Filebeat resource and how to deploy it
	// across clusters.
	Spec FilebeatClusterSetSpec `json:"spec"`
	// Status is the most recent observation of the deployment status of Filebeat resource
	// across the clusters.
	Status FilebeatClusterSetStatus `json:"status,omitempty"`
}

// FilebeatClusterSetSpec defines the specifications of a FilebeatClusterSet object.
type FilebeatClusterSetSpec struct {
	// ClusterSelector selects Cluster resources; Filebeat resources will be deployed onto
	// the clusters corresponding to these Cluster resources. If left empty, nothings will
	// be selected.
	// +optional
	ClusterSelector *metav1.LabelSelector `json:"clusterSelector,omitempty"`
	// InputsSelector selects Kubernetes Secrets in the same namespace as the FilebeatClusterSet object
	// , which holding the filebeat input configuration; all Secrets will merge into single Kubernetes Secrets
	// in the same Cluster and namespace as the Filebeat object.
	// If left empty, nothing whill be selected.
	//
	// The referenced secret should contain the following:
	//
	// - `filebeat.inputs.yaml`: input configurations for filebeat. Ref: https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-configuration-reloading.html
	//
	// Example:
	// ---
	// kind: Secret
	// apiVersion: v1
	// metadata:
	//	 name: my-roles
	// stringData:
	//   filebeat.inputs.yaml: |-
	//     - type: log
	//       paths:
	//       - /var/log/mysql.log
	//       scan_frequency: 10s
	//
	//     - type: log
	//       paths:
	//       - /var/log/apache.log
	//       scan_frequency: 5s
	// ---
	//
	// +optional
	InputsSelector *metav1.LabelSelector `json:"inputsSelector,omitempty"`
	// Settings  is a list of Secrets in the same namespace as the FilebeatClusterSet
	// object, which holding the filebeat configuration; all configuration will merge into single Kubernetes Secret
	// in the same Cluster and namespace as the Filebeat object.
	// this Secret shall be auto append to template.spec.secureSettings
	//
	// The referenced secret should contain the following:
	//
	// - `filebeat.yaml`: configuration for filebeat. Ref: https://www.elastic.co/guide/en/beats/filebeat/current/configuring-howto-filebeat.html
	//
	// +optional
	Settings []ConfigSource `json:"settings,omitempty"`
	// ElasticsearchEndpointRef contains a reference to an existing Elasticsearch Cluster.
	// It allows automatioc setup of Index Template and ILM.
	// +optional
	ElasticsearchEndpointRef *ElasticsearchEndpointRef `json:"elasticsearchEndpointRef,omitempty"`
	// Template defines the Filebeat resources that will be created across clusters.
	// +optional
	Template *FilebeatTemplate `json:"template,omitempty"`
}

// ElasticsearchEndpointRef is a references to ElasticsearchEndpoint.
// It allows automatioc setup of Index Template and ILM.
type ElasticsearchEndpointRef struct {
	// Name is the name of a ElasticsearchEndpointRef object.
	Name string `json:"name"`
}

// FilebeatTemplate describes the Filebeat resources that will be created.
type FilebeatTemplate struct {
	// Metadata describes the metadata of the Filebeat resources that will be created.
	// +optional
	Metadata *custom_metav1.EmbeddedObjectMeta `json:"metadata,omitempty"`
	// Spec describes the specification of the Filebeat resources that will be created.
	Spec *beatv1b1.BeatSpec `json:"spec,omitempty"`
}

// FilebeatClusterSetStatus describes the status of a FilebeatClusterSet object.
type FilebeatClusterSetStatus struct {
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions       []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	ExpectedClusters int                `json:"expectedClusters,omitempty"`
	UpdatedClusters  int                `json:"updatedClusters,omitempty"`
	// +optional
	// +patchMergeKey=cluster
	// +patchStrategy=merge
	FilebeatStatus []ClusterStatus `json:"filebeatStatus,omitempty" patchStrategy:"merge" patchMergeKey:"cluster"`
}

// FilebeatClusterSetConditionType is the type of the condition.
type FilebeatClusterSetConditionType = string

const (
	// FilebeatClusterSetIndexSettingsReady indicates whether all Index`s Settings are Ready
	FilebeatClusterSetIndexSettingsReady FilebeatClusterSetConditionType = "IndexSettingsReady"
	// FilebeatClusterSetConfigurationReady indicates whether all Filebeat`s Configuration are Ready
	FilebeatClusterSetConfigurationReady FilebeatClusterSetConditionType = "ConfigurationReady"
	// FilebeatClusterSetFilebeatsUpdated indicates whether all Filebeats on the Cluster are Updated
	FilebeatClusterSetFilebeatsUpdated FilebeatClusterSetConditionType = "FilebeatsUpdated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FilebeatClusterSetList is a list of FilebeatClusterSet resources
type FilebeatClusterSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []FilebeatClusterSet `json:"items"`
}
