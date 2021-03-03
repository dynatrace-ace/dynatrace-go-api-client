# \LogMonitoringCustomDevicesApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomDeviceLogJobDelete**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobDelete) | **Delete** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
[**CustomDeviceLogJobRecords**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobRecords) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
[**CustomDeviceLogJobRecordsFiltered**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobRecordsFiltered) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records | Gets the filtered content of the analyzed log
[**CustomDeviceLogJobRecordsTop**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobRecordsTop) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
[**CustomDeviceLogJobStart**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobStart) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/{logPath} | Starts the analysis job for the specified custom device log
[**CustomDeviceLogJobStatus**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogJobStatus) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
[**CustomDeviceLogList**](LogMonitoringCustomDevicesApi.md#CustomDeviceLogList) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs | Lists all the available logs on the specified custom device



## CustomDeviceLogJobDelete

> LogJobDeleteResult CustomDeviceLogJobDelete(ctx, customDeviceId, jobId).Execute()

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    jobId := "jobId_example" // string | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobDelete(context.Background(), customDeviceId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobDelete`: LogJobDeleteResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**jobId** | **string** | The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobDeleteRequest struct via the builder pattern


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


## CustomDeviceLogJobRecords

> LogJobRecordsResult CustomDeviceLogJobRecords(ctx, customDeviceId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecords(context.Background(), customDeviceId, jobId).ScrollToken(scrollToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobRecords`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobRecordsRequest struct via the builder pattern


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


## CustomDeviceLogJobRecordsFiltered

> LogJobRecordsResult CustomDeviceLogJobRecordsFiltered(ctx, customDeviceId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request.
    scrollToken := "scrollToken_example" // string | The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned. (optional)
    pageSize := int32(56) // int32 | The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is `10000`. (optional)
    filterLogContent := *openapiclient.NewFilterLogContent() // FilterLogContent | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsFiltered(context.Background(), customDeviceId, jobId).ScrollToken(scrollToken).PageSize(pageSize).FilterLogContent(filterLogContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsFiltered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobRecordsFiltered`: LogJobRecordsResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsFiltered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobRecordsFilteredRequest struct via the builder pattern


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


## CustomDeviceLogJobRecordsTop

> LogJobRecordsTopValuesRestResult CustomDeviceLogJobRecordsTop(ctx, customDeviceId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request.
    filterTopLogRecords := *openapiclient.NewFilterTopLogRecords() // FilterTopLogRecords | Filter the log content by the specified criteria.   See the [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) help page for the syntax definition and examples. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsTop(context.Background(), customDeviceId, jobId).FilterTopLogRecords(filterTopLogRecords).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsTop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobRecordsTop`: LogJobRecordsTopValuesRestResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobRecordsTop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobRecordsTopRequest struct via the builder pattern


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


## CustomDeviceLogJobStart

> string CustomDeviceLogJobStart(ctx, customDeviceId, logPath).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()

Starts the analysis job for the specified custom device log



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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    logPath := "logPath_example" // string | The full pathname of the log.
    query := "query_example" // string | Narrows down the scope of the analysis to the entries, matching the specified criteria.   The criteria must use the [text pattern query syntax](https://dt-url.net/vv83rhp). (optional)
    startTimestamp := int64(789) // int64 | The start timestamp of the analysis range, in UTC milliseconds.    If not set, then 2 hours behind from current timestamp is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the analysis range, in UTC milliseconds.    If not set, then the current timestamp is used. (optional)
    extractFields := *openapiclient.NewExtractFields() // ExtractFields | Extract fields from the log content to form custom columns.    See the [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) help page for the syntax definition and examples.   The special characters must be escaped. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobStart(context.Background(), customDeviceId, logPath).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).ExtractFields(extractFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobStart`: string
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**logPath** | **string** | The full pathname of the log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobStartRequest struct via the builder pattern


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


## CustomDeviceLogJobStatus

> LogJobStatusResult CustomDeviceLogJobStatus(ctx, customDeviceId, jobId).Execute()

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.
    jobId := "jobId_example" // string | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogJobStatus(context.Background(), customDeviceId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogJobStatus`: LogJobStatusResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 
**jobId** | **string** | The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/kzi3rb8) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogJobStatusRequest struct via the builder pattern


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


## CustomDeviceLogList

> LogListForCustomDeviceResult CustomDeviceLogList(ctx, customDeviceId).Execute()

Lists all the available logs on the specified custom device

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
    customDeviceId := "customDeviceId_example" // string | The Dynatrace entity ID of the required custom device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogMonitoringCustomDevicesApi.CustomDeviceLogList(context.Background(), customDeviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogMonitoringCustomDevicesApi.CustomDeviceLogList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomDeviceLogList`: LogListForCustomDeviceResult
    fmt.Fprintf(os.Stdout, "Response from `LogMonitoringCustomDevicesApi.CustomDeviceLogList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDeviceId** | **string** | The Dynatrace entity ID of the required custom device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomDeviceLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogListForCustomDeviceResult**](LogListForCustomDeviceResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

