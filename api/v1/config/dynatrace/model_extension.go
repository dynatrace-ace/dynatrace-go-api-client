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

// Extension General configuration of an extension.
type Extension struct {
	// The ID of the extension, for example `custom.remote.python.demo`.
	Id *string `json:"id,omitempty"`
	// The name of the extension, displayed in Dynatrace.
	Name *string `json:"name,omitempty"`
	// The version of the extension, displayed in Dynatrace.
	Version *string `json:"version,omitempty"`
	// The type of the extension. It indicates the runtime environment of the extension (for example, ACTIVEGATE).
	Type *string `json:"type,omitempty"`
	// The metricGroup of the extension used for grouping custom metrics into a hierarchical namespace.
	MetricGroup *string `json:"metricGroup,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// A list of extension properties.
	Properties *[]ExtensionProperty `json:"properties,omitempty"`
}

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension() *Extension {
	this := Extension{}
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Extension) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Extension) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Extension) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Extension) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Extension) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Extension) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Extension) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Extension) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Extension) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Extension) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Extension) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Extension) SetType(v string) {
	o.Type = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *Extension) GetMetricGroup() string {
	if o == nil || o.MetricGroup == nil {
		var ret string
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetMetricGroupOk() (*string, bool) {
	if o == nil || o.MetricGroup == nil {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *Extension) HasMetricGroup() bool {
	if o != nil && o.MetricGroup != nil {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given string and assigns it to the MetricGroup field.
func (o *Extension) SetMetricGroup(v string) {
	o.MetricGroup = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Extension) GetMetadata() ConfigurationMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Extension) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *Extension) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Extension) GetProperties() []ExtensionProperty {
	if o == nil || o.Properties == nil {
		var ret []ExtensionProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetPropertiesOk() (*[]ExtensionProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Extension) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ExtensionProperty and assigns it to the Properties field.
func (o *Extension) SetProperties(v []ExtensionProperty) {
	o.Properties = &v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MetricGroup != nil {
		toSerialize["metricGroup"] = o.MetricGroup
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


