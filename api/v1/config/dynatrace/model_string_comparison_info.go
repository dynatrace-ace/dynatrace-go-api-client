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

// StringComparisonInfo Comparison for `STRING` attributes.
type StringComparisonInfo struct {
	ComparisonInfo
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison string `json:"comparison"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
	// The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// NewStringComparisonInfo instantiates a new StringComparisonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringComparisonInfo(comparison string, ) *StringComparisonInfo {
	this := StringComparisonInfo{}
	this.Comparison = comparison
	return &this
}

// NewStringComparisonInfoWithDefaults instantiates a new StringComparisonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringComparisonInfoWithDefaults() *StringComparisonInfo {
	this := StringComparisonInfo{}
	return &this
}

// GetComparison returns the Comparison field value
func (o *StringComparisonInfo) GetComparison() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value
// and a boolean to check if the value has been set.
func (o *StringComparisonInfo) GetComparisonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Comparison, true
}

// SetComparison sets field value
func (o *StringComparisonInfo) SetComparison(v string) {
	o.Comparison = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringComparisonInfo) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringComparisonInfo) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringComparisonInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StringComparisonInfo) SetValue(v string) {
	o.Value = &v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *StringComparisonInfo) GetCaseSensitive() bool {
	if o == nil || o.CaseSensitive == nil {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringComparisonInfo) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || o.CaseSensitive == nil {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *StringComparisonInfo) HasCaseSensitive() bool {
	if o != nil && o.CaseSensitive != nil {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *StringComparisonInfo) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

func (o StringComparisonInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedComparisonInfo, errComparisonInfo := json.Marshal(o.ComparisonInfo)
	if errComparisonInfo != nil {
		return []byte{}, errComparisonInfo
	}
	errComparisonInfo = json.Unmarshal([]byte(serializedComparisonInfo), &toSerialize)
	if errComparisonInfo != nil {
		return []byte{}, errComparisonInfo
	}
	if true {
		toSerialize["comparison"] = o.Comparison
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.CaseSensitive != nil {
		toSerialize["caseSensitive"] = o.CaseSensitive
	}
	return json.Marshal(toSerialize)
}

type NullableStringComparisonInfo struct {
	value *StringComparisonInfo
	isSet bool
}

func (v NullableStringComparisonInfo) Get() *StringComparisonInfo {
	return v.value
}

func (v *NullableStringComparisonInfo) Set(val *StringComparisonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStringComparisonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStringComparisonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringComparisonInfo(val *StringComparisonInfo) *NullableStringComparisonInfo {
	return &NullableStringComparisonInfo{value: val, isSet: true}
}

func (v NullableStringComparisonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringComparisonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


