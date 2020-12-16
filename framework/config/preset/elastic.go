package preset

import (
	"errors"
	"fmt"
	"net/http"
)

type ElasticRawInfo struct {
	Enabled bool
	Address string
}

type ElasticInfo struct {
	// build is true if resource is created by framework
	build bool
	Addr  string
}

// Elastic checks whether the elastic is existed. If it does not exist, build a elastic service for all insight-e2e cases.
func Elastic(raw ElasticRawInfo, c *http.Client) (elastic *ElasticInfo, err error) {
	if !raw.Enabled {
		return &ElasticInfo{}, errors.New(fmt.Sprintf("Elastic configuration is disabled!!"))
	}
	if raw.Address != "" {
		// TODO：check whether elastic works, if it is true then return nil
		return &ElasticInfo{
			build: false,
			Addr:  raw.Address,
		}, nil
	}
	// Specified elastic does not work, then build a elastic service.
	// TODO：build a new elastic
	return nil, errors.New(fmt.Sprintf("Elastic cannot be integrated!!!! configuration: %v", raw))
}

func (elastic *ElasticInfo) Delete(c *http.Client) error {
	if elastic == nil || !elastic.build {
		return nil
	}
	// TODO separate elastic service
	return nil
}
