# TagComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to [**TagInfo**](TagInfo.md) |  | [optional] 

## Methods

### NewTagComparisonInfo

`func NewTagComparisonInfo(comparison string, ) *TagComparisonInfo`

NewTagComparisonInfo instantiates a new TagComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagComparisonInfoWithDefaults

`func NewTagComparisonInfoWithDefaults() *TagComparisonInfo`

NewTagComparisonInfoWithDefaults instantiates a new TagComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *TagComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *TagComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *TagComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *TagComparisonInfo) GetValue() TagInfo`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagComparisonInfo) GetValueOk() (*TagInfo, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagComparisonInfo) SetValue(v TagInfo)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


