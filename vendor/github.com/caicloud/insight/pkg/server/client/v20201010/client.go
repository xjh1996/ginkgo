package v20201010

import (
	"context"
	io "io"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// BindLogEndpointCluster description:
	// Update a log endpoint
	BindLogEndpointCluster(ctx context.Context, binding *Binding) (bindingResults *BindingResults, err error)
	// CreateDashboard does not have any description.
	CreateDashboard(ctx context.Context, dashboard *Dashboard) (dashboard1 *Dashboard, err error)
	// CreateGraph does not have any description.
	CreateGraph(ctx context.Context, graph *Graph) (graph1 *Graph, err error)
	// CreateLogEndpoint description:
	// Create a log endpoint
	CreateLogEndpoint(ctx context.Context, logEndpoint *LogEndpoint) (bindingResults *BindingResults, err error)
	// DeleteDashboard does not have any description.
	DeleteDashboard(ctx context.Context, dashboardId string) (err error)
	// DeleteGraph does not have any description.
	DeleteGraph(ctx context.Context, graph string) (err error)
	// DeleteLogEndpoint description:
	// Delete a log endpoint
	DeleteLogEndpoint(ctx context.Context, name string, namespace string) (err error)
	// GetClustersResourceStats description:
	// Get statistics on Clusters resource usage
	GetClustersResourceStats(ctx context.Context, filters *ClustersFilters) (clusterResourceStates []ClusterResourceStats, err error)
	// GetContainerResourceStats description:
	// Get statistics on container resource usage
	GetContainerResourceStats(ctx context.Context, filters *ContainerFilters, statsOptions *ContainerStatsOptions) (containerResourceStates []ContainerResourceStats, err error)
	// GetDashboard does not have any description.
	GetDashboard(ctx context.Context, dashboardId string) (dashboard *Dashboard, err error)
	// GetLoadBalancerStats description:
	// Get statistics on LoadBalancer operation
	GetLoadBalancerStats(ctx context.Context, filters *LoadBalancerFilters) (loadBalancerStats *LoadBalancerStats, err error)
	// GetLogEndpoint description:
	// Get a log endpoint
	GetLogEndpoint(ctx context.Context, name string, namespace string) (logEndpoint *LogEndpoint, err error)
	// GetLogEndpointHealthStatus description:
	// get log endpoint status
	GetLogEndpointHealthStatus(ctx context.Context, cluster string, namespace string, timeout int, verbose int) (healthStatus *HealthStatus, err error)
	// GetLogFiles description:
	// Retrieve collected log files for the given pod
	GetLogFiles(ctx context.Context, cid string, namespace string, pod string) (fileGroups *FileGroups, err error)
	// GetLogStream description:
	// Stream logs by filters
	GetLogStream(ctx context.Context, options *StreamOptions, k8s bool) (err error)
	// GetMetric does not have any description.
	GetMetric(ctx context.Context, metricID string) (metric *Metric, err error)
	// GetNodeCPUStats description:
	// Get statistics on Nodes CPU usage
	GetNodeCPUStats(ctx context.Context, filters *ClustersFilters, listOptions *ListOptions) (nodeCPUStates []NodeCPUStats, err error)
	// GetSystemLogStream description:
	// Stream system component logs by filters
	GetSystemLogStream(ctx context.Context, options *SystemComponentOptions) (err error)
	// ListDashboardGraphs does not have any description.
	ListDashboardGraphs(ctx context.Context, graphIDs *GraphIDs) (graphList *GraphList, err error)
	// ListDashboards does not have any description.
	ListDashboards(ctx context.Context, filter *DashboardFilter) (dashboardList *DashboardList, err error)
	// ListEvents description:
	// Get events by filters
	ListEvents(ctx context.Context, options *EventListOptions) (eventListResult *EventListResult, err error)
	// ListGraphs does not have any description.
	ListGraphs(ctx context.Context, filter *GraphFilter) (graphList *GraphList, err error)
	// ListLabelValues description:
	// Get label values from label name with given expression
	ListLabelValues(ctx context.Context, labelName string, expr string) (labelValueResult *LabelValueResult, err error)
	// ListLogEndpoint description:
	// List log endpoints
	ListLogEndpoint(ctx context.Context, keyword string) (listResult *ListResult, err error)
	// ListMetrics does not have any description.
	ListMetrics(ctx context.Context) (metricList *MetricList, err error)
	// ListUnitSeries does not have any description.
	ListUnitSeries(ctx context.Context, includeUnit *string) (unitSeriesList *UnitSeriesList, err error)
	// QueryRange description:
	// Prometheus query range
	QueryRange(ctx context.Context, promQuery *PromQuery) (stats *Stats, err error)
	// SearchEvents description:
	// Get events by filters
	SearchEvents(ctx context.Context, options *EventSearchOptions) (eventSearchResult *EventSearchResult, err error)
	// SearchLog description:
	// Get logs by filters
	SearchLog(ctx context.Context, options *LogSearchOptions) (resourceLogItems ResourceLogItems, err error)
	// SearchLogContexts description:
	// Get logs context by filters
	SearchLogContexts(ctx context.Context, options *ContextSearchOptions) (resourceLogItems ResourceLogItems, err error)
	// SearchLogContextsRaw description:
	// Download logs context by filters
	SearchLogContextsRaw(ctx context.Context, options *ContextSearchOptions) (strings map[string]string, ioReader io.Reader, err error)
	// SearchLogRaw description:
	// Download logs by filters
	SearchLogRaw(ctx context.Context, options *LogSearchOptions) (strings map[string]string, ioReader io.Reader, err error)
	// UpdateDashboard does not have any description.
	UpdateDashboard(ctx context.Context, dashboardId string, dashboard *Dashboard) (dashboard1 *Dashboard, err error)
	// UpdateGraph does not have any description.
	UpdateGraph(ctx context.Context, graph string, graph1 *Graph) (graph2 *Graph, err error)
	// UpdateLogEndpoint description:
	// Update a log endpoint
	UpdateLogEndpoint(ctx context.Context, logEndpoint *LogEndpoint) (bindingResults *BindingResults, err error)
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

// BindLogEndpointCluster description:
// Update a log endpoint
func (c *Client) BindLogEndpointCluster(ctx context.Context, binding *Binding) (bindingResults *BindingResults, err error) {
	bindingResults = new(BindingResults)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=BindLogEndpointCluster").
		Body("application/json", binding).
		TOPRPCData(bindingResults).
		Do(ctx)
	return
}

// CreateDashboard does not have any description.
func (c *Client) CreateDashboard(ctx context.Context, dashboard *Dashboard) (dashboard1 *Dashboard, err error) {
	dashboard1 = new(Dashboard)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateDashboard").
		Body("application/json", dashboard).
		TOPRPCData(dashboard1).
		Do(ctx)
	return
}

// CreateGraph does not have any description.
func (c *Client) CreateGraph(ctx context.Context, graph *Graph) (graph1 *Graph, err error) {
	graph1 = new(Graph)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateGraph").
		Body("application/json", graph).
		TOPRPCData(graph1).
		Do(ctx)
	return
}

// CreateLogEndpoint description:
// Create a log endpoint
func (c *Client) CreateLogEndpoint(ctx context.Context, logEndpoint *LogEndpoint) (bindingResults *BindingResults, err error) {
	bindingResults = new(BindingResults)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateLogEndpoint").
		Body("application/json", logEndpoint).
		TOPRPCData(bindingResults).
		Do(ctx)
	return
}

// DeleteDashboard does not have any description.
func (c *Client) DeleteDashboard(ctx context.Context, dashboardId string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteDashboard").
		Query("DashboardId", dashboardId).
		Do(ctx)
	return
}

// DeleteGraph does not have any description.
func (c *Client) DeleteGraph(ctx context.Context, graph string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteGraph").
		Query("Graph", graph).
		Do(ctx)
	return
}

// DeleteLogEndpoint description:
// Delete a log endpoint
func (c *Client) DeleteLogEndpoint(ctx context.Context, name string, namespace string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteLogEndpoint").
		Query("Name", name).
		Query("Namespace", namespace).
		Do(ctx)
	return
}

// GetClustersResourceStats description:
// Get statistics on Clusters resource usage
func (c *Client) GetClustersResourceStats(ctx context.Context, filters *ClustersFilters) (clusterResourceStates []ClusterResourceStats, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetClustersResourceStats").
		Query("Clusters", filters.Clusters).
		Data(&clusterResourceStates).
		Do(ctx)
	return
}

// GetContainerResourceStats description:
// Get statistics on container resource usage
func (c *Client) GetContainerResourceStats(ctx context.Context, filters *ContainerFilters, statsOptions *ContainerStatsOptions) (containerResourceStates []ContainerResourceStats, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetContainerResourceStats").
		Query("Owner", filters.Owner).
		Query("OwnerKind", filters.OwnerKind).
		Query("Namespace", filters.Namespace).
		Query("Cluster", filters.Cluster).
		Query("StartTime", statsOptions.StartTime).
		Query("EndTime", statsOptions.EndTime).
		Query("Quantile", statsOptions.Quantile).
		Data(&containerResourceStates).
		Do(ctx)
	return
}

// GetDashboard does not have any description.
func (c *Client) GetDashboard(ctx context.Context, dashboardId string) (dashboard *Dashboard, err error) {
	dashboard = new(Dashboard)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDashboard").
		Query("DashboardId", dashboardId).
		TOPRPCData(dashboard).
		Do(ctx)
	return
}

// GetLoadBalancerStats description:
// Get statistics on LoadBalancer operation
func (c *Client) GetLoadBalancerStats(ctx context.Context, filters *LoadBalancerFilters) (loadBalancerStats *LoadBalancerStats, err error) {
	loadBalancerStats = new(LoadBalancerStats)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetLoadBalancerStats").
		Query("LoadBalancer", filters.LoadBalancer).
		Query("Cluster", filters.Cluster).
		Data(loadBalancerStats).
		Do(ctx)
	return
}

// GetLogEndpoint description:
// Get a log endpoint
func (c *Client) GetLogEndpoint(ctx context.Context, name string, namespace string) (logEndpoint *LogEndpoint, err error) {
	logEndpoint = new(LogEndpoint)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetLogEndpoint").
		Query("Name", name).
		Query("Namespace", namespace).
		TOPRPCData(logEndpoint).
		Do(ctx)
	return
}

// GetLogEndpointHealthStatus description:
// get log endpoint status
func (c *Client) GetLogEndpointHealthStatus(ctx context.Context, cluster string, namespace string, timeout int, verbose int) (healthStatus *HealthStatus, err error) {
	healthStatus = new(HealthStatus)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetLogEndpointHealthStatus").
		Query("Cluster", cluster).
		Query("Namespace", namespace).
		Query("Timeout", timeout).
		Query("Verbose", verbose).
		TOPRPCData(healthStatus).
		Do(ctx)
	return
}

// GetLogFiles description:
// Retrieve collected log files for the given pod
func (c *Client) GetLogFiles(ctx context.Context, cid string, namespace string, pod string) (fileGroups *FileGroups, err error) {
	fileGroups = new(FileGroups)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetLogFiles").
		Query("Cid", cid).
		Query("Namespace", namespace).
		Query("Pod", pod).
		TOPRPCData(fileGroups).
		Do(ctx)
	return
}

// GetLogStream description:
// Stream logs by filters
func (c *Client) GetLogStream(ctx context.Context, options *StreamOptions, k8s bool) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetLogStream").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("Targets", options.Targets).
		Query("Pod", options.PodName).
		Query("PodWildcard", options.PodWildcard).
		Query("PodPrefix", options.PodPrefix).
		Query("Container", options.ContainerName).
		Query("Keyword", options.Keyword).
		Query("FilePath", options.FilePath).
		Query("Verbose", options.Verbose).
		Query("SinceTime", options.SinceTime).
		Query("SinceSecond", options.SinceSecond).
		Query("TailLines", options.TailLines).
		Query("Timestamps", options.Timestamps).
		Query("K8s", k8s).
		Do(ctx)
	return
}

