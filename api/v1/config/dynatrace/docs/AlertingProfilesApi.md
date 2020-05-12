# \AlertingProfilesApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertingProfile**](AlertingProfilesApi.md#CreateAlertingProfile) | **Post** /alertingProfiles | Creates a new alerting profile | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateAlertingProfile**](AlertingProfilesApi.md#CreateOrUpdateAlertingProfile) | **Put** /alertingProfiles/{id} | Updates an existing alerting profile | maturity&#x3D;EARLY_ADOPTER
[**DeleteAlertingProfile**](AlertingProfilesApi.md#DeleteAlertingProfile) | **Delete** /alertingProfiles/{id} | Deletes the specified alerting profile | maturity&#x3D;EARLY_ADOPTER
[**GetAlertingProfile**](AlertingProfilesApi.md#GetAlertingProfile) | **Get** /alertingProfiles/{id} | Gets the properties of the specified alerting profile | maturity&#x3D;EARLY_ADOPTER
[**GetAlertingProfiles**](AlertingProfilesApi.md#GetAlertingProfiles) | **Get** /alertingProfiles | Lists all alerting profiles available in your environment | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateAlertingProfile**](AlertingProfilesApi.md#ValidateCreateAlertingProfile) | **Post** /alertingProfiles/validator | Validates the payload the &#x60;POST /alertingProfiles&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateOrUpdateAlertingProfile**](AlertingProfilesApi.md#ValidateCreateOrUpdateAlertingProfile) | **Post** /alertingProfiles/{id}/validator | Validates the payload the &#x60;PUT /alertingProfiles/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateAlertingProfile

> EntityShortRepresentation CreateAlertingProfile(ctx, optional)

Creates a new alerting profile | maturity=EARLY_ADOPTER

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAlertingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAlertingProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfile** | [**optional.Interface of AlertingProfile**](AlertingProfile.md)| The JSON body of the request. Contains parameters of the new alerting profile. | 

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


## CreateOrUpdateAlertingProfile

> EntityShortRepresentation CreateOrUpdateAlertingProfile(ctx, id, optional)

Updates an existing alerting profile | maturity=EARLY_ADOPTER

If an alerting profile with the specified ID doesn't exist, a new one is created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the alerting profile to be updated. | 
 **optional** | ***CreateOrUpdateAlertingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateAlertingProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingProfile** | [**optional.Interface of AlertingProfile**](AlertingProfile.md)| The JSON body of the request. Contains updated parameters of the alerting profile. | 

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


## DeleteAlertingProfile

> DeleteAlertingProfile(ctx, id)

Deletes the specified alerting profile | maturity=EARLY_ADOPTER

The default alerting profile cannot be deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the alerting profile to be deleted. | 

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


## GetAlertingProfile

> AlertingProfile GetAlertingProfile(ctx, id)

Gets the properties of the specified alerting profile | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the required alerting profile. | 

### Return type

[**AlertingProfile**](AlertingProfile.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingProfiles

> StubList GetAlertingProfiles(ctx, )

Lists all alerting profiles available in your environment | maturity=EARLY_ADOPTER

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


## ValidateCreateAlertingProfile

> ValidateCreateAlertingProfile(ctx, optional)

Validates the payload the `POST /alertingProfiles` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateCreateAlertingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateAlertingProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfile** | [**optional.Interface of AlertingProfile**](AlertingProfile.md)| The JSON body of the request. Contains parameters of the alerting profile to be validated. | 

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


## ValidateCreateOrUpdateAlertingProfile

> ValidateCreateOrUpdateAlertingProfile(ctx, id, optional)

Validates the payload the `PUT /alertingProfiles/{id}` request | maturity=EARLY_ADOPTER

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The ID of the alerting profile to be validated. | 
 **optional** | ***ValidateCreateOrUpdateAlertingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCreateOrUpdateAlertingProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingProfile** | [**optional.Interface of AlertingProfile**](AlertingProfile.md)| The JSON body of the request. Contains parameters of the alerting profile to be validated. | 

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

