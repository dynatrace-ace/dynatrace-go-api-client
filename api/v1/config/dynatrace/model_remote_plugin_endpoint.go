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

// RemotePluginEndpoint Configuration of a plugin endpoint.
type RemotePluginEndpoint struct {
	// The ID of the endpoint.
	Id *string `json:"id,omitempty"`
	// The ID of the plugin to which the endpoint belongs.
	PluginId *string `json:"pluginId,omitempty"`
	// The name of the endpoint, displayed in Dynatrace.
	Name *string `json:"name,omitempty"`
	// The endpoint is enabled (`true`) or disabled (`false`).
	Enabled *bool `json:"enabled,omitempty"`
	// The list of endpoint parameters.    Each parameter is a property-value pair.
	Properties *map[string]string `json:"properties,omitempty"`
	ActiveGatePluginModule EntityShortRepresentation `json:"activeGatePluginModule"`
}

// NewRemotePluginEndpoint instantiates a new RemotePluginEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemotePluginEndpoint(activeGatePluginModule EntityShortRepresentation, ) *RemotePluginEndpoint {
	this := RemotePluginEndpoint{}
	this.ActiveGatePluginModule = activeGatePluginModule
	return &this
}

// NewRemotePluginEndpointWithDefaults instantiates a new RemotePluginEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemotePluginEndpointWithDefaults() *RemotePluginEndpoint {
	this := RemotePluginEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemotePluginEndpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemotePluginEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemotePluginEndpoint) SetId(v string) {
	o.Id = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *RemotePluginEndpoint) GetPluginId() string {
	if o == nil || o.PluginId == nil {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetPluginIdOk() (*string, bool) {
	if o == nil || o.PluginId == nil {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *RemotePluginEndpoint) HasPluginId() bool {
	if o != nil && o.PluginId != nil {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *RemotePluginEndpoint) SetPluginId(v string) {
	o.PluginId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemotePluginEndpoint) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemotePluginEndpoint) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemotePluginEndpoint) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RemotePluginEndpoint) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RemotePluginEndpoint) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RemotePluginEndpoint) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *RemotePluginEndpoint) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *RemotePluginEndpoint) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *RemotePluginEndpoint) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetActiveGatePluginModule returns the ActiveGatePluginModule field value
func (o *RemotePluginEndpoint) GetActiveGatePluginModule() EntityShortRepresentation {
	if o == nil  {
		var ret EntityShortRepresentation
		return ret
	}

	return o.ActiveGatePluginModule
}

// GetActiveGatePluginModuleOk returns a tuple with the ActiveGatePluginModule field value
// and a boolean to check if the value has been set.
func (o *RemotePluginEndpoint) GetActiveGatePluginModuleOk() (*EntityShortRepresentation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ActiveGatePluginModule, true
}

// SetActiveGatePluginModule sets field value
func (o *RemotePluginEndpoint) SetActiveGatePluginModule(v EntityShortRepresentation) {
	o.ActiveGatePluginModule = v
}

func (o RemotePluginEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PluginId != nil {
		toSerialize["pluginId"] = o.PluginId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["activeGatePluginModule"] = o.ActiveGatePluginModule
	}
	return json.Marshal(toSerialize)
}

type NullableRemotePluginEndpoint struct {
	value *RemotePluginEndpoint
	isSet bool
}

func (v NullableRemotePluginEndpoint) Get() *RemotePluginEndpoint {
	return v.value
}

func (v *NullableRemotePluginEndpoint) Set(val *RemotePluginEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableRemotePluginEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableRemotePluginEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemotePluginEndpoint(val *RemotePluginEndpoint) *NullableRemotePluginEndpoint {
	return &NullableRemotePluginEndpoint{value: val, isSet: true}
}

func (v NullableRemotePluginEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemotePluginEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


