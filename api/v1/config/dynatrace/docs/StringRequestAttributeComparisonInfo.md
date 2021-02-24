# StringRequestAttributeComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**RequestAttribute** | **string** |  | 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 
**MatchOnChildCalls** | Pointer to **bool** | If &#x60;true&#x60;, the request attribute is matched on child service calls.    Default is &#x60;false&#x60;. | [optional] 
**Source** | Pointer to [**PropagationSource**](PropagationSource.md) |  | [optional] 

## Methods

### NewStringRequestAttributeComparisonInfo

`func NewStringRequestAttributeComparisonInfo(comparison string, requestAttribute string, ) *StringRequestAttributeComparisonInfo`

NewStringRequestAttributeComparisonInfo instantiates a new StringRequestAttributeComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringRequestAttributeComparisonInfoWithDefaults

`func NewStringRequestAttributeComparisonInfoWithDefaults() *StringRequestAttributeComparisonInfo`

NewStringRequestAttributeComparisonInfoWithDefaults instantiates a new StringRequestAttributeComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *StringRequestAttributeComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *StringRequestAttributeComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *StringRequestAttributeComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *StringRequestAttributeComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringRequestAttributeComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringRequestAttributeComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringRequestAttributeComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequestAttribute

`func (o *StringRequestAttributeComparisonInfo) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *StringRequestAttributeComparisonInfo) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *StringRequestAttributeComparisonInfo) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.


### GetCaseSensitive

`func (o *StringRequestAttributeComparisonInfo) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *StringRequestAttributeComparisonInfo) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *StringRequestAttributeComparisonInfo) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *StringRequestAttributeComparisonInfo) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfo) GetMatchOnChildCalls() bool`

GetMatchOnChildCalls returns the MatchOnChildCalls field if non-nil, zero value otherwise.

### GetMatchOnChildCallsOk

`func (o *StringRequestAttributeComparisonInfo) GetMatchOnChildCallsOk() (*bool, bool)`

GetMatchOnChildCallsOk returns a tuple with the MatchOnChildCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfo) SetMatchOnChildCalls(v bool)`

SetMatchOnChildCalls sets MatchOnChildCalls field to given value.

### HasMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfo) HasMatchOnChildCalls() bool`

HasMatchOnChildCalls returns a boolean if a field has been set.

### GetSource

`func (o *StringRequestAttributeComparisonInfo) GetSource() PropagationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StringRequestAttributeComparisonInfo) GetSourceOk() (*PropagationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StringRequestAttributeComparisonInfo) SetSource(v PropagationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *StringRequestAttributeComparisonInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


