# AlertingCustomTextFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The filter is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Value** | **string** | The value to compare to. | 
**Operator** | **string** | Operator of the comparison.    You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Negate** | **bool** | Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**. | 
**CaseInsensitive** | **bool** | The condition is case sensitive (&#x60;false&#x60;) or case insensitive (&#x60;true&#x60;).    If not set, then &#x60;false&#x60; is used, making the condition case sensitive. | 

## Methods

### NewAlertingCustomTextFilter

`func NewAlertingCustomTextFilter(enabled bool, value string, operator string, negate bool, caseInsensitive bool, ) *AlertingCustomTextFilter`

NewAlertingCustomTextFilter instantiates a new AlertingCustomTextFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingCustomTextFilterWithDefaults

`func NewAlertingCustomTextFilterWithDefaults() *AlertingCustomTextFilter`

NewAlertingCustomTextFilterWithDefaults instantiates a new AlertingCustomTextFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AlertingCustomTextFilter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertingCustomTextFilter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertingCustomTextFilter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValue

`func (o *AlertingCustomTextFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AlertingCustomTextFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AlertingCustomTextFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *AlertingCustomTextFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AlertingCustomTextFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AlertingCustomTextFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetNegate

`func (o *AlertingCustomTextFilter) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AlertingCustomTextFilter) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AlertingCustomTextFilter) SetNegate(v bool)`

SetNegate sets Negate field to given value.


### GetCaseInsensitive

`func (o *AlertingCustomTextFilter) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *AlertingCustomTextFilter) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *AlertingCustomTextFilter) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


