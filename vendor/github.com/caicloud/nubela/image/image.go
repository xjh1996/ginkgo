package image

import (
	"fmt"
)

// Config holds an images registry, name, and version
type Config struct {
	Registry string
	Name     string
	Version  string
}

// SetRegistry sets an image registry in a Config struct
func (i *Config) SetRegistry(registry string) {
	i.Registry = registry
}

// SetName sets an image name in a Config struct
func (i *Config) SetName(name string) {
	i.Name = name
}

// SetVersion sets an image version in a Config struct
func (i *Config) SetVersion(version string) {
	i.Version = version
}

// GetE2EImage returns the fully qualified URI to an image (including version)
func (i *Config) GetE2EImage() string {
	return fmt.Sprintf("%s/%s:%s", i.Registry, i.Name, i.Version)
}
