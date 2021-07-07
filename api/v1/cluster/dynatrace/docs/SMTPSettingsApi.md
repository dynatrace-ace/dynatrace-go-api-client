# \SMTPSettingsApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmtpConfiguration**](SMTPSettingsApi.md#GetSmtpConfiguration) | **Get** /smtp | Get SMTP configuration
[**SendTestEmail**](SMTPSettingsApi.md#SendTestEmail) | **Post** /smtp/sendTestMessage/{emailAddress} | Send test email
[**UpdateSmtpConfiguration**](SMTPSettingsApi.md#UpdateSmtpConfiguration) | **Post** /smtp | Update SMTP configuration



## GetSmtpConfiguration

> SmtpConfiguration GetSmtpConfiguration(ctx).Execute()

Get SMTP configuration

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
    resp, r, err := api_client.SMTPSettingsApi.GetSmtpConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPSettingsApi.GetSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmtpConfiguration`: SmtpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SMTPSettingsApi.GetSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpConfigurationRequest struct via the builder pattern


### Return type

[**SmtpConfiguration**](SmtpConfiguration.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestEmail

> SendTestEmail(ctx, emailAddress).Execute()

Send test email



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
    emailAddress := "emailAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPSettingsApi.SendTestEmail(context.Background(), emailAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPSettingsApi.SendTestEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmtpConfiguration

> SmtpConfiguration UpdateSmtpConfiguration(ctx).SmtpConfiguration(smtpConfiguration).Execute()

Update SMTP configuration

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
    smtpConfiguration := *openapiclient.NewSmtpConfiguration() // SmtpConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SMTPSettingsApi.UpdateSmtpConfiguration(context.Background()).SmtpConfiguration(smtpConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMTPSettingsApi.UpdateSmtpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmtpConfiguration`: SmtpConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SMTPSettingsApi.UpdateSmtpConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmtpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpConfiguration** | [**SmtpConfiguration**](SmtpConfiguration.md) |  | 

### Return type

[**SmtpConfiguration**](SmtpConfiguration.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

