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

// PaasTypeComparisonAllOf struct for PaasTypeComparisonAllOf
type PaasTypeComparisonAllOf struct {
	// Operator of the comparison. You can reverse it by setting **negate** to `true`.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Operator *string `json:"operator,omitempty"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
}

// NewPaasTypeComparisonAllOf instantiates a new PaasTypeComparisonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaasTypeComparisonAllOf() *PaasTypeComparisonAllOf {
	this := PaasTypeComparisonAllOf{}
	return &this
}

// NewPaasTypeComparisonAllOfWithDefaults instantiates a new PaasTypeComparisonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaasTypeComparisonAllOfWithDefaults() *PaasTypeComparisonAllOf {
	this := PaasTypeComparisonAllOf{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *PaasTypeComparisonAllOf) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaasTypeComparisonAllOf) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *PaasTypeComparisonAllOf) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *PaasTypeComparisonAllOf) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PaasTypeComparisonAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaasTypeComparisonAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PaasTypeComparisonAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PaasTypeComparisonAllOf) SetValue(v string) {
	o.Value = &v
}

func (o PaasTypeComparisonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePaasTypeComparisonAllOf struct {
	value *PaasTypeComparisonAllOf
	isSet bool
}

func (v NullablePaasTypeComparisonAllOf) Get() *PaasTypeComparisonAllOf {
	return v.value
}

func (v *NullablePaasTypeComparisonAllOf) Set(val *PaasTypeComparisonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaasTypeComparisonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaasTypeComparisonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaasTypeComparisonAllOf(val *PaasTypeComparisonAllOf) *NullablePaasTypeComparisonAllOf {
	return &NullablePaasTypeComparisonAllOf{value: val, isSet: true}
}

func (v NullablePaasTypeComparisonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaasTypeComparisonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


