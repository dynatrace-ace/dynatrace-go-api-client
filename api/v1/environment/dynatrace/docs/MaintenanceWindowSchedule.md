# MaintenanceWindowSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Recurrence of the schedule. | 
**TimezoneId** | Pointer to **string** | The time zone of the start and end time. Default time zone is UTC.    You can user either UTC offset &#x60;UTC+01:00&#x60; format or the IANA Time Zone Database format. | [optional] 
**MaintenanceStart** | **string** | The start date and time of the maintenance window in the &#x60;yyyy-MM-dd HH:mm&#x60; format. | 
**MaintenanceEnd** | **string** | The end date and time of the maintenance window in the &#x60;yyyy-MM-dd HH:mm&#x60; format. | 
**Recurrence** | Pointer to [**MaintenanceWindowRecurrence**](MaintenanceWindowRecurrence.md) |  | [optional] 

## Methods

### NewMaintenanceWindowSchedule

`func NewMaintenanceWindowSchedule(type_ string, maintenanceStart string, maintenanceEnd string, ) *MaintenanceWindowSchedule`

NewMaintenanceWindowSchedule instantiates a new MaintenanceWindowSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowScheduleWithDefaults

`func NewMaintenanceWindowScheduleWithDefaults() *MaintenanceWindowSchedule`

NewMaintenanceWindowScheduleWithDefaults instantiates a new MaintenanceWindowSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MaintenanceWindowSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaintenanceWindowSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaintenanceWindowSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetTimezoneId

`func (o *MaintenanceWindowSchedule) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *MaintenanceWindowSchedule) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *MaintenanceWindowSchedule) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *MaintenanceWindowSchedule) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetMaintenanceStart

`func (o *MaintenanceWindowSchedule) GetMaintenanceStart() string`

GetMaintenanceStart returns the MaintenanceStart field if non-nil, zero value otherwise.

### GetMaintenanceStartOk

`func (o *MaintenanceWindowSchedule) GetMaintenanceStartOk() (*string, bool)`

GetMaintenanceStartOk returns a tuple with the MaintenanceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceStart

`func (o *MaintenanceWindowSchedule) SetMaintenanceStart(v string)`

SetMaintenanceStart sets MaintenanceStart field to given value.


### GetMaintenanceEnd

`func (o *MaintenanceWindowSchedule) GetMaintenanceEnd() string`

GetMaintenanceEnd returns the MaintenanceEnd field if non-nil, zero value otherwise.

### GetMaintenanceEndOk

`func (o *MaintenanceWindowSchedule) GetMaintenanceEndOk() (*string, bool)`

GetMaintenanceEndOk returns a tuple with the MaintenanceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEnd

`func (o *MaintenanceWindowSchedule) SetMaintenanceEnd(v string)`

SetMaintenanceEnd sets MaintenanceEnd field to given value.


### GetRecurrence

`func (o *MaintenanceWindowSchedule) GetRecurrence() MaintenanceWindowRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MaintenanceWindowSchedule) GetRecurrenceOk() (*MaintenanceWindowRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MaintenanceWindowSchedule) SetRecurrence(v MaintenanceWindowRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MaintenanceWindowSchedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


