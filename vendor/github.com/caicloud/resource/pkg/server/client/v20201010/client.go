package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// BindMachine description:
	// Bind the machine to cluster and returns an object.
	//
	// Old: POST /clusters/{cluster}/machinesbind
	BindMachine(ctx context.Context, bindMachineRequest *BindMachineRequest) (bindMachineResponse *BindMachineResponse, err error)
	// CheckStorageParameters description:
	// Check StorageClass parameters
	CheckStorageParameters(ctx context.Context, storageClusterParameters *StorageClusterParameters) (err error)
	// CreateCluster description:
	// Create the cluster and returns it
	CreateCluster(ctx context.Context, createClusterRequest *CreateClusterRequest) (cluster *Cluster, err error)
	// CreateMachine description:
	// Create the machine and returns an object.
	//
	// Old: POST /machines
	CreateMachine(ctx context.Context, createMachineRequest *CreateMachineRequest) (createMachineResponse *CreateMachineResponse, err error)
	// CreatePersistentVolumeClaim description:
	// Create PersistentVolumeClaim
	CreatePersistentVolumeClaim(ctx context.Context, createPVCRequest *CreatePVCRequest) (pVCObject *PVCObject, err error)
	// CreateStorageClass description:
	// Create StorageClass
	CreateStorageClass(ctx context.Context, createStorageClassRequest *CreateStorageClassRequest) (storageClass *StorageClass, err error)
	// CreateStorageService description:
	// Create StorageService
	CreateStorageService(ctx context.Context, createStorageServiceRequest *CreateStorageServiceRequest) (storageService *StorageService, err error)
	// DeleteCluster description:
	// Delete the cluster and returns it
	DeleteCluster(ctx context.Context, deleteClusterRequest *DeleteClusterRequest) (cluster *Cluster, err error)
	// DeleteMachine description:
	// Delete the machine and returns an object.
	//
	// Old: DELETE /machines/{machine}
	DeleteMachine(ctx context.Context, deleteMachineRequest *DeleteMachineRequest) (machine *Machine, err error)
	// DeletePersistentVolumeClaim description:
	// Delete PersistentVolumeClaim
	DeletePersistentVolumeClaim(ctx context.Context, cluster string, namespace string, name string) (err error)
	// DeleteStorageClass description:
	// Delete StorageClass
	DeleteStorageClass(ctx context.Context, cluster string, name string) (err error)
	// DeleteStorageService description:
	// Delete StorageService
	DeleteStorageService(ctx context.Context, name string) (err error)
	// GetCluster description:
	// Get the cluster and returns it
	GetCluster(ctx context.Context, getClusterRequest *GetClusterRequest) (cluster *Cluster, err error)
	// GetMachine description:
	// Get the machine and returns an object.
	//
	// Old: GET /machines/{machine}
	GetMachine(ctx context.Context, getMachineRequest *GetMachineRequest) (machine *Machine, err error)
	// GetMachineAuth description:
	// Get the machine auth and returns an object.
	//
	// Old: GET /machines/{machine}/auth
	GetMachineAuth(ctx context.Context, getMachineAuthRequest *GetMachineAuthRequest) (auth *Auth, err error)
	// GetNode description:
	// Get the node and returns an object.
	//
	// Old: GET /clusters/{cluster}/nodes/{node}
	GetNode(ctx context.Context, getNodeRequest *GetNodeRequest) (node *Node, err error)
	// GetPersistentVolumeClaim description:
	// Get PersistentVolumeClaim
	GetPersistentVolumeClaim(ctx context.Context, cluster string, namespace string, name string) (pVCObject *PVCObject, err error)
	// GetStorageClass description:
	// Get StorageClass
	GetStorageClass(ctx context.Context, cluster string, name string) (storageClass *StorageClass, err error)
	// GetStorageService description:
	// Get StorageService
	GetStorageService(ctx context.Context, name string) (storageService *StorageService, err error)
	// ListCluster description:
	// List the cluster and returns an array
	ListCluster(ctx context.Context, listClusterRequest *ListClusterRequest) (clusterList *ClusterList, err error)
	// ListMachine description:
	// List the machine and returns an array.
	//
	// Old: GET /machines
	ListMachine(ctx context.Context, listMachineRequest *ListMachineRequest) (machineList *MachineList, err error)
	// ListNode description:
	// List the nodes and returns an array.
	//
	// Old: GET /clusters/{cluster}/nodes
	ListNode(ctx context.Context, listNodeRequest *ListNodeRequest) (nodeList *NodeList, err error)
	// ListNodeImage description:
	// List the node images and returns an array.
	//
	// Old: GET /clusters/{cluster}/nodes/{node}/images
	ListNodeImage(ctx context.Context, listNodeImagesRequest *ListNodeImagesRequest) (nodeImageList *NodeImageList, err error)
	// ListNodePod description:
	// List the node pods and returns an array.
	//
	// Old: GET /clusters/{cluster}/nodes/{node}/pods
	ListNodePod(ctx context.Context, listNodePodsRequest *ListNodePodsRequest) (nodePodList *NodePodList, err error)
	// ListPersistentVolumeClaim description:
	// List PersistentVolumeClaim
	ListPersistentVolumeClaim(ctx context.Context, cluster string, namespace string, storageClass string, name string, isMountable *bool, start int, limit int) (pVCList *PVCList, err error)
	// ListStorageClass description:
	// List StorageClass
	ListStorageClass(ctx context.Context, cluster string, typeName string, service string, name string, start int, limit int) (storageClassList *StorageClassList, err error)
	// ListStorageService description:
	// List StorageService
	ListStorageService(ctx context.Context, typeName string, name string, start int, limit int) (storageServiceList *StorageServiceList, err error)
	// ListStorageType description:
	// List StorageType
	ListStorageType(ctx context.Context, start int, limit int) (storageTypeList *StorageTypeList, err error)
	// UnbindMachine description:
	// Unbind the machine from cluster and returns an object.
	//
	// Old: POST /clusters/{cluster}/machinesunbind
	UnbindMachine(ctx context.Context, unbindMachineRequest *UnbindMachineRequest) (unbindMachineResponse *UnbindMachineResponse, err error)
	// UpdateClusterAlias description:
	// Update the cluster alias and returns it
	UpdateClusterAlias(ctx context.Context, updateClusterAliasRequest *UpdateClusterAliasRequest) (cluster *Cluster, err error)
	// UpdateMachineAuth description:
	// Update the machine auth and returns an object.
	//
	// Old: PUT /machines/{machine}/auth
	UpdateMachineAuth(ctx context.Context, updateMachineAuthRequest *UpdateMachineAuthRequest) (auth *Auth, err error)
	// UpdateMachineTag description:
	// Update the machine tags and returns an object.
	//
	// Old: PUT /machines/{machine}/tags
	UpdateMachineTag(ctx context.Context, updateMachineTagsRequest *UpdateMachineTagsRequest) (machine *Machine, err error)
	// UpdateNodeOnline description:
	// Update the node online status and returns an object.
	//
	// Old: PUT /clusters/{cluster}/nodes/{node}/online
	UpdateNodeOnline(ctx context.Context, updateNodeOnlineRequest *UpdateNodeOnlineRequest) (node *Node, err error)
	// UpdateNodeTag description:
	// Update the node tags and returns an object.
	//
	// Old: PUT /clusters/{cluster}/nodes/{node}/tags
	UpdateNodeTag(ctx context.Context, updateNodeTagsRequest *UpdateNodeTagsRequest) (node *Node, err error)
	// UpdateNodeTaint description:
	// Update the node taints and returns an object.
	//
	// Old: PUT /clusters/{cluster}/nodes/{node}/taints
	UpdateNodeTaint(ctx context.Context, updateNodeTaintsRequest *UpdateNodeTaintsRequest) (node *Node, err error)
	// UpdatePersistentVolumeClaim description:
	// Update PersistentVolumeClaim
	UpdatePersistentVolumeClaim(ctx context.Context, updatePVCRequest *UpdatePVCRequest) (pVCObject *PVCObject, err error)
	// UpdateStorageClass description:
	// Update StorageClass
	UpdateStorageClass(ctx context.Context, updateStorageClassRequest *UpdateStorageClassRequest) (storageClass *StorageClass, err error)
	// UpdateStorageService description:
	// Update StorageService
	UpdateStorageService(ctx context.Context, updateStorageServiceRequest *UpdateStorageServiceRequest) (storageService *StorageService, err error)
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

