# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurrenceType** | **string** | The type of the schedule recurrence. | 
**Recurrence** | Pointer to [**Recurrence**](Recurrence.md) |  | [optional] 
**Start** | **string** | The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format. | 
**End** | **string** | The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format. | 
**ZoneId** | **string** | The time zone of the start and end time. Default time zone is UTC.   You can use either UTC offset &#x60;UTC+01:00&#x60; format or the IANA Time Zone Database format (for example, &#x60;Europe/Vienna&#x60;). | 

## Methods

### NewSchedule

`func NewSchedule(recurrenceType string, start string, end string, zoneId string, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurrenceType

`func (o *Schedule) GetRecurrenceType() string`

GetRecurrenceType returns the RecurrenceType field if non-nil, zero value otherwise.

### GetRecurrenceTypeOk

`func (o *Schedule) GetRecurrenceTypeOk() (*string, bool)`

GetRecurrenceTypeOk returns a tuple with the RecurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceType

`func (o *Schedule) SetRecurrenceType(v string)`

SetRecurrenceType sets RecurrenceType field to given value.


### GetRecurrence

`func (o *Schedule) GetRecurrence() Recurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *Schedule) GetRecurrenceOk() (*Recurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *Schedule) SetRecurrence(v Recurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *Schedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetStart

`func (o *Schedule) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Schedule) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Schedule) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Schedule) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Schedule) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Schedule) SetEnd(v string)`

SetEnd sets End field to given value.


### GetZoneId

`func (o *Schedule) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *Schedule) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *Schedule) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


