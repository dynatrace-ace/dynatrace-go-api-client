# \SecurityProblemsApi

All URIs are relative to *http://https:/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityProblem**](SecurityProblemsApi.md#GetSecurityProblem) | **Get** /securityProblems/{id} | Gets the specified security problem | maturity&#x3D;EARLY_ADOPTER
[**GetSecurityProblems**](SecurityProblemsApi.md#GetSecurityProblems) | **Get** /securityProblems | Lists all security problems | maturity&#x3D;EARLY_ADOPTER
[**MuteSecurityProblem**](SecurityProblemsApi.md#MuteSecurityProblem) | **Post** /securityProblems/{id}/mute | Mutes the specified security problem. | maturity&#x3D;EARLY_ADOPTER
[**UnmuteSecurityProblem**](SecurityProblemsApi.md#UnmuteSecurityProblem) | **Post** /securityProblems/{id}/unmute | Un-mutes the specified security problem. | maturity&#x3D;EARLY_ADOPTER



## GetSecurityProblem

> SecurityProblemDetails GetSecurityProblem(ctx, id).Fields(fields).Execute()

Gets the specified security problem | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required security problem.
    fields := "fields_example" // string | Defines the list of problem properties to be added to the response.  `securityProblemId`, `displayId`, `status`, `muted`, `externalVulnerabilityId`, `vulnerabilityType`, `title`, `packageName`, `url`, `technology`, `firstSeenTimestamp`, `lastUpdateTimestamp`, `cveIds` are **always** included in the result. To add more properties, list them with a leading plus `+`. Available fields are listed below. You can specify several properties, separated by a comma (for example `+riskAssessment,+managementZones`).   * `riskAssessment`: A risk assessment of the security problem. * `managementZones`: The management zone where the security problem occurred. * `description`: The description of the vulnerability. * `events`: A list of events of the security problem. * `vulnerableComponents`: A list of vulnerable components of the security problem.  * `affectedEntities`: A list of affected entities of the security problem.  * `exposedEntities`: A list of exposed entities of the security problem.  * `reachableDataAssets`: A list of data assets reachable by affected entities of the security problem.  * `relatedEntities`: A list of related entities of the security problem.  * `relatedContainerImages`: A list of related container images of the security problem.   (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityProblemsApi.GetSecurityProblem(context.Background(), id).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsApi.GetSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityProblem`: SecurityProblemDetails
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsApi.GetSecurityProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Defines the list of problem properties to be added to the response.  &#x60;securityProblemId&#x60;, &#x60;displayId&#x60;, &#x60;status&#x60;, &#x60;muted&#x60;, &#x60;externalVulnerabilityId&#x60;, &#x60;vulnerabilityType&#x60;, &#x60;title&#x60;, &#x60;packageName&#x60;, &#x60;url&#x60;, &#x60;technology&#x60;, &#x60;firstSeenTimestamp&#x60;, &#x60;lastUpdateTimestamp&#x60;, &#x60;cveIds&#x60; are **always** included in the result. To add more properties, list them with a leading plus &#x60;+&#x60;. Available fields are listed below. You can specify several properties, separated by a comma (for example &#x60;+riskAssessment,+managementZones&#x60;).   * &#x60;riskAssessment&#x60;: A risk assessment of the security problem. * &#x60;managementZones&#x60;: The management zone where the security problem occurred. * &#x60;description&#x60;: The description of the vulnerability. * &#x60;events&#x60;: A list of events of the security problem. * &#x60;vulnerableComponents&#x60;: A list of vulnerable components of the security problem.  * &#x60;affectedEntities&#x60;: A list of affected entities of the security problem.  * &#x60;exposedEntities&#x60;: A list of exposed entities of the security problem.  * &#x60;reachableDataAssets&#x60;: A list of data assets reachable by affected entities of the security problem.  * &#x60;relatedEntities&#x60;: A list of related entities of the security problem.  * &#x60;relatedContainerImages&#x60;: A list of related container images of the security problem.   | 

### Return type

[**SecurityProblemDetails**](SecurityProblemDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityProblems

> SecurityProblemList GetSecurityProblems(ctx).NextPageKey(nextPageKey).PageSize(pageSize).SecurityProblemSelector(securityProblemSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()

Lists all security problems | maturity=EARLY_ADOPTER

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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of security problems in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)
    securityProblemSelector := "securityProblemSelector_example" // string | Defines the scope of the query. Only security problems matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the `EQUALS` operator is used unless otherwise specified.  * Status: `status(\"value\")`. Find the possible values in the description of the **status** field of the response. If not set, all security problems are returned. * Muted: `muted(\"value\")`. Possible values are `TRUE` or `FALSE`. * Risk level: `riskLevel(\"value\")`. The Davis Risk Level. Find the possible values in the description of the **riskLevel** field of the response. * Minimum risk score: `minRiskScore(\"5.5\")`. The Davis minimum Risk Score. The `GREATER THAN OR EQUAL TO` operator is used. Specify a number between `1.0` and `10.0`. * Maximum risk score: `maxRiskScore(\"5.5\")`. The Davis maximum Risk Score. The `LESS THAN` operator is used. Specify a number between `1.0` and `10.0`. * Base risk level: `baseRiskLevel(\"value\")`. The Base Risk Level from the CVSS. Find the possible values in the description of the **riskLevel** field of the response. * Minimum base risk score: `minBaseRiskScore(\"5.5\")`. The Base minimum Risk Score from the CVSS. The `GREATER THAN OR EQUAL TO` operator is used. Specify a number between `1.0` and `10.0`. * Maximum base risk score: `maxBaseRiskScore(\"5.5\")`. The Base maximum Risk Score from the CVSS. The `LESS THAN` operator is used. Specify a number between `1.0` and `10.0`. * External vulnerability ID contains: `externalVulnerabilityIdContains(\"id-1\")`. The `CONTAINS` operator is used. * External vulnerability ID: `externalVulnerabilityId(\"id-1\",\"id-2\")`. Case insensitive `EQUALS` operator is used. * CVE ID: `cveId(\"id\")`. * Risk assessment `riskAssessment(\"value-1\",\"value-2\")` Possible values are `EXPOSED`, `SENSITIVE`, and `EXPLOIT`. * Related host ID: `relatedHostIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related host name: `relatedHostNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related host name contains: `relatedHostNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Related Kubernetes cluster ID: `relatedKubernetesClusterIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related Kubernetes cluster name: `relatedKubernetesClusterNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related Kubernetes cluster name contains: `relatedKubernetesClusterNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Related Kubernetes workload ID: `relatedKubernetesWorkloadIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related Kubernetes workload name: `relatedKubernetesWorkloadNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related Kubernetes workload name contains: `relatedKubernetesWorkloadNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Management zone ID: `managementZoneIds(\"mzId-1\",\"mzId-2\")`. * Management zone name: `managementZones(\"name-1\",\"name-2\")`. Values are case-sensitive. * Affected process group ID: `affectedPgIds(\"pgId-1\", \"pgId-2\")`. Specify Dynatrace entity IDs here. * Affected process group name: `affectedPgNames(\"name-1\", \"name-2\")`. Values are case-sensitive. * Affected process group name contains: `affectedPgNameContains(\"name-1\")`. The `CONTAINS` operator is used. * Vulnerable component ID: `vulnerableComponentIds(\"componentId-1\", \"componentId-2\")`. Specify component IDs here. * Vulnerable component name: `vulnerableComponentNames(\"name-1\", \"name-2\")`. Values are case-sensitive. * Vulnerable component name contains: `vulnerableComponentNameContains(\"name-1\")`. The `CONTAINS` operator is used. * Host tags: `hostTags(\"hostTag-1\")`. The `CONTAINS` operator is used. * Process group tags: `pgTags(\"pgTag-1\")`. The `CONTAINS` operator is used. * Process group instance tags: `pgiTags(\"pgiTag-1\")`. The `CONTAINS` operator is used. * Tags: `tags(\"tag-1\")`. The `CONTAINS` operator is used. This selector picks hosts, process groups, and process group instances at the same time. * Display ID: `displayIds(\"S-1234\",\"S-5678\")`. The `EQUALS` operator is used. * Technology: `technology(\"technology-1\",\"technology-2\")`. Find the possible values in the description of the **technology** field of the response. The `EQUALS` operator is used.  Risk score and risk category are mutually exclusive (cannot be used at the same time).  To set several criteria, separate them with a comma (`,`). Only results matching (**all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)
    sort := "sort_example" // string | Specifies a field for sorting the security problem list.  You can sort by the following properties with a sign prefix for the sorting order.   * `status`: The security problem status (`+` open first or `-` resolved first)  * `muted`: The security problem mute state (`+` muted first or `-` unmuted first)  * `technology`: The security problem technology (`+` ascending or `-` descending)  * `firstSeenTimestamp`: The timestamp of the first occurrence of the security problem (`+` new problems first or `-` old problems first)  * `securityProblemId`: The auto-generated ID of the security problem (`+` lower number first or `-` higher number first)  * `externalVulnerabilityId`: The ID of the external vulnerability (`+` lower number first or `-` higher number first)  * `displayId`: The display ID (`+` lower number first or `-` higher number first)  * `riskAssessment.riskScore`: The Davis security score (`+` lower score first or `-` higher score first)  * `riskAssessment.riskLevel`: The Davis security level (`+` lower level first or `-` higher level first)  * `riskAssessment.exposure`: Whether the problem is exposed to the internet (`+` unexposed first or `-` exposed first)  * `riskAssessment.dataAssets`: Whether data assets are affected (`+` not affected first or `-` affected first)   If no prefix is set, `+` is used. (optional)
    fields := "fields_example" // string | Defines the list of problem properties to be added to the response.  `securityProblemId`, `displayId`, `status`, `muted`, `externalVulnerabilityId`, `vulnerabilityType`, `title`, `packageName`, `url`, `technology`, `firstSeenTimestamp`, `lastUpdateTimestamp`, `cveIds` are **always** included in the result. To add more properties, list them with a leading plus `+`. Available fields are listed below. You can specify several properties, separated by a comma (for example `+riskAssessment,+managementZones`).   * `riskAssessment`: A risk assessment of the security problem. * `managementZones`: The management zone where the security problem occurred.  (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of thirty days is used (`now-30d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityProblemsApi.GetSecurityProblems(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).SecurityProblemSelector(securityProblemSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsApi.GetSecurityProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityProblems`: SecurityProblemList
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsApi.GetSecurityProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of security problems in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 
 **securityProblemSelector** | **string** | Defines the scope of the query. Only security problems matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * Status: &#x60;status(\&quot;value\&quot;)&#x60;. Find the possible values in the description of the **status** field of the response. If not set, all security problems are returned. * Muted: &#x60;muted(\&quot;value\&quot;)&#x60;. Possible values are &#x60;TRUE&#x60; or &#x60;FALSE&#x60;. * Risk level: &#x60;riskLevel(\&quot;value\&quot;)&#x60;. The Davis Risk Level. Find the possible values in the description of the **riskLevel** field of the response. * Minimum risk score: &#x60;minRiskScore(\&quot;5.5\&quot;)&#x60;. The Davis minimum Risk Score. The &#x60;GREATER THAN OR EQUAL TO&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Maximum risk score: &#x60;maxRiskScore(\&quot;5.5\&quot;)&#x60;. The Davis maximum Risk Score. The &#x60;LESS THAN&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Base risk level: &#x60;baseRiskLevel(\&quot;value\&quot;)&#x60;. The Base Risk Level from the CVSS. Find the possible values in the description of the **riskLevel** field of the response. * Minimum base risk score: &#x60;minBaseRiskScore(\&quot;5.5\&quot;)&#x60;. The Base minimum Risk Score from the CVSS. The &#x60;GREATER THAN OR EQUAL TO&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Maximum base risk score: &#x60;maxBaseRiskScore(\&quot;5.5\&quot;)&#x60;. The Base maximum Risk Score from the CVSS. The &#x60;LESS THAN&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * External vulnerability ID contains: &#x60;externalVulnerabilityIdContains(\&quot;id-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * External vulnerability ID: &#x60;externalVulnerabilityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;. Case insensitive &#x60;EQUALS&#x60; operator is used. * CVE ID: &#x60;cveId(\&quot;id\&quot;)&#x60;. * Risk assessment &#x60;riskAssessment(\&quot;value-1\&quot;,\&quot;value-2\&quot;)&#x60; Possible values are &#x60;EXPOSED&#x60;, &#x60;SENSITIVE&#x60;, and &#x60;EXPLOIT&#x60;. * Related host ID: &#x60;relatedHostIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related host name: &#x60;relatedHostNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related host name contains: &#x60;relatedHostNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Related Kubernetes cluster ID: &#x60;relatedKubernetesClusterIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related Kubernetes cluster name: &#x60;relatedKubernetesClusterNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related Kubernetes cluster name contains: &#x60;relatedKubernetesClusterNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Related Kubernetes workload ID: &#x60;relatedKubernetesWorkloadIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related Kubernetes workload name: &#x60;relatedKubernetesWorkloadNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related Kubernetes workload name contains: &#x60;relatedKubernetesWorkloadNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Management zone ID: &#x60;managementZoneIds(\&quot;mzId-1\&quot;,\&quot;mzId-2\&quot;)&#x60;. * Management zone name: &#x60;managementZones(\&quot;name-1\&quot;,\&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Affected process group ID: &#x60;affectedPgIds(\&quot;pgId-1\&quot;, \&quot;pgId-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Affected process group name: &#x60;affectedPgNames(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Affected process group name contains: &#x60;affectedPgNameContains(\&quot;name-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Vulnerable component ID: &#x60;vulnerableComponentIds(\&quot;componentId-1\&quot;, \&quot;componentId-2\&quot;)&#x60;. Specify component IDs here. * Vulnerable component name: &#x60;vulnerableComponentNames(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Vulnerable component name contains: &#x60;vulnerableComponentNameContains(\&quot;name-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Host tags: &#x60;hostTags(\&quot;hostTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Process group tags: &#x60;pgTags(\&quot;pgTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Process group instance tags: &#x60;pgiTags(\&quot;pgiTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Tags: &#x60;tags(\&quot;tag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. This selector picks hosts, process groups, and process group instances at the same time. * Display ID: &#x60;displayIds(\&quot;S-1234\&quot;,\&quot;S-5678\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator is used. * Technology: &#x60;technology(\&quot;technology-1\&quot;,\&quot;technology-2\&quot;)&#x60;. Find the possible values in the description of the **technology** field of the response. The &#x60;EQUALS&#x60; operator is used.  Risk score and risk category are mutually exclusive (cannot be used at the same time).  To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching (**all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 
 **sort** | **string** | Specifies a field for sorting the security problem list.  You can sort by the following properties with a sign prefix for the sorting order.   * &#x60;status&#x60;: The security problem status (&#x60;+&#x60; open first or &#x60;-&#x60; resolved first)  * &#x60;muted&#x60;: The security problem mute state (&#x60;+&#x60; muted first or &#x60;-&#x60; unmuted first)  * &#x60;technology&#x60;: The security problem technology (&#x60;+&#x60; ascending or &#x60;-&#x60; descending)  * &#x60;firstSeenTimestamp&#x60;: The timestamp of the first occurrence of the security problem (&#x60;+&#x60; new problems first or &#x60;-&#x60; old problems first)  * &#x60;securityProblemId&#x60;: The auto-generated ID of the security problem (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;externalVulnerabilityId&#x60;: The ID of the external vulnerability (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;displayId&#x60;: The display ID (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;riskAssessment.riskScore&#x60;: The Davis security score (&#x60;+&#x60; lower score first or &#x60;-&#x60; higher score first)  * &#x60;riskAssessment.riskLevel&#x60;: The Davis security level (&#x60;+&#x60; lower level first or &#x60;-&#x60; higher level first)  * &#x60;riskAssessment.exposure&#x60;: Whether the problem is exposed to the internet (&#x60;+&#x60; unexposed first or &#x60;-&#x60; exposed first)  * &#x60;riskAssessment.dataAssets&#x60;: Whether data assets are affected (&#x60;+&#x60; not affected first or &#x60;-&#x60; affected first)   If no prefix is set, &#x60;+&#x60; is used. | 
 **fields** | **string** | Defines the list of problem properties to be added to the response.  &#x60;securityProblemId&#x60;, &#x60;displayId&#x60;, &#x60;status&#x60;, &#x60;muted&#x60;, &#x60;externalVulnerabilityId&#x60;, &#x60;vulnerabilityType&#x60;, &#x60;title&#x60;, &#x60;packageName&#x60;, &#x60;url&#x60;, &#x60;technology&#x60;, &#x60;firstSeenTimestamp&#x60;, &#x60;lastUpdateTimestamp&#x60;, &#x60;cveIds&#x60; are **always** included in the result. To add more properties, list them with a leading plus &#x60;+&#x60;. Available fields are listed below. You can specify several properties, separated by a comma (for example &#x60;+riskAssessment,+managementZones&#x60;).   * &#x60;riskAssessment&#x60;: A risk assessment of the security problem. * &#x60;managementZones&#x60;: The management zone where the security problem occurred.  | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of thirty days is used (&#x60;now-30d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 

### Return type

[**SecurityProblemList**](SecurityProblemList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteSecurityProblem

> MuteSecurityProblem(ctx, id).Mute(mute).Execute()

Mutes the specified security problem. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required security problem.
    mute := *openapiclient.NewMute("Reason_example", "Comment_example") // Mute | The JSON body of the request. Contains the muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityProblemsApi.MuteSecurityProblem(context.Background(), id).Mute(mute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsApi.MuteSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mute** | [**Mute**](Mute.md) | The JSON body of the request. Contains the muting information. | 

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


## UnmuteSecurityProblem

> UnmuteSecurityProblem(ctx, id).Unmute(unmute).Execute()

Un-mutes the specified security problem. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required security problem.
    unmute := *openapiclient.NewUnmute("Reason_example", "Comment_example") // Unmute | The JSON body of the request. Contains the un-muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityProblemsApi.UnmuteSecurityProblem(context.Background(), id).Unmute(unmute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsApi.UnmuteSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unmute** | [**Unmute**](Unmute.md) | The JSON body of the request. Contains the un-muting information. | 

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

