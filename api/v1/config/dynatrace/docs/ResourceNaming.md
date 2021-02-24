# ResourceNaming

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Image** | **[]string** | The list of image extensions that will be renamed to &#x60;Image&#x60;. | 
**Binary** | **[]string** | The list of binary extensions that will be renamed to &#x60;Binary&#x60;. | 

## Methods

### NewResourceNaming

`func NewResourceNaming(image []string, binary []string, ) *ResourceNaming`

NewResourceNaming instantiates a new ResourceNaming object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNamingWithDefaults

`func NewResourceNamingWithDefaults() *ResourceNaming`

NewResourceNamingWithDefaults instantiates a new ResourceNaming object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ResourceNaming) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceNaming) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceNaming) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceNaming) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImage

`func (o *ResourceNaming) GetImage() []string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ResourceNaming) GetImageOk() (*[]string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ResourceNaming) SetImage(v []string)`

SetImage sets Image field to given value.


### GetBinary

`func (o *ResourceNaming) GetBinary() []string`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *ResourceNaming) GetBinaryOk() (*[]string, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *ResourceNaming) SetBinary(v []string)`

SetBinary sets Binary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


