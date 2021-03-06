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

// ProcessGroupIdAlertingScope A scope filter for a process group identifier.
type ProcessGroupIdAlertingScope struct {
	MetricEventAlertingScope
	// The process groups id to match on.
	ProcessGroupId string `json:"processGroupId"`
}

// NewProcessGroupIdAlertingScope instantiates a new ProcessGroupIdAlertingScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupIdAlertingScope(processGroupId string, ) *ProcessGroupIdAlertingScope {
	this := ProcessGroupIdAlertingScope{}
	this.ProcessGroupId = processGroupId
	return &this
}

// NewProcessGroupIdAlertingScopeWithDefaults instantiates a new ProcessGroupIdAlertingScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupIdAlertingScopeWithDefaults() *ProcessGroupIdAlertingScope {
	this := ProcessGroupIdAlertingScope{}
	return &this
}

// GetProcessGroupId returns the ProcessGroupId field value
func (o *ProcessGroupIdAlertingScope) GetProcessGroupId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProcessGroupId
}

// GetProcessGroupIdOk returns a tuple with the ProcessGroupId field value
// and a boolean to check if the value has been set.
func (o *ProcessGroupIdAlertingScope) GetProcessGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessGroupId, true
}

// SetProcessGroupId sets field value
func (o *ProcessGroupIdAlertingScope) SetProcessGroupId(v string) {
	o.ProcessGroupId = v
}

func (o ProcessGroupIdAlertingScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMetricEventAlertingScope, errMetricEventAlertingScope := json.Marshal(o.MetricEventAlertingScope)
	if errMetricEventAlertingScope != nil {
		return []byte{}, errMetricEventAlertingScope
	}
	errMetricEventAlertingScope = json.Unmarshal([]byte(serializedMetricEventAlertingScope), &toSerialize)
	if errMetricEventAlertingScope != nil {
		return []byte{}, errMetricEventAlertingScope
	}
	if true {
		toSerialize["processGroupId"] = o.ProcessGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableProcessGroupIdAlertingScope struct {
	value *ProcessGroupIdAlertingScope
	isSet bool
}

func (v NullableProcessGroupIdAlertingScope) Get() *ProcessGroupIdAlertingScope {
	return v.value
}

func (v *NullableProcessGroupIdAlertingScope) Set(val *ProcessGroupIdAlertingScope) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupIdAlertingScope) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupIdAlertingScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupIdAlertingScope(val *ProcessGroupIdAlertingScope) *NullableProcessGroupIdAlertingScope {
	return &NullableProcessGroupIdAlertingScope{value: val, isSet: true}
}

func (v NullableProcessGroupIdAlertingScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupIdAlertingScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


