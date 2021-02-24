# AnsibleTowerNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTemplateURL** | Pointer to **string** | The URL of the target Ansible Tower job template. | [optional] 
**AcceptAnyCertificate** | Pointer to **bool** | Accept any, including self-signed and invalid, SSL certificate (&#x60;true&#x60;) or only trusted (&#x60;false&#x60;) certificates. | [optional] 
**Username** | Pointer to **string** | The username of the Ansible Tower account. | [optional] 
**Password** | Pointer to **string** | The password for the Ansible Tower account. | [optional] 
**JobTemplateID** | Pointer to **int32** | The ID of the target Ansible Tower job template. | [optional] 
**CustomMessage** | Pointer to **string** | The custom message of the notification.    This message will be displayed in the extra variables **Message** field of your job template.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 

## Methods

### NewAnsibleTowerNotificationConfigAllOf

`func NewAnsibleTowerNotificationConfigAllOf() *AnsibleTowerNotificationConfigAllOf`

NewAnsibleTowerNotificationConfigAllOf instantiates a new AnsibleTowerNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleTowerNotificationConfigAllOfWithDefaults

`func NewAnsibleTowerNotificationConfigAllOfWithDefaults() *AnsibleTowerNotificationConfigAllOf`

NewAnsibleTowerNotificationConfigAllOfWithDefaults instantiates a new AnsibleTowerNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTemplateURL

`func (o *AnsibleTowerNotificationConfigAllOf) GetJobTemplateURL() string`

GetJobTemplateURL returns the JobTemplateURL field if non-nil, zero value otherwise.

### GetJobTemplateURLOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetJobTemplateURLOk() (*string, bool)`

GetJobTemplateURLOk returns a tuple with the JobTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateURL

`func (o *AnsibleTowerNotificationConfigAllOf) SetJobTemplateURL(v string)`

SetJobTemplateURL sets JobTemplateURL field to given value.

### HasJobTemplateURL

`func (o *AnsibleTowerNotificationConfigAllOf) HasJobTemplateURL() bool`

HasJobTemplateURL returns a boolean if a field has been set.

### GetAcceptAnyCertificate

`func (o *AnsibleTowerNotificationConfigAllOf) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *AnsibleTowerNotificationConfigAllOf) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.

### HasAcceptAnyCertificate

`func (o *AnsibleTowerNotificationConfigAllOf) HasAcceptAnyCertificate() bool`

HasAcceptAnyCertificate returns a boolean if a field has been set.

### GetUsername

`func (o *AnsibleTowerNotificationConfigAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnsibleTowerNotificationConfigAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AnsibleTowerNotificationConfigAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AnsibleTowerNotificationConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AnsibleTowerNotificationConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AnsibleTowerNotificationConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetJobTemplateID

`func (o *AnsibleTowerNotificationConfigAllOf) GetJobTemplateID() int32`

GetJobTemplateID returns the JobTemplateID field if non-nil, zero value otherwise.

### GetJobTemplateIDOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetJobTemplateIDOk() (*int32, bool)`

GetJobTemplateIDOk returns a tuple with the JobTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTemplateID

`func (o *AnsibleTowerNotificationConfigAllOf) SetJobTemplateID(v int32)`

SetJobTemplateID sets JobTemplateID field to given value.

### HasJobTemplateID

`func (o *AnsibleTowerNotificationConfigAllOf) HasJobTemplateID() bool`

HasJobTemplateID returns a boolean if a field has been set.

### GetCustomMessage

`func (o *AnsibleTowerNotificationConfigAllOf) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *AnsibleTowerNotificationConfigAllOf) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *AnsibleTowerNotificationConfigAllOf) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *AnsibleTowerNotificationConfigAllOf) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


