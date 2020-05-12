# \ServiceIBMMQTracingApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

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

> EntityShortRepresentation CreateImsEntryQueue(ctx, optional)

Creates an IBM IMS entry queue

The body must not provide an ID as it will be automatically assigned by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateImsEntryQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateImsEntryQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibmMqImsEntryQueue** | [**optional.Interface of IbmMqImsEntryQueue**](IbmMqImsEntryQueue.md)| JSON body of the request, containing parameters of the new IBM IMS entry queue. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateImsEntryQueue

> EntityShortRepresentation CreateOrUpdateImsEntryQueue(ctx, id, optional)

Updates an existing IBM IMS entry queue or creates a new one

If the IBM IMS entry queue with the specified ID does not exist, a new IBM IMS entry queue will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the IBM IMS entry queue to update.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateImsEntryQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateImsEntryQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ibmMqImsEntryQueue** | [**optional.Interface of IbmMqImsEntryQueue**](IbmMqImsEntryQueue.md)| JSON body of the request, containing updated parameters of the IBM IMS entry queue. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImsEntryQueue

> DeleteImsEntryQueue(ctx, id)

Deletes the specified IBM IMS entry queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the IBM IMS entry queue to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueueManager

> DeleteQueueManager(ctx, name)

Deletes the specified queue manager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the queue manager to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllImsEntryQueues

> StubList GetAllImsEntryQueues(ctx, )

Lists all IBM IMS entry queues

This endpoint is used to provide Dynatrace with all IBM MQ queues (defined by QueueManagerName and QueueName) which are used to send messages to IBM IMS on the mainframe.    This is required to facilitate end to end tracing for messages sent via IBM MQ to IBM IMS.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StubList**](StubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImsEntryQueue

> IbmMqImsEntryQueue GetImsEntryQueue(ctx, id)

Gets the properties of the specified IBM IMS entry queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required IBM IMS entry queue. | 

### Return type

[**IbmMqImsEntryQueue**](IbmMQImsEntryQueue.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueManager

> QueueManager GetQueueManager(ctx, name)

Gets the parameters of the specified queue manager

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the queue manager you&#39;re inquiring. | 

### Return type

[**QueueManager**](QueueManager.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueManagers

> StubList GetQueueManagers(ctx, )

Lists all available queue managers

This endpoint is used to provide Dynatrace with your IBM MQ setup regarding alias, remote and cluster queues.    This is required to facilitate end to end tracing for messages send via IBM MQ where sender and receiver use different queue names. Without this information Dynatrace would still trace all requests, but would not be able to stitch service calls that use these IBM MQ features.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StubList**](StubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutQueueManager

> EntityShortRepresentation PutQueueManager(ctx, name, optional)

Updates the specified queue manager or creates a new one

If the queue manager with the specified ID doesn’t exist, a new queue manager will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the queue manager to be updated.    The name of the queue manager in the body of the request must match this name. | 
 **optional** | ***PutQueueManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutQueueManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queueManager** | [**optional.Interface of QueueManager**](QueueManager.md)| The JSON body of the request containing updated parameters of the queue manager. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateImsEntryQueueForPost

> ValidateImsEntryQueueForPost(ctx, optional)

Validates new IBM IMS entry queues for the `POST /service/ibmMQTracing/imsEntryQueue` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateImsEntryQueueForPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateImsEntryQueueForPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibmMqImsEntryQueue** | [**optional.Interface of IbmMqImsEntryQueue**](IbmMqImsEntryQueue.md)| JSON body of the request, containing IBM IMS entry queue configuration to validate. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateImsEntryQueueForPut

> ValidateImsEntryQueueForPut(ctx, id, optional)

Validates update of existing IBM IMS entry queues for the `PUT /service/ibmMQTracing/imsEntryQueue/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the IBM IMS entry queue to validate. | 
 **optional** | ***ValidateImsEntryQueueForPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateImsEntryQueueForPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ibmMqImsEntryQueue** | [**optional.Interface of IbmMqImsEntryQueue**](IbmMqImsEntryQueue.md)| JSON body of the request, containing IBM IMS entry queue configuration to validate. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateQueueManager

> ValidateQueueManager(ctx, name, optional)

Validates the queue manager update for the `PUT /service/ibmMQTracing/queueManager/{name}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the queue manager to be validated.    The name of the queue manager in the body of the request must match this name. | 
 **optional** | ***ValidateQueueManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateQueueManagerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queueManager** | [**optional.Interface of QueueManager**](QueueManager.md)| The JSON body of the request containing updated parameters of the queue manager. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

