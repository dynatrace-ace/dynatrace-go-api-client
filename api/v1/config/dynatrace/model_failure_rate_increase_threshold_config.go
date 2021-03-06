/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// FailureRateIncreaseThresholdConfig Fixed thresholds for failure rate increase detection.    Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type FailureRateIncreaseThresholdConfig struct {
	// Failure rate during any 5-minute period to trigger an alert, %.
	Threshold int32 `json:"threshold"`
	// Sensitivity of the threshold.   With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.   With `high` sensitivity, no statistical confidence is used. Each violation triggers alert.
	Sensitivity string `json:"sensitivity"`
}

// NewFailureRateIncreaseThresholdConfig instantiates a new FailureRateIncreaseThresholdConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureRateIncreaseThresholdConfig(threshold int32, sensitivity string, ) *FailureRateIncreaseThresholdConfig {
	this := FailureRateIncreaseThresholdConfig{}
	this.Threshold = threshold
	this.Sensitivity = sensitivity
	return &this
}

// NewFailureRateIncreaseThresholdConfigWithDefaults instantiates a new FailureRateIncreaseThresholdConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureRateIncreaseThresholdConfigWithDefaults() *FailureRateIncreaseThresholdConfig {
	this := FailureRateIncreaseThresholdConfig{}
	return &this
}

// GetThreshold returns the Threshold field value
func (o *FailureRateIncreaseThresholdConfig) GetThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *FailureRateIncreaseThresholdConfig) GetThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *FailureRateIncreaseThresholdConfig) SetThreshold(v int32) {
	o.Threshold = v
}

// GetSensitivity returns the Sensitivity field value
func (o *FailureRateIncreaseThresholdConfig) GetSensitivity() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Sensitivity
}

// GetSensitivityOk returns a tuple with the Sensitivity field value
// and a boolean to check if the value has been set.
func (o *FailureRateIncreaseThresholdConfig) GetSensitivityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sensitivity, true
}

// SetSensitivity sets field value
func (o *FailureRateIncreaseThresholdConfig) SetSensitivity(v string) {
	o.Sensitivity = v
}

func (o FailureRateIncreaseThresholdConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["sensitivity"] = o.Sensitivity
	}
	return json.Marshal(toSerialize)
}

type NullableFailureRateIncreaseThresholdConfig struct {
	value *FailureRateIncreaseThresholdConfig
	isSet bool
}

func (v NullableFailureRateIncreaseThresholdConfig) Get() *FailureRateIncreaseThresholdConfig {
	return v.value
}

func (v *NullableFailureRateIncreaseThresholdConfig) Set(val *FailureRateIncreaseThresholdConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureRateIncreaseThresholdConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureRateIncreaseThresholdConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureRateIncreaseThresholdConfig(val *FailureRateIncreaseThresholdConfig) *NullableFailureRateIncreaseThresholdConfig {
	return &NullableFailureRateIncreaseThresholdConfig{value: val, isSet: true}
}

func (v NullableFailureRateIncreaseThresholdConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureRateIncreaseThresholdConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


