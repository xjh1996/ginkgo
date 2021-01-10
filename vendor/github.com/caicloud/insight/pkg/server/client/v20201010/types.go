package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	v1alpha1 "github.com/caicloud/containeros/pkg/apis/logging/v1alpha1"
	model "github.com/prometheus/common/model"
	time "time"
)

// AggregationOp describes the aggregation operator itself of promQL.
type AggregationOp string

// AlertLevel describes the alert level.
type AlertLevel struct {
	Level     int64   `json:"Level"`
	Operator  string  `json:"Operator"`
	Threshold float64 `json:"Threshold"`
	Duration  string  `json:"Duration"`
}

// Auth defines auth info for elasticsearch
type Auth struct {
	Username    []uint8
	Password    []uint8
	APIKey      []uint8 `json:"ApiKey"`
	Certificate []uint8
	KeyPem      []uint8
}

// Binding defines a bind of LogEndpoint And clusters
type Binding struct {
	LogEndpointName string
	Namespace       string
	Default         *bool
	Clusters        []string
}

// BindingResult defines a cluster log endpoint binding result
type BindingResult struct {
	Cluster      string
	Failed       bool
	ErrorMessage string
}

// BindingResults define a bunch of status of cluster log endpoint binding
type BindingResults struct {
	FailCount            int
	BindingStatus        []BindingResult
	DefaultBindingStatus BindingResult
}

// Clause describes the clause of aggregation operator of promQL.
type Clause string

// ClusterResourceStats ...
type ClusterResourceStats struct {
	Cluster     string `json:"Cluster"`
	CPUCores    string `json:"CPUCores"`
	MemoryBytes string `json:"MemoryBytes"`
}

// ClustersFilters ...
type ClustersFilters struct {
	Clusters []string `source:"Query,Clusters"`
}

// CommonOptions ...
type CommonOptions struct {
	Verbose        bool  `source:"query,Verbose"`
	SinceTime      int64 `source:"query,SinceTime"`
	SinceSecond    int   `source:"query,SinceSecond"`
	TailLines      int   `source:"query,TailLines"`
	Timestamps     bool  `source:"query,Timestamps"`
	ShowTargetName bool
}

// ContainerFilters ...
type ContainerFilters struct {
	Owner     string `source:"Query,Owner"`
	OwnerKind string `source:"Query,OwnerKind"`
	Namespace string `source:"Query,Namespace"`
	Cluster   string `source:"Query,Cluster"`
}

// ContainerResourceStats ...
type ContainerResourceStats struct {
	Container        string   `json:"Container"`
	Pod              string   `json:"Pod"`
	Owner            string   `json:"Owner"`
	OwnerKind        string   `json:"OwnerKind"`
	Namespace        string   `json:"Namespace"`
	Cluster          string   `json:"Cluster"`
	StartTime        *int64   `json:"StartTime,omitempty"`
	CPUUsageCores    *float64 `json:"CPUUsageCores,omitempty"`
	MemoryUsageBytes *float64 `json:"MemoryUsageBytes,omitempty"`
}

// ContainerStatsOptions ...
type ContainerStatsOptions struct {
	StartTime int64   `source:"Query,StartTime"`
	EndTime   int64   `source:"Query,EndTime"`
	Quantile  float64 `source:"Query,Quantile"`
}

// ContextSearchOptions ...
type ContextSearchOptions struct {
	GeneralOptions
	ID           string `source:"query,Id"`
	BeforeLines  int    `source:"query,BeforeLines" validate:"lte=100"`
	AfterLines   int    `source:"query,AfterLines" validate:"lte=100"`
	DownloadType string
}

// Dashboard is a monitoring dashboard.
type Dashboard struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	GraphSpecs    []DashboardGraphSpec `json:"Graphs" bson:"Graphs"`
}

// DashboardFilter describes the filter when listing the dashboard.
type DashboardFilter struct {
	Tenant string `source:"Header,X-Tenant"`
	User   string `source:"Header,X-User"`
	Preset bool   `source:"Query,Preset"`
	Labels string `source:"Query,Labels"`
}

// DashboardGraphSpec describes a graph and its layout in a dashboard
type DashboardGraphSpec struct {
	Layout GraphLayout `json:"Layout"`
	Graph  Graph       `json:"Graph"`
}

