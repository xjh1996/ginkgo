package v20201010

import (
	"context"
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	delegation "github.com/caicloud/cyclone/pkg/workflow/workload/delegation"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// CreateCICDConfig description:
	// Create the configuration for tenant
	CreateCICDConfig(ctx context.Context, xTenant string, setting *Setting) (setting1 *Setting, err error)
	// CreateIntegration description:
	// Create integration
	CreateIntegration(ctx context.Context, xTenant string, xPublic bool, integration *Integration, xDryRun bool) (integration1 *Integration, err error)
	// CreateJobTemplate description:
	// Create job template for the tenant
	CreateJobTemplate(ctx context.Context, xTenant string, jobTemplate *JobTemplate) (jobTemplate1 *JobTemplate, err error)
	// CreatePipelineRecord description:
	// Create record of the pipeline
	CreatePipelineRecord(ctx context.Context, xTenant string, xUser string, workspace string, pipeline string, record *Record, xDryRun bool) (record1 *Record, err error)
	// CreateWorkspace description:
	// Create workspace for the tenant
	CreateWorkspace(ctx context.Context, xTenant string, xUser string, workspace *Workspace) (workspace1 *Workspace, err error)
	// CreateWorkspacePipeline description:
	// Create pipeline of the workspace for the tenant
	CreateWorkspacePipeline(ctx context.Context, xTenant string, xUser string, workspace string, pipeline *Pipeline) (pipeline1 *Pipeline, err error)
	// DeleteCICDConfig description:
	// Delete the configuration for CI/CD of the tenant, it will delete ALL CI/CD resources(things) of the tenant
	DeleteCICDConfig(ctx context.Context, xTenant string) (err error)
	// DeleteIntegration description:
	// Delete integration
	DeleteIntegration(ctx context.Context, xTenant string, integration string, xPublic bool) (err error)
	// DeleteJobTemplate description:
	// Delete job template for the tenant
	DeleteJobTemplate(ctx context.Context, xTenant string, jobTemplate string) (err error)
	// DeletePipelineRecord description:
	// Delete record of the pipeline
	DeletePipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (err error)
	// DeletePipelineRecordArtifact description:
	// Delete a specific artifact
	DeletePipelineRecordArtifact(ctx context.Context, workspace string, pipeline string, record string, artifact string, xTenant string, stage string) (err error)
	// DeleteWorkspace description:
	// Delete workspace for the tenant
	DeleteWorkspace(ctx context.Context, xTenant string, workspace string) (err error)
	// DeleteWorkspacePipeline description:
	// Delete pipeline of the workspace for the tenant
	DeleteWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string) (err error)
	// GetCICDConfig description:
	// Get the configuration for tenant
	GetCICDConfig(ctx context.Context, xTenant string) (setting *Setting, err error)
	// GetCICDSummary description:
	// Get CICD summary
	GetCICDSummary(ctx context.Context, xTenant string, xUser string) (cICDSummary *CICDSummary, err error)
	// GetExecutionContext description:
	// Get execution context for a tenant
	GetExecutionContext(ctx context.Context, xTenant string) (executionContext *ExecutionContext, err error)
	// GetIntegration description:
	// Get integration
	GetIntegration(ctx context.Context, xTenant string, integration string, xPublic bool) (integration1 *Integration, err error)
	// GetJobTemplate description:
	// Get job template for the tenant
	GetJobTemplate(ctx context.Context, xTenant string, jobTemplate string) (jobTemplate1 *JobTemplate, err error)
	// GetPipelineRecord description:
	// Get record of the pipeline
	GetPipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (record1 *Record, err error)
	// GetPipelineRecordArtifact description:
	// Download a specific artifact
	GetPipelineRecordArtifact(ctx context.Context, workspace string, pipeline string, record string, artifact string, xTenant string, stage string) (err error)
	// GetPipelineRecordContainerLogStream description:
	// Get log stream of container
	GetPipelineRecordContainerLogStream(ctx context.Context, workspace string, pipeline string, record string, xTenant string, stage string, tenant string) (err error)
	// GetPipelineRecordLogs description:
	// Get record logs of the pipeline
	GetPipelineRecordLogs(ctx context.Context, workspace string, pipeline string, record string, xTenant string, stage string, download bool) (err error)
	// GetWorkspace description:
	// Get workspace for the tenant
	GetWorkspace(ctx context.Context, xTenant string, workspace string, countUsingCacheRecord bool) (workspace1 *Workspace, err error)
	// GetWorkspacePipeline description:
	// Get pipeline of the workspace for the tenant
	GetWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string, recentCountParams RecentCountParams) (pipeline1 *Pipeline, err error)
	// GetWorkspacePipelineStats description:
	// get statistics of the pipeline
	GetWorkspacePipelineStats(ctx context.Context, xTenant string, workspace string, pipeline string, startTime string, endTime string) (statistic *Statistic, err error)
	// GetWorkspaceStats description:
	// get statistics of the workspace
	GetWorkspaceStats(ctx context.Context, xTenant string, workspace string, startTime string, endTime string) (statistic *Statistic, err error)
	// GetWorkspaceTemplateType description:
	// Get template type of the repo for the project
	GetWorkspaceTemplateType(ctx context.Context, xTenant string, workspace string, repo string) (templateType *TemplateType, err error)
	// HandleTenantWebhook description:
	// handle webhooks from integrated systems
	HandleTenantWebhook(ctx context.Context, tenant string, sourceType string, integration string) (webhookResponse *WebhookResponse, err error)
	// ListIntegrations description:
	// List integrations
	ListIntegrations(ctx context.Context, xTenant string, includePublic bool, paginationParams PaginationParams) (integrationList *IntegrationList, err error)
	// ListJobTemplates description:
	// List job templates for the tenant
	ListJobTemplates(ctx context.Context, xTenant string, paginationParams PaginationParams) (jobTemplateList *JobTemplateList, err error)
	// ListPipelineRecordArtifacts description:
	// List artifacts
	ListPipelineRecordArtifacts(ctx context.Context, workspace string, pipeline string, record string, xTenant string) (stageArtifactList *StageArtifactList, err error)
	// ListPipelineRecords description:
	// List records of the pipeline
	ListPipelineRecords(ctx context.Context, xTenant string, workspace string, pipeline string, paginationParams PaginationParams) (recordList *RecordList, err error)
	// ListPipelineTemplates description:
	// List pipeline templates
	ListPipelineTemplates(ctx context.Context) (pipelineList *PipelineList, err error)
	// ListSonarQubeLanguages description:
	// List SonarQube supported languages
	ListSonarQubeLanguages(ctx context.Context, xTenant string, integration string) (languageList *LanguageList, err error)
	// ListSonarQubeQualityGates description:
	// List SonarQube quality gates
	ListSonarQubeQualityGates(ctx context.Context, xTenant string, integration string) (qualityGateList *QualityGateList, err error)
	// ListWorkingPods description:
	// list working pods
	ListWorkingPods(ctx context.Context, xTenant string, paginationParams PaginationParams) (podList *PodList, err error)
	// ListWorkspaceBranches description:
	// List SCM branches for specified SCM repo accessible by workspace for the tenant
	ListWorkspaceBranches(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error)
	// ListWorkspaceDockerfiles description:
	// List SCM tags for specified SCM repo accessible by workspace for the tenant
	ListWorkspaceDockerfiles(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error)
	// ListWorkspacePipelines description:
	// List pipelines of the workspaces for the tenant
	ListWorkspacePipelines(ctx context.Context, xTenant string, workspace string, recentCountParams RecentCountParams) (pipelineList *PipelineList, err error)
	// ListWorkspacePullRequests description:
	// List SCM pull requests for specified SCM repo accessible by workspace for the tenant
	ListWorkspacePullRequests(ctx context.Context, xTenant string, workspace string, repo string, state string) (pullRequestList *PullRequestList, err error)
	// ListWorkspaceRepos description:
	// List SCM repos accessible by workspace for the tenant
	ListWorkspaceRepos(ctx context.Context, xTenant string, workspace string) (repositoryList *RepositoryList, err error)
	// ListWorkspaceTags description:
	// List SCM dockerfiles for specified SCM repo accessible by workspace for the tenant
	ListWorkspaceTags(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error)
	// ListWorkspaces description:
	// List workspaces for the tenant
	ListWorkspaces(ctx context.Context, xTenant string, paginationParams PaginationParams) (workspaceList *WorkspaceList, err error)
	// NotifyCargo description:
	// Notify that a change happened in docker registry
	NotifyCargo(ctx context.Context, notification *CargoNotification) (err error)
	// ReceiveNotifications description:
	// Receive pipeline record notifications and send out their results
	ReceiveNotifications(ctx context.Context, v1alpha1WorkflowRun *v1alpha1.WorkflowRun) (err error)
	// RunStage description:
	// Run Cyclone's special stages, e.g. CD, Approval
	RunStage(ctx context.Context, request *delegation.Request) (err error)
	// StopPipelineRecord description:
	// Stop a pipeline execution
	StopPipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (err error)
	// TriggerCleanCacheTask description:
	// Trigger a task to clean up dependency cache
	TriggerCleanCacheTask(ctx context.Context, xTenant string, workspace string) (accelerationCacheCleanupStatus *AccelerationCacheCleanupStatus, err error)
	// UpdateCICDConfig description:
	// Update the configuration for tenant
	UpdateCICDConfig(ctx context.Context, xTenant string, setting *Setting) (setting1 *Setting, err error)
	// UpdateIntegration description:
	// Update integration
	UpdateIntegration(ctx context.Context, xTenant string, integration string, xPublic bool, integration1 *Integration) (integration2 *Integration, err error)
	// UpdateJobTemplate description:
	// Update job template for the tenant
	UpdateJobTemplate(ctx context.Context, xTenant string, jobTemplate string, jobTemplate1 *JobTemplate) (jobTemplate2 *JobTemplate, err error)
	// UpdatePipelineApprovalStatus description:
	// Update approval status
	UpdatePipelineApprovalStatus(ctx context.Context, tenant string, record string, stage string, operation string) (string string, err error)
	// UpdateWorkspace does not have any description.
	UpdateWorkspace(ctx context.Context, xTenant string, workspace string, workspace1 *Workspace) (workspace2 *Workspace, err error)
	// UpdateWorkspaceBasicConfig description:
	// Update workspace basic info for the tenant
	UpdateWorkspaceBasicConfig(ctx context.Context, xTenant string, workspace string, workspace1 *Workspace) (workspace2 *Workspace, err error)
	// UpdateWorkspacePipeline description:
	// Update pipeline of the workspace for the tenant
	UpdateWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string, pipeline1 *Pipeline) (pipeline2 *Pipeline, err error)
	// ValidateResource description:
	// Validate a resource
	ValidateResource(ctx context.Context, xResourceType string, integration *Integration) (bool bool, err error)
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

