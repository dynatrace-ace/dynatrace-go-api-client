# ProblemComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the comment. | [optional] 
**CreatedAtTimestamp** | Pointer to **int64** | The timestamp of the comment creation, in UTC milliseconds. | [optional] 
**Content** | Pointer to **string** | The text of the comment. | [optional] 
**UserName** | Pointer to **string** | The author of the comment. | [optional] 
**Context** | Pointer to **string** | The context of the comment.   Could be any textual comment. You can only set it via REST API. | [optional] 

## Methods

### NewProblemComment

`func NewProblemComment() *ProblemComment`

NewProblemComment instantiates a new ProblemComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemCommentWithDefaults

`func NewProblemCommentWithDefaults() *ProblemComment`

NewProblemCommentWithDefaults instantiates a new ProblemComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProblemComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProblemComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProblemComment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProblemComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAtTimestamp

`func (o *ProblemComment) GetCreatedAtTimestamp() int64`

GetCreatedAtTimestamp returns the CreatedAtTimestamp field if non-nil, zero value otherwise.

### GetCreatedAtTimestampOk

`func (o *ProblemComment) GetCreatedAtTimestampOk() (*int64, bool)`

GetCreatedAtTimestampOk returns a tuple with the CreatedAtTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimestamp

`func (o *ProblemComment) SetCreatedAtTimestamp(v int64)`

SetCreatedAtTimestamp sets CreatedAtTimestamp field to given value.

### HasCreatedAtTimestamp

`func (o *ProblemComment) HasCreatedAtTimestamp() bool`

HasCreatedAtTimestamp returns a boolean if a field has been set.

### GetContent

`func (o *ProblemComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProblemComment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProblemComment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ProblemComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetUserName

`func (o *ProblemComment) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ProblemComment) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ProblemComment) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ProblemComment) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetContext

`func (o *ProblemComment) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ProblemComment) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ProblemComment) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ProblemComment) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


