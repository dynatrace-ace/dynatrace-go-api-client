# \TimeseriesApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomTimeseries**](TimeseriesApi.md#CreateCustomTimeseries) | **Put** /timeseries/{timeseriesIdentifier} | Creates a new custom metric
[**DeleteCustomTimeseries**](TimeseriesApi.md#DeleteCustomTimeseries) | **Delete** /timeseries/{timeseriesIdentifier} | Deletes the specified custom metric
[**GetAllTimeseriesDefinitions**](TimeseriesApi.md#GetAllTimeseriesDefinitions) | **Get** /timeseries | Lists all metric definitions, with the parameters of each metric
[**ReadTimeseriesComplex**](TimeseriesApi.md#ReadTimeseriesComplex) | **Post** /timeseries/{timeseriesIdentifier} | Lists all available metric data points, matching the specified parameters
[**ReadTimeseriesData**](TimeseriesApi.md#ReadTimeseriesData) | **Get** /timeseries/{timeseriesIdentifier} | Gets the parameters of the specified metric and, optionally, its data points



## CreateCustomTimeseries

> TimeseriesDefinition CreateCustomTimeseries(ctx, timeseriesIdentifier).TimeseriesRegistrationMessage(timeseriesRegistrationMessage).Execute()

Creates a new custom metric

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
    timeseriesIdentifier := "timeseriesIdentifier_example" // string | The ID for the new metric. It must start with the `custom:` prefix.   You can use alphanumeric characters and the following punctuation marks: periods (`.`), hyphens (`-`), and commas (`,`). A number cannot follow a punctuation mark.   If you use the ID of an existing metric the respective parameters will be updated.   The length of ID is limited to **512 characters**.
    timeseriesRegistrationMessage := *openapiclient.NewTimeseriesRegistrationMessage() // TimeseriesRegistrationMessage | The JSON body of the request. Contains parameters of the new custom metric. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeseriesApi.CreateCustomTimeseries(context.Background(), timeseriesIdentifier).TimeseriesRegistrationMessage(timeseriesRegistrationMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesApi.CreateCustomTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomTimeseries`: TimeseriesDefinition
    fmt.Fprintf(os.Stdout, "Response from `TimeseriesApi.CreateCustomTimeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeseriesIdentifier** | **string** | The ID for the new metric. It must start with the &#x60;custom:&#x60; prefix.   You can use alphanumeric characters and the following punctuation marks: periods (&#x60;.&#x60;), hyphens (&#x60;-&#x60;), and commas (&#x60;,&#x60;). A number cannot follow a punctuation mark.   If you use the ID of an existing metric the respective parameters will be updated.   The length of ID is limited to **512 characters**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeseriesRegistrationMessage** | [**TimeseriesRegistrationMessage**](TimeseriesRegistrationMessage.md) | The JSON body of the request. Contains parameters of the new custom metric. | 

### Return type

[**TimeseriesDefinition**](TimeseriesDefinition.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomTimeseries

> DeleteCustomTimeseries(ctx, timeseriesIdentifier).Execute()

Deletes the specified custom metric

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
    timeseriesIdentifier := "timeseriesIdentifier_example" // string | The ID of the metric to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeseriesApi.DeleteCustomTimeseries(context.Background(), timeseriesIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesApi.DeleteCustomTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeseriesIdentifier** | **string** | The ID of the metric to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomTimeseriesRequest struct via the builder pattern


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


## GetAllTimeseriesDefinitions

> []TimeseriesDefinition GetAllTimeseriesDefinitions(ctx).Source(source).DetailedSource(detailedSource).Execute()

Lists all metric definitions, with the parameters of each metric



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
    source := "source_example" // string | The type of the metric, such as BUILTIN or CUSTOM. (optional)
    detailedSource := "detailedSource_example" // string | The feature where metrics originates, such as Synthetic or RUM. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeseriesApi.GetAllTimeseriesDefinitions(context.Background()).Source(source).DetailedSource(detailedSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesApi.GetAllTimeseriesDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTimeseriesDefinitions`: []TimeseriesDefinition
    fmt.Fprintf(os.Stdout, "Response from `TimeseriesApi.GetAllTimeseriesDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTimeseriesDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | The type of the metric, such as BUILTIN or CUSTOM. | 
 **detailedSource** | **string** | The feature where metrics originates, such as Synthetic or RUM. | 

### Return type

[**[]TimeseriesDefinition**](TimeseriesDefinition.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTimeseriesComplex

> TimeseriesQueryResultWrapper ReadTimeseriesComplex(ctx, timeseriesIdentifier).TimeseriesQueryMessage(timeseriesQueryMessage).Execute()

Lists all available metric data points, matching the specified parameters



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
    timeseriesIdentifier := "timeseriesIdentifier_example" // string | The case-sensitive ID of the timeseries from which you want to read parameters and data points.
    timeseriesQueryMessage := *openapiclient.NewTimeseriesQueryMessage() // TimeseriesQueryMessage | The JSON body of the request, containing parameters to identify the required data points. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeseriesApi.ReadTimeseriesComplex(context.Background(), timeseriesIdentifier).TimeseriesQueryMessage(timeseriesQueryMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesApi.ReadTimeseriesComplex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTimeseriesComplex`: TimeseriesQueryResultWrapper
    fmt.Fprintf(os.Stdout, "Response from `TimeseriesApi.ReadTimeseriesComplex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeseriesIdentifier** | **string** | The case-sensitive ID of the timeseries from which you want to read parameters and data points. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTimeseriesComplexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeseriesQueryMessage** | [**TimeseriesQueryMessage**](TimeseriesQueryMessage.md) | The JSON body of the request, containing parameters to identify the required data points. | 

### Return type

[**TimeseriesQueryResultWrapper**](TimeseriesQueryResultWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTimeseriesData

> TimeseriesQueryResult ReadTimeseriesData(ctx, timeseriesIdentifier).IncludeData(includeData).AggregationType(aggregationType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Predict(predict).RelativeTime(relativeTime).QueryMode(queryMode).Entity(entity).Tag(tag).Percentile(percentile).IncludeParentIds(includeParentIds).ConsiderMaintenanceWindowsForAvailability(considerMaintenanceWindowsForAvailability).Execute()

Gets the parameters of the specified metric and, optionally, its data points



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
    timeseriesIdentifier := "timeseriesIdentifier_example" // string | The case-sensitive ID of the timeseries, from which you want to read parameters and data points.
    includeData := true // bool | The flag to include data points in the response. Set to `true` to obtain data points. You must also specify the timeframe and aggregation type.   Defaults to `false`, meaning data points are not included. (optional)
    aggregationType := "aggregationType_example" // string | The aggregation type for the resulting data points.   If the requested metric doesn't support the specified aggregation, the request will result in an error. (optional)
    startTimestamp := int64(789) // int64 | The start timestamp of the requested timeframe, in UTC milliseconds. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the requested timeframe, in milliseconds (UTC).   If later than the current time, Dynatrace automatically uses the current time instead.   The timeframe must not exceed 6 months. (optional)
    predict := true // bool | The flag to predict future data points. (optional)
    relativeTime := "relativeTime_example" // string | The relative timeframe, back from the current time. (optional)
    queryMode := "queryMode_example" // string | The type of result that the call should return. Valid result modes are:   `series`: returns all the data points of the timeseries in specified timeframe.  `total`: returns one scalar value for the specified timeframe.   By default, the `series` mode is used. (optional)
    entity := []string{"Inner_example"} // []string | Filters requested data points by the entity that should deliver them. Allowed values are Dynatrace entity IDs.   You can specify several entities in the following format: `entity=entity1&entity=entity2`. The entity has to match at least one of the specified IDs.   If the selected entity doesn't support the requested timeseries, the request will result in an error. (optional)
    tag := []string{"Inner_example"} // []string | Filters requested data points by entity which should deliver them. Only data from entities with the specified tag is delivered.   You can specify several tags in the following format: `tag=tag1&tag=tag2`. The entity has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use the `key:value` format. If there's also a context, use the `[context]key:value` format. (optional)
    percentile := int32(56) // int32 | Specifies which percentile of the selected response time metric should be delivered.  Applicable only to the `PERCENTILE` aggregation type.   Valid values for percentile are between 1 and 99.   Keep in mind that percentile export is possible only for response-time-based metrics such as application and service response times. (optional)
    includeParentIds := true // bool | If set to true the result exposes dimension mappings between parent entities and their children.  For instance: SERVICE-0000000000000001, SERVICE_METHOD-0000000000000001 (optional)
    considerMaintenanceWindowsForAvailability := true // bool | Exclude (`true`) or include (`false`) data points from any [maintenance window](https://dt-url.net/b2123rg0), defined in your environment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TimeseriesApi.ReadTimeseriesData(context.Background(), timeseriesIdentifier).IncludeData(includeData).AggregationType(aggregationType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Predict(predict).RelativeTime(relativeTime).QueryMode(queryMode).Entity(entity).Tag(tag).Percentile(percentile).IncludeParentIds(includeParentIds).ConsiderMaintenanceWindowsForAvailability(considerMaintenanceWindowsForAvailability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeseriesApi.ReadTimeseriesData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTimeseriesData`: TimeseriesQueryResult
    fmt.Fprintf(os.Stdout, "Response from `TimeseriesApi.ReadTimeseriesData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeseriesIdentifier** | **string** | The case-sensitive ID of the timeseries, from which you want to read parameters and data points. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTimeseriesDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeData** | **bool** | The flag to include data points in the response. Set to &#x60;true&#x60; to obtain data points. You must also specify the timeframe and aggregation type.   Defaults to &#x60;false&#x60;, meaning data points are not included. | 
 **aggregationType** | **string** | The aggregation type for the resulting data points.   If the requested metric doesn&#39;t support the specified aggregation, the request will result in an error. | 
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in UTC milliseconds. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If later than the current time, Dynatrace automatically uses the current time instead.   The timeframe must not exceed 6 months. | 
 **predict** | **bool** | The flag to predict future data points. | 
 **relativeTime** | **string** | The relative timeframe, back from the current time. | 
 **queryMode** | **string** | The type of result that the call should return. Valid result modes are:   &#x60;series&#x60;: returns all the data points of the timeseries in specified timeframe.  &#x60;total&#x60;: returns one scalar value for the specified timeframe.   By default, the &#x60;series&#x60; mode is used. | 
 **entity** | **[]string** | Filters requested data points by the entity that should deliver them. Allowed values are Dynatrace entity IDs.   You can specify several entities in the following format: &#x60;entity&#x3D;entity1&amp;entity&#x3D;entity2&#x60;. The entity has to match at least one of the specified IDs.   If the selected entity doesn&#39;t support the requested timeseries, the request will result in an error. | 
 **tag** | **[]string** | Filters requested data points by entity which should deliver them. Only data from entities with the specified tag is delivered.   You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The entity has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use the &#x60;key:value&#x60; format. If there&#39;s also a context, use the &#x60;[context]key:value&#x60; format. | 
 **percentile** | **int32** | Specifies which percentile of the selected response time metric should be delivered.  Applicable only to the &#x60;PERCENTILE&#x60; aggregation type.   Valid values for percentile are between 1 and 99.   Keep in mind that percentile export is possible only for response-time-based metrics such as application and service response times. | 
 **includeParentIds** | **bool** | If set to true the result exposes dimension mappings between parent entities and their children.  For instance: SERVICE-0000000000000001, SERVICE_METHOD-0000000000000001 | 
 **considerMaintenanceWindowsForAvailability** | **bool** | Exclude (&#x60;true&#x60;) or include (&#x60;false&#x60;) data points from any [maintenance window](https://dt-url.net/b2123rg0), defined in your environment. | 

### Return type

[**TimeseriesQueryResult**](TimeseriesQueryResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

