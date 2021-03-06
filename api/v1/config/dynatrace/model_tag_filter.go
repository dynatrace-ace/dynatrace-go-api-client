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

// TagFilter A tag-based filter of monitored entities.
type TagFilter struct {
	// The origin of the tag, such as AWS or Cloud Foundry.   Custom tags use the `CONTEXTLESS` value.
	Context string `json:"context"`
	// The key of the tag.   Custom tags have the tag value here.
	Key string `json:"key"`
	// The value of the tag.    Not applicable to custom tags.
	Value *string `json:"value,omitempty"`
}

// NewTagFilter instantiates a new TagFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagFilter(context string, key string, ) *TagFilter {
	this := TagFilter{}
	this.Context = context
	this.Key = key
	return &this
}

// NewTagFilterWithDefaults instantiates a new TagFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagFilterWithDefaults() *TagFilter {
	this := TagFilter{}
	return &this
}

// GetContext returns the Context field value
func (o *TagFilter) GetContext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *TagFilter) GetContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *TagFilter) SetContext(v string) {
	o.Context = v
}

// GetKey returns the Key field value
func (o *TagFilter) GetKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TagFilter) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TagFilter) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TagFilter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagFilter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TagFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TagFilter) SetValue(v string) {
	o.Value = &v
}

func (o TagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTagFilter struct {
	value *TagFilter
	isSet bool
}

func (v NullableTagFilter) Get() *TagFilter {
	return v.value
}

func (v *NullableTagFilter) Set(val *TagFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTagFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTagFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagFilter(val *TagFilter) *NullableTagFilter {
	return &NullableTagFilter{value: val, isSet: true}
}

func (v NullableTagFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


