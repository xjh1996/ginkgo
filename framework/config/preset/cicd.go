package preset

import (
	"fmt"
	"net/http"
)

type CICDRawInfo struct {
	Enabled bool
}

type CICDInfo struct {
	// build is true if CICD is created by framework
	build     bool
	Namespace string
}

// CICD checks whether the specified namespace for CICD is existed. If the namespace does not exist,
// then build a CICD namespace for all CICD-e2e cases.
func CICD(raw CICDRawInfo, c *http.Client) (cicd *CICDInfo, err error) {
	if !raw.Enabled {
		return &CICDInfo{}, fmt.Errorf("CICD configuration is disabled!!")
	}
	// TODO：check whether the cicd namespace exists, if it is true then return namespace, nil
	//return &CICDInfo{
	//	build:     false,
	//	namespace: "",
	//}, nil

	// CICD namespace does not exist, then build one.
	// TODO：build a new CICD namespace
	return nil, fmt.Errorf("Namespace of CICD cannot be created!!!!")
}

func (cicd *CICDInfo) Delete(c *http.Client) error {
	if cicd == nil || !cicd.build {
		return nil
	}
	// TODO separate cicd service
	return nil
}
