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

// LoadSpikeDetectionConfig The configuration of load spikes detection.
type LoadSpikeDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Alert if the observed load is more than *X* % of the expected value.
	LoadSpikePercent *int32 `json:"loadSpikePercent,omitempty"`
	// Alert if the service stays in abnormal state for at least *X* minutes.
	MinAbnormalStateDurationInMinutes *int32 `json:"minAbnormalStateDurationInMinutes,omitempty"`
}

// NewLoadSpikeDetectionConfig instantiates a new LoadSpikeDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadSpikeDetectionConfig(enabled bool, ) *LoadSpikeDetectionConfig {
	this := LoadSpikeDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewLoadSpikeDetectionConfigWithDefaults instantiates a new LoadSpikeDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadSpikeDetectionConfigWithDefaults() *LoadSpikeDetectionConfig {
	this := LoadSpikeDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *LoadSpikeDetectionConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LoadSpikeDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LoadSpikeDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoadSpikePercent returns the LoadSpikePercent field value if set, zero value otherwise.
func (o *LoadSpikeDetectionConfig) GetLoadSpikePercent() int32 {
	if o == nil || o.LoadSpikePercent == nil {
		var ret int32
		return ret
	}
	return *o.LoadSpikePercent
}

// GetLoadSpikePercentOk returns a tuple with the LoadSpikePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSpikeDetectionConfig) GetLoadSpikePercentOk() (*int32, bool) {
	if o == nil || o.LoadSpikePercent == nil {
		return nil, false
	}
	return o.LoadSpikePercent, true
}

// HasLoadSpikePercent returns a boolean if a field has been set.
func (o *LoadSpikeDetectionConfig) HasLoadSpikePercent() bool {
	if o != nil && o.LoadSpikePercent != nil {
		return true
	}

	return false
}

// SetLoadSpikePercent gets a reference to the given int32 and assigns it to the LoadSpikePercent field.
func (o *LoadSpikeDetectionConfig) SetLoadSpikePercent(v int32) {
	o.LoadSpikePercent = &v
}

// GetMinAbnormalStateDurationInMinutes returns the MinAbnormalStateDurationInMinutes field value if set, zero value otherwise.
func (o *LoadSpikeDetectionConfig) GetMinAbnormalStateDurationInMinutes() int32 {
	if o == nil || o.MinAbnormalStateDurationInMinutes == nil {
		var ret int32
		return ret
	}
	return *o.MinAbnormalStateDurationInMinutes
}

// GetMinAbnormalStateDurationInMinutesOk returns a tuple with the MinAbnormalStateDurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSpikeDetectionConfig) GetMinAbnormalStateDurationInMinutesOk() (*int32, bool) {
	if o == nil || o.MinAbnormalStateDurationInMinutes == nil {
		return nil, false
	}
	return o.MinAbnormalStateDurationInMinutes, true
}

// HasMinAbnormalStateDurationInMinutes returns a boolean if a field has been set.
func (o *LoadSpikeDetectionConfig) HasMinAbnormalStateDurationInMinutes() bool {
	if o != nil && o.MinAbnormalStateDurationInMinutes != nil {
		return true
	}

	return false
}

// SetMinAbnormalStateDurationInMinutes gets a reference to the given int32 and assigns it to the MinAbnormalStateDurationInMinutes field.
func (o *LoadSpikeDetectionConfig) SetMinAbnormalStateDurationInMinutes(v int32) {
	o.MinAbnormalStateDurationInMinutes = &v
}

func (o LoadSpikeDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.LoadSpikePercent != nil {
		toSerialize["loadSpikePercent"] = o.LoadSpikePercent
	}
	if o.MinAbnormalStateDurationInMinutes != nil {
		toSerialize["minAbnormalStateDurationInMinutes"] = o.MinAbnormalStateDurationInMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableLoadSpikeDetectionConfig struct {
	value *LoadSpikeDetectionConfig
	isSet bool
}

func (v NullableLoadSpikeDetectionConfig) Get() *LoadSpikeDetectionConfig {
	return v.value
}

func (v *NullableLoadSpikeDetectionConfig) Set(val *LoadSpikeDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadSpikeDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadSpikeDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadSpikeDetectionConfig(val *LoadSpikeDetectionConfig) *NullableLoadSpikeDetectionConfig {
	return &NullableLoadSpikeDetectionConfig{value: val, isSet: true}
}

func (v NullableLoadSpikeDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadSpikeDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


