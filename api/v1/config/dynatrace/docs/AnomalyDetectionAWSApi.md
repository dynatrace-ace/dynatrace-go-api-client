# \AnomalyDetectionAWSApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#GetAwsAnomalyDetectionConfig) | **Get** /anomalyDetection/aws | Gets the configuration of anomaly detection for AWS
[**UpdateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#UpdateAwsAnomalyDetectionConfig) | **Put** /anomalyDetection/aws | Updates the configuration of anomaly detection for AWS
[**ValidateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#ValidateAwsAnomalyDetectionConfig) | **Post** /anomalyDetection/aws/validator | Validates the configuration of anomaly detection for AWS for the &#x60;PUT /anomalyDetection/aws&#x60; request



## GetAwsAnomalyDetectionConfig

> AwsAnomalyDetectionConfig GetAwsAnomalyDetectionConfig(ctx, )

Gets the configuration of anomaly detection for AWS

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsAnomalyDetectionConfig

> UpdateAwsAnomalyDetectionConfig(ctx, optional)

Updates the configuration of anomaly detection for AWS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateAwsAnomalyDetectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAwsAnomalyDetectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**optional.Interface of AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)| JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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


## ValidateAwsAnomalyDetectionConfig

> ValidateAwsAnomalyDetectionConfig(ctx, optional)

Validates the configuration of anomaly detection for AWS for the `PUT /anomalyDetection/aws` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateAwsAnomalyDetectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAwsAnomalyDetectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**optional.Interface of AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)| JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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

