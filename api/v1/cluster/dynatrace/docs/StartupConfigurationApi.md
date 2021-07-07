# \StartupConfigurationApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateNewNodeId**](StartupConfigurationApi.md#GenerateNewNodeId) | **Post** /bootstrapManagement/newNodeId | Get new node id
[**GetActiveGateConfigProperties**](StartupConfigurationApi.md#GetActiveGateConfigProperties) | **Get** /bootstrapManagement/files/ag/configProperties | Get ActiveGate config properties
[**GetActiveGateSecureConfigProperties**](StartupConfigurationApi.md#GetActiveGateSecureConfigProperties) | **Get** /bootstrapManagement/files/ag/secureConfigProperties | Get ActiveGate secure config properties
[**GetConfigProperties**](StartupConfigurationApi.md#GetConfigProperties) | **Get** /bootstrapManagement/files/configProperties | Get config properties
[**GetInstallerMetadata**](StartupConfigurationApi.md#GetInstallerMetadata) | **Get** /bootstrapManagement/files/installerMetadata | Get installer metadata
[**GetKeyStore**](StartupConfigurationApi.md#GetKeyStore) | **Get** /bootstrapManagement/files/keyStore | Get key store
[**GetRuntimeProperties**](StartupConfigurationApi.md#GetRuntimeProperties) | **Get** /bootstrapManagement/files/runtimeProperties | Get runtime properties
[**GetSecureConfigProperties**](StartupConfigurationApi.md#GetSecureConfigProperties) | **Get** /bootstrapManagement/files/secureConfigProperties | Get secure config properties



## GenerateNewNodeId

> FileDto GenerateNewNodeId(ctx).Execute()

Get new node id

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
    resp, r, err := api_client.StartupConfigurationApi.GenerateNewNodeId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GenerateNewNodeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateNewNodeId`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GenerateNewNodeId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateNewNodeIdRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveGateConfigProperties

> FileDto GetActiveGateConfigProperties(ctx).Execute()

Get ActiveGate config properties

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
    resp, r, err := api_client.StartupConfigurationApi.GetActiveGateConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetActiveGateConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGateConfigProperties`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetActiveGateConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGateConfigPropertiesRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveGateSecureConfigProperties

> FileDto GetActiveGateSecureConfigProperties(ctx).Execute()

Get ActiveGate secure config properties

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
    resp, r, err := api_client.StartupConfigurationApi.GetActiveGateSecureConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetActiveGateSecureConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGateSecureConfigProperties`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetActiveGateSecureConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGateSecureConfigPropertiesRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigProperties

> FileDto GetConfigProperties(ctx).Execute()

Get config properties

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
    resp, r, err := api_client.StartupConfigurationApi.GetConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigProperties`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigPropertiesRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallerMetadata

> InstallerMetadata GetInstallerMetadata(ctx).Execute()

Get installer metadata

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
    resp, r, err := api_client.StartupConfigurationApi.GetInstallerMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetInstallerMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstallerMetadata`: InstallerMetadata
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetInstallerMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallerMetadataRequest struct via the builder pattern


### Return type

[**InstallerMetadata**](InstallerMetadata.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyStore

> FileDto GetKeyStore(ctx).Execute()

Get key store

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
    resp, r, err := api_client.StartupConfigurationApi.GetKeyStore(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetKeyStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyStore`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetKeyStore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyStoreRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeProperties

> FileDto GetRuntimeProperties(ctx).Execute()

Get runtime properties

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
    resp, r, err := api_client.StartupConfigurationApi.GetRuntimeProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetRuntimeProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeProperties`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetRuntimeProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimePropertiesRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecureConfigProperties

> FileDto GetSecureConfigProperties(ctx).Execute()

Get secure config properties

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
    resp, r, err := api_client.StartupConfigurationApi.GetSecureConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupConfigurationApi.GetSecureConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecureConfigProperties`: FileDto
    fmt.Fprintf(os.Stdout, "Response from `StartupConfigurationApi.GetSecureConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecureConfigPropertiesRequest struct via the builder pattern


### Return type

[**FileDto**](FileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

