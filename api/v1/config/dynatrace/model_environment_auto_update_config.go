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

// EnvironmentAutoUpdateConfig Environment-wide configuration of OneAgents auto-updates.   Applies to all OneAgents connecting to the environment if their **setting** parameter is set to `INHERITED`. Otherwise, the host group or host level setting applies.
type EnvironmentAutoUpdateConfig struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The auto-update state of OneAgents connecting to the environment:   * `ENABLED`: OneAgents automatically update to the most recent version.  * `DISABLED`: OneAgents update to the version specified in the **version** field.  OneAgents that connect to the environment use this configuration only when their **setting** parameter is set to `INHERITED`.
	Setting string `json:"setting"`
	// The version to which the OneAgent must be updated.   Specify the version in the `<major>.<minor>.<revision>` format (for example `1.181.0`). You can fetch the list of available versions with the [GET available versions](https://www.dynatrace.com/support/help/shortlink/api-deployment-get-versions) call. If no suitable installer is found for the provided version or the value is set to `null`, OneAgent won't be updated.   Only applicable when the **setting** parameter is set to `DISABLED`.
	Version *string `json:"version,omitempty"`
}

// NewEnvironmentAutoUpdateConfig instantiates a new EnvironmentAutoUpdateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentAutoUpdateConfig(setting string, ) *EnvironmentAutoUpdateConfig {
	this := EnvironmentAutoUpdateConfig{}
	this.Setting = setting
	return &this
}

// NewEnvironmentAutoUpdateConfigWithDefaults instantiates a new EnvironmentAutoUpdateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentAutoUpdateConfigWithDefaults() *EnvironmentAutoUpdateConfig {
	this := EnvironmentAutoUpdateConfig{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EnvironmentAutoUpdateConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EnvironmentAutoUpdateConfig) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *EnvironmentAutoUpdateConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetSetting returns the Setting field value
func (o *EnvironmentAutoUpdateConfig) GetSetting() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *EnvironmentAutoUpdateConfig) GetSettingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *EnvironmentAutoUpdateConfig) SetSetting(v string) {
	o.Setting = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EnvironmentAutoUpdateConfig) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAutoUpdateConfig) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnvironmentAutoUpdateConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EnvironmentAutoUpdateConfig) SetVersion(v string) {
	o.Version = &v
}

func (o EnvironmentAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["setting"] = o.Setting
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentAutoUpdateConfig struct {
	value *EnvironmentAutoUpdateConfig
	isSet bool
}

func (v NullableEnvironmentAutoUpdateConfig) Get() *EnvironmentAutoUpdateConfig {
	return v.value
}

func (v *NullableEnvironmentAutoUpdateConfig) Set(val *EnvironmentAutoUpdateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentAutoUpdateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentAutoUpdateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentAutoUpdateConfig(val *EnvironmentAutoUpdateConfig) *NullableEnvironmentAutoUpdateConfig {
	return &NullableEnvironmentAutoUpdateConfig{value: val, isSet: true}
}

func (v NullableEnvironmentAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentAutoUpdateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


