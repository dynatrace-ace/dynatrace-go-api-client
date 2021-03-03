# \ProblemApi

All URIs are relative to *http://https:/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseProblem**](ProblemApi.md#CloseProblem) | **Post** /problem/details/{problemId}/close | Closes the specified problem and adds a closing comment to it
[**DeleteComment**](ProblemApi.md#DeleteComment) | **Delete** /problem/details/{problemId}/comments/{commentId} | Deletes an existing comment to the specified problem.
[**GetComment**](ProblemApi.md#GetComment) | **Get** /problem/details/{problemId}/comments | Gets all the comments to the specified problem
[**GetDetails**](ProblemApi.md#GetDetails) | **Get** /problem/details/{problemId} | Gets the properties of the specified problem
[**GetFeed**](ProblemApi.md#GetFeed) | **Get** /problem/feed | Gets the information about problems within the specified timeframe
[**GetProblemCloseResult**](ProblemApi.md#GetProblemCloseResult) | **Get** /problem/details/{problemId}/close | Closes the specified problem
[**GetProblemStatus**](ProblemApi.md#GetProblemStatus) | **Get** /problem/status | Lists the number of open problems, split by impact level
[**PushComment**](ProblemApi.md#PushComment) | **Post** /problem/details/{problemId}/comments | Adds a new comment to the specified problem
[**UpdateComment**](ProblemApi.md#UpdateComment) | **Put** /problem/details/{problemId}/comments/{commentId} | Updates an existing comment to the specified problem



## CloseProblem

> ProblemCloseResult CloseProblem(ctx, problemId).Content(content).Execute()

Closes the specified problem and adds a closing comment to it

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
    problemId := "problemId_example" // string | The ID of the problem to be closed.
    content := "content_example" // string | The closing comment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.CloseProblem(context.Background(), problemId).Content(content).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.CloseProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseProblem`: ProblemCloseResult
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.CloseProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem to be closed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | The closing comment. | 

### Return type

[**ProblemCloseResult**](ProblemCloseResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment(ctx, problemId, commentId).Execute()

Deletes an existing comment to the specified problem.

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
    problemId := "problemId_example" // string | The ID of the problem where you want to delete the comment.
    commentId := "commentId_example" // string | The ID of the comment to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.DeleteComment(context.Background(), problemId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.DeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem where you want to delete the comment. | 
**commentId** | **string** | The ID of the comment to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


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


## GetComment

> ProblemCommentList GetComment(ctx, problemId).Execute()

Gets all the comments to the specified problem

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
    problemId := "problemId_example" // string | The ID of the problem where you want to read the comments.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.GetComment(context.Background(), problemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComment`: ProblemCommentList
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem where you want to read the comments. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProblemCommentList**](ProblemCommentList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetails

> ProblemDetailsResultWrapper GetDetails(ctx, problemId).Execute()

Gets the properties of the specified problem

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
    problemId := "problemId_example" // string | The ID of the problem you're inquiring.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.GetDetails(context.Background(), problemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetails`: ProblemDetailsResultWrapper
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProblemDetailsResultWrapper**](ProblemDetailsResultWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeed

> ProblemFeedResultWrapper GetFeed(ctx).RelativeTime(relativeTime).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Status(status).ImpactLevel(impactLevel).SeverityLevel(severityLevel).Tag(tag).ExpandDetails(expandDetails).Execute()

Gets the information about problems within the specified timeframe



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
    relativeTime := "relativeTime_example" // string | The relative timeframe of the inquiry, back from the current time. (optional)
    startTimestamp := int64(789) // int64 | The start timestamp of the requested timeframe, in UTC milliseconds. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the requested timeframe, in UTC milliseconds.   If `endTimestamp` is later than the current time, the current time is used.   The timeframe must not exceed 31 days. (optional)
    status := "status_example" // string | Filters the result problems by the status. (optional)
    impactLevel := "impactLevel_example" // string | Filters the result problems by the impact level. (optional)
    severityLevel := "severityLevel_example" // string | Filters the result problems by the severity level. (optional)
    tag := []string{"Inner_example"} // []string | Filters the result problems by the tags of affected entities.You can specify several tags in the following format: `tag=tag1&tag=tag2`. The problem has to match *all* the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: `[context]key:value`. (optional)
    expandDetails := true // bool | Includes(`true`) or excludes(`false`) related events to the response.    Defaults to `false`, excluding the related events. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.GetFeed(context.Background()).RelativeTime(relativeTime).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Status(status).ImpactLevel(impactLevel).SeverityLevel(severityLevel).Tag(tag).ExpandDetails(expandDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeed`: ProblemFeedResultWrapper
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relativeTime** | **string** | The relative timeframe of the inquiry, back from the current time. | 
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in UTC milliseconds. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in UTC milliseconds.   If &#x60;endTimestamp&#x60; is later than the current time, the current time is used.   The timeframe must not exceed 31 days. | 
 **status** | **string** | Filters the result problems by the status. | 
 **impactLevel** | **string** | Filters the result problems by the impact level. | 
 **severityLevel** | **string** | Filters the result problems by the severity level. | 
 **tag** | **[]string** | Filters the result problems by the tags of affected entities.You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The problem has to match *all* the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: &#x60;[context]key:value&#x60;. | 
 **expandDetails** | **bool** | Includes(&#x60;true&#x60;) or excludes(&#x60;false&#x60;) related events to the response.    Defaults to &#x60;false&#x60;, excluding the related events. | 

### Return type

[**ProblemFeedResultWrapper**](ProblemFeedResultWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProblemCloseResult

> ProblemCloseResult GetProblemCloseResult(ctx, problemId).Execute()

Closes the specified problem

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
    problemId := "problemId_example" // string | The ID of the problem to be closed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.GetProblemCloseResult(context.Background(), problemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetProblemCloseResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProblemCloseResult`: ProblemCloseResult
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetProblemCloseResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem to be closed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemCloseResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProblemCloseResult**](ProblemCloseResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProblemStatus

> ProblemStatusResultWrapper GetProblemStatus(ctx).Execute()

Lists the number of open problems, split by impact level

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
    resp, r, err := api_client.ProblemApi.GetProblemStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.GetProblemStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProblemStatus`: ProblemStatusResultWrapper
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.GetProblemStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemStatusRequest struct via the builder pattern


### Return type

[**ProblemStatusResultWrapper**](ProblemStatusResultWrapper.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushComment

> ProblemComment PushComment(ctx, problemId).PushProblemComment(pushProblemComment).Execute()

Adds a new comment to the specified problem

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
    problemId := "problemId_example" // string | The ID of the problem where you want to add the comment.
    pushProblemComment := *openapiclient.NewPushProblemComment("Comment_example", "User_example") // PushProblemComment | JSON body of the request, containing the comment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.PushComment(context.Background(), problemId).PushProblemComment(pushProblemComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.PushComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushComment`: ProblemComment
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.PushComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem where you want to add the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushProblemComment** | [**PushProblemComment**](PushProblemComment.md) | JSON body of the request, containing the comment. | 

### Return type

[**ProblemComment**](ProblemComment.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment

> ProblemComment UpdateComment(ctx, problemId, commentId).PushProblemComment(pushProblemComment).Execute()

Updates an existing comment to the specified problem

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
    problemId := "problemId_example" // string | The ID of the problem where you want to edit the comment.
    commentId := "commentId_example" // string | The ID of the comment you want to edit.
    pushProblemComment := *openapiclient.NewPushProblemComment("Comment_example", "User_example") // PushProblemComment | JSON body of the request, containing the updated comment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemApi.UpdateComment(context.Background(), problemId, commentId).PushProblemComment(pushProblemComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemApi.UpdateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComment`: ProblemComment
    fmt.Fprintf(os.Stdout, "Response from `ProblemApi.UpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the problem where you want to edit the comment. | 
**commentId** | **string** | The ID of the comment you want to edit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pushProblemComment** | [**PushProblemComment**](PushProblemComment.md) | JSON body of the request, containing the updated comment. | 

### Return type

[**ProblemComment**](ProblemComment.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

