# \MaintenanceWindowsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceWindow**](MaintenanceWindowsApi.md#DeleteMaintenanceWindow) | **Delete** /maintenanceWindows/{id} | Deletes the specified maintenance window
[**GetMaintenanceWindow**](MaintenanceWindowsApi.md#GetMaintenanceWindow) | **Get** /maintenanceWindows/{id} | Gets the properties of the specified maintenance window
[**ReturnAllMaintenanceWindows**](MaintenanceWindowsApi.md#ReturnAllMaintenanceWindows) | **Get** /maintenanceWindows | Lists all available maintenance windows
[**StoreMaintenanceWindow**](MaintenanceWindowsApi.md#StoreMaintenanceWindow) | **Put** /maintenanceWindows/{id} | Updates an existing maintenance window
[**StoreMaintenanceWindow1**](MaintenanceWindowsApi.md#StoreMaintenanceWindow1) | **Post** /maintenanceWindows | Creates a new maintenance window
[**ValidateConfiguration7**](MaintenanceWindowsApi.md#ValidateConfiguration7) | **Post** /maintenanceWindows/{id}/validator | Validates the payload for the &#x60;PUT /maintenancewindow/{id}&#x60; request
[**ValidateConfiguration8**](MaintenanceWindowsApi.md#ValidateConfiguration8) | **Post** /maintenanceWindows/validator | Validates the payload for the &#x60;POST /maintenancewindow&#x60; request



## DeleteMaintenanceWindow

> DeleteMaintenanceWindow(ctx, id)

Deletes the specified maintenance window

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be deleted. | 

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


## GetMaintenanceWindow

> MaintenanceWindow GetMaintenanceWindow(ctx, id)

Gets the properties of the specified maintenance window

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required maintenance window. | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllMaintenanceWindows

> StubList ReturnAllMaintenanceWindows(ctx, )

Lists all available maintenance windows

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


## StoreMaintenanceWindow

> EntityShortRepresentation StoreMaintenanceWindow(ctx, id, optional)

Updates an existing maintenance window

If a maintenance window with the specified ID doesn't exist, a new one is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be updated. | 
 **optional** | ***StoreMaintenanceWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreMaintenanceWindowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains updated parameters of the maintenance window. | 

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


## StoreMaintenanceWindow1

> EntityShortRepresentation StoreMaintenanceWindow1(ctx, optional)

Creates a new maintenance window

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StoreMaintenanceWindow1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreMaintenanceWindow1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the new maintenance window. | 

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


## ValidateConfiguration7

> ValidateConfiguration7(ctx, id, optional)

Validates the payload for the `PUT /maintenancewindow/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be validated. | 
 **optional** | ***ValidateConfiguration7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the maintenance window to be validated. | 

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


## ValidateConfiguration8

> ValidateConfiguration8(ctx, optional)

Validates the payload for the `POST /maintenancewindow` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration8Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the maintenance window be validated. | 

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

