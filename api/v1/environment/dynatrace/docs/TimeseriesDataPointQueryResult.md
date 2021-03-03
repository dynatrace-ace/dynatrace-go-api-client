# TimeseriesDataPointQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | Pointer to [**map[string][][]float32**](array.md) | Data points of a metric.   A JSON object that maps the ID of the entity that delivered the data points and an array, which consists of arrays of the data point float values.   May contain more than one entity ID per record (for example, a host and its network interface). In such cases, entity IDs are separated by commas.   A datapoint contains a value and a timestamp, at which the value was recorded.    Dynatrace stores data in time slots. The **dataPoints** object shows the *starting* timestamp of the slot. If the **startTimestamp** or **endTimestamp** of your query lies inside of the data time slot, this time slot is included into response. Due to the fact that the timestamp of the first data point lies outside of the specified timeframe, you will see *earlier* timestamp than the specified **startTimestamp** in the first data point of the response.   There are three versions of data points:   * Numeric datapoint: Contains a numeric value.   * Enum datapoint: Contains an enum value, currently only for availability timeseries.   * Prediction datapoint: Similar to the numeric datapoint, but it contains a confidence interval, within which the future values are expected to be. | [optional] 
**TimeseriesId** | Pointer to **string** | The ID of the metric. | [optional] 
**Unit** | Pointer to **string** | The unit of data points. | [optional] 
**ResolutionInMillisUTC** | Pointer to **int64** | The resolution of data points. | [optional] 
**AggregationType** | Pointer to **string** | The type of data points aggregation. | [optional] 
**Entities** | Pointer to **map[string]string** | The list of entities where the data points originate.  A JSON object that maps the entity ID in Dynatrace and the actual name of the entity. | [optional] 

## Methods

### NewTimeseriesDataPointQueryResult

`func NewTimeseriesDataPointQueryResult() *TimeseriesDataPointQueryResult`

NewTimeseriesDataPointQueryResult instantiates a new TimeseriesDataPointQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesDataPointQueryResultWithDefaults

`func NewTimeseriesDataPointQueryResultWithDefaults() *TimeseriesDataPointQueryResult`

NewTimeseriesDataPointQueryResultWithDefaults instantiates a new TimeseriesDataPointQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *TimeseriesDataPointQueryResult) GetDataPoints() map[string][][]float32`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *TimeseriesDataPointQueryResult) GetDataPointsOk() (*map[string][][]float32, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *TimeseriesDataPointQueryResult) SetDataPoints(v map[string][][]float32)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *TimeseriesDataPointQueryResult) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.

### GetTimeseriesId

`func (o *TimeseriesDataPointQueryResult) GetTimeseriesId() string`

GetTimeseriesId returns the TimeseriesId field if non-nil, zero value otherwise.

### GetTimeseriesIdOk

`func (o *TimeseriesDataPointQueryResult) GetTimeseriesIdOk() (*string, bool)`

GetTimeseriesIdOk returns a tuple with the TimeseriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesId

`func (o *TimeseriesDataPointQueryResult) SetTimeseriesId(v string)`

SetTimeseriesId sets TimeseriesId field to given value.

### HasTimeseriesId

`func (o *TimeseriesDataPointQueryResult) HasTimeseriesId() bool`

HasTimeseriesId returns a boolean if a field has been set.

### GetUnit

`func (o *TimeseriesDataPointQueryResult) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeseriesDataPointQueryResult) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeseriesDataPointQueryResult) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeseriesDataPointQueryResult) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetResolutionInMillisUTC

`func (o *TimeseriesDataPointQueryResult) GetResolutionInMillisUTC() int64`

GetResolutionInMillisUTC returns the ResolutionInMillisUTC field if non-nil, zero value otherwise.

### GetResolutionInMillisUTCOk

`func (o *TimeseriesDataPointQueryResult) GetResolutionInMillisUTCOk() (*int64, bool)`

GetResolutionInMillisUTCOk returns a tuple with the ResolutionInMillisUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionInMillisUTC

`func (o *TimeseriesDataPointQueryResult) SetResolutionInMillisUTC(v int64)`

SetResolutionInMillisUTC sets ResolutionInMillisUTC field to given value.

### HasResolutionInMillisUTC

`func (o *TimeseriesDataPointQueryResult) HasResolutionInMillisUTC() bool`

HasResolutionInMillisUTC returns a boolean if a field has been set.

### GetAggregationType

`func (o *TimeseriesDataPointQueryResult) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *TimeseriesDataPointQueryResult) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *TimeseriesDataPointQueryResult) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *TimeseriesDataPointQueryResult) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetEntities

`func (o *TimeseriesDataPointQueryResult) GetEntities() map[string]string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TimeseriesDataPointQueryResult) GetEntitiesOk() (*map[string]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TimeseriesDataPointQueryResult) SetEntities(v map[string]string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *TimeseriesDataPointQueryResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


