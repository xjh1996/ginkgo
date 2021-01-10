package v20201010

import (
	"context"
	io "io"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// BuildImage description:
	// build image
	BuildImage(ctx context.Context, xUser string, xTenant string, registry string, project string, xTag string, ioReader io.Reader) (imageBuildRecordResp *ImageBuildRecordResp, err error)
	// CheckCargoAccount description:
	// check user cargo account
	CheckCargoAccount(ctx context.Context) (err error)
	// CheckCargoPermissions description:
	// check cargo permissions
	CheckCargoPermissions(ctx context.Context, permReq PermReq) (permResp *PermResp, err error)
	// CreateArtifactCleanPolicy description:
	// create artifact clean policy
	CreateArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, imageCleanPolicy *ImageCleanPolicy) (cleanPolicyResp *CleanPolicyResp, err error)
	// CreateArtifactProject description:
	// create artifact project
	CreateArtifactProject(ctx context.Context, xTenant string, registry string, createProjectReq *CreateProjectReq) (project *Project, err error)
	// CreateArtifactPublicProject description:
	// create artifact public project
	CreateArtifactPublicProject(ctx context.Context, xTenant string, registry string, createPublicProjectReq *CreatePublicProjectReq) (publicProject *PublicProject, err error)
	// CreateRegistry description:
	// create registry
	CreateRegistry(ctx context.Context, xTenant string, xUser string, createRegistryReq *CreateRegistryReq) (registryResp *RegistryResp, err error)
	// CreateReplication description:
	// create replication
	CreateReplication(ctx context.Context, xTenant string, createReplicationReq *CreateReplicationReq) (replication *Replication, err error)
	// DeleteArtifactCleanPolicy description:
	// delete artifact clean policy
	DeleteArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (err error)
	// DeleteArtifactProject description:
	// delete artifact project
	DeleteArtifactProject(ctx context.Context, xTenant string, registry string, project string) (err error)
	// DeleteArtifactPublicProject description:
	// delete artifact public project
	DeleteArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string) (err error)
	// DeleteArtifactPublicTag description:
	// delete public tag
	DeleteArtifactPublicTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error)
	// DeleteArtifactTag description:
	// delete artifact tag
	DeleteArtifactTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error)
	// DeletePublicRepository description:
	// delete public repository
	DeletePublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (err error)
	// DeleteRegistry description:
	// delete registry
	DeleteRegistry(ctx context.Context, xTenant string, registry string) (err error)
	// DeleteReplication description:
	// delete replication
	DeleteReplication(ctx context.Context, xTenant string, replication string) (err error)
	// DeleteRepository description:
	// delete repository
	DeleteRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (err error)
	// DownloadImage description:
	// download image
	DownloadImage(ctx context.Context, xTenant string, registry string, project string, download string, filename string) (ioReadCloser io.ReadCloser, strings map[string]string, err error)
	// DryRunArtifactCleanPolicy description:
	// dry run artifact clean policy
	DryRunArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (imageCleanDryRunResp *ImageCleanDryRunResp, err error)
	// GetArtifactCleanPolicy description:
	// get artifact clean policy
	GetArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (cleanPolicyResp *CleanPolicyResp, err error)
	// GetArtifactProject description:
	// get artifact project
	GetArtifactProject(ctx context.Context, xTenant string, registry string, project string) (project1 *Project, err error)
	// GetArtifactPublicProject description:
	// get artifact public project
	GetArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string) (project *Project, err error)
	// GetArtifactPublicTag description:
	// get artifact public tag
	GetArtifactPublicTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (artifactTagResp *ArtifactTagResp, err error)
	// GetArtifactTag description:
	// get artifact tag
	GetArtifactTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (artifactTagResp *ArtifactTagResp, err error)
	// GetCICDSummary description:
	// Get CICD summary
	GetCICDSummary(ctx context.Context, xTenant string, xUser string, pagination *Pagination) (registryListResp *RegistryListResp, err error)
	// GetCargoAccount description:
	// get cargo account
	GetCargoAccount(ctx context.Context, registry string, xUser string) (cargoAccount *CargoAccount, err error)
	// GetCargoConfig description:
	// get cargo config
	GetCargoConfig(ctx context.Context, xTenant string, config string) (configInfoResp *ConfigInfoResp, err error)
	// GetCargoPartitionStats description:
	// get partition stats
	GetCargoPartitionStats(ctx context.Context, xTenant string) (partitionsInfo *PartitionsInfo, err error)
	// GetCargoServerConfig description:
	// get system config
	GetCargoServerConfig(ctx context.Context) (systemConfigResp *SystemConfigResp, err error)
	// GetImageBuildLog description:
	// get image build record log
	GetImageBuildLog(ctx context.Context, registry string, project string, buildRecord string) (imageBuildLogResp *ImageBuildLogResp, err error)
	// GetImageBuildLogStream description:
	// get image build log stream
	GetImageBuildLogStream(ctx context.Context, registry string, project string, buildRecord string) (err error)
	// GetImageBuildRecord description:
	// get image build record
	GetImageBuildRecord(ctx context.Context, registry string, project string, buildRecord string) (imageBuildRecordResp *ImageBuildRecordResp, err error)
	// GetImageDownloadRecord description:
	// get image download record
	GetImageDownloadRecord(ctx context.Context, xTenant string, registry string, project string, download string) (imageDownloadRecordResp *ImageDownloadRecordResp, err error)
	// GetImageUploadRecord description:
	// get image upload record
	GetImageUploadRecord(ctx context.Context, xTenant string, registry string, project string, uploadRecord string) (imageUploadRecord *ImageUploadRecord, err error)
	// GetPublicRepository description:
	// get public repository
	GetPublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (repository1 *Repository, err error)
	// GetRegistry description:
	// get registry
	GetRegistry(ctx context.Context, xTenant string, xUser string, registry string) (registryResp *RegistryResp, err error)
	// GetReplication description:
	// get replication
	GetReplication(ctx context.Context, xTenant string, replication string) (replication1 *Replication, err error)
	// GetReplicationRecord description:
	// get replication record
	GetReplicationRecord(ctx context.Context, xTenant string, replication string, record string) (record1 *Record, err error)
	// GetRepository description:
	// get repository
	GetRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (repository1 *Repository, err error)
	// ListArtifactCleanPolicies description:
	// list artifact clean policies
	ListArtifactCleanPolicies(ctx context.Context, xTenant string, registry string, project string) (cleanPolicyListResp *CleanPolicyListResp, err error)
	// ListArtifactProjectStats description:
	// list artifact project stats
	ListArtifactProjectStats(ctx context.Context, xTenant string, registry string, project string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error)
	// ListArtifactProjects description:
	// list artifact projects
	ListArtifactProjects(ctx context.Context, xSequenceID string, xTenant string, xUser string, registry string, includePublic bool, q string, pagination *Pagination) (projectListResp *ProjectListResp, strings map[string]string, err error)
	// ListArtifactPublicProjectStats description:
	// list artifact public project stats
	ListArtifactPublicProjectStats(ctx context.Context, xTenant string, registry string, publicProject string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error)
	// ListArtifactPublicProjects description:
	// list artifact public projects
	ListArtifactPublicProjects(ctx context.Context, xSequenceID string, xTenant string, registry string, pagination *Pagination) (projectListResp *ProjectListResp, strings map[string]string, err error)
	// ListArtifactPublicTags description:
	// list artifact public tags
	ListArtifactPublicTags(ctx context.Context, xTenant string, registry string, project string, repository string, q string, baseImageCheck string, pagination *Pagination) (artifactTagListResp *ArtifactTagListResp, err error)
	// ListArtifactTags description:
	// list artifact tags
	ListArtifactTags(ctx context.Context, xTenant string, registry string, project string, repository string, q string, baseImageCheck string, pagination *Pagination) (artifactTagListResp *ArtifactTagListResp, err error)
	// ListImageApps description:
	// list image apps
	ListImageApps(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, filters *AppQuery, pagination *Pagination) (appList *AppList, err error)
	// ListImageBuildRecords description:
	// list image build records
	ListImageBuildRecords(ctx context.Context, registry string, project string, q string, pagination *Pagination) (imageBuildRecordList *ImageBuildRecordList, err error)
	// ListImageDockerfiles description:
	// list image dockerfiles
	ListImageDockerfiles(ctx context.Context, xTenant string) (dockerfileListResp *DockerfileListResp, err error)
	// ListPublicRepositories description:
	// list public repositories
	ListPublicRepositories(ctx context.Context, xSequenceID string, xTenant string, registry string, project string, q string, sort string, pagination *Pagination) (repositoryListResp *RepositoryListResp, strings map[string]string, err error)
	// ListRegistries description:
	// list registries
	ListRegistries(ctx context.Context, xTenant string, xUser string, pagination *Pagination) (registryListResp *RegistryListResp, err error)
	// ListRegistryStats description:
	// list registry stats
	ListRegistryStats(ctx context.Context, xTenant string, registry string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error)
	// ListRegistryUsages description:
	// list registry usages
	ListRegistryUsages(ctx context.Context, xTenant string, registry string, pagination *Pagination) (tenantStatisticListResp *TenantStatisticListResp, err error)
	// ListReplicationRecordArtifacts description:
	// list replication record artifacts
	ListReplicationRecordArtifacts(ctx context.Context, xTenant string, replication string, record string, status string, pagination *Pagination) (recordArtifactListResp *RecordArtifactListResp, err error)
	// ListReplicationRecords description:
	// list replication records
	ListReplicationRecords(ctx context.Context, xTenant string, replication string, status string, triggerType string, pagination *Pagination) (recordListResp *RecordListResp, err error)
	// ListReplications description:
	// list replications
	ListReplications(ctx context.Context, xSequenceID string, xTenant string, direction string, registry string, project string, triggerType string, q string, pagination *Pagination) (replicationListResp *ReplicationListResp, strings map[string]string, err error)
	// ListRepositories description:
	// list repositories
	ListRepositories(ctx context.Context, xSequenceID string, xTenant string, registry string, project string, q string, sort string, pagination *Pagination) (repositoryListResp *RepositoryListResp, strings map[string]string, err error)
	// PrepareImageDownload description:
	// prepare image download
	PrepareImageDownload(ctx context.Context, xTenant string, registry string, project string, imageDownloadReq *ImageDownloadReq) (imageDownloadRecordResp *ImageDownloadRecordResp, err error)
	// PrepareImageUpload description:
	// prepare image upload
	PrepareImageUpload(ctx context.Context, xTenant string, registry string, project string) (imageUploadRecord *ImageUploadRecord, err error)
	// ScanImage description:
	// scan image
	ScanImage(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error)
	// ScanPublicImage description:
	// scan public image
	ScanPublicImage(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error)
	// TriggerArtifactCleanPolicy description:
	// trigger artifact clean policy
	TriggerArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (err error)
	// TriggerArtifactImageCopy description:
	// trigger artifact image copy
	TriggerArtifactImageCopy(ctx context.Context, xTenant string, xUser string, triggerImageCopyReq *TriggerImageCopyReq) (err error)
	// TriggerReplication description:
	// Trigger replication
	TriggerReplication(ctx context.Context, xTenant string, replication string, triggerReplicationReq *TriggerReplicationReq) (err error)
	// UpdateArtifactCleanPolicy description:
	// update artifact clean policy
	UpdateArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string, imageCleanPolicy *ImageCleanPolicy) (err error)
	// UpdateArtifactProject description:
	// update artifact project
	UpdateArtifactProject(ctx context.Context, xTenant string, registry string, project string, updateProjectReq *UpdateProjectReq) (project1 *Project, err error)
	// UpdateArtifactPublicProject description:
	// update artifact public project
	UpdateArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string, updatePublicProjectReq *UpdatePublicProjectReq) (publicProject1 *PublicProject, err error)
	// UpdateCargoConfig description:
	// update cargo config
	UpdateCargoConfig(ctx context.Context, xTenant string, configInfo *ConfigInfo) (configInfoResp *ConfigInfoResp, err error)
	// UpdatePublicRepository description:
	// update public repository
	UpdatePublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string, updateRepositoryReq *UpdateRepositoryReq) (repository1 *Repository, err error)
	// UpdateRegistry description:
	// update registry
	UpdateRegistry(ctx context.Context, xTenant string, xUser string, registry string, updateRegistryReq *UpdateRegistryReq) (registryResp *RegistryResp, err error)
	// UpdateReplication description:
	// update replication
	UpdateReplication(ctx context.Context, xTenant string, replication string, updateReplicationReq *UpdateReplicationReq) (replication1 *Replication, err error)
	// UpdateRepository description:
	// update repository
	UpdateRepository(ctx context.Context, xTenant string, registry string, project string, repository string, updateRepositoryReq *UpdateRepositoryReq) (repository1 *Repository, err error)
	// UploadImage description:
	// upload image
	UploadImage(ctx context.Context, xTenant string, xUploadID string, registry string, project string, ioReader io.Reader) (imageUploadStatsResp *ImageUploadStatsResp, err error)
	// healthcheck description:
	// health check
	healthcheck(ctx context.Context) (err error)
}

