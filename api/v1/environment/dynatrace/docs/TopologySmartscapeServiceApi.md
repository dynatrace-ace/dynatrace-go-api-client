# \TopologySmartscapeServiceApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaselineValuesForSingleService**](TopologySmartscapeServiceApi.md#GetBaselineValuesForSingleService) | **Get** /entity/services/{meIdentifier}/baseline | Gets baseline data for the specified service | maturity&#x3D;EARLY_ADOPTER
[**GetServices**](TopologySmartscapeServiceApi.md#GetServices) | **Get** /entity/services | Lists all available services in your environment
[**GetSingleService**](TopologySmartscapeServiceApi.md#GetSingleService) | **Get** /entity/services/{meIdentifier} | Gets parameters of the specified service
[**UpdateService**](TopologySmartscapeServiceApi.md#UpdateService) | **Post** /entity/services/{meIdentifier} | Updates parameters of the specified service



## GetBaselineValuesForSingleService

> ServiceBaselineValues GetBaselineValuesForSingleService(ctx, meIdentifier).Execute()

Gets baseline data for the specified service | maturity=EARLY_ADOPTER

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeServiceApi.GetBaselineValuesForSingleService(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeServiceApi.GetBaselineValuesForSingleService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBaselineValuesForSingleService`: ServiceBaselineValues
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeServiceApi.GetBaselineValuesForSingleService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaselineValuesForSingleServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceBaselineValues**](ServiceBaselineValues.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []Service GetServices(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all available services in your environment



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
    tag := []string{"Inner_example"} // []string | Filters the resulting set of services by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The service has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified services only.    To specify several services use the following format: `entity=ID1&entity=ID2`. (optional)
    managementZone := int64(789) // int64 | Only return services that are part of the specified management zone. (optional)
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    pageSize := int32(56) // int32 | The number of services per result page.    If not set, pagination is not used and the result contains all services fitting the specified filtering criteria. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you're using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeServiceApi.GetServices(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeServiceApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: []Service
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeServiceApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. | 
 **relativeTime** | **string** | The relative timeframe, back from now. | 
 **tag** | **[]string** | Filters the resulting set of services by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The service has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified services only.    To specify several services use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **managementZone** | **int64** | Only return services that are part of the specified management zone. | 
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **pageSize** | **int32** | The number of services per result page.    If not set, pagination is not used and the result contains all services fitting the specified filtering criteria. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**[]Service**](Service.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleService

> Service GetSingleService(ctx, meIdentifier).Execute()

Gets parameters of the specified service

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeServiceApi.GetSingleService(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeServiceApi.GetSingleService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleService`: Service
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeServiceApi.GetSingleService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> UpdateService(ctx, meIdentifier).UpdateEntity(updateEntity).Execute()

Updates parameters of the specified service

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the service you're inquiring.
    updateEntity := *openapiclient.NewUpdateEntity([]string{"Tags_example"}) // UpdateEntity |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeServiceApi.UpdateService(context.Background(), meIdentifier).UpdateEntity(updateEntity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeServiceApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the service you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


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

