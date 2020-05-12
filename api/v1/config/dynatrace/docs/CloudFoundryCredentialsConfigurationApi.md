# \CloudFoundryCredentialsConfigurationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig1**](CloudFoundryCredentialsConfigurationApi.md#CreateConfig1) | **Post** /cloudFoundry/credentials | Create new credentials for a single foundation. | maturity&#x3D;EARLY_ADOPTER
[**DeleteConfig1**](CloudFoundryCredentialsConfigurationApi.md#DeleteConfig1) | **Delete** /cloudFoundry/credentials/{id} | Delete the specified Cloud Foundry foundation credentials. | maturity&#x3D;EARLY_ADOPTER
[**GetConfig2**](CloudFoundryCredentialsConfigurationApi.md#GetConfig2) | **Get** /cloudFoundry/credentials/{id} | Show the properties for the specified Cloud Foundry foundation credentials. | maturity&#x3D;EARLY_ADOPTER
[**ListConfigs**](CloudFoundryCredentialsConfigurationApi.md#ListConfigs) | **Get** /cloudFoundry/credentials | List all the Cloud Foundry foundations credentials. | maturity&#x3D;EARLY_ADOPTER
[**UpdateConfig**](CloudFoundryCredentialsConfigurationApi.md#UpdateConfig) | **Put** /cloudFoundry/credentials/{id} | Create or update credentials for a single Cloud Foundry foundation. | maturity&#x3D;EARLY_ADOPTER
[**ValidatePOSTConfig**](CloudFoundryCredentialsConfigurationApi.md#ValidatePOSTConfig) | **Post** /cloudFoundry/credentials/validator | Validate that creating credentials would be successful. | maturity&#x3D;EARLY_ADOPTER
[**ValidatePUTConfig**](CloudFoundryCredentialsConfigurationApi.md#ValidatePUTConfig) | **Post** /cloudFoundry/credentials/{id}/validator | Validate that updating or creating credentials would be successful. | maturity&#x3D;EARLY_ADOPTER



## CreateConfig1

> EntityShortRepresentation CreateConfig1(ctx, cloudFoundryCredentials)

Create new credentials for a single foundation. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md)| &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique. | 

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


## DeleteConfig1

> DeleteConfig1(ctx, id)

Delete the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Cloud Foundry foundation credentials to be deleted. | 

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


## GetConfig2

> CloudFoundryCredentials GetConfig2(ctx, id)

Show the properties for the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the required Cloud Foundry foundation credentials. | 

### Return type

[**CloudFoundryCredentials**](CloudFoundryCredentials.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigs

> StubList ListConfigs(ctx, )

List all the Cloud Foundry foundations credentials. | maturity=EARLY_ADOPTER

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


## UpdateConfig

> EntityShortRepresentation UpdateConfig(ctx, id, cloudFoundryCredentials)

Create or update credentials for a single Cloud Foundry foundation. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Cloud Foundry foundation credentials. | 
**cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md)| &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates, the existing one will be used. &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be set and may not differ from the existing config if it exists. Use this endpoint for copying credentials between environments while keeping their IDs and for updating existing credentials. You can *not* use this to create new credentials with an arbitrary ID, use POST instead. | 

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


## ValidatePOSTConfig

> ValidatePOSTConfig(ctx, cloudFoundryCredentials)

Validate that creating credentials would be successful. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md)| &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique. | 

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


## ValidatePUTConfig

> ValidatePUTConfig(ctx, id, cloudFoundryCredentials)

Validate that updating or creating credentials would be successful. | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the Cloud Foundry foundation credentials. | 
**cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md)| &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates. See the constraints for the PUT /cloudFoundry/credentials/{id} request. | 

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

