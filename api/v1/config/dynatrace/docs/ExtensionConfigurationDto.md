# ExtensionConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | **string** | The ID of the extension. | [optional] 
**Enabled** | **bool** | The extension is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**UseGlobal** | **bool** | Allows to skip current configuration and use global one. | 
**Properties** | **map[string]string** | The list of extension parameters.    Each parameter is a key-value pair. | [optional] 
**HostId** | **string** | The ID of the host on which the extension runs. | [optional] 
**ActiveGate** | [**EntityShortRepresentation**](EntityShortRepresentation.md) |  | [optional] 
**EndpointId** | **string** | The ID of the endpoint. | [optional] 
**EndpointName** | **string** | The name of the endpoint, displayed in Dynatrace. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


