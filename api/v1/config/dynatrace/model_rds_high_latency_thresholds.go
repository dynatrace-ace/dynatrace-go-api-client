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

// RdsHighLatencyThresholds Custom thresholds for high RDS write/read latency. If not set, automatic mode is used
type RdsHighLatencyThresholds struct {
	// Alert if read/write latency is higher than *X* milliseconds in 3 out of 5 samples.
	WriteReadLatency int32 `json:"writeReadLatency"`
}

// NewRdsHighLatencyThresholds instantiates a new RdsHighLatencyThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsHighLatencyThresholds(writeReadLatency int32, ) *RdsHighLatencyThresholds {
	this := RdsHighLatencyThresholds{}
	this.WriteReadLatency = writeReadLatency
	return &this
}

// NewRdsHighLatencyThresholdsWithDefaults instantiates a new RdsHighLatencyThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsHighLatencyThresholdsWithDefaults() *RdsHighLatencyThresholds {
	this := RdsHighLatencyThresholds{}
	return &this
}

// GetWriteReadLatency returns the WriteReadLatency field value
func (o *RdsHighLatencyThresholds) GetWriteReadLatency() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.WriteReadLatency
}

// GetWriteReadLatencyOk returns a tuple with the WriteReadLatency field value
// and a boolean to check if the value has been set.
func (o *RdsHighLatencyThresholds) GetWriteReadLatencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WriteReadLatency, true
}

// SetWriteReadLatency sets field value
func (o *RdsHighLatencyThresholds) SetWriteReadLatency(v int32) {
	o.WriteReadLatency = v
}

func (o RdsHighLatencyThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["writeReadLatency"] = o.WriteReadLatency
	}
	return json.Marshal(toSerialize)
}

type NullableRdsHighLatencyThresholds struct {
	value *RdsHighLatencyThresholds
	isSet bool
}

func (v NullableRdsHighLatencyThresholds) Get() *RdsHighLatencyThresholds {
	return v.value
}

func (v *NullableRdsHighLatencyThresholds) Set(val *RdsHighLatencyThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsHighLatencyThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsHighLatencyThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsHighLatencyThresholds(val *RdsHighLatencyThresholds) *NullableRdsHighLatencyThresholds {
	return &NullableRdsHighLatencyThresholds{value: val, isSet: true}
}

func (v NullableRdsHighLatencyThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsHighLatencyThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


