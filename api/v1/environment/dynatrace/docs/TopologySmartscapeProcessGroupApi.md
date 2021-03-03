# \TopologySmartscapeProcessGroupApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProcessGroups**](TopologySmartscapeProcessGroupApi.md#GetProcessGroups) | **Get** /entity/infrastructure/process-groups | Lists all process groups of your environment, along with their parameters
[**GetSingleProcessGroup**](TopologySmartscapeProcessGroupApi.md#GetSingleProcessGroup) | **Get** /entity/infrastructure/process-groups/{meIdentifier} | List properties of the specified process group
[**UpdateProcessGroup**](TopologySmartscapeProcessGroupApi.md#UpdateProcessGroup) | **Post** /entity/infrastructure/process-groups/{meIdentifier} | Updates properties of the specified process group



## GetProcessGroups

> []ProcessGroup GetProcessGroups(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).Host(host).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all process groups of your environment, along with their parameters



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
    tag := []string{"Inner_example"} // []string | Filters the resulting set of process groups by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The process group has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified process groups only.    To specify several process groups use the following format: `entity=ID1&entity=ID2`. (optional)
    host := []string{"Inner_example"} // []string | Filters process groups by the host they're running at.    Specify Dynatrace IDs of the host you're interested in.    To specify several hosts use the following format: `host=hostID1&host=hostID2`.    The **OR** logic applies. (optional)
    managementZone := int64(789) // int64 | Only return process groups that are part of the specified management zone. (optional)
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    pageSize := int32(56) // int32 | The number of process groups per result page.    If not set, pagination is not used and the result contains all process groups fitting the specified filtering criteria. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you're using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeProcessGroupApi.GetProcessGroups(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).Host(host).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeProcessGroupApi.GetProcessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessGroups`: []ProcessGroup
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeProcessGroupApi.GetProcessGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. | 
 **relativeTime** | **string** | The relative timeframe, back from now. | 
 **tag** | **[]string** | Filters the resulting set of process groups by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The process group has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified process groups only.    To specify several process groups use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **host** | **[]string** | Filters process groups by the host they&#39;re running at.    Specify Dynatrace IDs of the host you&#39;re interested in.    To specify several hosts use the following format: &#x60;host&#x3D;hostID1&amp;host&#x3D;hostID2&#x60;.    The **OR** logic applies. | 
 **managementZone** | **int64** | Only return process groups that are part of the specified management zone. | 
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **pageSize** | **int32** | The number of process groups per result page.    If not set, pagination is not used and the result contains all process groups fitting the specified filtering criteria. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**[]ProcessGroup**](ProcessGroup.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleProcessGroup

> ProcessGroup GetSingleProcessGroup(ctx, meIdentifier).Execute()

List properties of the specified process group

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required process group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeProcessGroupApi.GetSingleProcessGroup(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeProcessGroupApi.GetSingleProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleProcessGroup`: ProcessGroup
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeProcessGroupApi.GetSingleProcessGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProcessGroup**](ProcessGroup.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProcessGroup

> UpdateProcessGroup(ctx, meIdentifier).UpdateEntity(updateEntity).Execute()

Updates properties of the specified process group

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the process group to be updated.
    updateEntity := *openapiclient.NewUpdateEntity([]string{"Tags_example"}) // UpdateEntity | The JSON body of the request. Contains tags to be added to the process group. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeProcessGroupApi.UpdateProcessGroup(context.Background(), meIdentifier).UpdateEntity(updateEntity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeProcessGroupApi.UpdateProcessGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the process group to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntity** | [**UpdateEntity**](UpdateEntity.md) | The JSON body of the request. Contains tags to be added to the process group. | 

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

