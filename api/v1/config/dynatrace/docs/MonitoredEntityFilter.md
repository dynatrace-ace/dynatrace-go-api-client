# MonitoredEntityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching. | [optional] 
**MzId** | **string** | The ID of a management zone to which the matched entities must belong. | [optional] 
**Tags** | [**[]TagInfo**](TagInfo.md) | The tag you want to use for matching.   You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables. | 
**TagCombination** | **string** | The logic that applies when several tags are specified: AND/OR.   If not set, the OR logic is used. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


