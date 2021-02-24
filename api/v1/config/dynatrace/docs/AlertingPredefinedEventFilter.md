# AlertingPredefinedEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The type of the predefined event. | 
**Negate** | **bool** | The alert triggers when the problem of specified severity arises while the specified event **is** happening (&#x60;false&#x60;) or while the specified event is **not** happening (&#x60;true&#x60;).    For example, if you chose the Slowdown (&#x60;PERFORMANCE&#x60;) severity and Unexpected high traffic (&#x60;APPLICATION_UNEXPECTED_HIGH_LOAD&#x60;) event with **negate** set to &#x60;true&#x60;, the alerting profile will trigger only when the slowdown problem is raised while there is no unexpected high traffic event.   Consider the following use case as an example. The Slowdown (&#x60;PERFORMANCE&#x60;) severity rule is set. Depending on the configuration of the event filter (Unexpected high traffic (&#x60;APPLICATION_UNEXPECTED_HIGH_LOAD&#x60;) event is used as an example), the behavior of the alerting profile is one of the following:* **negate** is set to &#x60;false&#x60;: The alert triggers when the slowdown problem is raised while unexpected high traffic event is happening.  * **negate** is set to &#x60;true&#x60;: The alert triggers when the slowdown problem is raised while there is no unexpected high traffic event.   * no event rule is set: The alert triggers when the slowdown problem is raised, regardless of any events. | 

## Methods

### NewAlertingPredefinedEventFilter

`func NewAlertingPredefinedEventFilter(eventType string, negate bool, ) *AlertingPredefinedEventFilter`

NewAlertingPredefinedEventFilter instantiates a new AlertingPredefinedEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingPredefinedEventFilterWithDefaults

`func NewAlertingPredefinedEventFilterWithDefaults() *AlertingPredefinedEventFilter`

NewAlertingPredefinedEventFilterWithDefaults instantiates a new AlertingPredefinedEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *AlertingPredefinedEventFilter) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AlertingPredefinedEventFilter) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AlertingPredefinedEventFilter) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetNegate

`func (o *AlertingPredefinedEventFilter) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AlertingPredefinedEventFilter) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AlertingPredefinedEventFilter) SetNegate(v bool)`

SetNegate sets Negate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


