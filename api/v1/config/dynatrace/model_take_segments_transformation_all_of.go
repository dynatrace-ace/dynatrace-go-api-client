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

// TakeSegmentsTransformationAllOf struct for TakeSegmentsTransformationAllOf
type TakeSegmentsTransformationAllOf struct {
	// The number of elements to be kept.
	SegmentCount *int32 `json:"segmentCount,omitempty"`
	// The delimiter for splitting the detected value. The delimiter itself is not kept.
	Delimiter *string `json:"delimiter,omitempty"`
	// Keeps the first (`false`) or last (`true`) elements.    If not set, then `false` is used, keeping the first elements.
	TakeFromEnd *bool `json:"takeFromEnd,omitempty"`
}

// NewTakeSegmentsTransformationAllOf instantiates a new TakeSegmentsTransformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTakeSegmentsTransformationAllOf() *TakeSegmentsTransformationAllOf {
	this := TakeSegmentsTransformationAllOf{}
	return &this
}

// NewTakeSegmentsTransformationAllOfWithDefaults instantiates a new TakeSegmentsTransformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTakeSegmentsTransformationAllOfWithDefaults() *TakeSegmentsTransformationAllOf {
	this := TakeSegmentsTransformationAllOf{}
	return &this
}

// GetSegmentCount returns the SegmentCount field value if set, zero value otherwise.
func (o *TakeSegmentsTransformationAllOf) GetSegmentCount() int32 {
	if o == nil || o.SegmentCount == nil {
		var ret int32
		return ret
	}
	return *o.SegmentCount
}

// GetSegmentCountOk returns a tuple with the SegmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakeSegmentsTransformationAllOf) GetSegmentCountOk() (*int32, bool) {
	if o == nil || o.SegmentCount == nil {
		return nil, false
	}
	return o.SegmentCount, true
}

// HasSegmentCount returns a boolean if a field has been set.
func (o *TakeSegmentsTransformationAllOf) HasSegmentCount() bool {
	if o != nil && o.SegmentCount != nil {
		return true
	}

	return false
}

// SetSegmentCount gets a reference to the given int32 and assigns it to the SegmentCount field.
func (o *TakeSegmentsTransformationAllOf) SetSegmentCount(v int32) {
	o.SegmentCount = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *TakeSegmentsTransformationAllOf) GetDelimiter() string {
	if o == nil || o.Delimiter == nil {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakeSegmentsTransformationAllOf) GetDelimiterOk() (*string, bool) {
	if o == nil || o.Delimiter == nil {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *TakeSegmentsTransformationAllOf) HasDelimiter() bool {
	if o != nil && o.Delimiter != nil {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *TakeSegmentsTransformationAllOf) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetTakeFromEnd returns the TakeFromEnd field value if set, zero value otherwise.
func (o *TakeSegmentsTransformationAllOf) GetTakeFromEnd() bool {
	if o == nil || o.TakeFromEnd == nil {
		var ret bool
		return ret
	}
	return *o.TakeFromEnd
}

// GetTakeFromEndOk returns a tuple with the TakeFromEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TakeSegmentsTransformationAllOf) GetTakeFromEndOk() (*bool, bool) {
	if o == nil || o.TakeFromEnd == nil {
		return nil, false
	}
	return o.TakeFromEnd, true
}

// HasTakeFromEnd returns a boolean if a field has been set.
func (o *TakeSegmentsTransformationAllOf) HasTakeFromEnd() bool {
	if o != nil && o.TakeFromEnd != nil {
		return true
	}

	return false
}

// SetTakeFromEnd gets a reference to the given bool and assigns it to the TakeFromEnd field.
func (o *TakeSegmentsTransformationAllOf) SetTakeFromEnd(v bool) {
	o.TakeFromEnd = &v
}

func (o TakeSegmentsTransformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SegmentCount != nil {
		toSerialize["segmentCount"] = o.SegmentCount
	}
	if o.Delimiter != nil {
		toSerialize["delimiter"] = o.Delimiter
	}
	if o.TakeFromEnd != nil {
		toSerialize["takeFromEnd"] = o.TakeFromEnd
	}
	return json.Marshal(toSerialize)
}

type NullableTakeSegmentsTransformationAllOf struct {
	value *TakeSegmentsTransformationAllOf
	isSet bool
}

func (v NullableTakeSegmentsTransformationAllOf) Get() *TakeSegmentsTransformationAllOf {
	return v.value
}

func (v *NullableTakeSegmentsTransformationAllOf) Set(val *TakeSegmentsTransformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTakeSegmentsTransformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTakeSegmentsTransformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTakeSegmentsTransformationAllOf(val *TakeSegmentsTransformationAllOf) *NullableTakeSegmentsTransformationAllOf {
	return &NullableTakeSegmentsTransformationAllOf{value: val, isSet: true}
}

func (v NullableTakeSegmentsTransformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakeSegmentsTransformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


