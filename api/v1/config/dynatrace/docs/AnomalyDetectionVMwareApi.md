# \AnomalyDetectionVMwareApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#GetVMwareAnomalyDetectionConfig) | **Get** /anomalyDetection/vmware | Gets the configuration of anomaly detection for VMware
[**UpdateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#UpdateVMwareAnomalyDetectionConfig) | **Put** /anomalyDetection/vmware | Updates the configuration of anomaly detection for VMware
[**ValidateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#ValidateVMwareAnomalyDetectionConfig) | **Post** /anomalyDetection/vmware/validator | Validates the configuration of anomaly detection for VMware for the &#x60;PUT /anomalyDetection/vmware&#x60; request



## GetVMwareAnomalyDetectionConfig

> VMwareAnomalyDetectionConfig GetVMwareAnomalyDetectionConfig(ctx, )

Gets the configuration of anomaly detection for VMware

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMwareAnomalyDetectionConfig

> UpdateVMwareAnomalyDetectionConfig(ctx, optional)

Updates the configuration of anomaly detection for VMware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateVMwareAnomalyDetectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVMwareAnomalyDetectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**optional.Interface of VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)| JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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


## ValidateVMwareAnomalyDetectionConfig

> ValidateVMwareAnomalyDetectionConfig(ctx, optional)

Validates the configuration of anomaly detection for VMware for the `PUT /anomalyDetection/vmware` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateVMwareAnomalyDetectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateVMwareAnomalyDetectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**optional.Interface of VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)| JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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

