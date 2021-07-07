# CommentRequestDtoImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The text of the comment. | 
**Context** | Pointer to **string** | The context of the comment. | [optional] 

## Methods

### NewCommentRequestDtoImpl

`func NewCommentRequestDtoImpl(message string, ) *CommentRequestDtoImpl`

NewCommentRequestDtoImpl instantiates a new CommentRequestDtoImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRequestDtoImplWithDefaults

`func NewCommentRequestDtoImplWithDefaults() *CommentRequestDtoImpl`

NewCommentRequestDtoImplWithDefaults instantiates a new CommentRequestDtoImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CommentRequestDtoImpl) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommentRequestDtoImpl) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommentRequestDtoImpl) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetContext

`func (o *CommentRequestDtoImpl) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CommentRequestDtoImpl) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CommentRequestDtoImpl) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CommentRequestDtoImpl) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


