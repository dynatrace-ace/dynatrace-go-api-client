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

// HttpStatusClassComparisonInfoAllOf struct for HttpStatusClassComparisonInfoAllOf
type HttpStatusClassComparisonInfoAllOf struct {
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison *string `json:"comparison,omitempty"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
}

// NewHttpStatusClassComparisonInfoAllOf instantiates a new HttpStatusClassComparisonInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpStatusClassComparisonInfoAllOf() *HttpStatusClassComparisonInfoAllOf {
	this := HttpStatusClassComparisonInfoAllOf{}
	return &this
}

// NewHttpStatusClassComparisonInfoAllOfWithDefaults instantiates a new HttpStatusClassComparisonInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpStatusClassComparisonInfoAllOfWithDefaults() *HttpStatusClassComparisonInfoAllOf {
	this := HttpStatusClassComparisonInfoAllOf{}
	return &this
}

// GetComparison returns the Comparison field value if set, zero value otherwise.
func (o *HttpStatusClassComparisonInfoAllOf) GetComparison() string {
	if o == nil || o.Comparison == nil {
		var ret string
		return ret
	}
	return *o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpStatusClassComparisonInfoAllOf) GetComparisonOk() (*string, bool) {
	if o == nil || o.Comparison == nil {
		return nil, false
	}
	return o.Comparison, true
}

// HasComparison returns a boolean if a field has been set.
func (o *HttpStatusClassComparisonInfoAllOf) HasComparison() bool {
	if o != nil && o.Comparison != nil {
		return true
	}

	return false
}

// SetComparison gets a reference to the given string and assigns it to the Comparison field.
func (o *HttpStatusClassComparisonInfoAllOf) SetComparison(v string) {
	o.Comparison = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HttpStatusClassComparisonInfoAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpStatusClassComparisonInfoAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HttpStatusClassComparisonInfoAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HttpStatusClassComparisonInfoAllOf) SetValue(v string) {
	o.Value = &v
}

func (o HttpStatusClassComparisonInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comparison != nil {
		toSerialize["comparison"] = o.Comparison
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableHttpStatusClassComparisonInfoAllOf struct {
	value *HttpStatusClassComparisonInfoAllOf
	isSet bool
}

func (v NullableHttpStatusClassComparisonInfoAllOf) Get() *HttpStatusClassComparisonInfoAllOf {
	return v.value
}

func (v *NullableHttpStatusClassComparisonInfoAllOf) Set(val *HttpStatusClassComparisonInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpStatusClassComparisonInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpStatusClassComparisonInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpStatusClassComparisonInfoAllOf(val *HttpStatusClassComparisonInfoAllOf) *NullableHttpStatusClassComparisonInfoAllOf {
	return &NullableHttpStatusClassComparisonInfoAllOf{value: val, isSet: true}
}

func (v NullableHttpStatusClassComparisonInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpStatusClassComparisonInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


