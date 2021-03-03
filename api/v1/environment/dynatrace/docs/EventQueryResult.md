# EventQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** | The cursor for the next 150 events, fitting the specified criteria.    Set this value for the **cursor** query parameter. Without it you&#39;ll get the first 150 events again.   You don&#39;t have to specify any additional parameters, because the cursor already contains all of them. | [optional] 
**From** | Pointer to **int64** | Start of the query timeframe. | [optional] 
**To** | Pointer to **int64** | End of the query timeframe. | [optional] 
**TotalEventCount** | Pointer to **int64** | The total amount of events, fitting the specified criteria. | [optional] 
**Events** | Pointer to [**[]EventRestEntry**](EventRestEntry.md) | The list of events. | [optional] 

## Methods

### NewEventQueryResult

`func NewEventQueryResult() *EventQueryResult`

NewEventQueryResult instantiates a new EventQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryResultWithDefaults

`func NewEventQueryResultWithDefaults() *EventQueryResult`

NewEventQueryResultWithDefaults instantiates a new EventQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *EventQueryResult) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *EventQueryResult) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *EventQueryResult) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *EventQueryResult) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetFrom

`func (o *EventQueryResult) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EventQueryResult) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EventQueryResult) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EventQueryResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *EventQueryResult) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EventQueryResult) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EventQueryResult) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *EventQueryResult) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalEventCount

`func (o *EventQueryResult) GetTotalEventCount() int64`

GetTotalEventCount returns the TotalEventCount field if non-nil, zero value otherwise.

### GetTotalEventCountOk

`func (o *EventQueryResult) GetTotalEventCountOk() (*int64, bool)`

GetTotalEventCountOk returns a tuple with the TotalEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEventCount

`func (o *EventQueryResult) SetTotalEventCount(v int64)`

SetTotalEventCount sets TotalEventCount field to given value.

### HasTotalEventCount

`func (o *EventQueryResult) HasTotalEventCount() bool`

HasTotalEventCount returns a boolean if a field has been set.

### GetEvents

`func (o *EventQueryResult) GetEvents() []EventRestEntry`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventQueryResult) GetEventsOk() (*[]EventRestEntry, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventQueryResult) SetEvents(v []EventRestEntry)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventQueryResult) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


