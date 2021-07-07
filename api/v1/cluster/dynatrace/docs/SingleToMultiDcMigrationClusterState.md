# SingleToMultiDcMigrationClusterState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleToMultiDcMigration** | Pointer to [**MigrationState**](MigrationState.md) |  | [optional] 
**MigrationSteps** | Pointer to [**map[string]MigrationState**](MigrationState.md) | Status of the various sub steps of the migration | [optional] 

## Methods

### NewSingleToMultiDcMigrationClusterState

`func NewSingleToMultiDcMigrationClusterState() *SingleToMultiDcMigrationClusterState`

NewSingleToMultiDcMigrationClusterState instantiates a new SingleToMultiDcMigrationClusterState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleToMultiDcMigrationClusterStateWithDefaults

`func NewSingleToMultiDcMigrationClusterStateWithDefaults() *SingleToMultiDcMigrationClusterState`

NewSingleToMultiDcMigrationClusterStateWithDefaults instantiates a new SingleToMultiDcMigrationClusterState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleToMultiDcMigration

`func (o *SingleToMultiDcMigrationClusterState) GetSingleToMultiDcMigration() MigrationState`

GetSingleToMultiDcMigration returns the SingleToMultiDcMigration field if non-nil, zero value otherwise.

### GetSingleToMultiDcMigrationOk

`func (o *SingleToMultiDcMigrationClusterState) GetSingleToMultiDcMigrationOk() (*MigrationState, bool)`

GetSingleToMultiDcMigrationOk returns a tuple with the SingleToMultiDcMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleToMultiDcMigration

`func (o *SingleToMultiDcMigrationClusterState) SetSingleToMultiDcMigration(v MigrationState)`

SetSingleToMultiDcMigration sets SingleToMultiDcMigration field to given value.

### HasSingleToMultiDcMigration

`func (o *SingleToMultiDcMigrationClusterState) HasSingleToMultiDcMigration() bool`

HasSingleToMultiDcMigration returns a boolean if a field has been set.

### GetMigrationSteps

`func (o *SingleToMultiDcMigrationClusterState) GetMigrationSteps() map[string]MigrationState`

GetMigrationSteps returns the MigrationSteps field if non-nil, zero value otherwise.

### GetMigrationStepsOk

`func (o *SingleToMultiDcMigrationClusterState) GetMigrationStepsOk() (*map[string]MigrationState, bool)`

GetMigrationStepsOk returns a tuple with the MigrationSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationSteps

`func (o *SingleToMultiDcMigrationClusterState) SetMigrationSteps(v map[string]MigrationState)`

SetMigrationSteps sets MigrationSteps field to given value.

### HasMigrationSteps

`func (o *SingleToMultiDcMigrationClusterState) HasMigrationSteps() bool`

HasMigrationSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


