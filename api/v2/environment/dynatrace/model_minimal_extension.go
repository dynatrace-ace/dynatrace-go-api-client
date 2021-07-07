/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// MinimalExtension A list of extensions.
type MinimalExtension struct {
	// Extension name
	ExtensionName *string `json:"extensionName,omitempty"`
	// Extension version
	Version *string `json:"version,omitempty"`
}

// NewMinimalExtension instantiates a new MinimalExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalExtension() *MinimalExtension {
	this := MinimalExtension{}
	return &this
}

// NewMinimalExtensionWithDefaults instantiates a new MinimalExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalExtensionWithDefaults() *MinimalExtension {
	this := MinimalExtension{}
	return &this
}

// GetExtensionName returns the ExtensionName field value if set, zero value otherwise.
func (o *MinimalExtension) GetExtensionName() string {
	if o == nil || o.ExtensionName == nil {
		var ret string
		return ret
	}
	return *o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalExtension) GetExtensionNameOk() (*string, bool) {
	if o == nil || o.ExtensionName == nil {
		return nil, false
	}
	return o.ExtensionName, true
}

// HasExtensionName returns a boolean if a field has been set.
func (o *MinimalExtension) HasExtensionName() bool {
	if o != nil && o.ExtensionName != nil {
		return true
	}

	return false
}

// SetExtensionName gets a reference to the given string and assigns it to the ExtensionName field.
func (o *MinimalExtension) SetExtensionName(v string) {
	o.ExtensionName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MinimalExtension) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalExtension) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MinimalExtension) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MinimalExtension) SetVersion(v string) {
	o.Version = &v
}

func (o MinimalExtension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtensionName != nil {
		toSerialize["extensionName"] = o.ExtensionName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableMinimalExtension struct {
	value *MinimalExtension
	isSet bool
}

func (v NullableMinimalExtension) Get() *MinimalExtension {
	return v.value
}

func (v *NullableMinimalExtension) Set(val *MinimalExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalExtension(val *MinimalExtension) *NullableMinimalExtension {
	return &NullableMinimalExtension{value: val, isSet: true}
}

func (v NullableMinimalExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

