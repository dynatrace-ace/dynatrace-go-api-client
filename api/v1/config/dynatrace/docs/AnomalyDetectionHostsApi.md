# \AnomalyDetectionHostsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostEventsConfig**](AnomalyDetectionHostsApi.md#GetHostEventsConfig) | **Get** /anomalyDetection/hosts | Gets the configuration of anomaly detection for hosts
[**UpdateHostEventsConfig**](AnomalyDetectionHostsApi.md#UpdateHostEventsConfig) | **Put** /anomalyDetection/hosts | Updates the configuration of anomaly detection for hosts
[**ValidateHostEventsConfig**](AnomalyDetectionHostsApi.md#ValidateHostEventsConfig) | **Post** /anomalyDetection/hosts/validator | Validates the configuration of anomaly detection for hosts for the &#x60;PUT /anomalyDetection/hosts&#x60; request



## GetHostEventsConfig

> HostsAnomalyDetectionConfig GetHostEventsConfig(ctx, )

Gets the configuration of anomaly detection for hosts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostEventsConfig

> UpdateHostEventsConfig(ctx, optional)

Updates the configuration of anomaly detection for hosts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateHostEventsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHostEventsConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**optional.Interface of HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the host anomaly detection configuration. | 

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


## ValidateHostEventsConfig

> ValidateHostEventsConfig(ctx, optional)

Validates the configuration of anomaly detection for hosts for the `PUT /anomalyDetection/hosts` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateHostEventsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateHostEventsConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**optional.Interface of HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)| The JSON body of the request. Contains parameters of the host anomaly detection configuration. | 

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

