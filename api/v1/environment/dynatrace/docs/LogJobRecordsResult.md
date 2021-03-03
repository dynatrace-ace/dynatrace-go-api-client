# LogJobRecordsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]LogRecord**](LogRecord.md) | The list of log analysis results.    The last page contains empty list. | [optional] 
**ScrollToken** | Pointer to **string** | The *scroll token* for the next page of results.    Without it you&#39;ll get the first page again. | [optional] 

## Methods

### NewLogJobRecordsResult

`func NewLogJobRecordsResult() *LogJobRecordsResult`

NewLogJobRecordsResult instantiates a new LogJobRecordsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogJobRecordsResultWithDefaults

`func NewLogJobRecordsResultWithDefaults() *LogJobRecordsResult`

NewLogJobRecordsResultWithDefaults instantiates a new LogJobRecordsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *LogJobRecordsResult) GetRecords() []LogRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *LogJobRecordsResult) GetRecordsOk() (*[]LogRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *LogJobRecordsResult) SetRecords(v []LogRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *LogJobRecordsResult) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetScrollToken

`func (o *LogJobRecordsResult) GetScrollToken() string`

GetScrollToken returns the ScrollToken field if non-nil, zero value otherwise.

### GetScrollTokenOk

`func (o *LogJobRecordsResult) GetScrollTokenOk() (*string, bool)`

GetScrollTokenOk returns a tuple with the ScrollToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollToken

`func (o *LogJobRecordsResult) SetScrollToken(v string)`

SetScrollToken sets ScrollToken field to given value.

### HasScrollToken

`func (o *LogJobRecordsResult) HasScrollToken() bool`

HasScrollToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


