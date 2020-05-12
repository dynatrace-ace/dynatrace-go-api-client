# RequestNaming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the request naming rule. | [optional] 
**Order** | **string** | The order string. Sorting request namings alphabetically by their order string determines their relative ordering.  Typically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise. | [optional] 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**NamingPattern** | **string** | The name to be assigned to matching requests. | 
**ManagementZones** | **[]string** | Specifies the management zones for which this rule should be applied. | [optional] 
**Conditions** | [**[]Condition**](Condition.md) | The set of conditions for the request naming rule usage.    You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger. | 
**Placeholders** | [**[]Placeholder**](Placeholder.md) | The list of custom placeholders to be used in the naming pattern.    It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


