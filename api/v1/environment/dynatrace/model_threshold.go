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

// Threshold Parameters of a single plugin or custom event threshold.
type Threshold struct {
	// The ID of the threshold.
	ThresholdId *string `json:"thresholdId,omitempty"`
	// The case-sensitive ID of the metric evaluated by threshold.   You can use it to find the metric via the `timeseries` endpoint of the API.
	TimeseriesId *string `json:"timeseriesId,omitempty"`
	// The value of the threshold.
	Threshold *float64 `json:"threshold,omitempty"`
	// The condition for the threshold value check: above or below.
	AlertCondition *string `json:"alertCondition,omitempty"`
	// The number of one-minute samples that form sliding evaluation window.
	Samples *int32 `json:"samples,omitempty"`
	// How many one-minute samples within the evaluation window should violate the threshold to trigger an event.
	ViolatingSamples *int32 `json:"violatingSamples,omitempty"`
	// The type of the event to trigger on the threshold violation.
	EventType *string `json:"eventType,omitempty"`
	// The name of the event to trigger on the threshold violation.
	EventName *string `json:"eventName,omitempty"`
	// The source of the threshold.
	Filter *string `json:"filter,omitempty"`
	// The threshold is enabled/disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// How many one-minute samples within the evaluation window should go back to normal to close the event.
	DealertingSamples *int32 `json:"dealertingSamples,omitempty"`
	// A description of the event, triggered by a threshold violation.    You can use the following placeholders:  {severity}: the actual metric value,  {alert_condition}: above or below threshold condition,  {threshold}: the threshold value,{violating_samples}: the number of samples, violating the threshold,  {dimensions}: metric's dimension that violated the threshold.
	Description *string `json:"description,omitempty"`
}

// NewThreshold instantiates a new Threshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreshold() *Threshold {
	this := Threshold{}
	return &this
}

// NewThresholdWithDefaults instantiates a new Threshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdWithDefaults() *Threshold {
	this := Threshold{}
	return &this
}

// GetThresholdId returns the ThresholdId field value if set, zero value otherwise.
func (o *Threshold) GetThresholdId() string {
	if o == nil || o.ThresholdId == nil {
		var ret string
		return ret
	}
	return *o.ThresholdId
}

// GetThresholdIdOk returns a tuple with the ThresholdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetThresholdIdOk() (*string, bool) {
	if o == nil || o.ThresholdId == nil {
		return nil, false
	}
	return o.ThresholdId, true
}

// HasThresholdId returns a boolean if a field has been set.
func (o *Threshold) HasThresholdId() bool {
	if o != nil && o.ThresholdId != nil {
		return true
	}

	return false
}

// SetThresholdId gets a reference to the given string and assigns it to the ThresholdId field.
func (o *Threshold) SetThresholdId(v string) {
	o.ThresholdId = &v
}

// GetTimeseriesId returns the TimeseriesId field value if set, zero value otherwise.
func (o *Threshold) GetTimeseriesId() string {
	if o == nil || o.TimeseriesId == nil {
		var ret string
		return ret
	}
	return *o.TimeseriesId
}

// GetTimeseriesIdOk returns a tuple with the TimeseriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetTimeseriesIdOk() (*string, bool) {
	if o == nil || o.TimeseriesId == nil {
		return nil, false
	}
	return o.TimeseriesId, true
}

// HasTimeseriesId returns a boolean if a field has been set.
func (o *Threshold) HasTimeseriesId() bool {
	if o != nil && o.TimeseriesId != nil {
		return true
	}

	return false
}

