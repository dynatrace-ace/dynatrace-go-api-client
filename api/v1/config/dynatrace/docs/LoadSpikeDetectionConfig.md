# LoadSpikeDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**LoadSpikePercent** | Pointer to **int32** | Alert if the observed load is more than *X* % of the expected value. | [optional] 
**MinAbnormalStateDurationInMinutes** | Pointer to **int32** | Alert if the service stays in abnormal state for at least *X* minutes. | [optional] 

## Methods

### NewLoadSpikeDetectionConfig

`func NewLoadSpikeDetectionConfig(enabled bool, ) *LoadSpikeDetectionConfig`

NewLoadSpikeDetectionConfig instantiates a new LoadSpikeDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadSpikeDetectionConfigWithDefaults

`func NewLoadSpikeDetectionConfigWithDefaults() *LoadSpikeDetectionConfig`

NewLoadSpikeDetectionConfigWithDefaults instantiates a new LoadSpikeDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LoadSpikeDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoadSpikeDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoadSpikeDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoadSpikePercent

`func (o *LoadSpikeDetectionConfig) GetLoadSpikePercent() int32`

GetLoadSpikePercent returns the LoadSpikePercent field if non-nil, zero value otherwise.

### GetLoadSpikePercentOk

`func (o *LoadSpikeDetectionConfig) GetLoadSpikePercentOk() (*int32, bool)`

GetLoadSpikePercentOk returns a tuple with the LoadSpikePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadSpikePercent

`func (o *LoadSpikeDetectionConfig) SetLoadSpikePercent(v int32)`

SetLoadSpikePercent sets LoadSpikePercent field to given value.

### HasLoadSpikePercent

`func (o *LoadSpikeDetectionConfig) HasLoadSpikePercent() bool`

HasLoadSpikePercent returns a boolean if a field has been set.

### GetMinAbnormalStateDurationInMinutes

`func (o *LoadSpikeDetectionConfig) GetMinAbnormalStateDurationInMinutes() int32`

GetMinAbnormalStateDurationInMinutes returns the MinAbnormalStateDurationInMinutes field if non-nil, zero value otherwise.

### GetMinAbnormalStateDurationInMinutesOk

`func (o *LoadSpikeDetectionConfig) GetMinAbnormalStateDurationInMinutesOk() (*int32, bool)`

GetMinAbnormalStateDurationInMinutesOk returns a tuple with the MinAbnormalStateDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAbnormalStateDurationInMinutes

`func (o *LoadSpikeDetectionConfig) SetMinAbnormalStateDurationInMinutes(v int32)`

SetMinAbnormalStateDurationInMinutes sets MinAbnormalStateDurationInMinutes field to given value.

### HasMinAbnormalStateDurationInMinutes

`func (o *LoadSpikeDetectionConfig) HasMinAbnormalStateDurationInMinutes() bool`

HasMinAbnormalStateDurationInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


