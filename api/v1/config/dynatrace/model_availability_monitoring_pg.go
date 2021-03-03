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

// AvailabilityMonitoringPG Configuration of the availability monitoring for the process group.
type AvailabilityMonitoringPG struct {
	// How to monitor the availability of the process group:   * `PROCESS_IMPACT`: Alert if any process of the group becomes unavailable.  * `MINIMUM_THRESHOLD`: Alert if the number of active processes in the group falls below the specified threshold.  * `OFF`: Availability monitoring is disabled.
	Method string `json:"method"`
	// Alert if the number of active processes in the group is lower than this value.
	MinimumThreshold *int32 `json:"minimumThreshold,omitempty"`
}

// NewAvailabilityMonitoringPG instantiates a new AvailabilityMonitoringPG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityMonitoringPG(method string, ) *AvailabilityMonitoringPG {
	this := AvailabilityMonitoringPG{}
	this.Method = method
	return &this
}

// NewAvailabilityMonitoringPGWithDefaults instantiates a new AvailabilityMonitoringPG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityMonitoringPGWithDefaults() *AvailabilityMonitoringPG {
	this := AvailabilityMonitoringPG{}
	return &this
}

// GetMethod returns the Method field value
func (o *AvailabilityMonitoringPG) GetMethod() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *AvailabilityMonitoringPG) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *AvailabilityMonitoringPG) SetMethod(v string) {
	o.Method = v
}

// GetMinimumThreshold returns the MinimumThreshold field value if set, zero value otherwise.
func (o *AvailabilityMonitoringPG) GetMinimumThreshold() int32 {
	if o == nil || o.MinimumThreshold == nil {
		var ret int32
		return ret
	}
	return *o.MinimumThreshold
}

// GetMinimumThresholdOk returns a tuple with the MinimumThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityMonitoringPG) GetMinimumThresholdOk() (*int32, bool) {
	if o == nil || o.MinimumThreshold == nil {
		return nil, false
	}
	return o.MinimumThreshold, true
}

// HasMinimumThreshold returns a boolean if a field has been set.
func (o *AvailabilityMonitoringPG) HasMinimumThreshold() bool {
	if o != nil && o.MinimumThreshold != nil {
		return true
	}

	return false
}

// SetMinimumThreshold gets a reference to the given int32 and assigns it to the MinimumThreshold field.
func (o *AvailabilityMonitoringPG) SetMinimumThreshold(v int32) {
	o.MinimumThreshold = &v
}

func (o AvailabilityMonitoringPG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.MinimumThreshold != nil {
		toSerialize["minimumThreshold"] = o.MinimumThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityMonitoringPG struct {
	value *AvailabilityMonitoringPG
	isSet bool
}

func (v NullableAvailabilityMonitoringPG) Get() *AvailabilityMonitoringPG {
	return v.value
}

func (v *NullableAvailabilityMonitoringPG) Set(val *AvailabilityMonitoringPG) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityMonitoringPG) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityMonitoringPG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityMonitoringPG(val *AvailabilityMonitoringPG) *NullableAvailabilityMonitoringPG {
	return &NullableAvailabilityMonitoringPG{value: val, isSet: true}
}

func (v NullableAvailabilityMonitoringPG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityMonitoringPG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

