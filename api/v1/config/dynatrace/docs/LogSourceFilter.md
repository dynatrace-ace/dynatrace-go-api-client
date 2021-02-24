# LogSourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathDefinitions** | Pointer to [**[]PathDefinition**](PathDefinition.md) | A list of filtering criteria for log path.   If several criteria are specified, the OR logic applies. | [optional] 
**SourceEntities** | Pointer to **[]string** | A list of Dynatrace entities, where the log can originate from. Specify Dynatrace entity IDs here.    Allowed types of entities are &#x60;PROCESS_GROUP&#x60; and &#x60;PROCESS_GROUP_INSTANCE&#x60;. You can&#39;t mix entity types.   If several entities are specified, the OR logic applies.   This field is mutually exclusive with the **osTypes** field. | [optional] 
**HostFilters** | Pointer to **[]string** | A list of hosts, where the log can originate from. Specify Dynatrace entity IDs here.   If several hosts are specified, the OR logic applies. | [optional] 
**OsTypes** | Pointer to **[]string** | A list of operating system types, where the log can originate from.   If set, **only OS logs** are included in the result.   If several OS are specified, the OR logic applies.   This field is mutually exclusive with the **sourceEntities** field. | [optional] 

## Methods

### NewLogSourceFilter

`func NewLogSourceFilter() *LogSourceFilter`

NewLogSourceFilter instantiates a new LogSourceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSourceFilterWithDefaults

`func NewLogSourceFilterWithDefaults() *LogSourceFilter`

NewLogSourceFilterWithDefaults instantiates a new LogSourceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathDefinitions

`func (o *LogSourceFilter) GetPathDefinitions() []PathDefinition`

GetPathDefinitions returns the PathDefinitions field if non-nil, zero value otherwise.

### GetPathDefinitionsOk

`func (o *LogSourceFilter) GetPathDefinitionsOk() (*[]PathDefinition, bool)`

GetPathDefinitionsOk returns a tuple with the PathDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathDefinitions

`func (o *LogSourceFilter) SetPathDefinitions(v []PathDefinition)`

SetPathDefinitions sets PathDefinitions field to given value.

### HasPathDefinitions

`func (o *LogSourceFilter) HasPathDefinitions() bool`

HasPathDefinitions returns a boolean if a field has been set.

### GetSourceEntities

`func (o *LogSourceFilter) GetSourceEntities() []string`

GetSourceEntities returns the SourceEntities field if non-nil, zero value otherwise.

### GetSourceEntitiesOk

`func (o *LogSourceFilter) GetSourceEntitiesOk() (*[]string, bool)`

GetSourceEntitiesOk returns a tuple with the SourceEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntities

`func (o *LogSourceFilter) SetSourceEntities(v []string)`

SetSourceEntities sets SourceEntities field to given value.

### HasSourceEntities

`func (o *LogSourceFilter) HasSourceEntities() bool`

HasSourceEntities returns a boolean if a field has been set.

### GetHostFilters

`func (o *LogSourceFilter) GetHostFilters() []string`

GetHostFilters returns the HostFilters field if non-nil, zero value otherwise.

### GetHostFiltersOk

`func (o *LogSourceFilter) GetHostFiltersOk() (*[]string, bool)`

GetHostFiltersOk returns a tuple with the HostFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFilters

`func (o *LogSourceFilter) SetHostFilters(v []string)`

SetHostFilters sets HostFilters field to given value.

### HasHostFilters

`func (o *LogSourceFilter) HasHostFilters() bool`

HasHostFilters returns a boolean if a field has been set.

### GetOsTypes

`func (o *LogSourceFilter) GetOsTypes() []string`

GetOsTypes returns the OsTypes field if non-nil, zero value otherwise.

### GetOsTypesOk

`func (o *LogSourceFilter) GetOsTypesOk() (*[]string, bool)`

GetOsTypesOk returns a tuple with the OsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypes

`func (o *LogSourceFilter) SetOsTypes(v []string)`

SetOsTypes sets OsTypes field to given value.

### HasOsTypes

`func (o *LogSourceFilter) HasOsTypes() bool`

HasOsTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


