# ApplicationDetectionRuleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadataDtoImpl**](ConfigurationMetadataDtoImpl.md) |  | [optional] 
**Id** | **string** | The ID of the rule. | [optional] 
**Order** | **string** | The order of the rule in the rules list.   The rules are evaluated from top to bottom. The first matching rule applies. | [optional] 
**ApplicationIdentifier** | **string** | The Dynatrace entity ID of the application, for example &#x60;APPLICATION-4A3B43&#x60;.    You must use an existing ID. If you need to create a rule for an application that doesn&#39;t exist yet, [create an application first](https://www.dynatrace.com/support/help/shortlink/api-config-web-app-post-web-app) and then configure detection rules for it. | 
**FilterConfig** | [**ApplicationFilter**](ApplicationFilter.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


