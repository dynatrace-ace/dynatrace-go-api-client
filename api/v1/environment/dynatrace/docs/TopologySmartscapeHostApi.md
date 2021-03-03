# \TopologySmartscapeHostApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHosts**](TopologySmartscapeHostApi.md#GetHosts) | **Get** /entity/infrastructure/hosts | Lists all available hosts in your environment
[**GetSingleHost**](TopologySmartscapeHostApi.md#GetSingleHost) | **Get** /entity/infrastructure/hosts/{meIdentifier} | Gets parameters of the specified host
[**RemoveTags**](TopologySmartscapeHostApi.md#RemoveTags) | **Delete** /entity/infrastructure/hosts/{meIdentifier}/tags/{tag} | Remove tag of the specified host
[**UpdateHost**](TopologySmartscapeHostApi.md#UpdateHost) | **Post** /entity/infrastructure/hosts/{meIdentifier} | Updates properties of the specified host



## GetHosts

> []Host GetHosts(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).ShowMonitoringCandidates(showMonitoringCandidates).Entity(entity).ManagementZone(managementZone).HostGroupId(hostGroupId).HostGroupName(hostGroupName).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all available hosts in your environment



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
    tag := []string{"Inner_example"} // []string | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    showMonitoringCandidates := true // bool | Includes (`true`) or excludes (`false`) a monitoring candidate in the response.   Monitoring candidates are network entities that are detected but not monitored. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified hosts only.    To specify several hosts use the following format: `entity=ID1&entity=ID2`. (optional)
    managementZone := int64(789) // int64 | Only return hosts that are part of the specified management zone. (optional)
    hostGroupId := "hostGroupId_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you're interested in. (optional)
    hostGroupName := "hostGroupName_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you're interested in. (optional)
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    pageSize := int32(56) // int32 | The number of hosts per result page.    If not set, pagination is not used and the result contains all hosts fitting the specified filtering criteria. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you're using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeHostApi.GetHosts(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).ShowMonitoringCandidates(showMonitoringCandidates).Entity(entity).ManagementZone(managementZone).HostGroupId(hostGroupId).HostGroupName(hostGroupName).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeHostApi.GetHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHosts`: []Host
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeHostApi.GetHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. | 
 **relativeTime** | **string** | The relative timeframe, back from now. | 
 **tag** | **[]string** | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **showMonitoringCandidates** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) a monitoring candidate in the response.   Monitoring candidates are network entities that are detected but not monitored. | 
 **entity** | **[]string** | Filters result to the specified hosts only.    To specify several hosts use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **managementZone** | **int64** | Only return hosts that are part of the specified management zone. | 
 **hostGroupId** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you&#39;re interested in. | 
 **hostGroupName** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you&#39;re interested in. | 
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **pageSize** | **int32** | The number of hosts per result page.    If not set, pagination is not used and the result contains all hosts fitting the specified filtering criteria. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**[]Host**](Host.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleHost

> Host GetSingleHost(ctx, meIdentifier).Execute()

Gets parameters of the specified host

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeHostApi.GetSingleHost(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeHostApi.GetSingleHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleHost`: Host
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeHostApi.GetSingleHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Host**](Host.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTags

> RemoveTags(ctx, meIdentifier, tag).Execute()

Remove tag of the specified host

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the host.
    tag := "tag_example" // string | The tag to be removed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeHostApi.RemoveTags(context.Background(), meIdentifier, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeHostApi.RemoveTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the host. | 
**tag** | **string** | The tag to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> UpdateHost(ctx, meIdentifier).UpdateEntity(updateEntity).Execute()

Updates properties of the specified host

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the host to be updated.
    updateEntity := *openapiclient.NewUpdateEntity([]string{"Tags_example"}) // UpdateEntity |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeHostApi.UpdateHost(context.Background(), meIdentifier).UpdateEntity(updateEntity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeHostApi.UpdateHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the host to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntity** | [**UpdateEntity**](UpdateEntity.md) |  | 

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

