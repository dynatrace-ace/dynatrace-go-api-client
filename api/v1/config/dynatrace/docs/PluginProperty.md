# PluginProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the property. | [optional] 
**Type** | Pointer to **string** | The type of the property. | [optional] 
**DefaultValue** | Pointer to **string** | The default value of the property. | [optional] 
**DropdownValues** | Pointer to **[]string** | The list of possible values of the property.    If such a list is defined, only values from this list can be assigned to the property. | [optional] 

## Methods

### NewPluginProperty

`func NewPluginProperty() *PluginProperty`

NewPluginProperty instantiates a new PluginProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginPropertyWithDefaults

`func NewPluginPropertyWithDefaults() *PluginProperty`

NewPluginPropertyWithDefaults instantiates a new PluginProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PluginProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PluginProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PluginProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PluginProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *PluginProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PluginProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PluginProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PluginProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PluginProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PluginProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PluginProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PluginProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDropdownValues

`func (o *PluginProperty) GetDropdownValues() []string`

GetDropdownValues returns the DropdownValues field if non-nil, zero value otherwise.

### GetDropdownValuesOk

`func (o *PluginProperty) GetDropdownValuesOk() (*[]string, bool)`

GetDropdownValuesOk returns a tuple with the DropdownValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropdownValues

`func (o *PluginProperty) SetDropdownValues(v []string)`

SetDropdownValues sets DropdownValues field to given value.

### HasDropdownValues

`func (o *PluginProperty) HasDropdownValues() bool`

HasDropdownValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


