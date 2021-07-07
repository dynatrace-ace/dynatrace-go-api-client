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

// ActiveGateModule Information about ActiveGate module
type ActiveGateModule struct {
	// The module is misconfigured (`true`) or not (`false`).
	Misconfigured *bool `json:"misconfigured,omitempty"`
	// The version of the ActiveGate module.
	Version *string `json:"version,omitempty"`
	// The module is enabled (`true`) or disabled (`false`).
	Enabled *bool `json:"enabled,omitempty"`
	// The type of ActiveGate module.
	Type *string `json:"type,omitempty"`
	// The attributes of the ActiveGate module.
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// NewActiveGateModule instantiates a new ActiveGateModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateModule() *ActiveGateModule {
	this := ActiveGateModule{}
	return &this
}

// NewActiveGateModuleWithDefaults instantiates a new ActiveGateModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateModuleWithDefaults() *ActiveGateModule {
	this := ActiveGateModule{}
	return &this
}

// GetMisconfigured returns the Misconfigured field value if set, zero value otherwise.
func (o *ActiveGateModule) GetMisconfigured() bool {
	if o == nil || o.Misconfigured == nil {
		var ret bool
		return ret
	}
	return *o.Misconfigured
}

// GetMisconfiguredOk returns a tuple with the Misconfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateModule) GetMisconfiguredOk() (*bool, bool) {
	if o == nil || o.Misconfigured == nil {
		return nil, false
	}
	return o.Misconfigured, true
}

// HasMisconfigured returns a boolean if a field has been set.
func (o *ActiveGateModule) HasMisconfigured() bool {
	if o != nil && o.Misconfigured != nil {
		return true
	}

	return false
}

// SetMisconfigured gets a reference to the given bool and assigns it to the Misconfigured field.
func (o *ActiveGateModule) SetMisconfigured(v bool) {
	o.Misconfigured = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ActiveGateModule) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateModule) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ActiveGateModule) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ActiveGateModule) SetVersion(v string) {
	o.Version = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ActiveGateModule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateModule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ActiveGateModule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ActiveGateModule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActiveGateModule) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateModule) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActiveGateModule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActiveGateModule) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ActiveGateModule) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateModule) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ActiveGateModule) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ActiveGateModule) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o ActiveGateModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Misconfigured != nil {
		toSerialize["misconfigured"] = o.Misconfigured
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableActiveGateModule struct {
	value *ActiveGateModule
	isSet bool
}

func (v NullableActiveGateModule) Get() *ActiveGateModule {
	return v.value
}

func (v *NullableActiveGateModule) Set(val *ActiveGateModule) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateModule) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateModule(val *ActiveGateModule) *NullableActiveGateModule {
	return &NullableActiveGateModule{value: val, isSet: true}
}

func (v NullableActiveGateModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

