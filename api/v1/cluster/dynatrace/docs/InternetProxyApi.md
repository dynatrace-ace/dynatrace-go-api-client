# \InternetProxyApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOneOfMultiDCConfiguration**](InternetProxyApi.md#DeleteOneOfMultiDCConfiguration) | **Delete** /proxy/configurations/{dc} | Remove Data Center proxy configuration (Multi Data Center deployment)
[**DeleteSingleDCConfiguration**](InternetProxyApi.md#DeleteSingleDCConfiguration) | **Delete** /proxy/configuration | Remove cluster proxy configuration
[**GetAllMultiDCConfigurations**](InternetProxyApi.md#GetAllMultiDCConfigurations) | **Get** /proxy/configurations | Get proxy configurations for all Data Centers (Multi Data Center deployment)
[**GetOneOfMultiDCConfiguration**](InternetProxyApi.md#GetOneOfMultiDCConfiguration) | **Get** /proxy/configurations/{dc} | Get Data Center proxy configuration (Multi Data Center deployment)
[**GetSingleDCConfiguration**](InternetProxyApi.md#GetSingleDCConfiguration) | **Get** /proxy/configuration | Get cluster proxy configuration
[**ModifyOneOfMultiDCConfiguration**](InternetProxyApi.md#ModifyOneOfMultiDCConfiguration) | **Put** /proxy/configurations/{dc} | Set/update Data Center proxy configuration (Multi Data Center deployment)
[**ModifySingleDCConfiguration**](InternetProxyApi.md#ModifySingleDCConfiguration) | **Put** /proxy/configuration | Set/update cluster proxy configuration
[**TestConnectionInMultiDCMode**](InternetProxyApi.md#TestConnectionInMultiDCMode) | **Put** /proxy/test/{dc} | Test Internet connection from specific Data Center using given proxy configuration (Multi Data Center deployment)
[**TestConnectionInSingleDCMode**](InternetProxyApi.md#TestConnectionInSingleDCMode) | **Put** /proxy/test | Test Internet connection using given proxy configuration



## DeleteOneOfMultiDCConfiguration

> InternetProxy DeleteOneOfMultiDCConfiguration(ctx, dc).Execute()

Remove Data Center proxy configuration (Multi Data Center deployment)

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
    dc := "dc_example" // string | Data Center

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.DeleteOneOfMultiDCConfiguration(context.Background(), dc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.DeleteOneOfMultiDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOneOfMultiDCConfiguration`: InternetProxy
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.DeleteOneOfMultiDCConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dc** | **string** | Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneOfMultiDCConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternetProxy**](InternetProxy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSingleDCConfiguration

> InternetProxy DeleteSingleDCConfiguration(ctx).Execute()

Remove cluster proxy configuration

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
    resp, r, err := api_client.InternetProxyApi.DeleteSingleDCConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.DeleteSingleDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSingleDCConfiguration`: InternetProxy
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.DeleteSingleDCConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSingleDCConfigurationRequest struct via the builder pattern


### Return type

[**InternetProxy**](InternetProxy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMultiDCConfigurations

> ProxyConfigurations GetAllMultiDCConfigurations(ctx).Execute()

Get proxy configurations for all Data Centers (Multi Data Center deployment)

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
    resp, r, err := api_client.InternetProxyApi.GetAllMultiDCConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.GetAllMultiDCConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMultiDCConfigurations`: ProxyConfigurations
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.GetAllMultiDCConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMultiDCConfigurationsRequest struct via the builder pattern


### Return type

[**ProxyConfigurations**](ProxyConfigurations.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneOfMultiDCConfiguration

> InternetProxy GetOneOfMultiDCConfiguration(ctx, dc).Execute()

Get Data Center proxy configuration (Multi Data Center deployment)

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
    dc := "dc_example" // string | Data Center

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.GetOneOfMultiDCConfiguration(context.Background(), dc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.GetOneOfMultiDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOneOfMultiDCConfiguration`: InternetProxy
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.GetOneOfMultiDCConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dc** | **string** | Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneOfMultiDCConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternetProxy**](InternetProxy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleDCConfiguration

> InternetProxy GetSingleDCConfiguration(ctx).Execute()

Get cluster proxy configuration

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
    resp, r, err := api_client.InternetProxyApi.GetSingleDCConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.GetSingleDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleDCConfiguration`: InternetProxy
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.GetSingleDCConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleDCConfigurationRequest struct via the builder pattern


### Return type

[**InternetProxy**](InternetProxy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOneOfMultiDCConfiguration

> ModifyOneOfMultiDCConfiguration(ctx, dc).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()

Set/update Data Center proxy configuration (Multi Data Center deployment)

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
    dc := "dc_example" // string | Data Center
    internetProxyChangeRequest := *openapiclient.NewInternetProxyChangeRequest("Scheme_example", "Server_example", int32(123)) // InternetProxyChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.ModifyOneOfMultiDCConfiguration(context.Background(), dc).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.ModifyOneOfMultiDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dc** | **string** | Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOneOfMultiDCConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **internetProxyChangeRequest** | [**InternetProxyChangeRequest**](InternetProxyChangeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySingleDCConfiguration

> ModifySingleDCConfiguration(ctx).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()

Set/update cluster proxy configuration

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
    internetProxyChangeRequest := *openapiclient.NewInternetProxyChangeRequest("Scheme_example", "Server_example", int32(123)) // InternetProxyChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.ModifySingleDCConfiguration(context.Background()).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.ModifySingleDCConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifySingleDCConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internetProxyChangeRequest** | [**InternetProxyChangeRequest**](InternetProxyChangeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionInMultiDCMode

> ConnectionStatus TestConnectionInMultiDCMode(ctx, dc).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()

Test Internet connection from specific Data Center using given proxy configuration (Multi Data Center deployment)

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
    dc := "dc_example" // string | Data Center
    internetProxyChangeRequest := *openapiclient.NewInternetProxyChangeRequest("Scheme_example", "Server_example", int32(123)) // InternetProxyChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.TestConnectionInMultiDCMode(context.Background(), dc).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.TestConnectionInMultiDCMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnectionInMultiDCMode`: ConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.TestConnectionInMultiDCMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dc** | **string** | Data Center | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionInMultiDCModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **internetProxyChangeRequest** | [**InternetProxyChangeRequest**](InternetProxyChangeRequest.md) |  | 

### Return type

[**ConnectionStatus**](ConnectionStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionInSingleDCMode

> ConnectionStatus TestConnectionInSingleDCMode(ctx).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()

Test Internet connection using given proxy configuration

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
    internetProxyChangeRequest := *openapiclient.NewInternetProxyChangeRequest("Scheme_example", "Server_example", int32(123)) // InternetProxyChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetProxyApi.TestConnectionInSingleDCMode(context.Background()).InternetProxyChangeRequest(internetProxyChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetProxyApi.TestConnectionInSingleDCMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnectionInSingleDCMode`: ConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `InternetProxyApi.TestConnectionInSingleDCMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionInSingleDCModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internetProxyChangeRequest** | [**InternetProxyChangeRequest**](InternetProxyChangeRequest.md) |  | 

### Return type

[**ConnectionStatus**](ConnectionStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

