# UserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User id | 
**NodeId** | **int32** | Node on which user session exists | 
**SessionId** | **string** | User session id | 
**CreationTime** | **int64** | User session creation timestamp | 
**LastAccessedTimestamp** | Pointer to **int64** | Timestamp when session was recently accessed | [optional] 
**TenantUuid** | **string** | UUID of tenant to which user has logged in (or cluster UUID if user has logged in to CMC) | 
**LoginType** | **string** | The way user has logged in | 
**Device** | Pointer to **string** | Device on which user has logged in | [optional] 
**Ip** | Pointer to **string** | IP from which has logged in | [optional] 

## Methods

### NewUserSession

`func NewUserSession(userId string, nodeId int32, sessionId string, creationTime int64, tenantUuid string, loginType string, ) *UserSession`

NewUserSession instantiates a new UserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionWithDefaults

`func NewUserSessionWithDefaults() *UserSession`

NewUserSessionWithDefaults instantiates a new UserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserSession) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSession) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSession) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNodeId

`func (o *UserSession) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *UserSession) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *UserSession) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.


### GetSessionId

`func (o *UserSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetCreationTime

`func (o *UserSession) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *UserSession) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *UserSession) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.


### GetLastAccessedTimestamp

`func (o *UserSession) GetLastAccessedTimestamp() int64`

GetLastAccessedTimestamp returns the LastAccessedTimestamp field if non-nil, zero value otherwise.

### GetLastAccessedTimestampOk

`func (o *UserSession) GetLastAccessedTimestampOk() (*int64, bool)`

GetLastAccessedTimestampOk returns a tuple with the LastAccessedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedTimestamp

`func (o *UserSession) SetLastAccessedTimestamp(v int64)`

SetLastAccessedTimestamp sets LastAccessedTimestamp field to given value.

### HasLastAccessedTimestamp

`func (o *UserSession) HasLastAccessedTimestamp() bool`

HasLastAccessedTimestamp returns a boolean if a field has been set.

### GetTenantUuid

`func (o *UserSession) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *UserSession) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *UserSession) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.


### GetLoginType

`func (o *UserSession) GetLoginType() string`

GetLoginType returns the LoginType field if non-nil, zero value otherwise.

### GetLoginTypeOk

`func (o *UserSession) GetLoginTypeOk() (*string, bool)`

GetLoginTypeOk returns a tuple with the LoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginType

`func (o *UserSession) SetLoginType(v string)`

SetLoginType sets LoginType field to given value.


### GetDevice

`func (o *UserSession) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserSession) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserSession) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIp

`func (o *UserSession) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *UserSession) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *UserSession) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *UserSession) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


