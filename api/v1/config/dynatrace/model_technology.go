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

// Technology Configuration of technology monitoring.
type Technology struct {
	// The type of the technology.
	Type string `json:"type"`
	// The monitoring of the technology is enabled (`true`) or disabled (`false`).
	MonitoringEnabled bool `json:"monitoringEnabled"`
	// The validity of the configuration:   * `HOST`: The setting is valid for OneAgent on host only. Other OneAgents, connected to the same Dynatrace server may have different setting.  * `ENVIRONMENT`: The setting is valid for all OneAgents, connected to the Dynatrace server.
	Scope *string `json:"scope,omitempty"`
}

// NewTechnology instantiates a new Technology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechnology(type_ string, monitoringEnabled bool, ) *Technology {
	this := Technology{}
	this.Type = type_
	this.MonitoringEnabled = monitoringEnabled
	return &this
}

// NewTechnologyWithDefaults instantiates a new Technology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechnologyWithDefaults() *Technology {
	this := Technology{}
	return &this
}

// GetType returns the Type field value
func (o *Technology) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Technology) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Technology) SetType(v string) {
	o.Type = v
}

// GetMonitoringEnabled returns the MonitoringEnabled field value
func (o *Technology) GetMonitoringEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.MonitoringEnabled
}

// GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field value
// and a boolean to check if the value has been set.
func (o *Technology) GetMonitoringEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonitoringEnabled, true
}

// SetMonitoringEnabled sets field value
func (o *Technology) SetMonitoringEnabled(v bool) {
	o.MonitoringEnabled = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Technology) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Technology) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Technology) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Technology) SetScope(v string) {
	o.Scope = &v
}

func (o Technology) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["monitoringEnabled"] = o.MonitoringEnabled
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableTechnology struct {
	value *Technology
	isSet bool
}

func (v NullableTechnology) Get() *Technology {
	return v.value
}

func (v *NullableTechnology) Set(val *Technology) {
	v.value = val
	v.isSet = true
}

func (v NullableTechnology) IsSet() bool {
	return v.isSet
}

func (v *NullableTechnology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechnology(val *Technology) *NullableTechnology {
	return &NullableTechnology{value: val, isSet: true}
}

func (v NullableTechnology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechnology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