// CreateCICDConfig description:
// Create the configuration for tenant
func (c *Client) CreateCICDConfig(ctx context.Context, xTenant string, setting *Setting) (setting1 *Setting, err error) {
	setting1 = new(Setting)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateCICDConfig").
		Header("X-Tenant", xTenant).
		Body("application/json", setting).
		TOPRPCData(setting1).
		Do(ctx)
	return
}

// CreateIntegration description:
// Create integration
func (c *Client) CreateIntegration(ctx context.Context, xTenant string, xPublic bool, integration *Integration, xDryRun bool) (integration1 *Integration, err error) {
	integration1 = new(Integration)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateIntegration").
		Header("X-Tenant", xTenant).
		Header("X-Public", xPublic).
		Body("application/json", integration).
		Header("X-Dry-Run", xDryRun).
		TOPRPCData(integration1).
		Do(ctx)
	return
}

// CreateJobTemplate description:
// Create job template for the tenant
func (c *Client) CreateJobTemplate(ctx context.Context, xTenant string, jobTemplate *JobTemplate) (jobTemplate1 *JobTemplate, err error) {
	jobTemplate1 = new(JobTemplate)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateJobTemplate").
		Header("X-Tenant", xTenant).
		Body("application/json", jobTemplate).
		TOPRPCData(jobTemplate1).
		Do(ctx)
	return
}

