# \ServiceRequestNamingApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete6**](ServiceRequestNamingApi.md#Delete6) | **Delete** /service/requestNaming/{id} | Deletes the specified request naming rule
[**Get**](ServiceRequestNamingApi.md#Get) | **Get** /service/resourceNaming | Lists the global service resource requests.
[**GetItem5**](ServiceRequestNamingApi.md#GetItem5) | **Get** /service/requestNaming/{id} | Gets the parameters of the specified request naming rule
[**GetList5**](ServiceRequestNamingApi.md#GetList5) | **Get** /service/requestNaming | Lists all request naming rules along with their parameters
[**Post2**](ServiceRequestNamingApi.md#Post2) | **Post** /service/requestNaming | Creates a new request naming rule
[**Put1**](ServiceRequestNamingApi.md#Put1) | **Put** /service/requestNaming/{id} | Updates the specified request naming rule
[**Put2**](ServiceRequestNamingApi.md#Put2) | **Put** /service/resourceNaming | Updates the global service resource requests.
[**ReorderList2**](ServiceRequestNamingApi.md#ReorderList2) | **Put** /service/requestNaming/order | Reorders the request namings
[**Validate**](ServiceRequestNamingApi.md#Validate) | **Post** /service/resourceNaming/validator | Validates new resource requests settings for the &#x60;PUT /service/resourceRequest&#x60; request.
[**ValidateItem3**](ServiceRequestNamingApi.md#ValidateItem3) | **Post** /service/requestNaming/validator | Validates the new request naming rule for the &#x60;POST /requestNaming&#x60; request
[**ValidatePut1**](ServiceRequestNamingApi.md#ValidatePut1) | **Post** /service/requestNaming/{id}/validator | Validates the new request naming for the &#x60;PUT /requestNaming/{id}&#x60; request



## Delete6

> Delete6(ctx, id)

Deletes the specified request naming rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request naming rule to be deleted. | 

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


## Get

> ResourceNaming Get(ctx, )

Lists the global service resource requests.

Lists all extensions which currently are renamed to e.g., `Image`

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ResourceNaming**](ResourceNaming.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem5

> RequestNaming GetItem5(ctx, id)

Gets the parameters of the specified request naming rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request naming rule you&#39;re inquiring. | 

### Return type

[**RequestNaming**](RequestNaming.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList5

> StubList GetList5(ctx, )

Lists all request naming rules along with their parameters

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


## Post2

> EntityShortRepresentation Post2(ctx, optional)

Creates a new request naming rule

The new rule goes to the end of the rules list and will be the last to evaluate. Existing rules remain unaffected.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Post2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Post2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **optional.String**| Order of the new request naming rule. Set to &#x60;PREPEND&#x60; to prepend it to the list, &#x60;APPEND&#x60; to append it. Defaults to &#x60;APPEND&#x60;. | [default to APPEND]
 **requestNaming** | [**optional.Interface of RequestNaming**](RequestNaming.md)| The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## Put1

> EntityShortRepresentation Put1(ctx, id, optional)

Updates the specified request naming rule

If the rule with the specified ID doesn't exist, a new rule will be created at the end of the rules list and will be the last to evaluate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 
 **optional** | ***Put1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Put1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**optional.Interface of RequestNaming**](RequestNaming.md)| The JSON body of the request containing updated parameters of the request naming. | 

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


## Put2

> Put2(ctx, resourceNaming)

Updates the global service resource requests.

Update all extensions which are renamed to `Image` or `Binary`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceNaming** | [**ResourceNaming**](ResourceNaming.md)|  | 

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


## ReorderList2

> ReorderList2(ctx, optional)

Reorders the request namings

This request reorders the request namings according to the given list of IDs. Request namings not present in the body of the request will retain their relative ordering but will be ordered *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReorderList2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderList2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**optional.Interface of StubList**](StubList.md)| JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored. | 

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


## Validate

> Validate(ctx, resourceNaming)

Validates new resource requests settings for the `PUT /service/resourceRequest` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceNaming** | [**ResourceNaming**](ResourceNaming.md)|  | 

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


## ValidateItem3

> ValidateItem3(ctx, optional)

Validates the new request naming rule for the `POST /requestNaming` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateItem3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateItem3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestNaming** | [**optional.Interface of RequestNaming**](RequestNaming.md)| The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## ValidatePut1

> ValidatePut1(ctx, id, optional)

Validates the new request naming for the `PUT /requestNaming/{id}` request

If the rule with the specified ID doesn't exist, a new rule will be created at the end of the rules list and will be the last to evaluate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 
 **optional** | ***ValidatePut1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePut1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**optional.Interface of RequestNaming**](RequestNaming.md)| The JSON body of the request containing updated parameters of the request naming. | 

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

