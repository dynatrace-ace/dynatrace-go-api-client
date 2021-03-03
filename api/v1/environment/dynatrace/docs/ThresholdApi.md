# \ThresholdApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomThreshold**](ThresholdApi.md#CreateCustomThreshold) | **Put** /thresholds/{thresholdId} | Creates or updates an existing threshold for a plugin or a custom event
[**DeleteCustomThreshold**](ThresholdApi.md#DeleteCustomThreshold) | **Delete** /thresholds/{thresholdId} | Deletes the specified threshold
[**ReadCustomThresholds**](ThresholdApi.md#ReadCustomThresholds) | **Get** /thresholds | Gets all configured thresholds for plugins and custom events in your environment



## CreateCustomThreshold

> Threshold CreateCustomThreshold(ctx, thresholdId).ThresholdRegistrationMessage(thresholdRegistrationMessage).Execute()

Creates or updates an existing threshold for a plugin or a custom event

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
    thresholdId := "thresholdId_example" // string | The ID of the threshold to create or update.
    thresholdRegistrationMessage := *openapiclient.NewThresholdRegistrationMessage() // ThresholdRegistrationMessage | JSON body of the request, containing threshold parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThresholdApi.CreateCustomThreshold(context.Background(), thresholdId).ThresholdRegistrationMessage(thresholdRegistrationMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThresholdApi.CreateCustomThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomThreshold`: Threshold
    fmt.Fprintf(os.Stdout, "Response from `ThresholdApi.CreateCustomThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** | The ID of the threshold to create or update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thresholdRegistrationMessage** | [**ThresholdRegistrationMessage**](ThresholdRegistrationMessage.md) | JSON body of the request, containing threshold parameters. | 

### Return type

[**Threshold**](Threshold.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomThreshold

> DeleteCustomThreshold(ctx, thresholdId).Execute()

Deletes the specified threshold

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
    thresholdId := "thresholdId_example" // string | The ID of the threshold to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThresholdApi.DeleteCustomThreshold(context.Background(), thresholdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThresholdApi.DeleteCustomThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thresholdId** | **string** | The ID of the threshold to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomThresholdRequest struct via the builder pattern


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


## ReadCustomThresholds

> []Threshold ReadCustomThresholds(ctx).Filter(filter).Execute()

Gets all configured thresholds for plugins and custom events in your environment

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
    filter := "filter_example" // string | Filters thresholds by the source. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThresholdApi.ReadCustomThresholds(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThresholdApi.ReadCustomThresholds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCustomThresholds`: []Threshold
    fmt.Fprintf(os.Stdout, "Response from `ThresholdApi.ReadCustomThresholds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadCustomThresholdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters thresholds by the source. | 

### Return type

[**[]Threshold**](Threshold.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

