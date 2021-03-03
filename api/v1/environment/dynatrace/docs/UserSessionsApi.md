# \UserSessionsApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDTAQLResultAsTable**](UserSessionsApi.md#GetDTAQLResultAsTable) | **Get** /userSessionQueryLanguage/table | Returns the result of the query as a table structure
[**GetDTAQLResultAsTree**](UserSessionsApi.md#GetDTAQLResultAsTree) | **Get** /userSessionQueryLanguage/tree | Returns the result of the query as a tree structure



## GetDTAQLResultAsTable

> DtaqlResultAsTable GetDTAQLResultAsTable(ctx).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).OffsetUTC(offsetUTC).PageSize(pageSize).PageOffset(pageOffset).AddDeepLinkFields(addDeepLinkFields).Explain(explain).Execute()

Returns the result of the query as a table structure



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
    query := "query_example" // string | The user session query to be executed. See the [USQL help page](https://dt-url.net/dtusql) for syntax details.    You can find the available columns of the **usersession** table in the `UserSession` object.    Here is an example of the query: `SELECT country, city, COUNT(*) FROM usersession GROUP BY country, city`.
    startTimestamp := int64(789) // int64 | The start timestamp of the query, in UTC milliseconds.   If not set or set as `0`, 2 hours behind the current time is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the query, in UTC milliseconds.   If not set or set as `0`, the current timestamp is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). (optional)
    offsetUTC := int32(56) // int32 | Optional offset of local time to UTC time in minutes. Offset will be applied to Date fields encountered in the query.   Can be positive or negative. E.g. if the local time is UTC+02:00, the timeOffset is 120. If it is UTC-05:00, timeOffset is -300. (optional)
    pageSize := int32(56) // int32 | Optional limit on how many of the actual query results should be returned in the tabular result. (optional)
    pageOffset := int32(56) // int32 | Optional offset of the requested results from the start of tabular results. Relates to pageSize.   E.g. on a query that might return 500 results, you might want to receive results in chunks of 50 rows.   this can be achieved by using pageSize=50, and setting pageOffset in subsequent calls.In the example adding pageOffset=50 returns result rows 51-100. (optional)
    addDeepLinkFields := true // bool | Add (`true`) to enable deep linking of additional fields in the query.   If not set, then `false` is used (optional) (default to false)
    explain := true // bool | Add (`true`) or don't add (`false`) some additional information about the result to the response.   It helps to understand the query and how the result was calculated.   If not set, then `false` is used (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSessionsApi.GetDTAQLResultAsTable(context.Background()).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).OffsetUTC(offsetUTC).PageSize(pageSize).PageOffset(pageOffset).AddDeepLinkFields(addDeepLinkFields).Explain(explain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionsApi.GetDTAQLResultAsTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDTAQLResultAsTable`: DtaqlResultAsTable
    fmt.Fprintf(os.Stdout, "Response from `UserSessionsApi.GetDTAQLResultAsTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDTAQLResultAsTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The user session query to be executed. See the [USQL help page](https://dt-url.net/dtusql) for syntax details.    You can find the available columns of the **usersession** table in the &#x60;UserSession&#x60; object.    Here is an example of the query: &#x60;SELECT country, city, COUNT(*) FROM usersession GROUP BY country, city&#x60;. | 
 **startTimestamp** | **int64** | The start timestamp of the query, in UTC milliseconds.   If not set or set as &#x60;0&#x60;, 2 hours behind the current time is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). | 
 **endTimestamp** | **int64** | The end timestamp of the query, in UTC milliseconds.   If not set or set as &#x60;0&#x60;, the current timestamp is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). | 
 **offsetUTC** | **int32** | Optional offset of local time to UTC time in minutes. Offset will be applied to Date fields encountered in the query.   Can be positive or negative. E.g. if the local time is UTC+02:00, the timeOffset is 120. If it is UTC-05:00, timeOffset is -300. | 
 **pageSize** | **int32** | Optional limit on how many of the actual query results should be returned in the tabular result. | 
 **pageOffset** | **int32** | Optional offset of the requested results from the start of tabular results. Relates to pageSize.   E.g. on a query that might return 500 results, you might want to receive results in chunks of 50 rows.   this can be achieved by using pageSize&#x3D;50, and setting pageOffset in subsequent calls.In the example adding pageOffset&#x3D;50 returns result rows 51-100. | 
 **addDeepLinkFields** | **bool** | Add (&#x60;true&#x60;) to enable deep linking of additional fields in the query.   If not set, then &#x60;false&#x60; is used | [default to false]
 **explain** | **bool** | Add (&#x60;true&#x60;) or don&#39;t add (&#x60;false&#x60;) some additional information about the result to the response.   It helps to understand the query and how the result was calculated.   If not set, then &#x60;false&#x60; is used | [default to false]

### Return type

[**DtaqlResultAsTable**](DtaqlResultAsTable.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDTAQLResultAsTree

> DtaqlResultAsTree GetDTAQLResultAsTree(ctx).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).OffsetUTC(offsetUTC).AddDeepLinkFields(addDeepLinkFields).Explain(explain).Execute()

Returns the result of the query as a tree structure



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
    query := "query_example" // string | The user session query to be executed. See the [USQL help page](https://dt-url.net/dtusql) for syntax details.    You can find the available columns of the **usersession** table in the `UserSession` object.    Here is an example of the query: `SELECT country, city, COUNT(*) FROM usersession GROUP BY country, city`.
    startTimestamp := int64(789) // int64 | The start timestamp of the query, in UTC milliseconds.   If not set or set as `0`, 2 hours behind the current time is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the query, in UTC milliseconds.   If not set or set as `0`, the current timestamp is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). (optional)
    offsetUTC := int32(56) // int32 | Optional offset of local time to UTC time in minutes. Offset will be applied to Date fields encountered in the query.   Can be positive or negative. E.g. if the local time is UTC+02:00, the timeOffset is 120. If it is UTC-05:00, timeOffset is -300. (optional)
    addDeepLinkFields := true // bool | Add (`true`) to enable deep linking of additional fields in the query.   If not set, then `false` is used (optional) (default to false)
    explain := true // bool | Add (`true`) or don't add (`false`) some additional information about the result to the response.   It helps to understand the query and how the result was calculated.   If not set, then `false` is used (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSessionsApi.GetDTAQLResultAsTree(context.Background()).Query(query).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).OffsetUTC(offsetUTC).AddDeepLinkFields(addDeepLinkFields).Explain(explain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionsApi.GetDTAQLResultAsTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDTAQLResultAsTree`: DtaqlResultAsTree
    fmt.Fprintf(os.Stdout, "Response from `UserSessionsApi.GetDTAQLResultAsTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDTAQLResultAsTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The user session query to be executed. See the [USQL help page](https://dt-url.net/dtusql) for syntax details.    You can find the available columns of the **usersession** table in the &#x60;UserSession&#x60; object.    Here is an example of the query: &#x60;SELECT country, city, COUNT(*) FROM usersession GROUP BY country, city&#x60;. | 
 **startTimestamp** | **int64** | The start timestamp of the query, in UTC milliseconds.   If not set or set as &#x60;0&#x60;, 2 hours behind the current time is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). | 
 **endTimestamp** | **int64** | The end timestamp of the query, in UTC milliseconds.   If not set or set as &#x60;0&#x60;, the current timestamp is used.   If the exact times are important, set the timeframe in the query itself (**query** parameter). | 
 **offsetUTC** | **int32** | Optional offset of local time to UTC time in minutes. Offset will be applied to Date fields encountered in the query.   Can be positive or negative. E.g. if the local time is UTC+02:00, the timeOffset is 120. If it is UTC-05:00, timeOffset is -300. | 
 **addDeepLinkFields** | **bool** | Add (&#x60;true&#x60;) to enable deep linking of additional fields in the query.   If not set, then &#x60;false&#x60; is used | [default to false]
 **explain** | **bool** | Add (&#x60;true&#x60;) or don&#39;t add (&#x60;false&#x60;) some additional information about the result to the response.   It helps to understand the query and how the result was calculated.   If not set, then &#x60;false&#x60; is used | [default to false]

### Return type

[**DtaqlResultAsTree**](DtaqlResultAsTree.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

