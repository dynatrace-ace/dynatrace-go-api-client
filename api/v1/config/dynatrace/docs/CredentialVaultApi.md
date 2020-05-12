# \CredentialVaultApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredentials**](CredentialVaultApi.md#AddCredentials) | **Post** /credentials | Creates a new credentials set | maturity&#x3D;EARLY_ADOPTER
[**GetCredentials**](CredentialVaultApi.md#GetCredentials) | **Get** /credentials | Lists all sets of credentials for synthetic monitors stored in your environment | maturity&#x3D;EARLY_ADOPTER
[**GetCredentials1**](CredentialVaultApi.md#GetCredentials1) | **Get** /credentials/{id} | Gets the metadata of the specified credentials set | maturity&#x3D;EARLY_ADOPTER
[**RemoveCredentials**](CredentialVaultApi.md#RemoveCredentials) | **Delete** /credentials/{id} | Deletes the specified credentials set | maturity&#x3D;EARLY_ADOPTER
[**UpdateCredentials**](CredentialVaultApi.md#UpdateCredentials) | **Put** /credentials/{id} | Updates the specified credentials set | maturity&#x3D;EARLY_ADOPTER



## AddCredentials

> CredentialsResponseElement AddCredentials(ctx, optional)

Creates a new credentials set | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**optional.Interface of Credentials**](Credentials.md)| The JSON body of the request. Contains parameters of the new credentials set. | 

### Return type

[**CredentialsResponseElement**](CredentialsResponseElement.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> CredentialsList GetCredentials(ctx, optional)

Lists all sets of credentials for synthetic monitors stored in your environment | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Filters the result by the specified credentials type. | 

### Return type

[**CredentialsList**](CredentialsList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials1

> CredentialsResponseElement GetCredentials1(ctx, id)

Gets the metadata of the specified credentials set | maturity=EARLY_ADOPTER

The credentials set itself (username/certificate and password) is not included in the response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Dynatrace entity ID of the required credentials set. | 

### Return type

[**CredentialsResponseElement**](CredentialsResponseElement.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCredentials

> RemoveCredentials(ctx, id)

Deletes the specified credentials set | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the credentials set to be deleted. | 

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


## UpdateCredentials

> CredentialsResponseElement UpdateCredentials(ctx, id, optional)

Updates the specified credentials set | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The Dynatrace entity ID of the credentials set to be updated. | 
 **optional** | ***UpdateCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentials** | [**optional.Interface of Credentials**](Credentials.md)| The JSON body of the request. Contains updated parameters of the credentials set. | 

### Return type

[**CredentialsResponseElement**](CredentialsResponseElement.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

