# \AnomalyDetectionDiskEventsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#CreateDiskEventConfig) | **Post** /anomalyDetection/diskEvents | Creates a new disk event rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#DeleteDiskEventConfig) | **Delete** /anomalyDetection/diskEvents/{id} | Deletes the specified disk event rule | maturity&#x3D;EARLY_ADOPTER
[**GetDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#GetDiskEventConfig) | **Get** /anomalyDetection/diskEvents/{id} | Gets the properties of the specified disk event rule | maturity&#x3D;EARLY_ADOPTER
[**ListDiskEventConfigs**](AnomalyDetectionDiskEventsApi.md#ListDiskEventConfigs) | **Get** /anomalyDetection/diskEvents | Lists all existing disk event rules | maturity&#x3D;EARLY_ADOPTER
[**UpdateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#UpdateDiskEventConfig) | **Put** /anomalyDetection/diskEvents/{id} | Updates or creates a disk event rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#ValidateCreateDiskEventConfig) | **Post** /anomalyDetection/diskEvents/validator | Validates the payload for the &#x60;POST /anomalyDetection/diskEvents&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateDiskEventConfig**](AnomalyDetectionDiskEventsApi.md#ValidateUpdateDiskEventConfig) | **Post** /anomalyDetection/diskEvents/{id}/validator | Validates the payload for the &#x60;PUT /anomalyDetection/diskEvents/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateDiskEventConfig

> EntityShortRepresentation CreateDiskEventConfig(ctx).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()

Creates a new disk event rule | maturity=EARLY_ADOPTER



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
    diskEventAnomalyDetectionConfig := *openapiclient.NewDiskEventAnomalyDetectionConfig("Name_example", false, "Metric_example", float64(123), int32(10), int32(8)) // DiskEventAnomalyDetectionConfig | JSON body of the request. Contains parameters of the new disk event rule. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.CreateDiskEventConfig(context.Background()).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.CreateDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiskEventConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDiskEventsApi.CreateDiskEventConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiskEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md) | JSON body of the request. Contains parameters of the new disk event rule. | 

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


## DeleteDiskEventConfig

> DeleteDiskEventConfig(ctx, id).Execute()

Deletes the specified disk event rule | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the disk event rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.DeleteDiskEventConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.DeleteDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the disk event rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiskEventConfigRequest struct via the builder pattern


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


## GetDiskEventConfig

> DiskEventAnomalyDetectionConfig GetDiskEventConfig(ctx, id).Execute()

Gets the properties of the specified disk event rule | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the required disk event rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.GetDiskEventConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.GetDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiskEventConfig`: DiskEventAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDiskEventsApi.GetDiskEventConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required disk event rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiskEventConfigs

> StubList ListDiskEventConfigs(ctx).Execute()

Lists all existing disk event rules | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.ListDiskEventConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.ListDiskEventConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiskEventConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDiskEventsApi.ListDiskEventConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDiskEventConfigsRequest struct via the builder pattern


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


## UpdateDiskEventConfig

> EntityShortRepresentation UpdateDiskEventConfig(ctx, id).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()

Updates or creates a disk event rule | maturity=EARLY_ADOPTER



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
    id := TODO // string | The ID of the disk event rule to be updated.
    diskEventAnomalyDetectionConfig := *openapiclient.NewDiskEventAnomalyDetectionConfig("Name_example", false, "Metric_example", float64(123), int32(10), int32(8)) // DiskEventAnomalyDetectionConfig | JSON body of the request. Contains updated disk event rule parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.UpdateDiskEventConfig(context.Background(), id).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.UpdateDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDiskEventConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDiskEventsApi.UpdateDiskEventConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the disk event rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiskEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskEventAnomalyDetectionConfig** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md) | JSON body of the request. Contains updated disk event rule parameters. | 

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


## ValidateCreateDiskEventConfig

> ValidateCreateDiskEventConfig(ctx).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()

Validates the payload for the `POST /anomalyDetection/diskEvents` request | maturity=EARLY_ADOPTER

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
    diskEventAnomalyDetectionConfig := *openapiclient.NewDiskEventAnomalyDetectionConfig("Name_example", false, "Metric_example", float64(123), int32(10), int32(8)) // DiskEventAnomalyDetectionConfig | JSON body of the request. Contains the disk event rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.ValidateCreateDiskEventConfig(context.Background()).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.ValidateCreateDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateDiskEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md) | JSON body of the request. Contains the disk event rule to be validated. | 

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


## ValidateUpdateDiskEventConfig

> ValidateUpdateDiskEventConfig(ctx, id).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()

Validates the payload for the `PUT /anomalyDetection/diskEvents/{id}` request | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the disk event rule to be validated.
    diskEventAnomalyDetectionConfig := *openapiclient.NewDiskEventAnomalyDetectionConfig("Name_example", false, "Metric_example", float64(123), int32(10), int32(8)) // DiskEventAnomalyDetectionConfig | JSON body of the request. Contains the disk event rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDiskEventsApi.ValidateUpdateDiskEventConfig(context.Background(), id).DiskEventAnomalyDetectionConfig(diskEventAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDiskEventsApi.ValidateUpdateDiskEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the disk event rule to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateDiskEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskEventAnomalyDetectionConfig** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md) | JSON body of the request. Contains the disk event rule to be validated. | 

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

