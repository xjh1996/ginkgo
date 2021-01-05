package v20201010

import (
	v1 "github.com/caicloud/api/meta/v1"
	time "time"
)

// AddTenantMemberReq add member to tenant
type AddTenantMemberReq struct {
	// TenantID when operation tenant is systemTenant must take it
	TenantID string `json:"TenantId"`
	// Items add tenant members
	Items []OperateTenantMember `json:"Items" validate:"dive"`
}

// AddTenantResourceReq  use for tenant start cluster
type AddTenantResourceReq struct {
	TenantID      string         `json:"Tenant" validate:"required,min=1"`
	ResourceKind  string         `json:"FeatureType" validate:"required,min=1"`
	ResourceValue []ResourceMeta `json:"Datas" validate:"dive"`
}

// ChildMeta child resource define, at now compass only two level resource, if have more, need add children field
type ChildMeta struct {
	Kind  string `json:"Key" validate:"required,min=1"`
	Value string `json:"Value" validate:"required,min=1"`
}

// ClusterQuota defines the response for cluster quota
type ClusterQuota struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	Ratio         string               `json:"Ratio"`
	Logical       ClusterQuotaLogical  `json:"Logical"`
	Physical      ClusterQuotaPhysical `json:"Physical"`
}

// nolint
// ClusterQuotaLogical defines the logical quota for cluster quota
type ClusterQuotaLogical struct {
	Total      string `json:"Total"`
	Allocated  string `json:"Allocated"`
	SystemUsed string `json:"SystemUsed"`
	Used       string `json:"Used"`
}

// nolint
// ClusterQuotaPhysical defines the physical quota for cluster quota
type ClusterQuotaPhysical struct {
	Allocatable   string `json:"Allocatable"`
	Capacity      string `json:"Capacity"`
	Unavailable   string `json:"Unavailable"`
	Unschedulable string `json:"Unschedulable"`
}

// CreateNamespaceRequest defines the request for create namespace
type CreateNamespaceRequest struct {
	Tenant                     string `source:"Header,X-Tenant" json:"Tenant" validate:"required"`
	CreateNamespaceRequestBody `source:"Body"`
}

// CreateNamespaceRequestBody defines the request body for create namespace
type CreateNamespaceRequestBody struct {
	// Name format is cluster/{cid}/namespaceName, here will not do auth， for unified rule
	Name    string `json:"Name" validate:"required"`
	Quota   string `json:"Quota" validate:"required"`
	Cluster string `json:"Cluster" validate:"required"`
}

// CreateRoleReq create role req
type CreateRoleReq struct {
	Name        string   `json:"Name" validate:"required"`
	Description string   `json:"Description,omitempty"`
	PolicyNames []string `json:"PolicyNames"`
	Resource    []string `json:"Resource,omitempty"`
}

// CreateTenantQuotaRequest defines the request for create tenant quota
type CreateTenantQuotaRequest struct {
	Cluster string `json:"Cluster,omitempty" validate:"required"`
	Name    string `json:"Name,omitempty" validate:"required"`
	Quota   string `json:"Quota,omitempty" validate:"required"`
}

// CreateTenantReq create tenant request
type CreateTenantReq struct {
	// Name of the tenant
	Name string `json:"Name" validate:"required"`
	// description tenant description
	Description string `json:"Description" validate:"min=0,max=256"`
}

// CreateUserReq create users request
type CreateUserReq struct {
	// Items create user slice
	Items []UserReq `json:"Items" validate:"dive"`
}

// CreateUserResp create user response
type CreateUserResp struct {
}

// DeleteNamespaceRequest defines the request for delete namespace
type DeleteNamespaceRequest struct {
	Tenant  string `source:"Header,X-Tenant" json:"Tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
	// Name format is cluster/{cid}/namespaceName, this name need auth
	Name string `source:"Query,Namespace" validate:"required"`
}

// DeleteRoleReq delete role req
type DeleteRoleReq struct {
	// UID role id, query key Id
	UID string `source:"query,Id" validate:"required"`
}

// DeleteTenantQuotaRequest defines the request for delete tenant quota
type DeleteTenantQuotaRequest struct {
	Tenant  string `source:"Query,Tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
}

