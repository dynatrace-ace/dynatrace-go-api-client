# \ActiveGatesAutoUpdateConfigurationApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationApi.md#GetAutoUpdateConfigById) | **Get** /activeGates/{agId}/autoUpdate | Gets the configuration of auto-update for the specified ActiveGate
[**GetGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationApi.md#GetGlobalAutoUpdateConfigForTenant) | **Get** /activeGates/autoUpdate | Gets the global auto-update configuration of environment ActiveGates.
[**PutAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationApi.md#PutAutoUpdateConfigById) | **Put** /activeGates/{agId}/autoUpdate | Updates the configuration of auto-update for the specified ActiveGate
[**PutGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationApi.md#PutGlobalAutoUpdateConfigForTenant) | **Put** /activeGates/autoUpdate | Puts the global auto-update configuration of environment ActiveGates.
[**ValidateAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationApi.md#ValidateAutoUpdateConfigById) | **Post** /activeGates/{agId}/autoUpdate/validator | Validates the payload for the &#x60;POST /activeGates/{agId}/autoUpdate&#x60; request.
[**ValidateGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationApi.md#ValidateGlobalAutoUpdateConfigForTenant) | **Post** /activeGates/autoUpdate/validator | Validates the payload for the &#x60;POST /activeGates/autoUpdate&#x60; request.



## GetAutoUpdateConfigById

> ActiveGateAutoUpdateConfig GetAutoUpdateConfigById(ctx, agId).Execute()

Gets the configuration of auto-update for the specified ActiveGate

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
    agId := "agId_example" // string | The ID of the required ActiveGate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.GetAutoUpdateConfigById(context.Background(), agId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.GetAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfigById`: ActiveGateAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateConfigurationApi.GetAutoUpdateConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalAutoUpdateConfigForTenant

> ActiveGateGlobalAutoUpdateConfig GetGlobalAutoUpdateConfigForTenant(ctx).Execute()

Gets the global auto-update configuration of environment ActiveGates.

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
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.GetGlobalAutoUpdateConfigForTenant(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.GetGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalAutoUpdateConfigForTenant`: ActiveGateGlobalAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateConfigurationApi.GetGlobalAutoUpdateConfigForTenant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


### Return type

[**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutoUpdateConfigById

> PutAutoUpdateConfigById(ctx, agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()

Updates the configuration of auto-update for the specified ActiveGate

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
    agId := "agId_example" // string | The ID of the required ActiveGate.
    activeGateAutoUpdateConfig := *openapiclient.NewActiveGateAutoUpdateConfig("INHERITED") // ActiveGateAutoUpdateConfig | JSON body of the request, containing auto update parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.PutAutoUpdateConfigById(context.Background(), agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.PutAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeGateAutoUpdateConfig** | [**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md) | JSON body of the request, containing auto update parameters. | 

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


## PutGlobalAutoUpdateConfigForTenant

> PutGlobalAutoUpdateConfigForTenant(ctx).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()

Puts the global auto-update configuration of environment ActiveGates.

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
    activeGateGlobalAutoUpdateConfig := *openapiclient.NewActiveGateGlobalAutoUpdateConfig("GlobalSetting_example") // ActiveGateGlobalAutoUpdateConfig | JSON body of the request, containing global auto update parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.PutGlobalAutoUpdateConfigForTenant(context.Background()).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.PutGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeGateGlobalAutoUpdateConfig** | [**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md) | JSON body of the request, containing global auto update parameters. | 

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


## ValidateAutoUpdateConfigById

> ValidateAutoUpdateConfigById(ctx, agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()

Validates the payload for the `POST /activeGates/{agId}/autoUpdate` request.

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
    agId := "agId_example" // string | The ID of the required ActiveGate.
    activeGateAutoUpdateConfig := *openapiclient.NewActiveGateAutoUpdateConfig("INHERITED") // ActiveGateAutoUpdateConfig | JSON body of the request, containing auto update parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.ValidateAutoUpdateConfigById(context.Background(), agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.ValidateAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeGateAutoUpdateConfig** | [**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md) | JSON body of the request, containing auto update parameters. | 

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


## ValidateGlobalAutoUpdateConfigForTenant

> ValidateGlobalAutoUpdateConfigForTenant(ctx).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()

Validates the payload for the `POST /activeGates/autoUpdate` request.

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
    activeGateGlobalAutoUpdateConfig := *openapiclient.NewActiveGateGlobalAutoUpdateConfig("GlobalSetting_example") // ActiveGateGlobalAutoUpdateConfig | JSON body of the request, containing global auto update parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateConfigurationApi.ValidateGlobalAutoUpdateConfigForTenant(context.Background()).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationApi.ValidateGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeGateGlobalAutoUpdateConfig** | [**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md) | JSON body of the request, containing global auto update parameters. | 

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

