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

// ConditionsOpaqueAndExternalWebRequestAttributeTypeDto A condition of the service detection rule.
type ConditionsOpaqueAndExternalWebRequestAttributeTypeDto struct {
	// The type of the attribute to be checked.
	AttributeType string `json:"attributeType"`
	// A list of conditions for the rule.   If several conditions are specified, the AND logic applies.
	CompareOperations *[]CompareOperation `json:"compareOperations,omitempty"`
}

// NewConditionsOpaqueAndExternalWebRequestAttributeTypeDto instantiates a new ConditionsOpaqueAndExternalWebRequestAttributeTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionsOpaqueAndExternalWebRequestAttributeTypeDto(attributeType string, ) *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto {
	this := ConditionsOpaqueAndExternalWebRequestAttributeTypeDto{}
	this.AttributeType = attributeType
	return &this
}

// NewConditionsOpaqueAndExternalWebRequestAttributeTypeDtoWithDefaults instantiates a new ConditionsOpaqueAndExternalWebRequestAttributeTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionsOpaqueAndExternalWebRequestAttributeTypeDtoWithDefaults() *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto {
	this := ConditionsOpaqueAndExternalWebRequestAttributeTypeDto{}
	return &this
}

// GetAttributeType returns the AttributeType field value
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetAttributeType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetAttributeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetCompareOperations returns the CompareOperations field value if set, zero value otherwise.
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetCompareOperations() []CompareOperation {
	if o == nil || o.CompareOperations == nil {
		var ret []CompareOperation
		return ret
	}
	return *o.CompareOperations
}

// GetCompareOperationsOk returns a tuple with the CompareOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) GetCompareOperationsOk() (*[]CompareOperation, bool) {
	if o == nil || o.CompareOperations == nil {
		return nil, false
	}
	return o.CompareOperations, true
}

// HasCompareOperations returns a boolean if a field has been set.
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) HasCompareOperations() bool {
	if o != nil && o.CompareOperations != nil {
		return true
	}

	return false
}

// SetCompareOperations gets a reference to the given []CompareOperation and assigns it to the CompareOperations field.
func (o *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) SetCompareOperations(v []CompareOperation) {
	o.CompareOperations = &v
}

func (o ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeType"] = o.AttributeType
	}
	if o.CompareOperations != nil {
		toSerialize["compareOperations"] = o.CompareOperations
	}
	return json.Marshal(toSerialize)
}

type NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto struct {
	value *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto
	isSet bool
}

func (v NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) Get() *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto {
	return v.value
}

func (v *NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) Set(val *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto(val *ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) *NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto {
	return &NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto{value: val, isSet: true}
}

func (v NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionsOpaqueAndExternalWebRequestAttributeTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


