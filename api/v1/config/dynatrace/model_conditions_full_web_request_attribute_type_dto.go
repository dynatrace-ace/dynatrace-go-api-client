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

// ConditionsFullWebRequestAttributeTypeDto A condition of the service detection rule.
type ConditionsFullWebRequestAttributeTypeDto struct {
	// The type of the attribute to be checked.
	AttributeType string `json:"attributeType"`
	// A list of conditions for the rule.   If several conditions are specified, the AND logic applies.
	CompareOperations *[]CompareOperation `json:"compareOperations,omitempty"`
}

// NewConditionsFullWebRequestAttributeTypeDto instantiates a new ConditionsFullWebRequestAttributeTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionsFullWebRequestAttributeTypeDto(attributeType string, ) *ConditionsFullWebRequestAttributeTypeDto {
	this := ConditionsFullWebRequestAttributeTypeDto{}
	this.AttributeType = attributeType
	return &this
}

// NewConditionsFullWebRequestAttributeTypeDtoWithDefaults instantiates a new ConditionsFullWebRequestAttributeTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionsFullWebRequestAttributeTypeDtoWithDefaults() *ConditionsFullWebRequestAttributeTypeDto {
	this := ConditionsFullWebRequestAttributeTypeDto{}
	return &this
}

// GetAttributeType returns the AttributeType field value
func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *ConditionsFullWebRequestAttributeTypeDto) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetCompareOperations returns the CompareOperations field value if set, zero value otherwise.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperations() []CompareOperation {
	if o == nil || o.CompareOperations == nil {
		var ret []CompareOperation
		return ret
	}
	return *o.CompareOperations
}

// GetCompareOperationsOk returns a tuple with the CompareOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperationsOk() (*[]CompareOperation, bool) {
	if o == nil || o.CompareOperations == nil {
		return nil, false
	}
	return o.CompareOperations, true
}

// HasCompareOperations returns a boolean if a field has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) HasCompareOperations() bool {
	if o != nil && o.CompareOperations != nil {
		return true
	}

	return false
}

// SetCompareOperations gets a reference to the given []CompareOperation and assigns it to the CompareOperations field.
func (o *ConditionsFullWebRequestAttributeTypeDto) SetCompareOperations(v []CompareOperation) {
	o.CompareOperations = &v
}

func (o ConditionsFullWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeType"] = o.AttributeType
	}
	if o.CompareOperations != nil {
		toSerialize["compareOperations"] = o.CompareOperations
	}
	return json.Marshal(toSerialize)
}

type NullableConditionsFullWebRequestAttributeTypeDto struct {
	value *ConditionsFullWebRequestAttributeTypeDto
	isSet bool
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) Get() *ConditionsFullWebRequestAttributeTypeDto {
	return v.value
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) Set(val *ConditionsFullWebRequestAttributeTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionsFullWebRequestAttributeTypeDto(val *ConditionsFullWebRequestAttributeTypeDto) *NullableConditionsFullWebRequestAttributeTypeDto {
	return &NullableConditionsFullWebRequestAttributeTypeDto{value: val, isSet: true}
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


