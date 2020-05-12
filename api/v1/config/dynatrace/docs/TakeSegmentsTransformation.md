# TakeSegmentsTransformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BEFORE&#x60; -&gt; BeforeTransformation  * &#x60;AFTER&#x60; -&gt; AfterTransformation  * &#x60;BETWEEN&#x60; -&gt; BetweenTransformation  * &#x60;REPLACE_BETWEEN&#x60; -&gt; ReplaceBetweenTransformation  * &#x60;REMOVE_NUMBERS&#x60; -&gt; RemoveNumbersTransformation  * &#x60;REMOVE_CREDIT_CARDS&#x60; -&gt; RemoveCreditCardNumbersTransformation  * &#x60;REMOVE_IBANS&#x60; -&gt; RemoveIBANsTransformation  * &#x60;REMOVE_IPS&#x60; -&gt; RemoveIPsTransformation  * &#x60;SPLIT_SELECT&#x60; -&gt; SplitSelectTransformation  * &#x60;TAKE_SEGMENTS&#x60; -&gt; TakeSegmentsTransformation   | 
**SegmentCount** | **int32** | The number of elements to be kept. | 
**Delimiter** | **string** | The delimiter for splitting the detected value. The delimiter itself is not kept. | 
**TakeFromEnd** | **bool** | Keeps the first (&#x60;false&#x60;) or last (&#x60;true&#x60;) elements.    If not set, then &#x60;false&#x60; is used, keeping the first elements. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


