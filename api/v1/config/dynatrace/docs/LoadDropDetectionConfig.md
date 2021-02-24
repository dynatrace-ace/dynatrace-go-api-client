# LoadDropDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**LoadDropPercent** | Pointer to **int32** | Alert if the observed load is less than *X* % of the expected value. | [optional] 
**MinAbnormalStateDurationInMinutes** | Pointer to **int32** | Alert if the service stays in abnormal state for at least *X* minutes. | [optional] 

## Methods

### NewLoadDropDetectionConfig

`func NewLoadDropDetectionConfig(enabled bool, ) *LoadDropDetectionConfig`

NewLoadDropDetectionConfig instantiates a new LoadDropDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadDropDetectionConfigWithDefaults

`func NewLoadDropDetectionConfigWithDefaults() *LoadDropDetectionConfig`

NewLoadDropDetectionConfigWithDefaults instantiates a new LoadDropDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LoadDropDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoadDropDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoadDropDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLoadDropPercent

`func (o *LoadDropDetectionConfig) GetLoadDropPercent() int32`

GetLoadDropPercent returns the LoadDropPercent field if non-nil, zero value otherwise.

### GetLoadDropPercentOk

`func (o *LoadDropDetectionConfig) GetLoadDropPercentOk() (*int32, bool)`

GetLoadDropPercentOk returns a tuple with the LoadDropPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadDropPercent

`func (o *LoadDropDetectionConfig) SetLoadDropPercent(v int32)`

SetLoadDropPercent sets LoadDropPercent field to given value.

### HasLoadDropPercent

`func (o *LoadDropDetectionConfig) HasLoadDropPercent() bool`

HasLoadDropPercent returns a boolean if a field has been set.

### GetMinAbnormalStateDurationInMinutes

`func (o *LoadDropDetectionConfig) GetMinAbnormalStateDurationInMinutes() int32`

GetMinAbnormalStateDurationInMinutes returns the MinAbnormalStateDurationInMinutes field if non-nil, zero value otherwise.

### GetMinAbnormalStateDurationInMinutesOk

`func (o *LoadDropDetectionConfig) GetMinAbnormalStateDurationInMinutesOk() (*int32, bool)`

GetMinAbnormalStateDurationInMinutesOk returns a tuple with the MinAbnormalStateDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAbnormalStateDurationInMinutes

`func (o *LoadDropDetectionConfig) SetMinAbnormalStateDurationInMinutes(v int32)`

SetMinAbnormalStateDurationInMinutes sets MinAbnormalStateDurationInMinutes field to given value.

### HasMinAbnormalStateDurationInMinutes

`func (o *LoadDropDetectionConfig) HasMinAbnormalStateDurationInMinutes() bool`

HasMinAbnormalStateDurationInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


