# UserSessionSyntheticEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the synthetic event. | [optional] 
**SyntheticEventId** | Pointer to **string** | The Dynatrace entity ID for the synthetic event. | [optional] 
**SequenceNumber** | Pointer to **int32** | The sequence number of the synthetic event in scope of the complete browser monitor. | [optional] 
**Timestamp** | Pointer to **int64** | The timestamp when the synthetic event was simulated, in UTC milliseconds. | [optional] 
**Type** | Pointer to **string** | The type of the synthetic event. For example click or keystroke. | [optional] 
**ErrorCode** | Pointer to **int32** | The error code of the error that occurred during this event. | [optional] 
**ErrorName** | Pointer to **string** | Description of the error that occurred during this event. | [optional] 

## Methods

### NewUserSessionSyntheticEvent

`func NewUserSessionSyntheticEvent() *UserSessionSyntheticEvent`

NewUserSessionSyntheticEvent instantiates a new UserSessionSyntheticEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionSyntheticEventWithDefaults

`func NewUserSessionSyntheticEventWithDefaults() *UserSessionSyntheticEvent`

NewUserSessionSyntheticEventWithDefaults instantiates a new UserSessionSyntheticEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserSessionSyntheticEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSessionSyntheticEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSessionSyntheticEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSessionSyntheticEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSyntheticEventId

`func (o *UserSessionSyntheticEvent) GetSyntheticEventId() string`

GetSyntheticEventId returns the SyntheticEventId field if non-nil, zero value otherwise.

### GetSyntheticEventIdOk

`func (o *UserSessionSyntheticEvent) GetSyntheticEventIdOk() (*string, bool)`

GetSyntheticEventIdOk returns a tuple with the SyntheticEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEventId

`func (o *UserSessionSyntheticEvent) SetSyntheticEventId(v string)`

SetSyntheticEventId sets SyntheticEventId field to given value.

### HasSyntheticEventId

`func (o *UserSessionSyntheticEvent) HasSyntheticEventId() bool`

HasSyntheticEventId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *UserSessionSyntheticEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *UserSessionSyntheticEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *UserSessionSyntheticEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *UserSessionSyntheticEvent) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetTimestamp

`func (o *UserSessionSyntheticEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserSessionSyntheticEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserSessionSyntheticEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserSessionSyntheticEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *UserSessionSyntheticEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionSyntheticEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionSyntheticEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionSyntheticEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrorCode

`func (o *UserSessionSyntheticEvent) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *UserSessionSyntheticEvent) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *UserSessionSyntheticEvent) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *UserSessionSyntheticEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorName

`func (o *UserSessionSyntheticEvent) GetErrorName() string`

GetErrorName returns the ErrorName field if non-nil, zero value otherwise.

### GetErrorNameOk

`func (o *UserSessionSyntheticEvent) GetErrorNameOk() (*string, bool)`

GetErrorNameOk returns a tuple with the ErrorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorName

`func (o *UserSessionSyntheticEvent) SetErrorName(v string)`

SetErrorName sets ErrorName field to given value.

### HasErrorName

`func (o *UserSessionSyntheticEvent) HasErrorName() bool`

HasErrorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


