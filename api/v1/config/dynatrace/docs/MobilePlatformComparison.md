# MobilePlatformComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator of the comparison.   You can reverse it by setting **negate** to &#x60;true&#x60;.   Example: EQUALS, EXISTS   Example operators might not necessarily be available for the selected type | 
**Value** | **string** | The value to compare to. | [optional] 
**Negate** | **bool** | Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**. | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING&#x60; -&gt; StringComparison  * &#x60;INDEXED_NAME&#x60; -&gt; IndexedNameComparison  * &#x60;INDEXED_STRING&#x60; -&gt; IndexedStringComparison  * &#x60;INTEGER&#x60; -&gt; IntegerComparison  * &#x60;SERVICE_TYPE&#x60; -&gt; ServiceTypeComparison  * &#x60;PAAS_TYPE&#x60; -&gt; PaasTypeComparison  * &#x60;CLOUD_TYPE&#x60; -&gt; CloudTypeComparison  * &#x60;AZURE_SKU&#x60; -&gt; AzureSkuComparision  * &#x60;AZURE_COMPUTE_MODE&#x60; -&gt; AzureComputeModeComparison  * &#x60;ENTITY_ID&#x60; -&gt; EntityIdComparison  * &#x60;SIMPLE_TECH&#x60; -&gt; SimpleTechComparison  * &#x60;SIMPLE_HOST_TECH&#x60; -&gt; SimpleHostTechComparison  * &#x60;SERVICE_TOPOLOGY&#x60; -&gt; ServiceTopologyComparison  * &#x60;DATABASE_TOPOLOGY&#x60; -&gt; DatabaseTopologyComparison  * &#x60;OS_TYPE&#x60; -&gt; OsTypeComparison  * &#x60;HYPERVISOR_TYPE&#x60; -&gt; HypervisorTypeComparision  * &#x60;IP_ADDRESS&#x60; -&gt; IpAddressComparison  * &#x60;OS_ARCHITECTURE&#x60; -&gt; OsArchitectureComparison  * &#x60;BITNESS&#x60; -&gt; BitnessComparision  * &#x60;APPLICATION_TYPE&#x60; -&gt; ApplicationTypeComparison  * &#x60;MOBILE_PLATFORM&#x60; -&gt; MobilePlatformComparison  * &#x60;CUSTOM_APPLICATION_TYPE&#x60; -&gt; CustomApplicationTypeComparison  * &#x60;DCRUM_DECODER_TYPE&#x60; -&gt; DcrumDecoderComparison  * &#x60;SYNTHETIC_ENGINE_TYPE&#x60; -&gt; SyntheticEngineTypeComparison  * &#x60;TAG&#x60; -&gt; TagComparison  * &#x60;INDEXED_TAG&#x60; -&gt; IndexedTagComparison   | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


