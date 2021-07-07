# AuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | **string** | The ID of the log entry. | [readonly] 
**EventType** | **string** | The type of the recorded operation. | [readonly] 
**Category** | **string** | The category of the recorded operation. | [readonly] 
**EntityId** | Pointer to **string** | The ID of an entity from the **category**.   For example, it can be config ID for the &#x60;CONFIG&#x60; category or token ID for the &#x60;TOKEN&#x60; category. | [optional] [readonly] 
**EnvironmentId** | **string** | The ID of the Dynatrace environment where the recorded operation occurred. | [readonly] 
**User** | **string** | The ID of the user who performed the recorded operation. | [readonly] 
**UserType** | **string** | The type of the authentication of the **user**. | [readonly] 
**UserOrigin** | Pointer to **string** | The origin and the IP address of the **user**. | [optional] [readonly] 
**Timestamp** | **int64** | The timestamp of the record creation, in UTC milliseconds. | [readonly] 
**Success** | **bool** | The recorded operation is successful (&#x60;true&#x60;) or failed (&#x60;false&#x60;). | [readonly] 
**Message** | Pointer to **string** | The logged message. | [optional] [readonly] 
**Patch** | Pointer to **map[string]interface{}** | The patch of the recorded operation as the JSON representation.   The format is an enhanced RFC 6902. The patch also carries the previous value in the **oldValue** field. | [optional] [readonly] 

## Methods

### NewAuditLogEntry

`func NewAuditLogEntry(logId string, eventType string, category string, environmentId string, user string, userType string, timestamp int64, success bool, ) *AuditLogEntry`

NewAuditLogEntry instantiates a new AuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryWithDefaults

`func NewAuditLogEntryWithDefaults() *AuditLogEntry`

NewAuditLogEntryWithDefaults instantiates a new AuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *AuditLogEntry) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *AuditLogEntry) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *AuditLogEntry) SetLogId(v string)`

SetLogId sets LogId field to given value.


### GetEventType

`func (o *AuditLogEntry) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AuditLogEntry) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AuditLogEntry) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetCategory

`func (o *AuditLogEntry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditLogEntry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditLogEntry) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetEntityId

`func (o *AuditLogEntry) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AuditLogEntry) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AuditLogEntry) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AuditLogEntry) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *AuditLogEntry) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *AuditLogEntry) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *AuditLogEntry) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetUser

`func (o *AuditLogEntry) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditLogEntry) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditLogEntry) SetUser(v string)`

SetUser sets User field to given value.


### GetUserType

`func (o *AuditLogEntry) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *AuditLogEntry) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *AuditLogEntry) SetUserType(v string)`

SetUserType sets UserType field to given value.


### GetUserOrigin

`func (o *AuditLogEntry) GetUserOrigin() string`

GetUserOrigin returns the UserOrigin field if non-nil, zero value otherwise.

### GetUserOriginOk

`func (o *AuditLogEntry) GetUserOriginOk() (*string, bool)`

GetUserOriginOk returns a tuple with the UserOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrigin

`func (o *AuditLogEntry) SetUserOrigin(v string)`

SetUserOrigin sets UserOrigin field to given value.

### HasUserOrigin

`func (o *AuditLogEntry) HasUserOrigin() bool`

HasUserOrigin returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuditLogEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetSuccess

`func (o *AuditLogEntry) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AuditLogEntry) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AuditLogEntry) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *AuditLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPatch

`func (o *AuditLogEntry) GetPatch() map[string]interface{}`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *AuditLogEntry) GetPatchOk() (*map[string]interface{}, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *AuditLogEntry) SetPatch(v map[string]interface{})`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *AuditLogEntry) HasPatch() bool`

HasPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


