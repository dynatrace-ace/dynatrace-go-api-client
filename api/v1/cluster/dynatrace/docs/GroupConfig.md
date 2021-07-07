# GroupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClusterAdminGroup** | **bool** | If true, then the group has the cluster administrator rights. | 
**HasAccessAccountRole** | Pointer to **bool** | If true, then the group has the access account rights. | [optional] 
**HasManageAccountAndViewProductUsageRole** | Pointer to **bool** | If true, then the group has the manage account rights. | [optional] 
**IsAccessAccount** | Pointer to **bool** |  | [optional] 
**IsManageAccount** | Pointer to **bool** |  | [optional] 
**Id** | **string** | Group ID. Leave empty if creating group. Set if updating group. | 
**Name** | **string** | Group name | 
**LdapGroupNames** | Pointer to **[]string** | LDAP group names | [optional] 
**SsoGroupNames** | Pointer to **[]string** | SSO group names. If defined it&#39;s used to map SSO group name to Dynatrace group name, otherwise mapping is done by group name | [optional] 
**AccessRight** | Pointer to **map[string][]string** | Access rights | [optional] 

## Methods

### NewGroupConfig

`func NewGroupConfig(isClusterAdminGroup bool, id string, name string, ) *GroupConfig`

NewGroupConfig instantiates a new GroupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupConfigWithDefaults

`func NewGroupConfigWithDefaults() *GroupConfig`

NewGroupConfigWithDefaults instantiates a new GroupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClusterAdminGroup

`func (o *GroupConfig) GetIsClusterAdminGroup() bool`

GetIsClusterAdminGroup returns the IsClusterAdminGroup field if non-nil, zero value otherwise.

### GetIsClusterAdminGroupOk

`func (o *GroupConfig) GetIsClusterAdminGroupOk() (*bool, bool)`

GetIsClusterAdminGroupOk returns a tuple with the IsClusterAdminGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterAdminGroup

`func (o *GroupConfig) SetIsClusterAdminGroup(v bool)`

SetIsClusterAdminGroup sets IsClusterAdminGroup field to given value.


### GetHasAccessAccountRole

`func (o *GroupConfig) GetHasAccessAccountRole() bool`

GetHasAccessAccountRole returns the HasAccessAccountRole field if non-nil, zero value otherwise.

### GetHasAccessAccountRoleOk

`func (o *GroupConfig) GetHasAccessAccountRoleOk() (*bool, bool)`

GetHasAccessAccountRoleOk returns a tuple with the HasAccessAccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessAccountRole

`func (o *GroupConfig) SetHasAccessAccountRole(v bool)`

SetHasAccessAccountRole sets HasAccessAccountRole field to given value.

### HasHasAccessAccountRole

`func (o *GroupConfig) HasHasAccessAccountRole() bool`

HasHasAccessAccountRole returns a boolean if a field has been set.

### GetHasManageAccountAndViewProductUsageRole

`func (o *GroupConfig) GetHasManageAccountAndViewProductUsageRole() bool`

GetHasManageAccountAndViewProductUsageRole returns the HasManageAccountAndViewProductUsageRole field if non-nil, zero value otherwise.

### GetHasManageAccountAndViewProductUsageRoleOk

`func (o *GroupConfig) GetHasManageAccountAndViewProductUsageRoleOk() (*bool, bool)`

GetHasManageAccountAndViewProductUsageRoleOk returns a tuple with the HasManageAccountAndViewProductUsageRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManageAccountAndViewProductUsageRole

`func (o *GroupConfig) SetHasManageAccountAndViewProductUsageRole(v bool)`

SetHasManageAccountAndViewProductUsageRole sets HasManageAccountAndViewProductUsageRole field to given value.

### HasHasManageAccountAndViewProductUsageRole

`func (o *GroupConfig) HasHasManageAccountAndViewProductUsageRole() bool`

HasHasManageAccountAndViewProductUsageRole returns a boolean if a field has been set.

### GetIsAccessAccount

`func (o *GroupConfig) GetIsAccessAccount() bool`

GetIsAccessAccount returns the IsAccessAccount field if non-nil, zero value otherwise.

### GetIsAccessAccountOk

`func (o *GroupConfig) GetIsAccessAccountOk() (*bool, bool)`

GetIsAccessAccountOk returns a tuple with the IsAccessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessAccount

`func (o *GroupConfig) SetIsAccessAccount(v bool)`

SetIsAccessAccount sets IsAccessAccount field to given value.

### HasIsAccessAccount

`func (o *GroupConfig) HasIsAccessAccount() bool`

HasIsAccessAccount returns a boolean if a field has been set.

### GetIsManageAccount

`func (o *GroupConfig) GetIsManageAccount() bool`

GetIsManageAccount returns the IsManageAccount field if non-nil, zero value otherwise.

### GetIsManageAccountOk

`func (o *GroupConfig) GetIsManageAccountOk() (*bool, bool)`

GetIsManageAccountOk returns a tuple with the IsManageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManageAccount

`func (o *GroupConfig) SetIsManageAccount(v bool)`

SetIsManageAccount sets IsManageAccount field to given value.

### HasIsManageAccount

`func (o *GroupConfig) HasIsManageAccount() bool`

HasIsManageAccount returns a boolean if a field has been set.

### GetId

`func (o *GroupConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupConfig) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupConfig) SetName(v string)`

SetName sets Name field to given value.


### GetLdapGroupNames

`func (o *GroupConfig) GetLdapGroupNames() []string`

GetLdapGroupNames returns the LdapGroupNames field if non-nil, zero value otherwise.

### GetLdapGroupNamesOk

`func (o *GroupConfig) GetLdapGroupNamesOk() (*[]string, bool)`

GetLdapGroupNamesOk returns a tuple with the LdapGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupNames

`func (o *GroupConfig) SetLdapGroupNames(v []string)`

SetLdapGroupNames sets LdapGroupNames field to given value.

### HasLdapGroupNames

`func (o *GroupConfig) HasLdapGroupNames() bool`

HasLdapGroupNames returns a boolean if a field has been set.

### GetSsoGroupNames

`func (o *GroupConfig) GetSsoGroupNames() []string`

GetSsoGroupNames returns the SsoGroupNames field if non-nil, zero value otherwise.

### GetSsoGroupNamesOk

`func (o *GroupConfig) GetSsoGroupNamesOk() (*[]string, bool)`

GetSsoGroupNamesOk returns a tuple with the SsoGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoGroupNames

`func (o *GroupConfig) SetSsoGroupNames(v []string)`

SetSsoGroupNames sets SsoGroupNames field to given value.

### HasSsoGroupNames

`func (o *GroupConfig) HasSsoGroupNames() bool`

HasSsoGroupNames returns a boolean if a field has been set.

### GetAccessRight

`func (o *GroupConfig) GetAccessRight() map[string][]string`

GetAccessRight returns the AccessRight field if non-nil, zero value otherwise.

### GetAccessRightOk

`func (o *GroupConfig) GetAccessRightOk() (*map[string][]string, bool)`

GetAccessRightOk returns a tuple with the AccessRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRight

`func (o *GroupConfig) SetAccessRight(v map[string][]string)`

SetAccessRight sets AccessRight field to given value.

### HasAccessRight

`func (o *GroupConfig) HasAccessRight() bool`

HasAccessRight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


