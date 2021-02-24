# NotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the notification configuration. | [optional] 
**Name** | **string** | The name of the notification configuration. | 
**AlertingProfile** | **string** | The ID of the associated alerting profile. | 
**Active** | **bool** | The configuration is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EMAIL&#x60; -&gt; EmailNotificationConfig  * &#x60;PAGER_DUTY&#x60; -&gt; PagerDutyNotificationConfig  * &#x60;WEBHOOK&#x60; -&gt; WebHookNotificationConfig  * &#x60;SLACK&#x60; -&gt; SlackNotificationConfig  * &#x60;HIPCHAT&#x60; -&gt; HipChatNotificationConfig  * &#x60;VICTOROPS&#x60; -&gt; VictorOpsNotificationConfig  * &#x60;SERVICE_NOW&#x60; -&gt; ServiceNowNotificationConfig  * &#x60;XMATTERS&#x60; -&gt; XMattersNotificationConfig  * &#x60;ANSIBLETOWER&#x60; -&gt; AnsibleTowerNotificationConfig  * &#x60;OPS_GENIE&#x60; -&gt; OpsGenieNotificationConfig  * &#x60;JIRA&#x60; -&gt; JiraNotificationConfig  * &#x60;TRELLO&#x60; -&gt; TrelloNotificationConfig   | 
**JobTemplateURL** | Pointer to **string** | The URL of the target Ansible Tower job template. | [optional] 
**AcceptAnyCertificate** | Pointer to **bool** | Accept any, including self-signed and invalid, SSL certificate (&#x60;true&#x60;) or only trusted (&#x60;false&#x60;) certificates. | [optional] 
**Username** | Pointer to **string** | The username required for authentication. | [optional] 
**Password** | Pointer to **string** | The password required for authentication. | [optional] 
**JobTemplateID** | Pointer to **int32** | The ID of the target Ansible Tower job template. | [optional] 
**CustomMessage** | Pointer to **string** | The custom message of the notification.    This message will be displayed in the extra variables **Message** field of your job template.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Subject** | Pointer to **string** | The subject of the email notifications. | [optional] 
**Body** | Pointer to **string** | The template of the email notification.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemDetailsJSON}&#x60;: All problem event details, including root cause, as a JSON object.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Receivers** | **[]string** | The list of the email recipients. | 
**CcReceivers** | **[]string** | The list of the email CC-recipients. | 
**BccReceivers** | **[]string** | The list of the email BCC-recipients. | 
**ProjectKey** | Pointer to **string** | The project key of the Jira issue to be created by this notification. | [optional] 
**IssueType** | Pointer to **string** | The type of the Jira issue to be created by this notification. | [optional] 
**Summary** | Pointer to **string** | The summary of the Jira issue to be created by this notification.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Description** | Pointer to **string** | The description of the notification. | [optional] 
**ApiKey** | Pointer to **string** | The API key of the target account. | [optional] 
**Domain** | Pointer to **string** | The region domain of the OpsGenie. | [optional] 
**Account** | Pointer to **string** | The name of the PagerDuty account. | [optional] 
**ServiceApiKey** | Pointer to **string** | The API key to access PagerDuty. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. | [optional] 
**InstanceName** | Pointer to **string** | The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.    This field is mutually exclusive with the **url** field. You can only use one of them. | [optional] 
**Url** | Pointer to **string** | The URL of the notification endpoint | [optional] 
**Message** | Pointer to **string** | The content of the message | [optional] 
**SendIncidents** | Pointer to **bool** | Send incidents into ServiceNow ITSM (&#x60;true&#x60;). | [optional] 
**SendEvents** | Pointer to **bool** | Send events into ServiceNow ITOM (&#x60;true&#x60;). | [optional] 
**Channel** | Pointer to **string** | The channel (for example, &#x60;#general&#x60;) or the user (for example, &#x60;@john.smith&#x60;) to send the message to. | [optional] 
**Title** | Pointer to **string** | The content of the message.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**ApplicationKey** | Pointer to **string** | The application key for the Trello account. | [optional] 
**AuthorizationToken** | Pointer to **string** | The application token for the Trello account. | [optional] 
**BoardId** | Pointer to **string** | The Trello board to which the card should be assigned. | [optional] 
**ListId** | Pointer to **string** | The Trello list to which the card should be assigned. | [optional] 
**ResolvedListId** | Pointer to **string** | The Trello list to which the card of the resolved problem should be assigned. | [optional] 
**Text** | Pointer to **string** | The text of the generated Trello card.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**RoutingKey** | Pointer to **string** | The routing key, defining the group to be notified. | [optional] 
**Payload** | Pointer to **string** | The content of the notification message.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemDetailsJSON}&#x60;: All problem event details, including root cause, as a JSON object.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Headers** | Pointer to [**[]HttpHeader**](HttpHeader.md) | A list of the additional HTTP headers. | [optional] 
**NotifyEventMergesEnabled** | Pointer to **bool** | Call webhook if new events merge into existing problems. | [optional] 

## Methods

### NewNotificationConfig

`func NewNotificationConfig(name string, alertingProfile string, active bool, type_ string, receivers []string, ccReceivers []string, bccReceivers []string, ) *NotificationConfig`

NewNotificationConfig instantiates a new NotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigWithDefaults

`func NewNotificationConfigWithDefaults() *NotificationConfig`

NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationConfig) SetName(v string)`

SetName sets Name field to given value.


### GetAlertingProfile

`func (o *NotificationConfig) GetAlertingProfile() string`

GetAlertingProfile returns the AlertingProfile field if non-nil, zero value otherwise.

### GetAlertingProfileOk

`func (o *NotificationConfig) GetAlertingProfileOk() (*string, bool)`

GetAlertingProfileOk returns a tuple with the AlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfile

`func (o *NotificationConfig) SetAlertingProfile(v string)`

SetAlertingProfile sets AlertingProfile field to given value.


### GetActive

`func (o *NotificationConfig) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NotificationConfig) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NotificationConfig) SetActive(v bool)`

SetActive sets Active field to given value.


### GetType

`func (o *NotificationConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationConfig) SetType(v string)`

SetType sets Type field to given value.


### GetJobTemplateURL

`func (o *NotificationConfig) GetJobTemplateURL() string`

GetJobTemplateURL returns the JobTemplateURL field if non-nil, zero value otherwise.

### GetJobTemplateURLOk

`func (o *NotificationConfig) GetJobTemplateURLOk() (*string, bool)`

GetJobTemplateURLOk returns a tuple with the JobTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateURL

`func (o *NotificationConfig) SetJobTemplateURL(v string)`

SetJobTemplateURL sets JobTemplateURL field to given value.

### HasJobTemplateURL

`func (o *NotificationConfig) HasJobTemplateURL() bool`

HasJobTemplateURL returns a boolean if a field has been set.

### GetAcceptAnyCertificate

`func (o *NotificationConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *NotificationConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *NotificationConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.

### HasAcceptAnyCertificate

`func (o *NotificationConfig) HasAcceptAnyCertificate() bool`

HasAcceptAnyCertificate returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetJobTemplateID

`func (o *NotificationConfig) GetJobTemplateID() int32`

GetJobTemplateID returns the JobTemplateID field if non-nil, zero value otherwise.

### GetJobTemplateIDOk

`func (o *NotificationConfig) GetJobTemplateIDOk() (*int32, bool)`

GetJobTemplateIDOk returns a tuple with the JobTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateID

`func (o *NotificationConfig) SetJobTemplateID(v int32)`

SetJobTemplateID sets JobTemplateID field to given value.

### HasJobTemplateID

`func (o *NotificationConfig) HasJobTemplateID() bool`

HasJobTemplateID returns a boolean if a field has been set.

### GetCustomMessage

`func (o *NotificationConfig) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *NotificationConfig) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *NotificationConfig) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *NotificationConfig) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetSubject

`func (o *NotificationConfig) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationConfig) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationConfig) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotificationConfig) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *NotificationConfig) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotificationConfig) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotificationConfig) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotificationConfig) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetReceivers

