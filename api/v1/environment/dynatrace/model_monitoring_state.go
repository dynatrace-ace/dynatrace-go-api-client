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

// MonitoringState Defines the current monitoring state of an entity.
type MonitoringState struct {
	// The current actual monitoring state on the entity.
	ActualMonitoringState *string `json:"actualMonitoringState,omitempty"`
	// The monitoring state that is expected from the configuration
	ExpectedMonitoringState *string `json:"expectedMonitoringState,omitempty"`
	// Defines whether or not the process has to restarted to enable monitoring
	RestartRequired *bool `json:"restartRequired,omitempty"`
}

// NewMonitoringState instantiates a new MonitoringState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringState() *MonitoringState {
	this := MonitoringState{}
	return &this
}

// NewMonitoringStateWithDefaults instantiates a new MonitoringState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringStateWithDefaults() *MonitoringState {
	this := MonitoringState{}
	return &this
}

// GetActualMonitoringState returns the ActualMonitoringState field value if set, zero value otherwise.
func (o *MonitoringState) GetActualMonitoringState() string {
	if o == nil || o.ActualMonitoringState == nil {
		var ret string
		return ret
	}
	return *o.ActualMonitoringState
}

// GetActualMonitoringStateOk returns a tuple with the ActualMonitoringState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringState) GetActualMonitoringStateOk() (*string, bool) {
	if o == nil || o.ActualMonitoringState == nil {
		return nil, false
	}
	return o.ActualMonitoringState, true
}

// HasActualMonitoringState returns a boolean if a field has been set.
func (o *MonitoringState) HasActualMonitoringState() bool {
	if o != nil && o.ActualMonitoringState != nil {
		return true
	}

	return false
}

// SetActualMonitoringState gets a reference to the given string and assigns it to the ActualMonitoringState field.
func (o *MonitoringState) SetActualMonitoringState(v string) {
	o.ActualMonitoringState = &v
}

// GetExpectedMonitoringState returns the ExpectedMonitoringState field value if set, zero value otherwise.
func (o *MonitoringState) GetExpectedMonitoringState() string {
	if o == nil || o.ExpectedMonitoringState == nil {
		var ret string
		return ret
	}
	return *o.ExpectedMonitoringState
}

// GetExpectedMonitoringStateOk returns a tuple with the ExpectedMonitoringState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringState) GetExpectedMonitoringStateOk() (*string, bool) {
	if o == nil || o.ExpectedMonitoringState == nil {
		return nil, false
	}
	return o.ExpectedMonitoringState, true
}

// HasExpectedMonitoringState returns a boolean if a field has been set.
func (o *MonitoringState) HasExpectedMonitoringState() bool {
	if o != nil && o.ExpectedMonitoringState != nil {
		return true
	}

	return false
}

// SetExpectedMonitoringState gets a reference to the given string and assigns it to the ExpectedMonitoringState field.
func (o *MonitoringState) SetExpectedMonitoringState(v string) {
	o.ExpectedMonitoringState = &v
}

// GetRestartRequired returns the RestartRequired field value if set, zero value otherwise.
func (o *MonitoringState) GetRestartRequired() bool {
	if o == nil || o.RestartRequired == nil {
		var ret bool
		return ret
	}
	return *o.RestartRequired
}

// GetRestartRequiredOk returns a tuple with the RestartRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringState) GetRestartRequiredOk() (*bool, bool) {
	if o == nil || o.RestartRequired == nil {
		return nil, false
	}
	return o.RestartRequired, true
}

// HasRestartRequired returns a boolean if a field has been set.
func (o *MonitoringState) HasRestartRequired() bool {
	if o != nil && o.RestartRequired != nil {
		return true
	}

	return false
}

// SetRestartRequired gets a reference to the given bool and assigns it to the RestartRequired field.
func (o *MonitoringState) SetRestartRequired(v bool) {
	o.RestartRequired = &v
}

func (o MonitoringState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualMonitoringState != nil {
		toSerialize["actualMonitoringState"] = o.ActualMonitoringState
	}
	if o.ExpectedMonitoringState != nil {
		toSerialize["expectedMonitoringState"] = o.ExpectedMonitoringState
	}
	if o.RestartRequired != nil {
		toSerialize["restartRequired"] = o.RestartRequired
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringState struct {
	value *MonitoringState
	isSet bool
}

func (v NullableMonitoringState) Get() *MonitoringState {
	return v.value
}

func (v *NullableMonitoringState) Set(val *MonitoringState) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringState) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringState(val *MonitoringState) *NullableMonitoringState {
	return &NullableMonitoringState{value: val, isSet: true}
}

func (v NullableMonitoringState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

