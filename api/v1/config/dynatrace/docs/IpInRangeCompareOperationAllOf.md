# IpInRangeCompareOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negate** | Pointer to **bool** | Inverts the operation of the condition. Set to &#x60;true&#x60; to turn **IP is in range** into **IP is not in range**.    If not set, then &#x60;false&#x60; is used. | [optional] 
**Lower** | Pointer to **string** | The lower boundary of the IP range. | [optional] 
**Upper** | Pointer to **string** | The upper boundary of the IP range. | [optional] 

## Methods

### NewIpInRangeCompareOperationAllOf

`func NewIpInRangeCompareOperationAllOf() *IpInRangeCompareOperationAllOf`

NewIpInRangeCompareOperationAllOf instantiates a new IpInRangeCompareOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpInRangeCompareOperationAllOfWithDefaults

`func NewIpInRangeCompareOperationAllOfWithDefaults() *IpInRangeCompareOperationAllOf`

NewIpInRangeCompareOperationAllOfWithDefaults instantiates a new IpInRangeCompareOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegate

`func (o *IpInRangeCompareOperationAllOf) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *IpInRangeCompareOperationAllOf) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *IpInRangeCompareOperationAllOf) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *IpInRangeCompareOperationAllOf) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetLower

`func (o *IpInRangeCompareOperationAllOf) GetLower() string`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *IpInRangeCompareOperationAllOf) GetLowerOk() (*string, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *IpInRangeCompareOperationAllOf) SetLower(v string)`

SetLower sets Lower field to given value.

### HasLower

`func (o *IpInRangeCompareOperationAllOf) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetUpper

`func (o *IpInRangeCompareOperationAllOf) GetUpper() string`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *IpInRangeCompareOperationAllOf) GetUpperOk() (*string, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *IpInRangeCompareOperationAllOf) SetUpper(v string)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *IpInRangeCompareOperationAllOf) HasUpper() bool`

HasUpper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


