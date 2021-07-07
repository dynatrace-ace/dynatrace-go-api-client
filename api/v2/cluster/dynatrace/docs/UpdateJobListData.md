# UpdateJobListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateJobs** | Pointer to [**[]UpdateJobData**](UpdateJobData.md) |  | [optional] 
**AgId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateJobListData

`func NewUpdateJobListData() *UpdateJobListData`

NewUpdateJobListData instantiates a new UpdateJobListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJobListDataWithDefaults

`func NewUpdateJobListDataWithDefaults() *UpdateJobListData`

NewUpdateJobListDataWithDefaults instantiates a new UpdateJobListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateJobs

`func (o *UpdateJobListData) GetUpdateJobs() []UpdateJobData`

GetUpdateJobs returns the UpdateJobs field if non-nil, zero value otherwise.

### GetUpdateJobsOk

`func (o *UpdateJobListData) GetUpdateJobsOk() (*[]UpdateJobData, bool)`

GetUpdateJobsOk returns a tuple with the UpdateJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateJobs

`func (o *UpdateJobListData) SetUpdateJobs(v []UpdateJobData)`

SetUpdateJobs sets UpdateJobs field to given value.

### HasUpdateJobs

`func (o *UpdateJobListData) HasUpdateJobs() bool`

HasUpdateJobs returns a boolean if a field has been set.

### GetAgId

`func (o *UpdateJobListData) GetAgId() string`

GetAgId returns the AgId field if non-nil, zero value otherwise.

### GetAgIdOk

`func (o *UpdateJobListData) GetAgIdOk() (*string, bool)`

GetAgIdOk returns a tuple with the AgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgId

`func (o *UpdateJobListData) SetAgId(v string)`

SetAgId sets AgId field to given value.

### HasAgId

`func (o *UpdateJobListData) HasAgId() bool`

HasAgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


