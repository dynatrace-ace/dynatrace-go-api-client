# FullWebRequestRule

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
**Conditions** | [**[]ConditionsFullWebRequestAttributeTypeDto**](ConditionsFullWebRequestAttributeTypeDto.md) | A list of conditions of the rule.   If several conditions are specified, the AND logic applies. | [optional] 
**ApplicationId** | [**ApplicationId**](ApplicationId.md) |  | [optional] 
**ContextRoot** | [**ContextRoot**](ContextRoot.md) |  | [optional] 
**ServerName** | [**ServerName**](ServerName.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


