/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// CustomApplicationImpact Analysis of problem impact to a custom application.
type CustomApplicationImpact struct {
	Impact
}

// NewCustomApplicationImpact instantiates a new CustomApplicationImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomApplicationImpact(impactType string, impactedEntity EntityStub, estimatedAffectedUsers int64) *CustomApplicationImpact {
	this := CustomApplicationImpact{}
	this.ImpactType = impactType
	this.ImpactedEntity = impactedEntity
	this.EstimatedAffectedUsers = estimatedAffectedUsers
	return &this
}

// NewCustomApplicationImpactWithDefaults instantiates a new CustomApplicationImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomApplicationImpactWithDefaults() *CustomApplicationImpact {
	this := CustomApplicationImpact{}
	return &this
}

func (o CustomApplicationImpact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedImpact, errImpact := json.Marshal(o.Impact)
	if errImpact != nil {
		return []byte{}, errImpact
	}
	errImpact = json.Unmarshal([]byte(serializedImpact), &toSerialize)
	if errImpact != nil {
		return []byte{}, errImpact
	}
	return json.Marshal(toSerialize)
}

type NullableCustomApplicationImpact struct {
	value *CustomApplicationImpact
	isSet bool
}

func (v NullableCustomApplicationImpact) Get() *CustomApplicationImpact {
	return v.value
}

func (v *NullableCustomApplicationImpact) Set(val *CustomApplicationImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomApplicationImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomApplicationImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomApplicationImpact(val *CustomApplicationImpact) *NullableCustomApplicationImpact {
	return &NullableCustomApplicationImpact{value: val, isSet: true}
}

func (v NullableCustomApplicationImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomApplicationImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


