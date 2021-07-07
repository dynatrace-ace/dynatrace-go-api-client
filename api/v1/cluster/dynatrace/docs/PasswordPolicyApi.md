# \PasswordPolicyApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultPasswordPolicy**](PasswordPolicyApi.md#GetDefaultPasswordPolicy) | **Get** /passwordPolicy | Get default&#39;s realm password policy configuration
[**UpdatePasswordPolicy**](PasswordPolicyApi.md#UpdatePasswordPolicy) | **Put** /passwordPolicy | Update password policy configuration



## GetDefaultPasswordPolicy

> PasswordPolicy GetDefaultPasswordPolicy(ctx).Execute()

Get default's realm password policy configuration

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
    resp, r, err := api_client.PasswordPolicyApi.GetDefaultPasswordPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.GetDefaultPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultPasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.GetDefaultPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultPasswordPolicyRequest struct via the builder pattern


### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordPolicy

> UpdatePasswordPolicy(ctx).PasswordPolicy(passwordPolicy).Execute()

Update password policy configuration

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
    passwordPolicy := *openapiclient.NewPasswordPolicy(int64(123), int64(123), int64(123), int64(123), int64(123)) // PasswordPolicy | The JSON body of the request. Contains parameters of password policy configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PasswordPolicyApi.UpdatePasswordPolicy(context.Background()).PasswordPolicy(passwordPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.UpdatePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) | The JSON body of the request. Contains parameters of password policy configuration. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

