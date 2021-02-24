# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The attribute to be matched. | 
**ComparisonInfo** | [**ComparisonInfo**](ComparisonInfo.md) |  | 

## Methods

### NewCondition

`func NewCondition(attribute string, comparisonInfo ComparisonInfo, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *Condition) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Condition) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Condition) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetComparisonInfo

`func (o *Condition) GetComparisonInfo() ComparisonInfo`

GetComparisonInfo returns the ComparisonInfo field if non-nil, zero value otherwise.

### GetComparisonInfoOk

`func (o *Condition) GetComparisonInfoOk() (*ComparisonInfo, bool)`

GetComparisonInfoOk returns a tuple with the ComparisonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonInfo

`func (o *Condition) SetComparisonInfo(v ComparisonInfo)`

SetComparisonInfo sets ComparisonInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


