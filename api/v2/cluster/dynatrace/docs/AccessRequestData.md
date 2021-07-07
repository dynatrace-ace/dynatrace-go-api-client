# AccessRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Request id | [optional] 
**UserId** | Pointer to **string** | User id | [optional] 
**Reason** | Pointer to **string** | Request reason description | [optional] 
**RequestedDays** | Pointer to **int32** | For how many days access is requested | [optional] 
**Role** | Pointer to **string** | Requested role | [optional] 
**CreatedTimestamp** | Pointer to **int64** | Access request created at (timestamp) | [optional] 
**ExpirationTimestamp** | Pointer to **int64** | Access expires at (timestamp) | [optional] 
**State** | Pointer to **string** | Access request state | [optional] 
**StateModifiedByUser** | Pointer to **string** | Access request state was modified by user | [optional] 

## Methods

### NewAccessRequestData

`func NewAccessRequestData() *AccessRequestData`

NewAccessRequestData instantiates a new AccessRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestDataWithDefaults

`func NewAccessRequestDataWithDefaults() *AccessRequestData`

NewAccessRequestDataWithDefaults instantiates a new AccessRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AccessRequestData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AccessRequestData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AccessRequestData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AccessRequestData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetUserId

`func (o *AccessRequestData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessRequestData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessRequestData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccessRequestData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReason

`func (o *AccessRequestData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccessRequestData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccessRequestData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccessRequestData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestedDays

`func (o *AccessRequestData) GetRequestedDays() int32`

GetRequestedDays returns the RequestedDays field if non-nil, zero value otherwise.

### GetRequestedDaysOk

`func (o *AccessRequestData) GetRequestedDaysOk() (*int32, bool)`

GetRequestedDaysOk returns a tuple with the RequestedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDays

`func (o *AccessRequestData) SetRequestedDays(v int32)`

SetRequestedDays sets RequestedDays field to given value.

### HasRequestedDays

`func (o *AccessRequestData) HasRequestedDays() bool`

HasRequestedDays returns a boolean if a field has been set.

### GetRole

`func (o *AccessRequestData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccessRequestData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccessRequestData) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccessRequestData) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *AccessRequestData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *AccessRequestData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *AccessRequestData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *AccessRequestData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *AccessRequestData) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *AccessRequestData) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *AccessRequestData) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *AccessRequestData) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetState

`func (o *AccessRequestData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessRequestData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessRequestData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccessRequestData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateModifiedByUser

`func (o *AccessRequestData) GetStateModifiedByUser() string`

GetStateModifiedByUser returns the StateModifiedByUser field if non-nil, zero value otherwise.

### GetStateModifiedByUserOk

`func (o *AccessRequestData) GetStateModifiedByUserOk() (*string, bool)`

GetStateModifiedByUserOk returns a tuple with the StateModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateModifiedByUser

`func (o *AccessRequestData) SetStateModifiedByUser(v string)`

SetStateModifiedByUser sets StateModifiedByUser field to given value.

### HasStateModifiedByUser

`func (o *AccessRequestData) HasStateModifiedByUser() bool`

HasStateModifiedByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


