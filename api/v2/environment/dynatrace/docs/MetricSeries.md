# MetricSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionMap** | **map[string]string** |  | 
**Timestamps** | Pointer to **[]int64** | A list of timestamps of data points.   The value of data point for each time from this array is located in **values** array at the same index. | [optional] 
**Dimensions** | Pointer to **[]string** | The ordered list of dimensions to which the data point list belongs.    Each metric can have a certain number of dimensions. Dimensions exceeding this number are aggregated into one, which is shown as &#x60;null&#x60; here. | [optional] 
**Values** | Pointer to **[]float64** | A list of values of data points.   The timestamp of data point for each value from this array is located in **timestamps** array at the same index. | [optional] 

## Methods

### NewMetricSeries

`func NewMetricSeries(dimensionMap map[string]string, ) *MetricSeries`

NewMetricSeries instantiates a new MetricSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSeriesWithDefaults

`func NewMetricSeriesWithDefaults() *MetricSeries`

NewMetricSeriesWithDefaults instantiates a new MetricSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionMap

`func (o *MetricSeries) GetDimensionMap() map[string]string`

GetDimensionMap returns the DimensionMap field if non-nil, zero value otherwise.

### GetDimensionMapOk

`func (o *MetricSeries) GetDimensionMapOk() (*map[string]string, bool)`

GetDimensionMapOk returns a tuple with the DimensionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionMap

`func (o *MetricSeries) SetDimensionMap(v map[string]string)`

SetDimensionMap sets DimensionMap field to given value.


### GetTimestamps

`func (o *MetricSeries) GetTimestamps() []int64`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *MetricSeries) GetTimestampsOk() (*[]int64, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *MetricSeries) SetTimestamps(v []int64)`

SetTimestamps sets Timestamps field to given value.

### HasTimestamps

`func (o *MetricSeries) HasTimestamps() bool`

HasTimestamps returns a boolean if a field has been set.

### GetDimensions

`func (o *MetricSeries) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MetricSeries) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MetricSeries) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *MetricSeries) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetValues

`func (o *MetricSeries) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricSeries) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricSeries) SetValues(v []float64)`

SetValues sets Values field to given value.

### HasValues

`func (o *MetricSeries) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


