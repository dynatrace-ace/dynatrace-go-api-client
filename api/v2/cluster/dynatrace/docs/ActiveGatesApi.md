# \ActiveGatesApi

All URIs are relative to *http://https:/api/cluster/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllActiveGates**](ActiveGatesApi.md#GetAllActiveGates) | **Get** /activeGates | Lists all available ActiveGates
[**GetOneActiveGateById**](ActiveGatesApi.md#GetOneActiveGateById) | **Get** /activeGates/{agId} | Gets the details of the specified ActiveGate



## GetAllActiveGates

> ActiveGateList GetAllActiveGates(ctx).Hostname(hostname).OsType(osType).NetworkAddress(networkAddress).LoadBalancerAddress(loadBalancerAddress).Type_(type_).NetworkZone(networkZone).UpdateStatus(updateStatus).VersionCompareType(versionCompareType).Version(version).AutoUpdate(autoUpdate).Group(group).Online(online).EnabledModule(enabledModule).DisabledModule(disabledModule).Containerized(containerized).Execute()

Lists all available ActiveGates



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
    hostname := "hostname_example" // string | Filters the resulting set of ActiveGates by the name of the host it's running on.    You can specify a partial name. In that case, the `CONTAINS` operator is used. (optional)
    osType := "osType_example" // string | Filters the resulting set of ActiveGates by the OS type of the host it's running on. (optional)
    networkAddress := "networkAddress_example" // string | Filters the resulting set of ActiveGates by the network address.    You can specify a partial address. In that case, the `CONTAINS` operator is used. (optional)
    loadBalancerAddress := "loadBalancerAddress_example" // string | Filters the resulting set of ActiveGates by the Load Balancer address.    You can specify a partial address. In that case, the `CONTAINS` operator is used. (optional)
    type_ := "type__example" // string | Filters the resulting set of ActiveGates by the ActiveGate type. (optional)
    networkZone := "networkZone_example" // string | Filters the resulting set of ActiveGates by the network zone.    You can specify a partial name. In that case, the `CONTAINS` operator is used. (optional)
    updateStatus := "updateStatus_example" // string | Filters the resulting set of ActiveGates by the auto-update status. (optional)
    versionCompareType := "versionCompareType_example" // string | Filters the resulting set of ActiveGates by the specified version.    Specify the comparison operator here. (optional) (default to "EQUAL")
    version := "version_example" // string | Filters the resulting set of ActiveGates by the specified version.    Specify the version in `<major>.<minor>.<revision>` format (for example, `1.195.0`) here. (optional)
    autoUpdate := "autoUpdate_example" // string | Filters the resulting set of ActiveGates by the actual state of auto-update. (optional)
    group := "group_example" // string | Filters the resulting set of ActiveGates by the group.    You can specify a partial name. In that case, the `CONTAINS` operator is used. (optional)
    online := true // bool | Filters the resulting set of ActiveGates by the communication status. (optional)
    enabledModule := []string{"EnabledModule_example"} // []string | Filters the resulting set of ActiveGates by the enabled modules. (optional)
    disabledModule := []string{"DisabledModule_example"} // []string | Filters the resulting set of ActiveGates by the disabled modules. (optional)
    containerized := true // bool | Filters the resulting set of ActiveGates to those which are running in container (`true`) or not (`false`). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesApi.GetAllActiveGates(context.Background()).Hostname(hostname).OsType(osType).NetworkAddress(networkAddress).LoadBalancerAddress(loadBalancerAddress).Type_(type_).NetworkZone(networkZone).UpdateStatus(updateStatus).VersionCompareType(versionCompareType).Version(version).AutoUpdate(autoUpdate).Group(group).Online(online).EnabledModule(enabledModule).DisabledModule(disabledModule).Containerized(containerized).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesApi.GetAllActiveGates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllActiveGates`: ActiveGateList
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesApi.GetAllActiveGates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllActiveGatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostname** | **string** | Filters the resulting set of ActiveGates by the name of the host it&#39;s running on.    You can specify a partial name. In that case, the &#x60;CONTAINS&#x60; operator is used. | 
 **osType** | **string** | Filters the resulting set of ActiveGates by the OS type of the host it&#39;s running on. | 
 **networkAddress** | **string** | Filters the resulting set of ActiveGates by the network address.    You can specify a partial address. In that case, the &#x60;CONTAINS&#x60; operator is used. | 
 **loadBalancerAddress** | **string** | Filters the resulting set of ActiveGates by the Load Balancer address.    You can specify a partial address. In that case, the &#x60;CONTAINS&#x60; operator is used. | 
 **type_** | **string** | Filters the resulting set of ActiveGates by the ActiveGate type. | 
 **networkZone** | **string** | Filters the resulting set of ActiveGates by the network zone.    You can specify a partial name. In that case, the &#x60;CONTAINS&#x60; operator is used. | 
 **updateStatus** | **string** | Filters the resulting set of ActiveGates by the auto-update status. | 
 **versionCompareType** | **string** | Filters the resulting set of ActiveGates by the specified version.    Specify the comparison operator here. | [default to &quot;EQUAL&quot;]
 **version** | **string** | Filters the resulting set of ActiveGates by the specified version.    Specify the version in &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format (for example, &#x60;1.195.0&#x60;) here. | 
 **autoUpdate** | **string** | Filters the resulting set of ActiveGates by the actual state of auto-update. | 
 **group** | **string** | Filters the resulting set of ActiveGates by the group.    You can specify a partial name. In that case, the &#x60;CONTAINS&#x60; operator is used. | 
 **online** | **bool** | Filters the resulting set of ActiveGates by the communication status. | 
 **enabledModule** | **[]string** | Filters the resulting set of ActiveGates by the enabled modules. | 
 **disabledModule** | **[]string** | Filters the resulting set of ActiveGates by the disabled modules. | 
 **containerized** | **bool** | Filters the resulting set of ActiveGates to those which are running in container (&#x60;true&#x60;) or not (&#x60;false&#x60;). | 

### Return type

[**ActiveGateList**](ActiveGateList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneActiveGateById

> ActiveGate GetOneActiveGateById(ctx, agId).Execute()

Gets the details of the specified ActiveGate

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
    agId := "agId_example" // string | The ID of the required ActiveGate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesApi.GetOneActiveGateById(context.Background(), agId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesApi.GetOneActiveGateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOneActiveGateById`: ActiveGate
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesApi.GetOneActiveGateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneActiveGateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveGate**](ActiveGate.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

