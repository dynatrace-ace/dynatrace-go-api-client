# SettingsObjectErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidValue** | Pointer to **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Code** | Pointer to **int32** | The HTTP status code for the object. | [optional] 

## Methods

### NewSettingsObjectErrorResponse

`func NewSettingsObjectErrorResponse() *SettingsObjectErrorResponse`

NewSettingsObjectErrorResponse instantiates a new SettingsObjectErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectErrorResponseWithDefaults

`func NewSettingsObjectErrorResponseWithDefaults() *SettingsObjectErrorResponse`

NewSettingsObjectErrorResponseWithDefaults instantiates a new SettingsObjectErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidValue

`func (o *SettingsObjectErrorResponse) GetInvalidValue() map[string]interface{}`

GetInvalidValue returns the InvalidValue field if non-nil, zero value otherwise.

### GetInvalidValueOk

`func (o *SettingsObjectErrorResponse) GetInvalidValueOk() (*map[string]interface{}, bool)`

GetInvalidValueOk returns a tuple with the InvalidValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidValue

`func (o *SettingsObjectErrorResponse) SetInvalidValue(v map[string]interface{})`

SetInvalidValue sets InvalidValue field to given value.

### HasInvalidValue

`func (o *SettingsObjectErrorResponse) HasInvalidValue() bool`

HasInvalidValue returns a boolean if a field has been set.

### GetError

`func (o *SettingsObjectErrorResponse) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SettingsObjectErrorResponse) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SettingsObjectErrorResponse) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SettingsObjectErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCode

`func (o *SettingsObjectErrorResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SettingsObjectErrorResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SettingsObjectErrorResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SettingsObjectErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


