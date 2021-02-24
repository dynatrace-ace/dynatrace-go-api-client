# \ServiceRequestAttributesApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestAttributesConfig**](ServiceRequestAttributesApi.md#CreateRequestAttributesConfig) | **Post** /service/requestAttributes | Creates a new request attribute
[**DeleteRequestAttributesConfig**](ServiceRequestAttributesApi.md#DeleteRequestAttributesConfig) | **Delete** /service/requestAttributes/{id} | Deletes the specified request attribute
[**GetRequestAttributesConfig**](ServiceRequestAttributesApi.md#GetRequestAttributesConfig) | **Get** /service/requestAttributes/{id} | Shows the properties of the specified request attribute
[**ListRequestAttributesConfigs**](ServiceRequestAttributesApi.md#ListRequestAttributesConfigs) | **Get** /service/requestAttributes | Lists all available request attributes
[**UpdateRequestAttributesConfig**](ServiceRequestAttributesApi.md#UpdateRequestAttributesConfig) | **Put** /service/requestAttributes/{id} | Updates an existing request attribute or creates a new one
[**ValidateCreateRequestAttributesConfig**](ServiceRequestAttributesApi.md#ValidateCreateRequestAttributesConfig) | **Post** /service/requestAttributes/validator | Validates new request attributes for the &#x60;POST /requestAttributes&#x60; request
[**ValidateUpdateRequestAttributesConfig**](ServiceRequestAttributesApi.md#ValidateUpdateRequestAttributesConfig) | **Post** /service/requestAttributes/{id}/validator | Validate updates of existing request attribute for the &#x60;PUT /requestAttributes/{id}&#x60; request



## CreateRequestAttributesConfig

> EntityShortRepresentation CreateRequestAttributesConfig(ctx).RequestAttribute(requestAttribute).Execute()

Creates a new request attribute



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
    requestAttribute := *openapiclient.NewRequestAttribute("Name_example", false, "DataType_example", []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "Normalization_example", "Aggregation_example", false, false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.CreateRequestAttributesConfig(context.Background()).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.CreateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestAttributesConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesApi.CreateRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## DeleteRequestAttributesConfig

> DeleteRequestAttributesConfig(ctx, id).Execute()

Deletes the specified request attribute

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
    id := TODO // string | The ID of the request attribute to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.DeleteRequestAttributesConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.DeleteRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request attribute to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestAttributesConfigRequest struct via the builder pattern


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


## GetRequestAttributesConfig

> RequestAttribute GetRequestAttributesConfig(ctx, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()

Shows the properties of the specified request attribute

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
    id := TODO // string | The ID of the required request attribute.
    includeProcessGroupReferences := true // bool | Flag to include process group references to the response.    Process Group group references aren't compatible across environments. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.GetRequestAttributesConfig(context.Background(), id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.GetRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestAttributesConfig`: RequestAttribute
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesApi.GetRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required request attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **bool** | Flag to include process group references to the response.    Process Group group references aren&#39;t compatible across environments. | [default to false]

### Return type

[**RequestAttribute**](RequestAttribute.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestAttributesConfigs

> StubList ListRequestAttributesConfigs(ctx).Execute()

Lists all available request attributes

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
    resp, r, err := api_client.ServiceRequestAttributesApi.ListRequestAttributesConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.ListRequestAttributesConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestAttributesConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesApi.ListRequestAttributesConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestAttributesConfigsRequest struct via the builder pattern


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


## UpdateRequestAttributesConfig

> EntityShortRepresentation UpdateRequestAttributesConfig(ctx, id).RequestAttribute(requestAttribute).Execute()

Updates an existing request attribute or creates a new one

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
    id := TODO // string | The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID.
    requestAttribute := *openapiclient.NewRequestAttribute("Name_example", false, "DataType_example", []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "Normalization_example", "Aggregation_example", false, false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.UpdateRequestAttributesConfig(context.Background(), id).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.UpdateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestAttributesConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesApi.UpdateRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## ValidateCreateRequestAttributesConfig

> ValidateCreateRequestAttributesConfig(ctx).RequestAttribute(requestAttribute).Execute()

Validates new request attributes for the `POST /requestAttributes` request

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
    requestAttribute := *openapiclient.NewRequestAttribute("Name_example", false, "DataType_example", []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "Normalization_example", "Aggregation_example", false, false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.ValidateCreateRequestAttributesConfig(context.Background()).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.ValidateCreateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## ValidateUpdateRequestAttributesConfig

> ValidateUpdateRequestAttributesConfig(ctx, id).RequestAttribute(requestAttribute).Execute()

Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request

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
    id := TODO // string | The ID of the request attribute to be validated.
    requestAttribute := *openapiclient.NewRequestAttribute("Name_example", false, "DataType_example", []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "Normalization_example", "Aggregation_example", false, false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestAttributesApi.ValidateUpdateRequestAttributesConfig(context.Background(), id).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesApi.ValidateUpdateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request attribute to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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

