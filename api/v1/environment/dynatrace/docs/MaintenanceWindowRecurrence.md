# MaintenanceWindowRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **string** | The day of the week for weekly maintenance.    The format is the full name of the day in upper case, for example &#x60;WEDNESDAY&#x60;. | [optional] 
**DayOfMonth** | Pointer to **int32** | The day of the month for monthly maintenance. | [optional] 
**Start** | **string** | The start time of the maintenance window. The format is &#x60;HH:mm&#x60;. | 
**Duration** | **int32** | The duration of the maintenance window in minutes. | 

## Methods

### NewMaintenanceWindowRecurrence

`func NewMaintenanceWindowRecurrence(start string, duration int32, ) *MaintenanceWindowRecurrence`

NewMaintenanceWindowRecurrence instantiates a new MaintenanceWindowRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowRecurrenceWithDefaults

`func NewMaintenanceWindowRecurrenceWithDefaults() *MaintenanceWindowRecurrence`

NewMaintenanceWindowRecurrenceWithDefaults instantiates a new MaintenanceWindowRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *MaintenanceWindowRecurrence) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MaintenanceWindowRecurrence) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MaintenanceWindowRecurrence) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *MaintenanceWindowRecurrence) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *MaintenanceWindowRecurrence) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MaintenanceWindowRecurrence) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MaintenanceWindowRecurrence) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MaintenanceWindowRecurrence) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetStart

`func (o *MaintenanceWindowRecurrence) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MaintenanceWindowRecurrence) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MaintenanceWindowRecurrence) SetStart(v string)`

SetStart sets Start field to given value.


### GetDuration

`func (o *MaintenanceWindowRecurrence) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MaintenanceWindowRecurrence) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MaintenanceWindowRecurrence) SetDuration(v int32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


