# LogRecordsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]LogRecord**](LogRecord.md) | A list of retrieved log records. | [optional] 
**SliceSize** | Pointer to **int64** | The total number of records in a slice. | [optional] 
**NextSliceKey** | Pointer to **string** | The cursor for the next slice of log records. | [optional] 

## Methods

### NewLogRecordsList

`func NewLogRecordsList() *LogRecordsList`

NewLogRecordsList instantiates a new LogRecordsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogRecordsListWithDefaults

`func NewLogRecordsListWithDefaults() *LogRecordsList`

NewLogRecordsListWithDefaults instantiates a new LogRecordsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *LogRecordsList) GetResults() []LogRecord`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LogRecordsList) GetResultsOk() (*[]LogRecord, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LogRecordsList) SetResults(v []LogRecord)`

SetResults sets Results field to given value.

### HasResults

`func (o *LogRecordsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSliceSize

`func (o *LogRecordsList) GetSliceSize() int64`

GetSliceSize returns the SliceSize field if non-nil, zero value otherwise.

### GetSliceSizeOk

`func (o *LogRecordsList) GetSliceSizeOk() (*int64, bool)`

GetSliceSizeOk returns a tuple with the SliceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSize

`func (o *LogRecordsList) SetSliceSize(v int64)`

SetSliceSize sets SliceSize field to given value.

### HasSliceSize

`func (o *LogRecordsList) HasSliceSize() bool`

HasSliceSize returns a boolean if a field has been set.

### GetNextSliceKey

`func (o *LogRecordsList) GetNextSliceKey() string`

GetNextSliceKey returns the NextSliceKey field if non-nil, zero value otherwise.

### GetNextSliceKeyOk

`func (o *LogRecordsList) GetNextSliceKeyOk() (*string, bool)`

GetNextSliceKeyOk returns a tuple with the NextSliceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSliceKey

`func (o *LogRecordsList) SetNextSliceKey(v string)`

SetNextSliceKey sets NextSliceKey field to given value.

### HasNextSliceKey

`func (o *LogRecordsList) HasNextSliceKey() bool`

HasNextSliceKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


