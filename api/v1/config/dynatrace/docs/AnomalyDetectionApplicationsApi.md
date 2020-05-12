# \AnomalyDetectionApplicationsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](AnomalyDetectionApplicationsApi.md#GetConfiguration) | **Get** /anomalyDetection/applications | Gets the configuration of anomaly detection for applications
[**UpdateConfiguration**](AnomalyDetectionApplicationsApi.md#UpdateConfiguration) | **Put** /anomalyDetection/applications | Updates the configuration of anomaly detection for applications
[**ValidateConfiguration**](AnomalyDetectionApplicationsApi.md#ValidateConfiguration) | **Post** /anomalyDetection/applications/validator | Validates the configuration of anomaly detection for applications for the &#x60;PUT /anomalyDetection/applications&#x60; request



## GetConfiguration

> ApplicationAnomalyDetectionConfig GetConfiguration(ctx, )

Gets the configuration of anomaly detection for applications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> UpdateConfiguration(ctx, optional)

Updates the configuration of anomaly detection for applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAnomalyDetectionConfig** | [**optional.Interface of ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md)| The JSON body of the request, containing parameters of the application anomaly detection configuration. | 

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


## ValidateConfiguration

> ValidateConfiguration(ctx, optional)

Validates the configuration of anomaly detection for applications for the `PUT /anomalyDetection/applications` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAnomalyDetectionConfig** | [**optional.Interface of ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md)| The JSON body of the request, containing parameters of the application anomaly detection configuration. | 

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

