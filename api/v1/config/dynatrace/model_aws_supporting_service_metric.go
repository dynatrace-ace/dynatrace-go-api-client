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

// AwsSupportingServiceMetric A metric of supporting service to be monitored.
type AwsSupportingServiceMetric struct {
	// The name of the metric of the supporting service.
	Name string `json:"name"`
	// The statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM
	Statistic string `json:"statistic"`
	// A list of metric's dimensions names.
	Dimensions []string `json:"dimensions"`
}

// NewAwsSupportingServiceMetric instantiates a new AwsSupportingServiceMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSupportingServiceMetric(name string, statistic string, dimensions []string, ) *AwsSupportingServiceMetric {
	this := AwsSupportingServiceMetric{}
	this.Name = name
	this.Statistic = statistic
	this.Dimensions = dimensions
	return &this
}

// NewAwsSupportingServiceMetricWithDefaults instantiates a new AwsSupportingServiceMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSupportingServiceMetricWithDefaults() *AwsSupportingServiceMetric {
	this := AwsSupportingServiceMetric{}
	return &this
}

// GetName returns the Name field value
func (o *AwsSupportingServiceMetric) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsSupportingServiceMetric) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsSupportingServiceMetric) SetName(v string) {
	o.Name = v
}

// GetStatistic returns the Statistic field value
func (o *AwsSupportingServiceMetric) GetStatistic() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Statistic
}

// GetStatisticOk returns a tuple with the Statistic field value
// and a boolean to check if the value has been set.
func (o *AwsSupportingServiceMetric) GetStatisticOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Statistic, true
}

// SetStatistic sets field value
func (o *AwsSupportingServiceMetric) SetStatistic(v string) {
	o.Statistic = v
}

// GetDimensions returns the Dimensions field value
func (o *AwsSupportingServiceMetric) GetDimensions() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *AwsSupportingServiceMetric) GetDimensionsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *AwsSupportingServiceMetric) SetDimensions(v []string) {
	o.Dimensions = v
}

func (o AwsSupportingServiceMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["statistic"] = o.Statistic
	}
	if true {
		toSerialize["dimensions"] = o.Dimensions
	}
	return json.Marshal(toSerialize)
}

type NullableAwsSupportingServiceMetric struct {
	value *AwsSupportingServiceMetric
	isSet bool
}

func (v NullableAwsSupportingServiceMetric) Get() *AwsSupportingServiceMetric {
	return v.value
}

func (v *NullableAwsSupportingServiceMetric) Set(val *AwsSupportingServiceMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSupportingServiceMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSupportingServiceMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSupportingServiceMetric(val *AwsSupportingServiceMetric) *NullableAwsSupportingServiceMetric {
	return &NullableAwsSupportingServiceMetric{value: val, isSet: true}
}

func (v NullableAwsSupportingServiceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSupportingServiceMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


