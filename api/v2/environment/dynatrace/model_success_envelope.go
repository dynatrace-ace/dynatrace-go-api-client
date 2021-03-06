/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// SuccessEnvelope struct for SuccessEnvelope
type SuccessEnvelope struct {
	Details *Success `json:"details,omitempty"`
}

// NewSuccessEnvelope instantiates a new SuccessEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessEnvelope() *SuccessEnvelope {
	this := SuccessEnvelope{}
	return &this
}

// NewSuccessEnvelopeWithDefaults instantiates a new SuccessEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessEnvelopeWithDefaults() *SuccessEnvelope {
	this := SuccessEnvelope{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SuccessEnvelope) GetDetails() Success {
	if o == nil || o.Details == nil {
		var ret Success
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessEnvelope) GetDetailsOk() (*Success, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SuccessEnvelope) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given Success and assigns it to the Details field.
func (o *SuccessEnvelope) SetDetails(v Success) {
	o.Details = &v
}

func (o SuccessEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessEnvelope struct {
	value *SuccessEnvelope
	isSet bool
}

func (v NullableSuccessEnvelope) Get() *SuccessEnvelope {
	return v.value
}

func (v *NullableSuccessEnvelope) Set(val *SuccessEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessEnvelope(val *SuccessEnvelope) *NullableSuccessEnvelope {
	return &NullableSuccessEnvelope{value: val, isSet: true}
}

func (v NullableSuccessEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