// BindMachine description:
// Bind the machine to cluster and returns an object.
//
// Old: POST /clusters/{cluster}/machinesbind
func (c *Client) BindMachine(ctx context.Context, bindMachineRequest *BindMachineRequest) (bindMachineResponse *BindMachineResponse, err error) {
	bindMachineResponse = new(BindMachineResponse)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=BindMachine").
		Body("application/json", bindMachineRequest.BindMachineRequestBody).
		TOPRPCData(bindMachineResponse).
		Do(ctx)
	return
}

// CheckStorageParameters description:
// Check StorageClass parameters
func (c *Client) CheckStorageParameters(ctx context.Context, storageClusterParameters *StorageClusterParameters) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CheckStorageParameters").
		Body("application/json", storageClusterParameters).
		Do(ctx)
	return
}

// CreateCluster description:
// Create the cluster and returns it
func (c *Client) CreateCluster(ctx context.Context, createClusterRequest *CreateClusterRequest) (cluster *Cluster, err error) {
	cluster = new(Cluster)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateCluster").
		Query("Dryrun", createClusterRequest.Dryrun).
		Body("application/json", createClusterRequest.CreateClusterRequestBody).
		TOPRPCData(cluster).
		Do(ctx)
	return
}

// CreateMachine description:
// Create the machine and returns an object.
//
// Old: POST /machines
func (c *Client) CreateMachine(ctx context.Context, createMachineRequest *CreateMachineRequest) (createMachineResponse *CreateMachineResponse, err error) {
	createMachineResponse = new(CreateMachineResponse)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateMachine").
		Body("application/json", createMachineRequest.CreateMachineRequestBody).
		TOPRPCData(createMachineResponse).
		Do(ctx)
	return
}

