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

// FlawStateComparisonInfo Comparison for `FLAW_STATE` attributes.
type FlawStateComparisonInfo struct {
	ComparisonInfo
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison string `json:"comparison"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
}

// NewFlawStateComparisonInfo instantiates a new FlawStateComparisonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlawStateComparisonInfo(comparison string, ) *FlawStateComparisonInfo {
	this := FlawStateComparisonInfo{}
	this.Comparison = comparison
	return &this
}

// NewFlawStateComparisonInfoWithDefaults instantiates a new FlawStateComparisonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlawStateComparisonInfoWithDefaults() *FlawStateComparisonInfo {
	this := FlawStateComparisonInfo{}
	return &this
}

// GetComparison returns the Comparison field value
func (o *FlawStateComparisonInfo) GetComparison() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value
// and a boolean to check if the value has been set.
func (o *FlawStateComparisonInfo) GetComparisonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Comparison, true
}

// SetComparison sets field value
func (o *FlawStateComparisonInfo) SetComparison(v string) {
	o.Comparison = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FlawStateComparisonInfo) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlawStateComparisonInfo) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FlawStateComparisonInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FlawStateComparisonInfo) SetValue(v string) {
	o.Value = &v
}

func (o FlawStateComparisonInfo) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableFlawStateComparisonInfo struct {
	value *FlawStateComparisonInfo
	isSet bool
}

func (v NullableFlawStateComparisonInfo) Get() *FlawStateComparisonInfo {
	return v.value
}

func (v *NullableFlawStateComparisonInfo) Set(val *FlawStateComparisonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFlawStateComparisonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFlawStateComparisonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlawStateComparisonInfo(val *FlawStateComparisonInfo) *NullableFlawStateComparisonInfo {
	return &NullableFlawStateComparisonInfo{value: val, isSet: true}
}

func (v NullableFlawStateComparisonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlawStateComparisonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