// Client for version v20201010.
type Client struct {
	rest *rest.Client
}

// NewClient creates a new client.
func NewClient(cfg *rest.Config) (*Client, error) {
	client, err := rest.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// MustNewClient creates a new client or panic if an error occurs.
func MustNewClient(cfg *rest.Config) *Client {
	client, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

// BuildImage description:
// build image
func (c *Client) BuildImage(ctx context.Context, xUser string, xTenant string, registry string, project string, xTag string, ioReader io.Reader) (imageBuildRecordResp *ImageBuildRecordResp, err error) {
	imageBuildRecordResp = new(ImageBuildRecordResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=BuildImage").
		Header("X-User", xUser).
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Header("X-Tag", xTag).
		Body("application/octet-stream", ioReader).
		TOPRPCData(imageBuildRecordResp).
		Do(ctx)
	return
}

// CheckCargoAccount description:
// check user cargo account
func (c *Client) CheckCargoAccount(ctx context.Context) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CheckCargoAccount").
		Do(ctx)
	return
}

// CheckCargoPermissions description:
// check cargo permissions
func (c *Client) CheckCargoPermissions(ctx context.Context, permReq PermReq) (permResp *PermResp, err error) {
	permResp = new(PermResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CheckCargoPermissions").
		Body("application/json", permReq).
		Data(permResp).
		Do(ctx)
	return
}

// CreateArtifactCleanPolicy description:
// create artifact clean policy
func (c *Client) CreateArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, imageCleanPolicy *ImageCleanPolicy) (cleanPolicyResp *CleanPolicyResp, err error) {
	cleanPolicyResp = new(CleanPolicyResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Body("application/json", imageCleanPolicy).
		TOPRPCData(cleanPolicyResp).
		Do(ctx)
	return
}

// CreateArtifactProject description:
// create artifact project
func (c *Client) CreateArtifactProject(ctx context.Context, xTenant string, registry string, createProjectReq *CreateProjectReq) (project *Project, err error) {
	project = new(Project)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateArtifactProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Body("application/json", createProjectReq).
		TOPRPCData(project).
		Do(ctx)
	return
}

// CreateArtifactPublicProject description:
// create artifact public project
func (c *Client) CreateArtifactPublicProject(ctx context.Context, xTenant string, registry string, createPublicProjectReq *CreatePublicProjectReq) (publicProject *PublicProject, err error) {
	publicProject = new(PublicProject)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateArtifactPublicProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Body("application/json", createPublicProjectReq).
		TOPRPCData(publicProject).
		Do(ctx)
	return
}

// CreateRegistry description:
// create registry
func (c *Client) CreateRegistry(ctx context.Context, xTenant string, xUser string, createRegistryReq *CreateRegistryReq) (registryResp *RegistryResp, err error) {
	registryResp = new(RegistryResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateRegistry").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Body("application/json", createRegistryReq).
		TOPRPCData(registryResp).
		Do(ctx)
	return
}

// CreateReplication description:
// create replication
func (c *Client) CreateReplication(ctx context.Context, xTenant string, createReplicationReq *CreateReplicationReq) (replication *Replication, err error) {
	replication = new(Replication)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateReplication").
		Header("X-Tenant", xTenant).
		Body("application/json", createReplicationReq).
		TOPRPCData(replication).
		Do(ctx)
	return
}

// DeleteArtifactCleanPolicy description:
// delete artifact clean policy
func (c *Client) DeleteArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Policy", policy).
		Do(ctx)
	return
}

// DeleteArtifactProject description:
// delete artifact project
func (c *Client) DeleteArtifactProject(ctx context.Context, xTenant string, registry string, project string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteArtifactProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Do(ctx)
	return
}

// DeleteArtifactPublicProject description:
// delete artifact public project
func (c *Client) DeleteArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteArtifactPublicProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("PublicProject", publicProject).
		Do(ctx)
	return
}

// DeleteArtifactPublicTag description:
// delete public tag
func (c *Client) DeleteArtifactPublicTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteArtifactPublicTag").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		Do(ctx)
	return
}

// DeleteArtifactTag description:
// delete artifact tag
func (c *Client) DeleteArtifactTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteArtifactTag").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		Do(ctx)
	return
}

