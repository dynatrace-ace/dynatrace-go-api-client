# \RUMContentResourcesApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentConfig**](RUMContentResourcesApi.md#GetContentConfig) | **Get** /contentResources | Gets the configuration of content resources
[**PutContentResourceConfig**](RUMContentResourcesApi.md#PutContentResourceConfig) | **Put** /contentResources | Updates the configuration of content resources
[**ValidateContentResourceConfig**](RUMContentResourcesApi.md#ValidateContentResourceConfig) | **Post** /contentResources/validator | Validates the payload for the &#x60;PUT /contentResources&#x60; request



## GetContentConfig

> ContentResources GetContentConfig(ctx).Execute()

Gets the configuration of content resources

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
    resp, r, err := api_client.RUMContentResourcesApi.GetContentConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMContentResourcesApi.GetContentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentConfig`: ContentResources
    fmt.Fprintf(os.Stdout, "Response from `RUMContentResourcesApi.GetContentConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentConfigRequest struct via the builder pattern


### Return type

[**ContentResources**](ContentResources.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutContentResourceConfig

> PutContentResourceConfig(ctx).ContentResources(contentResources).Execute()

Updates the configuration of content resources

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
    contentResources := *openapiclient.NewContentResources() // ContentResources | The JSON body of the request. Contains the configuration of content resources. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMContentResourcesApi.PutContentResourceConfig(context.Background()).ContentResources(contentResources).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMContentResourcesApi.PutContentResourceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutContentResourceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentResources** | [**ContentResources**](ContentResources.md) | The JSON body of the request. Contains the configuration of content resources. | 

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


## ValidateContentResourceConfig

> ValidateContentResourceConfig(ctx).ContentResources(contentResources).Execute()

Validates the payload for the `PUT /contentResources` request

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
    contentResources := *openapiclient.NewContentResources() // ContentResources | The JSON body of the request. Contains the configuration of content resources to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMContentResourcesApi.ValidateContentResourceConfig(context.Background()).ContentResources(contentResources).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMContentResourcesApi.ValidateContentResourceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateContentResourceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentResources** | [**ContentResources**](ContentResources.md) | The JSON body of the request. Contains the configuration of content resources to be validated. | 

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

