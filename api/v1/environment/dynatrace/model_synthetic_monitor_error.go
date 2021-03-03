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

// SyntheticMonitorError The error message of a synthetic monitor step.
type SyntheticMonitorError struct {
	// The error code.
	Code int32 `json:"code"`
	// The error message.
	Message string `json:"message"`
}

// NewSyntheticMonitorError instantiates a new SyntheticMonitorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticMonitorError(code int32, message string, ) *SyntheticMonitorError {
	this := SyntheticMonitorError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewSyntheticMonitorErrorWithDefaults instantiates a new SyntheticMonitorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticMonitorErrorWithDefaults() *SyntheticMonitorError {
	this := SyntheticMonitorError{}
	return &this
}

// GetCode returns the Code field value
func (o *SyntheticMonitorError) GetCode() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorError) GetCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SyntheticMonitorError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *SyntheticMonitorError) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SyntheticMonitorError) SetMessage(v string) {
	o.Message = v
}

func (o SyntheticMonitorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticMonitorError struct {
	value *SyntheticMonitorError
	isSet bool
}

func (v NullableSyntheticMonitorError) Get() *SyntheticMonitorError {
	return v.value
}

func (v *NullableSyntheticMonitorError) Set(val *SyntheticMonitorError) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticMonitorError) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticMonitorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticMonitorError(val *SyntheticMonitorError) *NullableSyntheticMonitorError {
	return &NullableSyntheticMonitorError{value: val, isSet: true}
}

func (v NullableSyntheticMonitorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticMonitorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

