# \EnvironmentsApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /environments | Creates a new environment.
[**CreateOrUpdateEnvironment**](EnvironmentsApi.md#CreateOrUpdateEnvironment) | **Put** /environments/{id} | Updates an existing environment or creates a new one.
[**CreateTokenManagementToken**](EnvironmentsApi.md#CreateTokenManagementToken) | **Post** /environments/{id}/tokenManagementToken | Creates a new tenant management token for an environment.
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /environments/{id} | Deletes the specified environment. An environment must be disabled before it can be deleted.
[**GetAllEnvironments**](EnvironmentsApi.md#GetAllEnvironments) | **Get** /environments | Lists all existing environments.
[**GetSingleEnvironment**](EnvironmentsApi.md#GetSingleEnvironment) | **Get** /environments/{id} | Gets the properties of the specified environment.



## CreateEnvironment

> EnvironmentShortRepresentation CreateEnvironment(ctx).Environment(environment).CreateToken(createToken).Execute()

Creates a new environment.

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
    environment := *openapiclient.NewEnvironment("Name_example") // Environment | The JSON body of the request. The body must not provide an ID as it will be automatically assigned by the Dynatrace server.
    createToken := true // bool | If true, a token management token with the scopes 'apiTokens.read' and 'apiTokens.write' (for usage with token API v2) and 'TenantTokenManagement' (for usage with token API v1) is created when creating a new environment. This token is then returned in the response body. It can be used within the newly created environment to create other tokens for configuring this environment. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.CreateEnvironment(context.Background()).Environment(environment).CreateToken(createToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: EnvironmentShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**Environment**](Environment.md) | The JSON body of the request. The body must not provide an ID as it will be automatically assigned by the Dynatrace server. | 
 **createToken** | **bool** | If true, a token management token with the scopes &#39;apiTokens.read&#39; and &#39;apiTokens.write&#39; (for usage with token API v2) and &#39;TenantTokenManagement&#39; (for usage with token API v1) is created when creating a new environment. This token is then returned in the response body. It can be used within the newly created environment to create other tokens for configuring this environment. | [default to false]

### Return type

[**EnvironmentShortRepresentation**](EnvironmentShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateEnvironment

> EnvironmentShortRepresentation CreateOrUpdateEnvironment(ctx, id).CreateToken(createToken).Environment(environment).Execute()

Updates an existing environment or creates a new one.



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
    id := "id_example" // string | The ID of the environment to update.   If you set the ID in the body as well, it must match this ID.
    createToken := true // bool | If true, a token management token with the scopes 'apiTokens.read' and 'apiTokens.write' (for usage with token API v2) and 'TenantTokenManagement' (for usage with token API v1) is created when creating a new environment. This token is then returned in the response body. It can be used within the newly created environment to create other tokens for configuring this environment. (optional) (default to false)
    environment := *openapiclient.NewEnvironment("Name_example") // Environment | JSON body of the request, containing updated parameters of the environment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.CreateOrUpdateEnvironment(context.Background(), id).CreateToken(createToken).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateOrUpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateEnvironment`: EnvironmentShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateOrUpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the environment to update.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createToken** | **bool** | If true, a token management token with the scopes &#39;apiTokens.read&#39; and &#39;apiTokens.write&#39; (for usage with token API v2) and &#39;TenantTokenManagement&#39; (for usage with token API v1) is created when creating a new environment. This token is then returned in the response body. It can be used within the newly created environment to create other tokens for configuring this environment. | [default to false]
 **environment** | [**Environment**](Environment.md) | JSON body of the request, containing updated parameters of the environment. | 

### Return type

[**EnvironmentShortRepresentation**](EnvironmentShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTokenManagementToken

> Token CreateTokenManagementToken(ctx, id).CreateEnvironmentTokenManagementToken(createEnvironmentTokenManagementToken).Execute()

Creates a new tenant management token for an environment.



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
    id := "id_example" // string | The ID of the environment where the token is valid.
    createEnvironmentTokenManagementToken := *openapiclient.NewCreateEnvironmentTokenManagementToken("Name_example") // CreateEnvironmentTokenManagementToken | The JSON body of the request. Contains parameters of the token. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.CreateTokenManagementToken(context.Background(), id).CreateEnvironmentTokenManagementToken(createEnvironmentTokenManagementToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateTokenManagementToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTokenManagementToken`: Token
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateTokenManagementToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the environment where the token is valid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenManagementTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEnvironmentTokenManagementToken** | [**CreateEnvironmentTokenManagementToken**](CreateEnvironmentTokenManagementToken.md) | The JSON body of the request. Contains parameters of the token. | 

### Return type

[**Token**](Token.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, id).Execute()

Deletes the specified environment. An environment must be disabled before it can be deleted.

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
    id := "id_example" // string | The ID of the environment to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.DeleteEnvironment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the environment to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## GetAllEnvironments

> EnvironmentList GetAllEnvironments(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Sort(sort).Filter(filter).IncludeConsumptionInfo(includeConsumptionInfo).IncludeStorageInfo(includeStorageInfo).IncludeUncachedConsumptionInfo(includeUncachedConsumptionInfo).Execute()

Lists all existing environments.

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
    pageSize := int64(789) // int64 | The amount of environments in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. (optional)
    sort := "sort_example" // string | The sort order. Possible sorts orders are:  * 'name' (without quotes): Sort by name ascending.  * '-name' (without quotes): Sort by name descending.  * 'creationDate' (without quotes): Sort by creation date ascending.  * '-creationDate' (without quotes): Sort by name descending.  (optional)
    filter := "filter_example" // string | Specifies the filter of the query.   You can set one or several of the following criteria: * Name: `name(string)`. The name can be either a substring or the full name of an environment. Case-insensitive. * Trial: `trial(true)` or `trial(false)`. Only includes trial environments if true is specified or only non-trial environments if false is specified. * State: `state(ENABLED)` or `state(DISABLED)`. * Tag: `tag(string)`. To filter by multiple tags by applying OR logic use `tag(string1,string2)`. To filter by multiple tags by applying AND logic use `tag(string1),tag(string2)`. To set several criteria, separate them with a comma (`,`). Only results, matching **all** criteria, are included into response.  (optional)
    includeConsumptionInfo := true // bool | If true, consumption information (limits, usage) is returned for each environment.  Returned usage is typically up to 1 hour old. To obtain fresher data, you can use **includeUncachedConsumptionInfo** parameter instead. (optional)
    includeStorageInfo := true // bool | If true, storage information (limits, usage) is returned for each environment. (optional)
    includeUncachedConsumptionInfo := true // bool | If true, uncached consumption information (limits, usage) is returned for each environment.  Up to date consumption will be calculated. Calculation may be time consuming, so it's recommended to use **includeConsumptionInfo** parameter instead.  If both this parameter and **includeConsumptionInfo** are set, **includeUncachedConsumptionInfo** will take priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetAllEnvironments(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Sort(sort).Filter(filter).IncludeConsumptionInfo(includeConsumptionInfo).IncludeStorageInfo(includeStorageInfo).IncludeUncachedConsumptionInfo(includeUncachedConsumptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetAllEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllEnvironments`: EnvironmentList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetAllEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of environments in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. | 
 **sort** | **string** | The sort order. Possible sorts orders are:  * &#39;name&#39; (without quotes): Sort by name ascending.  * &#39;-name&#39; (without quotes): Sort by name descending.  * &#39;creationDate&#39; (without quotes): Sort by creation date ascending.  * &#39;-creationDate&#39; (without quotes): Sort by name descending.  | 
 **filter** | **string** | Specifies the filter of the query.   You can set one or several of the following criteria: * Name: &#x60;name(string)&#x60;. The name can be either a substring or the full name of an environment. Case-insensitive. * Trial: &#x60;trial(true)&#x60; or &#x60;trial(false)&#x60;. Only includes trial environments if true is specified or only non-trial environments if false is specified. * State: &#x60;state(ENABLED)&#x60; or &#x60;state(DISABLED)&#x60;. * Tag: &#x60;tag(string)&#x60;. To filter by multiple tags by applying OR logic use &#x60;tag(string1,string2)&#x60;. To filter by multiple tags by applying AND logic use &#x60;tag(string1),tag(string2)&#x60;. To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results, matching **all** criteria, are included into response.  | 
 **includeConsumptionInfo** | **bool** | If true, consumption information (limits, usage) is returned for each environment.  Returned usage is typically up to 1 hour old. To obtain fresher data, you can use **includeUncachedConsumptionInfo** parameter instead. | 
 **includeStorageInfo** | **bool** | If true, storage information (limits, usage) is returned for each environment. | 
 **includeUncachedConsumptionInfo** | **bool** | If true, uncached consumption information (limits, usage) is returned for each environment.  Up to date consumption will be calculated. Calculation may be time consuming, so it&#39;s recommended to use **includeConsumptionInfo** parameter instead.  If both this parameter and **includeConsumptionInfo** are set, **includeUncachedConsumptionInfo** will take priority. | 

### Return type

[**EnvironmentList**](EnvironmentList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleEnvironment

> Environment GetSingleEnvironment(ctx, id).IncludeConsumptionInfo(includeConsumptionInfo).IncludeStorageInfo(includeStorageInfo).IncludeUncachedConsumptionInfo(includeUncachedConsumptionInfo).Execute()

Gets the properties of the specified environment.

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
    id := "id_example" // string | The ID of the environment.
    includeConsumptionInfo := true // bool | If true, consumption information (limits, usage) is returned for this environment.  Returned usage is typically up to 1 hour old. To obtain fresher data, you can use **includeUncachedConsumptionInfo** parameter instead. (optional)
    includeStorageInfo := true // bool | If true, storage information (limits, usage) is returned for this environment. (optional)
    includeUncachedConsumptionInfo := true // bool | If true, uncached consumption information (limits, usage) is returned for this environment.  Up to date consumption will be calculated. If both this parameter and **includeConsumptionInfo** are set, **includeUncachedConsumptionInfo** will take priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetSingleEnvironment(context.Background(), id).IncludeConsumptionInfo(includeConsumptionInfo).IncludeStorageInfo(includeStorageInfo).IncludeUncachedConsumptionInfo(includeUncachedConsumptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetSingleEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetSingleEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeConsumptionInfo** | **bool** | If true, consumption information (limits, usage) is returned for this environment.  Returned usage is typically up to 1 hour old. To obtain fresher data, you can use **includeUncachedConsumptionInfo** parameter instead. | 
 **includeStorageInfo** | **bool** | If true, storage information (limits, usage) is returned for this environment. | 
 **includeUncachedConsumptionInfo** | **bool** | If true, uncached consumption information (limits, usage) is returned for this environment.  Up to date consumption will be calculated. If both this parameter and **includeConsumptionInfo** are set, **includeUncachedConsumptionInfo** will take priority. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

