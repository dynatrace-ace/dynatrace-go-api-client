# \ServiceRequestAttributesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration1**](ServiceRequestAttributesApi.md#CreateConfiguration1) | **Post** /service/requestAttributes | Creates a new request attribute
[**CreateOrUpdateConfiguration1**](ServiceRequestAttributesApi.md#CreateOrUpdateConfiguration1) | **Put** /service/requestAttributes/{id} | Updates an existing request attribute or creates a new one
[**DeleteConfiguration2**](ServiceRequestAttributesApi.md#DeleteConfiguration2) | **Delete** /service/requestAttributes/{id} | Deletes the specified request attribute
[**GetConfiguration6**](ServiceRequestAttributesApi.md#GetConfiguration6) | **Get** /service/requestAttributes/{id} | Shows the properties of the specified request attribute
[**ListConfigurations2**](ServiceRequestAttributesApi.md#ListConfigurations2) | **Get** /service/requestAttributes | Lists all available request attributes
[**ValidateConfiguration10**](ServiceRequestAttributesApi.md#ValidateConfiguration10) | **Post** /service/requestAttributes/{id}/validator | Validate updates of existing request attribute for the &#x60;PUT /requestAttributes/{id}&#x60; request
[**ValidateConfiguration9**](ServiceRequestAttributesApi.md#ValidateConfiguration9) | **Post** /service/requestAttributes/validator | Validates new request attributes for the &#x60;POST /requestAttributes&#x60; request



## CreateConfiguration1

> EntityShortRepresentation CreateConfiguration1(ctx, requestAttribute)

Creates a new request attribute

The body must not provide an ID as IDs are automatically assigned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestAttribute** | [**RequestAttribute**](RequestAttribute.md)|  | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateConfiguration1

> EntityShortRepresentation CreateOrUpdateConfiguration1(ctx, id, requestAttribute)

Updates an existing request attribute or creates a new one

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID. | 
**requestAttribute** | [**RequestAttribute**](RequestAttribute.md)|  | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration2

> DeleteConfiguration2(ctx, id)

Deletes the specified request attribute

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request attribute to be deleted. | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration6

> RequestAttribute GetConfiguration6(ctx, id, optional)

Shows the properties of the specified request attribute

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required request attribute. | 
 **optional** | ***GetConfiguration6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConfiguration6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process Group group references aren&#39;t compatible across environments. | [default to false]

### Return type

[**RequestAttribute**](RequestAttribute.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations2

> StubList ListConfigurations2(ctx, )

Lists all available request attributes

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


## ValidateConfiguration10

> ValidateConfiguration10(ctx, id, requestAttribute)

Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request attribute to be validated. | 
**requestAttribute** | [**RequestAttribute**](RequestAttribute.md)|  | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConfiguration9

> ValidateConfiguration9(ctx, requestAttribute)

Validates new request attributes for the `POST /requestAttributes` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestAttribute** | [**RequestAttribute**](RequestAttribute.md)|  | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

