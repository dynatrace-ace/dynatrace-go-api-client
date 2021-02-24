# StringComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 

## Methods

### NewStringComparisonInfo

`func NewStringComparisonInfo(comparison string, ) *StringComparisonInfo`

NewStringComparisonInfo instantiates a new StringComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringComparisonInfoWithDefaults

`func NewStringComparisonInfoWithDefaults() *StringComparisonInfo`

NewStringComparisonInfoWithDefaults instantiates a new StringComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *StringComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *StringComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *StringComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *StringComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *StringComparisonInfo) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *StringComparisonInfo) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *StringComparisonInfo) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *StringComparisonInfo) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


