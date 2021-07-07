# RelatedContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **string** | The image ID of the related container image. | [optional] [readonly] 
**ImageName** | Pointer to **string** | The image name of the related container image. | [optional] [readonly] 
**NumberOfAffectedEntities** | Pointer to **int32** | The number of affected entities. | [optional] [readonly] 
**AffectedEntities** | Pointer to **[]string** | A list of affected entities. | [optional] [readonly] 

## Methods

### NewRelatedContainerImage

`func NewRelatedContainerImage() *RelatedContainerImage`

NewRelatedContainerImage instantiates a new RelatedContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedContainerImageWithDefaults

`func NewRelatedContainerImageWithDefaults() *RelatedContainerImage`

NewRelatedContainerImageWithDefaults instantiates a new RelatedContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *RelatedContainerImage) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *RelatedContainerImage) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *RelatedContainerImage) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *RelatedContainerImage) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageName

`func (o *RelatedContainerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *RelatedContainerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *RelatedContainerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *RelatedContainerImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetNumberOfAffectedEntities

`func (o *RelatedContainerImage) GetNumberOfAffectedEntities() int32`

GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field if non-nil, zero value otherwise.

### GetNumberOfAffectedEntitiesOk

`func (o *RelatedContainerImage) GetNumberOfAffectedEntitiesOk() (*int32, bool)`

GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAffectedEntities

`func (o *RelatedContainerImage) SetNumberOfAffectedEntities(v int32)`

SetNumberOfAffectedEntities sets NumberOfAffectedEntities field to given value.

### HasNumberOfAffectedEntities

`func (o *RelatedContainerImage) HasNumberOfAffectedEntities() bool`

HasNumberOfAffectedEntities returns a boolean if a field has been set.

### GetAffectedEntities

`func (o *RelatedContainerImage) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RelatedContainerImage) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RelatedContainerImage) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *RelatedContainerImage) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


