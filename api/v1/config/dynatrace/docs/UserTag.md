# UserTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueId** | **int32** | uniqueId, unique among all userTags and properties of this application | 
**MetadataId** | Pointer to **int32** | If it&#39;s of type metaData, metaData id of the userTag | [optional] 
**CleanupRule** | Pointer to **string** | Cleanup rule expression of the userTag | [optional] 
**ServerSideRequestAttribute** | Pointer to **string** | requestAttribute Id of the userTag | [optional] 
**IgnoreCase** | Pointer to **bool** | If true, the value of this tag will always be stored in lower case. Defaults to false. | [optional] 

## Methods

### NewUserTag

`func NewUserTag(uniqueId int32, ) *UserTag`

NewUserTag instantiates a new UserTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTagWithDefaults

`func NewUserTagWithDefaults() *UserTag`

NewUserTagWithDefaults instantiates a new UserTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueId

`func (o *UserTag) GetUniqueId() int32`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UserTag) GetUniqueIdOk() (*int32, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UserTag) SetUniqueId(v int32)`

SetUniqueId sets UniqueId field to given value.


### GetMetadataId

`func (o *UserTag) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *UserTag) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *UserTag) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.

### HasMetadataId

`func (o *UserTag) HasMetadataId() bool`

HasMetadataId returns a boolean if a field has been set.

### GetCleanupRule

`func (o *UserTag) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *UserTag) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *UserTag) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *UserTag) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetServerSideRequestAttribute

`func (o *UserTag) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *UserTag) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *UserTag) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *UserTag) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetIgnoreCase

`func (o *UserTag) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *UserTag) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *UserTag) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *UserTag) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


