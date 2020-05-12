# AlertingProfileTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMode** | **string** | The filtering mode:  * &#x60;INCLUDE_ANY&#x60;: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * &#x60;INCLUDE_ALL&#x60;: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * &#x60;NONE&#x60;: The rule applies to all monitored entities. | 
**TagFilters** | [**[]TagFilter**](TagFilter.md) | A list of required tags. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


