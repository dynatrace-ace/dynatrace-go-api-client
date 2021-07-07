# \AuditLogsApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLog**](AuditLogsApi.md#GetLog) | **Get** /auditlogs/{id} | Gets the specified entry of the audit log | maturity&#x3D;EARLY_ADOPTER
[**GetLogs**](AuditLogsApi.md#GetLogs) | **Get** /auditlogs | Gets the audit log of your Dynatrace environment | maturity&#x3D;EARLY_ADOPTER



## GetLog

> AuditLogEntry GetLog(ctx, id).Execute()

Gets the specified entry of the audit log | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required log entry.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsApi.GetLog(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.GetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLog`: AuditLogEntry
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required log entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLogEntry**](AuditLogEntry.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> AuditLog GetLogs(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Filter(filter).From(from).To(to).Sort(sort).Execute()

Gets the audit log of your Dynatrace environment | maturity=EARLY_ADOPTER



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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of log entries in a single response payload.   The maximal allowed page size is 5000.   If not set, 1000 is used. (optional)
    filter := "filter_example" // string | Filters the audit log. You can use the following criteria:  * User: `user(\"userIdentification\")`. The `EQUALS` operator applies.  * Event type: `eventType(\"value\")`. The `EQUALS` operator applies.  * Category of a logged operation: `category(\"value\")`. The `EQUALS` operator applies.  * Entity ID: `entityId(\"id\")`. The `CONTAINS` operator applies.   For each criterion, you can specify multiple alternatives with comma-separated values. In this case, the OR logic applies. For example, `eventType(\"CREATE\",\"UPDATE\")` means eventType can be \"CREATE\" or \"UPDATE\".   You can specify multiple comma-separated criteria, such as `eventType(\"CREATE\",\"UPDATE\"),category(\"CONFIG\")`. Only results matching **all** criteria are included in response.   Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes:  * Tilde `~`  * Quote `\"` (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    sort := "sort_example" // string | The sorting of audit log entries:  * `timestamp`: Oldest first.  * `-timestamp`: Newest first.   If not set, the newest first sorting is applied. (optional) (default to "-timestamp")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsApi.GetLogs(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Filter(filter).From(from).To(to).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: AuditLog
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.GetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of log entries in a single response payload.   The maximal allowed page size is 5000.   If not set, 1000 is used. | 
 **filter** | **string** | Filters the audit log. You can use the following criteria:  * User: &#x60;user(\&quot;userIdentification\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator applies.  * Event type: &#x60;eventType(\&quot;value\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator applies.  * Category of a logged operation: &#x60;category(\&quot;value\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator applies.  * Entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator applies.   For each criterion, you can specify multiple alternatives with comma-separated values. In this case, the OR logic applies. For example, &#x60;eventType(\&quot;CREATE\&quot;,\&quot;UPDATE\&quot;)&#x60; means eventType can be \&quot;CREATE\&quot; or \&quot;UPDATE\&quot;.   You can specify multiple comma-separated criteria, such as &#x60;eventType(\&quot;CREATE\&quot;,\&quot;UPDATE\&quot;),category(\&quot;CONFIG\&quot;)&#x60;. Only results matching **all** criteria are included in response.   Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes:  * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **sort** | **string** | The sorting of audit log entries:  * &#x60;timestamp&#x60;: Oldest first.  * &#x60;-timestamp&#x60;: Newest first.   If not set, the newest first sorting is applied. | [default to &quot;-timestamp&quot;]

### Return type

[**AuditLog**](AuditLog.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

