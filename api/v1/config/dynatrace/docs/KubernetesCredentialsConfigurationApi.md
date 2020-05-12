# \KubernetesCredentialsConfigurationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration3**](KubernetesCredentialsConfigurationApi.md#CreateConfiguration3) | **Post** /kubernetes/credentials | Creates a new Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**DeleteConfiguration4**](KubernetesCredentialsConfigurationApi.md#DeleteConfiguration4) | **Delete** /kubernetes/credentials/{id} | Deletes the specified Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**GetConfiguration8**](KubernetesCredentialsConfigurationApi.md#GetConfiguration8) | **Get** /kubernetes/credentials/{id} | Gets the configuration of the specified Kubernetes credentials | maturity&#x3D;EARLY_ADOPTER
[**GetConfigurations2**](KubernetesCredentialsConfigurationApi.md#GetConfigurations2) | **Get** /kubernetes/credentials | Lists all available Kubernetes credentials configurations | maturity&#x3D;EARLY_ADOPTER
[**UpdateConfiguration5**](KubernetesCredentialsConfigurationApi.md#UpdateConfiguration5) | **Put** /kubernetes/credentials/{id} | Updates an existing Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**ValidateConfiguration12**](KubernetesCredentialsConfigurationApi.md#ValidateConfiguration12) | **Post** /kubernetes/credentials/validator | Validates the payload for the &#x60;POST /kubernetes/credentials&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateConfigurationUpdate1**](KubernetesCredentialsConfigurationApi.md#ValidateConfigurationUpdate1) | **Post** /kubernetes/credentials/{id}/validator | Validates the payload for the &#x60;PUT /kubernetes/credentials/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateConfiguration3

> EntityShortRepresentation CreateConfiguration3(ctx, optional)

Creates a new Kubernetes credentials configuration | maturity=EARLY_ADOPTER

The body must not provide an ID as it will be automatically assigned by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfiguration3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfiguration3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCredentials** | [**optional.Interface of KubernetesCredentials**](KubernetesCredentials.md)| The JSON body of the request. Contains parameters of the new Kubernetes credentials configuration. | 

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


## DeleteConfiguration4

> DeleteConfiguration4(ctx, id)

Deletes the specified Kubernetes credentials configuration | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Kubernetes credentials configuration be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration8

> KubernetesCredentials GetConfiguration8(ctx, id)

Gets the configuration of the specified Kubernetes credentials | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required Kubernetes credentials configuration. | 

### Return type

[**KubernetesCredentials**](KubernetesCredentials.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurations2

> StubList GetConfigurations2(ctx, )

Lists all available Kubernetes credentials configurations | maturity=EARLY_ADOPTER

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


## UpdateConfiguration5

> EntityShortRepresentation UpdateConfiguration5(ctx, id, optional)

Updates an existing Kubernetes credentials configuration | maturity=EARLY_ADOPTER

You can't change the URL of the Kubernetes cluster.   If a configuration with the specified ID doesn't exist, a new configuration is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Kubernetes credentials configuration to be updated. | 
 **optional** | ***UpdateConfiguration5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfiguration5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredentials** | [**optional.Interface of KubernetesCredentials**](KubernetesCredentials.md)| The JSON body of the request. Contains updated parameters of the Kubernetes credentials configuration. | 

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


## ValidateConfiguration12

> ValidateConfiguration12(ctx, optional)

Validates the payload for the `POST /kubernetes/credentials` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration12Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration12Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCredentials** | [**optional.Interface of KubernetesCredentials**](KubernetesCredentials.md)| The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. | 

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


## ValidateConfigurationUpdate1

> ValidateConfigurationUpdate1(ctx, id, optional)

Validates the payload for the `PUT /kubernetes/credentials/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Kubernetes credentials configuration to be validated. | 
 **optional** | ***ValidateConfigurationUpdate1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfigurationUpdate1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredentials** | [**optional.Interface of KubernetesCredentials**](KubernetesCredentials.md)| The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. | 

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

