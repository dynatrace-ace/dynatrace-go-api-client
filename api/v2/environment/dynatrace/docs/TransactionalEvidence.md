# TransactionalEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueBeforeChangePoint** | **float32** | The metric&#39;s value before the problem start. | 
**ValueAfterChangePoint** | **float32** | The metric&#39;s value after the problem start. | 
**EndTime** | **int64** | The end time of the evidence, in UTC milliseconds | 
**Unit** | **string** | The unit of the metric. | 

## Methods

### NewTransactionalEvidence

`func NewTransactionalEvidence(valueBeforeChangePoint float32, valueAfterChangePoint float32, endTime int64, unit string, ) *TransactionalEvidence`

NewTransactionalEvidence instantiates a new TransactionalEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalEvidenceWithDefaults

`func NewTransactionalEvidenceWithDefaults() *TransactionalEvidence`

NewTransactionalEvidenceWithDefaults instantiates a new TransactionalEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueBeforeChangePoint

`func (o *TransactionalEvidence) GetValueBeforeChangePoint() float32`

GetValueBeforeChangePoint returns the ValueBeforeChangePoint field if non-nil, zero value otherwise.

### GetValueBeforeChangePointOk

`func (o *TransactionalEvidence) GetValueBeforeChangePointOk() (*float32, bool)`

GetValueBeforeChangePointOk returns a tuple with the ValueBeforeChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBeforeChangePoint

`func (o *TransactionalEvidence) SetValueBeforeChangePoint(v float32)`

SetValueBeforeChangePoint sets ValueBeforeChangePoint field to given value.


### GetValueAfterChangePoint

`func (o *TransactionalEvidence) GetValueAfterChangePoint() float32`

GetValueAfterChangePoint returns the ValueAfterChangePoint field if non-nil, zero value otherwise.

### GetValueAfterChangePointOk

`func (o *TransactionalEvidence) GetValueAfterChangePointOk() (*float32, bool)`

GetValueAfterChangePointOk returns a tuple with the ValueAfterChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAfterChangePoint

`func (o *TransactionalEvidence) SetValueAfterChangePoint(v float32)`

SetValueAfterChangePoint sets ValueAfterChangePoint field to given value.


### GetEndTime

`func (o *TransactionalEvidence) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TransactionalEvidence) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TransactionalEvidence) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetUnit

`func (o *TransactionalEvidence) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionalEvidence) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionalEvidence) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


