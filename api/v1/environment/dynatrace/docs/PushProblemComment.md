# PushProblemComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | A comment on the problem. | 
**User** | **string** | The author of the comment. | 
**Context** | Pointer to **string** | The context of the comment. It can contain any additional information. | [optional] 

## Methods

### NewPushProblemComment

`func NewPushProblemComment(comment string, user string, ) *PushProblemComment`

NewPushProblemComment instantiates a new PushProblemComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushProblemCommentWithDefaults

`func NewPushProblemCommentWithDefaults() *PushProblemComment`

NewPushProblemCommentWithDefaults instantiates a new PushProblemComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PushProblemComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PushProblemComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PushProblemComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetUser

`func (o *PushProblemComment) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PushProblemComment) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PushProblemComment) SetUser(v string)`

SetUser sets User field to given value.


### GetContext

`func (o *PushProblemComment) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PushProblemComment) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PushProblemComment) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PushProblemComment) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


