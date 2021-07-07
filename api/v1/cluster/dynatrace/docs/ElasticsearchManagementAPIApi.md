# \ElasticsearchManagementAPIApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckElasticsearchUpgradePossibility**](ElasticsearchManagementAPIApi.md#CheckElasticsearchUpgradePossibility) | **Get** /elastic/upgradeStatus | Checks if it is safe to upgrade elasticsearch 
[**CreateAutoFollowPattern**](ElasticsearchManagementAPIApi.md#CreateAutoFollowPattern) | **Post** /elastic/ccr/autoFollowPattern | Creates auto follow pattern for MultiDC environment
[**DeleteAutoFollowPattern**](ElasticsearchManagementAPIApi.md#DeleteAutoFollowPattern) | **Delete** /elastic/ccr/autoFollowPattern | Deletes auto follow pattern for MultiDC environment



## CheckElasticsearchUpgradePossibility

> ElasticsearchUpgradeDTO CheckElasticsearchUpgradePossibility(ctx).ExpectedElasticsearchNodes(expectedElasticsearchNodes).Execute()

Checks if it is safe to upgrade elasticsearch 

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
    expectedElasticsearchNodes := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElasticsearchManagementAPIApi.CheckElasticsearchUpgradePossibility(context.Background()).ExpectedElasticsearchNodes(expectedElasticsearchNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchManagementAPIApi.CheckElasticsearchUpgradePossibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckElasticsearchUpgradePossibility`: ElasticsearchUpgradeDTO
    fmt.Fprintf(os.Stdout, "Response from `ElasticsearchManagementAPIApi.CheckElasticsearchUpgradePossibility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckElasticsearchUpgradePossibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expectedElasticsearchNodes** | **int32** |  | [default to -1]

### Return type

[**ElasticsearchUpgradeDTO**](ElasticsearchUpgradeDTO.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoFollowPattern

> ElasticsearchOperationDto CreateAutoFollowPattern(ctx).Execute()

Creates auto follow pattern for MultiDC environment

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
    resp, r, err := api_client.ElasticsearchManagementAPIApi.CreateAutoFollowPattern(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchManagementAPIApi.CreateAutoFollowPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutoFollowPattern`: ElasticsearchOperationDto
    fmt.Fprintf(os.Stdout, "Response from `ElasticsearchManagementAPIApi.CreateAutoFollowPattern`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoFollowPatternRequest struct via the builder pattern


### Return type

[**ElasticsearchOperationDto**](ElasticsearchOperationDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoFollowPattern

> ElasticsearchOperationDto DeleteAutoFollowPattern(ctx).Execute()

Deletes auto follow pattern for MultiDC environment

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
    resp, r, err := api_client.ElasticsearchManagementAPIApi.DeleteAutoFollowPattern(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElasticsearchManagementAPIApi.DeleteAutoFollowPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAutoFollowPattern`: ElasticsearchOperationDto
    fmt.Fprintf(os.Stdout, "Response from `ElasticsearchManagementAPIApi.DeleteAutoFollowPattern`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoFollowPatternRequest struct via the builder pattern


### Return type

[**ElasticsearchOperationDto**](ElasticsearchOperationDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

