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

// SlowPhysicalStorageDetectionConfig The configuraiton of the physical storage device running slow detection.
type SlowPhysicalStorageDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	CustomThresholds *SlowPhysicalStorageThresholds `json:"customThresholds,omitempty"`
}

// NewSlowPhysicalStorageDetectionConfig instantiates a new SlowPhysicalStorageDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlowPhysicalStorageDetectionConfig(enabled bool, ) *SlowPhysicalStorageDetectionConfig {
	this := SlowPhysicalStorageDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewSlowPhysicalStorageDetectionConfigWithDefaults instantiates a new SlowPhysicalStorageDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlowPhysicalStorageDetectionConfigWithDefaults() *SlowPhysicalStorageDetectionConfig {
	this := SlowPhysicalStorageDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SlowPhysicalStorageDetectionConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SlowPhysicalStorageDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SlowPhysicalStorageDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *SlowPhysicalStorageDetectionConfig) GetCustomThresholds() SlowPhysicalStorageThresholds {
	if o == nil || o.CustomThresholds == nil {
		var ret SlowPhysicalStorageThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlowPhysicalStorageDetectionConfig) GetCustomThresholdsOk() (*SlowPhysicalStorageThresholds, bool) {
	if o == nil || o.CustomThresholds == nil {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *SlowPhysicalStorageDetectionConfig) HasCustomThresholds() bool {
	if o != nil && o.CustomThresholds != nil {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given SlowPhysicalStorageThresholds and assigns it to the CustomThresholds field.
func (o *SlowPhysicalStorageDetectionConfig) SetCustomThresholds(v SlowPhysicalStorageThresholds) {
	o.CustomThresholds = &v
}

func (o SlowPhysicalStorageDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CustomThresholds != nil {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	return json.Marshal(toSerialize)
}

type NullableSlowPhysicalStorageDetectionConfig struct {
	value *SlowPhysicalStorageDetectionConfig
	isSet bool
}

func (v NullableSlowPhysicalStorageDetectionConfig) Get() *SlowPhysicalStorageDetectionConfig {
	return v.value
}

func (v *NullableSlowPhysicalStorageDetectionConfig) Set(val *SlowPhysicalStorageDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSlowPhysicalStorageDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSlowPhysicalStorageDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlowPhysicalStorageDetectionConfig(val *SlowPhysicalStorageDetectionConfig) *NullableSlowPhysicalStorageDetectionConfig {
	return &NullableSlowPhysicalStorageDetectionConfig{value: val, isSet: true}
}

func (v NullableSlowPhysicalStorageDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlowPhysicalStorageDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


