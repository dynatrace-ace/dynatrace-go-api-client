# \ConditionalNamingApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamingRule**](ConditionalNamingApi.md#CreateNamingRule) | **Post** /conditionalNaming/{type} | Creates a new naming rule | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateNamingRule**](ConditionalNamingApi.md#CreateOrUpdateNamingRule) | **Put** /conditionalNaming/{type}/{id} | Updates the specified naming rule or creates it | maturity&#x3D;EARLY_ADOPTER
[**DeleteNamingRule**](ConditionalNamingApi.md#DeleteNamingRule) | **Delete** /conditionalNaming/{type}/{id} | Deletes the specified naming rule | maturity&#x3D;EARLY_ADOPTER
[**GetAllNamingRules**](ConditionalNamingApi.md#GetAllNamingRules) | **Get** /conditionalNaming/{type} | Lists all configured naming rules of one type | maturity&#x3D;EARLY_ADOPTER
[**GetSingleNamingRule**](ConditionalNamingApi.md#GetSingleNamingRule) | **Get** /conditionalNaming/{type}/{id} | Lists the parameters of the specified naming rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateNamingRule**](ConditionalNamingApi.md#ValidateNamingRule) | **Post** /conditionalNaming/{type}/validator | Validates a new naming rule for the &#x60;POST /conditionalNaming/{type}&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateNamingRule1**](ConditionalNamingApi.md#ValidateNamingRule1) | **Post** /conditionalNaming/{type}/{id}/validator | Validate updates of existing naming rules for the &#x60;PUT /conditionalNaming/{type}/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateNamingRule

> EntityShortRepresentation CreateNamingRule(ctx, type_, optional)

Creates a new naming rule | maturity=EARLY_ADOPTER

The body must not provide an ID as IDs are automatically assigned. The type of the rule and the type of the resource must match.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rule should be created. | 
 **optional** | ***CreateNamingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNamingRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditionalNamingRule** | [**optional.Interface of ConditionalNamingRule**](ConditionalNamingRule.md)| The JSON body of the request, containing parameters of the new naming rule. | 

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


## CreateOrUpdateNamingRule

> EntityShortRepresentation CreateOrUpdateNamingRule(ctx, type_, id, optional)

Updates the specified naming rule or creates it | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rule should be updated. | 
**id** | [**string**](.md)| The ID of the naming rule to update. | 
 **optional** | ***CreateOrUpdateNamingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateNamingRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditionalNamingRule** | [**optional.Interface of ConditionalNamingRule**](ConditionalNamingRule.md)| The JSON body of the request, containing updated parameters of the naming rule. | 

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


## DeleteNamingRule

> DeleteNamingRule(ctx, type_, id)

Deletes the specified naming rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type from which the naming rule should be deleted. | 
**id** | [**string**](.md)| The ID of the naming rule to delete. | 

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


## GetAllNamingRules

> StubList GetAllNamingRules(ctx, type_)

Lists all configured naming rules of one type | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rules should be queried. | 

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


## GetSingleNamingRule

> ConditionalNamingRule GetSingleNamingRule(ctx, type_, id)

Lists the parameters of the specified naming rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rule should be queried. | 
**id** | [**string**](.md)| The ID of the naming rule to query. | 

### Return type

[**ConditionalNamingRule**](ConditionalNamingRule.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateNamingRule

> ValidateNamingRule(ctx, type_, optional)

Validates a new naming rule for the `POST /conditionalNaming/{type}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rule should be validated. | 
 **optional** | ***ValidateNamingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateNamingRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditionalNamingRule** | [**optional.Interface of ConditionalNamingRule**](ConditionalNamingRule.md)| The JSON body of the request, containing the naming rule parameters to validate. | 

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


## ValidateNamingRule1

> ValidateNamingRule1(ctx, type_, id, optional)

Validate updates of existing naming rules for the `PUT /conditionalNaming/{type}/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The entity type for which the naming rule should be validated. | 
**id** | [**string**](.md)| The ID of the naming rule to validate. | 
 **optional** | ***ValidateNamingRule1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateNamingRule1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditionalNamingRule** | [**optional.Interface of ConditionalNamingRule**](ConditionalNamingRule.md)| The JSON body of the request, containing the naming rule parameters to validate. | 

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

