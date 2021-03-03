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

// ConnectionInfo OneAgent connectivity information.
type ConnectionInfo struct {
	// The ID of your Dynatrace environment.
	TenantUUID *string `json:"tenantUUID,omitempty"`
	// The internal token that is used for authentication when OneAgent connects to the Dynatrace cluster to send data.
	TenantToken *string `json:"tenantToken,omitempty"`
	// The list of endpoints to connect to the Dynatrace environment. The list is sorted by endpoint priority, descending.
	CommunicationEndpoints *[]string `json:"communicationEndpoints,omitempty"`
}

// NewConnectionInfo instantiates a new ConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionInfo() *ConnectionInfo {
	this := ConnectionInfo{}
	return &this
}

// NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionInfoWithDefaults() *ConnectionInfo {
	this := ConnectionInfo{}
	return &this
}

// GetTenantUUID returns the TenantUUID field value if set, zero value otherwise.
func (o *ConnectionInfo) GetTenantUUID() string {
	if o == nil || o.TenantUUID == nil {
		var ret string
		return ret
	}
	return *o.TenantUUID
}

// GetTenantUUIDOk returns a tuple with the TenantUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetTenantUUIDOk() (*string, bool) {
	if o == nil || o.TenantUUID == nil {
		return nil, false
	}
	return o.TenantUUID, true
}

// HasTenantUUID returns a boolean if a field has been set.
func (o *ConnectionInfo) HasTenantUUID() bool {
	if o != nil && o.TenantUUID != nil {
		return true
	}

	return false
}

// SetTenantUUID gets a reference to the given string and assigns it to the TenantUUID field.
func (o *ConnectionInfo) SetTenantUUID(v string) {
	o.TenantUUID = &v
}

// GetTenantToken returns the TenantToken field value if set, zero value otherwise.
func (o *ConnectionInfo) GetTenantToken() string {
	if o == nil || o.TenantToken == nil {
		var ret string
		return ret
	}
	return *o.TenantToken
}

// GetTenantTokenOk returns a tuple with the TenantToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetTenantTokenOk() (*string, bool) {
	if o == nil || o.TenantToken == nil {
		return nil, false
	}
	return o.TenantToken, true
}

// HasTenantToken returns a boolean if a field has been set.
func (o *ConnectionInfo) HasTenantToken() bool {
	if o != nil && o.TenantToken != nil {
		return true
	}

	return false
}

// SetTenantToken gets a reference to the given string and assigns it to the TenantToken field.
func (o *ConnectionInfo) SetTenantToken(v string) {
	o.TenantToken = &v
}

// GetCommunicationEndpoints returns the CommunicationEndpoints field value if set, zero value otherwise.
func (o *ConnectionInfo) GetCommunicationEndpoints() []string {
	if o == nil || o.CommunicationEndpoints == nil {
		var ret []string
		return ret
	}
	return *o.CommunicationEndpoints
}

// GetCommunicationEndpointsOk returns a tuple with the CommunicationEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInfo) GetCommunicationEndpointsOk() (*[]string, bool) {
	if o == nil || o.CommunicationEndpoints == nil {
		return nil, false
	}
	return o.CommunicationEndpoints, true
}

// HasCommunicationEndpoints returns a boolean if a field has been set.
func (o *ConnectionInfo) HasCommunicationEndpoints() bool {
	if o != nil && o.CommunicationEndpoints != nil {
		return true
	}

	return false
}

// SetCommunicationEndpoints gets a reference to the given []string and assigns it to the CommunicationEndpoints field.
func (o *ConnectionInfo) SetCommunicationEndpoints(v []string) {
	o.CommunicationEndpoints = &v
}

func (o ConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantUUID != nil {
		toSerialize["tenantUUID"] = o.TenantUUID
	}
	if o.TenantToken != nil {
		toSerialize["tenantToken"] = o.TenantToken
	}
	if o.CommunicationEndpoints != nil {
		toSerialize["communicationEndpoints"] = o.CommunicationEndpoints
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionInfo struct {
	value *ConnectionInfo
	isSet bool
}

func (v NullableConnectionInfo) Get() *ConnectionInfo {
	return v.value
}

func (v *NullableConnectionInfo) Set(val *ConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionInfo(val *ConnectionInfo) *NullableConnectionInfo {
	return &NullableConnectionInfo{value: val, isSet: true}
}

func (v NullableConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

