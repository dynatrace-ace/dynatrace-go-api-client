# \Extensions20Api

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateExtensionEnvironmentConfiguration**](Extensions20Api.md#ActivateExtensionEnvironmentConfiguration) | **Post** /extensions/{extensionName}/environmentConfiguration | Activates the environment configuration from the specified version of the extension 2.0
[**CreateMonitoringConfiguration**](Extensions20Api.md#CreateMonitoringConfiguration) | **Post** /extensions/{extensionName}/monitoringConfigurations | Creates new monitoring configuration for the specified extension 2.0
[**DeleteEnvironmentConfiguration**](Extensions20Api.md#DeleteEnvironmentConfiguration) | **Delete** /extensions/{extensionName}/environmentConfiguration | Deactivates the environment configuration of the specified extension 2.0
[**ExtensionConfigurationSchema**](Extensions20Api.md#ExtensionConfigurationSchema) | **Get** /extensions/{extensionName}/{extensionVersion}/schema | Gets the configuration schema of the specified version of the extension 2.0
[**ExtensionDetails**](Extensions20Api.md#ExtensionDetails) | **Get** /extensions/{extensionName}/{extensionVersion} | Gets details of the specified version of the extension 2.0
[**ExtensionMonitoringConfigurations**](Extensions20Api.md#ExtensionMonitoringConfigurations) | **Get** /extensions/{extensionName}/monitoringConfigurations | Lists all the monitoring configurations of the specified extension 2.0
[**GetActiveEnvironmentConfiguration**](Extensions20Api.md#GetActiveEnvironmentConfiguration) | **Get** /extensions/{extensionName}/environmentConfiguration | Gets the active environment configuration version of the specified extension 2.0
[**GetEnvironmentConfigurationEvents**](Extensions20Api.md#GetEnvironmentConfigurationEvents) | **Get** /extensions/{extensionName}/environmentConfiguration/events | List of the latest extension environment configuration events
[**GetExtensionMonitoringConfigurationEvents**](Extensions20Api.md#GetExtensionMonitoringConfigurationEvents) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId}/events | Gets the list of the events linked to specific monitoring configuration
[**GetExtensionMonitoringConfigurationStatus**](Extensions20Api.md#GetExtensionMonitoringConfigurationStatus) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId}/status | Gets the most recent status of the execution of given monitoring configuration
[**GetSchemaFile**](Extensions20Api.md#GetSchemaFile) | **Get** /extensions/schemas/{schemaVersion}/{fileName} | Gets the extension 2.0 schema file in the specified version
[**ListExtensionVersions**](Extensions20Api.md#ListExtensionVersions) | **Get** /extensions/{extensionName} | Lists all versions of the extension 2.0
[**ListExtensions**](Extensions20Api.md#ListExtensions) | **Get** /extensions | Lists all the extensions 2.0 available in your environment
[**ListSchemaFiles**](Extensions20Api.md#ListSchemaFiles) | **Get** /extensions/schemas/{schemaVersion} | Lists all the files available for the specified extension 2.0 schema version
[**ListSchemas**](Extensions20Api.md#ListSchemas) | **Get** /extensions/schemas | Lists all the extension 2.0 schemas versions available in your environment
[**MonitoringConfigurationDetails**](Extensions20Api.md#MonitoringConfigurationDetails) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Gets the details of the specified monitoring configuration
[**RemoveExtension**](Extensions20Api.md#RemoveExtension) | **Delete** /extensions/{extensionName}/{extensionVersion} | Deletes the specified version of the extension 2.0
[**RemoveMonitoringConfiguration**](Extensions20Api.md#RemoveMonitoringConfiguration) | **Delete** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Deletes the specified monitoring configuration
[**UpdateExtensionEnvironmentConfiguration**](Extensions20Api.md#UpdateExtensionEnvironmentConfiguration) | **Put** /extensions/{extensionName}/environmentConfiguration | Updates the active environment configuration version of the extension 2.0
[**UpdateMonitoringConfiguration**](Extensions20Api.md#UpdateMonitoringConfiguration) | **Put** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Updates the specified monitoring configuration
[**UploadExtension**](Extensions20Api.md#UploadExtension) | **Post** /extensions | Uploads or verifies a new extension 2.0



## ActivateExtensionEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion ActivateExtensionEnvironmentConfiguration(ctx, extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()

Activates the environment configuration from the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionEnvironmentConfigurationVersion := *openapiclient.NewExtensionEnvironmentConfigurationVersion("1.2.3") // ExtensionEnvironmentConfigurationVersion | The version of the requested environment configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ActivateExtensionEnvironmentConfiguration(context.Background(), extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ActivateExtensionEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateExtensionEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ActivateExtensionEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateExtensionEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionEnvironmentConfigurationVersion** | [**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md) | The version of the requested environment configuration. | 

### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringConfiguration

> []MonitoringConfigurationResponse CreateMonitoringConfiguration(ctx, extensionName).MonitoringConfigurationDto(monitoringConfigurationDto).Execute()

Creates new monitoring configuration for the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    monitoringConfigurationDto := []openapiclient.MonitoringConfigurationDto{*openapiclient.NewMonitoringConfigurationDto("HOST-D3A3C5A146830A79")} // []MonitoringConfigurationDto | JSON body of the request, containing monitoring configuration parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.CreateMonitoringConfiguration(context.Background(), extensionName).MonitoringConfigurationDto(monitoringConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.CreateMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringConfiguration`: []MonitoringConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.CreateMonitoringConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfigurationDto** | [**[]MonitoringConfigurationDto**](MonitoringConfigurationDto.md) | JSON body of the request, containing monitoring configuration parameters. | 

### Return type

[**[]MonitoringConfigurationResponse**](MonitoringConfigurationResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion DeleteEnvironmentConfiguration(ctx, extensionName).Execute()

Deactivates the environment configuration of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.DeleteEnvironmentConfiguration(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.DeleteEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.DeleteEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionConfigurationSchema

> SchemaDefinitionRestDto ExtensionConfigurationSchema(ctx, extensionName, extensionVersion).Execute()

Gets the configuration schema of the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ExtensionConfigurationSchema(context.Background(), extensionName, extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ExtensionConfigurationSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionConfigurationSchema`: SchemaDefinitionRestDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ExtensionConfigurationSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionConfigurationSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SchemaDefinitionRestDto**](SchemaDefinitionRestDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionDetails

> Extension ExtensionDetails(ctx, extensionName, extensionVersion).Accept(accept).Execute()

Gets details of the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0
    accept := "accept_example" // string | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset=utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ExtensionDetails(context.Background(), extensionName, extensionVersion).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ExtensionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionDetails`: Extension
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ExtensionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset&#x3D;utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. | 

### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionMonitoringConfigurations

> ExtensionMonitoringConfigurationsList ExtensionMonitoringConfigurations(ctx, extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Version(version).Active(active).Execute()

Lists all the monitoring configurations of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    version := "version_example" // string | Filters the resulting set of configurations by extension 2.0 version. (optional)
    active := true // bool | Filters the resulting set of configurations by the active state. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ExtensionMonitoringConfigurations(context.Background(), extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Version(version).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ExtensionMonitoringConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionMonitoringConfigurations`: ExtensionMonitoringConfigurationsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ExtensionMonitoringConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionMonitoringConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **version** | **string** | Filters the resulting set of configurations by extension 2.0 version. | 
 **active** | **bool** | Filters the resulting set of configurations by the active state. | 

### Return type

[**ExtensionMonitoringConfigurationsList**](ExtensionMonitoringConfigurationsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion GetActiveEnvironmentConfiguration(ctx, extensionName).Execute()

Gets the active environment configuration version of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.GetActiveEnvironmentConfiguration(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.GetActiveEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.GetActiveEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentConfigurationEvents

> ExtensionEventsList GetEnvironmentConfigurationEvents(ctx, extensionName).Execute()

List of the latest extension environment configuration events

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.GetEnvironmentConfigurationEvents(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.GetEnvironmentConfigurationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentConfigurationEvents`: ExtensionEventsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.GetEnvironmentConfigurationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentConfigurationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionEventsList**](ExtensionEventsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionMonitoringConfigurationEvents

> ExtensionEventsList GetExtensionMonitoringConfigurationEvents(ctx, extensionName, configurationId).Execute()

Gets the list of the events linked to specific monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.GetExtensionMonitoringConfigurationEvents(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.GetExtensionMonitoringConfigurationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionMonitoringConfigurationEvents`: ExtensionEventsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.GetExtensionMonitoringConfigurationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionMonitoringConfigurationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionEventsList**](ExtensionEventsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionMonitoringConfigurationStatus

> ExtensionStatusDto GetExtensionMonitoringConfigurationStatus(ctx, extensionName, configurationId).Execute()

Gets the most recent status of the execution of given monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.GetExtensionMonitoringConfigurationStatus(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.GetExtensionMonitoringConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionMonitoringConfigurationStatus`: ExtensionStatusDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.GetExtensionMonitoringConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionMonitoringConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionStatusDto**](ExtensionStatusDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaFile

> map[string]interface{} GetSchemaFile(ctx, schemaVersion, fileName).Execute()

Gets the extension 2.0 schema file in the specified version

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
    schemaVersion := "schemaVersion_example" // string | The version of the schema.
    fileName := "fileName_example" // string | The name of the schema file.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.GetSchemaFile(context.Background(), schemaVersion, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.GetSchemaFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.GetSchemaFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaVersion** | **string** | The version of the schema. | 
**fileName** | **string** | The name of the schema file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionVersions

> ExtensionList ListExtensionVersions(ctx, extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all versions of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ListExtensionVersions(context.Background(), extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ListExtensionVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtensionVersions`: ExtensionList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ListExtensionVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 

### Return type

[**ExtensionList**](ExtensionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensions

> ExtensionList ListExtensions(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Name(name).Execute()

Lists all the extensions 2.0 available in your environment

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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    name := "name_example" // string | Filters the resulting set of extensions 2.0 by name. You can specify a partial name. In that case, the CONTAINS operator is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ListExtensions(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtensions`: ExtensionList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ListExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **name** | **string** | Filters the resulting set of extensions 2.0 by name. You can specify a partial name. In that case, the CONTAINS operator is used. | 

### Return type

[**ExtensionList**](ExtensionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemaFiles

> SchemaFiles ListSchemaFiles(ctx, schemaVersion).Execute()

Lists all the files available for the specified extension 2.0 schema version

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
    schemaVersion := "schemaVersion_example" // string | The version of the schema.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.ListSchemaFiles(context.Background(), schemaVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ListSchemaFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchemaFiles`: SchemaFiles
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ListSchemaFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaVersion** | **string** | The version of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemaFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SchemaFiles**](SchemaFiles.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemas

> SchemasList ListSchemas(ctx).Execute()

Lists all the extension 2.0 schemas versions available in your environment

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
    resp, r, err := api_client.Extensions20Api.ListSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.ListSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchemas`: SchemasList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.ListSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemasRequest struct via the builder pattern


### Return type

[**SchemasList**](SchemasList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringConfigurationDetails

> ExtensionMonitoringConfiguration MonitoringConfigurationDetails(ctx, extensionName, configurationId).Execute()

Gets the details of the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.MonitoringConfigurationDetails(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.MonitoringConfigurationDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringConfigurationDetails`: ExtensionMonitoringConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.MonitoringConfigurationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringConfigurationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionMonitoringConfiguration**](ExtensionMonitoringConfiguration.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExtension

> Extension RemoveExtension(ctx, extensionName, extensionVersion).Execute()

Deletes the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.RemoveExtension(context.Background(), extensionName, extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.RemoveExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExtension`: Extension
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.RemoveExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMonitoringConfiguration

> RemoveMonitoringConfiguration(ctx, extensionName, configurationId).Execute()

Deletes the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.RemoveMonitoringConfiguration(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.RemoveMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMonitoringConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtensionEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion UpdateExtensionEnvironmentConfiguration(ctx, extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()

Updates the active environment configuration version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionEnvironmentConfigurationVersion := *openapiclient.NewExtensionEnvironmentConfigurationVersion("1.2.3") // ExtensionEnvironmentConfigurationVersion | The version of the requested environment configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.UpdateExtensionEnvironmentConfiguration(context.Background(), extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.UpdateExtensionEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExtensionEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.UpdateExtensionEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionEnvironmentConfigurationVersion** | [**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md) | The version of the requested environment configuration. | 

### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringConfiguration

> MonitoringConfigurationResponse UpdateMonitoringConfiguration(ctx, extensionName, configurationId).MonitoringConfigurationUpdateDto(monitoringConfigurationUpdateDto).Execute()

Updates the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.
    monitoringConfigurationUpdateDto := *openapiclient.NewMonitoringConfigurationUpdateDto() // MonitoringConfigurationUpdateDto | JSON body of the request, containing monitoring configuration parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.UpdateMonitoringConfiguration(context.Background(), extensionName, configurationId).MonitoringConfigurationUpdateDto(monitoringConfigurationUpdateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.UpdateMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringConfiguration`: MonitoringConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.UpdateMonitoringConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitoringConfigurationUpdateDto** | [**MonitoringConfigurationUpdateDto**](MonitoringConfigurationUpdateDto.md) | JSON body of the request, containing monitoring configuration parameters. | 

### Return type

[**MonitoringConfigurationResponse**](MonitoringConfigurationResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadExtension

> Extension UploadExtension(ctx).ValidateOnly(validateOnly).Body(body).Execute()

Uploads or verifies a new extension 2.0

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
    validateOnly := true // bool | Only run validation but do not persist the extension even if validation was successful. (optional) (default to false)
    body := InlineObject1(987) // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Extensions20Api.UploadExtension(context.Background()).ValidateOnly(validateOnly).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20Api.UploadExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadExtension`: Extension
    fmt.Fprintf(os.Stdout, "Response from `Extensions20Api.UploadExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateOnly** | **bool** | Only run validation but do not persist the extension even if validation was successful. | [default to false]
 **body** | **InlineObject1** |  | 

### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

