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

// ApiTokenUpdate The update of the API token.
type ApiTokenUpdate struct {
	// The name of the token.
	Name *string `json:"name,omitempty"`
	// The token is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty"`
}

// NewApiTokenUpdate instantiates a new ApiTokenUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokenUpdate() *ApiTokenUpdate {
	this := ApiTokenUpdate{}
	return &this
}

// NewApiTokenUpdateWithDefaults instantiates a new ApiTokenUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenUpdateWithDefaults() *ApiTokenUpdate {
	this := ApiTokenUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiTokenUpdate) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiTokenUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiTokenUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiTokenUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ApiTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableApiTokenUpdate struct {
	value *ApiTokenUpdate
	isSet bool
}

func (v NullableApiTokenUpdate) Get() *ApiTokenUpdate {
	return v.value
}

func (v *NullableApiTokenUpdate) Set(val *ApiTokenUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenUpdate(val *ApiTokenUpdate) *NullableApiTokenUpdate {
	return &NullableApiTokenUpdate{value: val, isSet: true}
}

func (v NullableApiTokenUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


