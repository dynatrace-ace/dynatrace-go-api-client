# \CalculatedMetricsLogMonitoringApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#DeleteLogMetricConfig) | **Delete** /calculatedMetrics/log/{metricKey} | Deletes the specified custom log metric definition | maturity&#x3D;EARLY_ADOPTER
[**GetLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#GetLogMetricConfig) | **Get** /calculatedMetrics/log/{metricKey} | Gets the definition of the specified custom log metric | maturity&#x3D;EARLY_ADOPTER
[**ListLogMetricConfigs**](CalculatedMetricsLogMonitoringApi.md#ListLogMetricConfigs) | **Get** /calculatedMetrics/log | Lists all custom log metrics configured in your environment | maturity&#x3D;EARLY_ADOPTER
[**UpdateOrCreateLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#UpdateOrCreateLogMetricConfig) | **Put** /calculatedMetrics/log/{metricKey} | Creates a new custom log metric | maturity&#x3D;EARLY_ADOPTER
[**ValidateMetric**](CalculatedMetricsLogMonitoringApi.md#ValidateMetric) | **Post** /calculatedMetrics/log/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/log/{metricKey}&#x60; request. | maturity&#x3D;EARLY_ADOPTER



## DeleteLogMetricConfig

> DeleteLogMetricConfig(ctx, metricKey)

Deletes the specified custom log metric definition | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the custom log metric to be deleted. | 

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


## GetLogMetricConfig

> LogMetricConfig GetLogMetricConfig(ctx, metricKey)

Gets the definition of the specified custom log metric | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the required custom log metric. | 

### Return type

[**LogMetricConfig**](LogMetricConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogMetricConfigs

> StubList ListLogMetricConfigs(ctx, )

Lists all custom log metrics configured in your environment | maturity=EARLY_ADOPTER

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


## UpdateOrCreateLogMetricConfig

> EntityShortRepresentation UpdateOrCreateLogMetricConfig(ctx, metricKey, optional)

Creates a new custom log metric | maturity=EARLY_ADOPTER

If the metric definition with the specified key already exists, it is updated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The required key of the custom log metric. The key must have the &#x60;calc:log.&#x60; prefix.    The key in the body of the request must match this key. | 
 **optional** | ***UpdateOrCreateLogMetricConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateLogMetricConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logMetricConfig** | [**optional.Interface of LogMetricConfig**](LogMetricConfig.md)| The JSON body of the request. Contains the definition of the custom log metric. | 

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


## ValidateMetric

> ValidateMetric(ctx, metricKey, optional)

Validates the payload for the `PUT /calculatedMetrics/log/{metricKey}` request. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the custom log metric to be validated. | 
 **optional** | ***ValidateMetricOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateMetricOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logMetricConfig** | [**optional.Interface of LogMetricConfig**](LogMetricConfig.md)| The JSON body of the request. Contains definition of the custom log metric to be validated. | 

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

