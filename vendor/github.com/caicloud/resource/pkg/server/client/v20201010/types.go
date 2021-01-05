package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	time "time"
)

// Address defines the machine address
type Address struct {
	Type     string `json:"Type"`
	Address  string `json:"Address"`
	IPV6Addr string `json:"IPV6Addr"`
}

// Auth defines the machine auth
type Auth struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Key      string `json:"Key"`
	Port     int32  `json:"Port"`
}

// BindMachineRequest defines the request for bind machine to cluster
type BindMachineRequest struct {
	BindMachineRequestBody `source:"Body" validate:"required,dive"`
}

// BindMachineRequestBody defines the request body for bind machine to cluster
type BindMachineRequestBody struct {
	Cluster  string   `json:"Cluster,omitempty" valdiate:"required,min=1"`
	Machines []string `json:"Machines,omitempty" validate:"required,min=1,dive,min=1"`
	// +optional
	Taints Taints `json:"Taints,omitempty"`
}

// BindMachineResponse defines the response for bind machine to cluster
type BindMachineResponse struct {
	Machines []string `json:"Machines,omitempty"`
}

// CephMetaData is the data structure for Ceph metadata.
type CephMetaData struct {
	Pools []CephPool `json:"Pools"`
}

// CephPool is the data structure for single Ceph storage pool.
type CephPool struct {
	Name        string `json:"Name"`
	ReplicaSize int    `json:"ReplicaSize"`
	// Deprecated
	// total capacity of the current pool
	Capacity     string  `json:"Capacity"`
	MaxAvailable string  `json:"MaxAvailable"`
	PercentUsed  float32 `json:"PercentUsed"`
	// capacity used
	Used string `json:"Used"`
	// number of objects in the pool
	Objects int `json:"Objects"`
}

// Cluster defines the response for cluster
type Cluster struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	// Alias is the display name
	Alias string `json:"Alias"`
	// IsControlCluster defines the cluster role
	IsControlCluster bool `json:"IsControlCluster"`
	// Provider defines the provider name
	// Available: containeros.dev/baremetal
	Provider string `json:"Provider"`
	// IsHighAvailable defines the cluster high available
	IsHighAvailable bool `json:"IsHighAvailable"`
	// IsEnhanced defines the cluster is enhanced
	IsEnhanced bool `json:"IsEnhanced"`
	// MastersVIP defines the vip when cluster is high available
	MastersVIP string `json:"MastersVIP"`
	// KubernetesVersion defines the cluster k8s version
	// Available: v1.19.3
	KubernetesVersion string `json:"KubernetesVersion"`
	// Network defines the cluster network attr
	Network ClusterNetwork `json:"ClusterNetwork"`
	// Phase defines the cluster phase
	// Available: Pending, Provisioning, Provisioned, Bootstraping, Running, Deleting, Failed, Unknown
	// New: -> Pending
	// InstallMaster: -> Provisioning, Provisioned, Bootstraping
	// InstallAddon: ->
	// Ready: -> Running
	// Deleting: -> Deleting
	// Failed: -> Failed
	Phase string `json:"Phase"`
}

// nolint
// ClusterList defines the response for list cluster
type ClusterList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []Cluster `json:"Items"`
}

// nolint
// ClusterNetwork defines the response for cluster network
type ClusterNetwork struct {
	// DNSServiceIP defines the ip of dns service
	DNSServiceIP string `json:"DNSServiceIP`
	// ClusterCIDR defines the cluster cidr
	ClusterCIDR string `json:"ClusterCIDR"`
	// ServiceCIDR defines the service cidr
	ServiceCIDR string `json:"ServiceCIDR"`
}

// CreateClusterRequest defines the request for create cluster
type CreateClusterRequest struct {
	Dryrun                   bool `json:"Dryrun" source:"Query,Dryrun"`
	CreateClusterRequestBody `source:"Body"`
}

