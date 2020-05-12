# \CalculatedMetricsMobileCustomApplicationsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CalculatedMetricsMobileCustomApplicationsApi.md#Create) | **Post** /calculatedMetrics/mobile | Stores the provided calculated mobile metric configuration. | maturity&#x3D;EARLY_ADOPTER
[**Delete2**](CalculatedMetricsMobileCustomApplicationsApi.md#Delete2) | **Delete** /calculatedMetrics/mobile/{metricKey} | Deletes the calculated mobile metric configuration with the given id. | maturity&#x3D;EARLY_ADOPTER
[**GetItem1**](CalculatedMetricsMobileCustomApplicationsApi.md#GetItem1) | **Get** /calculatedMetrics/mobile/{metricKey} | Gets the definition of the specified calculated mobile metric. | maturity&#x3D;EARLY_ADOPTER
[**GetList1**](CalculatedMetricsMobileCustomApplicationsApi.md#GetList1) | **Get** /calculatedMetrics/mobile | Lists all calculated mobile metric configurations. | maturity&#x3D;EARLY_ADOPTER
[**Update**](CalculatedMetricsMobileCustomApplicationsApi.md#Update) | **Put** /calculatedMetrics/mobile/{metricKey} | Updates the specified calculated mobile metric. | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateMobileMetric**](CalculatedMetricsMobileCustomApplicationsApi.md#ValidateCreateMobileMetric) | **Post** /calculatedMetrics/mobile/validator | Validates the payload for the &#x60;POST /calculatedMetrics/mobile&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateMobileMetric**](CalculatedMetricsMobileCustomApplicationsApi.md#ValidateUpdateMobileMetric) | **Post** /calculatedMetrics/mobile/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/mobile/{metricKey}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## Create

> EntityShortRepresentation Create(ctx, optional)

Stores the provided calculated mobile metric configuration. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileMetric** | [**optional.Interface of MobileMetric**](MobileMetric.md)| JSON body of the request containing definition of the new calculated mobile metric. | 

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


## Delete2

> Delete2(ctx, metricKey)

Deletes the calculated mobile metric configuration with the given id. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated mobile metric to delete. | 

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


## GetItem1

> MobileMetric GetItem1(ctx, metricKey)

Gets the definition of the specified calculated mobile metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated mobile metric you&#39;re inquiring. | 

### Return type

[**MobileMetric**](MobileMetric.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList1

> StubList GetList1(ctx, )

Lists all calculated mobile metric configurations. | maturity=EARLY_ADOPTER

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


## Update

> EntityShortRepresentation Update(ctx, metricKey, optional)

Updates the specified calculated mobile metric. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**| The ID of the calculated mobile metric to update. | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileMetricUpdate** | [**optional.Interface of MobileMetricUpdate**](MobileMetricUpdate.md)| JSON body of the request containing updated definition of the calculated mobile metric. | 

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


## ValidateCreateMobileMetric

> ValidateCreateMobileMetric(ctx, mobileMetric)

Validates the payload for the `POST /calculatedMetrics/mobile` request | maturity=EARLY_ADOPTER

The body must not provide an ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileMetric** | [**MobileMetric**](MobileMetric.md)|  | 

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


## ValidateUpdateMobileMetric

> ValidateUpdateMobileMetric(ctx, metricKey, mobileMetricUpdate)

Validates the payload for the `PUT /calculatedMetrics/mobile/{metricKey}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string**|  | 
**mobileMetricUpdate** | [**MobileMetricUpdate**](MobileMetricUpdate.md)|  | 

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

