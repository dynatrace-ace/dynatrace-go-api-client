# ExistsCompareOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EQUALS&#x60; -&gt; EqualsCompareOperation  * &#x60;STRING_CONTAINS&#x60; -&gt; StringContainsCompareOperation  * &#x60;STARTS_WITH&#x60; -&gt; StartsWithCompareOperation  * &#x60;ENDS_WITH&#x60; -&gt; EndsWithCompareOperation  * &#x60;EXISTS&#x60; -&gt; ExistsCompareOperation  * &#x60;IP_IN_RANGE&#x60; -&gt; IpInRangeCompareOperation  * &#x60;LESS_THAN&#x60; -&gt; LessThanCompareOperation  * &#x60;GREATER_THAN&#x60; -&gt; GreaterThanCompareOperation  * &#x60;INT_EQUALS&#x60; -&gt; IntEqualsCompareOperation  * &#x60;STRING_EQUALS&#x60; -&gt; StringEqualsCompareOperation  * &#x60;TAG&#x60; -&gt; TagCompareOperation   | 
**Negate** | **bool** | Inverts the operation of the condition. Set to &#x60;true&#x60; to turn **exists** into **does not exist**.    If not set, then &#x60;false&#x60; is used. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


