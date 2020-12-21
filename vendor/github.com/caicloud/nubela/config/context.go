package config

import (
	"flag"
	"time"

	"github.com/onsi/ginkgo/config"
)

const (
	// DefaultNumNodes is the number of nodes. If not specified, then number of nodes is auto-detected
	DefaultNumNodes = -1
	Scheme          = "http"
)

// TestContextType contains test settings and global state. It is a item
// managed by the test framework itself.
//
// However, individual tests have specific settings. The recommendation way for
// those settings is:
// - They are stored in their own context structure or local
//   variables.
// - The standard `flag` package is used to register them.
//   The flag name should follow the pattern <part1>.<part2>....<partn>
//   where the prefix is unlikely to conflict with other tests or
//   standard packages and each part is in lower camel case. For
//   example, test/e2e/storage/csi/context.go could define
//   storage.csi.numIterations.
// - common/config is used to simplify the registration of
//   multiple options with a single function call:
//   var storageCSI {
//       NumIterations `default:"1" usage:"number of iterations"`
//   }
//   _ config.AddOptions(&storageCSI, "storage.csi")
// - The direct use Viper in tests is possible, but discouraged because
//   it only works in test suites which use Viper (which is not
//   required) and the supported options cannot be
//   discovered by a test suite user.
//
// Test suite authors can use framework/viper to make all command line
// parameters also configurable via a configuration file.

type TestContextType struct {
	BaseUrl    string
	Scheme     string
	KubeConfig string

	ViperConfig string

	ReportDir    string
	ReportPrefix string

	// Output directory for interesting/useful test data
	OutputDir string

	// Timeout for waiting for system pods to be running
	SystemPodsStartupTimeout time.Duration

	// SystemdServices are comma separated list of systemd services the test framework
	// will dump logs for.
	SystemdServices string
	// DumpSystemdJournal controls whether to dump the full systemd journal.
	DumpSystemdJournal bool

	DeleteNamespace          bool
	DeleteNamespaceOnFailure bool
	AllowedNotReadyNodes     int
	CleanStart               bool
	// If set to 'true' or 'all' framework will start a goroutine monitoring resource usage of system add-ons.
	// It will read the data every 30 seconds from all Nodes and print summary during afterEach. If set to 'master'
	// only master Node will be monitored.
	GatherKubeSystemResourceUsageData string
	GatherLogsSizes                   bool
	GatherMetricsAfterTest            string
	GatherSuiteMetricsAfterTest       bool
	AllowGatheringProfiles            bool

	// Currently supported values are 'hr' for human-readable and 'json'. It's a comma separated list.
	OutputPrintType string
	// NodeSchedulableTimeout is the timeout for waiting for all nodes to be schedulable.
	NodeSchedulableTimeout time.Duration
	// SystemDaemonsetStartupTimeout is the timeout for waiting for all system daemonsets to be ready.
	SystemDaemonsetStartupTimeout time.Duration

	// If set to true test will dump data about the namespace in which test was running.
	DumpLogsOnFailure bool
	// Disables dumping cluster log from master and nodes after all tests.
	DisableLogDump bool
	// SpecSummaryOutput is the file to write ginkgo.SpecSummary objects to as tests complete. Useful for debugging and test introspection.
	SpecSummaryOutputDir string
}

// TestContext should be used by all tests to access common context data.
var TestContext TestContextType

