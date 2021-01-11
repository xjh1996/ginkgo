package config

import (
	"net/http"

	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework/config/preset"
)

type PresetCompassResource struct {
	Auth    *preset.AuthInfo
	Storage *preset.StorageInfo
	Cargo   *preset.CargoInfo
	CICD    *preset.CICDInfo
	Elastic *preset.ElasticInfo
}

// SetupPresetActions are actions that preset compass resource for all-e2e cases.
func SetupCompassPreset(raw PresetResourceRaw) error {
	var err error
	p := PresetCompassResource{}
	p.Auth, err = preset.Auth(raw.AuthRawInfo, &http.Client{})
	if err != nil {
		return err
	}
	logger.Infof("Preset Auth: tenant(%v) user(%v)", p.Auth.TenantID, p.Auth.User)

	user := p.Auth.AdminTenantID
	passwd := p.Auth.AdminPassword
	tenant := p.Auth.AdminTenantID
	adminClient := NewAPIClient(tenant, user, passwd)

	p.Storage, err = preset.Storage(raw.StorageRawInfo, &http.Client{})
	if err != nil {
		logger.Warningf("%v", err)
	} else {
		logger.Infof("Preset Storage: storageclass(%v)", p.Storage.StorageClass)
	}

	cargoAPI, err := adminClient.Cargo()
	if err != nil {
		logger.Warningf("%v", err)
	}
	p.Cargo, err = preset.Cargo(raw.CargoRawInfo, p.Auth.AdminTenantID, cargoAPI)
	if err != nil {
		logger.Warningf("%v", err)
	} else {
		logger.Infof("Preset Cargo: %v", p.Cargo.ID)
	}

	p.CICD, err = preset.CICD(raw.CICDRawInfo, &http.Client{})
	if err != nil {
		logger.Warningf("%v", err)
	} else {
		logger.Infof("Preset CICD: %v", p.CICD.Namespace)
	}

	p.Elastic, err = preset.Elastic(raw.ElasticRawInfo, &http.Client{})
	if err != nil {
		logger.Warningf("%v", err)
	} else {
		logger.Infof("Preset Elastic: %v", p.Elastic.Addr)
	}
	Context.PresetCompassResource = p
	return nil
}

// CleanPresetActions are actions that clean compass preset resource.
func (p *PresetCompassResource) CleanPresetActions() {
	if err := p.Elastic.Delete(&http.Client{}); err != nil {
		logger.Warningf("%v", err)
	}

	if err := p.CICD.Delete(&http.Client{}); err != nil {
		logger.Warningf("%v", err)
	}

	if err := p.Cargo.Delete(&http.Client{}); err != nil {
		logger.Warningf("%v", err)
	}

	if err := p.Storage.Delete(&http.Client{}); err != nil {
		logger.Warningf("%v", err)
	}

	if err := p.Auth.Delete(&http.Client{}); err != nil {
		logger.Warningf("%v", err)
	}
}
