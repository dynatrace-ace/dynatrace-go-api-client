# Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the plugin, for example &#x60;custom.remote.python.demo&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the plugin, displayed in Dynatrace. | [optional] 
**Version** | Pointer to **string** | The version of the plugin, displayed in Dynatrace. | [optional] 
**Type** | Pointer to **string** | The type of the plugin. It indicates the runtime environment of the plugin (for example, ActiveGate). | [optional] 
**MetricGroup** | Pointer to **string** | The metric group of the plugin. All the metrics, captured by the plugin are children of this group. | [optional] 
**Properties** | Pointer to [**[]PluginProperty**](PluginProperty.md) | A list of plugin properties. | [optional] 

## Methods

### NewPlugin

`func NewPlugin() *Plugin`

NewPlugin instantiates a new Plugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginWithDefaults

`func NewPluginWithDefaults() *Plugin`

NewPluginWithDefaults instantiates a new Plugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Plugin) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Plugin) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Plugin) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Plugin) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *Plugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Plugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Plugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Plugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Plugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Plugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Plugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *Plugin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Plugin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Plugin) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Plugin) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetricGroup

`func (o *Plugin) GetMetricGroup() string`

GetMetricGroup returns the MetricGroup field if non-nil, zero value otherwise.

### GetMetricGroupOk

`func (o *Plugin) GetMetricGroupOk() (*string, bool)`

GetMetricGroupOk returns a tuple with the MetricGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroup

`func (o *Plugin) SetMetricGroup(v string)`

SetMetricGroup sets MetricGroup field to given value.

### HasMetricGroup

`func (o *Plugin) HasMetricGroup() bool`

HasMetricGroup returns a boolean if a field has been set.

### GetProperties

`func (o *Plugin) GetProperties() []PluginProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Plugin) GetPropertiesOk() (*[]PluginProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Plugin) SetProperties(v []PluginProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Plugin) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