// RegisterCommonFlags registers flags common to all e2e testsuites
func RegisterCommonFlags(flags *flag.FlagSet) {
	// Turn on verbose by default to get spec names
	config.DefaultReporterConfig.Verbose = true

	// Turn on EmitSpecProgress to get spec progress (especially on interrupt)
	config.GinkgoConfig.EmitSpecProgress = true

	// Randomize specs as well as suites
	config.GinkgoConfig.RandomizeAllSpecs = true

	flags.StringVar(&TestContext.BaseUrl, "baseurl", "", "The address of a test cluster, it always like 192.168.129.30:6060.")
	flags.StringVar(&TestContext.Scheme, "scheme", Scheme, "The communication way with platform. Default is http")

	flags.StringVar(&TestContext.KubeConfig, "kube-config", "~/.kube/config", "Path to the kubeconfig file.")

	flags.StringVar(&TestContext.GatherKubeSystemResourceUsageData, "gather-resource-usage", "false", "If set to 'true' or 'all' framework will be monitoring resource usage of system all add-ons in (some) e2e tests, if set to 'master' framework will be monitoring master node only, if set to 'none' of 'false' monitoring will be turned off.")
	flags.BoolVar(&TestContext.GatherLogsSizes, "gather-logs-sizes", false, "If set to true framework will be monitoring logs sizes on all machines running e2e tests.")
	flags.StringVar(&TestContext.GatherMetricsAfterTest, "gather-metrics-at-teardown", "false", "If set to 'true' framework will gather metrics from all components after each test. If set to 'master' only master component metrics would be gathered.")
	flags.BoolVar(&TestContext.GatherSuiteMetricsAfterTest, "gather-suite-metrics-at-teardown", false, "If set to true framwork will gather metrics from all components after the whole test suite completes.")
	flags.BoolVar(&TestContext.AllowGatheringProfiles, "allow-gathering-profiles", true, "If set to true framework will allow to gather CPU/memory allocation pprof profiles from the master.")
	flags.StringVar(&TestContext.OutputPrintType, "output-print-type", "json", "Format in which summaries should be printed: 'hr' for human readable, 'json' for JSON ones.")
	flags.BoolVar(&TestContext.DumpLogsOnFailure, "dump-logs-on-failure", true, "If set to true test will dump data about the namespace in which test was running.")
	flags.BoolVar(&TestContext.DisableLogDump, "disable-log-dump", false, "If set to true, logs from master and nodes won't be gathered after test run.")
	flags.StringVar(&TestContext.SystemdServices, "systemd-services", "docker", "The comma separated list of systemd services the framework will dump logs for.")
	flags.BoolVar(&TestContext.DumpSystemdJournal, "dump-systemd-journal", false, "Whether to dump the full systemd journal.")
	flags.StringVar(&TestContext.SpecSummaryOutputDir, "spec-dump", "./reports", "The file to dump all ginkgo.SpecSummary to after tests run. If empty, no objects are saved/printed.")

	flags.StringVar(&TestContext.OutputDir, "e2e-output-dir", "/tmp", "Output directory for interesting/useful test data, like performance data and other metrics. Default is ../tmp directory")
	flags.StringVar(&TestContext.ReportPrefix, "report-prefix", "", "Optional prefix for JUnit XML reports. Default is empty, which doesn't prepend anything to the default name.")
	flags.StringVar(&TestContext.ReportDir, "report-dir", "./reports", "Path to the directory where the JUnit XML reports should be saved. Default is ../reports directory.")

	flags.BoolVar(&TestContext.DeleteNamespace, "delete-namespace", true, "If true tests will delete namespace after completion. It is only designed to make debugging easier, DO NOT turn it off by default.")
	flags.BoolVar(&TestContext.DeleteNamespaceOnFailure, "delete-namespace-on-failure", true, "If true, framework will delete test namespace on failure. Used only during test debugging.")
	flags.IntVar(&TestContext.AllowedNotReadyNodes, "allowed-not-ready-nodes", 0, "If non-zero, framework will allow for that many non-ready nodes when checking for all ready nodes.")
	flags.DurationVar(&TestContext.SystemPodsStartupTimeout, "system-pods-startup-timeout", 10*time.Minute, "Timeout for waiting for all system pods to be running before starting tests.")
	flags.DurationVar(&TestContext.NodeSchedulableTimeout, "node-schedulable-timeout", 30*time.Minute, "Timeout for waiting for all nodes to be schedulable.")
	flags.DurationVar(&TestContext.SystemDaemonsetStartupTimeout, "system-daemonsets-startup-timeout", 5*time.Minute, "Timeout for waiting for all system daemonsets to be ready.")
	flags.BoolVar(&TestContext.CleanStart, "clean-start", false, "If true, purge all namespaces except default and system before running tests. This serves to Cleanup test namespaces from failed/interrupted e2e runs in a long-lived cluster.")
}
