# \SyntheticLocationsAndNodesApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocation**](SyntheticLocationsAndNodesApi.md#AddLocation) | **Post** /synthetic/locations | Creates a new private synthetic cluster location
[**GetLocation**](SyntheticLocationsAndNodesApi.md#GetLocation) | **Get** /synthetic/locations/{locationId} | Gets properties of the specified cluster location
[**GetLocations**](SyntheticLocationsAndNodesApi.md#GetLocations) | **Get** /synthetic/locations | Lists all cluster private synthetic locations
[**GetNode**](SyntheticLocationsAndNodesApi.md#GetNode) | **Get** /synthetic/nodes/{nodeId} | Lists properties of the specified synthetic cluster node | maturity&#x3D;EARLY_ADOPTER
[**GetNodes**](SyntheticLocationsAndNodesApi.md#GetNodes) | **Get** /synthetic/nodes | Lists all synthetic cluster nodes | maturity&#x3D;EARLY_ADOPTER
[**RemoveLocation**](SyntheticLocationsAndNodesApi.md#RemoveLocation) | **Delete** /synthetic/locations/{locationId} | Deletes the specified private synthetic cluster location
[**UpdateLocation**](SyntheticLocationsAndNodesApi.md#UpdateLocation) | **Put** /synthetic/locations/{locationId} | Updates the specified private synthetic cluster location



## AddLocation

> SyntheticLocationIdsDto AddLocation(ctx).PrivateSyntheticLocation(privateSyntheticLocation).Execute()

Creates a new private synthetic cluster location

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
    privateSyntheticLocation := *openapiclient.NewPrivateSyntheticLocation([]string{"Nodes_example"}, "EntityId_example", "Type_example", "Name_example", float64(123), float64(123)) // PrivateSyntheticLocation | The JSON body of the request. Contains parameters of the new private synthetic cluster location (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.AddLocation(context.Background()).PrivateSyntheticLocation(privateSyntheticLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.AddLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocation`: SyntheticLocationIdsDto
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.AddLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateSyntheticLocation** | [**PrivateSyntheticLocation**](PrivateSyntheticLocation.md) | The JSON body of the request. Contains parameters of the new private synthetic cluster location | 

### Return type

[**SyntheticLocationIdsDto**](SyntheticLocationIdsDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocation

> SyntheticLocation GetLocation(ctx, locationId).Execute()

Gets properties of the specified cluster location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the required cluster location.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetLocation(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocation`: SyntheticLocation
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the required cluster location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticLocation**](SyntheticLocation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocations

> SyntheticLocations GetLocations(ctx).Execute()

Lists all cluster private synthetic locations

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
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocations`: SyntheticLocations
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


### Return type

[**SyntheticLocations**](SyntheticLocations.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> Node GetNode(ctx, nodeId).Execute()

Lists properties of the specified synthetic cluster node | maturity=EARLY_ADOPTER

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
    nodeId := "nodeId_example" // string | The ID of the required synthetic cluster node.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: Node
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The ID of the required synthetic cluster node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Node**](Node.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> Nodes GetNodes(ctx).Execute()

Lists all synthetic cluster nodes | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodes`: Nodes
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


### Return type

[**Nodes**](Nodes.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLocation

> RemoveLocation(ctx, locationId).Execute()

Deletes the specified private synthetic cluster location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the private synthetic cluster location to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.RemoveLocation(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.RemoveLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the private synthetic cluster location to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> UpdateLocation(ctx, locationId).PrivateSyntheticLocation(privateSyntheticLocation).Execute()

Updates the specified private synthetic cluster location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the private synthetic cluster location to be updated.
    privateSyntheticLocation := *openapiclient.NewPrivateSyntheticLocation([]string{"Nodes_example"}, "EntityId_example", "Type_example", "Name_example", float64(123), float64(123)) // PrivateSyntheticLocation | The JSON body of the request. Contains updated parameters of the private synthetic cluster location. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.UpdateLocation(context.Background(), locationId).PrivateSyntheticLocation(privateSyntheticLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the private synthetic cluster location to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateSyntheticLocation** | [**PrivateSyntheticLocation**](PrivateSyntheticLocation.md) | The JSON body of the request. Contains updated parameters of the private synthetic cluster location. | 

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

