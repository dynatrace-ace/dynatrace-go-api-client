# \AnomalyDetectionMetricEventsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricEvent**](AnomalyDetectionMetricEventsApi.md#CreateMetricEvent) | **Post** /anomalyDetection/metricEvents | Creates a new metric event | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateMetricEvent**](AnomalyDetectionMetricEventsApi.md#CreateOrUpdateMetricEvent) | **Put** /anomalyDetection/metricEvents/{id} | Updates an existing metric event or creates a new one | maturity&#x3D;EARLY_ADOPTER
[**DeleteMetricEvent**](AnomalyDetectionMetricEventsApi.md#DeleteMetricEvent) | **Delete** /anomalyDetection/metricEvents/{id} | Deletes the specified metric event | maturity&#x3D;EARLY_ADOPTER
[**GetAllMetricEventConfigs**](AnomalyDetectionMetricEventsApi.md#GetAllMetricEventConfigs) | **Get** /anomalyDetection/metricEvents | Lists all configured metric events | maturity&#x3D;EARLY_ADOPTER
[**GetSingleMetricEventConfig**](AnomalyDetectionMetricEventsApi.md#GetSingleMetricEventConfig) | **Get** /anomalyDetection/metricEvents/{id} | Gets the properties of the specified metric event | maturity&#x3D;EARLY_ADOPTER
[**ValidateMetricEvent**](AnomalyDetectionMetricEventsApi.md#ValidateMetricEvent) | **Post** /anomalyDetection/metricEvents/{id}/validator | Validates the payload for the &#x60;PUT /anomalyDetection/metricEvents/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateMetricEvent1**](AnomalyDetectionMetricEventsApi.md#ValidateMetricEvent1) | **Post** /anomalyDetection/metricEvents/validator | Validates the payload for the &#x60;POST /anomalyDetection/metricEvents&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateMetricEvent

> EntityShortRepresentation CreateMetricEvent(ctx, optional)

Creates a new metric event | maturity=EARLY_ADOPTER

The response contains the ID of the newly created metric event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateMetricEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMetricEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricEvent** | [**optional.Interface of MetricEvent**](MetricEvent.md)| The JSON body of the request. Contains parameters of the new metric event. | 

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


## CreateOrUpdateMetricEvent

> EntityShortRepresentation CreateOrUpdateMetricEvent(ctx, id, optional)

Updates an existing metric event or creates a new one | maturity=EARLY_ADOPTER

If the metric event with the specified ID does not exist, a new metric event will be created.    You can't update the **type** of the metric event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the metric event to be updated.   If you also set the ID in the body, it must match this ID. | 
 **optional** | ***CreateOrUpdateMetricEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateMetricEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricEvent** | [**optional.Interface of MetricEvent**](MetricEvent.md)| The JSON body of the request. Contains updated parameters of the metric event. | 

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


## DeleteMetricEvent

> DeleteMetricEvent(ctx, id)

Deletes the specified metric event | maturity=EARLY_ADOPTER

You can't delete events created by plugins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the metric event to be deleted. | 

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


## GetAllMetricEventConfigs

> StubList GetAllMetricEventConfigs(ctx, optional)

Lists all configured metric events | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMetricEventConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllMetricEventConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeEntityFilterMetricEvents** | **optional.Bool**| Flag to include metric events with associated entities to the response.    Metric events with **entity** filters aren&#39;t compatible across environments. If set to &#x60;false&#x60;, metric events with entity filters will be removed. | [default to false]

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


## GetSingleMetricEventConfig

> MetricEvent GetSingleMetricEventConfig(ctx, id)

Gets the properties of the specified metric event | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required metric event. | 

### Return type

[**MetricEvent**](MetricEvent.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMetricEvent

> ValidateMetricEvent(ctx, id, optional)

Validates the payload for the `PUT /anomalyDetection/metricEvents/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the metric event to be validated. | 
 **optional** | ***ValidateMetricEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateMetricEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricEvent** | [**optional.Interface of MetricEvent**](MetricEvent.md)| The JSON body of the request. Contains the metric event configuration to validate. | 

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


## ValidateMetricEvent1

> ValidateMetricEvent1(ctx, optional)

Validates the payload for the `POST /anomalyDetection/metricEvents` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateMetricEvent1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateMetricEvent1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricEvent** | [**optional.Interface of MetricEvent**](MetricEvent.md)| The JSON body of the request. Contains the metric event configuration to validate. | 

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

