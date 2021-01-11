package config

import (
	"flag"

	commonconfig "github.com/caicloud/nubela/config"
	"github.com/caicloud/zeus/framework/config/preset"
)

const Scheme = "http"

type ContextType struct {
	*commonconfig.TestContextType

	// preset default resource, PresetResourceRaw is from configuration file.
	ClusterID             string
	PresetResourceRaw     PresetResourceRaw
	PresetCompassResource PresetCompassResource
}

type PresetResourceRaw struct {
	AuthRawInfo    preset.AuthRawInfo
	StorageRawInfo preset.StorageRawInfo
	CargoRawInfo   preset.CargoRawInfo
	CICDRawInfo    preset.CICDRawInfo
	ElasticRawInfo preset.ElasticRawInfo
}

// Context should be used by all tests to access common context data.
var Context ContextType

// RegisterClusterFlags registers flags specific to the cluster e2e test suite.
func RegisterFlags(flags *flag.FlagSet) {
	commonconfig.RegisterCommonFlags(flag.CommandLine)
	Context = ContextType{
		TestContextType: &commonconfig.TestContext,
	}

	flags.StringVar(&Context.ClusterID, "cluster", "compass-stack", "ID of the test cluster. Default is compass-stack that is the control cluster.")

	resourceInfo := &Context.PresetResourceRaw
	flags.StringVar(&resourceInfo.AuthRawInfo.TenantID, "tenant", "", "The tenant to run all e2e cases.")
	flags.StringVar(&resourceInfo.AuthRawInfo.User, "user", "", "The user to run all e2e cases.")
	flags.StringVar(&resourceInfo.AuthRawInfo.Password, "password", "", "The password of the user.")
	flags.StringVar(&resourceInfo.AuthRawInfo.AdminTenantID, "admin-tenant", "sys-tenant", "The admin tenant to run all e2e cases.")
	flags.StringVar(&resourceInfo.AuthRawInfo.AdminUser, "admin-user", "admin", "The admin user to run all e2e cases.")
	flags.StringVar(&resourceInfo.AuthRawInfo.AdminPassword, "admin-password", "Pwd123456", "The password of the admin user.")

	flags.BoolVar(&resourceInfo.StorageRawInfo.Enabled, "storage-enable", false, "If set true, must input a existed storageclass or storage-info.")
	flags.StringVar(&resourceInfo.StorageRawInfo.StorageClass, "storageclass", "", "Existed storageclass in the cluster.")
	flags.StringVar(&resourceInfo.StorageRawInfo.RawInfo, "storage-info", "", "Storage info that used to be integrated.")

	flags.BoolVar(&resourceInfo.CargoRawInfo.Enabled, "cargo-enable", false, "If set true, must input existed cargo-host or cargo-info.")
	flags.StringVar(&resourceInfo.CargoRawInfo.Domain, "cargo-host", "", "Host of the existed Cargo in the cluster.")

	flags.BoolVar(&resourceInfo.CICDRawInfo.Enabled, "cicd-enable", false, "If set true, the namespace of CICD will be created.")

	flags.BoolVar(&resourceInfo.ElasticRawInfo.Enabled, "elastic-enable", false, "If set true, must input elastic info.")
	flags.StringVar(&resourceInfo.ElasticRawInfo.Address, "elastic-address", "", "Specified elastic will be integrated.")
}
