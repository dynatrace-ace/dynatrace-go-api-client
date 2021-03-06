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

// AwsPrivateLinkConfig struct for AwsPrivateLinkConfig
type AwsPrivateLinkConfig struct {
	// Is AWS PrivateLink enabled
	Enabled bool `json:"enabled"`
	// The VirtualPrivateCluster-service name
	VpcEndpointServiceName *string `json:"vpcEndpointServiceName,omitempty"`
}

// NewAwsPrivateLinkConfig instantiates a new AwsPrivateLinkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsPrivateLinkConfig(enabled bool, ) *AwsPrivateLinkConfig {
	this := AwsPrivateLinkConfig{}
	this.Enabled = enabled
	return &this
}

// NewAwsPrivateLinkConfigWithDefaults instantiates a new AwsPrivateLinkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsPrivateLinkConfigWithDefaults() *AwsPrivateLinkConfig {
	this := AwsPrivateLinkConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AwsPrivateLinkConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AwsPrivateLinkConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AwsPrivateLinkConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetVpcEndpointServiceName returns the VpcEndpointServiceName field value if set, zero value otherwise.
func (o *AwsPrivateLinkConfig) GetVpcEndpointServiceName() string {
	if o == nil || o.VpcEndpointServiceName == nil {
		var ret string
		return ret
	}
	return *o.VpcEndpointServiceName
}

// GetVpcEndpointServiceNameOk returns a tuple with the VpcEndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsPrivateLinkConfig) GetVpcEndpointServiceNameOk() (*string, bool) {
	if o == nil || o.VpcEndpointServiceName == nil {
		return nil, false
	}
	return o.VpcEndpointServiceName, true
}

// HasVpcEndpointServiceName returns a boolean if a field has been set.
func (o *AwsPrivateLinkConfig) HasVpcEndpointServiceName() bool {
	if o != nil && o.VpcEndpointServiceName != nil {
		return true
	}

	return false
}

// SetVpcEndpointServiceName gets a reference to the given string and assigns it to the VpcEndpointServiceName field.
func (o *AwsPrivateLinkConfig) SetVpcEndpointServiceName(v string) {
	o.VpcEndpointServiceName = &v
}

func (o AwsPrivateLinkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.VpcEndpointServiceName != nil {
		toSerialize["vpcEndpointServiceName"] = o.VpcEndpointServiceName
	}
	return json.Marshal(toSerialize)
}

type NullableAwsPrivateLinkConfig struct {
	value *AwsPrivateLinkConfig
	isSet bool
}

func (v NullableAwsPrivateLinkConfig) Get() *AwsPrivateLinkConfig {
	return v.value
}

func (v *NullableAwsPrivateLinkConfig) Set(val *AwsPrivateLinkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsPrivateLinkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsPrivateLinkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsPrivateLinkConfig(val *AwsPrivateLinkConfig) *NullableAwsPrivateLinkConfig {
	return &NullableAwsPrivateLinkConfig{value: val, isSet: true}
}

func (v NullableAwsPrivateLinkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsPrivateLinkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


