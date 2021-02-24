# PluginState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginId** | Pointer to **string** | The ID of the plugin. | [optional] 
**Version** | Pointer to **string** | The version of the plugin (for example &#x60;1.0.0&#x60;). | [optional] 
**EndpointId** | Pointer to **string** | The ID of the endpoint where the state is detected - Active Gate only. | [optional] 
**State** | Pointer to **string** | The state of the plugin. | [optional] 
**StateDescription** | Pointer to **string** | A short description of the state. | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp when the state was detected, in UTC milliseconds. | [optional] 
**HostId** | Pointer to **string** | The ID of the host on which the plugin runs. | [optional] 
**ProcessId** | Pointer to **string** | The ID of the entity on which the plugin is active. | [optional] 

## Methods

### NewPluginState

`func NewPluginState() *PluginState`

NewPluginState instantiates a new PluginState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginStateWithDefaults

`func NewPluginStateWithDefaults() *PluginState`

NewPluginStateWithDefaults instantiates a new PluginState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginId

`func (o *PluginState) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *PluginState) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *PluginState) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *PluginState) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetVersion

`func (o *PluginState) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginState) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginState) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEndpointId

`func (o *PluginState) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *PluginState) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *PluginState) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *PluginState) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetState

`func (o *PluginState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PluginState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PluginState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PluginState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDescription

`func (o *PluginState) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *PluginState) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *PluginState) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *PluginState) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *PluginState) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PluginState) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PluginState) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PluginState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHostId

`func (o *PluginState) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *PluginState) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *PluginState) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *PluginState) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetProcessId

`func (o *PluginState) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *PluginState) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *PluginState) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *PluginState) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


