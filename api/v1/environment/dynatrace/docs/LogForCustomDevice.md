# LogForCustomDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The full path to the log. | [optional] 
**AvailableForAnalysis** | Pointer to **bool** | The log is available (&#x60;true&#x60;) or not available (&#x60;false&#x60;) for analysis. | [optional] 

## Methods

### NewLogForCustomDevice

`func NewLogForCustomDevice() *LogForCustomDevice`

NewLogForCustomDevice instantiates a new LogForCustomDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForCustomDeviceWithDefaults

`func NewLogForCustomDeviceWithDefaults() *LogForCustomDevice`

NewLogForCustomDeviceWithDefaults instantiates a new LogForCustomDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LogForCustomDevice) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogForCustomDevice) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogForCustomDevice) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogForCustomDevice) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetAvailableForAnalysis

`func (o *LogForCustomDevice) GetAvailableForAnalysis() bool`

GetAvailableForAnalysis returns the AvailableForAnalysis field if non-nil, zero value otherwise.

### GetAvailableForAnalysisOk

`func (o *LogForCustomDevice) GetAvailableForAnalysisOk() (*bool, bool)`

GetAvailableForAnalysisOk returns a tuple with the AvailableForAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForAnalysis

`func (o *LogForCustomDevice) SetAvailableForAnalysis(v bool)`

SetAvailableForAnalysis sets AvailableForAnalysis field to given value.

### HasAvailableForAnalysis

`func (o *LogForCustomDevice) HasAvailableForAnalysis() bool`

HasAvailableForAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


