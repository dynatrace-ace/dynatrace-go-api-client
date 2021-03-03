# EventSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** | The metric used in the event severity calculation. | [optional] 
**Value** | Pointer to **float64** | The value of the severity. | [optional] 
**Unit** | Pointer to **string** | The unit of the severity value. | [optional] 

## Methods

### NewEventSeverity

`func NewEventSeverity() *EventSeverity`

NewEventSeverity instantiates a new EventSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSeverityWithDefaults

`func NewEventSeverityWithDefaults() *EventSeverity`

NewEventSeverityWithDefaults instantiates a new EventSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *EventSeverity) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EventSeverity) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EventSeverity) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *EventSeverity) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetValue

`func (o *EventSeverity) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventSeverity) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventSeverity) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventSeverity) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *EventSeverity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *EventSeverity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *EventSeverity) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *EventSeverity) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


