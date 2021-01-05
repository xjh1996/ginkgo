package preset

import (
	"fmt"
	"net/http"
)

type StorageRawInfo struct {
	Enabled      bool
	StorageClass string
	RawInfo      string
}

type StorageInfo struct {
	// build is true if storage is created by framework
	build          bool
	StorageService string
	StorageClass   string
}

// Storage uses one of existed resource and the raw information of resource to service as an external storage.
// Existed resource is checked first, if it can not work then the raw storage information is used to build a new one.
func Storage(raw StorageRawInfo, c *http.Client) (storage *StorageInfo, err error) {
	if !raw.Enabled {
		return &StorageInfo{}, fmt.Errorf("Storage integration is disabled!!")
	}
	if raw.StorageClass != "" {
		// TODO：check whether the storage class exists, if it works then return nil
		return &StorageInfo{
			build:        false,
			StorageClass: raw.StorageClass,
		}, nil
	}
	// Specified resource does not exist, then build a new storage service with storage service information.
	if raw.RawInfo != "" {
		// TODO：
		// 1. 注册 storage service
		// 2. 创建 stroage class
		// 3. 分配给租户
	}
	return &StorageInfo{}, fmt.Errorf("Storage with info %v cannot be created!!!!", raw)
}

func (storage *StorageInfo) Delete(c *http.Client) error {
	if storage == nil || !storage.build {
		return nil
	}
	// TODO Delete storage service
	return nil
}
