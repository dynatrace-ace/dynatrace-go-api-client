# StringComparisonInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | [optional] 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 

## Methods

### NewStringComparisonInfoAllOf

`func NewStringComparisonInfoAllOf() *StringComparisonInfoAllOf`

NewStringComparisonInfoAllOf instantiates a new StringComparisonInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringComparisonInfoAllOfWithDefaults

`func NewStringComparisonInfoAllOfWithDefaults() *StringComparisonInfoAllOf`

NewStringComparisonInfoAllOfWithDefaults instantiates a new StringComparisonInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *StringComparisonInfoAllOf) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *StringComparisonInfoAllOf) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *StringComparisonInfoAllOf) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *StringComparisonInfoAllOf) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetValue

`func (o *StringComparisonInfoAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringComparisonInfoAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringComparisonInfoAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringComparisonInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *StringComparisonInfoAllOf) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *StringComparisonInfoAllOf) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *StringComparisonInfoAllOf) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *StringComparisonInfoAllOf) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


