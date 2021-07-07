# \ActiveGatesAutoUpdateJobsApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpdateJobForAg**](ActiveGatesAutoUpdateJobsApi.md#CreateUpdateJobForAg) | **Post** /activeGates/{agId}/updateJobs | Creates a new update job for the specified ActiveGate
[**DeleteUpdateJobByJobIdForAg**](ActiveGatesAutoUpdateJobsApi.md#DeleteUpdateJobByJobIdForAg) | **Delete** /activeGates/{agId}/updateJobs/{jobId} | Deletes the specified update job
[**GetAllUpdateJobList**](ActiveGatesAutoUpdateJobsApi.md#GetAllUpdateJobList) | **Get** /activeGates/updateJobs | List ActiveGates with update jobs
[**GetUpdateJobByJobIdForAg**](ActiveGatesAutoUpdateJobsApi.md#GetUpdateJobByJobIdForAg) | **Get** /activeGates/{agId}/updateJobs/{jobId} | Gets the parameters of the specified update job
[**GetUpdateJobListByAgId**](ActiveGatesAutoUpdateJobsApi.md#GetUpdateJobListByAgId) | **Get** /activeGates/{agId}/updateJobs | Lists update jobs for the specified ActiveGate
[**ValidateUpdateJobForAg**](ActiveGatesAutoUpdateJobsApi.md#ValidateUpdateJobForAg) | **Post** /activeGates/{agId}/updateJobs/validator | Validates the payload for the &#x60;POST /activeGates/{agId}/updateJobs&#x60; request.



## CreateUpdateJobForAg

> UpdateJob CreateUpdateJobForAg(ctx, agId).UpdateJob(updateJob).Execute()

Creates a new update job for the specified ActiveGate

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
    updateJob := *openapiclient.NewUpdateJob("1.190.0.20200301-130000") // UpdateJob | JSON body of the request, containing update-job parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.CreateUpdateJobForAg(context.Background(), agId).UpdateJob(updateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.CreateUpdateJobForAg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateJobForAg`: UpdateJob
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateJobsApi.CreateUpdateJobForAg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateJobForAgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateJob** | [**UpdateJob**](UpdateJob.md) | JSON body of the request, containing update-job parameters. | 

### Return type

[**UpdateJob**](UpdateJob.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpdateJobByJobIdForAg

> DeleteUpdateJobByJobIdForAg(ctx, agId, jobId).Execute()

Deletes the specified update job

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
    jobId := "jobId_example" // string | A unique identifier for a update-job of ActiveGate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.DeleteUpdateJobByJobIdForAg(context.Background(), agId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.DeleteUpdateJobByJobIdForAg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 
**jobId** | **string** | A unique identifier for a update-job of ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpdateJobByJobIdForAgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUpdateJobList

> UpdateJobsAll GetAllUpdateJobList(ctx).From(from).To(to).StartVersionCompareType(startVersionCompareType).StartVersion(startVersion).UpdateType(updateType).TargetVersionCompareType(targetVersionCompareType).TargetVersion(targetVersion).LastUpdates(lastUpdates).Execute()

List ActiveGates with update jobs



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
    from := "from_example" // string | The start of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of one day is used (`now-1d`).   Maximum timeframe is 31 days. (optional)
    to := "to_example" // string | The end of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    startVersionCompareType := "startVersionCompareType_example" // string | Filters the resulting set of update jobs by the specified initial version.    Specify the comparison operator here. (optional) (default to "EQUAL")
    startVersion := "startVersion_example" // string | Filters the resulting set of update-jobs by the initial version (required format `<major>.<minor>.<revision>`). (optional)
    updateType := "ACTIVE_GATE" // string | Filters the resulting set of update-jobs by the update type. (optional)
    targetVersionCompareType := "targetVersionCompareType_example" // string | Filters the resulting set of update jobs by the specified target version.    Specify the comparison operator here. (optional) (default to "EQUAL")
    targetVersion := "targetVersion_example" // string | Filters the resulting set of update-jobs by the target version (required format `<major>.<minor>.<revision>`). (optional)
    lastUpdates := true // bool | If `true`, filters the resulting set of update jobs to the most recent update of each type. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.GetAllUpdateJobList(context.Background()).From(from).To(to).StartVersionCompareType(startVersionCompareType).StartVersion(startVersion).UpdateType(updateType).TargetVersionCompareType(targetVersionCompareType).TargetVersion(targetVersion).LastUpdates(lastUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.GetAllUpdateJobList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllUpdateJobList`: UpdateJobsAll
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateJobsApi.GetAllUpdateJobList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUpdateJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of one day is used (&#x60;now-1d&#x60;).   Maximum timeframe is 31 days. | 
 **to** | **string** | The end of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **startVersionCompareType** | **string** | Filters the resulting set of update jobs by the specified initial version.    Specify the comparison operator here. | [default to &quot;EQUAL&quot;]
 **startVersion** | **string** | Filters the resulting set of update-jobs by the initial version (required format &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60;). | 
 **updateType** | **string** | Filters the resulting set of update-jobs by the update type. | 
 **targetVersionCompareType** | **string** | Filters the resulting set of update jobs by the specified target version.    Specify the comparison operator here. | [default to &quot;EQUAL&quot;]
 **targetVersion** | **string** | Filters the resulting set of update-jobs by the target version (required format &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60;). | 
 **lastUpdates** | **bool** | If &#x60;true&#x60;, filters the resulting set of update jobs to the most recent update of each type. | [default to false]

### Return type

[**UpdateJobsAll**](UpdateJobsAll.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateJobByJobIdForAg

> UpdateJob GetUpdateJobByJobIdForAg(ctx, agId, jobId).Execute()

Gets the parameters of the specified update job

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
    jobId := "jobId_example" // string | A unique identifier for a update-job of ActiveGate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.GetUpdateJobByJobIdForAg(context.Background(), agId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.GetUpdateJobByJobIdForAg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpdateJobByJobIdForAg`: UpdateJob
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateJobsApi.GetUpdateJobByJobIdForAg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 
**jobId** | **string** | A unique identifier for a update-job of ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateJobByJobIdForAgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateJob**](UpdateJob.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateJobListByAgId

> UpdateJobList GetUpdateJobListByAgId(ctx, agId).From(from).To(to).StartVersionCompareType(startVersionCompareType).StartVersion(startVersion).UpdateType(updateType).TargetVersionCompareType(targetVersionCompareType).TargetVersion(targetVersion).LastUpdates(lastUpdates).Execute()

Lists update jobs for the specified ActiveGate



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
    from := "from_example" // string | The start of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of one week is used (`now-1w`).   Maximum timeframe is 31 days. (optional)
    to := "to_example" // string | The end of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    startVersionCompareType := "startVersionCompareType_example" // string | Filters the resulting set of update jobs by the specified initial version.    Specify the comparison operator here. (optional) (default to "EQUAL")
    startVersion := "startVersion_example" // string | Filters the resulting set of update-jobs by the initial version (required format `<major>.<minor>.<revision>`). (optional)
    updateType := "ACTIVE_GATE" // string | Filters the resulting set of update-jobs by the update type. (optional)
    targetVersionCompareType := "targetVersionCompareType_example" // string | Filters the resulting set of update jobs by the specified target version.    Specify the comparison operator here. (optional) (default to "EQUAL")
    targetVersion := "targetVersion_example" // string | Filters the resulting set of update-jobs by the target version (required format `<major>.<minor>.<revision>`). (optional)
    lastUpdates := true // bool | If `true`, filters the resulting set of update jobs to the most recent update of each type. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.GetUpdateJobListByAgId(context.Background(), agId).From(from).To(to).StartVersionCompareType(startVersionCompareType).StartVersion(startVersion).UpdateType(updateType).TargetVersionCompareType(targetVersionCompareType).TargetVersion(targetVersion).LastUpdates(lastUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.GetUpdateJobListByAgId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpdateJobListByAgId`: UpdateJobList
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateJobsApi.GetUpdateJobListByAgId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateJobListByAgIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of one week is used (&#x60;now-1w&#x60;).   Maximum timeframe is 31 days. | 
 **to** | **string** | The end of the requested timeframe for update jobs.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **startVersionCompareType** | **string** | Filters the resulting set of update jobs by the specified initial version.    Specify the comparison operator here. | [default to &quot;EQUAL&quot;]
 **startVersion** | **string** | Filters the resulting set of update-jobs by the initial version (required format &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60;). | 
 **updateType** | **string** | Filters the resulting set of update-jobs by the update type. | 
 **targetVersionCompareType** | **string** | Filters the resulting set of update jobs by the specified target version.    Specify the comparison operator here. | [default to &quot;EQUAL&quot;]
 **targetVersion** | **string** | Filters the resulting set of update-jobs by the target version (required format &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60;). | 
 **lastUpdates** | **bool** | If &#x60;true&#x60;, filters the resulting set of update jobs to the most recent update of each type. | [default to false]

### Return type

[**UpdateJobList**](UpdateJobList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateUpdateJobForAg

> ValidateUpdateJobForAg(ctx, agId).UpdateJob(updateJob).Execute()

Validates the payload for the `POST /activeGates/{agId}/updateJobs` request.

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
    updateJob := *openapiclient.NewUpdateJob("1.190.0.20200301-130000") // UpdateJob | JSON body of the request, containing update-job parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActiveGatesAutoUpdateJobsApi.ValidateUpdateJobForAg(context.Background(), agId).UpdateJob(updateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateJobsApi.ValidateUpdateJobForAg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateJobForAgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateJob** | [**UpdateJob**](UpdateJob.md) | JSON body of the request, containing update-job parameters. | 

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

