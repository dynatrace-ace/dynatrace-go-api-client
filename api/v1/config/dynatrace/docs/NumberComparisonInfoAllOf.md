# NumberComparisonInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | [optional] 
**Value** | Pointer to **float32** | The value to compare to. | [optional] 

## Methods

### NewNumberComparisonInfoAllOf

`func NewNumberComparisonInfoAllOf() *NumberComparisonInfoAllOf`

NewNumberComparisonInfoAllOf instantiates a new NumberComparisonInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberComparisonInfoAllOfWithDefaults

`func NewNumberComparisonInfoAllOfWithDefaults() *NumberComparisonInfoAllOf`

NewNumberComparisonInfoAllOfWithDefaults instantiates a new NumberComparisonInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *NumberComparisonInfoAllOf) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *NumberComparisonInfoAllOf) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *NumberComparisonInfoAllOf) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *NumberComparisonInfoAllOf) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetValue

`func (o *NumberComparisonInfoAllOf) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberComparisonInfoAllOf) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberComparisonInfoAllOf) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *NumberComparisonInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


