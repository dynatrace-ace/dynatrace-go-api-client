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

// StringConditionKey The key for dynamic attributes of the `STRING` type.
type StringConditionKey struct {
	ConditionKey
	// The key of the attribute, which need dynamic keys.   Not applicable otherwise, as the attibute itself acts as a key.
	DynamicKey string `json:"dynamicKey"`
}

// NewStringConditionKey instantiates a new StringConditionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringConditionKey(dynamicKey string, ) *StringConditionKey {
	this := StringConditionKey{}
	this.DynamicKey = dynamicKey
	return &this
}

// NewStringConditionKeyWithDefaults instantiates a new StringConditionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringConditionKeyWithDefaults() *StringConditionKey {
	this := StringConditionKey{}
	return &this
}

// GetDynamicKey returns the DynamicKey field value
func (o *StringConditionKey) GetDynamicKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DynamicKey
}

// GetDynamicKeyOk returns a tuple with the DynamicKey field value
// and a boolean to check if the value has been set.
func (o *StringConditionKey) GetDynamicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DynamicKey, true
}

// SetDynamicKey sets field value
func (o *StringConditionKey) SetDynamicKey(v string) {
	o.DynamicKey = v
}

func (o StringConditionKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConditionKey, errConditionKey := json.Marshal(o.ConditionKey)
	if errConditionKey != nil {
		return []byte{}, errConditionKey
	}
	errConditionKey = json.Unmarshal([]byte(serializedConditionKey), &toSerialize)
	if errConditionKey != nil {
		return []byte{}, errConditionKey
	}
	if true {
		toSerialize["dynamicKey"] = o.DynamicKey
	}
	return json.Marshal(toSerialize)
}

type NullableStringConditionKey struct {
	value *StringConditionKey
	isSet bool
}

func (v NullableStringConditionKey) Get() *StringConditionKey {
	return v.value
}

func (v *NullableStringConditionKey) Set(val *StringConditionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableStringConditionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableStringConditionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringConditionKey(val *StringConditionKey) *NullableStringConditionKey {
	return &NullableStringConditionKey{value: val, isSet: true}
}

func (v NullableStringConditionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringConditionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


