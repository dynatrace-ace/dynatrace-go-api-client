# \DashboardsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /dashboards | Creates a dashboard | maturity&#x3D;EARLY_ADOPTER
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /dashboards/{id} | Deletes the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /dashboards/{id} | Gets the properties of the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**GetDashboardMetadata**](DashboardsApi.md#GetDashboardMetadata) | **Get** /dashboards | Lists all dashboards of the environment | maturity&#x3D;EARLY_ADOPTER
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Put** /dashboards/{id} | Updates or creates a dashboard | maturity&#x3D;EARLY_ADOPTER
[**ValidateDashboardCreation**](DashboardsApi.md#ValidateDashboardCreation) | **Post** /dashboards/validator | Validates the payload for the &#39;POST /dashboards&#39; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateDashboardUpdate**](DashboardsApi.md#ValidateDashboardUpdate) | **Post** /dashboards/{id}/validator | Validates the payload for the &#39;PUT /dashboards/{id}&#39; request | maturity&#x3D;EARLY_ADOPTER



## CreateDashboard

> EntityShortRepresentation CreateDashboard(ctx, optional)

Creates a dashboard | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The JSON body of the request. Contains parameters of the new dashboard. | 

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


## DeleteDashboard

> DeleteDashboard(ctx, id)

Deletes the specified dashboard | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the dashboard to be deleted. | 

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


## GetDashboard

> Dashboard GetDashboard(ctx, id)

Gets the properties of the specified dashboard | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardMetadata

> DashboardList GetDashboardMetadata(ctx, )

Lists all dashboards of the environment | maturity=EARLY_ADOPTER

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> EntityShortRepresentation UpdateDashboard(ctx, id, optional)

Updates or creates a dashboard | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the dashboard to be updated.    The ID in the request body provides must match this ID. | 
 **optional** | ***UpdateDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The JSON body of the request. Contains updated parameters of the dashboard. | 

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


## ValidateDashboardCreation

> ValidateDashboardCreation(ctx, optional)

Validates the payload for the 'POST /dashboards' request | maturity=EARLY_ADOPTER

The body must not provide an ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDashboardCreationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDashboardCreationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The JSON body of the request. Containing the dashboard to be validated. | 

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


## ValidateDashboardUpdate

> ValidateDashboardUpdate(ctx, id, optional)

Validates the payload for the 'PUT /dashboards/{id}' request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the dashboard to be validated.    The ID in the request body provides must match this ID. | 
 **optional** | ***ValidateDashboardUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDashboardUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The JSON body of the request. Contains the dashboard to be validated. | 

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

