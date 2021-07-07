# \PreferencesApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProperties**](PreferencesApi.md#GetProperties) | **Get** /preferences | Get Dynatrace Managed specific properties
[**UpdateProperties**](PreferencesApi.md#UpdateProperties) | **Post** /preferences | Update properties



## GetProperties

> Preferences GetProperties(ctx).Execute()

Get Dynatrace Managed specific properties

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
    resp, r, err := api_client.PreferencesApi.GetProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.GetProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProperties`: Preferences
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.GetProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesRequest struct via the builder pattern


### Return type

[**Preferences**](Preferences.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProperties

> Preferences UpdateProperties(ctx).Preferences(preferences).Execute()

Update properties

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
    preferences := *openapiclient.NewPreferences(false, false, false) // Preferences |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreferencesApi.UpdateProperties(context.Background()).Preferences(preferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.UpdateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProperties`: Preferences
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.UpdateProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preferences** | [**Preferences**](Preferences.md) |  | 

### Return type

[**Preferences**](Preferences.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

