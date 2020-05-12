# \MobileDeobfuscationAndSymbolicationApi

All URIs are relative to *https://nph08633.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdate**](MobileDeobfuscationAndSymbolicationApi.md#CreateOrUpdate) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Upload a symbolication file. Either a ProGuard file for Android or a zip file containing all the iOS dSYM files. 
[**CreateOrUpdatePinning**](MobileDeobfuscationAndSymbolicationApi.md#CreateOrUpdatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning | Pin or unpin all symbolication files of a app version. A pinned file will not be deleted automtically, when running out of space.
[**DeleteApp**](MobileDeobfuscationAndSymbolicationApi.md#DeleteApp) | **Delete** /symfiles/{applicationId} | Deletes all symbolication file belonging to the Dynatrace App specified
[**DeleteSingleFile**](MobileDeobfuscationAndSymbolicationApi.md#DeleteSingleFile) | **Delete** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Delete the symbolication file belonging to the given application, os and version
[**GetAll**](MobileDeobfuscationAndSymbolicationApi.md#GetAll) | **Get** /symfiles | Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) from the Symbol File Store.
[**GetAllPerApplication**](MobileDeobfuscationAndSymbolicationApi.md#GetAllPerApplication) | **Get** /symfiles/{applicationId} | Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) for one single mobile application from the Symbol File Store of this tenant.
[**GetInfo**](MobileDeobfuscationAndSymbolicationApi.md#GetInfo) | **Get** /symfiles/info | Retrieves information about used/empty diskspace, number of stored files,....
[**GetSingle**](MobileDeobfuscationAndSymbolicationApi.md#GetSingle) | **Get** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Gets the metadata of the symbolication file belonging to the specified app version. There always can exist only one file per os and version
[**GetSupportedVersion**](MobileDeobfuscationAndSymbolicationApi.md#GetSupportedVersion) | **Get** /symfiles/ios/supportedversion | Returns the supported file format version for iOS symbol files.
[**ValidatePinning**](MobileDeobfuscationAndSymbolicationApi.md#ValidatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning/validator | Validate updates of existing request attribute for the &#x60;PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning&#x60; request.



## CreateOrUpdate

> CreateOrUpdate(ctx, applicationId, packageName, os, versionCode, versionName, body, optional)

Upload a symbolication file. Either a ProGuard file for Android or a zip file containing all the iOS dSYM files. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace of the app this file belongs to | 
**packageName** | **string**| The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the uploaded file | 
**os** | **string**| The operating system the file belongs to | 
**versionCode** | **string**| The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string**| The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 
**body** | ***os.File*****os.File**| The file to be uploaded. A proguard file (*.txt) for Android or the zip file produced by the DTXDSSClient provided with the OneAgent for iOS.  | 
 **optional** | ***CreateOrUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **contentType** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: application/x-compressed, application/x-zip-compressed, application/zip, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdatePinning

> CreateOrUpdatePinning(ctx, applicationId, packageName, os, versionCode, versionName, optional)

Pin or unpin all symbolication files of a app version. A pinned file will not be deleted automtically, when running out of space.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace of the app which should be (un)pinned | 
**packageName** | **string**| The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned | 
**os** | **string**| The operating system for which the files will be (un)pinned | 
**versionCode** | **string**| The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string**| The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 
 **optional** | ***CreateOrUpdatePinningOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdatePinningOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**optional.Interface of SymbolFilePinning**](SymbolFilePinning.md)| JSON body of the request, containing the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, applicationId)

Deletes all symbolication file belonging to the Dynatrace App specified

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace of the mobile application the files will be deleted | 

### Return type

 (empty response body)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSingleFile

> DeleteSingleFile(ctx, applicationId, packageName, os, versionCode, versionName)

Delete the symbolication file belonging to the given application, os and version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace of the mobile application where a file should be deleted | 
**packageName** | **string**| The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the file to be deleted | 
**os** | **string**| The operating system the file belongs to | 
**versionCode** | **string**| The version code (Android) / CFBundleVersion (iOS) of the file to be deleted | 
**versionName** | **string**| The version name (Android) / CFBundleShortVersionString (iOS) of the file to be deleted | 

### Return type

 (empty response body)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> SymbolFileList GetAll(ctx, )

Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) from the Symbol File Store.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPerApplication

> SymbolFileList GetAllPerApplication(ctx, applicationId)

Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) for one single mobile application from the Symbol File Store of this tenant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace for the mobile application to be queried | 

### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> SymbolFileStorageInfo GetInfo(ctx, )

Retrieves information about used/empty diskspace, number of stored files,....

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SymbolFileStorageInfo**](SymbolFileStorageInfo.md)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingle

> SymbolFile GetSingle(ctx, applicationId, packageName, os, versionCode, versionName)

Gets the metadata of the symbolication file belonging to the specified app version. There always can exist only one file per os and version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace for the mobile application to be queried | 
**packageName** | **string**| The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app to be queried | 
**os** | **string**| The operating system for which the file should be displayed. | 
**versionCode** | **string**| The version code (Android) / CFBundleVersion (iOS) of the file to be retrieved | 
**versionName** | **string**| The version name (Android) / CFBundleShortVersionString (iOS) of the file to be retrieved | 

### Return type

[**SymbolFile**](SymbolFile.md)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedVersion

> SupportedVersion GetSupportedVersion(ctx, )

Returns the supported file format version for iOS symbol files.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SupportedVersion**](SupportedVersion.md)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePinning

> ValidatePinning(ctx, applicationId, packageName, os, versionCode, versionName, optional)

Validate updates of existing request attribute for the `PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md)| The application id used in Dynatrace of the app which should be (un)pinned | 
**packageName** | **string**| The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned | 
**os** | **string**| The operating system for which the files will be (un)pinned | 
**versionCode** | **string**| The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string**| The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 
 **optional** | ***ValidatePinningOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePinningOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**optional.Interface of SymbolFilePinning**](SymbolFilePinning.md)| JSON body of the request, containing the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[DssFileManagement](../README.md#DssFileManagement)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

