# \UserGroupsApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](UserGroupsApi.md#CreateGroup) | **Post** /groups | Create group
[**CreateGroups**](UserGroupsApi.md#CreateGroups) | **Post** /groups/bulk | Create groups
[**GetGroup**](UserGroupsApi.md#GetGroup) | **Get** /groups/{groupId} | Get group
[**GetGroups**](UserGroupsApi.md#GetGroups) | **Get** /groups | Get groups
[**GetManagementZonesPermissions**](UserGroupsApi.md#GetManagementZonesPermissions) | **Get** /groups/managementZones | Get management zone permissions for all groups
[**GetManagementZonesPermissionsForGroup**](UserGroupsApi.md#GetManagementZonesPermissionsForGroup) | **Get** /groups/managementZones/{groupId} | Get management zone permissions for a given group
[**RemoveGroup**](UserGroupsApi.md#RemoveGroup) | **Delete** /groups/{groupId} | Delete group
[**UpdateGroup**](UserGroupsApi.md#UpdateGroup) | **Put** /groups | Update group
[**UpdateManagementZonesPermissionsForGroup**](UserGroupsApi.md#UpdateManagementZonesPermissionsForGroup) | **Put** /groups/managementZones | Update management zone permissions for a single group



## CreateGroup

> GroupConfig CreateGroup(ctx).GroupConfig(groupConfig).Execute()

Create group

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
    groupConfig := *openapiclient.NewGroupConfig(false, "Id_example", "Name_example") // GroupConfig | Request body used for creating new user group.  For creating user group leave 'id' empty, setting 'id' will return 'Bad Request'.  Trying to create group with name that already exists will return 'Not Acceptable'. 'isAccessAccount' value is ignored when 'Dynatrace Platform Subscription' is not in use. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.CreateGroup(context.Background()).GroupConfig(groupConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupConfig** | [**GroupConfig**](GroupConfig.md) | Request body used for creating new user group.  For creating user group leave &#39;id&#39; empty, setting &#39;id&#39; will return &#39;Bad Request&#39;.  Trying to create group with name that already exists will return &#39;Not Acceptable&#39;. &#39;isAccessAccount&#39; value is ignored when &#39;Dynatrace Platform Subscription&#39; is not in use. | 

### Return type

[**GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroups

> []GroupConfig CreateGroups(ctx).GroupConfig(groupConfig).Execute()

Create groups

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
    groupConfig := []openapiclient.GroupConfig{*openapiclient.NewGroupConfig(false, "Id_example", "Name_example")} // []GroupConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.CreateGroups(context.Background()).GroupConfig(groupConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.CreateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroups`: []GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.CreateGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupConfig** | [**[]GroupConfig**](GroupConfig.md) |  | 

### Return type

[**[]GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> GroupConfig GetGroup(ctx, groupId).Execute()

Get group

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
    groupId := "groupId_example" // string | Group ID path parameter. Missing or empty values will return a 'Bad Request'.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.GetGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID path parameter. Missing or empty values will return a &#39;Bad Request&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> []GroupConfig GetGroups(ctx).Execute()

Get groups

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
    resp, r, err := api_client.UserGroupsApi.GetGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: []GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


### Return type

[**[]GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementZonesPermissions

> []MzPermissionsForGroup GetManagementZonesPermissions(ctx).Execute()

Get management zone permissions for all groups



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
    resp, r, err := api_client.UserGroupsApi.GetManagementZonesPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GetManagementZonesPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementZonesPermissions`: []MzPermissionsForGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GetManagementZonesPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementZonesPermissionsRequest struct via the builder pattern


### Return type

[**[]MzPermissionsForGroup**](MzPermissionsForGroup.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementZonesPermissionsForGroup

> MzPermissionsForGroup GetManagementZonesPermissionsForGroup(ctx, groupId).Execute()

Get management zone permissions for a given group



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
    groupId := "groupId_example" // string | Group ID path parameter. Missing or empty values will return a 'Bad Request'.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.GetManagementZonesPermissionsForGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GetManagementZonesPermissionsForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementZonesPermissionsForGroup`: MzPermissionsForGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GetManagementZonesPermissionsForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID path parameter. Missing or empty values will return a &#39;Bad Request&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementZonesPermissionsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MzPermissionsForGroup**](MzPermissionsForGroup.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroup

> GroupConfig RemoveGroup(ctx, groupId).Execute()

Delete group

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
    groupId := "groupId_example" // string | Group ID path parameter. Missing or empty values will return a 'Bad Request'.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.RemoveGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.RemoveGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveGroup`: GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.RemoveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID path parameter. Missing or empty values will return a &#39;Bad Request&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> GroupConfig UpdateGroup(ctx).GroupConfig(groupConfig).Execute()

Update group

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
    groupConfig := *openapiclient.NewGroupConfig(false, "Id_example", "Name_example") // GroupConfig | Request body used for updating existing user group.  For updating user group set correct 'id', not setting 'id' will return 'Bad Request'.  Trying to change group name to one that already exists will return 'Bad Request'.  Trying to update group that doesn't exist will return 'Not Acceptable'. 'isAccessAccount' value is ignored when 'Dynatrace Platform Subscription' is not in use.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UpdateGroup(context.Background()).GroupConfig(groupConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: GroupConfig
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupConfig** | [**GroupConfig**](GroupConfig.md) | Request body used for updating existing user group.  For updating user group set correct &#39;id&#39;, not setting &#39;id&#39; will return &#39;Bad Request&#39;.  Trying to change group name to one that already exists will return &#39;Bad Request&#39;.  Trying to update group that doesn&#39;t exist will return &#39;Not Acceptable&#39;. &#39;isAccessAccount&#39; value is ignored when &#39;Dynatrace Platform Subscription&#39; is not in use.  | 

### Return type

[**GroupConfig**](GroupConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManagementZonesPermissionsForGroup

> UpdateManagementZonesPermissionsForGroup(ctx).MzPermissionsForGroup(mzPermissionsForGroup).Execute()

Update management zone permissions for a single group

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
    mzPermissionsForGroup := *openapiclient.NewMzPermissionsForGroup() // MzPermissionsForGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UpdateManagementZonesPermissionsForGroup(context.Background()).MzPermissionsForGroup(mzPermissionsForGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UpdateManagementZonesPermissionsForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManagementZonesPermissionsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mzPermissionsForGroup** | [**MzPermissionsForGroup**](MzPermissionsForGroup.md) |  | 

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