// CreateClusterRequestBody defines the request body for create cluster
type CreateClusterRequestBody struct {
	// Alias is the display name
	Alias string `json:"Alias" validate:"required,min=2,max=32"`
	// Provider defines the provider name
	// Available: containeros.dev/baremetal
	Provider string `json:"Provider" validate:"required,oneof=containeros.dev/baremetal"`
	// IsHighAvailable defines the high available attr
	// NOTE: if the field is true, the MastersVIP must be set and Masters must gt 1
	IsHighAvailable bool `json:"IsHighAvailable"`
	// MastersVIP defines the vip
	// +optional
	MastersVIP string `json:"MastersVIP" validate:"omitempty,ip"`
	// NetworkType defines the network type, calico_backend
	// Available: vxlan, bird
	NetworkType string `json:"NetworkType" validate:"required,oneof=vxlan bird"`
	// KubeServiceCIDR defines the service cidr
	KubeServiceCIDR string `json:"KubeServiceCIDR" validate:"required,cidr"`
	// KubeDNSServiceIP defines the dns service ip
	KubeDNSServiceIP string `json:"KubeDNSServiceIP" validate:"required,ip"`
	// ClusterCIDR defines the cluster cidr
	ClusterCIDR string `json:"ClusterCIDR" validate:"required,cidr"`
	// DefaultNetwork defines the default network spec
	DefaultNetwork NetworkSpec `json:"DefaultNetwork"`
	// Masters defines the master machine's name array
	Masters []string `json:"Masters" validate:"required,min=1"`
	// Nodes defines the node machine's name array
	// +optional
	Nodes []string `json:"Nodes"`
	// Etcds defines the etcd machine's name array
	// +optional
	Etcds []string `json:"Etcds"`
	// KubeType defines the kubernetes type
	// Available: enhanced, original
	KubeType string `json:"KubeType" validate:"required,oneof=enhanced original"`
}

// CreateMachineRequest defines the request for create machine
type CreateMachineRequest struct {
	CreateMachineRequestBody `source:"Body"`
}

// CreateMachineRequestBody defines the request body for create machine
type CreateMachineRequestBody struct {
	// Provider defines the provider name
	// Available: containeros.dev/baremetal
	Provider string `json:"Provider,omitempty" validate:"required,oneof=containeros.dev/baremetal"`
	// Addresses is the ip array.
	Addresses []string `json:"Addresses,omitempty" validate:"required,min=1,dive,ip"`
	User      string   `json:"User,omitempty" validate:"required,min=1"`
	Port      int32    `json:"Port,omitempty" validate:"required,gte=1"`
	// +optional
	Password string `json:"Password,omitempty"`
	// +optional
	// Available: ssh-login-config
	Key string `json:"Key,omitempty" validate:"omitempty,oneof=ssh-login-config"`
	// +optional
	Tags Tags `json:"Tags,omitempty" validate:"dive"`
}

// CreateMachineResponse defines the response body for create machine
type CreateMachineResponse struct {
	Machines []string `json:"Machines"`
	Done     []string `json:"Done"`
	Failed   []string `json:"Failed"`
}

// CreatePVCRequest defines the data structure for PVC POST API request.
type CreatePVCRequest struct {
	Cluster      string `json:"Cluster"`
	Namespace    string `json:"Namespace"`
	Name         string `json:"Name"`
	StorageClass string `json:"StorageClass"`
	Size         string `json:"Size"`
	// for manually create the PV
	Type string `json:"Type"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters string `json:"Parameters"`
}

// CreateStorageClassRequest defines the data structure for StorageClass POST API request.
type CreateStorageClassRequest struct {
	Cluster     string `json:"Cluster"`
	Name        string `json:"Name"`
	Alias       string `json:"Alias,omitempty"`
	Description string `json:"Description,omitempty"`
	Service     string `json:"Service"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters string              `json:"Parameters"`
	Type       string              `json:"Type"`
	Quota      StorageQuotaRequest `json:"Quota"`
}

