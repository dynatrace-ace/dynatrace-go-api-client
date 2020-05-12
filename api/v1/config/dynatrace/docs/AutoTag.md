# AutoTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the auto-tag. | [optional] 
**Name** | **string** | The name of the auto-tag, which is applied to entities.   Additionally you can specify a **valueFormat** in the tag rule. In that case the tag is used in the &#x60;name:valueFormat&#x60; format.   For example you can extend the &#x60;Infrastructure&#x60; tag to &#x60;Infrastructure:Windows&#x60; and &#x60;Infrastructure:Linux&#x60;. | 
**Rules** | [**[]AutoTagRule**](AutoTagRule.md) | The list of rules for tag usage.   When there are multiple rules, the OR logic applies. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


