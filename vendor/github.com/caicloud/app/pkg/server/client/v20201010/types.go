package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	time "time"
)

// Application describes an application entry.
type Application struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ApplicationSpec   `json:"spec"`
	Status        ApplicationStatus `json:"status"`
}

// ApplicationList is a list of Application entry
type ApplicationList struct {
	v1.ListMeta `json:",inline"`
	Items       []Application `json:"items,omitempty"`
}

// ApplicationRevision describes the revision of an application.
type ApplicationRevision struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ApplicationRevisionSpec   `json:"spec,omitempty"`
	Status        ApplicationRevisionStatus `json:"status,omitempty"`
}

// ApplicationRevisionList is a list of ApplicationRevision entry
type ApplicationRevisionList struct {
	v1.ListMeta `json:",inline"`
	Items       []ApplicationRevision `json:"items,omitempty"`
}

// ApplicationRevisionSpec describes the application revision infos those can not be changed.
type ApplicationRevisionSpec struct {
	Revision     int    `json:"revision"`
	Repo         string `json:"repo"`
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion,omitempty"`
	Values       string `json:"values,omitempty"`
}

// ApplicationRevisionStatus describes the application revision
type ApplicationRevisionStatus struct {
}

// ApplicationSpec describes the application spec
type ApplicationSpec struct {
	Repo         string `json:"repo"`
	ChartName    string `json:"chartName"`
	ChartVersion string `json:"chartVersion,omitempty"`
	Values       string `json:"values,omitempty"`
}

// ApplicationStatus describes the application status
type ApplicationStatus struct {
	Phase           string    `json:"phase"`
	UpdateTimestamp time.Time `json:"updateTimestamp,omitempty"`
}

// Cluster ...
type Cluster struct {
	ClusterName string `source:"query,ClusterName"`
	Namespace   string `source:"query,Namespace"`
	Name        string `source:"query,Name"`
}

// ConfigMap describes a configmap entry.
type ConfigMap struct {
	v1.ObjectMeta `json:",inline"`
	Type          string `json:"Type,omitempty"`
	// KV || FILE
	Data       []ConfigMapData      `json:"Data,omitempty"`
	YAML       string               `json:"Yaml,omitempty"`
	References []ConfigMapReference `json:"References,omitempty"`
}

// ConfigMapData describes a kv pair.
//
// +nirvana:api=origin:"Data"
type ConfigMapData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// ConfigMapDeleteOption has some options for configmap delete API
//
// +nirvana:api=origin:"DeleteOption"
type ConfigMapDeleteOption struct {
	Cluster
}

// ConfigMapGetOption has some options for configmap get API
//
// +nirvana:api=origin:"GetOption"
type ConfigMapGetOption struct {
	Cluster
}

// ConfigMapList is a list of configmap entries.
//
// +nirvana:api=origin:"List"
type ConfigMapList struct {
	v1.ListMeta `json:",inline"`
	Items       []ConfigMap `json:"Items,omitempty"`
}

// ConfigMapListOption has some options for configmap list API
//
// +nirvana:api=origin:"ListOption"
type ConfigMapListOption struct {
	Pagination
	Filter
	Cluster
}

// ConfigMapReference provides a workload's minimum info
//
// +nirvana:api=origin:"Reference"
type ConfigMapReference struct {
	Name string `json:"Name"`
	// workload name
	Kind string `json:"Kind"`
}

// Exporter is the options when export native resource
type Exporter struct {
	ContentType string `source:"query,ContentType"`
}

// Filter ...
type Filter struct {
	Query string `source:"query,Query"`
}

// GetWorkloadOption has some options for get workload API
type GetWorkloadOption struct {
	Pagination
	Cluster
}

// Pagination ...
type Pagination struct {
	Start uint `source:"query,Start,default=0"`
	Limit uint `source:"query,Limit,default=99999"`
}

// Port represents the port on which the service is exposed
type Port struct {
	Protocol string `json:"Protocol"`
	Port     int32  `json:"Port"`
	NodePort int32  `json:"NodePort,omitempty"`
}

// Secret describes a secret entry.
type Secret struct {
	v1.ObjectMeta `json:",inline"`
	Data          []SecretData      `json:"Data,omitempty"`
	YAML          string            `json:"YAML,omitempty"`
	References    []SecretReference `json:"References,omitempty"`
}

// SecretData describes a kv pair.
//
// +nirvana:api=origin:"Data"
type SecretData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// SecretDeleteOption has some options for secret delete API
//
// +nirvana:api=origin:"DeleteOption"
type SecretDeleteOption struct {
	Cluster
}

