# IntEqualsCompareOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negate** | Pointer to **bool** | Inverts the operation of the condition. Set to &#x60;true&#x60; to turn **equals** into **does not equal**.    If not set, then &#x60;false&#x60; is used. | [optional] 
**Values** | Pointer to **[]int32** | The value to compare to.   If several values are specified, the OR logic applies. | [optional] 

## Methods

### NewIntEqualsCompareOperation

`func NewIntEqualsCompareOperation() *IntEqualsCompareOperation`

NewIntEqualsCompareOperation instantiates a new IntEqualsCompareOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntEqualsCompareOperationWithDefaults

`func NewIntEqualsCompareOperationWithDefaults() *IntEqualsCompareOperation`

NewIntEqualsCompareOperationWithDefaults instantiates a new IntEqualsCompareOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegate

`func (o *IntEqualsCompareOperation) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *IntEqualsCompareOperation) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *IntEqualsCompareOperation) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *IntEqualsCompareOperation) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetValues

`func (o *IntEqualsCompareOperation) GetValues() []int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IntEqualsCompareOperation) GetValuesOk() (*[]int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IntEqualsCompareOperation) SetValues(v []int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *IntEqualsCompareOperation) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


