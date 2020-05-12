# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The Dynatrace entity ID of the Azure credentials configuration. | [optional] [readonly] 
**Label** | **string** | The unique name of the Azure credentials configuration.   Allowed characters are letters, numbers, and spaces. Also the special characters &#x60;.+-_&#x60; are allowed. | 
**AppId** | **string** | The Application ID (also referred to as Client ID)   The combination of Application ID and Directory ID must be unique. | [readonly] 
**DirectoryId** | **string** | The Directory ID (also referred to as Tenant ID)   The combination of Application ID and Directory ID must be unique. | [readonly] 
**Key** | **string** | The secret key associated with the Application ID.   For security reasons, GET requests return this field as &#x60;null&#x60;.    Submit your key on creation or update of the configuration. If the field is omitted during an update, the old value remains unaffected. | [optional] 
**Active** | **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**AutoTagging** | **bool** | The automatic capture of Azure tags is on (&#x60;true&#x60;) or off (&#x60;false&#x60;). | 
**MonitorOnlyTaggedEntities** | **bool** | Monitor only resources that have specified Azure tags (&#x60;true&#x60;) or all resources (&#x60;false&#x60;). | 
**MonitorOnlyTagPairs** | [**[]CloudTag**](CloudTag.md) | A list of Azure tags to be monitored.   You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to &#x60;true&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


