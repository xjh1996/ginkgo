package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	token "github.com/docker/distribution/registry/auth/token"
	time "time"
)

// App ...
type App struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *AppSpec `json:"Spec"`
	Status         string   `json:"Status"`
}

// AppList ...
type AppList struct {
	*v1.ListMeta `json:",inline"`
	Items        []*App `json:"Items"`
}

// AppQuery describe app query
type AppQuery struct {
	Tenant    string `source:"query,Tenant"`
	Cluster   string `source:"query,Cluster"`
	Partition string `source:"query,Partition"`
	Type      string `source:"query,Type"`
	Query     string `source:"query,Query"`
}

// AppSpec describe app
type AppSpec struct {
	// Alias of the cluster
	ClusterAlias string `json:"ClusterAlias"`
	// Partition name
	Partition string `json:"Partition"`
	// Type of the application, 'deployments', 'statefulsets', 'jobs'
	Type string `json:"Type"`
}

// ArtifactTagListResp ...
type ArtifactTagListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*ArtifactTagResp `json:"Items"`
}

// ArtifactTagResp ...
type ArtifactTagResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *TagSpec   `json:"Spec"`
	Status         *TagStatus `json:"Status"`
}

// CargoAccount ...
type CargoAccount struct {
	v1.ObjectMeta `json:",inline"`
	Spec          CargoAccountSpec `json:"Spec"`
}

