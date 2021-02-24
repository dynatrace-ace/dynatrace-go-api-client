# \RUMMobileAndCustomApplicationConfigurationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#CreateMobileApplicationConfig) | **Post** /applications/mobile | Creates a new mobile or custom application
[**CreateMobileKeyUserAction**](RUMMobileAndCustomApplicationConfigurationApi.md#CreateMobileKeyUserAction) | **Post** /applications/mobile/{applicationId}/keyUserActions/{actionName} | Marks the user action as a key user action in the specified application
[**CreateSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#CreateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties | Creates a new session property for the specified application
[**DeleteMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#DeleteMobileApplicationConfig) | **Delete** /applications/mobile/{id} | Deletes the specified mobile or custom application
[**DeleteMobileKeyUserAction**](RUMMobileAndCustomApplicationConfigurationApi.md#DeleteMobileKeyUserAction) | **Delete** /applications/mobile/{applicationId}/keyUserActions/{actionName} | Removes the specified user action from the list of key user actions in the specified application
[**DeleteSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#DeleteSessionProperty) | **Delete** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Deletes the specified session property for an application
[**GetMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#GetMobileApplicationConfig) | **Get** /applications/mobile/{id} | Gets the configuration of the specified mobile or custom application
[**GetSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#GetSessionProperty) | **Get** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Gets the specified  session property of an application
[**ListMobileApplicationConfigs**](RUMMobileAndCustomApplicationConfigurationApi.md#ListMobileApplicationConfigs) | **Get** /applications/mobile | Lists all existing mobile and custom applications
[**ListMobileKeyUserActions**](RUMMobileAndCustomApplicationConfigurationApi.md#ListMobileKeyUserActions) | **Get** /applications/mobile/{applicationId}/keyUserActions | Gets the list of key user actions in the specified application
[**ListSessionProperties**](RUMMobileAndCustomApplicationConfigurationApi.md#ListSessionProperties) | **Get** /applications/mobile/{applicationId}/userActionAndSessionProperties | Lists all session properties for the specified application
[**UpdateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#UpdateMobileApplicationConfig) | **Put** /applications/mobile/{id} | Updates the configuration of the specified mobile or custom application
[**UpdateSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#UpdateSessionProperty) | **Put** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Updates the specified session property for an application
[**ValidateCreateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#ValidateCreateMobileApplicationConfig) | **Post** /applications/mobile/validator | Validates the payload for the &#x60;POST /applications/mobile&#x60; request
[**ValidateCreateSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#ValidateCreateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties/validator | Validates the payload for the &#x60;POST /applications/mobile/{applicationId}/sessionProperties&#x60; request
[**ValidateUpdateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationApi.md#ValidateUpdateMobileApplicationConfig) | **Post** /applications/mobile/{id}/validator | Validates the payload for the &#x60;PUT /applications/mobile/{id}&#x60; request.
[**ValidateUpdateSessionProperty**](RUMMobileAndCustomApplicationConfigurationApi.md#ValidateUpdateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key}/validator | Validates the payload for the &#x60;PUT /applications/mobile/{applicationId}/sessionProperties/{key}&#x60; request



## CreateMobileApplicationConfig

> EntityShortRepresentation CreateMobileApplicationConfig(ctx).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Creates a new mobile or custom application



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
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig("Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains the parameters of the new mobile or custom application. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.CreateMobileApplicationConfig(context.Background()).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.CreateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.CreateMobileApplicationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains the parameters of the new mobile or custom application. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMobileKeyUserAction

> KeyUserActionMobile CreateMobileKeyUserAction(ctx, applicationId, actionName).Execute()

Marks the user action as a key user action in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    actionName := "actionName_example" // string | The name of the user action to be marked as a key user action.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.CreateMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.CreateMobileKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileKeyUserAction`: KeyUserActionMobile
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.CreateMobileKeyUserAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**actionName** | **string** | The name of the user action to be marked as a key user action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileKeyUserActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyUserActionMobile**](KeyUserActionMobile.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionProperty

> EntityShortRepresentation CreateSessionProperty(ctx, applicationId).MobileSessionProperty(mobileSessionProperty).Execute()

Creates a new session property for the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    mobileSessionProperty := *openapiclient.NewMobileSessionProperty("Key_example", "Type_example", "Origin_example") // MobileSessionProperty | The JSON body of the request. Contains the configuration of the mobile session property. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.CreateSessionProperty(context.Background(), applicationId).MobileSessionProperty(mobileSessionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.CreateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionProperty`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.CreateSessionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileSessionProperty** | [**MobileSessionProperty**](MobileSessionProperty.md) | The JSON body of the request. Contains the configuration of the mobile session property. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMobileApplicationConfig

> DeleteMobileApplicationConfig(ctx, id).Execute()

Deletes the specified mobile or custom application

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
    id := "id_example" // string | The ID of the application to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.DeleteMobileApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.DeleteMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobileApplicationConfigRequest struct via the builder pattern


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


## DeleteMobileKeyUserAction

> DeleteMobileKeyUserAction(ctx, applicationId, actionName).Execute()

Removes the specified user action from the list of key user actions in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    actionName := "actionName_example" // string | The ID of the user action to be removed from the key user actions list.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.DeleteMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.DeleteMobileKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**actionName** | **string** | The ID of the user action to be removed from the key user actions list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobileKeyUserActionRequest struct via the builder pattern


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


## DeleteSessionProperty

> DeleteSessionProperty(ctx, applicationId, key).Execute()

Deletes the specified session property for an application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required session property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.DeleteSessionProperty(context.Background(), applicationId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.DeleteSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required session property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionPropertyRequest struct via the builder pattern


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


## GetMobileApplicationConfig

> MobileCustomAppConfig GetMobileApplicationConfig(ctx, id).Execute()

Gets the configuration of the specified mobile or custom application

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
    id := "id_example" // string | The ID of the required mobile or custom application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.GetMobileApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.GetMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileApplicationConfig`: MobileCustomAppConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.GetMobileApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required mobile or custom application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileCustomAppConfig**](MobileCustomAppConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionProperty

> MobileSessionProperty GetSessionProperty(ctx, applicationId, key).Execute()

Gets the specified  session property of an application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required session property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.GetSessionProperty(context.Background(), applicationId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.GetSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionProperty`: MobileSessionProperty
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.GetSessionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required session property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MobileSessionProperty**](MobileSessionProperty.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMobileApplicationConfigs

> StubList ListMobileApplicationConfigs(ctx).Execute()

Lists all existing mobile and custom applications

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
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ListMobileApplicationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ListMobileApplicationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobileApplicationConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.ListMobileApplicationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMobileApplicationConfigsRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMobileKeyUserActions

> KeyUserActionMobileList ListMobileKeyUserActions(ctx, applicationId).Execute()

Gets the list of key user actions in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ListMobileKeyUserActions(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ListMobileKeyUserActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobileKeyUserActions`: KeyUserActionMobileList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.ListMobileKeyUserActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMobileKeyUserActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyUserActionMobileList**](KeyUserActionMobileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionProperties

> MobileSessionPropertyList ListSessionProperties(ctx, applicationId).Execute()

Lists all session properties for the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ListSessionProperties(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ListSessionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionProperties`: MobileSessionPropertyList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationApi.ListSessionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileSessionPropertyList**](MobileSessionPropertyList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMobileApplicationConfig

> UpdateMobileApplicationConfig(ctx, id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Updates the configuration of the specified mobile or custom application



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
    id := "id_example" // string | The ID of the application to be updated.
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig("Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains updated configuration of the mobile or custom application.   Do not set the identifier in the body. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.UpdateMobileApplicationConfig(context.Background(), id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.UpdateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains updated configuration of the mobile or custom application.   Do not set the identifier in the body. | 

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


## UpdateSessionProperty

> UpdateSessionProperty(ctx, applicationId, key).MobileSessionPropertyUpdate(mobileSessionPropertyUpdate).Execute()

Updates the specified session property for an application



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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required session property.
    mobileSessionPropertyUpdate := *openapiclient.NewMobileSessionPropertyUpdate("Type_example", "Origin_example") // MobileSessionPropertyUpdate | The JSON body of the request. Contains the configuration of the mobile session property. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.UpdateSessionProperty(context.Background(), applicationId, key).MobileSessionPropertyUpdate(mobileSessionPropertyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.UpdateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required session property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mobileSessionPropertyUpdate** | [**MobileSessionPropertyUpdate**](MobileSessionPropertyUpdate.md) | The JSON body of the request. Contains the configuration of the mobile session property. | 

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


## ValidateCreateMobileApplicationConfig

> ValidateCreateMobileApplicationConfig(ctx).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Validates the payload for the `POST /applications/mobile` request

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
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig("Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains the mobile or custom application configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ValidateCreateMobileApplicationConfig(context.Background()).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ValidateCreateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains the mobile or custom application configuration to be validated. | 

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


## ValidateCreateSessionProperty

> ValidateCreateSessionProperty(ctx, applicationId).MobileSessionProperty(mobileSessionProperty).Execute()

Validates the payload for the `POST /applications/mobile/{applicationId}/sessionProperties` request

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    mobileSessionProperty := *openapiclient.NewMobileSessionProperty("Key_example", "Type_example", "Origin_example") // MobileSessionProperty | The JSON body of the request. Contains the configuration of the mobile session property to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ValidateCreateSessionProperty(context.Background(), applicationId).MobileSessionProperty(mobileSessionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ValidateCreateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileSessionProperty** | [**MobileSessionProperty**](MobileSessionProperty.md) | The JSON body of the request. Contains the configuration of the mobile session property to be validated. | 

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


## ValidateUpdateMobileApplicationConfig

> ValidateUpdateMobileApplicationConfig(ctx, id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Validates the payload for the `PUT /applications/mobile/{id}` request.

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
    id := "id_example" // string | The ID of the mobile or custom application to be validated.
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig("Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains the mobile or custom application configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ValidateUpdateMobileApplicationConfig(context.Background(), id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ValidateUpdateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the mobile or custom application to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains the mobile or custom application configuration to be validated. | 

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


## ValidateUpdateSessionProperty

> ValidateUpdateSessionProperty(ctx, applicationId, key).MobileSessionPropertyUpdate(mobileSessionPropertyUpdate).Execute()

Validates the payload for the `PUT /applications/mobile/{applicationId}/sessionProperties/{key}` request

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required session property.
    mobileSessionPropertyUpdate := *openapiclient.NewMobileSessionPropertyUpdate("Type_example", "Origin_example") // MobileSessionPropertyUpdate | The JSON body of the request. Contains the configuration of the mobile session property to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMMobileAndCustomApplicationConfigurationApi.ValidateUpdateSessionProperty(context.Background(), applicationId, key).MobileSessionPropertyUpdate(mobileSessionPropertyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationApi.ValidateUpdateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required session property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mobileSessionPropertyUpdate** | [**MobileSessionPropertyUpdate**](MobileSessionPropertyUpdate.md) | The JSON body of the request. Contains the configuration of the mobile session property to be validated. | 

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