// CreateStorageServiceRequest defines the data structure for StorageService POST API request.
type CreateStorageServiceRequest struct {
	Name        string `json:"Name"`
	Alias       string `json:"Alias,omitempty"`
	Description string `json:"Description,omitempty"`
	Type        string `json:"Type"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters string              `json:"Parameters"`
	Quota      StorageQuotaRequest `json:"Quota"`
}

// DeleteClusterRequest defines the request for delete cluster
type DeleteClusterRequest struct {
	Name string `source:"Query,Cluster" validate:"required,min=1"`
}

// DeleteMachineRequest defines the request for delete machine
type DeleteMachineRequest struct {
	Name         string `json:"Name" source:"Query,Name" validate:"required,min=1"`
	DeletePolicy string `json:"DeletePolicy" source:"Query,DeletePolicy,optional"`
}

// GetClusterRequest defines the request for get cluster
type GetClusterRequest struct {
	Name string `source:"Query,Cluster" validate:"required,min=1"`
}

// GetMachineAuthRequest defines the request for get machine auth
type GetMachineAuthRequest struct {
	Name string `json:"Name" source:"Query,Name" valdiate:"required,min=1"`
}

// GetMachineRequest defines the request for get machine
type GetMachineRequest struct {
	Name string `json:"Name" source:"Query,Name" validate:"required,min=1"`
}

// GetNodeRequest defines the request for get node
type GetNodeRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster" validate:"required,min=1"`
	Name    string `json:"Name" source:"Query,Name" validate:"required,min=1"`
}

// ListClusterRequest defines the request for list cluster
type ListClusterRequest struct {
	// Status defines the cluster status for filter
	// EX: Pending,Running
	Status string `json:"Status" source:"Query,Status" validate:"omitempty,min=1"`
	Start  int    `json:"Start" source:"Query,Start,default=0"`
	Limit  int    `json:"Limit" source:"Query,Limit,default=0"`
}

// ListMachineRequest defines the request for list machine
type ListMachineRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster,omitempty" source:"Query,Cluster,optional" validate:"omitempty,min=1"`
	// RawTags defines the selected tags.
	// EX: 'k1=v1,k2=v2'
	RawTags string `json:"RawTags" source:"Query,RawTags,optional"`
	IP      string `json:"IP,omitempty" source:"Query,IP,optional"`
	// Status defines the status of machine.
	// Available: Allocate, New, Failed, Removing, Free
	Status string `json:"Status" source:"Query,Status,optional" validate:"omitempty,oneof=Allocate New Failed Removing Free"`
	// Provider defines the provider name
	// Available: containeros.dev/baremetal
	Provider string `json:"Provider" source:"Query,Provider,optional" validate:"omitempty,oneof=containeros.dev/baremetal"`
	Start    int    `source:"Query,Start,default=0" json:"Start"`
	Limit    int    `source:"Query,Limit,default=10" json:"Limit"`
}

// ListNodeImagesRequest defines the request for list node images
type ListNodeImagesRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster" validate:"required,min=1"`
	Name    string `json:"Name" source:"Query,Name" validate:"required,min=1"`
	Start   int    `source:"Query,Start,default=0"`
	Limit   int    `source:"Query,Limit,default=10"`
}

// ListNodePodsRequest defines the request for list node pods
type ListNodePodsRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster" validate:"required,min=1"`
	Name    string `json:"Name" source:"Query,Name" validate:"required,min=1"`
	Start   int    `source:"Query,Start,default=0"`
	Limit   int    `source:"Query,Limit,default=10"`
}

// ListNodeRequest defines the request for list node
type ListNodeRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster     string `json:"Cluster" source:"Query,Cluster" validate:"required,min=1"`
	Schedulable *bool  `json:"Schedulable" source:"Query,Schedulable,optional"`
	IP          string `json:"IP" source:"Query,IP,optional"`
	Status      string `json:"Status" source:"Query,Status,optional"`
	// RawTags defines the selected tags.
	// EX: 'k1=v1,k2=v2'
	RawTags string `json:"RawTags" source:"Query,RawTags,optional"`
	Start   int    `source:"Query,Start,default=0"`
	Limit   int    `source:"Query,Limit,default=10"`
}

