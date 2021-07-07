# \ServiceLevelObjectivesApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlo**](ServiceLevelObjectivesApi.md#CreateSlo) | **Post** /slo | Creates a new SLO
[**DeleteSlo**](ServiceLevelObjectivesApi.md#DeleteSlo) | **Delete** /slo/{id} | Deletes an SLO
[**GetSlo**](ServiceLevelObjectivesApi.md#GetSlo) | **Get** /slo | Lists all available SLOs along with calculated values
[**GetSloById**](ServiceLevelObjectivesApi.md#GetSloById) | **Get** /slo/{id} | Gets parameters and the calculated value of an SLO
[**UpdateSloById**](ServiceLevelObjectivesApi.md#UpdateSloById) | **Put** /slo/{id} | Updates an existing SLO



## CreateSlo

> CreateSlo(ctx).SloCreate(sloCreate).Execute()

Creates a new SLO

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
    sloCreate := *openapiclient.NewSloCreate() // SloCreate | The JSON body of the request. Contains the parameters of the new SLO.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.CreateSlo(context.Background()).SloCreate(sloCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sloCreate** | [**SloCreate**](SloCreate.md) | The JSON body of the request. Contains the parameters of the new SLO. | 

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


## DeleteSlo

> DeleteSlo(ctx, id).Execute()

Deletes an SLO

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
    id := TODO // string | The ID of the required SLO.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.DeleteSlo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSloRequest struct via the builder pattern


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


## GetSlo

> SLOs GetSlo(ctx).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).SloSelector(sloSelector).Sort(sort).TimeFrame(timeFrame).PageIdx(pageIdx).Demo(demo).Evaluate(evaluate).Execute()

Lists all available SLOs along with calculated values



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
    pageSize := int64(789) // int64 | The amount of SLOs in a single response payload.   The maximal allowed page size is 10000.   If not set, 10 is used. (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    sloSelector := "sloSelector_example" // string | The scope of the query. Only SLOs matching the provided criteria are included in the response.   You can add one or several of the criteria listed below. * ID: id(\"id\"). Filters for an SLO with the given ID. * Name: name(\"name\"). Filters for an SLO with the given name. The filter is case-sensitive. * Health state: healthState(\"HEALTHY\") or healthState(\"UNHEALTHY\"). You can specify only one health state. * Text: text(\"value\"). Filters for all SLOs that contain the given value in their name or description. The filter is case-insensitive.  To set several criteria, separate them with comma (,). Only results matching all criteria are included in the response. e.g., .../api/v2/slo?sloSelector=name(\"Service Availability\"), .../api/v2/slo?sloSelector=id(\"id\"), .../api/v2/slo?sloSelector=text(\"Description\"),healthState(\"HEALTHY\"). (optional) (default to "")
    sort := "sort_example" // string | The sorting of SLO entries:  * `name`: Names in ascending order.  * `-name`: Names in descending order.   If not set, the ascending order is used. (optional) (default to "name")
    timeFrame := "timeFrame_example" // string | The timeframe to calculate the SLO values:   * `CURRENT`: SLO's own timeframe.  * `GTF`: timeframe specified by **from** and **to** parameters.   If not set, the `CURRENT` value is used. (optional) (default to "CURRENT")
    pageIdx := int32(56) // int32 | Only SLOs on the given page are included in the response. The first page has the index '1'. (optional) (default to 1)
    demo := true // bool | Get your SLOs (`false`) or a set of demo SLOs (`true`). (optional) (default to false)
    evaluate := true // bool | Get your SLOs without them being evaluated (`false`) or with evaluations (`true`). (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSlo(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).SloSelector(sloSelector).Sort(sort).TimeFrame(timeFrame).PageIdx(pageIdx).Demo(demo).Evaluate(evaluate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlo`: SLOs
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSlo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of SLOs in a single response payload.   The maximal allowed page size is 10000.   If not set, 10 is used. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **sloSelector** | **string** | The scope of the query. Only SLOs matching the provided criteria are included in the response.   You can add one or several of the criteria listed below. * ID: id(\&quot;id\&quot;). Filters for an SLO with the given ID. * Name: name(\&quot;name\&quot;). Filters for an SLO with the given name. The filter is case-sensitive. * Health state: healthState(\&quot;HEALTHY\&quot;) or healthState(\&quot;UNHEALTHY\&quot;). You can specify only one health state. * Text: text(\&quot;value\&quot;). Filters for all SLOs that contain the given value in their name or description. The filter is case-insensitive.  To set several criteria, separate them with comma (,). Only results matching all criteria are included in the response. e.g., .../api/v2/slo?sloSelector&#x3D;name(\&quot;Service Availability\&quot;), .../api/v2/slo?sloSelector&#x3D;id(\&quot;id\&quot;), .../api/v2/slo?sloSelector&#x3D;text(\&quot;Description\&quot;),healthState(\&quot;HEALTHY\&quot;). | [default to &quot;&quot;]
 **sort** | **string** | The sorting of SLO entries:  * &#x60;name&#x60;: Names in ascending order.  * &#x60;-name&#x60;: Names in descending order.   If not set, the ascending order is used. | [default to &quot;name&quot;]
 **timeFrame** | **string** | The timeframe to calculate the SLO values:   * &#x60;CURRENT&#x60;: SLO&#39;s own timeframe.  * &#x60;GTF&#x60;: timeframe specified by **from** and **to** parameters.   If not set, the &#x60;CURRENT&#x60; value is used. | [default to &quot;CURRENT&quot;]
 **pageIdx** | **int32** | Only SLOs on the given page are included in the response. The first page has the index &#39;1&#39;. | [default to 1]
 **demo** | **bool** | Get your SLOs (&#x60;false&#x60;) or a set of demo SLOs (&#x60;true&#x60;). | [default to false]
 **evaluate** | **bool** | Get your SLOs without them being evaluated (&#x60;false&#x60;) or with evaluations (&#x60;true&#x60;). | [default to false]

### Return type

[**SLOs**](SLOs.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSloById

> SLO GetSloById(ctx, id).From(from).To(to).TimeFrame(timeFrame).Execute()

Gets parameters and the calculated value of an SLO



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
    id := TODO // string | The ID of the required SLO.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    timeFrame := "timeFrame_example" // string | The timeframe to calculate the SLO values:   * `CURRENT`: SLO's own timeframe.  * `GTF`: timeframe specified by **from** and **to** parameters.   If not set, the `CURRENT` value is used. (optional) (default to "CURRENT")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSloById(context.Background(), id).From(from).To(to).TimeFrame(timeFrame).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSloById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSloById`: SLO
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSloById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSloByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **timeFrame** | **string** | The timeframe to calculate the SLO values:   * &#x60;CURRENT&#x60;: SLO&#39;s own timeframe.  * &#x60;GTF&#x60;: timeframe specified by **from** and **to** parameters.   If not set, the &#x60;CURRENT&#x60; value is used. | [default to &quot;CURRENT&quot;]

### Return type

[**SLO**](SLO.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSloById

> UpdateSloById(ctx, id).SloCreate(sloCreate).Execute()

Updates an existing SLO

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
    id := TODO // string | The ID of the required SLO.
    sloCreate := *openapiclient.NewSloCreate() // SloCreate | The JSON body of the request. Contains the updated parameters of the SLO.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.UpdateSloById(context.Background(), id).SloCreate(sloCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.UpdateSloById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSloByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sloCreate** | [**SloCreate**](SloCreate.md) | The JSON body of the request. Contains the updated parameters of the SLO. | 

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