// DeleteTenantReq delete tenant Request
type DeleteTenantReq struct {
	// UID of the tenant
	UID string `source:"query,Id" validate:"required"`
}

// DeleteTenantResp delete tenant request
type DeleteTenantResp struct {
}

// GetClusterQuotaRequest defines the request for get cluster quota
type GetClusterQuotaRequest struct {
	Name    string `json:"Name" source:"Query,Name"`
	Cluster string `json:"Cluster" source:"Query,Cluster"`
}

// GetNamespaceRequest defines the request for get namespace
type GetNamespaceRequest struct {
	Tenant  string `source:"Header,X-Tenant" json:"Tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
	// Name format is cluster/{cid}/namespaceName, this name need auth
	Name string `source:"Query,Namespace" validate:"required"`
}

// GetRoleReq get role request
type GetRoleReq struct {
	// UID role id, query key Id
	UID string `source:"query,Id" validate:"required"`
}

// GetTenantQuotaRequest defines the request for get tenant quota
type GetTenantQuotaRequest struct {
	Tenant  string `source:"Query,Tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
}

// GetTenantReq get tenant request
type GetTenantReq struct {
	// ID of the tenant
	ID string `source:"query,Id" json:"Id,omitempty"`
	// Name of tenant
	Name string `source:"query,Name" json:"Name,omitempty"`
}

// GetUserPolicyReq get use all policy for authZ
type GetUserPolicyReq struct {
	Paginator
}

// GetUserPolicyResp return user deny and allow policy
type GetUserPolicyResp struct {
	Deny  []StatementMeta `json:"Deny"`
	Allow []StatementMeta `json:"Allow"`
}

// GetUserReq get  user request
type GetUserReq struct {
	Name string `source:"query,Name" json:"Name" validate:"required"`
}

// ListMeta list meta info
type ListMeta struct {
	Total int `json:"Total"`
}

