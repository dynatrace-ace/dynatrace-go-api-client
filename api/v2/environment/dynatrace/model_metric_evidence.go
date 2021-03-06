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

// MetricEvidence The metric evidence of the problem.   A change of metric behavior that indicates the problem and/or is its root cause.
type MetricEvidence struct {
	Evidence
	// The ID of the metric.
	MetricId string `json:"metricId"`
	// The metric's value before the problem start.
	ValueBeforeChangePoint float32 `json:"valueBeforeChangePoint"`
	// The metric's value after the problem start.
	ValueAfterChangePoint float32 `json:"valueAfterChangePoint"`
	// The end time of the evidence, in UTC milliseconds.  The value `null` indicates that the evidence is still open.
	EndTime int64 `json:"endTime"`
	// The unit of the metric.
	Unit string `json:"unit"`
}

// NewMetricEvidence instantiates a new MetricEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricEvidence(metricId string, valueBeforeChangePoint float32, valueAfterChangePoint float32, endTime int64, unit string, evidenceType string, displayName string, entity EntityStub, rootCauseRelevant bool, startTime int64) *MetricEvidence {
	this := MetricEvidence{}
	this.EvidenceType = evidenceType
	this.DisplayName = displayName
	this.Entity = entity
	this.RootCauseRelevant = rootCauseRelevant
	this.StartTime = startTime
	this.MetricId = metricId
	this.ValueBeforeChangePoint = valueBeforeChangePoint
	this.ValueAfterChangePoint = valueAfterChangePoint
	this.EndTime = endTime
	this.Unit = unit
	return &this
}

// NewMetricEvidenceWithDefaults instantiates a new MetricEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricEvidenceWithDefaults() *MetricEvidence {
	this := MetricEvidence{}
	return &this
}

// GetMetricId returns the MetricId field value
func (o *MetricEvidence) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *MetricEvidence) GetMetricIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *MetricEvidence) SetMetricId(v string) {
	o.MetricId = v
}

// GetValueBeforeChangePoint returns the ValueBeforeChangePoint field value
func (o *MetricEvidence) GetValueBeforeChangePoint() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ValueBeforeChangePoint
}

// GetValueBeforeChangePointOk returns a tuple with the ValueBeforeChangePoint field value
// and a boolean to check if the value has been set.
func (o *MetricEvidence) GetValueBeforeChangePointOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValueBeforeChangePoint, true
}

// SetValueBeforeChangePoint sets field value
func (o *MetricEvidence) SetValueBeforeChangePoint(v float32) {
	o.ValueBeforeChangePoint = v
}

// GetValueAfterChangePoint returns the ValueAfterChangePoint field value
func (o *MetricEvidence) GetValueAfterChangePoint() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ValueAfterChangePoint
}

// GetValueAfterChangePointOk returns a tuple with the ValueAfterChangePoint field value
// and a boolean to check if the value has been set.
func (o *MetricEvidence) GetValueAfterChangePointOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValueAfterChangePoint, true
}

// SetValueAfterChangePoint sets field value
func (o *MetricEvidence) SetValueAfterChangePoint(v float32) {
	o.ValueAfterChangePoint = v
}

// GetEndTime returns the EndTime field value
func (o *MetricEvidence) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *MetricEvidence) GetEndTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *MetricEvidence) SetEndTime(v int64) {
	o.EndTime = v
}

// GetUnit returns the Unit field value
func (o *MetricEvidence) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *MetricEvidence) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *MetricEvidence) SetUnit(v string) {
	o.Unit = v
}

func (o MetricEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEvidence, errEvidence := json.Marshal(o.Evidence)
	if errEvidence != nil {
		return []byte{}, errEvidence
	}
	errEvidence = json.Unmarshal([]byte(serializedEvidence), &toSerialize)
	if errEvidence != nil {
		return []byte{}, errEvidence
	}
	if true {
		toSerialize["metricId"] = o.MetricId
	}
	if true {
		toSerialize["valueBeforeChangePoint"] = o.ValueBeforeChangePoint
	}
	if true {
		toSerialize["valueAfterChangePoint"] = o.ValueAfterChangePoint
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableMetricEvidence struct {
	value *MetricEvidence
	isSet bool
}

func (v NullableMetricEvidence) Get() *MetricEvidence {
	return v.value
}

func (v *NullableMetricEvidence) Set(val *MetricEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricEvidence(val *MetricEvidence) *NullableMetricEvidence {
	return &NullableMetricEvidence{value: val, isSet: true}
}

func (v NullableMetricEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


