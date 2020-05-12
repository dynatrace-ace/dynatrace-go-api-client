# \AzureCredentialsConfigurationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration2**](AzureCredentialsConfigurationApi.md#CreateConfiguration2) | **Post** /azure/credentials | Creates a new Azure credentials configuration
[**DeleteConfiguration3**](AzureCredentialsConfigurationApi.md#DeleteConfiguration3) | **Delete** /azure/credentials/{id} | Deletes the specified Azure credentials configuration
[**GetConfiguration7**](AzureCredentialsConfigurationApi.md#GetConfiguration7) | **Get** /azure/credentials/{id} | Gets the configuration of the specified Azure credentials
[**GetConfigurations1**](AzureCredentialsConfigurationApi.md#GetConfigurations1) | **Get** /azure/credentials | Lists all available Azure credentials configurations
[**UpdateConfiguration4**](AzureCredentialsConfigurationApi.md#UpdateConfiguration4) | **Put** /azure/credentials/{id} | Updates an existing Azure credentials configuration
[**ValidateConfiguration11**](AzureCredentialsConfigurationApi.md#ValidateConfiguration11) | **Post** /azure/credentials/validator | Validates the payload for the &#x60;POST /azure/credentials&#x60; request
[**ValidateConfigurationUpdate**](AzureCredentialsConfigurationApi.md#ValidateConfigurationUpdate) | **Post** /azure/credentials/{id}/validator | Validates the payload for the &#x60;PUT /azure/credentials/{id}&#x60; request



## CreateConfiguration2

> EntityShortRepresentation CreateConfiguration2(ctx, optional)

Creates a new Azure credentials configuration

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfiguration2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentials** | [**optional.Interface of AzureCredentials**](AzureCredentials.md)| The JSON body of the request. Contains parameters of the new Azure credentials configuration. | 

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


## DeleteConfiguration3

> DeleteConfiguration3(ctx, id)

Deletes the specified Azure credentials configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Azure credentials configuration to be deleted. | 

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


## GetConfiguration7

> AzureCredentials GetConfiguration7(ctx, id)

Gets the configuration of the specified Azure credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required Azure credentials configuration. | 

### Return type

[**AzureCredentials**](AzureCredentials.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurations1

> StubList GetConfigurations1(ctx, )

Lists all available Azure credentials configurations

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


## UpdateConfiguration4

> EntityShortRepresentation UpdateConfiguration4(ctx, id, optional)

Updates an existing Azure credentials configuration

* You can't change the Application ID or Directory ID of an Azure configuration. If you need to change them, create a new credentials configuration.  * If a configuration with the specified ID doesn't exist, a new configuration is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Azure credentials configuration to be updated. | 
 **optional** | ***UpdateConfiguration4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfiguration4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**optional.Interface of AzureCredentials**](AzureCredentials.md)| The JSON body of the request. Contains updated parameters of the Azure credentials configuration. | 

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


## ValidateConfiguration11

> ValidateConfiguration11(ctx, optional)

Validates the payload for the `POST /azure/credentials` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration11Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration11Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentials** | [**optional.Interface of AzureCredentials**](AzureCredentials.md)| The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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


## ValidateConfigurationUpdate

> ValidateConfigurationUpdate(ctx, id, optional)

Validates the payload for the `PUT /azure/credentials/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Azure credentials configuration to be validated. | 
 **optional** | ***ValidateConfigurationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfigurationUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**optional.Interface of AzureCredentials**](AzureCredentials.md)| The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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

