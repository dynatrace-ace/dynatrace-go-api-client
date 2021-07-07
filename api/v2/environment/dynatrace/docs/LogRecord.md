# LogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Type of event | [optional] 
**AdditionalColumns** | Pointer to **map[string]string** | Additional columns of the log record. | [optional] 
**Timestamp** | Pointer to **time.Time** | The timestamp of the log record. | [optional] 
**Content** | Pointer to **string** | The content of the log record. | [optional] 
**Status** | Pointer to **string** | The log status (based on the log level). | [optional] 

## Methods

### NewLogRecord

`func NewLogRecord() *LogRecord`

NewLogRecord instantiates a new LogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogRecordWithDefaults

`func NewLogRecordWithDefaults() *LogRecord`

NewLogRecordWithDefaults instantiates a new LogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *LogRecord) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LogRecord) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LogRecord) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *LogRecord) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAdditionalColumns

`func (o *LogRecord) GetAdditionalColumns() map[string]string`

GetAdditionalColumns returns the AdditionalColumns field if non-nil, zero value otherwise.

### GetAdditionalColumnsOk

`func (o *LogRecord) GetAdditionalColumnsOk() (*map[string]string, bool)`

GetAdditionalColumnsOk returns a tuple with the AdditionalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalColumns

`func (o *LogRecord) SetAdditionalColumns(v map[string]string)`

SetAdditionalColumns sets AdditionalColumns field to given value.

### HasAdditionalColumns

`func (o *LogRecord) HasAdditionalColumns() bool`

HasAdditionalColumns returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetContent

`func (o *LogRecord) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *LogRecord) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *LogRecord) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *LogRecord) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatus

`func (o *LogRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


