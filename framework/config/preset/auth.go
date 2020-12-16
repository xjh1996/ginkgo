package preset

import (
	"net/http"

	"k8s.io/apimachinery/pkg/util/rand"
)

type Interface interface {
	Delete(c *http.Client) (err error)
}

type AuthRawInfo struct {
	TenantID      string
	User          string
	Password      string
	AdminTenantID string
	AdminUser     string
	AdminPassword string
}

type AuthInfo struct {
	// build is true if auth is created by framework
	build         bool
	TenantID      string
	User          string
	Password      string
	AdminTenantID string
	AdminUser     string
	AdminPassword string
}

// Auth checks whether the preset auth is existed and works. If preset auth cannot work, then a random
// user and tenant is create for all e2e test case.
// Return build (if resource is created by framework, it is true) and auth information
func Auth(raw AuthRawInfo, c *http.Client) (auth *AuthInfo, err error) {
	if len(raw.Password) > 0 && len(raw.TenantID) > 0 && len(raw.User) > 0 {
		// TODO: check whether the auth is existed, if it is true then return nil
		// 检查用户是否可用
		// 租户是否存在
		// 是否有足够资源
		return &AuthInfo{
			build:         false,
			TenantID:      raw.TenantID,
			User:          raw.User,
			Password:      raw.User,
			AdminTenantID: raw.AdminTenantID,
			AdminUser:     raw.AdminUser,
			AdminPassword: raw.AdminPassword,
		}, nil
	}
	// Specified resource does not work, then create a new tenant and user.
	tenant := rand.String(10)
	user := rand.String(10)
	password := "12345678"
	// 1. 创建租户
	// 2. 分配资源
	// 3. 创建用户并绑定租户
	return &AuthInfo{
		build:         true,
		TenantID:      tenant,
		User:          user,
		Password:      password,
		AdminTenantID: raw.AdminTenantID,
		AdminUser:     raw.AdminUser,
		AdminPassword: raw.AdminPassword,
	}, nil
}

func (auth *AuthInfo) Delete(c *http.Client) error {
	if !auth.build {
		return nil
	}
	// TODO delete preset auth resource
	return nil
}
