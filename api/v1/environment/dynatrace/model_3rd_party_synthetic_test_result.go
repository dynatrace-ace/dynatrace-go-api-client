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

// Model3rdPartySyntheticTestResult The results of third-party synthetic monitor execution.
type Model3rdPartySyntheticTestResult struct {
	// The ID of the third-party synthetic monitor.
	Id string `json:"id"`
	ScheduleIntervalInSeconds *int32 `json:"scheduleIntervalInSeconds,omitempty"`
	// Number of steps in the monitor. Defaults to number of SyntheticTestSteps.
	TotalStepCount *int32 `json:"totalStepCount,omitempty"`
	// Results of third-party monitor executions per location.
	LocationResults []Model3rdPartySyntheticLocationTestResult `json:"locationResults"`
}

// NewModel3rdPartySyntheticTestResult instantiates a new Model3rdPartySyntheticTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartySyntheticTestResult(id string, locationResults []Model3rdPartySyntheticLocationTestResult) *Model3rdPartySyntheticTestResult {
	this := Model3rdPartySyntheticTestResult{}
	this.Id = id
	this.LocationResults = locationResults
	return &this
}

// NewModel3rdPartySyntheticTestResultWithDefaults instantiates a new Model3rdPartySyntheticTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartySyntheticTestResultWithDefaults() *Model3rdPartySyntheticTestResult {
	this := Model3rdPartySyntheticTestResult{}
	return &this
}

// GetId returns the Id field value
func (o *Model3rdPartySyntheticTestResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTestResult) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Model3rdPartySyntheticTestResult) SetId(v string) {
	o.Id = v
}

// GetScheduleIntervalInSeconds returns the ScheduleIntervalInSeconds field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticTestResult) GetScheduleIntervalInSeconds() int32 {
	if o == nil || o.ScheduleIntervalInSeconds == nil {
		var ret int32
		return ret
	}
	return *o.ScheduleIntervalInSeconds
}

// GetScheduleIntervalInSecondsOk returns a tuple with the ScheduleIntervalInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTestResult) GetScheduleIntervalInSecondsOk() (*int32, bool) {
	if o == nil || o.ScheduleIntervalInSeconds == nil {
		return nil, false
	}
	return o.ScheduleIntervalInSeconds, true
}

// HasScheduleIntervalInSeconds returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticTestResult) HasScheduleIntervalInSeconds() bool {
	if o != nil && o.ScheduleIntervalInSeconds != nil {
		return true
	}

	return false
}

// SetScheduleIntervalInSeconds gets a reference to the given int32 and assigns it to the ScheduleIntervalInSeconds field.
func (o *Model3rdPartySyntheticTestResult) SetScheduleIntervalInSeconds(v int32) {
	o.ScheduleIntervalInSeconds = &v
}

// GetTotalStepCount returns the TotalStepCount field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticTestResult) GetTotalStepCount() int32 {
	if o == nil || o.TotalStepCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalStepCount
}

// GetTotalStepCountOk returns a tuple with the TotalStepCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTestResult) GetTotalStepCountOk() (*int32, bool) {
	if o == nil || o.TotalStepCount == nil {
		return nil, false
	}
	return o.TotalStepCount, true
}

// HasTotalStepCount returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticTestResult) HasTotalStepCount() bool {
	if o != nil && o.TotalStepCount != nil {
		return true
	}

	return false
}

// SetTotalStepCount gets a reference to the given int32 and assigns it to the TotalStepCount field.
func (o *Model3rdPartySyntheticTestResult) SetTotalStepCount(v int32) {
	o.TotalStepCount = &v
}

// GetLocationResults returns the LocationResults field value
func (o *Model3rdPartySyntheticTestResult) GetLocationResults() []Model3rdPartySyntheticLocationTestResult {
	if o == nil {
		var ret []Model3rdPartySyntheticLocationTestResult
		return ret
	}

	return o.LocationResults
}

// GetLocationResultsOk returns a tuple with the LocationResults field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTestResult) GetLocationResultsOk() (*[]Model3rdPartySyntheticLocationTestResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocationResults, true
}

// SetLocationResults sets field value
func (o *Model3rdPartySyntheticTestResult) SetLocationResults(v []Model3rdPartySyntheticLocationTestResult) {
	o.LocationResults = v
}

func (o Model3rdPartySyntheticTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ScheduleIntervalInSeconds != nil {
		toSerialize["scheduleIntervalInSeconds"] = o.ScheduleIntervalInSeconds
	}
	if o.TotalStepCount != nil {
		toSerialize["totalStepCount"] = o.TotalStepCount
	}
	if true {
		toSerialize["locationResults"] = o.LocationResults
	}
	return json.Marshal(toSerialize)
}

type NullableModel3rdPartySyntheticTestResult struct {
	value *Model3rdPartySyntheticTestResult
	isSet bool
}

func (v NullableModel3rdPartySyntheticTestResult) Get() *Model3rdPartySyntheticTestResult {
	return v.value
}

func (v *NullableModel3rdPartySyntheticTestResult) Set(val *Model3rdPartySyntheticTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartySyntheticTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartySyntheticTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartySyntheticTestResult(val *Model3rdPartySyntheticTestResult) *NullableModel3rdPartySyntheticTestResult {
	return &NullableModel3rdPartySyntheticTestResult{value: val, isSet: true}
}

func (v NullableModel3rdPartySyntheticTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartySyntheticTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


