package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// CreateConfigMap does not have any description.
	CreateConfigMap(ctx context.Context, cluster_ Cluster, configMap_ *ConfigMap) (configMap_1 *ConfigMap, err error)
	// CreateDeployment does not have any description.
	CreateDeployment(ctx context.Context, cluster_ Cluster, deployment_ *Deployment) (deployment_1 *Deployment, err error)
	// CreateHelmApp does not have any description.
	CreateHelmApp(ctx context.Context, cluster_ Cluster, helmApp_ *HelmApp) (helmApp_1 *HelmApp, err error)
	// CreateSecret does not have any description.
	CreateSecret(ctx context.Context, cluster_ Cluster, secret_ *Secret) (secret_1 *Secret, err error)
	// CreateService does not have any description.
	CreateService(ctx context.Context, cluster_ Cluster, service_ *Service) (service_1 *Service, err error)
	// CreateStatefulSet does not have any description.
	CreateStatefulSet(ctx context.Context, cluster_ Cluster, statefulSet_ *StatefulSet) (statefulSet_1 *StatefulSet, err error)
	// CreateWithYAML does not have any description.
	CreateWithYAML(ctx context.Context, cluster_ Cluster, network_ string, yAML_ *YAML) (yAML_1 *YAML, err error)
	// DeleteConfigMap does not have any description.
	DeleteConfigMap(ctx context.Context, cluster_ Cluster) (err error)
	// DeleteDeployment does not have any description.
	DeleteDeployment(ctx context.Context, cluster_ Cluster) (err error)
	// DeleteHelmApp does not have any description.
	DeleteHelmApp(ctx context.Context, cluster_ Cluster) (err error)
	// DeletePod does not have any description.
	DeletePod(ctx context.Context, cluster_ Cluster) (err error)
	// DeleteSecret does not have any description.
	DeleteSecret(ctx context.Context, cluster_ Cluster) (err error)
	// DeleteService does not have any description.
	DeleteService(ctx context.Context, cluster_ Cluster) (err error)
	// DeleteStatefulSet does not have any description.
	DeleteStatefulSet(ctx context.Context, cluster_ Cluster) (err error)
	// GetConfigMap does not have any description.
	GetConfigMap(ctx context.Context, cluster_ Cluster) (configMap_ *ConfigMap, err error)
	// GetDeployment does not have any description.
	GetDeployment(ctx context.Context, cluster_ Cluster) (deployment_ *Deployment, err error)
	// GetDeploymentPodExecSession does not have any description.
	GetDeploymentPodExecSession(ctx context.Context, cluster_ Cluster, podExecOption_ PodExecOption) (terminalSession_ *TerminalSession, err error)
	// GetHelmApp does not have any description.
	GetHelmApp(ctx context.Context, cluster_ Cluster) (helmApp_ *HelmApp, err error)
	// GetOverview does not have any description.
	GetOverview(ctx context.Context) (overview_ *Overview, err error)
	// GetSecret does not have any description.
	GetSecret(ctx context.Context, cluster_ Cluster) (secret_ *Secret, err error)
	// GetService does not have any description.
	GetService(ctx context.Context, cluster_ Cluster) (service_ *Service, err error)
	// GetStatefulSet does not have any description.
	GetStatefulSet(ctx context.Context, cluster_ Cluster) (statefulSet_ *StatefulSet, err error)
	// GetStatefulSetPodExecSession does not have any description.
	GetStatefulSetPodExecSession(ctx context.Context, cluster_ Cluster, podExecOption_ PodExecOption) (terminalSession_ *TerminalSession, err error)
	// ListConfigMaps does not have any description.
	ListConfigMaps(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (configMapList_ *ConfigMapList, err error)
	// ListDeployments does not have any description.
	ListDeployments(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (deploymentList_ *DeploymentList, err error)
	// ListHelmApp does not have any description.
	ListHelmApp(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (helmAppList_ *HelmAppList, err error)
	// ListHelmAppRevisions does not have any description.
	ListHelmAppRevisions(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (helmAppRevisionList_ *HelmAppRevisionList, err error)
	// ListPodsForDeployment does not have any description.
	ListPodsForDeployment(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (podList_ *PodList, err error)
	// ListPodsForWorkload does not have any description.
	ListPodsForWorkload(ctx context.Context, cluster_ Cluster, kind_ string) (podList_ *PodList, err error)
	// ListSecrets does not have any description.
	ListSecrets(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (secretList_ *SecretList, err error)
	// ListServices does not have any description.
	ListServices(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error)
	// ListServicesForDeployment does not have any description.
	ListServicesForDeployment(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error)
	// ListServicesForStatefulSet does not have any description.
	ListServicesForStatefulSet(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error)
	// ListStatefulSets does not have any description.
	ListStatefulSets(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (statefulSetList_ *StatefulSetList, err error)
	// PodExec does not have any description.
	PodExec(ctx context.Context, sessionOption_ SessionOption) (err error)
	// RestartDeployment does not have any description.
	RestartDeployment(ctx context.Context, cluster_ Cluster) (err error)
	// RestartStatefulSet does not have any description.
	RestartStatefulSet(ctx context.Context, cluster_ Cluster) (err error)
	// RollbackHelmAppToRevision does not have any description.
	RollbackHelmAppToRevision(ctx context.Context, cluster_ Cluster, revision_ int) (err error)
	// UpdateConfigMap does not have any description.
	UpdateConfigMap(ctx context.Context, cluster_ Cluster, configMap_ *ConfigMap) (configMap_1 *ConfigMap, err error)
	// UpdateDeployment does not have any description.
	UpdateDeployment(ctx context.Context, cluster_ Cluster, deployment_ *Deployment) (deployment_1 *Deployment, err error)
	// UpdateHelmApp does not have any description.
	UpdateHelmApp(ctx context.Context, cluster_ Cluster, helmApp_ *HelmApp) (helmApp_1 *HelmApp, err error)
	// UpdateReferencesForConfigMap does not have any description.
	UpdateReferencesForConfigMap(ctx context.Context, cluster_ Cluster, configMapReferences_ []ConfigMapReference) (err error)
	// UpdateReferencesForSecret does not have any description.
	UpdateReferencesForSecret(ctx context.Context, cluster_ Cluster, secretReferences_ []SecretReference) (err error)
	// UpdateSecret does not have any description.
	UpdateSecret(ctx context.Context, cluster_ Cluster, secret_ *Secret) (secret_1 *Secret, err error)
	// UpdateService does not have any description.
	UpdateService(ctx context.Context, cluster_ Cluster, service_ *Service) (service_1 *Service, err error)
	// UpdateStatefulSet does not have any description.
	UpdateStatefulSet(ctx context.Context, cluster_ Cluster, statefulSet_ *StatefulSet) (statefulSet_1 *StatefulSet, err error)
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

// CreateConfigMap does not have any description.
func (c *Client) CreateConfigMap(ctx context.Context, cluster_ Cluster, configMap_ *ConfigMap) (configMap_1 *ConfigMap, err error) {
	configMap_1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateConfigMap").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", configMap_).
		TOPRPCData(configMap_1).
		Do(ctx)
	return
}

// CreateDeployment does not have any description.
func (c *Client) CreateDeployment(ctx context.Context, cluster_ Cluster, deployment_ *Deployment) (deployment_1 *Deployment, err error) {
	deployment_1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", deployment_).
		TOPRPCData(deployment_1).
		Do(ctx)
	return
}

// CreateHelmApp does not have any description.
func (c *Client) CreateHelmApp(ctx context.Context, cluster_ Cluster, helmApp_ *HelmApp) (helmApp_1 *HelmApp, err error) {
	helmApp_1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateHelmApp").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", helmApp_).
		TOPRPCData(helmApp_1).
		Do(ctx)
	return
}

// CreateSecret does not have any description.
func (c *Client) CreateSecret(ctx context.Context, cluster_ Cluster, secret_ *Secret) (secret_1 *Secret, err error) {
	secret_1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateSecret").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", secret_).
		TOPRPCData(secret_1).
		Do(ctx)
	return
}

// CreateService does not have any description.
func (c *Client) CreateService(ctx context.Context, cluster_ Cluster, service_ *Service) (service_1 *Service, err error) {
	service_1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateService").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", service_).
		TOPRPCData(service_1).
		Do(ctx)
	return
}

// CreateStatefulSet does not have any description.
func (c *Client) CreateStatefulSet(ctx context.Context, cluster_ Cluster, statefulSet_ *StatefulSet) (statefulSet_1 *StatefulSet, err error) {
	statefulSet_1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", statefulSet_).
		TOPRPCData(statefulSet_1).
		Do(ctx)
	return
}

// CreateWithYAML does not have any description.
func (c *Client) CreateWithYAML(ctx context.Context, cluster_ Cluster, network_ string, yAML_ *YAML) (yAML_1 *YAML, err error) {
	yAML_1 = new(YAML)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateWithYAML").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Network", network_).
		Body("application/json", yAML_).
		TOPRPCData(yAML_1).
		Do(ctx)
	return
}

// DeleteConfigMap does not have any description.
func (c *Client) DeleteConfigMap(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteConfigMap").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeleteDeployment does not have any description.
func (c *Client) DeleteDeployment(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeleteHelmApp does not have any description.
func (c *Client) DeleteHelmApp(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteHelmApp").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeletePod does not have any description.
func (c *Client) DeletePod(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePod").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeleteSecret does not have any description.
func (c *Client) DeleteSecret(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteSecret").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeleteService does not have any description.
func (c *Client) DeleteService(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteService").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// DeleteStatefulSet does not have any description.
func (c *Client) DeleteStatefulSet(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// GetConfigMap does not have any description.
func (c *Client) GetConfigMap(ctx context.Context, cluster_ Cluster) (configMap_ *ConfigMap, err error) {
	configMap_ = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetConfigMap").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(configMap_).
		Do(ctx)
	return
}

// GetDeployment does not have any description.
func (c *Client) GetDeployment(ctx context.Context, cluster_ Cluster) (deployment_ *Deployment, err error) {
	deployment_ = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(deployment_).
		Do(ctx)
	return
}

// GetDeploymentPodExecSession does not have any description.
func (c *Client) GetDeploymentPodExecSession(ctx context.Context, cluster_ Cluster, podExecOption_ PodExecOption) (terminalSession_ *TerminalSession, err error) {
	terminalSession_ = new(TerminalSession)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDeploymentPodExecSession").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Pod", podExecOption_.Pod).
		Query("Container", podExecOption_.Container).
		Query("Shell", podExecOption_.Shell).
		TOPRPCData(terminalSession_).
		Do(ctx)
	return
}

// GetHelmApp does not have any description.
func (c *Client) GetHelmApp(ctx context.Context, cluster_ Cluster) (helmApp_ *HelmApp, err error) {
	helmApp_ = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetHelmApp").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(helmApp_).
		Do(ctx)
	return
}

// GetOverview does not have any description.
func (c *Client) GetOverview(ctx context.Context) (overview_ *Overview, err error) {
	overview_ = new(Overview)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetOverview").
		TOPRPCData(overview_).
		Do(ctx)
	return
}

// GetSecret does not have any description.
func (c *Client) GetSecret(ctx context.Context, cluster_ Cluster) (secret_ *Secret, err error) {
	secret_ = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetSecret").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(secret_).
		Do(ctx)
	return
}

// GetService does not have any description.
func (c *Client) GetService(ctx context.Context, cluster_ Cluster) (service_ *Service, err error) {
	service_ = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetService").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(service_).
		Do(ctx)
	return
}

// GetStatefulSet does not have any description.
func (c *Client) GetStatefulSet(ctx context.Context, cluster_ Cluster) (statefulSet_ *StatefulSet, err error) {
	statefulSet_ = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		TOPRPCData(statefulSet_).
		Do(ctx)
	return
}

// GetStatefulSetPodExecSession does not have any description.
func (c *Client) GetStatefulSetPodExecSession(ctx context.Context, cluster_ Cluster, podExecOption_ PodExecOption) (terminalSession_ *TerminalSession, err error) {
	terminalSession_ = new(TerminalSession)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStatefulSetPodExecSession").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Pod", podExecOption_.Pod).
		Query("Container", podExecOption_.Container).
		Query("Shell", podExecOption_.Shell).
		TOPRPCData(terminalSession_).
		Do(ctx)
	return
}

// ListConfigMaps does not have any description.
func (c *Client) ListConfigMaps(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (configMapList_ *ConfigMapList, err error) {
	configMapList_ = new(ConfigMapList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListConfigMaps").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(configMapList_).
		Do(ctx)
	return
}

// ListDeployments does not have any description.
func (c *Client) ListDeployments(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (deploymentList_ *DeploymentList, err error) {
	deploymentList_ = new(DeploymentList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListDeployments").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(deploymentList_).
		Do(ctx)
	return
}

// ListHelmApp does not have any description.
func (c *Client) ListHelmApp(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (helmAppList_ *HelmAppList, err error) {
	helmAppList_ = new(HelmAppList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmApp").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(helmAppList_).
		Do(ctx)
	return
}

// ListHelmAppRevisions does not have any description.
func (c *Client) ListHelmAppRevisions(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (helmAppRevisionList_ *HelmAppRevisionList, err error) {
	helmAppRevisionList_ = new(HelmAppRevisionList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmAppRevisions").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(helmAppRevisionList_).
		Do(ctx)
	return
}

// ListPodsForDeployment does not have any description.
func (c *Client) ListPodsForDeployment(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (podList_ *PodList, err error) {
	podList_ = new(PodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPodsForDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(podList_).
		Do(ctx)
	return
}

// ListPodsForWorkload does not have any description.
func (c *Client) ListPodsForWorkload(ctx context.Context, cluster_ Cluster, kind_ string) (podList_ *PodList, err error) {
	podList_ = new(PodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPodsForWorkload").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Kind", kind_).
		TOPRPCData(podList_).
		Do(ctx)
	return
}

// ListSecrets does not have any description.
func (c *Client) ListSecrets(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (secretList_ *SecretList, err error) {
	secretList_ = new(SecretList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListSecrets").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(secretList_).
		Do(ctx)
	return
}

// ListServices does not have any description.
func (c *Client) ListServices(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error) {
	serviceList_ = new(ServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListServices").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(serviceList_).
		Do(ctx)
	return
}

// ListServicesForDeployment does not have any description.
func (c *Client) ListServicesForDeployment(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error) {
	serviceList_ = new(ServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListServicesForDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(serviceList_).
		Do(ctx)
	return
}

// ListServicesForStatefulSet does not have any description.
func (c *Client) ListServicesForStatefulSet(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (serviceList_ *ServiceList, err error) {
	serviceList_ = new(ServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListServicesForStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(serviceList_).
		Do(ctx)
	return
}

// ListStatefulSets does not have any description.
func (c *Client) ListStatefulSets(ctx context.Context, cluster_ Cluster, pagination_ Pagination) (statefulSetList_ *StatefulSetList, err error) {
	statefulSetList_ = new(StatefulSetList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStatefulSets").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Start", pagination_.Start).
		Query("Limit", pagination_.Limit).
		Query("Query", pagination_.Query).
		TOPRPCData(statefulSetList_).
		Do(ctx)
	return
}

// PodExec does not have any description.
func (c *Client) PodExec(ctx context.Context, sessionOption_ SessionOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=PodExec").
		Query("Server", sessionOption_.Server).
		Query("Session", sessionOption_.Session).
		Do(ctx)
	return
}

// RestartDeployment does not have any description.
func (c *Client) RestartDeployment(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// RestartStatefulSet does not have any description.
func (c *Client) RestartStatefulSet(ctx context.Context, cluster_ Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Do(ctx)
	return
}

// RollbackHelmAppToRevision does not have any description.
func (c *Client) RollbackHelmAppToRevision(ctx context.Context, cluster_ Cluster, revision_ int) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RollbackHelmAppToRevision").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Query("Revision", revision_).
		Do(ctx)
	return
}

// UpdateConfigMap does not have any description.
func (c *Client) UpdateConfigMap(ctx context.Context, cluster_ Cluster, configMap_ *ConfigMap) (configMap_1 *ConfigMap, err error) {
	configMap_1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateConfigMap").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", configMap_).
		TOPRPCData(configMap_1).
		Do(ctx)
	return
}

// UpdateDeployment does not have any description.
func (c *Client) UpdateDeployment(ctx context.Context, cluster_ Cluster, deployment_ *Deployment) (deployment_1 *Deployment, err error) {
	deployment_1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateDeployment").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", deployment_).
		TOPRPCData(deployment_1).
		Do(ctx)
	return
}

// UpdateHelmApp does not have any description.
func (c *Client) UpdateHelmApp(ctx context.Context, cluster_ Cluster, helmApp_ *HelmApp) (helmApp_1 *HelmApp, err error) {
	helmApp_1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateHelmApp").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", helmApp_).
		TOPRPCData(helmApp_1).
		Do(ctx)
	return
}

// UpdateReferencesForConfigMap does not have any description.
func (c *Client) UpdateReferencesForConfigMap(ctx context.Context, cluster_ Cluster, configMapReferences_ []ConfigMapReference) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateReferencesForConfigMap").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", configMapReferences_).
		Do(ctx)
	return
}

// UpdateReferencesForSecret does not have any description.
func (c *Client) UpdateReferencesForSecret(ctx context.Context, cluster_ Cluster, secretReferences_ []SecretReference) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateReferencesForSecret").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", secretReferences_).
		Do(ctx)
	return
}

// UpdateSecret does not have any description.
func (c *Client) UpdateSecret(ctx context.Context, cluster_ Cluster, secret_ *Secret) (secret_1 *Secret, err error) {
	secret_1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateSecret").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", secret_).
		TOPRPCData(secret_1).
		Do(ctx)
	return
}

// UpdateService does not have any description.
func (c *Client) UpdateService(ctx context.Context, cluster_ Cluster, service_ *Service) (service_1 *Service, err error) {
	service_1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateService").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", service_).
		TOPRPCData(service_1).
		Do(ctx)
	return
}

// UpdateStatefulSet does not have any description.
func (c *Client) UpdateStatefulSet(ctx context.Context, cluster_ Cluster, statefulSet_ *StatefulSet) (statefulSet_1 *StatefulSet, err error) {
	statefulSet_1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStatefulSet").
		Query("ClusterName", cluster_.ClusterName).
		Query("Namespace", cluster_.Namespace).
		Query("Name", cluster_.Name).
		Body("application/json", statefulSet_).
		TOPRPCData(statefulSet_1).
		Do(ctx)
	return
}
