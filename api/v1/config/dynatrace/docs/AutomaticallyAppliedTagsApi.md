# \AutomaticallyAppliedTagsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoTag**](AutomaticallyAppliedTagsApi.md#CreateAutoTag) | **Post** /autoTags | Creates a new auto-tag
[**CreateOrUpdateAutoTag**](AutomaticallyAppliedTagsApi.md#CreateOrUpdateAutoTag) | **Put** /autoTags/{id} | Updates an existing auto-tag
[**DeleteAutoTag**](AutomaticallyAppliedTagsApi.md#DeleteAutoTag) | **Delete** /autoTags/{id} | Deletes the specified auto-tag
[**GetAllAutoTagConfigs**](AutomaticallyAppliedTagsApi.md#GetAllAutoTagConfigs) | **Get** /autoTags | Lists all configured auto-tags
[**GetSingleAutoTagConfig**](AutomaticallyAppliedTagsApi.md#GetSingleAutoTagConfig) | **Get** /autoTags/{id} | Gets the properties of the specified auto-tag
[**ValidateAutoTag**](AutomaticallyAppliedTagsApi.md#ValidateAutoTag) | **Post** /autoTags/validator | Validates the payload for the &#x60;POST /autoTags&#x60; request
[**ValidateAutoTag1**](AutomaticallyAppliedTagsApi.md#ValidateAutoTag1) | **Post** /autoTags/{id}/validator | Validates the payload for the &#x60;PUT /autoTags/{id}&#x60; request



## CreateAutoTag

> EntityShortRepresentation CreateAutoTag(ctx, optional)

Creates a new auto-tag

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAutoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAutoTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)| The JSON body of the request. Contains parameters of the new auto-tag. | 

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


## CreateOrUpdateAutoTag

> EntityShortRepresentation CreateOrUpdateAutoTag(ctx, id, optional)

Updates an existing auto-tag

If the auto-tag with the specified ID does not exist, a new auto-tag is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the auto-tag to be updated.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateAutoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateAutoTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)| The JSON body of the request. Contains updated parameters of the auto-tag. | 

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


## DeleteAutoTag

> DeleteAutoTag(ctx, id)

Deletes the specified auto-tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the auto-tag to be deleted. | 

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


## GetAllAutoTagConfigs

> StubList GetAllAutoTagConfigs(ctx, )

Lists all configured auto-tags

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


## GetSingleAutoTagConfig

> AutoTag GetSingleAutoTagConfig(ctx, id, optional)

Gets the properties of the specified auto-tag

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required auto-tag. | 
 **optional** | ***GetSingleAutoTagConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSingleAutoTagConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   When this flag is set to &#x60;false&#x60;, conditions with process group references are removed. If that results in a rule having no conditions, the entire rule is removed. | [default to false]

### Return type

[**AutoTag**](AutoTag.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAutoTag

> ValidateAutoTag(ctx, optional)

Validates the payload for the `POST /autoTags` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateAutoTagOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAutoTagOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)| The JSON body of the request. Contains auto-tag configuration to be validated. | 

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


## ValidateAutoTag1

> ValidateAutoTag1(ctx, id, optional)

Validates the payload for the `PUT /autoTags/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the auto-tag to be validated. | 
 **optional** | ***ValidateAutoTag1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAutoTag1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)| The JSON body of the request. Contains auto-tag configuration to be validated. | 

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

