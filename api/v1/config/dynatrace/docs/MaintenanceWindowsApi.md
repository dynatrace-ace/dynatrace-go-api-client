# \MaintenanceWindowsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceWindow**](MaintenanceWindowsApi.md#DeleteMaintenanceWindow) | **Delete** /maintenanceWindows/{id} | Deletes the specified maintenance window
[**GetMaintenanceWindow**](MaintenanceWindowsApi.md#GetMaintenanceWindow) | **Get** /maintenanceWindows/{id} | Gets the properties of the specified maintenance window
[**ReturnAllMaintenanceWindows**](MaintenanceWindowsApi.md#ReturnAllMaintenanceWindows) | **Get** /maintenanceWindows | Lists all available maintenance windows
[**StoreMaintenanceWindow**](MaintenanceWindowsApi.md#StoreMaintenanceWindow) | **Put** /maintenanceWindows/{id} | Updates an existing maintenance window
[**StoreMaintenanceWindow1**](MaintenanceWindowsApi.md#StoreMaintenanceWindow1) | **Post** /maintenanceWindows | Creates a new maintenance window
[**ValidateConfiguration7**](MaintenanceWindowsApi.md#ValidateConfiguration7) | **Post** /maintenanceWindows/{id}/validator | Validates the payload for the &#x60;PUT /maintenancewindow/{id}&#x60; request
[**ValidateConfiguration8**](MaintenanceWindowsApi.md#ValidateConfiguration8) | **Post** /maintenanceWindows/validator | Validates the payload for the &#x60;POST /maintenancewindow&#x60; request



## DeleteMaintenanceWindow

> DeleteMaintenanceWindow(ctx, id)

Deletes the specified maintenance window

### Example

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/antihax/optional"
    dynatrace "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

dynatraceURL := "DT_ENVIRONMENT_URL"
dynatraceAPIToken := "DT_API_TOKEN"

ctx := context.WithValue(
    context.Background(),
    dynatrace.ContextAPIKey,
    dynatrace.APIKey{
        Key: dynatraceAPIToken,
        Prefix: "Api-token",
    },
)

configuration := dynatrace.NewConfiguration()
configuration.BasePath = dynatraceURL + "/api/config/v1"
apiClient := dynatrace.NewAPIClient(configuration)

maintenanceWindowID := "7b61a9cd-fe9a-4679-8408-bad669cc9912"
resp, err := apiClient.MaintenanceWindowsApi.DeleteMaintenanceWindow(ctx, maintenanceWindowID)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.DeleteMaintenanceWindow: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
}
fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.DeleteMaintenanceWindow`: %v\n", resp)

```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceWindow

> MaintenanceWindow GetMaintenanceWindow(ctx, id)

Gets the properties of the specified maintenance window

### Example

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/antihax/optional"
    dynatrace "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

dynatraceURL := "DT_ENVIRONMENT_URL"
dynatraceAPIToken := "DT_API_TOKEN"

ctx := context.WithValue(
    context.Background(),
    dynatrace.ContextAPIKey,
    dynatrace.APIKey{
        Key: dynatraceAPIToken,
        Prefix: "Api-token",
    },
)

configuration := dynatrace.NewConfiguration()
configuration.BasePath = dynatraceURL + "/api/config/v1"
apiClient := dynatrace.NewAPIClient(configuration)

maintenanceWindowID := "7b61a9cd-fe9a-4679-8408-bad669cc9912"
resp, r, err := apiClient.MaintenanceWindowsApi.GetMaintenanceWindow(ctx, maintenanceWindowID)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.GetMaintenanceWindow: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.GetMaintenanceWindow`: %v\n", resp)
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required maintenance window. | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllMaintenanceWindows

> StubList ReturnAllMaintenanceWindows(ctx, )

Lists all available maintenance windows

### Example

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/antihax/optional"
    dynatrace "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

dynatraceURL := "DT_ENVIRONMENT_URL"
dynatraceAPIToken := "DT_API_TOKEN"

ctx := context.WithValue(
    context.Background(),
    dynatrace.ContextAPIKey,
    dynatrace.APIKey{
        Key: dynatraceAPIToken,
        Prefix: "Api-token",
    },
)

configuration := dynatrace.NewConfiguration()
configuration.BasePath = dynatraceURL + "/api/config/v1"
apiClient := dynatrace.NewAPIClient(configuration)

resp, r, err := apiClient.MaintenanceWindowsApi.ReturnAllMaintenanceWindows(ctx)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ReturnAllMaintenanceWindows: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.ReturnAllMaintenanceWindows`: %v\n", resp)
```

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StubList**](StubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreMaintenanceWindow

> EntityShortRepresentation StoreMaintenanceWindow(ctx, id, optional)

Updates an existing maintenance window

If a maintenance window with the specified ID doesn't exist, a new one is created.

### Example

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/antihax/optional"
    dynatrace "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

dynatraceURL := "DT_ENVIRONMENT_URL"
dynatraceAPIToken := "DT_API_TOKEN"

ctx := context.WithValue(
    context.Background(),
    dynatrace.ContextAPIKey,
    dynatrace.APIKey{
        Key: dynatraceAPIToken,
        Prefix: "Api-token",
    },
)

configuration := dynatrace.NewConfiguration()
configuration.BasePath = dynatraceURL + "/api/config/v1"
apiClient := dynatrace.NewAPIClient(configuration)

mw := dynatrace.MaintenanceWindow{
    Name:        "Weekly updates",
    Description: "Weekly update of windows servers",
    Type:        "PLANNED",
    Suppression: "DETECT_PROBLEMS_DONT_ALERT",
    Scope: dynatrace.Scope{
        Entities: []string{},
        Matches: []dynatrace.MonitoredEntityFilter{
            dynatrace.MonitoredEntityFilter{
                Type: "HOST",
                Tags: []dynatrace.TagInfo{
                    dynatrace.TagInfo{
                        Context: "CONTEXTLESS",
                        Key:     "OS",
                        Value:   "Windows",
                    },
                },
            },
        },
    },
    Schedule: dynatrace.Schedule{
        RecurrenceType: "WEEKLY",
        Recurrence: dynatrace.Recurrence{
            DayOfWeek:       "FRIDAY",
            StartTime:       "19:21",
            DurationMinutes: 60,
        },
        Start:  "2020-10-13 15:38",
        End:    "2020-10-14 16:38",
        ZoneId: "America/Chicago",
    },
}

body := dynatrace.StoreMaintenanceWindowOpts{
    MaintenanceWindow: optional.NewInterface(mw),
}

maintenanceWindowID := "0581c1b2-b6a9-4209-b62d-281576a74426"

resp, r, err := apiClient.MaintenanceWindowsApi.StoreMaintenanceWindow(ctx, maintenanceWindowID, &body)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.StoreMaintenanceWindow: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.StoreMaintenanceWindow`: %v\n", resp)
```


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be updated. | 
 **optional** | ***StoreMaintenanceWindowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreMaintenanceWindowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains updated parameters of the maintenance window. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreMaintenanceWindow1

> EntityShortRepresentation StoreMaintenanceWindow1(ctx, optional)

Creates a new maintenance window

### Example

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/antihax/optional"
    dynatrace "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

dynatraceURL := "DT_ENVIRONMENT_URL"
dynatraceAPIToken := "DT_API_TOKEN"

ctx := context.WithValue(
    context.Background(),
    dynatrace.ContextAPIKey,
    dynatrace.APIKey{
        Key: dynatraceAPIToken,
        Prefix: "Api-token",
    },
)

configuration := dynatrace.NewConfiguration()
configuration.BasePath = dynatraceURL + "/api/config/v1"
apiClient := dynatrace.NewAPIClient(configuration)

mw := dynatrace.MaintenanceWindow{
    Name:        "Weekly updates",
    Description: "Weekly update of windows servers",
    Type:        "PLANNED",
    Suppression: "DETECT_PROBLEMS_DONT_ALERT",
    Scope: dynatrace.Scope{
        Entities: []string{},
        Matches: []dynatrace.MonitoredEntityFilter{
            dynatrace.MonitoredEntityFilter{
                Type: "HOST",
                Tags: []dynatrace.TagInfo{
                    dynatrace.TagInfo{
                        Context: "CONTEXTLESS",
                        Key:     "OS",
                        Value:   "Windows",
                    },
                },
            },
        },
    },
    Schedule: dynatrace.Schedule{
        RecurrenceType: "WEEKLY",
        Recurrence: dynatrace.Recurrence{
            DayOfWeek:       "FRIDAY",
            StartTime:       "19:21",
            DurationMinutes: 60,
        },
        Start:  "2020-10-13 15:38",
        End:    "2020-10-14 16:38",
        ZoneId: "America/Chicago",
    },
}

body := dynatrace.StoreMaintenanceWindow1Opts{
    MaintenanceWindow: optional.NewInterface(mw),
}

resp, r, err := apiClient.MaintenanceWindowsApi.StoreMaintenanceWindow1(ctx, &body)
if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.StoreMaintenanceWindow1: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.StoreMaintenanceWindow1`: %v\n", resp)

}

```

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StoreMaintenanceWindow1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreMaintenanceWindow1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the new maintenance window. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConfiguration7

> ValidateConfiguration7(ctx, id, optional)

Validates the payload for the `PUT /maintenancewindow/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the maintenance window to be validated. | 
 **optional** | ***ValidateConfiguration7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the maintenance window to be validated. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConfiguration8

> ValidateConfiguration8(ctx, optional)

Validates the payload for the `POST /maintenancewindow` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration8Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**optional.Interface of MaintenanceWindow**](MaintenanceWindow.md)| The JSON body of the request. Contains parameters of the maintenance window be validated. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

