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

// ImpactAnalysis A list of all impacts of the problem.
type ImpactAnalysis struct {
	// A list of all impacts of the problem.
	Impacts *[]Impact `json:"impacts,omitempty"`
}

// NewImpactAnalysis instantiates a new ImpactAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpactAnalysis() *ImpactAnalysis {
	this := ImpactAnalysis{}
	return &this
}

// NewImpactAnalysisWithDefaults instantiates a new ImpactAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpactAnalysisWithDefaults() *ImpactAnalysis {
	this := ImpactAnalysis{}
	return &this
}

// GetImpacts returns the Impacts field value if set, zero value otherwise.
func (o *ImpactAnalysis) GetImpacts() []Impact {
	if o == nil || o.Impacts == nil {
		var ret []Impact
		return ret
	}
	return *o.Impacts
}

// GetImpactsOk returns a tuple with the Impacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImpactAnalysis) GetImpactsOk() (*[]Impact, bool) {
	if o == nil || o.Impacts == nil {
		return nil, false
	}
	return o.Impacts, true
}

// HasImpacts returns a boolean if a field has been set.
func (o *ImpactAnalysis) HasImpacts() bool {
	if o != nil && o.Impacts != nil {
		return true
	}

	return false
}

// SetImpacts gets a reference to the given []Impact and assigns it to the Impacts field.
func (o *ImpactAnalysis) SetImpacts(v []Impact) {
	o.Impacts = &v
}

func (o ImpactAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Impacts != nil {
		toSerialize["impacts"] = o.Impacts
	}
	return json.Marshal(toSerialize)
}

type NullableImpactAnalysis struct {
	value *ImpactAnalysis
	isSet bool
}

func (v NullableImpactAnalysis) Get() *ImpactAnalysis {
	return v.value
}

func (v *NullableImpactAnalysis) Set(val *ImpactAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableImpactAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableImpactAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpactAnalysis(val *ImpactAnalysis) *NullableImpactAnalysis {
	return &NullableImpactAnalysis{value: val, isSet: true}
}

func (v NullableImpactAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImpactAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


