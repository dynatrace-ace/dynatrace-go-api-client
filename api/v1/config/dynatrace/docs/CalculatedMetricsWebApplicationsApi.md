# \CalculatedMetricsWebApplicationsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create1**](CalculatedMetricsWebApplicationsApi.md#Create1) | **Post** /calculatedMetrics/rum | Stores the provided calculated RUM metric configuration. | maturity&#x3D;EARLY_ADOPTER
[**Delete3**](CalculatedMetricsWebApplicationsApi.md#Delete3) | **Delete** /calculatedMetrics/rum/{metricKey} | Deletes the calculated RUM metric configuration with the given id. | maturity&#x3D;EARLY_ADOPTER
[**GetItem2**](CalculatedMetricsWebApplicationsApi.md#GetItem2) | **Get** /calculatedMetrics/rum/{metricKey} | Gets the definition of the specified calculated RUM metric. | maturity&#x3D;EARLY_ADOPTER
[**GetList2**](CalculatedMetricsWebApplicationsApi.md#GetList2) | **Get** /calculatedMetrics/rum | Lists all calculated RUM metric configurations. | maturity&#x3D;EARLY_ADOPTER
[**Update1**](CalculatedMetricsWebApplicationsApi.md#Update1) | **Put** /calculatedMetrics/rum/{metricKey} | Updates the specified calculated RUM metric. | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateRumMetric**](CalculatedMetricsWebApplicationsApi.md#ValidateCreateRumMetric) | **Post** /calculatedMetrics/rum/validator | Validates the payload for the &#x60;POST /calculatedMetrics/rum&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateRumMetric**](CalculatedMetricsWebApplicationsApi.md#ValidateUpdateRumMetric) | **Post** /calculatedMetrics/rum/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/rum/{metricKey}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## Create1

> EntityShortRepresentation Create1(ctx, optional)

Stores the provided calculated RUM metric configuration. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Create1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rumMetric** | [**optional.Interface of RumMetric**](RumMetric.md)| JSON body of the request containing definition of the new calculated RUM metric. | 

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


## Delete3

> Delete3(ctx, metricKey)

Deletes the calculated RUM metric configuration with the given id. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated RUM metric to delete. | 

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


## GetItem2

> RumMetric GetItem2(ctx, metricKey)

Gets the definition of the specified calculated RUM metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated RUM metric you&#39;re inquiring. | 

### Return type

[**RumMetric**](RumMetric.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList2

> StubList GetList2(ctx, )

Lists all calculated RUM metric configurations. | maturity=EARLY_ADOPTER

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


## Update1

> EntityShortRepresentation Update1(ctx, metricKey, optional)

Updates the specified calculated RUM metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated RUM metric to update. | 
 **optional** | ***Update1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rumMetricUpdate** | [**optional.Interface of RumMetricUpdate**](RumMetricUpdate.md)| JSON body of the request containing updated definition of the calculated RUM metric. | 

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


## ValidateCreateRumMetric

> ValidateCreateRumMetric(ctx, rumMetric)

Validates the payload for the `POST /calculatedMetrics/rum` request | maturity=EARLY_ADOPTER

The body must not provide an ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rumMetric** | [**RumMetric**](RumMetric.md)|  | 

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


## ValidateUpdateRumMetric

> ValidateUpdateRumMetric(ctx, metricKey, rumMetricUpdate)

Validates the payload for the `PUT /calculatedMetrics/rum/{metricKey}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**|  | 
**rumMetricUpdate** | [**RumMetricUpdate**](RumMetricUpdate.md)|  | 

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

