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

// SchemaFiles struct for SchemaFiles
type SchemaFiles struct {
	// A list of schema files.
	Files *[]string `json:"files,omitempty"`
}

// NewSchemaFiles instantiates a new SchemaFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaFiles() *SchemaFiles {
	this := SchemaFiles{}
	return &this
}

// NewSchemaFilesWithDefaults instantiates a new SchemaFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaFilesWithDefaults() *SchemaFiles {
	this := SchemaFiles{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SchemaFiles) GetFiles() []string {
	if o == nil || o.Files == nil {
		var ret []string
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaFiles) GetFilesOk() (*[]string, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SchemaFiles) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []string and assigns it to the Files field.
func (o *SchemaFiles) SetFiles(v []string) {
	o.Files = &v
}

func (o SchemaFiles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaFiles struct {
	value *SchemaFiles
	isSet bool
}

func (v NullableSchemaFiles) Get() *SchemaFiles {
	return v.value
}

func (v *NullableSchemaFiles) Set(val *SchemaFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaFiles(val *SchemaFiles) *NullableSchemaFiles {
	return &NullableSchemaFiles{value: val, isSet: true}
}

func (v NullableSchemaFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

