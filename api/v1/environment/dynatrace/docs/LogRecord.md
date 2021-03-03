# LogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | The timestamp of the log entry, in UTC milliseconds. | [optional] 
**LogLevel** | Pointer to **string** | The severity level of the log entry. | [optional] 
**HostId** | Pointer to **string** | The entity ID of the host that produced the log.    Not applicable to OS logs. | [optional] 
**Text** | Pointer to **string** | The text of the log entry. | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** | The map of the log entry custom fields. | [optional] 

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

### GetTimestamp

`func (o *LogRecord) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogRecord) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogRecord) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLogLevel

`func (o *LogRecord) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *LogRecord) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *LogRecord) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *LogRecord) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetHostId

`func (o *LogRecord) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *LogRecord) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *LogRecord) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *LogRecord) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetText

`func (o *LogRecord) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LogRecord) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LogRecord) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *LogRecord) HasText() bool`

HasText returns a boolean if a field has been set.

### GetCustomFields

`func (o *LogRecord) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *LogRecord) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *LogRecord) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *LogRecord) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


