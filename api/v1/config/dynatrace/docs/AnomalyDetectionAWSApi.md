# \AnomalyDetectionAWSApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#GetAwsAnomalyDetectionConfig) | **Get** /anomalyDetection/aws | Gets the configuration of anomaly detection for AWS
[**UpdateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#UpdateAwsAnomalyDetectionConfig) | **Put** /anomalyDetection/aws | Updates the configuration of anomaly detection for AWS
[**ValidateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSApi.md#ValidateAwsAnomalyDetectionConfig) | **Post** /anomalyDetection/aws/validator | Validates the configuration of anomaly detection for AWS for the &#x60;PUT /anomalyDetection/aws&#x60; request



## GetAwsAnomalyDetectionConfig

> AwsAnomalyDetectionConfig GetAwsAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for AWS

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
    resp, r, err := api_client.AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAwsAnomalyDetectionConfig`: AwsAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsAnomalyDetectionConfig

> UpdateAwsAnomalyDetectionConfig(ctx).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for AWS

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
    awsAnomalyDetectionConfig := *openapiclient.NewAwsAnomalyDetectionConfig(*openapiclient.NewRdsHighCpuDetectionConfig(false), *openapiclient.NewRdsHighWriteReadLatencyDetectionConfig(false), *openapiclient.NewRdsLowStorageDetectionConfig(false), *openapiclient.NewRdsHighMemoryDetectionConfig(false), *openapiclient.NewElbHighConnectionErrorsDetectionConfig(false), *openapiclient.NewRdsRestartsSequenceDetectionConfig(false), *openapiclient.NewLambdaHighErrorRateDetectionConfig(false)) // AwsAnomalyDetectionConfig | JSON body of the request, containing parameters of the AWS anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionAWSApi.UpdateAwsAnomalyDetectionConfig(context.Background()).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSApi.UpdateAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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


## ValidateAwsAnomalyDetectionConfig

> ValidateAwsAnomalyDetectionConfig(ctx).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for AWS for the `PUT /anomalyDetection/aws` request

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
    awsAnomalyDetectionConfig := *openapiclient.NewAwsAnomalyDetectionConfig(*openapiclient.NewRdsHighCpuDetectionConfig(false), *openapiclient.NewRdsHighWriteReadLatencyDetectionConfig(false), *openapiclient.NewRdsLowStorageDetectionConfig(false), *openapiclient.NewRdsHighMemoryDetectionConfig(false), *openapiclient.NewElbHighConnectionErrorsDetectionConfig(false), *openapiclient.NewRdsRestartsSequenceDetectionConfig(false), *openapiclient.NewLambdaHighErrorRateDetectionConfig(false)) // AwsAnomalyDetectionConfig | JSON body of the request, containing parameters of the AWS anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionAWSApi.ValidateAwsAnomalyDetectionConfig(context.Background()).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSApi.ValidateAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAwsAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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

