# \ServiceCustomServicesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete1**](ServiceCustomServicesApi.md#Delete1) | **Delete** /service/customServices/{technology}/{id} | Deletes the specified custom service
[**GetItem**](ServiceCustomServicesApi.md#GetItem) | **Get** /service/customServices/{technology}/{id} | Gets the definition of the specified custom service
[**GetList**](ServiceCustomServicesApi.md#GetList) | **Get** /service/customServices/{technology} | Lists all custom services of the specified technology
[**Post**](ServiceCustomServicesApi.md#Post) | **Post** /service/customServices/{technology} | Creates a custom service
[**PutItem**](ServiceCustomServicesApi.md#PutItem) | **Put** /service/customServices/{technology}/{id} | Updates the specified custom service or create a new one.
[**ReorderList1**](ServiceCustomServicesApi.md#ReorderList1) | **Put** /service/customServices/{technology}/order | Reorders the custom services of the specified technology
[**ValidateItem**](ServiceCustomServicesApi.md#ValidateItem) | **Post** /service/customServices/{technology}/validator | Validate the new custom service for the &#x60;POST /customServices/{technology}&#x60; request
[**ValidateItem1**](ServiceCustomServicesApi.md#ValidateItem1) | **Post** /service/customServices/{technology}/{id}/validator | Validate the new custom service for the &#x60;PUT /customServices/{technology}/{id}&#x60; request



## Delete1

> Delete1(ctx, technology, id)

Deletes the specified custom service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the custom service to delete. | 
**id** | [**string**](.md)| The ID of the custom service to delete. | 

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


## GetItem

> CustomService GetItem(ctx, technology, id, optional)

Gets the definition of the specified custom service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the custom service you&#39;re inquiring. | 
**id** | [**string**](.md)| The ID of the custom service you&#39;re inquiring. | 
 **optional** | ***GetItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   &#x60;false&#x60; is used if the value is not set. | [default to false]

### Return type

[**CustomService**](CustomService.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList

> StubList GetList(ctx, technology)

Lists all custom services of the specified technology

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the required custom services. | 

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


## Post

> EntityShortRepresentation Post(ctx, technology, optional)

Creates a custom service

In the body of the request, neither custom service nor its rules can have the ID. All IDs will be generated automatically by Dynatrace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the new custom service. | 
 **optional** | ***PostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **optional.String**| Order of the new custom service. Set to &#x60;PREPEND&#x60; to prepend it to the list, &#x60;APPEND&#x60; to append it. Defaults to &#x60;APPEND&#x60;. | [default to APPEND]
 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing definition of the new custom service.   You must not specify the IDs for the custom service or any of its rules. The *order* field is not allowed either. | 

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


## PutItem

> EntityShortRepresentation PutItem(ctx, technology, id, optional)

Updates the specified custom service or create a new one.

Will use the config's ´order´ attribute if supplied, otherwise keeps the order of the existing config or appends if no existing config with the supplied ID was found.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the custom service to update. | 
**id** | [**string**](.md)| The ID of the custom service to update.   The ID of the custom service in the body of the request must match this ID. | 
 **optional** | ***PutItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing updated definition of the custom service. If *order* is present, it will be used. | 

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


## ReorderList1

> ReorderList1(ctx, technology, optional)

Reorders the custom services of the specified technology

This request reorders the custom services of the specified technology according to the given list of IDs. Custom services not present in the body of the request will retain their relative ordering but will be ordered *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of custom services to update. | 
 **optional** | ***ReorderList1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderList1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stubList** | [**optional.Interface of StubList**](StubList.md)| JSON body of the request containing the IDs of the custom services in the desired order. Any further properties (*name*, *description*) will be ignored. | 

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


## ValidateItem

> ValidateItem(ctx, technology, optional)

Validate the new custom service for the `POST /customServices/{technology}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the custom service to validate. | 
 **optional** | ***ValidateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing definition of the custom service to validate. | 

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


## ValidateItem1

> ValidateItem1(ctx, technology, id, optional)

Validate the new custom service for the `PUT /customServices/{technology}/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string**| Technology of the custom service to validate. | 
**id** | [**string**](.md)| The ID of the custom service to validate. | 
 **optional** | ***ValidateItem1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateItem1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing definition of the custom service to validate. | 

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

