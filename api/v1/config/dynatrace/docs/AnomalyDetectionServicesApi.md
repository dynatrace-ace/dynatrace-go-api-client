# \AnomalyDetectionServicesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration2**](AnomalyDetectionServicesApi.md#GetConfiguration2) | **Get** /anomalyDetection/services | Gets the service anomaly detection configuration
[**UpdateConfiguration2**](AnomalyDetectionServicesApi.md#UpdateConfiguration2) | **Put** /anomalyDetection/services | Updates the service anomaly detection configuration
[**ValidateConfiguration2**](AnomalyDetectionServicesApi.md#ValidateConfiguration2) | **Post** /anomalyDetection/services/validator | Validates the payload for the &#x60;PUT /anomalyDetection/services&#x60; request



## GetConfiguration2

> ServiceAnomalyDetectionConfig GetConfiguration2(ctx, )

Gets the service anomaly detection configuration

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration2

> UpdateConfiguration2(ctx, optional)

Updates the service anomaly detection configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfiguration2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAnomalyDetectionConfig** | [**optional.Interface of ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the service anomaly detection configuration. | 

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


## ValidateConfiguration2

> ValidateConfiguration2(ctx, optional)

Validates the payload for the `PUT /anomalyDetection/services` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAnomalyDetectionConfig** | [**optional.Interface of ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the service anomaly detection configuration. | 

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

