# LogSourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathDefinitions** | [**[]PathDefinition**](PathDefinition.md) | A list of filtering criteria for log path.   If several criteria are specified, the OR logic applies. | [optional] 
**SourceEntities** | **[]string** | A list of Dynatrace entities, where the log can originate from. Specify Dynatrace entity IDs here.    Allowed types of entities are &#x60;PROCESS_GROUP&#x60; and &#x60;PROCESS_GROUP_INSTANCE&#x60;. You can&#39;t mix entity types.   If several entities are specified, the OR logic applies.   This field is mutually exclusive with the **osTypes** field. | [optional] 
**HostFilters** | **[]string** | A list of hosts, where the log can originate from. Specify Dynatrace entity IDs here.   If several hosts are specified, the OR logic applies. | [optional] 
**OsTypes** | **[]string** | A list of operating system types, where the log can originate from.   If set, **only OS logs** are included in the result.   If several OS are specified, the OR logic applies.   This field is mutually exclusive with the **sourceEntities** field. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


