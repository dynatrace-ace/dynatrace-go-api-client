# PluginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | Pointer to **string** | The name of the plugin. | [optional] 
**Instances** | Pointer to [**[]PluginInstance**](PluginInstance.md) | A list of instances of the plugin. | [optional] 

## Methods

### NewPluginInfo

`func NewPluginInfo() *PluginInfo`

NewPluginInfo instantiates a new PluginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInfoWithDefaults

`func NewPluginInfoWithDefaults() *PluginInfo`

NewPluginInfoWithDefaults instantiates a new PluginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *PluginInfo) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *PluginInfo) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *PluginInfo) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.

### HasPluginName

`func (o *PluginInfo) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.

### GetInstances

`func (o *PluginInfo) GetInstances() []PluginInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *PluginInfo) GetInstancesOk() (*[]PluginInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *PluginInfo) SetInstances(v []PluginInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *PluginInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


