package client

import (
	appclient "github.com/caicloud/app/pkg/server/client"
	authclient "github.com/caicloud/auth/pkg/server/client"
	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
	insightclient "github.com/caicloud/insight/pkg/server/client"
	"github.com/caicloud/nirvana/rest"
	"github.com/caicloud/nubela/baseclient"
	pipelineclient "github.com/caicloud/pipeline/pkg/server/client"
	resourceclient "github.com/caicloud/resource/pkg/server/client"
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
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Cargo retrieves the cargoClient
func (u *User) Cargo() (cargoclient.Interface, error) {
	return cargoclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Auth retrieves the authClient
func (u *User) Auth() (authclient.Interface, error) {
	return authclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Resource retrieves the ResourceClient
func (u *User) Resource() (resourceclient.Interface, error) {
	return resourceclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// Insight retrieves the InsightClient
func (u *User) Insight() (insightclient.Interface, error) {
	return insightclient.NewClient(&rest.Config{
		Scheme:   config.Context.Scheme,
		Host:     config.Context.BaseUrl + "/hodor",
		Executor: baseclient.NewRequestExecutorWithAuth(u.Tenant, u.Username, u.Password),
	})
}

// NewAPIClient return a rest client with specified user
func NewAPIClient(tenant, username, password string) User {
	return User{
		Tenant:   tenant,
		Username: username,
		Password: password,
	}
}
