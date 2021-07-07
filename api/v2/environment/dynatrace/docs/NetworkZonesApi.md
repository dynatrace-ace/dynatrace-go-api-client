# \NetworkZonesApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateNetworkZone**](NetworkZonesApi.md#CreateOrUpdateNetworkZone) | **Put** /networkZones/{id} | Updates an existing network zone or creates a new one | maturity&#x3D;EARLY_ADOPTER
[**DeleteNetworkZone**](NetworkZonesApi.md#DeleteNetworkZone) | **Delete** /networkZones/{id} | Deletes the specified network zone | maturity&#x3D;EARLY_ADOPTER
[**GetAllNetworkZones**](NetworkZonesApi.md#GetAllNetworkZones) | **Get** /networkZones | Lists all existing network zones | maturity&#x3D;EARLY_ADOPTER
[**GetNetworkZoneSettings**](NetworkZonesApi.md#GetNetworkZoneSettings) | **Get** /networkZoneSettings | Gets the global configuration of network zones | maturity&#x3D;EARLY_ADOPTER
[**GetSingleNetworkZone**](NetworkZonesApi.md#GetSingleNetworkZone) | **Get** /networkZones/{id} | Gets parameters of the specified network zone | maturity&#x3D;EARLY_ADOPTER
[**UpdateNetworkZoneSettings**](NetworkZonesApi.md#UpdateNetworkZoneSettings) | **Put** /networkZoneSettings | Updates the global configuration of network zones | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateNetworkZone

> EntityShortRepresentation CreateOrUpdateNetworkZone(ctx, id).NetworkZone(networkZone).Execute()

Updates an existing network zone or creates a new one | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the network zone to be updated.    If you set the ID in the body as well, it must match this ID.    The ID is not case sensitive. Dynatrace stores the ID in lowercase.
    networkZone := *openapiclient.NewNetworkZone() // NetworkZone | The JSON body of the request. Contains parameters of the network zone.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkZonesApi.CreateOrUpdateNetworkZone(context.Background(), id).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.CreateOrUpdateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateNetworkZone`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesApi.CreateOrUpdateNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the network zone to be updated.    If you set the ID in the body as well, it must match this ID.    The ID is not case sensitive. Dynatrace stores the ID in lowercase. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkZone** | [**NetworkZone**](NetworkZone.md) | The JSON body of the request. Contains parameters of the network zone. | 

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


## DeleteNetworkZone

> DeleteNetworkZone(ctx, id).Execute()

Deletes the specified network zone | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the network zone to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkZonesApi.DeleteNetworkZone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.DeleteNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the network zone to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkZoneRequest struct via the builder pattern


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


## GetAllNetworkZones

> NetworkZoneList GetAllNetworkZones(ctx).Execute()

Lists all existing network zones | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.NetworkZonesApi.GetAllNetworkZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.GetAllNetworkZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNetworkZones`: NetworkZoneList
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesApi.GetAllNetworkZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworkZonesRequest struct via the builder pattern


### Return type

[**NetworkZoneList**](NetworkZoneList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkZoneSettings

> NetworkZoneSettings GetNetworkZoneSettings(ctx).Execute()

Gets the global configuration of network zones | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.NetworkZonesApi.GetNetworkZoneSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.GetNetworkZoneSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkZoneSettings`: NetworkZoneSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesApi.GetNetworkZoneSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkZoneSettingsRequest struct via the builder pattern


### Return type

[**NetworkZoneSettings**](NetworkZoneSettings.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleNetworkZone

> NetworkZone GetSingleNetworkZone(ctx, id).Execute()

Gets parameters of the specified network zone | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required network zone.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkZonesApi.GetSingleNetworkZone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.GetSingleNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleNetworkZone`: NetworkZone
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesApi.GetSingleNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required network zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkZoneSettings

> UpdateNetworkZoneSettings(ctx).NetworkZoneSettings(networkZoneSettings).Execute()

Updates the global configuration of network zones | maturity=EARLY_ADOPTER

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
    networkZoneSettings := *openapiclient.NewNetworkZoneSettings() // NetworkZoneSettings | The JSON body of the request. Contains global configuration of network zones.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkZonesApi.UpdateNetworkZoneSettings(context.Background()).NetworkZoneSettings(networkZoneSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesApi.UpdateNetworkZoneSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkZoneSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZoneSettings** | [**NetworkZoneSettings**](NetworkZoneSettings.md) | The JSON body of the request. Contains global configuration of network zones. | 

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