// DashboardList is the list of dashboards
type DashboardList struct {
	v1.ListMeta `json:",inline"`
	Items       []Dashboard `json:"Items"`
}

// ESHealthStatus defines elasticsearch status
type ESHealthStatus struct {
	Host       string
	StatusCode int
	ErrorMsg   string `json:"Error,omitempty"`
}

// ESPingResults defines a bunch of ESHealthStatus
type ESPingResults []ESHealthStatus

// Elasticsearch defines a elasticsearch config
type Elasticsearch struct {
	// Auth is Elasticsearch Auth for connecting, if needed.
	Auth *Auth `json:"Auth,omitempty"`
	// Hosts is a list of Elasticsearch nodes to connect to.
	Hosts []string
	// Protocol is a name of the protocol Elasticsearch is reachable on.
	Protocol v1alpha1.ElasticsearchEndpointProtocol
	// Path is an HTTP path prefix that is prepended to the HTTP API calls.
	Path string
}

// EventListOptions ...
type EventListOptions struct {
	// nolint
	ClusterID string `source:"query,Cid"`
	Namespace string `source:"query,Namespace"`
	UID       string `source:"query,Uid"`
	FromTime  int64  `source:"query,FromTime" validate:"required"`
	ToTime    int64  `source:"query,ToTime" validate:"required"`
	Limit     int64  `source:"query,Limit" validate:"lte=100"`
}

// EventListResult ...
type EventListResult struct {
	// nolint
	Matrix  []EventMatrix `json:"Matrix"`
	Legends []string      `json:"Legends"`
}

// EventMatrix ...
type EventMatrix struct {
	// nolint
	Metric EventMetric `json:"Metric"`
	// Values format as follows
	// [
	// [string, int],
	// [string, int]
	// ]
	Values [][]interface{} `json:"Values"`
}

// EventMetric ...
type EventMetric struct {
	// nolint
	UID            string         `json:"Uid"`
	InvolvedObject InvolvedObject `json:"InvolvedObject"`
	Source         Source         `json:"Source"`
	Message        string         `json:"Message"`
	Reason         string         `json:"Reason"`
	Type           string         `json:"Type"`
}

// EventSearchOptions ...
// nolint
type EventSearchOptions struct {
	ClusterID          string `source:"query,Cid"`
	Namespace          string `source:"query,Namespace"`
	LoadName           string `source:"query,LoadName"`
	LoadType           string `source:"query,LoadType"`
	InvolvedObjectName string `source:"query,InvolvedObjectName"`
	InvolvedObjectKind string `source:"query,InvolvedObjectKind"`
	Type               string `source:"query,Type"`
	Reason             string `source:"query,Reason"`
	Keyword            string `source:"query,Keyword"`
	FromTime           int64  `source:"query,FromTime" validate:"required"`
	ToTime             int64  `source:"query,ToTime" validate:"required"`
	Start              int64  `source:"query,Start" validate:"gte=0"`
	Limit              int64  `source:"query,Limit" validate:"lte=100"`
	SortBy             string `source:"query,SortBy,default=-lastOccurrenceTimestamp"`
	UID                string `json:"-"`
	NewestOnly         bool   `json:"-"`
}

// EventSearchResult ...
type EventSearchResult struct {
	// nolint
	Total int64  `json:"Total"`
	Items []Item `json:"Items"`
}

// FileGroup log file group
type FileGroup struct {
	Container string   `json:"Container"`
	Files     []string `json:"Files"`
}

// FileGroups log file groups
type FileGroups struct {
	Groups []FileGroup `json:"Groups"`
}

// Filebeat defines a filebeat config
type Filebeat struct {
}

// GeneralOptions ...
type GeneralOptions struct {
	ClusterID string `source:"query,Cid"`
	Namespace string `source:"query,Namespace"`
	LoadName  string `source:"query,LoadName"`
	// LoadType can be "deployment" or "statefulset"
	LoadType      string `source:"query,LoadType,optional"`
	Targets       string `source:"query,Targets"`
	PodName       string `source:"query,Pod"`
	PodWildcard   string `source:"query,PodWildcard"`
	PodPrefix     string `source:"query,PodPrefix"`
	ContainerName string `source:"query,Container"`
	Keyword       string `source:"query,Keyword"`
	FilePath      string `source:"query,FilePath"`
}

