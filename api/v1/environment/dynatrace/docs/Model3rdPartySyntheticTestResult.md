# Model3rdPartySyntheticTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the third-party synthetic monitor. | 
**ScheduleIntervalInSeconds** | Pointer to **int32** |  | [optional] 
**TotalStepCount** | Pointer to **int32** | Number of steps in the monitor. Defaults to number of SyntheticTestSteps. | [optional] 
**LocationResults** | [**[]Model3rdPartySyntheticLocationTestResult**](3rdPartySyntheticLocationTestResult.md) | Results of third-party monitor executions per location. | 

## Methods

### NewModel3rdPartySyntheticTestResult

`func NewModel3rdPartySyntheticTestResult(id string, locationResults []Model3rdPartySyntheticLocationTestResult, ) *Model3rdPartySyntheticTestResult`

NewModel3rdPartySyntheticTestResult instantiates a new Model3rdPartySyntheticTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartySyntheticTestResultWithDefaults

`func NewModel3rdPartySyntheticTestResultWithDefaults() *Model3rdPartySyntheticTestResult`

NewModel3rdPartySyntheticTestResultWithDefaults instantiates a new Model3rdPartySyntheticTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model3rdPartySyntheticTestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model3rdPartySyntheticTestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model3rdPartySyntheticTestResult) SetId(v string)`

SetId sets Id field to given value.


### GetScheduleIntervalInSeconds

`func (o *Model3rdPartySyntheticTestResult) GetScheduleIntervalInSeconds() int32`

GetScheduleIntervalInSeconds returns the ScheduleIntervalInSeconds field if non-nil, zero value otherwise.

### GetScheduleIntervalInSecondsOk

`func (o *Model3rdPartySyntheticTestResult) GetScheduleIntervalInSecondsOk() (*int32, bool)`

GetScheduleIntervalInSecondsOk returns a tuple with the ScheduleIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleIntervalInSeconds

`func (o *Model3rdPartySyntheticTestResult) SetScheduleIntervalInSeconds(v int32)`

SetScheduleIntervalInSeconds sets ScheduleIntervalInSeconds field to given value.

### HasScheduleIntervalInSeconds

`func (o *Model3rdPartySyntheticTestResult) HasScheduleIntervalInSeconds() bool`

HasScheduleIntervalInSeconds returns a boolean if a field has been set.

### GetTotalStepCount

`func (o *Model3rdPartySyntheticTestResult) GetTotalStepCount() int32`

GetTotalStepCount returns the TotalStepCount field if non-nil, zero value otherwise.

### GetTotalStepCountOk

`func (o *Model3rdPartySyntheticTestResult) GetTotalStepCountOk() (*int32, bool)`

GetTotalStepCountOk returns a tuple with the TotalStepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStepCount

`func (o *Model3rdPartySyntheticTestResult) SetTotalStepCount(v int32)`

SetTotalStepCount sets TotalStepCount field to given value.

### HasTotalStepCount

`func (o *Model3rdPartySyntheticTestResult) HasTotalStepCount() bool`

HasTotalStepCount returns a boolean if a field has been set.

### GetLocationResults

`func (o *Model3rdPartySyntheticTestResult) GetLocationResults() []Model3rdPartySyntheticLocationTestResult`

GetLocationResults returns the LocationResults field if non-nil, zero value otherwise.

### GetLocationResultsOk

`func (o *Model3rdPartySyntheticTestResult) GetLocationResultsOk() (*[]Model3rdPartySyntheticLocationTestResult, bool)`

GetLocationResultsOk returns a tuple with the LocationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationResults

`func (o *Model3rdPartySyntheticTestResult) SetLocationResults(v []Model3rdPartySyntheticLocationTestResult)`

SetLocationResults sets LocationResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


