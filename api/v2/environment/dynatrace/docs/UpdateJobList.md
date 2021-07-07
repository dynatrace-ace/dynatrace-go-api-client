# UpdateJobList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgId** | Pointer to **string** | The ID of the ActiveGate. | [optional] [readonly] 
**UpdateJobs** | Pointer to [**[]UpdateJob**](UpdateJob.md) | A list of update jobs of the ActiveGate. | [optional] [readonly] 

## Methods

### NewUpdateJobList

`func NewUpdateJobList() *UpdateJobList`

NewUpdateJobList instantiates a new UpdateJobList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJobListWithDefaults

`func NewUpdateJobListWithDefaults() *UpdateJobList`

NewUpdateJobListWithDefaults instantiates a new UpdateJobList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgId

`func (o *UpdateJobList) GetAgId() string`

GetAgId returns the AgId field if non-nil, zero value otherwise.

### GetAgIdOk

`func (o *UpdateJobList) GetAgIdOk() (*string, bool)`

GetAgIdOk returns a tuple with the AgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgId

`func (o *UpdateJobList) SetAgId(v string)`

SetAgId sets AgId field to given value.

### HasAgId

`func (o *UpdateJobList) HasAgId() bool`

HasAgId returns a boolean if a field has been set.

### GetUpdateJobs

`func (o *UpdateJobList) GetUpdateJobs() []UpdateJob`

GetUpdateJobs returns the UpdateJobs field if non-nil, zero value otherwise.

### GetUpdateJobsOk

`func (o *UpdateJobList) GetUpdateJobsOk() (*[]UpdateJob, bool)`

GetUpdateJobsOk returns a tuple with the UpdateJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateJobs

`func (o *UpdateJobList) SetUpdateJobs(v []UpdateJob)`

SetUpdateJobs sets UpdateJobs field to given value.

### HasUpdateJobs

`func (o *UpdateJobList) HasUpdateJobs() bool`

HasUpdateJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


