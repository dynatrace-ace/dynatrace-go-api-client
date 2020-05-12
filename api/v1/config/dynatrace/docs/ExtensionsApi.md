# \ExtensionsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocalExtensionConfiguration**](ExtensionsApi.md#CreateLocalExtensionConfiguration) | **Post** /extensions/{id}/instances | Creates instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**DeleteExtension**](ExtensionsApi.md#DeleteExtension) | **Delete** /extensions/{id} | Deletes the ZIP file of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**DeleteLocalExtensionConfiguration**](ExtensionsApi.md#DeleteLocalExtensionConfiguration) | **Delete** /extensions/{id}/instances/{configurationId} | Deletes an existing configuration of the extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtension**](ExtensionsApi.md#GetExtension) | **Get** /extensions/{id} | Lists the properties of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionBinary**](ExtensionsApi.md#GetExtensionBinary) | **Get** /extensions/{id}/binary | Downloads the ZIP file of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionConfigurations**](ExtensionsApi.md#GetExtensionConfigurations) | **Get** /extensions/{id}/instances | Returns list of all local configuration instances for given extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionGlobalConfiguration**](ExtensionsApi.md#GetExtensionGlobalConfiguration) | **Get** /extensions/{id}/global | Get the global configuration of the specified OneAgent or JMX extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionStates**](ExtensionsApi.md#GetExtensionStates) | **Get** /extensions/{id}/states | Lists the states of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensions**](ExtensionsApi.md#GetExtensions) | **Get** /extensions | Lists all uploaded extensions | maturity&#x3D;EARLY_ADOPTER
[**GetHostsForTechnology**](ExtensionsApi.md#GetHostsForTechnology) | **Get** /extensions/{technology}/availableHosts | Lists all available hosts that have specified technology running | maturity&#x3D;EARLY_ADOPTER
[**GetLocalExtensionConfiguration**](ExtensionsApi.md#GetLocalExtensionConfiguration) | **Get** /extensions/{id}/instances/{configurationId} | Returns instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**GetRemoteExtensionModules**](ExtensionsApi.md#GetRemoteExtensionModules) | **Get** /extensions/activeGateExtensionModules | List available ActiveGate extension modules | maturity&#x3D;EARLY_ADOPTER
[**UpdateGlobalExtensionConfiguration**](ExtensionsApi.md#UpdateGlobalExtensionConfiguration) | **Put** /extensions/{id}/global | Updates the configuration of the specified OneAgent or JMX extension | maturity&#x3D;EARLY_ADOPTER
[**UpdateLocalExtensionConfiguration**](ExtensionsApi.md#UpdateLocalExtensionConfiguration) | **Put** /extensions/{id}/instances/{configurationId} | Updates instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**UploadExtension**](ExtensionsApi.md#UploadExtension) | **Post** /extensions | Uploads a ZIP extension file | maturity&#x3D;EARLY_ADOPTER
[**ValidateExtension**](ExtensionsApi.md#ValidateExtension) | **Post** /extensions/validator | Validates a ZIP extension file for &#x60;POST/extensions&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateLocalExtensionConfiguration**](ExtensionsApi.md#ValidateLocalExtensionConfiguration) | **Post** /extensions/{id}/instances/validator | Validates the payload for the &#x60;POST /extensions/{id}/instances&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateLocalExtensionConfiguration

> EntityShortRepresentation CreateLocalExtensionConfiguration(ctx, id, optional)

Creates instance of local configuration for given extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension | 
 **optional** | ***CreateLocalExtensionConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLocalExtensionConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionConfigurationDto** | [**optional.Interface of ExtensionConfigurationDto**](ExtensionConfigurationDto.md)| The JSON body of the request. Contains new configuration of the extension. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtension

> DeleteExtension(ctx, id)

Deletes the ZIP file of the specified extension | maturity=EARLY_ADOPTER

Deletion of the extension file uninstalls the extension, making it unavailable for use.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalExtensionConfiguration

> DeleteLocalExtensionConfiguration(ctx, id, configurationId)

Deletes an existing configuration of the extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension where you want to delete the configuration. | 
**configurationId** | **string**| The ID of the configuration to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtension

> Extension GetExtension(ctx, id)

Lists the properties of the specified extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required extension. | 

### Return type

[**Extension**](Extension.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionBinary

> GetExtensionBinary(ctx, id)

Downloads the ZIP file of the specified extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension you want to download. | 

### Return type

 (empty response body)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionConfigurations

> ExtensionConfigurationList GetExtensionConfigurations(ctx, id, optional)

Returns list of all local configuration instances for given extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required extension. | 
 **optional** | ***GetExtensionConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExtensionConfigurationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **optional.String**| The cursor for the next page of results. | 

### Return type

[**ExtensionConfigurationList**](ExtensionConfigurationList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionGlobalConfiguration

> GlobalExtensionConfiguration GetExtensionGlobalConfiguration(ctx, id)

Get the global configuration of the specified OneAgent or JMX extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension to be updated. | 

### Return type

[**GlobalExtensionConfiguration**](GlobalExtensionConfiguration.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionStates

> ExtensionStateList GetExtensionStates(ctx, id, optional)

Lists the states of the specified extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required extension. | 
 **optional** | ***GetExtensionStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExtensionStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **optional.String**| The cursor for the next page of results. | 
 **state** | **optional.String**| Extension state to filter. | 

### Return type

[**ExtensionStateList**](ExtensionStateList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensions

> ExtensionListDto GetExtensions(ctx, optional)

Lists all uploaded extensions | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **optional.String**| The cursor for the next page of results. | 

### Return type

[**ExtensionListDto**](ExtensionListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsForTechnology

> HostList GetHostsForTechnology(ctx, technology, optional)

Lists all available hosts that have specified technology running | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Name of requested technology | 
 **optional** | ***GetHostsForTechnologyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHostsForTechnologyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**optional.Interface of []string**](string.md)| Filters the resulting set of hosts by the specified tag.    You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags. | 
 **managementZone** | **optional.Int64**| Only return hosts that are part of the specified management zone. | 
 **hostGroupId** | **optional.String**| Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you&#39;re interested in. | 
 **hostGroupName** | **optional.String**| Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you&#39;re interested in. | 
 **pageSize** | **optional.Int32**| The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **optional.String**| The cursor for the next page of results. | 

### Return type

[**HostList**](HostList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalExtensionConfiguration

> ExtensionConfigurationDto GetLocalExtensionConfiguration(ctx, id, configurationId)

Returns instance of local configuration for given extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension. | 
**configurationId** | **string**| The ID of the configuration. | 

### Return type

[**ExtensionConfigurationDto**](ExtensionConfigurationDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteExtensionModules

> StubList GetRemoteExtensionModules(ctx, )

List available ActiveGate extension modules | maturity=EARLY_ADOPTER

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StubList**](StubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalExtensionConfiguration

> UpdateGlobalExtensionConfiguration(ctx, id, optional)

Updates the configuration of the specified OneAgent or JMX extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension to be updated. | 
 **optional** | ***UpdateGlobalExtensionConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateGlobalExtensionConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalExtensionConfiguration** | [**optional.Interface of GlobalExtensionConfiguration**](GlobalExtensionConfiguration.md)| The JSON body of the request. Contains updated configuration of the extension. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalExtensionConfiguration

> UpdateLocalExtensionConfiguration(ctx, id, configurationId, optional)

Updates instance of local configuration for given extension | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension where you want to update the configuration.   If you set the extension ID in the body as well, it must match this ID. | 
**configurationId** | **string**| The ID of the configuration to be updated. | 
 **optional** | ***UpdateLocalExtensionConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLocalExtensionConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionConfigurationDto** | [**optional.Interface of ExtensionConfigurationDto**](ExtensionConfigurationDto.md)| The JSON body of the request. Contains updated parameters of the extension configuration. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadExtension

> EntityShortRepresentation UploadExtension(ctx, file, optional)

Uploads a ZIP extension file | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| Extension ZIP file to be uploaded.    File name must match the **name** field in the &#x60;plugin.json&#x60; file.   For example, for the extension whose **name** is &#x60;custom.remote.python.demo&#x60;, the name of the extension file must be &#x60;custom.remote.python.demo.zip&#x60;. | 
 **optional** | ***UploadExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideAlerts** | **optional.Bool**| Use extension-defined thresholds for alerts (&#x60;true&#x60;) or user-defined thresholds (&#x60;false&#x60;).    Extension-defined thresholds are stored in the &#x60;plugin.json&#x60; file.   If not set, user-defined thresholds are used. | [default to false]

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateExtension

> ValidateExtension(ctx, file, optional)

Validates a ZIP extension file for `POST/extensions` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| The ZIP extension file to be uploaded.    The file name must match the ID of the extension. Example: &#x60;custom.remote.python.demo.zip&#x60; | 
 **optional** | ***ValidateExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonOnly** | **optional.Bool**| Validate only the &#x60;plugin.json&#x60; file (&#x60;true&#x60;) or the entire extension structure (&#x60;false&#x60;).    If not set, the entire extension is validated.  | [default to false]

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateLocalExtensionConfiguration

> ValidateLocalExtensionConfiguration(ctx, id, optional)

Validates the payload for the `POST /extensions/{id}/instances` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the extension. | 
 **optional** | ***ValidateLocalExtensionConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateLocalExtensionConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionConfigurationDto** | [**optional.Interface of ExtensionConfigurationDto**](ExtensionConfigurationDto.md)| The JSON body of the request. Contains new configuration of the extension to be validated. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