// CreatePipelineRecord description:
// Create record of the pipeline
func (c *Client) CreatePipelineRecord(ctx context.Context, xTenant string, xUser string, workspace string, pipeline string, record *Record, xDryRun bool) (record1 *Record, err error) {
	record1 = new(Record)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreatePipelineRecord").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Body("application/json", record).
		Header("X-Dry-Run", xDryRun).
		TOPRPCData(record1).
		Do(ctx)
	return
}

// CreateWorkspace description:
// Create workspace for the tenant
func (c *Client) CreateWorkspace(ctx context.Context, xTenant string, xUser string, workspace *Workspace) (workspace1 *Workspace, err error) {
	workspace1 = new(Workspace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateWorkspace").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Body("application/json", workspace).
		TOPRPCData(workspace1).
		Do(ctx)
	return
}

// CreateWorkspacePipeline description:
// Create pipeline of the workspace for the tenant
func (c *Client) CreateWorkspacePipeline(ctx context.Context, xTenant string, xUser string, workspace string, pipeline *Pipeline) (pipeline1 *Pipeline, err error) {
	pipeline1 = new(Pipeline)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateWorkspacePipeline").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		Query("Workspace", workspace).
		Body("application/json", pipeline).
		TOPRPCData(pipeline1).
		Do(ctx)
	return
}

// DeleteCICDConfig description:
// Delete the configuration for CI/CD of the tenant, it will delete ALL CI/CD resources(things) of the tenant
func (c *Client) DeleteCICDConfig(ctx context.Context, xTenant string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteCICDConfig").
		Header("X-Tenant", xTenant).
		Do(ctx)
	return
}

// DeleteIntegration description:
// Delete integration
func (c *Client) DeleteIntegration(ctx context.Context, xTenant string, integration string, xPublic bool) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteIntegration").
		Header("X-Tenant", xTenant).
		Query("Integration", integration).
		Header("X-Public", xPublic).
		Do(ctx)
	return
}

