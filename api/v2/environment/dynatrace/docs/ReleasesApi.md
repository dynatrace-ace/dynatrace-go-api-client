# \ReleasesApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleases**](ReleasesApi.md#GetReleases) | **Get** /releases | Returns all releases | maturity&#x3D;EARLY_ADOPTER



## GetReleases

> Releases GetReleases(ctx).From(from).To(to).NextPageKey(nextPageKey).PageSize(pageSize).Demo(demo).ReleasesSelector(releasesSelector).Sort(sort).Execute()

Returns all releases | maturity=EARLY_ADOPTER

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
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of Releases in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. (optional)
    demo := true // bool | Get your Releases (`false`) or a set of demo Releases (`true`). (optional) (default to false)
    releasesSelector := "releasesSelector_example" // string | Defines the scope of the query. Only Releases matching the provided criteria are included in the response.  You can add one or several of the criteria listed below. * Monitoring state: monitoringState(\"Active\") or monitoringState(\"Inactive\"). You can specify only one monitoring state.  * Health state: healthState(\"HEALTHY\") or healthState(\"UNHEALTHY\"). You can specify only one health state. * Name: entityName(\"name\"). Filters for all releases that contain the given value in their name. The filter is case-insensitive. * Entity ID: entityId(\"id\"). * Product: releasesProduct(\"product\"). Filters for all releases that contain the given value in their product. The filter is case-insensitive. * Stage: releasesStage(\"stage\"). Filters for all releases that contain the given value in their stage. The filter is case-insensitive. * Version: releasesVersion(\"version\"). Filters for all releases that contain the given value in their version. The filter is case-insensitive.  To set several criteria, separate them with comma (,). Only results matching all criteria are included in the response. e.g., .../api/v2/releases?releasesSelector=name(\"Server\"),monitoringState(\"Active\"),healthState(\"HEALTHY\"),releasesVersion(\"1.0.7\"). (optional) (default to "")
    sort := "sort_example" // string | Specifies the field that is used for sorting the releases list. The field has a sign prefix (+/-) which corresponds to the sorting order ('+' for ascending and '-' for descending). If no sign prefix is set, then the default ascending sorting order will be applied. You can sort by the following properties: * 'product': Product name * 'name': Release name * 'stage': Stage name * 'version': Version * 'instances': Instances * 'traffic': Traffic  If not set, the ascending order sorting for name is applied. (optional) (default to "name")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleasesApi.GetReleases(context.Background()).From(from).To(to).NextPageKey(nextPageKey).PageSize(pageSize).Demo(demo).ReleasesSelector(releasesSelector).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.GetReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleases`: Releases
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.GetReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of Releases in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. | 
 **demo** | **bool** | Get your Releases (&#x60;false&#x60;) or a set of demo Releases (&#x60;true&#x60;). | [default to false]
 **releasesSelector** | **string** | Defines the scope of the query. Only Releases matching the provided criteria are included in the response.  You can add one or several of the criteria listed below. * Monitoring state: monitoringState(\&quot;Active\&quot;) or monitoringState(\&quot;Inactive\&quot;). You can specify only one monitoring state.  * Health state: healthState(\&quot;HEALTHY\&quot;) or healthState(\&quot;UNHEALTHY\&quot;). You can specify only one health state. * Name: entityName(\&quot;name\&quot;). Filters for all releases that contain the given value in their name. The filter is case-insensitive. * Entity ID: entityId(\&quot;id\&quot;). * Product: releasesProduct(\&quot;product\&quot;). Filters for all releases that contain the given value in their product. The filter is case-insensitive. * Stage: releasesStage(\&quot;stage\&quot;). Filters for all releases that contain the given value in their stage. The filter is case-insensitive. * Version: releasesVersion(\&quot;version\&quot;). Filters for all releases that contain the given value in their version. The filter is case-insensitive.  To set several criteria, separate them with comma (,). Only results matching all criteria are included in the response. e.g., .../api/v2/releases?releasesSelector&#x3D;name(\&quot;Server\&quot;),monitoringState(\&quot;Active\&quot;),healthState(\&quot;HEALTHY\&quot;),releasesVersion(\&quot;1.0.7\&quot;). | [default to &quot;&quot;]
 **sort** | **string** | Specifies the field that is used for sorting the releases list. The field has a sign prefix (+/-) which corresponds to the sorting order (&#39;+&#39; for ascending and &#39;-&#39; for descending). If no sign prefix is set, then the default ascending sorting order will be applied. You can sort by the following properties: * &#39;product&#39;: Product name * &#39;name&#39;: Release name * &#39;stage&#39;: Stage name * &#39;version&#39;: Version * &#39;instances&#39;: Instances * &#39;traffic&#39;: Traffic  If not set, the ascending order sorting for name is applied. | [default to &quot;name&quot;]

### Return type

[**Releases**](Releases.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

