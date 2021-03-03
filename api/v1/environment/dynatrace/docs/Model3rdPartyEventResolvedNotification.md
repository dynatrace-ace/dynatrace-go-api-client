# Model3rdPartyEventResolvedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** | The ID of the third-party synthetic monitor. | 
**EventId** | **string** | The unique ID of the event. | 
**EndTimestamp** | **int64** | The end timestamp of the event, in UTC milliseconds. | 

## Methods

### NewModel3rdPartyEventResolvedNotification

`func NewModel3rdPartyEventResolvedNotification(testId string, eventId string, endTimestamp int64, ) *Model3rdPartyEventResolvedNotification`

NewModel3rdPartyEventResolvedNotification instantiates a new Model3rdPartyEventResolvedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartyEventResolvedNotificationWithDefaults

`func NewModel3rdPartyEventResolvedNotificationWithDefaults() *Model3rdPartyEventResolvedNotification`

NewModel3rdPartyEventResolvedNotificationWithDefaults instantiates a new Model3rdPartyEventResolvedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *Model3rdPartyEventResolvedNotification) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *Model3rdPartyEventResolvedNotification) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *Model3rdPartyEventResolvedNotification) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetEventId

`func (o *Model3rdPartyEventResolvedNotification) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Model3rdPartyEventResolvedNotification) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Model3rdPartyEventResolvedNotification) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEndTimestamp

`func (o *Model3rdPartyEventResolvedNotification) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *Model3rdPartyEventResolvedNotification) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *Model3rdPartyEventResolvedNotification) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


