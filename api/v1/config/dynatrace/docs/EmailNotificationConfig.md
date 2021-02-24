# EmailNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The subject of the email notifications. | 
**Body** | **string** | The template of the email notification.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemDetailsJSON}&#x60;: All problem event details, including root cause, as a JSON object.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 
**Receivers** | **[]string** | The list of the email recipients. | 
**CcReceivers** | **[]string** | The list of the email CC-recipients. | 
**BccReceivers** | **[]string** | The list of the email BCC-recipients. | 

## Methods

### NewEmailNotificationConfig

`func NewEmailNotificationConfig(subject string, body string, receivers []string, ccReceivers []string, bccReceivers []string, ) *EmailNotificationConfig`

NewEmailNotificationConfig instantiates a new EmailNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationConfigWithDefaults

`func NewEmailNotificationConfigWithDefaults() *EmailNotificationConfig`

NewEmailNotificationConfigWithDefaults instantiates a new EmailNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailNotificationConfig) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailNotificationConfig) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailNotificationConfig) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *EmailNotificationConfig) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailNotificationConfig) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailNotificationConfig) SetBody(v string)`

SetBody sets Body field to given value.


### GetReceivers

`func (o *EmailNotificationConfig) GetReceivers() []string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *EmailNotificationConfig) GetReceiversOk() (*[]string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *EmailNotificationConfig) SetReceivers(v []string)`

SetReceivers sets Receivers field to given value.


### GetCcReceivers

`func (o *EmailNotificationConfig) GetCcReceivers() []string`

GetCcReceivers returns the CcReceivers field if non-nil, zero value otherwise.

### GetCcReceiversOk

`func (o *EmailNotificationConfig) GetCcReceiversOk() (*[]string, bool)`

GetCcReceiversOk returns a tuple with the CcReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcReceivers

`func (o *EmailNotificationConfig) SetCcReceivers(v []string)`

SetCcReceivers sets CcReceivers field to given value.


### GetBccReceivers

`func (o *EmailNotificationConfig) GetBccReceivers() []string`

GetBccReceivers returns the BccReceivers field if non-nil, zero value otherwise.

### GetBccReceiversOk

`func (o *EmailNotificationConfig) GetBccReceiversOk() (*[]string, bool)`

GetBccReceiversOk returns a tuple with the BccReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccReceivers

`func (o *EmailNotificationConfig) SetBccReceivers(v []string)`

SetBccReceivers sets BccReceivers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


