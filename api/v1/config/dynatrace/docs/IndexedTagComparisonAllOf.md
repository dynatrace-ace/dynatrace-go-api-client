# IndexedTagComparisonAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Operator of the comparison. You can reverse it by setting **negate** to &#x60;true&#x60;.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need. | [optional] 
**Value** | Pointer to [**TagInfo**](TagInfo.md) |  | [optional] 

## Methods

### NewIndexedTagComparisonAllOf

`func NewIndexedTagComparisonAllOf() *IndexedTagComparisonAllOf`

NewIndexedTagComparisonAllOf instantiates a new IndexedTagComparisonAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedTagComparisonAllOfWithDefaults

`func NewIndexedTagComparisonAllOfWithDefaults() *IndexedTagComparisonAllOf`

NewIndexedTagComparisonAllOfWithDefaults instantiates a new IndexedTagComparisonAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *IndexedTagComparisonAllOf) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IndexedTagComparisonAllOf) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IndexedTagComparisonAllOf) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *IndexedTagComparisonAllOf) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *IndexedTagComparisonAllOf) GetValue() TagInfo`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IndexedTagComparisonAllOf) GetValueOk() (*TagInfo, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IndexedTagComparisonAllOf) SetValue(v TagInfo)`

SetValue sets Value field to given value.

### HasValue

`func (o *IndexedTagComparisonAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


