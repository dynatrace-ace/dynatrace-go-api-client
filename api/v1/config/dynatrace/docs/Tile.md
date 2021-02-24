# Tile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the tile. | 
**TileType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;CUSTOM_CHARTING&#x60; -&gt; CustomChartingTile  * &#x60;DTAQL&#x60; -&gt; UserSessionQueryTile  * &#x60;MARKDOWN&#x60; -&gt; MarkdownTile  * &#x60;HOSTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATIONS&#x60; -&gt; FilterableEntityTile  * &#x60;SERVICES&#x60; -&gt; FilterableEntityTile  * &#x60;DATABASES_OVERVIEW&#x60; -&gt; FilterableEntityTile  * &#x60;SYNTHETIC_TESTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATION_WORLDMAP&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;RESOURCES&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;THIRD_PARTY_MOST_ACTIVE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;UEM_CONVERSIONS_PER_GOAL&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;HOST&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;PROCESS_GROUPS_ONE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;SYNTHETIC_SINGLE_WEBCHECK&#x60; -&gt; SyntheticSingleWebcheckTile  * &#x60;APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;VIRTUALIZATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;AWS&#x60; -&gt; AssignedEntitiesTile  * &#x60;SERVICE_VERSATILE&#x60; -&gt; AssignedEntitiesTile  * &#x60;SESSION_METRICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;USERS&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_KEY_USER_ACTIONS&#x60; -&gt; AssignedEntitiesTile  * &#x60;BOUNCE_RATE&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_CONVERSIONS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_JSERRORS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;MOBILE_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_SINGLE_EXT_TEST&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_HTTP_MONITOR&#x60; -&gt; AssignedEntitiesTile  * &#x60;DATABASE&#x60; -&gt; AssignedEntitiesTile  * &#x60;CUSTOM_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;LOG_ANALYTICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_PROJECT&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_AV_ZONE&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEVICE_APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEM_KEY_USER_ACTION&#x60; -&gt; AssignedEntitiesTile  * &#x60;SLO&#x60; -&gt; AssignedEntitiesTile   | 
**Configured** | Pointer to **bool** | The tile is configured and ready to use (&#x60;true&#x60;) or just placed on the dashboard (&#x60;false&#x60;). | [optional] 
**Bounds** | [**TileBounds**](TileBounds.md) |  | 
**TileFilter** | Pointer to [**TileFilter**](TileFilter.md) |  | [optional] 
**AssignedEntities** | Pointer to **[]string** | The list of Dynatrace entities, assigned to the tile. | [optional] 
**Metric** | Pointer to **string** | The metric assigned to the tile. | [optional] 
**FilterConfig** | Pointer to [**CustomFilterConfig**](CustomFilterConfig.md) |  | [optional] 
**ChartVisible** | Pointer to **bool** |  | [optional] 
**Markdown** | Pointer to **string** | The markdown-formatted content of the tile. | [optional] 
**ExcludeMaintenanceWindows** | Pointer to **bool** | Include (&#x60;false&#39;) or exclude (&#x60;true&#x60;) maintenance windows from availability calculations. | [optional] 
**CustomName** | Pointer to **string** | The name of the tile, set by user. | [optional] 
**Query** | Pointer to **string** | A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile. | [optional] 
**Type** | Pointer to **string** | The visualization of the tile. | [optional] 
**TimeFrameShift** | Pointer to **string** | The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift. | [optional] 
**VisualizationConfig** | Pointer to [**UserSessionQueryTileConfiguration**](UserSessionQueryTileConfiguration.md) |  | [optional] 
**Limit** | Pointer to **int32** | The limit of the results, if not set will use the default value of the system | [optional] 

## Methods

### NewTile

`func NewTile(name string, tileType string, bounds TileBounds, ) *Tile`

NewTile instantiates a new Tile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTileWithDefaults

`func NewTileWithDefaults() *Tile`

NewTileWithDefaults instantiates a new Tile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tile) SetName(v string)`

SetName sets Name field to given value.


### GetTileType

`func (o *Tile) GetTileType() string`

GetTileType returns the TileType field if non-nil, zero value otherwise.

### GetTileTypeOk

`func (o *Tile) GetTileTypeOk() (*string, bool)`

GetTileTypeOk returns a tuple with the TileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileType

`func (o *Tile) SetTileType(v string)`

SetTileType sets TileType field to given value.


### GetConfigured

`func (o *Tile) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *Tile) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *Tile) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *Tile) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetBounds

`func (o *Tile) GetBounds() TileBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *Tile) GetBoundsOk() (*TileBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *Tile) SetBounds(v TileBounds)`

SetBounds sets Bounds field to given value.


### GetTileFilter

`func (o *Tile) GetTileFilter() TileFilter`

GetTileFilter returns the TileFilter field if non-nil, zero value otherwise.

### GetTileFilterOk

`func (o *Tile) GetTileFilterOk() (*TileFilter, bool)`

