# \ConditionalNamingApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamingRule**](ConditionalNamingApi.md#CreateNamingRule) | **Post** /conditionalNaming/{type} | Creates a new naming rule
[**DeleteNamingRule**](ConditionalNamingApi.md#DeleteNamingRule) | **Delete** /conditionalNaming/{type}/{id} | Deletes the specified naming rule
[**GetNamingRule**](ConditionalNamingApi.md#GetNamingRule) | **Get** /conditionalNaming/{type}/{id} | Lists the parameters of the specified naming rule
[**ListNamingRules**](ConditionalNamingApi.md#ListNamingRules) | **Get** /conditionalNaming/{type} | Lists all configured naming rules of the specified type
[**UpdateNamingRule**](ConditionalNamingApi.md#UpdateNamingRule) | **Put** /conditionalNaming/{type}/{id} | Updates the specified naming rule
[**ValidateCreateNamingRule**](ConditionalNamingApi.md#ValidateCreateNamingRule) | **Post** /conditionalNaming/{type}/validator | Validates the payload for the &#x60;POST /conditionalNaming/{type}&#x60; request
[**ValidateUpdateNamingRule**](ConditionalNamingApi.md#ValidateUpdateNamingRule) | **Post** /conditionalNaming/{type}/{id}/validator | Validates the payload for the &#x60;PUT /conditionalNaming/{type}/{id}&#x60; request



## CreateNamingRule

> EntityShortRepresentation CreateNamingRule(ctx, type_).ConditionalNamingRule(conditionalNamingRule).Execute()

Creates a new naming rule



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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    conditionalNamingRule := *openapiclient.NewConditionalNamingRule("PROCESS_GROUP", "Host with {Host:CpuCores} cores", "Host Naming Rule using Cores", false, []openapiclient.EntityRuleEngineCondition{*openapiclient.NewEntityRuleEngineCondition(*openapiclient.NewConditionKey("Attribute_example"), *openapiclient.NewComparisonBasic("Operator_example", false, "Type_example"))}) // ConditionalNamingRule | The JSON body of the request. Contains parameters of the new naming rule.   Do not specify the **id** of the rule.   The value of the **type** field must match the **type** path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.CreateNamingRule(context.Background(), type_).ConditionalNamingRule(conditionalNamingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.CreateNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNamingRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ConditionalNamingApi.CreateNamingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditionalNamingRule** | [**ConditionalNamingRule**](ConditionalNamingRule.md) | The JSON body of the request. Contains parameters of the new naming rule.   Do not specify the **id** of the rule.   The value of the **type** field must match the **type** path parameter. | 

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


## DeleteNamingRule

> DeleteNamingRule(ctx, type_, id).Execute()

Deletes the specified naming rule

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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    id := TODO // string | The ID of the naming rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.DeleteNamingRule(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.DeleteNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 
**id** | [**string**](.md) | The ID of the naming rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamingRuleRequest struct via the builder pattern


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


## GetNamingRule

> ConditionalNamingRule GetNamingRule(ctx, type_, id).Execute()

Lists the parameters of the specified naming rule

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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    id := TODO // string | The ID of the required naming rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.GetNamingRule(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.GetNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamingRule`: ConditionalNamingRule
    fmt.Fprintf(os.Stdout, "Response from `ConditionalNamingApi.GetNamingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 
**id** | [**string**](.md) | The ID of the required naming rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConditionalNamingRule**](ConditionalNamingRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamingRules

> StubList ListNamingRules(ctx, type_).Execute()

Lists all configured naming rules of the specified type

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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.ListNamingRules(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.ListNamingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNamingRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ConditionalNamingApi.ListNamingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamingRulesRequest struct via the builder pattern


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


## UpdateNamingRule

> EntityShortRepresentation UpdateNamingRule(ctx, type_, id).ConditionalNamingRule(conditionalNamingRule).Execute()

Updates the specified naming rule



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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    id := TODO // string | The ID of the naming rule to be updated.
    conditionalNamingRule := *openapiclient.NewConditionalNamingRule("PROCESS_GROUP", "Host with {Host:CpuCores} cores", "Host Naming Rule using Cores", false, []openapiclient.EntityRuleEngineCondition{*openapiclient.NewEntityRuleEngineCondition(*openapiclient.NewConditionKey("Attribute_example"), *openapiclient.NewComparisonBasic("Operator_example", false, "Type_example"))}) // ConditionalNamingRule | The JSON body of the request. Contains updated parameters of the naming rule.   The value of the **type** field must match the **type** path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.UpdateNamingRule(context.Background(), type_, id).ConditionalNamingRule(conditionalNamingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.UpdateNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNamingRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ConditionalNamingApi.UpdateNamingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 
**id** | [**string**](.md) | The ID of the naming rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditionalNamingRule** | [**ConditionalNamingRule**](ConditionalNamingRule.md) | The JSON body of the request. Contains updated parameters of the naming rule.   The value of the **type** field must match the **type** path parameter. | 

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


## ValidateCreateNamingRule

> ValidateCreateNamingRule(ctx, type_).ConditionalNamingRule(conditionalNamingRule).Execute()

Validates the payload for the `POST /conditionalNaming/{type}` request

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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    conditionalNamingRule := *openapiclient.NewConditionalNamingRule("PROCESS_GROUP", "Host with {Host:CpuCores} cores", "Host Naming Rule using Cores", false, []openapiclient.EntityRuleEngineCondition{*openapiclient.NewEntityRuleEngineCondition(*openapiclient.NewConditionKey("Attribute_example"), *openapiclient.NewComparisonBasic("Operator_example", false, "Type_example"))}) // ConditionalNamingRule | The JSON body of the request. Contains the naming rule parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.ValidateCreateNamingRule(context.Background(), type_).ConditionalNamingRule(conditionalNamingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.ValidateCreateNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateNamingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditionalNamingRule** | [**ConditionalNamingRule**](ConditionalNamingRule.md) | The JSON body of the request. Contains the naming rule parameters to be validated. | 

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


## ValidateUpdateNamingRule

> ValidateUpdateNamingRule(ctx, type_, id).ConditionalNamingRule(conditionalNamingRule).Execute()

Validates the payload for the `PUT /conditionalNaming/{type}/{id}` request

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
    type_ := "type__example" // string | The type of the rule, defined by the type of Dynatrace entities to which the rule applies.
    id := TODO // string | The ID of the naming rule to be validated.
    conditionalNamingRule := *openapiclient.NewConditionalNamingRule("PROCESS_GROUP", "Host with {Host:CpuCores} cores", "Host Naming Rule using Cores", false, []openapiclient.EntityRuleEngineCondition{*openapiclient.NewEntityRuleEngineCondition(*openapiclient.NewConditionKey("Attribute_example"), *openapiclient.NewComparisonBasic("Operator_example", false, "Type_example"))}) // ConditionalNamingRule | The JSON body of the request. Contains the naming rule parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConditionalNamingApi.ValidateUpdateNamingRule(context.Background(), type_, id).ConditionalNamingRule(conditionalNamingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConditionalNamingApi.ValidateUpdateNamingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of the rule, defined by the type of Dynatrace entities to which the rule applies. | 
**id** | [**string**](.md) | The ID of the naming rule to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateNamingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditionalNamingRule** | [**ConditionalNamingRule**](ConditionalNamingRule.md) | The JSON body of the request. Contains the naming rule parameters to be validated. | 

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

