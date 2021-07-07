# \UpdatesApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditInstallationBatch**](UpdatesApi.md#EditInstallationBatch) | **Put** /upgradeManagement/installationFiles/{type}/{version} | Modify update package state. Particular package might be excluded from download or re-downloaded.
[**GetClusterUpgradeStartupState**](UpdatesApi.md#GetClusterUpgradeStartupState) | **Get** /upgradeManagement/clusterUpgradeStartupState | Get system precondition check state for the cluster
[**GetInstallationFileList**](UpdatesApi.md#GetInstallationFileList) | **Get** /upgradeManagement/installationFiles | Get list of installation files and their cluster-wide availability information
[**RemoveInstallationBatch**](UpdatesApi.md#RemoveInstallationBatch) | **Delete** /upgradeManagement/installationFiles/{type}/{version} | Trigger removal of installation package
[**TriggerManualUpgrade**](UpdatesApi.md#TriggerManualUpgrade) | **Post** /upgradeManagement/triggerUpgrade | Trigger manual cluster upgrade



## EditInstallationBatch

> EditInstallationBatch(ctx, type_, version).UpdateEntryChangeRequestDto(updateEntryChangeRequestDto).Execute()

Modify update package state. Particular package might be excluded from download or re-downloaded.

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
    type_ := "type__example" // string | Installation package type, possible values = \"SERVER, AGENT, JS_AGENT, SG, NGINX\" 
    version := "version_example" // string | 
    updateEntryChangeRequestDto := *openapiclient.NewUpdateEntryChangeRequestDto() // UpdateEntryChangeRequestDto | The JSON body of the request. Contains parameters of update entry configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UpdatesApi.EditInstallationBatch(context.Background(), type_, version).UpdateEntryChangeRequestDto(updateEntryChangeRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatesApi.EditInstallationBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Installation package type, possible values &#x3D; \&quot;SERVER, AGENT, JS_AGENT, SG, NGINX\&quot;  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditInstallationBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEntryChangeRequestDto** | [**UpdateEntryChangeRequestDto**](UpdateEntryChangeRequestDto.md) | The JSON body of the request. Contains parameters of update entry configuration. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterUpgradeStartupState

> UpgradeStartupState GetClusterUpgradeStartupState(ctx).Execute()

Get system precondition check state for the cluster

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
    resp, r, err := api_client.UpdatesApi.GetClusterUpgradeStartupState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatesApi.GetClusterUpgradeStartupState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterUpgradeStartupState`: UpgradeStartupState
    fmt.Fprintf(os.Stdout, "Response from `UpdatesApi.GetClusterUpgradeStartupState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterUpgradeStartupStateRequest struct via the builder pattern


### Return type

[**UpgradeStartupState**](UpgradeStartupState.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallationFileList

> []InstallationPackage GetInstallationFileList(ctx).Execute()

Get list of installation files and their cluster-wide availability information

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
    resp, r, err := api_client.UpdatesApi.GetInstallationFileList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatesApi.GetInstallationFileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstallationFileList`: []InstallationPackage
    fmt.Fprintf(os.Stdout, "Response from `UpdatesApi.GetInstallationFileList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallationFileListRequest struct via the builder pattern


### Return type

[**[]InstallationPackage**](InstallationPackage.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInstallationBatch

> RemoveInstallationBatch(ctx, type_, version).Execute()

Trigger removal of installation package

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
    type_ := "type__example" // string | Installation package type, possible values = \"SERVER, AGENT, JS_AGENT, SG, NGINX\" 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UpdatesApi.RemoveInstallationBatch(context.Background(), type_, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatesApi.RemoveInstallationBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Installation package type, possible values &#x3D; \&quot;SERVER, AGENT, JS_AGENT, SG, NGINX\&quot;  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInstallationBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerManualUpgrade

> TriggerManualUpgrade(ctx).Execute()

Trigger manual cluster upgrade

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
    resp, r, err := api_client.UpdatesApi.TriggerManualUpgrade(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatesApi.TriggerManualUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerManualUpgradeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

