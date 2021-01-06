package baseclient

import (
	"encoding/base64"
	"net/http"

	"github.com/caicloud/nirvana/rest"

	log "k8s.io/klog/v2"
)

//TODO get Oauth token
func GetToken(username, password string) (string, error) {
	log.Infof("Get token with user: %v; password: %v", username, password)
	return "", nil
}

type requestExecutorWithAuth struct {
	c        *http.Client
	tenant   string
	username string
	password string
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Do uses the http.Client to send an HTTP request and set basic authentication before sending.
func (r *requestExecutorWithAuth) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "basic "+basicAuth(r.username, r.password))
	req.Header.Add("X-Tenant", r.tenant)
	return r.c.Do(req)
}

// NewRequestExecutorWithAuth implements a custom nirvana client.RequestExecutor to set authentication information on each request.
// Usage:
//   xxx.NewClient(&client.Config{
//	     Scheme:   "...",
//	     Host:     "...",
//	     Executor: NewRequestExecutorWithAuth("system-tenant", "admin", "Pwd123456"),
//   })
func NewRequestExecutorWithAuth(tenant, username, password string) rest.RequestExecutor {
	return &requestExecutorWithAuth{
		c:        &http.Client{},
		tenant:   tenant,
		username: username,
		password: password,
	}
}
