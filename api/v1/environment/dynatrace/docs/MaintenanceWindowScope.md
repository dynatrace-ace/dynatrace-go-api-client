# MaintenanceWindowScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to **[]string** | Defines Dynatrace entities to be included in scope, for example hosts, services, process groups.   Allowed values are Dynatrace entity IDs. | [optional] 
**Matches** | Pointer to [**[]MonitoredEntityFilter**](MonitoredEntityFilter.md) | An object defining a matching rule for dynamic scope formation. An empty rule matches for all entities. | [optional] 

## Methods

### NewMaintenanceWindowScope

`func NewMaintenanceWindowScope() *MaintenanceWindowScope`

NewMaintenanceWindowScope instantiates a new MaintenanceWindowScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowScopeWithDefaults

`func NewMaintenanceWindowScopeWithDefaults() *MaintenanceWindowScope`

NewMaintenanceWindowScopeWithDefaults instantiates a new MaintenanceWindowScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *MaintenanceWindowScope) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *MaintenanceWindowScope) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *MaintenanceWindowScope) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *MaintenanceWindowScope) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMatches

`func (o *MaintenanceWindowScope) GetMatches() []MonitoredEntityFilter`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *MaintenanceWindowScope) GetMatchesOk() (*[]MonitoredEntityFilter, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *MaintenanceWindowScope) SetMatches(v []MonitoredEntityFilter)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *MaintenanceWindowScope) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


