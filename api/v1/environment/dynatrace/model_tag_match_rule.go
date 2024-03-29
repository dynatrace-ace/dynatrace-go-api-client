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

// TagMatchRule The list of tags to be used for matching Dynatrace entities.
type TagMatchRule struct {
	// The list of types of the Dynatrace entities (for example hosts or services) you want to pick up by matching.
	MeTypes []string `json:"meTypes"`
	// The list of tags you want to use for matching. At least one tag is required.    You can use custom tags from the UI, imported tags, and tags based on environment variables.
	Tags []TagInfo `json:"tags"`
}

// NewTagMatchRule instantiates a new TagMatchRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagMatchRule(meTypes []string, tags []TagInfo) *TagMatchRule {
	this := TagMatchRule{}
	this.MeTypes = meTypes
	this.Tags = tags
	return &this
}

// NewTagMatchRuleWithDefaults instantiates a new TagMatchRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagMatchRuleWithDefaults() *TagMatchRule {
	this := TagMatchRule{}
	return &this
}

// GetMeTypes returns the MeTypes field value
func (o *TagMatchRule) GetMeTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MeTypes
}

// GetMeTypesOk returns a tuple with the MeTypes field value
// and a boolean to check if the value has been set.
func (o *TagMatchRule) GetMeTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MeTypes, true
}

// SetMeTypes sets field value
func (o *TagMatchRule) SetMeTypes(v []string) {
	o.MeTypes = v
}

// GetTags returns the Tags field value
func (o *TagMatchRule) GetTags() []TagInfo {
	if o == nil {
		var ret []TagInfo
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagMatchRule) GetTagsOk() (*[]TagInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *TagMatchRule) SetTags(v []TagInfo) {
	o.Tags = v
}

func (o TagMatchRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["meTypes"] = o.MeTypes
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTagMatchRule struct {
	value *TagMatchRule
	isSet bool
}

func (v NullableTagMatchRule) Get() *TagMatchRule {
	return v.value
}

func (v *NullableTagMatchRule) Set(val *TagMatchRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTagMatchRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTagMatchRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagMatchRule(val *TagMatchRule) *NullableTagMatchRule {
	return &NullableTagMatchRule{value: val, isSet: true}
}

func (v NullableTagMatchRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagMatchRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


