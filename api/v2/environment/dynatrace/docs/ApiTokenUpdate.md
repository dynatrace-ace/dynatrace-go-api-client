# ApiTokenUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Enabled** | Pointer to **bool** | The token is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;) | [optional] 

## Methods

### NewApiTokenUpdate

`func NewApiTokenUpdate() *ApiTokenUpdate`

NewApiTokenUpdate instantiates a new ApiTokenUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenUpdateWithDefaults

`func NewApiTokenUpdateWithDefaults() *ApiTokenUpdate`

NewApiTokenUpdateWithDefaults instantiates a new ApiTokenUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTokenUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTokenUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTokenUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTokenUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiTokenUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiTokenUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiTokenUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiTokenUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


