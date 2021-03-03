# LogJobRecordsTopValuesRestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParsingFieldTopValues** | Pointer to [**[]ParsingFieldTopValue**](ParsingFieldTopValue.md) | Log analysis parsing result top values | [optional] 
**ValuesCount** | Pointer to **int32** | Log analysis parsing result top values count | [optional] 

## Methods

### NewLogJobRecordsTopValuesRestResult

`func NewLogJobRecordsTopValuesRestResult() *LogJobRecordsTopValuesRestResult`

NewLogJobRecordsTopValuesRestResult instantiates a new LogJobRecordsTopValuesRestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogJobRecordsTopValuesRestResultWithDefaults

`func NewLogJobRecordsTopValuesRestResultWithDefaults() *LogJobRecordsTopValuesRestResult`

NewLogJobRecordsTopValuesRestResultWithDefaults instantiates a new LogJobRecordsTopValuesRestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParsingFieldTopValues

`func (o *LogJobRecordsTopValuesRestResult) GetParsingFieldTopValues() []ParsingFieldTopValue`

GetParsingFieldTopValues returns the ParsingFieldTopValues field if non-nil, zero value otherwise.

### GetParsingFieldTopValuesOk

`func (o *LogJobRecordsTopValuesRestResult) GetParsingFieldTopValuesOk() (*[]ParsingFieldTopValue, bool)`

GetParsingFieldTopValuesOk returns a tuple with the ParsingFieldTopValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingFieldTopValues

`func (o *LogJobRecordsTopValuesRestResult) SetParsingFieldTopValues(v []ParsingFieldTopValue)`

SetParsingFieldTopValues sets ParsingFieldTopValues field to given value.

### HasParsingFieldTopValues

`func (o *LogJobRecordsTopValuesRestResult) HasParsingFieldTopValues() bool`

HasParsingFieldTopValues returns a boolean if a field has been set.

### GetValuesCount

`func (o *LogJobRecordsTopValuesRestResult) GetValuesCount() int32`

GetValuesCount returns the ValuesCount field if non-nil, zero value otherwise.

### GetValuesCountOk

`func (o *LogJobRecordsTopValuesRestResult) GetValuesCountOk() (*int32, bool)`

GetValuesCountOk returns a tuple with the ValuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesCount

`func (o *LogJobRecordsTopValuesRestResult) SetValuesCount(v int32)`

SetValuesCount sets ValuesCount field to given value.

### HasValuesCount

`func (o *LogJobRecordsTopValuesRestResult) HasValuesCount() bool`

HasValuesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


