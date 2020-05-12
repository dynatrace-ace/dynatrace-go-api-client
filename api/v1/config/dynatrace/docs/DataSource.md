# DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The data source is enbled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Source** | **string** | The source of the attribute to capture. Works in conjunction with **parameterName** or **methods** and **technology**. | 
**ValueProcessing** | [**ValueProcessing**](ValueProcessing.md) |  | [optional] 
**Technology** | **string** | The technology of the method to capture if the **source** value is &#x60;METHOD_PARAM&#x60;. \\n\\n Not applicable in other cases. | [optional] 
**SessionAttributeTechnology** | **string** | The technology of the session attribute to capture if the **source** value is &#x60;SESSION_ATTRIBUTE&#x60;. \\n\\n Not applicable in other cases. | [optional] 
**Methods** | [**[]CapturedMethod**](CapturedMethod.md) | The method specification if the **source** value is &#x60;METHOD_PARAM&#x60;.    Not applicable in other cases. | [optional] 
**ParameterName** | **string** | The name of the web request parameter to capture.   Required if the **source** is one of the following: &#x60;POST_PARAMETER&#x60;, &#x60;GET_PARAMETER&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;, &#x60;CUSTOM_ATTRIBUTE&#x60;.   Not applicable in other cases. | [optional] 
**Scope** | [**ScopeConditions**](ScopeConditions.md) |  | [optional] 
**CapturingAndStorageLocation** | **string** | Specifies the location where the values are captured and stored.   Required if the **source** is one of the following: &#x60;GET_PARAMETER&#x60;, &#x60;URI&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;.    Not applicable in other cases.    If the **source** value is &#x60;REQUEST_HEADER&#x60; or &#x60;RESPONSE_HEADER&#x60;, the &#x60;CAPTURE_AND_STORE_ON_BOTH&#x60; location is not allowed. | [optional] 
**IibNodeType** | **string** | The IBM integration bus node type for which the value is captured.   This or &#x60;iibMethodNodeCondition&#x60; is required if the **source** is: &#x60;IIB_NODE&#x60;.   Not applicable in other cases. | [optional] 
**IibMethodNodeCondition** | [**ValueCondition**](ValueCondition.md) |  | [optional] 
**CicsSDKMethodNodeCondition** | [**ValueCondition**](ValueCondition.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


