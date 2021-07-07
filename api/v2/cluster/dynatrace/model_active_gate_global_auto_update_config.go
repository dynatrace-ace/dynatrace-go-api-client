/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster-wide functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// ActiveGateGlobalAutoUpdateConfig Global configuration of ActiveGates auto-update.
type ActiveGateGlobalAutoUpdateConfig struct {
	// The state of auto-updates for all ActiveGates connected to the environment or Managed cluster.   This setting is inherited by all ActiveGates that have the `INHERITED` setting.
	GlobalSetting string `json:"globalSetting"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
}

// NewActiveGateGlobalAutoUpdateConfig instantiates a new ActiveGateGlobalAutoUpdateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateGlobalAutoUpdateConfig(globalSetting string) *ActiveGateGlobalAutoUpdateConfig {
	this := ActiveGateGlobalAutoUpdateConfig{}
	this.GlobalSetting = globalSetting
	return &this
}

// NewActiveGateGlobalAutoUpdateConfigWithDefaults instantiates a new ActiveGateGlobalAutoUpdateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateGlobalAutoUpdateConfigWithDefaults() *ActiveGateGlobalAutoUpdateConfig {
	this := ActiveGateGlobalAutoUpdateConfig{}
	return &this
}

// GetGlobalSetting returns the GlobalSetting field value
func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalSetting
}

// GetGlobalSettingOk returns a tuple with the GlobalSetting field value
// and a boolean to check if the value has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSettingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GlobalSetting, true
}

// SetGlobalSetting sets field value
func (o *ActiveGateGlobalAutoUpdateConfig) SetGlobalSetting(v string) {
	o.GlobalSetting = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *ActiveGateGlobalAutoUpdateConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

func (o ActiveGateGlobalAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["globalSetting"] = o.GlobalSetting
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableActiveGateGlobalAutoUpdateConfig struct {
	value *ActiveGateGlobalAutoUpdateConfig
	isSet bool
}

func (v NullableActiveGateGlobalAutoUpdateConfig) Get() *ActiveGateGlobalAutoUpdateConfig {
	return v.value
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) Set(val *ActiveGateGlobalAutoUpdateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateGlobalAutoUpdateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateGlobalAutoUpdateConfig(val *ActiveGateGlobalAutoUpdateConfig) *NullableActiveGateGlobalAutoUpdateConfig {
	return &NullableActiveGateGlobalAutoUpdateConfig{value: val, isSet: true}
}

func (v NullableActiveGateGlobalAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

