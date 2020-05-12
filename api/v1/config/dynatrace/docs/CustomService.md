# CustomService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the custom service. | [optional] 
**Name** | **string** | The name of the custom service, displayed in the UI. | 
**Order** | **string** | The order string. Sorting custom services alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses. | [optional] 
**Enabled** | **bool** | Custom service enabled/disabled. | 
**Rules** | [**[]DetectionRule**](DetectionRule.md) | The list of rules defining the custom service. | 
**QueueEntryPoint** | **bool** | The queue entry point flag.   Set to &#x60;true&#x60; for custom messaging services. | 
**QueueEntryPointType** | **string** | The queue entry point type.. | [optional] 
**ProcessGroups** | **[]string** | The list of process groups the custom service should belong to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


