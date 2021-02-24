# \FrequentIssueDetectionApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFrequentIssueDetectionConfig**](FrequentIssueDetectionApi.md#GetFrequentIssueDetectionConfig) | **Get** /frequentIssueDetection | Gets the configuration of frequent issue detection
[**UpdateFrequentIssueDetectionConfig**](FrequentIssueDetectionApi.md#UpdateFrequentIssueDetectionConfig) | **Put** /frequentIssueDetection | Updates the configuration of frequent issue detection
[**ValidateFrequentIssueDetectionConfig**](FrequentIssueDetectionApi.md#ValidateFrequentIssueDetectionConfig) | **Post** /frequentIssueDetection/validator | Validates the payload for the &#x60;PUT /frequentIssueDetection&#x60; request



## GetFrequentIssueDetectionConfig

> FrequentIssueDetectionConfig GetFrequentIssueDetectionConfig(ctx).Execute()

Gets the configuration of frequent issue detection

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
    resp, r, err := api_client.FrequentIssueDetectionApi.GetFrequentIssueDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrequentIssueDetectionApi.GetFrequentIssueDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrequentIssueDetectionConfig`: FrequentIssueDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `FrequentIssueDetectionApi.GetFrequentIssueDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrequentIssueDetectionConfigRequest struct via the builder pattern


### Return type

[**FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFrequentIssueDetectionConfig

> UpdateFrequentIssueDetectionConfig(ctx).FrequentIssueDetectionConfig(frequentIssueDetectionConfig).Execute()

Updates the configuration of frequent issue detection

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
    frequentIssueDetectionConfig := *openapiclient.NewFrequentIssueDetectionConfig(false, false, false) // FrequentIssueDetectionConfig | The JSON body of the request, containing parameters of the frequent issue detection configuration (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FrequentIssueDetectionApi.UpdateFrequentIssueDetectionConfig(context.Background()).FrequentIssueDetectionConfig(frequentIssueDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrequentIssueDetectionApi.UpdateFrequentIssueDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFrequentIssueDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frequentIssueDetectionConfig** | [**FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md) | The JSON body of the request, containing parameters of the frequent issue detection configuration | 

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


## ValidateFrequentIssueDetectionConfig

> ValidateFrequentIssueDetectionConfig(ctx).FrequentIssueDetectionConfig(frequentIssueDetectionConfig).Execute()

Validates the payload for the `PUT /frequentIssueDetection` request

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
    frequentIssueDetectionConfig := *openapiclient.NewFrequentIssueDetectionConfig(false, false, false) // FrequentIssueDetectionConfig | The JSON body of the request, containing parameters of the frequent issue detection configuration (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FrequentIssueDetectionApi.ValidateFrequentIssueDetectionConfig(context.Background()).FrequentIssueDetectionConfig(frequentIssueDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrequentIssueDetectionApi.ValidateFrequentIssueDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFrequentIssueDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frequentIssueDetectionConfig** | [**FrequentIssueDetectionConfig**](FrequentIssueDetectionConfig.md) | The JSON body of the request, containing parameters of the frequent issue detection configuration | 

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

