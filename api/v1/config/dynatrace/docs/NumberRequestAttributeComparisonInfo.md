# NumberRequestAttributeComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **float32** | The value to compare to. | [optional] 
**RequestAttribute** | **string** |  | 
**MatchOnChildCalls** | Pointer to **bool** | If &#x60;true&#x60;, the request attribute is matched on child service calls.     Default is &#x60;false&#x60;. | [optional] 
**Source** | Pointer to [**PropagationSource**](PropagationSource.md) |  | [optional] 

## Methods

### NewNumberRequestAttributeComparisonInfo

`func NewNumberRequestAttributeComparisonInfo(comparison string, requestAttribute string, ) *NumberRequestAttributeComparisonInfo`

NewNumberRequestAttributeComparisonInfo instantiates a new NumberRequestAttributeComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberRequestAttributeComparisonInfoWithDefaults

`func NewNumberRequestAttributeComparisonInfoWithDefaults() *NumberRequestAttributeComparisonInfo`

NewNumberRequestAttributeComparisonInfoWithDefaults instantiates a new NumberRequestAttributeComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *NumberRequestAttributeComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *NumberRequestAttributeComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *NumberRequestAttributeComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *NumberRequestAttributeComparisonInfo) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberRequestAttributeComparisonInfo) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberRequestAttributeComparisonInfo) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *NumberRequestAttributeComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequestAttribute

`func (o *NumberRequestAttributeComparisonInfo) GetRequestAttribute() string`

GetRequestAttribute returns the RequestAttribute field if non-nil, zero value otherwise.

### GetRequestAttributeOk

`func (o *NumberRequestAttributeComparisonInfo) GetRequestAttributeOk() (*string, bool)`

GetRequestAttributeOk returns a tuple with the RequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAttribute

`func (o *NumberRequestAttributeComparisonInfo) SetRequestAttribute(v string)`

SetRequestAttribute sets RequestAttribute field to given value.


### GetMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfo) GetMatchOnChildCalls() bool`

GetMatchOnChildCalls returns the MatchOnChildCalls field if non-nil, zero value otherwise.

### GetMatchOnChildCallsOk

`func (o *NumberRequestAttributeComparisonInfo) GetMatchOnChildCallsOk() (*bool, bool)`

GetMatchOnChildCallsOk returns a tuple with the MatchOnChildCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfo) SetMatchOnChildCalls(v bool)`

SetMatchOnChildCalls sets MatchOnChildCalls field to given value.

### HasMatchOnChildCalls

`func (o *NumberRequestAttributeComparisonInfo) HasMatchOnChildCalls() bool`

HasMatchOnChildCalls returns a boolean if a field has been set.

### GetSource

`func (o *NumberRequestAttributeComparisonInfo) GetSource() PropagationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NumberRequestAttributeComparisonInfo) GetSourceOk() (*PropagationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NumberRequestAttributeComparisonInfo) SetSource(v PropagationSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *NumberRequestAttributeComparisonInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


