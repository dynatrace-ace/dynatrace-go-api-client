/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// PluginInstance An instance of the OneAgent plugin.
type PluginInstance struct {
	// The version of the plugin.
	PluginVersion *string `json:"pluginVersion,omitempty"`
	// The state of the plugin instance.
	State *string `json:"state,omitempty"`
}

// NewPluginInstance instantiates a new PluginInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginInstance() *PluginInstance {
	this := PluginInstance{}
	return &this
}

// NewPluginInstanceWithDefaults instantiates a new PluginInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginInstanceWithDefaults() *PluginInstance {
	this := PluginInstance{}
	return &this
}

// GetPluginVersion returns the PluginVersion field value if set, zero value otherwise.
func (o *PluginInstance) GetPluginVersion() string {
	if o == nil || o.PluginVersion == nil {
		var ret string
		return ret
	}
	return *o.PluginVersion
}

// GetPluginVersionOk returns a tuple with the PluginVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstance) GetPluginVersionOk() (*string, bool) {
	if o == nil || o.PluginVersion == nil {
		return nil, false
	}
	return o.PluginVersion, true
}

// HasPluginVersion returns a boolean if a field has been set.
func (o *PluginInstance) HasPluginVersion() bool {
	if o != nil && o.PluginVersion != nil {
		return true
	}

	return false
}

// SetPluginVersion gets a reference to the given string and assigns it to the PluginVersion field.
func (o *PluginInstance) SetPluginVersion(v string) {
	o.PluginVersion = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PluginInstance) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstance) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PluginInstance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PluginInstance) SetState(v string) {
	o.State = &v
}

func (o PluginInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PluginVersion != nil {
		toSerialize["pluginVersion"] = o.PluginVersion
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullablePluginInstance struct {
	value *PluginInstance
	isSet bool
}

func (v NullablePluginInstance) Get() *PluginInstance {
	return v.value
}

func (v *NullablePluginInstance) Set(val *PluginInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginInstance(val *PluginInstance) *NullablePluginInstance {
	return &NullablePluginInstance{value: val, isSet: true}
}

func (v NullablePluginInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


