# \CalculatedMetricsServicesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete4**](CalculatedMetricsServicesApi.md#Delete4) | **Delete** /calculatedMetrics/service/{metricKey} | Deletes the specified calculated service metric | maturity&#x3D;EARLY_ADOPTER
[**GetItem3**](CalculatedMetricsServicesApi.md#GetItem3) | **Get** /calculatedMetrics/service/{metricKey} | Gets the descriptor of the specified calculated service metric | maturity&#x3D;EARLY_ADOPTER
[**GetList3**](CalculatedMetricsServicesApi.md#GetList3) | **Get** /calculatedMetrics/service | Lists all calculated service metrics configured in your environment | maturity&#x3D;EARLY_ADOPTER
[**Post1**](CalculatedMetricsServicesApi.md#Post1) | **Post** /calculatedMetrics/service | Creates a new calculated service metric | maturity&#x3D;EARLY_ADOPTER
[**Put**](CalculatedMetricsServicesApi.md#Put) | **Put** /calculatedMetrics/service/{metricKey} | Updates the specified calculated service metric | maturity&#x3D;EARLY_ADOPTER
[**ValidateItem2**](CalculatedMetricsServicesApi.md#ValidateItem2) | **Post** /calculatedMetrics/service/validator | Validates the payload for the &#x60;POST /calculatedMetric/service&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidatePut**](CalculatedMetricsServicesApi.md#ValidatePut) | **Post** /calculatedMetrics/service/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetric/service/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## Delete4

> Delete4(ctx, metricKey)

Deletes the specified calculated service metric | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the calculated service metric to be deleted. | 

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


## GetItem3

> CalculatedServiceMetric GetItem3(ctx, metricKey)

Gets the descriptor of the specified calculated service metric | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the required calculated service metric. | 

### Return type

[**CalculatedServiceMetric**](CalculatedServiceMetric.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList3

> StubList GetList3(ctx, )

Lists all calculated service metrics configured in your environment | maturity=EARLY_ADOPTER

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


## Post1

> EntityShortRepresentation Post1(ctx, optional)

Creates a new calculated service metric | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Post1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Post1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedServiceMetric** | [**optional.Interface of CalculatedServiceMetric**](CalculatedServiceMetric.md)| The JSON body of the request. Contains parameters of the new calculated service metric. | 

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


## Put

> EntityShortRepresentation Put(ctx, metricKey, optional)

Updates the specified calculated service metric | maturity=EARLY_ADOPTER

If the calculated service metric with the specified key doesn't exist, a new metric is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the calculated service metric to be updated.    The key of the calculated service metric in the body of the request must match this key. | 
 **optional** | ***PutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedServiceMetric** | [**optional.Interface of CalculatedServiceMetric**](CalculatedServiceMetric.md)| The JSON body of the request. Contains updated parameters of the calculated service metric. | 

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


## ValidateItem2

> ValidateItem2(ctx, optional)

Validates the payload for the `POST /calculatedMetric/service` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateItem2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateItem2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedServiceMetric** | [**optional.Interface of CalculatedServiceMetric**](CalculatedServiceMetric.md)| The JSON body of the request. Contains the calculated service metric to be validated. | 

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


## ValidatePut

> ValidatePut(ctx, metricKey, optional)

Validates the payload for the `PUT /calculatedMetric/service/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The key of the calculated service metric to be validated .   The key of the metric in the body of the request must match this key. | 
 **optional** | ***ValidatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedServiceMetric** | [**optional.Interface of CalculatedServiceMetric**](CalculatedServiceMetric.md)| The JSON body of the request. Contains the calculated service metric to be validated. | 

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

