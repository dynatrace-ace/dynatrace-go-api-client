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

// AddEntityTags A list of tags to be added to monitored entities.
type AddEntityTags struct {
	// A list of tags to be added to monitored entities.
	Tags []AddEntityTag `json:"tags"`
}

// NewAddEntityTags instantiates a new AddEntityTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEntityTags(tags []AddEntityTag) *AddEntityTags {
	this := AddEntityTags{}
	this.Tags = tags
	return &this
}

// NewAddEntityTagsWithDefaults instantiates a new AddEntityTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEntityTagsWithDefaults() *AddEntityTags {
	this := AddEntityTags{}
	return &this
}

// GetTags returns the Tags field value
func (o *AddEntityTags) GetTags() []AddEntityTag {
	if o == nil {
		var ret []AddEntityTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *AddEntityTags) GetTagsOk() (*[]AddEntityTag, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *AddEntityTags) SetTags(v []AddEntityTag) {
	o.Tags = v
}

func (o AddEntityTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableAddEntityTags struct {
	value *AddEntityTags
	isSet bool
}

func (v NullableAddEntityTags) Get() *AddEntityTags {
	return v.value
}

func (v *NullableAddEntityTags) Set(val *AddEntityTags) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEntityTags) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEntityTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEntityTags(val *AddEntityTags) *NullableAddEntityTags {
	return &NullableAddEntityTags{value: val, isSet: true}
}

func (v NullableAddEntityTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEntityTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

