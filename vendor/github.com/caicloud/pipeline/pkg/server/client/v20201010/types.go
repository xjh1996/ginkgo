package v20201010

import (
	v12 "github.com/caicloud/api/meta/v1"
	v1 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	time "time"
)

// AccelerationCacheCleanupOverallStatus holds latest succeeded cache cleanup time and latest cleanup status of acceleration.
type AccelerationCacheCleanupOverallStatus struct {
	LatestSucceededTimestamp v11.Time                       `json:"LatestSucceededTimestamp"`
	LatestStatus             AccelerationCacheCleanupStatus `json:"LatestStatus"`
}

// AccelerationCacheCleanupStatus ...
type AccelerationCacheCleanupStatus struct {
	// Cyclone will launch a pod to cleanup acceleration cache, and the pod name will be used as TaskID.
	TaskID             string            `json:"TaskId"`
	Phase              CacheCleanupPhase `json:"Phase"`
	StartTime          v11.Time          `json:"StartTime"`
	LastTransitionTime v11.Time          `json:"LastTransitionTime"`
	// Reason holds information of why the task failed.
	Reason string `json:"Reason"`
}

// Actor holds information about actor.
type Actor struct {
	Name string `json:"Name"`
}

// ArgumentValue defines a argument value
type ArgumentValue struct {
	// Name of the parameter
	Name string `json:"Name"`
	// Value of the parameter
	Value *string `json:"Value"`
	// Description of the parameter
	Description string `json:"Description,omitempty"`
	// Required indicates whether this parameter is required
	Required bool `json:"Required,omitempty"`
}

// CICDSummary summary of CICD.
type CICDSummary struct {
	WorkspaceNum    int              `json:"WorkspaceNum"`
	PipelineNum     int              `json:"PipelineNum"`
	WorkspacesStats []WorkspaceStats `json:"WorkspacesStats"`
}

// CacheCleanupPhase defines phases of cache cleanup
type CacheCleanupPhase string

// Cargo contains information about cargo registry
type Cargo struct {
	// Name is the name of the cargo
	Name string `json:"Name"`
	// Project is the registry project
	Project string `json:"Project"`
}

// CargoNotification holds all events.
type CargoNotification struct {
	Events []Event `json:"Events"`
}

// Code describes code that the pipeline will use
type Code struct {
	// URL is url of code, e.g. https://github.com/caicloud/cyclone.git
	URL string `json:"Url"`
}

// CodeSource defines a certain revision code
type CodeSource struct {
	// Ref is a reference of git repo, support:
	// - branches   refs/heads/{branchName}
	// - tags       refs/tags/{tagName}
	// - pull
	// - GitLab: refs/merge-requests/{mrID}/head:{targetBranch}
	// - GitHub: refs/pull/{prID}/merge
	Ref string `json:"Ref"`
}

// Container describes the main workload
type Container struct {
	// Name is the name of the container
	Name string `json:"Name"`
	// Image is the image used to run the workload
	Image string `json:"Image"`
	// Commands contains a list of commands to run for this job
	Commands []string `json:"Commands"`
	// Env contains a list of env to run the workload
	Env []EnvVar `json:"Env,omitempty"`
}

// CronTrigger represents the cron trigger policy.
type CronTrigger struct {
	Schedule string `json:"Schedule"`
}

