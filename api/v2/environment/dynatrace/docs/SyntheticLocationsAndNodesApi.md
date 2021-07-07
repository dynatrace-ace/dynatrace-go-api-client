# \SyntheticLocationsAndNodesApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocation**](SyntheticLocationsAndNodesApi.md#AddLocation) | **Post** /synthetic/locations | Creates a new private synthetic location
[**GetLocation**](SyntheticLocationsAndNodesApi.md#GetLocation) | **Get** /synthetic/locations/{locationId} | Gets properties of the specified location
[**GetLocations**](SyntheticLocationsAndNodesApi.md#GetLocations) | **Get** /synthetic/locations | Lists all synthetic locations (both public and private) available for your environment
[**GetLocationsStatus**](SyntheticLocationsAndNodesApi.md#GetLocationsStatus) | **Get** /synthetic/locations/status | Checks the status of public synthetic locations
[**GetNode**](SyntheticLocationsAndNodesApi.md#GetNode) | **Get** /synthetic/nodes/{nodeId} | Lists properties of the specified synthetic node
[**GetNodes**](SyntheticLocationsAndNodesApi.md#GetNodes) | **Get** /synthetic/nodes | Lists all synthetic nodes available in your environment
[**RemoveLocation**](SyntheticLocationsAndNodesApi.md#RemoveLocation) | **Delete** /synthetic/locations/{locationId} | Deletes the specified private synthetic location
[**UpdateLocation**](SyntheticLocationsAndNodesApi.md#UpdateLocation) | **Put** /synthetic/locations/{locationId} | Updates the specified synthetic location
[**UpdateLocationsStatus**](SyntheticLocationsAndNodesApi.md#UpdateLocationsStatus) | **Put** /synthetic/locations/status | Changes the status of public synthetic locations



## AddLocation

> SyntheticLocationIdsDto AddLocation(ctx).PrivateSyntheticLocation(privateSyntheticLocation).Execute()

Creates a new private synthetic location

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
    privateSyntheticLocation := *openapiclient.NewPrivateSyntheticLocation([]string{"Nodes_example"}, "EntityId_example", "Type_example", "Name_example", float64(123), float64(123)) // PrivateSyntheticLocation | The JSON body of the request. Contains parameters of the new private synthetic location. (optional)

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
 **privateSyntheticLocation** | [**PrivateSyntheticLocation**](PrivateSyntheticLocation.md) | The JSON body of the request. Contains parameters of the new private synthetic location. | 

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

Gets properties of the specified location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the required location.

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
**locationId** | **string** | The Dynatrace entity ID of the required location. | 

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

> SyntheticLocations GetLocations(ctx).CloudPlatform(cloudPlatform).Type_(type_).Execute()

Lists all synthetic locations (both public and private) available for your environment

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
    cloudPlatform := "cloudPlatform_example" // string | Filters the resulting set of locations to those which are hosted on a specific cloud platform. (optional)
    type_ := "type__example" // string | Filters the resulting set of locations to those of a specific type. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetLocations(context.Background()).CloudPlatform(cloudPlatform).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocations`: SyntheticLocations
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPlatform** | **string** | Filters the resulting set of locations to those which are hosted on a specific cloud platform. | 
 **type_** | **string** | Filters the resulting set of locations to those of a specific type. | 

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


## GetLocationsStatus

> SyntheticPublicLocationsStatus GetLocationsStatus(ctx).Execute()

Checks the status of public synthetic locations

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
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.GetLocationsStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.GetLocationsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationsStatus`: SyntheticPublicLocationsStatus
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsAndNodesApi.GetLocationsStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsStatusRequest struct via the builder pattern


### Return type

[**SyntheticPublicLocationsStatus**](SyntheticPublicLocationsStatus.md)

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

Lists properties of the specified synthetic node

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
    nodeId := "nodeId_example" // string | The ID of the required synthetic node.

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
**nodeId** | **string** | The ID of the required synthetic node. | 

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

Lists all synthetic nodes available in your environment

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

Deletes the specified private synthetic location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the private synthetic location to be deleted.

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
**locationId** | **string** | The Dynatrace entity ID of the private synthetic location to be deleted. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> UpdateLocation(ctx, locationId).SyntheticLocationUpdate(syntheticLocationUpdate).Execute()

Updates the specified synthetic location



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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the synthetic location to be updated.
    syntheticLocationUpdate := *openapiclient.NewSyntheticLocationUpdate() // SyntheticLocationUpdate | The JSON body of the request. Contains updated parameters of the location. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.UpdateLocation(context.Background(), locationId).SyntheticLocationUpdate(syntheticLocationUpdate).Execute()
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
**locationId** | **string** | The Dynatrace entity ID of the synthetic location to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticLocationUpdate** | [**SyntheticLocationUpdate**](SyntheticLocationUpdate.md) | The JSON body of the request. Contains updated parameters of the location. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocationsStatus

> UpdateLocationsStatus(ctx).SyntheticPublicLocationsStatus(syntheticPublicLocationsStatus).Execute()

Changes the status of public synthetic locations

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
    syntheticPublicLocationsStatus := *openapiclient.NewSyntheticPublicLocationsStatus(false) // SyntheticPublicLocationsStatus | The new status of public synthetic locations. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticLocationsAndNodesApi.UpdateLocationsStatus(context.Background()).SyntheticPublicLocationsStatus(syntheticPublicLocationsStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsAndNodesApi.UpdateLocationsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticPublicLocationsStatus** | [**SyntheticPublicLocationsStatus**](SyntheticPublicLocationsStatus.md) | The new status of public synthetic locations. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

