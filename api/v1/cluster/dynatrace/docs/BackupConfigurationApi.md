# \BackupConfigurationApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeBackupConfig**](BackupConfigurationApi.md#ChangeBackupConfig) | **Put** /backup/config | Change backup configuration
[**CheckBackupDirForCluster**](BackupConfigurationApi.md#CheckBackupDirForCluster) | **Get** /backup/clusterCheckDir | Check if given directory is valid for backup in the cluster
[**GetBackupConfig**](BackupConfigurationApi.md#GetBackupConfig) | **Get** /backup/config | Return backup configuration overview
[**GetStatusOfChangeBackupConfig**](BackupConfigurationApi.md#GetStatusOfChangeBackupConfig) | **Get** /backup/config/status | Check status of change backup configuration



## ChangeBackupConfig

> ChangeBackupConfig(ctx).BackupConfigDto(backupConfigDto).Execute()

Change backup configuration

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
    backupConfigDto := *openapiclient.NewBackupConfigDto() // BackupConfigDto | The JSON body of the request, containing new state of backup configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupConfigurationApi.ChangeBackupConfig(context.Background()).BackupConfigDto(backupConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigurationApi.ChangeBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeBackupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupConfigDto** | [**BackupConfigDto**](BackupConfigDto.md) | The JSON body of the request, containing new state of backup configuration. | 

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


## CheckBackupDirForCluster

> StorageTestImpl CheckBackupDirForCluster(ctx).Dir(dir).Datacenter(datacenter).Execute()

Check if given directory is valid for backup in the cluster

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
    dir := "dir_example" // string | Backup store directory path parameter. Missing or empty values will return a 'Bad Request' (optional)
    datacenter := "datacenter_example" // string | Datacenter where backup should be executed. Should be omitted for Single Datacenter Cluster (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupConfigurationApi.CheckBackupDirForCluster(context.Background()).Dir(dir).Datacenter(datacenter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigurationApi.CheckBackupDirForCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckBackupDirForCluster`: StorageTestImpl
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigurationApi.CheckBackupDirForCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckBackupDirForClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dir** | **string** | Backup store directory path parameter. Missing or empty values will return a &#39;Bad Request&#39; | 
 **datacenter** | **string** | Datacenter where backup should be executed. Should be omitted for Single Datacenter Cluster | 

### Return type

[**StorageTestImpl**](StorageTestImpl.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupConfig

> BackupConfigDto GetBackupConfig(ctx).Execute()

Return backup configuration overview

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
    resp, r, err := api_client.BackupConfigurationApi.GetBackupConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigurationApi.GetBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupConfig`: BackupConfigDto
    fmt.Fprintf(os.Stdout, "Response from `BackupConfigurationApi.GetBackupConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupConfigRequest struct via the builder pattern


### Return type

[**BackupConfigDto**](BackupConfigDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusOfChangeBackupConfig

> GetStatusOfChangeBackupConfig(ctx).RequestId(requestId).Execute()

Check status of change backup configuration

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
    requestId := "requestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupConfigurationApi.GetStatusOfChangeBackupConfig(context.Background()).RequestId(requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupConfigurationApi.GetStatusOfChangeBackupConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusOfChangeBackupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** |  | 

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

