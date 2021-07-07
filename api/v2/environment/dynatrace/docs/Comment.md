# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorName** | Pointer to **string** | The user who wrote the comment. | [optional] 
**CreatedAtTimestamp** | **int64** | The timestamp of comment creation, in UTC milliseconds. | 
**Context** | Pointer to **string** | The context of the comment. | [optional] 
**Id** | Pointer to **string** | The ID of the comment. | [optional] 
**Content** | Pointer to **string** | The text of the comment. | [optional] 

## Methods

### NewComment

`func NewComment(createdAtTimestamp int64, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorName

`func (o *Comment) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Comment) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Comment) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *Comment) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetCreatedAtTimestamp

`func (o *Comment) GetCreatedAtTimestamp() int64`

GetCreatedAtTimestamp returns the CreatedAtTimestamp field if non-nil, zero value otherwise.

### GetCreatedAtTimestampOk

`func (o *Comment) GetCreatedAtTimestampOk() (*int64, bool)`

GetCreatedAtTimestampOk returns a tuple with the CreatedAtTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimestamp

`func (o *Comment) SetCreatedAtTimestamp(v int64)`

SetCreatedAtTimestamp sets CreatedAtTimestamp field to given value.


### GetContext

`func (o *Comment) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Comment) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Comment) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Comment) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *Comment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Comment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Comment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Comment) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


