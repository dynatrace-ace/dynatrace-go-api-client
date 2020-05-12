# StringRequestAttributeComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | **string** | The value to compare to. | [optional] 
**Negate** | **bool** | Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**. | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING&#x60; -&gt; StringComparisonInfo  * &#x60;NUMBER&#x60; -&gt; NumberComparisonInfo  * &#x60;BOOLEAN&#x60; -&gt; BooleanComparisonInfo  * &#x60;HTTP_METHOD&#x60; -&gt; HttpMethodComparisonInfo  * &#x60;STRING_REQUEST_ATTRIBUTE&#x60; -&gt; StringRequestAttributeComparisonInfo  * &#x60;NUMBER_REQUEST_ATTRIBUTE&#x60; -&gt; NumberRequestAttributeComparisonInfo  * &#x60;ZOS_CALL_TYPE&#x60; -&gt; ZosComparisonInfo  * &#x60;IIB_INPUT_NODE_TYPE&#x60; -&gt; IIBInputNodeTypeComparisonInfo  * &#x60;ESB_INPUT_NODE_TYPE&#x60; -&gt; ESBInputNodeTypeComparisonInfo  * &#x60;FAILED_STATE&#x60; -&gt; FailedStateComparisonInfo  * &#x60;FLAW_STATE&#x60; -&gt; FlawStateComparisonInfo  * &#x60;FAILURE_REASON&#x60; -&gt; FailureReasonComparisonInfo  * &#x60;HTTP_STATUS_CLASS&#x60; -&gt; HttpStatusClassComparisonInfo  * &#x60;TAG&#x60; -&gt; TagComparisonInfo  * &#x60;FAST_STRING&#x60; -&gt; FastStringComparisonInfo  * &#x60;SERVICE_TYPE&#x60; -&gt; ServiceTypeComparisonInfo   | 
**RequestAttribute** | **string** |  | 
**CaseSensitive** | **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 
**MatchOnChildCalls** | **bool** | If &#x60;true&#x60;, the request attribute is matched on child service calls.    Default is &#x60;false&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


