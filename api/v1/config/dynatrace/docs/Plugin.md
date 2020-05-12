# Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the plugin, for example &#x60;custom.remote.python.demo&#x60;. | [optional] 
**Name** | **string** | The name of the plugin, displayed in Dynatrace. | [optional] 
**Version** | **string** | The version of the plugin, displayed in Dynatrace. | [optional] 
**Type** | **string** | The type of the plugin. It indicates the runtime environment of the plugin (for example, ActiveGate). | [optional] 
**MetricGroup** | **string** | The metric group of the plugin. All the metrics, captured by the plugin are children of this group. | [optional] 
**Properties** | [**[]PluginProperty**](PluginProperty.md) | A list of plugin properties. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


