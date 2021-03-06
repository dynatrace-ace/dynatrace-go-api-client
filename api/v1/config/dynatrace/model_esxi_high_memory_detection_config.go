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

// EsxiHighMemoryDetectionConfig The configuration of the memory saturation on ESXi host detection.
type EsxiHighMemoryDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	CustomThresholds *EsxiHighMemoryThresholds `json:"customThresholds,omitempty"`
}

// NewEsxiHighMemoryDetectionConfig instantiates a new EsxiHighMemoryDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsxiHighMemoryDetectionConfig(enabled bool, ) *EsxiHighMemoryDetectionConfig {
	this := EsxiHighMemoryDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewEsxiHighMemoryDetectionConfigWithDefaults instantiates a new EsxiHighMemoryDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsxiHighMemoryDetectionConfigWithDefaults() *EsxiHighMemoryDetectionConfig {
	this := EsxiHighMemoryDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *EsxiHighMemoryDetectionConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EsxiHighMemoryDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EsxiHighMemoryDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *EsxiHighMemoryDetectionConfig) GetCustomThresholds() EsxiHighMemoryThresholds {
	if o == nil || o.CustomThresholds == nil {
		var ret EsxiHighMemoryThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsxiHighMemoryDetectionConfig) GetCustomThresholdsOk() (*EsxiHighMemoryThresholds, bool) {
	if o == nil || o.CustomThresholds == nil {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *EsxiHighMemoryDetectionConfig) HasCustomThresholds() bool {
	if o != nil && o.CustomThresholds != nil {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given EsxiHighMemoryThresholds and assigns it to the CustomThresholds field.
func (o *EsxiHighMemoryDetectionConfig) SetCustomThresholds(v EsxiHighMemoryThresholds) {
	o.CustomThresholds = &v
}

func (o EsxiHighMemoryDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.CustomThresholds != nil {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	return json.Marshal(toSerialize)
}

type NullableEsxiHighMemoryDetectionConfig struct {
	value *EsxiHighMemoryDetectionConfig
	isSet bool
}

func (v NullableEsxiHighMemoryDetectionConfig) Get() *EsxiHighMemoryDetectionConfig {
	return v.value
}

func (v *NullableEsxiHighMemoryDetectionConfig) Set(val *EsxiHighMemoryDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEsxiHighMemoryDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEsxiHighMemoryDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsxiHighMemoryDetectionConfig(val *EsxiHighMemoryDetectionConfig) *NullableEsxiHighMemoryDetectionConfig {
	return &NullableEsxiHighMemoryDetectionConfig{value: val, isSet: true}
}

func (v NullableEsxiHighMemoryDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsxiHighMemoryDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


