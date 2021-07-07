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

// SchemaStub The short representation of the settings schema.
type SchemaStub struct {
	// The most recent version of the schema.
	LatestSchemaVersion *string `json:"latestSchemaVersion,omitempty"`
	// The ID of the schema.
	SchemaId *string `json:"schemaId,omitempty"`
	// The name of the schema.
	DisplayName *string `json:"displayName,omitempty"`
}

// NewSchemaStub instantiates a new SchemaStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaStub() *SchemaStub {
	this := SchemaStub{}
	return &this
}

// NewSchemaStubWithDefaults instantiates a new SchemaStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaStubWithDefaults() *SchemaStub {
	this := SchemaStub{}
	return &this
}

// GetLatestSchemaVersion returns the LatestSchemaVersion field value if set, zero value otherwise.
func (o *SchemaStub) GetLatestSchemaVersion() string {
	if o == nil || o.LatestSchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.LatestSchemaVersion
}

// GetLatestSchemaVersionOk returns a tuple with the LatestSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStub) GetLatestSchemaVersionOk() (*string, bool) {
	if o == nil || o.LatestSchemaVersion == nil {
		return nil, false
	}
	return o.LatestSchemaVersion, true
}

// HasLatestSchemaVersion returns a boolean if a field has been set.
func (o *SchemaStub) HasLatestSchemaVersion() bool {
	if o != nil && o.LatestSchemaVersion != nil {
		return true
	}

	return false
}

// SetLatestSchemaVersion gets a reference to the given string and assigns it to the LatestSchemaVersion field.
func (o *SchemaStub) SetLatestSchemaVersion(v string) {
	o.LatestSchemaVersion = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *SchemaStub) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStub) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *SchemaStub) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *SchemaStub) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchemaStub) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStub) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchemaStub) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchemaStub) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o SchemaStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LatestSchemaVersion != nil {
		toSerialize["latestSchemaVersion"] = o.LatestSchemaVersion
	}
	if o.SchemaId != nil {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaStub struct {
	value *SchemaStub
	isSet bool
}

func (v NullableSchemaStub) Get() *SchemaStub {
	return v.value
}

func (v *NullableSchemaStub) Set(val *SchemaStub) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaStub) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaStub(val *SchemaStub) *NullableSchemaStub {
	return &NullableSchemaStub{value: val, isSet: true}
}

func (v NullableSchemaStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

