# \AnonymizationApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Anonymize**](AnonymizationApi.md#Anonymize) | **Put** /anonymize/anonymizationJobs | Creates user session anonymization job
[**GetStatus**](AnonymizationApi.md#GetStatus) | **Get** /anonymize/anonymizationJobs/{requestId} | Shows the progress of the specified anonymization job



## Anonymize

> AnonymizationIdResult Anonymize(ctx).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).UserIds(userIds).Ips(ips).AdditionalField(additionalField).Execute()

Creates user session anonymization job



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
    startTimestamp := int64(789) // int64 | The start timestamp of the user session to anonymize, in UTC milliseconds.   If not set the earliest available time is used. (optional) (default to 0)
    endTimestamp := int64(789) // int64 | The end timestamp of the user session to anonymize, in UTC milliseconds.   If not set the current time is used. (optional) (default to 0)
    userIds := []string{"Inner_example"} // []string | The UserID of the user to anonymize.    You can specify several IDs, in the following format: `userIds=user1&userIds=user2`. (optional)
    ips := []string{"Inner_example"} // []string | The IP address of the user to anonymize. All user sessions from this IP will be anonymized.   You can specify several IPs, in the following format: `ips=ip1&ips=ip2`. (optional)
    additionalField := []string{"AdditionalField_example"} // []string | A list of fields to be anonymized.   You can specify several fields, in the following format: `additionalField=field1&additionalField=field2`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnonymizationApi.Anonymize(context.Background()).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).UserIds(userIds).Ips(ips).AdditionalField(additionalField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnonymizationApi.Anonymize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Anonymize`: AnonymizationIdResult
    fmt.Fprintf(os.Stdout, "Response from `AnonymizationApi.Anonymize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnonymizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **int64** | The start timestamp of the user session to anonymize, in UTC milliseconds.   If not set the earliest available time is used. | [default to 0]
 **endTimestamp** | **int64** | The end timestamp of the user session to anonymize, in UTC milliseconds.   If not set the current time is used. | [default to 0]
 **userIds** | **[]string** | The UserID of the user to anonymize.    You can specify several IDs, in the following format: &#x60;userIds&#x3D;user1&amp;userIds&#x3D;user2&#x60;. | 
 **ips** | **[]string** | The IP address of the user to anonymize. All user sessions from this IP will be anonymized.   You can specify several IPs, in the following format: &#x60;ips&#x3D;ip1&amp;ips&#x3D;ip2&#x60;. | 
 **additionalField** | **[]string** | A list of fields to be anonymized.   You can specify several fields, in the following format: &#x60;additionalField&#x3D;field1&amp;additionalField&#x3D;field2&#x60;. | 

### Return type

[**AnonymizationIdResult**](AnonymizationIdResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> AnonymizationProgressResult GetStatus(ctx, requestId).Execute()

Shows the progress of the specified anonymization job

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
    requestId := "requestId_example" // string | The ID of the required anonymization job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnonymizationApi.GetStatus(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnonymizationApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: AnonymizationProgressResult
    fmt.Fprintf(os.Stdout, "Response from `AnonymizationApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The ID of the required anonymization job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnonymizationProgressResult**](AnonymizationProgressResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

