# LogJobStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogAnalysisStatus** | Pointer to **string** | The status of the log analysis job. | [optional] 
**StatusChangeTimestamp** | Pointer to **int64** | The timestamp of the last status change, in UTC milliseconds. | [optional] 
**LogHandlingError** | Pointer to **string** | The cause of the job failure.    A successful job has the &#x60;NONE&#x60; value. | [optional] 
**RecordsTotal** | Pointer to **int32** | The number of analyzed log entries. | [optional] 
**SortableFields** | Pointer to **[]string** | The map of the log entry sortable fields. | [optional] 
**FilterableFields** | Pointer to **[]string** | The map of the log entry filterable fields. | [optional] 

## Methods

### NewLogJobStatusResult

`func NewLogJobStatusResult() *LogJobStatusResult`

NewLogJobStatusResult instantiates a new LogJobStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogJobStatusResultWithDefaults

`func NewLogJobStatusResultWithDefaults() *LogJobStatusResult`

NewLogJobStatusResultWithDefaults instantiates a new LogJobStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogAnalysisStatus

`func (o *LogJobStatusResult) GetLogAnalysisStatus() string`

GetLogAnalysisStatus returns the LogAnalysisStatus field if non-nil, zero value otherwise.

### GetLogAnalysisStatusOk

`func (o *LogJobStatusResult) GetLogAnalysisStatusOk() (*string, bool)`

GetLogAnalysisStatusOk returns a tuple with the LogAnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalysisStatus

`func (o *LogJobStatusResult) SetLogAnalysisStatus(v string)`

SetLogAnalysisStatus sets LogAnalysisStatus field to given value.

### HasLogAnalysisStatus

`func (o *LogJobStatusResult) HasLogAnalysisStatus() bool`

HasLogAnalysisStatus returns a boolean if a field has been set.

### GetStatusChangeTimestamp

`func (o *LogJobStatusResult) GetStatusChangeTimestamp() int64`

GetStatusChangeTimestamp returns the StatusChangeTimestamp field if non-nil, zero value otherwise.

### GetStatusChangeTimestampOk

`func (o *LogJobStatusResult) GetStatusChangeTimestampOk() (*int64, bool)`

GetStatusChangeTimestampOk returns a tuple with the StatusChangeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeTimestamp

`func (o *LogJobStatusResult) SetStatusChangeTimestamp(v int64)`

SetStatusChangeTimestamp sets StatusChangeTimestamp field to given value.

### HasStatusChangeTimestamp

`func (o *LogJobStatusResult) HasStatusChangeTimestamp() bool`

HasStatusChangeTimestamp returns a boolean if a field has been set.

### GetLogHandlingError

`func (o *LogJobStatusResult) GetLogHandlingError() string`

GetLogHandlingError returns the LogHandlingError field if non-nil, zero value otherwise.

### GetLogHandlingErrorOk

`func (o *LogJobStatusResult) GetLogHandlingErrorOk() (*string, bool)`

GetLogHandlingErrorOk returns a tuple with the LogHandlingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHandlingError

`func (o *LogJobStatusResult) SetLogHandlingError(v string)`

SetLogHandlingError sets LogHandlingError field to given value.

### HasLogHandlingError

`func (o *LogJobStatusResult) HasLogHandlingError() bool`

HasLogHandlingError returns a boolean if a field has been set.

### GetRecordsTotal

`func (o *LogJobStatusResult) GetRecordsTotal() int32`

GetRecordsTotal returns the RecordsTotal field if non-nil, zero value otherwise.

### GetRecordsTotalOk

`func (o *LogJobStatusResult) GetRecordsTotalOk() (*int32, bool)`

GetRecordsTotalOk returns a tuple with the RecordsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsTotal

`func (o *LogJobStatusResult) SetRecordsTotal(v int32)`

SetRecordsTotal sets RecordsTotal field to given value.

### HasRecordsTotal

`func (o *LogJobStatusResult) HasRecordsTotal() bool`

HasRecordsTotal returns a boolean if a field has been set.

### GetSortableFields

`func (o *LogJobStatusResult) GetSortableFields() []string`

GetSortableFields returns the SortableFields field if non-nil, zero value otherwise.

### GetSortableFieldsOk

`func (o *LogJobStatusResult) GetSortableFieldsOk() (*[]string, bool)`

GetSortableFieldsOk returns a tuple with the SortableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortableFields

`func (o *LogJobStatusResult) SetSortableFields(v []string)`

SetSortableFields sets SortableFields field to given value.

### HasSortableFields

`func (o *LogJobStatusResult) HasSortableFields() bool`

HasSortableFields returns a boolean if a field has been set.

### GetFilterableFields

`func (o *LogJobStatusResult) GetFilterableFields() []string`

GetFilterableFields returns the FilterableFields field if non-nil, zero value otherwise.

### GetFilterableFieldsOk

`func (o *LogJobStatusResult) GetFilterableFieldsOk() (*[]string, bool)`

GetFilterableFieldsOk returns a tuple with the FilterableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterableFields

`func (o *LogJobStatusResult) SetFilterableFields(v []string)`

SetFilterableFields sets FilterableFields field to given value.

### HasFilterableFields

`func (o *LogJobStatusResult) HasFilterableFields() bool`

HasFilterableFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


