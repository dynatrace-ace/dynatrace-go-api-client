# MonitorExecutionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | Pointer to **string** | Monitor id. | [optional] 
**LocationsExecutionResults** | Pointer to [**[]LocationExecutionResults**](LocationExecutionResults.md) | The list with the results of the requests executed on assigned locations. | [optional] 

## Methods

### NewMonitorExecutionResults

`func NewMonitorExecutionResults() *MonitorExecutionResults`

NewMonitorExecutionResults instantiates a new MonitorExecutionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorExecutionResultsWithDefaults

`func NewMonitorExecutionResultsWithDefaults() *MonitorExecutionResults`

NewMonitorExecutionResultsWithDefaults instantiates a new MonitorExecutionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *MonitorExecutionResults) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorExecutionResults) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorExecutionResults) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *MonitorExecutionResults) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetLocationsExecutionResults

`func (o *MonitorExecutionResults) GetLocationsExecutionResults() []LocationExecutionResults`

GetLocationsExecutionResults returns the LocationsExecutionResults field if non-nil, zero value otherwise.

### GetLocationsExecutionResultsOk

`func (o *MonitorExecutionResults) GetLocationsExecutionResultsOk() (*[]LocationExecutionResults, bool)`

GetLocationsExecutionResultsOk returns a tuple with the LocationsExecutionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsExecutionResults

`func (o *MonitorExecutionResults) SetLocationsExecutionResults(v []LocationExecutionResults)`

SetLocationsExecutionResults sets LocationsExecutionResults field to given value.

### HasLocationsExecutionResults

`func (o *MonitorExecutionResults) HasLocationsExecutionResults() bool`

HasLocationsExecutionResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


