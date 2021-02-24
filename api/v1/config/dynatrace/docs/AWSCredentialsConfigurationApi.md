# \AWSCredentialsConfigurationApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#CreateAwsCredentialsConfig) | **Post** /aws/credentials | Creates a new AWS credentials configuration
[**DeleteAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#DeleteAwsCredentialsConfig) | **Delete** /aws/credentials/{id} | Deletes the specified AWS credentials configuration
[**GetAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#GetAwsCredentialsConfig) | **Get** /aws/credentials/{id} | Gets the configuration of the specified AWS credentials
[**ListAwsCredentialConfigs**](AWSCredentialsConfigurationApi.md#ListAwsCredentialConfigs) | **Get** /aws/credentials | Lists all available AWS credentials configurations
[**ReadIamExternalIdToken**](AWSCredentialsConfigurationApi.md#ReadIamExternalIdToken) | **Get** /aws/iamExternalId | Gets the external ID token for setting an IAM role
[**UpdateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#UpdateAwsCredentialsConfig) | **Put** /aws/credentials/{id} | Updates an existing AWS credentials configuration
[**ValidateCreateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#ValidateCreateAwsCredentialsConfig) | **Post** /aws/credentials/validator | Validates the payload for the &#x60;POST /aws/credentials&#x60; request
[**ValidateUpdateAwsCredentialsConfig**](AWSCredentialsConfigurationApi.md#ValidateUpdateAwsCredentialsConfig) | **Post** /aws/credentials/{id}/validator | Validates the payload for the &#x60;PUT /aws/credentials/{id}&#x60; request



## CreateAwsCredentialsConfig

> EntityShortRepresentation CreateAwsCredentialsConfig(ctx).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Creates a new AWS credentials configuration



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
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig("Label_example", "PartitionType_example", *openapiclient.NewAwsAuthenticationData("Type_example"), false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains parameters of the new AWS credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.CreateAwsCredentialsConfig(context.Background()).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.CreateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAwsCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationApi.CreateAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains parameters of the new AWS credentials configuration. | 

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


## DeleteAwsCredentialsConfig

> DeleteAwsCredentialsConfig(ctx, id).Execute()

Deletes the specified AWS credentials configuration

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
    id := "id_example" // string | The ID of AWS credentials configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.DeleteAwsCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.DeleteAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of AWS credentials configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAwsCredentialsConfig

> AwsCredentialsConfig GetAwsCredentialsConfig(ctx, id).Execute()

Gets the configuration of the specified AWS credentials

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
    id := "id_example" // string | The ID of the specified AWS credentials configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.GetAwsCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.GetAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAwsCredentialsConfig`: AwsCredentialsConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationApi.GetAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the specified AWS credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsCredentialsConfig**](AwsCredentialsConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAwsCredentialConfigs

> []EntityShortRepresentation ListAwsCredentialConfigs(ctx).Execute()

Lists all available AWS credentials configurations

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
    resp, r, err := api_client.AWSCredentialsConfigurationApi.ListAwsCredentialConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.ListAwsCredentialConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsCredentialConfigs`: []EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationApi.ListAwsCredentialConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsCredentialConfigsRequest struct via the builder pattern


### Return type

[**[]EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIamExternalIdToken

> AwsIamToken ReadIamExternalIdToken(ctx).Execute()

Gets the external ID token for setting an IAM role



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
    resp, r, err := api_client.AWSCredentialsConfigurationApi.ReadIamExternalIdToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.ReadIamExternalIdToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIamExternalIdToken`: AwsIamToken
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationApi.ReadIamExternalIdToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadIamExternalIdTokenRequest struct via the builder pattern


### Return type

[**AwsIamToken**](AwsIamToken.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsCredentialsConfig

> EntityShortRepresentation UpdateAwsCredentialsConfig(ctx, id).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Updates an existing AWS credentials configuration

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
    id := "id_example" // string | The ID of the AWS credentials configuration to be updated.
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig("Label_example", "PartitionType_example", *openapiclient.NewAwsAuthenticationData("Type_example"), false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains updated parameters of the AWS credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.UpdateAwsCredentialsConfig(context.Background(), id).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.UpdateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAwsCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationApi.UpdateAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains updated parameters of the AWS credentials configuration. | 

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


## ValidateCreateAwsCredentialsConfig

> ValidateCreateAwsCredentialsConfig(ctx).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Validates the payload for the `POST /aws/credentials` request

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
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig("Label_example", "PartitionType_example", *openapiclient.NewAwsAuthenticationData("Type_example"), false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains the AWS credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.ValidateCreateAwsCredentialsConfig(context.Background()).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.ValidateCreateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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


## ValidateUpdateAwsCredentialsConfig

> ValidateUpdateAwsCredentialsConfig(ctx, id).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Validates the payload for the `PUT /aws/credentials/{id}` request

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
    id := "id_example" // string | The ID of the AWS credentials configuration to be validated.
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig("Label_example", "PartitionType_example", *openapiclient.NewAwsAuthenticationData("Type_example"), false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains the AWS credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AWSCredentialsConfigurationApi.ValidateUpdateAwsCredentialsConfig(context.Background(), id).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationApi.ValidateUpdateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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