// GetMetric does not have any description.
func (c *Client) GetMetric(ctx context.Context, metricID string) (metric *Metric, err error) {
	metric = new(Metric)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetMetric").
		Query("MetricID", metricID).
		TOPRPCData(metric).
		Do(ctx)
	return
}

// GetNodeCPUStats description:
// Get statistics on Nodes CPU usage
func (c *Client) GetNodeCPUStats(ctx context.Context, filters *ClustersFilters, listOptions *ListOptions) (nodeCPUStates []NodeCPUStats, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetNodeCPUStats").
		Query("Clusters", filters.Clusters).
		Query("Start", listOptions.Start).
		Query("Limit", listOptions.Limit).
		Query("SortBy", listOptions.SortBy).
		Query("ReverseOrder", listOptions.ReverseOrder).
		Data(&nodeCPUStates).
		Do(ctx)
	return
}

// GetSystemLogStream description:
// Stream system component logs by filters
func (c *Client) GetSystemLogStream(ctx context.Context, options *SystemComponentOptions) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetSystemLogStream").
		Query("Cid", options.ClusterID).
		Query("Component", options.Component).
		Query("Node", options.NodeName).
		Query("Verbose", options.Verbose).
		Query("SinceTime", options.SinceTime).
		Query("SinceSecond", options.SinceSecond).
		Query("TailLines", options.TailLines).
		Query("Timestamps", options.Timestamps).
		Do(ctx)
	return
}