// DelegationWorkload describes the delegation job details
type DelegationWorkload struct {
	// Type indicates what kinds of delegation workload, 'cd', 'approval' supported
	Type string `json:"Type"`
	// Config is json string to define the config used in the workload
	Config string `json:"Config"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Event holds the details of a event.
type Event struct {
	ID        string    `json:"Id"`
	TimeStamp time.Time `json:"TimeStamp"`
	Action    string    `json:"Action"`
	Target    *Target   `json:"Target"`
	Request   *Request  `json:"Request"`
	Actor     *Actor    `json:"Actor"`
}

// ExecutionContext represtents a context used to execute workflows.
type ExecutionContext struct {
	Spec   ExecutionContextSpec   `json:"Spec"`
	Status ExecutionContextStatus `json:"Status"`
}

// ExecutionContextPhase represents the phase of ExecutionContext.
type ExecutionContextPhase string

// ExecutionContextSpec describes the execution context
type ExecutionContextSpec struct {
	Cluster   string `json:"Cluster"`
	Namespace string `json:"Namespace"`
	PVC       string `json:"Pvc"`
}

// ExecutionContextStatus describe the status of execution context, it contains information that affects
// pipeline execution, like reserved resources, pvc status.
type ExecutionContextStatus struct {
	// Phase of the execution context, could be 'Ready', 'NotReady' or 'Unknown'
	Phase ExecutionContextPhase `json:"Phase"`
	// ReservedResources indicate resources that will be used by system components, like pvc watcher, and
	// can not be used by workflows execution.
	ReservedResources []ResourceValue `json:"ReservedResources"`
	// PVC describes status of PVC
	PVC PVCOverallStatus `json:"Pvc"`
}

// GenerationRule ...
type GenerationRule string

// HTTP describes a HTTP type resource
type HTTP struct {
	// Path to find output files, support * to match all files under a directory.
	Path string `json:"Path,omitempty"`
}

// Image describes an image
type Image struct {
	// Name is the image name
	Name string `json:"Name,omitempty"`
}

// Integration contains information about external systems
type Integration struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains integration spec
	Spec IntegrationSpec `json:"Spec"`
}

// IntegrationList represents a list of Integration
type IntegrationList struct {
	v12.ListMeta `json:",inline"`
	Items        []Integration `json:"Items,omitempty"`
}

// IntegrationSource contains various external systems.
// exactly one of its members must be set, and the member must equal with the integration's type.
type IntegrationSource struct {
	// SonarQube describes info about external system sonar qube, and is used for code scanning in CI.
	SonarQube *SonarQubeSource `json:"SonarQube,omitempty"`
	// SCM describes info about external Source Code Management system, and is used to manager code.
	SCM *SCMSource `json:"Scm,omitempty"`
}

// IntegrationSpec contains the integration spec information
type IntegrationSpec struct {
	// Type of integration
	Type IntegrationType `json:"Type"`
	// The actual info about various external systems.
	IntegrationSource `json:",inline"`
}

// IntegrationType defines the type of integration
type IntegrationType string

// Job is the unit of the pipeline
type Job struct {
	// Name is the name of job
	Name string `json:"Name"`
	// Alias is the alias of job
	Alias string `json:"Alias"`
	// Template represents which template the job created from, empty represents customized job.
	// BUT, devops-web used it to store template kind, instead of template name.
	// +optional
	Template string `json:"Template"`
	// Pod contains the workload info to describe a job
	Pod PodWorkload `json:"Pod"`
	// Delegation contains workload that delegated to external services, instead of executed in Cyclone as pod.
	Delegation *DelegationWorkload `json:"Delegation,omitempty"`
	// AllowFailure describes whether this stage can accept failure
	AllowFailure bool `json:"AllowFailure"`
	// ExtraInfo stores the information external system uses.
	// for example, front-end devops-web can store image source type(value could be 'Registry' or 'ImageBuild')
	// in CD stage into this field.
	// +optional
	ExtraInfo map[string]string `json:"ExtraInfo,omitempty"`
}

// JobStatus contains status of a devops job
type JobStatus struct {
	// Name of the job
	Name string `json:"Name"`
	// Alias of the job
	Alias string `json:"Alias"`
	// Status of the job
	Status Status `json:"Status,omitempty"`
	// Outputs of the job
	Outputs map[string]string `json:"Outputs,omitempty"`
	// Events of the job
	Events []StageEvent `json:"Events,omitempty"`
}

// JobTemplate represents the template information of devops job
type JobTemplate struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains job spec
	Spec JobTemplateSpec `json:"Spec"`
}

// JobTemplateList represents a list of JobTemplate
type JobTemplateList struct {
	v12.ListMeta `json:",inline"`
	Items        []JobTemplate `json:"Items,omitempty"`
}

// JobTemplateSpec is the job spec
type JobTemplateSpec struct {
	Job `json:",inline"`
	// Builtin represents if the job template is built-in template
	Builtin bool `json:"Builtin"`
	// Kind is the kind of the job
	Kind string `json:"Kind"`
}

// KeyValue represents a simple object's name and its value
type KeyValue struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Language ...
type Language struct {
	Key  string `json:"Key,omitempty"`
	Name string `json:"Name,omitempty"`
}

// LanguageList represents a list of Language
type LanguageList struct {
	v12.ListMeta `json:",inline"`
	Items        []Language `json:"Items,omitempty"`
}

// Notification represents notifications for workflowrun results.
type Notification struct {
	// Policy represents the policy to send notifications.
	Policy NotificationPolicy `json:"Policy"`
	// Receivers represents the receivers of notifications.
	Receivers []NotificationReceiver `json:"Receivers"`
}

// NotificationPolicy represents the policy to send notifications.
type NotificationPolicy string

// NotificationReceiver represents the receiver of notifications.
type NotificationReceiver struct {
	// Type represents the way to send notifications.
	Type NotificationType `json:"Type"`
	// Addresses represents the addresses to receive notifications.
	Addresses []string `json:"Addresses"`
}

// NotificationType represents the way to send notifications.
type NotificationType string

// Output defines job output.
type Output struct {
	// Type is the output type, only supports 'Image' and "Http" by now
	Type OutputResourceType `json:"Type,omitempty"`
	// Image describe image type output
	Image *Image `json:"Image,omitempty"`
	// HTTP describe http type output
	HTTP *HTTP `json:"Http,omitempty"`
}

// OutputResourceType represtents supported output resource type
type OutputResourceType string

// PVCOverallStatus includes upstream kubernetes PersistentVolumeClaimStatus and human readable pvc usage profile
type PVCOverallStatus struct {
	Usage *PVCUsageStatus               `json:"Usage"`
	Phase v1.PersistentVolumeClaimPhase `json:"Phase"`
}

// PVCUsage represents PVC usages in a tenant, values are in human readable format, for example, '8K', '1.2G'.
type PVCUsage struct {
	// Total is total space
	Total string `json:"Total"`
	// Used is space used
	Used string `json:"Used"`
	// Items are space used by each folder, for example, 'caches' -> '1.2G'
	Items []KeyValue `json:"Items"`
}

// PVCUsageStatus describe PVSUsage and calculate the used percentage of PVC.
type PVCUsageStatus struct {
	PVCUsage       `json:"PvcUsage"`
	UsedPercentage float64 `json:"UsedPercentage"`
}

// PaginationParams describes pagination of request, start and limit defined.
type PaginationParams struct {
	Start  uint64 `source:"query,Start,default=0"`
	Limit  uint64 `source:"query,Limit,default=99999"`
	Filter string `source:"query,Filter"`
	// Sort will sorts the results by metadata.creationTimestamp
	Sort bool `source:"query,Sort,default=false"`
	// Ascending will sort the results by ascending order, otherwise by descending order
	Ascending bool `source:"query,Ascending,default=false"`
	// Detail determines whether to return details of list items.
	Detail bool `source:"query,Detail,default=false"`
}

// PersistentVolumeClaim describes information about pvc belongs to a tenant
type PersistentVolumeClaim struct {
	// StorageClass represents the strorageclass used to create pvc
	StorageClass string `json:"StorageClass"`
	// Size represents the capacity of the pvc, unit supports 'Gi' or 'Mi'
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Size string `json:"Size"`
}

// Pipeline contains information to describe CI/CD pipeline
type Pipeline struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains pipeline spec
	Spec PipelineSpec `json:"Spec"`
	// Status of a pipeline
	Status PipelineStatus `json:"Status"`
}

// PipelineList represents a list of Pipeline
type PipelineList struct {
	v12.ListMeta `json:",inline"`
	Items        []Pipeline `json:"Items,omitempty"`
}

// PipelineSpec contains the pipeline spec information
type PipelineSpec struct {
	// Code contains info about code that the pipeline will use
	Code Code `json:"Code"`
	// CacheDependency controls whether to cache pipeline dependencies
	CacheDependency bool `json:"CacheDependency"`
	// Stages manages a set of stages
	Stages []Stage `json:"Stages,omitempty"`
	// Trigger represents the auto trigger policy for pipelines.
	// Supports time trigger and SCM event trigger.
	Trigger *Trigger `json:"Trigger,omitempty"`
	// Notification represents the notification config of pipeline results.
	Notification `json:"Notification"`
	// CustomQuota indicates whether user overrides the workspace quota
	CustomQuota bool `json:"CustomQuota"`
	// Quota represents the default quota used to run workflow
	Quota []ResourceValue `json:"Quota,omitempty"`
	// Owner represents the owner of pipeline, is the creator when pipeline created.
	Owner string `json:"Owner,omitempty"`
	// ImageTagRule is a rule to generate image tag
	ImageTagRule GenerationRule `json:"ImageTagRule,omitempty"`
}

// PipelineStatus ...
type PipelineStatus struct {
	// RecentRecords represents recent records of the pipeline
	RecentRecords []Record `json:"RecentRecords,omitempty"`
	// RecentSuccessRecords represents recent success records of the pipeline
	RecentSuccessRecords []Record `json:"RecentSuccessRecords,omitempty"`
	// RecentFailedRecords represents recent failed records of the pipeline
	RecentFailedRecords []Record `json:"RecentFailedRecords,omitempty"`
	// TotalApprovalTimeout represents the total timeout of all approval jobs.
	// It is calculated automatically, users does not concern about.
	// Unit: second.
	TotalApprovalTimeout int `json:"TotalApprovalTimeout,omitempty"`
}

// Pod ...
type Pod struct {
	v12.ObjectMeta `json:",inline"`
	Status         PodStatus `json:"Status"`
}

// PodList represents a list of Pod
type PodList struct {
	v12.ListMeta `json:",inline"`
	Items        []Pod `json:"Items,omitempty"`
}

// PodStatus ...
type PodStatus struct {
	Phase v1.PodPhase `json:"Phase"`
}

// PodWorkload describes pod type workload, a complete pod spec is included.
type PodWorkload struct {
	// Arguments of the job
	Arguments []ArgumentValue `json:"Arguments,omitempty"`
	// Outputs of the job outputs
	Outputs []Output `json:"Outputs,omitempty"`
	// Container describes the main workload
	Container Container `json:"Container,omitempty"`
	// Services are a set of services on which main workload container depends
	Services []Container `json:"Services,omitempty"`
}

// PullRequest describes pull requests of SCM repositories.
type PullRequest struct {
	ID          int    `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	State       string `json:"State"`
	// TargetBranch used for GitLab to indicate to which branch the merge-request should merge.
	TargetBranch string `json:"TargetBranch"`
}