// ListNamespaceRequest defines the request for list namespace
type ListNamespaceRequest struct {
	Tenant  string `source:"Header,X-Tenant" json:"Tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
	// Category defines namespace category: system/non-system/all
	Category string `source:"Query,Category,default=all"`
	Start    int    `source:"Query,Start,default=0"`
	Limit    int    `source:"Query,Limit,default=10"`
}

// ListPolicyOptions list policy Req
type ListPolicyOptions struct {
	Paginator
}

// ListRoleBindObjectsOptions list role bind user or teams
type ListRoleBindObjectsOptions struct {
	Paginator
	// Sorts object sort by fields, exp: [id, -createTime]
	// +optional
	Sorts []string `source:"query,Sorts" json:",omitempty"`
	// UID role id
	UID string `source:"query,Id" validate:"required"`
}

// ListRoleBindObjectsResp list role bind user or team
type ListRoleBindObjectsResp struct {
	ListMeta
	Items []RoleBindObject `json:"Items"`
}

// ListRoleOptions list role options
type ListRoleOptions struct {
	Paginator
	// Sorts object sort by fields, exp: [id, -createTime]
	// +optional
	Sorts []string `source:"query,Sorts" json:",omitempty"`
	// Name search with role name
	Name string `source:"query,Name" json:",omitempty"`
}

// ListTenantMemberOptions list tenant members
type ListTenantMemberOptions struct {
	Paginator
	// Sorts object sort by fields, exp: [id, -createTime]
	// +optional
	Sorts []string `source:"query,Sorts" json:"Sorts,omitempty"`
	// TenantFilter for system-tenant list normal tenant members
	// +optional
	TenantFilter string `source:"query,TenantFilter" json:",omitempty"`
}

// ListTenantOptions list tenant options
type ListTenantOptions struct {
	Paginator
	// Sorts object sort by fields, exp: [id, -createTime]
	// +optional
	Sorts []string `source:"query,Sorts" json:"Sorts,omitempty"`
	// State search with state, web field is Status
	// +optional
	State string `source:"query,Status" json:"State,omitempty"`
	// Name regex tenant name
	// +optional
	Name string `source:"query,Name" json:"Name,omitempty"`
	// ListAll return all tenant
	// +optional
	ListAll string `source:"query,ListAll" json:"ListAll,omitempty"`
}

// ListTenantQuotaRequest defines the request for list tenant quota
type ListTenantQuotaRequest struct {
	Tenant  string `source:"Query,Tenant" json:"tenant" validate:"required"`
	Cluster string `source:"Query,Cluster" validate:"required"`
	Start   int    `source:"Query,Start,default=0"`
	Limit   int    `source:"Query,Limit,default=10"`
}

// ListUserOptions list users options
type ListUserOptions struct {
	Paginator
	// Sorts object sort by fields, exp: [id, -createTime]
	// +optional
	Sorts []string `source:"query,Sorts" json:",omitempty"`
	// Fields if fields not nil, only return fields value
	// +optional
	Fields []string `source:"query,Fields" json:",omitempty"`
	// NotInTenant use for add tenant member
	// +optional
	NotInTenant string `source:"query,NotInTenant" json:",omitempty"`
	// NotInTeam use for add team member
	// +optional
	NotInTeam string `source:"query,NotInTeam" json:",omitempty"`
	// NotInRole use for role bind
	// +optional
	NotInRole string `source:"query,NotInRole" json:",omitempty"`
	// State search with state
	// +optional
	State string `source:"query,State,omitempty" json:",omitempty"`
	// Role search with role name
	// +optional
	RoleName string `source:"query,RoleName" json:",omitempty"`
	// NoRole search user no roles
	NoRole bool `source:"query,NoRole" json:",omitempty"`
	// Name search with user name
	// +optional
	Name string `source:"query,Name" json:",omitempty"`
	// Nick search with user nick
	// +optional
	Nick string `source:"query,Nick" json:",omitempty"`
	// NameOrNick search with user name or nick
	// +optional
	NameOrNick string `source:"query,NameOrNick" json:",omitempty"`
	// TenantFilter for system-tenant list normal tenant users
	// +optional
	TenantFilter string `source:"query,TenantFilter" json:",omitempty"`
}

// ListUserRoleResp return user role and role from, role bind or team bind
type ListUserRoleResp struct {
	ListMeta
	Items []UserRole `json:"Items"`
}

// Meta object comm info
type Meta struct {
	// UID object id
	UID string `json:"Id"`
	// CreateTime object create time
	CreateTime time.Time `json:"CreateTime"`
	// DeleteTime object delete time
	DeleteTime *time.Time `json:"DeleteTime,omitempty"`
}

// Namespace defines the response for namespace
type Namespace struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	Phase         string            `json:"Phase"`
	Conditions    []StatusCondition `json:"Conditions"`
	Quota         string            `json:"Quota"`
	Used          string            `json:"Used"`
}

// nolint
// NamespaceList defines the response for namespace list
type NamespaceList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []Namespace `json:"Items"`
}

// NilResp nothing to return
type NilResp struct {
}

// OperateTenantMember operation tenant member object
type OperateTenantMember struct {
	// UserName member name
	UserName string `json:"Name" validate:"required"`
	// RoleNames member role name
	RoleNames []string `json:"Role"`
}

// OperationPolicyList operation policy name list
type OperationPolicyList struct {
	ListMeta
	Items []string `json:"Items"`
}

// Paginator is a list option for page
type Paginator struct {
	Start int `json:"Start" source:"Query,Start"`
	Limit int `json:"Limit" source:"Query,Limit"`
}

// RemoveTenantMemberReq remove tenant member
type RemoveTenantMemberReq struct {
	// TenantID when operation tenant is systemTenant must take it
	TenantID string `json:"TenantId"`
	// UserName delete member name
	UserName string `json:"Name"`
}

// RemoveTenantResourceReq use for tenant close cluster
type RemoveTenantResourceReq struct {
	TenantID      string `json:"Tenant" validate:"required,min=1"`
	ResourceKind  string `json:"FeatureType" validate:"required,min=1"`
	ResourceValue string `json:"FeatureValue" validate:"required,min=1"`
}

// ResetPasswordReq manager reset other user's password
type ResetPasswordReq struct {
	UserName string `json:"Username" validate:"required,min=1"`
	Password string `json:"Password" validate:"required,min=3"`
}

// ResourceMeta resource define, include resource and sub resource
type ResourceMeta struct {
	Kind     string      `json:"FeatureType"`
	Value    string      `json:"FeatureValue" validate:"required,min=1"`
	Children []ChildMeta `json:"Modules" validate:"dive"`
}

// Role role info
type Role struct {
	Meta        `json:",inline"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	PolicyNames []string `json:"PolicyNames"`
	Resource    []string `json:"Resource"`
	Parent      string   `json:"Parent"`
}

