# \AccessTokensAPITokensApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](AccessTokensAPITokensApi.md#CreateApiToken) | **Post** /apiTokens | Creates a new API token
[**DeleteApiToken**](AccessTokensAPITokensApi.md#DeleteApiToken) | **Delete** /apiTokens/{id} | Deletes an API token
[**GetApiToken**](AccessTokensAPITokensApi.md#GetApiToken) | **Get** /apiTokens/{id} | Gets API token metadata by token ID
[**ListApiTokens**](AccessTokensAPITokensApi.md#ListApiTokens) | **Get** /apiTokens | Lists all available API tokens
[**LookupApiToken**](AccessTokensAPITokensApi.md#LookupApiToken) | **Post** /apiTokens/lookup | Gets API token metadata by token secret
[**UpdateApiToken**](AccessTokensAPITokensApi.md#UpdateApiToken) | **Put** /apiTokens/{id} | Updates an API token



## CreateApiToken

> ApiTokenCreated CreateApiToken(ctx).ApiTokenCreate(apiTokenCreate).Execute()

Creates a new API token



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
    apiTokenCreate := *openapiclient.NewApiTokenCreate([]string{"metrics.read"}, "tokenName") // ApiTokenCreate | The JSON body of the request. Contains parameters of the new API token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.CreateApiToken(context.Background()).ApiTokenCreate(apiTokenCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.CreateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiToken`: ApiTokenCreated
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPITokensApi.CreateApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTokenCreate** | [**ApiTokenCreate**](ApiTokenCreate.md) | The JSON body of the request. Contains parameters of the new API token. | 

### Return type

[**ApiTokenCreated**](ApiTokenCreated.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiToken

> DeleteApiToken(ctx, id).Execute()

Deletes an API token

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
    id := "id_example" // string | The ID of the token to be deleted.   You can specify either the ID or the secret of the token.   You can't delete the token you're using for authentication of the request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.DeleteApiToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.DeleteApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the token to be deleted.   You can specify either the ID or the secret of the token.   You can&#39;t delete the token you&#39;re using for authentication of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


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


## GetApiToken

> ApiToken GetApiToken(ctx, id).Execute()

Gets API token metadata by token ID



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
    id := "id_example" // string | The ID of the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.GetApiToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.GetApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPITokensApi.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiTokens

> ApiTokenList ListApiTokens(ctx).NextPageKey(nextPageKey).PageSize(pageSize).ApiTokenSelector(apiTokenSelector).Fields(fields).From(from).To(to).Sort(sort).Execute()

Lists all available API tokens



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
    pageSize := int64(789) // int64 | The amount of API tokens in a single response payload.   The maximal allowed page size is 10000 and the minimal allowed page size is 100.   If not set, 200 is used. (optional)
    apiTokenSelector := "apiTokenSelector_example" // string | Filters the resulting sets of tokens. Only tokens matching the specified criteria are included into response.   You can set one or more of the following criteria:   * Owner: `owner(\"value\")`. The user that owns the token. Case-sensitive. * Personal access token: `personalAccessToken(false)`. Set to `true` to include only personal access tokens or to `false` to include only API tokens. * Token scope: `scope(\"scope1\",\"scope2\")`. If several values are specified, the **OR** logic applies.   To set multiple criteria, separate them with commas (`,`). Only results matching **all** criteria are included into response. (optional)
    fields := "fields_example" // string | Specifies the fields to be included in the response.  The following fields are included by default:   * `id`  * `name`  * `enabled`  * `owner`  * `creationDate`   To remove fields from the response, specify them with the minus (`-`) operator as a comma-separated list (for example, `-creationDate,-owner`).   You can include additional fields:   * `personalAccessToken`  * `expirationDate`  * `lastUsedDate`  * `lastUsedIpAddress`  * `scopes`   To add fields to the response, specify them with the plus (`+`) operator as a comma-separated list (for example, `+expirationDate,+scopes`). You can combine adding and removing of fields (for example, `+scopes,-creationDate`).   Alternatively, you can define the desired set of fields in the response. Specify the required fields as a comma-separated list, without operators (for example, `creationDate,expirationDate,owner`). The ID is always included in the response.   The **fields** string must be URL-encoded. (optional)
    from := "from_example" // string | Filters tokens based on the last usage time.  The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years    (optional)
    to := "to_example" // string | Filters tokens based on the last usage time.  The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    sort := "sort_example" // string | The sort order of the token list.   You can sort by the following properties with a sign prefix for the sort order:    * `name`: token name (`+` a...z or `-` z...a)   * `lastUsedDate` last used (`+` never used tokens first `-` most recently used tokens first)   * `creationDate` (`+` oldest tokens first `-` newest tokens first)   * `expirationDate` (`+` tokens that expire soon first `-` unlimited tokens first)   If no prefix is set, + is used.   If not set, tokens are sorted by creation date with newest first. (optional) (default to "-creationDate")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.ListApiTokens(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).ApiTokenSelector(apiTokenSelector).Fields(fields).From(from).To(to).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.ListApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiTokens`: ApiTokenList
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPITokensApi.ListApiTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of API tokens in a single response payload.   The maximal allowed page size is 10000 and the minimal allowed page size is 100.   If not set, 200 is used. | 
 **apiTokenSelector** | **string** | Filters the resulting sets of tokens. Only tokens matching the specified criteria are included into response.   You can set one or more of the following criteria:   * Owner: &#x60;owner(\&quot;value\&quot;)&#x60;. The user that owns the token. Case-sensitive. * Personal access token: &#x60;personalAccessToken(false)&#x60;. Set to &#x60;true&#x60; to include only personal access tokens or to &#x60;false&#x60; to include only API tokens. * Token scope: &#x60;scope(\&quot;scope1\&quot;,\&quot;scope2\&quot;)&#x60;. If several values are specified, the **OR** logic applies.   To set multiple criteria, separate them with commas (&#x60;,&#x60;). Only results matching **all** criteria are included into response. | 
 **fields** | **string** | Specifies the fields to be included in the response.  The following fields are included by default:   * &#x60;id&#x60;  * &#x60;name&#x60;  * &#x60;enabled&#x60;  * &#x60;owner&#x60;  * &#x60;creationDate&#x60;   To remove fields from the response, specify them with the minus (&#x60;-&#x60;) operator as a comma-separated list (for example, &#x60;-creationDate,-owner&#x60;).   You can include additional fields:   * &#x60;personalAccessToken&#x60;  * &#x60;expirationDate&#x60;  * &#x60;lastUsedDate&#x60;  * &#x60;lastUsedIpAddress&#x60;  * &#x60;scopes&#x60;   To add fields to the response, specify them with the plus (&#x60;+&#x60;) operator as a comma-separated list (for example, &#x60;+expirationDate,+scopes&#x60;). You can combine adding and removing of fields (for example, &#x60;+scopes,-creationDate&#x60;).   Alternatively, you can define the desired set of fields in the response. Specify the required fields as a comma-separated list, without operators (for example, &#x60;creationDate,expirationDate,owner&#x60;). The ID is always included in the response.   The **fields** string must be URL-encoded. | 
 **from** | **string** | Filters tokens based on the last usage time.  The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years    | 
 **to** | **string** | Filters tokens based on the last usage time.  The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **sort** | **string** | The sort order of the token list.   You can sort by the following properties with a sign prefix for the sort order:    * &#x60;name&#x60;: token name (&#x60;+&#x60; a...z or &#x60;-&#x60; z...a)   * &#x60;lastUsedDate&#x60; last used (&#x60;+&#x60; never used tokens first &#x60;-&#x60; most recently used tokens first)   * &#x60;creationDate&#x60; (&#x60;+&#x60; oldest tokens first &#x60;-&#x60; newest tokens first)   * &#x60;expirationDate&#x60; (&#x60;+&#x60; tokens that expire soon first &#x60;-&#x60; unlimited tokens first)   If no prefix is set, + is used.   If not set, tokens are sorted by creation date with newest first. | [default to &quot;-creationDate&quot;]

### Return type

[**ApiTokenList**](ApiTokenList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupApiToken

> ApiToken LookupApiToken(ctx).ApiTokenSecret(apiTokenSecret).Execute()

Gets API token metadata by token secret

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
    apiTokenSecret := *openapiclient.NewApiTokenSecret("dt0c01.ST2EY72KQINMH574WMNVI7YN.G3DFPBEJYMODIDAEX454M7YWBUVEFOWKPRVMWFASS64NFH52PX6BNDVFFM572RZM") // ApiTokenSecret | The JSON body of the request. Contains the required token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.LookupApiToken(context.Background()).ApiTokenSecret(apiTokenSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.LookupApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPITokensApi.LookupApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTokenSecret** | [**ApiTokenSecret**](ApiTokenSecret.md) | The JSON body of the request. Contains the required token. | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiToken

> UpdateApiToken(ctx, id).ApiTokenUpdate(apiTokenUpdate).Execute()

Updates an API token

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
    id := "id_example" // string | The ID of the token to be updated.    You can't disable the token you're using for authentication of the request.
    apiTokenUpdate := *openapiclient.NewApiTokenUpdate() // ApiTokenUpdate | The JSON body of the request. Contains updated parameters of the API token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokensAPITokensApi.UpdateApiToken(context.Background(), id).ApiTokenUpdate(apiTokenUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPITokensApi.UpdateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the token to be updated.    You can&#39;t disable the token you&#39;re using for authentication of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTokenUpdate** | [**ApiTokenUpdate**](ApiTokenUpdate.md) | The JSON body of the request. Contains updated parameters of the API token. | 

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

