package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// CreateApplication does not have any description.
	CreateApplication(ctx context.Context, application *Application) (application1 *Application, err error)
	// CreateConfigMap does not have any description.
	CreateConfigMap(ctx context.Context, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// CreateSecret does not have any description.
	CreateSecret(ctx context.Context, secret *Secret) (secret1 *Secret, err error)
	// CreateService does not have any description.
	CreateService(ctx context.Context, service *Service) (service1 *Service, err error)
	// CreateStatefulSet does not have any description.
	CreateStatefulSet(ctx context.Context, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
	// DeleteApplication does not have any description.
	DeleteApplication(ctx context.Context, application string) (application1 *Application, err error)
	// DeleteConfigMap does not have any description.
	DeleteConfigMap(ctx context.Context, configMapDeleteOption ConfigMapDeleteOption) (err error)
	// DeleteSecret does not have any description.
	DeleteSecret(ctx context.Context, secretDeleteOption SecretDeleteOption) (err error)
	// DeleteService does not have any description.
	DeleteService(ctx context.Context, serviceDeleteOption *ServiceDeleteOption) (err error)
	// DeleteStatefulSet does not have any description.
	DeleteStatefulSet(ctx context.Context, statefulSetDeleteOption StatefulSetDeleteOption) (err error)
	// GetApplication does not have any description.
	GetApplication(ctx context.Context, application string) (application1 *Application, err error)
	// GetApplicationRevision does not have any description.
	GetApplicationRevision(ctx context.Context, application string, revision int) (applicationRevision *ApplicationRevision, err error)
	// GetConfigMap does not have any description.
	GetConfigMap(ctx context.Context, configMapGetOption ConfigMapGetOption) (configMap *ConfigMap, err error)
	// GetSecret does not have any description.
	GetSecret(ctx context.Context, secretGetOption SecretGetOption) (secret *Secret, err error)
	// GetService does not have any description.
	GetService(ctx context.Context, serviceGetOption *ServiceGetOption) (service *Service, err error)
	// GetStatefulSet does not have any description.
	GetStatefulSet(ctx context.Context, statefulSetGetOption StatefulSetGetOption) (statefulSet *StatefulSet, err error)
	// GetWorkload description:
	// Get a workload by id
	GetWorkload(ctx context.Context, workload string) (workload1 *Workload, err error)
	// GetWorkloadsByService does not have any description.
	GetWorkloadsByService(ctx context.Context, getWorkloadOption GetWorkloadOption) (service *Service, err error)
	// ListApplicationRevisions does not have any description.
	ListApplicationRevisions(ctx context.Context, application string) (applicationRevisionList *ApplicationRevisionList, err error)
	// ListApplications does not have any description.
	ListApplications(ctx context.Context, count int) (applicationList *ApplicationList, err error)
	// ListConfigMaps does not have any description.
	ListConfigMaps(ctx context.Context, configMapListOption ConfigMapListOption) (configMapList *ConfigMapList, err error)
	// ListSecrets does not have any description.
	ListSecrets(ctx context.Context, secretListOption SecretListOption) (secretList *SecretList, err error)
	// ListServices does not have any description.
	ListServices(ctx context.Context, serviceListOption *ServiceListOption) (serviceList *ServiceList, err error)
	// ListStatefulSets does not have any description.
	ListStatefulSets(ctx context.Context, statefulSetListOption StatefulSetListOption) (statefulSetList *StatefulSetList, err error)
	// ListWorkloads description:
	// Query a specified number of workloads and returns an array
	ListWorkloads(ctx context.Context, count int) (workloads []Workload, err error)
	// RestartStatefulSet does not have any description.
	RestartStatefulSet(ctx context.Context, statefulSetRestartOption StatefulSetRestartOption) (err error)
	// RollbackApplicationToRevision does not have any description.
	RollbackApplicationToRevision(ctx context.Context, application string, revision int) (application1 *Application, err error)
	// UpdateApplication does not have any description.
	UpdateApplication(ctx context.Context, application string, application1 *Application) (application2 *Application, err error)
	// UpdateConfigMap does not have any description.
	UpdateConfigMap(ctx context.Context, configMap *ConfigMap) (configMap1 *ConfigMap, err error)
	// UpdateSecret does not have any description.
	UpdateSecret(ctx context.Context, secret *Secret) (secret1 *Secret, err error)
	// UpdateService does not have any description.
	UpdateService(ctx context.Context, service *Service) (service1 *Service, err error)
	// UpdateStatefulSet does not have any description.
	UpdateStatefulSet(ctx context.Context, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error)
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

// CreateApplication does not have any description.
func (c *Client) CreateApplication(ctx context.Context, application *Application) (application1 *Application, err error) {
	application1 = new(Application)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateApplication").
		Body("application/json", application).
		TOPRPCData(application1).
		Do(ctx)
	return
}

// CreateConfigMap does not have any description.
func (c *Client) CreateConfigMap(ctx context.Context, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateConfigMap").
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// CreateSecret does not have any description.
func (c *Client) CreateSecret(ctx context.Context, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateSecret").
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// CreateService does not have any description.
func (c *Client) CreateService(ctx context.Context, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateService").
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// CreateStatefulSet does not have any description.
func (c *Client) CreateStatefulSet(ctx context.Context, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStatefulSet").
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}

// DeleteApplication does not have any description.
func (c *Client) DeleteApplication(ctx context.Context, application string) (application1 *Application, err error) {
	application1 = new(Application)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteApplication").
		Query("Application", application).
		TOPRPCData(application1).
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
func (c *Client) DeleteService(ctx context.Context, serviceDeleteOption *ServiceDeleteOption) (err error) {
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

// GetApplication does not have any description.
func (c *Client) GetApplication(ctx context.Context, application string) (application1 *Application, err error) {
	application1 = new(Application)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetApplication").
		Query("Application", application).
		TOPRPCData(application1).
		Do(ctx)
	return
}

// GetApplicationRevision does not have any description.
func (c *Client) GetApplicationRevision(ctx context.Context, application string, revision int) (applicationRevision *ApplicationRevision, err error) {
	applicationRevision = new(ApplicationRevision)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetApplicationRevision").
		Query("Application", application).
		Query("Revision", revision).
		TOPRPCData(applicationRevision).
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

// GetSecret does not have any description.
func (c *Client) GetSecret(ctx context.Context, secretGetOption SecretGetOption) (secret *Secret, err error) {
	secret = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetSecret").
		Query("ClusterName", secretGetOption.ClusterName).
		Query("Namespace", secretGetOption.Namespace).
		Query("Name", secretGetOption.Name).
		Query("Style", secretGetOption.Style).
		TOPRPCData(secret).
		Do(ctx)
	return
}

// GetService does not have any description.
func (c *Client) GetService(ctx context.Context, serviceGetOption *ServiceGetOption) (service *Service, err error) {
	service = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetService").
		Query("ClusterName", serviceGetOption.ClusterName).
		Query("Namespace", serviceGetOption.Namespace).
		Query("Name", serviceGetOption.Name).
		Query("ContentType", serviceGetOption.ContentType).
		Data(service).
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
		Query("Style", statefulSetGetOption.Style).
		TOPRPCData(statefulSet).
		Do(ctx)
	return
}

// GetWorkload description:
// Get a workload by id
func (c *Client) GetWorkload(ctx context.Context, workload string) (workload1 *Workload, err error) {
	workload1 = new(Workload)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkload").
		Query("Workload", workload).
		TOPRPCData(workload1).
		Do(ctx)
	return
}

// GetWorkloadsByService does not have any description.
func (c *Client) GetWorkloadsByService(ctx context.Context, getWorkloadOption GetWorkloadOption) (service *Service, err error) {
	service = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetWorkloadsByService").
		Query("Start", getWorkloadOption.Start).
		Query("Limit", getWorkloadOption.Limit).
		Query("ClusterName", getWorkloadOption.ClusterName).
		Query("Namespace", getWorkloadOption.Namespace).
		Query("Name", getWorkloadOption.Name).
		TOPRPCData(service).
		Do(ctx)
	return
}

// ListApplicationRevisions does not have any description.
func (c *Client) ListApplicationRevisions(ctx context.Context, application string) (applicationRevisionList *ApplicationRevisionList, err error) {
	applicationRevisionList = new(ApplicationRevisionList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListApplicationRevisions").
		Query("Application", application).
		TOPRPCData(applicationRevisionList).
		Do(ctx)
	return
}

// ListApplications does not have any description.
func (c *Client) ListApplications(ctx context.Context, count int) (applicationList *ApplicationList, err error) {
	applicationList = new(ApplicationList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListApplications").
		Query("Count", count).
		TOPRPCData(applicationList).
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
		Query("Style", secretListOption.Style).
		TOPRPCData(secretList).
		Do(ctx)
	return
}

// ListServices does not have any description.
func (c *Client) ListServices(ctx context.Context, serviceListOption *ServiceListOption) (serviceList *ServiceList, err error) {
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
		Query("Style", statefulSetListOption.Style).
		TOPRPCData(statefulSetList).
		Do(ctx)
	return
}

// ListWorkloads description:
// Query a specified number of workloads and returns an array
func (c *Client) ListWorkloads(ctx context.Context, count int) (workloads []Workload, err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListWorkloads").
		Query("Count", count).
		TOPRPCData(&workloads).
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

// RollbackApplicationToRevision does not have any description.
func (c *Client) RollbackApplicationToRevision(ctx context.Context, application string, revision int) (application1 *Application, err error) {
	application1 = new(Application)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RollbackApplicationToRevision").
		Query("Application", application).
		Query("Revision", revision).
		TOPRPCData(application1).
		Do(ctx)
	return
}

// UpdateApplication does not have any description.
func (c *Client) UpdateApplication(ctx context.Context, application string, application1 *Application) (application2 *Application, err error) {
	application2 = new(Application)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateApplication").
		Query("Application", application).
		Body("application/json", application1).
		TOPRPCData(application2).
		Do(ctx)
	return
}

// UpdateConfigMap does not have any description.
func (c *Client) UpdateConfigMap(ctx context.Context, configMap *ConfigMap) (configMap1 *ConfigMap, err error) {
	configMap1 = new(ConfigMap)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateConfigMap").
		Body("application/json", configMap).
		TOPRPCData(configMap1).
		Do(ctx)
	return
}

// UpdateSecret does not have any description.
func (c *Client) UpdateSecret(ctx context.Context, secret *Secret) (secret1 *Secret, err error) {
	secret1 = new(Secret)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateSecret").
		Body("application/json", secret).
		TOPRPCData(secret1).
		Do(ctx)
	return
}

// UpdateService does not have any description.
func (c *Client) UpdateService(ctx context.Context, service *Service) (service1 *Service, err error) {
	service1 = new(Service)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateService").
		Body("application/json", service).
		TOPRPCData(service1).
		Do(ctx)
	return
}

// UpdateStatefulSet does not have any description.
func (c *Client) UpdateStatefulSet(ctx context.Context, statefulSet *StatefulSet) (statefulSet1 *StatefulSet, err error) {
	statefulSet1 = new(StatefulSet)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStatefulSet").
		Body("application/json", statefulSet).
		TOPRPCData(statefulSet1).
		Do(ctx)
	return
}
