# ExtensionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the property. | [optional] 
**Type** | Pointer to **string** | The type of the property. | [optional] 
**DefaultValue** | Pointer to **string** | The default value of the property. | [optional] 
**DropdownValues** | Pointer to **[]string** | The list of possible values of the property.    If such a list is defined, only values from this list can be assigned to the property. | [optional] 

## Methods

### NewExtensionProperty

`func NewExtensionProperty() *ExtensionProperty`

NewExtensionProperty instantiates a new ExtensionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionPropertyWithDefaults

`func NewExtensionPropertyWithDefaults() *ExtensionProperty`

NewExtensionPropertyWithDefaults instantiates a new ExtensionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExtensionProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExtensionProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExtensionProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExtensionProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *ExtensionProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtensionProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtensionProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExtensionProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ExtensionProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExtensionProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExtensionProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExtensionProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDropdownValues

`func (o *ExtensionProperty) GetDropdownValues() []string`

GetDropdownValues returns the DropdownValues field if non-nil, zero value otherwise.

### GetDropdownValuesOk

`func (o *ExtensionProperty) GetDropdownValuesOk() (*[]string, bool)`

GetDropdownValuesOk returns a tuple with the DropdownValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropdownValues

`func (o *ExtensionProperty) SetDropdownValues(v []string)`

SetDropdownValues sets DropdownValues field to given value.

### HasDropdownValues

`func (o *ExtensionProperty) HasDropdownValues() bool`

HasDropdownValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


