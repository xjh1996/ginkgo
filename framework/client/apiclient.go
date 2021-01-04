package client

import (
	"net/http"

	appclient "github.com/caicloud/app/pkg/server/client"
	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
	"github.com/caicloud/nirvana/rest"
	"github.com/caicloud/nubela/baseclient"
	pipelineclient "github.com/caicloud/pipeline/pkg/server/client"
	"github.com/caicloud/zeus/framework/config"
)

// User is used to login.
type User struct {
	Tenant   string
	Username string
	Password string
}

// App retrieves the appClient
func (u *User) App() (appclient.Interface, error) {
	return appclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Pipeline retrieves the pipelineClient
func (u *User) Pipeline() (pipelineclient.Interface, error) {
	return pipelineclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl,
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Cargo retrieves the cargoClient
func (u *User) Cargo() (cargoclient.Interface, error) {
	return cargoclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl,
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Auth retrieves the authClient
func (u *User) Auth() (*http.Client, error) {
	return nil, nil
}

// NewAPIClient return a rest client with specified user
func NewAPIClient(tenant, username, password string) User {
	return User{
		Tenant:   tenant,
		Username: username,
		Password: password,
	}
}
