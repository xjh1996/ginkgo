package baseclient

import (
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
	username string
	password string
}

// Do uses the http.Client to send an HTTP request and set basic authentication before sending.
func (r *requestExecutorWithAuth) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(r.username, r.password)
	return r.c.Do(req)
}

// NewRequestExecutorWithAuth implements a custom nirvana client.RequestExecutor to set authentication information on each request.
// Usage:
//   xxx.NewClient(&client.Config{
//	     Scheme:   "...",
//	     Host:     "...",
//	     Executor: NewRequestExecutorWithAuth("admin", "Pwd123456"),
//   })
func NewRequestExecutorWithAuth(username, password string) rest.RequestExecutor {
	return &requestExecutorWithAuth{
		c:        &http.Client{},
		username: username,
		password: password,
	}
}
