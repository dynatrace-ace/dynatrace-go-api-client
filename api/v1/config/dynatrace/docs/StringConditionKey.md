# StringConditionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The attribute to be used for comparision. | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;PROCESS_CUSTOM_METADATA_KEY&#x60; -&gt; CustomProcessMetadataConditionKey  * &#x60;HOST_CUSTOM_METADATA_KEY&#x60; -&gt; CustomHostMetadataConditionKey  * &#x60;PROCESS_PREDEFINED_METADATA_KEY&#x60; -&gt; ProcessMetadataConditionKey  * &#x60;STRING&#x60; -&gt; StringConditionKey   | [optional] 
**DynamicKey** | **string** | The key of the attribute, which need dynamic keys.   Not applicable otherwise, as the attibute itself acts as a key. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


