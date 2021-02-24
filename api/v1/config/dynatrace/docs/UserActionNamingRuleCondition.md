# UserActionNamingRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operand1** | **string** | Must be a defined placeholder wrapped in curly braces | 
**Operand2** | Pointer to **string** | Must be null if operator is \&quot;IS_EMPTY\&quot;, a regex if operator is \&quot;MATCHES_REGULAR_ERPRESSION\&quot;. In all other cases the value can be a freetext or a placeholder wrapped in curly braces | [optional] 
**Operator** | **string** | The operator of the condition | 

## Methods

### NewUserActionNamingRuleCondition

`func NewUserActionNamingRuleCondition(operand1 string, operator string, ) *UserActionNamingRuleCondition`

NewUserActionNamingRuleCondition instantiates a new UserActionNamingRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingRuleConditionWithDefaults

`func NewUserActionNamingRuleConditionWithDefaults() *UserActionNamingRuleCondition`

NewUserActionNamingRuleConditionWithDefaults instantiates a new UserActionNamingRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperand1

`func (o *UserActionNamingRuleCondition) GetOperand1() string`

GetOperand1 returns the Operand1 field if non-nil, zero value otherwise.

### GetOperand1Ok

`func (o *UserActionNamingRuleCondition) GetOperand1Ok() (*string, bool)`

GetOperand1Ok returns a tuple with the Operand1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand1

`func (o *UserActionNamingRuleCondition) SetOperand1(v string)`

SetOperand1 sets Operand1 field to given value.


### GetOperand2

`func (o *UserActionNamingRuleCondition) GetOperand2() string`

GetOperand2 returns the Operand2 field if non-nil, zero value otherwise.

### GetOperand2Ok

`func (o *UserActionNamingRuleCondition) GetOperand2Ok() (*string, bool)`

GetOperand2Ok returns a tuple with the Operand2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand2

`func (o *UserActionNamingRuleCondition) SetOperand2(v string)`

SetOperand2 sets Operand2 field to given value.

### HasOperand2

`func (o *UserActionNamingRuleCondition) HasOperand2() bool`

HasOperand2 returns a boolean if a field has been set.

### GetOperator

`func (o *UserActionNamingRuleCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UserActionNamingRuleCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UserActionNamingRuleCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


