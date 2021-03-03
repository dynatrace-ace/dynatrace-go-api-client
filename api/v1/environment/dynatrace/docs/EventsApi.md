# \EventsApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventById**](EventsApi.md#GetEventById) | **Get** /events/{eventId} | Gets the parameters of the specified event
[**PostEvent**](EventsApi.md#PostEvent) | **Post** /events | Pushes custom events to one or more monitored entities
[**QueryEvents**](EventsApi.md#QueryEvents) | **Get** /events | Lists all the events of your environment, along with their parameters



## GetEventById

> EventRestEntry GetEventById(ctx, eventId).Execute()

Gets the parameters of the specified event

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
    eventId := "eventId_example" // string | The ID of the required event.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.GetEventById(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventById`: EventRestEntry
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The ID of the required event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventRestEntry**](EventRestEntry.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEvent

> EventStoreResult PostEvent(ctx).EventCreation(eventCreation).Execute()

Pushes custom events to one or more monitored entities



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
    eventCreation := *openapiclient.NewEventCreation("EventType_example", *openapiclient.NewPushEventAttachRules(), "Source_example") // EventCreation | The JSON body of the request, containing parameters of the event. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.PostEvent(context.Background()).EventCreation(eventCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.PostEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostEvent`: EventStoreResult
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.PostEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventCreation** | [**EventCreation**](EventCreation.md) | The JSON body of the request, containing parameters of the event. | 

### Return type

[**EventStoreResult**](EventStoreResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryEvents

> EventQueryResult QueryEvents(ctx).From(from).To(to).RelativeTime(relativeTime).EventType(eventType).EntityId(entityId).Cursor(cursor).Execute()

Lists all the events of your environment, along with their parameters



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
    from := int64(789) // int64 | Start timestamp of the query, in UTC milliseconds.    If not set, 30 days ago from now is used, unless **relativeTime** is set. (optional)
    to := int64(789) // int64 | End timestamp of the query, in UTC milliseconds.    If not set, the current timestamp is used, unless **relativeTime** is set.   The timeframe must not exceed 2 years. (optional)
    relativeTime := "relativeTime_example" // string | Relative timeframe, back from the current time. (optional)
    eventType := "eventType_example" // string | Filters the resulting set of events by the event type. (optional)
    entityId := "entityId_example" // string | Filters the resulting set of events to the events, related to the specified Dynatrace entity. (optional)
    cursor := "cursor_example" // string | The response is limited to 150 events, so if you want to receive more you can use the cursor to get the next 150. This parameter should be empty when querying for the first time.    The cursor key can then be found in the **nextCursor** field of the previous response.   When using the cursor string, it is not necessary to specify the additional parameters, as that info is already encoded within the cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.QueryEvents(context.Background()).From(from).To(to).RelativeTime(relativeTime).EventType(eventType).EntityId(entityId).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.QueryEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryEvents`: EventQueryResult
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.QueryEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start timestamp of the query, in UTC milliseconds.    If not set, 30 days ago from now is used, unless **relativeTime** is set. | 
 **to** | **int64** | End timestamp of the query, in UTC milliseconds.    If not set, the current timestamp is used, unless **relativeTime** is set.   The timeframe must not exceed 2 years. | 
 **relativeTime** | **string** | Relative timeframe, back from the current time. | 
 **eventType** | **string** | Filters the resulting set of events by the event type. | 
 **entityId** | **string** | Filters the resulting set of events to the events, related to the specified Dynatrace entity. | 
 **cursor** | **string** | The response is limited to 150 events, so if you want to receive more you can use the cursor to get the next 150. This parameter should be empty when querying for the first time.    The cursor key can then be found in the **nextCursor** field of the previous response.   When using the cursor string, it is not necessary to specify the additional parameters, as that info is already encoded within the cursor. | 

### Return type

[**EventQueryResult**](EventQueryResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

