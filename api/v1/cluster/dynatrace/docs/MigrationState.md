# MigrationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current status of migration | [optional] 
**StartedAt** | Pointer to **int64** | Timestamp (milliseconds format) of migration start | [optional] 
**FinishedAt** | Pointer to **int64** | Timestamp (milliseconds format) of migration finish | [optional] 
**Details** | Pointer to **string** | Additional information about migration state | [optional] 

## Methods

### NewMigrationState

`func NewMigrationState() *MigrationState`

NewMigrationState instantiates a new MigrationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationStateWithDefaults

`func NewMigrationStateWithDefaults() *MigrationState`

NewMigrationStateWithDefaults instantiates a new MigrationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MigrationState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MigrationState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *MigrationState) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MigrationState) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MigrationState) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MigrationState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *MigrationState) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *MigrationState) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *MigrationState) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *MigrationState) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetDetails

`func (o *MigrationState) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MigrationState) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MigrationState) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MigrationState) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


