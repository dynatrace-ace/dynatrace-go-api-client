# \TopologySmartscapeCustomDeviceApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomDataPoints**](TopologySmartscapeCustomDeviceApi.md#CreateCustomDataPoints) | **Post** /entity/infrastructure/custom/{customDeviceId} | Creates or updates a custom device, or reports metric data points to the custom device.



## CreateCustomDataPoints

> CustomDevicePushResult CreateCustomDataPoints(ctx, customDeviceId).CustomDevicePushMessage(customDevicePushMessage).Execute()

Creates or updates a custom device, or reports metric data points to the custom device.

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
    customDeviceId := "customDeviceId_example" // string | The ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated.
    customDevicePushMessage := *openapiclient.NewCustomDevicePushMessage() // CustomDevicePushMessage | The JSON body of the request. Contains parameters of a custom device. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints(context.Background(), customDeviceId).CustomDevicePushMessage(customDevicePushMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomDataPoints`: CustomDevicePushResult
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomDataPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDevicePushMessage** | [**CustomDevicePushMessage**](CustomDevicePushMessage.md) | The JSON body of the request. Contains parameters of a custom device. | 

### Return type

[**CustomDevicePushResult**](CustomDevicePushResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

