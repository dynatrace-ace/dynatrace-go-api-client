# \AnomalyDetectionDiskEventsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#CreateDiskEventConfig) | **Post** /anomalyDetection/diskEvents | Creates a new disk event rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#DeleteDiskEventConfig) | **Delete** /anomalyDetection/diskEvents/{id} | Deletes the specified disk event rule | maturity&#x3D;EARLY_ADOPTER
[**GetDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#GetDiskEventConfig) | **Get** /anomalyDetection/diskEvents/{id} | Gets the properties of the specified disk event rule | maturity&#x3D;EARLY_ADOPTER
[**ListDiskEventConfigs**](AnomalyDetectionDiskEventsApi.md#ListDiskEventConfigs) | **Get** /anomalyDetection/diskEvents | Lists all existing disk event rules | maturity&#x3D;EARLY_ADOPTER
[**UpdateOrCreateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#UpdateOrCreateDiskEventConfig) | **Put** /anomalyDetection/diskEvents/{id} | Updates or creates a disk event rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#ValidateDiskEventConfig) | **Post** /anomalyDetection/diskEvents/validator | Validates the payload for the &#x60;POST /anomalyDetection/diskEvents&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#ValidateDiskEventConfig1) | **Post** /anomalyDetection/diskEvents/{id}/validator | Validates the payload for the &#x60;PUT /anomalyDetection/diskEvents/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateDiskEventConfig

> EntityShortRepresentation CreateDiskEventConfig(ctx, optional)

Creates a new disk event rule | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDiskEventConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDiskEventConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)| JSON body of the request. Contains parameters of the new disk event rule. | 

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


## DeleteDiskEventConfig

> DeleteDiskEventConfig(ctx, id)

Deletes the specified disk event rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the disk event rule to be deleted. | 

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


## GetDiskEventConfig

> DiskEventAnomalyDetectionConfig GetDiskEventConfig(ctx, id)

Gets the properties of the specified disk event rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required disk event rule. | 

### Return type

[**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskEventConfigs

> StubList ListDiskEventConfigs(ctx, )

Lists all existing disk event rules | maturity=EARLY_ADOPTER

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


## UpdateOrCreateDiskEventConfig

> EntityShortRepresentation UpdateOrCreateDiskEventConfig(ctx, id, optional)

Updates or creates a disk event rule | maturity=EARLY_ADOPTER

If a disk event rule with the specified ID does not exist, a new rule is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the disk event rule to be updated. | 
 **optional** | ***UpdateOrCreateDiskEventConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateDiskEventConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)| JSON body of the request. Contains updated disk event rule parameters. | 

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


## ValidateDiskEventConfig

> ValidateDiskEventConfig(ctx, optional)

Validates the payload for the `POST /anomalyDetection/diskEvents` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDiskEventConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDiskEventConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)| JSON body of the request. Contains the disk event rule to be validated. | 

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


## ValidateDiskEventConfig1

> ValidateDiskEventConfig1(ctx, id, optional)

Validates the payload for the `PUT /anomalyDetection/diskEvents/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the disk event rule to be validated. | 
 **optional** | ***ValidateDiskEventConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDiskEventConfig1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)| JSON body of the request. Contains the disk event rule to be validated. | 

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

