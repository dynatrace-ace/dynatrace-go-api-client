# \FirewallConfigurationApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClusterNodes**](FirewallConfigurationApi.md#AddClusterNodes) | **Post** /firewallManagement/addClusterNode | Add cluster node
[**GetAsyncNodeRemovalStatus**](FirewallConfigurationApi.md#GetAsyncNodeRemovalStatus) | **Get** /firewallManagement/removeClusterNode | 
[**GetClusterNodes**](FirewallConfigurationApi.md#GetClusterNodes) | **Get** /firewallManagement/clusterNodes | Get cluster nodes
[**RemoveClusterNodes**](FirewallConfigurationApi.md#RemoveClusterNodes) | **Post** /firewallManagement/removeClusterNode | Remove cluster node
[**RemoveClusterNodesDeleteMethod**](FirewallConfigurationApi.md#RemoveClusterNodesDeleteMethod) | **Delete** /firewallManagement/removeClusterNode/{ip} | 



## AddClusterNodes

> AddClusterNodes(ctx).NodeConfigDto(nodeConfigDto).Execute()

Add cluster node

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
    nodeConfigDto := *openapiclient.NewNodeConfigDto() // NodeConfigDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallConfigurationApi.AddClusterNodes(context.Background()).NodeConfigDto(nodeConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallConfigurationApi.AddClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeConfigDto** | [**NodeConfigDto**](NodeConfigDto.md) |  | 

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


## GetAsyncNodeRemovalStatus

> GetAsyncNodeRemovalStatus(ctx).Execute()



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
    resp, r, err := api_client.FirewallConfigurationApi.GetAsyncNodeRemovalStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallConfigurationApi.GetAsyncNodeRemovalStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncNodeRemovalStatusRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodes

> ClusterNodesConfigDto GetClusterNodes(ctx).Execute()

Get cluster nodes

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
    resp, r, err := api_client.FirewallConfigurationApi.GetClusterNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallConfigurationApi.GetClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNodes`: ClusterNodesConfigDto
    fmt.Fprintf(os.Stdout, "Response from `FirewallConfigurationApi.GetClusterNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodesRequest struct via the builder pattern


### Return type

[**ClusterNodesConfigDto**](ClusterNodesConfigDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClusterNodes

> RemoveClusterNodes(ctx).FirewallNodeIp(firewallNodeIp).Execute()

Remove cluster node

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
    firewallNodeIp := *openapiclient.NewFirewallNodeIp() // FirewallNodeIp |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallConfigurationApi.RemoveClusterNodes(context.Background()).FirewallNodeIp(firewallNodeIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallConfigurationApi.RemoveClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallNodeIp** | [**FirewallNodeIp**](FirewallNodeIp.md) |  | 

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


## RemoveClusterNodesDeleteMethod

> RemoveClusterNodesDeleteMethod(ctx, ip).Scope(scope).AsyncCrossNodeCommunication(asyncCrossNodeCommunication).Execute()



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
    ip := "ip_example" // string | 
    scope := "scope_example" // string |  (optional)
    asyncCrossNodeCommunication := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallConfigurationApi.RemoveClusterNodesDeleteMethod(context.Background(), ip).Scope(scope).AsyncCrossNodeCommunication(asyncCrossNodeCommunication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallConfigurationApi.RemoveClusterNodesDeleteMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveClusterNodesDeleteMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** |  | 
 **asyncCrossNodeCommunication** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

