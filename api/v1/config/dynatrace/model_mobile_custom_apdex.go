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

// MobileCustomApdex Apdex configuration of a mobile or custom application.   A duration less than the **tolerable** threshold is considered satisfied.
type MobileCustomApdex struct {
	// Apdex **tolerable** threshold, in milliseconds: a duration greater than or equal to this value is considered tolerable.
	ToleratedThreshold int32 `json:"toleratedThreshold"`
	// Apdex **frustrated** threshold, in milliseconds: a duration greater than or equal to this value is considered frustrated.
	FrustratingThreshold int32 `json:"frustratingThreshold"`
	// Apdex error condition: if `true` the user session is considered frustrated when an error is reported.
	FrustratedOnError bool `json:"frustratedOnError"`
}

// NewMobileCustomApdex instantiates a new MobileCustomApdex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileCustomApdex(toleratedThreshold int32, frustratingThreshold int32, frustratedOnError bool, ) *MobileCustomApdex {
	this := MobileCustomApdex{}
	this.ToleratedThreshold = toleratedThreshold
	this.FrustratingThreshold = frustratingThreshold
	this.FrustratedOnError = frustratedOnError
	return &this
}

// NewMobileCustomApdexWithDefaults instantiates a new MobileCustomApdex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileCustomApdexWithDefaults() *MobileCustomApdex {
	this := MobileCustomApdex{}
	return &this
}

// GetToleratedThreshold returns the ToleratedThreshold field value
func (o *MobileCustomApdex) GetToleratedThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ToleratedThreshold
}

// GetToleratedThresholdOk returns a tuple with the ToleratedThreshold field value
// and a boolean to check if the value has been set.
func (o *MobileCustomApdex) GetToleratedThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToleratedThreshold, true
}

// SetToleratedThreshold sets field value
func (o *MobileCustomApdex) SetToleratedThreshold(v int32) {
	o.ToleratedThreshold = v
}

// GetFrustratingThreshold returns the FrustratingThreshold field value
func (o *MobileCustomApdex) GetFrustratingThreshold() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FrustratingThreshold
}

// GetFrustratingThresholdOk returns a tuple with the FrustratingThreshold field value
// and a boolean to check if the value has been set.
func (o *MobileCustomApdex) GetFrustratingThresholdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FrustratingThreshold, true
}

// SetFrustratingThreshold sets field value
func (o *MobileCustomApdex) SetFrustratingThreshold(v int32) {
	o.FrustratingThreshold = v
}

// GetFrustratedOnError returns the FrustratedOnError field value
func (o *MobileCustomApdex) GetFrustratedOnError() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.FrustratedOnError
}

// GetFrustratedOnErrorOk returns a tuple with the FrustratedOnError field value
// and a boolean to check if the value has been set.
func (o *MobileCustomApdex) GetFrustratedOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FrustratedOnError, true
}

// SetFrustratedOnError sets field value
func (o *MobileCustomApdex) SetFrustratedOnError(v bool) {
	o.FrustratedOnError = v
}

func (o MobileCustomApdex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["toleratedThreshold"] = o.ToleratedThreshold
	}
	if true {
		toSerialize["frustratingThreshold"] = o.FrustratingThreshold
	}
	if true {
		toSerialize["frustratedOnError"] = o.FrustratedOnError
	}
	return json.Marshal(toSerialize)
}

type NullableMobileCustomApdex struct {
	value *MobileCustomApdex
	isSet bool
}

func (v NullableMobileCustomApdex) Get() *MobileCustomApdex {
	return v.value
}

func (v *NullableMobileCustomApdex) Set(val *MobileCustomApdex) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileCustomApdex) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileCustomApdex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileCustomApdex(val *MobileCustomApdex) *NullableMobileCustomApdex {
	return &NullableMobileCustomApdex{value: val, isSet: true}
}

func (v NullableMobileCustomApdex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileCustomApdex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


