# \MigrationToMultiDatacenterDeploymentApi

All URIs are relative to *http://https:/api/v1.0/onpremise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMultiDCTopologyDraft**](MigrationToMultiDatacenterDeploymentApi.md#CreateMultiDCTopologyDraft) | **Post** /multiDc/migration/datacenterTopology | Create cluster topology draft | maturity&#x3D;EARLY_ADOPTER
[**DeleteMultiDCTopologyDraft**](MigrationToMultiDatacenterDeploymentApi.md#DeleteMultiDCTopologyDraft) | **Delete** /multiDc/migration/datacenterTopology | Delete cluster topology draft | maturity&#x3D;EARLY_ADOPTER
[**FinishMigration**](MigrationToMultiDatacenterDeploymentApi.md#FinishMigration) | **Post** /multiDc/migration/finish | Finish migration | maturity&#x3D;EARLY_ADOPTER
[**GetClusterState**](MigrationToMultiDatacenterDeploymentApi.md#GetClusterState) | **Get** /multiDc/migration/clusterState | Get cluster state of migration from single to multi datacenter cluster and its sub steps | maturity&#x3D;EARLY_ADOPTER
[**GetMigrateCassandraNewDatacenter**](MigrationToMultiDatacenterDeploymentApi.md#GetMigrateCassandraNewDatacenter) | **Get** /multiDc/migration/cassandra/newDc/{requestId} | Get cassandra migration status in new datacenter | maturity&#x3D;EARLY_ADOPTER
[**GetMigrateCassandraOldDatacenter**](MigrationToMultiDatacenterDeploymentApi.md#GetMigrateCassandraOldDatacenter) | **Get** /multiDc/migration/cassandra/currentDc/{requestId} | Get cassandra migration status in current datacenter | maturity&#x3D;EARLY_ADOPTER
[**GetMigrateElasticsearchStatus**](MigrationToMultiDatacenterDeploymentApi.md#GetMigrateElasticsearchStatus) | **Get** /multiDc/migration/elasticsearch/{requestId} | Get elasticsearch migration status | maturity&#x3D;EARLY_ADOPTER
[**GetMigrateServerStatus**](MigrationToMultiDatacenterDeploymentApi.md#GetMigrateServerStatus) | **Get** /multiDc/migration/server/{requestId} | Get server migration status | maturity&#x3D;EARLY_ADOPTER
[**GetMigrationStateForAllComponents**](MigrationToMultiDatacenterDeploymentApi.md#GetMigrationStateForAllComponents) | **Get** /multiDc/migration/inServerconfigState | Get state of in-server config migration (from single to multi datacenter cluster) | maturity&#x3D;EARLY_ADOPTER
[**GetMultiDCTopologyDraft**](MigrationToMultiDatacenterDeploymentApi.md#GetMultiDCTopologyDraft) | **Get** /multiDc/migration/datacenterTopology | Get cluster topology draft | maturity&#x3D;EARLY_ADOPTER
[**GetMultiDcNodekeeperHealthcheck**](MigrationToMultiDatacenterDeploymentApi.md#GetMultiDcNodekeeperHealthcheck) | **Get** /multiDc/migration/nodekeeper/healthCheck | Get nodekeeper helathcheck | maturity&#x3D;EARLY_ADOPTER
[**GetPrepareClusterForReplicationProgress**](MigrationToMultiDatacenterDeploymentApi.md#GetPrepareClusterForReplicationProgress) | **Get** /multiDc/migration/clusterReplicationPreparation | Get progress status of cluster preparation for a replication | maturity&#x3D;EARLY_ADOPTER
[**GetSubStepState**](MigrationToMultiDatacenterDeploymentApi.md#GetSubStepState) | **Get** /multiDc/migration/clusterState/{subStep} | Get sub step of migration from single to multi datacenter cluster | maturity&#x3D;EARLY_ADOPTER
[**GetTopologyConfiguration**](MigrationToMultiDatacenterDeploymentApi.md#GetTopologyConfiguration) | **Get** /multiDc/migration/clusterNodes/currentDc/{requestId} | Get topology configuration status | maturity&#x3D;EARLY_ADOPTER
[**GetVerifyCassandraStatus**](MigrationToMultiDatacenterDeploymentApi.md#GetVerifyCassandraStatus) | **Get** /multiDc/migration/cassandra/rebuildStatus | Verify cassandra rebuild status | maturity&#x3D;EARLY_ADOPTER
[**GetVerifyElasticsearchMigration**](MigrationToMultiDatacenterDeploymentApi.md#GetVerifyElasticsearchMigration) | **Get** /multiDc/migration/elasticsearch/indexMigrationStatus | Verify elasticsearch migration status | maturity&#x3D;EARLY_ADOPTER
[**InitDatacenterCleanUp**](MigrationToMultiDatacenterDeploymentApi.md#InitDatacenterCleanUp) | **Post** /multiDc/migration/lostDatacenterCleanUp | Clean up lost datacenter settings | maturity&#x3D;EARLY_ADOPTER
[**InitMigrateCassandraNewDatacenter**](MigrationToMultiDatacenterDeploymentApi.md#InitMigrateCassandraNewDatacenter) | **Post** /multiDc/migration/cassandra/newDc | Start cassandra migration in new datacenter | maturity&#x3D;EARLY_ADOPTER
[**InitMigrateCassandraOldDatacenter**](MigrationToMultiDatacenterDeploymentApi.md#InitMigrateCassandraOldDatacenter) | **Post** /multiDc/migration/cassandra/currentDc | Start cassandra migration in current datacenter | maturity&#x3D;EARLY_ADOPTER
[**InitPrepareClusterForReplication**](MigrationToMultiDatacenterDeploymentApi.md#InitPrepareClusterForReplication) | **Post** /multiDc/migration/clusterReplicationPreparation | Start procedure of cluster preparation for a replication | maturity&#x3D;EARLY_ADOPTER
[**InitTopologyConfiguration**](MigrationToMultiDatacenterDeploymentApi.md#InitTopologyConfiguration) | **Post** /multiDc/migration/clusterNodes/currentDc | Start cluster topology configuration - open firewall ports | maturity&#x3D;EARLY_ADOPTER
[**MigrateElasticsearch**](MigrationToMultiDatacenterDeploymentApi.md#MigrateElasticsearch) | **Post** /multiDc/migration/elasticsearch | Start elasticsearch migration | maturity&#x3D;EARLY_ADOPTER
[**MigrateServer**](MigrationToMultiDatacenterDeploymentApi.md#MigrateServer) | **Post** /multiDc/migration/server | Start server migration in current datacenter | maturity&#x3D;EARLY_ADOPTER
[**ModifyClusterState**](MigrationToMultiDatacenterDeploymentApi.md#ModifyClusterState) | **Put** /multiDc/migration/clusterState | Modify the overall cluster state of migration from single to multi datacenter cluster | maturity&#x3D;EARLY_ADOPTER
[**ModifySubStepState**](MigrationToMultiDatacenterDeploymentApi.md#ModifySubStepState) | **Put** /multiDc/migration/clusterState/{subStep} | Modify sub step of migration from single to multi datacenter cluster | maturity&#x3D;EARLY_ADOPTER



## CreateMultiDCTopologyDraft

> CreateMultiDCTopologyDraft(ctx).DatacenterMigrationDto(datacenterMigrationDto).Execute()

Create cluster topology draft | maturity=EARLY_ADOPTER

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
    datacenterMigrationDto := *openapiclient.NewDatacenterMigrationDto() // DatacenterMigrationDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.CreateMultiDCTopologyDraft(context.Background()).DatacenterMigrationDto(datacenterMigrationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.CreateMultiDCTopologyDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultiDCTopologyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datacenterMigrationDto** | [**DatacenterMigrationDto**](DatacenterMigrationDto.md) |  | 

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


## DeleteMultiDCTopologyDraft

> DeleteMultiDCTopologyDraft(ctx).Execute()

Delete cluster topology draft | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.DeleteMultiDCTopologyDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.DeleteMultiDCTopologyDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultiDCTopologyDraftRequest struct via the builder pattern


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


## FinishMigration

> FinishMigration(ctx).Execute()

Finish migration | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.FinishMigration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.FinishMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinishMigrationRequest struct via the builder pattern


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


## GetClusterState

> SingleToMultiDcMigrationClusterState GetClusterState(ctx).Execute()

Get cluster state of migration from single to multi datacenter cluster and its sub steps | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetClusterState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetClusterState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterState`: SingleToMultiDcMigrationClusterState
    fmt.Fprintf(os.Stdout, "Response from `MigrationToMultiDatacenterDeploymentApi.GetClusterState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStateRequest struct via the builder pattern


### Return type

[**SingleToMultiDcMigrationClusterState**](SingleToMultiDcMigrationClusterState.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMigrateCassandraNewDatacenter

> GetMigrateCassandraNewDatacenter(ctx, requestId).Execute()

Get cassandra migration status in new datacenter | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMigrateCassandraNewDatacenter(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMigrateCassandraNewDatacenter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMigrateCassandraNewDatacenterRequest struct via the builder pattern


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


## GetMigrateCassandraOldDatacenter

> GetMigrateCassandraOldDatacenter(ctx, requestId).Execute()

Get cassandra migration status in current datacenter | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMigrateCassandraOldDatacenter(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMigrateCassandraOldDatacenter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMigrateCassandraOldDatacenterRequest struct via the builder pattern


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


## GetMigrateElasticsearchStatus

> GetMigrateElasticsearchStatus(ctx, requestId).Execute()

Get elasticsearch migration status | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMigrateElasticsearchStatus(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMigrateElasticsearchStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMigrateElasticsearchStatusRequest struct via the builder pattern


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


## GetMigrateServerStatus

> GetMigrateServerStatus(ctx, requestId).Execute()

Get server migration status | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMigrateServerStatus(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMigrateServerStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMigrateServerStatusRequest struct via the builder pattern


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


## GetMigrationStateForAllComponents

> InServerConfigDatacenterMigrationState GetMigrationStateForAllComponents(ctx).Execute()

Get state of in-server config migration (from single to multi datacenter cluster) | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMigrationStateForAllComponents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMigrationStateForAllComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMigrationStateForAllComponents`: InServerConfigDatacenterMigrationState
    fmt.Fprintf(os.Stdout, "Response from `MigrationToMultiDatacenterDeploymentApi.GetMigrationStateForAllComponents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigrationStateForAllComponentsRequest struct via the builder pattern


### Return type

[**InServerConfigDatacenterMigrationState**](InServerConfigDatacenterMigrationState.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiDCTopologyDraft

> MultiDatacenterTopology GetMultiDCTopologyDraft(ctx).Execute()

Get cluster topology draft | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMultiDCTopologyDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMultiDCTopologyDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultiDCTopologyDraft`: MultiDatacenterTopology
    fmt.Fprintf(os.Stdout, "Response from `MigrationToMultiDatacenterDeploymentApi.GetMultiDCTopologyDraft`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiDCTopologyDraftRequest struct via the builder pattern


### Return type

[**MultiDatacenterTopology**](MultiDatacenterTopology.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultiDcNodekeeperHealthcheck

> GetMultiDcNodekeeperHealthcheck(ctx).Execute()

Get nodekeeper helathcheck | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetMultiDcNodekeeperHealthcheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetMultiDcNodekeeperHealthcheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiDcNodekeeperHealthcheckRequest struct via the builder pattern


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


## GetPrepareClusterForReplicationProgress

> ClusterPreparationForReplicationDto GetPrepareClusterForReplicationProgress(ctx).Execute()

Get progress status of cluster preparation for a replication | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetPrepareClusterForReplicationProgress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetPrepareClusterForReplicationProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrepareClusterForReplicationProgress`: ClusterPreparationForReplicationDto
    fmt.Fprintf(os.Stdout, "Response from `MigrationToMultiDatacenterDeploymentApi.GetPrepareClusterForReplicationProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrepareClusterForReplicationProgressRequest struct via the builder pattern


### Return type

[**ClusterPreparationForReplicationDto**](ClusterPreparationForReplicationDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubStepState

> MigrationState GetSubStepState(ctx, subStep).Execute()

Get sub step of migration from single to multi datacenter cluster | maturity=EARLY_ADOPTER

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
    subStep := "subStep_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetSubStepState(context.Background(), subStep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetSubStepState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubStepState`: MigrationState
    fmt.Fprintf(os.Stdout, "Response from `MigrationToMultiDatacenterDeploymentApi.GetSubStepState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subStep** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubStepStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MigrationState**](MigrationState.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyConfiguration

> GetTopologyConfiguration(ctx, requestId).Execute()

Get topology configuration status | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetTopologyConfiguration(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetTopologyConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetTopologyConfigurationRequest struct via the builder pattern


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


## GetVerifyCassandraStatus

> GetVerifyCassandraStatus(ctx).Execute()

Verify cassandra rebuild status | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetVerifyCassandraStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetVerifyCassandraStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifyCassandraStatusRequest struct via the builder pattern


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


## GetVerifyElasticsearchMigration

> GetVerifyElasticsearchMigration(ctx).Execute()

Verify elasticsearch migration status | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.GetVerifyElasticsearchMigration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.GetVerifyElasticsearchMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifyElasticsearchMigrationRequest struct via the builder pattern


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


## InitDatacenterCleanUp

> InitDatacenterCleanUp(ctx).Execute()

Clean up lost datacenter settings | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.InitDatacenterCleanUp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.InitDatacenterCleanUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitDatacenterCleanUpRequest struct via the builder pattern


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


## InitMigrateCassandraNewDatacenter

> InitMigrateCassandraNewDatacenter(ctx).Execute()

Start cassandra migration in new datacenter | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.InitMigrateCassandraNewDatacenter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.InitMigrateCassandraNewDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitMigrateCassandraNewDatacenterRequest struct via the builder pattern


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


## InitMigrateCassandraOldDatacenter

> InitMigrateCassandraOldDatacenter(ctx).Execute()

Start cassandra migration in current datacenter | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.InitMigrateCassandraOldDatacenter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.InitMigrateCassandraOldDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitMigrateCassandraOldDatacenterRequest struct via the builder pattern


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


## InitPrepareClusterForReplication

> InitPrepareClusterForReplication(ctx).Execute()

Start procedure of cluster preparation for a replication | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.InitPrepareClusterForReplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.InitPrepareClusterForReplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitPrepareClusterForReplicationRequest struct via the builder pattern


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


## InitTopologyConfiguration

> InitTopologyConfiguration(ctx).Execute()

Start cluster topology configuration - open firewall ports | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.InitTopologyConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.InitTopologyConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitTopologyConfigurationRequest struct via the builder pattern


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


## MigrateElasticsearch

> MigrateElasticsearch(ctx).Execute()

Start elasticsearch migration | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.MigrateElasticsearch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.MigrateElasticsearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateElasticsearchRequest struct via the builder pattern


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


## MigrateServer

> MigrateServer(ctx).Execute()

Start server migration in current datacenter | maturity=EARLY_ADOPTER

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
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.MigrateServer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.MigrateServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateServerRequest struct via the builder pattern


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


## ModifyClusterState

> ModifyClusterState(ctx).Status(status).Details(details).Execute()

Modify the overall cluster state of migration from single to multi datacenter cluster | maturity=EARLY_ADOPTER

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
    status := "status_example" // string |  (optional)
    details := "details_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.ModifyClusterState(context.Background()).Status(status).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.ModifyClusterState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyClusterStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **details** | **string** |  | [default to &quot;&quot;]

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


## ModifySubStepState

> ModifySubStepState(ctx, subStep).Status(status).Details(details).Execute()

Modify sub step of migration from single to multi datacenter cluster | maturity=EARLY_ADOPTER

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
    subStep := "subStep_example" // string | 
    status := "status_example" // string |  (optional)
    details := "details_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MigrationToMultiDatacenterDeploymentApi.ModifySubStepState(context.Background(), subStep).Status(status).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationToMultiDatacenterDeploymentApi.ModifySubStepState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subStep** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySubStepStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **details** | **string** |  | 

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

