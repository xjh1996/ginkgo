package compass

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	commonconfig "github.com/caicloud/nubela/config"
	_ "github.com/caicloud/zeus/app"
	_ "github.com/caicloud/zeus/auth"
	_ "github.com/caicloud/zeus/demo"
	_ "github.com/caicloud/zeus/devops"
	"github.com/caicloud/zeus/framework/config"
	_ "github.com/caicloud/zeus/net"
	_ "github.com/caicloud/zeus/resource"
)

var viperConfig = flag.String("config", "e2e-config.yaml", "The name of a viper config file (https://github.com/spf13/viper#what-is-viper). All e2e command line parameters can also be configured in such a file. May contain a path and may or may not contain the file suffix. The default is to look for an optional file with `e2e` as base name. If a file is specified explicitly, it must be present.")

// handleFlags sets up all flags and parses the command line.
func handleFlags() {
	commonconfig.CopyFlags(commonconfig.Flags, flag.CommandLine)
	config.RegisterFlags(flag.CommandLine)
	flag.Parse()
}

func TestMain(m *testing.M) {
	var versionFlag bool
	flag.CommandLine.BoolVar(&versionFlag, "version", false, "Displays version information.")

	// Resister test flags, then parse flags
	handleFlags()

	if err := commonconfig.ViperizeFlags(*viperConfig, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	RunCPSE2ETests(t)
}