`func (o *NotificationConfig) GetReceivers() []string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *NotificationConfig) GetReceiversOk() (*[]string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *NotificationConfig) SetReceivers(v []string)`

SetReceivers sets Receivers field to given value.


### GetCcReceivers

`func (o *NotificationConfig) GetCcReceivers() []string`

GetCcReceivers returns the CcReceivers field if non-nil, zero value otherwise.

### GetCcReceiversOk

`func (o *NotificationConfig) GetCcReceiversOk() (*[]string, bool)`

GetCcReceiversOk returns a tuple with the CcReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcReceivers

`func (o *NotificationConfig) SetCcReceivers(v []string)`

SetCcReceivers sets CcReceivers field to given value.


### GetBccReceivers

`func (o *NotificationConfig) GetBccReceivers() []string`

GetBccReceivers returns the BccReceivers field if non-nil, zero value otherwise.

### GetBccReceiversOk

`func (o *NotificationConfig) GetBccReceiversOk() (*[]string, bool)`

GetBccReceiversOk returns a tuple with the BccReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccReceivers

`func (o *NotificationConfig) SetBccReceivers(v []string)`

SetBccReceivers sets BccReceivers field to given value.


### GetProjectKey

`func (o *NotificationConfig) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *NotificationConfig) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *NotificationConfig) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *NotificationConfig) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetIssueType

`func (o *NotificationConfig) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *NotificationConfig) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *NotificationConfig) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *NotificationConfig) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetSummary

`func (o *NotificationConfig) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NotificationConfig) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NotificationConfig) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NotificationConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiKey

`func (o *NotificationConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *NotificationConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *NotificationConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *NotificationConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDomain

`func (o *NotificationConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NotificationConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NotificationConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NotificationConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAccount

`func (o *NotificationConfig) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NotificationConfig) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NotificationConfig) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NotificationConfig) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetServiceApiKey

