/*
Copyright 2020 bytedance authors. All rights reserved.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// SecretRef is a reference to a secret that exists in the same namespace.
type SecretRef struct {
	// SecretName is the name of the secret.
	SecretName string `json:"secretName,omitempty"`
}

// ConfigSource references configuration settings.
type ConfigSource struct {
	SecretRef `json:",inline"`
}

// +kubebuilder:validation:Enum=OK;Pending;Error

// State describes the deployment status of the desired resource.
type State string

const (
	// StateOK means that the resource has been successfully deployed without errors.
	StateOK State = "OK"
	// StatePending means that the resource cannot be deployed due to some missing
	// prerequisites.
	StatePending State = "Pending"
	// StateError means that the resource cannot be deployed due to an error.
	StateError State = "Error"
)

// ClusterStatus describes the deployment status of the desired resources on a particular
// cluster.
type ClusterStatus struct {
	// Cluster is the name of the Cluster resource corresponding to the cluster.
	Cluster string `json:"cluster,omitempty"`
	// State describes the current state of the user resource on the cluster.
	State State `json:"state,omitempty"`
	// Message gives a full explanation on why the user resource is in its current
	// state.
	Message string `json:"message,omitempty"`
}

// DeploymentSpec provides customisation options (podTemplate and replicas) for the Pods.
type DeploymentSpec struct {
	PodTemplate corev1.PodTemplateSpec `json:"podTemplate,omitempty"`
	Replicas    *int32                 `json:"replicas,omitempty"`
}
