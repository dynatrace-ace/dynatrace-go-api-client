# ZosComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 

## Methods

### NewZosComparisonInfo

`func NewZosComparisonInfo(comparison string, ) *ZosComparisonInfo`

NewZosComparisonInfo instantiates a new ZosComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZosComparisonInfoWithDefaults

`func NewZosComparisonInfoWithDefaults() *ZosComparisonInfo`

NewZosComparisonInfoWithDefaults instantiates a new ZosComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *ZosComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *ZosComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *ZosComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *ZosComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ZosComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ZosComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ZosComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