// Graph represents the monitoring graph.
type Graph struct {
	v1.ObjectMeta   `json:",inline" bson:",inline"`
	Style           string        `json:"Style" bson:"Style"`
	Query           PromQuerySpec `json:"Query" bson:"Query"`
	LegendTemplate  string        `json:"LegendTemplate,omitempty" bson:"LegendTemplate,omitempty"`
	AlertingRuleUID string        `json:"AlertingRuleUid,omitempty" bson:"AlertingRuleUid,omitempty"`
}

// GraphFilter is the filter for listing graphs.
type GraphFilter struct {
	IDs    []string
	Tenant string `source:"Header,X-Tenant"`
	Preset bool   `source:"Query,Preset"`
}

// GraphIDs describes a set of graph IDs for filtering.
type GraphIDs struct {
	IDs []string `json:"IDs"`
}

// GraphLayout describes the layout of a graph in Dashboard.
type GraphLayout struct {
	X      int `json:"X" bson:"X"`
	Y      int `json:"Y" bson:"Y"`
	Width  int `json:"Width" bson:"Width"`
	Height int `json:"Height" bson:"Height"`
}

// GraphList is a list of Graph.
type GraphList struct {
	v1.ListMeta `json:",inline"`
	Items       []Graph `json:"Items"`
}

// HealthStatus defines log endpoint health status
type HealthStatus struct {
	TotalCount   int
	HealthyCount int
	LogEndpoint  *LogEndpoint   `json:"LogEndpoint,omitempty"`
	Results      *ESPingResults `json:"Results,omitempty"`
}

// InvolvedObject ...
type InvolvedObject struct {
	Kind      string `json:"Kind" bson:"kind"`
	Namespace string `json:"Namespace" bson:"namespace"`
	Name      string `json:"Name" bson:"name"`
	LoadName  string `json:"LoadName" bson:"loadName"`
	LoadType  string `json:"LoadType" bson:"loadType"`
	FieldPath string `json:"FieldPath" bson:"fieldPath"`
	UID       string `json:"Uid" bson:"uid"`
}

// Item ...
type Item struct {
	ID                       string         `json:"-" bson:"_id"`
	UID                      string         `json:"Uid" bson:"uid"`
	Count                    int            `json:"Count" bson:"count"`
	FirstOccurrenceTimestamp time.Time      `json:"FirstOccurrenceTimestamp" bson:"firstOccurrenceTimestamp"`
	LastOccurrenceTimestamp  time.Time      `json:"LastOccurrenceTimestamp" bson:"lastOccurrenceTimestamp"`
	InvolvedObject           InvolvedObject `json:"InvolvedObject" bson:"involvedObject"`
	Source                   Source         `json:"Source" bson:"source"`
	Message                  string         `json:"Message" bson:"message"`
	Reason                   string         `json:"Reason" bson:"reason"`
	Type                     string         `json:"Type" bson:"type"`
	Newest                   bool           `json:"-" bson:"newest"`
	// new
	Cluster string `json:"-" bson:"cluster"`
}

// LabelValueResult represents the result of list label values
type LabelValueResult struct {
	Label  string   `json:"Label"`
	Values []string `json:"Values"`
}

// ListOptions contains generic list options
type ListOptions struct {
	Start        int64  `source:"Query,Start,default=0" validate:"gte=0"`
	Limit        int64  `source:"Query,Limit,default=100" validate:"gt=0"`
	SortBy       string `source:"Query,SortBy"`
	ReverseOrder bool   `source:"Query,ReverseOrder"`
}

// ListResult defines list log endpoint response
type ListResult struct {
	Total int64         `json:"Total"`
	Items []LogEndpoint `json:"Items"`
}

// LoadBalancerFilters ...
type LoadBalancerFilters struct {
	LoadBalancer string `source:"Query,LoadBalancer"`
	Cluster      string `source:"Query,Cluster"`
}