// CreatePersistentVolumeClaim description:
// Create PersistentVolumeClaim
func (c *Client) CreatePersistentVolumeClaim(ctx context.Context, createPVCRequest *CreatePVCRequest) (pVCObject *PVCObject, err error) {
	pVCObject = new(PVCObject)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreatePersistentVolumeClaim").
		Body("application/json", createPVCRequest).
		TOPRPCData(pVCObject).
		Do(ctx)
	return
}

// CreateStorageClass description:
// Create StorageClass
func (c *Client) CreateStorageClass(ctx context.Context, createStorageClassRequest *CreateStorageClassRequest) (storageClass *StorageClass, err error) {
	storageClass = new(StorageClass)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStorageClass").
		Body("application/json", createStorageClassRequest).
		TOPRPCData(storageClass).
		Do(ctx)
	return
}

// CreateStorageService description:
// Create StorageService
func (c *Client) CreateStorageService(ctx context.Context, createStorageServiceRequest *CreateStorageServiceRequest) (storageService *StorageService, err error) {
	storageService = new(StorageService)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateStorageService").
		Body("application/json", createStorageServiceRequest).
		TOPRPCData(storageService).
		Do(ctx)
	return
}

// DeleteCluster description:
// Delete the cluster and returns it
func (c *Client) DeleteCluster(ctx context.Context, deleteClusterRequest *DeleteClusterRequest) (cluster *Cluster, err error) {
	cluster = new(Cluster)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteCluster").
		Query("Cluster", deleteClusterRequest.Name).
		TOPRPCData(cluster).
		Do(ctx)
	return
}

// DeleteMachine description:
// Delete the machine and returns an object.
//
// Old: DELETE /machines/{machine}
func (c *Client) DeleteMachine(ctx context.Context, deleteMachineRequest *DeleteMachineRequest) (machine *Machine, err error) {
	machine = new(Machine)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteMachine").
		Query("Name", deleteMachineRequest.Name).
		Query("DeletePolicy", deleteMachineRequest.DeletePolicy).
		TOPRPCData(machine).
		Do(ctx)
	return
}

// DeletePersistentVolumeClaim description:
// Delete PersistentVolumeClaim
func (c *Client) DeletePersistentVolumeClaim(ctx context.Context, cluster string, namespace string, name string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeletePersistentVolumeClaim").
		Query("Cluster", cluster).
		Query("Namespace", namespace).
		Query("Name", name).
		Do(ctx)
	return
}

// DeleteStorageClass description:
// Delete StorageClass
func (c *Client) DeleteStorageClass(ctx context.Context, cluster string, name string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteStorageClass").
		Query("Cluster", cluster).
		Query("Name", name).
		Do(ctx)
	return
}

