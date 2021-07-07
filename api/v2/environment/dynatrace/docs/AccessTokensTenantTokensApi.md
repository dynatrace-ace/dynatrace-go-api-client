# \AccessTokensTenantTokensApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRotation**](AccessTokensTenantTokensApi.md#CancelRotation) | **Post** /tenantTokenRotation/cancel | Cancels tenant token rotation
[**FinishRotation**](AccessTokensTenantTokensApi.md#FinishRotation) | **Post** /tenantTokenRotation/finish | Finishes tenant token rotation
[**StartRotation**](AccessTokensTenantTokensApi.md#StartRotation) | **Post** /tenantTokenRotation/start | Starts tenant token rotation



## CancelRotation

> TenantTokenConfig CancelRotation(ctx).Execute()

Cancels tenant token rotation



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
    resp, r, err := api_client.AccessTokensTenantTokensApi.CancelRotation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensTenantTokensApi.CancelRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRotation`: TenantTokenConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensTenantTokensApi.CancelRotation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRotationRequest struct via the builder pattern


### Return type

[**TenantTokenConfig**](TenantTokenConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishRotation

> TenantTokenConfig FinishRotation(ctx).Execute()

Finishes tenant token rotation



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
    resp, r, err := api_client.AccessTokensTenantTokensApi.FinishRotation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensTenantTokensApi.FinishRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinishRotation`: TenantTokenConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensTenantTokensApi.FinishRotation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinishRotationRequest struct via the builder pattern


### Return type

[**TenantTokenConfig**](TenantTokenConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRotation

> TenantTokenConfig StartRotation(ctx).Execute()

Starts tenant token rotation



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
    resp, r, err := api_client.AccessTokensTenantTokensApi.StartRotation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensTenantTokensApi.StartRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartRotation`: TenantTokenConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensTenantTokensApi.StartRotation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartRotationRequest struct via the builder pattern


### Return type

[**TenantTokenConfig**](TenantTokenConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

