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

// HttpSyntheticMonitorAllOf struct for HttpSyntheticMonitorAllOf
type HttpSyntheticMonitorAllOf struct {
	// A list of events for this monitor
	Requests *[]RequestDto `json:"requests,omitempty"`
}

// NewHttpSyntheticMonitorAllOf instantiates a new HttpSyntheticMonitorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpSyntheticMonitorAllOf() *HttpSyntheticMonitorAllOf {
	this := HttpSyntheticMonitorAllOf{}
	return &this
}

// NewHttpSyntheticMonitorAllOfWithDefaults instantiates a new HttpSyntheticMonitorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpSyntheticMonitorAllOfWithDefaults() *HttpSyntheticMonitorAllOf {
	this := HttpSyntheticMonitorAllOf{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *HttpSyntheticMonitorAllOf) GetRequests() []RequestDto {
	if o == nil || o.Requests == nil {
		var ret []RequestDto
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpSyntheticMonitorAllOf) GetRequestsOk() (*[]RequestDto, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *HttpSyntheticMonitorAllOf) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []RequestDto and assigns it to the Requests field.
func (o *HttpSyntheticMonitorAllOf) SetRequests(v []RequestDto) {
	o.Requests = &v
}

func (o HttpSyntheticMonitorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	return json.Marshal(toSerialize)
}

type NullableHttpSyntheticMonitorAllOf struct {
	value *HttpSyntheticMonitorAllOf
	isSet bool
}

func (v NullableHttpSyntheticMonitorAllOf) Get() *HttpSyntheticMonitorAllOf {
	return v.value
}

func (v *NullableHttpSyntheticMonitorAllOf) Set(val *HttpSyntheticMonitorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpSyntheticMonitorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpSyntheticMonitorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpSyntheticMonitorAllOf(val *HttpSyntheticMonitorAllOf) *NullableHttpSyntheticMonitorAllOf {
	return &NullableHttpSyntheticMonitorAllOf{value: val, isSet: true}
}

func (v NullableHttpSyntheticMonitorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpSyntheticMonitorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


