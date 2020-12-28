package v20201010

import (
	"github.com/caicloud/api/meta/v1"
	time "time"
)

// Application describes an application entry.
type Application struct {
	v1.ObjectMeta `json:",inline"`
	Spec          Spec   `json:"Spec"`
	Status        Status `json:"Status"`
}

// ApplicationDeleteOption has some options for application delete API
//
// +nirvana:api=origin:"DeleteOption"
type ApplicationDeleteOption struct {
	Cluster
}

// ApplicationGetOption has some options for application get API
//
// +nirvana:api=origin:"GetOption"
type ApplicationGetOption struct {
	Cluster
}

// ApplicationListOption has some options for application list API
//
// +nirvana:api=origin:"ListOption"
type ApplicationListOption struct {
	Pagination
	Filter
	Cluster
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

// Deployment describes a deployment entry.
type Deployment struct {
	v1.ObjectMeta `json:",inline"`
	Spec          DeploymentSpec `json:"Spec,omitempty"`
	YAML          string         `json:"Yaml,omitempty"`
}

// DeploymentList is a list of deployment entries.
//
// +nirvana:api=origin:"List"
type DeploymentList struct {
	v1.ListMeta `json:",inline"`
	Items       []Deployment `json:"Items,omitempty"`
}

// DeploymentSpec describes the attributes that a user uses to create a deployment
//
// +nirvana:api=origin:"Spec"
type DeploymentSpec struct {
	// TODO: 容器网络
	Replicas *int32
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

// List is a list of Application entry
type List struct {
	v1.ListMeta `json:",inline"`
	Items       []Application `json:"Items"`
}

// Overview contains the workload info in overview page.
type Overview struct {
	v1.ObjectMeta `json:",inline"`
	Status        *OverviewStatus `json:"Status,omitempty"`
}

// OverviewStatus represents the simple status of all workloads.
type OverviewStatus struct {
	Total    int `json:"Total"`
	Running  int `json:"Running"`
	Updating int `json:"Updating"`
	Error    int `json:"Error"`
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

// Revision describes the revision of an application.
type Revision struct {
	v1.ObjectMeta `json:",inline"`
	Spec          RevisionSpec   `json:"Spec"`
	Status        RevisionStatus `json:"Status,omitempty"`
}

// RevisionList is a list of Revision entry
type RevisionList struct {
	v1.ListMeta `json:",inline"`
	Items       []Revision `json:"items,omitempty"`
}

// RevisionSpec describes the application revision which can not be changed.
type RevisionSpec struct {
	Revision     int    `json:"Revision"`
	Repo         string `json:"Repo"`
	ChartName    string `json:"ChartName"`
	ChartVersion string `json:"ChartVersion"`
	Values       string `json:"Values"`
}

// RevisionStatus describes the application revision status.
type RevisionStatus struct {
}

// Secret describes a secret entry.
type Secret struct {
	v1.ObjectMeta `json:",inline"`
	Type          string `json:"Type,omitempty"`
	// KV || FILE
	Encryption string `json:"Encryption,omitempty"`
	// Kubernetes Secret types (e.g., Opaque, kubernetes.io/tls)
	Data       []SecretData      `json:"Data,omitempty"`
	YAML       string            `json:"Yaml,omitempty"`
	References []SecretReference `json:"References,omitempty"`
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
	YAML          string        `json:"Yaml,omitempty"`
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

// Spec describes the application spec
type Spec struct {
	Repo         string `json:"Repo"`
	ChartName    string `json:"ChartName"`
	ChartVersion string `json:"ChartVersion"`
	Values       string `json:"Values"`
}

// StatefulSet describes a statefulset entry.
type StatefulSet struct {
	v1.ObjectMeta `json:",inline"`
	Spec          StatefulSetSpec `json:"Spec,omitempty"`
	YAML          string          `json:"Yaml,omitempty"`
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
}

// StatefulSetList is a list of statefulset entries.
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

// Status describes the application status
type Status struct {
	Phase           string    `json:"Phase"`
	UpdateTimestamp time.Time `json:"UpdateTimestamp"`
}

// WorkloadList is a list of Workload entry
type WorkloadList struct {
	v1.ListMeta `json:",inline"`
	Items       []ServiceWorkload `json:"items,omitempty"`
}
