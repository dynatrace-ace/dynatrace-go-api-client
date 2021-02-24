# TagCompareOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareKeyOnly** | Pointer to **bool** | If &#x60;true&#x60; ignores the tag values and only validates that the tag key is matching. Defaults to &#x60;false&#x60;. | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | The value to compare to.   If several values are specified, the OR logic applies. | [optional] 

## Methods

### NewTagCompareOperation

`func NewTagCompareOperation() *TagCompareOperation`

NewTagCompareOperation instantiates a new TagCompareOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCompareOperationWithDefaults

`func NewTagCompareOperationWithDefaults() *TagCompareOperation`

NewTagCompareOperationWithDefaults instantiates a new TagCompareOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareKeyOnly

`func (o *TagCompareOperation) GetCompareKeyOnly() bool`

GetCompareKeyOnly returns the CompareKeyOnly field if non-nil, zero value otherwise.

### GetCompareKeyOnlyOk

`func (o *TagCompareOperation) GetCompareKeyOnlyOk() (*bool, bool)`

GetCompareKeyOnlyOk returns a tuple with the CompareKeyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareKeyOnly

`func (o *TagCompareOperation) SetCompareKeyOnly(v bool)`

SetCompareKeyOnly sets CompareKeyOnly field to given value.

### HasCompareKeyOnly

`func (o *TagCompareOperation) HasCompareKeyOnly() bool`

HasCompareKeyOnly returns a boolean if a field has been set.

### GetTags

`func (o *TagCompareOperation) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagCompareOperation) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagCompareOperation) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagCompareOperation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


