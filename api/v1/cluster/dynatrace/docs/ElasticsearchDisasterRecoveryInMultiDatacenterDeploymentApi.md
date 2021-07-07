# \ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecoverStatus**](ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.md#GetRecoverStatus) | **Get** /multiDc/restore/elasticsearch/recover | Get status of disaster recovery. | maturity&#x3D;EARLY_ADOPTER
[**StartComponentsAfterMigration**](ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.md#StartComponentsAfterMigration) | **Get** /multiDc/restore/server/recovery/{requestId} | Get status after server migration and start components in disaster recovery | maturity&#x3D;EARLY_ADOPTER
[**StartComponentsAfterMigration1**](ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.md#StartComponentsAfterMigration1) | **Post** /multiDc/restore/server/recovery | Migrate servers and start components after disaster recovery | maturity&#x3D;EARLY_ADOPTER
[**StartRecover**](ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.md#StartRecover) | **Post** /multiDc/restore/elasticsearch/recover | Recover elasticsearch from a snapshot on one datacenter. | maturity&#x3D;EARLY_ADOPTER



## GetRecoverStatus

> GetRecoverStatus(ctx).Execute()

Get status of disaster recovery. | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.GetRecoverStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.GetRecoverStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoverStatusRequest struct via the builder pattern


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


## StartComponentsAfterMigration

> StartComponentsAfterMigration(ctx, requestId).Execute()

Get status after server migration and start components in disaster recovery | maturity=EARLY_ADOPTER

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
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartComponentsAfterMigration(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartComponentsAfterMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartComponentsAfterMigrationRequest struct via the builder pattern


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


## StartComponentsAfterMigration1

> StartComponentsAfterMigration1(ctx).Execute()

Migrate servers and start components after disaster recovery | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartComponentsAfterMigration1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartComponentsAfterMigration1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartComponentsAfterMigration1Request struct via the builder pattern


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


## StartRecover

> StartRecover(ctx).Execute()

Recover elasticsearch from a snapshot on one datacenter. | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartRecover(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchDisasterRecoveryInMultiDatacenterDeploymentApi.StartRecover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartRecoverRequest struct via the builder pattern


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

