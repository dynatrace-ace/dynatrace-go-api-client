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

// DatabaseConnectionFailureDetectionConfig Parameters of the failed database connections detection.   The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
type DatabaseConnectionFailureDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert.
	ConnectionFailsCount *int32 `json:"connectionFailsCount,omitempty"`
	// The *X* minutes time period during which the **connectionFailsCount** is evaluated.
	TimePeriodMinutes *int32 `json:"timePeriodMinutes,omitempty"`
}

// NewDatabaseConnectionFailureDetectionConfig instantiates a new DatabaseConnectionFailureDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConnectionFailureDetectionConfig(enabled bool, ) *DatabaseConnectionFailureDetectionConfig {
	this := DatabaseConnectionFailureDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewDatabaseConnectionFailureDetectionConfigWithDefaults instantiates a new DatabaseConnectionFailureDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConnectionFailureDetectionConfigWithDefaults() *DatabaseConnectionFailureDetectionConfig {
	this := DatabaseConnectionFailureDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DatabaseConnectionFailureDetectionConfig) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseConnectionFailureDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DatabaseConnectionFailureDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetConnectionFailsCount returns the ConnectionFailsCount field value if set, zero value otherwise.
func (o *DatabaseConnectionFailureDetectionConfig) GetConnectionFailsCount() int32 {
	if o == nil || o.ConnectionFailsCount == nil {
		var ret int32
		return ret
	}
	return *o.ConnectionFailsCount
}

// GetConnectionFailsCountOk returns a tuple with the ConnectionFailsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConnectionFailureDetectionConfig) GetConnectionFailsCountOk() (*int32, bool) {
	if o == nil || o.ConnectionFailsCount == nil {
		return nil, false
	}
	return o.ConnectionFailsCount, true
}

// HasConnectionFailsCount returns a boolean if a field has been set.
func (o *DatabaseConnectionFailureDetectionConfig) HasConnectionFailsCount() bool {
	if o != nil && o.ConnectionFailsCount != nil {
		return true
	}

	return false
}

// SetConnectionFailsCount gets a reference to the given int32 and assigns it to the ConnectionFailsCount field.
func (o *DatabaseConnectionFailureDetectionConfig) SetConnectionFailsCount(v int32) {
	o.ConnectionFailsCount = &v
}

// GetTimePeriodMinutes returns the TimePeriodMinutes field value if set, zero value otherwise.
func (o *DatabaseConnectionFailureDetectionConfig) GetTimePeriodMinutes() int32 {
	if o == nil || o.TimePeriodMinutes == nil {
		var ret int32
		return ret
	}
	return *o.TimePeriodMinutes
}

// GetTimePeriodMinutesOk returns a tuple with the TimePeriodMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConnectionFailureDetectionConfig) GetTimePeriodMinutesOk() (*int32, bool) {
	if o == nil || o.TimePeriodMinutes == nil {
		return nil, false
	}
	return o.TimePeriodMinutes, true
}

// HasTimePeriodMinutes returns a boolean if a field has been set.
func (o *DatabaseConnectionFailureDetectionConfig) HasTimePeriodMinutes() bool {
	if o != nil && o.TimePeriodMinutes != nil {
		return true
	}

	return false
}

// SetTimePeriodMinutes gets a reference to the given int32 and assigns it to the TimePeriodMinutes field.
func (o *DatabaseConnectionFailureDetectionConfig) SetTimePeriodMinutes(v int32) {
	o.TimePeriodMinutes = &v
}

func (o DatabaseConnectionFailureDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ConnectionFailsCount != nil {
		toSerialize["connectionFailsCount"] = o.ConnectionFailsCount
	}
	if o.TimePeriodMinutes != nil {
		toSerialize["timePeriodMinutes"] = o.TimePeriodMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseConnectionFailureDetectionConfig struct {
	value *DatabaseConnectionFailureDetectionConfig
	isSet bool
}

func (v NullableDatabaseConnectionFailureDetectionConfig) Get() *DatabaseConnectionFailureDetectionConfig {
	return v.value
}

func (v *NullableDatabaseConnectionFailureDetectionConfig) Set(val *DatabaseConnectionFailureDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConnectionFailureDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConnectionFailureDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConnectionFailureDetectionConfig(val *DatabaseConnectionFailureDetectionConfig) *NullableDatabaseConnectionFailureDetectionConfig {
	return &NullableDatabaseConnectionFailureDetectionConfig{value: val, isSet: true}
}

func (v NullableDatabaseConnectionFailureDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConnectionFailureDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


