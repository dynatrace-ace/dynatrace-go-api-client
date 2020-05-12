# MetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the metric event. | [optional] 
**MetricId** | **string** | The ID of the metric evaluated by the metric event. | 
**Name** | **string** | The name of the metric event displayed in the UI. | 
**Description** | **string** | The description of the metric event. | 
**AggregationType** | **string** | How the metric data points are aggregated for the evaluation.    The timeseries must support this aggregation. | [optional] 
**Severity** | **string** | The type of the event to trigger on the threshold violation.   The &#x60;CUSTOM_ALERT&#x60; type is not correlated with other alerts. The &#x60;INFO&#x60; type does not open a problem. | [optional] 
**AlertCondition** | **string** | The condition for the **threshold** value check: above or below. | 
**Samples** | **int32** | The number of one-minute samples that form sliding evaluation window. | 
**ViolatingSamples** | **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event | 
**DealertingSamples** | **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | 
**Threshold** | **float64** | The value of the threshold. | 
**Enabled** | **bool** | The metric event is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**AlertingScope** | [**[]MetricEventAlertingScope**](MetricEventAlertingScope.md) | Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction. | [optional] 
**MetricDimensions** | [**[]MetricEventDimensions**](MetricEventDimensions.md) | Defines the dimensions of the metric to alert on. The filters are combined by conjunction. | [optional] 
**Unit** | **string** | The unit of the threshold, matching the metric definition.  If not set the metrics unit is picked. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


