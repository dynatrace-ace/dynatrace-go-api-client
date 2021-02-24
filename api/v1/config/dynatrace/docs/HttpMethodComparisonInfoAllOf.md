# HttpMethodComparisonInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | [optional] 
**Value** | Pointer to **string** | The value to compare to. | [optional] 

## Methods

### NewHttpMethodComparisonInfoAllOf

`func NewHttpMethodComparisonInfoAllOf() *HttpMethodComparisonInfoAllOf`

NewHttpMethodComparisonInfoAllOf instantiates a new HttpMethodComparisonInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpMethodComparisonInfoAllOfWithDefaults

`func NewHttpMethodComparisonInfoAllOfWithDefaults() *HttpMethodComparisonInfoAllOf`

NewHttpMethodComparisonInfoAllOfWithDefaults instantiates a new HttpMethodComparisonInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *HttpMethodComparisonInfoAllOf) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *HttpMethodComparisonInfoAllOf) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *HttpMethodComparisonInfoAllOf) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *HttpMethodComparisonInfoAllOf) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetValue

`func (o *HttpMethodComparisonInfoAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpMethodComparisonInfoAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpMethodComparisonInfoAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HttpMethodComparisonInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


