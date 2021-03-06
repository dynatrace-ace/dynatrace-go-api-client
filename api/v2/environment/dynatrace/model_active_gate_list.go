/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// ActiveGateList A list of ActiveGates.
type ActiveGateList struct {
	// A list of ActiveGates.
	ActiveGates *[]ActiveGate `json:"activeGates,omitempty"`
}

// NewActiveGateList instantiates a new ActiveGateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateList() *ActiveGateList {
	this := ActiveGateList{}
	return &this
}

// NewActiveGateListWithDefaults instantiates a new ActiveGateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateListWithDefaults() *ActiveGateList {
	this := ActiveGateList{}
	return &this
}

// GetActiveGates returns the ActiveGates field value if set, zero value otherwise.
func (o *ActiveGateList) GetActiveGates() []ActiveGate {
	if o == nil || o.ActiveGates == nil {
		var ret []ActiveGate
		return ret
	}
	return *o.ActiveGates
}

// GetActiveGatesOk returns a tuple with the ActiveGates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateList) GetActiveGatesOk() (*[]ActiveGate, bool) {
	if o == nil || o.ActiveGates == nil {
		return nil, false
	}
	return o.ActiveGates, true
}

// HasActiveGates returns a boolean if a field has been set.
func (o *ActiveGateList) HasActiveGates() bool {
	if o != nil && o.ActiveGates != nil {
		return true
	}

	return false
}

// SetActiveGates gets a reference to the given []ActiveGate and assigns it to the ActiveGates field.
func (o *ActiveGateList) SetActiveGates(v []ActiveGate) {
	o.ActiveGates = &v
}

func (o ActiveGateList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveGates != nil {
		toSerialize["activeGates"] = o.ActiveGates
	}
	return json.Marshal(toSerialize)
}

type NullableActiveGateList struct {
	value *ActiveGateList
	isSet bool
}

func (v NullableActiveGateList) Get() *ActiveGateList {
	return v.value
}

func (v *NullableActiveGateList) Set(val *ActiveGateList) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateList) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateList(val *ActiveGateList) *NullableActiveGateList {
	return &NullableActiveGateList{value: val, isSet: true}
}

func (v NullableActiveGateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


