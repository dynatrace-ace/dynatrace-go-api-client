# StringRequestAttributeComparisonInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | [optional] 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**RequestAttribute** | Pointer to **string** |  | [optional] 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 
**MatchOnChildCalls** | Pointer to **bool** | If &#x60;true&#x60;, the request attribute is matched on child service calls.    Default is &#x60;false&#x60;. | [optional] 
**Source** | Pointer to [**PropagationSource**](PropagationSource.md) |  | [optional] 

## Methods

### NewStringRequestAttributeComparisonInfoAllOf

`func NewStringRequestAttributeComparisonInfoAllOf() *StringRequestAttributeComparisonInfoAllOf`

NewStringRequestAttributeComparisonInfoAllOf instantiates a new StringRequestAttributeComparisonInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringRequestAttributeComparisonInfoAllOfWithDefaults

`func NewStringRequestAttributeComparisonInfoAllOfWithDefaults() *StringRequestAttributeComparisonInfoAllOf`

NewStringRequestAttributeComparisonInfoAllOfWithDefaults instantiates a new StringRequestAttributeComparisonInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *StringRequestAttributeComparisonInfoAllOf) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *StringRequestAttributeComparisonInfoAllOf) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *StringRequestAttributeComparisonInfoAllOf) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetValue

`func (o *StringRequestAttributeComparisonInfoAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringRequestAttributeComparisonInfoAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringRequestAttributeComparisonInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequestAttribute

`func (o *StringRequestAttributeComparisonInfoAllOf) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *StringRequestAttributeComparisonInfoAllOf) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.

### HasRequestAttribute

`func (o *StringRequestAttributeComparisonInfoAllOf) HasRequestAttribute() bool`

HasRequestAttribute returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *StringRequestAttributeComparisonInfoAllOf) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *StringRequestAttributeComparisonInfoAllOf) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *StringRequestAttributeComparisonInfoAllOf) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfoAllOf) GetMatchOnChildCalls() bool`

GetMatchOnChildCalls returns the MatchOnChildCalls field if non-nil, zero value otherwise.

### GetMatchOnChildCallsOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetMatchOnChildCallsOk() (*bool, bool)`

GetMatchOnChildCallsOk returns a tuple with the MatchOnChildCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfoAllOf) SetMatchOnChildCalls(v bool)`

SetMatchOnChildCalls sets MatchOnChildCalls field to given value.

### HasMatchOnChildCalls

`func (o *StringRequestAttributeComparisonInfoAllOf) HasMatchOnChildCalls() bool`

HasMatchOnChildCalls returns a boolean if a field has been set.

### GetSource

`func (o *StringRequestAttributeComparisonInfoAllOf) GetSource() PropagationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StringRequestAttributeComparisonInfoAllOf) GetSourceOk() (*PropagationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StringRequestAttributeComparisonInfoAllOf) SetSource(v PropagationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *StringRequestAttributeComparisonInfoAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


