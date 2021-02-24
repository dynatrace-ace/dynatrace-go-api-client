# \ServiceIBMMQTracingApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImsEntryQueue**](ServiceIBMMQTracingApi.md#CreateImsEntryQueue) | **Post** /service/ibmMQTracing/imsEntryQueue | Creates an IBM IMS entry queue
[**CreateOrUpdateImsEntryQueue**](ServiceIBMMQTracingApi.md#CreateOrUpdateImsEntryQueue) | **Put** /service/ibmMQTracing/imsEntryQueue/{id} | Updates an existing IBM IMS entry queue or creates a new one
[**DeleteImsEntryQueue**](ServiceIBMMQTracingApi.md#DeleteImsEntryQueue) | **Delete** /service/ibmMQTracing/imsEntryQueue/{id} | Deletes the specified IBM IMS entry queue
[**DeleteQueueManager**](ServiceIBMMQTracingApi.md#DeleteQueueManager) | **Delete** /service/ibmMQTracing/queueManager/{name} | Deletes the specified queue manager
[**GetAllImsEntryQueues**](ServiceIBMMQTracingApi.md#GetAllImsEntryQueues) | **Get** /service/ibmMQTracing/imsEntryQueue | Lists all IBM IMS entry queues
[**GetImsEntryQueue**](ServiceIBMMQTracingApi.md#GetImsEntryQueue) | **Get** /service/ibmMQTracing/imsEntryQueue/{id} | Gets the properties of the specified IBM IMS entry queue
[**GetQueueManager**](ServiceIBMMQTracingApi.md#GetQueueManager) | **Get** /service/ibmMQTracing/queueManager/{name} | Gets the parameters of the specified queue manager
[**GetQueueManagers**](ServiceIBMMQTracingApi.md#GetQueueManagers) | **Get** /service/ibmMQTracing/queueManager | Lists all available queue managers
[**PutQueueManager**](ServiceIBMMQTracingApi.md#PutQueueManager) | **Put** /service/ibmMQTracing/queueManager/{name} | Updates the specified queue manager or creates a new one
[**ValidateImsEntryQueueForPost**](ServiceIBMMQTracingApi.md#ValidateImsEntryQueueForPost) | **Post** /service/ibmMQTracing/imsEntryQueue/validator | Validates new IBM IMS entry queues for the &#x60;POST /service/ibmMQTracing/imsEntryQueue&#x60; request
[**ValidateImsEntryQueueForPut**](ServiceIBMMQTracingApi.md#ValidateImsEntryQueueForPut) | **Post** /service/ibmMQTracing/imsEntryQueue/{id}/validator | Validates update of existing IBM IMS entry queues for the &#x60;PUT /service/ibmMQTracing/imsEntryQueue/{id}&#x60; request
[**ValidateQueueManager**](ServiceIBMMQTracingApi.md#ValidateQueueManager) | **Post** /service/ibmMQTracing/queueManager/{name}/validator | Validates the queue manager update for the &#x60;PUT /service/ibmMQTracing/queueManager/{name}&#x60; request



## CreateImsEntryQueue

> EntityShortRepresentation CreateImsEntryQueue(ctx).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()

Creates an IBM IMS entry queue



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
    ibmMQImsEntryQueue := *openapiclient.NewIbmMQImsEntryQueue("QueueManagerName_example", "QueueName_example") // IbmMQImsEntryQueue | JSON body of the request, containing parameters of the new IBM IMS entry queue. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.CreateImsEntryQueue(context.Background()).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.CreateImsEntryQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImsEntryQueue`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.CreateImsEntryQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImsEntryQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibmMQImsEntryQueue** | [**IbmMQImsEntryQueue**](IbmMQImsEntryQueue.md) | JSON body of the request, containing parameters of the new IBM IMS entry queue. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateImsEntryQueue

> EntityShortRepresentation CreateOrUpdateImsEntryQueue(ctx, id).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()

Updates an existing IBM IMS entry queue or creates a new one



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
    id := TODO // string | The ID of the IBM IMS entry queue to update.   If you set the ID in the body as well, it must match this ID.
    ibmMQImsEntryQueue := *openapiclient.NewIbmMQImsEntryQueue("QueueManagerName_example", "QueueName_example") // IbmMQImsEntryQueue | JSON body of the request, containing updated parameters of the IBM IMS entry queue. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.CreateOrUpdateImsEntryQueue(context.Background(), id).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.CreateOrUpdateImsEntryQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateImsEntryQueue`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.CreateOrUpdateImsEntryQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the IBM IMS entry queue to update.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateImsEntryQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ibmMQImsEntryQueue** | [**IbmMQImsEntryQueue**](IbmMQImsEntryQueue.md) | JSON body of the request, containing updated parameters of the IBM IMS entry queue. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImsEntryQueue

> DeleteImsEntryQueue(ctx, id).Execute()

Deletes the specified IBM IMS entry queue

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
    id := TODO // string | The ID of the IBM IMS entry queue to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.DeleteImsEntryQueue(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.DeleteImsEntryQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the IBM IMS entry queue to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImsEntryQueueRequest struct via the builder pattern


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


## DeleteQueueManager

> DeleteQueueManager(ctx, name).Execute()

Deletes the specified queue manager

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
    name := "name_example" // string | The name of the queue manager to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.DeleteQueueManager(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.DeleteQueueManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the queue manager to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueManagerRequest struct via the builder pattern


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


## GetAllImsEntryQueues

> StubList GetAllImsEntryQueues(ctx).Execute()

Lists all IBM IMS entry queues



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.GetAllImsEntryQueues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.GetAllImsEntryQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllImsEntryQueues`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.GetAllImsEntryQueues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllImsEntryQueuesRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImsEntryQueue

> IbmMQImsEntryQueue GetImsEntryQueue(ctx, id).Execute()

Gets the properties of the specified IBM IMS entry queue

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
    id := TODO // string | The ID of the required IBM IMS entry queue.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.GetImsEntryQueue(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.GetImsEntryQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImsEntryQueue`: IbmMQImsEntryQueue
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.GetImsEntryQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required IBM IMS entry queue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImsEntryQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IbmMQImsEntryQueue**](IbmMQImsEntryQueue.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueManager

> QueueManager GetQueueManager(ctx, name).Execute()

Gets the parameters of the specified queue manager

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
    name := "name_example" // string | The name of the queue manager you're inquiring.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.GetQueueManager(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.GetQueueManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueManager`: QueueManager
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.GetQueueManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the queue manager you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueManager**](QueueManager.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueManagers

> StubList GetQueueManagers(ctx).Execute()

Lists all available queue managers



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.GetQueueManagers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.GetQueueManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueManagers`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.GetQueueManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueManagersRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutQueueManager

> EntityShortRepresentation PutQueueManager(ctx, name).QueueManager(queueManager).Execute()

Updates the specified queue manager or creates a new one



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
    name := "name_example" // string | The name of the queue manager to be updated.    The name of the queue manager in the body of the request must match this name.
    queueManager := *openapiclient.NewQueueManager("Name_example", []string{"Clusters_example"}, []openapiclient.AliasQueue{*openapiclient.NewAliasQueue("AliasQueue_example", "BaseQueue_example", []string{"ClusterVisibility_example"})}, []openapiclient.RemoteQueue{*openapiclient.NewRemoteQueue("LocalQueue_example", "RemoteQueue_example", "RemoteQueueManager_example", []string{"ClusterVisibility_example"})}, []openapiclient.ClusterQueue{*openapiclient.NewClusterQueue("LocalQueue_example", []string{"ClusterVisibility_example"})}) // QueueManager | The JSON body of the request containing updated parameters of the queue manager. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.PutQueueManager(context.Background(), name).QueueManager(queueManager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.PutQueueManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutQueueManager`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceIBMMQTracingApi.PutQueueManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the queue manager to be updated.    The name of the queue manager in the body of the request must match this name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutQueueManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queueManager** | [**QueueManager**](QueueManager.md) | The JSON body of the request containing updated parameters of the queue manager. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateImsEntryQueueForPost

> ValidateImsEntryQueueForPost(ctx).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()

Validates new IBM IMS entry queues for the `POST /service/ibmMQTracing/imsEntryQueue` request

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
    ibmMQImsEntryQueue := *openapiclient.NewIbmMQImsEntryQueue("QueueManagerName_example", "QueueName_example") // IbmMQImsEntryQueue | JSON body of the request, containing IBM IMS entry queue configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.ValidateImsEntryQueueForPost(context.Background()).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.ValidateImsEntryQueueForPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateImsEntryQueueForPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibmMQImsEntryQueue** | [**IbmMQImsEntryQueue**](IbmMQImsEntryQueue.md) | JSON body of the request, containing IBM IMS entry queue configuration to validate. | 

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


## ValidateImsEntryQueueForPut

> ValidateImsEntryQueueForPut(ctx, id).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()

Validates update of existing IBM IMS entry queues for the `PUT /service/ibmMQTracing/imsEntryQueue/{id}` request

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
    id := TODO // string | The ID of the IBM IMS entry queue to validate.
    ibmMQImsEntryQueue := *openapiclient.NewIbmMQImsEntryQueue("QueueManagerName_example", "QueueName_example") // IbmMQImsEntryQueue | JSON body of the request, containing IBM IMS entry queue configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.ValidateImsEntryQueueForPut(context.Background(), id).IbmMQImsEntryQueue(ibmMQImsEntryQueue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.ValidateImsEntryQueueForPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the IBM IMS entry queue to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateImsEntryQueueForPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ibmMQImsEntryQueue** | [**IbmMQImsEntryQueue**](IbmMQImsEntryQueue.md) | JSON body of the request, containing IBM IMS entry queue configuration to validate. | 

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


## ValidateQueueManager

> ValidateQueueManager(ctx, name).QueueManager(queueManager).Execute()

Validates the queue manager update for the `PUT /service/ibmMQTracing/queueManager/{name}` request

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
    name := "name_example" // string | The name of the queue manager to be validated.    The name of the queue manager in the body of the request must match this name.
    queueManager := *openapiclient.NewQueueManager("Name_example", []string{"Clusters_example"}, []openapiclient.AliasQueue{*openapiclient.NewAliasQueue("AliasQueue_example", "BaseQueue_example", []string{"ClusterVisibility_example"})}, []openapiclient.RemoteQueue{*openapiclient.NewRemoteQueue("LocalQueue_example", "RemoteQueue_example", "RemoteQueueManager_example", []string{"ClusterVisibility_example"})}, []openapiclient.ClusterQueue{*openapiclient.NewClusterQueue("LocalQueue_example", []string{"ClusterVisibility_example"})}) // QueueManager | The JSON body of the request containing updated parameters of the queue manager. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceIBMMQTracingApi.ValidateQueueManager(context.Background(), name).QueueManager(queueManager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceIBMMQTracingApi.ValidateQueueManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the queue manager to be validated.    The name of the queue manager in the body of the request must match this name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateQueueManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queueManager** | [**QueueManager**](QueueManager.md) | The JSON body of the request containing updated parameters of the queue manager. | 

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

