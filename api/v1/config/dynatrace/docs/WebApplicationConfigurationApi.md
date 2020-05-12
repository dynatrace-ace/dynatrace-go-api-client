# \WebApplicationConfigurationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](WebApplicationConfigurationApi.md#CreateConfiguration) | **Post** /applications/web | Creates a new web application
[**CreateOrUpdateConfiguration**](WebApplicationConfigurationApi.md#CreateOrUpdateConfiguration) | **Put** /applications/web/{id} | Updates the configuration of the specified web application or creates a new one
[**CreateOrUpdateDefaultConfiguration**](WebApplicationConfigurationApi.md#CreateOrUpdateDefaultConfiguration) | **Put** /applications/web/default | Updates the configuration of the default web application
[**DeleteConfiguration**](WebApplicationConfigurationApi.md#DeleteConfiguration) | **Delete** /applications/web/{id} | Deletes the specified web application
[**GetConfiguration3**](WebApplicationConfigurationApi.md#GetConfiguration3) | **Get** /applications/web/{id} | Gets the configuration of the specified web application
[**GetDataPrivacySettings**](WebApplicationConfigurationApi.md#GetDataPrivacySettings) | **Get** /applications/web/{id}/dataPrivacy | Gets the data privacy settings of the specified web application
[**GetDefaultApplication**](WebApplicationConfigurationApi.md#GetDefaultApplication) | **Get** /applications/web/default | Gets the configuration of the default web application
[**GetDefaultApplicationDataPrivacySettings**](WebApplicationConfigurationApi.md#GetDefaultApplicationDataPrivacySettings) | **Get** /applications/web/default/dataPrivacy | Gets the data privacy settings of the default web application
[**ListConfigurations**](WebApplicationConfigurationApi.md#ListConfigurations) | **Get** /applications/web | Lists all existing web applications
[**ListDataPrivacySettings**](WebApplicationConfigurationApi.md#ListDataPrivacySettings) | **Get** /applications/web/dataPrivacy | Lists data privacy settings of all web applications
[**UpdateDataPrivacySettings**](WebApplicationConfigurationApi.md#UpdateDataPrivacySettings) | **Put** /applications/web/{id}/dataPrivacy | Updates the data privacy settings of the specified web application
[**UpdateDefaultApplicationDataPrivacySettings**](WebApplicationConfigurationApi.md#UpdateDefaultApplicationDataPrivacySettings) | **Put** /applications/web/default/dataPrivacy | Updates the data privacy settings of the default web application
[**ValidateConfiguration3**](WebApplicationConfigurationApi.md#ValidateConfiguration3) | **Post** /applications/web/validator | Validates the configuration of the web application for the &#x60;POST /applications/web&#x60; request
[**ValidateConfiguration4**](WebApplicationConfigurationApi.md#ValidateConfiguration4) | **Post** /applications/web/{id}/validator | Validates the configuration of the web application for the &#x60;PUT /applications/web/{id}&#x60; request.
[**ValidateDataPrivacySettings**](WebApplicationConfigurationApi.md#ValidateDataPrivacySettings) | **Post** /applications/web/{id}/dataPrivacy/validator | Validates data privacy settings for the &#x60;PUT /applications/web/{id}/dataPrivacy&#x60; request
[**ValidateDefaultApplicationDataPrivacySettings**](WebApplicationConfigurationApi.md#ValidateDefaultApplicationDataPrivacySettings) | **Post** /applications/web/default/dataPrivacy/validator | Validates data privacy settings of the default web application for the &#x60;PUT /applications/web/default/dataPrivacy&#x60; request
[**ValidateDefaultConfiguration**](WebApplicationConfigurationApi.md#ValidateDefaultConfiguration) | **Post** /applications/web/default/validator | Validates the configuration of the default web application for the &#x60;PUT /applications/web/default&#x60; request



## CreateConfiguration

> EntityShortRepresentation CreateConfiguration(ctx, optional)

Creates a new web application

The body must not provide an ID as that will be automatically assigned by Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing parameters of the new web application configuraiton. | 

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


## CreateOrUpdateConfiguration

> EntityShortRepresentation CreateOrUpdateConfiguration(ctx, id, optional)

Updates the configuration of the specified web application or creates a new one

If the application with the specified ID does not exist, a new application will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application to update.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing updated configuration of the web application. | 

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


## CreateOrUpdateDefaultConfiguration

> CreateOrUpdateDefaultConfiguration(ctx, optional)

Updates the configuration of the default web application

Default application is pre-configured in your Dynatrace environment. By default all traffic goes to this application.   After you configure your own applications, all the traffic, which doesn't fit to any of your applications, goes to the default one.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOrUpdateDefaultConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateDefaultConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing the new parameters of the default web application. | 

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


## DeleteConfiguration

> DeleteConfiguration(ctx, id)

Deletes the specified web application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application to delete. | 

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


## GetConfiguration3

> WebApplicationConfig GetConfiguration3(ctx, id)

Gets the configuration of the specified web application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the requested web application. | 

### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataPrivacySettings

> ApplicationDataPrivacy GetDataPrivacySettings(ctx, id)

Gets the data privacy settings of the specified web application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application where you want to check data privacy settings. | 

### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplication

> WebApplicationConfig GetDefaultApplication(ctx, )

Gets the configuration of the default web application

Default application is pre-configured in your Dynatrace environment. By default all traffic goes to this application.   After you configure your own applications, all the traffic, which doesn't fit to any of your applications, goes to the default one.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplicationDataPrivacySettings

> ApplicationDataPrivacy GetDefaultApplicationDataPrivacySettings(ctx, )

Gets the data privacy settings of the default web application

Default application is pre-configured in your Dynatrace environment. By default all the traffic goes to this application.   After you configure your own applications, all the traffic, which doesn't fit to any of your applications, goes to the default one.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations

> StubList ListConfigurations(ctx, )

Lists all existing web applications

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


## ListDataPrivacySettings

> ApplicationDataPrivacyList ListDataPrivacySettings(ctx, )

Lists data privacy settings of all web applications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ApplicationDataPrivacyList**](ApplicationDataPrivacyList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataPrivacySettings

> UpdateDataPrivacySettings(ctx, id, optional)

Updates the data privacy settings of the specified web application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application, where you want to update data privacy settings. | 
 **optional** | ***UpdateDataPrivacySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDataPrivacySettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)| JSON body of the request, containing new data privacy settings. | 

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


## UpdateDefaultApplicationDataPrivacySettings

> UpdateDefaultApplicationDataPrivacySettings(ctx, optional)

Updates the data privacy settings of the default web application

Default application is pre-configured in your Dynatrace environment. By default all traffic goes to this application.   After you configure your own applications, all the traffic, which doesn't fit to any of your applications, goes to the default one.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateDefaultApplicationDataPrivacySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDefaultApplicationDataPrivacySettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)| JSON body of the request, containing new data privacy settings. | 

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


## ValidateConfiguration3

> ValidateConfiguration3(ctx, optional)

Validates the configuration of the web application for the `POST /applications/web` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing the web application configuration to validate. | 

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


## ValidateConfiguration4

> ValidateConfiguration4(ctx, id, optional)

Validates the configuration of the web application for the `PUT /applications/web/{id}` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application to validate. | 
 **optional** | ***ValidateConfiguration4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateConfiguration4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing the web application configuration to validate. | 

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


## ValidateDataPrivacySettings

> ValidateDataPrivacySettings(ctx, id, optional)

Validates data privacy settings for the `PUT /applications/web/{id}/dataPrivacy` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the web application, where you want to validate data privacy settings. | 
 **optional** | ***ValidateDataPrivacySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDataPrivacySettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)| JSON body of the request, containing new data privacy settings. | 

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


## ValidateDefaultApplicationDataPrivacySettings

> ValidateDefaultApplicationDataPrivacySettings(ctx, optional)

Validates data privacy settings of the default web application for the `PUT /applications/web/default/dataPrivacy` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDefaultApplicationDataPrivacySettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDefaultApplicationDataPrivacySettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)| JSON body of the request, containing the data privacy settings to validate. | 

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


## ValidateDefaultConfiguration

> ValidateDefaultConfiguration(ctx, optional)

Validates the configuration of the default web application for the `PUT /applications/web/default` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDefaultConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateDefaultConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)| JSON body of the request, containing web application configuration to validate. | 

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

