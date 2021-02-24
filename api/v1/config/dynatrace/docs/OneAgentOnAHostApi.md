# \OneAgentOnAHostApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfig1**](OneAgentOnAHostApi.md#GetAutoUpdateConfig1) | **Get** /hosts/{id}/autoupdate | Gets the configuration of OneAgent auto-update on the specified host
[**GetHostConfig**](OneAgentOnAHostApi.md#GetHostConfig) | **Get** /hosts/{id} | Gets the OneAgent configuration on the specified host
[**GetMonitoringConfig**](OneAgentOnAHostApi.md#GetMonitoringConfig) | **Get** /hosts/{id}/monitoring | Gets the monitoring configuration of OneAgent on the specified host
[**GetTechHostConfigs**](OneAgentOnAHostApi.md#GetTechHostConfigs) | **Get** /hosts/{id}/technologies | Gets the configuration of monitored technologies on the specified host
[**UpdateAutoUpdateConfig**](OneAgentOnAHostApi.md#UpdateAutoUpdateConfig) | **Put** /hosts/{id}/autoupdate | Updates the configuration of OneAgent auto-update on the specified host
[**UpdateMonitoringConfig**](OneAgentOnAHostApi.md#UpdateMonitoringConfig) | **Put** /hosts/{id}/monitoring | Updates the monitoring configuration of OneAgent on the specified host
[**ValidateAutoUpdateConfig1**](OneAgentOnAHostApi.md#ValidateAutoUpdateConfig1) | **Post** /hosts/{id}/autoupdate/validator | Validates the payload for the &#x60;PUT /hosts/{id}/autoupdate&#x60; request
[**ValidateMonitoringConfig**](OneAgentOnAHostApi.md#ValidateMonitoringConfig) | **Post** /hosts/{id}/monitoring/validator | Validates the payload for the &#x60;PUT /hosts/{id}/monitoring&#x60; request



## GetAutoUpdateConfig1

> HostAutoUpdateConfig GetAutoUpdateConfig1(ctx, id).Execute()

Gets the configuration of OneAgent auto-update on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.GetAutoUpdateConfig1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.GetAutoUpdateConfig1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfig1`: HostAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostApi.GetAutoUpdateConfig1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostAutoUpdateConfig**](HostAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostConfig

> HostConfig GetHostConfig(ctx, id).Execute()

Gets the OneAgent configuration on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.GetHostConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.GetHostConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostConfig`: HostConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostApi.GetHostConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostConfig**](HostConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringConfig

> MonitoringConfig GetMonitoringConfig(ctx, id).Execute()

Gets the monitoring configuration of OneAgent on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.GetMonitoringConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.GetMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringConfig`: MonitoringConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostApi.GetMonitoringConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringConfig**](MonitoringConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechHostConfigs

> TechMonitoringList GetTechHostConfigs(ctx, id).Execute()

Gets the configuration of monitored technologies on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.GetTechHostConfigs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.GetTechHostConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTechHostConfigs`: TechMonitoringList
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostApi.GetTechHostConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechHostConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TechMonitoringList**](TechMonitoringList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoUpdateConfig

> UpdateAutoUpdateConfig(ctx, id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()

Updates the configuration of OneAgent auto-update on the specified host



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    hostAutoUpdateConfig := *openapiclient.NewHostAutoUpdateConfig("DISABLED") // HostAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.UpdateAutoUpdateConfig(context.Background(), id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.UpdateAutoUpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostAutoUpdateConfig** | [**HostAutoUpdateConfig**](HostAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters. | 

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


## UpdateMonitoringConfig

> UpdateMonitoringConfig(ctx, id).MonitoringConfig(monitoringConfig).Execute()

Updates the monitoring configuration of OneAgent on the specified host



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    monitoringConfig := *openapiclient.NewMonitoringConfig(true, "FULL_STACK") // MonitoringConfig | The JSON body of the request. Contains OneAgent monitoring parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.UpdateMonitoringConfig(context.Background(), id).MonitoringConfig(monitoringConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.UpdateMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfig** | [**MonitoringConfig**](MonitoringConfig.md) | The JSON body of the request. Contains OneAgent monitoring parameters. | 

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


## ValidateAutoUpdateConfig1

> ValidateAutoUpdateConfig1(ctx, id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()

Validates the payload for the `PUT /hosts/{id}/autoupdate` request

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    hostAutoUpdateConfig := *openapiclient.NewHostAutoUpdateConfig("DISABLED") // HostAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.ValidateAutoUpdateConfig1(context.Background(), id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.ValidateAutoUpdateConfig1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostAutoUpdateConfig** | [**HostAutoUpdateConfig**](HostAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. | 

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


## ValidateMonitoringConfig

> ValidateMonitoringConfig(ctx, id).MonitoringConfig(monitoringConfig).Execute()

Validates the payload for the `PUT /hosts/{id}/monitoring` request

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    monitoringConfig := *openapiclient.NewMonitoringConfig(true, "FULL_STACK") // MonitoringConfig | The JSON body of the request. Contains OneAgent monitoring parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentOnAHostApi.ValidateMonitoringConfig(context.Background(), id).MonitoringConfig(monitoringConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostApi.ValidateMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfig** | [**MonitoringConfig**](MonitoringConfig.md) | The JSON body of the request. Contains OneAgent monitoring parameters. | 

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

