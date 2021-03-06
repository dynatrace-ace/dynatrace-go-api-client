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

// Log4host The list of available OS logs.
type Log4host struct {
	// The full path to the log.
	Path *string `json:"path,omitempty"`
	// The size of the log, bytes.
	Size *int64 `json:"size,omitempty"`
	// The log is available (`true`) or not available (`false`) for analysis.
	AvailableForAnalysis *bool `json:"availableForAnalysis,omitempty"`
}

// NewLog4host instantiates a new Log4host object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog4host() *Log4host {
	this := Log4host{}
	return &this
}

// NewLog4hostWithDefaults instantiates a new Log4host object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLog4hostWithDefaults() *Log4host {
	this := Log4host{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Log4host) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log4host) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Log4host) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Log4host) SetPath(v string) {
	o.Path = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Log4host) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log4host) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Log4host) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Log4host) SetSize(v int64) {
	o.Size = &v
}

// GetAvailableForAnalysis returns the AvailableForAnalysis field value if set, zero value otherwise.
func (o *Log4host) GetAvailableForAnalysis() bool {
	if o == nil || o.AvailableForAnalysis == nil {
		var ret bool
		return ret
	}
	return *o.AvailableForAnalysis
}

// GetAvailableForAnalysisOk returns a tuple with the AvailableForAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log4host) GetAvailableForAnalysisOk() (*bool, bool) {
	if o == nil || o.AvailableForAnalysis == nil {
		return nil, false
	}
	return o.AvailableForAnalysis, true
}

// HasAvailableForAnalysis returns a boolean if a field has been set.
func (o *Log4host) HasAvailableForAnalysis() bool {
	if o != nil && o.AvailableForAnalysis != nil {
		return true
	}

	return false
}

// SetAvailableForAnalysis gets a reference to the given bool and assigns it to the AvailableForAnalysis field.
func (o *Log4host) SetAvailableForAnalysis(v bool) {
	o.AvailableForAnalysis = &v
}

func (o Log4host) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.AvailableForAnalysis != nil {
		toSerialize["availableForAnalysis"] = o.AvailableForAnalysis
	}
	return json.Marshal(toSerialize)
}

type NullableLog4host struct {
	value *Log4host
	isSet bool
}

func (v NullableLog4host) Get() *Log4host {
	return v.value
}

func (v *NullableLog4host) Set(val *Log4host) {
	v.value = val
	v.isSet = true
}

func (v NullableLog4host) IsSet() bool {
	return v.isSet
}

func (v *NullableLog4host) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog4host(val *Log4host) *NullableLog4host {
	return &NullableLog4host{value: val, isSet: true}
}

func (v NullableLog4host) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog4host) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


