# CreateAccessRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User id | [optional] 
**Reason** | Pointer to **string** | Request reason description | [optional] 
**RequestedDays** | Pointer to **int32** | For how many days access is requested | [optional] 
**Role** | Pointer to **string** | Requested role | [optional] 

## Methods

### NewCreateAccessRequestDto

`func NewCreateAccessRequestDto() *CreateAccessRequestDto`

NewCreateAccessRequestDto instantiates a new CreateAccessRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessRequestDtoWithDefaults

`func NewCreateAccessRequestDtoWithDefaults() *CreateAccessRequestDto`

NewCreateAccessRequestDtoWithDefaults instantiates a new CreateAccessRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateAccessRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAccessRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAccessRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateAccessRequestDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReason

`func (o *CreateAccessRequestDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateAccessRequestDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateAccessRequestDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateAccessRequestDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestedDays

`func (o *CreateAccessRequestDto) GetRequestedDays() int32`

GetRequestedDays returns the RequestedDays field if non-nil, zero value otherwise.

### GetRequestedDaysOk

`func (o *CreateAccessRequestDto) GetRequestedDaysOk() (*int32, bool)`

GetRequestedDaysOk returns a tuple with the RequestedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDays

`func (o *CreateAccessRequestDto) SetRequestedDays(v int32)`

SetRequestedDays sets RequestedDays field to given value.

### HasRequestedDays

`func (o *CreateAccessRequestDto) HasRequestedDays() bool`

HasRequestedDays returns a boolean if a field has been set.

### GetRole

`func (o *CreateAccessRequestDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateAccessRequestDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateAccessRequestDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateAccessRequestDto) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


