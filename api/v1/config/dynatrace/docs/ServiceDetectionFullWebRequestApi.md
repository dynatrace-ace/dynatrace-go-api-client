# \ServiceDetectionFullWebRequestApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#CreateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#DeleteRequestDetectionRule) | **Delete** /service/detectionRules/FULL_WEB_REQUEST/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#GetRequestDetectionRule) | **Get** /service/detectionRules/FULL_WEB_REQUEST/{id} | Gets the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListRequestDetectionRules**](ServiceDetectionFullWebRequestApi.md#ListRequestDetectionRules) | **Get** /service/detectionRules/FULL_WEB_REQUEST | Lists all full web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**OrderRequestDetectionRules**](ServiceDetectionFullWebRequestApi.md#OrderRequestDetectionRules) | **Put** /service/detectionRules/FULL_WEB_REQUEST/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**UpdateRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#UpdateRequestDetectionRule) | **Put** /service/detectionRules/FULL_WEB_REQUEST/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#ValidateCreateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_REQUEST&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateRequestDetectionRule**](ServiceDetectionFullWebRequestApi.md#ValidateUpdateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST/{id}/validator | Validates the payload for the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateRequestDetectionRule

> EntityShortRepresentation CreateRequestDetectionRule(ctx).Position(position).FullWebRequestRule(fullWebRequestRule).Execute()

Creates a new service detection rule | maturity=EARLY_ADOPTER



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
    position := "position_example" // string | The position of the new rule:    * `APPEND`: at the bottom of the rule list.   * `PREPEND`: at the top of the rule list.    If not set, the `APPEND` is used. (optional) (default to "APPEND")
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule("Type_example", "Name_example", false) // FullWebRequestRule | The JSON body of the request. Contains parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order, use the `PUT /service/detectionRules/FULL_WEB_REQUEST/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.CreateRequestDetectionRule(context.Background()).Position(position).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.CreateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestApi.CreateRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used. | [default to &quot;APPEND&quot;]
 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order, use the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/reorder&#x60; request. | 

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


## DeleteRequestDetectionRule

> DeleteRequestDetectionRule(ctx, id).Execute()

Deletes the specified service detection rule | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the service detection rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.DeleteRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.DeleteRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the service detection rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestDetectionRuleRequest struct via the builder pattern


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


## GetRequestDetectionRule

> FullWebRequestRule GetRequestDetectionRule(ctx, id).Execute()

Gets the properties of the specified service detection rule | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the required service detection rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.GetRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.GetRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestDetectionRule`: FullWebRequestRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestApi.GetRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullWebRequestRule**](FullWebRequestRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestDetectionRules

> StubList ListRequestDetectionRules(ctx).Execute()

Lists all full web service detection rules | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.ListRequestDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.ListRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestDetectionRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestApi.ListRequestDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestDetectionRulesRequest struct via the builder pattern


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


## OrderRequestDetectionRules

> OrderRequestDetectionRules(ctx).StubList(stubList).Execute()

Reorders the service detection rules of the specified type | maturity=EARLY_ADOPTER



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
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | The JSON body of the request containing the service detection rules in the required order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.OrderRequestDetectionRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.OrderRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRequestDetectionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**StubList**](StubList.md) | The JSON body of the request containing the service detection rules in the required order. | 

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


## UpdateRequestDetectionRule

> EntityShortRepresentation UpdateRequestDetectionRule(ctx, id).FullWebRequestRule(fullWebRequestRule).Execute()

Updates an existing service detection rule | maturity=EARLY_ADOPTER



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
    id := TODO // string | The ID of the rule to be updated.
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule("Type_example", "Name_example", false) // FullWebRequestRule | The JSON body of the request. Contains updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.UpdateRequestDetectionRule(context.Background(), id).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.UpdateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestApi.UpdateRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains updated parameters of the service detection rule. | 

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


## ValidateCreateRequestDetectionRule

> ValidateCreateRequestDetectionRule(ctx).FullWebRequestRule(fullWebRequestRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_REQUEST` request | maturity=EARLY_ADOPTER

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
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule("Type_example", "Name_example", false) // FullWebRequestRule | The JSON body of the request. Contains parameters of the service detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.ValidateCreateRequestDetectionRule(context.Background()).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.ValidateCreateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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


## ValidateUpdateRequestDetectionRule

> ValidateUpdateRequestDetectionRule(ctx, id).FullWebRequestRule(fullWebRequestRule).Execute()

Validates the payload for the `PUT /service/detectionRules/FULL_WEB_REQUEST/{id}` request | maturity=EARLY_ADOPTER

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
    id := TODO // string | The ID of the service detection rule to be validated.
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule("Type_example", "Name_example", false) // FullWebRequestRule | The JSON body of the request. Contains parameters of the service detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebRequestApi.ValidateUpdateRequestDetectionRule(context.Background(), id).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestApi.ValidateUpdateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the service detection rule to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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

