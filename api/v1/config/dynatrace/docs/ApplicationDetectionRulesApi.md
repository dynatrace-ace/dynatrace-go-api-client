# \ApplicationDetectionRulesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ApplicationDetectionRulesApi.md#Delete) | **Delete** /applicationDetectionRules/{id} | Deletes the specified application detection rule
[**GetConfig**](ApplicationDetectionRulesApi.md#GetConfig) | **Get** /applicationDetectionRules/{id} | Gets the parameters of the specified application detection rule
[**ListConfigurations1**](ApplicationDetectionRulesApi.md#ListConfigurations1) | **Get** /applicationDetectionRules | Lists all available application detection rules
[**PostConfig**](ApplicationDetectionRulesApi.md#PostConfig) | **Post** /applicationDetectionRules | Creates a new application detection rule
[**PutConfig**](ApplicationDetectionRulesApi.md#PutConfig) | **Put** /applicationDetectionRules/{id} | Updates the specified application detection rule
[**ReorderList**](ApplicationDetectionRulesApi.md#ReorderList) | **Put** /applicationDetectionRules/order | Reorders the application detection rules
[**ValidateConfig**](ApplicationDetectionRulesApi.md#ValidateConfig) | **Post** /applicationDetectionRules/{id}/validator | Validate the payload for the &#x60;PUT /applicationDetection/{id}&#x60; request
[**ValidateConfig1**](ApplicationDetectionRulesApi.md#ValidateConfig1) | **Post** /applicationDetectionRules/validator | Validates the payload for the &#x60;POST /applicationDetection&#x60; request



## Delete

> Delete(ctx, id)

Deletes the specified application detection rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the application detection rule to be deleted. | 

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


## GetConfig

> ApplicationDetectionRuleConfig GetConfig(ctx, id)

Gets the parameters of the specified application detection rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required application detection rule. | 

### Return type

[**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations1

> StubList ListConfigurations1(ctx, )

Lists all available application detection rules

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


## PostConfig

> EntityShortRepresentation PostConfig(ctx, optional)

Creates a new application detection rule

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.   You can only create detection rules for an existing application. If you need to create a rule for an application that doesn't exist yet, [create an application first](https://www.dynatrace.com/support/help/shortlink/api-config-web-app-post-web-app) and then configure detection rules for it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **optional.String**| The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used. | 
 **applicationDetectionRuleConfig** | [**optional.Interface of ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)| The JSON body of the request. Contains configuration of the new application detection rule.    You must not specify the ID of the rule.   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /applicationDetectionRules/order&#x60; request. | 

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


## PutConfig

> EntityShortRepresentation PutConfig(ctx, id, optional)

Updates the specified application detection rule

If the application detection rule with the specified ID doesn't exist, a new application is created and appended to the end of the rule list.   If the **order** parameter is set for an existing rule, the request uses this value. Otherwise it keeps the existing order of rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the application detection rule to be updated.    If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***PutConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDetectionRuleConfig** | [**optional.Interface of ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)| The JSON body of the request. Contains updated parameters of the application detection rule.    If the **order** parameter is set, the rule is placed to this position. | 

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


## ReorderList

> ReorderList(ctx, optional)

Reorders the application detection rules

This request reorders the application detection rules according to the submitted list of IDs. Application detection rules not present in the body of the request will retain their relative ordering but are placed *after* all those present in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReorderListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReorderListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**optional.Interface of StubList**](StubList.md)| The JSON body of the request. Contains the IDs of the application detection rules in the desired order. Any further properties (**name**, **description**) are ignored. | 

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


## ValidateConfig

> ValidateConfig(ctx, id, optional)

Validate the payload for the `PUT /applicationDetection/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the application detection rule to be validated.    If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***ValidateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDetectionRuleConfig** | [**optional.Interface of ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)| The JSON body of the request. Contains the application detection rule to be validated. | 

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


## ValidateConfig1

> ValidateConfig1(ctx, optional)

Validates the payload for the `POST /applicationDetection` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfig1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionRuleConfig** | [**optional.Interface of ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)| The JSON body of the request. Contains the application detection rule to be validated. | 

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

