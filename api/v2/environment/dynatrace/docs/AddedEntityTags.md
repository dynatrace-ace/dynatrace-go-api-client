# AddedEntityTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedTags** | Pointer to [**[]METag**](METag.md) | A list of added custom tags. | [optional] 
**MatchedEntitiesCount** | Pointer to **int64** | The number of monitored entities where the tags have been added. | [optional] 

## Methods

### NewAddedEntityTags

`func NewAddedEntityTags() *AddedEntityTags`

NewAddedEntityTags instantiates a new AddedEntityTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedEntityTagsWithDefaults

`func NewAddedEntityTagsWithDefaults() *AddedEntityTags`

NewAddedEntityTagsWithDefaults instantiates a new AddedEntityTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedTags

`func (o *AddedEntityTags) GetAppliedTags() []METag`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *AddedEntityTags) GetAppliedTagsOk() (*[]METag, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *AddedEntityTags) SetAppliedTags(v []METag)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *AddedEntityTags) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### GetMatchedEntitiesCount

`func (o *AddedEntityTags) GetMatchedEntitiesCount() int64`

GetMatchedEntitiesCount returns the MatchedEntitiesCount field if non-nil, zero value otherwise.

### GetMatchedEntitiesCountOk

`func (o *AddedEntityTags) GetMatchedEntitiesCountOk() (*int64, bool)`

GetMatchedEntitiesCountOk returns a tuple with the MatchedEntitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedEntitiesCount

`func (o *AddedEntityTags) SetMatchedEntitiesCount(v int64)`

SetMatchedEntitiesCount sets MatchedEntitiesCount field to given value.

### HasMatchedEntitiesCount

`func (o *AddedEntityTags) HasMatchedEntitiesCount() bool`

HasMatchedEntitiesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


