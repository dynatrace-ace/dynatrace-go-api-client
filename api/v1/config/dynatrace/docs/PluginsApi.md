# \PluginsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemotePluginEndpoint**](PluginsApi.md#CreateRemotePluginEndpoint) | **Post** /plugins/{id}/endpoints | Creates a new endpoint for the specified ActiveGate plugin
[**DeletePlugin**](PluginsApi.md#DeletePlugin) | **Delete** /plugins/{id}/binary | Deletes the ZIP file of the specified plugin
[**DeleteRemotePluginEndpoint**](PluginsApi.md#DeleteRemotePluginEndpoint) | **Delete** /plugins/{id}/endpoints/{endpointId} | Deletes an existing endpoint of the ActiveGate plugin
[**GetPlugin**](PluginsApi.md#GetPlugin) | **Get** /plugins/{id} | Lists the properties of the specified plugin
[**GetPluginBinary**](PluginsApi.md#GetPluginBinary) | **Get** /plugins/{id}/binary | Downloads the ZIP file of the specified plugin
[**GetPluginStates**](PluginsApi.md#GetPluginStates) | **Get** /plugins/{id}/states | Lists the states of the specified plugin
[**GetPlugins**](PluginsApi.md#GetPlugins) | **Get** /plugins | Lists all uploaded plugins
[**GetRemotePluginEndpoint**](PluginsApi.md#GetRemotePluginEndpoint) | **Get** /plugins/{id}/endpoints/{endpointId} | Gets parameters of the specified endpoint of the ActiveGate plugin
[**GetRemotePluginEndpoints**](PluginsApi.md#GetRemotePluginEndpoints) | **Get** /plugins/{id}/endpoints | Lists endpoints of the specified ActiveGate plugin
[**GetRemotePluginModules**](PluginsApi.md#GetRemotePluginModules) | **Get** /plugins/activeGatePluginModules | List available ActiveGate plugin modules
[**UpdateRemotePluginEndpoint**](PluginsApi.md#UpdateRemotePluginEndpoint) | **Put** /plugins/{id}/endpoints/{endpointId} | Updates an existing endpoint of the ActiveGate plugin
[**UploadPlugin**](PluginsApi.md#UploadPlugin) | **Post** /plugins | Uploads a ZIP plugin file
[**ValidatePlugin**](PluginsApi.md#ValidatePlugin) | **Post** /plugins/validator | Validates a ZIP plugin file for &#x60;POST/plugins&#x60; request
[**ValidateRemotePluginEndpoint**](PluginsApi.md#ValidateRemotePluginEndpoint) | **Post** /plugins/{id}/endpoints/validator | Validates the payload for the &#x60;POST /plugins/{id}/endpoints&#x60; request



## CreateRemotePluginEndpoint

> EntityShortRepresentation CreateRemotePluginEndpoint(ctx, id, optional)

Creates a new endpoint for the specified ActiveGate plugin

The body must not provide an ID of the endpoint, as IDs are automatically assigned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the plugin where you want to create an endpoint. | 
 **optional** | ***CreateRemotePluginEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRemotePluginEndpointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePluginEndpoint** | [**optional.Interface of RemotePluginEndpoint**](RemotePluginEndpoint.md)| The JSON body of the request. Contains parameters of the new plugin endpoint. | 

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


## DeletePlugin

> DeletePlugin(ctx, id)

Deletes the ZIP file of the specified plugin

Deletion of the plugin file uninstalls the plugin, making it unavailable for use.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the plugin to be deleted. | 

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


## DeleteRemotePluginEndpoint

> DeleteRemotePluginEndpoint(ctx, id, endpointId)

Deletes an existing endpoint of the ActiveGate plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the plugin where you want to delete an endpoint. | 
**endpointId** | **string**| The ID of the endpoint to be deleted. | 

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


## GetPlugin

> Plugin GetPlugin(ctx, id)

Lists the properties of the specified plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required plugin. | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginBinary

> map[string]interface{} GetPluginBinary(ctx, id)

Downloads the ZIP file of the specified plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the plugin you want to download. | 

### Return type

**map[string]interface{}**

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginStates

> PluginStateList GetPluginStates(ctx, id)

Lists the states of the specified plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required plugin. | 

### Return type

[**PluginStateList**](PluginStateList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> StubList GetPlugins(ctx, )

Lists all uploaded plugins

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


## GetRemotePluginEndpoint

> RemotePluginEndpoint GetRemotePluginEndpoint(ctx, id, endpointId)

Gets parameters of the specified endpoint of the ActiveGate plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required plugin. | 
**endpointId** | **string**| The ID of the required endpoint. | 

### Return type

[**RemotePluginEndpoint**](RemotePluginEndpoint.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemotePluginEndpoints

> StubList GetRemotePluginEndpoints(ctx, id)

Lists endpoints of the specified ActiveGate plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required plugin. | 

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


## GetRemotePluginModules

> StubList GetRemotePluginModules(ctx, )

List available ActiveGate plugin modules

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


## UpdateRemotePluginEndpoint

> UpdateRemotePluginEndpoint(ctx, id, endpointId, optional)

Updates an existing endpoint of the ActiveGate plugin

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the plugin where you want to update an endpoint.   If you set the plugin ID in the body as well, it must match this ID. | 
**endpointId** | **string**| The ID of the endpoint to be updated.   If you set the endpoint ID in the body as well, it must match this ID. | 
 **optional** | ***UpdateRemotePluginEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRemotePluginEndpointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remotePluginEndpoint** | [**optional.Interface of RemotePluginEndpoint**](RemotePluginEndpoint.md)| The JSON body of the request. Contains updated parameters of the plugin endpoint. | 

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


## UploadPlugin

> EntityShortRepresentation UploadPlugin(ctx, file, optional)

Uploads a ZIP plugin file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| Plugin ZIP file to be uploaded.    File name must match the **name** field in the &#x60;plugin.json&#x60; file.   For example, for the plugin whose **name** is &#x60;custom.remote.python.demo&#x60;, the name of the plugin file must be &#x60;custom.remote.python.demo.zip&#x60;. | 
 **optional** | ***UploadPluginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadPluginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideAlerts** | **optional.Bool**| Use plugin-defined thresholds for alerts (&#x60;true&#x60;) or user-defined thresholds (&#x60;false&#x60;).    Plugin-defined thresholds are stored in the &#x60;plugin.json&#x60; file.   If not set, user-defined thresholds are used. | [default to false]

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


## ValidatePlugin

> ValidatePlugin(ctx, file, optional)

Validates a ZIP plugin file for `POST/plugins` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| The ZIP plugin file to be uploaded.    The file name must match the ID of the plugin. Example: &#x60;custom.remote.python.demo.zip&#x60; | 
 **optional** | ***ValidatePluginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePluginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonOnly** | **optional.Bool**| Validate only the &#x60;plugin.json&#x60; file (&#x60;true&#x60;) or the entire plugin structure (&#x60;false&#x60;).    If not set, the entire plugin is validated.  | [default to false]

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


## ValidateRemotePluginEndpoint

> ValidateRemotePluginEndpoint(ctx, id, optional)

Validates the payload for the `POST /plugins/{id}/endpoints` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required plugin. | 
 **optional** | ***ValidateRemotePluginEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateRemotePluginEndpointOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePluginEndpoint** | [**optional.Interface of RemotePluginEndpoint**](RemotePluginEndpoint.md)| The JSON body of the request. Contains parameters of the new plugin endpoint. | 

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

