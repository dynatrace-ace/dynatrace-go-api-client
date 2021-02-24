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

// OneAgentHostGroupConfig Configuration of OneAgent in a host group.
type OneAgentHostGroupConfig struct {
	// The Dynatrace entity ID of the host group.
	Id *string `json:"id,omitempty"`
	AutoUpdateConfig *HostGroupAutoUpdateConfig `json:"autoUpdateConfig,omitempty"`
}

// NewOneAgentHostGroupConfig instantiates a new OneAgentHostGroupConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneAgentHostGroupConfig() *OneAgentHostGroupConfig {
	this := OneAgentHostGroupConfig{}
	return &this
}

// NewOneAgentHostGroupConfigWithDefaults instantiates a new OneAgentHostGroupConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneAgentHostGroupConfigWithDefaults() *OneAgentHostGroupConfig {
	this := OneAgentHostGroupConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OneAgentHostGroupConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneAgentHostGroupConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OneAgentHostGroupConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OneAgentHostGroupConfig) SetId(v string) {
	o.Id = &v
}

// GetAutoUpdateConfig returns the AutoUpdateConfig field value if set, zero value otherwise.
func (o *OneAgentHostGroupConfig) GetAutoUpdateConfig() HostGroupAutoUpdateConfig {
	if o == nil || o.AutoUpdateConfig == nil {
		var ret HostGroupAutoUpdateConfig
		return ret
	}
	return *o.AutoUpdateConfig
}

// GetAutoUpdateConfigOk returns a tuple with the AutoUpdateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneAgentHostGroupConfig) GetAutoUpdateConfigOk() (*HostGroupAutoUpdateConfig, bool) {
	if o == nil || o.AutoUpdateConfig == nil {
		return nil, false
	}
	return o.AutoUpdateConfig, true
}

// HasAutoUpdateConfig returns a boolean if a field has been set.
func (o *OneAgentHostGroupConfig) HasAutoUpdateConfig() bool {
	if o != nil && o.AutoUpdateConfig != nil {
		return true
	}

	return false
}

// SetAutoUpdateConfig gets a reference to the given HostGroupAutoUpdateConfig and assigns it to the AutoUpdateConfig field.
func (o *OneAgentHostGroupConfig) SetAutoUpdateConfig(v HostGroupAutoUpdateConfig) {
	o.AutoUpdateConfig = &v
}

func (o OneAgentHostGroupConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AutoUpdateConfig != nil {
		toSerialize["autoUpdateConfig"] = o.AutoUpdateConfig
	}
	return json.Marshal(toSerialize)
}

type NullableOneAgentHostGroupConfig struct {
	value *OneAgentHostGroupConfig
	isSet bool
}

func (v NullableOneAgentHostGroupConfig) Get() *OneAgentHostGroupConfig {
	return v.value
}

func (v *NullableOneAgentHostGroupConfig) Set(val *OneAgentHostGroupConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOneAgentHostGroupConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOneAgentHostGroupConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneAgentHostGroupConfig(val *OneAgentHostGroupConfig) *NullableOneAgentHostGroupConfig {
	return &NullableOneAgentHostGroupConfig{value: val, isSet: true}
}

func (v NullableOneAgentHostGroupConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneAgentHostGroupConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


