package compass

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/caicloud/nubela/logger"
	e2ereporter "github.com/caicloud/nubela/reporter"
	"github.com/caicloud/zeus/framework/client"
	e2econfig "github.com/caicloud/zeus/framework/config"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
)

var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	setupSuite()
	return nil
}, func(data []byte) {
	// Run on all Ginkgo nodes
	setupSuitePerGinkgoNode()
})

var _ = ginkgo.SynchronizedAfterSuite(func() {
	cleanupSuite()
}, func() {
	afterSuiteActions()
})

// RunCPSE2ETests checks configuration parameters (specified through flags ) and then runs
// E2E tests using the Ginkgo runner.
// TODO:
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory, and cluster logs will also be saved.
// This function is called on each Ginkgo node in parallel mode.
func RunCPSE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	gomega.RegisterFailHandler(ginkgo.Fail)

	// Run tests through the Ginkgo runner with output to console + JUnit for Jenkins
	var r []ginkgo.Reporter
	if e2econfig.Context.ReportDir != "" {
		if err := os.MkdirAll(e2econfig.Context.ReportDir, 0755); err != nil {
			klog.Errorf("Failed creating report directory: %v", err)
		} else {
			r = append(r, reporters.NewJUnitReporter(path.Join(e2econfig.Context.ReportDir, fmt.Sprintf("junit_%v%02d.xml", e2econfig.Context.ReportPrefix, config.GinkgoConfig.ParallelNode))))
		}

	}

	// The DetailsRepoerter will output details about every test (name, files, lines, etc) which helps
	// when documenting our tests.
	if e2econfig.Context.SpecSummaryOutputDir != "" {
		if err := os.MkdirAll(e2econfig.Context.SpecSummaryOutputDir, 0755); err != nil {
			klog.Errorf("Failed creating report directory: %v", err)
		} else {
			r = append(r, e2ereporter.NewDetailsReporterFile(path.Join(e2econfig.Context.SpecSummaryOutputDir, fmt.Sprintf("spec_details.log"))))
		}
	}

	klog.Infof("Starting e2e run on Ginkgo node %d", config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "Kubernetes e2e suite", r)
}

// setupSuite is the boilerplate that can be used to setup ginkgo test suites, on the SynchronizedBeforeSuite step.
// There are certain operations we only want to run once per overall test invocation
// (such as deleting old namespaces, or verifying that all system pods are running.
// Because of the way Ginkgo runs tests in parallel, we must use SynchronizedBeforeSuite
// to ensure that these operations only run on the first parallel Ginkgo node.
//
// This function takes two parameters: one function which runs on only the first Ginkgo node,
// returning an opaque byte array, and then a second function which runs on all Ginkgo nodes,
// accepting the byte array.
func setupSuite() {
	// Run only on Ginkgo node 1
	klog.Infof("Running setupSuite actions on node 1")

	err := client.LoadClientsetFromConfig(e2econfig.Context.KubeConfig, e2econfig.Context.ControlClusterConfig, e2econfig.Context.UserClusterConfigs)
	if err != nil {
		logger.Failf("Error loading Clientset: %v", err)
	}

	dc := client.BaseClient.K8S.DiscoveryClient

	serverVersion, serverErr := dc.ServerVersion()
	if serverErr != nil {
		logger.Failf("Unexpected server error retrieving version: %v", serverErr)
	}
	if serverVersion != nil {
		klog.Infof("kube-apiserver version: %s", serverVersion.GitVersion)
	}

	// preset compass resource
	err = e2econfig.SetupCompassPreset(e2econfig.Context.PresetResourceRaw)
	if err != nil {
		logger.Failf("Error preset resources: %v", err)
	}
}

// setupSuitePerGinkgoNode is the boilerplate that can be used to setup ginkgo test suites, on the SynchronizedBeforeSuite step.
// There are certain operations we only want to run once per overall test invocation on each Ginkgo node
// such as making some global variables accessible to all parallel executions
// Because of the way Ginkgo runs tests in parallel, we must use SynchronizedBeforeSuite
// Ref: https://onsi.github.io/ginkgo/#parallel-specs
func setupSuitePerGinkgoNode() {
	klog.Info("Running setupSuite actions on all nodes")
}

// CleanupSuite is the boilerplate that can be used after tests on ginkgo were run, on the SynchronizedAfterSuite step.
// Similar to SynchronizedBeforeSuite, we want to run some operations only once (such as collecting cluster logs).
// Here, the order of functions is reversed; first, the function which runs everywhere,
// and then the function that only runs on the first Ginkgo node.
func cleanupSuite() {
	// Run on all Ginkgo nodes
	klog.Info("Running afterSuite actions on all nodes")
}

// AfterSuiteActions are actions that are run on ginkgo's SynchronizedAfterSuite
func afterSuiteActions() {
	// Run only Ginkgo on node 1
	klog.Info("Running afterSuite actions on node 1")

	e2econfig.Context.PresetCompassResource.CleanPresetActions()
}