// ListDashboardGraphs does not have any description.
func (c *Client) ListDashboardGraphs(ctx context.Context, graphIDs *GraphIDs) (graphList *GraphList, err error) {
	graphList = new(GraphList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListDashboardGraphs").
		Body("application/json", graphIDs).
		TOPRPCData(graphList).
		Do(ctx)
	return
}

// ListDashboards does not have any description.
func (c *Client) ListDashboards(ctx context.Context, filter *DashboardFilter) (dashboardList *DashboardList, err error) {
	dashboardList = new(DashboardList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListDashboards").
		Header("X-Tenant", filter.Tenant).
		Header("X-User", filter.User).
		Query("Preset", filter.Preset).
		Query("Labels", filter.Labels).
		TOPRPCData(dashboardList).
		Do(ctx)
	return
}

// ListEvents description:
// Get events by filters
func (c *Client) ListEvents(ctx context.Context, options *EventListOptions) (eventListResult *EventListResult, err error) {
	eventListResult = new(EventListResult)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListEvents").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("Uid", options.UID).
		Query("FromTime", options.FromTime).
		Query("ToTime", options.ToTime).
		Query("Limit", options.Limit).
		TOPRPCData(eventListResult).
		Do(ctx)
	return
}

// ListGraphs does not have any description.
func (c *Client) ListGraphs(ctx context.Context, filter *GraphFilter) (graphList *GraphList, err error) {
	graphList = new(GraphList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListGraphs").
		Header("X-Tenant", filter.Tenant).
		Query("Preset", filter.Preset).
		TOPRPCData(graphList).
		Do(ctx)
	return
}

// ListLabelValues description:
// Get label values from label name with given expression
func (c *Client) ListLabelValues(ctx context.Context, labelName string, expr string) (labelValueResult *LabelValueResult, err error) {
	labelValueResult = new(LabelValueResult)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListLabelValues").
		Query("LabelName", labelName).
		Query("Expr", expr).
		TOPRPCData(labelValueResult).
		Do(ctx)
	return
}

// ListLogEndpoint description:
// List log endpoints
func (c *Client) ListLogEndpoint(ctx context.Context, keyword string) (listResult *ListResult, err error) {
	listResult = new(ListResult)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListLogEndpoint").
		Query("Keyword", keyword).
		TOPRPCData(listResult).
		Do(ctx)
	return
}

// ListMetrics does not have any description.
func (c *Client) ListMetrics(ctx context.Context) (metricList *MetricList, err error) {
	metricList = new(MetricList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListMetrics").
		TOPRPCData(metricList).
		Do(ctx)
	return
}

// ListUnitSeries does not have any description.
func (c *Client) ListUnitSeries(ctx context.Context, includeUnit *string) (unitSeriesList *UnitSeriesList, err error) {
	unitSeriesList = new(UnitSeriesList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListUnitSeries").
		Query("IncludeUnit", includeUnit).
		TOPRPCData(unitSeriesList).
		Do(ctx)
	return
}

// QueryRange description:
// Prometheus query range
func (c *Client) QueryRange(ctx context.Context, promQuery *PromQuery) (stats *Stats, err error) {
	stats = new(Stats)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=QueryRange").
		Body("application/json", promQuery).
		TOPRPCData(stats).
		Do(ctx)
	return
}

// SearchEvents description:
// Get events by filters
func (c *Client) SearchEvents(ctx context.Context, options *EventSearchOptions) (eventSearchResult *EventSearchResult, err error) {
	eventSearchResult = new(EventSearchResult)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=SearchEvents").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("InvolvedObjectName", options.InvolvedObjectName).
		Query("InvolvedObjectKind", options.InvolvedObjectKind).
		Query("Type", options.Type).
		Query("Reason", options.Reason).
		Query("Keyword", options.Keyword).
		Query("FromTime", options.FromTime).
		Query("ToTime", options.ToTime).
		Query("Start", options.Start).
		Query("Limit", options.Limit).
		Query("SortBy", options.SortBy).
		TOPRPCData(eventSearchResult).
		Do(ctx)
	return
}

// SearchLog description:
// Get logs by filters
func (c *Client) SearchLog(ctx context.Context, options *LogSearchOptions) (resourceLogItems ResourceLogItems, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=SearchLog").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("Targets", options.Targets).
		Query("Pod", options.PodName).
		Query("PodWildcard", options.PodWildcard).
		Query("PodPrefix", options.PodPrefix).
		Query("Container", options.ContainerName).
		Query("Keyword", options.Keyword).
		Query("FilePath", options.FilePath).
		Query("FromTime", options.FromTime).
		Query("ToTime", options.ToTime).
		Query("StartID", options.StartID).
		Query("Limit", options.Limit).
		TOPRPCData(&resourceLogItems).
		Do(ctx)
	return
}

// SearchLogContexts description:
// Get logs context by filters
func (c *Client) SearchLogContexts(ctx context.Context, options *ContextSearchOptions) (resourceLogItems ResourceLogItems, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=SearchLogContexts").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("Targets", options.Targets).
		Query("Pod", options.PodName).
		Query("PodWildcard", options.PodWildcard).
		Query("PodPrefix", options.PodPrefix).
		Query("Container", options.ContainerName).
		Query("Keyword", options.Keyword).
		Query("FilePath", options.FilePath).
		Query("Id", options.ID).
		Query("BeforeLines", options.BeforeLines).
		Query("AfterLines", options.AfterLines).
		TOPRPCData(&resourceLogItems).
		Do(ctx)
	return
}

// SearchLogContextsRaw description:
// Download logs context by filters
func (c *Client) SearchLogContextsRaw(ctx context.Context, options *ContextSearchOptions) (strings map[string]string, ioReader io.Reader, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=SearchLogContextsRaw").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("Targets", options.Targets).
		Query("Pod", options.PodName).
		Query("PodWildcard", options.PodWildcard).
		Query("PodPrefix", options.PodPrefix).
		Query("Container", options.ContainerName).
		Query("Keyword", options.Keyword).
		Query("FilePath", options.FilePath).
		Query("Id", options.ID).
		Query("BeforeLines", options.BeforeLines).
		Query("AfterLines", options.AfterLines).
		Meta(&strings).
		Data(&ioReader).
		Do(ctx)
	return
}

// SearchLogRaw description:
// Download logs by filters
func (c *Client) SearchLogRaw(ctx context.Context, options *LogSearchOptions) (strings map[string]string, ioReader io.Reader, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=SearchLogRaw").
		Query("Cid", options.ClusterID).
		Query("Namespace", options.Namespace).
		Query("LoadName", options.LoadName).
		Query("LoadType", options.LoadType).
		Query("Targets", options.Targets).
		Query("Pod", options.PodName).
		Query("PodWildcard", options.PodWildcard).
		Query("PodPrefix", options.PodPrefix).
		Query("Container", options.ContainerName).
		Query("Keyword", options.Keyword).
		Query("FilePath", options.FilePath).
		Query("FromTime", options.FromTime).
		Query("ToTime", options.ToTime).
		Query("StartID", options.StartID).
		Query("Limit", options.Limit).
		Meta(&strings).
		Data(&ioReader).
		Do(ctx)
	return
}

// UpdateDashboard does not have any description.
func (c *Client) UpdateDashboard(ctx context.Context, dashboardId string, dashboard *Dashboard) (dashboard1 *Dashboard, err error) {
	dashboard1 = new(Dashboard)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateDashboard").
		Query("DashboardId", dashboardId).
		Body("application/json", dashboard).
		TOPRPCData(dashboard1).
		Do(ctx)
	return
}

// UpdateGraph does not have any description.
func (c *Client) UpdateGraph(ctx context.Context, graph string, graph1 *Graph) (graph2 *Graph, err error) {
	graph2 = new(Graph)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateGraph").
		Query("Graph", graph).
		Body("application/json", graph1).
		TOPRPCData(graph2).
		Do(ctx)
	return
}

// UpdateLogEndpoint description:
// Update a log endpoint
func (c *Client) UpdateLogEndpoint(ctx context.Context, logEndpoint *LogEndpoint) (bindingResults *BindingResults, err error) {
	bindingResults = new(BindingResults)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateLogEndpoint").
		Body("application/json", logEndpoint).
		TOPRPCData(bindingResults).
		Do(ctx)
	return
}
