# ConditionsOpaqueAndExternalWebRequestAttributeTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | The type of the attribute to be checked. | 
**CompareOperations** | Pointer to [**[]CompareOperation**](CompareOperation.md) | A list of conditions for the rule.   If several conditions are specified, the AND logic applies. | [optional] 

## Methods

### NewConditionsOpaqueAndExternalWebRequestAttributeTypeDto

`func NewConditionsOpaqueAndExternalWebRequestAttributeTypeDto(attributeType string, ) *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto`

NewConditionsOpaqueAndExternalWebRequestAttributeTypeDto instantiates a new ConditionsOpaqueAndExternalWebRequestAttributeTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionsOpaqueAndExternalWebRequestAttributeTypeDtoWithDefaults

`func NewConditionsOpaqueAndExternalWebRequestAttributeTypeDtoWithDefaults() *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto`

NewConditionsOpaqueAndExternalWebRequestAttributeTypeDtoWithDefaults instantiates a new ConditionsOpaqueAndExternalWebRequestAttributeTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetCompareOperations

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetCompareOperations() []CompareOperation`

GetCompareOperations returns the CompareOperations field if non-nil, zero value otherwise.

### GetCompareOperationsOk

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetCompareOperationsOk() (*[]CompareOperation, bool)`

GetCompareOperationsOk returns a tuple with the CompareOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareOperations

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) SetCompareOperations(v []CompareOperation)`

SetCompareOperations sets CompareOperations field to given value.

### HasCompareOperations

`func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) HasCompareOperations() bool`

HasCompareOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