// DeleteJobTemplate description:
// Delete job template for the tenant
func (c *Client) DeleteJobTemplate(ctx context.Context, xTenant string, jobTemplate string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteJobTemplate").
		Header("X-Tenant", xTenant).
		Query("JobTemplate", jobTemplate).
		Do(ctx)
	return
}

// DeletePipelineRecord description:
// Delete record of the pipeline
func (c *Client) DeletePipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePipelineRecord").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Do(ctx)
	return
}

// DeletePipelineRecordArtifact description:
// Delete a specific artifact
func (c *Client) DeletePipelineRecordArtifact(ctx context.Context, workspace string, pipeline string, record string, artifact string, xTenant string, stage string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePipelineRecordArtifact").
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Query("Artifact", artifact).
		Header("X-Tenant", xTenant).
		Query("Stage", stage).
		Do(ctx)
	return
}

// DeleteWorkspace description:
// Delete workspace for the tenant
func (c *Client) DeleteWorkspace(ctx context.Context, xTenant string, workspace string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteWorkspace").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Do(ctx)
	return
}

// DeleteWorkspacePipeline description:
// Delete pipeline of the workspace for the tenant
func (c *Client) DeleteWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteWorkspacePipeline").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Do(ctx)
	return
}

// GetCICDConfig description:
// Get the configuration for tenant
func (c *Client) GetCICDConfig(ctx context.Context, xTenant string) (setting *Setting, err error) {
	setting = new(Setting)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCICDConfig").
		Header("X-Tenant", xTenant).
		TOPRPCData(setting).
		Do(ctx)
	return
}

// GetCICDSummary description:
// Get CICD summary
func (c *Client) GetCICDSummary(ctx context.Context, xTenant string, xUser string) (cICDSummary *CICDSummary, err error) {
	cICDSummary = new(CICDSummary)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCICDSummary").
		Header("X-Tenant", xTenant).
		Header("X-User", xUser).
		TOPRPCData(cICDSummary).
		Do(ctx)
	return
}

// GetExecutionContext description:
// Get execution context for a tenant
func (c *Client) GetExecutionContext(ctx context.Context, xTenant string) (executionContext *ExecutionContext, err error) {
	executionContext = new(ExecutionContext)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetExecutionContext").
		Header("X-Tenant", xTenant).
		TOPRPCData(executionContext).
		Do(ctx)
	return
}

// GetIntegration description:
// Get integration
func (c *Client) GetIntegration(ctx context.Context, xTenant string, integration string, xPublic bool) (integration1 *Integration, err error) {
	integration1 = new(Integration)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetIntegration").
		Header("X-Tenant", xTenant).
		Query("Integration", integration).
		Header("X-Public", xPublic).
		TOPRPCData(integration1).
		Do(ctx)
	return
}

