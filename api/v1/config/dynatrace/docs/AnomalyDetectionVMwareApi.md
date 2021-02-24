# \AnomalyDetectionVMwareApi

All URIs are relative to *http://https:/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#GetVMwareAnomalyDetectionConfig) | **Get** /anomalyDetection/vmware | Gets the configuration of anomaly detection for VMware
[**UpdateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#UpdateVMwareAnomalyDetectionConfig) | **Put** /anomalyDetection/vmware | Updates the configuration of anomaly detection for VMware
[**ValidateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareApi.md#ValidateVMwareAnomalyDetectionConfig) | **Post** /anomalyDetection/vmware/validator | Validates the configuration of anomaly detection for VMware for the &#x60;PUT /anomalyDetection/vmware&#x60; request



## GetVMwareAnomalyDetectionConfig

> VMwareAnomalyDetectionConfig GetVMwareAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for VMware

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
    resp, r, err := api_client.AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMwareAnomalyDetectionConfig`: VMwareAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMwareAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMwareAnomalyDetectionConfig

> UpdateVMwareAnomalyDetectionConfig(ctx).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for VMware

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
    vMwareAnomalyDetectionConfig := *openapiclient.NewVMwareAnomalyDetectionConfig(*openapiclient.NewEsxiHighCpuSaturationConfig(false), *openapiclient.NewEsxiHighMemoryDetectionConfig(false), *openapiclient.NewOverloadedStorageDetectionConfig(false), *openapiclient.NewUndersizedStorageDetectionConfig(false), *openapiclient.NewSlowPhysicalStorageDetectionConfig(false), *openapiclient.NewDroppedPacketsDetectionConfig(false), *openapiclient.NewLowDatastoreSpaceDetectionConfig(false)) // VMwareAnomalyDetectionConfig | JSON body of the request, containing parameters of the VMware anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionVMwareApi.UpdateVMwareAnomalyDetectionConfig(context.Background()).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareApi.UpdateVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMwareAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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


## ValidateVMwareAnomalyDetectionConfig

> ValidateVMwareAnomalyDetectionConfig(ctx).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for VMware for the `PUT /anomalyDetection/vmware` request

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
    vMwareAnomalyDetectionConfig := *openapiclient.NewVMwareAnomalyDetectionConfig(*openapiclient.NewEsxiHighCpuSaturationConfig(false), *openapiclient.NewEsxiHighMemoryDetectionConfig(false), *openapiclient.NewOverloadedStorageDetectionConfig(false), *openapiclient.NewUndersizedStorageDetectionConfig(false), *openapiclient.NewSlowPhysicalStorageDetectionConfig(false), *openapiclient.NewDroppedPacketsDetectionConfig(false), *openapiclient.NewLowDatastoreSpaceDetectionConfig(false)) // VMwareAnomalyDetectionConfig | JSON body of the request, containing parameters of the VMware anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnomalyDetectionVMwareApi.ValidateVMwareAnomalyDetectionConfig(context.Background()).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareApi.ValidateVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateVMwareAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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

