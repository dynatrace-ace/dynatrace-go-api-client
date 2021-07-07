# EnumValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnumInstance** | Pointer to **string** | The name of the value in an existing Java enum class. | [optional] 
**Icon** | Pointer to **string** | The icon of the value. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The allowed value of the enum. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the value. | [optional] 
**Description** | Pointer to **string** | A short description of the value. | [optional] 

## Methods

### NewEnumValue

`func NewEnumValue() *EnumValue`

NewEnumValue instantiates a new EnumValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumValueWithDefaults

`func NewEnumValueWithDefaults() *EnumValue`

NewEnumValueWithDefaults instantiates a new EnumValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnumInstance

`func (o *EnumValue) GetEnumInstance() string`

GetEnumInstance returns the EnumInstance field if non-nil, zero value otherwise.

### GetEnumInstanceOk

`func (o *EnumValue) GetEnumInstanceOk() (*string, bool)`

GetEnumInstanceOk returns a tuple with the EnumInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumInstance

`func (o *EnumValue) SetEnumInstance(v string)`

SetEnumInstance sets EnumInstance field to given value.

### HasEnumInstance

`func (o *EnumValue) HasEnumInstance() bool`

HasEnumInstance returns a boolean if a field has been set.

### GetIcon

`func (o *EnumValue) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EnumValue) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EnumValue) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EnumValue) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetValue

`func (o *EnumValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *EnumValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnumValue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnumValue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnumValue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnumValue) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *EnumValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnumValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


