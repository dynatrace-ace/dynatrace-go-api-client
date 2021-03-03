# \LogMonitoringProcessGroupsApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessGroupLogJobDelete**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobDelete) | **Delete** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
[**ProcessGroupLogJobRecords**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobRecords) | **Get** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
[**ProcessGroupLogJobRecordsFiltered**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobRecordsFiltered) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
[**ProcessGroupLogJobRecordsTop**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobRecordsTop) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
[**ProcessGroupLogJobStart**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobStart) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/{logPath} | Starts analysis job for the specified process group log
[**ProcessGroupLogJobStatus**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogJobStatus) | **Get** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
[**ProcessGroupLogList**](LogMonitoringProcessGroupsApi.md#ProcessGroupLogList) | **Get** /entity/infrastructure/process-groups/{pgId}/logs | Lists all the available logs of the specified process group



## ProcessGroupLogJobDelete

> LogJobDeleteResult ProcessGroupLogJobDelete(ctx, pgId, jobId).Execute()

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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    jobId := "jobId_example" // string | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobDelete(context.Background(), pgId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobDelete`: LogJobDeleteResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**jobId** | **string** | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobDeleteRequest struct via the builder pattern


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


## ProcessGroupLogJobRecords

> LogJobRecordsResult ProcessGroupLogJobRecords(ctx, pgId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()

Gets the content of the analyzed log



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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecords(context.Background(), pgId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobRecords`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobRecordsRequest struct via the builder pattern


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


## ProcessGroupLogJobRecordsFiltered

> LogJobRecordsResult ProcessGroupLogJobRecordsFiltered(ctx, pgId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()

Gets the content of the analyzed log



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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)
    filterLogContent := *openapiclient.NewFilterLogContent() // FilterLogContent | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsFiltered(context.Background(), pgId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsFiltered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobRecordsFiltered`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsFiltered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobRecordsFilteredRequest struct via the builder pattern


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


## ProcessGroupLogJobRecordsTop

> LogJobRecordsTopValuesRestResult ProcessGroupLogJobRecordsTop(ctx, pgId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()

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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
    filterTopLogRecords := *openapiclient.NewFilterTopLogRecords() // FilterTopLogRecords | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsTop(context.Background(), pgId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsTop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobRecordsTop`: LogJobRecordsTopValuesRestResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobRecordsTop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobRecordsTopRequest struct via the builder pattern


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


## ProcessGroupLogJobStart

> string ProcessGroupLogJobStart(ctx, pgId, logPath).HostFilter(hostFilter).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()

Starts analysis job for the specified process group log

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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    logPath := "logPath_example" // string | The full pathname of the log.
    hostFilter := "hostFilter_example" // string | Narrows down the scope of the analysis to process groups, running at the specified hosts.   Specify the entity ID of the required host here. To specify several IDs, separate them with a comma. (optional)
    query := "query_example" // string | Narrows down the scope of the analysis to the entries, matching the specified criteria.   The criteria must use the [text pattern query syntax](https://dt-url.net/vv83rhp). (optional)
    startTimestamp := int64(789) // int64 | The start timestamp of the analysis range, in UTC milliseconds.    If not set, then 2 hours behind from current timestamp is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the analysis range, in UTC milliseconds.    If not set, then the current timestamp is used. (optional)
    extractFields := *openapiclient.NewExtractFields() // ExtractFields | Extract fields from the log content to form custom columns.    See the [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) help page for the syntax definition and examples.   The special characters must be escaped. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobStart(context.Background(), pgId, logPath).HostFilter(hostFilter).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobStart`: string
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**logPath** | **string** | The full pathname of the log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostFilter** | **string** | Narrows down the scope of the analysis to process groups, running at the specified hosts.   Specify the entity ID of the required host here. To specify several IDs, separate them with a comma. | 
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


## ProcessGroupLogJobStatus

> LogJobStatusResult ProcessGroupLogJobStatus(ctx, pgId, jobId).Execute()

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
    pgId := "pgId_example" // string | The Dynatrace entity ID of the required process group.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogJobStatus(context.Background(), pgId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogJobStatus`: LogJobStatusResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The Dynatrace entity ID of the required process group. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogJobStatusRequest struct via the builder pattern


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


## ProcessGroupLogList

> LogList4pgResult ProcessGroupLogList(ctx, pgId).Execute()

Lists all the available logs of the specified process group

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
    pgId := "pgId_example" // string | The entity ID of the process group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringProcessGroupsApi.ProcessGroupLogList(context.Background(), pgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringProcessGroupsApi.ProcessGroupLogList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessGroupLogList`: LogList4pgResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringProcessGroupsApi.ProcessGroupLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pgId** | **string** | The entity ID of the process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessGroupLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogList4pgResult**](LogList4pgResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