// Machine defines the response for machine
type Machine struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	Alias         string    `json:"Alias"`
	Provider      string    `json:"Provider"`
	Cluster       string    `json:"Cluster"`
	IsMaster      bool      `json:"IsMaster"`
	IsEtcd        bool      `json:"IsEtcd"`
	Tags          Tags      `json:"Tags,omitempty"`
	Taints        Taints    `json:"Taints,omitempty"`
	Addresses     []Address `json:"Addresses"`
	CPUCores      string    `json:"CPUCores"`
	Memory        string    `json:"Memory"`
	Disk          string    `json:"Disk"`
	GPUCores      string    `json:"GPUCores"`
	Auth          Auth      `json:"Auth"`
	// Available: Allocated, New, Failed, Deleting, Free
	Phase     string `json:"Phase"`
	NodeRefer string `json:"NodeRefer"`
}

// nolint
// MachineList defines the response for list machine
type MachineList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []Machine `json:"Items"`
}

// NetworkSpec defines the network advance options
type NetworkSpec struct {
	// IsFixedIP defines the fixip attr
	// NOTE: available when type is bird
	// +optional
	IsFixedIP bool `json:"IsFixedIP"`
	// IPRecycleTimeout defines the ip reuse when node down
	// NOTE: -1 when this is disable, available when type is bird and fixedip is enabled
	// +Default: -1
	// +optional
	IPRecycleTimeout *int32 `json:"IPRecycleTimeout"`
	// NodeCIDRMaskSize defines the node mask
	// +Default: 24
	// +optional
	NodeCIDRMaskSize int `json:"NodeCIDRMaskSize"`
	// RouteReflectorIPs defines the ip when type is bird router
	// EX: 1.1.1.1,2.2.2.2
	// +optional
	RouteReflectorIPs string `json:"RouteReflectorIPs"`
	// AsNumber defines the as number when bird router is enabled
	// NOTE: [64512, 65534]
	// +optional
	AsNumber int `json:"AsNumber" validate:"omitempty,min=64512,max=65534"`
}

// Node defines the response for node
type Node struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	// Status defines the node status
	// Available: NodeNotReady, SetOffline, Offline, Ready, Binding, Unbinding
	Status           string        `json:"Status"`
	ClusterName      string        `json:"ClusterName"`
	MachineName      string        `json:"MachineName"`
	IsSchedulable    bool          `json:"IsSchedulable"`
	IsMaster         bool          `json:"IsMaster"`
	IsEtcd           bool          `json:"IsEtcd"`
	Addresses        []NodeAddress `json:"Addresses"`
	CPUCores         string        `json:"CPUCores"`
	Memory           string        `json:"Memory"`
	GPUCores         string        `json:"GPUCores"`
	Disk             string        `json:"Disk"`
	EphemeralStorage string        `json:"EphemeralStorage"`
	Taints           Taints        `json:"Taints"`
	NodeInfo         SystemInfo    `json:"NodeInfo"`
}

// nolint
// NodeAddress defines the node address
type NodeAddress struct {
	Type    string `json:"Type"`
	Address string `json:"Address"`
}

// nolint
// NodeImage defines the response for get node image
type NodeImage struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	Names         []string `json:"Names"`
	SizeBytes     int64    `json:"SizeBytes"`
}

// nolint
// NodeImageList defines the response for list node images
type NodeImageList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []NodeImage `json:"Items"`
}

// nolint
// NodeList defines the response for list node
type NodeList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []Node `json:"Items"`
}

// nolint
// NodePod defines the response for get node pod
type NodePod struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	Status        string    `json:"Status"`
	IP            string    `json:"IP"`
	IPv6          string    `json:"IPv6"`
	StartTime     time.Time `json:"StartTime"`
	Images        []string  `json:"Images"`
	RestartCounts []int32   `json:"RestartCounts"`
}

// nolint
// NodePodList defines the response for list node pods
type NodePodList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []NodePod `json:"Items"`
}

// PVCList defines the data structure for PVC LIST API response.
type PVCList struct {
	v1.ListMeta `json:",inline"`
	Items       []PVCObject `json:"Items"`
}

