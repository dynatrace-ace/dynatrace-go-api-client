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

// SimpleTechComparisonAllOf struct for SimpleTechComparisonAllOf
type SimpleTechComparisonAllOf struct {
	// Operator of the comparison. You can reverse it by setting **negate** to `true`.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Operator *string `json:"operator,omitempty"`
	Value *SimpleTech `json:"value,omitempty"`
}

// NewSimpleTechComparisonAllOf instantiates a new SimpleTechComparisonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleTechComparisonAllOf() *SimpleTechComparisonAllOf {
	this := SimpleTechComparisonAllOf{}
	return &this
}

// NewSimpleTechComparisonAllOfWithDefaults instantiates a new SimpleTechComparisonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleTechComparisonAllOfWithDefaults() *SimpleTechComparisonAllOf {
	this := SimpleTechComparisonAllOf{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SimpleTechComparisonAllOf) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTechComparisonAllOf) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SimpleTechComparisonAllOf) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SimpleTechComparisonAllOf) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SimpleTechComparisonAllOf) GetValue() SimpleTech {
	if o == nil || o.Value == nil {
		var ret SimpleTech
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleTechComparisonAllOf) GetValueOk() (*SimpleTech, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SimpleTechComparisonAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given SimpleTech and assigns it to the Value field.
func (o *SimpleTechComparisonAllOf) SetValue(v SimpleTech) {
	o.Value = &v
}

func (o SimpleTechComparisonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleTechComparisonAllOf struct {
	value *SimpleTechComparisonAllOf
	isSet bool
}

func (v NullableSimpleTechComparisonAllOf) Get() *SimpleTechComparisonAllOf {
	return v.value
}

func (v *NullableSimpleTechComparisonAllOf) Set(val *SimpleTechComparisonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleTechComparisonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleTechComparisonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleTechComparisonAllOf(val *SimpleTechComparisonAllOf) *NullableSimpleTechComparisonAllOf {
	return &NullableSimpleTechComparisonAllOf{value: val, isSet: true}
}

func (v NullableSimpleTechComparisonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleTechComparisonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