// PullRequestList represents a list of PullRequest
type PullRequestList struct {
	v12.ListMeta `json:",inline"`
	Items        []PullRequest `json:"Items,omitempty"`
}

// QualityGate ...
type QualityGate struct {
	ID        int    `json:"Id,omitempty"`
	Name      string `json:"Name,omitempty"`
	IsDefault bool   `json:"IsDefault,omitempty"`
	IsBuiltIn bool   `json:"IsBuiltIn,omitempty"`
	// `nirvana client` will choke if we use anonymous struct here
	Actions QualityGateActions `json:"Actions"`
}

// QualityGateActions ...
type QualityGateActions struct {
	Rename            bool `json:"Rename,omitempty"`
	SetAsDefault      bool `json:"SetAsDefault,omitempty"`
	Copy              bool `json:"Copy,omitempty"`
	AssociateProjects bool `json:"AssociateProjects,omitempty"`
	Delete            bool `json:"Delete,omitempty"`
	ManageConditions  bool `json:"ManageConditions,omitempty"`
}

// QualityGateList represents a list of QualityGate
type QualityGateList struct {
	v12.ListMeta `json:",inline"`
	Items        []QualityGate `json:"Items,omitempty"`
}

// RecentCountParams describes how many recent all/success/failed counts of records to return, used in pipeline request.
type RecentCountParams struct {
	All     int  `source:"query,RecentCount,default=0"`
	Success int  `source:"query,RecentSuccessCount,default=0"`
	Failed  int  `source:"query,RecentFailedCount,default=0"`
	Sort    bool `source:"query,Sort,default=false"`
}

