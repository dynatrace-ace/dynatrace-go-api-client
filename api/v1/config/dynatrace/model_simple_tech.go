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

// SimpleTech The value to compare to.
type SimpleTech struct {
	// Predefined technology, if technology is not predefined, then the verbatim type must be set
	Type *string `json:"type,omitempty"`
	// Non-predefined technology, use for custom technologies.
	VerbatimType *string `json:"verbatimType,omitempty"`
}

// NewSimpleTech instantiates a new SimpleTech object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleTech() *SimpleTech {
	this := SimpleTech{}
	return &this
}

// NewSimpleTechWithDefaults instantiates a new SimpleTech object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleTechWithDefaults() *SimpleTech {
	this := SimpleTech{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimpleTech) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTech) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimpleTech) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimpleTech) SetType(v string) {
	o.Type = &v
}

// GetVerbatimType returns the VerbatimType field value if set, zero value otherwise.
func (o *SimpleTech) GetVerbatimType() string {
	if o == nil || o.VerbatimType == nil {
		var ret string
		return ret
	}
	return *o.VerbatimType
}

// GetVerbatimTypeOk returns a tuple with the VerbatimType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTech) GetVerbatimTypeOk() (*string, bool) {
	if o == nil || o.VerbatimType == nil {
		return nil, false
	}
	return o.VerbatimType, true
}

// HasVerbatimType returns a boolean if a field has been set.
func (o *SimpleTech) HasVerbatimType() bool {
	if o != nil && o.VerbatimType != nil {
		return true
	}

	return false
}

// SetVerbatimType gets a reference to the given string and assigns it to the VerbatimType field.
func (o *SimpleTech) SetVerbatimType(v string) {
	o.VerbatimType = &v
}

func (o SimpleTech) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.VerbatimType != nil {
		toSerialize["verbatimType"] = o.VerbatimType
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleTech struct {
	value *SimpleTech
	isSet bool
}

func (v NullableSimpleTech) Get() *SimpleTech {
	return v.value
}

func (v *NullableSimpleTech) Set(val *SimpleTech) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleTech) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleTech) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleTech(val *SimpleTech) *NullableSimpleTech {
	return &NullableSimpleTech{value: val, isSet: true}
}

func (v NullableSimpleTech) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleTech) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