// PVCObject defines the type of PVC object.
type PVCObject struct {
	v1.ObjectMeta `json:",inline"`
	Spec          PVCSpec   `json:"Spec"`
	Status        PVCStatus `json:"Status"`
}

// PVCSpec describes the common attributes of the data volume object.
type PVCSpec struct {
	AccessModes       []string             `json:"AccessModes"`
	Resources         ResourceRequirements `json:"Resources"`
	VolumeName        string               `json:"VolumeName"`
	StorageClassName  *string              `json:"StorageClassName"`
	StorageClassAlias string               `json:"StorageClassAlias"`
	Type              string               `json:"Type,omitempty"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters string `json:"Parameters,omitempty"`
}

// PVCStatus is the current status of a data volume object.
type PVCStatus struct {
	Phase       string   `json:"Phase"`
	AccessModes []string `json:"AccessModes"`
	// The underlying value is map[string]string, eg:
	// {
	// "storage": "1Gi"
	// }
	Capacity string `json:"Capacity"`
	// The underlying value is map[string]string, eg:
	// {
	// "storage": "1Gi"
	// }
	Used string `json:"Used,omitempty"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// The underlying value is map[string]string, eg:
	// {
	// "storage": "1Gi"
	// }
	Requests string `json:"Requests,omitempty"`
}

// StorageClass is the data structure for StorageClass.
// nolint
type StorageClass struct {
	v1.ObjectMeta `json:",inline"`
	Provisioner   string `json:"Provisioner"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters  string       `json:"Parameters"`
	Quota       StorageQuota `json:"Quota"`
	TypeName    string       `json:"TypeName"`
	ServiceName string       `json:"ServiceName"`
}

// StorageClassList defines the data structure for StorageClass LIST API response.
// nolint
type StorageClassList struct {
	v1.ListMeta `json:",inline"`
	Items       []StorageClass `json:"Items"`
}

// StorageClusterParameters defines the API request structure for checking storage cluster parameters.
type StorageClusterParameters struct {
	Type              string            `json:"Type"`
	NFS               nfs               `json:"Nfs,omitempty"`
	GlusterFS         glusterFS         `json:"GlusterFS,omitempty"`
	ExistingGlusterFS existingGlusterFS `json:"ExistingGlusterFS,omitempty"`
	Ceph              ceph              `json:"Ceph,omitempty"`
}

// StorageMetaData is the data structure for each storage backend metadata.
// nolint
type StorageMetaData struct {
	Ceph *CephMetaData `json:"Ceph,omitempty"`
}

// StorageQuota defines the data structure for full storage quota information.
// nolint
type StorageQuota struct {
	// The underlying value is map[string]string, eg:
	// {
	// "persistentvolumeclaims": "",
	// "requests.storage": ""
	// }
	// persistentvolumeclaims represents the number of PVC that can be allocated, 0 means no limit.
	// requests.storage represents the total storage capacity that can be allocated, must be greater than 0.
	Hard string `json:"Hard"`
	// The underlying value is map[string]string, eg:
	// {
	// "persistentvolumeclaims": "",
	// "requests.storage": ""
	// }
	// persistentvolumeclaims represents the number of PVC that can be allocated, 0 means no limit.
	// requests.storage represents the total storage capacity that can be allocated, must be greater than 0.
	Allocated string `json:"Allocated"`
	// compatible with StorageService/StorageClass created by previous versions of Compass 2.9,
	// they always allow setting the amount of PVC.
	//
	// For StorageClass, this field is only for StorageClasses that do not have a StorageService
	// (local storage, azure disk, azure file)
	// The rest are judged by their own StorageService.
	PVCAllowUpdate bool `json:"PvcAllowUpdate"`
	// similar to the `PVCAllowUpdate` field, the front end uses this field to prompt the user to set the quota.
	NoQuotaSet bool `json:"NoQuotaSet"`
}

// StorageQuotaRequest defines the data structure storage quota.
// nolint
type StorageQuotaRequest struct {
	// PersistentVolumeClaims represents the number of PVC that can be allocated, 0 means no limit.
	PersistentVolumeClaims string `json:"PersistentVolumeClaims"`
	// RequestsStorage represents the total storage capacity that can be allocated, must be greater than 0.
	RequestsStorage string `json:"RequestsStorage"`
}

// StorageService is the data structure for StorageService.
// nolint
type StorageService struct {
	v1.ObjectMeta `json:",inline"`
	Type          StorageTypeInfoObject `json:"Type"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	Parameters      string          `json:"Parameters"`
	StorageMetaData StorageMetaData `json:"StorageMetaData"`
	Quota           StorageQuota    `json:"Quota"`
}

