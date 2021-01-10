/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// DefaultParametersVersion is the version of Parameters Object
	DefaultParametersVersion = runtime.APIVersionInternal
	// DefaultParametersGroup is the group of Parameters Object
	DefaultParametersGroup = ""
	// IPVSDRParametersKind is the name of the IPVSDRParameters internal kind.
	IPVSDRParametersKind = "IPVSDRParameters"
	// GeneralParametersKind is the name of the GeneralParameters internal kind.
	GeneralParametersKind = "GeneralParameters"
	// IngressNginxParametersKind is the name of the IngressNginxParameters internal kind.
	IngressNginxParametersKind = "IngressNginxParameters"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IPVSDRParameters is an ipvs direct-routing provider.
type IPVSDRParameters struct {
	metav1.TypeMeta     `json:",inline"`
	KeepalivedProviders []KeepalivedProvider `json:"keepalivedProviders,omitempty"`
}

// KeepalivedProvider is a keepalived provider.
type KeepalivedProvider struct {
	// Virtual IP Addresses.
	VIPs []string `json:"vips,omitempty"`
	// virtual server shceduler algorithm type.
	Scheduler IPVSScheduler `json:"scheduler"`
	// ActiveActive or ActivePassive.
	HAMode HAMode `json:"haMode,omitempty"`
	// vip bound to.
	Bind *KeepalivedBind `json:"bind,omitempty"`
}

// HAMode represents the mode of High-Availability.
type HAMode string

const (
	// ActiveActiveHA ...
	ActiveActiveHA HAMode = "ActiveActive"
	// ActivePassiveHA ...
	ActivePassiveHA HAMode = "ActivePassive"
)

// IPVSScheduler is ipvs shceduler algorithm type.
type IPVSScheduler string

const (
	// IPVSSchedulerRR - Round Robin
	IPVSSchedulerRR IPVSScheduler = "rr"
	// IPVSSchedulerWRR - Weighted Round Robin
	IPVSSchedulerWRR IPVSScheduler = "wrr"
	// IPVSSchedulerLC - Round Robin
	IPVSSchedulerLC IPVSScheduler = "lc"
	// IPVSSchedulerWLC - Weighted Least Connections
	IPVSSchedulerWLC IPVSScheduler = "wlc"
	// IPVSSchedulerLBLC - Locality-Based Least Connections
	IPVSSchedulerLBLC IPVSScheduler = "lblc"
	// IPVSSchedulerDH - Destination Hashing
	IPVSSchedulerDH IPVSScheduler = "dh"
	// IPVSSchedulerSH - Source Hashing
	IPVSSchedulerSH IPVSScheduler = "sh"
)

// KeepalivedBind is vip binding information.
type KeepalivedBind struct {
	// bind to interface
	Iface string `json:"iface,omitempty"`
	// bind to interface which in subnet
	// CIDR string `json:"cidr,omitempty"`
	// bind to ip from node annotation
	NodeIPAnnotation string `json:"nodeIPAnnotation,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GeneralParameters is a provider interface for cloud.
type GeneralParameters struct {
	metav1.TypeMeta `json:",inline"`
	Data            map[string]string `json:"data,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressNginxParameters defines parameters object for ingress-nginx controller.
type IngressNginxParameters struct {
	metav1.TypeMeta `json:",inline"`
	Nginx           map[string]string `json:"nginx,omitempty"`
	Args            map[string]string `json:"args,omitempty"`
}
