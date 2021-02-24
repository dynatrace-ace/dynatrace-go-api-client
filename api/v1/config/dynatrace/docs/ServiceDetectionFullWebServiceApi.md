# \ServiceDetectionFullWebServiceApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#CreateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#DeleteServiceDetectionRule) | **Delete** /service/detectionRules/FULL_WEB_SERVICE/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#GetServiceDetectionRule) | **Get** /service/detectionRules/FULL_WEB_SERVICE/{id} | Shows the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListServiceDetectionRules**](ServiceDetectionFullWebServiceApi.md#ListServiceDetectionRules) | **Get** /service/detectionRules/FULL_WEB_SERVICE | Lists all full web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**OrderServiceDetectionRules**](ServiceDetectionFullWebServiceApi.md#OrderServiceDetectionRules) | **Put** /service/detectionRules/FULL_WEB_SERVICE/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**UpdateServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#UpdateServiceDetectionRule) | **Put** /service/detectionRules/FULL_WEB_SERVICE/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#ValidateCreateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_SERVICE&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateServiceDetectionRule**](ServiceDetectionFullWebServiceApi.md#ValidateUpdateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateServiceDetectionRule

> EntityShortRepresentation CreateServiceDetectionRule(ctx).Position(position).FullWebServiceRule(fullWebServiceRule).Execute()

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
    position := "position_example" // string | The position of the new rule:    * `APPEND`: at the end of the rule list.   * `PREPEND`: on top of the rule list.    (optional) (default to "APPEND")
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule("Type_example", "Name_example", false) // FullWebServiceRule | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.CreateServiceDetectionRule(context.Background()).Position(position).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.CreateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceApi.CreateServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to &quot;APPEND&quot;]
 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder&#x60; request. | 

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


## DeleteServiceDetectionRule

> DeleteServiceDetectionRule(ctx, id).Execute()

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
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.DeleteServiceDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.DeleteServiceDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServiceDetectionRuleRequest struct via the builder pattern


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


## GetServiceDetectionRule

> FullWebServiceRule GetServiceDetectionRule(ctx, id).Execute()

Shows the properties of the specified service detection rule | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.GetServiceDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.GetServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDetectionRule`: FullWebServiceRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceApi.GetServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullWebServiceRule**](FullWebServiceRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceDetectionRules

> StubList ListServiceDetectionRules(ctx).Execute()

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
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.ListServiceDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.ListServiceDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceDetectionRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceApi.ListServiceDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceDetectionRulesRequest struct via the builder pattern


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


## OrderServiceDetectionRules

> OrderServiceDetectionRules(ctx).StubList(stubList).Execute()

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
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.OrderServiceDetectionRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.OrderServiceDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderServiceDetectionRulesRequest struct via the builder pattern


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


## UpdateServiceDetectionRule

> EntityShortRepresentation UpdateServiceDetectionRule(ctx, id).FullWebServiceRule(fullWebServiceRule).Execute()

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
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule("Type_example", "Name_example", false) // FullWebServiceRule | The JSON body of the request containing updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.UpdateServiceDetectionRule(context.Background(), id).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.UpdateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceApi.UpdateServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) | The JSON body of the request containing updated parameters of the service detection rule. | 

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


## ValidateCreateServiceDetectionRule

> ValidateCreateServiceDetectionRule(ctx).FullWebServiceRule(fullWebServiceRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_SERVICE` request | maturity=EARLY_ADOPTER

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
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule("Type_example", "Name_example", false) // FullWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.ValidateCreateServiceDetectionRule(context.Background()).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.ValidateCreateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) |  | 

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


## ValidateUpdateServiceDetectionRule

> ValidateUpdateServiceDetectionRule(ctx, id).FullWebServiceRule(fullWebServiceRule).Execute()

Validate the payload for the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}` request | maturity=EARLY_ADOPTER

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
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule("Type_example", "Name_example", false) // FullWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionFullWebServiceApi.ValidateUpdateServiceDetectionRule(context.Background(), id).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceApi.ValidateUpdateServiceDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiValidateUpdateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) |  | 

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

