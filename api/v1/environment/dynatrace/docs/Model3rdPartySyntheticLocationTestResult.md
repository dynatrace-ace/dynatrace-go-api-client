# Model3rdPartySyntheticLocationTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the location. | 
**StartTimestamp** | **int64** | The timestamp of text execution start, in UTC milliseconds. | 
**SuccessRate** | Pointer to **float64** | The overall availability of the monitor from this location, percent.    If absent, calculated as the number of successful steps compared to the overall number of steps. | [optional] 
**Success** | **bool** | If the test was successful (&#x60;true&#x60;) or failed (&#x60;false&#x60;) - will influence availability timeseries. | 
**ResponseTimeMillis** | Pointer to **int64** | The overall response time of the monitor from this location, in milliseconds.    If absent, it is calculated as the sum of response times of all steps. | [optional] 
**StepResults** | Pointer to [**[]SyntheticMonitorStepResult**](SyntheticMonitorStepResult.md) | Results of individual monitor steps. | [optional] 

## Methods

### NewModel3rdPartySyntheticLocationTestResult

`func NewModel3rdPartySyntheticLocationTestResult(id string, startTimestamp int64, success bool, ) *Model3rdPartySyntheticLocationTestResult`

NewModel3rdPartySyntheticLocationTestResult instantiates a new Model3rdPartySyntheticLocationTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartySyntheticLocationTestResultWithDefaults

`func NewModel3rdPartySyntheticLocationTestResultWithDefaults() *Model3rdPartySyntheticLocationTestResult`

NewModel3rdPartySyntheticLocationTestResultWithDefaults instantiates a new Model3rdPartySyntheticLocationTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model3rdPartySyntheticLocationTestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model3rdPartySyntheticLocationTestResult) SetId(v string)`

SetId sets Id field to given value.


### GetStartTimestamp

`func (o *Model3rdPartySyntheticLocationTestResult) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Model3rdPartySyntheticLocationTestResult) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetSuccessRate

`func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *Model3rdPartySyntheticLocationTestResult) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *Model3rdPartySyntheticLocationTestResult) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetSuccess

`func (o *Model3rdPartySyntheticLocationTestResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Model3rdPartySyntheticLocationTestResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetResponseTimeMillis

`func (o *Model3rdPartySyntheticLocationTestResult) GetResponseTimeMillis() int64`

GetResponseTimeMillis returns the ResponseTimeMillis field if non-nil, zero value otherwise.

### GetResponseTimeMillisOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetResponseTimeMillisOk() (*int64, bool)`

GetResponseTimeMillisOk returns a tuple with the ResponseTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMillis

`func (o *Model3rdPartySyntheticLocationTestResult) SetResponseTimeMillis(v int64)`

SetResponseTimeMillis sets ResponseTimeMillis field to given value.

### HasResponseTimeMillis

`func (o *Model3rdPartySyntheticLocationTestResult) HasResponseTimeMillis() bool`

HasResponseTimeMillis returns a boolean if a field has been set.

### GetStepResults

`func (o *Model3rdPartySyntheticLocationTestResult) GetStepResults() []SyntheticMonitorStepResult`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *Model3rdPartySyntheticLocationTestResult) GetStepResultsOk() (*[]SyntheticMonitorStepResult, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *Model3rdPartySyntheticLocationTestResult) SetStepResults(v []SyntheticMonitorStepResult)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *Model3rdPartySyntheticLocationTestResult) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


