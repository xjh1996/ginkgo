package framework

import (
	"fmt"
	"strings"
	"time"

	"github.com/caicloud/zeus/framework/auth"

	"github.com/caicloud/nubela/expect"

	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework/client"
	e2econfig "github.com/caicloud/zeus/framework/config"
	"github.com/caicloud/zeus/framework/util"
	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
)

// Framework supports common operations used by e2e tests;
// it will keep business clients & a namespace & kubernetes clients for you.
// Eventual goal is to merge this with integration test framework.
type Framework struct {
	BaseName string

	// Guaranteed to be unique in the cluster even when running the same
	// test multiple times in parallel.
	UniqueName string

	// Set the Clientset for kubernetes
	skipK8sClientsetCreation bool                   // Whether to skip creationg a k8s clientset
	ClientSet                *client.BaseClientType // return backend clientset

	// cluster info
	RestClient      client.User
	AdminRestClient client.User
	ClusterID       string
	PresetResource  e2econfig.PresetCompassResource

	skipNamespaceCreation    bool            // Whether to skip creating a namespace
	Namespace                *v1.Namespace   // Every test has at least one namespace unless creation is skipped
	namespacesToDelete       []*v1.Namespace // Some tests have more than one.
	NamespaceDeletionTimeout time.Duration
	NamespceMetadate         *auth.NamespceMetadate
}

// NewDefaultFramework makes a new framework.
func NewDefaultFramework(baseName string) *Framework {
	return NewFramework(baseName, false, false)
}

// NewFramework creates a test framework and sets up a BeforeEach/AfterEach for
// you (you can write additional before/after each functions).
func NewFramework(baseName string, skipK8sClientsetCreation, skipNamespaceCreation bool) *Framework {
	f := &Framework{
		BaseName:                 baseName,
		skipK8sClientsetCreation: skipK8sClientsetCreation,
		skipNamespaceCreation:    skipNamespaceCreation,
	}
	ginkgo.BeforeEach(f.BeforeEach)
	ginkgo.AfterEach(f.AfterEach)
	return f
}

// BeforeEach sets clients makes a namespace.
func (f *Framework) BeforeEach() {
	f.ClusterID = e2econfig.Context.ClusterID
	f.PresetResource = e2econfig.Context.PresetCompassResource
	f.RestClient = client.NewAPIClient(f.PresetResource.Auth.User, f.PresetResource.Auth.Password)
	f.AdminRestClient = client.NewAPIClient(f.PresetResource.Auth.AdminUser, f.PresetResource.Auth.Password)
	f.ClientSet = client.BaseClient

	if !f.skipNamespaceCreation {
		ginkgo.By(fmt.Sprintf("Building a namespace, basename %s", f.BaseName))
		if f.NamespceMetadate == nil {
			f.NamespceMetadate = auth.DefaultNamespaceMeta()
		}
		namespace, err := f.CreateNamespace(f.NamespceMetadate)
		expect.NoError(err)

		f.Namespace = namespace
		//f.UniqueName = f.Namespace.GetName()
	}
}

// CreateNamespace creates a namespace for e2e testing.
func (f *Framework) CreateNamespace(metadate *auth.NamespceMetadate) (*v1.Namespace, error) {
	ns, err := util.CreateTestingNS(f.BaseName, f.ClusterID, f.PresetResource.Auth.TenantID, metadate, f.RestClient)
	f.AddNamespacesToDelete(ns)
	return ns, err
}

// AddNamespacesToDelete adds one or more namespaces to be deleted when the test completes.
func (f *Framework) AddNamespacesToDelete(namespaces ...*v1.Namespace) {
	for _, ns := range namespaces {
		if ns == nil {
			continue
		}
		f.namespacesToDelete = append(f.namespacesToDelete, ns)
	}
}

// AfterEach deletes the namespace, after reading its events.
func (f *Framework) AfterEach() {

	// DeleteNamespace at the very end in defer, to avoid any
	// expectation failures preventing deleting the namespace.
	defer func() {
		nsDeletionErrors := map[string]error{}
		// Whether to delete namespace is determined by 3 factors: delete-namespace flag, delete-namespace-on-failure flag and the test result
		// if delete-namespace set to false, namespace will always be preserved.
		// if delete-namespace is true and delete-namespace-on-failure is false, namespace will be preserved if test failed.
		if e2econfig.Context.DeleteNamespace && (e2econfig.Context.DeleteNamespaceOnFailure || !ginkgo.CurrentGinkgoTestDescription().Failed) {
			for _, ns := range f.namespacesToDelete {
				ginkgo.By(fmt.Sprintf("Destroying namespace %q for this suite.", ns.Name))
				// TODO 删除分区
				if _, err := f.RestClient.Auth(); err != nil {
					//if !apierrors.IsNotFound(err) {
					//	nsDeletionErrors[ns.Name] = err
					//} else {
					//	logger.Infof("Namespace %v was already deleted", ns.Name)
					//}
				}
			}
		} else {
			if !e2econfig.Context.DeleteNamespace {
				logger.Warningf("Found DeleteNamespace=false, skipping namespace deletion!")
			} else {
				logger.Warningf("Found DeleteNamespaceOnFailure=false and current test failed, skipping namespace deletion!")
			}
		}

		// Paranoia-- prevent reuse!
		f.Namespace = nil
		f.ClientSet = nil
		f.namespacesToDelete = nil

		// if we had errors deleting, report them now.
		if len(nsDeletionErrors) != 0 {
			messages := []string{}
			for namespaceKey, namespaceErr := range nsDeletionErrors {
				messages = append(messages, fmt.Sprintf("Couldn't delete ns: %q: %s (%#v)", namespaceKey, namespaceErr, namespaceErr))
			}
			logger.Failf(strings.Join(messages, ","))
		}
	}()

	// Do something when the test ended before clean namespaces.
}
