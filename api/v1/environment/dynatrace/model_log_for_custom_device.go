/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// LogForCustomDevice The list of available logs.
type LogForCustomDevice struct {
	// The full path to the log.
	Path *string `json:"path,omitempty"`
	// The log is available (`true`) or not available (`false`) for analysis.
	AvailableForAnalysis *bool `json:"availableForAnalysis,omitempty"`
}

// NewLogForCustomDevice instantiates a new LogForCustomDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForCustomDevice() *LogForCustomDevice {
	this := LogForCustomDevice{}
	return &this
}

// NewLogForCustomDeviceWithDefaults instantiates a new LogForCustomDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForCustomDeviceWithDefaults() *LogForCustomDevice {
	this := LogForCustomDevice{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LogForCustomDevice) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForCustomDevice) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LogForCustomDevice) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LogForCustomDevice) SetPath(v string) {
	o.Path = &v
}

// GetAvailableForAnalysis returns the AvailableForAnalysis field value if set, zero value otherwise.
func (o *LogForCustomDevice) GetAvailableForAnalysis() bool {
	if o == nil || o.AvailableForAnalysis == nil {
		var ret bool
		return ret
	}
	return *o.AvailableForAnalysis
}

// GetAvailableForAnalysisOk returns a tuple with the AvailableForAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForCustomDevice) GetAvailableForAnalysisOk() (*bool, bool) {
	if o == nil || o.AvailableForAnalysis == nil {
		return nil, false
	}
	return o.AvailableForAnalysis, true
}

// HasAvailableForAnalysis returns a boolean if a field has been set.
func (o *LogForCustomDevice) HasAvailableForAnalysis() bool {
	if o != nil && o.AvailableForAnalysis != nil {
		return true
	}

	return false
}

// SetAvailableForAnalysis gets a reference to the given bool and assigns it to the AvailableForAnalysis field.
func (o *LogForCustomDevice) SetAvailableForAnalysis(v bool) {
	o.AvailableForAnalysis = &v
}

func (o LogForCustomDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.AvailableForAnalysis != nil {
		toSerialize["availableForAnalysis"] = o.AvailableForAnalysis
	}
	return json.Marshal(toSerialize)
}

type NullableLogForCustomDevice struct {
	value *LogForCustomDevice
	isSet bool
}

func (v NullableLogForCustomDevice) Get() *LogForCustomDevice {
	return v.value
}

func (v *NullableLogForCustomDevice) Set(val *LogForCustomDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForCustomDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForCustomDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForCustomDevice(val *LogForCustomDevice) *NullableLogForCustomDevice {
	return &NullableLogForCustomDevice{value: val, isSet: true}
}

func (v NullableLogForCustomDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForCustomDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

