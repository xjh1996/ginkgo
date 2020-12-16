package client

import (
	"fmt"

	cosclientset "github.com/caicloud/containeros/pkg/client/clientset"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var BaseClient *BaseClientType

type BaseClientType struct {
	kubeconfig *rest.Config
	K8S        *clientset.Clientset
	COSCRD     *cosclientset.Clientset
}

// LoadClientFromConfig returns basic baseclient and crd clients for the kubernetes cluster.
func LoadClientsetFromConfig(kubeConfig string) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return fmt.Errorf("error building configs with kubeconfig: %v", err.Error())
	}

	c := &BaseClientType{kubeconfig: config}

	c.K8S, err = clientset.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error load k8s baseclient: %v", err.Error())
	}
	c.COSCRD, err = cosclientset.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error load cos crd baseclient: %v", err.Error())
	}
	BaseClient = c
	return nil
}
