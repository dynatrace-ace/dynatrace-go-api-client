# EventRestImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The ID of the affected Dynatrace entity. | [optional] 
**EntityName** | Pointer to **string** | The name of the affected Dynatrace entity. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity of the event. | [optional] 
**ImpactLevel** | Pointer to **string** | The impact level of the event. It shows what is affected by the problem: infrastructure, service, or application. | [optional] 
**EventType** | Pointer to **string** | The type of the event. | [optional] 
**ResourceId** | Pointer to **string** | The id of the resource the event occurred on. | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource the event occurred on. | [optional] 

## Methods

### NewEventRestImpact

`func NewEventRestImpact() *EventRestImpact`

NewEventRestImpact instantiates a new EventRestImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRestImpactWithDefaults

`func NewEventRestImpactWithDefaults() *EventRestImpact`

NewEventRestImpactWithDefaults instantiates a new EventRestImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EventRestImpact) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EventRestImpact) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EventRestImpact) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EventRestImpact) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *EventRestImpact) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EventRestImpact) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EventRestImpact) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *EventRestImpact) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *EventRestImpact) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *EventRestImpact) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *EventRestImpact) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *EventRestImpact) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetImpactLevel

`func (o *EventRestImpact) GetImpactLevel() string`

GetImpactLevel returns the ImpactLevel field if non-nil, zero value otherwise.

### GetImpactLevelOk

`func (o *EventRestImpact) GetImpactLevelOk() (*string, bool)`

GetImpactLevelOk returns a tuple with the ImpactLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactLevel

`func (o *EventRestImpact) SetImpactLevel(v string)`

SetImpactLevel sets ImpactLevel field to given value.

### HasImpactLevel

`func (o *EventRestImpact) HasImpactLevel() bool`

HasImpactLevel returns a boolean if a field has been set.

### GetEventType

`func (o *EventRestImpact) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventRestImpact) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventRestImpact) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventRestImpact) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetResourceId

`func (o *EventRestImpact) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventRestImpact) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventRestImpact) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *EventRestImpact) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *EventRestImpact) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *EventRestImpact) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *EventRestImpact) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *EventRestImpact) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


