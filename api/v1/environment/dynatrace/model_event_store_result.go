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

// EventStoreResult Contains IDs of all custom events, created by an event push call.
type EventStoreResult struct {
	// List of event IDs for information-only-events.    This field is provided for compatibility reasons. You should use the values from the **storedIds** field instead.
	StoredEventIds *[]int64 `json:"storedEventIds,omitempty"`
	// List of **encoded** event IDs for information-only-events.
	StoredIds *[]string `json:"storedIds,omitempty"`
	// List of correlation IDs for problem-opening-events.
	StoredCorrelationIds *[]string `json:"storedCorrelationIds,omitempty"`
}

// NewEventStoreResult instantiates a new EventStoreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventStoreResult() *EventStoreResult {
	this := EventStoreResult{}
	return &this
}

// NewEventStoreResultWithDefaults instantiates a new EventStoreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventStoreResultWithDefaults() *EventStoreResult {
	this := EventStoreResult{}
	return &this
}

// GetStoredEventIds returns the StoredEventIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredEventIds() []int64 {
	if o == nil || o.StoredEventIds == nil {
		var ret []int64
		return ret
	}
	return *o.StoredEventIds
}

// GetStoredEventIdsOk returns a tuple with the StoredEventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredEventIdsOk() (*[]int64, bool) {
	if o == nil || o.StoredEventIds == nil {
		return nil, false
	}
	return o.StoredEventIds, true
}

// HasStoredEventIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredEventIds() bool {
	if o != nil && o.StoredEventIds != nil {
		return true
	}

	return false
}

// SetStoredEventIds gets a reference to the given []int64 and assigns it to the StoredEventIds field.
func (o *EventStoreResult) SetStoredEventIds(v []int64) {
	o.StoredEventIds = &v
}

// GetStoredIds returns the StoredIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredIds() []string {
	if o == nil || o.StoredIds == nil {
		var ret []string
		return ret
	}
	return *o.StoredIds
}

// GetStoredIdsOk returns a tuple with the StoredIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredIdsOk() (*[]string, bool) {
	if o == nil || o.StoredIds == nil {
		return nil, false
	}
	return o.StoredIds, true
}

// HasStoredIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredIds() bool {
	if o != nil && o.StoredIds != nil {
		return true
	}

	return false
}

// SetStoredIds gets a reference to the given []string and assigns it to the StoredIds field.
func (o *EventStoreResult) SetStoredIds(v []string) {
	o.StoredIds = &v
}

// GetStoredCorrelationIds returns the StoredCorrelationIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredCorrelationIds() []string {
	if o == nil || o.StoredCorrelationIds == nil {
		var ret []string
		return ret
	}
	return *o.StoredCorrelationIds
}

// GetStoredCorrelationIdsOk returns a tuple with the StoredCorrelationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredCorrelationIdsOk() (*[]string, bool) {
	if o == nil || o.StoredCorrelationIds == nil {
		return nil, false
	}
	return o.StoredCorrelationIds, true
}

// HasStoredCorrelationIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredCorrelationIds() bool {
	if o != nil && o.StoredCorrelationIds != nil {
		return true
	}

	return false
}

// SetStoredCorrelationIds gets a reference to the given []string and assigns it to the StoredCorrelationIds field.
func (o *EventStoreResult) SetStoredCorrelationIds(v []string) {
	o.StoredCorrelationIds = &v
}

func (o EventStoreResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoredEventIds != nil {
		toSerialize["storedEventIds"] = o.StoredEventIds
	}
	if o.StoredIds != nil {
		toSerialize["storedIds"] = o.StoredIds
	}
	if o.StoredCorrelationIds != nil {
		toSerialize["storedCorrelationIds"] = o.StoredCorrelationIds
	}
	return json.Marshal(toSerialize)
}

type NullableEventStoreResult struct {
	value *EventStoreResult
	isSet bool
}

func (v NullableEventStoreResult) Get() *EventStoreResult {
	return v.value
}

func (v *NullableEventStoreResult) Set(val *EventStoreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventStoreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventStoreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventStoreResult(val *EventStoreResult) *NullableEventStoreResult {
	return &NullableEventStoreResult{value: val, isSet: true}
}

func (v NullableEventStoreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventStoreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


