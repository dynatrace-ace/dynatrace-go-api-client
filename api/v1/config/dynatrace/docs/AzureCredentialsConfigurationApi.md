# \AzureCredentialsConfigurationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureCredentialsConfig**](AzureCredentialsConfigurationApi.md#CreateAzureCredentialsConfig) | **Post** /azure/credentials | Creates a new Azure credentials configuration
[**DeleteAzureCredentialsConfig**](AzureCredentialsConfigurationApi.md#DeleteAzureCredentialsConfig) | **Delete** /azure/credentials/{id} | Deletes the specified Azure credentials configuration
[**GetAzureCredentialsConfig**](AzureCredentialsConfigurationApi.md#GetAzureCredentialsConfig) | **Get** /azure/credentials/{id} | Gets the configuration of the specified Azure credentials
[**ListAzureCredentialsConfigs**](AzureCredentialsConfigurationApi.md#ListAzureCredentialsConfigs) | **Get** /azure/credentials | Lists all available Azure credentials configurations
[**UpdateAzureCredentialsConfig**](AzureCredentialsConfigurationApi.md#UpdateAzureCredentialsConfig) | **Put** /azure/credentials/{id} | Updates an existing Azure credentials configuration
[**ValidateAzureCredentialsConfig**](AzureCredentialsConfigurationApi.md#ValidateAzureCredentialsConfig) | **Post** /azure/credentials/validator | Validates the payload for the &#x60;POST /azure/credentials&#x60; request
[**ValidateConfigurationUpdate**](AzureCredentialsConfigurationApi.md#ValidateConfigurationUpdate) | **Post** /azure/credentials/{id}/validator | Validates the payload for the &#x60;PUT /azure/credentials/{id}&#x60; request



## CreateAzureCredentialsConfig

> EntityShortRepresentation CreateAzureCredentialsConfig(ctx).AzureCredentials(azureCredentials).Execute()

Creates a new Azure credentials configuration



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
    azureCredentials := *openapiclient.NewAzureCredentials("Label_example", "AppId_example", "DirectoryId_example", false, false, []openapiclient.CloudTag{*openapiclient.NewCloudTag()}) // AzureCredentials | The JSON body of the request. Contains parameters of the new Azure credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.CreateAzureCredentialsConfig(context.Background()).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.CreateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAzureCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationApi.CreateAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains parameters of the new Azure credentials configuration. | 

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


## DeleteAzureCredentialsConfig

> DeleteAzureCredentialsConfig(ctx, id).Execute()

Deletes the specified Azure credentials configuration

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
    id := "id_example" // string | The ID of the Azure credentials configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.DeleteAzureCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.DeleteAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureCredentialsConfig

> AzureCredentials GetAzureCredentialsConfig(ctx, id).Execute()

Gets the configuration of the specified Azure credentials

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
    id := "id_example" // string | The ID of the required Azure credentials configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.GetAzureCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.GetAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureCredentialsConfig`: AzureCredentials
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationApi.GetAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required Azure credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureCredentials**](AzureCredentials.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureCredentialsConfigs

> StubList ListAzureCredentialsConfigs(ctx).Execute()

Lists all available Azure credentials configurations

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
    resp, r, err := api_client.AzureCredentialsConfigurationApi.ListAzureCredentialsConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.ListAzureCredentialsConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAzureCredentialsConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationApi.ListAzureCredentialsConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureCredentialsConfigsRequest struct via the builder pattern


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


## UpdateAzureCredentialsConfig

> EntityShortRepresentation UpdateAzureCredentialsConfig(ctx, id).AzureCredentials(azureCredentials).Execute()

Updates an existing Azure credentials configuration



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
    id := "id_example" // string | The ID of the Azure credentials configuration to be updated.
    azureCredentials := *openapiclient.NewAzureCredentials("Label_example", "AppId_example", "DirectoryId_example", false, false, []openapiclient.CloudTag{*openapiclient.NewCloudTag()}) // AzureCredentials | The JSON body of the request. Contains updated parameters of the Azure credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.UpdateAzureCredentialsConfig(context.Background(), id).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.UpdateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationApi.UpdateAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains updated parameters of the Azure credentials configuration. | 

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


## ValidateAzureCredentialsConfig

> ValidateAzureCredentialsConfig(ctx).AzureCredentials(azureCredentials).Execute()

Validates the payload for the `POST /azure/credentials` request

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
    azureCredentials := *openapiclient.NewAzureCredentials("Label_example", "AppId_example", "DirectoryId_example", false, false, []openapiclient.CloudTag{*openapiclient.NewCloudTag()}) // AzureCredentials | The JSON body of the request. Contains the Azure credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.ValidateAzureCredentialsConfig(context.Background()).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.ValidateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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


## ValidateConfigurationUpdate

> ValidateConfigurationUpdate(ctx, id).AzureCredentials(azureCredentials).Execute()

Validates the payload for the `PUT /azure/credentials/{id}` request

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
    id := "id_example" // string | The ID of the Azure credentials configuration to be validated.
    azureCredentials := *openapiclient.NewAzureCredentials("Label_example", "AppId_example", "DirectoryId_example", false, false, []openapiclient.CloudTag{*openapiclient.NewCloudTag()}) // AzureCredentials | The JSON body of the request. Contains the Azure credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureCredentialsConfigurationApi.ValidateConfigurationUpdate(context.Background(), id).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationApi.ValidateConfigurationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConfigurationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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

