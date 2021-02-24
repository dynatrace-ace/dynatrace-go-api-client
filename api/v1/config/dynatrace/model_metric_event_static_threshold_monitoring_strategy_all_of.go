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

// MetricEventStaticThresholdMonitoringStrategyAllOf struct for MetricEventStaticThresholdMonitoringStrategyAllOf
type MetricEventStaticThresholdMonitoringStrategyAllOf struct {
	// The number of one-minute samples that form the sliding evaluation window.
	Samples *int32 `json:"samples,omitempty"`
	// The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event.
	ViolatingSamples *int32 `json:"violatingSamples,omitempty"`
	// The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	DealertingSamples *int32 `json:"dealertingSamples,omitempty"`
	// The condition for the **threshold** value check: above or below.
	AlertCondition *string `json:"alertCondition,omitempty"`
	// The value of the static threshold based on the specified unit.
	Threshold *float64 `json:"threshold,omitempty"`
	// The unit of the threshold, matching the metric definition.
	Unit *string `json:"unit,omitempty"`
}

// NewMetricEventStaticThresholdMonitoringStrategyAllOf instantiates a new MetricEventStaticThresholdMonitoringStrategyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricEventStaticThresholdMonitoringStrategyAllOf() *MetricEventStaticThresholdMonitoringStrategyAllOf {
	this := MetricEventStaticThresholdMonitoringStrategyAllOf{}
	return &this
}

// NewMetricEventStaticThresholdMonitoringStrategyAllOfWithDefaults instantiates a new MetricEventStaticThresholdMonitoringStrategyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricEventStaticThresholdMonitoringStrategyAllOfWithDefaults() *MetricEventStaticThresholdMonitoringStrategyAllOf {
	this := MetricEventStaticThresholdMonitoringStrategyAllOf{}
	return &this
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetSamples() int32 {
	if o == nil || o.Samples == nil {
		var ret int32
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetSamplesOk() (*int32, bool) {
	if o == nil || o.Samples == nil {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasSamples() bool {
	if o != nil && o.Samples != nil {
		return true
	}

	return false
}

// SetSamples gets a reference to the given int32 and assigns it to the Samples field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetSamples(v int32) {
	o.Samples = &v
}

// GetViolatingSamples returns the ViolatingSamples field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetViolatingSamples() int32 {
	if o == nil || o.ViolatingSamples == nil {
		var ret int32
		return ret
	}
	return *o.ViolatingSamples
}

// GetViolatingSamplesOk returns a tuple with the ViolatingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetViolatingSamplesOk() (*int32, bool) {
	if o == nil || o.ViolatingSamples == nil {
		return nil, false
	}
	return o.ViolatingSamples, true
}

// HasViolatingSamples returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasViolatingSamples() bool {
	if o != nil && o.ViolatingSamples != nil {
		return true
	}

	return false
}

// SetViolatingSamples gets a reference to the given int32 and assigns it to the ViolatingSamples field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetViolatingSamples(v int32) {
	o.ViolatingSamples = &v
}

// GetDealertingSamples returns the DealertingSamples field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetDealertingSamples() int32 {
	if o == nil || o.DealertingSamples == nil {
		var ret int32
		return ret
	}
	return *o.DealertingSamples
}

// GetDealertingSamplesOk returns a tuple with the DealertingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetDealertingSamplesOk() (*int32, bool) {
	if o == nil || o.DealertingSamples == nil {
		return nil, false
	}
	return o.DealertingSamples, true
}

// HasDealertingSamples returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasDealertingSamples() bool {
	if o != nil && o.DealertingSamples != nil {
		return true
	}

	return false
}

// SetDealertingSamples gets a reference to the given int32 and assigns it to the DealertingSamples field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetDealertingSamples(v int32) {
	o.DealertingSamples = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetAlertCondition() string {
	if o == nil || o.AlertCondition == nil {
		var ret string
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetAlertConditionOk() (*string, bool) {
	if o == nil || o.AlertCondition == nil {
		return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasAlertCondition() bool {
	if o != nil && o.AlertCondition != nil {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given string and assigns it to the AlertCondition field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetAlertCondition(v string) {
	o.AlertCondition = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricEventStaticThresholdMonitoringStrategyAllOf) SetUnit(v string) {
	o.Unit = &v
}

func (o MetricEventStaticThresholdMonitoringStrategyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Samples != nil {
		toSerialize["samples"] = o.Samples
	}
	if o.ViolatingSamples != nil {
		toSerialize["violatingSamples"] = o.ViolatingSamples
	}
	if o.DealertingSamples != nil {
		toSerialize["dealertingSamples"] = o.DealertingSamples
	}
	if o.AlertCondition != nil {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableMetricEventStaticThresholdMonitoringStrategyAllOf struct {
	value *MetricEventStaticThresholdMonitoringStrategyAllOf
	isSet bool
}

func (v NullableMetricEventStaticThresholdMonitoringStrategyAllOf) Get() *MetricEventStaticThresholdMonitoringStrategyAllOf {
	return v.value
}

func (v *NullableMetricEventStaticThresholdMonitoringStrategyAllOf) Set(val *MetricEventStaticThresholdMonitoringStrategyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricEventStaticThresholdMonitoringStrategyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricEventStaticThresholdMonitoringStrategyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricEventStaticThresholdMonitoringStrategyAllOf(val *MetricEventStaticThresholdMonitoringStrategyAllOf) *NullableMetricEventStaticThresholdMonitoringStrategyAllOf {
	return &NullableMetricEventStaticThresholdMonitoringStrategyAllOf{value: val, isSet: true}
}

func (v NullableMetricEventStaticThresholdMonitoringStrategyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricEventStaticThresholdMonitoringStrategyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


