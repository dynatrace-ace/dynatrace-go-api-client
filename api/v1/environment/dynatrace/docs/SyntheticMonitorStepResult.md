# SyntheticMonitorStepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the step. It is unique within the test definition. | 
**StartTimestamp** | **int64** | The timestamp of text step execution, UTC milliseconds. | 
**ResponseTimeMillis** | Pointer to **int64** | The response time of the step, in milliseconds.    Absent when no meaningful response time is available (as may be the case for certain error conditions such as a misconfigured step script). | [optional] 
**Error** | Pointer to [**SyntheticMonitorError**](SyntheticMonitorError.md) |  | [optional] 

## Methods

### NewSyntheticMonitorStepResult

`func NewSyntheticMonitorStepResult(id int64, startTimestamp int64, ) *SyntheticMonitorStepResult`

NewSyntheticMonitorStepResult instantiates a new SyntheticMonitorStepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticMonitorStepResultWithDefaults

`func NewSyntheticMonitorStepResultWithDefaults() *SyntheticMonitorStepResult`

NewSyntheticMonitorStepResultWithDefaults instantiates a new SyntheticMonitorStepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyntheticMonitorStepResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticMonitorStepResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticMonitorStepResult) SetId(v int64)`

SetId sets Id field to given value.


### GetStartTimestamp

`func (o *SyntheticMonitorStepResult) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SyntheticMonitorStepResult) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SyntheticMonitorStepResult) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetResponseTimeMillis

`func (o *SyntheticMonitorStepResult) GetResponseTimeMillis() int64`

GetResponseTimeMillis returns the ResponseTimeMillis field if non-nil, zero value otherwise.

### GetResponseTimeMillisOk

`func (o *SyntheticMonitorStepResult) GetResponseTimeMillisOk() (*int64, bool)`

GetResponseTimeMillisOk returns a tuple with the ResponseTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMillis

`func (o *SyntheticMonitorStepResult) SetResponseTimeMillis(v int64)`

SetResponseTimeMillis sets ResponseTimeMillis field to given value.

### HasResponseTimeMillis

`func (o *SyntheticMonitorStepResult) HasResponseTimeMillis() bool`

HasResponseTimeMillis returns a boolean if a field has been set.

### GetError

`func (o *SyntheticMonitorStepResult) GetError() SyntheticMonitorError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SyntheticMonitorStepResult) GetErrorOk() (*SyntheticMonitorError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SyntheticMonitorStepResult) SetError(v SyntheticMonitorError)`

SetError sets Error field to given value.

### HasError

`func (o *SyntheticMonitorStepResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


