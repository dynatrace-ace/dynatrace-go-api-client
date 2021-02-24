# \MobileDeobfuscationAndSymbolicationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdate**](MobileDeobfuscationAndSymbolicationApi.md#CreateOrUpdate) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Upload a symbolication file. Either a ProGuard file for Android or a zip file containing all the iOS dSYM files. 
[**CreateOrUpdatePinning**](MobileDeobfuscationAndSymbolicationApi.md#CreateOrUpdatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning | Pin or unpin all symbolication files of a app version. A pinned file will not be deleted automtically, when running out of space.
[**DeleteApp**](MobileDeobfuscationAndSymbolicationApi.md#DeleteApp) | **Delete** /symfiles/{applicationId} | Deletes all symbolication file belonging to the Dynatrace App specified
[**DeleteSingleFile**](MobileDeobfuscationAndSymbolicationApi.md#DeleteSingleFile) | **Delete** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Delete the symbolication file belonging to the given application, os and version
[**GetAll**](MobileDeobfuscationAndSymbolicationApi.md#GetAll) | **Get** /symfiles | Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) from the Symbol File Store.
[**GetAllPerApplication**](MobileDeobfuscationAndSymbolicationApi.md#GetAllPerApplication) | **Get** /symfiles/{applicationId} | Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) for one single mobile application from the Symbol File Store of this tenant.
[**GetDssClientUrl**](MobileDeobfuscationAndSymbolicationApi.md#GetDssClientUrl) | **Get** /symfiles/dtxdss-download | Gets a download link for the latest version of the DTXDSSClient.
[**GetInfo**](MobileDeobfuscationAndSymbolicationApi.md#GetInfo) | **Get** /symfiles/info | Retrieves information about used/empty diskspace, number of stored files,....
[**GetSingle**](MobileDeobfuscationAndSymbolicationApi.md#GetSingle) | **Get** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Gets the metadata of the symbolication file belonging to the specified app version. There always can exist only one file per os and version
[**GetSupportedVersion**](MobileDeobfuscationAndSymbolicationApi.md#GetSupportedVersion) | **Get** /symfiles/ios/supportedversion | Returns the supported file format version for iOS symbol files.
[**ValidatePinning**](MobileDeobfuscationAndSymbolicationApi.md#ValidatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning/validator | Validate updates of existing request attribute for the &#x60;PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning&#x60; request.



## CreateOrUpdate

> CreateOrUpdate(ctx, applicationId, packageName, os, versionCode, versionName).Body(body).ContentType(contentType).Execute()

Upload a symbolication file. Either a ProGuard file for Android or a zip file containing all the iOS dSYM files. 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace of the app this file belongs to
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the uploaded file
    os := "os_example" // string | The operating system the file belongs to
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) the file belongs to
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to
    body := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded. A proguard file (*.txt) for Android or the zip file produced by the DTXDSSClient provided with the OneAgent for iOS. 
    contentType := "contentType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.CreateOrUpdate(context.Background(), applicationId, packageName, os, versionCode, versionName).Body(body).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.CreateOrUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace of the app this file belongs to | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the uploaded file | 
**os** | **string** | The operating system the file belongs to | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | ***os.File** | The file to be uploaded. A proguard file (*.txt) for Android or the zip file produced by the DTXDSSClient provided with the OneAgent for iOS.  | 
 **contentType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/x-compressed, application/x-zip-compressed, application/zip, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdatePinning

> CreateOrUpdatePinning(ctx, applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()

Pin or unpin all symbolication files of a app version. A pinned file will not be deleted automtically, when running out of space.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace of the app which should be (un)pinned
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned
    os := "os_example" // string | The operating system for which the files will be (un)pinned
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) the file belongs to
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to
    symbolFilePinning := *openapiclient.NewSymbolFilePinning(false) // SymbolFilePinning | JSON body of the request, containing the pinning status to set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.CreateOrUpdatePinning(context.Background(), applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.CreateOrUpdatePinning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace of the app which should be (un)pinned | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned | 
**os** | **string** | The operating system for which the files will be (un)pinned | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdatePinningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**SymbolFilePinning**](SymbolFilePinning.md) | JSON body of the request, containing the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, applicationId).Execute()

Deletes all symbolication file belonging to the Dynatrace App specified

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace of the mobile application the files will be deleted

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.DeleteApp(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace of the mobile application the files will be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSingleFile

> DeleteSingleFile(ctx, applicationId, packageName, os, versionCode, versionName).Execute()

Delete the symbolication file belonging to the given application, os and version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace of the mobile application where a file should be deleted
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the file to be deleted
    os := "os_example" // string | The operating system the file belongs to
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the file to be deleted
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the file to be deleted

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.DeleteSingleFile(context.Background(), applicationId, packageName, os, versionCode, versionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.DeleteSingleFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace of the mobile application where a file should be deleted | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the file to be deleted | 
**os** | **string** | The operating system the file belongs to | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the file to be deleted | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the file to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSingleFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> SymbolFileList GetAll(ctx).Execute()

Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) from the Symbol File Store.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: SymbolFileList
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPerApplication

> SymbolFileList GetAllPerApplication(ctx, applicationId).Execute()

Lists the metadata of all symbolication files (ProGuard files for Android or dSYM files for iOS Apps) for one single mobile application from the Symbol File Store of this tenant.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace for the mobile application to be queried

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetAllPerApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetAllPerApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPerApplication`: SymbolFileList
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetAllPerApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace for the mobile application to be queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPerApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDssClientUrl

> SymbolFileClientLinkDto GetDssClientUrl(ctx).Execute()

Gets a download link for the latest version of the DTXDSSClient.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetDssClientUrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetDssClientUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDssClientUrl`: SymbolFileClientLinkDto
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetDssClientUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDssClientUrlRequest struct via the builder pattern


### Return type

[**SymbolFileClientLinkDto**](SymbolFileClientLinkDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> SymbolFileStorageInfo GetInfo(ctx).Execute()

Retrieves information about used/empty diskspace, number of stored files,....

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfo`: SymbolFileStorageInfo
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern


### Return type

[**SymbolFileStorageInfo**](SymbolFileStorageInfo.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingle

> SymbolFile GetSingle(ctx, applicationId, packageName, os, versionCode, versionName).Execute()

Gets the metadata of the symbolication file belonging to the specified app version. There always can exist only one file per os and version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace for the mobile application to be queried
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app to be queried
    os := "os_example" // string | The operating system for which the file should be displayed.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the file to be retrieved
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the file to be retrieved

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetSingle(context.Background(), applicationId, packageName, os, versionCode, versionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingle`: SymbolFile
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace for the mobile application to be queried | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app to be queried | 
**os** | **string** | The operating system for which the file should be displayed. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the file to be retrieved | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the file to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**SymbolFile**](SymbolFile.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedVersion

> SupportedVersion GetSupportedVersion(ctx).Execute()

Returns the supported file format version for iOS symbol files.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.GetSupportedVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.GetSupportedVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedVersion`: SupportedVersion
    fmt.Fprintf(os.Stdout, "Response from `MobileDeobfuscationAndSymbolicationApi.GetSupportedVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedVersionRequest struct via the builder pattern


### Return type

[**SupportedVersion**](SupportedVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePinning

> ValidatePinning(ctx, applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()

Validate updates of existing request attribute for the `PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning` request.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := TODO // string | The application id used in Dynatrace of the app which should be (un)pinned
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned
    os := "os_example" // string | The operating system for which the files will be (un)pinned
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) the file belongs to
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to
    symbolFilePinning := *openapiclient.NewSymbolFilePinning(false) // SymbolFilePinning | JSON body of the request, containing the pinning status to set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MobileDeobfuscationAndSymbolicationApi.ValidatePinning(context.Background(), applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeobfuscationAndSymbolicationApi.ValidatePinning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | The application id used in Dynatrace of the app which should be (un)pinned | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned | 
**os** | **string** | The operating system for which the files will be (un)pinned | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) the file belongs to | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidatePinningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**SymbolFilePinning**](SymbolFilePinning.md) | JSON body of the request, containing the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

