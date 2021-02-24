# SlackNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the Slack WebHook.   This is confidential information, therefore GET requests return this field with the &#x60;null&#x60; value, and it is optional for PUT requests. | [optional] 
**Channel** | Pointer to **string** | The channel (for example, &#x60;#general&#x60;) or the user (for example, &#x60;@john.smith&#x60;) to send the message to. | [optional] 
**Title** | Pointer to **string** | The content of the message.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 

## Methods

### NewSlackNotificationConfigAllOf

`func NewSlackNotificationConfigAllOf() *SlackNotificationConfigAllOf`

NewSlackNotificationConfigAllOf instantiates a new SlackNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackNotificationConfigAllOfWithDefaults

`func NewSlackNotificationConfigAllOfWithDefaults() *SlackNotificationConfigAllOf`

NewSlackNotificationConfigAllOfWithDefaults instantiates a new SlackNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SlackNotificationConfigAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SlackNotificationConfigAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SlackNotificationConfigAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SlackNotificationConfigAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetChannel

`func (o *SlackNotificationConfigAllOf) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SlackNotificationConfigAllOf) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SlackNotificationConfigAllOf) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SlackNotificationConfigAllOf) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTitle

`func (o *SlackNotificationConfigAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SlackNotificationConfigAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SlackNotificationConfigAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SlackNotificationConfigAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


