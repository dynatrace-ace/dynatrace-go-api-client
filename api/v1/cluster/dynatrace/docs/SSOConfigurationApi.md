# \SSOConfigurationApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSamlSpCert**](SSOConfigurationApi.md#GetSamlSpCert) | **Get** /sso/saml/sp/cert | Get SAML SP X.509 certificate details
[**UpdateSamlSpCert**](SSOConfigurationApi.md#UpdateSamlSpCert) | **Put** /sso/saml/sp/cert | Update SAML SP X.509 certificate



## GetSamlSpCert

> CertificateDetails GetSamlSpCert(ctx).Execute()

Get SAML SP X.509 certificate details

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
    resp, r, err := api_client.SSOConfigurationApi.GetSamlSpCert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSOConfigurationApi.GetSamlSpCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSamlSpCert`: CertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `SSOConfigurationApi.GetSamlSpCert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlSpCertRequest struct via the builder pattern


### Return type

[**CertificateDetails**](CertificateDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSamlSpCert

> UpdateSamlSpCert(ctx).CertDto(certDto).Execute()

Update SAML SP X.509 certificate

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
    certDto := *openapiclient.NewCertDto("PrivateKeyEncoded_example", "PublicKeyCertificateEncoded_example") // CertDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SSOConfigurationApi.UpdateSamlSpCert(context.Background()).CertDto(certDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSOConfigurationApi.UpdateSamlSpCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSamlSpCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certDto** | [**CertDto**](CertDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

