# FailureRateIncreaseDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMode** | **string** | How to detect failure rate increase: automatically, or based on fixed thresholds, or do not detect. | 
**AutomaticDetection** | Pointer to [**FailureRateIncreaseAutodetectionConfig**](FailureRateIncreaseAutodetectionConfig.md) |  | [optional] 
**Thresholds** | Pointer to [**FailureRateIncreaseThresholdConfig**](FailureRateIncreaseThresholdConfig.md) |  | [optional] 

## Methods

### NewFailureRateIncreaseDetectionConfig

`func NewFailureRateIncreaseDetectionConfig(detectionMode string, ) *FailureRateIncreaseDetectionConfig`

NewFailureRateIncreaseDetectionConfig instantiates a new FailureRateIncreaseDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureRateIncreaseDetectionConfigWithDefaults

`func NewFailureRateIncreaseDetectionConfigWithDefaults() *FailureRateIncreaseDetectionConfig`

NewFailureRateIncreaseDetectionConfigWithDefaults instantiates a new FailureRateIncreaseDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionMode

`func (o *FailureRateIncreaseDetectionConfig) GetDetectionMode() string`

GetDetectionMode returns the DetectionMode field if non-nil, zero value otherwise.

### GetDetectionModeOk

`func (o *FailureRateIncreaseDetectionConfig) GetDetectionModeOk() (*string, bool)`

GetDetectionModeOk returns a tuple with the DetectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMode

`func (o *FailureRateIncreaseDetectionConfig) SetDetectionMode(v string)`

SetDetectionMode sets DetectionMode field to given value.


### GetAutomaticDetection

`func (o *FailureRateIncreaseDetectionConfig) GetAutomaticDetection() FailureRateIncreaseAutodetectionConfig`

GetAutomaticDetection returns the AutomaticDetection field if non-nil, zero value otherwise.

### GetAutomaticDetectionOk

`func (o *FailureRateIncreaseDetectionConfig) GetAutomaticDetectionOk() (*FailureRateIncreaseAutodetectionConfig, bool)`

GetAutomaticDetectionOk returns a tuple with the AutomaticDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDetection

`func (o *FailureRateIncreaseDetectionConfig) SetAutomaticDetection(v FailureRateIncreaseAutodetectionConfig)`

SetAutomaticDetection sets AutomaticDetection field to given value.

### HasAutomaticDetection

`func (o *FailureRateIncreaseDetectionConfig) HasAutomaticDetection() bool`

HasAutomaticDetection returns a boolean if a field has been set.

### GetThresholds

`func (o *FailureRateIncreaseDetectionConfig) GetThresholds() FailureRateIncreaseThresholdConfig`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *FailureRateIncreaseDetectionConfig) GetThresholdsOk() (*FailureRateIncreaseThresholdConfig, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *FailureRateIncreaseDetectionConfig) SetThresholds(v FailureRateIncreaseThresholdConfig)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *FailureRateIncreaseDetectionConfig) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


