# CompareOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EQUALS&#x60; -&gt; EqualsCompareOperation  * &#x60;STRING_CONTAINS&#x60; -&gt; StringContainsCompareOperation  * &#x60;STARTS_WITH&#x60; -&gt; StartsWithCompareOperation  * &#x60;ENDS_WITH&#x60; -&gt; EndsWithCompareOperation  * &#x60;EXISTS&#x60; -&gt; ExistsCompareOperation  * &#x60;IP_IN_RANGE&#x60; -&gt; IpInRangeCompareOperation  * &#x60;LESS_THAN&#x60; -&gt; LessThanCompareOperation  * &#x60;GREATER_THAN&#x60; -&gt; GreaterThanCompareOperation  * &#x60;INT_EQUALS&#x60; -&gt; IntEqualsCompareOperation  * &#x60;STRING_EQUALS&#x60; -&gt; StringEqualsCompareOperation  * &#x60;TAG&#x60; -&gt; TagCompareOperation   | 

## Methods

### NewCompareOperation

`func NewCompareOperation(type_ string, ) *CompareOperation`

NewCompareOperation instantiates a new CompareOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareOperationWithDefaults

`func NewCompareOperationWithDefaults() *CompareOperation`

NewCompareOperationWithDefaults instantiates a new CompareOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompareOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompareOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompareOperation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


