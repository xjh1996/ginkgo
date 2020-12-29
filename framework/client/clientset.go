package client

import (
	"fmt"
	"strings"

	cosclientset "github.com/caicloud/containeros/pkg/client/clientset"
	"github.com/caicloud/nubela/logger"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var BaseClient *BaseClientType
var ControlClient *BaseClientType
var UserClients []BaseClientType
var err error

type BaseClientType struct {
	kubeconfig *rest.Config
	K8S        *clientset.Clientset
	COSCRD     *cosclientset.Clientset
}

// LoadClientFromConfig returns basic baseclient and crd clients for the kubernetes cluster.
func LoadClientsetFromConfig(kubeConfig, controlConfig, userConfigs string) error {
	if kubeConfig == "" {
		logger.Infof("kubeconfig file is not set")
	}
	if BaseClient, err = loadClient(kubeConfig); err != nil {
		return err
	}

	if controlConfig == "" {
		logger.Infof("controlConfig file is not set")
	}
	if ControlClient, err = loadClient(controlConfig); err != nil {
		return err
	}

	if userConfigs == "" {
		logger.Infof("userConfigs file is not set")
	}
	// load multi userClusterClient, config files are separated by comma.
	configs := strings.Split(userConfigs, ", ")
	for _, c := range configs {
		if client, err := loadClient(c); err != nil {
			return err
		} else {
			UserClients = append(UserClients, *client)
		}
	}
	return nil
}

func loadClient(kubeConfig string) (*BaseClientType, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("error building configs with kubeconfig: %v", err.Error())
	}

	c := &BaseClientType{kubeconfig: config}

	c.K8S, err = clientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error load k8s baseclient: %v", err.Error())
	}
	c.COSCRD, err = cosclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error load cos crd baseclient: %v", err.Error())
	}
	return c, nil
}
