/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/caicloud/containeros/pkg/apis/logging"
	crdutils "github.com/caicloud/containeros/pkg/utils/crd"
)

const (
	// ElasticsearchEndpointsKind is the name of the ElasticsearchEndpointKind resource kind.
	ElasticsearchEndpointsKind = "ElasticsearchEndpoint"
	// ElasticsearchEndpointName is the name of the ElasticsearchEndpointKind resource (plural).
	ElasticsearchEndpointName = "elasticsearchendpoints"
	// ElasticsearchEndpointKindKey is used as the key when mapping to the ElasticsearchEndpoint resource.
	ElasticsearchEndpointKindKey = "elasticsearchendpoint"
)

// ElasticsearchEndpointCRDChartFilename is the filename of the ElasticsearchEndpointName CRD chart.
var ElasticsearchEndpointCRDChartFilename = crdutils.GenerateCRDChartFilename(logging.GroupName, ElasticsearchEndpointName)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName="esendpoint"
// +kubebuilder:printcolumn:name="IndexTemplate",type="integer",JSONPath=".status.expectedIndexTemplate",description="The expected number of Index Templates to be ensure"
// +kubebuilder:printcolumn:name="ILM",type="integer",JSONPath=".status.expectedIndexLifecycleManagement",description="The expected number of Index Lifecycle Managements to be ensure"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// ElasticsearchEndpoint declares a Elasticsearch cluster connection.
type ElasticsearchEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of a Elasticsearch cluster.
	Spec ElasticsearchEndpointSpec `json:"spec"`
	// Status defines the observed state of a Elasticsearch cluster.
	Status ElasticsearchEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum=http;https

// ElasticsearchEndpointProtocol declares Elasticsearch cluster connection's protocol.
type ElasticsearchEndpointProtocol string

const (
	// ElasticsearchEndpointHTTP means using Http Protocol
	ElasticsearchEndpointHTTP ElasticsearchEndpointProtocol = "http"
	// ElasticsearchEndpointHTTPS means using Https Protocol
	ElasticsearchEndpointHTTPS ElasticsearchEndpointProtocol = "https"
)

// ElasticsearchEndpointSpec defines the specifications of a ElasticsearchEndpoint object.
type ElasticsearchEndpointSpec struct {
	// IndexTemplateSelector selects selects a list of Kubernetes Secrets in the same namespace as the ElasticsearchEndpoint object
	// , which holding the index template object;
	//
	// The referenced secret should contain the following:
	//
	// - `template`: index template configuration in JSON format. Ref: https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
	//
	// +optional
	IndexTemplateSelector *metav1.LabelSelector `json:"indexTemplateSelector,omitempty"`
	// IndexLifecycleManagementSelector selects a list of Kubernetes Secrets in the same namespace as the ElasticsearchEndpoint object
	// , which holding the index lifecycle management object;
	//
	// The referenced secret should contain the following:
	//
	// - `ilm`: ilm configuration in JSON format. Ref: https://www.elastic.co/guide/en/elasticsearch/reference/current/overview-index-lifecycle-management.html
	//
	// +optional
	IndexLifecycleManagementSelector *metav1.LabelSelector `json:"indexLifecycleManagementSelector,omitempty"`
	// Hosts is a list of Elasticsearch nodes to connect to.
	Hosts []string `json:"hosts"`
	// Protocol is a name of the protocol Elasticsearch is reachable on.
	Protocol ElasticsearchEndpointProtocol `json:"protocol,omitempty"`
	// Path is an HTTP path prefix that is prepended to the HTTP API calls.
	// +optional
	Path string `json:"path,omitempty"`
	// AuthConfigRef contains a reference to an existing Kubernetes Secret holding the auth configuration.
	// The referenced secret should contain the following:
	//
	// - `username`: The basic authentication username for connecting to Elasticsearch.
	// - `password`: The basic authentication password for connecting to Elasticsearch.
	// - `api_key`: Instead of using a username and password, you can use API keys to secure communication with Elasticsearch. The value must be the ID of the API key and the API key joined by a colon: id:api_key.
	// - `certificate`: The SSL client authentication in PEM format.
	// - `key`: The client certificate key in PEM format.
	// +optional
	AuthConfigRef *SecretRef `json:"authConfigRef,omitempty"`
	// Paused can be used to stop the controller from reacting to this ElasitcsearchEndpoint.
	// +optional
	Paused bool `json:"paused,omitempty"`
}

// ElasticsearchEndpointStatus describes the status of a ElasticsearchEndpoint object.
type ElasticsearchEndpointStatus struct {
	ExpectedIndexTemplate            int                `json:"expectedIndexTemplate"`
	EnsuredIndexTemplate             int                `json:"ensuredIndexTemplate"`
	ExpectedIndexLifecycleManagement int                `json:"expectedIndexLifecycleManagement"`
	EnsuredIndexLifecycleManagement  int                `json:"ensuredIndexLifecycleManagement"`
	Conditions                       []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:validation:Enum=IndexTemplatesReady;IndexLifecycleManagementsReady

// ElasticsearchEndpointConditionType is the type of the condition.
type ElasticsearchEndpointConditionType = string

const (
	// ElasticsearchEndpointIndexTemplatesReady indicates whether all Index`s Template are Ready
	ElasticsearchEndpointIndexTemplatesReady ElasticsearchEndpointConditionType = "IndexTemplatesReady"
	// ElasticsearchEndpointIndexLifecycleManagementsReady indicates whether all Index Lifecycle Management are Ready
	ElasticsearchEndpointIndexLifecycleManagementsReady ElasticsearchEndpointConditionType = "IndexLifecycleManagementsReady"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticsearchEndpointList is a list of ElasticsearchEndpoint resources
type ElasticsearchEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ElasticsearchEndpoint `json:"items"`
}
