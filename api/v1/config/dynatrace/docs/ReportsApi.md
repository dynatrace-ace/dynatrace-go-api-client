# \ReportsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateReport**](ReportsApi.md#CreateOrUpdateReport) | **Put** /reports/{id} | Updates an existing report | maturity&#x3D;EARLY_ADOPTER
[**CreateReport**](ReportsApi.md#CreateReport) | **Post** /reports | Creates a report | maturity&#x3D;EARLY_ADOPTER
[**DeleteReport**](ReportsApi.md#DeleteReport) | **Delete** /reports/{id} | Deletes the specified report | maturity&#x3D;EARLY_ADOPTER
[**GetReport**](ReportsApi.md#GetReport) | **Get** /reports/{id} | Gets the properties of the specified report | maturity&#x3D;EARLY_ADOPTER
[**ListReports**](ReportsApi.md#ListReports) | **Get** /reports | Lists all reports available in your environment | maturity&#x3D;EARLY_ADOPTER
[**SubscribeReport**](ReportsApi.md#SubscribeReport) | **Post** /reports/{id}/subscribe | Subscribes to the specified report | maturity&#x3D;EARLY_ADOPTER
[**UnsubscribeReport**](ReportsApi.md#UnsubscribeReport) | **Post** /reports/{id}/unsubscribe | Unsubscribes from the specified report | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateOrUpdateReport**](ReportsApi.md#ValidateCreateOrUpdateReport) | **Post** /reports/{id}/validator | Validates the payload for the &#x60;PUT /reports/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateReport**](ReportsApi.md#ValidateCreateReport) | **Post** /reports/validator | Validates the payload for the &#x60;POST /reports&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateReport

> EntityShortRepresentation CreateOrUpdateReport(ctx, id, optional)

Updates an existing report | maturity=EARLY_ADOPTER

If a report with the specified ID doesn't exist, a new report is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the report to be updated.    The ID in the request body must match this ID. | 
 **optional** | ***CreateOrUpdateReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardReport** | [**optional.Interface of DashboardReport**](DashboardReport.md)| The JSON body of the request. Contains updated parameters of the report. | 

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


## CreateReport

> EntityShortRepresentation CreateReport(ctx, optional)

Creates a report | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardReport** | [**optional.Interface of DashboardReport**](DashboardReport.md)| The JSON body of the request. Contains parameters of the new report. | 

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


## DeleteReport

> DeleteReport(ctx, id)

Deletes the specified report | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the report to be deleted. | 

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


## GetReport

> DashboardReport GetReport(ctx, id)

Gets the properties of the specified report | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required report. | 

### Return type

[**DashboardReport**](DashboardReport.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportStubList ListReports(ctx, optional)

Lists all reports available in your environment | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListReportsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Type of a report. | 
 **sourceId** | **optional.String**| Referencing source entity of a report (e.g. dashboard). | 

### Return type

[**ReportStubList**](ReportStubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeReport

> EntityShortRepresentation SubscribeReport(ctx, id, optional)

Subscribes to the specified report | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the report to subscribe to. | 
 **optional** | ***SubscribeReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscribeReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportSubscriptions** | [**optional.Interface of ReportSubscriptions**](ReportSubscriptions.md)| The JSON body of the request. Contains a list of new subscribers. | 

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


## UnsubscribeReport

> EntityShortRepresentation UnsubscribeReport(ctx, id, optional)

Unsubscribes from the specified report | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the report to unsubscribe from. | 
 **optional** | ***UnsubscribeReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UnsubscribeReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportSubscriptions** | [**optional.Interface of ReportSubscriptions**](ReportSubscriptions.md)| The JSON body of the request. Contains a list of recipients to be unsubscribed. | 

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


## ValidateCreateOrUpdateReport

> ValidateCreateOrUpdateReport(ctx, id, optional)

Validates the payload for the `PUT /reports/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the report to be validated.    The ID in the request body must match this ID. | 
 **optional** | ***ValidateCreateOrUpdateReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateOrUpdateReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardReport** | [**optional.Interface of DashboardReport**](DashboardReport.md)| The JSON body of the request. Contains the report to be validated. | 

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


## ValidateCreateReport

> ValidateCreateReport(ctx, optional)

Validates the payload for the `POST /reports` request | maturity=EARLY_ADOPTER

The body must not provide an ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateCreateReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardReport** | [**optional.Interface of DashboardReport**](DashboardReport.md)| The JSON body of the request. Contains the report to be validated. | 

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

