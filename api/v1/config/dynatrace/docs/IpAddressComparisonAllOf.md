# IpAddressComparisonAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Operator of the comparison. You can reverse it by setting **negate** to &#x60;true&#x60;.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need. | [optional] 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or insensitive (&#x60;false&#x60;). | [optional] 

## Methods

### NewIpAddressComparisonAllOf

`func NewIpAddressComparisonAllOf() *IpAddressComparisonAllOf`

NewIpAddressComparisonAllOf instantiates a new IpAddressComparisonAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressComparisonAllOfWithDefaults

`func NewIpAddressComparisonAllOfWithDefaults() *IpAddressComparisonAllOf`

NewIpAddressComparisonAllOfWithDefaults instantiates a new IpAddressComparisonAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *IpAddressComparisonAllOf) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IpAddressComparisonAllOf) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IpAddressComparisonAllOf) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *IpAddressComparisonAllOf) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *IpAddressComparisonAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpAddressComparisonAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpAddressComparisonAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpAddressComparisonAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *IpAddressComparisonAllOf) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *IpAddressComparisonAllOf) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *IpAddressComparisonAllOf) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *IpAddressComparisonAllOf) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


