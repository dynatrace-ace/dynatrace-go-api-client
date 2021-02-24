# ConditionsFullWebRequestAttributeTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | The type of the attribute to be checked. | 
**CompareOperations** | Pointer to [**[]CompareOperation**](CompareOperation.md) | A list of conditions for the rule.   If several conditions are specified, the AND logic applies. | [optional] 

## Methods

### NewConditionsFullWebRequestAttributeTypeDto

`func NewConditionsFullWebRequestAttributeTypeDto(attributeType string, ) *ConditionsFullWebRequestAttributeTypeDto`

NewConditionsFullWebRequestAttributeTypeDto instantiates a new ConditionsFullWebRequestAttributeTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsFullWebRequestAttributeTypeDtoWithDefaults

`func NewConditionsFullWebRequestAttributeTypeDtoWithDefaults() *ConditionsFullWebRequestAttributeTypeDto`

NewConditionsFullWebRequestAttributeTypeDtoWithDefaults instantiates a new ConditionsFullWebRequestAttributeTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ConditionsFullWebRequestAttributeTypeDto) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetCompareOperations

`func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperations() []CompareOperation`

GetCompareOperations returns the CompareOperations field if non-nil, zero value otherwise.

### GetCompareOperationsOk

`func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperationsOk() (*[]CompareOperation, bool)`

GetCompareOperationsOk returns a tuple with the CompareOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareOperations

`func (o *ConditionsFullWebRequestAttributeTypeDto) SetCompareOperations(v []CompareOperation)`

SetCompareOperations sets CompareOperations field to given value.

### HasCompareOperations

`func (o *ConditionsFullWebRequestAttributeTypeDto) HasCompareOperations() bool`

HasCompareOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


