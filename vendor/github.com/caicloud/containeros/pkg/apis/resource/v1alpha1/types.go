/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
)

// CloudProvider represents the infrastructure provider name.
type CloudProvider = string

// NodeAutoscalerNotifySetting defines the notify information when scale
type NodeAutoscalerNotifySetting struct {
	// Methods is the notify methods, example: email
	Methods []string `json:"methods"`

	// Groups is the receiver group id
	Groups []string `json:"groups"`
}

const (
	// ConditionReady defines the condition type Ready with ref object
	ConditionReady = "Ready"
)

// StorageQuota represents the quota for StorageClass object.
type StorageQuota struct {
	// PersistentVolumeClaims represents the number of PVC that can be allocated, 0 means no limit.
	PersistentVolumeClaims resource.Quantity `json:"persistentvolumeclaims"`

	// RequestsStorage represents the total storage capacity that can be allocated, must be greater than 0.
	RequestsStorage resource.Quantity `json:"requests.storage"`
}