// DeleteStorageService description:
// Delete StorageService
func (c *Client) DeleteStorageService(ctx context.Context, name string) (err error) {
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteStorageService").
		Query("Name", name).
		Do(ctx)
	return
}

// GetCluster description:
// Get the cluster and returns it
func (c *Client) GetCluster(ctx context.Context, getClusterRequest *GetClusterRequest) (cluster *Cluster, err error) {
	cluster = new(Cluster)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetCluster").
		Query("Cluster", getClusterRequest.Name).
		TOPRPCData(cluster).
		Do(ctx)
	return
}

// GetMachine description:
// Get the machine and returns an object.
//
// Old: GET /machines/{machine}
func (c *Client) GetMachine(ctx context.Context, getMachineRequest *GetMachineRequest) (machine *Machine, err error) {
	machine = new(Machine)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetMachine").
		Query("Name", getMachineRequest.Name).
		TOPRPCData(machine).
		Do(ctx)
	return
}

// GetMachineAuth description:
// Get the machine auth and returns an object.
//
// Old: GET /machines/{machine}/auth
func (c *Client) GetMachineAuth(ctx context.Context, getMachineAuthRequest *GetMachineAuthRequest) (auth *Auth, err error) {
	auth = new(Auth)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetMachineAuth").
		Query("Name", getMachineAuthRequest.Name).
		TOPRPCData(auth).
		Do(ctx)
	return
}

// GetNode description:
// Get the node and returns an object.
//
// Old: GET /clusters/{cluster}/nodes/{node}
func (c *Client) GetNode(ctx context.Context, getNodeRequest *GetNodeRequest) (node *Node, err error) {
	node = new(Node)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetNode").
		Query("Cluster", getNodeRequest.Cluster).
		Query("Name", getNodeRequest.Name).
		TOPRPCData(node).
		Do(ctx)
	return
}

// GetPersistentVolumeClaim description:
// Get PersistentVolumeClaim
func (c *Client) GetPersistentVolumeClaim(ctx context.Context, cluster string, namespace string, name string) (pVCObject *PVCObject, err error) {
	pVCObject = new(PVCObject)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetPersistentVolumeClaim").
		Query("Cluster", cluster).
		Query("Namespace", namespace).
		Query("Name", name).
		TOPRPCData(pVCObject).
		Do(ctx)
	return
}

// GetStorageClass description:
// Get StorageClass
func (c *Client) GetStorageClass(ctx context.Context, cluster string, name string) (storageClass *StorageClass, err error) {
	storageClass = new(StorageClass)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStorageClass").
		Query("Cluster", cluster).
		Query("Name", name).
		TOPRPCData(storageClass).
		Do(ctx)
	return
}

// GetStorageService description:
// Get StorageService
func (c *Client) GetStorageService(ctx context.Context, name string) (storageService *StorageService, err error) {
	storageService = new(StorageService)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetStorageService").
		Query("Name", name).
		TOPRPCData(storageService).
		Do(ctx)
	return
}

