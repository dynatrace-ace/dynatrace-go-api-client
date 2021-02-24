# NumberRequestAttributeComparisonInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | [optional] 
**Value** | Pointer to **float32** | The value to compare to. | [optional] 
**RequestAttribute** | Pointer to **string** |  | [optional] 
**MatchOnChildCalls** | Pointer to **bool** | If &#x60;true&#x60;, the request attribute is matched on child service calls.     Default is &#x60;false&#x60;. | [optional] 
**Source** | Pointer to [**PropagationSource**](PropagationSource.md) |  | [optional] 

## Methods

### NewNumberRequestAttributeComparisonInfoAllOf

`func NewNumberRequestAttributeComparisonInfoAllOf() *NumberRequestAttributeComparisonInfoAllOf`

NewNumberRequestAttributeComparisonInfoAllOf instantiates a new NumberRequestAttributeComparisonInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberRequestAttributeComparisonInfoAllOfWithDefaults

`func NewNumberRequestAttributeComparisonInfoAllOfWithDefaults() *NumberRequestAttributeComparisonInfoAllOf`

NewNumberRequestAttributeComparisonInfoAllOfWithDefaults instantiates a new NumberRequestAttributeComparisonInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *NumberRequestAttributeComparisonInfoAllOf) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *NumberRequestAttributeComparisonInfoAllOf) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetValue

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberRequestAttributeComparisonInfoAllOf) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *NumberRequestAttributeComparisonInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequestAttribute

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *NumberRequestAttributeComparisonInfoAllOf) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.

### HasRequestAttribute

`func (o *NumberRequestAttributeComparisonInfoAllOf) HasRequestAttribute() bool`

HasRequestAttribute returns a boolean if a field has been set.

### GetMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetMatchOnChildCalls() bool`

GetMatchOnChildCalls returns the MatchOnChildCalls field if non-nil, zero value otherwise.

### GetMatchOnChildCallsOk

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetMatchOnChildCallsOk() (*bool, bool)`

GetMatchOnChildCallsOk returns a tuple with the MatchOnChildCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfoAllOf) SetMatchOnChildCalls(v bool)`

SetMatchOnChildCalls sets MatchOnChildCalls field to given value.

### HasMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfoAllOf) HasMatchOnChildCalls() bool`

HasMatchOnChildCalls returns a boolean if a field has been set.

### GetSource

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetSource() PropagationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NumberRequestAttributeComparisonInfoAllOf) GetSourceOk() (*PropagationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NumberRequestAttributeComparisonInfoAllOf) SetSource(v PropagationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *NumberRequestAttributeComparisonInfoAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


