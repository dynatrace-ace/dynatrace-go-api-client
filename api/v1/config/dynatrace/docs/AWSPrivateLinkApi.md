# \AWSPrivateLinkApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllAccounts**](AWSPrivateLinkApi.md#GetAllAccounts) | **Get** /aws/privateLink/whitelistedAccounts | Gets the information about all whitelisted accounts in AWS PrivateLink | maturity&#x3D;EARLY_ADOPTER
[**GetPrivateLinkConfig**](AWSPrivateLinkApi.md#GetPrivateLinkConfig) | **Get** /aws/privateLink | Gets the configuration information about AWS PrivateLink | maturity&#x3D;EARLY_ADOPTER
[**PutAccount**](AWSPrivateLinkApi.md#PutAccount) | **Put** /aws/privateLink/whitelistedAccounts/{id} | Updates the given AWS account id in the whitelist in AWS PrivateLink | maturity&#x3D;EARLY_ADOPTER
[**PutPrivateLinkConfig**](AWSPrivateLinkApi.md#PutPrivateLinkConfig) | **Put** /aws/privateLink | Updates the configuration information about AWS PrivateLink | maturity&#x3D;EARLY_ADOPTER
[**RemoveAccount**](AWSPrivateLinkApi.md#RemoveAccount) | **Delete** /aws/privateLink/whitelistedAccounts/{id} | Removes one AWS account from the whitelist in AWS PrivateLink | maturity&#x3D;EARLY_ADOPTER



## GetAllAccounts

> WhitelistedAwsAccountList GetAllAccounts(ctx).Execute()

Gets the information about all whitelisted accounts in AWS PrivateLink | maturity=EARLY_ADOPTER



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
    resp, r, err := api_client.AWSPrivateLinkApi.GetAllAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkApi.GetAllAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAccounts`: WhitelistedAwsAccountList
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkApi.GetAllAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccountsRequest struct via the builder pattern


### Return type

[**WhitelistedAwsAccountList**](WhitelistedAwsAccountList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateLinkConfig

> AwsPrivateLinkConfig GetPrivateLinkConfig(ctx).Execute()

Gets the configuration information about AWS PrivateLink | maturity=EARLY_ADOPTER



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
    resp, r, err := api_client.AWSPrivateLinkApi.GetPrivateLinkConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkApi.GetPrivateLinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateLinkConfig`: AwsPrivateLinkConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkApi.GetPrivateLinkConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateLinkConfigRequest struct via the builder pattern


### Return type

[**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccount

> WhitelistedAwsAccount PutAccount(ctx, id).WhitelistedAwsAccount(whitelistedAwsAccount).Execute()

Updates the given AWS account id in the whitelist in AWS PrivateLink | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The AWS account id to be updated in the AWS PrivateLink whitelist, has to match the id in the provided payload.
    whitelistedAwsAccount := *openapiclient.NewWhitelistedAwsAccount("Id_example") // WhitelistedAwsAccount | The AWS account id to be updated in the AWS PrivateLink whitelist, has to match the id in the path.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSPrivateLinkApi.PutAccount(context.Background(), id).WhitelistedAwsAccount(whitelistedAwsAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkApi.PutAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccount`: WhitelistedAwsAccount
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkApi.PutAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The AWS account id to be updated in the AWS PrivateLink whitelist, has to match the id in the provided payload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whitelistedAwsAccount** | [**WhitelistedAwsAccount**](WhitelistedAwsAccount.md) | The AWS account id to be updated in the AWS PrivateLink whitelist, has to match the id in the path. | 

### Return type

[**WhitelistedAwsAccount**](WhitelistedAwsAccount.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrivateLinkConfig

> AwsPrivateLinkConfig PutPrivateLinkConfig(ctx).AwsPrivateLinkConfig(awsPrivateLinkConfig).Execute()

Updates the configuration information about AWS PrivateLink | maturity=EARLY_ADOPTER



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
    awsPrivateLinkConfig := *openapiclient.NewAwsPrivateLinkConfig(false) // AwsPrivateLinkConfig | The AWS PrivateLink configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSPrivateLinkApi.PutPrivateLinkConfig(context.Background()).AwsPrivateLinkConfig(awsPrivateLinkConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkApi.PutPrivateLinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPrivateLinkConfig`: AwsPrivateLinkConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkApi.PutPrivateLinkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPrivateLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsPrivateLinkConfig** | [**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md) | The AWS PrivateLink configuration. | 

### Return type

[**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccount

> RemoveAccount(ctx, id).Execute()

Removes one AWS account from the whitelist in AWS PrivateLink | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The AWS account id to be deleted from the AWS PrivateLink whitelist

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSPrivateLinkApi.RemoveAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkApi.RemoveAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The AWS account id to be deleted from the AWS PrivateLink whitelist | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountRequest struct via the builder pattern


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

