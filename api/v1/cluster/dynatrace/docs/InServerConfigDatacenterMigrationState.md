# InServerConfigDatacenterMigrationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentMigrationStates** | Pointer to [**map[string]MigrationState**](MigrationState.md) | Map of (component, migration state) pairs | [optional] 

## Methods

### NewInServerConfigDatacenterMigrationState

`func NewInServerConfigDatacenterMigrationState() *InServerConfigDatacenterMigrationState`

NewInServerConfigDatacenterMigrationState instantiates a new InServerConfigDatacenterMigrationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInServerConfigDatacenterMigrationStateWithDefaults

`func NewInServerConfigDatacenterMigrationStateWithDefaults() *InServerConfigDatacenterMigrationState`

NewInServerConfigDatacenterMigrationStateWithDefaults instantiates a new InServerConfigDatacenterMigrationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentMigrationStates

`func (o *InServerConfigDatacenterMigrationState) GetComponentMigrationStates() map[string]MigrationState`

GetComponentMigrationStates returns the ComponentMigrationStates field if non-nil, zero value otherwise.

### GetComponentMigrationStatesOk

`func (o *InServerConfigDatacenterMigrationState) GetComponentMigrationStatesOk() (*map[string]MigrationState, bool)`

GetComponentMigrationStatesOk returns a tuple with the ComponentMigrationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMigrationStates

`func (o *InServerConfigDatacenterMigrationState) SetComponentMigrationStates(v map[string]MigrationState)`

SetComponentMigrationStates sets ComponentMigrationStates field to given value.

### HasComponentMigrationStates

`func (o *InServerConfigDatacenterMigrationState) HasComponentMigrationStates() bool`

HasComponentMigrationStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