// CargoAccountSpec ...
type CargoAccountSpec struct {
	Domain   string `json:"domain"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CleanPolicyListResp ...
type CleanPolicyListResp struct {
	*v1.ListMeta `json:",inline"`
	Item         []*ImageCleanPolicy `json:"Item"`
}

// CleanPolicyResp ...
type CleanPolicyResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ImageCleanPolicy `json:"Spec"`
}

// ClusterPartitions ...
type ClusterPartitions struct {
	Name       string   `json:"name"`
	Alias      string   `json:"alias"`
	Partitions []string `json:"partitions"`
}

// ComponentsOverview has the total number and a list of components number of different serverity level.
type ComponentsOverview struct {
	Total   int64                      `json:"Total"`
	Summary []*ComponentsOverviewEntry `json:"Summary"`
}

// ComponentsOverviewEntry ...
type ComponentsOverviewEntry struct {
	Severity string `json:"Severity"`
	Count    int64  `json:"Count"`
}

// ConfigInfo ...
type ConfigInfo struct {
	ConfigType string `json:"ConfigType"`
	Value      string `json:"Value"`
}

// ConfigInfoResp ...
type ConfigInfoResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ConfigInfo `json:"Spec"`
}

// CreateProjectReq ...
type CreateProjectReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec `json:"Spec"`
}

// CreatePublicProjectReq ...
type CreatePublicProjectReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec `json:"Spec"`
}

// CreateRegistryReq ...
type CreateRegistryReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *RegistrySpec `json:"Spec"`
}

// CreateReplicationReq ...
type CreateReplicationReq struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ReplicationSpec `json:"Spec"`
}

// DefaultPublicProject ...
type DefaultPublicProject struct {
	Name string `json:"Name"`
}

// DefaultRegistryInfo ...
type DefaultRegistryInfo struct {
	Host     string `json:"Host"`
	Username string `json:"Username"`
}

// Dockerfile ...
type Dockerfile struct {
	Group   string `yaml:"group" json:"Group"`
	Image   string `yaml:"image" json:"Image"`
	Content string `yaml:"content" json:"Content"`
}

// DockerfileGroup ...
type DockerfileGroup struct {
	Group       string        `yaml:"group" json:"Group"`
	Dockerfiles []*Dockerfile `yaml:"dockerfiles" json:"Dockerfiles"`
}

// DockerfileListResp ...
type DockerfileListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*DockerfileGroup `json:"Items"`
}

// ErrorMessage ...
type ErrorMessage struct {
	Reason  string `json:"Reason"`
	Message string `json:"Message"`
}

// ImageBuild ...
type ImageBuild struct {
	BuildLogPreserveDays  int `json:"BuildLogPreserveDays"`
	CleanBuildLogInterval int `json:"CleanBuildLogInterval"`
}

// ImageBuildLogResp ...
type ImageBuildLogResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ImageBuildRecordSpec `json:"Spec"`
}

// ImageBuildRecordList ...
type ImageBuildRecordList struct {
	*v1.ListMeta `json:",inline"`
	Items        []*ImageBuildRecordResp `json:"Items"`
}

// ImageBuildRecordResp ...
type ImageBuildRecordResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ImageBuildRecordSpec   `json:"Spec"`
	Status         *ImageBuildRecordStatus `json:"Status"`
}

// ImageBuildRecordSpec ...
type ImageBuildRecordSpec struct {
	Registry   string `json:"Registry"`
	Project    string `json:"Project"`
	Repository string `json:"Repository"`
	Tag        string `json:"Tag"`
	Dockerfile string `json:"Dockerfile"`
	User       string `json:"User"`
	LogDeleted bool   `json:"LogDeleted"`
}

// ImageBuildRecordStatus ...
type ImageBuildRecordStatus struct {
	LastUpdateTime time.Time `json:"LastUpdateTime"`
	Status         string    `json:"Status"`
}

// ImageCleanDryRunResp ...
type ImageCleanDryRunResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*RepoImagesResp `json:"Items"`
}

// ImageCleanPolicy image clean policy table
type ImageCleanPolicy struct {
	// Name of the policy
	Name string `bson:"name" json:"Name"`
	// Enabled indicates whether this policy is enabled
	Enabled bool `bason:"enabled" json:"Enabled"`
	// Type is type of the policy
	Type ImageCleanPolicyType `bson:"type" json:"Type"`
	// Registry is name of the registry
	Registry string `bson:"registry" json:"Registry"`
	// Project is project this policy is created for
	Project string `bson:"project" json:"Project"`
	// Time gives how long an image hasn't been touched that should be cleaned, in second.
	Time int64 `bson:"time" json:"Time"`
	// Number gives how many tags to retain for each repo
	Number int `bson:"number" json:"Number"`
	// RetainTags holds tag patterns for tags that should not been cleaned.
	RetainTags []string `bson:"retains" json:"Retains"`
	// TriggerCron is a cron expression to trigger the policy
	TriggerCron string `bson:"cron" json:"Cron"`
	// LastCleanTime the last time this policy is triggered
	LastCleanTime time.Time `bson:"lastCleanTime" json:"LastCleanTime"`
}

// ImageCleanPolicyType ...
type ImageCleanPolicyType string

// ImageDownloadRecordResp ...
type ImageDownloadRecordResp struct {
	*v1.ObjectMeta `json:",inline"`
	Status         *ImageDownloadRecordStatus `json:"Status"`
}

// ImageDownloadRecordStatus ...
type ImageDownloadRecordStatus struct {
	Status         string    `json:"Status"`
	LastUpdateTime time.Time `json:"LastUpdateTime"`
}

// ImageDownloadReq ...
type ImageDownloadReq struct {
	// List of images to be downloaded, in format <repo>:<tag>
	Images []string `json:"images"`
}

// ImageUploadRecord ...
type ImageUploadRecord struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ImageUploadRecordSpec   `json:"Spec"`
	Status         *ImageUploadRecordStatus `json:"Status"`
}

// ImageUploadRecordSpec ...
type ImageUploadRecordSpec struct {
	ErrorMessage      *ErrorMessage     `json:"ErrorMessage"`
	ImageUploadStatus *ImageUploadStats `json:"ImageUploadStats"`
}

// ImageUploadRecordStatus ...
type ImageUploadRecordStatus struct {
	Status         string    `json:"Status"`
	LastUpdateTime time.Time `json:"LastUpdateTime"`
}

// ImageUploadStats ...
type ImageUploadStats struct {
	Succeed []string `json:"succeed"`
	Failed  []string `json:"failed"`
}

// ImageUploadStatsResp ...
type ImageUploadStatsResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ImageUploadStats `json:"Spec"`
}

// Pagination describe pagination default value
type Pagination struct {
	Start int `source:"query,Start,default=0"`
	Limit int `source:"query,Limit,default=99999"`
}

// PartitionsInfo ...
type PartitionsInfo struct {
	Tenants map[string]*TenantPartitions `json:"tenants"`
}

// PermReq ...
type PermReq struct {
	Registry string                   `json:"Registry"`
	Username string                   `json:"Username"`
	Accesses []*token.ResourceActions `json:"Accesses"`
}

// PermResp ...
type PermResp struct {
	Accesses []*token.ResourceActions `json:"Accesses"`
}

// Project ...
type Project struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec   `json:"Spec"`
	Status         *ProjectStatus `json:"Status"`
}

// ProjectCount ...
type ProjectCount struct {
	Public  int64 `json:"Public"`
	Private int64 `json:"Private"`
}

// ProjectListResp ...
type ProjectListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*Project `json:"Items"`
}

// ProjectSpec ...
type ProjectSpec struct {
	IsPublic    bool   `json:"IsPublic"`
	IsProtected bool   `json:"IsProtected"`
	Registry    string `json:"Registry"`
}

// ProjectStatus ...
type ProjectStatus struct {
	Synced              bool              `json:"Synced"`
	RepositoryCount     int               `json:"RepositoryCount"`
	ReplicationCount    int               `json:"ReplicationCount"`
	CleanPolicy         *ImageCleanPolicy `json:"CleanPolicy,omitempty"`
	LastImageUpdateTime time.Time         `json:"LastImageUpdateTime"`
	LastUpdateTime      time.Time         `json:"LastUpdateTime"`
}

// PublicProject ...
type PublicProject struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec   `json:"Spec"`
	Status         *ProjectStatus `json:"Status"`
}

// Record ...
type Record struct {
	v1.ObjectMeta `json:",inline"`
	Spec          RecordSpec    `json:"Spec"`
	Status        *RecordStatus `json:"Status"`
}

// RecordArtifact ...
type RecordArtifact struct {
	v1.ObjectMeta `json:",inline"`
	Spec          RecordArtifactSpec    `json:"Spec"`
	Status        *RecordArtifactStatus `json:"Status"`
}

// RecordArtifactListResp ...
type RecordArtifactListResp struct {
	v1.ListMeta `json:",inline"`
	Items       []*RecordArtifact `json:"Items"`
}

// RecordArtifactSpec ...
type RecordArtifactSpec struct {
	Artifact  string `json:"Artifact"`
	Operation string `json:"Operation"`
}

// RecordArtifactStatus ...
type RecordArtifactStatus struct {
	Status string `json:"Status"`
}

// RecordListResp ...
type RecordListResp struct {
	v1.ListMeta `json:",inline"`
	Items       []*Record `json:"Items"`
}

// RecordSpec replication record table
type RecordSpec struct {
	Replication *Replication        `json:"Replication"`
	Trigger     *ReplicationTrigger `json:"Trigger"`
}

// RecordStatus ...
type RecordStatus struct {
	Total        int64     `json:"Total"`
	SuccessCount int64     `json:"SuccessCount"`
	FailedCount  int64     `json:"FailedCount"`
	Status       string    `json:"Status"`
	Reason       string    `json:"Reason"`
	StartTime    time.Time `json:"StartTime"`
	EndTime      time.Time `json:"EndTime"`
}

// RegistryListResp ...
type RegistryListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*RegistryResp `json:"Items"`
}

// RegistryResp ...
type RegistryResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *RegistrySpec   `json:"Spec"`
	Status         *RegistryStatus `json:"Status"`
}

// RegistrySpec ...
type RegistrySpec struct {
	Host     string `json:"Host"`
	Domain   string `json:"Domain"`
	Username string `json:"Username,omitempty"`
	Password string `json:"Password,omitempty"`
}

// RegistryStatus ...
type RegistryStatus struct {
	LastUpdateTime  time.Time        `json:"LastUpdateTime"`
	ProjectCount    *ProjectCount    `json:"ProjectCount"`
	RepositoryCount *RepositoryCount `json:"RepositoryCount"`
	StorageStatics  *StorageStatics  `json:"StorageStatics"`
	Healthy         bool             `json:"Healthy"`
}

// Replication ...
type Replication struct {
	v1.ObjectMeta `json:",inline"`
	Spec          *ReplicationSpec   `json:"Spec"`
	Status        *ReplicationStatus `json:"Status"`
}

// ReplicationFilter ...
type ReplicationFilter struct {
	Kind  string `json:"Kind"`
	Value string `json:"Value"`
}

// ReplicationListResp ...
type ReplicationListResp struct {
	v1.ListMeta `json:",inline"`
	Items       []*Replication `json:"Items"`
}

// ReplicationObject ...
type ReplicationObject struct {
	Name   string `json:"Name"`
	Alias  string `json:"Alias"`
	Domain string `json:"Domain"`
}

// ReplicationSpec ...
type ReplicationSpec struct {
	Project           string               `json:"Project"`
	ReplicateNow      bool                 `json:"ReplicateNow"`
	ReplicateDeletion bool                 `json:"ReplicateDeletion"`
	Source            *ReplicationObject   `json:"Source"`
	Target            *ReplicationObject   `json:"Target"`
	Trigger           *ReplicationTrigger  `json:"Trigger"`
	Filters           []*ReplicationFilter `json:"Filters"`
}

// ReplicationStatus ...
type ReplicationStatus struct {
	LastUpdateTime      time.Time `json:"LastUpdateTime"`
	IsSyncing           bool      `json:"IsSyncing"`
	LastReplicationTime time.Time `json:"LastReplicationTime"`
}

// ReplicationTrigger ...
type ReplicationTrigger struct {
	Kind     string           `json:"Type"`
	Settings *TriggerSettings `json:"Settings,omitempty"`
}

// RepoImagesResp ...
type RepoImagesResp struct {
	Project   string `json:"Project"`
	Repo      string `json:"Repo"`
	Tags      []Tag  `json:"Tags"`
	Protected string `json:"Protected"`
}

// Repository ...
type Repository struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *RepositorySpec   `json:"spec"`
	Status         *RepositoryStatus `json:"status"`
}

// RepositoryCount ...
type RepositoryCount struct {
	Public  int64 `json:"Public"`
	Private int64 `json:"Private"`
}

// RepositoryListResp ...
type RepositoryListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*Repository `json:"Items"`
}

// RepositorySpec ...
type RepositorySpec struct {
	Project  string `json:"Project"`
	FullName string `json:"FullName"`
}

// RepositoryStatus ...
type RepositoryStatus struct {
	// json still use TagCount
	ArtifactCount  int64     `json:"TagCount"`
	PullCount      int64     `json:"PullCount"`
	LastUpdateTime time.Time `json:"LastUpdateTime"`
}

// ScanOverview ...
type ScanOverview struct {
	Status       string              `json:"Status"`
	Severity     string              `json:"Severity"`
	ReportID     string              `json:"ReportId"`
	CompOverview *ComponentsOverview `json:"Components,omitempty"`
	CreationTime time.Time           `json:"CreationTime,omitempty"`
	UpdateTime   time.Time           `json:"UpdateTime,omitempty"`
}

// StatItem ...
type StatItem struct {
	TimeStamp int64 `json:"TimeStamp"`
	Count     int   `json:"Count"`
}

// StatItemListResp ...
type StatItemListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*StatItemResp `json:"Items"`
}

// StatItemResp ...
type StatItemResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *StatItem `json:"Spec"`
}

// StorageStatics ...
type StorageStatics struct {
	Used  uint64 `json:"Used"`
	Total uint64 `json:"Total"`
}

// SystemConfig ...
type SystemConfig struct {
	DefaultPublicProjects []*DefaultPublicProject `json:"DefaultPublicProjects"`
	BaseImageProject      string                  `json:"BaseImageProject"`
	DefaultRegistry       *DefaultRegistryInfo    `json:"DefaultRegistry"`
	ImageBuild            *ImageBuild             `json:"ImageBuild"`
}

// SystemConfigResp ...
type SystemConfigResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *SystemConfig `json:"Spec"`
}

// Tag describe tag name and digest
type Tag struct {
	Name   string `json:"Name"`
	Digest string `json:"Digest"`
}

// TagHistoryLine describes image build history
// each layer is a application/vnd.docker.image.rootfs.diff.tar.gzip file in registry
type TagHistoryLine struct {
	// .tar.gzip file digest of this layer, which is different from Image ID
	Digest    string    `json:"Digest"`
	Created   time.Time `json:"Created"`
	CreatedBy string    `json:"CreatedBy"`
	// if this field is true, "Size" filed is 0, and no .tar.gzip file exists in registry
	EmptyLayer bool `json:"EmptyLayer"`
	// .tar.gzip file size of this layer
	Size    int    `json:"Size"`
	Comment string `json:"Comment"`
}

// TagSpec ...
type TagSpec struct {
	ArtifactDigest string            `json:"ArtifactDigest"`
	Image          string            `json:"Image"`
	BaseImage      string            `json:"BaseImage"`
	BuildHistory   []*TagHistoryLine `json:"History"`
}

// TagStatus ...
type TagStatus struct {
	Author               string           `json:"Author"`
	VulnerabilitiesCount int64            `json:"VulnerabilitiesCount"`
	Vulnerabilities      []*Vulnerability `json:"Vulnerabilities"`
	ScanOverview         *ScanOverview    `json:"ScanOverview,omitempty"`
}

// TenantPartitions ...
type TenantPartitions struct {
	Name     string                        `json:"name"`
	Alias    string                        `json:"alias"`
	Clusters map[string]*ClusterPartitions `json:"clusters"`
}

// TenantStatistic ...
type TenantStatistic struct {
	Tenant       string           `json:"Tenant"`
	ProjectCount *ProjectCount    `json:"ProjectCount"`
	RepoCount    *RepositoryCount `json:"RepositoryCount"`
}

// TenantStatisticListResp ...
type TenantStatisticListResp struct {
	*v1.ListMeta `json:",inline"`
	Items        []*TenantStatisticResp `json:"Items"`
}

// TenantStatisticResp ...
type TenantStatisticResp struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *TenantStatistic `json:"Spec"`
}

// TriggerImageCopyReq ...
type TriggerImageCopyReq struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// TriggerReplicationReq ...
type TriggerReplicationReq struct {
	Action string `json:"Action"`
}

// TriggerSettings is the setting about the trigger
type TriggerSettings struct {
	Cron string `json:"Cron"`
}

// UpdateProjectReq ...
type UpdateProjectReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec `json:"Spec"`
}

// UpdatePublicProjectReq ...
type UpdatePublicProjectReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *ProjectSpec `json:"Spec"`
}

// UpdateRegistryReq ...
type UpdateRegistryReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *RegistrySpec `json:"Spec"`
}

// UpdateReplicationReq ...
type UpdateReplicationReq struct {
	v1.ObjectMeta `json:",inline"`
	Spec          ReplicationSpec `json:"Spec"`
}

// UpdateRepositoryReq ...
type UpdateRepositoryReq struct {
	*v1.ObjectMeta `json:",inline"`
	Spec           *RepositorySpec `json:"Spec"`
}

// Vulnerability ...
type Vulnerability struct {
	Name        string   `json:"Name"`
	Package     string   `json:"Package"`
	Description string   `json:"Description"`
	Links       []string `json:"Links"`
	Severity    string   `json:"Severity"`
	Version     string   `json:"Version"`
	Fixed       string   `json:"FixedVersion,omitempty"`
}
