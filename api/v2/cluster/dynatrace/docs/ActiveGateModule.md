# ActiveGateModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Misconfigured** | Pointer to **bool** | The module is misconfigured (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the ActiveGate module. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | The module is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] [readonly] 
**Type** | Pointer to **string** | The type of ActiveGate module. | [optional] [readonly] 
**Attributes** | Pointer to **map[string]string** | The attributes of the ActiveGate module. | [optional] [readonly] 

## Methods

### NewActiveGateModule

`func NewActiveGateModule() *ActiveGateModule`

NewActiveGateModule instantiates a new ActiveGateModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateModuleWithDefaults

`func NewActiveGateModuleWithDefaults() *ActiveGateModule`

NewActiveGateModuleWithDefaults instantiates a new ActiveGateModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMisconfigured

`func (o *ActiveGateModule) GetMisconfigured() bool`

GetMisconfigured returns the Misconfigured field if non-nil, zero value otherwise.

### GetMisconfiguredOk

`func (o *ActiveGateModule) GetMisconfiguredOk() (*bool, bool)`

GetMisconfiguredOk returns a tuple with the Misconfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisconfigured

`func (o *ActiveGateModule) SetMisconfigured(v bool)`

SetMisconfigured sets Misconfigured field to given value.

### HasMisconfigured

`func (o *ActiveGateModule) HasMisconfigured() bool`

HasMisconfigured returns a boolean if a field has been set.

### GetVersion

`func (o *ActiveGateModule) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActiveGateModule) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActiveGateModule) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActiveGateModule) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *ActiveGateModule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActiveGateModule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActiveGateModule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActiveGateModule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ActiveGateModule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActiveGateModule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActiveGateModule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActiveGateModule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *ActiveGateModule) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ActiveGateModule) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ActiveGateModule) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ActiveGateModule) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