// SecretGetOption has some options for secret get API
//
// +nirvana:api=origin:"GetOption"
type SecretGetOption struct {
	Cluster
	Style string `source:"query,Style"`
}

// SecretList is a list of secret entries.
//
// +nirvana:api=origin:"List"
type SecretList struct {
	v1.ListMeta `json:",inline"`
	Items       []Secret `json:"Items,omitempty"`
}

// SecretListOption has some options for secret list API
//
// +nirvana:api=origin:"ListOption"
type SecretListOption struct {
	Pagination
	Filter
	Cluster
	Style string `source:"query,Style"`
}

// SecretReference provides a workload's minimum info
//
// +nirvana:api=origin:"Reference"
type SecretReference struct {
	Name string `json:"Name"`
	// workload name
	Kind string `json:"Kind"`
}

// Service describes a service entry
type Service struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ServiceSpec   `json:"Spec,omitempty"`
	YAML          []uint8       `json:"YAML,omitempty"`
	Workloads     *WorkloadList `json:"WorkloadList,omitempty"`
}

// ServiceDeleteOption has some options for service delete API
//
// +nirvana:api=origin:"DeleteOption"
type ServiceDeleteOption struct {
	Cluster
}

// ServiceGetOption has some options for service get API
//
// +nirvana:api=origin:"GetOption"
type ServiceGetOption struct {
	Cluster
	Exporter
}

// ServiceList is a list of Service entry
//
// +nirvana:api=origin:"List"
type ServiceList struct {
	v1.ListMeta `json:",inline"`
	Items       []Service `json:"items,omitempty"`
}

// ServiceListOption has some options for service list API
//
// +nirvana:api=origin:"ListOption"
type ServiceListOption struct {
	Pagination
	Filter
	Cluster
}

// ServiceSpec describes the attributes that a user creates on a service
//
// +nirvana:api=origin:"Spec"
type ServiceSpec struct {
	Selector        map[string]string `json:"Selector,omitempty"`
	Type            string            `json:"Type"`
	Ports           []Port            `json:"Ports"`
	ClusterIP       string            `json:"ClusterIP,omitempty"`
	ExternalName    string            `json:"ExternalName,omitempty"`
	SessionAffinity *SessionAffinity  `json:"SessionAffinity,omitempty"`
}

// ServiceWorkload contains the basic info of a workload
//
// +nirvana:api=origin:"Workload"
type ServiceWorkload struct {
	v1.ObjectMeta   `json:",inline"`
	ApplicationName string `json:"ApplicationName,omitempty"`
}

// SessionAffinity contains the configurations of session affinity
type SessionAffinity struct {
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty"`
}

// StatefulSet describes a statefulset entry.
type StatefulSet struct {
	v1.ObjectMeta `json:",inline"`
	Spec          StatefulSetSpec `json:"Spec,omitempty"`
	YAML          []uint8         `json:"YAML,omitempty"`
}

// StatefulSetDeleteOption has some options for statefulset delete API
//
// +nirvana:api=origin:"DeleteOption"
type StatefulSetDeleteOption struct {
	Cluster
}

// StatefulSetGetOption has some options for statefulset get API
//
// +nirvana:api=origin:"GetOption"
type StatefulSetGetOption struct {
	Cluster
	Style string `source:"query,Style"`
}

// StatefulSetList is a list of configmap entries.
//
// +nirvana:api=origin:"List"
type StatefulSetList struct {
	v1.ListMeta `json:",inline"`
	Items       []StatefulSet `json:"Items,omitempty"`
}

// StatefulSetListOption has some options for statefulset list API
//
// +nirvana:api=origin:"ListOption"
type StatefulSetListOption struct {
	Pagination
	Filter
	Cluster
	Style string `source:"query,Style"`
}

// StatefulSetRestartOption has some options for statefulset delete API
//
// +nirvana:api=origin:"RestartOption"
type StatefulSetRestartOption struct {
	Cluster
}

// StatefulSetSpec describes the attributes that a user uses to create a statefulset
//
// +nirvana:api=origin:"Spec"
type StatefulSetSpec struct {
	// TODO: 容器网络
	Replicas *int32
}

// Workload describes a workload entry.
type Workload struct {
	Name string `json:"name"`
}

// WorkloadList is a list of Workload entry
type WorkloadList struct {
	v1.ListMeta `json:",inline"`
	Items       []ServiceWorkload `json:"items,omitempty"`
}
