package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// CreateConfigMap does not have any description.
	CreateConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// CreateDeployment does not have any description.
	CreateDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption, deployment *Deployment) (deployment1 *Deployment, err error)
	// CreateHelmApp does not have any description.
	CreateHelmApp(ctx context.Context, createOption CreateOption, helmApp *HelmApp) (helmApp1 *HelmApp, err error)
	// CreateSecret does not have any description.
	CreateSecret(ctx context.Context, secretGetOption SecretGetOption, secret *Secret) (secret1 *Secret, err error)
	// CreateService does not have any description.
	CreateService(ctx context.Context, serviceGetOption ServiceGetOption, service *Service) (service1 *Service, err error)
	// CreateStatefulSet does not have any description.
	CreateStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
	// DeleteConfigMap does not have any description.
	DeleteConfigMap(ctx context.Context, configMapDeleteOption ConfigMapDeleteOption) (err error)
	// DeleteDeployment does not have any description.
	DeleteDeployment(ctx context.Context, deploymentDeleteOption DeploymentDeleteOption) (err error)
	// DeleteHelmApp does not have any description.
	DeleteHelmApp(ctx context.Context, deleteOption DeleteOption) (err error)
	// DeleteSecret does not have any description.
	DeleteSecret(ctx context.Context, secretDeleteOption SecretDeleteOption) (err error)
	// DeleteService does not have any description.
	DeleteService(ctx context.Context, serviceDeleteOption ServiceDeleteOption) (err error)
	// DeleteStatefulSet does not have any description.
	DeleteStatefulSet(ctx context.Context, statefulSetDeleteOption StatefulSetDeleteOption) (err error)
	// GetConfigMap does not have any description.
	GetConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption) (configMap *ConfigMap, err error)
	// GetDeployment does not have any description.
	GetDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption) (deployment *Deployment, err error)
	// GetHelmApp does not have any description.
	GetHelmApp(ctx context.Context, getOption GetOption) (helmApp *HelmApp, err error)
	// GetHelmAppRevision does not have any description.
	GetHelmAppRevision(ctx context.Context, getHelmRevisionOption GetHelmRevisionOption) (helmRevision *HelmRevision, err error)
	// GetOverview does not have any description.
	GetOverview(ctx context.Context) (overview *Overview, err error)
	// GetSecret does not have any description.
	GetSecret(ctx context.Context, secretGetOption SecretGetOption) (secret *Secret, err error)
	// GetService does not have any description.
	GetService(ctx context.Context, serviceGetOption ServiceGetOption) (service *Service, err error)
	// GetStatefulSet does not have any description.
	GetStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption) (statefulSet *StatefulSet, err error)
	// ListConfigMaps does not have any description.
	ListConfigMaps(ctx context.Context, configMapListOption ConfigMapListOption) (configMapList *ConfigMapList, err error)
	// ListDeployments does not have any description.
	ListDeployments(ctx context.Context, deploymentListOption DeploymentListOption) (deploymentList *DeploymentList, err error)
	// ListHelmApp does not have any description.
	ListHelmApp(ctx context.Context, listOption ListOption) (helmAppList *HelmAppList, err error)
	// ListHelmAppRevisions does not have any description.
	ListHelmAppRevisions(ctx context.Context, listOption ListOption) (revisionList *RevisionList, err error)
	// ListSecrets does not have any description.
	ListSecrets(ctx context.Context, secretListOption SecretListOption) (secretList *SecretList, err error)
	// ListServices does not have any description.
	ListServices(ctx context.Context, serviceListOption ServiceListOption) (serviceList *ServiceList, err error)
	// ListStatefulSets does not have any description.
	ListStatefulSets(ctx context.Context, statefulSetListOption StatefulSetListOption) (statefulSetList *StatefulSetList, err error)
	// RestartDeployment does not have any description.
	RestartDeployment(ctx context.Context, deploymentRestartOption DeploymentRestartOption) (err error)
	// RestartStatefulSet does not have any description.
	RestartStatefulSet(ctx context.Context, statefulSetRestartOption StatefulSetRestartOption) (err error)
	// RollbackHelmAppToRevision does not have any description.
	RollbackHelmAppToRevision(ctx context.Context, rollbackHelmAppToRevisionOption RollbackHelmAppToRevisionOption) (helmApp *HelmApp, err error)
	// UpdateConfigMap does not have any description.
	UpdateConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// UpdateDeployment does not have any description.
	UpdateDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption, deployment *Deployment) (deployment1 *Deployment, err error)
	// UpdateHelmApp does not have any description.
	UpdateHelmApp(ctx context.Context, updateOption UpdateOption, helmApp *HelmApp) (helmApp1 *HelmApp, err error)
	// UpdateSecret does not have any description.
	UpdateSecret(ctx context.Context, secretGetOption SecretGetOption, secret *Secret) (secret1 *Secret, err error)
	// UpdateService does not have any description.
	UpdateService(ctx context.Context, serviceGetOption ServiceGetOption, service *Service) (service1 *Service, err error)
	// UpdateStatefulSet does not have any description.
	UpdateStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
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
func (c *Client) CreateConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateConfigMap").
		Query("ClusterName", configMapGetOption.ClusterName).
		Query("Namespace", configMapGetOption.Namespace).
		Query("Name", configMapGetOption.Name).
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// CreateDeployment does not have any description.
func (c *Client) CreateDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption, deployment *Deployment) (deployment1 *Deployment, err error) {
	deployment1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateDeployment").
		Query("ClusterName", deploymentGetOption.ClusterName).
		Query("Namespace", deploymentGetOption.Namespace).
		Query("Name", deploymentGetOption.Name).
		Body("application/json", deployment).
		TOPRPCData(deployment1).
		Do(ctx)
	return
}

