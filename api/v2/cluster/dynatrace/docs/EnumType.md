# EnumType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]EnumValue**](EnumValue.md) | A list of allowed values of the enum. | [optional] 
**EnumClass** | Pointer to **string** | An existing Java enum class that holds the allowed values of the enum. | [optional] 
**Documentation** | Pointer to **string** | An extended description and/or links to documentation. | [optional] 
**Description** | Pointer to **string** | A short description of the property. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 

## Methods

### NewEnumType

`func NewEnumType() *EnumType`

NewEnumType instantiates a new EnumType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumTypeWithDefaults

`func NewEnumTypeWithDefaults() *EnumType`

NewEnumTypeWithDefaults instantiates a new EnumType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EnumType) GetItems() []EnumValue`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnumType) GetItemsOk() (*[]EnumValue, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnumType) SetItems(v []EnumValue)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnumType) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetEnumClass

`func (o *EnumType) GetEnumClass() string`

GetEnumClass returns the EnumClass field if non-nil, zero value otherwise.

### GetEnumClassOk

`func (o *EnumType) GetEnumClassOk() (*string, bool)`

GetEnumClassOk returns a tuple with the EnumClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumClass

`func (o *EnumType) SetEnumClass(v string)`

SetEnumClass sets EnumClass field to given value.

### HasEnumClass

`func (o *EnumType) HasEnumClass() bool`

HasEnumClass returns a boolean if a field has been set.

### GetDocumentation

`func (o *EnumType) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *EnumType) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *EnumType) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *EnumType) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetDescription

`func (o *EnumType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnumType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnumType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnumType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnumType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EnumType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


