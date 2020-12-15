package framework

import (
	"fmt"

	. "github.com/caicloud/nubela/image"
)

// registryList holds image registries, typically is the default cargo and one outside cargo.
type registryList struct {
	DefaultLibraryRegistry string
	DefaultReleaseRegistry string
	E2eRegistry            string
}

// InitReg
func InitReg(defaultReg, e2eReg string) {
	registry = registryList{
		DefaultLibraryRegistry: defaultReg + "/library",
		DefaultReleaseRegistry: defaultReg + "/release",
		E2eRegistry:            e2eReg}
}

var (
	registry               registryList
	defaultLibraryRegistry = registry.DefaultLibraryRegistry
	defaultReleaseRegistry = registry.DefaultReleaseRegistry
	e2eRegistry            = registry.E2eRegistry
	imageConfigs           = initImageConfigs()
)

const (
	// Agnhost image
	Nginx = iota
	// BUsyBox image
	BusyBox
	// Apline image
	Apline
)

func initImageConfigs() map[int]Config {
	configs := map[int]Config{}
	configs[Nginx] = Config{defaultLibraryRegistry, "nginx", "1.12.2"}
	configs[BusyBox] = Config{defaultLibraryRegistry, "busybox", "1.30.0"}
	configs[Apline] = Config{defaultLibraryRegistry, "alpine", "3.6"}
	return configs
}

// GetConfig returns the Config object for an image
func GetConfig(image int) Config {
	return imageConfigs[image]
}

// GetE2EImage returns the fully qualified URI to an image (including version)
func GetE2EImage(image int) string {
	return fmt.Sprintf("%s/%s:%s", imageConfigs[image].Registry, imageConfigs[image].Name, imageConfigs[image].Version)
}