// Record contains pipeline record configuration
type Record struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains record spec
	Spec RecordSpec `json:"Spec"`
	// Status contains statuses of record,stages and jobs
	Status RecordStatus `json:"Status"`
}

// RecordList represents a list of Record
type RecordList struct {
	v12.ListMeta `json:",inline"`
	Items        []Record `json:"Items,omitempty"`
}

// RecordSpec contains the pipeline record spec information
type RecordSpec struct {
	// Pipeline to run
	Pipeline string `json:"Pipeline"`
	// Trigger triggered the pipeline
	Trigger string `json:"Trigger"`
	// CodeSource defines a certain revision code
	CodeSource CodeSource `json:"CodeSource,omitempty"`
	// SCMTag defines whether to create a tag and a release in SCM
	SCMTag *SCMTagInfo `json:"ScmTag,omitempty"`
	// CacheDependency defines whether to use dependency caches to speedup
	CacheDependency bool `json:"CacheDependency"`
	// ImageTag represents name of docker images built with the pipeline
	ImageTag string `json:"ImageTag,omitempty"`
	// Note describes the execution of the pipeline
	Note string `json:"Note"`
}

// RecordStats record of stats.
type RecordStats struct {
	Timestamp int64 `json:"Timestamp"`
	Total     int   `json:"Total"`
}

