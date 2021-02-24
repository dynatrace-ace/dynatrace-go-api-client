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

// AnomalyDetectionPG Configuration of anomaly detection for the process group.
type AnomalyDetectionPG struct {
	AvailabilityMonitoring *AvailabilityMonitoringPG `json:"availabilityMonitoring,omitempty"`
}

// NewAnomalyDetectionPG instantiates a new AnomalyDetectionPG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnomalyDetectionPG() *AnomalyDetectionPG {
	this := AnomalyDetectionPG{}
	return &this
}

// NewAnomalyDetectionPGWithDefaults instantiates a new AnomalyDetectionPG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnomalyDetectionPGWithDefaults() *AnomalyDetectionPG {
	this := AnomalyDetectionPG{}
	return &this
}

// GetAvailabilityMonitoring returns the AvailabilityMonitoring field value if set, zero value otherwise.
func (o *AnomalyDetectionPG) GetAvailabilityMonitoring() AvailabilityMonitoringPG {
	if o == nil || o.AvailabilityMonitoring == nil {
		var ret AvailabilityMonitoringPG
		return ret
	}
	return *o.AvailabilityMonitoring
}

// GetAvailabilityMonitoringOk returns a tuple with the AvailabilityMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyDetectionPG) GetAvailabilityMonitoringOk() (*AvailabilityMonitoringPG, bool) {
	if o == nil || o.AvailabilityMonitoring == nil {
		return nil, false
	}
	return o.AvailabilityMonitoring, true
}

// HasAvailabilityMonitoring returns a boolean if a field has been set.
func (o *AnomalyDetectionPG) HasAvailabilityMonitoring() bool {
	if o != nil && o.AvailabilityMonitoring != nil {
		return true
	}

	return false
}

// SetAvailabilityMonitoring gets a reference to the given AvailabilityMonitoringPG and assigns it to the AvailabilityMonitoring field.
func (o *AnomalyDetectionPG) SetAvailabilityMonitoring(v AvailabilityMonitoringPG) {
	o.AvailabilityMonitoring = &v
}

func (o AnomalyDetectionPG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityMonitoring != nil {
		toSerialize["availabilityMonitoring"] = o.AvailabilityMonitoring
	}
	return json.Marshal(toSerialize)
}

type NullableAnomalyDetectionPG struct {
	value *AnomalyDetectionPG
	isSet bool
}

func (v NullableAnomalyDetectionPG) Get() *AnomalyDetectionPG {
	return v.value
}

func (v *NullableAnomalyDetectionPG) Set(val *AnomalyDetectionPG) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomalyDetectionPG) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomalyDetectionPG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomalyDetectionPG(val *AnomalyDetectionPG) *NullableAnomalyDetectionPG {
	return &NullableAnomalyDetectionPG{value: val, isSet: true}
}

func (v NullableAnomalyDetectionPG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomalyDetectionPG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


