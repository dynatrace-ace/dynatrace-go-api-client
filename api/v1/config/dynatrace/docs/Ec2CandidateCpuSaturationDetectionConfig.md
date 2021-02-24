# Ec2CandidateCpuSaturationDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**Ec2CandidateCpuSaturationThresholds**](Ec2CandidateCpuSaturationThresholds.md) |  | [optional] 

## Methods

### NewEc2CandidateCpuSaturationDetectionConfig

`func NewEc2CandidateCpuSaturationDetectionConfig(enabled bool, ) *Ec2CandidateCpuSaturationDetectionConfig`

NewEc2CandidateCpuSaturationDetectionConfig instantiates a new Ec2CandidateCpuSaturationDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2CandidateCpuSaturationDetectionConfigWithDefaults

`func NewEc2CandidateCpuSaturationDetectionConfigWithDefaults() *Ec2CandidateCpuSaturationDetectionConfig`

NewEc2CandidateCpuSaturationDetectionConfigWithDefaults instantiates a new Ec2CandidateCpuSaturationDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Ec2CandidateCpuSaturationDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Ec2CandidateCpuSaturationDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Ec2CandidateCpuSaturationDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *Ec2CandidateCpuSaturationDetectionConfig) GetCustomThresholds() Ec2CandidateCpuSaturationThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *Ec2CandidateCpuSaturationDetectionConfig) GetCustomThresholdsOk() (*Ec2CandidateCpuSaturationThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *Ec2CandidateCpuSaturationDetectionConfig) SetCustomThresholds(v Ec2CandidateCpuSaturationThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *Ec2CandidateCpuSaturationDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


