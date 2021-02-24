# RoleBasedAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamRole** | **string** | The IAM role to be used by Dynatrace to get monitoring data. | 
**AccountId** | **string** | The ID of the Amazon account. | 
**ExternalId** | Pointer to **string** | The external ID token for setting an IAM role.    You can obtain it with the &#x60;GET /aws/iamExternalId&#x60; request. | [optional] [readonly] 

## Methods

### NewRoleBasedAuthentication

`func NewRoleBasedAuthentication(iamRole string, accountId string, ) *RoleBasedAuthentication`

NewRoleBasedAuthentication instantiates a new RoleBasedAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBasedAuthenticationWithDefaults

`func NewRoleBasedAuthenticationWithDefaults() *RoleBasedAuthentication`

NewRoleBasedAuthenticationWithDefaults instantiates a new RoleBasedAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIamRole

`func (o *RoleBasedAuthentication) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *RoleBasedAuthentication) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *RoleBasedAuthentication) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.


### GetAccountId

`func (o *RoleBasedAuthentication) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoleBasedAuthentication) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RoleBasedAuthentication) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetExternalId

`func (o *RoleBasedAuthentication) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *RoleBasedAuthentication) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *RoleBasedAuthentication) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *RoleBasedAuthentication) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


