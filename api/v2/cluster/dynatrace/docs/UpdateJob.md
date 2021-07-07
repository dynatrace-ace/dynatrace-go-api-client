# UpdateJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | The ID of the update job. | [optional] [readonly] 
**JobState** | Pointer to **string** | The status of the update job. | [optional] [readonly] 
**UpdateMethod** | Pointer to **string** | The method of updating the ActiveGate or its component. | [optional] [readonly] 
**UpdateType** | Pointer to **string** | The component to be updated. | [optional] [readonly] 
**Cancelable** | Pointer to **bool** | The job can (&#x60;true&#x60;) or can&#39;t (&#x60;false&#x60;) be cancelled at the moment. | [optional] [readonly] 
**StartVersion** | Pointer to **string** | The initial version of the ActiveGate. | [optional] [readonly] 
**TargetVersion** | **string** | The target version of the update.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;.&lt;timestamp&gt;&#x60; format.   To update to the latest available version, use the &#x60;latest&#x60; value. | 
**Timestamp** | Pointer to **int64** | The timestamp of the update job completion.    The &#x60;null&#x60; value means the job is still running. | [optional] [readonly] 
**AgType** | Pointer to **string** | The type of the ActiveGate. | [optional] [readonly] 
**Environments** | Pointer to **[]string** | A list of environments (specified by IDs) the ActiveGate can connect to. | [optional] [readonly] 
**Error** | Pointer to **string** | The information about update error. | [optional] [readonly] 
**Duration** | Pointer to **int64** | The duration of the update, in milliseconds. | [optional] [readonly] 

## Methods

### NewUpdateJob

`func NewUpdateJob(targetVersion string, ) *UpdateJob`

NewUpdateJob instantiates a new UpdateJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJobWithDefaults

`func NewUpdateJobWithDefaults() *UpdateJob`

NewUpdateJobWithDefaults instantiates a new UpdateJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *UpdateJob) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UpdateJob) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UpdateJob) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *UpdateJob) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobState

`func (o *UpdateJob) GetJobState() string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *UpdateJob) GetJobStateOk() (*string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *UpdateJob) SetJobState(v string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *UpdateJob) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetUpdateMethod

`func (o *UpdateJob) GetUpdateMethod() string`

GetUpdateMethod returns the UpdateMethod field if non-nil, zero value otherwise.

### GetUpdateMethodOk

`func (o *UpdateJob) GetUpdateMethodOk() (*string, bool)`

GetUpdateMethodOk returns a tuple with the UpdateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMethod

`func (o *UpdateJob) SetUpdateMethod(v string)`

SetUpdateMethod sets UpdateMethod field to given value.

### HasUpdateMethod

`func (o *UpdateJob) HasUpdateMethod() bool`

HasUpdateMethod returns a boolean if a field has been set.

### GetUpdateType

`func (o *UpdateJob) GetUpdateType() string`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *UpdateJob) GetUpdateTypeOk() (*string, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *UpdateJob) SetUpdateType(v string)`

SetUpdateType sets UpdateType field to given value.

### HasUpdateType

`func (o *UpdateJob) HasUpdateType() bool`

HasUpdateType returns a boolean if a field has been set.

### GetCancelable

`func (o *UpdateJob) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *UpdateJob) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *UpdateJob) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *UpdateJob) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetStartVersion

`func (o *UpdateJob) GetStartVersion() string`

GetStartVersion returns the StartVersion field if non-nil, zero value otherwise.

### GetStartVersionOk

`func (o *UpdateJob) GetStartVersionOk() (*string, bool)`

GetStartVersionOk returns a tuple with the StartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVersion

`func (o *UpdateJob) SetStartVersion(v string)`

SetStartVersion sets StartVersion field to given value.

### HasStartVersion

`func (o *UpdateJob) HasStartVersion() bool`

HasStartVersion returns a boolean if a field has been set.

### GetTargetVersion

`func (o *UpdateJob) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *UpdateJob) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *UpdateJob) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetTimestamp

`func (o *UpdateJob) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdateJob) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdateJob) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UpdateJob) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAgType

`func (o *UpdateJob) GetAgType() string`

GetAgType returns the AgType field if non-nil, zero value otherwise.

### GetAgTypeOk

`func (o *UpdateJob) GetAgTypeOk() (*string, bool)`

GetAgTypeOk returns a tuple with the AgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgType

`func (o *UpdateJob) SetAgType(v string)`

SetAgType sets AgType field to given value.

### HasAgType

`func (o *UpdateJob) HasAgType() bool`

HasAgType returns a boolean if a field has been set.

### GetEnvironments

`func (o *UpdateJob) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *UpdateJob) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *UpdateJob) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *UpdateJob) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetError

`func (o *UpdateJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDuration

`func (o *UpdateJob) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateJob) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateJob) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UpdateJob) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


