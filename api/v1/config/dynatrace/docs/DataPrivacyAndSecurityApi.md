# \DataPrivacyAndSecurityApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataPrivacySettings1**](DataPrivacyAndSecurityApi.md#GetDataPrivacySettings1) | **Get** /dataPrivacy | Lists the global data privacy and security settings.
[**UpdateDataPrivacySettings1**](DataPrivacyAndSecurityApi.md#UpdateDataPrivacySettings1) | **Put** /dataPrivacy | Updates the global data privacy and security settings.
[**ValidateDataPrivacySettings1**](DataPrivacyAndSecurityApi.md#ValidateDataPrivacySettings1) | **Post** /dataPrivacy/validator | Validates new data privacy and security settings for the &#x60;PUT /dataPrivacy&#x60; request.



## GetDataPrivacySettings1

> DataPrivacyAndSecurity GetDataPrivacySettings1(ctx).Execute()

Lists the global data privacy and security settings.

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
    resp, r, err := api_client.DataPrivacyAndSecurityApi.GetDataPrivacySettings1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPrivacyAndSecurityApi.GetDataPrivacySettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataPrivacySettings1`: DataPrivacyAndSecurity
    fmt.Fprintf(os.Stdout, "Response from `DataPrivacyAndSecurityApi.GetDataPrivacySettings1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPrivacySettings1Request struct via the builder pattern


### Return type

[**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataPrivacySettings1

> UpdateDataPrivacySettings1(ctx).DataPrivacyAndSecurity(dataPrivacyAndSecurity).Execute()

Updates the global data privacy and security settings.



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
    dataPrivacyAndSecurity := *openapiclient.NewDataPrivacyAndSecurity(false, false, false) // DataPrivacyAndSecurity | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPrivacyAndSecurityApi.UpdateDataPrivacySettings1(context.Background()).DataPrivacyAndSecurity(dataPrivacyAndSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPrivacyAndSecurityApi.UpdateDataPrivacySettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataPrivacySettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataPrivacyAndSecurity** | [**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md) |  | 

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


## ValidateDataPrivacySettings1

> ValidateDataPrivacySettings1(ctx).DataPrivacyAndSecurity(dataPrivacyAndSecurity).Execute()

Validates new data privacy and security settings for the `PUT /dataPrivacy` request.

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
    dataPrivacyAndSecurity := *openapiclient.NewDataPrivacyAndSecurity(false, false, false) // DataPrivacyAndSecurity | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPrivacyAndSecurityApi.ValidateDataPrivacySettings1(context.Background()).DataPrivacyAndSecurity(dataPrivacyAndSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPrivacyAndSecurityApi.ValidateDataPrivacySettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDataPrivacySettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataPrivacyAndSecurity** | [**DataPrivacyAndSecurity**](DataPrivacyAndSecurity.md) |  | 

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

