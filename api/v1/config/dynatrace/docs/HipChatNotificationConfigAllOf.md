# HipChatNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the HipChat WebHook.   This is confidential information, therefore GET requests return this field with the &#x60;null&#x60; value, and it is optional for PUT requests. | [optional] 
**Message** | Pointer to **string** | The content of the notification message.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 

## Methods

### NewHipChatNotificationConfigAllOf

`func NewHipChatNotificationConfigAllOf() *HipChatNotificationConfigAllOf`

NewHipChatNotificationConfigAllOf instantiates a new HipChatNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHipChatNotificationConfigAllOfWithDefaults

`func NewHipChatNotificationConfigAllOfWithDefaults() *HipChatNotificationConfigAllOf`

NewHipChatNotificationConfigAllOfWithDefaults instantiates a new HipChatNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HipChatNotificationConfigAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HipChatNotificationConfigAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HipChatNotificationConfigAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HipChatNotificationConfigAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMessage

`func (o *HipChatNotificationConfigAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HipChatNotificationConfigAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HipChatNotificationConfigAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HipChatNotificationConfigAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


