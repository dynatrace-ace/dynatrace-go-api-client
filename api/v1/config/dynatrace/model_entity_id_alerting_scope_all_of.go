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

// EntityIdAlertingScopeAllOf struct for EntityIdAlertingScopeAllOf
type EntityIdAlertingScopeAllOf struct {
	// The monitored entities id to match on.
	EntityId *string `json:"entityId,omitempty"`
}

// NewEntityIdAlertingScopeAllOf instantiates a new EntityIdAlertingScopeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityIdAlertingScopeAllOf() *EntityIdAlertingScopeAllOf {
	this := EntityIdAlertingScopeAllOf{}
	return &this
}

// NewEntityIdAlertingScopeAllOfWithDefaults instantiates a new EntityIdAlertingScopeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityIdAlertingScopeAllOfWithDefaults() *EntityIdAlertingScopeAllOf {
	this := EntityIdAlertingScopeAllOf{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntityIdAlertingScopeAllOf) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityIdAlertingScopeAllOf) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntityIdAlertingScopeAllOf) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntityIdAlertingScopeAllOf) SetEntityId(v string) {
	o.EntityId = &v
}

func (o EntityIdAlertingScopeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	return json.Marshal(toSerialize)
}

type NullableEntityIdAlertingScopeAllOf struct {
	value *EntityIdAlertingScopeAllOf
	isSet bool
}

func (v NullableEntityIdAlertingScopeAllOf) Get() *EntityIdAlertingScopeAllOf {
	return v.value
}

func (v *NullableEntityIdAlertingScopeAllOf) Set(val *EntityIdAlertingScopeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityIdAlertingScopeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityIdAlertingScopeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityIdAlertingScopeAllOf(val *EntityIdAlertingScopeAllOf) *NullableEntityIdAlertingScopeAllOf {
	return &NullableEntityIdAlertingScopeAllOf{value: val, isSet: true}
}

func (v NullableEntityIdAlertingScopeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityIdAlertingScopeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


