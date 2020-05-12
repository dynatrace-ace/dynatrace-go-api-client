# TagCompareOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EQUALS&#x60; -&gt; EqualsCompareOperation  * &#x60;STRING_CONTAINS&#x60; -&gt; StringContainsCompareOperation  * &#x60;STARTS_WITH&#x60; -&gt; StartsWithCompareOperation  * &#x60;ENDS_WITH&#x60; -&gt; EndsWithCompareOperation  * &#x60;EXISTS&#x60; -&gt; ExistsCompareOperation  * &#x60;IP_IN_RANGE&#x60; -&gt; IpInRangeCompareOperation  * &#x60;LESS_THAN&#x60; -&gt; LessThanCompareOperation  * &#x60;GREATER_THAN&#x60; -&gt; GreaterThanCompareOperation  * &#x60;INT_EQUALS&#x60; -&gt; IntEqualsCompareOperation  * &#x60;STRING_EQUALS&#x60; -&gt; StringEqualsCompareOperation  * &#x60;TAG&#x60; -&gt; TagCompareOperation   | 
**CompareKeyOnly** | **bool** | If &#x60;true&#x60; ignores the tag values and only validates that the tag key is matching. Defaults to &#x60;false&#x60;. | [optional] 
**Tags** | [**[]TagInfo**](TagInfo.md) | The value to compare to.   If several values are specified, the OR logic applies. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