// LoadBalancerStats ...
type LoadBalancerStats struct {
	LoadBalancer                    string   `json:"LoadBalancer"`
	Cluster                         string   `json:"Cluster"`
	IngressTotalReadBytesPerSecond  *float64 `json:"IngressTotalReadBytesPerSeconds,omitempty"`
	IngressTotalWriteBytesPerSecond *float64 `json:"IngressTotalWriteBytesPerSeconds,omitempty"`
}

// LogEndpoint defines a log endpoint for api
type LogEndpoint struct {
	Name          string
	Namespace     string
	Alias         string
	Description   string
	Default       *bool
	Clusters      []string
	Elasticsearch *Elasticsearch `json:"Elasticsearch,omitempty"`
	Filebeat      *Filebeat      `json:"Filebeat,omitempty"`
}

// LogSearchOptions ...
// nolint
type LogSearchOptions struct {
	GeneralOptions
	FromTime     int64  `source:"query,FromTime" validate:"required"`
	ToTime       int64  `source:"query,ToTime" validate:"required"`
	StartID      string `source:"query,StartID"`
	Limit        int    `source:"query,Limit" validate:"lte=100"`
	DownloadType string
}

// Matrix is a list of time series.
type Matrix []*SampleStream

// Metric describes a monitoring metric.
type Metric struct {
	v1.ObjectMeta `json:",inline" yaml:",inline" bson:",inline"`
	MetricSpec    `json:",inline" yaml:",inline" bson:",inline" `
}

// MetricList is a list of metrics.
type MetricList struct {
	v1.ListMeta `json:",inline"`
	Items       []Metric `json:"Items"`
}

// MetricSpec is the spec of a metric.
type MetricSpec struct {
	QueryTemplate string            `json:"QueryTemplate" yaml:"QueryTemplate" bson:"QueryTemplate"`
	LabelNames    []string          `json:"LabelNames,omitempty" yaml:"LabelNames,omitempty" bson:"LabelNames,omitempty"`
	Tags          []string          `json:"Tags,omitempty" yaml:"Tags,omitempty" bson:"Tags,omitempty"`
	UnitUID       string            `json:"UnitUid,omitempty" yaml:"UnitUid,omitempty" bson:"UnitUid,omitempty"`
	ValueMapping  map[string]string `json:"ValueMapping,omitempty" yaml:"ValueMapping,omitempty" bson:"ValueMapping,omitempty"`
	AlertTemplate AlertLevel        `json:"AlertTemplate,omitempty" yaml:"AlertTemplate,omitempty" bson:"AlertTemplate,omitempty"`
}

// NodeCPUStats ...
type NodeCPUStats struct {
	Node       string   `json:"Node"`
	Cluster    string   `json:"Cluster"`
	TotalCores *int     `json:"TotalCores,omitempty"`
	LoadCores  *float64 `json:"LoadCores,omitempty"`
	UsageCores *float64 `json:"UsageCores,omitempty"`
	Saturation *float64 `json:"Saturation,omitempty"`
}

// PromQuery performs a real prometheus query.
type PromQuery struct {
	Start          int64         `json:"Start"`
	End            int64         `json:"End"`
	Step           string        `json:"Step"`
	LegendTemplate string        `json:"LegendTemplate"`
	QuerySpec      PromQuerySpec `json:"QuerySpec"`
	MetricSpec     *MetricSpec   `json:"MetricSpec,omitempty"`
}

// PromQueryAggregation is the aggregation rule of the query,
// for syntax, see https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
type PromQueryAggregation struct {
	// Aggregator is the aggregation operator, e.g. sum, min, max, topk.
	Aggregator AggregationOp `json:"Aggregator" bson:"Aggregator"`
	// Parameters are the parameters of the given Aggregator
	Parameters []string `json:"Parameters,omitempty" bson:"Parameters,omitempty"`
	// Clause is one of [by, without], by means keep the labels in LabelList, without means the opposite.
	Clause Clause `json:"Clause,omitempty" bson:"Clause,omitempty"`
	// LabelList contains the labels going to be kept or removed in the aggregation.
	LabelList []string `json:"LabelList,omitempty" bson:"LabelList,omitempty"`
}

