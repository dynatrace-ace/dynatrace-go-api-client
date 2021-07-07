# \DeploymentApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveGateInstallerAvailableVersions**](DeploymentApi.md#GetActiveGateInstallerAvailableVersions) | **Get** /deployment/installer/gateway/versions/{osType} | Lists all available versions of ActiveGate installer



## GetActiveGateInstallerAvailableVersions

> ActiveGateInstallerVersions GetActiveGateInstallerAvailableVersions(ctx, osType).Execute()

Lists all available versions of ActiveGate installer

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
    osType := "osType_example" // string | The operating system of the installer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.GetActiveGateInstallerAvailableVersions(context.Background(), osType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetActiveGateInstallerAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGateInstallerAvailableVersions`: ActiveGateInstallerVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetActiveGateInstallerAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGateInstallerAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveGateInstallerVersions**](ActiveGateInstallerVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

