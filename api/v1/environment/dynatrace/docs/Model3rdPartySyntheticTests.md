# Model3rdPartySyntheticTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyntheticEngineName** | **string** | The type of the third-party synthetic monitor. | 
**SyntheticEngineIconUrl** | Pointer to **string** | The URL of the third-party synthetic monitor icon. | [optional] 
**MessageTimestamp** | **int64** | The timestamp of the message creation, in UTC milliseconds. | 
**Locations** | [**[]Model3rdPartySyntheticLocation**](Model3rdPartySyntheticLocation.md) | The list of third-party synthetic locations. | 
**Tests** | [**[]Model3rdPartySyntheticMonitor**](Model3rdPartySyntheticMonitor.md) | The list of third-party synthetic monitors. | 
**TestResults** | Pointer to [**[]Model3rdPartySyntheticTestResult**](Model3rdPartySyntheticTestResult.md) | The list of results of third-party synthetic monitor execution. | [optional] 

## Methods

### NewModel3rdPartySyntheticTests

`func NewModel3rdPartySyntheticTests(syntheticEngineName string, messageTimestamp int64, locations []Model3rdPartySyntheticLocation, tests []Model3rdPartySyntheticMonitor, ) *Model3rdPartySyntheticTests`

NewModel3rdPartySyntheticTests instantiates a new Model3rdPartySyntheticTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel3rdPartySyntheticTestsWithDefaults

`func NewModel3rdPartySyntheticTestsWithDefaults() *Model3rdPartySyntheticTests`

NewModel3rdPartySyntheticTestsWithDefaults instantiates a new Model3rdPartySyntheticTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyntheticEngineName

`func (o *Model3rdPartySyntheticTests) GetSyntheticEngineName() string`

GetSyntheticEngineName returns the SyntheticEngineName field if non-nil, zero value otherwise.

### GetSyntheticEngineNameOk

`func (o *Model3rdPartySyntheticTests) GetSyntheticEngineNameOk() (*string, bool)`

GetSyntheticEngineNameOk returns a tuple with the SyntheticEngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEngineName

`func (o *Model3rdPartySyntheticTests) SetSyntheticEngineName(v string)`

SetSyntheticEngineName sets SyntheticEngineName field to given value.


### GetSyntheticEngineIconUrl

`func (o *Model3rdPartySyntheticTests) GetSyntheticEngineIconUrl() string`

GetSyntheticEngineIconUrl returns the SyntheticEngineIconUrl field if non-nil, zero value otherwise.

### GetSyntheticEngineIconUrlOk

`func (o *Model3rdPartySyntheticTests) GetSyntheticEngineIconUrlOk() (*string, bool)`

GetSyntheticEngineIconUrlOk returns a tuple with the SyntheticEngineIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEngineIconUrl

`func (o *Model3rdPartySyntheticTests) SetSyntheticEngineIconUrl(v string)`

SetSyntheticEngineIconUrl sets SyntheticEngineIconUrl field to given value.

### HasSyntheticEngineIconUrl

`func (o *Model3rdPartySyntheticTests) HasSyntheticEngineIconUrl() bool`

HasSyntheticEngineIconUrl returns a boolean if a field has been set.

### GetMessageTimestamp

`func (o *Model3rdPartySyntheticTests) GetMessageTimestamp() int64`

GetMessageTimestamp returns the MessageTimestamp field if non-nil, zero value otherwise.

### GetMessageTimestampOk

`func (o *Model3rdPartySyntheticTests) GetMessageTimestampOk() (*int64, bool)`

GetMessageTimestampOk returns a tuple with the MessageTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTimestamp

`func (o *Model3rdPartySyntheticTests) SetMessageTimestamp(v int64)`

SetMessageTimestamp sets MessageTimestamp field to given value.


### GetLocations

`func (o *Model3rdPartySyntheticTests) GetLocations() []Model3rdPartySyntheticLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Model3rdPartySyntheticTests) GetLocationsOk() (*[]Model3rdPartySyntheticLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Model3rdPartySyntheticTests) SetLocations(v []Model3rdPartySyntheticLocation)`

SetLocations sets Locations field to given value.


### GetTests

`func (o *Model3rdPartySyntheticTests) GetTests() []Model3rdPartySyntheticMonitor`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *Model3rdPartySyntheticTests) GetTestsOk() (*[]Model3rdPartySyntheticMonitor, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *Model3rdPartySyntheticTests) SetTests(v []Model3rdPartySyntheticMonitor)`

SetTests sets Tests field to given value.


### GetTestResults

`func (o *Model3rdPartySyntheticTests) GetTestResults() []Model3rdPartySyntheticTestResult`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *Model3rdPartySyntheticTests) GetTestResultsOk() (*[]Model3rdPartySyntheticTestResult, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *Model3rdPartySyntheticTests) SetTestResults(v []Model3rdPartySyntheticTestResult)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *Model3rdPartySyntheticTests) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


