# CustomFilterChartSeriesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The name of the charted metric. | 
**Aggregation** | **string** | The charted aggregation of the metric. | 
**Percentile** | Pointer to **int64** | The charted percentile.    Only applicable if the **aggregation** is set to &#x60;PERCENTILE&#x60;. | [optional] 
**Type** | **string** | The visualization of the timeseries chart. | 
**EntityType** | **string** | The type of the Dynatrace entity that delivered the charted metric. | 
**Dimensions** | [**[]CustomFilterChartSeriesDimensionConfig**](CustomFilterChartSeriesDimensionConfig.md) | Configuration of the charted metric splitting. | 
**SortAscending** | Pointer to **bool** | Sort ascending (&#x60;true&#x60;) or descending (&#x60;false&#x60;).  | [optional] 
**SortColumn** | Pointer to **bool** |  | [optional] 
**AggregationRate** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFilterChartSeriesConfig

`func NewCustomFilterChartSeriesConfig(metric string, aggregation string, type_ string, entityType string, dimensions []CustomFilterChartSeriesDimensionConfig, ) *CustomFilterChartSeriesConfig`

NewCustomFilterChartSeriesConfig instantiates a new CustomFilterChartSeriesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterChartSeriesConfigWithDefaults

`func NewCustomFilterChartSeriesConfigWithDefaults() *CustomFilterChartSeriesConfig`

NewCustomFilterChartSeriesConfigWithDefaults instantiates a new CustomFilterChartSeriesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *CustomFilterChartSeriesConfig) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CustomFilterChartSeriesConfig) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CustomFilterChartSeriesConfig) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetAggregation

`func (o *CustomFilterChartSeriesConfig) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *CustomFilterChartSeriesConfig) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *CustomFilterChartSeriesConfig) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetPercentile

`func (o *CustomFilterChartSeriesConfig) GetPercentile() int64`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *CustomFilterChartSeriesConfig) GetPercentileOk() (*int64, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *CustomFilterChartSeriesConfig) SetPercentile(v int64)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *CustomFilterChartSeriesConfig) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetType

`func (o *CustomFilterChartSeriesConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFilterChartSeriesConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFilterChartSeriesConfig) SetType(v string)`

SetType sets Type field to given value.


### GetEntityType

`func (o *CustomFilterChartSeriesConfig) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomFilterChartSeriesConfig) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomFilterChartSeriesConfig) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDimensions

`func (o *CustomFilterChartSeriesConfig) GetDimensions() []CustomFilterChartSeriesDimensionConfig`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CustomFilterChartSeriesConfig) GetDimensionsOk() (*[]CustomFilterChartSeriesDimensionConfig, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CustomFilterChartSeriesConfig) SetDimensions(v []CustomFilterChartSeriesDimensionConfig)`

SetDimensions sets Dimensions field to given value.


### GetSortAscending

`func (o *CustomFilterChartSeriesConfig) GetSortAscending() bool`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *CustomFilterChartSeriesConfig) GetSortAscendingOk() (*bool, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *CustomFilterChartSeriesConfig) SetSortAscending(v bool)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *CustomFilterChartSeriesConfig) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.

### GetSortColumn

`func (o *CustomFilterChartSeriesConfig) GetSortColumn() bool`

GetSortColumn returns the SortColumn field if non-nil, zero value otherwise.

### GetSortColumnOk

`func (o *CustomFilterChartSeriesConfig) GetSortColumnOk() (*bool, bool)`

GetSortColumnOk returns a tuple with the SortColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumn

`func (o *CustomFilterChartSeriesConfig) SetSortColumn(v bool)`

SetSortColumn sets SortColumn field to given value.

### HasSortColumn

`func (o *CustomFilterChartSeriesConfig) HasSortColumn() bool`

HasSortColumn returns a boolean if a field has been set.

### GetAggregationRate

`func (o *CustomFilterChartSeriesConfig) GetAggregationRate() string`

GetAggregationRate returns the AggregationRate field if non-nil, zero value otherwise.

### GetAggregationRateOk

`func (o *CustomFilterChartSeriesConfig) GetAggregationRateOk() (*string, bool)`

GetAggregationRateOk returns a tuple with the AggregationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRate

`func (o *CustomFilterChartSeriesConfig) SetAggregationRate(v string)`

SetAggregationRate sets AggregationRate field to given value.

### HasAggregationRate

`func (o *CustomFilterChartSeriesConfig) HasAggregationRate() bool`

HasAggregationRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


