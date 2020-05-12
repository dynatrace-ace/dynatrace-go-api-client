# \AnomalyDetectionDatabaseServicesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration1**](AnomalyDetectionDatabaseServicesApi.md#GetConfiguration1) | **Get** /anomalyDetection/databaseServices | Gets the configuration of anomaly detection for database services
[**UpdateConfiguration1**](AnomalyDetectionDatabaseServicesApi.md#UpdateConfiguration1) | **Put** /anomalyDetection/databaseServices | Updates the configuration of anomaly detection for database services
[**ValidateConfiguration1**](AnomalyDetectionDatabaseServicesApi.md#ValidateConfiguration1) | **Post** /anomalyDetection/databaseServices/validator | Validates the payload for the &#x60;PUT /anomalyDetection/databaseServices&#x60; request



## GetConfiguration1

> DatabaseAnomalyDetectionConfig GetConfiguration1(ctx, )

Gets the configuration of anomaly detection for database services

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration1

> UpdateConfiguration1(ctx, optional)

Updates the configuration of anomaly detection for database services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfiguration1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**optional.Interface of DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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


## ValidateConfiguration1

> ValidateConfiguration1(ctx, optional)

Validates the payload for the `PUT /anomalyDetection/databaseServices` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**optional.Interface of DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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

