# \CalculatedMetricsWebApplicationsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRumMetric**](CalculatedMetricsWebApplicationsApi.md#CreateRumMetric) | **Post** /calculatedMetrics/rum | Creates a new calculated RUM metric
[**DeleteRumMetric**](CalculatedMetricsWebApplicationsApi.md#DeleteRumMetric) | **Delete** /calculatedMetrics/rum/{metricKey} | Deletes the specified calculated RUM metric
[**GetRumMetric**](CalculatedMetricsWebApplicationsApi.md#GetRumMetric) | **Get** /calculatedMetrics/rum/{metricKey} | Gets the descriptor of the specified calculated RUM metric
[**ListRumMetrics**](CalculatedMetricsWebApplicationsApi.md#ListRumMetrics) | **Get** /calculatedMetrics/rum | Lists all calculated RUM metrics
[**UpdateRumMetric**](CalculatedMetricsWebApplicationsApi.md#UpdateRumMetric) | **Put** /calculatedMetrics/rum/{metricKey} | Updates the specified calculated RUM metric
[**ValidateCreateRumMetric**](CalculatedMetricsWebApplicationsApi.md#ValidateCreateRumMetric) | **Post** /calculatedMetrics/rum/validator | Validates the payload for the &#x60;POST /calculatedMetrics/rum&#x60; request
[**ValidateUpdateRumMetric**](CalculatedMetricsWebApplicationsApi.md#ValidateUpdateRumMetric) | **Post** /calculatedMetrics/rum/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/rum/{metricKey}&#x60; request



## CreateRumMetric

> EntityShortRepresentation CreateRumMetric(ctx).RumMetric(rumMetric).Execute()

Creates a new calculated RUM metric

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
    rumMetric := *openapiclient.NewRumMetric("ApplicationIdentifier_example", "Name_example", "MetricKey_example", false, *openapiclient.NewRumMetricDefinition("Metric_example")) // RumMetric | The JSON body of the request. Contains the definition of the new calculated RUM metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.CreateRumMetric(context.Background()).RumMetric(rumMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.CreateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRumMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsApi.CreateRumMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rumMetric** | [**RumMetric**](RumMetric.md) | The JSON body of the request. Contains the definition of the new calculated RUM metric. | 

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


## DeleteRumMetric

> DeleteRumMetric(ctx, metricKey).Execute()

Deletes the specified calculated RUM metric

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
    metricKey := "metricKey_example" // string | The key of the metric to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.DeleteRumMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.DeleteRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the metric to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRumMetric

> RumMetric GetRumMetric(ctx, metricKey).Execute()

Gets the descriptor of the specified calculated RUM metric

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
    metricKey := "metricKey_example" // string | The key of the required metric.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.GetRumMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.GetRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRumMetric`: RumMetric
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsApi.GetRumMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RumMetric**](RumMetric.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRumMetrics

> StubList ListRumMetrics(ctx).Execute()

Lists all calculated RUM metrics

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
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.ListRumMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.ListRumMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRumMetrics`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsApi.ListRumMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRumMetricsRequest struct via the builder pattern


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


## UpdateRumMetric

> UpdateRumMetric(ctx, metricKey).RumMetricUpdate(rumMetricUpdate).Execute()

Updates the specified calculated RUM metric

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
    metricKey := "metricKey_example" // string | The key of the calculated RUM metric to be updated.
    rumMetricUpdate := *openapiclient.NewRumMetricUpdate() // RumMetricUpdate | The JSON body of the request. Contains the updated parameters of the calculated RUM metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.UpdateRumMetric(context.Background(), metricKey).RumMetricUpdate(rumMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.UpdateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated RUM metric to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rumMetricUpdate** | [**RumMetricUpdate**](RumMetricUpdate.md) | The JSON body of the request. Contains the updated parameters of the calculated RUM metric. | 

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


## ValidateCreateRumMetric

> ValidateCreateRumMetric(ctx).RumMetric(rumMetric).Execute()

Validates the payload for the `POST /calculatedMetrics/rum` request



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
    rumMetric := *openapiclient.NewRumMetric("ApplicationIdentifier_example", "Name_example", "MetricKey_example", false, *openapiclient.NewRumMetricDefinition("Metric_example")) // RumMetric | The JSON body of the request. Contains the parameters of the calculated RUM metric to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.ValidateCreateRumMetric(context.Background()).RumMetric(rumMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.ValidateCreateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rumMetric** | [**RumMetric**](RumMetric.md) | The JSON body of the request. Contains the parameters of the calculated RUM metric to be validated. | 

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


## ValidateUpdateRumMetric

> ValidateUpdateRumMetric(ctx, metricKey).RumMetricUpdate(rumMetricUpdate).Execute()

Validates the payload for the `PUT /calculatedMetrics/rum/{metricKey}` request

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
    metricKey := "metricKey_example" // string | The key of the calculated RUM metric to be validated.
    rumMetricUpdate := *openapiclient.NewRumMetricUpdate() // RumMetricUpdate | The JSON body of the request. Contains the parameters of the calculated RUM metric to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsWebApplicationsApi.ValidateUpdateRumMetric(context.Background(), metricKey).RumMetricUpdate(rumMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsApi.ValidateUpdateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated RUM metric to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rumMetricUpdate** | [**RumMetricUpdate**](RumMetricUpdate.md) | The JSON body of the request. Contains the parameters of the calculated RUM metric to be validated. | 

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

