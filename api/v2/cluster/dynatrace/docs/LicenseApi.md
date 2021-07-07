# \LicenseApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingArchive**](LicenseApi.md#GetBillingArchive) | **Get** /license/consumption | Export license consumption data
[**GetBillingHour**](LicenseApi.md#GetBillingHour) | **Get** /license/consumption/hour | Retrieve license consumption



## GetBillingArchive

> GetBillingArchive(ctx).StartTs(startTs).EndTs(endTs).Execute()

Export license consumption data



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
    startTs := int64(789) // int64 | Start timestamp (30 days ago by default) (optional)
    endTs := int64(789) // int64 | End timestamp (2 hours ago by default, values from the last 2 hours are not allowed) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetBillingArchive(context.Background()).StartTs(startTs).EndTs(endTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetBillingArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTs** | **int64** | Start timestamp (30 days ago by default) | 
 **endTs** | **int64** | End timestamp (2 hours ago by default, values from the last 2 hours are not allowed) | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingHour

> BillingRequestDto GetBillingHour(ctx).StartTs(startTs).Execute()

Retrieve license consumption

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
    startTs := int64(789) // int64 | Begin timestamp (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicenseApi.GetBillingHour(context.Background()).StartTs(startTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.GetBillingHour``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingHour`: BillingRequestDto
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.GetBillingHour`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingHourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTs** | **int64** | Begin timestamp | 

### Return type

[**BillingRequestDto**](BillingRequestDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

