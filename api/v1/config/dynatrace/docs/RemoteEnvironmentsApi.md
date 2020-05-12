# \RemoteEnvironmentsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](RemoteEnvironmentsApi.md#CreateConfig) | **Post** /remoteEnvironments | Creates a new remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateConfig**](RemoteEnvironmentsApi.md#CreateOrUpdateConfig) | **Put** /remoteEnvironments/{id} | Updates an existing remote environment configuration or creates a new one | maturity&#x3D;EARLY_ADOPTER
[**DeleteConfig**](RemoteEnvironmentsApi.md#DeleteConfig) | **Delete** /remoteEnvironments/{id} | Deletes the specified remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**GetConfig1**](RemoteEnvironmentsApi.md#GetConfig1) | **Get** /remoteEnvironments/{id} | Gets the properties of the specified remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**ReturnAllConfigs**](RemoteEnvironmentsApi.md#ReturnAllConfigs) | **Get** /remoteEnvironments | Lists all remote environment configurations | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateConfiguration1**](RemoteEnvironmentsApi.md#ValidateCreateConfiguration1) | **Post** /remoteEnvironments/validator | Validates the payload for the &#x60;POST /remoteEnvironments&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateOrUpdateConfiguration1**](RemoteEnvironmentsApi.md#ValidateCreateOrUpdateConfiguration1) | **Post** /remoteEnvironments/{id}/validator | Validates the payload for the &#x60;PUT /remoteEnvironments/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateConfig

> RemoteEnvironmentConfigStub CreateConfig(ctx, optional)

Creates a new remote environment configuration | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.   Be sure to prepare a token with the **Fetch data from a remote environment** (`RestRequestForwarding`) scope. You can create such a token via [Tokens API](https://www.dynatrace.com/support/help/shortlink/api-tokens-post-new).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteEnvironmentConfigDto** | [**optional.Interface of RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)| The JSON body of the request. Contains parameters of the new remote environment configuration. | 

### Return type

[**RemoteEnvironmentConfigStub**](RemoteEnvironmentConfigStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateConfig

> RemoteEnvironmentConfigStub CreateOrUpdateConfig(ctx, id, optional)

Updates an existing remote environment configuration or creates a new one | maturity=EARLY_ADOPTER

If a remote environment configuration with the specified ID doesn't exist, a new configuration is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the configuration to be updated.    If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteEnvironmentConfigDto** | [**optional.Interface of RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)| The JSON body of the request. Contains updated parameters of the remote environment configuration. | 

### Return type

[**RemoteEnvironmentConfigStub**](RemoteEnvironmentConfigStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> DeleteConfig(ctx, id)

Deletes the specified remote environment configuration | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the configuration to be deleted. | 

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


## GetConfig1

> RemoteEnvironmentConfigDto GetConfig1(ctx, id)

Gets the properties of the specified remote environment configuration | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required configuration. | 

### Return type

[**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllConfigs

> RemoteEnvironmentConfigListDto ReturnAllConfigs(ctx, )

Lists all remote environment configurations | maturity=EARLY_ADOPTER

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RemoteEnvironmentConfigListDto**](RemoteEnvironmentConfigListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateConfiguration1

> ValidateCreateConfiguration1(ctx, optional)

Validates the payload for the `POST /remoteEnvironments` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateCreateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateConfiguration1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteEnvironmentConfigDto** | [**optional.Interface of RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)| The JSON body of the request. Contains the remote environment configuration to be validated. | 

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


## ValidateCreateOrUpdateConfiguration1

> ValidateCreateOrUpdateConfiguration1(ctx, id, optional)

Validates the payload for the `PUT /remoteEnvironments/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the remote environment configuration to be validated. | 
 **optional** | ***ValidateCreateOrUpdateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateOrUpdateConfiguration1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteEnvironmentConfigDto** | [**optional.Interface of RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)| The JSON body of the request. Contains the remote environment configuration to be validated. | 

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

