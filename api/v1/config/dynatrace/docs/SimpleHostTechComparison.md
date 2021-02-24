# SimpleHostTechComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator of the comparison. You can reverse it by setting **negate** to &#x60;true&#x60;.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need. | 
**Value** | Pointer to [**SimpleHostTech**](SimpleHostTech.md) |  | [optional] 

## Methods

### NewSimpleHostTechComparison

`func NewSimpleHostTechComparison(operator string, ) *SimpleHostTechComparison`

NewSimpleHostTechComparison instantiates a new SimpleHostTechComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleHostTechComparisonWithDefaults

`func NewSimpleHostTechComparisonWithDefaults() *SimpleHostTechComparison`

NewSimpleHostTechComparisonWithDefaults instantiates a new SimpleHostTechComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SimpleHostTechComparison) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SimpleHostTechComparison) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SimpleHostTechComparison) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *SimpleHostTechComparison) GetValue() SimpleHostTech`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SimpleHostTechComparison) GetValueOk() (*SimpleHostTech, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SimpleHostTechComparison) SetValue(v SimpleHostTech)`

SetValue sets Value field to given value.

### HasValue

`func (o *SimpleHostTechComparison) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