`func (o *NotificationConfig) GetServiceApiKey() string`

GetServiceApiKey returns the ServiceApiKey field if non-nil, zero value otherwise.

### GetServiceApiKeyOk

`func (o *NotificationConfig) GetServiceApiKeyOk() (*string, bool)`

GetServiceApiKeyOk returns a tuple with the ServiceApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiKey

`func (o *NotificationConfig) SetServiceApiKey(v string)`

SetServiceApiKey sets ServiceApiKey field to given value.

### HasServiceApiKey

`func (o *NotificationConfig) HasServiceApiKey() bool`

HasServiceApiKey returns a boolean if a field has been set.

### GetServiceName

`func (o *NotificationConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NotificationConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NotificationConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *NotificationConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetInstanceName

`func (o *NotificationConfig) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NotificationConfig) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NotificationConfig) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *NotificationConfig) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMessage

`func (o *NotificationConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSendIncidents

`func (o *NotificationConfig) GetSendIncidents() bool`

GetSendIncidents returns the SendIncidents field if non-nil, zero value otherwise.

### GetSendIncidentsOk

`func (o *NotificationConfig) GetSendIncidentsOk() (*bool, bool)`

GetSendIncidentsOk returns a tuple with the SendIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendIncidents

`func (o *NotificationConfig) SetSendIncidents(v bool)`

SetSendIncidents sets SendIncidents field to given value.

### HasSendIncidents

`func (o *NotificationConfig) HasSendIncidents() bool`

HasSendIncidents returns a boolean if a field has been set.

### GetSendEvents

`func (o *NotificationConfig) GetSendEvents() bool`

GetSendEvents returns the SendEvents field if non-nil, zero value otherwise.

### GetSendEventsOk

`func (o *NotificationConfig) GetSendEventsOk() (*bool, bool)`

GetSendEventsOk returns a tuple with the SendEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEvents

`func (o *NotificationConfig) SetSendEvents(v bool)`

SetSendEvents sets SendEvents field to given value.

### HasSendEvents

`func (o *NotificationConfig) HasSendEvents() bool`

HasSendEvents returns a boolean if a field has been set.

### GetChannel

`func (o *NotificationConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *NotificationConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *NotificationConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *NotificationConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTitle

`func (o *NotificationConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetApplicationKey

`func (o *NotificationConfig) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *NotificationConfig) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *NotificationConfig) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *NotificationConfig) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### GetAuthorizationToken

`func (o *NotificationConfig) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *NotificationConfig) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *NotificationConfig) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *NotificationConfig) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### GetBoardId

`func (o *NotificationConfig) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *NotificationConfig) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *NotificationConfig) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *NotificationConfig) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetListId

`func (o *NotificationConfig) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *NotificationConfig) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *NotificationConfig) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *NotificationConfig) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetResolvedListId

`func (o *NotificationConfig) GetResolvedListId() string`

GetResolvedListId returns the ResolvedListId field if non-nil, zero value otherwise.

### GetResolvedListIdOk

`func (o *NotificationConfig) GetResolvedListIdOk() (*string, bool)`

GetResolvedListIdOk returns a tuple with the ResolvedListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedListId

`func (o *NotificationConfig) SetResolvedListId(v string)`

SetResolvedListId sets ResolvedListId field to given value.

### HasResolvedListId

`func (o *NotificationConfig) HasResolvedListId() bool`

HasResolvedListId returns a boolean if a field has been set.

### GetText

`func (o *NotificationConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NotificationConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NotificationConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *NotificationConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetRoutingKey

`func (o *NotificationConfig) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *NotificationConfig) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *NotificationConfig) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *NotificationConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetPayload

`func (o *NotificationConfig) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NotificationConfig) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NotificationConfig) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NotificationConfig) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetHeaders

`func (o *NotificationConfig) GetHeaders() []HttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NotificationConfig) GetHeadersOk() (*[]HttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NotificationConfig) SetHeaders(v []HttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *NotificationConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetNotifyEventMergesEnabled

`func (o *NotificationConfig) GetNotifyEventMergesEnabled() bool`

GetNotifyEventMergesEnabled returns the NotifyEventMergesEnabled field if non-nil, zero value otherwise.

### GetNotifyEventMergesEnabledOk

`func (o *NotificationConfig) GetNotifyEventMergesEnabledOk() (*bool, bool)`

GetNotifyEventMergesEnabledOk returns a tuple with the NotifyEventMergesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEventMergesEnabled

`func (o *NotificationConfig) SetNotifyEventMergesEnabled(v bool)`

SetNotifyEventMergesEnabled sets NotifyEventMergesEnabled field to given value.

### HasNotifyEventMergesEnabled

`func (o *NotificationConfig) HasNotifyEventMergesEnabled() bool`

HasNotifyEventMergesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


