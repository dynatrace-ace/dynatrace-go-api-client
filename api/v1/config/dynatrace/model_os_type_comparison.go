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

// OsTypeComparison Comparison for `OS_TYPE` attributes.
type OsTypeComparison struct {
	ComparisonBasic
	// Operator of the comparison. You can reverse it by setting **negate** to `true`.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Operator string `json:"operator"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
}

// NewOsTypeComparison instantiates a new OsTypeComparison object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsTypeComparison(operator string, ) *OsTypeComparison {
	this := OsTypeComparison{}
	this.Operator = operator
	return &this
}

// NewOsTypeComparisonWithDefaults instantiates a new OsTypeComparison object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsTypeComparisonWithDefaults() *OsTypeComparison {
	this := OsTypeComparison{}
	return &this
}

// GetOperator returns the Operator field value
func (o *OsTypeComparison) GetOperator() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *OsTypeComparison) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *OsTypeComparison) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OsTypeComparison) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeComparison) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OsTypeComparison) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OsTypeComparison) SetValue(v string) {
	o.Value = &v
}

func (o OsTypeComparison) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedComparisonBasic, errComparisonBasic := json.Marshal(o.ComparisonBasic)
	if errComparisonBasic != nil {
		return []byte{}, errComparisonBasic
	}
	errComparisonBasic = json.Unmarshal([]byte(serializedComparisonBasic), &toSerialize)
	if errComparisonBasic != nil {
		return []byte{}, errComparisonBasic
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOsTypeComparison struct {
	value *OsTypeComparison
	isSet bool
}

func (v NullableOsTypeComparison) Get() *OsTypeComparison {
	return v.value
}

func (v *NullableOsTypeComparison) Set(val *OsTypeComparison) {
	v.value = val
	v.isSet = true
}

func (v NullableOsTypeComparison) IsSet() bool {
	return v.isSet
}

func (v *NullableOsTypeComparison) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsTypeComparison(val *OsTypeComparison) *NullableOsTypeComparison {
	return &NullableOsTypeComparison{value: val, isSet: true}
}

func (v NullableOsTypeComparison) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsTypeComparison) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


