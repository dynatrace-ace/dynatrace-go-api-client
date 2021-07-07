# LocationExecutionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | Location id. | [optional] 
**ExecutionId** | Pointer to **string** | Execution id. | [optional] 
**RequestResults** | Pointer to [**[]MonitorRequestExecutionResult**](MonitorRequestExecutionResult.md) | The list of the monitor&#39;s request results executed on this location. | [optional] 

## Methods

### NewLocationExecutionResults

`func NewLocationExecutionResults() *LocationExecutionResults`

NewLocationExecutionResults instantiates a new LocationExecutionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationExecutionResultsWithDefaults

`func NewLocationExecutionResultsWithDefaults() *LocationExecutionResults`

NewLocationExecutionResultsWithDefaults instantiates a new LocationExecutionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *LocationExecutionResults) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *LocationExecutionResults) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *LocationExecutionResults) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *LocationExecutionResults) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetExecutionId

`func (o *LocationExecutionResults) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *LocationExecutionResults) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *LocationExecutionResults) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *LocationExecutionResults) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetRequestResults

`func (o *LocationExecutionResults) GetRequestResults() []MonitorRequestExecutionResult`

GetRequestResults returns the RequestResults field if non-nil, zero value otherwise.

### GetRequestResultsOk

`func (o *LocationExecutionResults) GetRequestResultsOk() (*[]MonitorRequestExecutionResult, bool)`

GetRequestResultsOk returns a tuple with the RequestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestResults

`func (o *LocationExecutionResults) SetRequestResults(v []MonitorRequestExecutionResult)`

SetRequestResults sets RequestResults field to given value.

### HasRequestResults

`func (o *LocationExecutionResults) HasRequestResults() bool`

HasRequestResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


