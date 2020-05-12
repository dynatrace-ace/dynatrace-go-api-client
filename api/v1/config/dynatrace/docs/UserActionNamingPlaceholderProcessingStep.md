# UserActionNamingPlaceholderProcessingStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | An action to be taken by the processing:   * &#x60;SUBSTRING&#x60;: Extracts the string between **patternBefore** and **patternAfter**.  * &#x60;REPLACEMENT&#x60;: Replaces the string between **patternBefore** and **patternAfter** with the specified **replacement**. * &#x60;REPLACE_WITH_PATTERN&#x60;: Replaces the **patternToReplace** with the specified **replacement**.  * &#x60;EXTRACT_BY_REGULAR_EXPRESSION&#x60;: Extracts the part of the string that matches the **regularExpression**.  * &#x60;REPLACE_WITH_REGULAR_EXPRESSION&#x60;: Replaces all occurrences that match **regularExpression** with the specified **replacement**.  * &#x60;REPLACE_IDS&#x60;: Replaces all IDs and UUIDs with the specified **replacement**. | 
**PatternBefore** | **string** | The pattern before the required value. It will be removed. | [optional] 
**PatternBeforeSearchType** | **string** | The required occurrence of **patternBefore**. | [optional] 
**PatternAfter** | **string** | The pattern after the required value. It will be removed. | [optional] 
**PatternAfterSearchType** | **string** | The required occurrence of **patternAfter**. | [optional] 
**Replacement** | **string** | Replacement for the original value. | [optional] 
**PatternToReplace** | **string** | The pattern to be replaced.    Only applicable if the **type** is &#x60;REPLACE_WITH_PATTERN&#x60;. | [optional] 
**RegularExpression** | **string** | A regular expression for the string to be extracted or replaced.    Only applicable if the **type** is &#x60;EXTRACT_BY_REGULAR_EXPRESSION&#x60; or &#x60;REPLACE_WITH_REGULAR_EXPRESSION&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


