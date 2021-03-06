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

// ParsingFieldTopValue struct for ParsingFieldTopValue
type ParsingFieldTopValue struct {
	// Top value parsing field name
	FieldName *string `json:"fieldName,omitempty"`
	// Top value parsing field occurrences
	Occurrences *[]Occurrence `json:"occurrences,omitempty"`
}

// NewParsingFieldTopValue instantiates a new ParsingFieldTopValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsingFieldTopValue() *ParsingFieldTopValue {
	this := ParsingFieldTopValue{}
	return &this
}

// NewParsingFieldTopValueWithDefaults instantiates a new ParsingFieldTopValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsingFieldTopValueWithDefaults() *ParsingFieldTopValue {
	this := ParsingFieldTopValue{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *ParsingFieldTopValue) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsingFieldTopValue) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *ParsingFieldTopValue) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *ParsingFieldTopValue) SetFieldName(v string) {
	o.FieldName = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *ParsingFieldTopValue) GetOccurrences() []Occurrence {
	if o == nil || o.Occurrences == nil {
		var ret []Occurrence
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsingFieldTopValue) GetOccurrencesOk() (*[]Occurrence, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *ParsingFieldTopValue) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []Occurrence and assigns it to the Occurrences field.
func (o *ParsingFieldTopValue) SetOccurrences(v []Occurrence) {
	o.Occurrences = &v
}

func (o ParsingFieldTopValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	return json.Marshal(toSerialize)
}

type NullableParsingFieldTopValue struct {
	value *ParsingFieldTopValue
	isSet bool
}

func (v NullableParsingFieldTopValue) Get() *ParsingFieldTopValue {
	return v.value
}

func (v *NullableParsingFieldTopValue) Set(val *ParsingFieldTopValue) {
	v.value = val
	v.isSet = true
}

func (v NullableParsingFieldTopValue) IsSet() bool {
	return v.isSet
}

func (v *NullableParsingFieldTopValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsingFieldTopValue(val *ParsingFieldTopValue) *NullableParsingFieldTopValue {
	return &NullableParsingFieldTopValue{value: val, isSet: true}
}

func (v NullableParsingFieldTopValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsingFieldTopValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


