# \MetricsApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllMetrics**](MetricsApi.md#AllMetrics) | **Get** /metrics | Lists all available metrics
[**Delete**](MetricsApi.md#Delete) | **Delete** /metrics/{metricId} | Deletes the specified metric
[**Ingest**](MetricsApi.md#Ingest) | **Post** /metrics/ingest | Pushes metric data points to Dynatrace
[**Metric**](MetricsApi.md#Metric) | **Get** /metrics/{metricId} | Gets the descriptor of the specified metric
[**Query**](MetricsApi.md#Query) | **Get** /metrics/query | Gets data points of the specified metrics



## AllMetrics

> MetricDescriptorCollection AllMetrics(ctx).NextPageKey(nextPageKey).PageSize(pageSize).MetricSelector(metricSelector).Text(text).Fields(fields).WrittenSince(writtenSince).MetadataSelector(metadataSelector).Execute()

Lists all available metrics



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
    pageSize := int64(789) // int64 | The amount of primary entities in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used.   If a value higher than 500 is used, only 500 results per page are returned. (optional)
    metricSelector := "metricSelector_example" // string | Selects metrics for the query by their keys.   You can specify multiple metric keys separated by commas (for example, `metrickey1,metrickey2`). To select multiple metrics belonging to the same parent, list the last part of the required metric keys in parentheses, separated by commas, while keeping the common part untouched. For example, to list the `builtin:host.cpu.idle` and `builtin:host.cpu.user` metric, write: `builtin:host.cpu.(idle,user)`.   You can select a full set of related metrics by using a trailing asterisk (`*`) wildcard. For example, `builtin:host.*` selects all host-based metrics and `builtin:*` selects all Dynatrace-provided metrics.   You can set additional transformation operators, separated by a colon (`:`). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   Only `aggregation`, `merge`, `parents`, and `splitBy` transformations are supported by this endpoint.   The length of the string is limited to 5,000 characters.   To find metrics based on a search term, rather than metricId, use the **text** query parameter instead of this one. (optional)
    text := "text_example" // string | Metric registry search term. Only show metrics that contain the term in their key, display name, or description. Use the `metricSelector` parameter instead of this one to select a complete metric hierarchy instead of doing a text-based search. (optional)
    fields := "fields_example" // string | Defines the list of metric properties included in the response.   `metricId` is **always** included in the result. The following additional properties are available:   * `displayName`: The name of the metric in the user interface. Enabled by default.  * `description`: A short description of the metric. Enabled by default.  * `unit`: The unit of the metric. Enabled by default.  * `tags`: The tags of the metric.  + `dduBillable`:  An indicator whether the usage of metric consumes [Davis data units](https://dt-url.net/ddu).  * `created`:  The timestamp (UTC milliseconds) when the metrics has been created.  * `lastWritten`:  The timestamp (UTC milliseconds) when metric data points have been written for the last time.  * `aggregationTypes`: The list of allowed aggregations for the metric. Note that it may be different after a [transformation](https://dt-url.net/metricSelector) is applied.  * `defaultAggregation`: The default aggregation of the metric. It is used when no aggregation is specified or the `:auto` transformation is set.  * `dimensionDefinitions`: The fine metric division (for example, process group and process ID for some process-related metric).  * `transformations`: A list of [transformations](https://dt-url.net/metricSelector) that can be applied to the metric. * `entityType`: A list of entity types supported by the metric.  To add properties, list them with leading plus `+`. To exclude default properties, list them with leading minus `-`.   To specify several properties, join them with a comma (for example `fields=+aggregationTypes,-description`).  If you specify just one property, the response contains the metric key and the specified property. To return metric keys only, specify `metricId` here. (optional)
    writtenSince := "writtenSince_example" // string | Filters the resulted set of metrics to those that have data points within the specified timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years    (optional)
    metadataSelector := "metadataSelector_example" // string | The metadata scope of the query. Only metrics with specified properties are included to the response.   You can set one or more of the following criteria. Values are case-sensitive and the `EQUALS` operator is used. If several values are specified, the **OR** logic applies.   * `unit(\"unit-1\",\"unit-2\")`  * `tags(\"tag-1\",\"tag-2\")`   To set several criteria, separate them with a comma (`,`). For example, `tags(\"feature\",\"cloud\"),unit(\"Percent\")`. Only results matching **all** criteria are included in response.   For example, to list metrics that have the tags **feature** AND **cloud** with a unit of **Percent** OR **MegaByte**, use this **metadataSelector**: `tags(\"feature\"),unit(\"Percent\",\"MegaByte\"),tags(\"cloud\")`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.AllMetrics(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).MetricSelector(metricSelector).Text(text).Fields(fields).WrittenSince(writtenSince).MetadataSelector(metadataSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.AllMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllMetrics`: MetricDescriptorCollection
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.AllMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of primary entities in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used.   If a value higher than 500 is used, only 500 results per page are returned. | 
 **metricSelector** | **string** | Selects metrics for the query by their keys.   You can specify multiple metric keys separated by commas (for example, &#x60;metrickey1,metrickey2&#x60;). To select multiple metrics belonging to the same parent, list the last part of the required metric keys in parentheses, separated by commas, while keeping the common part untouched. For example, to list the &#x60;builtin:host.cpu.idle&#x60; and &#x60;builtin:host.cpu.user&#x60; metric, write: &#x60;builtin:host.cpu.(idle,user)&#x60;.   You can select a full set of related metrics by using a trailing asterisk (&#x60;*&#x60;) wildcard. For example, &#x60;builtin:host.*&#x60; selects all host-based metrics and &#x60;builtin:*&#x60; selects all Dynatrace-provided metrics.   You can set additional transformation operators, separated by a colon (&#x60;:&#x60;). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   Only &#x60;aggregation&#x60;, &#x60;merge&#x60;, &#x60;parents&#x60;, and &#x60;splitBy&#x60; transformations are supported by this endpoint.   The length of the string is limited to 5,000 characters.   To find metrics based on a search term, rather than metricId, use the **text** query parameter instead of this one. | 
 **text** | **string** | Metric registry search term. Only show metrics that contain the term in their key, display name, or description. Use the &#x60;metricSelector&#x60; parameter instead of this one to select a complete metric hierarchy instead of doing a text-based search. | 
 **fields** | **string** | Defines the list of metric properties included in the response.   &#x60;metricId&#x60; is **always** included in the result. The following additional properties are available:   * &#x60;displayName&#x60;: The name of the metric in the user interface. Enabled by default.  * &#x60;description&#x60;: A short description of the metric. Enabled by default.  * &#x60;unit&#x60;: The unit of the metric. Enabled by default.  * &#x60;tags&#x60;: The tags of the metric.  + &#x60;dduBillable&#x60;:  An indicator whether the usage of metric consumes [Davis data units](https://dt-url.net/ddu).  * &#x60;created&#x60;:  The timestamp (UTC milliseconds) when the metrics has been created.  * &#x60;lastWritten&#x60;:  The timestamp (UTC milliseconds) when metric data points have been written for the last time.  * &#x60;aggregationTypes&#x60;: The list of allowed aggregations for the metric. Note that it may be different after a [transformation](https://dt-url.net/metricSelector) is applied.  * &#x60;defaultAggregation&#x60;: The default aggregation of the metric. It is used when no aggregation is specified or the &#x60;:auto&#x60; transformation is set.  * &#x60;dimensionDefinitions&#x60;: The fine metric division (for example, process group and process ID for some process-related metric).  * &#x60;transformations&#x60;: A list of [transformations](https://dt-url.net/metricSelector) that can be applied to the metric. * &#x60;entityType&#x60;: A list of entity types supported by the metric.  To add properties, list them with leading plus &#x60;+&#x60;. To exclude default properties, list them with leading minus &#x60;-&#x60;.   To specify several properties, join them with a comma (for example &#x60;fields&#x3D;+aggregationTypes,-description&#x60;).  If you specify just one property, the response contains the metric key and the specified property. To return metric keys only, specify &#x60;metricId&#x60; here. | 
 **writtenSince** | **string** | Filters the resulted set of metrics to those that have data points within the specified timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years    | 
 **metadataSelector** | **string** | The metadata scope of the query. Only metrics with specified properties are included to the response.   You can set one or more of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used. If several values are specified, the **OR** logic applies.   * &#x60;unit(\&quot;unit-1\&quot;,\&quot;unit-2\&quot;)&#x60;  * &#x60;tags(\&quot;tag-1\&quot;,\&quot;tag-2\&quot;)&#x60;   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;tags(\&quot;feature\&quot;,\&quot;cloud\&quot;),unit(\&quot;Percent\&quot;)&#x60;. Only results matching **all** criteria are included in response.   For example, to list metrics that have the tags **feature** AND **cloud** with a unit of **Percent** OR **MegaByte**, use this **metadataSelector**: &#x60;tags(\&quot;feature\&quot;),unit(\&quot;Percent\&quot;,\&quot;MegaByte\&quot;),tags(\&quot;cloud\&quot;)&#x60;. | 

### Return type

[**MetricDescriptorCollection**](MetricDescriptorCollection.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, text/csv; header=present; charset=utf-8, text/csv; header=absent; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, metricId).Execute()

Deletes the specified metric



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
    metricId := "metricId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.Delete(context.Background(), metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ingest

> Ingest(ctx).Body(body).Execute()

Pushes metric data points to Dynatrace

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
    body := "server.cpu.temperature,cpu.id=0 42" // string | Data points, provided in the [line protocol](https://dt-url.net/5d63ic1). Each line represents a single data point.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.Ingest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.Ingest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Data points, provided in the [line protocol](https://dt-url.net/5d63ic1). Each line represents a single data point. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: text/plain; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Metric

> MetricDescriptor Metric(ctx, metricId).Execute()

Gets the descriptor of the specified metric

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
    metricId := "metricId_example" // string | The key of the required metric.   You can set additional transformation operators, separated by a colon (`:`). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   The length of the string is limited to 5,000 characters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.Metric(context.Background(), metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.Metric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Metric`: MetricDescriptor
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.Metric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The key of the required metric.   You can set additional transformation operators, separated by a colon (&#x60;:&#x60;). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   The length of the string is limited to 5,000 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricDescriptor**](MetricDescriptor.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, text/csv; header=present; charset=utf-8, text/csv; header=absent; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> MetricData Query(ctx).MetricSelector(metricSelector).Resolution(resolution).From(from).To(to).EntitySelector(entitySelector).Execute()

Gets data points of the specified metrics



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
    metricSelector := "metricSelector_example" // string | Selects metrics for the query by their keys. You can select up to 10 metrics for one query.   You can specify multiple metric keys separated by commas (for example, `metrickey1,metrickey2`). To select multiple metrics belonging to the same parent, list the last part of the required metric keys in parentheses, separated by commas, while keeping the common part untouched. For example, to list the `builtin:host.cpu.idle` and `builtin:host.cpu.user` metric, write: `builtin:host.cpu.(idle,user)`.   You can set additional transformation operators, separated by a colon (`:`). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   If the metric key contains whitespaces, they must be escaped with a tilde (`~`). For example, to query the metric with the key of **ext:selfmonitoring.jmx.Agents of Type 'APACHE'** you must specify this selector:   `ext:selfmonitoring.jmx.Agents~ of~ Type~ 'APACHE'`   The length of the string is limited to 5,000 characters. (optional)
    resolution := "resolution_example" // string | The desired resolution of data points.   You can use one of the following options:  * One aggregated data point of each series. Set `Inf` to use this option.  * The desired amount of data points. This is the default option. This is a reference number of points, which is not necessarily equal to the number of the returned data points.  * The desired timespan between data points. This is a reference timespan, which is not necessarily equal to the returned timespan. To use this option, specify the unit of the timespan.   Valid units for the timespan are:  * `m`: minutes  * `h`: hours  * `d`: days  * `w`: weeks  * `M`: months  * `y`: years   If not set, the default is **120 data points**.  For example:   * Get data points which are 10 minutes apart: `resolution=10m`  * Get data points which are 3 weeks apart: `resolution=3w` (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    entitySelector := "entitySelector_example" // string | Specifies the entity scope of the query. Only data points delivered by matched entities are included in response.   You must set one of these criteria:   * Entity type: `type(\"TYPE\")`  * Dynatrace entity ID: `entityId(\"id\")`. You can specify several IDs, separated by a comma (`entityId(\"id-1\",\"id-2\")`).   You can add one or several of the following criteria. Values are case-sensitive and the `EQUALS` operator is used unless otherwise specified.   * Tag: `tag(\"value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. Any colons (`:`) that are part of the key or value must be escaped with a backslash(`\\`), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: `mzId(123)`  * Management zone name: `mzName(\"value\")` * Entity name: `entityName(\"value\")`. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * `entityName.equals`: changes the operator to `EQUALS`.   * `entityName.startsWith`: changes the operator to `BEGINS WITH`.   * `entityName.in`: enables you to provide multiple values. The `EQUALS` operator applies.   * `caseSensitive(entityName(\"value\"))`: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): `healthState(\"HEALTHY\")` * First seen timestamp: `firstSeenTms.<operator>(now-3h)`. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * `lte`: earlier than or at the specified time  * `lt`: earlier than the specified time  * `gte`: later than or at the specified time  * `gt`: later than the specified time * Entity attribute: `<attribute>(\"value1\",\"value2\")` and `<attribute>.exists()`. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: `fromRelationships.<relationshipName>()` and `toRelationships.<relationshipName>()`. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: `not(<criterion>)`. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (`,`). For example, `type(\"HOST\"),healthState(\"HEALTHY\")`. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   Use the [`GET /metrics/{metricId}`](https://dt-url.net/6z23ifk) call to fetch the list of possible entity types for your metric.   To set a universal scope matching all entities, omit this parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.Query(context.Background()).MetricSelector(metricSelector).Resolution(resolution).From(from).To(to).EntitySelector(entitySelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.Query``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query`: MetricData
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.Query`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSelector** | **string** | Selects metrics for the query by their keys. You can select up to 10 metrics for one query.   You can specify multiple metric keys separated by commas (for example, &#x60;metrickey1,metrickey2&#x60;). To select multiple metrics belonging to the same parent, list the last part of the required metric keys in parentheses, separated by commas, while keeping the common part untouched. For example, to list the &#x60;builtin:host.cpu.idle&#x60; and &#x60;builtin:host.cpu.user&#x60; metric, write: &#x60;builtin:host.cpu.(idle,user)&#x60;.   You can set additional transformation operators, separated by a colon (&#x60;:&#x60;). See [Metrics selector transformations](https://dt-url.net/metricSelector) in Dynatrace Documentation for additional information on available result transformations.   If the metric key contains whitespaces, they must be escaped with a tilde (&#x60;~&#x60;). For example, to query the metric with the key of **ext:selfmonitoring.jmx.Agents of Type &#39;APACHE&#39;** you must specify this selector:   &#x60;ext:selfmonitoring.jmx.Agents~ of~ Type~ &#39;APACHE&#39;&#x60;   The length of the string is limited to 5,000 characters. | 
 **resolution** | **string** | The desired resolution of data points.   You can use one of the following options:  * One aggregated data point of each series. Set &#x60;Inf&#x60; to use this option.  * The desired amount of data points. This is the default option. This is a reference number of points, which is not necessarily equal to the number of the returned data points.  * The desired timespan between data points. This is a reference timespan, which is not necessarily equal to the returned timespan. To use this option, specify the unit of the timespan.   Valid units for the timespan are:  * &#x60;m&#x60;: minutes  * &#x60;h&#x60;: hours  * &#x60;d&#x60;: days  * &#x60;w&#x60;: weeks  * &#x60;M&#x60;: months  * &#x60;y&#x60;: years   If not set, the default is **120 data points**.  For example:   * Get data points which are 10 minutes apart: &#x60;resolution&#x3D;10m&#x60;  * Get data points which are 3 weeks apart: &#x60;resolution&#x3D;3w&#x60; | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **entitySelector** | **string** | Specifies the entity scope of the query. Only data points delivered by matched entities are included in response.   You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;).   You can add one or several of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name: &#x60;entityName(\&quot;value\&quot;)&#x60;. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * &#x60;entityName.equals&#x60;: changes the operator to &#x60;EQUALS&#x60;.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   Use the [&#x60;GET /metrics/{metricId}&#x60;](https://dt-url.net/6z23ifk) call to fetch the list of possible entity types for your metric.   To set a universal scope matching all entities, omit this parameter. | 

### Return type

[**MetricData**](MetricData.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, text/csv; header=present; charset=utf-8, text/csv; header=absent; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