// StorageServiceList defines the data structure for StorageService LIST API response.
// nolint
type StorageServiceList struct {
	v1.ListMeta `json:",inline"`
	Items       []StorageService `json:"Items"`
}

// StorageType is the data structure for StorageType.
// nolint
type StorageType struct {
	v1.ObjectMeta `json:",inline"`
	Provisioner   string `json:"Provisioner"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	CommonParameters string `json:"CommonParameters"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	OptionalParameters string `json:"OptionalParameters"`
}

// StorageTypeInfoObject is the streamlined data structure for StorageType, used in StorageServiceObject.
// nolint
type StorageTypeInfoObject struct {
	Name        string `json:"Name"`
	Provisioner string `json:"Provisioner"`
	// The underlying value is map[string]string, eg:
	// {
	// "key1": "",
	// "key2": ""
	// }
	OptionalParameters string `json:"OptionalParameters"`
}

// StorageTypeList defines the data structure for StorageType LIST API response.
// nolint
type StorageTypeList struct {
	v1.ListMeta `json:",inline"`
	Items       []StorageType `json:"Items"`
}

// SystemInfo defines the node system info
type SystemInfo struct {
	MachineID               string `json:"MachineID"`
	SystemUUID              string `json:"SystemUUID"`
	BootID                  string `json:"BootID"`
	KernelVersion           string `json:"KernelVersion"`
	OSImage                 string `json:"OSImage"`
	ContainerRuntimeVersion string `json:"ContainerRuntimeVersion"`
	KubeletVersion          string `json:"KubeletVersion"`
	KubeProxyVersion        string `json:"KubeProxyVersion"`
	OperatingSystem         string `json:"OperatingSystem"`
	Architecture            string `json:"Architecture"`
}

// Tag defines the tag object
type Tag struct {
	Key   string `json:"Key,omitempty" validate:"required"`
	Value string `json:"Value,omitempty" validate:"required"`
}

// Tags defines the tag array object
type Tags []Tag

// Taint defines the taint on object
type Taint struct {
	Key   string `json:"Key,omitempty" validate:"required,min=1"`
	Value string `json:"Value,omitempty" validate:"required,min=1"`
	// Available: NoExecute, NoSchedule, PreferNoSchedule
	Effect string `json:"Effect,omitempty" validate:"required,oneof=NoExecute NoSchedule PreferNoSchedule"`
}

// Taints defines the taint array object
type Taints []Taint

// UnbindMachineRequest defines the request for bind machine from cluster
type UnbindMachineRequest struct {
	UnbindMachineRequestBody `source:"Body" validate:"required,dive"`
}

// UnbindMachineRequestBody defines the request body for bind machine from cluster
type UnbindMachineRequestBody struct {
	Machines []string `json:"Machines,omitempty" validate:"required,min=1,dive,min=1"`
}

// UnbindMachineResponse defines the response for unbind machine to cluster
type UnbindMachineResponse struct {
	Machines []string `json:"Machines"`
}

// UpdateClusterAliasRequest defines the request for update cluster alias
type UpdateClusterAliasRequest struct {
	Alias string `json:"Alias" source:"Query,Alias" validate:"required,min=1"`
	Name  string `source:"Query,Cluster" validate:"required,min=1"`
}

