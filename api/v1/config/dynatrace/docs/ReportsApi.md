# \ReportsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateReport**](ReportsApi.md#CreateOrUpdateReport) | **Put** /reports/{id} | Updates an existing report | maturity&#x3D;EARLY_ADOPTER
[**CreateReport**](ReportsApi.md#CreateReport) | **Post** /reports | Creates a report | maturity&#x3D;EARLY_ADOPTER
[**DeleteReport**](ReportsApi.md#DeleteReport) | **Delete** /reports/{id} | Deletes the specified report | maturity&#x3D;EARLY_ADOPTER
[**GetReport**](ReportsApi.md#GetReport) | **Get** /reports/{id} | Gets the properties of the specified report | maturity&#x3D;EARLY_ADOPTER
[**ListReports**](ReportsApi.md#ListReports) | **Get** /reports | Lists all reports available in your environment | maturity&#x3D;EARLY_ADOPTER
[**SubscribeReport**](ReportsApi.md#SubscribeReport) | **Post** /reports/{id}/subscribe | Subscribes to the specified report | maturity&#x3D;EARLY_ADOPTER
[**UnsubscribeReport**](ReportsApi.md#UnsubscribeReport) | **Post** /reports/{id}/unsubscribe | Unsubscribes from the specified report | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateOrUpdateReport**](ReportsApi.md#ValidateCreateOrUpdateReport) | **Post** /reports/{id}/validator | Validates the payload for the &#x60;PUT /reports/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateReport**](ReportsApi.md#ValidateCreateReport) | **Post** /reports/validator | Validates the payload for the &#x60;POST /reports&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateReport

> EntityShortRepresentation CreateOrUpdateReport(ctx, id).DashboardReport(dashboardReport).Execute()

Updates an existing report | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the report to be updated.    The ID in the request body must match this ID.
    dashboardReport := *openapiclient.NewDashboardReport("Type_example", "DashboardId_example", *openapiclient.NewDashboardReportSubscription()) // DashboardReport | The JSON body of the request. Contains updated parameters of the report. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.CreateOrUpdateReport(context.Background(), id).DashboardReport(dashboardReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateOrUpdateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateReport`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateOrUpdateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the report to be updated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardReport** | [**DashboardReport**](DashboardReport.md) | The JSON body of the request. Contains updated parameters of the report. | 

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


## CreateReport

> EntityShortRepresentation CreateReport(ctx).DashboardReport(dashboardReport).Execute()

Creates a report | maturity=EARLY_ADOPTER



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
    dashboardReport := *openapiclient.NewDashboardReport("Type_example", "DashboardId_example", *openapiclient.NewDashboardReportSubscription()) // DashboardReport | The JSON body of the request. Contains parameters of the new report. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.CreateReport(context.Background()).DashboardReport(dashboardReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.CreateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReport`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.CreateReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardReport** | [**DashboardReport**](DashboardReport.md) | The JSON body of the request. Contains parameters of the new report. | 

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


## DeleteReport

> DeleteReport(ctx, id).Execute()

Deletes the specified report | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the report to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.DeleteReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.DeleteReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the report to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportRequest struct via the builder pattern


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


## GetReport

> DashboardReport GetReport(ctx, id).Execute()

Gets the properties of the specified report | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required report.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.GetReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: DashboardReport
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardReport**](DashboardReport.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportStubList ListReports(ctx).Type_(type_).SourceId(sourceId).Execute()

Lists all reports available in your environment | maturity=EARLY_ADOPTER

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
    type_ := "type__example" // string | Type of a report. (optional)
    sourceId := "sourceId_example" // string | Referencing source entity of a report (e.g. dashboard). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ListReports(context.Background()).Type_(type_).SourceId(sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ListReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReports`: ReportStubList
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of a report. | 
 **sourceId** | **string** | Referencing source entity of a report (e.g. dashboard). | 

### Return type

[**ReportStubList**](ReportStubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeReport

> EntityShortRepresentation SubscribeReport(ctx, id).ReportSubscriptions(reportSubscriptions).Execute()

Subscribes to the specified report | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the report to subscribe to.
    reportSubscriptions := *openapiclient.NewReportSubscriptions("Schedule_example", []string{"Recipients_example"}) // ReportSubscriptions | The JSON body of the request. Contains a list of new subscribers. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.SubscribeReport(context.Background(), id).ReportSubscriptions(reportSubscriptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.SubscribeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribeReport`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.SubscribeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the report to subscribe to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportSubscriptions** | [**ReportSubscriptions**](ReportSubscriptions.md) | The JSON body of the request. Contains a list of new subscribers. | 

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


## UnsubscribeReport

> EntityShortRepresentation UnsubscribeReport(ctx, id).ReportSubscriptions(reportSubscriptions).Execute()

Unsubscribes from the specified report | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the report to unsubscribe from.
    reportSubscriptions := *openapiclient.NewReportSubscriptions("Schedule_example", []string{"Recipients_example"}) // ReportSubscriptions | The JSON body of the request. Contains a list of recipients to be unsubscribed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.UnsubscribeReport(context.Background(), id).ReportSubscriptions(reportSubscriptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.UnsubscribeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsubscribeReport`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.UnsubscribeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the report to unsubscribe from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportSubscriptions** | [**ReportSubscriptions**](ReportSubscriptions.md) | The JSON body of the request. Contains a list of recipients to be unsubscribed. | 

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


## ValidateCreateOrUpdateReport

> ValidateCreateOrUpdateReport(ctx, id).DashboardReport(dashboardReport).Execute()

Validates the payload for the `PUT /reports/{id}` request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the report to be validated.    The ID in the request body must match this ID.
    dashboardReport := *openapiclient.NewDashboardReport("Type_example", "DashboardId_example", *openapiclient.NewDashboardReportSubscription()) // DashboardReport | The JSON body of the request. Contains the report to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ValidateCreateOrUpdateReport(context.Background(), id).DashboardReport(dashboardReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ValidateCreateOrUpdateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the report to be validated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateOrUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardReport** | [**DashboardReport**](DashboardReport.md) | The JSON body of the request. Contains the report to be validated. | 

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


## ValidateCreateReport

> ValidateCreateReport(ctx).DashboardReport(dashboardReport).Execute()

Validates the payload for the `POST /reports` request | maturity=EARLY_ADOPTER



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
    dashboardReport := *openapiclient.NewDashboardReport("Type_example", "DashboardId_example", *openapiclient.NewDashboardReportSubscription()) // DashboardReport | The JSON body of the request. Contains the report to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ValidateCreateReport(context.Background()).DashboardReport(dashboardReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ValidateCreateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardReport** | [**DashboardReport**](DashboardReport.md) | The JSON body of the request. Contains the report to be validated. | 

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