// RecordStatus contains the status of a pipeline record
type RecordStatus struct {
	// Overall status
	Overall Status `json:"Overall,omitempty"`
	// Stages status
	Stages []*StageStatus `json:"Stages,omitempty"`
}

// ReleaseInfo represents infomatin ablout release.
type ReleaseInfo struct {
	Title string `json:"Title"`
	Note  string `json:"Note"`
}

// Repository represents the information of a repository.
type Repository struct {
	Name string `json:"Name,omitempty"`
	URL  string `json:"Url,omitempty"`
}

// RepositoryList represents a list of Repository
type RepositoryList struct {
	v12.ListMeta `json:",inline"`
	Items        []Repository `json:"Items,omitempty"`
}

// Request holds information about a request.
type Request struct {
	ID        string `json:"Id"`
	Method    string `json:"Method"`
	UserAgent string `json:"UserAgent"`
}

// ResourceValue contains the resource name and its value
type ResourceValue struct {
	// Name represents the resource name, its value could be:
	// * requests.cpu
	// * requests.memory
	// * limits.cpu
	// * limits.memory
	// * {storage class name}/requests.storage
	// * {storage class name}/persistentvolumeclaims
	Name  v1.ResourceName `json:"Name"`
	Value string          `json:"Value"`
}

// SCMAuthType represents the type of SCM auth, support password and token.
type SCMAuthType string

// SCMSource represents Source Code Management to manage code.
type SCMSource struct {
	// Type is the type of scm, e.g. GitLab, GitHub, SVN
	Type SCMType `json:"Type"`
	// Server represents the domain of docker registry.
	Server string `json:"Server"`
	// User is a user of the SCM.
	User string `json:"User"`
	// Password is the password of the corresponding user.
	Password string `json:"Password"`
	// Token is the credential to access SCM.
	Token string `json:"Token"`
	// AuthType is the type of auth way, can be Token or Password
	AuthType SCMAuthType `json:"AuthType"`
}

// SCMTagInfo represents infomatin ablout scm tag and release.
type SCMTagInfo struct {
	// Name of the tag we will create after pipeline running successfully
	Name string `json:"Name,omitempty"`
	// Release info for creating a release in SCM
	Release *ReleaseInfo `json:"Release,omitempty"`
}

// SCMTriggerBasic represents basic config for SCM trigger policy.
type SCMTriggerBasic struct {
	// Enabled represents whether enable this policy.
	Enabled bool `json:"Enabled"`
}

// SCMTriggerPolicy represents trigger policies for SCM events.
// Supports 4 events: push, tag release, pull request and pull request comment.
type SCMTriggerPolicy struct {
	// Push represents trigger policy for push events.
	Push SCMTriggerPush `json:"Push"`
	// TagRelease represents trigger policy for tag release events.
	TagRelease SCMTriggerTagRelease `json:"TagRelease"`
	// PullRequest represents trigger policy for pull request events.
	PullRequest SCMTriggerPullRequest `json:"PullRequest"`
	// PullRequestComment represents trigger policy for pull request comment events.
	PullRequestComment SCMTriggerPullRequestComment `json:"PullRequestComment"`
	// PostCommit represents trigger policy for post commit events.
	PostCommit SCMTriggerPostCommit `json:"PostCommit"`
}

// SCMTriggerPostCommit represents trigger policy for post commit events.
type SCMTriggerPostCommit struct {
	SCMTriggerBasic `json:",inline"`
	// RootURL represents SVN repository root url, this root is retrieved by
	//
	// 'svn info --show-item repos-root-url --username {user} --password {password} --non-interactive
	// --trust-server-cert-failures unknown-ca,cn-mismatch,expired,not-yet-valid,other
	// --no-auth-cache {remote-svn-address}'
	//
	// e.g: http://192.168.21.97/svn/caicloud
	RootURL string `json:"RootUrl"`
	// WorkflowURL represents repository url of the workflow that the wrokflowTrigger related to,
	// Cyclone will checkout code from this URL while executing WorkflowRun.
	// e.g: http://192.168.21.97/svn/caicloud/cyclone
	WorkflowURL string `json:"WorkflowUrl"`
}

