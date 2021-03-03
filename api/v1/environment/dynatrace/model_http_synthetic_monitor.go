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

// HttpSyntheticMonitor HTTP synthetic monitor. Some fields are inherited from base `SyntheticMonitor` model.
type HttpSyntheticMonitor struct {
	SyntheticMonitor
	// A list of events for this monitor
	Requests *[]RequestDto `json:"requests,omitempty"`
}

// NewHttpSyntheticMonitor instantiates a new HttpSyntheticMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpSyntheticMonitor() *HttpSyntheticMonitor {
	this := HttpSyntheticMonitor{}
	return &this
}

// NewHttpSyntheticMonitorWithDefaults instantiates a new HttpSyntheticMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpSyntheticMonitorWithDefaults() *HttpSyntheticMonitor {
	this := HttpSyntheticMonitor{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *HttpSyntheticMonitor) GetRequests() []RequestDto {
	if o == nil || o.Requests == nil {
		var ret []RequestDto
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpSyntheticMonitor) GetRequestsOk() (*[]RequestDto, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *HttpSyntheticMonitor) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []RequestDto and assigns it to the Requests field.
func (o *HttpSyntheticMonitor) SetRequests(v []RequestDto) {
	o.Requests = &v
}

func (o HttpSyntheticMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntheticMonitor, errSyntheticMonitor := json.Marshal(o.SyntheticMonitor)
	if errSyntheticMonitor != nil {
		return []byte{}, errSyntheticMonitor
	}
	errSyntheticMonitor = json.Unmarshal([]byte(serializedSyntheticMonitor), &toSerialize)
	if errSyntheticMonitor != nil {
		return []byte{}, errSyntheticMonitor
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	return json.Marshal(toSerialize)
}

type NullableHttpSyntheticMonitor struct {
	value *HttpSyntheticMonitor
	isSet bool
}

func (v NullableHttpSyntheticMonitor) Get() *HttpSyntheticMonitor {
	return v.value
}

func (v *NullableHttpSyntheticMonitor) Set(val *HttpSyntheticMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpSyntheticMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpSyntheticMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpSyntheticMonitor(val *HttpSyntheticMonitor) *NullableHttpSyntheticMonitor {
	return &NullableHttpSyntheticMonitor{value: val, isSet: true}
}

func (v NullableHttpSyntheticMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpSyntheticMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

