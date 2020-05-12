# Placeholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the placeholder. Use it in the naming pattern as &#x60;{name}&#x60;. | 
**Attribute** | **string** | The attribute to extract from. You can only use attributes of the **string** type. | 
**Kind** | **string** | The type of extraction.    Defines either usage of regular expression (&#x60;regex&#x60;) or the position of request attribute value to be extracted.   When the **attribute** is &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute and **aggregation** is &#x60;COUNT&#x60;, needs to be set to &#x60;ORIGINAL_TEXT&#x60; | 
**DelimiterOrRegex** | **string** | Depending on the **type** value:    * &#x60;REGEX_EXTRACTION&#x60;: The regular expression.   * &#x60;BETWEEN_DELIMITER&#x60;: The opening delimiter string to look for.   * All other values: The delimiter string to look for. | [optional] 
**EndDelimiter** | **string** | The closing delimiter string to look for.    Required if the **kind** value is &#x60;BETWEEN_DELIMITER&#x60;. Not applicable otherwise. | [optional] 
**RequestAttribute** | **string** | The request attribute to extract from.    Required if the **kind** value is &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60;. Not applicable otherwise. | [optional] 
**Normalization** | **string** | The format of the extracted string. | [optional] 
**UseFromChildCalls** | **bool** | If &#x60;true&#x60; request attribute will be taken from a child service call.    Only applicable for the &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute. Defaults to &#x60;false&#x60;. | [optional] 
**Aggregation** | **string** | Which value of the request attribute must be used when it occurs across multiple child requests.   Only applicable for the &#x60;SERVICE_REQUEST_ATTRIBUTE&#x60; attribute, when **useFromChildCalls** is &#x60;true&#x60;.   For the &#x60;COUNT&#x60; aggregation, the **kind** field is not applicable. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


