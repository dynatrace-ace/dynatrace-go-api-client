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

// LogJobRecordsResult The results of the log analysis job.
type LogJobRecordsResult struct {
	// The list of log analysis results.    The last page contains empty list.
	Records *[]LogRecord `json:"records,omitempty"`
	// The *scroll token* for the next page of results.    Without it you'll get the first page again.
	ScrollToken *string `json:"scrollToken,omitempty"`
}

// NewLogJobRecordsResult instantiates a new LogJobRecordsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogJobRecordsResult() *LogJobRecordsResult {
	this := LogJobRecordsResult{}
	return &this
}

// NewLogJobRecordsResultWithDefaults instantiates a new LogJobRecordsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogJobRecordsResultWithDefaults() *LogJobRecordsResult {
	this := LogJobRecordsResult{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *LogJobRecordsResult) GetRecords() []LogRecord {
	if o == nil || o.Records == nil {
		var ret []LogRecord
		return ret
	}
	return *o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobRecordsResult) GetRecordsOk() (*[]LogRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *LogJobRecordsResult) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []LogRecord and assigns it to the Records field.
func (o *LogJobRecordsResult) SetRecords(v []LogRecord) {
	o.Records = &v
}

// GetScrollToken returns the ScrollToken field value if set, zero value otherwise.
func (o *LogJobRecordsResult) GetScrollToken() string {
	if o == nil || o.ScrollToken == nil {
		var ret string
		return ret
	}
	return *o.ScrollToken
}

// GetScrollTokenOk returns a tuple with the ScrollToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobRecordsResult) GetScrollTokenOk() (*string, bool) {
	if o == nil || o.ScrollToken == nil {
		return nil, false
	}
	return o.ScrollToken, true
}

// HasScrollToken returns a boolean if a field has been set.
func (o *LogJobRecordsResult) HasScrollToken() bool {
	if o != nil && o.ScrollToken != nil {
		return true
	}

	return false
}

// SetScrollToken gets a reference to the given string and assigns it to the ScrollToken field.
func (o *LogJobRecordsResult) SetScrollToken(v string) {
	o.ScrollToken = &v
}

func (o LogJobRecordsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	if o.ScrollToken != nil {
		toSerialize["scrollToken"] = o.ScrollToken
	}
	return json.Marshal(toSerialize)
}

type NullableLogJobRecordsResult struct {
	value *LogJobRecordsResult
	isSet bool
}

func (v NullableLogJobRecordsResult) Get() *LogJobRecordsResult {
	return v.value
}

func (v *NullableLogJobRecordsResult) Set(val *LogJobRecordsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLogJobRecordsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLogJobRecordsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogJobRecordsResult(val *LogJobRecordsResult) *NullableLogJobRecordsResult {
	return &NullableLogJobRecordsResult{value: val, isSet: true}
}

func (v NullableLogJobRecordsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogJobRecordsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

