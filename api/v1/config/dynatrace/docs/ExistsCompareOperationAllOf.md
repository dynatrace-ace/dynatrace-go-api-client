# ExistsCompareOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Negate** | Pointer to **bool** | Inverts the operation of the condition. Set to &#x60;true&#x60; to turn **exists** into **does not exist**.    If not set, then &#x60;false&#x60; is used. | [optional] 

## Methods

### NewExistsCompareOperationAllOf

`func NewExistsCompareOperationAllOf() *ExistsCompareOperationAllOf`

NewExistsCompareOperationAllOf instantiates a new ExistsCompareOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExistsCompareOperationAllOfWithDefaults

`func NewExistsCompareOperationAllOfWithDefaults() *ExistsCompareOperationAllOf`

NewExistsCompareOperationAllOfWithDefaults instantiates a new ExistsCompareOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNegate

`func (o *ExistsCompareOperationAllOf) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ExistsCompareOperationAllOf) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ExistsCompareOperationAllOf) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *ExistsCompareOperationAllOf) HasNegate() bool`

HasNegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


