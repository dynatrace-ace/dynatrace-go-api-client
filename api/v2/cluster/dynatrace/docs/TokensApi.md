# \TokensApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensApi.md#CreateToken) | **Post** /tokens | Creates a new token
[**DeleteToken**](TokensApi.md#DeleteToken) | **Delete** /tokens/{id} | Deletes the specified token
[**GetTokenMetadata**](TokensApi.md#GetTokenMetadata) | **Get** /tokens/{id} | Lists token metadata by token ID
[**GetTokenMetadataBySecret**](TokensApi.md#GetTokenMetadataBySecret) | **Post** /tokens/lookup | Lists token metadata by token itself
[**ListTokens**](TokensApi.md#ListTokens) | **Get** /tokens | Lists available tokens in your environment
[**UpdateToken**](TokensApi.md#UpdateToken) | **Put** /tokens/{id} | Updates the specified token



## CreateToken

> Token CreateToken(ctx).CreateToken(createToken).Execute()

Creates a new token



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
    createToken := *openapiclient.NewCreateToken("Name_example", []string{"Scopes_example"}) // CreateToken | The JSON body of the request. Contains parameters of the new token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.CreateToken(context.Background()).CreateToken(createToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createToken** | [**CreateToken**](CreateToken.md) | The JSON body of the request. Contains parameters of the new token. | 

### Return type

[**Token**](Token.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8, text/csv; header=present; charset=utf-8, text/csv; header=absent; charset=utf-8, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, id).Execute()

Deletes the specified token

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
    id := "id_example" // string | The ID of the token to be deleted. Can either be the public identifier or the secret.   You can't delete the token you're using for authentication of the request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.DeleteToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.DeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the token to be deleted. Can either be the public identifier or the secret.   You can&#39;t delete the token you&#39;re using for authentication of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


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


## GetTokenMetadata

> TokenMetadata GetTokenMetadata(ctx, id).Execute()

Lists token metadata by token ID



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
    id := "id_example" // string | The ID of the required token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.GetTokenMetadata(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.GetTokenMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenMetadata`: TokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.GetTokenMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenMetadata**](TokenMetadata.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenMetadataBySecret

> TokenMetadata GetTokenMetadataBySecret(ctx).Token(token).Execute()

Lists token metadata by token itself

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
    token := *openapiclient.NewToken("Token_example") // Token | The JSON body of the request. Contains the required token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.GetTokenMetadataBySecret(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.GetTokenMetadataBySecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenMetadataBySecret`: TokenMetadata
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.GetTokenMetadataBySecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenMetadataBySecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | [**Token**](Token.md) | The JSON body of the request. Contains the required token. | 

### Return type

[**TokenMetadata**](TokenMetadata.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> StubList ListTokens(ctx).Limit(limit).User(user).Permissions(permissions).From(from).To(to).Execute()

Lists available tokens in your environment



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
    limit := int32(56) // int32 | Limits the maximum number of returned tokens.    If not set the value of `1000` is used.    Maximum value is 1000000. (optional) (default to 1000)
    user := "user_example" // string | Filters the resulting set of tokens by user, who owns the token. (optional)
    permissions := []string{"Permissions_example"} // []string | Filters the resulting set of tokens by scopes assigned to the token.    You can specify several permissions in the following format: `permissions=scope1&permissions=scope2`. The token must have *all* the specified scopes. (optional)
    from := int64(789) // int64 | Last used after this timestamp (UTC milliseconds). (optional)
    to := int64(789) // int64 | Last used before this timestamp (UTC milliseconds). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.ListTokens(context.Background()).Limit(limit).User(user).Permissions(permissions).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: StubList
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limits the maximum number of returned tokens.    If not set the value of &#x60;1000&#x60; is used.    Maximum value is 1000000. | [default to 1000]
 **user** | **string** | Filters the resulting set of tokens by user, who owns the token. | 
 **permissions** | **[]string** | Filters the resulting set of tokens by scopes assigned to the token.    You can specify several permissions in the following format: &#x60;permissions&#x3D;scope1&amp;permissions&#x3D;scope2&#x60;. The token must have *all* the specified scopes. | 
 **from** | **int64** | Last used after this timestamp (UTC milliseconds). | 
 **to** | **int64** | Last used before this timestamp (UTC milliseconds). | 

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


## UpdateToken

> UpdateToken(ctx, id).UpdateToken(updateToken).Execute()

Updates the specified token

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
    id := "id_example" // string | The ID of the token to be updated.    You can't update the token you're using for authentication of the request.
    updateToken := *openapiclient.NewUpdateToken() // UpdateToken | The JSON body of the request. Contains updated parameters of the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.UpdateToken(context.Background(), id).UpdateToken(updateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.UpdateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the token to be updated.    You can&#39;t update the token you&#39;re using for authentication of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateToken** | [**UpdateToken**](UpdateToken.md) | The JSON body of the request. Contains updated parameters of the token. | 

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

