# FailureRateIncreaseThresholdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **int32** | Failure rate during any 5-minute period to trigger an alert, %. | 
**Sensitivity** | **string** | Sensitivity of the threshold.   With &#x60;low&#x60; sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won&#39;t trigger alerts.   With &#x60;high&#x60; sensitivity, no statistical confidence is used. Each violation triggers alert. | 

## Methods

### NewFailureRateIncreaseThresholdConfig

`func NewFailureRateIncreaseThresholdConfig(threshold int32, sensitivity string, ) *FailureRateIncreaseThresholdConfig`

NewFailureRateIncreaseThresholdConfig instantiates a new FailureRateIncreaseThresholdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureRateIncreaseThresholdConfigWithDefaults

`func NewFailureRateIncreaseThresholdConfigWithDefaults() *FailureRateIncreaseThresholdConfig`

NewFailureRateIncreaseThresholdConfigWithDefaults instantiates a new FailureRateIncreaseThresholdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *FailureRateIncreaseThresholdConfig) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *FailureRateIncreaseThresholdConfig) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *FailureRateIncreaseThresholdConfig) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetSensitivity

`func (o *FailureRateIncreaseThresholdConfig) GetSensitivity() string`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *FailureRateIncreaseThresholdConfig) GetSensitivityOk() (*string, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *FailureRateIncreaseThresholdConfig) SetSensitivity(v string)`

SetSensitivity sets Sensitivity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


