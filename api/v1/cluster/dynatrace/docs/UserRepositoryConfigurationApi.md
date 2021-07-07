# \UserRepositoryConfigurationApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthenticationMode**](UserRepositoryConfigurationApi.md#GetAuthenticationMode) | **Get** /userRepository/authenticationMode | Get authentication mode
[**GetLdapConnectionConfiguration**](UserRepositoryConfigurationApi.md#GetLdapConnectionConfiguration) | **Get** /userRepository/ldap/connectionConfiguration | Get LDAP configuration
[**GetLdapGroupsQuery**](UserRepositoryConfigurationApi.md#GetLdapGroupsQuery) | **Get** /userRepository/ldap/queryGroups | Get LDAP group configuration
[**GetLdapLdapUsersQuery**](UserRepositoryConfigurationApi.md#GetLdapLdapUsersQuery) | **Get** /userRepository/ldap/queryUsers | Get LDAP users query configuration
[**UpdateAuthenticationMode**](UserRepositoryConfigurationApi.md#UpdateAuthenticationMode) | **Post** /userRepository/authenticationMode | Update authentication mode
[**UpdateLdapConnection**](UserRepositoryConfigurationApi.md#UpdateLdapConnection) | **Post** /userRepository/ldap/connectionConfiguration | Update LDAP connection
[**UpdateLdapGroupsQuery**](UserRepositoryConfigurationApi.md#UpdateLdapGroupsQuery) | **Post** /userRepository/ldap/queryGroups | Update LDAP groups query configuration
[**UpdateLdapUsersQueryDescImpl**](UserRepositoryConfigurationApi.md#UpdateLdapUsersQueryDescImpl) | **Post** /userRepository/ldap/queryUsers | Update LDAP users query configuration



## GetAuthenticationMode

> AuthenticationMode GetAuthenticationMode(ctx).Execute()

Get authentication mode

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
    resp, r, err := api_client.UserRepositoryConfigurationApi.GetAuthenticationMode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.GetAuthenticationMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationMode`: AuthenticationMode
    fmt.Fprintf(os.Stdout, "Response from `UserRepositoryConfigurationApi.GetAuthenticationMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationModeRequest struct via the builder pattern


### Return type

[**AuthenticationMode**](AuthenticationMode.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConnectionConfiguration

> LdapConnectionDescImpl GetLdapConnectionConfiguration(ctx).Execute()

Get LDAP configuration

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
    resp, r, err := api_client.UserRepositoryConfigurationApi.GetLdapConnectionConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.GetLdapConnectionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapConnectionConfiguration`: LdapConnectionDescImpl
    fmt.Fprintf(os.Stdout, "Response from `UserRepositoryConfigurationApi.GetLdapConnectionConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConnectionConfigurationRequest struct via the builder pattern


### Return type

[**LdapConnectionDescImpl**](LdapConnectionDescImpl.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapGroupsQuery

> LdapGroupsQueryDtoImpl GetLdapGroupsQuery(ctx).Execute()

Get LDAP group configuration

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
    resp, r, err := api_client.UserRepositoryConfigurationApi.GetLdapGroupsQuery(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.GetLdapGroupsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapGroupsQuery`: LdapGroupsQueryDtoImpl
    fmt.Fprintf(os.Stdout, "Response from `UserRepositoryConfigurationApi.GetLdapGroupsQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapGroupsQueryRequest struct via the builder pattern


### Return type

[**LdapGroupsQueryDtoImpl**](LdapGroupsQueryDtoImpl.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapLdapUsersQuery

> LdapUsersQueryDescImpl GetLdapLdapUsersQuery(ctx).Execute()

Get LDAP users query configuration

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
    resp, r, err := api_client.UserRepositoryConfigurationApi.GetLdapLdapUsersQuery(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.GetLdapLdapUsersQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapLdapUsersQuery`: LdapUsersQueryDescImpl
    fmt.Fprintf(os.Stdout, "Response from `UserRepositoryConfigurationApi.GetLdapLdapUsersQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapLdapUsersQueryRequest struct via the builder pattern


### Return type

[**LdapUsersQueryDescImpl**](LdapUsersQueryDescImpl.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationMode

> AuthenticationMode UpdateAuthenticationMode(ctx).AuthenticationModel(authenticationModel).Execute()

Update authentication mode

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
    authenticationModel := *openapiclient.NewAuthenticationModel("AuthenticationProvider_example") // AuthenticationModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserRepositoryConfigurationApi.UpdateAuthenticationMode(context.Background()).AuthenticationModel(authenticationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.UpdateAuthenticationMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticationMode`: AuthenticationMode
    fmt.Fprintf(os.Stdout, "Response from `UserRepositoryConfigurationApi.UpdateAuthenticationMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationModel** | [**AuthenticationModel**](AuthenticationModel.md) |  | 

### Return type

[**AuthenticationMode**](AuthenticationMode.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapConnection

> UpdateLdapConnection(ctx).LdapConnectionDescImpl(ldapConnectionDescImpl).Execute()

Update LDAP connection

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
    ldapConnectionDescImpl := *openapiclient.NewLdapConnectionDescImpl() // LdapConnectionDescImpl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserRepositoryConfigurationApi.UpdateLdapConnection(context.Background()).LdapConnectionDescImpl(ldapConnectionDescImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.UpdateLdapConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConnectionDescImpl** | [**LdapConnectionDescImpl**](LdapConnectionDescImpl.md) |  | 

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


## UpdateLdapGroupsQuery

> UpdateLdapGroupsQuery(ctx).LdapGroupsQueryDtoImpl(ldapGroupsQueryDtoImpl).Execute()

Update LDAP groups query configuration

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
    ldapGroupsQueryDtoImpl := *openapiclient.NewLdapGroupsQueryDtoImpl() // LdapGroupsQueryDtoImpl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserRepositoryConfigurationApi.UpdateLdapGroupsQuery(context.Background()).LdapGroupsQueryDtoImpl(ldapGroupsQueryDtoImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.UpdateLdapGroupsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapGroupsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapGroupsQueryDtoImpl** | [**LdapGroupsQueryDtoImpl**](LdapGroupsQueryDtoImpl.md) |  | 

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


## UpdateLdapUsersQueryDescImpl

> UpdateLdapUsersQueryDescImpl(ctx).LdapUsersQueryDescImpl(ldapUsersQueryDescImpl).Execute()

Update LDAP users query configuration

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
    ldapUsersQueryDescImpl := *openapiclient.NewLdapUsersQueryDescImpl() // LdapUsersQueryDescImpl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserRepositoryConfigurationApi.UpdateLdapUsersQueryDescImpl(context.Background()).LdapUsersQueryDescImpl(ldapUsersQueryDescImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRepositoryConfigurationApi.UpdateLdapUsersQueryDescImpl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapUsersQueryDescImplRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapUsersQueryDescImpl** | [**LdapUsersQueryDescImpl**](LdapUsersQueryDescImpl.md) |  | 

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

