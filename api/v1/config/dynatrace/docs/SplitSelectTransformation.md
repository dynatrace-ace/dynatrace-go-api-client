# SplitSelectTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BEFORE&#x60; -&gt; BeforeTransformation  * &#x60;AFTER&#x60; -&gt; AfterTransformation  * &#x60;BETWEEN&#x60; -&gt; BetweenTransformation  * &#x60;REPLACE_BETWEEN&#x60; -&gt; ReplaceBetweenTransformation  * &#x60;REMOVE_NUMBERS&#x60; -&gt; RemoveNumbersTransformation  * &#x60;REMOVE_CREDIT_CARDS&#x60; -&gt; RemoveCreditCardNumbersTransformation  * &#x60;REMOVE_IBANS&#x60; -&gt; RemoveIBANsTransformation  * &#x60;REMOVE_IPS&#x60; -&gt; RemoveIPsTransformation  * &#x60;SPLIT_SELECT&#x60; -&gt; SplitSelectTransformation  * &#x60;TAKE_SEGMENTS&#x60; -&gt; TakeSegmentsTransformation   | 
**Delimiter** | **string** | The delimiter for splitting the detected value. The delimiter itself is not kept. | 
**ItemIndex** | **int32** | The index of the element in the split array to be used. Indexing starts with &#x60;1&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


