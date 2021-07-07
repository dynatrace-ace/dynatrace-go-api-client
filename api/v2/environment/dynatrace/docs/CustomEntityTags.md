# CustomEntityTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** | The total number of tags in the response. | [optional] 
**Tags** | [**[]METag**](METag.md) | A list of custom tags. | 

## Methods

### NewCustomEntityTags

`func NewCustomEntityTags(tags []METag, ) *CustomEntityTags`

NewCustomEntityTags instantiates a new CustomEntityTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityTagsWithDefaults

`func NewCustomEntityTagsWithDefaults() *CustomEntityTags`

NewCustomEntityTagsWithDefaults instantiates a new CustomEntityTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CustomEntityTags) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CustomEntityTags) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CustomEntityTags) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CustomEntityTags) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTags

`func (o *CustomEntityTags) GetTags() []METag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CustomEntityTags) GetTagsOk() (*[]METag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CustomEntityTags) SetTags(v []METag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


