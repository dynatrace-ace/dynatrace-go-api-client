# AggregatedLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationResult** | Pointer to [**map[string]map[string]map[string]int64**](map.md) | Aggregated log records. | [optional] 

## Methods

### NewAggregatedLog

`func NewAggregatedLog() *AggregatedLog`

NewAggregatedLog instantiates a new AggregatedLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedLogWithDefaults

`func NewAggregatedLogWithDefaults() *AggregatedLog`

NewAggregatedLogWithDefaults instantiates a new AggregatedLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationResult

`func (o *AggregatedLog) GetAggregationResult() map[string]map[string]map[string]int64`

GetAggregationResult returns the AggregationResult field if non-nil, zero value otherwise.

### GetAggregationResultOk

`func (o *AggregatedLog) GetAggregationResultOk() (*map[string]map[string]map[string]int64, bool)`

GetAggregationResultOk returns a tuple with the AggregationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationResult

`func (o *AggregatedLog) SetAggregationResult(v map[string]map[string]map[string]int64)`

SetAggregationResult sets AggregationResult field to given value.

### HasAggregationResult

`func (o *AggregatedLog) HasAggregationResult() bool`

HasAggregationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


