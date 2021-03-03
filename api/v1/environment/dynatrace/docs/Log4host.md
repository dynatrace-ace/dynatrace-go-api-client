# Log4host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The full path to the log. | [optional] 
**Size** | Pointer to **int64** | The size of the log, bytes. | [optional] 
**AvailableForAnalysis** | Pointer to **bool** | The log is available (&#x60;true&#x60;) or not available (&#x60;false&#x60;) for analysis. | [optional] 

## Methods

### NewLog4host

`func NewLog4host() *Log4host`

NewLog4host instantiates a new Log4host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLog4hostWithDefaults

`func NewLog4hostWithDefaults() *Log4host`

NewLog4hostWithDefaults instantiates a new Log4host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Log4host) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Log4host) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Log4host) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Log4host) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *Log4host) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Log4host) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Log4host) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Log4host) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAvailableForAnalysis

`func (o *Log4host) GetAvailableForAnalysis() bool`

GetAvailableForAnalysis returns the AvailableForAnalysis field if non-nil, zero value otherwise.

### GetAvailableForAnalysisOk

`func (o *Log4host) GetAvailableForAnalysisOk() (*bool, bool)`

GetAvailableForAnalysisOk returns a tuple with the AvailableForAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForAnalysis

`func (o *Log4host) SetAvailableForAnalysis(v bool)`

SetAvailableForAnalysis sets AvailableForAnalysis field to given value.

### HasAvailableForAnalysis

`func (o *Log4host) HasAvailableForAnalysis() bool`

HasAvailableForAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


