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

// MetaDataCapturing Configuration to capture meta data with the JavaScript agent.
type MetaDataCapturing struct {
	// The type of the meta data to capture.
	Type string `json:"type"`
	// The name of the meta data to capture.
	CapturingName string `json:"capturingName"`
	// Name for displaying the captured values in Dynatrace.
	Name string `json:"name"`
	// The unique id of the meta data to capture.
	UniqueId *int32 `json:"uniqueId,omitempty"`
	// True if this metadata should be captured regardless of the privacy settings
	PublicMetadata *bool `json:"publicMetadata,omitempty"`
}

// NewMetaDataCapturing instantiates a new MetaDataCapturing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaDataCapturing(type_ string, capturingName string, name string, ) *MetaDataCapturing {
	this := MetaDataCapturing{}
	this.Type = type_
	this.CapturingName = capturingName
	this.Name = name
	return &this
}

// NewMetaDataCapturingWithDefaults instantiates a new MetaDataCapturing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaDataCapturingWithDefaults() *MetaDataCapturing {
	this := MetaDataCapturing{}
	return &this
}

// GetType returns the Type field value
func (o *MetaDataCapturing) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetaDataCapturing) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetaDataCapturing) SetType(v string) {
	o.Type = v
}

// GetCapturingName returns the CapturingName field value
func (o *MetaDataCapturing) GetCapturingName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CapturingName
}

// GetCapturingNameOk returns a tuple with the CapturingName field value
// and a boolean to check if the value has been set.
func (o *MetaDataCapturing) GetCapturingNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CapturingName, true
}

// SetCapturingName sets field value
func (o *MetaDataCapturing) SetCapturingName(v string) {
	o.CapturingName = v
}

// GetName returns the Name field value
func (o *MetaDataCapturing) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetaDataCapturing) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetaDataCapturing) SetName(v string) {
	o.Name = v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *MetaDataCapturing) GetUniqueId() int32 {
	if o == nil || o.UniqueId == nil {
		var ret int32
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataCapturing) GetUniqueIdOk() (*int32, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *MetaDataCapturing) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given int32 and assigns it to the UniqueId field.
func (o *MetaDataCapturing) SetUniqueId(v int32) {
	o.UniqueId = &v
}

// GetPublicMetadata returns the PublicMetadata field value if set, zero value otherwise.
func (o *MetaDataCapturing) GetPublicMetadata() bool {
	if o == nil || o.PublicMetadata == nil {
		var ret bool
		return ret
	}
	return *o.PublicMetadata
}

// GetPublicMetadataOk returns a tuple with the PublicMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataCapturing) GetPublicMetadataOk() (*bool, bool) {
	if o == nil || o.PublicMetadata == nil {
		return nil, false
	}
	return o.PublicMetadata, true
}

// HasPublicMetadata returns a boolean if a field has been set.
func (o *MetaDataCapturing) HasPublicMetadata() bool {
	if o != nil && o.PublicMetadata != nil {
		return true
	}

	return false
}

// SetPublicMetadata gets a reference to the given bool and assigns it to the PublicMetadata field.
func (o *MetaDataCapturing) SetPublicMetadata(v bool) {
	o.PublicMetadata = &v
}

func (o MetaDataCapturing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["capturingName"] = o.CapturingName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UniqueId != nil {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if o.PublicMetadata != nil {
		toSerialize["publicMetadata"] = o.PublicMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableMetaDataCapturing struct {
	value *MetaDataCapturing
	isSet bool
}

func (v NullableMetaDataCapturing) Get() *MetaDataCapturing {
	return v.value
}

func (v *NullableMetaDataCapturing) Set(val *MetaDataCapturing) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaDataCapturing) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaDataCapturing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaDataCapturing(val *MetaDataCapturing) *NullableMetaDataCapturing {
	return &NullableMetaDataCapturing{value: val, isSet: true}
}

func (v NullableMetaDataCapturing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaDataCapturing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