// ListCluster description:
// List the cluster and returns an array
func (c *Client) ListCluster(ctx context.Context, listClusterRequest *ListClusterRequest) (clusterList *ClusterList, err error) {
	clusterList = new(ClusterList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListCluster").
		Query("Status", listClusterRequest.Status).
		Query("Start", listClusterRequest.Start).
		Query("Limit", listClusterRequest.Limit).
		TOPRPCData(clusterList).
		Do(ctx)
	return
}

// ListMachine description:
// List the machine and returns an array.
//
// Old: GET /machines
func (c *Client) ListMachine(ctx context.Context, listMachineRequest *ListMachineRequest) (machineList *MachineList, err error) {
	machineList = new(MachineList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListMachine").
		Query("Cluster", listMachineRequest.Cluster).
		Query("RawTags", listMachineRequest.RawTags).
		Query("IP", listMachineRequest.IP).
		Query("Status", listMachineRequest.Status).
		Query("Provider", listMachineRequest.Provider).
		Query("Start", listMachineRequest.Start).
		Query("Limit", listMachineRequest.Limit).
		TOPRPCData(machineList).
		Do(ctx)
	return
}

// ListNode description:
// List the nodes and returns an array.
//
// Old: GET /clusters/{cluster}/nodes
func (c *Client) ListNode(ctx context.Context, listNodeRequest *ListNodeRequest) (nodeList *NodeList, err error) {
	nodeList = new(NodeList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListNode").
		Query("Cluster", listNodeRequest.Cluster).
		Query("Schedulable", listNodeRequest.Schedulable).
		Query("IP", listNodeRequest.IP).
		Query("Status", listNodeRequest.Status).
		Query("RawTags", listNodeRequest.RawTags).
		Query("Start", listNodeRequest.Start).
		Query("Limit", listNodeRequest.Limit).
		TOPRPCData(nodeList).
		Do(ctx)
	return
}

// ListNodeImage description:
// List the node images and returns an array.
//
// Old: GET /clusters/{cluster}/nodes/{node}/images
func (c *Client) ListNodeImage(ctx context.Context, listNodeImagesRequest *ListNodeImagesRequest) (nodeImageList *NodeImageList, err error) {
	nodeImageList = new(NodeImageList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListNodeImage").
		Query("Cluster", listNodeImagesRequest.Cluster).
		Query("Name", listNodeImagesRequest.Name).
		Query("Start", listNodeImagesRequest.Start).
		Query("Limit", listNodeImagesRequest.Limit).
		TOPRPCData(nodeImageList).
		Do(ctx)
	return
}

// ListNodePod description:
// List the node pods and returns an array.
//
// Old: GET /clusters/{cluster}/nodes/{node}/pods
func (c *Client) ListNodePod(ctx context.Context, listNodePodsRequest *ListNodePodsRequest) (nodePodList *NodePodList, err error) {
	nodePodList = new(NodePodList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListNodePod").
		Query("Cluster", listNodePodsRequest.Cluster).
		Query("Name", listNodePodsRequest.Name).
		Query("Start", listNodePodsRequest.Start).
		Query("Limit", listNodePodsRequest.Limit).
		TOPRPCData(nodePodList).
		Do(ctx)
	return
}

// ListPersistentVolumeClaim description:
// List PersistentVolumeClaim
func (c *Client) ListPersistentVolumeClaim(ctx context.Context, cluster string, namespace string, storageClass string, name string, isMountable *bool, start int, limit int) (pVCList *PVCList, err error) {
	pVCList = new(PVCList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListPersistentVolumeClaim").
		Query("Cluster", cluster).
		Query("Namespace", namespace).
		Query("StorageClass", storageClass).
		Query("Name", name).
		Query("IsMountable", isMountable).
		Query("Start", start).
		Query("Limit", limit).
		TOPRPCData(pVCList).
		Do(ctx)
	return
}

// ListStorageClass description:
// List StorageClass
func (c *Client) ListStorageClass(ctx context.Context, cluster string, typeName string, service string, name string, start int, limit int) (storageClassList *StorageClassList, err error) {
	storageClassList = new(StorageClassList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStorageClass").
		Query("Cluster", cluster).
		Query("TypeName", typeName).
		Query("Service", service).
		Query("Name", name).
		Query("Start", start).
		Query("Limit", limit).
		TOPRPCData(storageClassList).
		Do(ctx)
	return
}

// ListStorageService description:
// List StorageService
func (c *Client) ListStorageService(ctx context.Context, typeName string, name string, start int, limit int) (storageServiceList *StorageServiceList, err error) {
	storageServiceList = new(StorageServiceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStorageService").
		Query("TypeName", typeName).
		Query("Name", name).
		Query("Start", start).
		Query("Limit", limit).
		TOPRPCData(storageServiceList).
		Do(ctx)
	return
}

// ListStorageType description:
// List StorageType
func (c *Client) ListStorageType(ctx context.Context, start int, limit int) (storageTypeList *StorageTypeList, err error) {
	storageTypeList = new(StorageTypeList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListStorageType").
		Query("Start", start).
		Query("Limit", limit).
		TOPRPCData(storageTypeList).
		Do(ctx)
	return
}

// UnbindMachine description:
// Unbind the machine from cluster and returns an object.
//
// Old: POST /clusters/{cluster}/machinesunbind
func (c *Client) UnbindMachine(ctx context.Context, unbindMachineRequest *UnbindMachineRequest) (unbindMachineResponse *UnbindMachineResponse, err error) {
	unbindMachineResponse = new(UnbindMachineResponse)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UnbindMachine").
		Body("application/json", unbindMachineRequest.UnbindMachineRequestBody).
		TOPRPCData(unbindMachineResponse).
		Do(ctx)
	return
}

// UpdateClusterAlias description:
// Update the cluster alias and returns it
func (c *Client) UpdateClusterAlias(ctx context.Context, updateClusterAliasRequest *UpdateClusterAliasRequest) (cluster *Cluster, err error) {
	cluster = new(Cluster)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateClusterAlias").
		Query("Alias", updateClusterAliasRequest.Alias).
		Query("Cluster", updateClusterAliasRequest.Name).
		TOPRPCData(cluster).
		Do(ctx)
	return
}

// UpdateMachineAuth description:
// Update the machine auth and returns an object.
//
// Old: PUT /machines/{machine}/auth
func (c *Client) UpdateMachineAuth(ctx context.Context, updateMachineAuthRequest *UpdateMachineAuthRequest) (auth *Auth, err error) {
	auth = new(Auth)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateMachineAuth").
		Body("application/json", updateMachineAuthRequest.UpdateMachineAuthRequestBody).
		TOPRPCData(auth).
		Do(ctx)
	return
}

// UpdateMachineTag description:
// Update the machine tags and returns an object.
//
// Old: PUT /machines/{machine}/tags
func (c *Client) UpdateMachineTag(ctx context.Context, updateMachineTagsRequest *UpdateMachineTagsRequest) (machine *Machine, err error) {
	machine = new(Machine)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateMachineTag").
		Body("application/json", updateMachineTagsRequest.UpdateMachineTagsRequestBody).
		TOPRPCData(machine).
		Do(ctx)
	return
}

// UpdateNodeOnline description:
// Update the node online status and returns an object.
//
// Old: PUT /clusters/{cluster}/nodes/{node}/online
func (c *Client) UpdateNodeOnline(ctx context.Context, updateNodeOnlineRequest *UpdateNodeOnlineRequest) (node *Node, err error) {
	node = new(Node)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateNodeOnline").
		Query("Cluster", updateNodeOnlineRequest.Cluster).
		Query("Name", updateNodeOnlineRequest.Name).
		Query("OnlineStatus", updateNodeOnlineRequest.OnlineStatus).
		TOPRPCData(node).
		Do(ctx)
	return
}

// UpdateNodeTag description:
// Update the node tags and returns an object.
//
// Old: PUT /clusters/{cluster}/nodes/{node}/tags
func (c *Client) UpdateNodeTag(ctx context.Context, updateNodeTagsRequest *UpdateNodeTagsRequest) (node *Node, err error) {
	node = new(Node)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateNodeTag").
		Body("application/json", updateNodeTagsRequest.UpdateNodeTagsRequestBody).
		TOPRPCData(node).
		Do(ctx)
	return
}

// UpdateNodeTaint description:
// Update the node taints and returns an object.
//
// Old: PUT /clusters/{cluster}/nodes/{node}/taints
func (c *Client) UpdateNodeTaint(ctx context.Context, updateNodeTaintsRequest *UpdateNodeTaintsRequest) (node *Node, err error) {
	node = new(Node)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateNodeTaint").
		Body("application/json", updateNodeTaintsRequest.UpdateNodeTaintsRequestBody).
		TOPRPCData(node).
		Do(ctx)
	return
}

// UpdatePersistentVolumeClaim description:
// Update PersistentVolumeClaim
func (c *Client) UpdatePersistentVolumeClaim(ctx context.Context, updatePVCRequest *UpdatePVCRequest) (pVCObject *PVCObject, err error) {
	pVCObject = new(PVCObject)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdatePersistentVolumeClaim").
		Body("application/json", updatePVCRequest).
		TOPRPCData(pVCObject).
		Do(ctx)
	return
}

// UpdateStorageClass description:
// Update StorageClass
func (c *Client) UpdateStorageClass(ctx context.Context, updateStorageClassRequest *UpdateStorageClassRequest) (storageClass *StorageClass, err error) {
	storageClass = new(StorageClass)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStorageClass").
		Body("application/json", updateStorageClassRequest).
		TOPRPCData(storageClass).
		Do(ctx)
	return
}

// UpdateStorageService description:
// Update StorageService
func (c *Client) UpdateStorageService(ctx context.Context, updateStorageServiceRequest *UpdateStorageServiceRequest) (storageService *StorageService, err error) {
	storageService = new(StorageService)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateStorageService").
		Body("application/json", updateStorageServiceRequest).
		TOPRPCData(storageService).
		Do(ctx)
	return
}
