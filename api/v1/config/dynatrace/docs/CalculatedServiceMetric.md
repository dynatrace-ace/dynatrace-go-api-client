# CalculatedServiceMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**TsmMetricKey** | **string** | The key of the calculated service metric. | 
**Name** | **string** | The displayed name of the metric. | 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MetricDefinition** | [**CalculatedMetricDefinition**](CalculatedMetricDefinition.md) |  | 
**Unit** | **string** | The unit of the metric. | 
**UnitDisplayName** | **string** | The display name of the metric&#39;s unit.    Only applicable when the **unit** parameter is set to &#x60;UNSPECIFIED&#x60;. | [optional] 
**EntityId** | **string** | Restricts the metric usage to the specified service.    This field is mutually exclusive with the **managementZones** field. | [optional] 
**ManagementZones** | **[]string** | Restricts the metric usage to specified management zones.    This field is mutually exclusive with the **entityId** field. | [optional] 
**Conditions** | [**[]Condition**](Condition.md) | The set of conditions for the metric usage.    **All** the specified conditions must be fulfilled to use the metric. | [optional] 
**DimensionDefinition** | [**DimensionDefinition**](DimensionDefinition.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


