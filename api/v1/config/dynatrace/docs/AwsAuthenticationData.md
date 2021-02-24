# AwsAuthenticationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the authentication: role-based or key-based. | 
**KeyBasedAuthentication** | Pointer to [**KeyBasedAuthentication**](KeyBasedAuthentication.md) |  | [optional] 
**RoleBasedAuthentication** | Pointer to [**RoleBasedAuthentication**](RoleBasedAuthentication.md) |  | [optional] 

## Methods

### NewAwsAuthenticationData

`func NewAwsAuthenticationData(type_ string, ) *AwsAuthenticationData`

NewAwsAuthenticationData instantiates a new AwsAuthenticationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAuthenticationDataWithDefaults

`func NewAwsAuthenticationDataWithDefaults() *AwsAuthenticationData`

NewAwsAuthenticationDataWithDefaults instantiates a new AwsAuthenticationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwsAuthenticationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsAuthenticationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsAuthenticationData) SetType(v string)`

SetType sets Type field to given value.


### GetKeyBasedAuthentication

`func (o *AwsAuthenticationData) GetKeyBasedAuthentication() KeyBasedAuthentication`

GetKeyBasedAuthentication returns the KeyBasedAuthentication field if non-nil, zero value otherwise.

### GetKeyBasedAuthenticationOk

`func (o *AwsAuthenticationData) GetKeyBasedAuthenticationOk() (*KeyBasedAuthentication, bool)`

GetKeyBasedAuthenticationOk returns a tuple with the KeyBasedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBasedAuthentication

`func (o *AwsAuthenticationData) SetKeyBasedAuthentication(v KeyBasedAuthentication)`

SetKeyBasedAuthentication sets KeyBasedAuthentication field to given value.

### HasKeyBasedAuthentication

`func (o *AwsAuthenticationData) HasKeyBasedAuthentication() bool`

HasKeyBasedAuthentication returns a boolean if a field has been set.

### GetRoleBasedAuthentication

`func (o *AwsAuthenticationData) GetRoleBasedAuthentication() RoleBasedAuthentication`

GetRoleBasedAuthentication returns the RoleBasedAuthentication field if non-nil, zero value otherwise.

### GetRoleBasedAuthenticationOk

`func (o *AwsAuthenticationData) GetRoleBasedAuthenticationOk() (*RoleBasedAuthentication, bool)`

GetRoleBasedAuthenticationOk returns a tuple with the RoleBasedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBasedAuthentication

`func (o *AwsAuthenticationData) SetRoleBasedAuthentication(v RoleBasedAuthentication)`

SetRoleBasedAuthentication sets RoleBasedAuthentication field to given value.

### HasRoleBasedAuthentication

`func (o *AwsAuthenticationData) HasRoleBasedAuthentication() bool`

HasRoleBasedAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


