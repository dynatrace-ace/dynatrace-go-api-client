# \CalculatedMetricsSyntheticApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create2**](CalculatedMetricsSyntheticApi.md#Create2) | **Post** /calculatedMetrics/synthetic | Stores the provided calculated synthetic metric configuration. | maturity&#x3D;EARLY_ADOPTER
[**Delete5**](CalculatedMetricsSyntheticApi.md#Delete5) | **Delete** /calculatedMetrics/synthetic/{metricKey} | Deletes the calculated synthetic metric configuration with the given id. | maturity&#x3D;EARLY_ADOPTER
[**GetItem4**](CalculatedMetricsSyntheticApi.md#GetItem4) | **Get** /calculatedMetrics/synthetic/{metricKey} | Gets the definition of the specified calculated synthetic metric. | maturity&#x3D;EARLY_ADOPTER
[**GetList4**](CalculatedMetricsSyntheticApi.md#GetList4) | **Get** /calculatedMetrics/synthetic | Lists all calculated synthetic metric configurations. | maturity&#x3D;EARLY_ADOPTER
[**Update2**](CalculatedMetricsSyntheticApi.md#Update2) | **Put** /calculatedMetrics/synthetic/{metricKey} | Updates the specified calculated synthetic metric. | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#ValidateCreateSyntheticMetric) | **Post** /calculatedMetrics/synthetic/validator | Validates the payload for the &#x60;POST /calculatedMetrics/synthetic&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#ValidateUpdateSyntheticMetric) | **Post** /calculatedMetrics/synthetic/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/synthetic/{metricKey}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## Create2

> EntityShortRepresentation Create2(ctx, optional)

Stores the provided calculated synthetic metric configuration. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Create2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticMetric** | [**optional.Interface of SyntheticMetric**](SyntheticMetric.md)| JSON body of the request containing definition of the new calculated synthetic metric. | 

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


## Delete5

> Delete5(ctx, metricKey)

Deletes the calculated synthetic metric configuration with the given id. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated synthetic metric to delete. | 

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


## GetItem4

> SyntheticMetric GetItem4(ctx, metricKey)

Gets the definition of the specified calculated synthetic metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated synthetic metric you&#39;re inquiring. | 

### Return type

[**SyntheticMetric**](SyntheticMetric.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList4

> StubList GetList4(ctx, )

Lists all calculated synthetic metric configurations. | maturity=EARLY_ADOPTER

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


## Update2

> EntityShortRepresentation Update2(ctx, metricKey, optional)

Updates the specified calculated synthetic metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated synthetic metric to update. | 
 **optional** | ***Update2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticMetricUpdate** | [**optional.Interface of SyntheticMetricUpdate**](SyntheticMetricUpdate.md)| JSON body of the request containing updated definition of the calculated synthetic metric. | 

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


## ValidateCreateSyntheticMetric

> ValidateCreateSyntheticMetric(ctx, syntheticMetric)

Validates the payload for the `POST /calculatedMetrics/synthetic` request | maturity=EARLY_ADOPTER

The body must not provide an ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syntheticMetric** | [**SyntheticMetric**](SyntheticMetric.md)|  | 

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


## ValidateUpdateSyntheticMetric

> ValidateUpdateSyntheticMetric(ctx, metricKey, syntheticMetricUpdate)

Validates the payload for the `PUT /calculatedMetrics/synthetic/{metricKey}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**|  | 
**syntheticMetricUpdate** | [**SyntheticMetricUpdate**](SyntheticMetricUpdate.md)|  | 

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

