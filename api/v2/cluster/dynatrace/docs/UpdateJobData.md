# UpdateJobData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to **[]string** |  | [optional] 
**JobState** | Pointer to **string** |  | [optional] 
**UpdateMethod** | Pointer to **string** |  | [optional] 
**UpdateType** | Pointer to **string** |  | [optional] 
**AgType** | Pointer to **string** |  | [optional] 
**StartVersion** | Pointer to **string** |  | [optional] 
**Cancelable** | Pointer to **bool** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**TargetVersion** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateJobData

`func NewUpdateJobData() *UpdateJobData`

NewUpdateJobData instantiates a new UpdateJobData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJobDataWithDefaults

`func NewUpdateJobDataWithDefaults() *UpdateJobData`

NewUpdateJobDataWithDefaults instantiates a new UpdateJobData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *UpdateJobData) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *UpdateJobData) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *UpdateJobData) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *UpdateJobData) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetJobState

`func (o *UpdateJobData) GetJobState() string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *UpdateJobData) GetJobStateOk() (*string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *UpdateJobData) SetJobState(v string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *UpdateJobData) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetUpdateMethod

`func (o *UpdateJobData) GetUpdateMethod() string`

GetUpdateMethod returns the UpdateMethod field if non-nil, zero value otherwise.

### GetUpdateMethodOk

`func (o *UpdateJobData) GetUpdateMethodOk() (*string, bool)`

GetUpdateMethodOk returns a tuple with the UpdateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMethod

`func (o *UpdateJobData) SetUpdateMethod(v string)`

SetUpdateMethod sets UpdateMethod field to given value.

### HasUpdateMethod

`func (o *UpdateJobData) HasUpdateMethod() bool`

HasUpdateMethod returns a boolean if a field has been set.

### GetUpdateType

`func (o *UpdateJobData) GetUpdateType() string`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *UpdateJobData) GetUpdateTypeOk() (*string, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *UpdateJobData) SetUpdateType(v string)`

SetUpdateType sets UpdateType field to given value.

### HasUpdateType

`func (o *UpdateJobData) HasUpdateType() bool`

HasUpdateType returns a boolean if a field has been set.

### GetAgType

`func (o *UpdateJobData) GetAgType() string`

GetAgType returns the AgType field if non-nil, zero value otherwise.

### GetAgTypeOk

`func (o *UpdateJobData) GetAgTypeOk() (*string, bool)`

GetAgTypeOk returns a tuple with the AgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgType

`func (o *UpdateJobData) SetAgType(v string)`

SetAgType sets AgType field to given value.

### HasAgType

`func (o *UpdateJobData) HasAgType() bool`

HasAgType returns a boolean if a field has been set.

### GetStartVersion

`func (o *UpdateJobData) GetStartVersion() string`

GetStartVersion returns the StartVersion field if non-nil, zero value otherwise.

### GetStartVersionOk

`func (o *UpdateJobData) GetStartVersionOk() (*string, bool)`

GetStartVersionOk returns a tuple with the StartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVersion

`func (o *UpdateJobData) SetStartVersion(v string)`

SetStartVersion sets StartVersion field to given value.

### HasStartVersion

`func (o *UpdateJobData) HasStartVersion() bool`

HasStartVersion returns a boolean if a field has been set.

### GetCancelable

`func (o *UpdateJobData) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *UpdateJobData) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *UpdateJobData) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *UpdateJobData) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetJobId

`func (o *UpdateJobData) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UpdateJobData) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UpdateJobData) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *UpdateJobData) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetTimestamp

`func (o *UpdateJobData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdateJobData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdateJobData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UpdateJobData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDuration

`func (o *UpdateJobData) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateJobData) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateJobData) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UpdateJobData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTargetVersion

`func (o *UpdateJobData) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *UpdateJobData) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *UpdateJobData) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *UpdateJobData) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetError

`func (o *UpdateJobData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateJobData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateJobData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateJobData) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