// SCMTriggerPullRequest represents trigger policy for pull request events.
type SCMTriggerPullRequest struct {
	SCMTriggerBasic `json:",inline"`
	// Branches represents the pr target branches list to filter PullRequest events.
	Branches []string `json:"Branches"`
}

// SCMTriggerPullRequestComment represents trigger policy for pull request comment events.
type SCMTriggerPullRequestComment struct {
	SCMTriggerBasic `json:",inline"`
	// Comments represents the comment lists to filter pull request comment events.
	Comments []string `json:"Comments"`
}

// SCMTriggerPush represents trigger policy for push events.
type SCMTriggerPush struct {
	SCMTriggerBasic `json:",inline"`
	// Branches represents the branch lists to filter push events.
	Branches []string `json:"Branches"`
}

// SCMTriggerTagRelease represents trigger policy for tag release events.
type SCMTriggerTagRelease struct {
	SCMTriggerBasic `json:",inline"`
}

// SCMType defines the type of Source Code Management
type SCMType string

// Setting represents the config of Tenant.
type Setting struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains tenant spec
	Spec SettingSpec `json:"Spec"`
}

// SettingSpec contains the setting information
type SettingSpec struct {
	// PersistentVolumeClaim describes information about persistent volume claim
	PersistentVolumeClaim PersistentVolumeClaim `json:"PersistentVolumeClaim"`
	// ResourceQuota describes the resource quota of the namespace,
	ResourceQuota []ResourceValue `json:"ResourceQuota"`
	// Cluster is the cluster used to run workflow
	Cluster string `json:"Cluster"`
	// Partition is the partition used to run workflow
	Partition string `json:"Partition"`
	// UsedResourceQuota represents quota used in a partition
	UsedResourceQuota []ResourceValue `json:"UsedResourceQuota,omitempty"`
}

// SonarQubeSource represents a code scanning tool for CI.
type SonarQubeSource struct {
	// Server represents the server address of sonar qube .
	Server string `json:"Server"`
	// Token is the credential to access sonar server.
	Token string `json:"Token"`
}

// Stage contains compass-devops stage
type Stage struct {
	// Alias is the alias of stage
	Alias string `json:"Alias"`
	// Jobs manages a set of jobs, all jobs in the same stage run in parallel
	Jobs []Job `json:"Jobs,omitempty"`
}

// StageArtifact describes artifacts produced by a workflowRun
type StageArtifact struct {
	// Stage name
	Stage string `json:"Stage"`
	// Files name
	File              string    `json:"File"`
	CreationTimestamp time.Time `json:"CreationTimestamp"`
}

// StageArtifactList represents a list of StageArtifact
type StageArtifactList struct {
	v12.ListMeta `json:",inline"`
	Items        []StageArtifact `json:"Items,omitempty"`
}

// StageEvent describes pod warning events for a stage
type StageEvent struct {
	// Event name
	Name string `json:"Name"`
	// This should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status.
	Reason string `json:"Reason,omitempty"`
	// A human-readable description of the status of this operation.
	Message string `json:"Message,omitempty"`
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp v11.Time `json:"LastTimestamp,omitempty"`
	// The number of times this event has occurred.
	Count int32 `json:"Count,omitempty"`
}

// StageStatus contains status of a devops stage
type StageStatus struct {
	// Alias of devops stage
	Alias string `json:"Alias"`
	// State of devops stage
	State StatusPhase `json:"State"`
	// Jobs status
	Jobs []*JobStatus `json:"Jobs,omitempty"`
}

// Statistic represents statistics of project or workflow.
type Statistic struct {
	// Overview statistics
	Overview StatsOverview `json:"Overview"`
	// Details statistics
	Details []*StatsDetail `json:"Details"`
}

// StatsDetail represents detailed statistics
type StatsDetail struct {
	Timestamp int64 `json:"Timestamp"`
	// StatsPhase ...
	StatsPhase `json:",inline"`
}

