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

// HostGroupAutoUpdateConfig Configuration of OneAgent auto-update in a host group.   Applies to all OneAgents installed on hosts of the host group if their **setting** parameter is set to `INHERITED`. Otherwise, the host level setting applies.
type HostGroupAutoUpdateConfig struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The Dynatrace entity ID of the host group.
	Id *string `json:"id,omitempty"`
	// The auto-update state of OneAgents in a host group:   * `ENABLED`: OneAgents automatically update to the most recent version.  * `DISABLED`: OneAgents update to the version specified in the **version** field. * `INHERITED`: The setting from the environment-wide configuration is used.  OneAgents installed on hosts of the host group use this configuration only when their **setting** parameter is set to `INHERITED`.
	Setting string `json:"setting"`
	// The version to which the OneAgent must be updated.   Specify the version in the `<major>.<minor>.<revision>` format (for example `1.181.0`). You can fetch the list of available versions with the [GET available versions](https://www.dynatrace.com/support/help/shortlink/api-deployment-get-versions) call. If no suitable installer is found for the provided version or the value is set to `null`, OneAgent won't be updated.   Only applicable when the **setting** parameter is set to `DISABLED`.
	Version *string `json:"version,omitempty"`
	// The actual state of the auto-update on the hosts in a host group.   Applicable only if the **setting** parameter is set to `INHERITED`. In that case the value is taken from the environment-wide setting.
	EffectiveSetting *string `json:"effectiveSetting,omitempty"`
	// The actual version to which the OneAgent must be updated.   Applicable only if the **setting** parameter is set to `INHERITED`. In that case the value is taken from the environment-wide setting.
	EffectiveVersion *string `json:"effectiveVersion,omitempty"`
}

// NewHostGroupAutoUpdateConfig instantiates a new HostGroupAutoUpdateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostGroupAutoUpdateConfig(setting string, ) *HostGroupAutoUpdateConfig {
	this := HostGroupAutoUpdateConfig{}
	this.Setting = setting
	return &this
}

// NewHostGroupAutoUpdateConfigWithDefaults instantiates a new HostGroupAutoUpdateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostGroupAutoUpdateConfigWithDefaults() *HostGroupAutoUpdateConfig {
	this := HostGroupAutoUpdateConfig{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HostGroupAutoUpdateConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HostGroupAutoUpdateConfig) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *HostGroupAutoUpdateConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostGroupAutoUpdateConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostGroupAutoUpdateConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostGroupAutoUpdateConfig) SetId(v string) {
	o.Id = &v
}

// GetSetting returns the Setting field value
func (o *HostGroupAutoUpdateConfig) GetSetting() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetSettingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *HostGroupAutoUpdateConfig) SetSetting(v string) {
	o.Setting = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HostGroupAutoUpdateConfig) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HostGroupAutoUpdateConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HostGroupAutoUpdateConfig) SetVersion(v string) {
	o.Version = &v
}

// GetEffectiveSetting returns the EffectiveSetting field value if set, zero value otherwise.
func (o *HostGroupAutoUpdateConfig) GetEffectiveSetting() string {
	if o == nil || o.EffectiveSetting == nil {
		var ret string
		return ret
	}
	return *o.EffectiveSetting
}

// GetEffectiveSettingOk returns a tuple with the EffectiveSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetEffectiveSettingOk() (*string, bool) {
	if o == nil || o.EffectiveSetting == nil {
		return nil, false
	}
	return o.EffectiveSetting, true
}

// HasEffectiveSetting returns a boolean if a field has been set.
func (o *HostGroupAutoUpdateConfig) HasEffectiveSetting() bool {
	if o != nil && o.EffectiveSetting != nil {
		return true
	}

	return false
}

// SetEffectiveSetting gets a reference to the given string and assigns it to the EffectiveSetting field.
func (o *HostGroupAutoUpdateConfig) SetEffectiveSetting(v string) {
	o.EffectiveSetting = &v
}

// GetEffectiveVersion returns the EffectiveVersion field value if set, zero value otherwise.
func (o *HostGroupAutoUpdateConfig) GetEffectiveVersion() string {
	if o == nil || o.EffectiveVersion == nil {
		var ret string
		return ret
	}
	return *o.EffectiveVersion
}

// GetEffectiveVersionOk returns a tuple with the EffectiveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupAutoUpdateConfig) GetEffectiveVersionOk() (*string, bool) {
	if o == nil || o.EffectiveVersion == nil {
		return nil, false
	}
	return o.EffectiveVersion, true
}

// HasEffectiveVersion returns a boolean if a field has been set.
func (o *HostGroupAutoUpdateConfig) HasEffectiveVersion() bool {
	if o != nil && o.EffectiveVersion != nil {
		return true
	}

	return false
}

// SetEffectiveVersion gets a reference to the given string and assigns it to the EffectiveVersion field.
func (o *HostGroupAutoUpdateConfig) SetEffectiveVersion(v string) {
	o.EffectiveVersion = &v
}

func (o HostGroupAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["setting"] = o.Setting
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.EffectiveSetting != nil {
		toSerialize["effectiveSetting"] = o.EffectiveSetting
	}
	if o.EffectiveVersion != nil {
		toSerialize["effectiveVersion"] = o.EffectiveVersion
	}
	return json.Marshal(toSerialize)
}

type NullableHostGroupAutoUpdateConfig struct {
	value *HostGroupAutoUpdateConfig
	isSet bool
}

func (v NullableHostGroupAutoUpdateConfig) Get() *HostGroupAutoUpdateConfig {
	return v.value
}

func (v *NullableHostGroupAutoUpdateConfig) Set(val *HostGroupAutoUpdateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHostGroupAutoUpdateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHostGroupAutoUpdateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostGroupAutoUpdateConfig(val *HostGroupAutoUpdateConfig) *NullableHostGroupAutoUpdateConfig {
	return &NullableHostGroupAutoUpdateConfig{value: val, isSet: true}
}

func (v NullableHostGroupAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostGroupAutoUpdateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


