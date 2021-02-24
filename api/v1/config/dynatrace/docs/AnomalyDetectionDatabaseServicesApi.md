# \AnomalyDetectionDatabaseServicesApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesApi.md#GetDatabaseServiceAnomalyDetectionConfig) | **Get** /anomalyDetection/databaseServices | Gets the configuration of anomaly detection for database services
[**UpdateDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesApi.md#UpdateDatabaseServiceAnomalyDetectionConfig) | **Put** /anomalyDetection/databaseServices | Updates the configuration of anomaly detection for database services
[**ValidateDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesApi.md#ValidateDatabaseServiceAnomalyDetectionConfig) | **Post** /anomalyDetection/databaseServices/validator | Validates the payload for the &#x60;PUT /anomalyDetection/databaseServices&#x60; request



## GetDatabaseServiceAnomalyDetectionConfig

> DatabaseAnomalyDetectionConfig GetDatabaseServiceAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for database services

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
    resp, r, err := api_client.AnomalyDetectionDatabaseServicesApi.GetDatabaseServiceAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesApi.GetDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServiceAnomalyDetectionConfig`: DatabaseAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDatabaseServicesApi.GetDatabaseServiceAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseServiceAnomalyDetectionConfig

> UpdateDatabaseServiceAnomalyDetectionConfig(ctx).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for database services

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
    databaseAnomalyDetectionConfig := *openapiclient.NewDatabaseAnomalyDetectionConfig(*openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example"), *openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewDatabaseConnectionFailureDetectionConfig(false)) // DatabaseAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDatabaseServicesApi.UpdateDatabaseServiceAnomalyDetectionConfig(context.Background()).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesApi.UpdateDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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


## ValidateDatabaseServiceAnomalyDetectionConfig

> ValidateDatabaseServiceAnomalyDetectionConfig(ctx).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()

Validates the payload for the `PUT /anomalyDetection/databaseServices` request

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
    databaseAnomalyDetectionConfig := *openapiclient.NewDatabaseAnomalyDetectionConfig(*openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example"), *openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewDatabaseConnectionFailureDetectionConfig(false)) // DatabaseAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionDatabaseServicesApi.ValidateDatabaseServiceAnomalyDetectionConfig(context.Background()).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesApi.ValidateDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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

