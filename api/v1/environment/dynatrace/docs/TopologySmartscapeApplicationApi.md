# \TopologySmartscapeApplicationApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplications**](TopologySmartscapeApplicationApi.md#GetApplications) | **Get** /entity/applications | Gets the list of all applications in your environment along with their parameters
[**GetBaselineValuesForSingleApplication**](TopologySmartscapeApplicationApi.md#GetBaselineValuesForSingleApplication) | **Get** /entity/applications/{meIdentifier}/baseline | Gets baseline data for the specified application | maturity&#x3D;EARLY_ADOPTER
[**GetSingleApplication**](TopologySmartscapeApplicationApi.md#GetSingleApplication) | **Get** /entity/applications/{meIdentifier} | Gets parameters of the specified application
[**UpdateApplication**](TopologySmartscapeApplicationApi.md#UpdateApplication) | **Post** /entity/applications/{meIdentifier} | Updates parameters of the specified application



## GetApplications

> []Application GetApplications(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Gets the list of all applications in your environment along with their parameters



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
    tag := []string{"Inner_example"} // []string | Filters the resulting set of applications by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The application has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified applications only.    To specify several applications use the following format: `entity=ID1&entity=ID2`. (optional)
    managementZone := int64(789) // int64 | Only return applications that are part of the specified management zone. (optional)
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    pageSize := int32(56) // int32 | The number of applications per result page.    If not set, pagination is not used and the result contains all applications fitting the specified filtering criteria. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you're using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeApplicationApi.GetApplications(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZone(managementZone).IncludeDetails(includeDetails).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeApplicationApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: []Application
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeApplicationApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days. | 
 **relativeTime** | **string** | The relative timeframe, back from now. | 
 **tag** | **[]string** | Filters the resulting set of applications by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The application has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified applications only.    To specify several applications use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **managementZone** | **int64** | Only return applications that are part of the specified management zone. | 
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **pageSize** | **int32** | The number of applications per result page.    If not set, pagination is not used and the result contains all applications fitting the specified filtering criteria. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages. | 

### Return type

[**[]Application**](Application.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaselineValuesForSingleApplication

> ApplicationBaselineValues GetBaselineValuesForSingleApplication(ctx, meIdentifier).Execute()

Gets baseline data for the specified application | maturity=EARLY_ADOPTER

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeApplicationApi.GetBaselineValuesForSingleApplication(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeApplicationApi.GetBaselineValuesForSingleApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBaselineValuesForSingleApplication`: ApplicationBaselineValues
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeApplicationApi.GetBaselineValuesForSingleApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaselineValuesForSingleApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationBaselineValues**](ApplicationBaselineValues.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleApplication

> Application GetSingleApplication(ctx, meIdentifier).Execute()

Gets parameters of the specified application

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the required application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeApplicationApi.GetSingleApplication(context.Background(), meIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeApplicationApi.GetSingleApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `TopologySmartscapeApplicationApi.GetSingleApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> UpdateApplication(ctx, meIdentifier).UpdateEntity(updateEntity).Execute()

Updates parameters of the specified application

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
    meIdentifier := "meIdentifier_example" // string | The Dynatrace entity ID of the application you want to update.
    updateEntity := *openapiclient.NewUpdateEntity([]string{"Tags_example"}) // UpdateEntity |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopologySmartscapeApplicationApi.UpdateApplication(context.Background(), meIdentifier).UpdateEntity(updateEntity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologySmartscapeApplicationApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meIdentifier** | **string** | The Dynatrace entity ID of the application you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


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

