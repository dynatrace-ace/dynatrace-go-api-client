# JiraNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username of the Jira profile. | [optional] 
**Password** | Pointer to **string** | The password for the Jira profile. | [optional] 
**Url** | Pointer to **string** | The URL of the Jira API endpoint. | [optional] 
**ProjectKey** | Pointer to **string** | The project key of the Jira issue to be created by this notification. | [optional] 
**IssueType** | Pointer to **string** | The type of the Jira issue to be created by this notification. | [optional] 
**Summary** | Pointer to **string** | The summary of the Jira issue to be created by this notification.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Description** | Pointer to **string** | The description of the Jira issue to be created by this notification.    You can use same placeholders as in issue summary. | [optional] 

## Methods

### NewJiraNotificationConfigAllOf

`func NewJiraNotificationConfigAllOf() *JiraNotificationConfigAllOf`

NewJiraNotificationConfigAllOf instantiates a new JiraNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraNotificationConfigAllOfWithDefaults

`func NewJiraNotificationConfigAllOfWithDefaults() *JiraNotificationConfigAllOf`

NewJiraNotificationConfigAllOfWithDefaults instantiates a new JiraNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *JiraNotificationConfigAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *JiraNotificationConfigAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *JiraNotificationConfigAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *JiraNotificationConfigAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *JiraNotificationConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JiraNotificationConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JiraNotificationConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *JiraNotificationConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *JiraNotificationConfigAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JiraNotificationConfigAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JiraNotificationConfigAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *JiraNotificationConfigAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProjectKey

`func (o *JiraNotificationConfigAllOf) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *JiraNotificationConfigAllOf) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *JiraNotificationConfigAllOf) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *JiraNotificationConfigAllOf) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetIssueType

`func (o *JiraNotificationConfigAllOf) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *JiraNotificationConfigAllOf) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *JiraNotificationConfigAllOf) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *JiraNotificationConfigAllOf) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetSummary

`func (o *JiraNotificationConfigAllOf) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *JiraNotificationConfigAllOf) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *JiraNotificationConfigAllOf) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *JiraNotificationConfigAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *JiraNotificationConfigAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JiraNotificationConfigAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JiraNotificationConfigAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JiraNotificationConfigAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


