# \DashboardsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /dashboards | Creates a dashboard | maturity&#x3D;EARLY_ADOPTER
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /dashboards/{id} | Deletes the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /dashboards/{id} | Gets the properties of the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**GetDashboardStubsList**](DashboardsApi.md#GetDashboardStubsList) | **Get** /dashboards | Lists all dashboards of the environment | maturity&#x3D;EARLY_ADOPTER
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Put** /dashboards/{id} | Updates the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**ValidateDashboardCreation**](DashboardsApi.md#ValidateDashboardCreation) | **Post** /dashboards/validator | Validates the payload for the &#39;POST /dashboards&#39; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateDashboardUpdate**](DashboardsApi.md#ValidateDashboardUpdate) | **Post** /dashboards/{id}/validator | Validates the payload for the &#39;PUT /dashboards/{id}&#39; request | maturity&#x3D;EARLY_ADOPTER



## CreateDashboard

> EntityShortRepresentation CreateDashboard(ctx).Dashboard(dashboard).Execute()

Creates a dashboard | maturity=EARLY_ADOPTER



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
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example"), []openapiclient.Tile{*openapiclient.NewTile("Name_example", "TileType_example", *openapiclient.NewTileBounds())}) // Dashboard | The JSON body of the request. Contains parameters of the new dashboard. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.CreateDashboard(context.Background()).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains parameters of the new dashboard. | 

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


## DeleteDashboard

> DeleteDashboard(ctx, id).Execute()

Deletes the specified dashboard | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the dashboard to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.DeleteDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GetDashboard

> Dashboard GetDashboard(ctx, id).Execute()

Gets the properties of the specified dashboard | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.GetDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardStubsList

> DashboardList GetDashboardStubsList(ctx).Owner(owner).Tags(tags).Execute()

Lists all dashboards of the environment | maturity=EARLY_ADOPTER

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
    owner := "owner_example" // string | The owner of the dashboard. (optional)
    tags := []string{"Inner_example"} // []string | A list of tags applied to the dashboard.    The dashboard must match **all** the specified tags. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.GetDashboardStubsList(context.Background()).Owner(owner).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboardStubsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardStubsList`: DashboardList
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboardStubsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardStubsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the dashboard. | 
 **tags** | **[]string** | A list of tags applied to the dashboard.    The dashboard must match **all** the specified tags. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> EntityShortRepresentation UpdateDashboard(ctx, id).Dashboard(dashboard).Execute()

Updates the specified dashboard | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the dashboard to be updated.    The ID in the request body must match this ID.
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example"), []openapiclient.Tile{*openapiclient.NewTile("Name_example", "TileType_example", *openapiclient.NewTileBounds())}) // Dashboard | The JSON body of the request. Contains updated parameters of the dashboard. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.UpdateDashboard(context.Background(), id).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be updated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains updated parameters of the dashboard. | 

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


## ValidateDashboardCreation

> ValidateDashboardCreation(ctx).Dashboard(dashboard).Execute()

Validates the payload for the 'POST /dashboards' request | maturity=EARLY_ADOPTER



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
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example"), []openapiclient.Tile{*openapiclient.NewTile("Name_example", "TileType_example", *openapiclient.NewTileBounds())}) // Dashboard | The JSON body of the request. Containing the dashboard to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.ValidateDashboardCreation(context.Background()).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ValidateDashboardCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDashboardCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Containing the dashboard to be validated. | 

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


## ValidateDashboardUpdate

> ValidateDashboardUpdate(ctx, id).Dashboard(dashboard).Execute()

Validates the payload for the 'PUT /dashboards/{id}' request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the dashboard to be validated.    The ID in the request body must match this ID.
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example"), []openapiclient.Tile{*openapiclient.NewTile("Name_example", "TileType_example", *openapiclient.NewTileBounds())}) // Dashboard | The JSON body of the request. Contains the dashboard to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.ValidateDashboardUpdate(context.Background(), id).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ValidateDashboardUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be validated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDashboardUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains the dashboard to be validated. | 

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

