# TenantTokenConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**TenantToken**](TenantToken.md) |  | [optional] 
**Old** | Pointer to [**TenantToken**](TenantToken.md) |  | [optional] 

## Methods

### NewTenantTokenConfig

`func NewTenantTokenConfig() *TenantTokenConfig`

NewTenantTokenConfig instantiates a new TenantTokenConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTokenConfigWithDefaults

`func NewTenantTokenConfigWithDefaults() *TenantTokenConfig`

NewTenantTokenConfigWithDefaults instantiates a new TenantTokenConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TenantTokenConfig) GetActive() TenantToken`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TenantTokenConfig) GetActiveOk() (*TenantToken, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TenantTokenConfig) SetActive(v TenantToken)`

SetActive sets Active field to given value.

### HasActive

`func (o *TenantTokenConfig) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOld

`func (o *TenantTokenConfig) GetOld() TenantToken`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *TenantTokenConfig) GetOldOk() (*TenantToken, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *TenantTokenConfig) SetOld(v TenantToken)`

SetOld sets Old field to given value.

### HasOld

`func (o *TenantTokenConfig) HasOld() bool`

HasOld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


