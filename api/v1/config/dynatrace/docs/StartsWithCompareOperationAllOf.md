# StartsWithCompareOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negate** | Pointer to **bool** | Inverts the operation of the condition. Set to &#x60;true&#x60; to turn **starts with** into **does not start with**.    If not set, then &#x60;false&#x60; is used. | [optional] 
**IgnoreCase** | Pointer to **bool** | The condition is case sensitive (&#x60;false&#x60;) or case insensitive (&#x60;true&#x60;).   If not set, then &#x60;false&#x60; is used, making the condition case sensitive. | [optional] 
**Values** | Pointer to **[]string** | The value to compare to.   If several values are specified, the OR logic applies. | [optional] 

## Methods

### NewStartsWithCompareOperationAllOf

`func NewStartsWithCompareOperationAllOf() *StartsWithCompareOperationAllOf`

NewStartsWithCompareOperationAllOf instantiates a new StartsWithCompareOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartsWithCompareOperationAllOfWithDefaults

`func NewStartsWithCompareOperationAllOfWithDefaults() *StartsWithCompareOperationAllOf`

NewStartsWithCompareOperationAllOfWithDefaults instantiates a new StartsWithCompareOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegate

`func (o *StartsWithCompareOperationAllOf) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *StartsWithCompareOperationAllOf) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *StartsWithCompareOperationAllOf) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *StartsWithCompareOperationAllOf) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *StartsWithCompareOperationAllOf) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *StartsWithCompareOperationAllOf) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *StartsWithCompareOperationAllOf) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *StartsWithCompareOperationAllOf) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetValues

`func (o *StartsWithCompareOperationAllOf) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StartsWithCompareOperationAllOf) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StartsWithCompareOperationAllOf) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StartsWithCompareOperationAllOf) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


