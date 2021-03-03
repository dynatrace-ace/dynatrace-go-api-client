# EventRestEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **int64** | The timestamp of the event detection, in UTC milliseconds. | [optional] 
**EndTime** | Pointer to **int64** | The timestamp of the event closure, in UTC milliseconds | [optional] 
**EntityId** | Pointer to **string** | The ID of the affected Dynatrace entity. | [optional] 
**EntityName** | Pointer to **string** | The name of the affected Dynatrace entity. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity of the event. | [optional] 
**ImpactLevel** | Pointer to **string** | The impact level of the event. It shows what is affected by the problem: infrastructure, service, or application. | [optional] 
**EventType** | Pointer to **string** | The type of the event. | [optional] 
**ResourceId** | Pointer to **string** | The id of the resource the event occurred on. | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource the event occurred on. | [optional] 
**EventStatus** | Pointer to **string** | The state of the event: open or closed. | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | Tags of the Dynatrace entity that raised the event. | [optional] 
**Id** | Pointer to **string** | The encoded ID of the event. The format is *eventID_startTime*.    You should use the value from this field when you need an event ID.  | [optional] 

## Methods

### NewEventRestEntry

`func NewEventRestEntry() *EventRestEntry`

NewEventRestEntry instantiates a new EventRestEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRestEntryWithDefaults

`func NewEventRestEntryWithDefaults() *EventRestEntry`

NewEventRestEntryWithDefaults instantiates a new EventRestEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *EventRestEntry) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EventRestEntry) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EventRestEntry) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EventRestEntry) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *EventRestEntry) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EventRestEntry) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EventRestEntry) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EventRestEntry) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityId

`func (o *EventRestEntry) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EventRestEntry) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EventRestEntry) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EventRestEntry) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *EventRestEntry) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EventRestEntry) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EventRestEntry) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *EventRestEntry) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *EventRestEntry) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *EventRestEntry) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *EventRestEntry) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *EventRestEntry) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetImpactLevel

`func (o *EventRestEntry) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *EventRestEntry) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *EventRestEntry) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *EventRestEntry) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetEventType

`func (o *EventRestEntry) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventRestEntry) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventRestEntry) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventRestEntry) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetResourceId

`func (o *EventRestEntry) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventRestEntry) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventRestEntry) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventRestEntry) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *EventRestEntry) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *EventRestEntry) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *EventRestEntry) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *EventRestEntry) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetEventStatus

`func (o *EventRestEntry) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *EventRestEntry) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *EventRestEntry) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *EventRestEntry) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetTags

`func (o *EventRestEntry) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventRestEntry) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventRestEntry) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventRestEntry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *EventRestEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRestEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRestEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventRestEntry) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