// UpdateMachineAuthRequest defines the request for update machine auth
type UpdateMachineAuthRequest struct {
	UpdateMachineAuthRequestBody `source:"Body" validate:"required,dive"`
}

// UpdateMachineAuthRequestBody defines the request for update machine auth
type UpdateMachineAuthRequestBody struct {
	Name string `json:"Name" validate:"required,min=1"`
	User string `json:"User" validate:"required,min=1"`
	// +optional
	Password string `json:"Password"`
	// +optional
	// Available: ssh-login-config
	Key string `json:"Key"`
}

// UpdateMachineTagsRequest defines the request for update machine tags
type UpdateMachineTagsRequest struct {
	UpdateMachineTagsRequestBody `source:"Body" vadidate:"required,dive"`
}

// UpdateMachineTagsRequestBody defines the request body for update machine tags
type UpdateMachineTagsRequestBody struct {
	Name string `json:"Name,omitempty" validate:"required,min=1"`
	Tags Tags   `json:"Tags,omitempty" validate:"omitempty,dive"`
}

// UpdateNodeOnlineRequest defines the request for update node online status
type UpdateNodeOnlineRequest struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster"`
	Name    string `json:"Name" source:"Query,Name" validate:"required,min=1"`
	// Available: Online, Offline
	OnlineStatus string `json:"OnlineStatus" source:"Query,OnlineStatus" validate:"required,min=1,oneof=Online Offline"`
}

// UpdateNodeTagsRequest defines the request for update node tags
type UpdateNodeTagsRequest struct {
	UpdateNodeTagsRequestBody `source:"Body" validate:"required,dive"`
}

// UpdateNodeTagsRequestBody defines the request body for update node tags
type UpdateNodeTagsRequestBody struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster" validate:"required,min=1"`
	Name    string `json:"Name,omitempty" validate:"required,min=1"`
	Tags    Tags   `json:"Tags,omitempty" validate:"omitempty,dive"`
}

// UpdateNodeTaintsRequest defines the request for update node taints
type UpdateNodeTaintsRequest struct {
	UpdateNodeTaintsRequestBody `source:"Body"`
}

// UpdateNodeTaintsRequestBody defines the request body for update node taints
type UpdateNodeTaintsRequestBody struct {
	// Cluster defines the Cluster name which machine belongs to
	Cluster string `json:"Cluster" source:"Query,Cluster"`
	Name    string `json:"Name,omitempty" validate:"required,min=1"`
	Taints  Taints `json:"Taints,omitempty" validate:"omitempty,dive"`
}

// UpdatePVCRequest defines the data structure for PVC PUT API request.
type UpdatePVCRequest struct {
	Cluster   string `json:"Cluster"`
	Namespace string `json:"Namespace"`
	Name      string `json:"Name"`
	Size      string `json:"Size"`
}

// UpdateStorageClassRequest defines the data structure for StorageClass PUT API request.
type UpdateStorageClassRequest struct {
	Cluster     string              `json:"Cluster"`
	Name        string              `json:"Name"`
	Alias       string              `json:"Alias,omitempty"`
	Description string              `json:"Description,omitempty"`
	Quota       StorageQuotaRequest `json:"Quota"`
}

// UpdateStorageServiceRequest defines the data structure for StorageService PUT API request.
type UpdateStorageServiceRequest struct {
	Name        string              `json:"Name"`
	Alias       string              `json:"Alias,omitempty"`
	Description string              `json:"Description,omitempty"`
	Quota       StorageQuotaRequest `json:"Quota"`
}

type ceph struct {
	Monitors    string `json:"Monitors"`
	AdminID     string `json:"AdminId"`
	AdminSecret string `json:"AdminSecret"`
}

type existingGlusterFS struct {
	Path    string   `json:"Path"`
	Servers []string `json:"Servers"`
}

type glusterFS struct {
	RESTURL     string `json:"RestURL"`
	RESTUser    string `json:"RestUser"`
	RESTUserKey string `json:"RestUserKey"`
}

type nfs struct {
	Server     string `json:"Server"`
	ExportPath string `json:"ExportPath"`
}
