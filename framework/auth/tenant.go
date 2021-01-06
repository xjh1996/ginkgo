package auth

import (
	"context"

	authclient "github.com/caicloud/auth/pkg/server/client"
	v20201010 "github.com/caicloud/auth/pkg/server/client/v20201010"
	"github.com/caicloud/nubela/logger"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"
)

func CreateTenantAndWait(authAPI authclient.Interface, name, description string) error {
	tenantReq := &v20201010.CreateTenantReq{
		Name:        name,
		Description: description,
	}
	tenant, err := authAPI.V20201010().CreateTenant(context.TODO(), tenantReq)
	if err != nil {
		return err
	}
	return wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		tenant, err = GetTenant(authAPI, tenant.Name, tenant.UID)
		if err != nil {
			return false, err
		}
		if tenant.State == "Active" { // FIXME: this field value need verify.
			return true, nil
		} else {
			logger.Infof("tenant is not ready, retrying...")
			return false, nil
		}
	})
}

func DeleteTenantAndWait(authAPI authclient.Interface, id string) error {
	delTenantReq := &v20201010.DeleteTenantReq{UID: id} // XXX: why only support id
	_, err := authAPI.V20201010().DeleteTenant(context.TODO(), delTenantReq)
	if err != nil {
		return err
	}
	return wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		_, err = GetTenant(authAPI, "", id)
		if err != nil {
			if apierrors.IsNotFound(err) { // FIXME: this error check may not useful
				return true, nil
			} else {
				return false, err
			}
		}
		return false, nil
	})
}
func GetTenant(authAPI authclient.Interface, name, id string) (*v20201010.Tenant, error) {
	getTenantReq := &v20201010.GetTenantReq{
		Name: name, // Name and ID only need supply one
		ID:   id,
	}
	return authAPI.V20201010().GetTenant(context.TODO(), getTenantReq)
}
