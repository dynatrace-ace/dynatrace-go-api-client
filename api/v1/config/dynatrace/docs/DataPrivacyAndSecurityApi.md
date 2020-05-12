# \DataPrivacyAndSecurityApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataPrivacySettings1**](DataPrivacyAndSecurityApi.md#GetDataPrivacySettings1) | **Get** /dataPrivacy | Lists the global data privacy and security settings.
[**UpdateDataPrivacySettings1**](DataPrivacyAndSecurityApi.md#UpdateDataPrivacySettings1) | **Put** /dataPrivacy | Updates the global data privacy and security settings.
[**ValidateConfiguration5**](DataPrivacyAndSecurityApi.md#ValidateConfiguration5) | **Post** /dataPrivacy/validator | Validates new data privacy and security settings for the &#x60;PUT /dataPrivacy&#x60; request.



## GetDataPrivacySettings1

> DataPrivacyAndSecurity GetDataPrivacySettings1(ctx, )

Lists the global data privacy and security settings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataPrivacySettings1

> UpdateDataPrivacySettings1(ctx, dataPrivacyAndSecurity)

Updates the global data privacy and security settings.

This request updates global settings, affecting all your applications. To update application-specific data privacy settings, use the `PUT /applications/web/{id}/dataPrivacy` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPrivacyAndSecurity** | [**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md)|  | 

### Return type

 (empty response body)

### Authorization

[DataPrivacy](../README.md#DataPrivacy)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConfiguration5

> ValidateConfiguration5(ctx, dataPrivacyAndSecurity)

Validates new data privacy and security settings for the `PUT /dataPrivacy` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPrivacyAndSecurity** | [**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md)|  | 

### Return type

 (empty response body)

### Authorization

[DataPrivacy](../README.md#DataPrivacy)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