// StatsOverview represents overview statistics
type StatsOverview struct {
	// Total represents number of workflowruns
	Total int `json:"Total"`
	// StatsPhase ...
	StatsPhase `json:",inline"`
	// SuccessRatio represents ratio of success workflowrun,
	// SuccessRatio == CompletedCount / Total
	SuccessRatio string `json:"SuccessRatio"`
}

// StatsPhase ...
type StatsPhase struct {
	// Pending wfr count
	Pending int `json:"Pending"`
	// Running wfr count
	Running int `json:"Running"`
	// Waiting wfr count
	Waiting int `json:"Waiting"`
	// Succeeded wfr count
	Succeeded int `json:"Succeeded"`
	// Failed wfr count
	Failed int `json:"Failed"`
	// Cancelled wfr count
	Cancelled int `json:"Cancelled"`
}

// Status of a Stage in a WorkflowRun or the whole WorkflowRun.
type Status struct {
	// Phase with value: Running, Waiting, Completed, Error
	Phase StatusPhase `json:"Phase"`
	// LastTransitionTime is the last time the status transitioned from one status to another.
	// +optional
	LastTransitionTime v11.Time `json:"LastTransitionTime,omitempty"`
	// The reason for the status's last transition.
	// +optional
	Reason string `json:"Reason,omitempty"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"Message,omitempty"`
	// StartTime is the start time of processing stage/workflowrun
	StartTime v11.Time `json:"StartTime,omitempty"`
}

// StatusPhase represents the phase of stage status or workflowrun status.
type StatusPhase string

// StringList represents a list of String
type StringList struct {
	v12.ListMeta `json:",inline"`
	Items        []string `json:"Items,omitempty"`
}

// Target holds information about the target of a event.
type Target struct {
	MediaType  string `json:"MediaType"`
	Digest     string `json:"Digest"`
	Repository string `json:"Repository"`
	URL        string `json:"Url"`
	Tag        string `json:"Tag"`
}

// TemplateType represents the type of the template
type TemplateType struct {
	Type string `json:"Type,omitempty" description:"type of the template"`
}

// Trigger represents the auto trigger policy for pipelines.
type Trigger struct {
	// Cron represents time trigger.
	Cron *CronTrigger `json:"Cron,omitempty"`
	// SCM represents webhook trigger config.
	SCM *SCMTriggerPolicy `json:"Scm,omitempty"`
}

// WebhookResponse represents response for webhooks.
type WebhookResponse struct {
	Message string `json:"Message,omitempty"`
}

// Workspace represents the isolated space for your work.
type Workspace struct {
	// Metadata for the particular object, including name, namespace, labels, etc
	v12.ObjectMeta `json:",inline"`
	// Spec contains workspace spec
	Spec WorkspaceSpec `json:"Spec"`
	// Status contains workspace status
	Status WorkspaceStatus `json:"Status"`
}

// WorkspaceList represents a list of Workspace
type WorkspaceList struct {
	v12.ListMeta `json:",inline"`
	Items        []Workspace `json:"Items,omitempty"`
}

// WorkspaceSpec contains the workspace spec information
type WorkspaceSpec struct {
	// SCM contains SCM auth information
	SCM SCMSource `json:"Scm"`
	// Cargo contains information about cargo
	Cargo Cargo `json:"Cargo"`
	// Quota represents the default quota used to run workflow
	Quota []ResourceValue `json:"Quota,omitempty"`
	// Owner represents the owner of workspace, is the creator when workspace created.
	Owner string `json:"Owner"`
}

// WorkspaceStats stats of workspace.
type WorkspaceStats struct {
	WorkspaceName string        `json:"WorkspaceName"`
	RecordStats   []RecordStats `json:"RecordStats"`
}

// WorkspaceStatus contains the workspace status information
type WorkspaceStatus struct {
	// PipelineCount represents the number of pipelines under the workspace
	PipelineCount int `json:"PipelineCount"`
	// UsingCachePipelineCount represents the number of pipeline records that is in Running
	// status and is using dependency cache in the cyclone-server PVC.
	UsingCacheRecordCount *int `json:"UsingCacheRecordCount,omitempty"`
	// CacheDependencyCleanup describes dependency cache cleanup status.
	CacheDependencyCleanup AccelerationCacheCleanupOverallStatus `json:"CacheDependencyCleanup"`
}
