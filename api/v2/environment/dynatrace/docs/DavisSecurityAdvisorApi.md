# \DavisSecurityAdvisorApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdviceForSecurityProblems**](DavisSecurityAdvisorApi.md#GetAdviceForSecurityProblems) | **Get** /davis/securityAdvices | Provides advice for security problems. | maturity&#x3D;EARLY_ADOPTER



## GetAdviceForSecurityProblems

> DavisSecurityAdviceList GetAdviceForSecurityProblems(ctx).ManagementZoneFilter(managementZoneFilter).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Provides advice for security problems. | maturity=EARLY_ADOPTER



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
    managementZoneFilter := "managementZoneFilter_example" // string | To specify management zones, use one of the options listed below. For each option you can specify multiple comma-separated values. If several values are specified, the **OR** logic applies. All values are case-sensitive and must be quoted.   * Management zone ID: ids(\"mzId-1\", \"mzId-2\").  * Management zone names: names(\"mz-1\", \"mz-2\").   You can specify several comma-separated criteria (for example, `names(\"myMz\"),ids(\"9130632296508575249\")`). (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of security advices in a single response payload.   The maximal allowed page size is 50.   If not set, 5 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DavisSecurityAdvisorApi.GetAdviceForSecurityProblems(context.Background()).ManagementZoneFilter(managementZoneFilter).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DavisSecurityAdvisorApi.GetAdviceForSecurityProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdviceForSecurityProblems`: DavisSecurityAdviceList
    fmt.Fprintf(os.Stdout, "Response from `DavisSecurityAdvisorApi.GetAdviceForSecurityProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdviceForSecurityProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZoneFilter** | **string** | To specify management zones, use one of the options listed below. For each option you can specify multiple comma-separated values. If several values are specified, the **OR** logic applies. All values are case-sensitive and must be quoted.   * Management zone ID: ids(\&quot;mzId-1\&quot;, \&quot;mzId-2\&quot;).  * Management zone names: names(\&quot;mz-1\&quot;, \&quot;mz-2\&quot;).   You can specify several comma-separated criteria (for example, &#x60;names(\&quot;myMz\&quot;),ids(\&quot;9130632296508575249\&quot;)&#x60;). | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of security advices in a single response payload.   The maximal allowed page size is 50.   If not set, 5 is used. | 

### Return type

[**DavisSecurityAdviceList**](DavisSecurityAdviceList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

