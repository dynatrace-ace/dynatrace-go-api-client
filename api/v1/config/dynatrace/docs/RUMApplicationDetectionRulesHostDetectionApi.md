# \RUMApplicationDetectionRulesHostDetectionApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostDetectionConfig**](RUMApplicationDetectionRulesHostDetectionApi.md#GetHostDetectionConfig) | **Get** /applicationDetectionRules/hostDetection | Gets the configuration of host detection headers
[**UpdateHostDetectionConfig**](RUMApplicationDetectionRulesHostDetectionApi.md#UpdateHostDetectionConfig) | **Put** /applicationDetectionRules/hostDetection | Updates the configuration of host detection headers
[**ValidateHostDetectionConfig**](RUMApplicationDetectionRulesHostDetectionApi.md#ValidateHostDetectionConfig) | **Post** /applicationDetectionRules/hostDetection/validator | Validate the payload for &#x60;PUT /applicationDetection/hostDetection&#x60; request.



## GetHostDetectionConfig

> ApplicationDetectionRulesHostDetectionSettings GetHostDetectionConfig(ctx).Execute()

Gets the configuration of host detection headers

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
    resp, r, err := api_client.RUMApplicationDetectionRulesHostDetectionApi.GetHostDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesHostDetectionApi.GetHostDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostDetectionConfig`: ApplicationDetectionRulesHostDetectionSettings
    fmt.Fprintf(os.Stdout, "Response from `RUMApplicationDetectionRulesHostDetectionApi.GetHostDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostDetectionConfigRequest struct via the builder pattern


### Return type

[**ApplicationDetectionRulesHostDetectionSettings**](ApplicationDetectionRulesHostDetectionSettings.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostDetectionConfig

> UpdateHostDetectionConfig(ctx).ApplicationDetectionRulesHostDetectionSettings(applicationDetectionRulesHostDetectionSettings).Execute()

Updates the configuration of host detection headers



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
    applicationDetectionRulesHostDetectionSettings := *openapiclient.NewApplicationDetectionRulesHostDetectionSettings([]string{"HostDetectionHeaders_example"}) // ApplicationDetectionRulesHostDetectionSettings | The JSON body of the request. Contains the configuration of host detection headers. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMApplicationDetectionRulesHostDetectionApi.UpdateHostDetectionConfig(context.Background()).ApplicationDetectionRulesHostDetectionSettings(applicationDetectionRulesHostDetectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesHostDetectionApi.UpdateHostDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionRulesHostDetectionSettings** | [**ApplicationDetectionRulesHostDetectionSettings**](ApplicationDetectionRulesHostDetectionSettings.md) | The JSON body of the request. Contains the configuration of host detection headers. | 

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


## ValidateHostDetectionConfig

> ValidateHostDetectionConfig(ctx).ApplicationDetectionRulesHostDetectionSettings(applicationDetectionRulesHostDetectionSettings).Execute()

Validate the payload for `PUT /applicationDetection/hostDetection` request.

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
    applicationDetectionRulesHostDetectionSettings := *openapiclient.NewApplicationDetectionRulesHostDetectionSettings([]string{"HostDetectionHeaders_example"}) // ApplicationDetectionRulesHostDetectionSettings | The JSON body of the request. Contains the configuration of host detection headers to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMApplicationDetectionRulesHostDetectionApi.ValidateHostDetectionConfig(context.Background()).ApplicationDetectionRulesHostDetectionSettings(applicationDetectionRulesHostDetectionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesHostDetectionApi.ValidateHostDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateHostDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionRulesHostDetectionSettings** | [**ApplicationDetectionRulesHostDetectionSettings**](ApplicationDetectionRulesHostDetectionSettings.md) | The JSON body of the request. Contains the configuration of host detection headers to be validated. | 

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

