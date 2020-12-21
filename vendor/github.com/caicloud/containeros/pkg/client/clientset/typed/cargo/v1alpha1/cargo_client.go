/*
Copyright 2020 bytedance authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/containeros/pkg/apis/cargo/v1alpha1"
	"github.com/caicloud/containeros/pkg/client/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type CargoV1alpha1Interface interface {
	RESTClient() rest.Interface
	CargosGetter
}

// CargoV1alpha1Client is used to interact with features provided by the cargo.containeros.dev group.
type CargoV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CargoV1alpha1Client) Cargos() CargoInterface {
	return newCargos(c)
}

// NewForConfig creates a new CargoV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CargoV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CargoV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CargoV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CargoV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CargoV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CargoV1alpha1Client {
	return &CargoV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CargoV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
