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

// OutageHandlingPolicy Outage handling configuration.
type OutageHandlingPolicy struct {
	// When enabled (`true`), generate a problem and send an alert when the monitor is unavailable at all configured locations.
	GlobalOutage bool `json:"globalOutage"`
	// When enabled (`true`), generate a problem and send an alert when the monitor is unavailable for one or more consecutive runs at any location.
	LocalOutage bool `json:"localOutage"`
	LocalOutagePolicy LocalOutagePolicy `json:"localOutagePolicy"`
	// Schedule retry if browser monitor execution results in a fail. For HTTP monitors this property is ignored.
	RetryOnError *bool `json:"retryOnError,omitempty"`
}

// NewOutageHandlingPolicy instantiates a new OutageHandlingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutageHandlingPolicy(globalOutage bool, localOutage bool, localOutagePolicy LocalOutagePolicy) *OutageHandlingPolicy {
	this := OutageHandlingPolicy{}
	this.GlobalOutage = globalOutage
	this.LocalOutage = localOutage
	this.LocalOutagePolicy = localOutagePolicy
	var retryOnError bool = true
	this.RetryOnError = &retryOnError
	return &this
}

// NewOutageHandlingPolicyWithDefaults instantiates a new OutageHandlingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutageHandlingPolicyWithDefaults() *OutageHandlingPolicy {
	this := OutageHandlingPolicy{}
	var retryOnError bool = true
	this.RetryOnError = &retryOnError
	return &this
}

// GetGlobalOutage returns the GlobalOutage field value
func (o *OutageHandlingPolicy) GetGlobalOutage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GlobalOutage
}

// GetGlobalOutageOk returns a tuple with the GlobalOutage field value
// and a boolean to check if the value has been set.
func (o *OutageHandlingPolicy) GetGlobalOutageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GlobalOutage, true
}

// SetGlobalOutage sets field value
func (o *OutageHandlingPolicy) SetGlobalOutage(v bool) {
	o.GlobalOutage = v
}

// GetLocalOutage returns the LocalOutage field value
func (o *OutageHandlingPolicy) GetLocalOutage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LocalOutage
}

// GetLocalOutageOk returns a tuple with the LocalOutage field value
// and a boolean to check if the value has been set.
func (o *OutageHandlingPolicy) GetLocalOutageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocalOutage, true
}

// SetLocalOutage sets field value
func (o *OutageHandlingPolicy) SetLocalOutage(v bool) {
	o.LocalOutage = v
}

// GetLocalOutagePolicy returns the LocalOutagePolicy field value
func (o *OutageHandlingPolicy) GetLocalOutagePolicy() LocalOutagePolicy {
	if o == nil {
		var ret LocalOutagePolicy
		return ret
	}

	return o.LocalOutagePolicy
}

// GetLocalOutagePolicyOk returns a tuple with the LocalOutagePolicy field value
// and a boolean to check if the value has been set.
func (o *OutageHandlingPolicy) GetLocalOutagePolicyOk() (*LocalOutagePolicy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocalOutagePolicy, true
}

// SetLocalOutagePolicy sets field value
func (o *OutageHandlingPolicy) SetLocalOutagePolicy(v LocalOutagePolicy) {
	o.LocalOutagePolicy = v
}

// GetRetryOnError returns the RetryOnError field value if set, zero value otherwise.
func (o *OutageHandlingPolicy) GetRetryOnError() bool {
	if o == nil || o.RetryOnError == nil {
		var ret bool
		return ret
	}
	return *o.RetryOnError
}

// GetRetryOnErrorOk returns a tuple with the RetryOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutageHandlingPolicy) GetRetryOnErrorOk() (*bool, bool) {
	if o == nil || o.RetryOnError == nil {
		return nil, false
	}
	return o.RetryOnError, true
}

// HasRetryOnError returns a boolean if a field has been set.
func (o *OutageHandlingPolicy) HasRetryOnError() bool {
	if o != nil && o.RetryOnError != nil {
		return true
	}

	return false
}

// SetRetryOnError gets a reference to the given bool and assigns it to the RetryOnError field.
func (o *OutageHandlingPolicy) SetRetryOnError(v bool) {
	o.RetryOnError = &v
}

func (o OutageHandlingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["globalOutage"] = o.GlobalOutage
	}
	if true {
		toSerialize["localOutage"] = o.LocalOutage
	}
	if true {
		toSerialize["localOutagePolicy"] = o.LocalOutagePolicy
	}
	if o.RetryOnError != nil {
		toSerialize["retryOnError"] = o.RetryOnError
	}
	return json.Marshal(toSerialize)
}

type NullableOutageHandlingPolicy struct {
	value *OutageHandlingPolicy
	isSet bool
}

func (v NullableOutageHandlingPolicy) Get() *OutageHandlingPolicy {
	return v.value
}

func (v *NullableOutageHandlingPolicy) Set(val *OutageHandlingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOutageHandlingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOutageHandlingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutageHandlingPolicy(val *OutageHandlingPolicy) *NullableOutageHandlingPolicy {
	return &NullableOutageHandlingPolicy{value: val, isSet: true}
}

func (v NullableOutageHandlingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutageHandlingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


