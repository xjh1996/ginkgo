package v20201010

import (
	"context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v20201010 client.
type Interface interface {
	// AddTenantMembers description:
	// Add tenant members
	AddTenantMembers(ctx context.Context, addTenantMemberReq_ *AddTenantMemberReq) (nilResp_ *NilResp, err error)
	// AddTenantResource description:
	// Add resource to tenant
	AddTenantResource(ctx context.Context, addTenantResourceReq_ *AddTenantResourceReq) (nilResp_ *NilResp, err error)
	// CreateNamespace description:
	// Create the namespace within cluster, tenant and returns it
	CreateNamespace(ctx context.Context, createNamespace_ *CreateNamespaceRequest) (namespace_ *Namespace, err error)
	// CreateRole description:
	// Create a role
	CreateRole(ctx context.Context, createRoleReq_ *CreateRoleReq) (role_ *Role, err error)
	// CreateTenant description:
	// Create a tenant
	CreateTenant(ctx context.Context, createTenantReq_ *CreateTenantReq) (tenant_ *Tenant, err error)
	// CreateTenantQuota description:
	// Create the tenant quota within cluster and returns it
	CreateTenantQuota(ctx context.Context, req_ *CreateTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error)
	// CreateUsers description:
	// Create multi users
	CreateUsers(ctx context.Context, createUserReq_ *CreateUserReq) (createUserResp_ *CreateUserResp, err error)
	// DeleteNamespace description:
	// Delete the namespace within cluster, tenant and returns it
	DeleteNamespace(ctx context.Context, deleteNamespace_ *DeleteNamespaceRequest) (nilResp_ *NilResp, err error)
	// DeleteRole description:
	// Delete role
	DeleteRole(ctx context.Context, deleteRoleReq_ *DeleteRoleReq) (nilResp_ *NilResp, err error)
	// DeleteTenant description:
	// Delete tenant
	DeleteTenant(ctx context.Context, deleteTenantReq_ *DeleteTenantReq) (deleteTenantResp_ *DeleteTenantResp, err error)
	// DeleteTenantQuota description:
	// delete the tenant quota within cluster and returns it
	DeleteTenantQuota(ctx context.Context, deleteTenantQuotaRequest_ *DeleteTenantQuotaRequest) (nilResp_ *NilResp, err error)
	// DeleteUser description:
	// Delete a user with name
	DeleteUser(ctx context.Context, getUserReq_ *GetUserReq) (nilResp_ *NilResp, err error)
	// GetClusterQuota description:
	// Get the cluster quota within cluster and returns it
	GetClusterQuota(ctx context.Context, getClusterQuota_ *GetClusterQuotaRequest) (clusterQuota_ *ClusterQuota, err error)
	// GetNamespace description:
	// Get the namespace within cluster, tenant and returns it
	GetNamespace(ctx context.Context, getNamespace_ *GetNamespaceRequest) (namespace_ *Namespace, err error)
	// GetRole description:
	// Get a role with name
	GetRole(ctx context.Context, getRoleReq_ *GetRoleReq) (role_ *Role, err error)
	// GetRoleBindState description:
	// Get role bind state
	GetRoleBindState(ctx context.Context, getRoleReq_ *GetRoleReq) (roleBindState_ *RoleBindState, err error)
	// GetTenant description:
	// Get a tenant with name
	GetTenant(ctx context.Context, getTenantReq_ *GetTenantReq) (tenant_ *Tenant, err error)
	// GetTenantQuota description:
	// get the tenant quota within cluster and returns it
	GetTenantQuota(ctx context.Context, getTenantQuotaRequest_ *GetTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error)
	// GetUser description:
	// Get a user with name
	GetUser(ctx context.Context, getUserReq_ *GetUserReq) (userResp_ *UserResp, err error)
	// GetUserPolicy description:
	// Get user policy
	GetUserPolicy(ctx context.Context, getUserPolicyReq_ *GetUserPolicyReq) (getUserPolicyResp_ *GetUserPolicyResp, err error)
	// GetUserProfile description:
	// Get user profile
	GetUserProfile(ctx context.Context) (userResp_ *UserResp, err error)
	// ListNamespace description:
	// List the namespace within cluster, tenant and returns an array
	ListNamespace(ctx context.Context, listNamespace_ *ListNamespaceRequest, labelSelector_ string) (namespaceList_ *NamespaceList, err error)
	// ListOperationPolicies description:
	// List operation policies
	ListOperationPolicies(ctx context.Context, listPolicyOptions_ *ListPolicyOptions) (operationPolicyList_ *OperationPolicyList, err error)
	// ListRoleBindTeams description:
	// List Role bind teams
	ListRoleBindTeams(ctx context.Context, listRoleBindObjectsOptions_ *ListRoleBindObjectsOptions) (listRoleBindObjectsResp_ *ListRoleBindObjectsResp, err error)
	// ListRoleBindUsers description:
	// Role unbind users
	ListRoleBindUsers(ctx context.Context, listRoleBindObjectsOptions_ *ListRoleBindObjectsOptions) (listRoleBindObjectsResp_ *ListRoleBindObjectsResp, err error)
	// ListRoles description:
	// List roles with options
	ListRoles(ctx context.Context, listRoleOptions_ *ListRoleOptions) (roleList_ *RoleList, err error)
	// ListTenantMembers description:
	// List tenant members
	ListTenantMembers(ctx context.Context, listTenantMemberOptions_ *ListTenantMemberOptions) (tenantMemberList_ *TenantMemberList, err error)
	// ListTenantQuota description:
	// List the tenant quota within cluster and returns an array
	ListTenantQuota(ctx context.Context, listTenantQuotaRequest_ *ListTenantQuotaRequest) (tenantQuotaList_ *TenantQuotaList, err error)
	// ListTenants description:
	// List tenants with options
	ListTenants(ctx context.Context, listTenantOptions_ *ListTenantOptions) (tenantList_ *TenantList, err error)
	// ListUserRole description:
	// List user role
	ListUserRole(ctx context.Context) (listUserRoleResp_ *ListUserRoleResp, err error)
	// ListUserTenant description:
	// List user tenants
	ListUserTenant(ctx context.Context) (tenantList_ *TenantList, err error)
	// ListUsers description:
	// List users with options
	ListUsers(ctx context.Context, listUserOptions_ *ListUserOptions) (userList_ *UserList, err error)
	// Login description:
	// User login
	Login(ctx context.Context, userLoginReq_ *UserLoginReq) (userResp_ *UserResp, err error)
	// RemoveTenantMembers description:
	// Remove tenant members
	RemoveTenantMembers(ctx context.Context, removeTenantMemberReq_ *RemoveTenantMemberReq) (nilResp_ *NilResp, err error)
	// RemoveTenantResource description:
	// Remove resource to tenant
	RemoveTenantResource(ctx context.Context, removeTenantResourceReq_ *RemoveTenantResourceReq) (nilResp_ *NilResp, err error)
	// ResetUserPassword description:
	// Reset other user password
	ResetUserPassword(ctx context.Context, resetPasswordReq_ *ResetPasswordReq) (nilResp_ *NilResp, err error)
	// RoleBindTeams description:
	// Role bind teams
	RoleBindTeams(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error)
	// RoleBindUsers description:
	// Role bind users
	RoleBindUsers(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error)
	// RoleUnbindTeams description:
	// Role unbind teams
	RoleUnbindTeams(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error)
	// RoleUnbindUsers description:
	// Role unbind users
	RoleUnbindUsers(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error)
	// UpdateNamespace description:
	// Update the namespace within cluster, tenant and returns it
	UpdateNamespace(ctx context.Context, updateNamespace_ *CreateNamespaceRequest) (namespace_ *Namespace, err error)
	// UpdateRole description:
	// Update role
	UpdateRole(ctx context.Context, updateRoleReq_ *UpdateRoleReq) (nilResp_ *NilResp, err error)
	// UpdateTenant description:
	// Update tenant
	UpdateTenant(ctx context.Context, updateTenantReq_ *UpdateTenantReq) (tenant_ *Tenant, err error)
	// UpdateTenantMember description:
	// Update tenant member to tenant owner or member
	UpdateTenantMember(ctx context.Context, updateTenantMemberRole_ *UpdateTenantMemberRole) (nilResp_ *NilResp, err error)
	// UpdateTenantQuota description:
	// Update the tenant quota within cluster and returns it
	UpdateTenantQuota(ctx context.Context, req_ *CreateTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error)
	// UpdateUser description:
	// Update user
	UpdateUser(ctx context.Context, userReq_ *UserReq) (userResp_ *UserResp, err error)
	// UpdateUserPassword description:
	// Get user password
	UpdateUserPassword(ctx context.Context, updatePasswordReq_ *UpdatePasswordReq) (nilResp_ *NilResp, err error)
	// UpdateUserProfile description:
	// Get user profile
	UpdateUserProfile(ctx context.Context, updateUserProfileReq_ *UpdateUserProfileReq) (userResp_ *UserResp, err error)
}

// Client for version v20201010.
type Client struct {
	rest *rest.Client
}

// NewClient creates a new client.
func NewClient(cfg *rest.Config) (*Client, error) {
	client, err := rest.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// MustNewClient creates a new client or panic if an error occurs.
func MustNewClient(cfg *rest.Config) *Client {
	client, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

// AddTenantMembers description:
// Add tenant members
func (c *Client) AddTenantMembers(ctx context.Context, addTenantMemberReq_ *AddTenantMemberReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=AddTenantMembers").
		Body("application/json", addTenantMemberReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// AddTenantResource description:
// Add resource to tenant
func (c *Client) AddTenantResource(ctx context.Context, addTenantResourceReq_ *AddTenantResourceReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=AddTenantResource").
		Body("application/json", addTenantResourceReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// CreateNamespace description:
// Create the namespace within cluster, tenant and returns it
func (c *Client) CreateNamespace(ctx context.Context, createNamespace_ *CreateNamespaceRequest) (namespace_ *Namespace, err error) {
	namespace_ = new(Namespace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateNamespace").
		Header("X-Tenant", createNamespace_.Tenant).
		Body("application/json", createNamespace_.CreateNamespaceRequestBody).
		TOPRPCData(namespace_).
		Do(ctx)
	return
}

// CreateRole description:
// Create a role
func (c *Client) CreateRole(ctx context.Context, createRoleReq_ *CreateRoleReq) (role_ *Role, err error) {
	role_ = new(Role)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateRole").
		Body("application/json", createRoleReq_).
		TOPRPCData(role_).
		Do(ctx)
	return
}

// CreateTenant description:
// Create a tenant
func (c *Client) CreateTenant(ctx context.Context, createTenantReq_ *CreateTenantReq) (tenant_ *Tenant, err error) {
	tenant_ = new(Tenant)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateTenant").
		Body("application/json", createTenantReq_).
		TOPRPCData(tenant_).
		Do(ctx)
	return
}

// CreateTenantQuota description:
// Create the tenant quota within cluster and returns it
func (c *Client) CreateTenantQuota(ctx context.Context, req_ *CreateTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error) {
	tenantQuota_ = new(TenantQuota)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateTenantQuota").
		Body("application/json", req_).
		TOPRPCData(tenantQuota_).
		Do(ctx)
	return
}

// CreateUsers description:
// Create multi users
func (c *Client) CreateUsers(ctx context.Context, createUserReq_ *CreateUserReq) (createUserResp_ *CreateUserResp, err error) {
	createUserResp_ = new(CreateUserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=CreateUsers").
		Body("application/json", createUserReq_).
		TOPRPCData(createUserResp_).
		Do(ctx)
	return
}

// DeleteNamespace description:
// Delete the namespace within cluster, tenant and returns it
func (c *Client) DeleteNamespace(ctx context.Context, deleteNamespace_ *DeleteNamespaceRequest) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteNamespace").
		Header("X-Tenant", deleteNamespace_.Tenant).
		Query("Cluster", deleteNamespace_.Cluster).
		Query("Namespace", deleteNamespace_.Name).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// DeleteRole description:
// Delete role
func (c *Client) DeleteRole(ctx context.Context, deleteRoleReq_ *DeleteRoleReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteRole").
		Query("Id", deleteRoleReq_.UID).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// DeleteTenant description:
// Delete tenant
func (c *Client) DeleteTenant(ctx context.Context, deleteTenantReq_ *DeleteTenantReq) (deleteTenantResp_ *DeleteTenantResp, err error) {
	deleteTenantResp_ = new(DeleteTenantResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteTenant").
		Query("Id", deleteTenantReq_.UID).
		TOPRPCData(deleteTenantResp_).
		Do(ctx)
	return
}

// DeleteTenantQuota description:
// delete the tenant quota within cluster and returns it
func (c *Client) DeleteTenantQuota(ctx context.Context, deleteTenantQuotaRequest_ *DeleteTenantQuotaRequest) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteTenantQuota").
		Query("Tenant", deleteTenantQuotaRequest_.Tenant).
		Query("Cluster", deleteTenantQuotaRequest_.Cluster).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// DeleteUser description:
// Delete a user with name
func (c *Client) DeleteUser(ctx context.Context, getUserReq_ *GetUserReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=DeleteUser").
		Query("Name", getUserReq_.Name).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// GetClusterQuota description:
// Get the cluster quota within cluster and returns it
func (c *Client) GetClusterQuota(ctx context.Context, getClusterQuota_ *GetClusterQuotaRequest) (clusterQuota_ *ClusterQuota, err error) {
	clusterQuota_ = new(ClusterQuota)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetClusterQuota").
		Query("Name", getClusterQuota_.Name).
		Query("Cluster", getClusterQuota_.Cluster).
		TOPRPCData(clusterQuota_).
		Do(ctx)
	return
}

// GetNamespace description:
// Get the namespace within cluster, tenant and returns it
func (c *Client) GetNamespace(ctx context.Context, getNamespace_ *GetNamespaceRequest) (namespace_ *Namespace, err error) {
	namespace_ = new(Namespace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetNamespace").
		Header("X-Tenant", getNamespace_.Tenant).
		Query("Cluster", getNamespace_.Cluster).
		Query("Namespace", getNamespace_.Name).
		TOPRPCData(namespace_).
		Do(ctx)
	return
}

// GetRole description:
// Get a role with name
func (c *Client) GetRole(ctx context.Context, getRoleReq_ *GetRoleReq) (role_ *Role, err error) {
	role_ = new(Role)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetRole").
		Query("Id", getRoleReq_.UID).
		TOPRPCData(role_).
		Do(ctx)
	return
}

// GetRoleBindState description:
// Get role bind state
func (c *Client) GetRoleBindState(ctx context.Context, getRoleReq_ *GetRoleReq) (roleBindState_ *RoleBindState, err error) {
	roleBindState_ = new(RoleBindState)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetRoleBindState").
		Query("Id", getRoleReq_.UID).
		TOPRPCData(roleBindState_).
		Do(ctx)
	return
}

// GetTenant description:
// Get a tenant with name
func (c *Client) GetTenant(ctx context.Context, getTenantReq_ *GetTenantReq) (tenant_ *Tenant, err error) {
	tenant_ = new(Tenant)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetTenant").
		Query("Id", getTenantReq_.ID).
		Query("Name", getTenantReq_.Name).
		TOPRPCData(tenant_).
		Do(ctx)
	return
}

// GetTenantQuota description:
// get the tenant quota within cluster and returns it
func (c *Client) GetTenantQuota(ctx context.Context, getTenantQuotaRequest_ *GetTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error) {
	tenantQuota_ = new(TenantQuota)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetTenantQuota").
		Query("Tenant", getTenantQuotaRequest_.Tenant).
		Query("Cluster", getTenantQuotaRequest_.Cluster).
		TOPRPCData(tenantQuota_).
		Do(ctx)
	return
}

// GetUser description:
// Get a user with name
func (c *Client) GetUser(ctx context.Context, getUserReq_ *GetUserReq) (userResp_ *UserResp, err error) {
	userResp_ = new(UserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetUser").
		Query("Name", getUserReq_.Name).
		TOPRPCData(userResp_).
		Do(ctx)
	return
}

// GetUserPolicy description:
// Get user policy
func (c *Client) GetUserPolicy(ctx context.Context, getUserPolicyReq_ *GetUserPolicyReq) (getUserPolicyResp_ *GetUserPolicyResp, err error) {
	getUserPolicyResp_ = new(GetUserPolicyResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetUserPolicy").
		Query("Start", getUserPolicyReq_.Start).
		Query("Limit", getUserPolicyReq_.Limit).
		TOPRPCData(getUserPolicyResp_).
		Do(ctx)
	return
}

// GetUserProfile description:
// Get user profile
func (c *Client) GetUserProfile(ctx context.Context) (userResp_ *UserResp, err error) {
	userResp_ = new(UserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=GetUserProfile").
		TOPRPCData(userResp_).
		Do(ctx)
	return
}

// ListNamespace description:
// List the namespace within cluster, tenant and returns an array
func (c *Client) ListNamespace(ctx context.Context, listNamespace_ *ListNamespaceRequest, labelSelector_ string) (namespaceList_ *NamespaceList, err error) {
	namespaceList_ = new(NamespaceList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListNamespace").
		Header("X-Tenant", listNamespace_.Tenant).
		Query("Cluster", listNamespace_.Cluster).
		Query("Category", listNamespace_.Category).
		Query("Start", listNamespace_.Start).
		Query("Limit", listNamespace_.Limit).
		Query("LabelSelector", labelSelector_).
		TOPRPCData(namespaceList_).
		Do(ctx)
	return
}

// ListOperationPolicies description:
// List operation policies
func (c *Client) ListOperationPolicies(ctx context.Context, listPolicyOptions_ *ListPolicyOptions) (operationPolicyList_ *OperationPolicyList, err error) {
	operationPolicyList_ = new(OperationPolicyList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListOperationPolicies").
		Query("Start", listPolicyOptions_.Start).
		Query("Limit", listPolicyOptions_.Limit).
		TOPRPCData(operationPolicyList_).
		Do(ctx)
	return
}

// ListRoleBindTeams description:
// List Role bind teams
func (c *Client) ListRoleBindTeams(ctx context.Context, listRoleBindObjectsOptions_ *ListRoleBindObjectsOptions) (listRoleBindObjectsResp_ *ListRoleBindObjectsResp, err error) {
	listRoleBindObjectsResp_ = new(ListRoleBindObjectsResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRoleBindTeams").
		Query("Start", listRoleBindObjectsOptions_.Start).
		Query("Limit", listRoleBindObjectsOptions_.Limit).
		Query("Sorts", listRoleBindObjectsOptions_.Sorts).
		Query("Id", listRoleBindObjectsOptions_.UID).
		TOPRPCData(listRoleBindObjectsResp_).
		Do(ctx)
	return
}

// ListRoleBindUsers description:
// Role unbind users
func (c *Client) ListRoleBindUsers(ctx context.Context, listRoleBindObjectsOptions_ *ListRoleBindObjectsOptions) (listRoleBindObjectsResp_ *ListRoleBindObjectsResp, err error) {
	listRoleBindObjectsResp_ = new(ListRoleBindObjectsResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRoleBindUsers").
		Query("Start", listRoleBindObjectsOptions_.Start).
		Query("Limit", listRoleBindObjectsOptions_.Limit).
		Query("Sorts", listRoleBindObjectsOptions_.Sorts).
		Query("Id", listRoleBindObjectsOptions_.UID).
		TOPRPCData(listRoleBindObjectsResp_).
		Do(ctx)
	return
}

// ListRoles description:
// List roles with options
func (c *Client) ListRoles(ctx context.Context, listRoleOptions_ *ListRoleOptions) (roleList_ *RoleList, err error) {
	roleList_ = new(RoleList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListRoles").
		Query("Start", listRoleOptions_.Start).
		Query("Limit", listRoleOptions_.Limit).
		Query("Sorts", listRoleOptions_.Sorts).
		Query("Name", listRoleOptions_.Name).
		TOPRPCData(roleList_).
		Do(ctx)
	return
}

// ListTenantMembers description:
// List tenant members
func (c *Client) ListTenantMembers(ctx context.Context, listTenantMemberOptions_ *ListTenantMemberOptions) (tenantMemberList_ *TenantMemberList, err error) {
	tenantMemberList_ = new(TenantMemberList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListTenantMembers").
		Query("Start", listTenantMemberOptions_.Start).
		Query("Limit", listTenantMemberOptions_.Limit).
		Query("Sorts", listTenantMemberOptions_.Sorts).
		Query("TenantFilter", listTenantMemberOptions_.TenantFilter).
		TOPRPCData(tenantMemberList_).
		Do(ctx)
	return
}

// ListTenantQuota description:
// List the tenant quota within cluster and returns an array
func (c *Client) ListTenantQuota(ctx context.Context, listTenantQuotaRequest_ *ListTenantQuotaRequest) (tenantQuotaList_ *TenantQuotaList, err error) {
	tenantQuotaList_ = new(TenantQuotaList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListTenantQuota").
		Query("Tenant", listTenantQuotaRequest_.Tenant).
		Query("Cluster", listTenantQuotaRequest_.Cluster).
		Query("Start", listTenantQuotaRequest_.Start).
		Query("Limit", listTenantQuotaRequest_.Limit).
		TOPRPCData(tenantQuotaList_).
		Do(ctx)
	return
}

// ListTenants description:
// List tenants with options
func (c *Client) ListTenants(ctx context.Context, listTenantOptions_ *ListTenantOptions) (tenantList_ *TenantList, err error) {
	tenantList_ = new(TenantList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListTenants").
		Query("Start", listTenantOptions_.Start).
		Query("Limit", listTenantOptions_.Limit).
		Query("Sorts", listTenantOptions_.Sorts).
		Query("Status", listTenantOptions_.State).
		Query("Name", listTenantOptions_.Name).
		Query("ListAll", listTenantOptions_.ListAll).
		TOPRPCData(tenantList_).
		Do(ctx)
	return
}

// ListUserRole description:
// List user role
func (c *Client) ListUserRole(ctx context.Context) (listUserRoleResp_ *ListUserRoleResp, err error) {
	listUserRoleResp_ = new(ListUserRoleResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListUserRole").
		TOPRPCData(listUserRoleResp_).
		Do(ctx)
	return
}

// ListUserTenant description:
// List user tenants
func (c *Client) ListUserTenant(ctx context.Context) (tenantList_ *TenantList, err error) {
	tenantList_ = new(TenantList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListUserTenant").
		TOPRPCData(tenantList_).
		Do(ctx)
	return
}

// ListUsers description:
// List users with options
func (c *Client) ListUsers(ctx context.Context, listUserOptions_ *ListUserOptions) (userList_ *UserList, err error) {
	userList_ = new(UserList)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ListUsers").
		Query("Start", listUserOptions_.Start).
		Query("Limit", listUserOptions_.Limit).
		Query("Sorts", listUserOptions_.Sorts).
		Query("Fields", listUserOptions_.Fields).
		Query("NotInTenant", listUserOptions_.NotInTenant).
		Query("NotInTeam", listUserOptions_.NotInTeam).
		Query("NotInRole", listUserOptions_.NotInRole).
		Query("State", listUserOptions_.State).
		Query("RoleName", listUserOptions_.RoleName).
		Query("NoRole", listUserOptions_.NoRole).
		Query("Name", listUserOptions_.Name).
		Query("Nick", listUserOptions_.Nick).
		Query("NameOrNick", listUserOptions_.NameOrNick).
		Query("TenantFilter", listUserOptions_.TenantFilter).
		TOPRPCData(userList_).
		Do(ctx)
	return
}

// Login description:
// User login
func (c *Client) Login(ctx context.Context, userLoginReq_ *UserLoginReq) (userResp_ *UserResp, err error) {
	userResp_ = new(UserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=Login").
		Body("application/json", userLoginReq_).
		TOPRPCData(userResp_).
		Do(ctx)
	return
}

// RemoveTenantMembers description:
// Remove tenant members
func (c *Client) RemoveTenantMembers(ctx context.Context, removeTenantMemberReq_ *RemoveTenantMemberReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RemoveTenantMembers").
		Body("application/json", removeTenantMemberReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// RemoveTenantResource description:
// Remove resource to tenant
func (c *Client) RemoveTenantResource(ctx context.Context, removeTenantResourceReq_ *RemoveTenantResourceReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RemoveTenantResource").
		Body("application/json", removeTenantResourceReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// ResetUserPassword description:
// Reset other user password
func (c *Client) ResetUserPassword(ctx context.Context, resetPasswordReq_ *ResetPasswordReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=ResetUserPassword").
		Body("application/json", resetPasswordReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// RoleBindTeams description:
// Role bind teams
func (c *Client) RoleBindTeams(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RoleBindTeams").
		Body("application/json", roleBindReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// RoleBindUsers description:
// Role bind users
func (c *Client) RoleBindUsers(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RoleBindUsers").
		Body("application/json", roleBindReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// RoleUnbindTeams description:
// Role unbind teams
func (c *Client) RoleUnbindTeams(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RoleUnbindTeams").
		Body("application/json", roleBindReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// RoleUnbindUsers description:
// Role unbind users
func (c *Client) RoleUnbindUsers(ctx context.Context, roleBindReq_ *RoleBindReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=RoleUnbindUsers").
		Body("application/json", roleBindReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// UpdateNamespace description:
// Update the namespace within cluster, tenant and returns it
func (c *Client) UpdateNamespace(ctx context.Context, updateNamespace_ *CreateNamespaceRequest) (namespace_ *Namespace, err error) {
	namespace_ = new(Namespace)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateNamespace").
		Header("X-Tenant", updateNamespace_.Tenant).
		Body("application/json", updateNamespace_.CreateNamespaceRequestBody).
		TOPRPCData(namespace_).
		Do(ctx)
	return
}

// UpdateRole description:
// Update role
func (c *Client) UpdateRole(ctx context.Context, updateRoleReq_ *UpdateRoleReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateRole").
		Body("application/json", updateRoleReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// UpdateTenant description:
// Update tenant
func (c *Client) UpdateTenant(ctx context.Context, updateTenantReq_ *UpdateTenantReq) (tenant_ *Tenant, err error) {
	tenant_ = new(Tenant)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateTenant").
		Body("application/json", updateTenantReq_).
		TOPRPCData(tenant_).
		Do(ctx)
	return
}

// UpdateTenantMember description:
// Update tenant member to tenant owner or member
func (c *Client) UpdateTenantMember(ctx context.Context, updateTenantMemberRole_ *UpdateTenantMemberRole) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateTenantMember").
		Body("application/json", updateTenantMemberRole_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// UpdateTenantQuota description:
// Update the tenant quota within cluster and returns it
func (c *Client) UpdateTenantQuota(ctx context.Context, req_ *CreateTenantQuotaRequest) (tenantQuota_ *TenantQuota, err error) {
	tenantQuota_ = new(TenantQuota)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateTenantQuota").
		Body("application/json", req_).
		TOPRPCData(tenantQuota_).
		Do(ctx)
	return
}

// UpdateUser description:
// Update user
func (c *Client) UpdateUser(ctx context.Context, userReq_ *UserReq) (userResp_ *UserResp, err error) {
	userResp_ = new(UserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateUser").
		Body("application/json", userReq_).
		TOPRPCData(userResp_).
		Do(ctx)
	return
}

// UpdateUserPassword description:
// Get user password
func (c *Client) UpdateUserPassword(ctx context.Context, updatePasswordReq_ *UpdatePasswordReq) (nilResp_ *NilResp, err error) {
	nilResp_ = new(NilResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateUserPassword").
		Body("application/json", updatePasswordReq_).
		TOPRPCData(nilResp_).
		Do(ctx)
	return
}

// UpdateUserProfile description:
// Get user profile
func (c *Client) UpdateUserProfile(ctx context.Context, updateUserProfileReq_ *UpdateUserProfileReq) (userResp_ *UserResp, err error) {
	userResp_ = new(UserResp)
	err = c.rest.Request("POST", 200, "/?Version=2020-10-10&Action=UpdateUserProfile").
		Body("application/json", updateUserProfileReq_).
		TOPRPCData(userResp_).
		Do(ctx)
	return
}
