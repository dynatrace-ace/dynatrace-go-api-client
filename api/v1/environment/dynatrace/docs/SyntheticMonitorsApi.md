# \SyntheticMonitorsApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitor**](SyntheticMonitorsApi.md#AddMonitor) | **Post** /synthetic/monitors | Creates a new synthetic monitor
[**DeleteMonitor**](SyntheticMonitorsApi.md#DeleteMonitor) | **Delete** /synthetic/monitors/{monitorId} | Deletes the specified synthetic monitor
[**GetMonitor**](SyntheticMonitorsApi.md#GetMonitor) | **Get** /synthetic/monitors/{monitorId} | Gets parameters of the specified synthetic monitor
[**GetMonitorsCollection**](SyntheticMonitorsApi.md#GetMonitorsCollection) | **Get** /synthetic/monitors | Lists all synthetic monitors in your Dynatrace environment
[**ReplaceMonitor**](SyntheticMonitorsApi.md#ReplaceMonitor) | **Put** /synthetic/monitors/{monitorId} | Updates parameters of the specified synthetic monitor



## AddMonitor

> EntityIdDto AddMonitor(ctx).SyntheticMonitorUpdate(syntheticMonitorUpdate).Execute()

Creates a new synthetic monitor

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
    syntheticMonitorUpdate := *openapiclient.NewSyntheticMonitorUpdate(int32(123), "Type_example", "Name_example", []string{"Locations_example"}, false, map[string]interface{}(123), []openapiclient.TagWithSourceInfo{*openapiclient.NewTagWithSourceInfo("Context_example", "Key_example")}, []string{"ManuallyAssignedApps_example"}) // SyntheticMonitorUpdate | The JSON body of the request, containing parameters of the new synthetic monitor. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticMonitorsApi.AddMonitor(context.Background()).SyntheticMonitorUpdate(syntheticMonitorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticMonitorsApi.AddMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMonitor`: EntityIdDto
    fmt.Fprintf(os.Stdout, "Response from `SyntheticMonitorsApi.AddMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticMonitorUpdate** | [**SyntheticMonitorUpdate**](SyntheticMonitorUpdate.md) | The JSON body of the request, containing parameters of the new synthetic monitor. | 

### Return type

[**EntityIdDto**](EntityIdDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeleteMonitor(ctx, monitorId).Execute()

Deletes the specified synthetic monitor

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
    monitorId := "monitorId_example" // string | The ID of the synthetic monitor to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticMonitorsApi.DeleteMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticMonitorsApi.DeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | The ID of the synthetic monitor to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


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


## GetMonitor

> SyntheticMonitor GetMonitor(ctx, monitorId).Execute()

Gets parameters of the specified synthetic monitor

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
    monitorId := "monitorId_example" // string | The ID of the required synthetic monitor

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticMonitorsApi.GetMonitor(context.Background(), monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticMonitorsApi.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: SyntheticMonitor
    fmt.Fprintf(os.Stdout, "Response from `SyntheticMonitorsApi.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | The ID of the required synthetic monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticMonitor**](SyntheticMonitor.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorsCollection

> Monitors GetMonitorsCollection(ctx).ManagementZone(managementZone).Tag(tag).Location(location).AssignedApps(assignedApps).Type_(type_).Enabled(enabled).CredentialId(credentialId).CredentialOwner(credentialOwner).Execute()

Lists all synthetic monitors in your Dynatrace environment



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
    managementZone := int64(789) // int64 | Filters the resulting set of monitors to those which are part of the specified management zone.    Specify the ID of the management zone here. (optional)
    tag := []string{"Inner_example"} // []string | Filters the resulting set of monitors by specified tags.   You can specify several tags in the following format: `tag=tag1&tag=tag2`. The monitor has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: `[context]key:value`. (optional)
    location := "location_example" // string | Filters the resulting set of monitors to those assigned to a specified Synthetic location.     Specify the ID of the location here. (optional)
    assignedApps := []string{"Inner_example"} // []string | Filters the resulting set of monitors to those assigned to the specified applications.   You can specify several applications in the following format: `assignedApps=app1&assignedApps=app2`. The monitor has to have **all** the specified applications assigned.   Specify Dynatrace entity IDs of applications here. (optional)
    type_ := "type__example" // string | Filters the resulting set of monitors to those of the specified type: `BROWSER` or `HTTP`. (optional)
    enabled := true // bool | Filters the resulting set of monitors to those which are enabled (`true`) or disabled (`false`). (optional)
    credentialId := "credentialId_example" // string | Filters the resulting set of monitors to those using the specified credential set.   Specify the ID of the credentials set here. (optional)
    credentialOwner := "credentialOwner_example" // string | Filters the resulting set of monitors to those using a credential owned by the specified user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticMonitorsApi.GetMonitorsCollection(context.Background()).ManagementZone(managementZone).Tag(tag).Location(location).AssignedApps(assignedApps).Type_(type_).Enabled(enabled).CredentialId(credentialId).CredentialOwner(credentialOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticMonitorsApi.GetMonitorsCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorsCollection`: Monitors
    fmt.Fprintf(os.Stdout, "Response from `SyntheticMonitorsApi.GetMonitorsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | **int64** | Filters the resulting set of monitors to those which are part of the specified management zone.    Specify the ID of the management zone here. | 
 **tag** | **[]string** | Filters the resulting set of monitors by specified tags.   You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The monitor has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: &#x60;[context]key:value&#x60;. | 
 **location** | **string** | Filters the resulting set of monitors to those assigned to a specified Synthetic location.     Specify the ID of the location here. | 
 **assignedApps** | **[]string** | Filters the resulting set of monitors to those assigned to the specified applications.   You can specify several applications in the following format: &#x60;assignedApps&#x3D;app1&amp;assignedApps&#x3D;app2&#x60;. The monitor has to have **all** the specified applications assigned.   Specify Dynatrace entity IDs of applications here. | 
 **type_** | **string** | Filters the resulting set of monitors to those of the specified type: &#x60;BROWSER&#x60; or &#x60;HTTP&#x60;. | 
 **enabled** | **bool** | Filters the resulting set of monitors to those which are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
 **credentialId** | **string** | Filters the resulting set of monitors to those using the specified credential set.   Specify the ID of the credentials set here. | 
 **credentialOwner** | **string** | Filters the resulting set of monitors to those using a credential owned by the specified user. | 

### Return type

[**Monitors**](Monitors.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMonitor

> ReplaceMonitor(ctx, monitorId).SyntheticMonitorUpdate(syntheticMonitorUpdate).Execute()

Updates parameters of the specified synthetic monitor

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
    monitorId := "monitorId_example" // string | The ID of the synthetic monitor to be updated.
    syntheticMonitorUpdate := *openapiclient.NewSyntheticMonitorUpdate(int32(123), "Type_example", "Name_example", []string{"Locations_example"}, false, map[string]interface{}(123), []openapiclient.TagWithSourceInfo{*openapiclient.NewTagWithSourceInfo("Context_example", "Key_example")}, []string{"ManuallyAssignedApps_example"}) // SyntheticMonitorUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticMonitorsApi.ReplaceMonitor(context.Background(), monitorId).SyntheticMonitorUpdate(syntheticMonitorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticMonitorsApi.ReplaceMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | The ID of the synthetic monitor to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticMonitorUpdate** | [**SyntheticMonitorUpdate**](SyntheticMonitorUpdate.md) |  | 

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

