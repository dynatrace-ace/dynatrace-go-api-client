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

// StartsWithCompareOperation The condition of the `STARTS_WITH` type.   The condition checks whether the string value starts with the specified text.
type StartsWithCompareOperation struct {
	CompareOperation
	// Inverts the operation of the condition. Set to `true` to turn **starts with** into **does not start with**.    If not set, then `false` is used.
	Negate *bool `json:"negate,omitempty"`
	// The condition is case sensitive (`false`) or case insensitive (`true`).   If not set, then `false` is used, making the condition case sensitive.
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// The value to compare to.   If several values are specified, the OR logic applies.
	Values *[]string `json:"values,omitempty"`
}

// NewStartsWithCompareOperation instantiates a new StartsWithCompareOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartsWithCompareOperation() *StartsWithCompareOperation {
	this := StartsWithCompareOperation{}
	return &this
}

// NewStartsWithCompareOperationWithDefaults instantiates a new StartsWithCompareOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartsWithCompareOperationWithDefaults() *StartsWithCompareOperation {
	this := StartsWithCompareOperation{}
	return &this
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *StartsWithCompareOperation) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartsWithCompareOperation) GetNegateOk() (*bool, bool) {
	if o == nil || o.Negate == nil {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *StartsWithCompareOperation) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *StartsWithCompareOperation) SetNegate(v bool) {
	o.Negate = &v
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *StartsWithCompareOperation) GetIgnoreCase() bool {
	if o == nil || o.IgnoreCase == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartsWithCompareOperation) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || o.IgnoreCase == nil {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *StartsWithCompareOperation) HasIgnoreCase() bool {
	if o != nil && o.IgnoreCase != nil {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *StartsWithCompareOperation) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *StartsWithCompareOperation) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartsWithCompareOperation) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *StartsWithCompareOperation) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *StartsWithCompareOperation) SetValues(v []string) {
	o.Values = &v
}

func (o StartsWithCompareOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCompareOperation, errCompareOperation := json.Marshal(o.CompareOperation)
	if errCompareOperation != nil {
		return []byte{}, errCompareOperation
	}
	errCompareOperation = json.Unmarshal([]byte(serializedCompareOperation), &toSerialize)
	if errCompareOperation != nil {
		return []byte{}, errCompareOperation
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
	if o.IgnoreCase != nil {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableStartsWithCompareOperation struct {
	value *StartsWithCompareOperation
	isSet bool
}

func (v NullableStartsWithCompareOperation) Get() *StartsWithCompareOperation {
	return v.value
}

func (v *NullableStartsWithCompareOperation) Set(val *StartsWithCompareOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableStartsWithCompareOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableStartsWithCompareOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartsWithCompareOperation(val *StartsWithCompareOperation) *NullableStartsWithCompareOperation {
	return &NullableStartsWithCompareOperation{value: val, isSet: true}
}

func (v NullableStartsWithCompareOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartsWithCompareOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


