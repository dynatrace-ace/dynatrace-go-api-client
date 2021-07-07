# \SSLCertificatesApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSslCertificateDetails**](SSLCertificatesApi.md#GetSslCertificateDetails) | **Get** /sslCertificate/{entityType}/{entityId} | SSL certificate details.
[**GetStoringSslCertificateStatus**](SSLCertificatesApi.md#GetStoringSslCertificateStatus) | **Get** /sslCertificate/store/{entityType}/{entityId} | Get certificate-storage status
[**StoreSslCertificateStatus**](SSLCertificatesApi.md#StoreSslCertificateStatus) | **Post** /sslCertificate/store/{entityType}/{entityId} | Store SSL certificate status



## GetSslCertificateDetails

> SSLDetails GetSslCertificateDetails(ctx, entityType, entityId).Execute()

SSL certificate details.

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
    entityType := "entityType_example" // string | entity type, possible values = \"SERVER, COLLECTOR\" 
    entityId := int32(56) // int32 | Node ID, which can be extracted from the URL in 'Node details' view.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.GetSslCertificateDetails(context.Background(), entityType, entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.GetSslCertificateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslCertificateDetails`: SSLDetails
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.GetSslCertificateDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | entity type, possible values &#x3D; \&quot;SERVER, COLLECTOR\&quot;  | 
**entityId** | **int32** | Node ID, which can be extracted from the URL in &#39;Node details&#39; view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslCertificateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SSLDetails**](SSLDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoringSslCertificateStatus

> CertificateStoreStatus GetStoringSslCertificateStatus(ctx, entityType, entityId).Execute()

Get certificate-storage status

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
    entityType := "entityType_example" // string | entity type, possible values = \"COLLECTOR\" 
    entityId := int32(56) // int32 | Node ID, which can be extracted from the URL in 'Node details' view.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.GetStoringSslCertificateStatus(context.Background(), entityType, entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.GetStoringSslCertificateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoringSslCertificateStatus`: CertificateStoreStatus
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.GetStoringSslCertificateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | entity type, possible values &#x3D; \&quot;COLLECTOR\&quot;  | 
**entityId** | **int32** | Node ID, which can be extracted from the URL in &#39;Node details&#39; view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoringSslCertificateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificateStoreStatus**](CertificateStoreStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreSslCertificateStatus

> CertificateStoreStatus StoreSslCertificateStatus(ctx, entityType, entityId).SslCertDto(sslCertDto).Execute()

Store SSL certificate status

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
    entityType := "entityType_example" // string | entity type, possible values = \"SERVER, COLLECTOR\" 
    entityId := int32(56) // int32 | Node ID, which can be extracted from the URL in 'Node details' view.
    sslCertDto := *openapiclient.NewSslCertDto("PrivateKeyEncoded_example", "PublicKeyCertificateEncoded_example") // SslCertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSLCertificatesApi.StoreSslCertificateStatus(context.Background(), entityType, entityId).SslCertDto(sslCertDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLCertificatesApi.StoreSslCertificateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreSslCertificateStatus`: CertificateStoreStatus
    fmt.Fprintf(os.Stdout, "Response from `SSLCertificatesApi.StoreSslCertificateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | entity type, possible values &#x3D; \&quot;SERVER, COLLECTOR\&quot;  | 
**entityId** | **int32** | Node ID, which can be extracted from the URL in &#39;Node details&#39; view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreSslCertificateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sslCertDto** | [**SslCertDto**](SslCertDto.md) |  | 

### Return type

[**CertificateStoreStatus**](CertificateStoreStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

