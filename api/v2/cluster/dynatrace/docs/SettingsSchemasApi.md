# \SettingsSchemasApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableSchemaDefinitions**](SettingsSchemasApi.md#GetAvailableSchemaDefinitions) | **Get** /settings/schemas | Lists available settings schemas | maturity&#x3D;EARLY_ADOPTER
[**GetSchemaDefinition**](SettingsSchemasApi.md#GetSchemaDefinition) | **Get** /settings/schemas/{schemaId} | Gets parameters of the specified settings schema | maturity&#x3D;EARLY_ADOPTER



## GetAvailableSchemaDefinitions

> SchemaList GetAvailableSchemaDefinitions(ctx).Execute()

Lists available settings schemas | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.SettingsSchemasApi.GetAvailableSchemaDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsSchemasApi.GetAvailableSchemaDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableSchemaDefinitions`: SchemaList
    fmt.Fprintf(os.Stdout, "Response from `SettingsSchemasApi.GetAvailableSchemaDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSchemaDefinitionsRequest struct via the builder pattern


### Return type

[**SchemaList**](SchemaList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaDefinition

> SchemaDefinitionRestDto GetSchemaDefinition(ctx, schemaId).SchemaVersion(schemaVersion).Execute()

Gets parameters of the specified settings schema | maturity=EARLY_ADOPTER

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
    schemaId := "schemaId_example" // string | The ID of the required schema.
    schemaVersion := "schemaVersion_example" // string | The version of the required schema.    If not set, the most recent version is returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsSchemasApi.GetSchemaDefinition(context.Background(), schemaId).SchemaVersion(schemaVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsSchemasApi.GetSchemaDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaDefinition`: SchemaDefinitionRestDto
    fmt.Fprintf(os.Stdout, "Response from `SettingsSchemasApi.GetSchemaDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | The ID of the required schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaVersion** | **string** | The version of the required schema.    If not set, the most recent version is returned. | 

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

