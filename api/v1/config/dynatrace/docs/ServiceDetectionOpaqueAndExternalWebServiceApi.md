# \ServiceDetectionOpaqueAndExternalWebServiceApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#CreateOpaqueAndExternalWebServiceRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE | Creates a new service detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#DeleteOpaqueAndExternalWebServiceRule) | **Delete** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id} | Deletes the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#GetOpaqueAndExternalWebServiceRule) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id} | Shows the properties of the specified service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ListOpaqueAndExternalWebServiceRules**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#ListOpaqueAndExternalWebServiceRules) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE | Lists all opaque and external web service detection rules | maturity&#x3D;EARLY_ADOPTER
[**OrderOpaqueAndExternalWebServiceRules**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#OrderOpaqueAndExternalWebServiceRules) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/order | Reorders the service detection rules of the specified type | maturity&#x3D;EARLY_ADOPTER
[**UpdateOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#UpdateOpaqueAndExternalWebServiceRule) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id} | Updates an existing service detection rule | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#ValidateCreateOpaqueAndExternalWebServiceRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateOpaqueAndExternalWebServiceRule**](ServiceDetectionOpaqueAndExternalWebServiceApi.md#ValidateUpdateOpaqueAndExternalWebServiceRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateOpaqueAndExternalWebServiceRule

> EntityShortRepresentation CreateOpaqueAndExternalWebServiceRule(ctx).Position(position).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()

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
    opaqueAndExternalWebServiceRule := *openapiclient.NewOpaqueAndExternalWebServiceRule("Type_example", "Name_example", false) // OpaqueAndExternalWebServiceRule | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.CreateOpaqueAndExternalWebServiceRule(context.Background()).Position(position).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.CreateOpaqueAndExternalWebServiceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOpaqueAndExternalWebServiceRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebServiceApi.CreateOpaqueAndExternalWebServiceRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to &quot;APPEND&quot;]
 **opaqueAndExternalWebServiceRule** | [**OpaqueAndExternalWebServiceRule**](OpaqueAndExternalWebServiceRule.md) | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE/reorder&#x60; request. | 

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


## DeleteOpaqueAndExternalWebServiceRule

> DeleteOpaqueAndExternalWebServiceRule(ctx, id).Execute()

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
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.DeleteOpaqueAndExternalWebServiceRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.DeleteOpaqueAndExternalWebServiceRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


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


## GetOpaqueAndExternalWebServiceRule

> OpaqueAndExternalWebServiceRule GetOpaqueAndExternalWebServiceRule(ctx, id).Execute()

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
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.GetOpaqueAndExternalWebServiceRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.GetOpaqueAndExternalWebServiceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpaqueAndExternalWebServiceRule`: OpaqueAndExternalWebServiceRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebServiceApi.GetOpaqueAndExternalWebServiceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpaqueAndExternalWebServiceRule**](OpaqueAndExternalWebServiceRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOpaqueAndExternalWebServiceRules

> StubList ListOpaqueAndExternalWebServiceRules(ctx).Execute()

Lists all opaque and external web service detection rules | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.ListOpaqueAndExternalWebServiceRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.ListOpaqueAndExternalWebServiceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOpaqueAndExternalWebServiceRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebServiceApi.ListOpaqueAndExternalWebServiceRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOpaqueAndExternalWebServiceRulesRequest struct via the builder pattern


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


## OrderOpaqueAndExternalWebServiceRules

> OrderOpaqueAndExternalWebServiceRules(ctx).StubList(stubList).Execute()

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
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.OrderOpaqueAndExternalWebServiceRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.OrderOpaqueAndExternalWebServiceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderOpaqueAndExternalWebServiceRulesRequest struct via the builder pattern


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


## UpdateOpaqueAndExternalWebServiceRule

> EntityShortRepresentation UpdateOpaqueAndExternalWebServiceRule(ctx, id).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()

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
    opaqueAndExternalWebServiceRule := *openapiclient.NewOpaqueAndExternalWebServiceRule("Type_example", "Name_example", false) // OpaqueAndExternalWebServiceRule | The JSON body of the request containing updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.UpdateOpaqueAndExternalWebServiceRule(context.Background(), id).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.UpdateOpaqueAndExternalWebServiceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOpaqueAndExternalWebServiceRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebServiceApi.UpdateOpaqueAndExternalWebServiceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaqueAndExternalWebServiceRule** | [**OpaqueAndExternalWebServiceRule**](OpaqueAndExternalWebServiceRule.md) | The JSON body of the request containing updated parameters of the service detection rule. | 

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


## ValidateCreateOpaqueAndExternalWebServiceRule

> ValidateCreateOpaqueAndExternalWebServiceRule(ctx).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE` request | maturity=EARLY_ADOPTER

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
    opaqueAndExternalWebServiceRule := *openapiclient.NewOpaqueAndExternalWebServiceRule("Type_example", "Name_example", false) // OpaqueAndExternalWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.ValidateCreateOpaqueAndExternalWebServiceRule(context.Background()).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.ValidateCreateOpaqueAndExternalWebServiceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaqueAndExternalWebServiceRule** | [**OpaqueAndExternalWebServiceRule**](OpaqueAndExternalWebServiceRule.md) |  | 

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


## ValidateUpdateOpaqueAndExternalWebServiceRule

> ValidateUpdateOpaqueAndExternalWebServiceRule(ctx, id).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()

Validate the payload for the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_SERVICE/{id}` request | maturity=EARLY_ADOPTER

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
    opaqueAndExternalWebServiceRule := *openapiclient.NewOpaqueAndExternalWebServiceRule("Type_example", "Name_example", false) // OpaqueAndExternalWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceDetectionOpaqueAndExternalWebServiceApi.ValidateUpdateOpaqueAndExternalWebServiceRule(context.Background(), id).OpaqueAndExternalWebServiceRule(opaqueAndExternalWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebServiceApi.ValidateUpdateOpaqueAndExternalWebServiceRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiValidateUpdateOpaqueAndExternalWebServiceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaqueAndExternalWebServiceRule** | [**OpaqueAndExternalWebServiceRule**](OpaqueAndExternalWebServiceRule.md) |  | 

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

