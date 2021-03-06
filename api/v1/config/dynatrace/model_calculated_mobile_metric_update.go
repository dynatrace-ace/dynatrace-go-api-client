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

// CalculatedMobileMetricUpdate Update of the calculated mobile or custom metric.
type CalculatedMobileMetricUpdate struct {
	// The metric is enabled (`true`) or disabled (`false`).
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCalculatedMobileMetricUpdate instantiates a new CalculatedMobileMetricUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculatedMobileMetricUpdate() *CalculatedMobileMetricUpdate {
	this := CalculatedMobileMetricUpdate{}
	return &this
}

// NewCalculatedMobileMetricUpdateWithDefaults instantiates a new CalculatedMobileMetricUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculatedMobileMetricUpdateWithDefaults() *CalculatedMobileMetricUpdate {
	this := CalculatedMobileMetricUpdate{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CalculatedMobileMetricUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculatedMobileMetricUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CalculatedMobileMetricUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CalculatedMobileMetricUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CalculatedMobileMetricUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCalculatedMobileMetricUpdate struct {
	value *CalculatedMobileMetricUpdate
	isSet bool
}

func (v NullableCalculatedMobileMetricUpdate) Get() *CalculatedMobileMetricUpdate {
	return v.value
}

func (v *NullableCalculatedMobileMetricUpdate) Set(val *CalculatedMobileMetricUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculatedMobileMetricUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculatedMobileMetricUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculatedMobileMetricUpdate(val *CalculatedMobileMetricUpdate) *NullableCalculatedMobileMetricUpdate {
	return &NullableCalculatedMobileMetricUpdate{value: val, isSet: true}
}

func (v NullableCalculatedMobileMetricUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculatedMobileMetricUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


