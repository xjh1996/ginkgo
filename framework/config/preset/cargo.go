package preset

import (
	"fmt"
	"net/http"

	"github.com/caicloud/zeus/framework/devops"

	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
)

const (
	cargoUserName = "admin"
	cargoPWD      = "Pwd123456"
	cargoName     = "cn-test"
)

type CargoRawInfo struct {
	Enabled bool
	Domain  string
}

type CargoInfo struct {
	// build is true if Cargo is created by framework
	build bool
	ID    string
}

// Cargo checks whether the preset cargo works. if it does not work,
// the raw information is used to integrate a new one.
func Cargo(raw CargoRawInfo, tenant string, cargoAPI cargoclient.Interface) (cargo *CargoInfo, err error) {
	if !raw.Enabled {
		return &CargoInfo{}, fmt.Errorf("Cargo integration is disabled!!")
	}
	if raw.Domain != "" {
		// 集成外部仓库
		presetCargo, err := devops.CreateCargoRegistry(cargoAPI, cargoName, tenant, raw.Domain, cargoUserName, cargoPWD)
		if err != nil {
			return nil, err
		}
		// TODO：check whether the cargo service exists, if it works then return nil
		// 1. 部署cargo服务，使用自签证书
		// 2. 在所有集群的coredns里加上域名解析
		// 3. 所有集群所有节点的本地域名解析
		return &CargoInfo{
			build: false,
			ID:    presetCargo.Name,
		}, nil
	}
	// Specified cargo does not work, then  a new cargo service with cargo information.
	// TODO：build a new cargo service
	return nil, fmt.Errorf("Cargo cannot be integrated!!!! configuration: %v", raw)
}

func (cargo *CargoInfo) Delete(c *http.Client) error {
	if cargo == nil || !cargo.build {
		return nil
	}
	// TODO separate cargo service
	return nil
}
