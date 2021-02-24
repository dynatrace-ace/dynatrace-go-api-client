# ValueCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator comparing the extracted value to the comparison value. | 
**Negate** | **bool** | Negate the comparison. | 
**Value** | **string** | The value to compare to. | 

## Methods

### NewValueCondition

`func NewValueCondition(operator string, negate bool, value string, ) *ValueCondition`

NewValueCondition instantiates a new ValueCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueConditionWithDefaults

`func NewValueConditionWithDefaults() *ValueCondition`

NewValueConditionWithDefaults instantiates a new ValueCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ValueCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ValueCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ValueCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetNegate

`func (o *ValueCondition) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ValueCondition) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ValueCondition) SetNegate(v bool)`

SetNegate sets Negate field to given value.


### GetValue

`func (o *ValueCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueCondition) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


