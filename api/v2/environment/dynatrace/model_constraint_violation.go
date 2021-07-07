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

// ConstraintViolation A list of constraint violations
type ConstraintViolation struct {
	ParameterLocation *string `json:"parameterLocation,omitempty"`
	Location *string `json:"location,omitempty"`
	Message *string `json:"message,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewConstraintViolation instantiates a new ConstraintViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintViolation() *ConstraintViolation {
	this := ConstraintViolation{}
	return &this
}

// NewConstraintViolationWithDefaults instantiates a new ConstraintViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintViolationWithDefaults() *ConstraintViolation {
	this := ConstraintViolation{}
	return &this
}

// GetParameterLocation returns the ParameterLocation field value if set, zero value otherwise.
func (o *ConstraintViolation) GetParameterLocation() string {
	if o == nil || o.ParameterLocation == nil {
		var ret string
		return ret
	}
	return *o.ParameterLocation
}

// GetParameterLocationOk returns a tuple with the ParameterLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintViolation) GetParameterLocationOk() (*string, bool) {
	if o == nil || o.ParameterLocation == nil {
		return nil, false
	}
	return o.ParameterLocation, true
}

// HasParameterLocation returns a boolean if a field has been set.
func (o *ConstraintViolation) HasParameterLocation() bool {
	if o != nil && o.ParameterLocation != nil {
		return true
	}

	return false
}

// SetParameterLocation gets a reference to the given string and assigns it to the ParameterLocation field.
func (o *ConstraintViolation) SetParameterLocation(v string) {
	o.ParameterLocation = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ConstraintViolation) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintViolation) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ConstraintViolation) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ConstraintViolation) SetLocation(v string) {
	o.Location = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConstraintViolation) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintViolation) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConstraintViolation) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConstraintViolation) SetMessage(v string) {
	o.Message = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ConstraintViolation) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintViolation) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ConstraintViolation) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ConstraintViolation) SetPath(v string) {
	o.Path = &v
}

func (o ConstraintViolation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParameterLocation != nil {
		toSerialize["parameterLocation"] = o.ParameterLocation
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableConstraintViolation struct {
	value *ConstraintViolation
	isSet bool
}

func (v NullableConstraintViolation) Get() *ConstraintViolation {
	return v.value
}

func (v *NullableConstraintViolation) Set(val *ConstraintViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintViolation(val *ConstraintViolation) *NullableConstraintViolation {
	return &NullableConstraintViolation{value: val, isSet: true}
}

func (v NullableConstraintViolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraintViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

