# AlertingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the alerting profile. | [optional] 
**DisplayName** | **string** | The name of the alerting profile, displayed in the UI. | 
**Rules** | [**[]AlertingProfileSeverityRule**](AlertingProfileSeverityRule.md) | A list of severity rules.    The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops.   If you specify both severity rule and event filter, the AND logic applies. | [optional] 
**MzId** | **string** | The ID of the management zone to which the alerting profile applies. | [optional] 
**EventTypeFilters** | [**[]AlertingEventTypeFilter**](AlertingEventTypeFilter.md) | The list of event filters.   For all filters that are *negated* inside of these event filters, that is all \&quot;Predefined\&quot; as well as \&quot;Custom\&quot; (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.   If you specify both severity rule and event filter, the AND logic applies. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


