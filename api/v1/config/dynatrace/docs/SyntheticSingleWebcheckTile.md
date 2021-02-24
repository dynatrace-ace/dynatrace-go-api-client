# SyntheticSingleWebcheckTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedEntities** | **[]string** | The list of Dynatrace entities, assigned to the tile. | 
**ExcludeMaintenanceWindows** | Pointer to **bool** | Include (&#x60;false&#39;) or exclude (&#x60;true&#x60;) maintenance windows from availability calculations. | [optional] 

## Methods

### NewSyntheticSingleWebcheckTile

`func NewSyntheticSingleWebcheckTile(assignedEntities []string, ) *SyntheticSingleWebcheckTile`

NewSyntheticSingleWebcheckTile instantiates a new SyntheticSingleWebcheckTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticSingleWebcheckTileWithDefaults

`func NewSyntheticSingleWebcheckTileWithDefaults() *SyntheticSingleWebcheckTile`

NewSyntheticSingleWebcheckTileWithDefaults instantiates a new SyntheticSingleWebcheckTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedEntities

`func (o *SyntheticSingleWebcheckTile) GetAssignedEntities() []string`

GetAssignedEntities returns the AssignedEntities field if non-nil, zero value otherwise.

### GetAssignedEntitiesOk

`func (o *SyntheticSingleWebcheckTile) GetAssignedEntitiesOk() (*[]string, bool)`

GetAssignedEntitiesOk returns a tuple with the AssignedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEntities

`func (o *SyntheticSingleWebcheckTile) SetAssignedEntities(v []string)`

SetAssignedEntities sets AssignedEntities field to given value.


### GetExcludeMaintenanceWindows

`func (o *SyntheticSingleWebcheckTile) GetExcludeMaintenanceWindows() bool`

GetExcludeMaintenanceWindows returns the ExcludeMaintenanceWindows field if non-nil, zero value otherwise.

### GetExcludeMaintenanceWindowsOk

`func (o *SyntheticSingleWebcheckTile) GetExcludeMaintenanceWindowsOk() (*bool, bool)`

GetExcludeMaintenanceWindowsOk returns a tuple with the ExcludeMaintenanceWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMaintenanceWindows

`func (o *SyntheticSingleWebcheckTile) SetExcludeMaintenanceWindows(v bool)`

SetExcludeMaintenanceWindows sets ExcludeMaintenanceWindows field to given value.

### HasExcludeMaintenanceWindows

`func (o *SyntheticSingleWebcheckTile) HasExcludeMaintenanceWindows() bool`

HasExcludeMaintenanceWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


