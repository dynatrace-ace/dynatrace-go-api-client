# FullWebServiceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [optional] 
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ManagementZones** | **[]string** | Specifies the management zones of the process group for which this service detection rule should be created. | [optional] 
**Id** | **string** | The ID of the service detection rule. | [optional] 
**Order** | **string** | The order of the rule in the rules list.    The rules are evaluated from top to bottom. The first matching rule applies. | [optional] 
**Name** | **string** | The name of the rule. | 
**Description** | **string** | A short description of the rule. | [optional] 
**Enabled** | **bool** | The rule is enabled(&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Conditions** | [**[]ConditionsFullWebServiceAttributeTypeDto**](ConditionsFullWebServiceAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**DetectAsWebRequestService** | **bool** | Detect the matching requests as full web services (&#x60;false&#x60;) or web request services (&#x60;true&#x60;).   Setting this field to &#x60;true&#x60; prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate rule of the &#x60;FULL_WEB_REQUEST&#x60; type.   Default is &#x60;false&#x60;, detecting matching requests as full web services. | [optional] 
**WebServiceName** | [**WebServiceName**](WebServiceName.md) |  | [optional] 
**WebServiceNameSpace** | [**WebServiceNameSpace**](WebServiceNameSpace.md) |  | [optional] 
**ApplicationId** | [**ApplicationId**](ApplicationId.md) |  | [optional] 
**ContextRoot** | [**ContextRoot**](ContextRoot.md) |  | [optional] 
**ServerName** | [**ServerName**](ServerName.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


