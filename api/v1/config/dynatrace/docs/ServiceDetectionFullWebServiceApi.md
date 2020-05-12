# \ServiceDetectionFullWebServiceApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateRule1**](ServiceDetectionFullWebServiceApi.md#CreateOrUpdateRule1) | **Put** /service/detectionRules/FULL_WEB_SERVICE/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**CreateRule1**](ServiceDetectionFullWebServiceApi.md#CreateRule1) | **Post** /service/detectionRules/FULL_WEB_SERVICE | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteRule1**](ServiceDetectionFullWebServiceApi.md#DeleteRule1) | **Delete** /service/detectionRules/FULL_WEB_SERVICE/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetRule1**](ServiceDetectionFullWebServiceApi.md#GetRule1) | **Get** /service/detectionRules/FULL_WEB_SERVICE/{id} | Shows the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListRulesSorted1**](ServiceDetectionFullWebServiceApi.md#ListRulesSorted1) | **Get** /service/detectionRules/FULL_WEB_SERVICE | Lists all full web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**ReorderList4**](ServiceDetectionFullWebServiceApi.md#ReorderList4) | **Put** /service/detectionRules/FULL_WEB_SERVICE/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**ValidatePostConfiguration1**](ServiceDetectionFullWebServiceApi.md#ValidatePostConfiguration1) | **Post** /service/detectionRules/FULL_WEB_SERVICE/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_SERVICE&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidatePutConfiguration1**](ServiceDetectionFullWebServiceApi.md#ValidatePutConfiguration1) | **Post** /service/detectionRules/FULL_WEB_SERVICE/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateRule1

> EntityShortRepresentation CreateOrUpdateRule1(ctx, id, optional)

Updates an existing service detection rule | maturity=EARLY_ADOPTER

If the rule with the specified ID doesn't exist, a new rule will be created and appended to the end of the rule list.    The request keeps an existing order of rules, unless the **order** parameter is set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the rule to be updated. | 
 **optional** | ***CreateOrUpdateRule1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateRule1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebServiceRule** | [**optional.Interface of FullWebServiceRule**](FullWebServiceRule.md)| The JSON body of the request containing updated parameters of the service detection rule. | 

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


## CreateRule1

> EntityShortRepresentation CreateRule1(ctx, optional)

Creates a new service detection rule | maturity=EARLY_ADOPTER

The body must not provide an ID as it will be automatically assigned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRule1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRule1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **optional.String**| The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to APPEND]
 **fullWebServiceRule** | [**optional.Interface of FullWebServiceRule**](FullWebServiceRule.md)| The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder&#x60; request. | 

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


## DeleteRule1

> DeleteRule1(ctx, id)

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


## GetRule1

> FullWebServiceRule GetRule1(ctx, id)

Shows the properties of the specified service detection rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required service detection rule. | 

### Return type

[**FullWebServiceRule**](FullWebServiceRule.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesSorted1

> StubList ListRulesSorted1(ctx, )

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


## ReorderList4

> ReorderList4(ctx, optional)

Reorders the service detection rules of the specified type | maturity=EARLY_ADOPTER

The request reorders the rules of the specified type according to the order of the IDs in the body of the request.    Rules that are omitted in the body of the request will retain their relative order but will be placed *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReorderList4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderList4Opts struct


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


## ValidatePostConfiguration1

> ValidatePostConfiguration1(ctx, fullWebServiceRule)

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_SERVICE` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md)|  | 

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


## ValidatePutConfiguration1

> ValidatePutConfiguration1(ctx, id, fullWebServiceRule)

Validate the payload for the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the service detection rule to be validated. | 
**fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md)|  | 

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

