# ConditionalNamingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the auto tag. | [optional] 
**Order** | **string** | The order string. Sorting naming rules alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses. | [optional] 
**Type** | **string** | The type of entities to which the naming applies. | 
**NameFormat** | **string** | The format of the applied name. Can contain placeholders for attributes of entities using this format ´{Entity:Attribute}´. | 
**DisplayName** | **string** | The name of the naming rule | 
**Enabled** | **bool** | Naming rule enabled/disabled. | 
**Rules** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | Matching rules of the conditional naming. The naming applies only if all conditions are fulfilled. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


