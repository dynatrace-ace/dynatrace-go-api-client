# CustomFilterChartSeriesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The name of the charted metric. | 
**Aggregation** | **string** | The charted aggregation of the metric. | 
**Percentile** | **int64** | The charted percentile.    Only applicable if the **aggregation** is set to &#x60;PERCENTILE&#x60;. | [optional] 
**Type** | **string** | The visualization of the timeseries chart. | 
**EntityType** | **string** | The type of the Dynatrace entity that delivered the charted metric. | 
**Dimensions** | [**[]CustomFilterChartSeriesDimensionConfig**](CustomFilterChartSeriesDimensionConfig.md) | Configuration of the charted metric splitting. | 
**SortAscending** | **bool** | Sort ascending (&#x60;true&#x60;) or descending (&#x60;false&#x60;).  | [optional] 
**SortColumn** | **bool** |  | [optional] 
**AggregationRate** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


