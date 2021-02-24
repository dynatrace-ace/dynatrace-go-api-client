# \NotificationsApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationConfig**](NotificationsApi.md#CreateNotificationConfig) | **Post** /notifications | Creates a new notification configuration
[**DeleteNotificationConfig**](NotificationsApi.md#DeleteNotificationConfig) | **Delete** /notifications/{id} | Deletes the specified notification configuration
[**GetNotificationConfig**](NotificationsApi.md#GetNotificationConfig) | **Get** /notifications/{id} | Gets the properties of the specified notification configuration
[**ListNotificationConfigs**](NotificationsApi.md#ListNotificationConfigs) | **Get** /notifications | Lists all notification configurations available in your environment
[**UpdateNotificationConfig**](NotificationsApi.md#UpdateNotificationConfig) | **Put** /notifications/{id} | Updates an existing notification configuration or creates a new one
[**ValidateCreateNotificationConfig**](NotificationsApi.md#ValidateCreateNotificationConfig) | **Post** /notifications/validator | Validates the payload for the &#x60;POST /notifications&#x60; request
[**ValidateUpdateNotificationConfig**](NotificationsApi.md#ValidateUpdateNotificationConfig) | **Post** /notifications/{id}/validator | Validates the payload the &#x60;PUT /notifications/{id}&#x60; request



## CreateNotificationConfig

> CreateNotificationConfig(ctx).NotificationConfig(notificationConfig).Execute()

Creates a new notification configuration

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
    notificationConfig := *openapiclient.NewNotificationConfig("Name_example", "AlertingProfile_example", false, "Type_example", []string{"Receivers_example"}, []string{"CcReceivers_example"}, []string{"BccReceivers_example"}) // NotificationConfig | The JSON body of the request. Contains parameters of the new notification configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.CreateNotificationConfig(context.Background()).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.CreateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains parameters of the new notification configuration. | 

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


## DeleteNotificationConfig

> DeleteNotificationConfig(ctx, id).Execute()

Deletes the specified notification configuration

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
    id := TODO // string | The ID of the notification configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.DeleteNotificationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the notification configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetNotificationConfig

> NotificationConfig GetNotificationConfig(ctx, id).Execute()

Gets the properties of the specified notification configuration

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
    id := TODO // string | The ID of the required notification configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.GetNotificationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfig`: NotificationConfig
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the required notification configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationConfig**](NotificationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationConfigs

> NotificationConfigStubListDto ListNotificationConfigs(ctx).Execute()

Lists all notification configurations available in your environment

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
    resp, r, err := api_client.NotificationsApi.ListNotificationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ListNotificationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationConfigs`: NotificationConfigStubListDto
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ListNotificationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationConfigsRequest struct via the builder pattern


### Return type

[**NotificationConfigStubListDto**](NotificationConfigStubListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationConfig

> NotificationConfigStub UpdateNotificationConfig(ctx, id).NotificationConfig(notificationConfig).Execute()

Updates an existing notification configuration or creates a new one



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
    id := TODO // string | The ID of the notification configuration to be updated.    If you set the ID in the body as well, it must match this ID.
    notificationConfig := *openapiclient.NewNotificationConfig("Name_example", "AlertingProfile_example", false, "Type_example", []string{"Receivers_example"}, []string{"CcReceivers_example"}, []string{"BccReceivers_example"}) // NotificationConfig | The JSON body of the request. Contains updated parameters of the notification configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.UpdateNotificationConfig(context.Background(), id).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.UpdateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationConfig`: NotificationConfigStub
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.UpdateNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the notification configuration to be updated.    If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains updated parameters of the notification configuration. | 

### Return type

[**NotificationConfigStub**](NotificationConfigStub.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateNotificationConfig

> ValidateCreateNotificationConfig(ctx).NotificationConfig(notificationConfig).Execute()

Validates the payload for the `POST /notifications` request

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
    notificationConfig := *openapiclient.NewNotificationConfig("Name_example", "AlertingProfile_example", false, "Type_example", []string{"Receivers_example"}, []string{"CcReceivers_example"}, []string{"BccReceivers_example"}) // NotificationConfig | The JSON body of the request. Contains the notification configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.ValidateCreateNotificationConfig(context.Background()).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ValidateCreateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains the notification configuration to be validated. | 

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


## ValidateUpdateNotificationConfig

> ValidateUpdateNotificationConfig(ctx, id).NotificationConfig(notificationConfig).Execute()

Validates the payload the `PUT /notifications/{id}` request

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
    id := TODO // string | The ID of the notification configuration to be validated.
    notificationConfig := *openapiclient.NewNotificationConfig("Name_example", "AlertingProfile_example", false, "Type_example", []string{"Receivers_example"}, []string{"CcReceivers_example"}, []string{"BccReceivers_example"}) // NotificationConfig | The JSON body of the request. Contains the notification configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.ValidateUpdateNotificationConfig(context.Background(), id).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ValidateUpdateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | The ID of the notification configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains the notification configuration to be validated. | 

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

