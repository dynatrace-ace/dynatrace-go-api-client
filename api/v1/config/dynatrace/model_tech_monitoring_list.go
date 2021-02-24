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

// TechMonitoringList A list of technology monitoring configurations.
type TechMonitoringList struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// A list of technology monitoring configurations.
	Technologies []Technology `json:"technologies"`
}

// NewTechMonitoringList instantiates a new TechMonitoringList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechMonitoringList(technologies []Technology, ) *TechMonitoringList {
	this := TechMonitoringList{}
	this.Technologies = technologies
	return &this
}

// NewTechMonitoringListWithDefaults instantiates a new TechMonitoringList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechMonitoringListWithDefaults() *TechMonitoringList {
	this := TechMonitoringList{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TechMonitoringList) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechMonitoringList) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TechMonitoringList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *TechMonitoringList) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetTechnologies returns the Technologies field value
func (o *TechMonitoringList) GetTechnologies() []Technology {
	if o == nil  {
		var ret []Technology
		return ret
	}

	return o.Technologies
}

// GetTechnologiesOk returns a tuple with the Technologies field value
// and a boolean to check if the value has been set.
func (o *TechMonitoringList) GetTechnologiesOk() (*[]Technology, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Technologies, true
}

// SetTechnologies sets field value
func (o *TechMonitoringList) SetTechnologies(v []Technology) {
	o.Technologies = v
}

func (o TechMonitoringList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["technologies"] = o.Technologies
	}
	return json.Marshal(toSerialize)
}

type NullableTechMonitoringList struct {
	value *TechMonitoringList
	isSet bool
}

func (v NullableTechMonitoringList) Get() *TechMonitoringList {
	return v.value
}

func (v *NullableTechMonitoringList) Set(val *TechMonitoringList) {
	v.value = val
	v.isSet = true
}

func (v NullableTechMonitoringList) IsSet() bool {
	return v.isSet
}

func (v *NullableTechMonitoringList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechMonitoringList(val *TechMonitoringList) *NullableTechMonitoringList {
	return &NullableTechMonitoringList{value: val, isSet: true}
}

func (v NullableTechMonitoringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechMonitoringList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


