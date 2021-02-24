# TrelloNotificationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationKey** | Pointer to **string** | The application key for the Trello account. | [optional] 
**AuthorizationToken** | Pointer to **string** | The application token for the Trello account. | [optional] 
**BoardId** | Pointer to **string** | The Trello board to which the card should be assigned. | [optional] 
**ListId** | Pointer to **string** | The Trello list to which the card should be assigned. | [optional] 
**ResolvedListId** | Pointer to **string** | The Trello list to which the card of the resolved problem should be assigned. | [optional] 
**Text** | Pointer to **string** | The text of the generated Trello card.   You can use the following placeholders:  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | [optional] 
**Description** | Pointer to **string** | The description of the Trello card.    You can use same placeholders as in card text. | [optional] 

## Methods

### NewTrelloNotificationConfigAllOf

`func NewTrelloNotificationConfigAllOf() *TrelloNotificationConfigAllOf`

NewTrelloNotificationConfigAllOf instantiates a new TrelloNotificationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrelloNotificationConfigAllOfWithDefaults

`func NewTrelloNotificationConfigAllOfWithDefaults() *TrelloNotificationConfigAllOf`

NewTrelloNotificationConfigAllOfWithDefaults instantiates a new TrelloNotificationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationKey

`func (o *TrelloNotificationConfigAllOf) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *TrelloNotificationConfigAllOf) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *TrelloNotificationConfigAllOf) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *TrelloNotificationConfigAllOf) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### GetAuthorizationToken

`func (o *TrelloNotificationConfigAllOf) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *TrelloNotificationConfigAllOf) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *TrelloNotificationConfigAllOf) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *TrelloNotificationConfigAllOf) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### GetBoardId

`func (o *TrelloNotificationConfigAllOf) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *TrelloNotificationConfigAllOf) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *TrelloNotificationConfigAllOf) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *TrelloNotificationConfigAllOf) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetListId

`func (o *TrelloNotificationConfigAllOf) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *TrelloNotificationConfigAllOf) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *TrelloNotificationConfigAllOf) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *TrelloNotificationConfigAllOf) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetResolvedListId

`func (o *TrelloNotificationConfigAllOf) GetResolvedListId() string`

GetResolvedListId returns the ResolvedListId field if non-nil, zero value otherwise.

### GetResolvedListIdOk

`func (o *TrelloNotificationConfigAllOf) GetResolvedListIdOk() (*string, bool)`

GetResolvedListIdOk returns a tuple with the ResolvedListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedListId

`func (o *TrelloNotificationConfigAllOf) SetResolvedListId(v string)`

SetResolvedListId sets ResolvedListId field to given value.

### HasResolvedListId

`func (o *TrelloNotificationConfigAllOf) HasResolvedListId() bool`

HasResolvedListId returns a boolean if a field has been set.

### GetText

`func (o *TrelloNotificationConfigAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TrelloNotificationConfigAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TrelloNotificationConfigAllOf) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TrelloNotificationConfigAllOf) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDescription

`func (o *TrelloNotificationConfigAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrelloNotificationConfigAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrelloNotificationConfigAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrelloNotificationConfigAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


