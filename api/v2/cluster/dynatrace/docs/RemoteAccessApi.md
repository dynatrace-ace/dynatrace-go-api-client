# \RemoteAccessApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessRequestChangeState**](RemoteAccessApi.md#AccessRequestChangeState) | **Put** /remoteaccess/requests/{requestId}/state | Change state of access request
[**AddAccessRequest**](RemoteAccessApi.md#AddAccessRequest) | **Post** /remoteaccess/requests | Grant remote access permission
[**GetAccessRequest**](RemoteAccessApi.md#GetAccessRequest) | **Get** /remoteaccess/requests/{requestId} | Get access request
[**GetAllAccessRequests**](RemoteAccessApi.md#GetAllAccessRequests) | **Get** /remoteaccess/requests | Get all access requests



## AccessRequestChangeState

> AccessRequestChangeState(ctx, requestId).AccessRequestStateData(accessRequestStateData).Execute()

Change state of access request

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
    requestId := "requestId_example" // string | Request id param
    accessRequestStateData := *openapiclient.NewAccessRequestStateData() // AccessRequestStateData | The JSON body of the request, containing new state of access request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteAccessApi.AccessRequestChangeState(context.Background(), requestId).AccessRequestStateData(accessRequestStateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessApi.AccessRequestChangeState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request id param | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestChangeStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRequestStateData** | [**AccessRequestStateData**](AccessRequestStateData.md) | The JSON body of the request, containing new state of access request. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequest

> AccessRequestData AddAccessRequest(ctx).CreateAccessRequestDto(createAccessRequestDto).Execute()

Grant remote access permission

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
    createAccessRequestDto := *openapiclient.NewCreateAccessRequestDto() // CreateAccessRequestDto | The JSON body of the request, containing parameters of access request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteAccessApi.AddAccessRequest(context.Background()).CreateAccessRequestDto(createAccessRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessApi.AddAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessRequest`: AccessRequestData
    fmt.Fprintf(os.Stdout, "Response from `RemoteAccessApi.AddAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessRequestDto** | [**CreateAccessRequestDto**](CreateAccessRequestDto.md) | The JSON body of the request, containing parameters of access request. | 

### Return type

[**AccessRequestData**](AccessRequestData.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequest

> AccessRequestData GetAccessRequest(ctx, requestId).Execute()

Get access request

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
    requestId := "requestId_example" // string | Request id param

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RemoteAccessApi.GetAccessRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessApi.GetAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequest`: AccessRequestData
    fmt.Fprintf(os.Stdout, "Response from `RemoteAccessApi.GetAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request id param | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequestData**](AccessRequestData.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessRequests

> []AccessRequestData GetAllAccessRequests(ctx).Execute()

Get all access requests

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
    resp, r, err := api_client.RemoteAccessApi.GetAllAccessRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessApi.GetAllAccessRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAccessRequests`: []AccessRequestData
    fmt.Fprintf(os.Stdout, "Response from `RemoteAccessApi.GetAllAccessRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessRequestsRequest struct via the builder pattern


### Return type

[**[]AccessRequestData**](AccessRequestData.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

