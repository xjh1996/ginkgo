package util

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/caicloud/zeus/framework/config"

	"github.com/caicloud/zeus/framework/auth"
	v1 "k8s.io/api/core/v1"
)

const (
	// Poll is how often to Poll claims.
	Poll = 2 * time.Second
)

// CreateTestingNS should be used by every test, note that we append a common prefix to the provided test name.
func CreateTestingNS(baseName, clusterID, tenantID string, metadate *auth.NamespceMetadate, user config.User) (*v1.Namespace, error) {

	// one namespace of a random name was created.
	//name := fmt.Sprintf("%v-%v", baseName, RandomSuffix())

	// TODO 1. 建一个 namespces 2.等待 namespace active

	return nil, nil
}

// RandomSuffix provides a random sequence to append to pods,services,rcs.
func RandomSuffix() string {
	return strconv.Itoa(rand.Intn(10000))
}
