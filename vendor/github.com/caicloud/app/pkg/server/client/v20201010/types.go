package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	time "time"
)

// Affinity 页面位置：高级配置 - 调度策略，不包括污点调度
type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"NodeAffinity,omitempty"`
	PodAffinity     *PodAffinity     `json:"PodAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"PodAntiAffinity,omitempty"`
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

// ConfigMapEnvSource ...
type ConfigMapEnvSource struct {
	Name string `json:"Name"`
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

// ConfigMapVolumeSource ...
type ConfigMapVolumeSource struct {
	Name string `json:"Name"`
}

// Container 页面位置：容器配置
type Container struct {
	Name  string `json:"Name"`
	Image string `json:"Image"`
	// image + version
	Command []string `json:"Command,omitempty"`
	// 页面位置：启动命令 - 运行命令
	Args []string `json:"Args,omitempty"`
	// 页面位置：启动命令 - 运行参数
	Resources ResourceRequirements `json:"Resources,omitempty"`
	// TODO: GPU usage
	Env []EnvVar `json:"Env,omitempty"`
	// custom env
	EnvFrom []EnvFromSource `json:"EnvFrom,omitempty"`
	// env from configmap or secret
	VolumeMounts  []VolumeMount `json:"VolumeMounts,omitempty"`
	Lifecycle     *Lifecycle    `json:"Lifecycle,omitempty"`
	LivenessProbe *Probe        `json:"LivenessProbe,omitempty"`
	// 页面位置：容器配置 - 健康检查 - 存活检查
	ReadinessProbe *Probe `json:"ReadinessProbe,omitempty"`
}

// ContainerState ...
type ContainerState struct {
	Waiting    *ContainerStateWaiting    `json:"Waiting,omitempty"`
	Running    *ContainerStateRunning    `json:"Running,omitempty"`
	Terminated *ContainerStateTerminated `json:"Terminated,omitempty"`
}

// ContainerStateRunning ...
type ContainerStateRunning struct {
	StartedAt time.Time `json:"StartedAt"`
}

// ContainerStateTerminated ...
type ContainerStateTerminated struct {
	ExitCode    int32     `json:"ExitCode"`
	Signal      int32     `json:"Signal"`
	Reason      string    `json:"Reason"`
	Message     string    `json:"Message"`
	StartedAt   time.Time `json:"StartedAt"`
	FinishedAt  time.Time `json:"FinishedAt"`
	ContainerID string    `json:"ContainerID"`
}

// ContainerStateWaiting ...
type ContainerStateWaiting struct {
	Reason  string `json:"Reason"`
	Message string `json:"Message"`
}

// ContainerStatus ...
type ContainerStatus struct {
	Name                 string         `json:"Name"`
	State                ContainerState `json:"State"`
	LastTerminationState ContainerState `json:"LastTerminationState,omitempty"`
	Ready                bool           `json:"Ready,omitempty"`
	RestartCount         int32          `json:"RestartCount"`
	Image                string         `json:"Image"`
	ImageID              string         `json:"ImageID"`
	ContainerID          string         `json:"ContainerID,omitempty"`
	Started              *bool          `json:"Started,omitempty"`
}

// CreateOption has some options for create API
type CreateOption struct {
	Cluster
}

// DeleteOption has some options for delete API
type DeleteOption struct {
	Cluster
}

// Deployment describes a deployment entry.
type Deployment struct {
	v1.ObjectMeta `json:",inline"`
	Network       string           `json:"Network,omitempty"`
	YAML          string           `json:"Yaml,omitempty"`
	Spec          DeploymentSpec   `json:"Spec,omitempty"`
	Status        DeploymentStatus `json:"Status,omitempty"`
}

// DeploymentDeleteOption has some options for deployment delete API
//
// +nirvana:api=origin:"DeleteOption"
type DeploymentDeleteOption struct {
	Cluster
}

// DeploymentGetOption has some options for deployment get API
//
// +nirvana:api=origin:"GetOption"
type DeploymentGetOption struct {
	Cluster
}

// DeploymentList is a list of deployment entries.
//
// +nirvana:api=origin:"List"
type DeploymentList struct {
	v1.ListMeta `json:",inline"`
	Items       []Deployment `json:"Items,omitempty"`
}

// DeploymentListOption has some options for deployment list API
//
// +nirvana:api=origin:"ListOption"
type DeploymentListOption struct {
	Pagination
	Filter
	Cluster
}

// DeploymentRestartOption has some options for deployment restart API
//
// +nirvana:api=origin:"RestartOption"
type DeploymentRestartOption struct {
	Cluster
}

// DeploymentSpec describes the attributes that a user uses to create a deployment
//
// +nirvana:api=origin:"Spec"
type DeploymentSpec struct {
	Replicas *int32 `json:"Replicas,omitempty"`
	// 页面位置：高级配置 - 实例控制
	Template TemplateSpec `json:"Template"`
	Strategy Strategy     `json:"Strategy,omitempty"`
}

// DeploymentStatus describes the status of a deployment
//
// +nirvana:api=origin:"Status"
type DeploymentStatus struct {
}

// EmptyDirVolumeSource ...
type EmptyDirVolumeSource struct {
	Medium string `json:"Medium"`
}

// EnvFromSource 页面位置： 容器配置 - 环境变量 - 配置
type EnvFromSource struct {
	ConfigMapRef *ConfigMapEnvSource `json:"ConfigMapRef,omitempty"`
	SecretRef    *SecretEnvSource    `json:"SecretRef,omitempty"`
}

// EnvVar 页面位置： 容器配置 - 环境变量 - 自定义
type EnvVar struct {
	Name  string `json:"Name"`
	Value string `json:"Value,omitempty"`
}

// ExecAction ...
type ExecAction struct {
	Command []string `json:"Command"`
}

// Filter ...
type Filter struct {
	Query string `source:"query,Query"`
}

// GetOption has some options for get API
type GetOption struct {
	Cluster
}

// HTTPGetAction ...
type HTTPGetAction struct {
	Scheme string `json:"Scheme"`
	// 协议
	Path        string       `json:"Path"`
	Port        int          `json:"Port"`
	HTTPHeaders []HTTPHeader `json:"HttpHeaders,omitempty"`
}

// HTTPHeader 页面位置：容器配置 - 健康检查 - HTTP 请求头
type HTTPHeader struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Handler ...
type Handler struct {
	Exec *ExecAction `json:"Exec,omitempty"`
	// 健康检查 - 执行命令检查
	HTTPGet *HTTPGetAction `json:"HttpGet,omitempty"`
	// 健康检查- HTTP 请求肩擦好
	TCPSocket *TCPSocketAction `json:"TcpSocket,omitempty"`
}

// HelmApp describes an application entry.
type HelmApp struct {
	v1.ObjectMeta `json:",inline"`
	Spec          HelmAppSpec   `json:"Spec"`
	Status        HelmAppStatus `json:"Status"`
}

// HelmAppList is a list of HelmApp entry
type HelmAppList struct {
	v1.ListMeta `json:",inline"`
	Items       []HelmApp `json:"Items"`
}

// HelmAppResourceConfig is the config info in application's details page.
type HelmAppResourceConfig struct {
	v1.ObjectMeta
	YAML string `json:"Yaml"`
}

// HelmAppResourceMisc includes other resource in application's details page
type HelmAppResourceMisc struct {
	v1.ObjectMeta
	YAML string `json:"Yaml"`
}

// HelmAppResourcePVC is the PVC info in application's details page.
type HelmAppResourcePVC struct {
	v1.ObjectMeta
	YAML     string `json:"Yaml"`
	Size     string `json:"Size"`
	Capacity string `json:"Capacity"`
}

// HelmAppResourceService is the service info in application's details page.
type HelmAppResourceService struct {
	v1.ObjectMeta
	YAML      string                       `json:"Yaml"`
	Type      string                       `json:"Type"`
	Ports     []HelmAppResourceServicePort `json:"Ports"`
	ClusterIP string                       `json:"ClusterIP,omitempty"`
	NodeIP    string                       `json:"NodeIP,omitempty"`
}

// HelmAppResourceServicePort is the port of service in application's details page.
type HelmAppResourceServicePort struct {
	Name     string `json:"Name,omitempty"`
	Protocol string `json:"Protocol"`
	Port     int32  `json:"Port"`
	NodePort int32  `json:"NodePort,omitempty"`
}

// HelmAppResourceWorkload is the workloads info in application's details page.
type HelmAppResourceWorkload struct {
	v1.ObjectMeta
	YAML  string `json:"Yaml"`
	Phase string `json:"Phase"`
}

// HelmAppResources is the resources created by helm application
type HelmAppResources struct {
	Workloads []HelmAppResourceWorkload `json:"Workloads"`
	Services  []HelmAppResourceService  `json:"Services"`
	Configs   []HelmAppResourceConfig   `json:"Configs"`
	Volumes   []HelmAppResourcePVC      `json:"Volumes"`
	Misc      []HelmAppResourceMisc     `json:"Misc"`
}

// HelmAppRevision describes the revision of an application.
type HelmAppRevision struct {
	v1.ObjectMeta `json:",inline"`
	Spec          HelmAppRevisionSpec   `json:"Spec"`
	Status        HelmAppRevisionStatus `json:"Status,omitempty"`
}

// HelmAppRevisionList is a list of HelmAppRevision entry
type HelmAppRevisionList struct {
	v1.ListMeta `json:",inline"`
	Items       []HelmAppRevision `json:"Items,omitempty"`
}

// HelmAppRevisionSpec describes the application revision which can not be changed.
type HelmAppRevisionSpec struct {
	ChartName    string `json:"ChartName"`
	ChartVersion string `json:"ChartVersion"`
	Values       string `json:"Values"`
	Revision     int    `json:"HelmAppRevision"`
}

// HelmAppRevisionStatus describes the application revision status.
type HelmAppRevisionStatus struct {
	UpdateTimestamp time.Time `json:"UpdateTimestamp"`
}

// HelmAppSpec describes the application spec
type HelmAppSpec struct {
	ChartName     string `json:"ChartName"`
	ChartVersion  string `json:"ChartVersion"`
	IsCustomChart bool   `json:"IsCustomChart"`
	Values        string `json:"Values"`
	Network       string `json:"Network"`
}

// HelmAppStatus describes the application status
type HelmAppStatus struct {
	Phase           string            `json:"Phase"`
	UpdateTimestamp time.Time         `json:"UpdateTimestamp"`
	Version         int               `json:"Version"`
	Resources       *HelmAppResources `json:"Resources,omitempty"`
}

// HostAlias 页面位置：高级配置 - DNS 配置 - Hosts 文件配置
type HostAlias struct {
	IP        string   `json:"Ip"`
	HostNames []string `json:"HostNames"`
}

// KV represents a single key-value pair in the form of a struct
type KV struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// Lifecycle 页面位置：容器配置 - 生命周期
type Lifecycle struct {
	PostStart []string `json:"PostStart,omitempty"`
	PreStop   []string `json:"PreStop,omitempty"`
}

// ListOption has some options for list API
type ListOption struct {
	Pagination
	Filter
	Cluster
}

// Metadata 使用在 PodTemplate 中的元数据
// 不直接使用 `cpsmetav1.ObjectMeta`，因为包含了太多无用字段
type Metadata struct {
	Labels      []KV `json:"Labels,omitempty"`
	Annotations []KV `json:"Annotations,omitempty"`
}

// NodeAffinity 页面位置：高级配置 - 调度策略 - 节点亲和性
type NodeAffinity struct {
	Type string `json:"Type"`
	// one of Required and Preferred
	Labels []KV `json:"Labels"`
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

// PersistentVolumeClaim ...
type PersistentVolumeClaim struct {
	v1.ObjectMeta `json:",inline"`
	Spec          PersistentVolumeClaimSpec `json:"spec,omitempty"`
}

// PersistentVolumeClaimSpec ...
type PersistentVolumeClaimSpec struct {
	StorageClassName *string `json:"StorageClassName,omitempty"`
	// 存储方案（使用 storageclass 的 name）
	Resources ResourceRequirements `json:"Resources,omitempty"`
}

// PersistentVolumeClaimVolumeSource ...
type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"ClaimName"`
}

// Pod describes a pod entry
type Pod struct {
	v1.ObjectMeta `json:",inline"`
	Spec          PodSpec   `json:"Spec,omitempty"`
	Status        PodStatus `json:"Status,omitempty"`
}

// PodAffinity 页面位置：高级配置 - 调度策略 - Pod 亲和性
// nolint
type PodAffinity struct {
	Type   string `json:"Type"`
	Labels []KV   `json:"Labels"`
}

// PodAntiAffinity 页面位置：高级配置 - 调度策略 - Pod 反亲和性
// nolint
type PodAntiAffinity struct {
	Type   string `json:"Type"`
	Labels []KV   `json:"Labels"`
}

// PodList ...
//
// +nirvana:api=origin:"List"
type PodList struct {
	v1.ListMeta `json:",inline"`
	Items       []Pod `json:"Items,omitempty"`
}

// PodSpec ...
//
// +nirvana:api=origin:"Spec"
type PodSpec struct {
	InitContainers                []Container `json:"InitContainers,omitempty"`
	Containers                    []Container `json:"Containers"`
	Volumes                       []Volume    `json:"Volumes,omitempty"`
	TerminationGracePeriodSeconds *int64      `json:"TerminationGracePeriodSeconds,omitempty"`
	// 页面位置：高级配置 - 实例控制
	DNSPolicy       string           `json:"DnsPolicy,omitempty"`
	Affinity        *Affinity        `json:"Affinity,omitempty"`
	Tolerations     []Toleration     `json:"Tolerations,omitempty"`
	HostAliases     []HostAlias      `json:"HostAliases,omitempty"`
	SecurityContext *SecurityContext `json:"PodSecurityContext,omitempty"`
}

// PodStatus ...
//
// +nirvana:api=origin:"Status"
type PodStatus struct {
	State                 State             `json:"State"`
	Phase                 string            `json:"Phase"`
	Reason                string            `json:"Reason,omitempty"`
	Message               string            `json:"Message,omitempty"`
	HostIP                string            `json:"HostIP,omitempty"`
	PodIPs                []string          `json:"PodIPs,omitempty"`
	StartTime             *time.Time        `json:"StartTime,omitempty"`
	InitContainerStatuses []ContainerStatus `json:"InitContainerStatuses,omitempty"`
	ContainerStatuses     []ContainerStatus `json:"ContainerStatuses,omitempty"`
}

// Port represents the port on which the service is exposed
type Port struct {
	Name     string `json:"Name,omitempty"`
	Protocol string `json:"Protocol"`
	Port     int32  `json:"Port"`
	NodePort int32  `json:"NodePort,omitempty"`
}

// Probe 页面位置：容器配置 - 健康检查
type Probe struct {
	Handler             `json:",inline"`
	InitialDelaySeconds int32 `json:"InitialDelaySeconds"`
	// 时间设置 - 初始等待
	TimeoutSeconds int32 `json:"TimeoutSeconds"`
	// 时间设置 - 超时
	PeriodSeconds int32 `json:"PeriodSeconds"`
	// 时间设置 - 检查间隔
	SuccessThreshold int32 `json:"SuccessThreshold"`
	// 阈值 - 成功
	FailureThreshold int32 `json:"FailureThreshold"`
}

// Remarks ...
type Remarks struct {
	SubjectName            string   `json:"SubjectName,omitempty"`
	IssuerName             []string `json:"IssuerName,omitempty"`
	NotBefore              string   `json:"NotBefore,omitempty"`
	NotAfter               string   `json:"NotAfter,omitempty"`
	SubjectAlternativeName []string `json:"SubjectAlternativeName,omitempty"`
}

// ResourceRequirements 页面位置： 容器配置 - 容器基本信息 - 资源配额
type ResourceRequirements struct {
	Limits   []KV `json:"Limits"`
	Requests []KV `json:"Requests"`
}

// RollbackHelmAppToRevisionOption has some options for RollbackHelmAppToRevision API
//
// +nirvana:api=alias:"RollbackHelmAppToRevisionOption"
type RollbackHelmAppToRevisionOption struct {
	Cluster
	Revision int `source:"query,Revision"`
}

// RollingUpdateDeployment 页面位置： 高级配置 - 更新策略（滚动更新） - 最大不可用 & 最大超量
type RollingUpdateDeployment struct {
	MaxUnavailable int `json:"MaxUnavailable,omitempty"`
	MaxSurge       int `json:"MaxSurge,omitempty"`
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
	Remarks    Remarks           `json:"Remarks,omitempty"`
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

// SecretEnvSource ...
type SecretEnvSource struct {
	Name string `json:"Name"`
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

// SecretVolumeSource ...
type SecretVolumeSource struct {
	SecretName string `json:"SecretName"`
}

// SecurityContext 页面位置：高级配置 - 安全
type SecurityContext struct {
	RunAsNonRoot *bool `json:"RunAsNonRoot,omitempty"`
}

// Service describes a service entry
type Service struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ServiceSpec        `json:"Spec,omitempty"`
	YAML          string             `json:"Yaml,omitempty"`
	Workloads     []*ServiceWorkload `json:"Workloads,omitempty"`
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

// ServiceLabelSelector describes a kv pair.
//
// +nirvana:api=origin:"LabelSelector"
type ServiceLabelSelector struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
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
	Selector        []ServiceLabelSelector `json:"Selector,omitempty"`
	Type            string                 `json:"Type"`
	Ports           []Port                 `json:"Ports"`
	ClusterIP       string                 `json:"ClusterIP,omitempty"`
	NodeIP          string                 `json:"NodeIP,omitempty"`
	SessionAffinity *SessionAffinity       `json:"SessionAffinity,omitempty"`
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

// State is the state of Pod.
type State string

// StatefulSet describes a statefulset entry.
type StatefulSet struct {
	v1.ObjectMeta `json:",inline"`
	Network       string            `json:"Network,omitempty"`
	YAML          string            `json:"Yaml,omitempty"`
	Spec          StatefulSetSpec   `json:"Spec,omitempty"`
	Status        StatefulSetStatus `json:"Status,omitempty"`
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

// StatefulSetRestartOption has some options for statefulset restart API
//
// +nirvana:api=origin:"RestartOption"
type StatefulSetRestartOption struct {
	Cluster
}

// StatefulSetSpec describes the attributes that a user uses to create a statefulset
//
// +nirvana:api=origin:"Spec"
type StatefulSetSpec struct {
	Replicas             *int32                  `json:"Replicas,omitempty"`
	Template             TemplateSpec            `json:"Template"`
	VolumeClaimTemplates []PersistentVolumeClaim `json:"VolumeClaimTemplates,omitempty"`
	ServiceName          string                  `json:"ServiceName"`
	// the service must be pre-defined
	UpdateStrategy UpdateStrategy `json:"UpdateStrategy,omitempty"`
}

// StatefulSetStatus ...
//
// +nirvana:api=origin:"Status"
type StatefulSetStatus struct {
}

// Strategy 页面位置： 高级配置 - 更新策略
type Strategy struct {
	Type          string                   `json:"Type,omitempty"`
	RollingUpdate *RollingUpdateDeployment `json:"RollingUpdate,omitempty"`
}

// TCPSocketAction ...
type TCPSocketAction struct {
	Port int    `json:"Port"`
	Host string `json:"Host,omitempty"`
}

// TemplateSpec ...
type TemplateSpec struct {
	Metadata Metadata `json:"Metadata,omitempty"`
	Spec     PodSpec  `json:"Spec,omitempty"`
}

// Toleration 页面位置：高级配置 - 调度策略 - 节点污染调度
type Toleration struct {
	Key    string `json:"Key"`
	Value  string `json:"Value"`
	Effect string `json:"Effect"`
}

// UpdateOption has some options for update API
type UpdateOption struct {
	Cluster
}

// UpdateStrategy 页面位置：高级配置 - 更新策略（目前只有滚动更新，无需配置）
type UpdateStrategy struct {
	Type string `json:"Type,omitempty"`
}

// Volume 页面位置：存储配置
type Volume struct {
	Name         string `json:"Name"`
	VolumeSource `json:",inline"`
}

// VolumeMount 页面位置：容器配置 - 配置 - 挂载配置 / 存储配置
type VolumeMount struct {
	Name string `json:"Name"`
	// name of volume mount or related volume
	ReadOnly bool `json:"ReadOnly"`
	// should be true
	MountPath string `json:"MountPath"`
	// suffix should be subpath
	SubPath    string `json:"SubPath"`
	SourceName string `json:"SourceName"`
	// name of configmap or secret
	Source string `json:"Source"`
}

// VolumeSource ...
type VolumeSource struct {
	EmptyDir              *EmptyDirVolumeSource              `json:"EmptyDir,omitempty"`
	Secret                *SecretVolumeSource                `json:"Secret,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"PersistentVolumeClaim,omitempty"`
	ConfigMap             *ConfigMapVolumeSource             `json:"ConfigMap,omitempty"`
}

// YAML describe a yaml resource
type YAML struct {
	v1.ObjectMeta `json:",inline"`
	Spec          YamlSpec `json:"Spec"`
}

// YamlCreateOption has some options for yaml create API
//
// +nirvana:api=origin:"CreateOption"
type YamlCreateOption struct {
	Cluster
	Network string `source:"query,Network"`
}

// YamlSpec describe yaml content
//
// +nirvana:api=origin:"Spec"
type YamlSpec struct {
	Content string `json:"Content"`
}
