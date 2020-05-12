# ManagementZoneRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of Dynatrace entities the management zone can be applied to. | 
**Enabled** | **bool** | The evaluation rule is enabled(&#x60;true&#x60;) or disabled(&#x60;false&#x60;). | 
**PropagationTypes** | **[]string** | How to apply the management zone to underlying entities:   &#x60;SERVICE_TO_HOST_LIKE&#x60;: Apply to underlying hosts of matching services.  &#x60;SERVICE_TO_PROCESS_GROUP_LIKE&#x60;: Apply to underlying process groups of matching services.   &#x60;PROCESS_GROUP_TO_HOST&#x60;: Apply to underlying hosts of matching process groups.  &#x60;PROCESS_GROUP_TO_SERVICE&#x60;: Apply to all services provided by matching process groups.   &#x60;HOST_TO_PROCESS_GROUP_INSTANCE&#x60;: Apply to processes running on matching hosts.   &#x60;CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE&#x60;: Apply to custom devices in matching custom device groups. | [optional] 
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | Matching rules of the management zone. The management zone applies only if all conditions are fulfilled. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