// CreateHelmApp does not have any description.
func (c *Client) CreateHelmApp(ctx context.Context, createOption CreateOption, helmApp *HelmApp) (helmApp1 *HelmApp, err error) {
	helmApp1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateHelmApp").
		Query("ClusterName", createOption.ClusterName).
		Query("Namespace", createOption.Namespace).
		Query("Name", createOption.Name).
		Body("application/json", helmApp).
		TOPRPCData(helmApp1).
		Do(ctx)
	return
}

// CreateSecret does not have any description.
func (c *Client) CreateSecret(ctx context.Context, secretGetOption SecretGetOption, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateSecret").
		Query("ClusterName", secretGetOption.ClusterName).
		Query("Namespace", secretGetOption.Namespace).
		Query("Name", secretGetOption.Name).
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// CreateService does not have any description.
func (c *Client) CreateService(ctx context.Context, serviceGetOption ServiceGetOption, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateService").
		Query("ClusterName", serviceGetOption.ClusterName).
		Query("Namespace", serviceGetOption.Namespace).
		Query("Name", serviceGetOption.Name).
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// CreateStatefulSet does not have any description.
func (c *Client) CreateStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStatefulSet").
		Query("ClusterName", statefulSetGetOption.ClusterName).
		Query("Namespace", statefulSetGetOption.Namespace).
		Query("Name", statefulSetGetOption.Name).
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}

// DeleteConfigMap does not have any description.
func (c *Client) DeleteConfigMap(ctx context.Context, configMapDeleteOption ConfigMapDeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteConfigMap").
		Query("ClusterName", configMapDeleteOption.ClusterName).
		Query("Namespace", configMapDeleteOption.Namespace).
		Query("Name", configMapDeleteOption.Name).
		Do(ctx)
	return
}

// DeleteDeployment does not have any description.
func (c *Client) DeleteDeployment(ctx context.Context, deploymentDeleteOption DeploymentDeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteDeployment").
		Query("ClusterName", deploymentDeleteOption.ClusterName).
		Query("Namespace", deploymentDeleteOption.Namespace).
		Query("Name", deploymentDeleteOption.Name).
		Do(ctx)
	return
}

// DeleteHelmApp does not have any description.
func (c *Client) DeleteHelmApp(ctx context.Context, deleteOption DeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteHelmApp").
		Query("ClusterName", deleteOption.ClusterName).
		Query("Namespace", deleteOption.Namespace).
		Query("Name", deleteOption.Name).
		Do(ctx)
	return
}

// DeleteSecret does not have any description.
func (c *Client) DeleteSecret(ctx context.Context, secretDeleteOption SecretDeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteSecret").
		Query("ClusterName", secretDeleteOption.ClusterName).
		Query("Namespace", secretDeleteOption.Namespace).
		Query("Name", secretDeleteOption.Name).
		Do(ctx)
	return
}

// DeleteService does not have any description.
func (c *Client) DeleteService(ctx context.Context, serviceDeleteOption ServiceDeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteService").
		Query("ClusterName", serviceDeleteOption.ClusterName).
		Query("Namespace", serviceDeleteOption.Namespace).
		Query("Name", serviceDeleteOption.Name).
		Do(ctx)
	return
}

// DeleteStatefulSet does not have any description.
func (c *Client) DeleteStatefulSet(ctx context.Context, statefulSetDeleteOption StatefulSetDeleteOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteStatefulSet").
		Query("ClusterName", statefulSetDeleteOption.ClusterName).
		Query("Namespace", statefulSetDeleteOption.Namespace).
		Query("Name", statefulSetDeleteOption.Name).
		Do(ctx)
	return
}

// GetConfigMap does not have any description.
func (c *Client) GetConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption) (configMap *ConfigMap, err error) {
	configMap = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetConfigMap").
		Query("ClusterName", configMapGetOption.ClusterName).
		Query("Namespace", configMapGetOption.Namespace).
		Query("Name", configMapGetOption.Name).
		TOPRPCData(configMap).
		Do(ctx)
	return
}

