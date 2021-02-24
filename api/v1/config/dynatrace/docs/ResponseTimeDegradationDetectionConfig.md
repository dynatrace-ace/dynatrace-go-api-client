# ResponseTimeDegradationDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMode** | **string** | How to detect response time degradation: automatically, or based on fixed thresholds, or do not detect. | 
**AutomaticDetection** | Pointer to [**ResponseTimeDegradationAutodetectionConfig**](ResponseTimeDegradationAutodetectionConfig.md) |  | [optional] 
**Thresholds** | Pointer to [**ResponseTimeDegradationThresholdConfig**](ResponseTimeDegradationThresholdConfig.md) |  | [optional] 

## Methods

### NewResponseTimeDegradationDetectionConfig

`func NewResponseTimeDegradationDetectionConfig(detectionMode string, ) *ResponseTimeDegradationDetectionConfig`

NewResponseTimeDegradationDetectionConfig instantiates a new ResponseTimeDegradationDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTimeDegradationDetectionConfigWithDefaults

`func NewResponseTimeDegradationDetectionConfigWithDefaults() *ResponseTimeDegradationDetectionConfig`

NewResponseTimeDegradationDetectionConfigWithDefaults instantiates a new ResponseTimeDegradationDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionMode

`func (o *ResponseTimeDegradationDetectionConfig) GetDetectionMode() string`

GetDetectionMode returns the DetectionMode field if non-nil, zero value otherwise.

### GetDetectionModeOk

`func (o *ResponseTimeDegradationDetectionConfig) GetDetectionModeOk() (*string, bool)`

GetDetectionModeOk returns a tuple with the DetectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMode

`func (o *ResponseTimeDegradationDetectionConfig) SetDetectionMode(v string)`

SetDetectionMode sets DetectionMode field to given value.


### GetAutomaticDetection

`func (o *ResponseTimeDegradationDetectionConfig) GetAutomaticDetection() ResponseTimeDegradationAutodetectionConfig`

GetAutomaticDetection returns the AutomaticDetection field if non-nil, zero value otherwise.

### GetAutomaticDetectionOk

`func (o *ResponseTimeDegradationDetectionConfig) GetAutomaticDetectionOk() (*ResponseTimeDegradationAutodetectionConfig, bool)`

GetAutomaticDetectionOk returns a tuple with the AutomaticDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDetection

`func (o *ResponseTimeDegradationDetectionConfig) SetAutomaticDetection(v ResponseTimeDegradationAutodetectionConfig)`

SetAutomaticDetection sets AutomaticDetection field to given value.

### HasAutomaticDetection

`func (o *ResponseTimeDegradationDetectionConfig) HasAutomaticDetection() bool`

HasAutomaticDetection returns a boolean if a field has been set.

### GetThresholds

`func (o *ResponseTimeDegradationDetectionConfig) GetThresholds() ResponseTimeDegradationThresholdConfig`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *ResponseTimeDegradationDetectionConfig) GetThresholdsOk() (*ResponseTimeDegradationThresholdConfig, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *ResponseTimeDegradationDetectionConfig) SetThresholds(v ResponseTimeDegradationThresholdConfig)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *ResponseTimeDegradationDetectionConfig) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


