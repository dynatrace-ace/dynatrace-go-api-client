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

// DiskLowInodesDetectionConfig Configuration of low disk inodes number detection.
type DiskLowInodesDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	CustomThresholds *DiskLowInodesThresholds `json:"customThresholds,omitempty"`
}

// NewDiskLowInodesDetectionConfig instantiates a new DiskLowInodesDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskLowInodesDetectionConfig(enabled bool, ) *DiskLowInodesDetectionConfig {
	this := DiskLowInodesDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewDiskLowInodesDetectionConfigWithDefaults instantiates a new DiskLowInodesDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskLowInodesDetectionConfigWithDefaults() *DiskLowInodesDetectionConfig {
	this := DiskLowInodesDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DiskLowInodesDetectionConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DiskLowInodesDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DiskLowInodesDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *DiskLowInodesDetectionConfig) GetCustomThresholds() DiskLowInodesThresholds {
	if o == nil || o.CustomThresholds == nil {
		var ret DiskLowInodesThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskLowInodesDetectionConfig) GetCustomThresholdsOk() (*DiskLowInodesThresholds, bool) {
	if o == nil || o.CustomThresholds == nil {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *DiskLowInodesDetectionConfig) HasCustomThresholds() bool {
	if o != nil && o.CustomThresholds != nil {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given DiskLowInodesThresholds and assigns it to the CustomThresholds field.
func (o *DiskLowInodesDetectionConfig) SetCustomThresholds(v DiskLowInodesThresholds) {
	o.CustomThresholds = &v
}

func (o DiskLowInodesDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CustomThresholds != nil {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	return json.Marshal(toSerialize)
}

type NullableDiskLowInodesDetectionConfig struct {
	value *DiskLowInodesDetectionConfig
	isSet bool
}

func (v NullableDiskLowInodesDetectionConfig) Get() *DiskLowInodesDetectionConfig {
	return v.value
}

func (v *NullableDiskLowInodesDetectionConfig) Set(val *DiskLowInodesDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskLowInodesDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskLowInodesDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskLowInodesDetectionConfig(val *DiskLowInodesDetectionConfig) *NullableDiskLowInodesDetectionConfig {
	return &NullableDiskLowInodesDetectionConfig{value: val, isSet: true}
}

func (v NullableDiskLowInodesDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskLowInodesDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