// PromQueryCondition is the filter condition for filtering label key & values of a metric.
type PromQueryCondition struct {
	// Reverse for reverse filtering, false for keep the values of LabelValues. default values is false.
	Reverse bool `json:"Reverse,omitempty" bson:"Reverse,omitempty"`
	// LabelName denotes the label name as the key for filtering.
	LabelName string `json:"LabelName" bson:"LabelName"`
	// LabelValues corresponds LabelName, as the value for filtering, empty means all values of the label.
	LabelValues []string `json:"LabelValues,omitempty" bson:"LabelValues,omitempty"`
}

// PromQuerySpec is the core concept for Graph and AlertingRule.
// It will eventually generate the PromQL.
type PromQuerySpec struct {
	// MetricUID is for generating the UID of the query object.
	MetricUID string `json:"MetricUid,omitempty" bson:"MetricUid,omitempty"`
	// Conditions are the filter conditions of the query. see PromQueryCondition for details.
	Conditions []PromQueryCondition `json:"Conditions,omitempty" bson:"Conditions,omitempty"`
	// Aggregations are the aggregation operators of the query. see PromQueryAggregation for details.
	Aggregations []PromQueryAggregation `json:"Aggregations,omitempty" bson:"Aggregations,omitempty"`
	// Query is the raw PromQL of the query. It will either be generated by other fields or hand written.
	// depends on the field Custom, when a hand written Query is given, other field wll be ignored.
	Query string `json:"Query,omitempty" bson:"Query,omitempty"`
	// Custom denotes the PromQuerySpec is whether a generated query from other fields or a hand written one.
	Custom bool `json:"Custom,omitempty" bson:"Custom,omitempty"`
}

// ResourceLogItem ...
type ResourceLogItem struct {
	ID  string `json:"Id"`
	Log string `json:"Log"`
	// prefer highlight log if there is
	Time              time.Time     `json:"Time"`
	Namespace         string        `json:"Namespace"`
	LoadName          string        `json:"LoadName"`
	LoadType          string        `json:"LoadType"`
	PodName           string        `json:"Pod"`
	ContainerName     string        `json:"Container"`
	RawLog            string        `json:"-"`
	Sort              []interface{} `json:"-"`
	logWithTimePrefix bool
	showPodName       bool
}

// ResourceLogItems ...
type ResourceLogItems struct {
	Total int64             `json:"Total"`
	Items []ResourceLogItem `json:"Items"`
}

// SampleStream is aka prometheus time series. but with TOP style.
type SampleStream struct {
	Metric string             `json:"Metric"`
	Values []model.SamplePair `json:"Values"`
}

// Source ...
type Source struct {
	Component string `json:"Component" bson:"component"`
	Host      string `json:"Host" bson:"host"`
}

// Stats is the result of prometheus query.
type Stats struct {
	QueryExpr string   `json:"QueryExpr"`
	Matrix    Matrix   `json:"Matrix"`
	Legends   []string `json:"Legends"`
}

// StreamOptions ...
type StreamOptions struct {
	GeneralOptions
	CommonOptions
}

// SystemComponentOptions ...
type SystemComponentOptions struct {
	ClusterID string `source:"query,Cid"`
	Component string `source:"query,Component"`
	NodeName  string `source:"query,Node"`
	CommonOptions
}

// Unit represents the unit of metrics
type Unit struct {
	v1.ObjectMeta `json:",inline" yaml:",inline" bson:",inline"`
	Prefix        string `json:"Prefix" yaml:"Prefix" bson:"Prefix"`
	Suffix        string `json:"Suffix" yaml:"Suffix" bson:"Suffix"`
	ScaleFactor   int64  `json:"ScaleFactor" yaml:"ScaleFactor" bson:"ScaleFactor"`
}

// UnitSeries describes a series of units that are associated.
type UnitSeries struct {
	v1.ObjectMeta `json:",inline" yaml:",inline" bson:",inline"`
	Tags          []string `json:"Tags,omitempty" yaml:"Tags,omitempty" bson:"Tags,omitempty"`
	Type          UnitType `json:"Type" yaml:"Type" bson:"Type"`
	Items         []Unit   `json:"Items" yaml:"Items" bson:"Items"`
}

// UnitSeriesList is the list of UnitSeries.
type UnitSeriesList struct {
	v1.ListMeta `json:",inline" yaml:",inline"`
	Items       []UnitSeries `json:"Items" yaml:",inline"`
}

// UnitType describes the type of unit.
type UnitType string
