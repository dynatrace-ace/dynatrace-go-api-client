# ServiceNowNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | Pointer to **string** | The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.    This field is mutually exclusive with the **url** field. You can only use one of them. | [optional] 
**Url** | Pointer to **string** | The URL of the on-premise ServiceNow installation.    This field is mutually exclusive with the **instanceName** field. You can only use one of them. | [optional] 
**Username** | **string** | The username of the ServiceNow account.    Make sure that your user account has the &#x60;rest_service&#x60;, &#x60;web_request_admin&#x60;, and &#x60;x_dynat_ruxit.Integration&#x60; roles. | 
**Password** | Pointer to **string** | The username to the ServiceNow account | [optional] 
**Message** | **string** | The content of the ServiceNow description.     You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 
**SendIncidents** | **bool** | Send incidents into ServiceNow ITSM (&#x60;true&#x60;). | 
**SendEvents** | **bool** | Send events into ServiceNow ITOM (&#x60;true&#x60;). | 

## Methods

### NewServiceNowNotificationConfig

`func NewServiceNowNotificationConfig(username string, message string, sendIncidents bool, sendEvents bool, ) *ServiceNowNotificationConfig`

NewServiceNowNotificationConfig instantiates a new ServiceNowNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowNotificationConfigWithDefaults

`func NewServiceNowNotificationConfigWithDefaults() *ServiceNowNotificationConfig`

NewServiceNowNotificationConfigWithDefaults instantiates a new ServiceNowNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *ServiceNowNotificationConfig) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *ServiceNowNotificationConfig) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *ServiceNowNotificationConfig) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *ServiceNowNotificationConfig) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceNowNotificationConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceNowNotificationConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceNowNotificationConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceNowNotificationConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *ServiceNowNotificationConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceNowNotificationConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceNowNotificationConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ServiceNowNotificationConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceNowNotificationConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceNowNotificationConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceNowNotificationConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceNowNotificationConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceNowNotificationConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceNowNotificationConfig) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSendIncidents

`func (o *ServiceNowNotificationConfig) GetSendIncidents() bool`

GetSendIncidents returns the SendIncidents field if non-nil, zero value otherwise.

### GetSendIncidentsOk

`func (o *ServiceNowNotificationConfig) GetSendIncidentsOk() (*bool, bool)`

GetSendIncidentsOk returns a tuple with the SendIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendIncidents

`func (o *ServiceNowNotificationConfig) SetSendIncidents(v bool)`

SetSendIncidents sets SendIncidents field to given value.


### GetSendEvents

`func (o *ServiceNowNotificationConfig) GetSendEvents() bool`

GetSendEvents returns the SendEvents field if non-nil, zero value otherwise.

### GetSendEventsOk

`func (o *ServiceNowNotificationConfig) GetSendEventsOk() (*bool, bool)`

GetSendEventsOk returns a tuple with the SendEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEvents

`func (o *ServiceNowNotificationConfig) SetSendEvents(v bool)`

SetSendEvents sets SendEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