// GetJobTemplate description:
// Get job template for the tenant
func (c *Client) GetJobTemplate(ctx context.Context, xTenant string, jobTemplate string) (jobTemplate1 *JobTemplate, err error) {
	jobTemplate1 = new(JobTemplate)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetJobTemplate").
		Header("X-Tenant", xTenant).
		Query("JobTemplate", jobTemplate).
		TOPRPCData(jobTemplate1).
		Do(ctx)
	return
}

// GetPipelineRecord description:
// Get record of the pipeline
func (c *Client) GetPipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (record1 *Record, err error) {
	record1 = new(Record)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPipelineRecord").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		TOPRPCData(record1).
		Do(ctx)
	return
}

// GetPipelineRecordArtifact description:
// Download a specific artifact
func (c *Client) GetPipelineRecordArtifact(ctx context.Context, workspace string, pipeline string, record string, artifact string, xTenant string, stage string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPipelineRecordArtifact").
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Query("Artifact", artifact).
		Header("X-Tenant", xTenant).
		Query("Stage", stage).
		Do(ctx)
	return
}

// GetPipelineRecordContainerLogStream description:
// Get log stream of container
func (c *Client) GetPipelineRecordContainerLogStream(ctx context.Context, workspace string, pipeline string, record string, xTenant string, stage string, tenant string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPipelineRecordContainerLogStream").
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Header("X-Tenant", xTenant).
		Query("Stage", stage).
		Query("Tenant", tenant).
		Do(ctx)
	return
}

// GetPipelineRecordLogs description:
// Get record logs of the pipeline
func (c *Client) GetPipelineRecordLogs(ctx context.Context, workspace string, pipeline string, record string, xTenant string, stage string, download bool) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPipelineRecordLogs").
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Header("X-Tenant", xTenant).
		Query("Stage", stage).
		Query("Download", download).
		Do(ctx)
	return
}

// GetWorkspace description:
// Get workspace for the tenant
func (c *Client) GetWorkspace(ctx context.Context, xTenant string, workspace string, countUsingCacheRecord bool) (workspace1 *Workspace, err error) {
	workspace1 = new(Workspace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkspace").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("CountUsingCacheRecord", countUsingCacheRecord).
		TOPRPCData(workspace1).
		Do(ctx)
	return
}

// GetWorkspacePipeline description:
// Get pipeline of the workspace for the tenant
func (c *Client) GetWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string, recentCountParams RecentCountParams) (pipeline1 *Pipeline, err error) {
	pipeline1 = new(Pipeline)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkspacePipeline").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("RecentCount", recentCountParams.All).
		Query("RecentSuccessCount", recentCountParams.Success).
		Query("RecentFailedCount", recentCountParams.Failed).
		Query("Sort", recentCountParams.Sort).
		TOPRPCData(pipeline1).
		Do(ctx)
	return
}

// GetWorkspacePipelineStats description:
// get statistics of the pipeline
func (c *Client) GetWorkspacePipelineStats(ctx context.Context, xTenant string, workspace string, pipeline string, startTime string, endTime string) (statistic *Statistic, err error) {
	statistic = new(Statistic)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkspacePipelineStats").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("StartTime", startTime).
		Query("EndTime", endTime).
		TOPRPCData(statistic).
		Do(ctx)
	return
}

// GetWorkspaceStats description:
// get statistics of the workspace
func (c *Client) GetWorkspaceStats(ctx context.Context, xTenant string, workspace string, startTime string, endTime string) (statistic *Statistic, err error) {
	statistic = new(Statistic)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkspaceStats").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("StartTime", startTime).
		Query("EndTime", endTime).
		TOPRPCData(statistic).
		Do(ctx)
	return
}

// GetWorkspaceTemplateType description:
// Get template type of the repo for the project
func (c *Client) GetWorkspaceTemplateType(ctx context.Context, xTenant string, workspace string, repo string) (templateType *TemplateType, err error) {
	templateType = new(TemplateType)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkspaceTemplateType").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Repo", repo).
		TOPRPCData(templateType).
		Do(ctx)
	return
}

