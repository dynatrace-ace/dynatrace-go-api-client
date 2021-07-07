# \SettingsObjectsApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSettingsObjectByObjectId**](SettingsObjectsApi.md#DeleteSettingsObjectByObjectId) | **Delete** /settings/objects/{objectId} | Deletes the specified settings object | maturity&#x3D;EARLY_ADOPTER
[**GetSettingsObjectByObjectId**](SettingsObjectsApi.md#GetSettingsObjectByObjectId) | **Get** /settings/objects/{objectId} | Gets the specified settings object | maturity&#x3D;EARLY_ADOPTER
[**GetSettingsObjects**](SettingsObjectsApi.md#GetSettingsObjects) | **Get** /settings/objects | Lists available settings objects | maturity&#x3D;EARLY_ADOPTER
[**PostSettingsObjects**](SettingsObjectsApi.md#PostSettingsObjects) | **Post** /settings/objects | Creates a new settings object | maturity&#x3D;EARLY_ADOPTER
[**PutSettingsObjectByObjectId**](SettingsObjectsApi.md#PutSettingsObjectByObjectId) | **Put** /settings/objects/{objectId} | Updates an existing settings object | maturity&#x3D;EARLY_ADOPTER



## DeleteSettingsObjectByObjectId

> DeleteSettingsObjectByObjectId(ctx, objectId).UpdateToken(updateToken).Execute()

Deletes the specified settings object | maturity=EARLY_ADOPTER

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
    objectId := "objectId_example" // string | The ID of the required settings object.
    updateToken := "updateToken_example" // string | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn't any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsObjectsApi.DeleteSettingsObjectByObjectId(context.Background(), objectId).UpdateToken(updateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsApi.DeleteSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateToken** | **string** | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn&#39;t any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsObjectByObjectId

> SettingsObject GetSettingsObjectByObjectId(ctx, objectId).Execute()

Gets the specified settings object | maturity=EARLY_ADOPTER

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
    objectId := "objectId_example" // string | The ID of the required settings object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsObjectsApi.GetSettingsObjectByObjectId(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsApi.GetSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingsObjectByObjectId`: SettingsObject
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsApi.GetSettingsObjectByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsObject**](SettingsObject.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsObjects

> ObjectsList GetSettingsObjects(ctx).SchemaIds(schemaIds).Scopes(scopes).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists available settings objects | maturity=EARLY_ADOPTER

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
    schemaIds := "schemaIds_example" // string | A list of comma-separated schema IDs to which the requested objects belong.   To load the first page, when the **nextPageKey** is not set, this parameter is required.  (optional)
    scopes := "scopes_example" // string | A list of comma-separated scopes, that the requested objects target.   The selection only matches objects directly targeting the specified scopes. For example, `environment` will not match objects that target a host within environment.   To load the first page, when the **nextPageKey** is not set, this parameter is required.  (optional)
    fields := "fields_example" // string | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, `objectId,value`). (optional) (default to "objectId,value")
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsObjectsApi.GetSettingsObjects(context.Background()).SchemaIds(schemaIds).Scopes(scopes).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsApi.GetSettingsObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingsObjects`: ObjectsList
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsApi.GetSettingsObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaIds** | **string** | A list of comma-separated schema IDs to which the requested objects belong.   To load the first page, when the **nextPageKey** is not set, this parameter is required.  | 
 **scopes** | **string** | A list of comma-separated scopes, that the requested objects target.   The selection only matches objects directly targeting the specified scopes. For example, &#x60;environment&#x60; will not match objects that target a host within environment.   To load the first page, when the **nextPageKey** is not set, this parameter is required.  | 
 **fields** | **string** | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, &#x60;objectId,value&#x60;). | [default to &quot;objectId,value&quot;]
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

### Return type

[**ObjectsList**](ObjectsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsObjects

> []SettingsObjectResponse PostSettingsObjects(ctx).ValidateOnly(validateOnly).SettingsObjectCreate(settingsObjectCreate).Execute()

Creates a new settings object | maturity=EARLY_ADOPTER



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
    validateOnly := true // bool | If `true`, the request runs only validation of the submitted settings objects, without saving them. (optional) (default to false)
    settingsObjectCreate := []openapiclient.SettingsObjectCreate{*openapiclient.NewSettingsObjectCreate("HOST-D3A3C5A146830A79", "builtin:anomaly.infrastructure", map[string]interface{}({"autoMonitoring":true}))} // []SettingsObjectCreate | The JSON body of the request. Contains the settings objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsObjectsApi.PostSettingsObjects(context.Background()).ValidateOnly(validateOnly).SettingsObjectCreate(settingsObjectCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsApi.PostSettingsObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSettingsObjects`: []SettingsObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsApi.PostSettingsObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateOnly** | **bool** | If &#x60;true&#x60;, the request runs only validation of the submitted settings objects, without saving them. | [default to false]
 **settingsObjectCreate** | [**[]SettingsObjectCreate**](SettingsObjectCreate.md) | The JSON body of the request. Contains the settings objects. | 

### Return type

[**[]SettingsObjectResponse**](SettingsObjectResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsObjectByObjectId

> SettingsObjectResponse PutSettingsObjectByObjectId(ctx, objectId).SettingsObjectUpdate(settingsObjectUpdate).Execute()

Updates an existing settings object | maturity=EARLY_ADOPTER

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
    objectId := "objectId_example" // string | The ID of the required settings object.
    settingsObjectUpdate := *openapiclient.NewSettingsObjectUpdate(map[string]interface{}({"autoMonitoring":true})) // SettingsObjectUpdate | The JSON body of the request. Contains updated parameters of the settings object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsObjectsApi.PutSettingsObjectByObjectId(context.Background(), objectId).SettingsObjectUpdate(settingsObjectUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsApi.PutSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSettingsObjectByObjectId`: SettingsObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsApi.PutSettingsObjectByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsObjectUpdate** | [**SettingsObjectUpdate**](SettingsObjectUpdate.md) | The JSON body of the request. Contains updated parameters of the settings object. | 

### Return type

[**SettingsObjectResponse**](SettingsObjectResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

