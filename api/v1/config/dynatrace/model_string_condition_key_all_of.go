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

// StringConditionKeyAllOf struct for StringConditionKeyAllOf
type StringConditionKeyAllOf struct {
	// The key of the attribute, which need dynamic keys.   Not applicable otherwise, as the attibute itself acts as a key.
	DynamicKey *string `json:"dynamicKey,omitempty"`
}

// NewStringConditionKeyAllOf instantiates a new StringConditionKeyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringConditionKeyAllOf() *StringConditionKeyAllOf {
	this := StringConditionKeyAllOf{}
	return &this
}

// NewStringConditionKeyAllOfWithDefaults instantiates a new StringConditionKeyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringConditionKeyAllOfWithDefaults() *StringConditionKeyAllOf {
	this := StringConditionKeyAllOf{}
	return &this
}

// GetDynamicKey returns the DynamicKey field value if set, zero value otherwise.
func (o *StringConditionKeyAllOf) GetDynamicKey() string {
	if o == nil || o.DynamicKey == nil {
		var ret string
		return ret
	}
	return *o.DynamicKey
}

// GetDynamicKeyOk returns a tuple with the DynamicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringConditionKeyAllOf) GetDynamicKeyOk() (*string, bool) {
	if o == nil || o.DynamicKey == nil {
		return nil, false
	}
	return o.DynamicKey, true
}

// HasDynamicKey returns a boolean if a field has been set.
func (o *StringConditionKeyAllOf) HasDynamicKey() bool {
	if o != nil && o.DynamicKey != nil {
		return true
	}

	return false
}

// SetDynamicKey gets a reference to the given string and assigns it to the DynamicKey field.
func (o *StringConditionKeyAllOf) SetDynamicKey(v string) {
	o.DynamicKey = &v
}

func (o StringConditionKeyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicKey != nil {
		toSerialize["dynamicKey"] = o.DynamicKey
	}
	return json.Marshal(toSerialize)
}

type NullableStringConditionKeyAllOf struct {
	value *StringConditionKeyAllOf
	isSet bool
}

func (v NullableStringConditionKeyAllOf) Get() *StringConditionKeyAllOf {
	return v.value
}

func (v *NullableStringConditionKeyAllOf) Set(val *StringConditionKeyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStringConditionKeyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStringConditionKeyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringConditionKeyAllOf(val *StringConditionKeyAllOf) *NullableStringConditionKeyAllOf {
	return &NullableStringConditionKeyAllOf{value: val, isSet: true}
}

func (v NullableStringConditionKeyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringConditionKeyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


