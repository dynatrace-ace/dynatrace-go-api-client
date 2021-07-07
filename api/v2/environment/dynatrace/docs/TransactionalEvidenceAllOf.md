# TransactionalEvidenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueBeforeChangePoint** | Pointer to **float32** | The metric&#39;s value before the problem start. | [optional] 
**ValueAfterChangePoint** | Pointer to **float32** | The metric&#39;s value after the problem start. | [optional] 
**EndTime** | Pointer to **int64** | The end time of the evidence, in UTC milliseconds | [optional] 
**Unit** | Pointer to **string** | The unit of the metric. | [optional] 

## Methods

### NewTransactionalEvidenceAllOf

`func NewTransactionalEvidenceAllOf() *TransactionalEvidenceAllOf`

NewTransactionalEvidenceAllOf instantiates a new TransactionalEvidenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalEvidenceAllOfWithDefaults

`func NewTransactionalEvidenceAllOfWithDefaults() *TransactionalEvidenceAllOf`

NewTransactionalEvidenceAllOfWithDefaults instantiates a new TransactionalEvidenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueBeforeChangePoint

`func (o *TransactionalEvidenceAllOf) GetValueBeforeChangePoint() float32`

GetValueBeforeChangePoint returns the ValueBeforeChangePoint field if non-nil, zero value otherwise.

### GetValueBeforeChangePointOk

`func (o *TransactionalEvidenceAllOf) GetValueBeforeChangePointOk() (*float32, bool)`

GetValueBeforeChangePointOk returns a tuple with the ValueBeforeChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBeforeChangePoint

`func (o *TransactionalEvidenceAllOf) SetValueBeforeChangePoint(v float32)`

SetValueBeforeChangePoint sets ValueBeforeChangePoint field to given value.

### HasValueBeforeChangePoint

`func (o *TransactionalEvidenceAllOf) HasValueBeforeChangePoint() bool`

HasValueBeforeChangePoint returns a boolean if a field has been set.

### GetValueAfterChangePoint

`func (o *TransactionalEvidenceAllOf) GetValueAfterChangePoint() float32`

GetValueAfterChangePoint returns the ValueAfterChangePoint field if non-nil, zero value otherwise.

### GetValueAfterChangePointOk

`func (o *TransactionalEvidenceAllOf) GetValueAfterChangePointOk() (*float32, bool)`

GetValueAfterChangePointOk returns a tuple with the ValueAfterChangePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAfterChangePoint

`func (o *TransactionalEvidenceAllOf) SetValueAfterChangePoint(v float32)`

SetValueAfterChangePoint sets ValueAfterChangePoint field to given value.

### HasValueAfterChangePoint

`func (o *TransactionalEvidenceAllOf) HasValueAfterChangePoint() bool`

HasValueAfterChangePoint returns a boolean if a field has been set.

### GetEndTime

`func (o *TransactionalEvidenceAllOf) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TransactionalEvidenceAllOf) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TransactionalEvidenceAllOf) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TransactionalEvidenceAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUnit

`func (o *TransactionalEvidenceAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TransactionalEvidenceAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TransactionalEvidenceAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TransactionalEvidenceAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


