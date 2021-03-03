# \TopologySmartscapeProcessApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProcesses**](TopologySmartscapeProcessApi.md#GetProcesses) | **Get** /entity/infrastructure/processes | Lists all monitored processes along with their parameters
[**GetSingleProcess**](TopologySmartscapeProcessApi.md#GetSingleProcess) | **Get** /entity/infrastructure/processes/{meIdentifier} | List properties of the specified process



## GetProcesses

> []ProcessGroupInstance GetProcesses(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).HostTag(hostTag).Host(host).ActualMonitoringState(actualMonitoringState).ExpectedMonitoringState(expectedMonitoringState).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all monitored processes along with their parameters



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
    startTimestamp := int64(789) // int64 | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. (optional)
    relativeTime := "relativeTime_example" // string | The relative timeframe, back from now. (optional)
    tag := []string{"Inner_example"} // []string | Filters the resulting set of processes by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The process has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified processes only.    To specify several processes use the following format: `entity=ID1&entity=ID2`. (optional)
    hostTag := []string{"Inner_example"} // []string | Filters processes by the host they're running at.    Specify tags of the host you're interested in. (optional)
    host := []string{"Inner_example"} // []string | Filters processes by the host they're running at.    Specify Dynatrace IDs of the host you're interested in.    To specify several hosts use the following format: `host=hostID1&host=hostID2`.    The **OR** logic applies. (optional)
    actualMonitoringState := "actualMonitoringState_example" // string | Filters processes by the actual monitoring state of the process. (optional)
    expectedMonitoringState := "expectedMonitoringState_example" // string | Filters processes by the expected monitoring state of the process. (optional)
    managementZone := int64(789) // int64 | Only return processes that are part of the specified management zone. (optional)
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    pageSize := int32(56) // int32 | The number of processes per result page.    If not set, pagination is not used and the result contains all processes fitting the specified filtering criteria. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you're using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeProcessApi.GetProcesses(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).HostTag(hostTag).Host(host).ActualMonitoringState(actualMonitoringState).ExpectedMonitoringState(expectedMonitoringState).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeProcessApi.GetProcesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcesses`: []ProcessGroupInstance
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeProcessApi.GetProcesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. | 
 **relativeTime** | **string** | The relative timeframe, back from now. | 
 **tag** | **[]string** | Filters the resulting set of processes by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The process has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified processes only.    To specify several processes use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **hostTag** | **[]string** | Filters processes by the host they&#39;re running at.    Specify tags of the host you&#39;re interested in. | 
 **host** | **[]string** | Filters processes by the host they&#39;re running at.    Specify Dynatrace IDs of the host you&#39;re interested in.    To specify several hosts use the following format: &#x60;host&#x3D;hostID1&amp;host&#x3D;hostID2&#x60;.    The **OR** logic applies. | 
 **actualMonitoringState** | **string** | Filters processes by the actual monitoring state of the process. | 
 **expectedMonitoringState** | **string** | Filters processes by the expected monitoring state of the process. | 
 **managementZone** | **int64** | Only return processes that are part of the specified management zone. | 
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **pageSize** | **int32** | The number of processes per result page.    If not set, pagination is not used and the result contains all processes fitting the specified filtering criteria. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**[]ProcessGroupInstance**](ProcessGroupInstance.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleProcess

> ProcessGroupInstance GetSingleProcess(ctx, meIdentifier).Execute()

List properties of the specified process

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required process.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeProcessApi.GetSingleProcess(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeProcessApi.GetSingleProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleProcess`: ProcessGroupInstance
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeProcessApi.GetSingleProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroupInstance**](ProcessGroupInstance.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

