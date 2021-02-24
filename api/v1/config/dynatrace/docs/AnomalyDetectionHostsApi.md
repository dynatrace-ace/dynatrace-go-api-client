# \AnomalyDetectionHostsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostEventsConfig**](AnomalyDetectionHostsApi.md#GetHostEventsConfig) | **Get** /anomalyDetection/hosts | Gets the configuration of anomaly detection for hosts
[**UpdateHostEventsConfig**](AnomalyDetectionHostsApi.md#UpdateHostEventsConfig) | **Put** /anomalyDetection/hosts | Updates the configuration of anomaly detection for hosts
[**ValidateHostEventsConfig**](AnomalyDetectionHostsApi.md#ValidateHostEventsConfig) | **Post** /anomalyDetection/hosts/validator | Validates the configuration of anomaly detection for hosts for the &#x60;PUT /anomalyDetection/hosts&#x60; request



## GetHostEventsConfig

> HostsAnomalyDetectionConfig GetHostEventsConfig(ctx).Execute()

Gets the configuration of anomaly detection for hosts

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
    resp, r, err := api_client.AnomalyDetectionHostsApi.GetHostEventsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionHostsApi.GetHostEventsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostEventsConfig`: HostsAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionHostsApi.GetHostEventsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostEventsConfigRequest struct via the builder pattern


### Return type

[**HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostEventsConfig

> UpdateHostEventsConfig(ctx).HostsAnomalyDetectionConfig(hostsAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for hosts

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
    hostsAnomalyDetectionConfig := *openapiclient.NewHostsAnomalyDetectionConfig(*openapiclient.NewConnectionLostDetectionConfig(false, false), *openapiclient.NewHighCpuSaturationDetectionConfig(false), *openapiclient.NewHighMemoryDetectionConfig(false), *openapiclient.NewHighGcActivityDetectionConfig(false), *openapiclient.NewOutOfMemoryDetectionConfig(false), *openapiclient.NewOutOfThreadsDetectionConfig(false), *openapiclient.NewNetworkDroppedPacketsDetectionConfig(false), *openapiclient.NewNetworkErrorsDetectionConfig(false), *openapiclient.NewHighNetworkDetectionConfig(false), *openapiclient.NewNetworkTcpProblemsDetectionConfig(false), *openapiclient.NewNetworkHighRetransmissionDetectionConfig(false), *openapiclient.NewDiskLowSpaceDetectionConfig(false), *openapiclient.NewDiskSlowWritesAndReadsDetectionConfig(false), *openapiclient.NewDiskLowInodesDetectionConfig(false)) // HostsAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the host anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionHostsApi.UpdateHostEventsConfig(context.Background()).HostsAnomalyDetectionConfig(hostsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionHostsApi.UpdateHostEventsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostEventsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the host anomaly detection configuration. | 

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


## ValidateHostEventsConfig

> ValidateHostEventsConfig(ctx).HostsAnomalyDetectionConfig(hostsAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for hosts for the `PUT /anomalyDetection/hosts` request

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
    hostsAnomalyDetectionConfig := *openapiclient.NewHostsAnomalyDetectionConfig(*openapiclient.NewConnectionLostDetectionConfig(false, false), *openapiclient.NewHighCpuSaturationDetectionConfig(false), *openapiclient.NewHighMemoryDetectionConfig(false), *openapiclient.NewHighGcActivityDetectionConfig(false), *openapiclient.NewOutOfMemoryDetectionConfig(false), *openapiclient.NewOutOfThreadsDetectionConfig(false), *openapiclient.NewNetworkDroppedPacketsDetectionConfig(false), *openapiclient.NewNetworkErrorsDetectionConfig(false), *openapiclient.NewHighNetworkDetectionConfig(false), *openapiclient.NewNetworkTcpProblemsDetectionConfig(false), *openapiclient.NewNetworkHighRetransmissionDetectionConfig(false), *openapiclient.NewDiskLowSpaceDetectionConfig(false), *openapiclient.NewDiskSlowWritesAndReadsDetectionConfig(false), *openapiclient.NewDiskLowInodesDetectionConfig(false)) // HostsAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the host anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionHostsApi.ValidateHostEventsConfig(context.Background()).HostsAnomalyDetectionConfig(hostsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionHostsApi.ValidateHostEventsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateHostEventsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the host anomaly detection configuration. | 

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