// HandleTenantWebhook description:
// handle webhooks from integrated systems
func (c *Client) HandleTenantWebhook(ctx context.Context, tenant string, sourceType string, integration string) (webhookResponse *WebhookResponse, err error) {
	webhookResponse = new(WebhookResponse)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=HandleTenantWebhook").
		Query("Tenant", tenant).
		Query("SourceType", sourceType).
		Query("Integration", integration).
		TOPRPCData(webhookResponse).
		Do(ctx)
	return
}

// ListIntegrations description:
// List integrations
func (c *Client) ListIntegrations(ctx context.Context, xTenant string, includePublic bool, paginationParams PaginationParams) (integrationList *IntegrationList, err error) {
	integrationList = new(IntegrationList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListIntegrations").
		Header("X-Tenant", xTenant).
		Query("IncludePublic", includePublic).
		Query("Start", paginationParams.Start).
		Query("Limit", paginationParams.Limit).
		Query("Filter", paginationParams.Filter).
		Query("Sort", paginationParams.Sort).
		Query("Ascending", paginationParams.Ascending).
		Query("Detail", paginationParams.Detail).
		TOPRPCData(integrationList).
		Do(ctx)
	return
}

// ListJobTemplates description:
// List job templates for the tenant
func (c *Client) ListJobTemplates(ctx context.Context, xTenant string, paginationParams PaginationParams) (jobTemplateList *JobTemplateList, err error) {
	jobTemplateList = new(JobTemplateList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListJobTemplates").
		Header("X-Tenant", xTenant).
		Query("Start", paginationParams.Start).
		Query("Limit", paginationParams.Limit).
		Query("Filter", paginationParams.Filter).
		Query("Sort", paginationParams.Sort).
		Query("Ascending", paginationParams.Ascending).
		Query("Detail", paginationParams.Detail).
		TOPRPCData(jobTemplateList).
		Do(ctx)
	return
}

// ListPipelineRecordArtifacts description:
// List artifacts
func (c *Client) ListPipelineRecordArtifacts(ctx context.Context, workspace string, pipeline string, record string, xTenant string) (stageArtifactList *StageArtifactList, err error) {
	stageArtifactList = new(StageArtifactList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPipelineRecordArtifacts").
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Header("X-Tenant", xTenant).
		TOPRPCData(stageArtifactList).
		Do(ctx)
	return
}

// ListPipelineRecords description:
// List records of the pipeline
func (c *Client) ListPipelineRecords(ctx context.Context, xTenant string, workspace string, pipeline string, paginationParams PaginationParams) (recordList *RecordList, err error) {
	recordList = new(RecordList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPipelineRecords").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Start", paginationParams.Start).
		Query("Limit", paginationParams.Limit).
		Query("Filter", paginationParams.Filter).
		Query("Sort", paginationParams.Sort).
		Query("Ascending", paginationParams.Ascending).
		Query("Detail", paginationParams.Detail).
		TOPRPCData(recordList).
		Do(ctx)
	return
}

// ListPipelineTemplates description:
// List pipeline templates
func (c *Client) ListPipelineTemplates(ctx context.Context) (pipelineList *PipelineList, err error) {
	pipelineList = new(PipelineList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPipelineTemplates").
		TOPRPCData(pipelineList).
		Do(ctx)
	return
}

// ListSonarQubeLanguages description:
// List SonarQube supported languages
func (c *Client) ListSonarQubeLanguages(ctx context.Context, xTenant string, integration string) (languageList *LanguageList, err error) {
	languageList = new(LanguageList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListSonarQubeLanguages").
		Header("X-Tenant", xTenant).
		Query("Integration", integration).
		TOPRPCData(languageList).
		Do(ctx)
	return
}

// ListSonarQubeQualityGates description:
// List SonarQube quality gates
func (c *Client) ListSonarQubeQualityGates(ctx context.Context, xTenant string, integration string) (qualityGateList *QualityGateList, err error) {
	qualityGateList = new(QualityGateList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListSonarQubeQualityGates").
		Header("X-Tenant", xTenant).
		Query("Integration", integration).
		TOPRPCData(qualityGateList).
		Do(ctx)
	return
}

// ListWorkingPods description:
// list working pods
func (c *Client) ListWorkingPods(ctx context.Context, xTenant string, paginationParams PaginationParams) (podList *PodList, err error) {
	podList = new(PodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkingPods").
		Header("X-Tenant", xTenant).
		Query("Start", paginationParams.Start).
		Query("Limit", paginationParams.Limit).
		Query("Filter", paginationParams.Filter).
		Query("Sort", paginationParams.Sort).
		Query("Ascending", paginationParams.Ascending).
		Query("Detail", paginationParams.Detail).
		TOPRPCData(podList).
		Do(ctx)
	return
}

// ListWorkspaceBranches description:
// List SCM branches for specified SCM repo accessible by workspace for the tenant
func (c *Client) ListWorkspaceBranches(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error) {
	stringList = new(StringList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspaceBranches").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Repo", repo).
		TOPRPCData(stringList).
		Do(ctx)
	return
}

// ListWorkspaceDockerfiles description:
// List SCM tags for specified SCM repo accessible by workspace for the tenant
func (c *Client) ListWorkspaceDockerfiles(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error) {
	stringList = new(StringList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspaceDockerfiles").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Repo", repo).
		TOPRPCData(stringList).
		Do(ctx)
	return
}

// ListWorkspacePipelines description:
// List pipelines of the workspaces for the tenant
func (c *Client) ListWorkspacePipelines(ctx context.Context, xTenant string, workspace string, recentCountParams RecentCountParams) (pipelineList *PipelineList, err error) {
	pipelineList = new(PipelineList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspacePipelines").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("RecentCount", recentCountParams.All).
		Query("RecentSuccessCount", recentCountParams.Success).
		Query("RecentFailedCount", recentCountParams.Failed).
		Query("Sort", recentCountParams.Sort).
		TOPRPCData(pipelineList).
		Do(ctx)
	return
}

// ListWorkspacePullRequests description:
// List SCM pull requests for specified SCM repo accessible by workspace for the tenant
func (c *Client) ListWorkspacePullRequests(ctx context.Context, xTenant string, workspace string, repo string, state string) (pullRequestList *PullRequestList, err error) {
	pullRequestList = new(PullRequestList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspacePullRequests").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Repo", repo).
		Query("State", state).
		TOPRPCData(pullRequestList).
		Do(ctx)
	return
}

// ListWorkspaceRepos description:
// List SCM repos accessible by workspace for the tenant
func (c *Client) ListWorkspaceRepos(ctx context.Context, xTenant string, workspace string) (repositoryList *RepositoryList, err error) {
	repositoryList = new(RepositoryList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspaceRepos").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		TOPRPCData(repositoryList).
		Do(ctx)
	return
}

// ListWorkspaceTags description:
// List SCM dockerfiles for specified SCM repo accessible by workspace for the tenant
func (c *Client) ListWorkspaceTags(ctx context.Context, xTenant string, workspace string, repo string) (stringList *StringList, err error) {
	stringList = new(StringList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspaceTags").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Repo", repo).
		TOPRPCData(stringList).
		Do(ctx)
	return
}

// ListWorkspaces description:
// List workspaces for the tenant
func (c *Client) ListWorkspaces(ctx context.Context, xTenant string, paginationParams PaginationParams) (workspaceList *WorkspaceList, err error) {
	workspaceList = new(WorkspaceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkspaces").
		Header("X-Tenant", xTenant).
		Query("Start", paginationParams.Start).
		Query("Limit", paginationParams.Limit).
		Query("Filter", paginationParams.Filter).
		Query("Sort", paginationParams.Sort).
		Query("Ascending", paginationParams.Ascending).
		Query("Detail", paginationParams.Detail).
		TOPRPCData(workspaceList).
		Do(ctx)
	return
}

// NotifyCargo description:
// Notify that a change happened in docker registry
func (c *Client) NotifyCargo(ctx context.Context, notification *CargoNotification) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=NotifyCargo").
		Body("application/json", notification).
		Do(ctx)
	return
}

// ReceiveNotifications description:
// Receive pipeline record notifications and send out their results
func (c *Client) ReceiveNotifications(ctx context.Context, v1alpha1WorkflowRun *v1alpha1.WorkflowRun) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ReceiveNotifications").
		Body("application/json", v1alpha1WorkflowRun).
		Do(ctx)
	return
}

// RunStage description:
// Run Cyclone's special stages, e.g. CD, Approval
func (c *Client) RunStage(ctx context.Context, request *delegation.Request) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RunStage").
		Body("application/json", request).
		Do(ctx)
	return
}

// StopPipelineRecord description:
// Stop a pipeline execution
func (c *Client) StopPipelineRecord(ctx context.Context, xTenant string, workspace string, pipeline string, record string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=StopPipelineRecord").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Query("Record", record).
		Do(ctx)
	return
}

// TriggerCleanCacheTask description:
// Trigger a task to clean up dependency cache
func (c *Client) TriggerCleanCacheTask(ctx context.Context, xTenant string, workspace string) (accelerationCacheCleanupStatus *AccelerationCacheCleanupStatus, err error) {
	accelerationCacheCleanupStatus = new(AccelerationCacheCleanupStatus)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=TriggerCleanCacheTask").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		TOPRPCData(accelerationCacheCleanupStatus).
		Do(ctx)
	return
}

// UpdateCICDConfig description:
// Update the configuration for tenant
func (c *Client) UpdateCICDConfig(ctx context.Context, xTenant string, setting *Setting) (setting1 *Setting, err error) {
	setting1 = new(Setting)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateCICDConfig").
		Header("X-Tenant", xTenant).
		Body("application/json", setting).
		TOPRPCData(setting1).
		Do(ctx)
	return
}

// UpdateIntegration description:
// Update integration
func (c *Client) UpdateIntegration(ctx context.Context, xTenant string, integration string, xPublic bool, integration1 *Integration) (integration2 *Integration, err error) {
	integration2 = new(Integration)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateIntegration").
		Header("X-Tenant", xTenant).
		Query("Integration", integration).
		Header("X-Public", xPublic).
		Body("application/json", integration1).
		TOPRPCData(integration2).
		Do(ctx)
	return
}

// UpdateJobTemplate description:
// Update job template for the tenant
func (c *Client) UpdateJobTemplate(ctx context.Context, xTenant string, jobTemplate string, jobTemplate1 *JobTemplate) (jobTemplate2 *JobTemplate, err error) {
	jobTemplate2 = new(JobTemplate)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateJobTemplate").
		Header("X-Tenant", xTenant).
		Query("JobTemplate", jobTemplate).
		Body("application/json", jobTemplate1).
		TOPRPCData(jobTemplate2).
		Do(ctx)
	return
}

// UpdatePipelineApprovalStatus description:
// Update approval status
func (c *Client) UpdatePipelineApprovalStatus(ctx context.Context, tenant string, record string, stage string, operation string) (string string, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdatePipelineApprovalStatus").
		Query("Tenant", tenant).
		Query("Record", record).
		Query("Stage", stage).
		Query("Operation", operation).
		TOPRPCData(&string).
		Do(ctx)
	return
}

// UpdateWorkspace does not have any description.
func (c *Client) UpdateWorkspace(ctx context.Context, xTenant string, workspace string, workspace1 *Workspace) (workspace2 *Workspace, err error) {
	workspace2 = new(Workspace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateWorkspace").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Body("application/json", workspace1).
		TOPRPCData(workspace2).
		Do(ctx)
	return
}

// UpdateWorkspaceBasicConfig description:
// Update workspace basic info for the tenant
func (c *Client) UpdateWorkspaceBasicConfig(ctx context.Context, xTenant string, workspace string, workspace1 *Workspace) (workspace2 *Workspace, err error) {
	workspace2 = new(Workspace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateWorkspaceBasicConfig").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Body("application/json", workspace1).
		TOPRPCData(workspace2).
		Do(ctx)
	return
}

// UpdateWorkspacePipeline description:
// Update pipeline of the workspace for the tenant
func (c *Client) UpdateWorkspacePipeline(ctx context.Context, xTenant string, workspace string, pipeline string, pipeline1 *Pipeline) (pipeline2 *Pipeline, err error) {
	pipeline2 = new(Pipeline)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateWorkspacePipeline").
		Header("X-Tenant", xTenant).
		Query("Workspace", workspace).
		Query("Pipeline", pipeline).
		Body("application/json", pipeline1).
		TOPRPCData(pipeline2).
		Do(ctx)
	return
}

// ValidateResource description:
// Validate a resource
func (c *Client) ValidateResource(ctx context.Context, xResourceType string, integration *Integration) (bool bool, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ValidateResource").
		Header("X-Resource-Type", xResourceType).
		Body("application/json", integration).
		TOPRPCData(&bool).
		Do(ctx)
	return
}
