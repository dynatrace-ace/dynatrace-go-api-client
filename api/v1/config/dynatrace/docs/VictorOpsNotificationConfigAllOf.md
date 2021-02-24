# VictorOpsNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | The API key for the target VictorOps account. | [optional] 
**RoutingKey** | Pointer to **string** | The routing key, defining the group to be notified. | [optional] 
**Message** | Pointer to **string** | The content of the message.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.   | [optional] 

## Methods

### NewVictorOpsNotificationConfigAllOf

`func NewVictorOpsNotificationConfigAllOf() *VictorOpsNotificationConfigAllOf`

NewVictorOpsNotificationConfigAllOf instantiates a new VictorOpsNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVictorOpsNotificationConfigAllOfWithDefaults

`func NewVictorOpsNotificationConfigAllOfWithDefaults() *VictorOpsNotificationConfigAllOf`

NewVictorOpsNotificationConfigAllOfWithDefaults instantiates a new VictorOpsNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *VictorOpsNotificationConfigAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *VictorOpsNotificationConfigAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *VictorOpsNotificationConfigAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *VictorOpsNotificationConfigAllOf) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetRoutingKey

`func (o *VictorOpsNotificationConfigAllOf) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *VictorOpsNotificationConfigAllOf) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *VictorOpsNotificationConfigAllOf) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *VictorOpsNotificationConfigAllOf) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetMessage

`func (o *VictorOpsNotificationConfigAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VictorOpsNotificationConfigAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VictorOpsNotificationConfigAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VictorOpsNotificationConfigAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


