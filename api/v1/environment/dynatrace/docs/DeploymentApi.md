# \DeploymentApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadAgentInstallerWithVersion**](DeploymentApi.md#DownloadAgentInstallerWithVersion) | **Get** /deployment/installer/agent/{osType}/{installerType}/version/{version} | Downloads OneAgent installer of the specified version
[**DownloadBoshReleaseWithVersion**](DeploymentApi.md#DownloadBoshReleaseWithVersion) | **Get** /deployment/boshrelease/agent/{osType}/version/{version} | Downloads BOSH release tarballs of the specified version, OneAgent included
[**DownloadGatewayInstallerWithVersion**](DeploymentApi.md#DownloadGatewayInstallerWithVersion) | **Get** /deployment/installer/gateway/{osType}/version/{version} | Downloads the ActiveGate installer, compatible with the specified **OneAgent** version
[**DownloadLatestAgentInstaller**](DeploymentApi.md#DownloadLatestAgentInstaller) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest | Downloads the latest OneAgent installer
[**DownloadLatestGatewayInstaller**](DeploymentApi.md#DownloadLatestGatewayInstaller) | **Get** /deployment/installer/gateway/{osType}/latest | Downloads the configured standard ActiveGate installer of the latest version for the specified OS
[**GetActiveGateInstallerAvailableVersions**](DeploymentApi.md#GetActiveGateInstallerAvailableVersions) | **Get** /deployment/installer/gateway/versions/{osType} | Lists all available versions of ActiveGate installer | maturity&#x3D;EARLY_ADOPTER
[**GetAgentInstallerAvailableVersions**](DeploymentApi.md#GetAgentInstallerAvailableVersions) | **Get** /deployment/installer/agent/versions/{osType}/{installerType} | Lists all available versions of OneAgent installer
[**GetAgentInstallerConnectionInfo**](DeploymentApi.md#GetAgentInstallerConnectionInfo) | **Get** /deployment/installer/agent/connectioninfo | Gets the connectivity information for OneAgent
[**GetAgentInstallerConnectionInfoEndpoints**](DeploymentApi.md#GetAgentInstallerConnectionInfoEndpoints) | **Get** /deployment/installer/agent/connectioninfo/endpoints | Gets the list of the ActiveGate-Endpoints to be used for Agents ordered by networkzone-priorities. | maturity&#x3D;EARLY_ADOPTER
[**GetAgentInstallerMetaInfo**](DeploymentApi.md#GetAgentInstallerMetaInfo) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest/metainfo | Gets the latest available version of the OneAgent installer of the specified type
[**GetBoshReleaseAvailableVersions**](DeploymentApi.md#GetBoshReleaseAvailableVersions) | **Get** /deployment/boshrelease/versions/{osType} | Gets the list of available OneAgent versions for BOSH release tarballs
[**GetBoshReleaseChecksum**](DeploymentApi.md#GetBoshReleaseChecksum) | **Get** /deployment/boshrelease/agent/{osType}/version/{version}/checksum | Gets the checksum of the specified BOSH release tarbell



## DownloadAgentInstallerWithVersion

> DownloadAgentInstallerWithVersion(ctx, osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).Execute()

Downloads OneAgent installer of the specified version



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
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    version := "version_example" // string | The required version of the OneAgent in `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    flavor := "flavor_example" // string | The flavor of your Linux distribution.    Set `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    include := []string{"Include_example"} // []string | The code modules to be included to the installer. You can specify several modules in the following format: `include=java&include=dotnet`.   Only applicable to the `paas` and `paas-sh` installer types. (optional)
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.DownloadAgentInstallerWithVersion(context.Background(), osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DownloadAgentInstallerWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 
**version** | **string** | The required version of the OneAgent in &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAgentInstallerWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **flavor** | **string** | The flavor of your Linux distribution.    Set &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **include** | **[]string** | The code modules to be included to the installer. You can specify several modules in the following format: &#x60;include&#x3D;java&amp;include&#x3D;dotnet&#x60;.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | 
 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to false]

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


## DownloadBoshReleaseWithVersion

> DownloadBoshReleaseWithVersion(ctx, osType, version).SkipMetadata(skipMetadata).Execute()

Downloads BOSH release tarballs of the specified version, OneAgent included



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
    version := "version_example" // string | The required version of the OneAgent in the `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call.
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    If not set, `false` is used. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.DownloadBoshReleaseWithVersion(context.Background(), osType, version).SkipMetadata(skipMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DownloadBoshReleaseWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the OneAgent in the &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBoshReleaseWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    If not set, &#x60;false&#x60; is used. | [default to false]

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


## DownloadGatewayInstallerWithVersion

> DownloadGatewayInstallerWithVersion(ctx, osType, version).IfNoneMatch(ifNoneMatch).Execute()

Downloads the ActiveGate installer, compatible with the specified **OneAgent** version



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
    version := "version_example" // string | The required version of the **OneAgent** installer, in `1.155.275.20181112-084458` format.     The request returns the ActiveGate installer of the version, compatible with the provided OneAgent version.   You can retrieve the list of available versions with the [**GET available versions of ActiveGate**](https://dt-url.net/kh43rha) call.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.DownloadGatewayInstallerWithVersion(context.Background(), osType, version).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DownloadGatewayInstallerWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the **OneAgent** installer, in &#x60;1.155.275.20181112-084458&#x60; format.     The request returns the ActiveGate installer of the version, compatible with the provided OneAgent version.   You can retrieve the list of available versions with the [**GET available versions of ActiveGate**](https://dt-url.net/kh43rha) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGatewayInstallerWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 

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


## DownloadLatestAgentInstaller

> DownloadLatestAgentInstaller(ctx, osType, installerType).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).Execute()

Downloads the latest OneAgent installer



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
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    flavor := "flavor_example" // string | The flavor of your Linux distribution.    Set `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    include := []string{"Include_example"} // []string | The code modules to be included to the installer. You can specify several modules in the following format: `include=java&include=dotnet`.   Only applicable to the `paas` and `paas-sh` installer types. (optional)
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.DownloadLatestAgentInstaller(context.Background(), osType, installerType).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DownloadLatestAgentInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestAgentInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **flavor** | **string** | The flavor of your Linux distribution.    Set &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **include** | **[]string** | The code modules to be included to the installer. You can specify several modules in the following format: &#x60;include&#x3D;java&amp;include&#x3D;dotnet&#x60;.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | 
 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to false]

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


## DownloadLatestGatewayInstaller

> DownloadLatestGatewayInstaller(ctx, osType).IfNoneMatch(ifNoneMatch).Execute()

Downloads the configured standard ActiveGate installer of the latest version for the specified OS

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
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.DownloadLatestGatewayInstaller(context.Background(), osType).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.DownloadLatestGatewayInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestGatewayInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 

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


## GetActiveGateInstallerAvailableVersions

> ActiveGateInstallerVersions GetActiveGateInstallerAvailableVersions(ctx, osType).Execute()

Lists all available versions of ActiveGate installer | maturity=EARLY_ADOPTER

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


## GetAgentInstallerAvailableVersions

> AgentInstallerVersions GetAgentInstallerAvailableVersions(ctx, osType, installerType).Flavor(flavor).Arch(arch).Execute()

Lists all available versions of OneAgent installer

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
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    flavor := "flavor_example" // string | The flavor of your Linux distribution.    Set `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.GetAgentInstallerAvailableVersions(context.Background(), osType, installerType).Flavor(flavor).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetAgentInstallerAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerAvailableVersions`: AgentInstallerVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetAgentInstallerAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flavor** | **string** | The flavor of your Linux distribution.    Set &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]

### Return type

[**AgentInstallerVersions**](AgentInstallerVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerConnectionInfo

> ConnectionInfo GetAgentInstallerConnectionInfo(ctx).Execute()

Gets the connectivity information for OneAgent

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
    resp, r, err := api_client.DeploymentApi.GetAgentInstallerConnectionInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetAgentInstallerConnectionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerConnectionInfo`: ConnectionInfo
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetAgentInstallerConnectionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerConnectionInfoRequest struct via the builder pattern


### Return type

[**ConnectionInfo**](ConnectionInfo.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerConnectionInfoEndpoints

> GetAgentInstallerConnectionInfoEndpoints(ctx).NetworkZone(networkZone).Execute()

Gets the list of the ActiveGate-Endpoints to be used for Agents ordered by networkzone-priorities. | maturity=EARLY_ADOPTER



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
    networkZone := "networkZone_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.GetAgentInstallerConnectionInfoEndpoints(context.Background()).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetAgentInstallerConnectionInfoEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerConnectionInfoEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZone** | **string** |  | 

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


## GetAgentInstallerMetaInfo

> InstallerMetaInfoDto GetAgentInstallerMetaInfo(ctx, osType, installerType).Flavor(flavor).Arch(arch).Bitness(bitness).Execute()

Gets the latest available version of the OneAgent installer of the specified type



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
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    flavor := "flavor_example" // string | The flavor of your Linux distribution.    Set `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.GetAgentInstallerMetaInfo(context.Background(), osType, installerType).Flavor(flavor).Arch(arch).Bitness(bitness).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetAgentInstallerMetaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerMetaInfo`: InstallerMetaInfoDto
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetAgentInstallerMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flavor** | **string** | The flavor of your Linux distribution.    Set &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]

### Return type

[**InstallerMetaInfoDto**](InstallerMetaInfoDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoshReleaseAvailableVersions

> BoshReleaseAvailableVersions GetBoshReleaseAvailableVersions(ctx, osType).Execute()

Gets the list of available OneAgent versions for BOSH release tarballs

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
    resp, r, err := api_client.DeploymentApi.GetBoshReleaseAvailableVersions(context.Background(), osType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetBoshReleaseAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoshReleaseAvailableVersions`: BoshReleaseAvailableVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetBoshReleaseAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoshReleaseAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BoshReleaseAvailableVersions**](BoshReleaseAvailableVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoshReleaseChecksum

> BoshReleaseChecksum GetBoshReleaseChecksum(ctx, osType, version).SkipMetadata(skipMetadata).Execute()

Gets the checksum of the specified BOSH release tarbell



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
    version := "version_example" // string | The required version of the OneAgent in the `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call.
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    If not set, `false` is used. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentApi.GetBoshReleaseChecksum(context.Background(), osType, version).SkipMetadata(skipMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.GetBoshReleaseChecksum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoshReleaseChecksum`: BoshReleaseChecksum
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.GetBoshReleaseChecksum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the OneAgent in the &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoshReleaseChecksumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    If not set, &#x60;false&#x60; is used. | [default to false]

### Return type

[**BoshReleaseChecksum**](BoshReleaseChecksum.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

