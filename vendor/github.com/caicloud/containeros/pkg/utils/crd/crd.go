package crd

import "fmt"

// GenerateCRDChartFilename generates the filename of the CRD chart for the given resource.
func GenerateCRDChartFilename(group, resource string) string {
	return fmt.Sprintf("%s_%s.yaml", group, resource)
}
