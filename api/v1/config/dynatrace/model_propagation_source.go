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

// PropagationSource Defines valid sources of request attributes for conditions or placeholders.
type PropagationSource struct {
	// Use only request attributes from services that belong to this management zone.. Use either this or `serviceTag`
	ManagementZone *string `json:"managementZone,omitempty"`
	ServiceTag *UniversalTag `json:"serviceTag,omitempty"`
}

// NewPropagationSource instantiates a new PropagationSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropagationSource() *PropagationSource {
	this := PropagationSource{}
	return &this
}

// NewPropagationSourceWithDefaults instantiates a new PropagationSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropagationSourceWithDefaults() *PropagationSource {
	this := PropagationSource{}
	return &this
}

// GetManagementZone returns the ManagementZone field value if set, zero value otherwise.
func (o *PropagationSource) GetManagementZone() string {
	if o == nil || o.ManagementZone == nil {
		var ret string
		return ret
	}
	return *o.ManagementZone
}

// GetManagementZoneOk returns a tuple with the ManagementZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationSource) GetManagementZoneOk() (*string, bool) {
	if o == nil || o.ManagementZone == nil {
		return nil, false
	}
	return o.ManagementZone, true
}

// HasManagementZone returns a boolean if a field has been set.
func (o *PropagationSource) HasManagementZone() bool {
	if o != nil && o.ManagementZone != nil {
		return true
	}

	return false
}

// SetManagementZone gets a reference to the given string and assigns it to the ManagementZone field.
func (o *PropagationSource) SetManagementZone(v string) {
	o.ManagementZone = &v
}

// GetServiceTag returns the ServiceTag field value if set, zero value otherwise.
func (o *PropagationSource) GetServiceTag() UniversalTag {
	if o == nil || o.ServiceTag == nil {
		var ret UniversalTag
		return ret
	}
	return *o.ServiceTag
}

// GetServiceTagOk returns a tuple with the ServiceTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationSource) GetServiceTagOk() (*UniversalTag, bool) {
	if o == nil || o.ServiceTag == nil {
		return nil, false
	}
	return o.ServiceTag, true
}

// HasServiceTag returns a boolean if a field has been set.
func (o *PropagationSource) HasServiceTag() bool {
	if o != nil && o.ServiceTag != nil {
		return true
	}

	return false
}

// SetServiceTag gets a reference to the given UniversalTag and assigns it to the ServiceTag field.
func (o *PropagationSource) SetServiceTag(v UniversalTag) {
	o.ServiceTag = &v
}

func (o PropagationSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManagementZone != nil {
		toSerialize["managementZone"] = o.ManagementZone
	}
	if o.ServiceTag != nil {
		toSerialize["serviceTag"] = o.ServiceTag
	}
	return json.Marshal(toSerialize)
}

type NullablePropagationSource struct {
	value *PropagationSource
	isSet bool
}

func (v NullablePropagationSource) Get() *PropagationSource {
	return v.value
}

func (v *NullablePropagationSource) Set(val *PropagationSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationSource(val *PropagationSource) *NullablePropagationSource {
	return &NullablePropagationSource{value: val, isSet: true}
}

func (v NullablePropagationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