GetTileFilterOk returns a tuple with the TileFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileFilter

`func (o *Tile) SetTileFilter(v TileFilter)`

SetTileFilter sets TileFilter field to given value.

### HasTileFilter

`func (o *Tile) HasTileFilter() bool`

HasTileFilter returns a boolean if a field has been set.

### GetAssignedEntities

`func (o *Tile) GetAssignedEntities() []string`

GetAssignedEntities returns the AssignedEntities field if non-nil, zero value otherwise.

### GetAssignedEntitiesOk

`func (o *Tile) GetAssignedEntitiesOk() (*[]string, bool)`

GetAssignedEntitiesOk returns a tuple with the AssignedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEntities

`func (o *Tile) SetAssignedEntities(v []string)`

SetAssignedEntities sets AssignedEntities field to given value.

### HasAssignedEntities

`func (o *Tile) HasAssignedEntities() bool`

HasAssignedEntities returns a boolean if a field has been set.

### GetMetric

`func (o *Tile) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Tile) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Tile) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Tile) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetFilterConfig

`func (o *Tile) GetFilterConfig() CustomFilterConfig`

GetFilterConfig returns the FilterConfig field if non-nil, zero value otherwise.

### GetFilterConfigOk

`func (o *Tile) GetFilterConfigOk() (*CustomFilterConfig, bool)`

GetFilterConfigOk returns a tuple with the FilterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConfig

`func (o *Tile) SetFilterConfig(v CustomFilterConfig)`

SetFilterConfig sets FilterConfig field to given value.

### HasFilterConfig

`func (o *Tile) HasFilterConfig() bool`

HasFilterConfig returns a boolean if a field has been set.

### GetChartVisible

`func (o *Tile) GetChartVisible() bool`

GetChartVisible returns the ChartVisible field if non-nil, zero value otherwise.

### GetChartVisibleOk

`func (o *Tile) GetChartVisibleOk() (*bool, bool)`

GetChartVisibleOk returns a tuple with the ChartVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVisible

`func (o *Tile) SetChartVisible(v bool)`

SetChartVisible sets ChartVisible field to given value.

### HasChartVisible

`func (o *Tile) HasChartVisible() bool`

HasChartVisible returns a boolean if a field has been set.

### GetMarkdown

`func (o *Tile) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *Tile) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *Tile) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *Tile) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.

### GetExcludeMaintenanceWindows

`func (o *Tile) GetExcludeMaintenanceWindows() bool`

GetExcludeMaintenanceWindows returns the ExcludeMaintenanceWindows field if non-nil, zero value otherwise.

### GetExcludeMaintenanceWindowsOk

`func (o *Tile) GetExcludeMaintenanceWindowsOk() (*bool, bool)`

GetExcludeMaintenanceWindowsOk returns a tuple with the ExcludeMaintenanceWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMaintenanceWindows

`func (o *Tile) SetExcludeMaintenanceWindows(v bool)`

SetExcludeMaintenanceWindows sets ExcludeMaintenanceWindows field to given value.

### HasExcludeMaintenanceWindows

`func (o *Tile) HasExcludeMaintenanceWindows() bool`

HasExcludeMaintenanceWindows returns a boolean if a field has been set.

### GetCustomName

`func (o *Tile) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Tile) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Tile) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Tile) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### GetQuery

`func (o *Tile) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Tile) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Tile) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Tile) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *Tile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Tile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeFrameShift

`func (o *Tile) GetTimeFrameShift() string`

GetTimeFrameShift returns the TimeFrameShift field if non-nil, zero value otherwise.

### GetTimeFrameShiftOk

`func (o *Tile) GetTimeFrameShiftOk() (*string, bool)`

GetTimeFrameShiftOk returns a tuple with the TimeFrameShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrameShift

`func (o *Tile) SetTimeFrameShift(v string)`

SetTimeFrameShift sets TimeFrameShift field to given value.

### HasTimeFrameShift

`func (o *Tile) HasTimeFrameShift() bool`

HasTimeFrameShift returns a boolean if a field has been set.

### GetVisualizationConfig

`func (o *Tile) GetVisualizationConfig() UserSessionQueryTileConfiguration`

GetVisualizationConfig returns the VisualizationConfig field if non-nil, zero value otherwise.

### GetVisualizationConfigOk

`func (o *Tile) GetVisualizationConfigOk() (*UserSessionQueryTileConfiguration, bool)`

GetVisualizationConfigOk returns a tuple with the VisualizationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationConfig

`func (o *Tile) SetVisualizationConfig(v UserSessionQueryTileConfiguration)`

SetVisualizationConfig sets VisualizationConfig field to given value.

### HasVisualizationConfig

`func (o *Tile) HasVisualizationConfig() bool`

HasVisualizationConfig returns a boolean if a field has been set.

### GetLimit

`func (o *Tile) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Tile) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Tile) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Tile) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


