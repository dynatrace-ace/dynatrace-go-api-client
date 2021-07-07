# \UserManagementApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConcurrentSessionPolicyConfig**](UserManagementApi.md#GetConcurrentSessionPolicyConfig) | **Get** /clusterConfig/userSessions | Get user sessions configuration
[**GetUserSessions**](UserManagementApi.md#GetUserSessions) | **Get** /userSessions | Get user sessions
[**RemoveUserSession**](UserManagementApi.md#RemoveUserSession) | **Delete** /userSessions | Remove user sessions for a given user
[**UpdateConcurrentSessionPolicyConfig**](UserManagementApi.md#UpdateConcurrentSessionPolicyConfig) | **Put** /clusterConfig/userSessions | Update user sessions configuration



## GetConcurrentSessionPolicyConfig

> UserSessionsConfig GetConcurrentSessionPolicyConfig(ctx).Execute()

Get user sessions configuration

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
    resp, r, err := api_client.UserManagementApi.GetConcurrentSessionPolicyConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.GetConcurrentSessionPolicyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConcurrentSessionPolicyConfig`: UserSessionsConfig
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.GetConcurrentSessionPolicyConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConcurrentSessionPolicyConfigRequest struct via the builder pattern


### Return type

[**UserSessionsConfig**](UserSessionsConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSessions

> []UserSession GetUserSessions(ctx).UserId(userId).Execute()

Get user sessions

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
    userId := "userId_example" // string | User ID (optional) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.GetUserSessions(context.Background()).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.GetUserSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSessions`: []UserSession
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.GetUserSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User ID (optional) | 

### Return type

[**[]UserSession**](UserSession.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserSession

> RemoveUserSession(ctx).UserId(userId).Execute()

Remove user sessions for a given user

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
    userId := "userId_example" // string | User ID (mandatory) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.RemoveUserSession(context.Background()).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.RemoveUserSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User ID (mandatory) | 

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


## UpdateConcurrentSessionPolicyConfig

> UpdateConcurrentSessionPolicyConfig(ctx).UserSessionsConfig(userSessionsConfig).Execute()

Update user sessions configuration

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
    userSessionsConfig := *openapiclient.NewUserSessionsConfig(*openapiclient.NewConcurrentSessionPolicy(int32(123), int32(123)), *openapiclient.NewAutomaticLogoutConfiguration(false, int32(123))) // UserSessionsConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.UpdateConcurrentSessionPolicyConfig(context.Background()).UserSessionsConfig(userSessionsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UpdateConcurrentSessionPolicyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConcurrentSessionPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSessionsConfig** | [**UserSessionsConfig**](UserSessionsConfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

