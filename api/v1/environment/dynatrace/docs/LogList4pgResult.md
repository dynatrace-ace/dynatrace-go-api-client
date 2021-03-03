# LogList4pgResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]LogFile4pg**](LogFile4pg.md) | The list of available process group logs. | [optional] 

## Methods

### NewLogList4pgResult

`func NewLogList4pgResult() *LogList4pgResult`

NewLogList4pgResult instantiates a new LogList4pgResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogList4pgResultWithDefaults

`func NewLogList4pgResultWithDefaults() *LogList4pgResult`

NewLogList4pgResultWithDefaults instantiates a new LogList4pgResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LogList4pgResult) GetLogs() []LogFile4pg`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogList4pgResult) GetLogsOk() (*[]LogFile4pg, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogList4pgResult) SetLogs(v []LogFile4pg)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *LogList4pgResult) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