// GetDeployment does not have any description.
func (c *Client) GetDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption) (deployment *Deployment, err error) {
	deployment = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetDeployment").
		Query("ClusterName", deploymentGetOption.ClusterName).
		Query("Namespace", deploymentGetOption.Namespace).
		Query("Name", deploymentGetOption.Name).
		TOPRPCData(deployment).
		Do(ctx)
	return
}

// GetHelmApp does not have any description.
func (c *Client) GetHelmApp(ctx context.Context, getOption GetOption) (helmApp *HelmApp, err error) {
	helmApp = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetHelmApp").
		Query("ClusterName", getOption.ClusterName).
		Query("Namespace", getOption.Namespace).
		Query("Name", getOption.Name).
		TOPRPCData(helmApp).
		Do(ctx)
	return
}

// GetHelmAppRevision does not have any description.
func (c *Client) GetHelmAppRevision(ctx context.Context, getHelmRevisionOption GetHelmRevisionOption) (helmRevision *HelmRevision, err error) {
	helmRevision = new(HelmRevision)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetHelmAppRevision").
		Query("ClusterName", getHelmRevisionOption.ClusterName).
		Query("Namespace", getHelmRevisionOption.Namespace).
		Query("Name", getHelmRevisionOption.Name).
		Query("Revision", getHelmRevisionOption.Revision).
		TOPRPCData(helmRevision).
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
func (c *Client) GetSecret(ctx context.Context, secretGetOption SecretGetOption) (secret *Secret, err error) {
	secret = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetSecret").
		Query("ClusterName", secretGetOption.ClusterName).
		Query("Namespace", secretGetOption.Namespace).
		Query("Name", secretGetOption.Name).
		TOPRPCData(secret).
		Do(ctx)
	return
}

// GetService does not have any description.
func (c *Client) GetService(ctx context.Context, serviceGetOption ServiceGetOption) (service *Service, err error) {
	service = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetService").
		Query("ClusterName", serviceGetOption.ClusterName).
		Query("Namespace", serviceGetOption.Namespace).
		Query("Name", serviceGetOption.Name).
		TOPRPCData(service).
		Do(ctx)
	return
}

// GetStatefulSet does not have any description.
func (c *Client) GetStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption) (statefulSet *StatefulSet, err error) {
	statefulSet = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStatefulSet").
		Query("ClusterName", statefulSetGetOption.ClusterName).
		Query("Namespace", statefulSetGetOption.Namespace).
		Query("Name", statefulSetGetOption.Name).
		TOPRPCData(statefulSet).
		Do(ctx)
	return
}

