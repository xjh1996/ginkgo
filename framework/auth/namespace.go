package auth

import (
	"context"
	"encoding/json"

	authclient "github.com/caicloud/auth/pkg/server/client"
	v20201010 "github.com/caicloud/auth/pkg/server/client/v20201010"
	"k8s.io/apimachinery/pkg/util/wait"
)

// Describe resource metadate for a namespace
type NamespceMetadate struct {
	LimitCPU         string
	LimitMem         string
	RequestCPU       string
	RequestMem       string
	GPU              string
	StorageClassName string
	StorageSize      string
	PVCSize          string
}

// DefaultNM returns default namespace metadata
func DefaultNamespaceMeta() *NamespceMetadate {
	return &NamespceMetadate{
		LimitCPU:   defaultLimitCPU,
		LimitMem:   defaultLimitMem,
		RequestCPU: defaultRequestCPU,
		RequestMem: defaultRequestMem,
	}
}

func GenerateNSQuotaString(quota NamespceMetadate) string {
	quotaMap := map[string]string{"limits.cpu": quota.LimitCPU, "limits.memory": quota.LimitMem, "requests.cpu": quota.RequestCPU, "requests.memory": quota.RequestMem,
		"requests.nvidia.com/gpu": quota.GPU, quota.StorageClassName + ".storageclass.storage.k8s.io/requests.storage": quota.StorageSize,
		quota.StorageClassName + ".storageclass.storage.k8s.io/persistentvolumeclaims": quota.PVCSize}
	quotaByte, err := json.Marshal(quotaMap)
	if err != nil {
		panic(err)
	}
	return string(quotaByte)
}

func CreateNamespaceAndWait(authAPI authclient.Interface, tenantID, name, quota, clusterID string) (*v20201010.Namespace, error) {
	createNSReq := &v20201010.CreateNamespaceRequest{
		Tenant: tenantID,
		CreateNamespaceRequestBody: v20201010.CreateNamespaceRequestBody{
			Name:    "cluster/" + clusterID + "/" + name, // Name format is cluster/{cid}/namespaceName
			Quota:   quota,
			Cluster: clusterID,
		},
	}
	ns, err := authAPI.V20201010().CreateNamespace(context.TODO(), createNSReq)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		ns, err = GetNamespace(authAPI, tenantID, clusterID, name)
		if err != nil {
			return false, err
		}
		if ns.Phase == "Active" {
			return true, nil
		} else {
			return false, nil
		}
	})
	return ns, err
}

func DeleteNamespace(authAPI authclient.Interface, tenantID, clusterName, name string) error {
	delNSReq := &v20201010.DeleteNamespaceRequest{
		Tenant:  tenantID,
		Cluster: clusterName,
		Name:    "cluster/" + clusterName + "/" + name, // cluster/cluster-id/namespace-name格式
	}
	_, err := authAPI.V20201010().DeleteNamespace(context.TODO(), delNSReq)
	return err
}

func UpdateNamespaceAndWait(authAPI authclient.Interface, tenantID, name, quota, clusterID string) (*v20201010.Namespace, error) {
	createNSReq := &v20201010.CreateNamespaceRequest{
		Tenant: tenantID,
		CreateNamespaceRequestBody: v20201010.CreateNamespaceRequestBody{
			Name:    "cluster/" + clusterID + "/" + name, // cluster/cluster-id/namespace-name格式
			Quota:   quota,
			Cluster: clusterID,
		},
	}
	ns, err := authAPI.V20201010().UpdateNamespace(context.TODO(), createNSReq)
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		ns, err = GetNamespace(authAPI, tenantID, clusterID, name)
		if err != nil {
			return false, err
		}
		if ns.Quota == quota {
			return true, nil
		} else {
			return false, nil
		}
	})
	return ns, err
}

func GetNamespace(authAPI authclient.Interface, tenantID, clusterName, name string) (*v20201010.Namespace, error) {
	getNSReq := &v20201010.GetNamespaceRequest{
		Tenant:  tenantID,
		Cluster: clusterName,
		Name:    "cluster/" + clusterName + "/" + name, // cluster/cluster-id/namespace-name格式
	}
	return authAPI.V20201010().GetNamespace(context.TODO(), getNSReq)
}

func ListNamespace(authAPI authclient.Interface, tenantID, clusterName string) (*v20201010.NamespaceList, error) {
	listNSReq := &v20201010.ListNamespaceRequest{
		Tenant:  tenantID,
		Cluster: clusterName,
	}
	return authAPI.V20201010().ListNamespace(context.TODO(), listNSReq, "")
}

func CreateTenantQuota(authAPI authclient.Interface, clusterName, name, quota string) (*v20201010.TenantQuota, error) {
	createQuotaReq := &v20201010.CreateTenantQuotaRequest{
		Cluster: clusterName,
		Name:    name,
		Quota:   quota,
	}
	return authAPI.V20201010().CreateTenantQuota(context.TODO(), createQuotaReq)
}

func DeleteTenantQuota(authAPI authclient.Interface, tenantName, cluster string) error {
	deleteTenantQuotaReq := &v20201010.DeleteTenantQuotaRequest{
		Tenant:  tenantName,
		Cluster: cluster,
	}
	_, err := authAPI.V20201010().DeleteTenantQuota(context.TODO(), deleteTenantQuotaReq)
	return err
}
