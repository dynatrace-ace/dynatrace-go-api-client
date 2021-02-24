# HighGcActivityDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**CustomThresholds** | Pointer to [**HighGcActivityThresholds**](HighGcActivityThresholds.md) |  | [optional] 

## Methods

### NewHighGcActivityDetectionConfig

`func NewHighGcActivityDetectionConfig(enabled bool, ) *HighGcActivityDetectionConfig`

NewHighGcActivityDetectionConfig instantiates a new HighGcActivityDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighGcActivityDetectionConfigWithDefaults

`func NewHighGcActivityDetectionConfigWithDefaults() *HighGcActivityDetectionConfig`

NewHighGcActivityDetectionConfigWithDefaults instantiates a new HighGcActivityDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HighGcActivityDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HighGcActivityDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HighGcActivityDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomThresholds

`func (o *HighGcActivityDetectionConfig) GetCustomThresholds() HighGcActivityThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *HighGcActivityDetectionConfig) GetCustomThresholdsOk() (*HighGcActivityThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *HighGcActivityDetectionConfig) SetCustomThresholds(v HighGcActivityThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *HighGcActivityDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


