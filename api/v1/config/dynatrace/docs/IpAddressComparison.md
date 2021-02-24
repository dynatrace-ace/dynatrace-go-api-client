# IpAddressComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator of the comparison. You can reverse it by setting **negate** to &#x60;true&#x60;.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or insensitive (&#x60;false&#x60;). | [optional] 

## Methods

### NewIpAddressComparison

`func NewIpAddressComparison(operator string, ) *IpAddressComparison`

NewIpAddressComparison instantiates a new IpAddressComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressComparisonWithDefaults

`func NewIpAddressComparisonWithDefaults() *IpAddressComparison`

NewIpAddressComparisonWithDefaults instantiates a new IpAddressComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *IpAddressComparison) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IpAddressComparison) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IpAddressComparison) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *IpAddressComparison) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IpAddressComparison) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IpAddressComparison) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IpAddressComparison) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *IpAddressComparison) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *IpAddressComparison) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *IpAddressComparison) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *IpAddressComparison) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


