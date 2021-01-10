package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineInfoKind is the name of the MachineInfo resource kind.
	MachineInfoKind = "MachineInfo"

	// MachineInfoName is the name of the MachineInfo resource (plural).
	MachineInfoName = "machineinfoes"

	// MachineInfoKindKey is used as the key when mapping to the MachineInfo resource.
	MachineInfoKindKey = "machineinfo"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,path=machineinfoes
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Machine",type="string",JSONPath=".spec.machineName",description="The Name of Machine object"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status

// MachineInfo is the Schema for the machineinfos API
type MachineInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of MachineInfo
	Spec MachineInfoSpec `json:"spec"`

	// Status defines the observed state of MachineInfo
	// +optional
	Status MachineInfoStatus `json:"status,omitempty"`
}

// MachineInfoSpec defines the desired state of MachineInfo
type MachineInfoSpec struct {
	// MachineName is the name of the Machine this object belongs to.
	MachineName string `json:"machineName,omitempty"`
}

// MachineInfoStatus defines the observed state of MachineInfo
type MachineInfoStatus struct {
	// Conditions represents the observations of machineinfo's current state.
	// Known .status.conditions.type are:
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	// System is the information about OS
	// +optional
	System *MachineSystemInfo `json:"system,omitempty"`

	// Hardware is the information about cpu、memory and etc
	// +optional
	Hardware *MachineHardwareInfo `json:"hardware,omitempty"`

	// Disk is the information about block device
	// +optional
	Disk []MachineDiskInfo `json:"disk,omitempty"`

	// Nic is the information about network interface
	// +optional
	Nic []MachineNicInfo `json:"nic,omitempty"`

	// GPU is the information about gpu device
	// +optional
	GPU []MachineGPUInfo `json:"gpu,omitempty"`
}

const (
	// MachineInfoConditionReady defines the machine info sync result
	MachineInfoConditionReady = "Ready"
)

// MachineSystemInfo is the information about OS
type MachineSystemInfo struct {
	// BootTime is system bootup time from Epoch (in Second)
	BootTime string `json:"bootTime"`

	// Hostname is the hostname of machine
	Hostname string `json:"hostname"`

	// OS is the name such as linux
	OS string `json:"os"`

	// Platform is the distribute version such as centos
	Platform string `json:"platform"`

	// PlatformFamily is the distribute version family such as rhel
	PlatformFamily string `json:"platformFamily"`

	// PlatformVersion is the distribute version number such as 7.7.1908
	PlatformVersion string `json:"platformVersion"`

	// KernelVersion is the number such as 3.10.0-1062.el7.x86_64
	KernelVersion string `json:"kernelVersion"`
}

// MachineHardwareInfo is the information about cpu、memory and etc
type MachineHardwareInfo struct {
	// CPUModel example: Intel Core Processor (Haswell, no TSX)
	CPUModel string `json:"cpuModel"`

	// CPUArch example: x86_64
	CPUArch string `json:"cpuArch"`

	// CPUMhz example: 2200
	CPUMHz string `json:"cpuMHz"`

	// CPUCores is the number of logical cores
	CPUCores int `json:"cpuCores"`

	// CPUPhysicalCores is the number of phsical cores
	CPUPhysicalCores int `json:"cpuPhysicalCores"`

	// MemoryTotal is the memory size (in Byte)
	MemoryTotal uint64 `json:"memoryTotal"`
}

// MachineDiskInfo is the information about block device
type MachineDiskInfo struct {
	// Device is the path to device
	Device string `json:"device"`

	// Capacity is the size of device (in Byte)
	Capacity uint64 `json:"capacity"`

	// Type is the block device type such as HDD/SDD
	Type string `json:"type"`

	// DeviceType is the block device type such as lvm
	DeviceType string `json:"deviceType"`

	// MountPoint is the path where the device mounted
	MountPoint string `json:"mountPoint"`
}

// MachineNicInfo is the information about network interface
type MachineNicInfo struct {
	// Name is the network interface name
	Name string `json:"name"`

	MTU          string   `json:"mtu"`
	Speed        string   `json:"speed"`
	HardwareAddr string   `json:"hardwareAddr"`
	Status       string   `json:"status"`
	Addrs        []string `json:"addrs"`
}

// MachineGPUInfo is the information about gpu device
type MachineGPUInfo struct {
	UUID string `json:"uuid"`

	// ProductName is the name of product such as GeForce GT 1030
	ProductName string `json:"productName"`

	// ProductBrand is the name of product such as GeForce
	ProductBrand     string `json:"productBrand"`
	PCIeGen          string `json:"pcieGen"`
	PCILinkWidths    string `json:"pciLinkWidths"`
	MemoryTotal      string `json:"memoryTotal"`
	MemoryClock      string `json:"memoryClock"`
	GraphicsAppClock string `json:"graphicsAppClock"`
	GraphicsMaxClock string `json:"graphicsMaxClock"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineInfoList is a list of MachineInfo resources
type MachineInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is list of MachineInfos.
	Items []MachineInfo `json:"items"`
}
