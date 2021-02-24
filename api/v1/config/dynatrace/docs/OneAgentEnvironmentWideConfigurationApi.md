# \OneAgentEnvironmentWideConfigurationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfig**](OneAgentEnvironmentWideConfigurationApi.md#GetAutoUpdateConfig) | **Get** /hosts/autoupdate | Gets the environment-wide configuration of OneAgents auto-update
[**GetTechGlobalConfigs**](OneAgentEnvironmentWideConfigurationApi.md#GetTechGlobalConfigs) | **Get** /technologies | Gets the global monitoring configuration of technologies.
[**SetAutoUpdateConfig**](OneAgentEnvironmentWideConfigurationApi.md#SetAutoUpdateConfig) | **Put** /hosts/autoupdate | Updates the environment-wide configuration of OneAgents auto-update
[**ValidateAutoUpdateConfig**](OneAgentEnvironmentWideConfigurationApi.md#ValidateAutoUpdateConfig) | **Post** /hosts/autoupdate/validator | Validates the payload for the &#x60;PUT hosts/autoupdate&#x60; request



## GetAutoUpdateConfig

> EnvironmentAutoUpdateConfig GetAutoUpdateConfig(ctx).Execute()

Gets the environment-wide configuration of OneAgents auto-update

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
    resp, r, err := api_client.OneAgentEnvironmentWideConfigurationApi.GetAutoUpdateConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentEnvironmentWideConfigurationApi.GetAutoUpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfig`: EnvironmentAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentEnvironmentWideConfigurationApi.GetAutoUpdateConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfigRequest struct via the builder pattern


### Return type

[**EnvironmentAutoUpdateConfig**](EnvironmentAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechGlobalConfigs

> TechMonitoringList GetTechGlobalConfigs(ctx).Execute()

Gets the global monitoring configuration of technologies.

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
    resp, r, err := api_client.OneAgentEnvironmentWideConfigurationApi.GetTechGlobalConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentEnvironmentWideConfigurationApi.GetTechGlobalConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTechGlobalConfigs`: TechMonitoringList
    fmt.Fprintf(os.Stdout, "Response from `OneAgentEnvironmentWideConfigurationApi.GetTechGlobalConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechGlobalConfigsRequest struct via the builder pattern


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


## SetAutoUpdateConfig

> SetAutoUpdateConfig(ctx).EnvironmentAutoUpdateConfig(environmentAutoUpdateConfig).Execute()

Updates the environment-wide configuration of OneAgents auto-update



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
    environmentAutoUpdateConfig := *openapiclient.NewEnvironmentAutoUpdateConfig("DISABLED") // EnvironmentAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentEnvironmentWideConfigurationApi.SetAutoUpdateConfig(context.Background()).EnvironmentAutoUpdateConfig(environmentAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentEnvironmentWideConfigurationApi.SetAutoUpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAutoUpdateConfig** | [**EnvironmentAutoUpdateConfig**](EnvironmentAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters. | 

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


## ValidateAutoUpdateConfig

> ValidateAutoUpdateConfig(ctx).EnvironmentAutoUpdateConfig(environmentAutoUpdateConfig).Execute()

Validates the payload for the `PUT hosts/autoupdate` request

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
    environmentAutoUpdateConfig := *openapiclient.NewEnvironmentAutoUpdateConfig("DISABLED") // EnvironmentAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentEnvironmentWideConfigurationApi.ValidateAutoUpdateConfig(context.Background()).EnvironmentAutoUpdateConfig(environmentAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentEnvironmentWideConfigurationApi.ValidateAutoUpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAutoUpdateConfig** | [**EnvironmentAutoUpdateConfig**](EnvironmentAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. | 

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

