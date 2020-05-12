# \ServiceDetectionFullWebRequestApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateRule**](ServiceDetectionFullWebRequestApi.md#CreateOrUpdateRule) | **Put** /service/detectionRules/FULL_WEB_REQUEST/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**CreateRule**](ServiceDetectionFullWebRequestApi.md#CreateRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteRule**](ServiceDetectionFullWebRequestApi.md#DeleteRule) | **Delete** /service/detectionRules/FULL_WEB_REQUEST/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetRule**](ServiceDetectionFullWebRequestApi.md#GetRule) | **Get** /service/detectionRules/FULL_WEB_REQUEST/{id} | Gets the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListRulesSorted**](ServiceDetectionFullWebRequestApi.md#ListRulesSorted) | **Get** /service/detectionRules/FULL_WEB_REQUEST | Lists all full web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**ReorderList3**](ServiceDetectionFullWebRequestApi.md#ReorderList3) | **Put** /service/detectionRules/FULL_WEB_REQUEST/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**ValidatePostConfiguration**](ServiceDetectionFullWebRequestApi.md#ValidatePostConfiguration) | **Post** /service/detectionRules/FULL_WEB_REQUEST/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_REQUEST&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidatePutConfiguration**](ServiceDetectionFullWebRequestApi.md#ValidatePutConfiguration) | **Post** /service/detectionRules/FULL_WEB_REQUEST/{id}/validator | Validates the payload for the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateRule

> EntityShortRepresentation CreateOrUpdateRule(ctx, id, optional)

Updates an existing service detection rule | maturity=EARLY_ADOPTER

If a rule with the specified ID doesn't exist, a new rule is created and appended to the end of the rule list.    The request keeps the existing order of rules, unless the **order** parameter is set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the rule to be updated. | 
 **optional** | ***CreateOrUpdateRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**optional.Interface of FullWebRequestRule**](FullWebRequestRule.md)| The JSON body of the request. Contains updated parameters of the service detection rule. | 

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


## CreateRule

> EntityShortRepresentation CreateRule(ctx, optional)

Creates a new service detection rule | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **optional.String**| The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used. | [default to APPEND]
 **fullWebRequestRule** | [**optional.Interface of FullWebRequestRule**](FullWebRequestRule.md)| The JSON body of the request. Contains parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order, use the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/reorder&#x60; request. | 

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


## DeleteRule

> DeleteRule(ctx, id)

Deletes the specified service detection rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the service detection rule to be deleted. | 

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


## GetRule

> FullWebRequestRule GetRule(ctx, id)

Gets the properties of the specified service detection rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required service detection rule. | 

### Return type

[**FullWebRequestRule**](FullWebRequestRule.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesSorted

> StubList ListRulesSorted(ctx, )

Lists all full web service detection rules | maturity=EARLY_ADOPTER

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


## ReorderList3

> ReorderList3(ctx, optional)

Reorders the service detection rules of the specified type | maturity=EARLY_ADOPTER

The request reorders the rules of the specified type according to the order of the IDs in the body of the request.    Rules that are omitted in the body of the request will retain their relative order but will be placed *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReorderList3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderList3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**optional.Interface of StubList**](StubList.md)| The JSON body of the request containing the service detection rules in the required order. | 

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


## ValidatePostConfiguration

> ValidatePostConfiguration(ctx, optional)

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_REQUEST` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatePostConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePostConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullWebRequestRule** | [**optional.Interface of FullWebRequestRule**](FullWebRequestRule.md)| The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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


## ValidatePutConfiguration

> ValidatePutConfiguration(ctx, id, optional)

Validates the payload for the `PUT /service/detectionRules/FULL_WEB_REQUEST/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the service detection rule to be validated. | 
 **optional** | ***ValidatePutConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePutConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**optional.Interface of FullWebRequestRule**](FullWebRequestRule.md)| The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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

