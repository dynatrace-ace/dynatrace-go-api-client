# \FrequentIssueDetectionApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration5**](FrequentIssueDetectionApi.md#GetConfiguration5) | **Get** /frequentIssueDetection | Gets the configuration of frequent issue detection
[**UpdateConfiguration3**](FrequentIssueDetectionApi.md#UpdateConfiguration3) | **Put** /frequentIssueDetection | Updates the configuration of frequent issue detection
[**ValidateConfiguration6**](FrequentIssueDetectionApi.md#ValidateConfiguration6) | **Post** /frequentIssueDetection/validator | Validates the frequent issue detection configuration for the &#x60;PUT /frequentIssueDetection&#x60; request



## GetConfiguration5

> FrequentIssueDetectionConfig GetConfiguration5(ctx, )

Gets the configuration of frequent issue detection

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration3

> UpdateConfiguration3(ctx, optional)

Updates the configuration of frequent issue detection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfiguration3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfiguration3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frequentIssueDetectionConfig** | [**optional.Interface of FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md)| The JSON body of the request, containing parameters of the frequent issue detection configuration | 

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


## ValidateConfiguration6

> ValidateConfiguration6(ctx, optional)

Validates the frequent issue detection configuration for the `PUT /frequentIssueDetection` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frequentIssueDetectionConfig** | [**optional.Interface of FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md)| The JSON body of the request, containing parameters of the frequent issue detection configuration | 

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

