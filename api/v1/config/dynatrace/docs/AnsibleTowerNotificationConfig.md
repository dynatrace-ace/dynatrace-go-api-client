# AnsibleTowerNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTemplateURL** | **string** | The URL of the target Ansible Tower job template. | 
**AcceptAnyCertificate** | **bool** | Accept any, including self-signed and invalid, SSL certificate (&#x60;true&#x60;) or only trusted (&#x60;false&#x60;) certificates. | 
**Username** | **string** | The username of the Ansible Tower account. | 
**Password** | Pointer to **string** | The password for the Ansible Tower account. | [optional] 
**JobTemplateID** | **int32** | The ID of the target Ansible Tower job template. | 
**CustomMessage** | **string** | The custom message of the notification.    This message will be displayed in the extra variables **Message** field of your job template.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 

## Methods

### NewAnsibleTowerNotificationConfig

`func NewAnsibleTowerNotificationConfig(jobTemplateURL string, acceptAnyCertificate bool, username string, jobTemplateID int32, customMessage string, ) *AnsibleTowerNotificationConfig`

NewAnsibleTowerNotificationConfig instantiates a new AnsibleTowerNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleTowerNotificationConfigWithDefaults

`func NewAnsibleTowerNotificationConfigWithDefaults() *AnsibleTowerNotificationConfig`

NewAnsibleTowerNotificationConfigWithDefaults instantiates a new AnsibleTowerNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTemplateURL

`func (o *AnsibleTowerNotificationConfig) GetJobTemplateURL() string`

GetJobTemplateURL returns the JobTemplateURL field if non-nil, zero value otherwise.

### GetJobTemplateURLOk

`func (o *AnsibleTowerNotificationConfig) GetJobTemplateURLOk() (*string, bool)`

GetJobTemplateURLOk returns a tuple with the JobTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateURL

`func (o *AnsibleTowerNotificationConfig) SetJobTemplateURL(v string)`

SetJobTemplateURL sets JobTemplateURL field to given value.


### GetAcceptAnyCertificate

`func (o *AnsibleTowerNotificationConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *AnsibleTowerNotificationConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *AnsibleTowerNotificationConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.


### GetUsername

`func (o *AnsibleTowerNotificationConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnsibleTowerNotificationConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnsibleTowerNotificationConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AnsibleTowerNotificationConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AnsibleTowerNotificationConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AnsibleTowerNotificationConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AnsibleTowerNotificationConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetJobTemplateID

`func (o *AnsibleTowerNotificationConfig) GetJobTemplateID() int32`

GetJobTemplateID returns the JobTemplateID field if non-nil, zero value otherwise.

### GetJobTemplateIDOk

`func (o *AnsibleTowerNotificationConfig) GetJobTemplateIDOk() (*int32, bool)`

GetJobTemplateIDOk returns a tuple with the JobTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateID

`func (o *AnsibleTowerNotificationConfig) SetJobTemplateID(v int32)`

SetJobTemplateID sets JobTemplateID field to given value.


### GetCustomMessage

`func (o *AnsibleTowerNotificationConfig) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *AnsibleTowerNotificationConfig) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *AnsibleTowerNotificationConfig) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


