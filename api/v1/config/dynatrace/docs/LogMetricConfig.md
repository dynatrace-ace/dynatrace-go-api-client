# LogMetricConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** | The unique key of the metric.   The key must have the &#x60;calc:log.&#x60; prefix. | 
**Active** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**DisplayName** | **string** | The name of the metric, displayed in the UI. | 
**Unit** | **string** | The unit of the metric. | 
**UnitDisplayName** | **string** | The display name of the unit.    Only applicable if the **unit** is set to &#x60;UNSPECIFIED&#x60;. | [optional] 
**SearchString** | **string** | The pattern to look for in logs.    Use the [Dynatrace search query language](https://www.dynatrace.com/support/help/shortlink/log-viewer#search-for-text-patterns-in-log-files) to specify it. Quotes must be escaped.    To return all results, leave the field blank. | 
**MetricValueType** | **string** | The type of the metric data points calculation. For now the only allowed value is &#x60;OCCURRENCES&#x60;. | 
**ColumnDefiningValue** | [**ColumnDefinition**](ColumnDefinition.md) |  | [optional] 
**LogSourceFilters** | [**[]LogSourceFilter**](LogSourceFilter.md) | A list of filters to define the logs to look into.    If several filters are specified, the OR logic applies. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


