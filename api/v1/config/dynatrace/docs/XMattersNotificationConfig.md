# XMattersNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the xMatters WebHook. | 
**AcceptAnyCertificate** | **bool** | Accept any, including self-signed and invalid, SSL certificate (&#x60;true&#x60;) or only trusted (&#x60;false&#x60;) certificates. | 
**Headers** | Pointer to [**[]HttpHeader**](HttpHeader.md) | A list of the additional HTTP headers. | [optional] 
**Payload** | **string** | The content of the message.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemDetailsJSON}&#x60;: All problem event details, including root cause, as a JSON object.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 

## Methods

### NewXMattersNotificationConfig

`func NewXMattersNotificationConfig(url string, acceptAnyCertificate bool, payload string, ) *XMattersNotificationConfig`

NewXMattersNotificationConfig instantiates a new XMattersNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXMattersNotificationConfigWithDefaults

`func NewXMattersNotificationConfigWithDefaults() *XMattersNotificationConfig`

NewXMattersNotificationConfigWithDefaults instantiates a new XMattersNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *XMattersNotificationConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *XMattersNotificationConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *XMattersNotificationConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAcceptAnyCertificate

`func (o *XMattersNotificationConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *XMattersNotificationConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *XMattersNotificationConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.


### GetHeaders

`func (o *XMattersNotificationConfig) GetHeaders() []HttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *XMattersNotificationConfig) GetHeadersOk() (*[]HttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *XMattersNotificationConfig) SetHeaders(v []HttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *XMattersNotificationConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPayload

`func (o *XMattersNotificationConfig) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *XMattersNotificationConfig) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *XMattersNotificationConfig) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


