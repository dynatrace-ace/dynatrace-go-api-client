# \NotificationsApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationConfig**](NotificationsApi.md#CreateNotificationConfig) | **Post** /notifications | Creates a new notification configuration
[**CreateOrUpdateNotificationConfig**](NotificationsApi.md#CreateOrUpdateNotificationConfig) | **Put** /notifications/{id} | Updates an existing notification configuration or creates a new one
[**DeleteNotificatonConfig**](NotificationsApi.md#DeleteNotificatonConfig) | **Delete** /notifications/{id} | Deletes the specified notification configuration
[**GetNotificationConfig**](NotificationsApi.md#GetNotificationConfig) | **Get** /notifications/{id} | Gets the properties of the specified notification configuration
[**ReturnAllNotificationCofigs**](NotificationsApi.md#ReturnAllNotificationCofigs) | **Get** /notifications | Lists all notification configurations available in your environment
[**ValidateCreateConfiguration**](NotificationsApi.md#ValidateCreateConfiguration) | **Post** /notifications/validator | Validates the payload for the &#x60;POST /notifications&#x60; request
[**ValidateCreateOrUpdateConfiguration**](NotificationsApi.md#ValidateCreateOrUpdateConfiguration) | **Post** /notifications/{id}/validator | Validates the payload the &#x60;PUT /notifications/{id}&#x60; request



## CreateNotificationConfig

> CreateNotificationConfig(ctx, optional)

Creates a new notification configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNotificationConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNotificationConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**optional.Interface of NotificationConfig**](NotificationConfig.md)| The JSON body of the request. Contains parameters of the new notification configuration. | 

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


## CreateOrUpdateNotificationConfig

> NotificationConfigStub CreateOrUpdateNotificationConfig(ctx, id, optional)

Updates an existing notification configuration or creates a new one

If a notification configuration with the specified ID doesn't exist, a new configuration is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the notification configuration to be updated.    If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateNotificationConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateNotificationConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**optional.Interface of NotificationConfig**](NotificationConfig.md)| The JSON body of the request. Contains updated parameters of the notification configuration. | 

### Return type

[**NotificationConfigStub**](NotificationConfigStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificatonConfig

> DeleteNotificatonConfig(ctx, id)

Deletes the specified notification configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the notification configuration to be deleted. | 

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


## GetNotificationConfig

> NotificationConfig GetNotificationConfig(ctx, id)

Gets the properties of the specified notification configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required notification configuration. | 

### Return type

[**NotificationConfig**](NotificationConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllNotificationCofigs

> NotificationConfigStubListDto ReturnAllNotificationCofigs(ctx, )

Lists all notification configurations available in your environment

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**NotificationConfigStubListDto**](NotificationConfigStubListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateConfiguration

> ValidateCreateConfiguration(ctx, optional)

Validates the payload for the `POST /notifications` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateCreateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**optional.Interface of NotificationConfig**](NotificationConfig.md)| The JSON body of the request. Contains the notification configuration to be validated. | 

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


## ValidateCreateOrUpdateConfiguration

> ValidateCreateOrUpdateConfiguration(ctx, id, optional)

Validates the payload the `PUT /notifications/{id}` request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the notification configuration to be validated. | 
 **optional** | ***ValidateCreateOrUpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateOrUpdateConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**optional.Interface of NotificationConfig**](NotificationConfig.md)| The JSON body of the request. Contains the notification configuration to be validated. | 

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

