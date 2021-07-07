# MetricEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | **string** | The ID of the metric. | 
**ValueBeforeChangePoint** | **float32** | The metric&#39;s value before the problem start. | 
**ValueAfterChangePoint** | **float32** | The metric&#39;s value after the problem start. | 
**EndTime** | **int64** | The end time of the evidence, in UTC milliseconds.  The value &#x60;null&#x60; indicates that the evidence is still open. | 
**Unit** | **string** | The unit of the metric. | 

## Methods

### NewMetricEvidence

`func NewMetricEvidence(metricId string, valueBeforeChangePoint float32, valueAfterChangePoint float32, endTime int64, unit string, ) *MetricEvidence`

NewMetricEvidence instantiates a new MetricEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEvidenceWithDefaults

`func NewMetricEvidenceWithDefaults() *MetricEvidence`

NewMetricEvidenceWithDefaults instantiates a new MetricEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricEvidence) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricEvidence) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricEvidence) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetValueBeforeChangePoint

`func (o *MetricEvidence) GetValueBeforeChangePoint() float32`

GetValueBeforeChangePoint returns the ValueBeforeChangePoint field if non-nil, zero value otherwise.

### GetValueBeforeChangePointOk

`func (o *MetricEvidence) GetValueBeforeChangePointOk() (*float32, bool)`

GetValueBeforeChangePointOk returns a tuple with the ValueBeforeChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBeforeChangePoint

`func (o *MetricEvidence) SetValueBeforeChangePoint(v float32)`

SetValueBeforeChangePoint sets ValueBeforeChangePoint field to given value.


### GetValueAfterChangePoint

`func (o *MetricEvidence) GetValueAfterChangePoint() float32`

GetValueAfterChangePoint returns the ValueAfterChangePoint field if non-nil, zero value otherwise.

### GetValueAfterChangePointOk

`func (o *MetricEvidence) GetValueAfterChangePointOk() (*float32, bool)`

GetValueAfterChangePointOk returns a tuple with the ValueAfterChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAfterChangePoint

`func (o *MetricEvidence) SetValueAfterChangePoint(v float32)`

SetValueAfterChangePoint sets ValueAfterChangePoint field to given value.


### GetEndTime

`func (o *MetricEvidence) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MetricEvidence) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MetricEvidence) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetUnit

`func (o *MetricEvidence) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricEvidence) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricEvidence) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


