# PluginInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginVersion** | Pointer to **string** | The version of the plugin. | [optional] 
**State** | Pointer to **string** | The state of the plugin instance. | [optional] 

## Methods

### NewPluginInstance

`func NewPluginInstance() *PluginInstance`

NewPluginInstance instantiates a new PluginInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInstanceWithDefaults

`func NewPluginInstanceWithDefaults() *PluginInstance`

NewPluginInstanceWithDefaults instantiates a new PluginInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginVersion

`func (o *PluginInstance) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *PluginInstance) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *PluginInstance) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *PluginInstance) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### GetState

`func (o *PluginInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PluginInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PluginInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PluginInstance) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


