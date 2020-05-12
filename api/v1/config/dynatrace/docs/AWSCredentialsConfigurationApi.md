# \AWSCredentialsConfigurationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#CreateAwsCredentialsConfig) | **Post** /aws/credentials | Creates a new AWS credentials configuration
[**CreateOrUpdateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#CreateOrUpdateAwsCredentialsConfig) | **Put** /aws/credentials/{id} | Updates an existing AWS credentials configuration
[**DeleteAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#DeleteAwsCredentialsConfig) | **Delete** /aws/credentials/{id} | Deletes the specified AWS credentials configuration
[**ListAwsCredentials**](AWSCredentialsConfigurationApi.md#ListAwsCredentials) | **Get** /aws/credentials | Lists all available AWS credentials configurations
[**ReadAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#ReadAwsCredentialsConfig) | **Get** /aws/credentials/{id} | Gets the configuration of the specified AWS credentials
[**ReadIamExternalIdToken**](AWSCredentialsConfigurationApi.md#ReadIamExternalIdToken) | **Get** /aws/iamExternalId | Gets the external ID token for setting an IAM role
[**ValidateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#ValidateAwsCredentialsConfig) | **Post** /aws/credentials/validator | Validates the payload for the &#x60;POST /aws/credentials&#x60; request
[**ValidateAwsCredentialsConfig1**](AWSCredentialsConfigurationApi.md#ValidateAwsCredentialsConfig1) | **Post** /aws/credentials/{id}/validator | Validates the payload for the &#x60;PUT /aws/credentials/{id}&#x60; request



## CreateAwsCredentialsConfig

> EntityShortRepresentation CreateAwsCredentialsConfig(ctx, optional)

Creates a new AWS credentials configuration

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAwsCredentialsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAwsCredentialsConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**optional.Interface of AwsCredentialsConfig**](AwsCredentialsConfig.md)| The JSON body of the request. Contains parameters of the new AWS credentials configuration. | 

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


## CreateOrUpdateAwsCredentialsConfig

> EntityShortRepresentation CreateOrUpdateAwsCredentialsConfig(ctx, id, optional)

Updates an existing AWS credentials configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the AWS credentials configuration to be updated. | 
 **optional** | ***CreateOrUpdateAwsCredentialsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateAwsCredentialsConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**optional.Interface of AwsCredentialsConfig**](AwsCredentialsConfig.md)| The JSON body of the request. Contains updated parameters of the AWS credentials configuration. | 

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


## DeleteAwsCredentialsConfig

> DeleteAwsCredentialsConfig(ctx, id)

Deletes the specified AWS credentials configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of AWS credentials configuration to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAwsCredentials

> StubList ListAwsCredentials(ctx, )

Lists all available AWS credentials configurations

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


## ReadAwsCredentialsConfig

> AwsCredentialsConfig ReadAwsCredentialsConfig(ctx, id)

Gets the configuration of the specified AWS credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the specified AWS credentials configuration. | 

### Return type

[**AwsCredentialsConfig**](AwsCredentialsConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIamExternalIdToken

> AwsIamToken ReadIamExternalIdToken(ctx, )

Gets the external ID token for setting an IAM role

You'll need it for the role-based AWS authentication.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AwsIamToken**](AwsIamToken.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAwsCredentialsConfig

> ValidateAwsCredentialsConfig(ctx, optional)

Validates the payload for the `POST /aws/credentials` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateAwsCredentialsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAwsCredentialsConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**optional.Interface of AwsCredentialsConfig**](AwsCredentialsConfig.md)| The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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


## ValidateAwsCredentialsConfig1

> ValidateAwsCredentialsConfig1(ctx, id, optional)

Validates the payload for the `PUT /aws/credentials/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the AWS credentials configuration to be validated. | 
 **optional** | ***ValidateAwsCredentialsConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAwsCredentialsConfig1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**optional.Interface of AwsCredentialsConfig**](AwsCredentialsConfig.md)| The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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

