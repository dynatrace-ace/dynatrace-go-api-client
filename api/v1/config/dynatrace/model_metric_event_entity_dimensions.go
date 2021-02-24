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

// MetricEventEntityDimensions A filter for the metrics entity dimensions.
type MetricEventEntityDimensions struct {
	MetricEventDimensions
	NameFilter MetricEventTextFilterMetricEventDimensionsFilterOperatorDto `json:"nameFilter"`
}

// NewMetricEventEntityDimensions instantiates a new MetricEventEntityDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricEventEntityDimensions(nameFilter MetricEventTextFilterMetricEventDimensionsFilterOperatorDto, ) *MetricEventEntityDimensions {
	this := MetricEventEntityDimensions{}
	this.NameFilter = nameFilter
	return &this
}

// NewMetricEventEntityDimensionsWithDefaults instantiates a new MetricEventEntityDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricEventEntityDimensionsWithDefaults() *MetricEventEntityDimensions {
	this := MetricEventEntityDimensions{}
	return &this
}

// GetNameFilter returns the NameFilter field value
func (o *MetricEventEntityDimensions) GetNameFilter() MetricEventTextFilterMetricEventDimensionsFilterOperatorDto {
	if o == nil  {
		var ret MetricEventTextFilterMetricEventDimensionsFilterOperatorDto
		return ret
	}

	return o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value
// and a boolean to check if the value has been set.
func (o *MetricEventEntityDimensions) GetNameFilterOk() (*MetricEventTextFilterMetricEventDimensionsFilterOperatorDto, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameFilter, true
}

// SetNameFilter sets field value
func (o *MetricEventEntityDimensions) SetNameFilter(v MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) {
	o.NameFilter = v
}

func (o MetricEventEntityDimensions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMetricEventDimensions, errMetricEventDimensions := json.Marshal(o.MetricEventDimensions)
	if errMetricEventDimensions != nil {
		return []byte{}, errMetricEventDimensions
	}
	errMetricEventDimensions = json.Unmarshal([]byte(serializedMetricEventDimensions), &toSerialize)
	if errMetricEventDimensions != nil {
		return []byte{}, errMetricEventDimensions
	}
	if true {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return json.Marshal(toSerialize)
}

type NullableMetricEventEntityDimensions struct {
	value *MetricEventEntityDimensions
	isSet bool
}

func (v NullableMetricEventEntityDimensions) Get() *MetricEventEntityDimensions {
	return v.value
}

func (v *NullableMetricEventEntityDimensions) Set(val *MetricEventEntityDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricEventEntityDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricEventEntityDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricEventEntityDimensions(val *MetricEventEntityDimensions) *NullableMetricEventEntityDimensions {
	return &NullableMetricEventEntityDimensions{value: val, isSet: true}
}

func (v NullableMetricEventEntityDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricEventEntityDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


