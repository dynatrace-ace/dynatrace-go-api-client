# \PublicEndpointsApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBeaconForwarderAddress**](PublicEndpointsApi.md#GetBeaconForwarderAddress) | **Get** /endpoint/beaconForwarderAddress | Get beacon forwarder address
[**GetNodeIpForAgents**](PublicEndpointsApi.md#GetNodeIpForAgents) | **Get** /endpoint/publicIp/agents/{nodeId} | Get node IP for OneAgents
[**GetNodeIpForDomain**](PublicEndpointsApi.md#GetNodeIpForDomain) | **Get** /endpoint/publicIp/domain/{nodeId} | Get node IP for domain
[**GetRootCDNAddress**](PublicEndpointsApi.md#GetRootCDNAddress) | **Get** /endpoint/cdnAddress | Get root CDN address
[**GetWebUiAddress**](PublicEndpointsApi.md#GetWebUiAddress) | **Get** /endpoint/webUiAddress | Get WebUi address
[**StoreNodeIpForAgents**](PublicEndpointsApi.md#StoreNodeIpForAgents) | **Put** /endpoint/publicIp/agents/{nodeId} | Store node IP for OneAgents
[**StoreNodeIpForDomain**](PublicEndpointsApi.md#StoreNodeIpForDomain) | **Put** /endpoint/publicIp/domain/{nodeId} | Store node IP for domain
[**UpdateBeaconForwarderAddress**](PublicEndpointsApi.md#UpdateBeaconForwarderAddress) | **Post** /endpoint/beaconForwarderAddress | Update beacon forwarder address
[**UpdateRootCDNAddress**](PublicEndpointsApi.md#UpdateRootCDNAddress) | **Post** /endpoint/cdnAddress | Update root CDN address
[**UpdateWebUiAddress**](PublicEndpointsApi.md#UpdateWebUiAddress) | **Post** /endpoint/webUiAddress | Update WebUi address



## GetBeaconForwarderAddress

> AddressWrapper GetBeaconForwarderAddress(ctx).Execute()

Get beacon forwarder address

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
    resp, r, err := api_client.PublicEndpointsApi.GetBeaconForwarderAddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.GetBeaconForwarderAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeaconForwarderAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.GetBeaconForwarderAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconForwarderAddressRequest struct via the builder pattern


### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeIpForAgents

> AddressWrapper GetNodeIpForAgents(ctx, nodeId).Execute()

Get node IP for OneAgents

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
    nodeId := int32(56) // int32 | Node ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.GetNodeIpForAgents(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.GetNodeIpForAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeIpForAgents`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.GetNodeIpForAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Node ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeIpForAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeIpForDomain

> AddressWrapper GetNodeIpForDomain(ctx, nodeId).Execute()

Get node IP for domain

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
    nodeId := int32(56) // int32 | Node ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.GetNodeIpForDomain(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.GetNodeIpForDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeIpForDomain`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.GetNodeIpForDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Node ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeIpForDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootCDNAddress

> AddressWrapper GetRootCDNAddress(ctx).Execute()

Get root CDN address

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
    resp, r, err := api_client.PublicEndpointsApi.GetRootCDNAddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.GetRootCDNAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootCDNAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.GetRootCDNAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootCDNAddressRequest struct via the builder pattern


### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebUiAddress

> AddressWrapper GetWebUiAddress(ctx).Execute()

Get WebUi address

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
    resp, r, err := api_client.PublicEndpointsApi.GetWebUiAddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.GetWebUiAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebUiAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.GetWebUiAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebUiAddressRequest struct via the builder pattern


### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreNodeIpForAgents

> StoreNodeIpForAgents(ctx, nodeId).Body(body).Execute()

Store node IP for OneAgents

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
    nodeId := int32(56) // int32 | Node ID
    body := "body_example" // string | String (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.StoreNodeIpForAgents(context.Background(), nodeId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.StoreNodeIpForAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Node ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreNodeIpForAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | String | 

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


## StoreNodeIpForDomain

> AddressWrapper StoreNodeIpForDomain(ctx, nodeId).Body(body).Execute()

Store node IP for domain

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
    nodeId := int32(56) // int32 | Node ID
    body := "body_example" // string | String (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.StoreNodeIpForDomain(context.Background(), nodeId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.StoreNodeIpForDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreNodeIpForDomain`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.StoreNodeIpForDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **int32** | Node ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreNodeIpForDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | String | 

### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBeaconForwarderAddress

> AddressWrapper UpdateBeaconForwarderAddress(ctx).AddressWrapper(addressWrapper).Execute()

Update beacon forwarder address

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
    addressWrapper := *openapiclient.NewAddressWrapper("Address_example") // AddressWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.UpdateBeaconForwarderAddress(context.Background()).AddressWrapper(addressWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.UpdateBeaconForwarderAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBeaconForwarderAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.UpdateBeaconForwarderAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBeaconForwarderAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressWrapper** | [**AddressWrapper**](AddressWrapper.md) |  | 

### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRootCDNAddress

> AddressWrapper UpdateRootCDNAddress(ctx).AddressWrapper(addressWrapper).Execute()

Update root CDN address

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
    addressWrapper := *openapiclient.NewAddressWrapper("Address_example") // AddressWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.UpdateRootCDNAddress(context.Background()).AddressWrapper(addressWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.UpdateRootCDNAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRootCDNAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.UpdateRootCDNAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRootCDNAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressWrapper** | [**AddressWrapper**](AddressWrapper.md) |  | 

### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebUiAddress

> AddressWrapper UpdateWebUiAddress(ctx).AddressWrapper(addressWrapper).Execute()

Update WebUi address

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
    addressWrapper := *openapiclient.NewAddressWrapper("Address_example") // AddressWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicEndpointsApi.UpdateWebUiAddress(context.Background()).AddressWrapper(addressWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicEndpointsApi.UpdateWebUiAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebUiAddress`: AddressWrapper
    fmt.Fprintf(os.Stdout, "Response from `PublicEndpointsApi.UpdateWebUiAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebUiAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressWrapper** | [**AddressWrapper**](AddressWrapper.md) |  | 

### Return type

[**AddressWrapper**](AddressWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

