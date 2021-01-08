package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// CreateConfigMap does not have any description.
	CreateConfigMap(ctx context.Context, cluster Cluster, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// CreateDeployment does not have any description.
	CreateDeployment(ctx context.Context, cluster Cluster, deployment *Deployment) (deployment1 *Deployment, err error)
	// CreateHelmApp does not have any description.
	CreateHelmApp(ctx context.Context, cluster Cluster, helmApp *HelmApp) (helmApp1 *HelmApp, err error)
	// CreateSecret does not have any description.
	CreateSecret(ctx context.Context, cluster Cluster, secret *Secret) (secret1 *Secret, err error)
	// CreateService does not have any description.
	CreateService(ctx context.Context, cluster Cluster, service *Service) (service1 *Service, err error)
	// CreateStatefulSet does not have any description.
	CreateStatefulSet(ctx context.Context, cluster Cluster, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
	// CreateWithYAML does not have any description.
	CreateWithYAML(ctx context.Context, cluster Cluster, network string, yAML *YAML) (yAML1 *YAML, err error)
	// DeleteConfigMap does not have any description.
	DeleteConfigMap(ctx context.Context, cluster Cluster) (err error)
	// DeleteDeployment does not have any description.
	DeleteDeployment(ctx context.Context, cluster Cluster) (err error)
	// DeleteHelmApp does not have any description.
	DeleteHelmApp(ctx context.Context, cluster Cluster) (err error)
	// DeletePod does not have any description.
	DeletePod(ctx context.Context, cluster Cluster) (err error)
	// DeleteSecret does not have any description.
	DeleteSecret(ctx context.Context, cluster Cluster) (err error)
	// DeleteService does not have any description.
	DeleteService(ctx context.Context, cluster Cluster) (err error)
	// DeleteStatefulSet does not have any description.
	DeleteStatefulSet(ctx context.Context, cluster Cluster) (err error)
	// GetConfigMap does not have any description.
	GetConfigMap(ctx context.Context, cluster Cluster) (configMap *ConfigMap, err error)
	// GetDeployment does not have any description.
	GetDeployment(ctx context.Context, cluster Cluster) (deployment *Deployment, err error)
	// GetDeploymentPodExecSession does not have any description.
	GetDeploymentPodExecSession(ctx context.Context, cluster Cluster, podExecOption PodExecOption) (terminalSession *TerminalSession, err error)
	// GetHelmApp does not have any description.
	GetHelmApp(ctx context.Context, cluster Cluster) (helmApp *HelmApp, err error)
	// GetOverview does not have any description.
	GetOverview(ctx context.Context) (overview *Overview, err error)
	// GetSecret does not have any description.
	GetSecret(ctx context.Context, cluster Cluster) (secret *Secret, err error)
	// GetService does not have any description.
	GetService(ctx context.Context, cluster Cluster) (service *Service, err error)
	// GetStatefulSet does not have any description.
	GetStatefulSet(ctx context.Context, cluster Cluster) (statefulSet *StatefulSet, err error)
	// GetStatefulSetPodExecSession does not have any description.
	GetStatefulSetPodExecSession(ctx context.Context, cluster Cluster, podExecOption PodExecOption) (terminalSession *TerminalSession, err error)
	// ListConfigMaps does not have any description.
	ListConfigMaps(ctx context.Context, cluster Cluster, pagination Pagination) (configMapList *ConfigMapList, err error)
	// ListDeployments does not have any description.
	ListDeployments(ctx context.Context, cluster Cluster, pagination Pagination) (deploymentList *DeploymentList, err error)
	// ListHelmApp does not have any description.
	ListHelmApp(ctx context.Context, cluster Cluster, pagination Pagination) (helmAppList *HelmAppList, err error)
	// ListHelmAppRevisions does not have any description.
	ListHelmAppRevisions(ctx context.Context, cluster Cluster, pagination Pagination) (helmAppRevisionList *HelmAppRevisionList, err error)
	// ListPodsForDeployment does not have any description.
	ListPodsForDeployment(ctx context.Context, cluster Cluster, pagination Pagination) (podList *PodList, err error)
	// ListPodsForWorkload does not have any description.
	ListPodsForWorkload(ctx context.Context, cluster Cluster, kind string) (podList *PodList, err error)
	// ListSecrets does not have any description.
	ListSecrets(ctx context.Context, cluster Cluster, pagination Pagination) (secretList *SecretList, err error)
	// ListServices does not have any description.
	ListServices(ctx context.Context, cluster Cluster, pagination Pagination) (serviceList *ServiceList, err error)
	// ListStatefulSets does not have any description.
	ListStatefulSets(ctx context.Context, cluster Cluster, pagination Pagination) (statefulSetList *StatefulSetList, err error)
	// PodExec does not have any description.
	PodExec(ctx context.Context, sessionOption SessionOption) (err error)
	// RestartDeployment does not have any description.
	RestartDeployment(ctx context.Context, cluster Cluster) (err error)
	// RestartStatefulSet does not have any description.
	RestartStatefulSet(ctx context.Context, cluster Cluster) (err error)
	// RollbackHelmAppToRevision does not have any description.
	RollbackHelmAppToRevision(ctx context.Context, cluster Cluster, revision int) (err error)
	// UpdateConfigMap does not have any description.
	UpdateConfigMap(ctx context.Context, cluster Cluster, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// UpdateDeployment does not have any description.
	UpdateDeployment(ctx context.Context, cluster Cluster, deployment *Deployment) (deployment1 *Deployment, err error)
	// UpdateHelmApp does not have any description.
	UpdateHelmApp(ctx context.Context, cluster Cluster, helmApp *HelmApp) (helmApp1 *HelmApp, err error)
	// UpdateReferencesForConfigMap does not have any description.
	UpdateReferencesForConfigMap(ctx context.Context, cluster Cluster, configMapReferences []ConfigMapReference) (err error)
	// UpdateReferencesForSecret does not have any description.
	UpdateReferencesForSecret(ctx context.Context, cluster Cluster, secretReferences []SecretReference) (err error)
	// UpdateSecret does not have any description.
	UpdateSecret(ctx context.Context, cluster Cluster, secret *Secret) (secret1 *Secret, err error)
	// UpdateService does not have any description.
	UpdateService(ctx context.Context, cluster Cluster, service *Service) (service1 *Service, err error)
	// UpdateStatefulSet does not have any description.
	UpdateStatefulSet(ctx context.Context, cluster Cluster, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
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
func (c *Client) CreateConfigMap(ctx context.Context, cluster Cluster, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateConfigMap").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// CreateDeployment does not have any description.
func (c *Client) CreateDeployment(ctx context.Context, cluster Cluster, deployment *Deployment) (deployment1 *Deployment, err error) {
	deployment1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", deployment).
		TOPRPCData(deployment1).
		Do(ctx)
	return
}

// CreateHelmApp does not have any description.
func (c *Client) CreateHelmApp(ctx context.Context, cluster Cluster, helmApp *HelmApp) (helmApp1 *HelmApp, err error) {
	helmApp1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateHelmApp").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", helmApp).
		TOPRPCData(helmApp1).
		Do(ctx)
	return
}

// CreateSecret does not have any description.
func (c *Client) CreateSecret(ctx context.Context, cluster Cluster, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateSecret").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// CreateService does not have any description.
func (c *Client) CreateService(ctx context.Context, cluster Cluster, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateService").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// CreateStatefulSet does not have any description.
func (c *Client) CreateStatefulSet(ctx context.Context, cluster Cluster, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStatefulSet").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}

// CreateWithYAML does not have any description.
func (c *Client) CreateWithYAML(ctx context.Context, cluster Cluster, network string, yAML *YAML) (yAML1 *YAML, err error) {
	yAML1 = new(YAML)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateWithYAML").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Network", network).
		Body("application/json", yAML).
		TOPRPCData(yAML1).
		Do(ctx)
	return
}

// DeleteConfigMap does not have any description.
func (c *Client) DeleteConfigMap(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteConfigMap").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeleteDeployment does not have any description.
func (c *Client) DeleteDeployment(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeleteHelmApp does not have any description.
func (c *Client) DeleteHelmApp(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteHelmApp").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeletePod does not have any description.
func (c *Client) DeletePod(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePod").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeleteSecret does not have any description.
func (c *Client) DeleteSecret(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteSecret").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeleteService does not have any description.
func (c *Client) DeleteService(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteService").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// DeleteStatefulSet does not have any description.
func (c *Client) DeleteStatefulSet(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteStatefulSet").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// GetConfigMap does not have any description.
func (c *Client) GetConfigMap(ctx context.Context, cluster Cluster) (configMap *ConfigMap, err error) {
	configMap = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetConfigMap").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(configMap).
		Do(ctx)
	return
}

// GetDeployment does not have any description.
func (c *Client) GetDeployment(ctx context.Context, cluster Cluster) (deployment *Deployment, err error) {
	deployment = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(deployment).
		Do(ctx)
	return
}

// GetDeploymentPodExecSession does not have any description.
func (c *Client) GetDeploymentPodExecSession(ctx context.Context, cluster Cluster, podExecOption PodExecOption) (terminalSession *TerminalSession, err error) {
	terminalSession = new(TerminalSession)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDeploymentPodExecSession").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Pod", podExecOption.Pod).
		Query("Container", podExecOption.Container).
		Query("Shell", podExecOption.Shell).
		TOPRPCData(terminalSession).
		Do(ctx)
	return
}

// GetHelmApp does not have any description.
func (c *Client) GetHelmApp(ctx context.Context, cluster Cluster) (helmApp *HelmApp, err error) {
	helmApp = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetHelmApp").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(helmApp).
		Do(ctx)
	return
}

// GetOverview does not have any description.
func (c *Client) GetOverview(ctx context.Context) (overview *Overview, err error) {
	overview = new(Overview)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetOverview").
		TOPRPCData(overview).
		Do(ctx)
	return
}

// GetSecret does not have any description.
func (c *Client) GetSecret(ctx context.Context, cluster Cluster) (secret *Secret, err error) {
	secret = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetSecret").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(secret).
		Do(ctx)
	return
}

// GetService does not have any description.
func (c *Client) GetService(ctx context.Context, cluster Cluster) (service *Service, err error) {
	service = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetService").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(service).
		Do(ctx)
	return
}

// GetStatefulSet does not have any description.
func (c *Client) GetStatefulSet(ctx context.Context, cluster Cluster) (statefulSet *StatefulSet, err error) {
	statefulSet = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStatefulSet").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		TOPRPCData(statefulSet).
		Do(ctx)
	return
}

// GetStatefulSetPodExecSession does not have any description.
func (c *Client) GetStatefulSetPodExecSession(ctx context.Context, cluster Cluster, podExecOption PodExecOption) (terminalSession *TerminalSession, err error) {
	terminalSession = new(TerminalSession)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStatefulSetPodExecSession").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Pod", podExecOption.Pod).
		Query("Container", podExecOption.Container).
		Query("Shell", podExecOption.Shell).
		TOPRPCData(terminalSession).
		Do(ctx)
	return
}

// ListConfigMaps does not have any description.
func (c *Client) ListConfigMaps(ctx context.Context, cluster Cluster, pagination Pagination) (configMapList *ConfigMapList, err error) {
	configMapList = new(ConfigMapList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListConfigMaps").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(configMapList).
		Do(ctx)
	return
}

// ListDeployments does not have any description.
func (c *Client) ListDeployments(ctx context.Context, cluster Cluster, pagination Pagination) (deploymentList *DeploymentList, err error) {
	deploymentList = new(DeploymentList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListDeployments").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(deploymentList).
		Do(ctx)
	return
}

// ListHelmApp does not have any description.
func (c *Client) ListHelmApp(ctx context.Context, cluster Cluster, pagination Pagination) (helmAppList *HelmAppList, err error) {
	helmAppList = new(HelmAppList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmApp").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(helmAppList).
		Do(ctx)
	return
}

// ListHelmAppRevisions does not have any description.
func (c *Client) ListHelmAppRevisions(ctx context.Context, cluster Cluster, pagination Pagination) (helmAppRevisionList *HelmAppRevisionList, err error) {
	helmAppRevisionList = new(HelmAppRevisionList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmAppRevisions").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(helmAppRevisionList).
		Do(ctx)
	return
}

// ListPodsForDeployment does not have any description.
func (c *Client) ListPodsForDeployment(ctx context.Context, cluster Cluster, pagination Pagination) (podList *PodList, err error) {
	podList = new(PodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPodsForDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(podList).
		Do(ctx)
	return
}

// ListPodsForWorkload does not have any description.
func (c *Client) ListPodsForWorkload(ctx context.Context, cluster Cluster, kind string) (podList *PodList, err error) {
	podList = new(PodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPodsForWorkload").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Kind", kind).
		TOPRPCData(podList).
		Do(ctx)
	return
}

// ListSecrets does not have any description.
func (c *Client) ListSecrets(ctx context.Context, cluster Cluster, pagination Pagination) (secretList *SecretList, err error) {
	secretList = new(SecretList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListSecrets").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(secretList).
		Do(ctx)
	return
}

// ListServices does not have any description.
func (c *Client) ListServices(ctx context.Context, cluster Cluster, pagination Pagination) (serviceList *ServiceList, err error) {
	serviceList = new(ServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListServices").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(serviceList).
		Do(ctx)
	return
}

// ListStatefulSets does not have any description.
func (c *Client) ListStatefulSets(ctx context.Context, cluster Cluster, pagination Pagination) (statefulSetList *StatefulSetList, err error) {
	statefulSetList = new(StatefulSetList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStatefulSets").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Start", pagination.Start).
		Query("Limit", pagination.Limit).
		Query("Query", pagination.Query).
		TOPRPCData(statefulSetList).
		Do(ctx)
	return
}

// PodExec does not have any description.
func (c *Client) PodExec(ctx context.Context, sessionOption SessionOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=PodExec").
		Query("Server", sessionOption.Server).
		Query("Session", sessionOption.Session).
		Do(ctx)
	return
}

// RestartDeployment does not have any description.
func (c *Client) RestartDeployment(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// RestartStatefulSet does not have any description.
func (c *Client) RestartStatefulSet(ctx context.Context, cluster Cluster) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartStatefulSet").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Do(ctx)
	return
}

// RollbackHelmAppToRevision does not have any description.
func (c *Client) RollbackHelmAppToRevision(ctx context.Context, cluster Cluster, revision int) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RollbackHelmAppToRevision").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Query("Revision", revision).
		Do(ctx)
	return
}

// UpdateConfigMap does not have any description.
func (c *Client) UpdateConfigMap(ctx context.Context, cluster Cluster, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateConfigMap").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// UpdateDeployment does not have any description.
func (c *Client) UpdateDeployment(ctx context.Context, cluster Cluster, deployment *Deployment) (deployment1 *Deployment, err error) {
	deployment1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateDeployment").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", deployment).
		TOPRPCData(deployment1).
		Do(ctx)
	return
}

// UpdateHelmApp does not have any description.
func (c *Client) UpdateHelmApp(ctx context.Context, cluster Cluster, helmApp *HelmApp) (helmApp1 *HelmApp, err error) {
	helmApp1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateHelmApp").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", helmApp).
		TOPRPCData(helmApp1).
		Do(ctx)
	return
}

// UpdateReferencesForConfigMap does not have any description.
func (c *Client) UpdateReferencesForConfigMap(ctx context.Context, cluster Cluster, configMapReferences []ConfigMapReference) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateReferencesForConfigMap").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", configMapReferences).
		Do(ctx)
	return
}

// UpdateReferencesForSecret does not have any description.
func (c *Client) UpdateReferencesForSecret(ctx context.Context, cluster Cluster, secretReferences []SecretReference) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateReferencesForSecret").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", secretReferences).
		Do(ctx)
	return
}

// UpdateSecret does not have any description.
func (c *Client) UpdateSecret(ctx context.Context, cluster Cluster, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateSecret").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// UpdateService does not have any description.
func (c *Client) UpdateService(ctx context.Context, cluster Cluster, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateService").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// UpdateStatefulSet does not have any description.
func (c *Client) UpdateStatefulSet(ctx context.Context, cluster Cluster, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStatefulSet").
		Query("ClusterName", cluster.ClusterName).
		Query("Namespace", cluster.Namespace).
		Query("Name", cluster.Name).
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}
