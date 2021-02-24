# \RUMWebApplicationConfigurationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeyUserAction**](RUMWebApplicationConfigurationApi.md#CreateKeyUserAction) | **Post** /applications/web/{id}/keyUserActions | Marks the user action as a key user action in the specified web application
[**CreateOrUpdateDefaultConfiguration**](RUMWebApplicationConfigurationApi.md#CreateOrUpdateDefaultConfiguration) | **Put** /applications/web/default | Updates the configuration of the default web application
[**CreateWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#CreateWebApplicationConfig) | **Post** /applications/web | Creates a new web application
[**DeleteKeyUserAction**](RUMWebApplicationConfigurationApi.md#DeleteKeyUserAction) | **Delete** /applications/web/{id}/keyUserActions/{keyUserActionId} | Removes the specified user action from the list of key user actions in the specified web application
[**DeleteWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#DeleteWebApplicationConfig) | **Delete** /applications/web/{id} | Deletes the specified web application
[**GetApplicationErrorConfig**](RUMWebApplicationConfigurationApi.md#GetApplicationErrorConfig) | **Get** /applications/web/{id}/errorRules | Gets the configuration of error rules in the specified application
[**GetDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#GetDataPrivacySettings) | **Get** /applications/web/{id}/dataPrivacy | Gets the data privacy settings of the specified web application
[**GetDefaultApplication**](RUMWebApplicationConfigurationApi.md#GetDefaultApplication) | **Get** /applications/web/default | Gets the configuration of the default web application
[**GetDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#GetDefaultApplicationDataPrivacySettings) | **Get** /applications/web/default/dataPrivacy | Gets the data privacy settings of the default web application
[**GetWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#GetWebApplicationConfig) | **Get** /applications/web/{id} | Gets the configuration of the specified web application
[**ListDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#ListDataPrivacySettings) | **Get** /applications/web/dataPrivacy | Lists data privacy settings of all web applications
[**ListKeyUserActions**](RUMWebApplicationConfigurationApi.md#ListKeyUserActions) | **Get** /applications/web/{id}/keyUserActions | Gets the list of key user actions in the specified web application
[**ListWebApplicationConfigs**](RUMWebApplicationConfigurationApi.md#ListWebApplicationConfigs) | **Get** /applications/web | Lists all existing web applications
[**UpdateApplicationErrorConfig**](RUMWebApplicationConfigurationApi.md#UpdateApplicationErrorConfig) | **Put** /applications/web/{id}/errorRules | Updates the configuration of error rules in the specified application
[**UpdateDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#UpdateDataPrivacySettings) | **Put** /applications/web/{id}/dataPrivacy | Updates the data privacy settings of the specified web application
[**UpdateDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#UpdateDefaultApplicationDataPrivacySettings) | **Put** /applications/web/default/dataPrivacy | Updates the data privacy settings of the default web application
[**UpdateWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#UpdateWebApplicationConfig) | **Put** /applications/web/{id} | Updates the configuration of the specified web application or creates a new one
[**ValidateCreateWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#ValidateCreateWebApplicationConfig) | **Post** /applications/web/validator | Validates the configuration of the web application for the &#x60;POST /applications/web&#x60; request
[**ValidateDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#ValidateDataPrivacySettings) | **Post** /applications/web/{id}/dataPrivacy/validator | Validates data privacy settings for the &#x60;PUT /applications/web/{id}/dataPrivacy&#x60; request
[**ValidateDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationApi.md#ValidateDefaultApplicationDataPrivacySettings) | **Post** /applications/web/default/dataPrivacy/validator | Validates data privacy settings of the default web application for the &#x60;PUT /applications/web/default/dataPrivacy&#x60; request
[**ValidateDefaultConfiguration**](RUMWebApplicationConfigurationApi.md#ValidateDefaultConfiguration) | **Post** /applications/web/default/validator | Validates the configuration of the default web application for the &#x60;PUT /applications/web/default&#x60; request
[**ValidateUpdateWebApplicationConfig**](RUMWebApplicationConfigurationApi.md#ValidateUpdateWebApplicationConfig) | **Post** /applications/web/{id}/validator | Validates the configuration of the web application for the &#x60;PUT /applications/web/{id}&#x60; request.



## CreateKeyUserAction

> EntityShortRepresentation CreateKeyUserAction(ctx, id).KeyUserAction(keyUserAction).Execute()

Marks the user action as a key user action in the specified web application

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
    id := "id_example" // string | The ID of the required web application.
    keyUserAction := *openapiclient.NewKeyUserAction("Name_example", "ActionType_example") // KeyUserAction | The JSON body of the request. Contains the action to be marked as a key user action. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.CreateKeyUserAction(context.Background(), id).KeyUserAction(keyUserAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.CreateKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeyUserAction`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.CreateKeyUserAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyUserActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyUserAction** | [**KeyUserAction**](KeyUserAction.md) | The JSON body of the request. Contains the action to be marked as a key user action. | 

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


## CreateOrUpdateDefaultConfiguration

> CreateOrUpdateDefaultConfiguration(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Updates the configuration of the default web application



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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing the new parameters of the default web application. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.CreateOrUpdateDefaultConfiguration(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.CreateOrUpdateDefaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDefaultConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the new parameters of the default web application. | 

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


## CreateWebApplicationConfig

> EntityShortRepresentation CreateWebApplicationConfig(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Creates a new web application



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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing parameters of the new web application configuraiton. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.CreateWebApplicationConfig(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.CreateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.CreateWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing parameters of the new web application configuraiton. | 

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


## DeleteKeyUserAction

> DeleteKeyUserAction(ctx, id, keyUserActionId).Execute()

Removes the specified user action from the list of key user actions in the specified web application

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
    id := "id_example" // string | The ID of the required web application.
    keyUserActionId := "keyUserActionId_example" // string | The ID of the user action to be removed from the key user actions list.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.DeleteKeyUserAction(context.Background(), id, keyUserActionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.DeleteKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 
**keyUserActionId** | **string** | The ID of the user action to be removed from the key user actions list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyUserActionRequest struct via the builder pattern


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


## DeleteWebApplicationConfig

> DeleteWebApplicationConfig(ctx, id).Execute()

Deletes the specified web application

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
    id := "id_example" // string | The ID of the web application to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.DeleteWebApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.DeleteWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebApplicationConfigRequest struct via the builder pattern


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


## GetApplicationErrorConfig

> ApplicationErrorRules GetApplicationErrorConfig(ctx, id).Execute()

Gets the configuration of error rules in the specified application

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
    id := "id_example" // string | The ID of the required web application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.GetApplicationErrorConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.GetApplicationErrorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationErrorConfig`: ApplicationErrorRules
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.GetApplicationErrorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationErrorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationErrorRules**](ApplicationErrorRules.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataPrivacySettings

> ApplicationDataPrivacy GetDataPrivacySettings(ctx, id).Execute()

Gets the data privacy settings of the specified web application

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
    id := "id_example" // string | The ID of the web application where you want to check data privacy settings.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.GetDataPrivacySettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.GetDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataPrivacySettings`: ApplicationDataPrivacy
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.GetDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application where you want to check data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplication

> WebApplicationConfig GetDefaultApplication(ctx).Execute()

Gets the configuration of the default web application



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
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.GetDefaultApplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.GetDefaultApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultApplication`: WebApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.GetDefaultApplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultApplicationRequest struct via the builder pattern


### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplicationDataPrivacySettings

> ApplicationDataPrivacy GetDefaultApplicationDataPrivacySettings(ctx).Execute()

Gets the data privacy settings of the default web application



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
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.GetDefaultApplicationDataPrivacySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.GetDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultApplicationDataPrivacySettings`: ApplicationDataPrivacy
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.GetDefaultApplicationDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebApplicationConfig

> WebApplicationConfig GetWebApplicationConfig(ctx, id).Execute()

Gets the configuration of the specified web application

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
    id := "id_example" // string | The ID of the requested web application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.GetWebApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.GetWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebApplicationConfig`: WebApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.GetWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the requested web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataPrivacySettings

> ApplicationDataPrivacyList ListDataPrivacySettings(ctx).Execute()

Lists data privacy settings of all web applications

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
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ListDataPrivacySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ListDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataPrivacySettings`: ApplicationDataPrivacyList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.ListDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataPrivacySettingsRequest struct via the builder pattern


### Return type

[**ApplicationDataPrivacyList**](ApplicationDataPrivacyList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyUserActions

> KeyUserActionList ListKeyUserActions(ctx, id).Execute()

Gets the list of key user actions in the specified web application

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
    id := "id_example" // string | The ID of the required web application.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ListKeyUserActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ListKeyUserActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeyUserActions`: KeyUserActionList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.ListKeyUserActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyUserActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyUserActionList**](KeyUserActionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebApplicationConfigs

> StubList ListWebApplicationConfigs(ctx).Execute()

Lists all existing web applications

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
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ListWebApplicationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ListWebApplicationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebApplicationConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.ListWebApplicationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebApplicationConfigsRequest struct via the builder pattern


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


## UpdateApplicationErrorConfig

> UpdateApplicationErrorConfig(ctx, id).ApplicationErrorRules(applicationErrorRules).Execute()

Updates the configuration of error rules in the specified application

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
    id := "id_example" // string | The ID of the required web application.
    applicationErrorRules := *openapiclient.NewApplicationErrorRules(false, false, false, []openapiclient.HttpErrorRule{*openapiclient.NewHttpErrorRule(false, false, false, false, false)}, []openapiclient.CustomErrorRule{*openapiclient.NewCustomErrorRule("KeyPattern_example", "KeyMatcher_example", false, false, false)}) // ApplicationErrorRules | The JSON body of the request. Contains the updated configuration of error rules. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.UpdateApplicationErrorConfig(context.Background(), id).ApplicationErrorRules(applicationErrorRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.UpdateApplicationErrorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationErrorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationErrorRules** | [**ApplicationErrorRules**](ApplicationErrorRules.md) | The JSON body of the request. Contains the updated configuration of error rules. | 

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


## UpdateDataPrivacySettings

> UpdateDataPrivacySettings(ctx, id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Updates the data privacy settings of the specified web application

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
    id := "id_example" // string | The ID of the web application, where you want to update data privacy settings.
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, false, "DoNotTrackBehaviour_example") // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.UpdateDataPrivacySettings(context.Background(), id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.UpdateDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application, where you want to update data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## UpdateDefaultApplicationDataPrivacySettings

> UpdateDefaultApplicationDataPrivacySettings(ctx).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Updates the data privacy settings of the default web application



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
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, false, "DoNotTrackBehaviour_example") // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.UpdateDefaultApplicationDataPrivacySettings(context.Background()).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.UpdateDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## UpdateWebApplicationConfig

> EntityShortRepresentation UpdateWebApplicationConfig(ctx, id).WebApplicationConfig(webApplicationConfig).Execute()

Updates the configuration of the specified web application or creates a new one



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
    id := "id_example" // string | The ID of the web application to update.   If you set the ID in the body as well, it must match this ID.
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing updated configuration of the web application. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.UpdateWebApplicationConfig(context.Background(), id).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.UpdateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationApi.UpdateWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to update.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing updated configuration of the web application. | 

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


## ValidateCreateWebApplicationConfig

> ValidateCreateWebApplicationConfig(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the web application for the `POST /applications/web` request

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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing the web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ValidateCreateWebApplicationConfig(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ValidateCreateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the web application configuration to validate. | 

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


## ValidateDataPrivacySettings

> ValidateDataPrivacySettings(ctx, id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Validates data privacy settings for the `PUT /applications/web/{id}/dataPrivacy` request

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
    id := "id_example" // string | The ID of the web application, where you want to validate data privacy settings.
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, false, "DoNotTrackBehaviour_example") // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ValidateDataPrivacySettings(context.Background(), id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ValidateDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application, where you want to validate data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## ValidateDefaultApplicationDataPrivacySettings

> ValidateDefaultApplicationDataPrivacySettings(ctx).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Validates data privacy settings of the default web application for the `PUT /applications/web/default/dataPrivacy` request

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
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, false, "DoNotTrackBehaviour_example") // ApplicationDataPrivacy | JSON body of the request, containing the data privacy settings to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ValidateDefaultApplicationDataPrivacySettings(context.Background()).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ValidateDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing the data privacy settings to validate. | 

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


## ValidateDefaultConfiguration

> ValidateDefaultConfiguration(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the default web application for the `PUT /applications/web/default` request

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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ValidateDefaultConfiguration(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ValidateDefaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDefaultConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing web application configuration to validate. | 

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


## ValidateUpdateWebApplicationConfig

> ValidateUpdateWebApplicationConfig(ctx, id).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the web application for the `PUT /applications/web/{id}` request.

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
    id := "id_example" // string | The ID of the web application to validate.
    webApplicationConfig := *openapiclient.NewWebApplicationConfig("Name_example", false, float32(123), "LoadActionKeyPerformanceMetric_example", "XhrActionKeyPerformanceMetric_example", *openapiclient.NewApdex(), *openapiclient.NewApdex(), , *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewMonitoringSettings(false, false, *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), *openapiclient.NewContentCapture(*openapiclient.NewResourceTimingSettings(false, false, int32(123), "ResourceTimingCaptureType_example", int32(123)), false, *openapiclient.NewTimeoutSettings(false, int32(123), int32(123)), false), "ExcludeXhrRegex_example", "InjectionMode_example", "CustomConfigurationProperties_example", "ServerRequestPathId_example", false, "CookiePlacementDomain_example", false, *openapiclient.NewAdvancedJavaScriptTagSettings(false, "SpecialCharactersToEscape_example", int32(123), int32(123), *openapiclient.NewAdditionalEventHandlers(false, false, false, false, false, false, int32(123)), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings(false, false, false, false, false, false, false, "AdditionalEventCapturedAsUserInput_example")))) // WebApplicationConfig | JSON body of the request, containing the web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RUMWebApplicationConfigurationApi.ValidateUpdateWebApplicationConfig(context.Background(), id).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationApi.ValidateUpdateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the web application configuration to validate. | 

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

