# \ManagementZonesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagementZone**](ManagementZonesApi.md#CreateManagementZone) | **Post** /managementZones | Creates a new management zone
[**CreateOrUpdateManagementZone**](ManagementZonesApi.md#CreateOrUpdateManagementZone) | **Put** /managementZones/{id} | Updates the specified management zone or creates it
[**DeleteManagementZone**](ManagementZonesApi.md#DeleteManagementZone) | **Delete** /managementZones/{id} | Deletes the specified management zone
[**GetAllManagementZoneConfigs**](ManagementZonesApi.md#GetAllManagementZoneConfigs) | **Get** /managementZones | Lists all configured management zones
[**GetSingleManagementZoneConfig**](ManagementZonesApi.md#GetSingleManagementZoneConfig) | **Get** /managementZones/{id} | Lists the parameters of the specified management zone
[**ValidateManagementZone**](ManagementZonesApi.md#ValidateManagementZone) | **Post** /managementZones/validator | Validates a new management zone for the &#x60;POST /managementZones&#x60; request
[**ValidateManagementZone1**](ManagementZonesApi.md#ValidateManagementZone1) | **Post** /managementZones/{id}/validator | Validate updates of existing management zone for the &#x60;PUT /managementZones/{id}&#x60; request



## CreateManagementZone

> EntityShortRepresentation CreateManagementZone(ctx, optional)

Creates a new management zone

The body must not provide an ID as IDs are automatically assigned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateManagementZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateManagementZoneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)| The JSON body of the request, containing parameters of the management zone. | 

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


## CreateOrUpdateManagementZone

> EntityShortRepresentation CreateOrUpdateManagementZone(ctx, id, optional)

Updates the specified management zone or creates it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the management zone to update.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateManagementZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateManagementZoneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)| The JSON body of the request, containing updated parameters of the management zone. | 

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


## DeleteManagementZone

> DeleteManagementZone(ctx, id)

Deletes the specified management zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the management zone to delete. | 

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


## GetAllManagementZoneConfigs

> StubList GetAllManagementZoneConfigs(ctx, )

Lists all configured management zones

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


## GetSingleManagementZoneConfig

> ManagementZone GetSingleManagementZoneConfig(ctx, id, optional)

Lists the parameters of the specified management zone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required management zone. | 
 **optional** | ***GetSingleManagementZoneConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSingleManagementZoneConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed. | [default to false]

### Return type

[**ManagementZone**](ManagementZone.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateManagementZone

> ValidateManagementZone(ctx, optional)

Validates a new management zone for the `POST /managementZones` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateManagementZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateManagementZoneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)| The JSON body of the request, containing the management zone parameters to validate. | 

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


## ValidateManagementZone1

> ValidateManagementZone1(ctx, id, optional)

Validate updates of existing management zone for the `PUT /managementZones/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the management zone to validate. | 
 **optional** | ***ValidateManagementZone1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateManagementZone1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)| The JSON body of the request, containing the management zone parameters to validate. | 

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

