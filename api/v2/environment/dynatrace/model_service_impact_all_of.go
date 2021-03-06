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

// ServiceImpactAllOf struct for ServiceImpactAllOf
type ServiceImpactAllOf struct {
	// The number of potentially impacted services.
	NumberOfPotentiallyAffectedServiceCalls *int64 `json:"numberOfPotentiallyAffectedServiceCalls,omitempty"`
}

// NewServiceImpactAllOf instantiates a new ServiceImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceImpactAllOf() *ServiceImpactAllOf {
	this := ServiceImpactAllOf{}
	return &this
}

// NewServiceImpactAllOfWithDefaults instantiates a new ServiceImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceImpactAllOfWithDefaults() *ServiceImpactAllOf {
	this := ServiceImpactAllOf{}
	return &this
}

// GetNumberOfPotentiallyAffectedServiceCalls returns the NumberOfPotentiallyAffectedServiceCalls field value if set, zero value otherwise.
func (o *ServiceImpactAllOf) GetNumberOfPotentiallyAffectedServiceCalls() int64 {
	if o == nil || o.NumberOfPotentiallyAffectedServiceCalls == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfPotentiallyAffectedServiceCalls
}

// GetNumberOfPotentiallyAffectedServiceCallsOk returns a tuple with the NumberOfPotentiallyAffectedServiceCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceImpactAllOf) GetNumberOfPotentiallyAffectedServiceCallsOk() (*int64, bool) {
	if o == nil || o.NumberOfPotentiallyAffectedServiceCalls == nil {
		return nil, false
	}
	return o.NumberOfPotentiallyAffectedServiceCalls, true
}

// HasNumberOfPotentiallyAffectedServiceCalls returns a boolean if a field has been set.
func (o *ServiceImpactAllOf) HasNumberOfPotentiallyAffectedServiceCalls() bool {
	if o != nil && o.NumberOfPotentiallyAffectedServiceCalls != nil {
		return true
	}

	return false
}

// SetNumberOfPotentiallyAffectedServiceCalls gets a reference to the given int64 and assigns it to the NumberOfPotentiallyAffectedServiceCalls field.
func (o *ServiceImpactAllOf) SetNumberOfPotentiallyAffectedServiceCalls(v int64) {
	o.NumberOfPotentiallyAffectedServiceCalls = &v
}

func (o ServiceImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumberOfPotentiallyAffectedServiceCalls != nil {
		toSerialize["numberOfPotentiallyAffectedServiceCalls"] = o.NumberOfPotentiallyAffectedServiceCalls
	}
	return json.Marshal(toSerialize)
}

type NullableServiceImpactAllOf struct {
	value *ServiceImpactAllOf
	isSet bool
}

func (v NullableServiceImpactAllOf) Get() *ServiceImpactAllOf {
	return v.value
}

func (v *NullableServiceImpactAllOf) Set(val *ServiceImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceImpactAllOf(val *ServiceImpactAllOf) *NullableServiceImpactAllOf {
	return &NullableServiceImpactAllOf{value: val, isSet: true}
}

func (v NullableServiceImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


