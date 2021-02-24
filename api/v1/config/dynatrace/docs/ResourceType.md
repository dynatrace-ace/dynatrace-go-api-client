# ResourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegularExpression** | **string** | The regular expression to detect the resource. | 
**PrimaryResourceType** | **string** | The primary type of the resource. | 
**SecondaryResourceType** | Pointer to **string** | The secondary type of the resource. | [optional] 

## Methods

### NewResourceType

`func NewResourceType(regularExpression string, primaryResourceType string, ) *ResourceType`

NewResourceType instantiates a new ResourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeWithDefaults

`func NewResourceTypeWithDefaults() *ResourceType`

NewResourceTypeWithDefaults instantiates a new ResourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegularExpression

`func (o *ResourceType) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *ResourceType) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *ResourceType) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.


### GetPrimaryResourceType

`func (o *ResourceType) GetPrimaryResourceType() string`

GetPrimaryResourceType returns the PrimaryResourceType field if non-nil, zero value otherwise.

### GetPrimaryResourceTypeOk

`func (o *ResourceType) GetPrimaryResourceTypeOk() (*string, bool)`

GetPrimaryResourceTypeOk returns a tuple with the PrimaryResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryResourceType

`func (o *ResourceType) SetPrimaryResourceType(v string)`

SetPrimaryResourceType sets PrimaryResourceType field to given value.


### GetSecondaryResourceType

`func (o *ResourceType) GetSecondaryResourceType() string`

GetSecondaryResourceType returns the SecondaryResourceType field if non-nil, zero value otherwise.

### GetSecondaryResourceTypeOk

`func (o *ResourceType) GetSecondaryResourceTypeOk() (*string, bool)`

GetSecondaryResourceTypeOk returns a tuple with the SecondaryResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryResourceType

`func (o *ResourceType) SetSecondaryResourceType(v string)`

SetSecondaryResourceType sets SecondaryResourceType field to given value.

### HasSecondaryResourceType

`func (o *ResourceType) HasSecondaryResourceType() bool`

HasSecondaryResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


