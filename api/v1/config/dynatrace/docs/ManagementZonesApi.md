# \ManagementZonesApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagementZone**](ManagementZonesApi.md#CreateManagementZone) | **Post** /managementZones | Creates a new management zone
[**DeleteManagementZone**](ManagementZonesApi.md#DeleteManagementZone) | **Delete** /managementZones/{id} | Deletes the specified management zone
[**GetManagementZone**](ManagementZonesApi.md#GetManagementZone) | **Get** /managementZones/{id} | Lists the parameters of the specified management zone
[**ListManagementZones**](ManagementZonesApi.md#ListManagementZones) | **Get** /managementZones | Lists all configured management zones
[**UpdateManagementZone**](ManagementZonesApi.md#UpdateManagementZone) | **Put** /managementZones/{id} | Updates the specified management zone
[**ValidateCreateManagementZone**](ManagementZonesApi.md#ValidateCreateManagementZone) | **Post** /managementZones/validator | Validates a new management zone for the &#x60;POST /managementZones&#x60; request
[**ValidateUpdateManagementZone**](ManagementZonesApi.md#ValidateUpdateManagementZone) | **Post** /managementZones/{id}/validator | Validate updates of existing management zone for the &#x60;PUT /managementZones/{id}&#x60; request



## CreateManagementZone

> EntityShortRepresentation CreateManagementZone(ctx).ManagementZone(managementZone).Execute()

Creates a new management zone



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
    managementZone := *openapiclient.NewManagementZone("Name_example") // ManagementZone | The JSON body of the request. Contains parameters of the new management zone. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.CreateManagementZone(context.Background()).ManagementZone(managementZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.CreateManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateManagementZone`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ManagementZonesApi.CreateManagementZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagementZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**ManagementZone**](ManagementZone.md) | The JSON body of the request. Contains parameters of the new management zone. | 

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


## DeleteManagementZone

> DeleteManagementZone(ctx, id).Execute()

Deletes the specified management zone

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
    id := "id_example" // string | The ID of the management zone to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.DeleteManagementZone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.DeleteManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the management zone to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagementZoneRequest struct via the builder pattern


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


## GetManagementZone

> ManagementZone GetManagementZone(ctx, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()

Lists the parameters of the specified management zone

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
    id := "id_example" // string | The ID of the required management zone.
    includeProcessGroupReferences := true // bool | Flag to include process group references to the response.    Process group references aren't compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.GetManagementZone(context.Background(), id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.GetManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementZone`: ManagementZone
    fmt.Fprintf(os.Stdout, "Response from `ManagementZonesApi.GetManagementZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required management zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **bool** | Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed. | [default to false]

### Return type

[**ManagementZone**](ManagementZone.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagementZones

> StubList ListManagementZones(ctx).Execute()

Lists all configured management zones

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
    resp, r, err := api_client.ManagementZonesApi.ListManagementZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.ListManagementZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagementZones`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ManagementZonesApi.ListManagementZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListManagementZonesRequest struct via the builder pattern


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


## UpdateManagementZone

> EntityShortRepresentation UpdateManagementZone(ctx, id).ManagementZone(managementZone).Execute()

Updates the specified management zone



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
    id := "id_example" // string | The ID of the management zone to be updated.   If you set the ID in the body as well, it must match this ID.
    managementZone := *openapiclient.NewManagementZone("Name_example") // ManagementZone | The JSON body of the request. Contains updated parameters of the management zone. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.UpdateManagementZone(context.Background(), id).ManagementZone(managementZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.UpdateManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManagementZone`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ManagementZonesApi.UpdateManagementZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the management zone to be updated.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagementZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**ManagementZone**](ManagementZone.md) | The JSON body of the request. Contains updated parameters of the management zone. | 

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


## ValidateCreateManagementZone

> ValidateCreateManagementZone(ctx).ManagementZone(managementZone).Execute()

Validates a new management zone for the `POST /managementZones` request

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
    managementZone := *openapiclient.NewManagementZone("Name_example") // ManagementZone | The JSON body of the request, containing the management zone parameters to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.ValidateCreateManagementZone(context.Background()).ManagementZone(managementZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.ValidateCreateManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateManagementZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**ManagementZone**](ManagementZone.md) | The JSON body of the request, containing the management zone parameters to validate. | 

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


## ValidateUpdateManagementZone

> ValidateUpdateManagementZone(ctx, id).ManagementZone(managementZone).Execute()

Validate updates of existing management zone for the `PUT /managementZones/{id}` request

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
    id := "id_example" // string | The ID of the management zone to validate.
    managementZone := *openapiclient.NewManagementZone("Name_example") // ManagementZone | The JSON body of the request, containing the management zone parameters to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementZonesApi.ValidateUpdateManagementZone(context.Background(), id).ManagementZone(managementZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementZonesApi.ValidateUpdateManagementZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the management zone to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateManagementZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**ManagementZone**](ManagementZone.md) | The JSON body of the request, containing the management zone parameters to validate. | 

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

