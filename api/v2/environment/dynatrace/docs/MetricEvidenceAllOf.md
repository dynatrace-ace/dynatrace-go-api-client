# MetricEvidenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | Pointer to **string** | The ID of the metric. | [optional] 
**ValueBeforeChangePoint** | Pointer to **float32** | The metric&#39;s value before the problem start. | [optional] 
**ValueAfterChangePoint** | Pointer to **float32** | The metric&#39;s value after the problem start. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the evidence, in UTC milliseconds.  The value &#x60;null&#x60; indicates that the evidence is still open. | [optional] 
**Unit** | Pointer to **string** | The unit of the metric. | [optional] 

## Methods

### NewMetricEvidenceAllOf

`func NewMetricEvidenceAllOf() *MetricEvidenceAllOf`

NewMetricEvidenceAllOf instantiates a new MetricEvidenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEvidenceAllOfWithDefaults

`func NewMetricEvidenceAllOfWithDefaults() *MetricEvidenceAllOf`

NewMetricEvidenceAllOfWithDefaults instantiates a new MetricEvidenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricEvidenceAllOf) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricEvidenceAllOf) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricEvidenceAllOf) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *MetricEvidenceAllOf) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetValueBeforeChangePoint

`func (o *MetricEvidenceAllOf) GetValueBeforeChangePoint() float32`

GetValueBeforeChangePoint returns the ValueBeforeChangePoint field if non-nil, zero value otherwise.

### GetValueBeforeChangePointOk

`func (o *MetricEvidenceAllOf) GetValueBeforeChangePointOk() (*float32, bool)`

GetValueBeforeChangePointOk returns a tuple with the ValueBeforeChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBeforeChangePoint

`func (o *MetricEvidenceAllOf) SetValueBeforeChangePoint(v float32)`

SetValueBeforeChangePoint sets ValueBeforeChangePoint field to given value.

### HasValueBeforeChangePoint

`func (o *MetricEvidenceAllOf) HasValueBeforeChangePoint() bool`

HasValueBeforeChangePoint returns a boolean if a field has been set.

### GetValueAfterChangePoint

`func (o *MetricEvidenceAllOf) GetValueAfterChangePoint() float32`

GetValueAfterChangePoint returns the ValueAfterChangePoint field if non-nil, zero value otherwise.

### GetValueAfterChangePointOk

`func (o *MetricEvidenceAllOf) GetValueAfterChangePointOk() (*float32, bool)`

GetValueAfterChangePointOk returns a tuple with the ValueAfterChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAfterChangePoint

`func (o *MetricEvidenceAllOf) SetValueAfterChangePoint(v float32)`

SetValueAfterChangePoint sets ValueAfterChangePoint field to given value.

### HasValueAfterChangePoint

`func (o *MetricEvidenceAllOf) HasValueAfterChangePoint() bool`

HasValueAfterChangePoint returns a boolean if a field has been set.

### GetEndTime

`func (o *MetricEvidenceAllOf) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MetricEvidenceAllOf) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MetricEvidenceAllOf) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MetricEvidenceAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUnit

`func (o *MetricEvidenceAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricEvidenceAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricEvidenceAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricEvidenceAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