// ListConfigMaps does not have any description.
func (c *Client) ListConfigMaps(ctx context.Context, configMapListOption ConfigMapListOption) (configMapList *ConfigMapList, err error) {
	configMapList = new(ConfigMapList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListConfigMaps").
		Query("Start", configMapListOption.Start).
		Query("Limit", configMapListOption.Limit).
		Query("Query", configMapListOption.Query).
		Query("ClusterName", configMapListOption.ClusterName).
		Query("Namespace", configMapListOption.Namespace).
		Query("Name", configMapListOption.Name).
		TOPRPCData(configMapList).
		Do(ctx)
	return
}

// ListDeployments does not have any description.
func (c *Client) ListDeployments(ctx context.Context, deploymentListOption DeploymentListOption) (deploymentList *DeploymentList, err error) {
	deploymentList = new(DeploymentList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListDeployments").
		Query("Start", deploymentListOption.Start).
		Query("Limit", deploymentListOption.Limit).
		Query("Query", deploymentListOption.Query).
		Query("ClusterName", deploymentListOption.ClusterName).
		Query("Namespace", deploymentListOption.Namespace).
		Query("Name", deploymentListOption.Name).
		TOPRPCData(deploymentList).
		Do(ctx)
	return
}

// ListHelmApp does not have any description.
func (c *Client) ListHelmApp(ctx context.Context, listOption ListOption) (helmAppList *HelmAppList, err error) {
	helmAppList = new(HelmAppList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmApp").
		Query("Start", listOption.Start).
		Query("Limit", listOption.Limit).
		Query("Query", listOption.Query).
		Query("ClusterName", listOption.ClusterName).
		Query("Namespace", listOption.Namespace).
		Query("Name", listOption.Name).
		TOPRPCData(helmAppList).
		Do(ctx)
	return
}

// ListHelmAppRevisions does not have any description.
func (c *Client) ListHelmAppRevisions(ctx context.Context, listOption ListOption) (revisionList *RevisionList, err error) {
	revisionList = new(RevisionList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListHelmAppRevisions").
		Query("Start", listOption.Start).
		Query("Limit", listOption.Limit).
		Query("Query", listOption.Query).
		Query("ClusterName", listOption.ClusterName).
		Query("Namespace", listOption.Namespace).
		Query("Name", listOption.Name).
		TOPRPCData(revisionList).
		Do(ctx)
	return
}

// ListSecrets does not have any description.
func (c *Client) ListSecrets(ctx context.Context, secretListOption SecretListOption) (secretList *SecretList, err error) {
	secretList = new(SecretList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListSecrets").
		Query("Start", secretListOption.Start).
		Query("Limit", secretListOption.Limit).
		Query("Query", secretListOption.Query).
		Query("ClusterName", secretListOption.ClusterName).
		Query("Namespace", secretListOption.Namespace).
		Query("Name", secretListOption.Name).
		TOPRPCData(secretList).
		Do(ctx)
	return
}

// ListServices does not have any description.
func (c *Client) ListServices(ctx context.Context, serviceListOption ServiceListOption) (serviceList *ServiceList, err error) {
	serviceList = new(ServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListServices").
		Query("Start", serviceListOption.Start).
		Query("Limit", serviceListOption.Limit).
		Query("Query", serviceListOption.Query).
		Query("ClusterName", serviceListOption.ClusterName).
		Query("Namespace", serviceListOption.Namespace).
		Query("Name", serviceListOption.Name).
		TOPRPCData(serviceList).
		Do(ctx)
	return
}

// ListStatefulSets does not have any description.
func (c *Client) ListStatefulSets(ctx context.Context, statefulSetListOption StatefulSetListOption) (statefulSetList *StatefulSetList, err error) {
	statefulSetList = new(StatefulSetList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStatefulSets").
		Query("Start", statefulSetListOption.Start).
		Query("Limit", statefulSetListOption.Limit).
		Query("Query", statefulSetListOption.Query).
		Query("ClusterName", statefulSetListOption.ClusterName).
		Query("Namespace", statefulSetListOption.Namespace).
		Query("Name", statefulSetListOption.Name).
		TOPRPCData(statefulSetList).
		Do(ctx)
	return
}

// RestartDeployment does not have any description.
func (c *Client) RestartDeployment(ctx context.Context, deploymentRestartOption DeploymentRestartOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartDeployment").
		Query("ClusterName", deploymentRestartOption.ClusterName).
		Query("Namespace", deploymentRestartOption.Namespace).
		Query("Name", deploymentRestartOption.Name).
		Do(ctx)
	return
}

// RestartStatefulSet does not have any description.
func (c *Client) RestartStatefulSet(ctx context.Context, statefulSetRestartOption StatefulSetRestartOption) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RestartStatefulSet").
		Query("ClusterName", statefulSetRestartOption.ClusterName).
		Query("Namespace", statefulSetRestartOption.Namespace).
		Query("Name", statefulSetRestartOption.Name).
		Do(ctx)
	return
}

// RollbackHelmAppToRevision does not have any description.
func (c *Client) RollbackHelmAppToRevision(ctx context.Context, rollbackHelmAppToRevisionOption RollbackHelmAppToRevisionOption) (helmApp *HelmApp, err error) {
	helmApp = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RollbackHelmAppToRevision").
		Query("ClusterName", rollbackHelmAppToRevisionOption.ClusterName).
		Query("Namespace", rollbackHelmAppToRevisionOption.Namespace).
		Query("Name", rollbackHelmAppToRevisionOption.Name).
		Query("Revision", rollbackHelmAppToRevisionOption.Revision).
		TOPRPCData(helmApp).
		Do(ctx)
	return
}

// UpdateConfigMap does not have any description.
func (c *Client) UpdateConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateConfigMap").
		Query("ClusterName", configMapGetOption.ClusterName).
		Query("Namespace", configMapGetOption.Namespace).
		Query("Name", configMapGetOption.Name).
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// UpdateDeployment does not have any description.
func (c *Client) UpdateDeployment(ctx context.Context, deploymentGetOption DeploymentGetOption, deployment *Deployment) (deployment1 *Deployment, err error) {
	deployment1 = new(Deployment)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateDeployment").
		Query("ClusterName", deploymentGetOption.ClusterName).
		Query("Namespace", deploymentGetOption.Namespace).
		Query("Name", deploymentGetOption.Name).
		Body("application/json", deployment).
		TOPRPCData(deployment1).
		Do(ctx)
	return
}

// UpdateHelmApp does not have any description.
func (c *Client) UpdateHelmApp(ctx context.Context, updateOption UpdateOption, helmApp *HelmApp) (helmApp1 *HelmApp, err error) {
	helmApp1 = new(HelmApp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateHelmApp").
		Query("ClusterName", updateOption.ClusterName).
		Query("Namespace", updateOption.Namespace).
		Query("Name", updateOption.Name).
		Body("application/json", helmApp).
		TOPRPCData(helmApp1).
		Do(ctx)
	return
}

// UpdateSecret does not have any description.
func (c *Client) UpdateSecret(ctx context.Context, secretGetOption SecretGetOption, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateSecret").
		Query("ClusterName", secretGetOption.ClusterName).
		Query("Namespace", secretGetOption.Namespace).
		Query("Name", secretGetOption.Name).
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// UpdateService does not have any description.
func (c *Client) UpdateService(ctx context.Context, serviceGetOption ServiceGetOption, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateService").
		Query("ClusterName", serviceGetOption.ClusterName).
		Query("Namespace", serviceGetOption.Namespace).
		Query("Name", serviceGetOption.Name).
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// UpdateStatefulSet does not have any description.
func (c *Client) UpdateStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStatefulSet").
		Query("ClusterName", statefulSetGetOption.ClusterName).
		Query("Namespace", statefulSetGetOption.Namespace).
		Query("Name", statefulSetGetOption.Name).
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}
