# PluginStateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**States** | Pointer to [**[]PluginState**](PluginState.md) | A list of plugin states. | [optional] 

## Methods

### NewPluginStateList

`func NewPluginStateList() *PluginStateList`

NewPluginStateList instantiates a new PluginStateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginStateListWithDefaults

`func NewPluginStateListWithDefaults() *PluginStateList`

NewPluginStateListWithDefaults instantiates a new PluginStateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStates

`func (o *PluginStateList) GetStates() []PluginState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *PluginStateList) GetStatesOk() (*[]PluginState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *PluginStateList) SetStates(v []PluginState)`

SetStates sets States field to given value.

### HasStates

`func (o *PluginStateList) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


