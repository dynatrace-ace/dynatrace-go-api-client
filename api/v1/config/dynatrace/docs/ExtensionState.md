# ExtensionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | Pointer to **string** | The ID of the extension. | [optional] 
**Version** | Pointer to **string** | The version of the extension (for example &#x60;1.0.0&#x60;). | [optional] 
**EndpointId** | Pointer to **string** | The ID of the endpoint where the state is detected - Active Gate only. | [optional] 
**State** | Pointer to **string** | The state of the extension. | [optional] 
**StateDescription** | Pointer to **string** | A short description of the state. | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp when the state was detected, in UTC milliseconds. | [optional] 
**HostId** | Pointer to **string** | The ID of the host on which the extension runs. | [optional] 
**ProcessId** | Pointer to **string** | The ID of the entity on which the extension is active. | [optional] 

## Methods

### NewExtensionState

`func NewExtensionState() *ExtensionState`

NewExtensionState instantiates a new ExtensionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionStateWithDefaults

`func NewExtensionStateWithDefaults() *ExtensionState`

NewExtensionStateWithDefaults instantiates a new ExtensionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionId

`func (o *ExtensionState) GetExtensionId() string`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *ExtensionState) GetExtensionIdOk() (*string, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *ExtensionState) SetExtensionId(v string)`

SetExtensionId sets ExtensionId field to given value.

### HasExtensionId

`func (o *ExtensionState) HasExtensionId() bool`

HasExtensionId returns a boolean if a field has been set.

### GetVersion

`func (o *ExtensionState) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionState) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionState) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEndpointId

`func (o *ExtensionState) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ExtensionState) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ExtensionState) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ExtensionState) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetState

`func (o *ExtensionState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExtensionState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExtensionState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExtensionState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDescription

`func (o *ExtensionState) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *ExtensionState) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *ExtensionState) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *ExtensionState) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExtensionState) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtensionState) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtensionState) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtensionState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHostId

`func (o *ExtensionState) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ExtensionState) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ExtensionState) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ExtensionState) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetProcessId

`func (o *ExtensionState) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ExtensionState) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ExtensionState) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ExtensionState) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


