# \LogMonitoringHostsApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostLogJobDelete**](LogMonitoringHostsApi.md#HostLogJobDelete) | **Delete** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
[**HostLogJobRecords**](LogMonitoringHostsApi.md#HostLogJobRecords) | **Get** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records | Gets the full content of the analyzed log
[**HostLogJobRecordsFiltered**](LogMonitoringHostsApi.md#HostLogJobRecordsFiltered) | **Post** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records | Gets the filtered content of the analyzed log
[**HostLogJobRecordsTop**](LogMonitoringHostsApi.md#HostLogJobRecordsTop) | **Post** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
[**HostLogJobStart**](LogMonitoringHostsApi.md#HostLogJobStart) | **Post** /entity/infrastructure/hosts/{hostId}/logs/{logPath} | Starts the analysis job for the specified OS log
[**HostLogJobStatus**](LogMonitoringHostsApi.md#HostLogJobStatus) | **Get** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
[**HostLogList**](LogMonitoringHostsApi.md#HostLogList) | **Get** /entity/infrastructure/hosts/{hostId}/logs | Lists all the available OS logs on the specified host



## HostLogJobDelete

> LogJobDeleteResult HostLogJobDelete(ctx, hostId, jobId).Execute()

Deletes or cancels the specified log analysis job

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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    jobId := "jobId_example" // string | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobDelete(context.Background(), hostId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobDelete`: LogJobDeleteResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**jobId** | **string** | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogJobDeleteResult**](LogJobDeleteResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogJobRecords

> LogJobRecordsResult HostLogJobRecords(ctx, hostId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()

Gets the full content of the analyzed log



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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobRecords(context.Background(), hostId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobRecords`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scrollToken** | **string** | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. | 
 **pageSize** | **int32** | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is &#x60;10000&#x60;. | 

### Return type

[**LogJobRecordsResult**](LogJobRecordsResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogJobRecordsFiltered

> LogJobRecordsResult HostLogJobRecordsFiltered(ctx, hostId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()

Gets the filtered content of the analyzed log



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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)
    filterLogContent := *openapiclient.NewFilterLogContent() // FilterLogContent | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobRecordsFiltered(context.Background(), hostId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobRecordsFiltered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobRecordsFiltered`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobRecordsFiltered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobRecordsFilteredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scrollToken** | **string** | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. | 
 **pageSize** | **int32** | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is &#x60;10000&#x60;. | 
 **filterLogContent** | [**FilterLogContent**](FilterLogContent.md) | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. | 

### Return type

[**LogJobRecordsResult**](LogJobRecordsResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogJobRecordsTop

> LogJobRecordsTopValuesRestResult HostLogJobRecordsTop(ctx, hostId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()

Gets the top values of fields present in the content of the analyzed log



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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request.
    filterTopLogRecords := *openapiclient.NewFilterTopLogRecords() // FilterTopLogRecords | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobRecordsTop(context.Background(), hostId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobRecordsTop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobRecordsTop`: LogJobRecordsTopValuesRestResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobRecordsTop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobRecordsTopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterTopLogRecords** | [**FilterTopLogRecords**](FilterTopLogRecords.md) | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. | 

### Return type

[**LogJobRecordsTopValuesRestResult**](LogJobRecordsTopValuesRestResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogJobStart

> string HostLogJobStart(ctx, hostId, logPath).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()

Starts the analysis job for the specified OS log



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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    logPath := "logPath_example" // string | The full pathname of the log.
    query := "query_example" // string | Narrows down the scope of the analysis to the entries, matching the specified criteria.   The criteria must use the [text pattern query syntax](https://dt-url.net/vv83rhp). (optional)
    startTimestamp := int64(789) // int64 | The start timestamp of the analysis range, in UTC milliseconds.    If not set, then 2 hours behind from current timestamp is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the analysis range, in UTC milliseconds.    If not set, then the current timestamp is used. (optional)
    extractFields := *openapiclient.NewExtractFields() // ExtractFields | Extract fields from the log content to form custom columns.    See the [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) help page for the syntax definition and examples.   The special characters must be escaped. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobStart(context.Background(), hostId, logPath).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobStart`: string
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**logPath** | **string** | The full pathname of the log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | Narrows down the scope of the analysis to the entries, matching the specified criteria.   The criteria must use the [text pattern query syntax](https://dt-url.net/vv83rhp). | 
 **startTimestamp** | **int64** | The start timestamp of the analysis range, in UTC milliseconds.    If not set, then 2 hours behind from current timestamp is used. | 
 **endTimestamp** | **int64** | The end timestamp of the analysis range, in UTC milliseconds.    If not set, then the current timestamp is used. | 
 **extractFields** | [**ExtractFields**](ExtractFields.md) | Extract fields from the log content to form custom columns.    See the [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) help page for the syntax definition and examples.   The special characters must be escaped. | 

### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogJobStatus

> LogJobStatusResult HostLogJobStatus(ctx, hostId, jobId).Execute()

Gets status of the specified log analysis job

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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogJobStatus(context.Background(), hostId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogJobStatus`: LogJobStatusResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/52k3r7f) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogJobStatusResult**](LogJobStatusResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostLogList

> LogList4hostResult HostLogList(ctx, hostId).Execute()

Lists all the available OS logs on the specified host

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
    hostId := "hostId_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringHostsApi.HostLogList(context.Background(), hostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringHostsApi.HostLogList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostLogList`: LogList4hostResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringHostsApi.HostLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogList4hostResult**](LogList4hostResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

