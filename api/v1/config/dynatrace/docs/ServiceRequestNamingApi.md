# \ServiceRequestNamingApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestNaming**](ServiceRequestNamingApi.md#CreateRequestNaming) | **Post** /service/requestNaming | Creates a new request naming rule
[**DeleteRequestNaming**](ServiceRequestNamingApi.md#DeleteRequestNaming) | **Delete** /service/requestNaming/{id} | Deletes the specified request naming rule
[**Get**](ServiceRequestNamingApi.md#Get) | **Get** /service/resourceNaming | Lists the global service resource requests.
[**GetRequestNaming**](ServiceRequestNamingApi.md#GetRequestNaming) | **Get** /service/requestNaming/{id} | Gets the parameters of the specified request naming rule
[**ListRequestNaming**](ServiceRequestNamingApi.md#ListRequestNaming) | **Get** /service/requestNaming | Lists all request naming rules along with their parameters
[**OrderRequestNaming**](ServiceRequestNamingApi.md#OrderRequestNaming) | **Put** /service/requestNaming/order | Reorders the request namings
[**Put**](ServiceRequestNamingApi.md#Put) | **Put** /service/resourceNaming | Updates the global service resource requests.
[**UpdateRequestNaming**](ServiceRequestNamingApi.md#UpdateRequestNaming) | **Put** /service/requestNaming/{id} | Updates the specified request naming rule
[**Validate**](ServiceRequestNamingApi.md#Validate) | **Post** /service/resourceNaming/validator | Validates new resource requests settings for the &#x60;PUT /service/resourceRequest&#x60; request.
[**ValidateCreateRequestNaming**](ServiceRequestNamingApi.md#ValidateCreateRequestNaming) | **Post** /service/requestNaming/validator | Validates the new request naming rule for the &#x60;POST /requestNaming&#x60; request
[**ValidateUpdateRequestNaming**](ServiceRequestNamingApi.md#ValidateUpdateRequestNaming) | **Post** /service/requestNaming/{id}/validator | Validates the new request naming for the &#x60;PUT /requestNaming/{id}&#x60; request



## CreateRequestNaming

> EntityShortRepresentation CreateRequestNaming(ctx).Position(position).RequestNaming(requestNaming).Execute()

Creates a new request naming rule



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
    position := "position_example" // string | Order of the new request naming rule. Set to `PREPEND` to prepend it to the list, `APPEND` to append it. Defaults to `APPEND`. (optional) (default to "APPEND")
    requestNaming := *openapiclient.NewRequestNaming(false, "NamingPattern_example", []openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo(map[string]interface{}(123), false, "Type_example"))}) // RequestNaming | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.CreateRequestNaming(context.Background()).Position(position).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.CreateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestNaming`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingApi.CreateRequestNaming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | Order of the new request naming rule. Set to &#x60;PREPEND&#x60; to prepend it to the list, &#x60;APPEND&#x60; to append it. Defaults to &#x60;APPEND&#x60;. | [default to &quot;APPEND&quot;]
 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## DeleteRequestNaming

> DeleteRequestNaming(ctx, id).Execute()

Deletes the specified request naming rule

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
    id := TODO // string | The ID of the request naming rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.DeleteRequestNaming(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.DeleteRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request naming rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestNamingRequest struct via the builder pattern


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


## Get

> ResourceNaming Get(ctx).Execute()

Lists the global service resource requests.



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
    resp, r, err := api_client.ServiceRequestNamingApi.Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: ResourceNaming
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingApi.Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


### Return type

[**ResourceNaming**](ResourceNaming.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestNaming

> RequestNaming GetRequestNaming(ctx, id).Execute()

Gets the parameters of the specified request naming rule

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
    id := TODO // string | The ID of the request naming rule you're inquiring.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.GetRequestNaming(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.GetRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestNaming`: RequestNaming
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingApi.GetRequestNaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request naming rule you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestNaming**](RequestNaming.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestNaming

> StubList ListRequestNaming(ctx).Execute()

Lists all request naming rules along with their parameters

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
    resp, r, err := api_client.ServiceRequestNamingApi.ListRequestNaming(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.ListRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestNaming`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingApi.ListRequestNaming`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestNamingRequest struct via the builder pattern


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


## OrderRequestNaming

> OrderRequestNaming(ctx).StubList(stubList).Execute()

Reorders the request namings



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
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.OrderRequestNaming(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.OrderRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**StubList**](StubList.md) | JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored. | 

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


## Put

> Put(ctx).ResourceNaming(resourceNaming).Execute()

Updates the global service resource requests.



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
    resourceNaming := *openapiclient.NewResourceNaming([]string{"Image_example"}, []string{"Binary_example"}) // ResourceNaming | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.Put(context.Background()).ResourceNaming(resourceNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.Put``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceNaming** | [**ResourceNaming**](ResourceNaming.md) |  | 

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


## UpdateRequestNaming

> EntityShortRepresentation UpdateRequestNaming(ctx, id).RequestNaming(requestNaming).Execute()

Updates the specified request naming rule



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
    id := TODO // string | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID.
    requestNaming := *openapiclient.NewRequestNaming(false, "NamingPattern_example", []openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo(map[string]interface{}(123), false, "Type_example"))}) // RequestNaming | The JSON body of the request containing updated parameters of the request naming. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.UpdateRequestNaming(context.Background(), id).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.UpdateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestNaming`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingApi.UpdateRequestNaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing updated parameters of the request naming. | 

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


## Validate

> Validate(ctx).ResourceNaming(resourceNaming).Execute()

Validates new resource requests settings for the `PUT /service/resourceRequest` request.

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
    resourceNaming := *openapiclient.NewResourceNaming([]string{"Image_example"}, []string{"Binary_example"}) // ResourceNaming | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.Validate(context.Background()).ResourceNaming(resourceNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceNaming** | [**ResourceNaming**](ResourceNaming.md) |  | 

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


## ValidateCreateRequestNaming

> ValidateCreateRequestNaming(ctx).RequestNaming(requestNaming).Execute()

Validates the new request naming rule for the `POST /requestNaming` request

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
    requestNaming := *openapiclient.NewRequestNaming(false, "NamingPattern_example", []openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo(map[string]interface{}(123), false, "Type_example"))}) // RequestNaming | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.ValidateCreateRequestNaming(context.Background()).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.ValidateCreateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## ValidateUpdateRequestNaming

> ValidateUpdateRequestNaming(ctx, id).RequestNaming(requestNaming).Execute()

Validates the new request naming for the `PUT /requestNaming/{id}` request



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
    id := TODO // string | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID.
    requestNaming := *openapiclient.NewRequestNaming(false, "NamingPattern_example", []openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo(map[string]interface{}(123), false, "Type_example"))}) // RequestNaming | The JSON body of the request containing updated parameters of the request naming. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceRequestNamingApi.ValidateUpdateRequestNaming(context.Background(), id).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingApi.ValidateUpdateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing updated parameters of the request naming. | 

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

