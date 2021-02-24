# Recurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **string** | The day of the week for weekly maintenance.   The format is the full name of the day in upper case, for example &#x60;THURSDAY&#x60;. | [optional] 
**DayOfMonth** | Pointer to **int32** | The day of the month for monthly maintenance.   The value of &#x60;31&#x60; is treated as the last day of the month for months that don&#39;t have a 31st day. The value of &#x60;30&#x60; is also treated as the last day of the month for February. | [optional] 
**StartTime** | **string** | The start time of the maintenance window in HH:mm format. | 
**DurationMinutes** | **int32** | The duration of the maintenance window in minutes. | 

## Methods

### NewRecurrence

`func NewRecurrence(startTime string, durationMinutes int32, ) *Recurrence`

NewRecurrence instantiates a new Recurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceWithDefaults

`func NewRecurrenceWithDefaults() *Recurrence`

NewRecurrenceWithDefaults instantiates a new Recurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *Recurrence) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *Recurrence) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *Recurrence) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *Recurrence) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *Recurrence) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *Recurrence) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *Recurrence) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *Recurrence) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetStartTime

`func (o *Recurrence) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Recurrence) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Recurrence) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetDurationMinutes

`func (o *Recurrence) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *Recurrence) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *Recurrence) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


