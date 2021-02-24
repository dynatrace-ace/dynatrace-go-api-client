# OpsGenieNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | The API key to access OpsGenie. | [optional] 
**Domain** | Pointer to **string** | The region domain of the OpsGenie. | [optional] 
**Message** | Pointer to **string** | The content of the message.   You can use the following placeholders:  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.   | [optional] 

## Methods

### NewOpsGenieNotificationConfigAllOf

`func NewOpsGenieNotificationConfigAllOf() *OpsGenieNotificationConfigAllOf`

NewOpsGenieNotificationConfigAllOf instantiates a new OpsGenieNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsGenieNotificationConfigAllOfWithDefaults

`func NewOpsGenieNotificationConfigAllOfWithDefaults() *OpsGenieNotificationConfigAllOf`

NewOpsGenieNotificationConfigAllOfWithDefaults instantiates a new OpsGenieNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OpsGenieNotificationConfigAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OpsGenieNotificationConfigAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OpsGenieNotificationConfigAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OpsGenieNotificationConfigAllOf) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDomain

`func (o *OpsGenieNotificationConfigAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OpsGenieNotificationConfigAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OpsGenieNotificationConfigAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OpsGenieNotificationConfigAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMessage

`func (o *OpsGenieNotificationConfigAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OpsGenieNotificationConfigAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OpsGenieNotificationConfigAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OpsGenieNotificationConfigAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