// DeletePublicRepository description:
// delete public repository
func (c *Client) DeletePublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePublicRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Do(ctx)
	return
}

// DeleteRegistry description:
// delete registry
func (c *Client) DeleteRegistry(ctx context.Context, xTenant string, registry string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteRegistry").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Do(ctx)
	return
}

// DeleteReplication description:
// delete replication
func (c *Client) DeleteReplication(ctx context.Context, xTenant string, replication string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteReplication").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Do(ctx)
	return
}

// DeleteRepository description:
// delete repository
func (c *Client) DeleteRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Do(ctx)
	return
}

// DownloadImage description:
// download image
func (c *Client) DownloadImage(ctx context.Context, xTenant string, registry string, project string, download string, filename string) (ioReadCloser io.ReadCloser, strings map[string]string, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DownloadImage").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Download", download).
		Query("Filename", filename).
		Data(&ioReadCloser).
		Meta(&strings).
		Do(ctx)
	return
}

// DryRunArtifactCleanPolicy description:
// dry run artifact clean policy
func (c *Client) DryRunArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (imageCleanDryRunResp *ImageCleanDryRunResp, err error) {
	imageCleanDryRunResp = new(ImageCleanDryRunResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DryRunArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Policy", policy).
		TOPRPCData(imageCleanDryRunResp).
		Do(ctx)
	return
}

// GetArtifactCleanPolicy description:
// get artifact clean policy
func (c *Client) GetArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (cleanPolicyResp *CleanPolicyResp, err error) {
	cleanPolicyResp = new(CleanPolicyResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Policy", policy).
		TOPRPCData(cleanPolicyResp).
		Do(ctx)
	return
}

// GetArtifactProject description:
// get artifact project
func (c *Client) GetArtifactProject(ctx context.Context, xTenant string, registry string, project string) (project1 *Project, err error) {
	project1 = new(Project)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetArtifactProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		TOPRPCData(project1).
		Do(ctx)
	return
}

// GetArtifactPublicProject description:
// get artifact public project
func (c *Client) GetArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string) (project *Project, err error) {
	project = new(Project)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetArtifactPublicProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("PublicProject", publicProject).
		TOPRPCData(project).
		Do(ctx)
	return
}

// GetArtifactPublicTag description:
// get artifact public tag
func (c *Client) GetArtifactPublicTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (artifactTagResp *ArtifactTagResp, err error) {
	artifactTagResp = new(ArtifactTagResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetArtifactPublicTag").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		TOPRPCData(artifactTagResp).
		Do(ctx)
	return
}

// GetArtifactTag description:
// get artifact tag
func (c *Client) GetArtifactTag(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (artifactTagResp *ArtifactTagResp, err error) {
	artifactTagResp = new(ArtifactTagResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetArtifactTag").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		TOPRPCData(artifactTagResp).
		Do(ctx)
	return
}

// GetCICDSummary description:
// Get CICD summary
func (c *Client) GetCICDSummary(ctx context.Context, xTenant string, xUser string, pagination *Pagination) (registryListResp *RegistryListResp, err error) {
	registryListResp = new(RegistryListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCICDSummary").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(registryListResp).
		Do(ctx)
	return
}

// GetCargoAccount description:
// get cargo account
func (c *Client) GetCargoAccount(ctx context.Context, registry string, xUser string) (cargoAccount *CargoAccount, err error) {
	cargoAccount = new(CargoAccount)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCargoAccount").
		Query("Registry", registry).
		Header("X-User", xUser).
		TOPRPCData(cargoAccount).
		Do(ctx)
	return
}

// GetCargoConfig description:
// get cargo config
func (c *Client) GetCargoConfig(ctx context.Context, xTenant string, config string) (configInfoResp *ConfigInfoResp, err error) {
	configInfoResp = new(ConfigInfoResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCargoConfig").
		Header("X-Tenant", xTenant).
		Query("Config", config).
		TOPRPCData(configInfoResp).
		Do(ctx)
	return
}

// GetCargoPartitionStats description:
// get partition stats
func (c *Client) GetCargoPartitionStats(ctx context.Context, xTenant string) (partitionsInfo *PartitionsInfo, err error) {
	partitionsInfo = new(PartitionsInfo)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCargoPartitionStats").
		Header("X-Tenant", xTenant).
		TOPRPCData(partitionsInfo).
		Do(ctx)
	return
}

// GetCargoServerConfig description:
// get system config
func (c *Client) GetCargoServerConfig(ctx context.Context) (systemConfigResp *SystemConfigResp, err error) {
	systemConfigResp = new(SystemConfigResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCargoServerConfig").
		TOPRPCData(systemConfigResp).
		Do(ctx)
	return
}

// GetImageBuildLog description:
// get image build record log
func (c *Client) GetImageBuildLog(ctx context.Context, registry string, project string, buildRecord string) (imageBuildLogResp *ImageBuildLogResp, err error) {
	imageBuildLogResp = new(ImageBuildLogResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetImageBuildLog").
		Query("Registry", registry).
		Query("Project", project).
		Query("BuildRecord", buildRecord).
		TOPRPCData(imageBuildLogResp).
		Do(ctx)
	return
}

// GetImageBuildLogStream description:
// get image build log stream
func (c *Client) GetImageBuildLogStream(ctx context.Context, registry string, project string, buildRecord string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetImageBuildLogStream").
		Query("Registry", registry).
		Query("Project", project).
		Query("BuildRecord", buildRecord).
		Do(ctx)
	return
}

// GetImageBuildRecord description:
// get image build record
func (c *Client) GetImageBuildRecord(ctx context.Context, registry string, project string, buildRecord string) (imageBuildRecordResp *ImageBuildRecordResp, err error) {
	imageBuildRecordResp = new(ImageBuildRecordResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetImageBuildRecord").
		Query("Registry", registry).
		Query("Project", project).
		Query("BuildRecord", buildRecord).
		TOPRPCData(imageBuildRecordResp).
		Do(ctx)
	return
}

// GetImageDownloadRecord description:
// get image download record
func (c *Client) GetImageDownloadRecord(ctx context.Context, xTenant string, registry string, project string, download string) (imageDownloadRecordResp *ImageDownloadRecordResp, err error) {
	imageDownloadRecordResp = new(ImageDownloadRecordResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetImageDownloadRecord").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Download", download).
		TOPRPCData(imageDownloadRecordResp).
		Do(ctx)
	return
}

// GetImageUploadRecord description:
// get image upload record
func (c *Client) GetImageUploadRecord(ctx context.Context, xTenant string, registry string, project string, uploadRecord string) (imageUploadRecord *ImageUploadRecord, err error) {
	imageUploadRecord = new(ImageUploadRecord)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetImageUploadRecord").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("UploadRecord", uploadRecord).
		TOPRPCData(imageUploadRecord).
		Do(ctx)
	return
}

// GetPublicRepository description:
// get public repository
func (c *Client) GetPublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (repository1 *Repository, err error) {
	repository1 = new(Repository)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPublicRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		TOPRPCData(repository1).
		Do(ctx)
	return
}

// GetRegistry description:
// get registry
func (c *Client) GetRegistry(ctx context.Context, xTenant string, xUser string, registry string) (registryResp *RegistryResp, err error) {
	registryResp = new(RegistryResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetRegistry").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Registry", registry).
		TOPRPCData(registryResp).
		Do(ctx)
	return
}

// GetReplication description:
// get replication
func (c *Client) GetReplication(ctx context.Context, xTenant string, replication string) (replication1 *Replication, err error) {
	replication1 = new(Replication)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetReplication").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		TOPRPCData(replication1).
		Do(ctx)
	return
}

// GetReplicationRecord description:
// get replication record
func (c *Client) GetReplicationRecord(ctx context.Context, xTenant string, replication string, record string) (record1 *Record, err error) {
	record1 = new(Record)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetReplicationRecord").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Query("Record", record).
		TOPRPCData(record1).
		Do(ctx)
	return
}

// GetRepository description:
// get repository
func (c *Client) GetRepository(ctx context.Context, xTenant string, registry string, project string, repository string) (repository1 *Repository, err error) {
	repository1 = new(Repository)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		TOPRPCData(repository1).
		Do(ctx)
	return
}

// ListArtifactCleanPolicies description:
// list artifact clean policies
func (c *Client) ListArtifactCleanPolicies(ctx context.Context, xTenant string, registry string, project string) (cleanPolicyListResp *CleanPolicyListResp, err error) {
	cleanPolicyListResp = new(CleanPolicyListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactCleanPolicies").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		TOPRPCData(cleanPolicyListResp).
		Do(ctx)
	return
}

// ListArtifactProjectStats description:
// list artifact project stats
func (c *Client) ListArtifactProjectStats(ctx context.Context, xTenant string, registry string, project string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error) {
	statItemListResp = new(StatItemListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactProjectStats").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Operation", operation).
		Query("StartTime", startTime).
		Query("EndTime", endTime).
		TOPRPCData(statItemListResp).
		Do(ctx)
	return
}

// ListArtifactProjects description:
// list artifact projects
func (c *Client) ListArtifactProjects(ctx context.Context, xSequenceID string, xTenant string, xUser string, registry string, includePublic bool, q string, pagination *Pagination) (projectListResp *ProjectListResp, strings map[string]string, err error) {
	projectListResp = new(ProjectListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactProjects").
		Header("X-Sequence-ID", xSequenceID).
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Registry", registry).
		Query("IncludePublic", includePublic).
		Query("Q", q).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(projectListResp).
		Meta(&strings).
		Do(ctx)
	return
}

// ListArtifactPublicProjectStats description:
// list artifact public project stats
func (c *Client) ListArtifactPublicProjectStats(ctx context.Context, xTenant string, registry string, publicProject string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error) {
	statItemListResp = new(StatItemListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactPublicProjectStats").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("PublicProject", publicProject).
		Query("Operation", operation).
		Query("StartTime", startTime).
		Query("EndTime", endTime).
		TOPRPCData(statItemListResp).
		Do(ctx)
	return
}

// ListArtifactPublicProjects description:
// list artifact public projects
func (c *Client) ListArtifactPublicProjects(ctx context.Context, xSequenceID string, xTenant string, registry string, pagination *Pagination) (projectListResp *ProjectListResp, strings map[string]string, err error) {
	projectListResp = new(ProjectListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactPublicProjects").
		Header("X-Sequence-ID", xSequenceID).
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(projectListResp).
		Meta(&strings).
		Do(ctx)
	return
}

// ListArtifactPublicTags description:
// list artifact public tags
func (c *Client) ListArtifactPublicTags(ctx context.Context, xTenant string, registry string, project string, repository string, q string, baseImageCheck string, pagination *Pagination) (artifactTagListResp *ArtifactTagListResp, err error) {
	artifactTagListResp = new(ArtifactTagListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactPublicTags").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Q", q).
		Query("BaseImageCheck", baseImageCheck).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(artifactTagListResp).
		Do(ctx)
	return
}

// ListArtifactTags description:
// list artifact tags
func (c *Client) ListArtifactTags(ctx context.Context, xTenant string, registry string, project string, repository string, q string, baseImageCheck string, pagination *Pagination) (artifactTagListResp *ArtifactTagListResp, err error) {
	artifactTagListResp = new(ArtifactTagListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListArtifactTags").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Q", q).
		Query("BaseImageCheck", baseImageCheck).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(artifactTagListResp).
		Do(ctx)
	return
}

// ListImageApps description:
// list image apps
func (c *Client) ListImageApps(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, filters *AppQuery, pagination *Pagination) (appList *AppList, err error) {
	appList = new(AppList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListImageApps").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("Tenant", filters.Tenant).
		Query("Cluster", filters.Cluster).
		Query("Partition", filters.Partition).
		Query("Type", filters.Type).
		Query("Query", filters.Query).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Data(appList).
		Do(ctx)
	return
}

// ListImageBuildRecords description:
// list image build records
func (c *Client) ListImageBuildRecords(ctx context.Context, registry string, project string, q string, pagination *Pagination) (imageBuildRecordList *ImageBuildRecordList, err error) {
	imageBuildRecordList = new(ImageBuildRecordList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListImageBuildRecords").
		Query("Registry", registry).
		Query("Project", project).
		Query("Q", q).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(imageBuildRecordList).
		Do(ctx)
	return
}

// ListImageDockerfiles description:
// list image dockerfiles
func (c *Client) ListImageDockerfiles(ctx context.Context, xTenant string) (dockerfileListResp *DockerfileListResp, err error) {
	dockerfileListResp = new(DockerfileListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListImageDockerfiles").
		Header("X-Tenant", xTenant).
		TOPRPCData(dockerfileListResp).
		Do(ctx)
	return
}

// ListPublicRepositories description:
// list public repositories
func (c *Client) ListPublicRepositories(ctx context.Context, xSequenceID string, xTenant string, registry string, project string, q string, sort string, pagination *Pagination) (repositoryListResp *RepositoryListResp, strings map[string]string, err error) {
	repositoryListResp = new(RepositoryListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPublicRepositories").
		Header("X-Sequence-ID", xSequenceID).
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Q", q).
		Query("Sort", sort).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(repositoryListResp).
		Meta(&strings).
		Do(ctx)
	return
}

// ListRegistries description:
// list registries
func (c *Client) ListRegistries(ctx context.Context, xTenant string, xUser string, pagination *Pagination) (registryListResp *RegistryListResp, err error) {
	registryListResp = new(RegistryListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRegistries").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(registryListResp).
		Do(ctx)
	return
}

// ListRegistryStats description:
// list registry stats
func (c *Client) ListRegistryStats(ctx context.Context, xTenant string, registry string, operation string, startTime int64, endTime int64) (statItemListResp *StatItemListResp, err error) {
	statItemListResp = new(StatItemListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRegistryStats").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Operation", operation).
		Query("StartTime", startTime).
		Query("EndTime", endTime).
		TOPRPCData(statItemListResp).
		Do(ctx)
	return
}

// ListRegistryUsages description:
// list registry usages
func (c *Client) ListRegistryUsages(ctx context.Context, xTenant string, registry string, pagination *Pagination) (tenantStatisticListResp *TenantStatisticListResp, err error) {
	tenantStatisticListResp = new(TenantStatisticListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRegistryUsages").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(tenantStatisticListResp).
		Do(ctx)
	return
}

// ListReplicationRecordArtifacts description:
// list replication record artifacts
func (c *Client) ListReplicationRecordArtifacts(ctx context.Context, xTenant string, replication string, record string, status string, pagination *Pagination) (recordArtifactListResp *RecordArtifactListResp, err error) {
	recordArtifactListResp = new(RecordArtifactListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListReplicationRecordArtifacts").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Query("Record", record).
		Query("Status", status).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(recordArtifactListResp).
		Do(ctx)
	return
}

// ListReplicationRecords description:
// list replication records
func (c *Client) ListReplicationRecords(ctx context.Context, xTenant string, replication string, status string, triggerType string, pagination *Pagination) (recordListResp *RecordListResp, err error) {
	recordListResp = new(RecordListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListReplicationRecords").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Query("Status", status).
		Query("TriggerType", triggerType).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(recordListResp).
		Do(ctx)
	return
}

// ListReplications description:
// list replications
func (c *Client) ListReplications(ctx context.Context, xSequenceID string, xTenant string, direction string, registry string, project string, triggerType string, q string, pagination *Pagination) (replicationListResp *ReplicationListResp, strings map[string]string, err error) {
	replicationListResp = new(ReplicationListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListReplications").
		Header("X-Sequence-ID", xSequenceID).
		Header("X-Tenant", xTenant).
		Query("Direction", direction).
		Query("Registry", registry).
		Query("Project", project).
		Query("TriggerType", triggerType).
		Query("Q", q).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(replicationListResp).
		Meta(&strings).
		Do(ctx)
	return
}

// ListRepositories description:
// list repositories
func (c *Client) ListRepositories(ctx context.Context, xSequenceID string, xTenant string, registry string, project string, q string, sort string, pagination *Pagination) (repositoryListResp *RepositoryListResp, strings map[string]string, err error) {
	repositoryListResp = new(RepositoryListResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRepositories").
		Header("X-Sequence-ID", xSequenceID).
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Q", q).
		Query("Sort", sort).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		TOPRPCData(repositoryListResp).
		Meta(&strings).
		Do(ctx)
	return
}

// PrepareImageDownload description:
// prepare image download
func (c *Client) PrepareImageDownload(ctx context.Context, xTenant string, registry string, project string, imageDownloadReq *ImageDownloadReq) (imageDownloadRecordResp *ImageDownloadRecordResp, err error) {
	imageDownloadRecordResp = new(ImageDownloadRecordResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=PrepareImageDownload").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Body("application/json", imageDownloadReq).
		TOPRPCData(imageDownloadRecordResp).
		Do(ctx)
	return
}

// PrepareImageUpload description:
// prepare image upload
func (c *Client) PrepareImageUpload(ctx context.Context, xTenant string, registry string, project string) (imageUploadRecord *ImageUploadRecord, err error) {
	imageUploadRecord = new(ImageUploadRecord)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=PrepareImageUpload").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		TOPRPCData(imageUploadRecord).
		Do(ctx)
	return
}

// ScanImage description:
// scan image
func (c *Client) ScanImage(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ScanImage").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		Do(ctx)
	return
}

// ScanPublicImage description:
// scan public image
func (c *Client) ScanPublicImage(ctx context.Context, xTenant string, registry string, project string, repository string, tag string, artifactDigest string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ScanPublicImage").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Query("Tag", tag).
		Query("ArtifactDigest", artifactDigest).
		Do(ctx)
	return
}

// TriggerArtifactCleanPolicy description:
// trigger artifact clean policy
func (c *Client) TriggerArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=TriggerArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Policy", policy).
		Do(ctx)
	return
}

// TriggerArtifactImageCopy description:
// trigger artifact image copy
func (c *Client) TriggerArtifactImageCopy(ctx context.Context, xTenant string, xUser string, triggerImageCopyReq *TriggerImageCopyReq) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=TriggerArtifactImageCopy").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Body("application/json", triggerImageCopyReq).
		Do(ctx)
	return
}

// TriggerReplication description:
// Trigger replication
func (c *Client) TriggerReplication(ctx context.Context, xTenant string, replication string, triggerReplicationReq *TriggerReplicationReq) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=TriggerReplication").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Body("application/json", triggerReplicationReq).
		Do(ctx)
	return
}

// UpdateArtifactCleanPolicy description:
// update artifact clean policy
func (c *Client) UpdateArtifactCleanPolicy(ctx context.Context, xTenant string, registry string, project string, policy string, imageCleanPolicy *ImageCleanPolicy) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateArtifactCleanPolicy").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Policy", policy).
		Body("application/json", imageCleanPolicy).
		Do(ctx)
	return
}

// UpdateArtifactProject description:
// update artifact project
func (c *Client) UpdateArtifactProject(ctx context.Context, xTenant string, registry string, project string, updateProjectReq *UpdateProjectReq) (project1 *Project, err error) {
	project1 = new(Project)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateArtifactProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Body("application/json", updateProjectReq).
		TOPRPCData(project1).
		Do(ctx)
	return
}

// UpdateArtifactPublicProject description:
// update artifact public project
func (c *Client) UpdateArtifactPublicProject(ctx context.Context, xTenant string, registry string, publicProject string, updatePublicProjectReq *UpdatePublicProjectReq) (publicProject1 *PublicProject, err error) {
	publicProject1 = new(PublicProject)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateArtifactPublicProject").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("PublicProject", publicProject).
		Body("application/json", updatePublicProjectReq).
		TOPRPCData(publicProject1).
		Do(ctx)
	return
}

// UpdateCargoConfig description:
// update cargo config
func (c *Client) UpdateCargoConfig(ctx context.Context, xTenant string, configInfo *ConfigInfo) (configInfoResp *ConfigInfoResp, err error) {
	configInfoResp = new(ConfigInfoResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateCargoConfig").
		Header("X-Tenant", xTenant).
		Body("application/json", configInfo).
		TOPRPCData(configInfoResp).
		Do(ctx)
	return
}

// UpdatePublicRepository description:
// update public repository
func (c *Client) UpdatePublicRepository(ctx context.Context, xTenant string, registry string, project string, repository string, updateRepositoryReq *UpdateRepositoryReq) (repository1 *Repository, err error) {
	repository1 = new(Repository)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdatePublicRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Body("application/json", updateRepositoryReq).
		TOPRPCData(repository1).
		Do(ctx)
	return
}

// UpdateRegistry description:
// update registry
func (c *Client) UpdateRegistry(ctx context.Context, xTenant string, xUser string, registry string, updateRegistryReq *UpdateRegistryReq) (registryResp *RegistryResp, err error) {
	registryResp = new(RegistryResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateRegistry").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Registry", registry).
		Body("application/json", updateRegistryReq).
		TOPRPCData(registryResp).
		Do(ctx)
	return
}

// UpdateReplication description:
// update replication
func (c *Client) UpdateReplication(ctx context.Context, xTenant string, replication string, updateReplicationReq *UpdateReplicationReq) (replication1 *Replication, err error) {
	replication1 = new(Replication)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateReplication").
		Header("X-Tenant", xTenant).
		Query("Replication", replication).
		Body("application/json", updateReplicationReq).
		TOPRPCData(replication1).
		Do(ctx)
	return
}

// UpdateRepository description:
// update repository
func (c *Client) UpdateRepository(ctx context.Context, xTenant string, registry string, project string, repository string, updateRepositoryReq *UpdateRepositoryReq) (repository1 *Repository, err error) {
	repository1 = new(Repository)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateRepository").
		Header("X-Tenant", xTenant).
		Query("Registry", registry).
		Query("Project", project).
		Query("Repository", repository).
		Body("application/json", updateRepositoryReq).
		TOPRPCData(repository1).
		Do(ctx)
	return
}

// UploadImage description:
// upload image
func (c *Client) UploadImage(ctx context.Context, xTenant string, xUploadID string, registry string, project string, ioReader io.Reader) (imageUploadStatsResp *ImageUploadStatsResp, err error) {
	imageUploadStatsResp = new(ImageUploadStatsResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UploadImage").
		Header("X-Tenant", xTenant).
		Header("X-Upload-ID", xUploadID).
		Query("Registry", registry).
		Query("Project", project).
		Body("application/octet-stream", ioReader).
		TOPRPCData(imageUploadStatsResp).
		Do(ctx)
	return
}

// healthcheck description:
// health check
func (c *Client) healthcheck(ctx context.Context) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=healthcheck").
		Do(ctx)
	return
}