// RoleBindObject role bind value item
type RoleBindObject struct {
	UID      string    `json:"Id"`
	Name     string    `json:"Name"`
	BindTime time.Time `json:"BindTime"`
}

// RoleBindReq role bind user or team request body
type RoleBindReq struct {
	UID   string   `json:"Id" validate:"required"`
	Items []string `json:"Items" validate:"required"`
}

// RoleBindState if the role have been bind
type RoleBindState struct {
	IsBinded bool `json:"IsBinded"`
}

// RoleList role list
type RoleList struct {
	ListMeta
	// Items list user slice
	Items []Role `json:"Items"`
}

// StatementMeta policy statement define, include Effect, Action, and Resource
type StatementMeta struct {
	// Effect value: Allow, Deny
	Effect string `json:"Effect" bson:"Effect" validate:"required"`
	// Action policy action
	Action []string `json:"Action" bson:"Action" validate:"required"`
	// Resource policy action
	Resource []string `json:"Resource,omitempty" bson:"Resource,omitempty"`
}

// StatusCondition namespace status condition
type StatusCondition struct {
	Type    string `json:"Type"`
	Status  string `json:"Status"`
	Reason  string `json:"Reason,omitempty"`
	Message string `json:"Message,omitempty"`
}

// Tenant web tenant info
type Tenant struct {
	Meta `json:",inline"`
	// Name of the tenant
	Name string `json:"Name"`
	// State of tenant, web field is Status
	State string `json:"Status"`
	// description tenant description
	Description string `json:"Description"`
	// OwnerCount owner number
	OwnerCount int64 `json:"OwnerCount"`
	// MemberCount member number
	MemberCount int64 `json:"MemberCount"`
}

// TenantList wet tenant list info
type TenantList struct {
	ListMeta
	// Items tenant list
	Items []Tenant `json:"Items"`
}

// TenantMember tenant member
type TenantMember struct {
	Meta `json:",inline"`
	// UserName tenant member name
	UserName string `json:"Name"`
	// Nick tenant member nick
	UserNick string `json:"Nickname"`
	// RoleNames tenant member names
	RoleNames []string `json:"Role,omitempty"`
	// Email member email
	UserEmail string `json:"Email"`
}

// TenantMemberList list tenant member
type TenantMemberList struct {
	ListMeta
	// Items tenant member list
	Items []TenantMember `json:"Items"`
}

