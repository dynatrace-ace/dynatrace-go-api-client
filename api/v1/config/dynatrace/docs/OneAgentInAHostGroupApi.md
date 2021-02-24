# \OneAgentInAHostGroupApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfig2**](OneAgentInAHostGroupApi.md#GetAutoUpdateConfig2) | **Get** /hostgroups/{id}/autoupdate | Gets the configuration of OneAgent auto-update in the specified host group
[**GetHostGroupConfig**](OneAgentInAHostGroupApi.md#GetHostGroupConfig) | **Get** /hostgroups/{id} | Gets the OneAgent configuration in the specified host group
[**UpdateAutoUpdateConfig1**](OneAgentInAHostGroupApi.md#UpdateAutoUpdateConfig1) | **Put** /hostgroups/{id}/autoupdate | Updates the configuration of OneAgent auto-update in the specified host group
[**ValidateAutoUpdateConfig2**](OneAgentInAHostGroupApi.md#ValidateAutoUpdateConfig2) | **Post** /hostgroups/{id}/autoupdate/validator | Validates the payload for the &#x60;PUT /hostgroups/{id}/autoupdate&#x60; request



## GetAutoUpdateConfig2

> HostGroupAutoUpdateConfig GetAutoUpdateConfig2(ctx, id).Execute()

Gets the configuration of OneAgent auto-update in the specified host group

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
    id := "id_example" // string | The Dynatrace entity ID of the required host group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentInAHostGroupApi.GetAutoUpdateConfig2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentInAHostGroupApi.GetAutoUpdateConfig2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfig2`: HostGroupAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentInAHostGroupApi.GetAutoUpdateConfig2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfig2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostGroupAutoUpdateConfig**](HostGroupAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostGroupConfig

> OneAgentHostGroupConfig GetHostGroupConfig(ctx, id).Execute()

Gets the OneAgent configuration in the specified host group

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
    id := "id_example" // string | The Dynatrace entity ID of the required host group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentInAHostGroupApi.GetHostGroupConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentInAHostGroupApi.GetHostGroupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostGroupConfig`: OneAgentHostGroupConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentInAHostGroupApi.GetHostGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OneAgentHostGroupConfig**](OneAgentHostGroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoUpdateConfig1

> UpdateAutoUpdateConfig1(ctx, id).HostGroupAutoUpdateConfig(hostGroupAutoUpdateConfig).Execute()

Updates the configuration of OneAgent auto-update in the specified host group



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
    id := "id_example" // string | The Dynatrace entity ID of the required host group.
    hostGroupAutoUpdateConfig := *openapiclient.NewHostGroupAutoUpdateConfig("DISABLED") // HostGroupAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentInAHostGroupApi.UpdateAutoUpdateConfig1(context.Background(), id).HostGroupAutoUpdateConfig(hostGroupAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentInAHostGroupApi.UpdateAutoUpdateConfig1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoUpdateConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostGroupAutoUpdateConfig** | [**HostGroupAutoUpdateConfig**](HostGroupAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters. | 

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


## ValidateAutoUpdateConfig2

> ValidateAutoUpdateConfig2(ctx, id).HostGroupAutoUpdateConfig(hostGroupAutoUpdateConfig).Execute()

Validates the payload for the `PUT /hostgroups/{id}/autoupdate` request

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
    id := "id_example" // string | The Dynatrace entity ID of the required host group.
    hostGroupAutoUpdateConfig := *openapiclient.NewHostGroupAutoUpdateConfig("DISABLED") // HostGroupAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OneAgentInAHostGroupApi.ValidateAutoUpdateConfig2(context.Background(), id).HostGroupAutoUpdateConfig(hostGroupAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentInAHostGroupApi.ValidateAutoUpdateConfig2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfig2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostGroupAutoUpdateConfig** | [**HostGroupAutoUpdateConfig**](HostGroupAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. | 

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

