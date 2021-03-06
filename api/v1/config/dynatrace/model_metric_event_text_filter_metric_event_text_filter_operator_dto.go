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

// MetricEventTextFilterMetricEventTextFilterOperatorDto A filter for a string value based on the given operator.
type MetricEventTextFilterMetricEventTextFilterOperatorDto struct {
	// The value to match on.
	Value string `json:"value"`
	// The operator to match on.
	Operator string `json:"operator"`
}

// NewMetricEventTextFilterMetricEventTextFilterOperatorDto instantiates a new MetricEventTextFilterMetricEventTextFilterOperatorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricEventTextFilterMetricEventTextFilterOperatorDto(value string, operator string, ) *MetricEventTextFilterMetricEventTextFilterOperatorDto {
	this := MetricEventTextFilterMetricEventTextFilterOperatorDto{}
	this.Value = value
	this.Operator = operator
	return &this
}

// NewMetricEventTextFilterMetricEventTextFilterOperatorDtoWithDefaults instantiates a new MetricEventTextFilterMetricEventTextFilterOperatorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricEventTextFilterMetricEventTextFilterOperatorDtoWithDefaults() *MetricEventTextFilterMetricEventTextFilterOperatorDto {
	this := MetricEventTextFilterMetricEventTextFilterOperatorDto{}
	return &this
}

// GetValue returns the Value field value
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) SetValue(v string) {
	o.Value = v
}

// GetOperator returns the Operator field value
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) GetOperator() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *MetricEventTextFilterMetricEventTextFilterOperatorDto) SetOperator(v string) {
	o.Operator = v
}

func (o MetricEventTextFilterMetricEventTextFilterOperatorDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	return json.Marshal(toSerialize)
}

type NullableMetricEventTextFilterMetricEventTextFilterOperatorDto struct {
	value *MetricEventTextFilterMetricEventTextFilterOperatorDto
	isSet bool
}

func (v NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) Get() *MetricEventTextFilterMetricEventTextFilterOperatorDto {
	return v.value
}

func (v *NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) Set(val *MetricEventTextFilterMetricEventTextFilterOperatorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricEventTextFilterMetricEventTextFilterOperatorDto(val *MetricEventTextFilterMetricEventTextFilterOperatorDto) *NullableMetricEventTextFilterMetricEventTextFilterOperatorDto {
	return &NullableMetricEventTextFilterMetricEventTextFilterOperatorDto{value: val, isSet: true}
}

func (v NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricEventTextFilterMetricEventTextFilterOperatorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


