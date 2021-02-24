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

// FailureRateIncreaseAutodetectionConfig Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.   The absolute and relative thresholds **both** must exceed to trigger an alert.   Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
type FailureRateIncreaseAutodetectionConfig struct {
	// Absolute increase of failing service calls to trigger an alert, %.
	FailingServiceCallPercentageIncreaseAbsolute int32 `json:"failingServiceCallPercentageIncreaseAbsolute"`
	// Relative increase of failing service calls to trigger an alert, %.
	FailingServiceCallPercentageIncreaseRelative int32 `json:"failingServiceCallPercentageIncreaseRelative"`
}

// NewFailureRateIncreaseAutodetectionConfig instantiates a new FailureRateIncreaseAutodetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureRateIncreaseAutodetectionConfig(failingServiceCallPercentageIncreaseAbsolute int32, failingServiceCallPercentageIncreaseRelative int32, ) *FailureRateIncreaseAutodetectionConfig {
	this := FailureRateIncreaseAutodetectionConfig{}
	this.FailingServiceCallPercentageIncreaseAbsolute = failingServiceCallPercentageIncreaseAbsolute
	this.FailingServiceCallPercentageIncreaseRelative = failingServiceCallPercentageIncreaseRelative
	return &this
}

// NewFailureRateIncreaseAutodetectionConfigWithDefaults instantiates a new FailureRateIncreaseAutodetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureRateIncreaseAutodetectionConfigWithDefaults() *FailureRateIncreaseAutodetectionConfig {
	this := FailureRateIncreaseAutodetectionConfig{}
	return &this
}

// GetFailingServiceCallPercentageIncreaseAbsolute returns the FailingServiceCallPercentageIncreaseAbsolute field value
func (o *FailureRateIncreaseAutodetectionConfig) GetFailingServiceCallPercentageIncreaseAbsolute() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FailingServiceCallPercentageIncreaseAbsolute
}

// GetFailingServiceCallPercentageIncreaseAbsoluteOk returns a tuple with the FailingServiceCallPercentageIncreaseAbsolute field value
// and a boolean to check if the value has been set.
func (o *FailureRateIncreaseAutodetectionConfig) GetFailingServiceCallPercentageIncreaseAbsoluteOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FailingServiceCallPercentageIncreaseAbsolute, true
}

// SetFailingServiceCallPercentageIncreaseAbsolute sets field value
func (o *FailureRateIncreaseAutodetectionConfig) SetFailingServiceCallPercentageIncreaseAbsolute(v int32) {
	o.FailingServiceCallPercentageIncreaseAbsolute = v
}

// GetFailingServiceCallPercentageIncreaseRelative returns the FailingServiceCallPercentageIncreaseRelative field value
func (o *FailureRateIncreaseAutodetectionConfig) GetFailingServiceCallPercentageIncreaseRelative() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FailingServiceCallPercentageIncreaseRelative
}

// GetFailingServiceCallPercentageIncreaseRelativeOk returns a tuple with the FailingServiceCallPercentageIncreaseRelative field value
// and a boolean to check if the value has been set.
func (o *FailureRateIncreaseAutodetectionConfig) GetFailingServiceCallPercentageIncreaseRelativeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FailingServiceCallPercentageIncreaseRelative, true
}

// SetFailingServiceCallPercentageIncreaseRelative sets field value
func (o *FailureRateIncreaseAutodetectionConfig) SetFailingServiceCallPercentageIncreaseRelative(v int32) {
	o.FailingServiceCallPercentageIncreaseRelative = v
}

func (o FailureRateIncreaseAutodetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["failingServiceCallPercentageIncreaseAbsolute"] = o.FailingServiceCallPercentageIncreaseAbsolute
	}
	if true {
		toSerialize["failingServiceCallPercentageIncreaseRelative"] = o.FailingServiceCallPercentageIncreaseRelative
	}
	return json.Marshal(toSerialize)
}

type NullableFailureRateIncreaseAutodetectionConfig struct {
	value *FailureRateIncreaseAutodetectionConfig
	isSet bool
}

func (v NullableFailureRateIncreaseAutodetectionConfig) Get() *FailureRateIncreaseAutodetectionConfig {
	return v.value
}

func (v *NullableFailureRateIncreaseAutodetectionConfig) Set(val *FailureRateIncreaseAutodetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureRateIncreaseAutodetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureRateIncreaseAutodetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureRateIncreaseAutodetectionConfig(val *FailureRateIncreaseAutodetectionConfig) *NullableFailureRateIncreaseAutodetectionConfig {
	return &NullableFailureRateIncreaseAutodetectionConfig{value: val, isSet: true}
}

func (v NullableFailureRateIncreaseAutodetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureRateIncreaseAutodetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


