# MuteState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | The user who has muted or unmuted the problem. | [optional] [readonly] 
**Reason** | Pointer to **string** | The reason for the mute state change. | [optional] [readonly] 
**Comment** | Pointer to **string** | A comment by the user. | [optional] [readonly] 

## Methods

### NewMuteState

`func NewMuteState() *MuteState`

NewMuteState instantiates a new MuteState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteStateWithDefaults

`func NewMuteStateWithDefaults() *MuteState`

NewMuteStateWithDefaults instantiates a new MuteState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *MuteState) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MuteState) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MuteState) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MuteState) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReason

`func (o *MuteState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MuteState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MuteState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MuteState) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *MuteState) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MuteState) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MuteState) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MuteState) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


