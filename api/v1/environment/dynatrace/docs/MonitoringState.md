# MonitoringState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualMonitoringState** | Pointer to **string** | The current actual monitoring state on the entity. | [optional] 
**ExpectedMonitoringState** | Pointer to **string** | The monitoring state that is expected from the configuration | [optional] 
**RestartRequired** | Pointer to **bool** | Defines whether or not the process has to restarted to enable monitoring | [optional] 

## Methods

### NewMonitoringState

`func NewMonitoringState() *MonitoringState`

NewMonitoringState instantiates a new MonitoringState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringStateWithDefaults

`func NewMonitoringStateWithDefaults() *MonitoringState`

NewMonitoringStateWithDefaults instantiates a new MonitoringState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualMonitoringState

`func (o *MonitoringState) GetActualMonitoringState() string`

GetActualMonitoringState returns the ActualMonitoringState field if non-nil, zero value otherwise.

### GetActualMonitoringStateOk

`func (o *MonitoringState) GetActualMonitoringStateOk() (*string, bool)`

GetActualMonitoringStateOk returns a tuple with the ActualMonitoringState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMonitoringState

`func (o *MonitoringState) SetActualMonitoringState(v string)`

SetActualMonitoringState sets ActualMonitoringState field to given value.

### HasActualMonitoringState

`func (o *MonitoringState) HasActualMonitoringState() bool`

HasActualMonitoringState returns a boolean if a field has been set.

### GetExpectedMonitoringState

`func (o *MonitoringState) GetExpectedMonitoringState() string`

GetExpectedMonitoringState returns the ExpectedMonitoringState field if non-nil, zero value otherwise.

### GetExpectedMonitoringStateOk

`func (o *MonitoringState) GetExpectedMonitoringStateOk() (*string, bool)`

GetExpectedMonitoringStateOk returns a tuple with the ExpectedMonitoringState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedMonitoringState

`func (o *MonitoringState) SetExpectedMonitoringState(v string)`

SetExpectedMonitoringState sets ExpectedMonitoringState field to given value.

### HasExpectedMonitoringState

`func (o *MonitoringState) HasExpectedMonitoringState() bool`

HasExpectedMonitoringState returns a boolean if a field has been set.

### GetRestartRequired

`func (o *MonitoringState) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *MonitoringState) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *MonitoringState) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *MonitoringState) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


