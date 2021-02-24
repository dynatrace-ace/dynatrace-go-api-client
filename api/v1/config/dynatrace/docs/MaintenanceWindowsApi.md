# \MaintenanceWindowsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMaintenanceWindow**](MaintenanceWindowsApi.md#CreateMaintenanceWindow) | **Post** /maintenanceWindows | Creates a new maintenance window
[**DeleteMaintenanceWindow**](MaintenanceWindowsApi.md#DeleteMaintenanceWindow) | **Delete** /maintenanceWindows/{id} | Deletes the specified maintenance window
[**GetMaintenanceWindow**](MaintenanceWindowsApi.md#GetMaintenanceWindow) | **Get** /maintenanceWindows/{id} | Gets the properties of the specified maintenance window
[**ListMaintenanceWindows**](MaintenanceWindowsApi.md#ListMaintenanceWindows) | **Get** /maintenanceWindows | Lists all available maintenance windows
[**UpdateMaintenanceWindow**](MaintenanceWindowsApi.md#UpdateMaintenanceWindow) | **Put** /maintenanceWindows/{id} | Updates an existing maintenance window
[**ValidateCreateMaintenanceWindow**](MaintenanceWindowsApi.md#ValidateCreateMaintenanceWindow) | **Post** /maintenanceWindows/validator | Validates the payload for the &#x60;POST /maintenancewindow&#x60; request
[**ValidateUpdateMaintenanceWindow**](MaintenanceWindowsApi.md#ValidateUpdateMaintenanceWindow) | **Post** /maintenanceWindows/{id}/validator | Validates the payload for the &#x60;PUT /maintenancewindow/{id}&#x60; request



## CreateMaintenanceWindow

> EntityShortRepresentation CreateMaintenanceWindow(ctx).MaintenanceWindow(maintenanceWindow).Execute()

Creates a new maintenance window



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
    maintenanceWindow := *openapiclient.NewMaintenanceWindow("Name_example", "Description_example", "Type_example", "Suppression_example", *openapiclient.NewSchedule("RecurrenceType_example", "Start_example", "End_example", "ZoneId_example")) // MaintenanceWindow | The JSON body of the request. Contains parameters of the new maintenance window. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.CreateMaintenanceWindow(context.Background()).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.CreateMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMaintenanceWindow`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.CreateMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) | The JSON body of the request. Contains parameters of the new maintenance window. | 

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


## DeleteMaintenanceWindow

> DeleteMaintenanceWindow(ctx, id).Execute()

Deletes the specified maintenance window

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
    id := TODO // string | The ID of the maintenance window to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.DeleteMaintenanceWindow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.DeleteMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the maintenance window to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaintenanceWindowRequest struct via the builder pattern


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


## GetMaintenanceWindow

> MaintenanceWindow GetMaintenanceWindow(ctx, id).Execute()

Gets the properties of the specified maintenance window

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
    id := TODO // string | The ID of the required maintenance window.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.GetMaintenanceWindow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.GetMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceWindow`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.GetMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required maintenance window. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMaintenanceWindows

> StubList ListMaintenanceWindows(ctx).Execute()

Lists all available maintenance windows

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
    resp, r, err := api_client.MaintenanceWindowsApi.ListMaintenanceWindows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ListMaintenanceWindows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMaintenanceWindows`: StubList
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.ListMaintenanceWindows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMaintenanceWindowsRequest struct via the builder pattern


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


## UpdateMaintenanceWindow

> EntityShortRepresentation UpdateMaintenanceWindow(ctx, id).MaintenanceWindow(maintenanceWindow).Execute()

Updates an existing maintenance window



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
    id := TODO // string | The ID of the maintenance window to be updated.
    maintenanceWindow := *openapiclient.NewMaintenanceWindow("Name_example", "Description_example", "Type_example", "Suppression_example", *openapiclient.NewSchedule("RecurrenceType_example", "Start_example", "End_example", "ZoneId_example")) // MaintenanceWindow | The JSON body of the request. Contains updated parameters of the maintenance window. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.UpdateMaintenanceWindow(context.Background(), id).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.UpdateMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMaintenanceWindow`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.UpdateMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the maintenance window to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) | The JSON body of the request. Contains updated parameters of the maintenance window. | 

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


## ValidateCreateMaintenanceWindow

> ValidateCreateMaintenanceWindow(ctx).MaintenanceWindow(maintenanceWindow).Execute()

Validates the payload for the `POST /maintenancewindow` request

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
    maintenanceWindow := *openapiclient.NewMaintenanceWindow("Name_example", "Description_example", "Type_example", "Suppression_example", *openapiclient.NewSchedule("RecurrenceType_example", "Start_example", "End_example", "ZoneId_example")) // MaintenanceWindow | The JSON body of the request. Contains parameters of the maintenance window be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.ValidateCreateMaintenanceWindow(context.Background()).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ValidateCreateMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) | The JSON body of the request. Contains parameters of the maintenance window be validated. | 

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


## ValidateUpdateMaintenanceWindow

> ValidateUpdateMaintenanceWindow(ctx, id).MaintenanceWindow(maintenanceWindow).Execute()

Validates the payload for the `PUT /maintenancewindow/{id}` request

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
    id := TODO // string | The ID of the maintenance window to be validated.
    maintenanceWindow := *openapiclient.NewMaintenanceWindow("Name_example", "Description_example", "Type_example", "Suppression_example", *openapiclient.NewSchedule("RecurrenceType_example", "Start_example", "End_example", "ZoneId_example")) // MaintenanceWindow | The JSON body of the request. Contains parameters of the maintenance window to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.ValidateUpdateMaintenanceWindow(context.Background(), id).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ValidateUpdateMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the maintenance window to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateMaintenanceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) | The JSON body of the request. Contains parameters of the maintenance window to be validated. | 

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

