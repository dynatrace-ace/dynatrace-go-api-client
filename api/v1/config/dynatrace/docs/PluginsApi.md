# \PluginsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemotePluginEndpoint**](PluginsApi.md#CreateRemotePluginEndpoint) | **Post** /plugins/{id}/endpoints | Creates a new endpoint for the specified ActiveGate plugin
[**DeletePlugin**](PluginsApi.md#DeletePlugin) | **Delete** /plugins/{id}/binary | Deletes the ZIP file of the specified plugin
[**DeleteRemotePluginEndpoint**](PluginsApi.md#DeleteRemotePluginEndpoint) | **Delete** /plugins/{id}/endpoints/{endpointId} | Deletes an existing endpoint of the ActiveGate plugin
[**GetPlugin**](PluginsApi.md#GetPlugin) | **Get** /plugins/{id} | Lists the properties of the specified plugin
[**GetPluginBinary**](PluginsApi.md#GetPluginBinary) | **Get** /plugins/{id}/binary | Downloads the ZIP file of the specified plugin
[**GetPluginStates**](PluginsApi.md#GetPluginStates) | **Get** /plugins/{id}/states | Lists the states of the specified plugin
[**GetPlugins**](PluginsApi.md#GetPlugins) | **Get** /plugins | Lists all uploaded plugins
[**GetRemotePluginEndpoint**](PluginsApi.md#GetRemotePluginEndpoint) | **Get** /plugins/{id}/endpoints/{endpointId} | Gets parameters of the specified endpoint of the ActiveGate plugin
[**GetRemotePluginEndpoints**](PluginsApi.md#GetRemotePluginEndpoints) | **Get** /plugins/{id}/endpoints | Lists endpoints of the specified ActiveGate plugin
[**GetRemotePluginModules**](PluginsApi.md#GetRemotePluginModules) | **Get** /plugins/activeGatePluginModules | List available ActiveGate plugin modules
[**UpdateRemotePluginEndpoint**](PluginsApi.md#UpdateRemotePluginEndpoint) | **Put** /plugins/{id}/endpoints/{endpointId} | Updates an existing endpoint of the ActiveGate plugin
[**UploadPlugin**](PluginsApi.md#UploadPlugin) | **Post** /plugins | Uploads a ZIP plugin file
[**ValidatePlugin**](PluginsApi.md#ValidatePlugin) | **Post** /plugins/validator | Validates a ZIP plugin file for &#x60;POST/plugins&#x60; request
[**ValidateRemotePluginEndpoint**](PluginsApi.md#ValidateRemotePluginEndpoint) | **Post** /plugins/{id}/endpoints/validator | Validates the payload for the &#x60;POST /plugins/{id}/endpoints&#x60; request



## CreateRemotePluginEndpoint

> EntityShortRepresentation CreateRemotePluginEndpoint(ctx, id).RemotePluginEndpoint(remotePluginEndpoint).Execute()

Creates a new endpoint for the specified ActiveGate plugin



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
    id := "id_example" // string | The ID of the plugin where you want to create an endpoint.
    remotePluginEndpoint := *openapiclient.NewRemotePluginEndpoint(*openapiclient.NewEntityShortRepresentation("Id_example")) // RemotePluginEndpoint | The JSON body of the request. Contains parameters of the new plugin endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.CreateRemotePluginEndpoint(context.Background(), id).RemotePluginEndpoint(remotePluginEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.CreateRemotePluginEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemotePluginEndpoint`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.CreateRemotePluginEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the plugin where you want to create an endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemotePluginEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePluginEndpoint** | [**RemotePluginEndpoint**](RemotePluginEndpoint.md) | The JSON body of the request. Contains parameters of the new plugin endpoint. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlugin

> DeletePlugin(ctx, id).Execute()

Deletes the ZIP file of the specified plugin



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
    id := "id_example" // string | The ID of the plugin to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.DeletePlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.DeletePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the plugin to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginRequest struct via the builder pattern


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


## DeleteRemotePluginEndpoint

> DeleteRemotePluginEndpoint(ctx, id, endpointId).Execute()

Deletes an existing endpoint of the ActiveGate plugin

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
    id := "id_example" // string | The ID of the plugin where you want to delete an endpoint.
    endpointId := "endpointId_example" // string | The ID of the endpoint to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.DeleteRemotePluginEndpoint(context.Background(), id, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.DeleteRemotePluginEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the plugin where you want to delete an endpoint. | 
**endpointId** | **string** | The ID of the endpoint to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemotePluginEndpointRequest struct via the builder pattern


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


## GetPlugin

> Plugin GetPlugin(ctx, id).Execute()

Lists the properties of the specified plugin

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
    id := "id_example" // string | The ID of the required plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.GetPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlugin`: Plugin
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plugin**](Plugin.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginBinary

> map[string]interface{} GetPluginBinary(ctx, id).Execute()

Downloads the ZIP file of the specified plugin

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
    id := "id_example" // string | The ID of the plugin you want to download.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.GetPluginBinary(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetPluginBinary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginBinary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetPluginBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the plugin you want to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginStates

> PluginStateList GetPluginStates(ctx, id).Execute()

Lists the states of the specified plugin

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
    id := "id_example" // string | The ID of the required plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.GetPluginStates(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetPluginStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginStates`: PluginStateList
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetPluginStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginStateList**](PluginStateList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> StubList GetPlugins(ctx).Execute()

Lists all uploaded plugins

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
    resp, r, err := api_client.PluginsApi.GetPlugins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlugins`: StubList
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemotePluginEndpoint

> RemotePluginEndpoint GetRemotePluginEndpoint(ctx, id, endpointId).Execute()

Gets parameters of the specified endpoint of the ActiveGate plugin

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
    id := "id_example" // string | The ID of the required plugin.
    endpointId := "endpointId_example" // string | The ID of the required endpoint.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.GetRemotePluginEndpoint(context.Background(), id, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetRemotePluginEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemotePluginEndpoint`: RemotePluginEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetRemotePluginEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required plugin. | 
**endpointId** | **string** | The ID of the required endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemotePluginEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemotePluginEndpoint**](RemotePluginEndpoint.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemotePluginEndpoints

> StubList GetRemotePluginEndpoints(ctx, id).Execute()

Lists endpoints of the specified ActiveGate plugin

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
    id := "id_example" // string | The ID of the required plugin.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.GetRemotePluginEndpoints(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetRemotePluginEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemotePluginEndpoints`: StubList
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetRemotePluginEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemotePluginEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemotePluginModules

> StubList GetRemotePluginModules(ctx).Execute()

List available ActiveGate plugin modules

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
    resp, r, err := api_client.PluginsApi.GetRemotePluginModules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.GetRemotePluginModules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemotePluginModules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.GetRemotePluginModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemotePluginModulesRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemotePluginEndpoint

> UpdateRemotePluginEndpoint(ctx, id, endpointId).RemotePluginEndpoint(remotePluginEndpoint).Execute()

Updates an existing endpoint of the ActiveGate plugin

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
    id := "id_example" // string | The ID of the plugin where you want to update an endpoint.   If you set the plugin ID in the body as well, it must match this ID.
    endpointId := "endpointId_example" // string | The ID of the endpoint to be updated.   If you set the endpoint ID in the body as well, it must match this ID.
    remotePluginEndpoint := *openapiclient.NewRemotePluginEndpoint(*openapiclient.NewEntityShortRepresentation("Id_example")) // RemotePluginEndpoint | The JSON body of the request. Contains updated parameters of the plugin endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.UpdateRemotePluginEndpoint(context.Background(), id, endpointId).RemotePluginEndpoint(remotePluginEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.UpdateRemotePluginEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the plugin where you want to update an endpoint.   If you set the plugin ID in the body as well, it must match this ID. | 
**endpointId** | **string** | The ID of the endpoint to be updated.   If you set the endpoint ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemotePluginEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remotePluginEndpoint** | [**RemotePluginEndpoint**](RemotePluginEndpoint.md) | The JSON body of the request. Contains updated parameters of the plugin endpoint. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPlugin

> EntityShortRepresentation UploadPlugin(ctx).OverrideAlerts(overrideAlerts).InlineObject3(inlineObject3).Execute()

Uploads a ZIP plugin file

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
    overrideAlerts := true // bool | Use plugin-defined thresholds for alerts (`true`) or user-defined thresholds (`false`).    Plugin-defined thresholds are stored in the `plugin.json` file.   If not set, user-defined thresholds are used. (optional) (default to false)
    inlineObject3 := *openapiclient.Newinline_object_3("TODO") // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.UploadPlugin(context.Background()).OverrideAlerts(overrideAlerts).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.UploadPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadPlugin`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `PluginsApi.UploadPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overrideAlerts** | **bool** | Use plugin-defined thresholds for alerts (&#x60;true&#x60;) or user-defined thresholds (&#x60;false&#x60;).    Plugin-defined thresholds are stored in the &#x60;plugin.json&#x60; file.   If not set, user-defined thresholds are used. | [default to false]
 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePlugin

> ValidatePlugin(ctx).JsonOnly(jsonOnly).InlineObject2(inlineObject2).Execute()

Validates a ZIP plugin file for `POST/plugins` request

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
    jsonOnly := true // bool | Validate only the `plugin.json` file (`true`) or the entire plugin structure (`false`).    If not set, the entire plugin is validated.  (optional) (default to false)
    inlineObject2 := *openapiclient.Newinline_object_2("TODO") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.ValidatePlugin(context.Background()).JsonOnly(jsonOnly).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.ValidatePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonOnly** | **bool** | Validate only the &#x60;plugin.json&#x60; file (&#x60;true&#x60;) or the entire plugin structure (&#x60;false&#x60;).    If not set, the entire plugin is validated.  | [default to false]
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRemotePluginEndpoint

> ValidateRemotePluginEndpoint(ctx, id).RemotePluginEndpoint(remotePluginEndpoint).Execute()

Validates the payload for the `POST /plugins/{id}/endpoints` request

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
    id := "id_example" // string | The ID of the required plugin.
    remotePluginEndpoint := *openapiclient.NewRemotePluginEndpoint(*openapiclient.NewEntityShortRepresentation("Id_example")) // RemotePluginEndpoint | The JSON body of the request. Contains parameters of the new plugin endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PluginsApi.ValidateRemotePluginEndpoint(context.Background(), id).RemotePluginEndpoint(remotePluginEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginsApi.ValidateRemotePluginEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required plugin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRemotePluginEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePluginEndpoint** | [**RemotePluginEndpoint**](RemotePluginEndpoint.md) | The JSON body of the request. Contains parameters of the new plugin endpoint. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

