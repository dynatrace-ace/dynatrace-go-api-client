# TagMatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeTypes** | **[]string** | The list of types of the Dynatrace entities (for example hosts or services) you want to pick up by matching. | 
**Tags** | [**[]TagInfo**](TagInfo.md) | The list of tags you want to use for matching. At least one tag is required.    You can use custom tags from the UI, imported tags, and tags based on environment variables. | 

## Methods

### NewTagMatchRule

`func NewTagMatchRule(meTypes []string, tags []TagInfo, ) *TagMatchRule`

NewTagMatchRule instantiates a new TagMatchRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagMatchRuleWithDefaults

`func NewTagMatchRuleWithDefaults() *TagMatchRule`

NewTagMatchRuleWithDefaults instantiates a new TagMatchRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeTypes

`func (o *TagMatchRule) GetMeTypes() []string`

GetMeTypes returns the MeTypes field if non-nil, zero value otherwise.

### GetMeTypesOk

`func (o *TagMatchRule) GetMeTypesOk() (*[]string, bool)`

GetMeTypesOk returns a tuple with the MeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeTypes

`func (o *TagMatchRule) SetMeTypes(v []string)`

SetMeTypes sets MeTypes field to given value.


### GetTags

`func (o *TagMatchRule) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagMatchRule) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagMatchRule) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


