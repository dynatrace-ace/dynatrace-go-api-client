# LogList4hostResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentAccess** | Pointer to **bool** | The access to the log content is granted (&#x60;true&#x60;) or denied (&#x60;false&#x60;). | [optional] 
**Logs** | Pointer to [**[]Log4host**](Log4host.md) | The list of available OS logs. | [optional] 

## Methods

### NewLogList4hostResult

`func NewLogList4hostResult() *LogList4hostResult`

NewLogList4hostResult instantiates a new LogList4hostResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogList4hostResultWithDefaults

`func NewLogList4hostResultWithDefaults() *LogList4hostResult`

NewLogList4hostResultWithDefaults instantiates a new LogList4hostResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentAccess

`func (o *LogList4hostResult) GetContentAccess() bool`

GetContentAccess returns the ContentAccess field if non-nil, zero value otherwise.

### GetContentAccessOk

`func (o *LogList4hostResult) GetContentAccessOk() (*bool, bool)`

GetContentAccessOk returns a tuple with the ContentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccess

`func (o *LogList4hostResult) SetContentAccess(v bool)`

SetContentAccess sets ContentAccess field to given value.

### HasContentAccess

`func (o *LogList4hostResult) HasContentAccess() bool`

HasContentAccess returns a boolean if a field has been set.

### GetLogs

`func (o *LogList4hostResult) GetLogs() []Log4host`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogList4hostResult) GetLogsOk() (*[]Log4host, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogList4hostResult) SetLogs(v []Log4host)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *LogList4hostResult) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


