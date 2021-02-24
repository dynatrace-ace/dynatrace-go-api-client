# AssignedEntitiesWithMetricTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedEntities** | **[]string** | The list of Dynatrace entities, assigned to the tile. | 
**Metric** | Pointer to **string** | The metric assigned to the tile. | [optional] 

## Methods

### NewAssignedEntitiesWithMetricTile

`func NewAssignedEntitiesWithMetricTile(assignedEntities []string, ) *AssignedEntitiesWithMetricTile`

NewAssignedEntitiesWithMetricTile instantiates a new AssignedEntitiesWithMetricTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedEntitiesWithMetricTileWithDefaults

`func NewAssignedEntitiesWithMetricTileWithDefaults() *AssignedEntitiesWithMetricTile`

NewAssignedEntitiesWithMetricTileWithDefaults instantiates a new AssignedEntitiesWithMetricTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedEntities

`func (o *AssignedEntitiesWithMetricTile) GetAssignedEntities() []string`

GetAssignedEntities returns the AssignedEntities field if non-nil, zero value otherwise.

### GetAssignedEntitiesOk

`func (o *AssignedEntitiesWithMetricTile) GetAssignedEntitiesOk() (*[]string, bool)`

GetAssignedEntitiesOk returns a tuple with the AssignedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEntities

`func (o *AssignedEntitiesWithMetricTile) SetAssignedEntities(v []string)`

SetAssignedEntities sets AssignedEntities field to given value.


### GetMetric

`func (o *AssignedEntitiesWithMetricTile) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AssignedEntitiesWithMetricTile) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AssignedEntitiesWithMetricTile) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AssignedEntitiesWithMetricTile) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