// SetTimeseriesId gets a reference to the given string and assigns it to the TimeseriesId field.
func (o *Threshold) SetTimeseriesId(v string) {
	o.TimeseriesId = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Threshold) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Threshold) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *Threshold) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *Threshold) GetAlertCondition() string {
	if o == nil || o.AlertCondition == nil {
		var ret string
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetAlertConditionOk() (*string, bool) {
	if o == nil || o.AlertCondition == nil {
		return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *Threshold) HasAlertCondition() bool {
	if o != nil && o.AlertCondition != nil {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given string and assigns it to the AlertCondition field.
func (o *Threshold) SetAlertCondition(v string) {
	o.AlertCondition = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *Threshold) GetSamples() int32 {
	if o == nil || o.Samples == nil {
		var ret int32
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetSamplesOk() (*int32, bool) {
	if o == nil || o.Samples == nil {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *Threshold) HasSamples() bool {
	if o != nil && o.Samples != nil {
		return true
	}

	return false
}

// SetSamples gets a reference to the given int32 and assigns it to the Samples field.
func (o *Threshold) SetSamples(v int32) {
	o.Samples = &v
}

// GetViolatingSamples returns the ViolatingSamples field value if set, zero value otherwise.
func (o *Threshold) GetViolatingSamples() int32 {
	if o == nil || o.ViolatingSamples == nil {
		var ret int32
		return ret
	}
	return *o.ViolatingSamples
}

// GetViolatingSamplesOk returns a tuple with the ViolatingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetViolatingSamplesOk() (*int32, bool) {
	if o == nil || o.ViolatingSamples == nil {
		return nil, false
	}
	return o.ViolatingSamples, true
}

// HasViolatingSamples returns a boolean if a field has been set.
func (o *Threshold) HasViolatingSamples() bool {
	if o != nil && o.ViolatingSamples != nil {
		return true
	}

	return false
}

// SetViolatingSamples gets a reference to the given int32 and assigns it to the ViolatingSamples field.
func (o *Threshold) SetViolatingSamples(v int32) {
	o.ViolatingSamples = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *Threshold) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *Threshold) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *Threshold) SetEventType(v string) {
	o.EventType = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *Threshold) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *Threshold) HasEventName() bool {
	if o != nil && o.EventName != nil {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *Threshold) SetEventName(v string) {
	o.EventName = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *Threshold) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *Threshold) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *Threshold) SetFilter(v string) {
	o.Filter = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Threshold) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Threshold) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Threshold) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDealertingSamples returns the DealertingSamples field value if set, zero value otherwise.
func (o *Threshold) GetDealertingSamples() int32 {
	if o == nil || o.DealertingSamples == nil {
		var ret int32
		return ret
	}
	return *o.DealertingSamples
}

// GetDealertingSamplesOk returns a tuple with the DealertingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetDealertingSamplesOk() (*int32, bool) {
	if o == nil || o.DealertingSamples == nil {
		return nil, false
	}
	return o.DealertingSamples, true
}

// HasDealertingSamples returns a boolean if a field has been set.
func (o *Threshold) HasDealertingSamples() bool {
	if o != nil && o.DealertingSamples != nil {
		return true
	}

	return false
}

// SetDealertingSamples gets a reference to the given int32 and assigns it to the DealertingSamples field.
func (o *Threshold) SetDealertingSamples(v int32) {
	o.DealertingSamples = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Threshold) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Threshold) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Threshold) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Threshold) SetDescription(v string) {
	o.Description = &v
}

func (o Threshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThresholdId != nil {
		toSerialize["thresholdId"] = o.ThresholdId
	}
	if o.TimeseriesId != nil {
		toSerialize["timeseriesId"] = o.TimeseriesId
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.AlertCondition != nil {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if o.Samples != nil {
		toSerialize["samples"] = o.Samples
	}
	if o.ViolatingSamples != nil {
		toSerialize["violatingSamples"] = o.ViolatingSamples
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.DealertingSamples != nil {
		toSerialize["dealertingSamples"] = o.DealertingSamples
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableThreshold struct {
	value *Threshold
	isSet bool
}

func (v NullableThreshold) Get() *Threshold {
	return v.value
}

func (v *NullableThreshold) Set(val *Threshold) {
	v.value = val
	v.isSet = true
}

func (v NullableThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreshold(val *Threshold) *NullableThreshold {
	return &NullableThreshold{value: val, isSet: true}
}

func (v NullableThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

