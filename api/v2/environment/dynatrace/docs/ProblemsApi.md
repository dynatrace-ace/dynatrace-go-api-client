# \ProblemsApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseProblem**](ProblemsApi.md#CloseProblem) | **Post** /problems/{problemId}/close | Closes the specified problem and adds a closing comment on it | maturity&#x3D;EARLY_ADOPTER
[**CreateComment**](ProblemsApi.md#CreateComment) | **Post** /problems/{problemId}/comments | Adds a new comment on the specified problem | maturity&#x3D;EARLY_ADOPTER
[**DeleteComment**](ProblemsApi.md#DeleteComment) | **Delete** /problems/{problemId}/comments/{commentId} | Deletes the specified comment from a problem | maturity&#x3D;EARLY_ADOPTER
[**GetComment**](ProblemsApi.md#GetComment) | **Get** /problems/{problemId}/comments/{commentId} | Gets the specified comment on a problem | maturity&#x3D;EARLY_ADOPTER
[**GetComments**](ProblemsApi.md#GetComments) | **Get** /problems/{problemId}/comments | Gets all comments on the specified problem | maturity&#x3D;EARLY_ADOPTER
[**GetProblem**](ProblemsApi.md#GetProblem) | **Get** /problems/{problemId} | Gets the properties of the specified problem | maturity&#x3D;EARLY_ADOPTER
[**GetProblems**](ProblemsApi.md#GetProblems) | **Get** /problems | Lists problems observed within the specified timeframe | maturity&#x3D;EARLY_ADOPTER
[**UpdateComment**](ProblemsApi.md#UpdateComment) | **Put** /problems/{problemId}/comments/{commentId} | Updates the specified comment on a problem | maturity&#x3D;EARLY_ADOPTER



## CloseProblem

> ProblemCloseResult CloseProblem(ctx, problemId).ClosingComment(closingComment).Execute()

Closes the specified problem and adds a closing comment on it | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    closingComment := *openapiclient.NewClosingComment() // ClosingComment | The JSON body of the request. Contains the closing comment on the problem. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.CloseProblem(context.Background(), problemId).ClosingComment(closingComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.CloseProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseProblem`: ProblemCloseResult
    fmt.Fprintf(os.Stdout, "Response from `ProblemsApi.CloseProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **closingComment** | [**ClosingComment**](ClosingComment.md) | The JSON body of the request. Contains the closing comment on the problem. | 

### Return type

[**ProblemCloseResult**](ProblemCloseResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment

> CreateComment(ctx, problemId).CommentRequestDtoImpl(commentRequestDtoImpl).Execute()

Adds a new comment on the specified problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    commentRequestDtoImpl := *openapiclient.NewCommentRequestDtoImpl("Message_example") // CommentRequestDtoImpl | The JSON body of the request. Contains the comment to be added. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.CreateComment(context.Background(), problemId).CommentRequestDtoImpl(commentRequestDtoImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.CreateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentRequestDtoImpl** | [**CommentRequestDtoImpl**](CommentRequestDtoImpl.md) | The JSON body of the request. Contains the comment to be added. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment(ctx, problemId, commentId).Execute()

Deletes the specified comment from a problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    commentId := "commentId_example" // string | The ID of the required comment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.DeleteComment(context.Background(), problemId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.DeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 
**commentId** | **string** | The ID of the required comment. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> Comment GetComment(ctx, problemId, commentId).Execute()

Gets the specified comment on a problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    commentId := "commentId_example" // string | The ID of the required comment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.GetComment(context.Background(), problemId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.GetComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComment`: Comment
    fmt.Fprintf(os.Stdout, "Response from `ProblemsApi.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 
**commentId** | **string** | The ID of the required comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Comment**](Comment.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments

> CommentsList GetComments(ctx, problemId).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Gets all comments on the specified problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters except the optional **fields** parameter.   (optional)
    pageSize := int64(789) // int64 | The amount of comments in a single response payload.   The maximal allowed page size is 500.   If not set, 10 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.GetComments(context.Background(), problemId).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.GetComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComments`: CommentsList
    fmt.Fprintf(os.Stdout, "Response from `ProblemsApi.GetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters except the optional **fields** parameter.   | 
 **pageSize** | **int64** | The amount of comments in a single response payload.   The maximal allowed page size is 500.   If not set, 10 is used. | 

### Return type

[**CommentsList**](CommentsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProblem

> Problem GetProblem(ctx, problemId).Fields(fields).Execute()

Gets the properties of the specified problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    fields := "fields_example" // string | A list of additional problem properties you can add to the response.   The following properties are available (all other properties are always included and you can't remove them from the response):   * `evidenceDetails`: The details of the root cause. * `impactAnalysis`: The impact analysis of the problem on other entities/users. * `recentComments`: A list of the most recent comments to the problem.  To add properties, specify them as a comma-separated list (for example, `evidenceDetails,impactAnalysis`). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.GetProblem(context.Background(), problemId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.GetProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProblem`: Problem
    fmt.Fprintf(os.Stdout, "Response from `ProblemsApi.GetProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of additional problem properties you can add to the response.   The following properties are available (all other properties are always included and you can&#39;t remove them from the response):   * &#x60;evidenceDetails&#x60;: The details of the root cause. * &#x60;impactAnalysis&#x60;: The impact analysis of the problem on other entities/users. * &#x60;recentComments&#x60;: A list of the most recent comments to the problem.  To add properties, specify them as a comma-separated list (for example, &#x60;evidenceDetails,impactAnalysis&#x60;). | 

### Return type

[**Problem**](Problem.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProblems

> Problems GetProblems(ctx).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).ProblemSelector(problemSelector).EntitySelector(entitySelector).Sort(sort).Execute()

Lists problems observed within the specified timeframe | maturity=EARLY_ADOPTER

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
    fields := "fields_example" // string | A list of additional problem properties you can add to the response.   The following properties are available (all other properties are always included and you can't remove them from the response):   * `evidenceDetails`: The details of the root cause. * `impactAnalysis`: The impact analysis of the problem on other entities/users. * `recentComments`: A list of the most recent comments to the problem.  To add properties, specify them as a comma-separated list (for example, `evidenceDetails,impactAnalysis`).   The field is valid only for the current page of results. You must set it for each page you're requesting. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters except the optional **fields** parameter.   (optional)
    pageSize := int64(789) // int64 | The amount of problems in a single response payload.   The maximal allowed page size is 500.   If not set, 50 is used. (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two hours is used (`now-2h`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    problemSelector := "problemSelector_example" // string | Defines the scope of the query. Only problems matching the specified criteria are included into response.   You can add one or several of the criteria listed below. For each criterion you can specify multiple comma-separated values, unless stated otherwise. If several values are specified, the **OR** logic applies. All values must be quoted.  * Status: `status(\"open\")` or `status(\"closed\")`. You can specify only one status.  * Severity level: `severityLevel(\"level-1\",\"level-2\")`. Find the possible values in the description of the **severityLevel** field of the response. * Impact level: `impactLevel(\"level-11\",\"level-2\")` Find the possible values in the description of the **impactLevel** field of the response. * Root cause entity: `rootCauseEntity(\"id-1\", \"id-2\")`.  * Management zone ID: `managementZoneIds(\"mZId-1\", \"mzId-2\")`.  * Management zone name: `managementZones(\"value-1\",\"value-2\")`.  * Impacted entities: `impactedEntities(\"id-1\", \"id-2\")`.  * Affected entities: `affectedEntities(\"id-1\", \"id-2\")`.  * Type of affected entities: `affectedEntityTypes(\"value-1\",\"value-2\")`.  * Problem ID: `problemId(\"id-1\", \"id-2\")`.  * Alerting profile ID: `problemFilterIds(\"id-1\", \"id-2\")`.  * Alerting profile name: `problemFilterNames(\"value-1\",\"value-2\")`.  * Entity tags: `entityTags(\"[context]key:value\",\"key:value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. If a value-only tag has a colon (`:`) in it, you must escape the colon with a backslash(`\\`). Otherwise, the tag will be parsed as a `key:value tag`. All tag values are case-sensitive. * Display ID of the problem: `displayId(\"id-1\", \"id-2\")`.  * Under maintenance: `underMaintenance(true|false)`. Shows (true) or hides (false) all problems created during maintenance mode. * Text search: `text(\"value\")`. Text search on the following fields: problem title, event titles, displayId and the id of affected and impacted entities. The text search is case insensitive, partial matching and based on a relevance score. Therefore the `relevance` sort option should be used to get the most relevant problems first. The special characters `~` and `\"` need to be escaped using a `~` (e.g. double quote search `text(\"~\"\")`).   To set several criteria, separate them with a comma (`,`). Only results matching **all** criteria are included in the response.    (optional)
    entitySelector := "entitySelector_example" // string | The entity scope of the query. You must set one of these criteria:   * Entity type: `type(\"TYPE\")`  * Dynatrace entity ID: `entityId(\"id\")`. You can specify several IDs, separated by a comma (`entityId(\"id-1\",\"id-2\")`).   You can add one or several of the following criteria. Values are case-sensitive and the `EQUALS` operator is used unless otherwise specified.   * Tag: `tag(\"value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. Any colons (`:`) that are part of the key or value must be escaped with a backslash(`\\`), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: `mzId(123)`  * Management zone name: `mzName(\"value\")` * Entity name: `entityName(\"value\")`. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * `entityName.equals`: changes the operator to `EQUALS`.   * `entityName.startsWith`: changes the operator to `BEGINS WITH`.   * `entityName.in`: enables you to provide multiple values. The `EQUALS` operator applies.   * `caseSensitive(entityName(\"value\"))`: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): `healthState(\"HEALTHY\")` * First seen timestamp: `firstSeenTms.<operator>(now-3h)`. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * `lte`: earlier than or at the specified time  * `lt`: earlier than the specified time  * `gte`: later than or at the specified time  * `gt`: later than the specified time * Entity attribute: `<attribute>(\"value1\",\"value2\")` and `<attribute>.exists()`. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: `fromRelationships.<relationshipName>()` and `toRelationships.<relationshipName>()`. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: `not(<criterion>)`. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (`,`). For example, `type(\"HOST\"),healthState(\"HEALTHY\")`. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   The maximum number of entities that may be selected is limited to 10000.   (optional)
    sort := "sort_example" // string | Specifies a set of comma-separated (`,`) fields for sorting in the problem list.  You can sort by the following properties with a sign prefix for the sorting order.    * `status`: problem status (`+` open problems first or `-` closed problems first)  * `startTime`: problem start time (`+` old problems first or `-` new problems first)   * `relevance`: problem relevance (`+` least relevant problems first or `-` most relevant problems first) - can be used only in combination with text search   If no prefix is set, `+` is used.   You can specify several levels of sorting. For example, `+status,-startTime` sorts problems by status, open problems first. Within the status, problems are sorted by start time, oldest first. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.GetProblems(context.Background()).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).ProblemSelector(problemSelector).EntitySelector(entitySelector).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.GetProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProblems`: Problems
    fmt.Fprintf(os.Stdout, "Response from `ProblemsApi.GetProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | A list of additional problem properties you can add to the response.   The following properties are available (all other properties are always included and you can&#39;t remove them from the response):   * &#x60;evidenceDetails&#x60;: The details of the root cause. * &#x60;impactAnalysis&#x60;: The impact analysis of the problem on other entities/users. * &#x60;recentComments&#x60;: A list of the most recent comments to the problem.  To add properties, specify them as a comma-separated list (for example, &#x60;evidenceDetails,impactAnalysis&#x60;).   The field is valid only for the current page of results. You must set it for each page you&#39;re requesting. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters except the optional **fields** parameter.   | 
 **pageSize** | **int64** | The amount of problems in a single response payload.   The maximal allowed page size is 500.   If not set, 50 is used. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two hours is used (&#x60;now-2h&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **problemSelector** | **string** | Defines the scope of the query. Only problems matching the specified criteria are included into response.   You can add one or several of the criteria listed below. For each criterion you can specify multiple comma-separated values, unless stated otherwise. If several values are specified, the **OR** logic applies. All values must be quoted.  * Status: &#x60;status(\&quot;open\&quot;)&#x60; or &#x60;status(\&quot;closed\&quot;)&#x60;. You can specify only one status.  * Severity level: &#x60;severityLevel(\&quot;level-1\&quot;,\&quot;level-2\&quot;)&#x60;. Find the possible values in the description of the **severityLevel** field of the response. * Impact level: &#x60;impactLevel(\&quot;level-11\&quot;,\&quot;level-2\&quot;)&#x60; Find the possible values in the description of the **impactLevel** field of the response. * Root cause entity: &#x60;rootCauseEntity(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Management zone ID: &#x60;managementZoneIds(\&quot;mZId-1\&quot;, \&quot;mzId-2\&quot;)&#x60;.  * Management zone name: &#x60;managementZones(\&quot;value-1\&quot;,\&quot;value-2\&quot;)&#x60;.  * Impacted entities: &#x60;impactedEntities(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Affected entities: &#x60;affectedEntities(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Type of affected entities: &#x60;affectedEntityTypes(\&quot;value-1\&quot;,\&quot;value-2\&quot;)&#x60;.  * Problem ID: &#x60;problemId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Alerting profile ID: &#x60;problemFilterIds(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Alerting profile name: &#x60;problemFilterNames(\&quot;value-1\&quot;,\&quot;value-2\&quot;)&#x60;.  * Entity tags: &#x60;entityTags(\&quot;[context]key:value\&quot;,\&quot;key:value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. If a value-only tag has a colon (&#x60;:&#x60;) in it, you must escape the colon with a backslash(&#x60;\\&#x60;). Otherwise, the tag will be parsed as a &#x60;key:value tag&#x60;. All tag values are case-sensitive. * Display ID of the problem: &#x60;displayId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * Under maintenance: &#x60;underMaintenance(true|false)&#x60;. Shows (true) or hides (false) all problems created during maintenance mode. * Text search: &#x60;text(\&quot;value\&quot;)&#x60;. Text search on the following fields: problem title, event titles, displayId and the id of affected and impacted entities. The text search is case insensitive, partial matching and based on a relevance score. Therefore the &#x60;relevance&#x60; sort option should be used to get the most relevant problems first. The special characters &#x60;~&#x60; and &#x60;\&quot;&#x60; need to be escaped using a &#x60;~&#x60; (e.g. double quote search &#x60;text(\&quot;~\&quot;\&quot;)&#x60;).   To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching **all** criteria are included in the response.    | 
 **entitySelector** | **string** | The entity scope of the query. You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;).   You can add one or several of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;), otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name: &#x60;entityName(\&quot;value\&quot;)&#x60;. By default this filters for entities, whose name **contains** the given value and is **not** case sensitive.   The following modifications are available:  * &#x60;entityName.equals&#x60;: changes the operator to &#x60;EQUALS&#x60;.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an arguments and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **properties** field.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. The criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET all entity types](https://dt-url.net/dw03s7h) and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in response.   The length of the string is limited to 10,000 characters.   The maximum number of entities that may be selected is limited to 10000.   | 
 **sort** | **string** | Specifies a set of comma-separated (&#x60;,&#x60;) fields for sorting in the problem list.  You can sort by the following properties with a sign prefix for the sorting order.    * &#x60;status&#x60;: problem status (&#x60;+&#x60; open problems first or &#x60;-&#x60; closed problems first)  * &#x60;startTime&#x60;: problem start time (&#x60;+&#x60; old problems first or &#x60;-&#x60; new problems first)   * &#x60;relevance&#x60;: problem relevance (&#x60;+&#x60; least relevant problems first or &#x60;-&#x60; most relevant problems first) - can be used only in combination with text search   If no prefix is set, &#x60;+&#x60; is used.   You can specify several levels of sorting. For example, &#x60;+status,-startTime&#x60; sorts problems by status, open problems first. Within the status, problems are sorted by start time, oldest first. | 

### Return type

[**Problems**](Problems.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment

> UpdateComment(ctx, problemId, commentId).CommentRequestDtoImpl(commentRequestDtoImpl).Execute()

Updates the specified comment on a problem | maturity=EARLY_ADOPTER

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
    problemId := "problemId_example" // string | The ID of the required problem.
    commentId := "commentId_example" // string | The ID of the required comment.
    commentRequestDtoImpl := *openapiclient.NewCommentRequestDtoImpl("Message_example") // CommentRequestDtoImpl | The JSON body of the request. Contains the updated comment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProblemsApi.UpdateComment(context.Background(), problemId, commentId).CommentRequestDtoImpl(commentRequestDtoImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProblemsApi.UpdateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**problemId** | **string** | The ID of the required problem. | 
**commentId** | **string** | The ID of the required comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commentRequestDtoImpl** | [**CommentRequestDtoImpl**](CommentRequestDtoImpl.md) | The JSON body of the request. Contains the updated comment. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

