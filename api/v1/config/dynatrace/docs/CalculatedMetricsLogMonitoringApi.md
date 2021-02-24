# \CalculatedMetricsLogMonitoringApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#DeleteLogMetricConfig) | **Delete** /calculatedMetrics/log/{metricKey} | Deletes the specified custom log metric definition | maturity&#x3D;EARLY_ADOPTER
[**GetLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#GetLogMetricConfig) | **Get** /calculatedMetrics/log/{metricKey} | Gets the definition of the specified custom log metric | maturity&#x3D;EARLY_ADOPTER
[**ListLogMetricConfigs**](CalculatedMetricsLogMonitoringApi.md#ListLogMetricConfigs) | **Get** /calculatedMetrics/log | Lists all custom log metrics configured in your environment | maturity&#x3D;EARLY_ADOPTER
[**UpdateOrCreateLogMetricConfig**](CalculatedMetricsLogMonitoringApi.md#UpdateOrCreateLogMetricConfig) | **Put** /calculatedMetrics/log/{metricKey} | Creates a new custom log metric | maturity&#x3D;EARLY_ADOPTER
[**ValidateMetric**](CalculatedMetricsLogMonitoringApi.md#ValidateMetric) | **Post** /calculatedMetrics/log/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/log/{metricKey}&#x60; request. | maturity&#x3D;EARLY_ADOPTER



## DeleteLogMetricConfig

> DeleteLogMetricConfig(ctx, metricKey).Execute()

Deletes the specified custom log metric definition | maturity=EARLY_ADOPTER

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricKey := "metricKey_example" // string | The key of the custom log metric to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsLogMonitoringApi.DeleteLogMetricConfig(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsLogMonitoringApi.DeleteLogMetricConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the custom log metric to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogMetricConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogMetricConfig

> LogMetricConfig GetLogMetricConfig(ctx, metricKey).Execute()

Gets the definition of the specified custom log metric | maturity=EARLY_ADOPTER

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricKey := "metricKey_example" // string | The key of the required custom log metric.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsLogMonitoringApi.GetLogMetricConfig(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsLogMonitoringApi.GetLogMetricConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogMetricConfig`: LogMetricConfig
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsLogMonitoringApi.GetLogMetricConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required custom log metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogMetricConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogMetricConfig**](LogMetricConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogMetricConfigs

> StubList ListLogMetricConfigs(ctx).Execute()

Lists all custom log metrics configured in your environment | maturity=EARLY_ADOPTER

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsLogMonitoringApi.ListLogMetricConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsLogMonitoringApi.ListLogMetricConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogMetricConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsLogMonitoringApi.ListLogMetricConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogMetricConfigsRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateLogMetricConfig

> EntityShortRepresentation UpdateOrCreateLogMetricConfig(ctx, metricKey).LogMetricConfig(logMetricConfig).Execute()

Creates a new custom log metric | maturity=EARLY_ADOPTER



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricKey := "metricKey_example" // string | The required key of the custom log metric. The key must have the `calc:log.` prefix.    The key in the body of the request must match this key.
    logMetricConfig := *openapiclient.NewLogMetricConfig("calc:log.metric", "DisplayName_example", "Unit_example", "prefix1* OR prefix2*", "OCCURRENCES", []openapiclient.LogSourceFilter{*openapiclient.NewLogSourceFilter()}) // LogMetricConfig | The JSON body of the request. Contains the definition of the custom log metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsLogMonitoringApi.UpdateOrCreateLogMetricConfig(context.Background(), metricKey).LogMetricConfig(logMetricConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsLogMonitoringApi.UpdateOrCreateLogMetricConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrCreateLogMetricConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsLogMonitoringApi.UpdateOrCreateLogMetricConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The required key of the custom log metric. The key must have the &#x60;calc:log.&#x60; prefix.    The key in the body of the request must match this key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateLogMetricConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logMetricConfig** | [**LogMetricConfig**](LogMetricConfig.md) | The JSON body of the request. Contains the definition of the custom log metric. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMetric

> ValidateMetric(ctx, metricKey).LogMetricConfig(logMetricConfig).Execute()

Validates the payload for the `PUT /calculatedMetrics/log/{metricKey}` request. | maturity=EARLY_ADOPTER

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricKey := "metricKey_example" // string | The key of the custom log metric to be validated.
    logMetricConfig := *openapiclient.NewLogMetricConfig("calc:log.metric", "DisplayName_example", "Unit_example", "prefix1* OR prefix2*", "OCCURRENCES", []openapiclient.LogSourceFilter{*openapiclient.NewLogSourceFilter()}) // LogMetricConfig | The JSON body of the request. Contains definition of the custom log metric to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsLogMonitoringApi.ValidateMetric(context.Background(), metricKey).LogMetricConfig(logMetricConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsLogMonitoringApi.ValidateMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the custom log metric to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logMetricConfig** | [**LogMetricConfig**](LogMetricConfig.md) | The JSON body of the request. Contains definition of the custom log metric to be validated. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

