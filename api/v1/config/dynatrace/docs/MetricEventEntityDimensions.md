# MetricEventEntityDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;ENTITY&#x60; -&gt; MetricEventEntityDimensions  * &#x60;STRING&#x60; -&gt; MetricEventStringDimensions   | 
**Name** | **string** | The dimensions name. Sending this has no effect while creating a configuration, as only the *index* of the dimension is used here -&gt; dimension names might change, indexes not. | [optional] 
**Index** | **int32** | The dimensions index on the metric. | 
**NameFilter** | [**MetricEventTextFilterMetricEventDimensionsFilterOperatorDto**](MetricEventTextFilterMetricEventDimensionsFilterOperatorDto.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