// TenantQuota defines the response for tenant quota
type TenantQuota struct {
	v1.ObjectMeta `json:",inline" bson:",inline"`
	// Quota defines the quota allocated to tenant
	Quota string `json:"Hard"`
	// Phase defines the tenant quota phase, corev1.ResourceList
	Phase          string `json:"Phase"`
	Allocated      string `json:"Allocated"`
	Used           string `json:"Used"`
	NamespaceCount int    `json:"NamespaceCount"`
}

// nolint
// TenantQuotaList defines the response for tenant quota list
type TenantQuotaList struct {
	v1.ListMeta `json:",inline" bson:"inline"`
	Items       []TenantQuota `json:"Items"`
}

// UpdatePasswordReq user update self password
type UpdatePasswordReq struct {
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword" validate:"required,min=3"`
}

// UpdateRoleReq udpat role Req
type UpdateRoleReq struct {
	UID         string   `json:"Id" validate:"required"`
	Name        string   `json:"Name" validate:"required"`
	Description string   `json:"Description"`
	PolicyNames []string `json:"PolicyNames"`
	Resource    []string `json:"Resource"`
}

// UpdateTenantMemberRole set member to tenant owner or remove owner
type UpdateTenantMemberRole struct {
	TenantID string `json:"TenantId" validate:"required"`
	UserName string `json:"Name" validate:"required"`
	RoleName string `json:"Role"`
}

// UpdateTenantReq update tenant request
type UpdateTenantReq struct {
	// UID tenant id
	UID string `json:"Id" validate:"required"`
	// Name of the tenant
	Name string `json:"Name" validate:"required"`
	// description tenant description
	Description string `json:"Description" validate:"min=0,max=256"`
}

// UpdateUserProfileReq update user profile request
type UpdateUserProfileReq struct {
	// Nick user nick
	Nick string `json:"Nickname" validate:"required"`
	// Email user mail
	Email string `json:"Email" validate:"required,email"`
}

// UserList list user result
type UserList struct {
	ListMeta
	// Items list user slice
	Items []UserResp `json:"Items"`
}

// UserLoginReq user login request
type UserLoginReq struct {
	Name     string `json:"Name" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

// UserReq web operate user common define
type UserReq struct {
	// Name user name
	Name string `json:"Username" validate:"required"`
	// Nick user nick
	Nick string `json:"Nickname" validate:"required"`
	// Password user password, create user must have
	// +optional
	Password string `json:"Password,omitempty"`
	// Email user mail
	Email string `json:"Email" validate:"required,email"`
	// Invalid user is invalid
	Invalid bool `json:"Invalid"`
	// State, user' state: [Normal, Invalid, Deleting]
	// +optional
	State string `json:"State"`
	// Remote third platform name
	// +optional
	Remote string `json:"Remote,omitempty"`
	// RemoteID third platform id, if have remote, remoteID must have
	// +optional
	RemoteID string `json:"RemoteId,omitempty"`
	// RoleNames create user alloc user roles
	// +optional
	RoleNames []string `json:"Roles,omitempty"`
}

// UserResp cauth return user info to web
type UserResp struct {
	Meta `json:",inline"`
	// Name, user name
	Name string `json:"Username"`
	// Nick, user nick
	Nick string `json:"Nickname"`
	// Email, user mail
	Email string `json:"Email"`
	// Invalid
	Invalid bool `json:"Invalid"`
	// State, user' state: [Normal, Invalid, Deleting]
	State string `json:"State"`
	// Remote, third platform name
	Remote string `json:"Remote,omitempty"`
	// RemoteID, third platform id
	RemoteID string `json:"RemoteId,omitempty"`
	// RoleNames, user role list
	RoleNames []string `json:"Roles,omitempty"`
	// PasswordState, user password state: [Normal, New, Reset ]
	PasswordState string `json:"PasswordReset"`
}

// UserRole role info and role origin
type UserRole struct {
	UID         string    `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Origin      string    `json:"Origin"`
	BindTime    time.Time `json:"BindTime"`
	Policy      []string  `json:"Policy"`
	Resource    []string  `json:"Resource"`
}
