# \ServiceDetectionOpaqueAndExternalWebRequestApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateRule2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#CreateOrUpdateRule2) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**CreateRule2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#CreateRule2) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteRule2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#DeleteRule2) | **Delete** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetRule2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#GetRule2) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Shows the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListRulesSorted2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#ListRulesSorted2) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST | Lists all full web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**ReorderList5**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#ReorderList5) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**ValidatePostConfiguration2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#ValidatePostConfiguration2) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidatePutConfiguration2**](ServiceDetectionOpaqueAndExternalWebRequestApi.md#ValidatePutConfiguration2) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOrUpdateRule2

> EntityShortRepresentation CreateOrUpdateRule2(ctx, id, optional)

Updates an existing service detection rule | maturity=EARLY_ADOPTER

If the rule with the specified ID doesn't exist, a new rule will be created and appended to the end of the rule list.    The request keeps an existing order of rules, unless the **order** parameter is set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the rule to be updated. | 
 **optional** | ***CreateOrUpdateRule2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateRule2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaqueAndExternalWebRequestRule** | [**optional.Interface of OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)| The JSON body of the request containing updated parameters of the service detection rule. | 

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


## CreateRule2

> EntityShortRepresentation CreateRule2(ctx, optional)

Creates a new service detection rule | maturity=EARLY_ADOPTER

The body must not provide an ID as it will be automatically assigned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRule2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRule2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **optional.String**| The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to APPEND]
 **opaqueAndExternalWebRequestRule** | [**optional.Interface of OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)| The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/reorder&#x60; request. | 

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


## DeleteRule2

> DeleteRule2(ctx, id)

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


## GetRule2

> OpaqueAndExternalWebRequestRule GetRule2(ctx, id)

Shows the properties of the specified service detection rule | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required service detection rule. | 

### Return type

[**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesSorted2

> StubList ListRulesSorted2(ctx, )

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


## ReorderList5

> ReorderList5(ctx, optional)

Reorders the service detection rules of the specified type | maturity=EARLY_ADOPTER

The request reorders the rules of the specified type according to the order of the IDs in the body of the request.    Rules that are omitted in the body of the request will retain their relative order but will be placed *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReorderList5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderList5Opts struct


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


## ValidatePostConfiguration2

> ValidatePostConfiguration2(ctx, opaqueAndExternalWebRequestRule)

Validates the payload for the `POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)|  | 

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


## ValidatePutConfiguration2

> ValidatePutConfiguration2(ctx, id, opaqueAndExternalWebRequestRule)

Validate the payload for the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the service detection rule to be validated. | 
**opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)|  | 

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

