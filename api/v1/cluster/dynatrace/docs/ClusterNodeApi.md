# \ClusterNodeApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeadNodeCleaning**](ClusterNodeApi.md#DeadNodeCleaning) | **Delete** /nodeManagement/deadNodeCleaning | 
[**FinalizeJoining**](ClusterNodeApi.md#FinalizeJoining) | **Put** /nodeManagement/finalizeJoining | 
[**ForceCassandraQuorumOverride**](ClusterNodeApi.md#ForceCassandraQuorumOverride) | **Put** /nodeManagement/cassandraQuorumOverride/{quorum} | Override cassandra read quorum at runtime only for this particular server instance
[**GetCassandraQuorumOverride**](ClusterNodeApi.md#GetCassandraQuorumOverride) | **Get** /nodeManagement/cassandraQuorumOverride | Get the cassandra read quorum override.
[**GetNodeJoin**](ClusterNodeApi.md#GetNodeJoin) | **Get** /nodeManagement/join/check/{requestId} | Verification if a node can be added to the cluster status
[**GetNodeJoining**](ClusterNodeApi.md#GetNodeJoining) | **Get** /nodeManagement/joining/{requestId} | 
[**GetNodeProductVersion**](ClusterNodeApi.md#GetNodeProductVersion) | **Get** /nodeManagement/productVersion | Get product version
[**GetNodeRemovalStatus**](ClusterNodeApi.md#GetNodeRemovalStatus) | **Get** /nodeManagement/nodeRemovalStatus | Get node removal status
[**GetNodeServerStatus**](ClusterNodeApi.md#GetNodeServerStatus) | **Get** /nodeManagement/nodeServerStatus | Get node&#39;s server status
[**GetNodekeeperProductVersion**](ClusterNodeApi.md#GetNodekeeperProductVersion) | **Get** /nodeManagement/nodekeeperProductVersion | 
[**InitializeNodeJoinCheck**](ClusterNodeApi.md#InitializeNodeJoinCheck) | **Post** /nodeManagement/join/check | Initialize verification if a node can be added to the cluster
[**InitializeNodeJoining**](ClusterNodeApi.md#InitializeNodeJoining) | **Post** /nodeManagement/joining | 
[**NodeRestart**](ClusterNodeApi.md#NodeRestart) | **Post** /nodeManagement/nodeRestart | Restart node
[**ResponsibilityOverride**](ClusterNodeApi.md#ResponsibilityOverride) | **Put** /nodeManagement/responsibilityOverride | Override which nodes can perform internal responsibilities
[**SetInstallerStatus**](ClusterNodeApi.md#SetInstallerStatus) | **Post** /nodeManagement/installerStatus | 
[**TriggerRemoveNode**](ClusterNodeApi.md#TriggerRemoveNode) | **Post** /nodeManagement/triggerRemoveNode | Remove node
[**TryLockJoining**](ClusterNodeApi.md#TryLockJoining) | **Put** /nodeManagement/tryLockJoining | 
[**TryLockRemoval**](ClusterNodeApi.md#TryLockRemoval) | **Put** /nodeManagement/tryLockRemoval | 
[**UnlockRemoval**](ClusterNodeApi.md#UnlockRemoval) | **Put** /nodeManagement/unlockRemoval | 



## DeadNodeCleaning

> DeadNodeCleaning(ctx).Ip(ip).RequestedOnNode(requestedOnNode).RequestedByUser(requestedByUser).Execute()



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
    ip := "ip_example" // string |  (optional)
    requestedOnNode := int32(56) // int32 |  (optional)
    requestedByUser := "requestedByUser_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.DeadNodeCleaning(context.Background()).Ip(ip).RequestedOnNode(requestedOnNode).RequestedByUser(requestedByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.DeadNodeCleaning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeadNodeCleaningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **requestedOnNode** | **int32** |  | 
 **requestedByUser** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeJoining

> FinalizeJoining(ctx).Ip(ip).Status(status).ForceUnlock(forceUnlock).FinalizeNodeJoiningRequestDto(finalizeNodeJoiningRequestDto).Execute()



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
    ip := "ip_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    forceUnlock := true // bool |  (optional)
    finalizeNodeJoiningRequestDto := *openapiclient.NewFinalizeNodeJoiningRequestDto() // FinalizeNodeJoiningRequestDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.FinalizeJoining(context.Background()).Ip(ip).Status(status).ForceUnlock(forceUnlock).FinalizeNodeJoiningRequestDto(finalizeNodeJoiningRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.FinalizeJoining``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeJoiningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **status** | **string** |  | 
 **forceUnlock** | **bool** |  | 
 **finalizeNodeJoiningRequestDto** | [**FinalizeNodeJoiningRequestDto**](FinalizeNodeJoiningRequestDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceCassandraQuorumOverride

> ForceCassandraQuorumOverride(ctx, quorum).Execute()

Override cassandra read quorum at runtime only for this particular server instance



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
    quorum := "quorum_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.ForceCassandraQuorumOverride(context.Background(), quorum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.ForceCassandraQuorumOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quorum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceCassandraQuorumOverrideRequest struct via the builder pattern


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


## GetCassandraQuorumOverride

> GetCassandraQuorumOverride(ctx).Execute()

Get the cassandra read quorum override.



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
    resp, r, err := api_client.ClusterNodeApi.GetCassandraQuorumOverride(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetCassandraQuorumOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCassandraQuorumOverrideRequest struct via the builder pattern


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


## GetNodeJoin

> NodeJoiningPreCheckStatusDto GetNodeJoin(ctx, requestId).Execute()

Verification if a node can be added to the cluster status

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
    resp, r, err := api_client.ClusterNodeApi.GetNodeJoin(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodeJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeJoin`: NodeJoiningPreCheckStatusDto
    fmt.Fprintf(os.Stdout, "Response from `ClusterNodeApi.GetNodeJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeJoiningPreCheckStatusDto**](NodeJoiningPreCheckStatusDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeJoining

> GetNodeJoining(ctx, requestId).Execute()



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
    resp, r, err := api_client.ClusterNodeApi.GetNodeJoining(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodeJoining``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetNodeJoiningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeProductVersion

> ProductVersion GetNodeProductVersion(ctx).Execute()

Get product version

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
    resp, r, err := api_client.ClusterNodeApi.GetNodeProductVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodeProductVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeProductVersion`: ProductVersion
    fmt.Fprintf(os.Stdout, "Response from `ClusterNodeApi.GetNodeProductVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeProductVersionRequest struct via the builder pattern


### Return type

[**ProductVersion**](ProductVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeRemovalStatus

> GetNodeRemovalStatus(ctx).Execute()

Get node removal status



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
    resp, r, err := api_client.ClusterNodeApi.GetNodeRemovalStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodeRemovalStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRemovalStatusRequest struct via the builder pattern


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


## GetNodeServerStatus

> string GetNodeServerStatus(ctx).Execute()

Get node's server status

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
    resp, r, err := api_client.ClusterNodeApi.GetNodeServerStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodeServerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeServerStatus`: string
    fmt.Fprintf(os.Stdout, "Response from `ClusterNodeApi.GetNodeServerStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeServerStatusRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodekeeperProductVersion

> GetNodekeeperProductVersion(ctx).Execute()



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
    resp, r, err := api_client.ClusterNodeApi.GetNodekeeperProductVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.GetNodekeeperProductVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodekeeperProductVersionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeNodeJoinCheck

> NodeJoiningInitializePreCheckResponseDto InitializeNodeJoinCheck(ctx).Ip(ip).Timestamp(timestamp).Datacenter(datacenter).Execute()

Initialize verification if a node can be added to the cluster

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
    ip := "ip_example" // string |  (optional)
    timestamp := int64(789) // int64 |  (optional)
    datacenter := "datacenter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.InitializeNodeJoinCheck(context.Background()).Ip(ip).Timestamp(timestamp).Datacenter(datacenter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.InitializeNodeJoinCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeNodeJoinCheck`: NodeJoiningInitializePreCheckResponseDto
    fmt.Fprintf(os.Stdout, "Response from `ClusterNodeApi.InitializeNodeJoinCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitializeNodeJoinCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **timestamp** | **int64** |  | 
 **datacenter** | **string** |  | 

### Return type

[**NodeJoiningInitializePreCheckResponseDto**](NodeJoiningInitializePreCheckResponseDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeNodeJoining

> InitializeNodeJoining(ctx).Ip(ip).Timestamp(timestamp).Datacenter(datacenter).Execute()



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
    ip := "ip_example" // string |  (optional)
    timestamp := int64(789) // int64 |  (optional)
    datacenter := "datacenter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.InitializeNodeJoining(context.Background()).Ip(ip).Timestamp(timestamp).Datacenter(datacenter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.InitializeNodeJoining``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitializeNodeJoiningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **timestamp** | **int64** |  | 
 **datacenter** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeRestart

> NodeRestart(ctx).Execute()

Restart node

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
    resp, r, err := api_client.ClusterNodeApi.NodeRestart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.NodeRestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNodeRestartRequest struct via the builder pattern


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


## ResponsibilityOverride

> ResponsibilityOverride(ctx).ResponsibilityOverrideDto(responsibilityOverrideDto).Execute()

Override which nodes can perform internal responsibilities



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
    responsibilityOverrideDto := *openapiclient.NewResponsibilityOverrideDto() // ResponsibilityOverrideDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.ResponsibilityOverride(context.Background()).ResponsibilityOverrideDto(responsibilityOverrideDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.ResponsibilityOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResponsibilityOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responsibilityOverrideDto** | [**ResponsibilityOverrideDto**](ResponsibilityOverrideDto.md) |  | 

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


## SetInstallerStatus

> SetInstallerStatus(ctx).Ip(ip).InstallerStatusDto(installerStatusDto).Execute()



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
    ip := "ip_example" // string |  (optional)
    installerStatusDto := *openapiclient.NewInstallerStatusDto() // InstallerStatusDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.SetInstallerStatus(context.Background()).Ip(ip).InstallerStatusDto(installerStatusDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.SetInstallerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetInstallerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **installerStatusDto** | [**InstallerStatusDto**](InstallerStatusDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerRemoveNode

> TriggerRemoveNode(ctx).RemoveNode(removeNode).Execute()

Remove node

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
    removeNode := *openapiclient.NewRemoveNode() // RemoveNode |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.TriggerRemoveNode(context.Background()).RemoveNode(removeNode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.TriggerRemoveNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerRemoveNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeNode** | [**RemoveNode**](RemoveNode.md) |  | 

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


## TryLockJoining

> TryLockJoining(ctx).Ip(ip).Execute()



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
    ip := "ip_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.TryLockJoining(context.Background()).Ip(ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.TryLockJoining``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTryLockJoiningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TryLockRemoval

> TryLockRemoval(ctx).Id(id).Ip(ip).RequestedOnNode(requestedOnNode).RequestedByUser(requestedByUser).Execute()



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
    id := int32(56) // int32 |  (optional)
    ip := "ip_example" // string |  (optional)
    requestedOnNode := int32(56) // int32 |  (optional)
    requestedByUser := "requestedByUser_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.TryLockRemoval(context.Background()).Id(id).Ip(ip).RequestedOnNode(requestedOnNode).RequestedByUser(requestedByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.TryLockRemoval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTryLockRemovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **ip** | **string** |  | 
 **requestedOnNode** | **int32** |  | 
 **requestedByUser** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockRemoval

> UnlockRemoval(ctx).Id(id).Ip(ip).RemovalStopReason(removalStopReason).Execute()



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
    id := int32(56) // int32 |  (optional)
    ip := "ip_example" // string |  (optional)
    removalStopReason := "removalStopReason_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterNodeApi.UnlockRemoval(context.Background()).Id(id).Ip(ip).RemovalStopReason(removalStopReason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNodeApi.UnlockRemoval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlockRemovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **ip** | **string** |  | 
 **removalStopReason** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

