# EventStoreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoredEventIds** | Pointer to **[]int64** | List of event IDs for information-only-events.    This field is provided for compatibility reasons. You should use the values from the **storedIds** field instead. | [optional] 
**StoredIds** | Pointer to **[]string** | List of **encoded** event IDs for information-only-events. | [optional] 
**StoredCorrelationIds** | Pointer to **[]string** | List of correlation IDs for problem-opening-events. | [optional] 

## Methods

### NewEventStoreResult

`func NewEventStoreResult() *EventStoreResult`

NewEventStoreResult instantiates a new EventStoreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStoreResultWithDefaults

`func NewEventStoreResultWithDefaults() *EventStoreResult`

NewEventStoreResultWithDefaults instantiates a new EventStoreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoredEventIds

`func (o *EventStoreResult) GetStoredEventIds() []int64`

GetStoredEventIds returns the StoredEventIds field if non-nil, zero value otherwise.

### GetStoredEventIdsOk

`func (o *EventStoreResult) GetStoredEventIdsOk() (*[]int64, bool)`

GetStoredEventIdsOk returns a tuple with the StoredEventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredEventIds

`func (o *EventStoreResult) SetStoredEventIds(v []int64)`

SetStoredEventIds sets StoredEventIds field to given value.

### HasStoredEventIds

`func (o *EventStoreResult) HasStoredEventIds() bool`

HasStoredEventIds returns a boolean if a field has been set.

### GetStoredIds

`func (o *EventStoreResult) GetStoredIds() []string`

GetStoredIds returns the StoredIds field if non-nil, zero value otherwise.

### GetStoredIdsOk

`func (o *EventStoreResult) GetStoredIdsOk() (*[]string, bool)`

GetStoredIdsOk returns a tuple with the StoredIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredIds

`func (o *EventStoreResult) SetStoredIds(v []string)`

SetStoredIds sets StoredIds field to given value.

### HasStoredIds

`func (o *EventStoreResult) HasStoredIds() bool`

HasStoredIds returns a boolean if a field has been set.

### GetStoredCorrelationIds

`func (o *EventStoreResult) GetStoredCorrelationIds() []string`

GetStoredCorrelationIds returns the StoredCorrelationIds field if non-nil, zero value otherwise.

### GetStoredCorrelationIdsOk

`func (o *EventStoreResult) GetStoredCorrelationIdsOk() (*[]string, bool)`

GetStoredCorrelationIdsOk returns a tuple with the StoredCorrelationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredCorrelationIds

`func (o *EventStoreResult) SetStoredCorrelationIds(v []string)`

SetStoredCorrelationIds sets StoredCorrelationIds field to given value.

### HasStoredCorrelationIds

`func (o *EventStoreResult) HasStoredCorrelationIds() bool`

HasStoredCorrelationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


