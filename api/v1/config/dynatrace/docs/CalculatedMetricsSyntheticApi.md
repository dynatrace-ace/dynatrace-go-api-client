# \CalculatedMetricsSyntheticApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#CreateSyntheticMetric) | **Post** /calculatedMetrics/synthetic | Creates a new calculated synthetic metric | maturity&#x3D;EARLY_ADOPTER
[**DeleteSyntheticMetric**](CalculatedMetricsSyntheticApi.md#DeleteSyntheticMetric) | **Delete** /calculatedMetrics/synthetic/{metricKey} | Deletes the specified calculated synthetic metric | maturity&#x3D;EARLY_ADOPTER
[**GetSyntheticMetric**](CalculatedMetricsSyntheticApi.md#GetSyntheticMetric) | **Get** /calculatedMetrics/synthetic/{metricKey} | Gets the descriptor of the specified calculated synthetic metric | maturity&#x3D;EARLY_ADOPTER
[**ListSyntheticMetrics**](CalculatedMetricsSyntheticApi.md#ListSyntheticMetrics) | **Get** /calculatedMetrics/synthetic | Lists all calculated synthetic metrics available in your environment | maturity&#x3D;EARLY_ADOPTER
[**UpdateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#UpdateSyntheticMetric) | **Put** /calculatedMetrics/synthetic/{metricKey} | Updates the specified calculated synthetic metric | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#ValidateCreateSyntheticMetric) | **Post** /calculatedMetrics/synthetic/validator | Validates the payload for the &#x60;POST /calculatedMetrics/synthetic&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateSyntheticMetric**](CalculatedMetricsSyntheticApi.md#ValidateUpdateSyntheticMetric) | **Post** /calculatedMetrics/synthetic/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/synthetic/{metricKey}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateSyntheticMetric

> EntityShortRepresentation CreateSyntheticMetric(ctx).CalculatedSyntheticMetric(calculatedSyntheticMetric).Execute()

Creates a new calculated synthetic metric | maturity=EARLY_ADOPTER

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
    calculatedSyntheticMetric := *openapiclient.NewCalculatedSyntheticMetric("MonitorIdentifier_example", "Name_example", "MetricKey_example", false, "Metric_example") // CalculatedSyntheticMetric | The JSON body of the request. Contains definition of the new calculated synthetic metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.CreateSyntheticMetric(context.Background()).CalculatedSyntheticMetric(calculatedSyntheticMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.CreateSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyntheticMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsSyntheticApi.CreateSyntheticMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyntheticMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedSyntheticMetric** | [**CalculatedSyntheticMetric**](CalculatedSyntheticMetric.md) | The JSON body of the request. Contains definition of the new calculated synthetic metric. | 

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


## DeleteSyntheticMetric

> DeleteSyntheticMetric(ctx, metricKey).Execute()

Deletes the specified calculated synthetic metric | maturity=EARLY_ADOPTER

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
    metricKey := "metricKey_example" // string | The key of the calculated synthetic metric to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.DeleteSyntheticMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.DeleteSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated synthetic metric to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyntheticMetricRequest struct via the builder pattern


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


## GetSyntheticMetric

> CalculatedSyntheticMetric GetSyntheticMetric(ctx, metricKey).Execute()

Gets the descriptor of the specified calculated synthetic metric | maturity=EARLY_ADOPTER

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
    metricKey := "metricKey_example" // string | The key of the required calculated synthetic metric.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.GetSyntheticMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.GetSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticMetric`: CalculatedSyntheticMetric
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsSyntheticApi.GetSyntheticMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required calculated synthetic metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalculatedSyntheticMetric**](CalculatedSyntheticMetric.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyntheticMetrics

> StubList ListSyntheticMetrics(ctx).Execute()

Lists all calculated synthetic metrics available in your environment | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.ListSyntheticMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.ListSyntheticMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyntheticMetrics`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsSyntheticApi.ListSyntheticMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSyntheticMetricsRequest struct via the builder pattern


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


## UpdateSyntheticMetric

> UpdateSyntheticMetric(ctx, metricKey).SyntheticMetricUpdate(syntheticMetricUpdate).Execute()

Updates the specified calculated synthetic metric | maturity=EARLY_ADOPTER

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
    metricKey := "metricKey_example" // string | The key of the calculated synthetic metric to be updated.
    syntheticMetricUpdate := *openapiclient.NewSyntheticMetricUpdate() // SyntheticMetricUpdate | The JSON body of the request. Contains the update to the calculated synthetic metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.UpdateSyntheticMetric(context.Background(), metricKey).SyntheticMetricUpdate(syntheticMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.UpdateSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated synthetic metric to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyntheticMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticMetricUpdate** | [**SyntheticMetricUpdate**](SyntheticMetricUpdate.md) | The JSON body of the request. Contains the update to the calculated synthetic metric. | 

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


## ValidateCreateSyntheticMetric

> ValidateCreateSyntheticMetric(ctx).CalculatedSyntheticMetric(calculatedSyntheticMetric).Execute()

Validates the payload for the `POST /calculatedMetrics/synthetic` request | maturity=EARLY_ADOPTER

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
    calculatedSyntheticMetric := *openapiclient.NewCalculatedSyntheticMetric("MonitorIdentifier_example", "Name_example", "MetricKey_example", false, "Metric_example") // CalculatedSyntheticMetric | The JSON body of the request. Contains the definition of the metric to be validated.   The key of the metric must be omitted. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.ValidateCreateSyntheticMetric(context.Background()).CalculatedSyntheticMetric(calculatedSyntheticMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.ValidateCreateSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateSyntheticMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedSyntheticMetric** | [**CalculatedSyntheticMetric**](CalculatedSyntheticMetric.md) | The JSON body of the request. Contains the definition of the metric to be validated.   The key of the metric must be omitted. | 

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


## ValidateUpdateSyntheticMetric

> ValidateUpdateSyntheticMetric(ctx, metricKey).SyntheticMetricUpdate(syntheticMetricUpdate).Execute()

Validates the payload for the `PUT /calculatedMetrics/synthetic/{metricKey}` request | maturity=EARLY_ADOPTER

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
    metricKey := "metricKey_example" // string | The key of the metric to be validated.
    syntheticMetricUpdate := *openapiclient.NewSyntheticMetricUpdate() // SyntheticMetricUpdate | The JSON body of the request. Contains the update to the metric to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CalculatedMetricsSyntheticApi.ValidateUpdateSyntheticMetric(context.Background(), metricKey).SyntheticMetricUpdate(syntheticMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsSyntheticApi.ValidateUpdateSyntheticMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the metric to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateSyntheticMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticMetricUpdate** | [**SyntheticMetricUpdate**](SyntheticMetricUpdate.md) | The JSON body of the request. Contains the update to the metric to be validated. | 

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

