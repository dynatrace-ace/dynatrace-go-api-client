# Tile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the tile. | 
**TileType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;DATA_EXPLORER&#x60; -&gt; DataExplorerTile  * &#x60;CUSTOM_CHARTING&#x60; -&gt; CustomChartingTile  * &#x60;DTAQL&#x60; -&gt; UserSessionQueryTile  * &#x60;MARKDOWN&#x60; -&gt; MarkdownTile  * &#x60;HOSTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATIONS&#x60; -&gt; FilterableEntityTile  * &#x60;SERVICES&#x60; -&gt; FilterableEntityTile  * &#x60;DATABASES_OVERVIEW&#x60; -&gt; FilterableEntityTile  * &#x60;SYNTHETIC_TESTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATION_WORLDMAP&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;RESOURCES&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;THIRD_PARTY_MOST_ACTIVE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;UEM_CONVERSIONS_PER_GOAL&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;HOST&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;PROCESS_GROUPS_ONE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;SYNTHETIC_SINGLE_WEBCHECK&#x60; -&gt; SyntheticSingleWebcheckTile  * &#x60;APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;VIRTUALIZATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;AWS&#x60; -&gt; AssignedEntitiesTile  * &#x60;SERVICE_VERSATILE&#x60; -&gt; AssignedEntitiesTile  * &#x60;SESSION_METRICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;USERS&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_KEY_USER_ACTIONS&#x60; -&gt; AssignedEntitiesTile  * &#x60;BOUNCE_RATE&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_CONVERSIONS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_JSERRORS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;MOBILE_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_SINGLE_EXT_TEST&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_HTTP_MONITOR&#x60; -&gt; AssignedEntitiesTile  * &#x60;DATABASE&#x60; -&gt; AssignedEntitiesTile  * &#x60;CUSTOM_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;LOG_ANALYTICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_PROJECT&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_AV_ZONE&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEVICE_APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEM_KEY_USER_ACTION&#x60; -&gt; AssignedEntitiesTile   | 
**Configured** | **bool** | The tile is configured and ready to use (&#x60;true&#x60;) or just placed on the dashboard (&#x60;false&#x60;). | [optional] 
**Bounds** | [**TileBounds**](TileBounds.md) |  | 
**TileFilter** | [**TileFilter**](TileFilter.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


