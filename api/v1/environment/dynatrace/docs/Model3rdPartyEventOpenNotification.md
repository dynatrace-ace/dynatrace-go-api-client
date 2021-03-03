# Model3rdPartyEventOpenNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** | The ID of the third-party synthetic monitor. | 
**EventId** | **string** | The unique ID of the event. | 
**Name** | **string** | The name of the event. | 
**EventType** | **string** | The type of the event. | 
**Reason** | **string** | The cause of the event. | 
**StartTimestamp** | **int64** | The start timestamp of the event, in UTC milliseconds. | 
**LocationIds** | **[]string** | The list of IDs of third-party synthetic locations where the event happens. | 

## Methods

### NewModel3rdPartyEventOpenNotification

`func NewModel3rdPartyEventOpenNotification(testId string, eventId string, name string, eventType string, reason string, startTimestamp int64, locationIds []string, ) *Model3rdPartyEventOpenNotification`

NewModel3rdPartyEventOpenNotification instantiates a new Model3rdPartyEventOpenNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartyEventOpenNotificationWithDefaults

`func NewModel3rdPartyEventOpenNotificationWithDefaults() *Model3rdPartyEventOpenNotification`

NewModel3rdPartyEventOpenNotificationWithDefaults instantiates a new Model3rdPartyEventOpenNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *Model3rdPartyEventOpenNotification) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *Model3rdPartyEventOpenNotification) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *Model3rdPartyEventOpenNotification) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetEventId

`func (o *Model3rdPartyEventOpenNotification) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Model3rdPartyEventOpenNotification) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Model3rdPartyEventOpenNotification) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetName

`func (o *Model3rdPartyEventOpenNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model3rdPartyEventOpenNotification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model3rdPartyEventOpenNotification) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *Model3rdPartyEventOpenNotification) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Model3rdPartyEventOpenNotification) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Model3rdPartyEventOpenNotification) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetReason

`func (o *Model3rdPartyEventOpenNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Model3rdPartyEventOpenNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Model3rdPartyEventOpenNotification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartTimestamp

`func (o *Model3rdPartyEventOpenNotification) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Model3rdPartyEventOpenNotification) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Model3rdPartyEventOpenNotification) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetLocationIds

`func (o *Model3rdPartyEventOpenNotification) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *Model3rdPartyEventOpenNotification) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *Model3rdPartyEventOpenNotification) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


