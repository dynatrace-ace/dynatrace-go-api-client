# UserSessionQueryTileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | **string** | The name of the tile, set by user. | [optional] 
**Query** | **string** | A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile. | [optional] 
**Type** | **string** | The visualization of the tile. | [optional] 
**TimeFrameShift** | **string** | The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift. | [optional] 
**VisualizationConfig** | [**UserSessionQueryTileConfiguration**](UserSessionQueryTileConfiguration.md) |  | [optional] 
**Limit** | **int32** | The limit of the results, if not set will use the default value of the system | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


