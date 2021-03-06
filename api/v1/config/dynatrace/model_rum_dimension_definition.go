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

// RumDimensionDefinition Dimension of the calculated RUM metrics.
type RumDimensionDefinition struct {
	// The number of top values to be calculated.
	TopX int32 `json:"topX"`
	// The dimension of the metric.
	Dimension string `json:"dimension"`
	// The key of the user action property.    Only applicable for the `StringProperty` dimension.
	PropertyKey *string `json:"propertyKey,omitempty"`
}

// NewRumDimensionDefinition instantiates a new RumDimensionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRumDimensionDefinition(topX int32, dimension string, ) *RumDimensionDefinition {
	this := RumDimensionDefinition{}
	this.TopX = topX
	this.Dimension = dimension
	return &this
}

// NewRumDimensionDefinitionWithDefaults instantiates a new RumDimensionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRumDimensionDefinitionWithDefaults() *RumDimensionDefinition {
	this := RumDimensionDefinition{}
	return &this
}

// GetTopX returns the TopX field value
func (o *RumDimensionDefinition) GetTopX() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TopX
}

// GetTopXOk returns a tuple with the TopX field value
// and a boolean to check if the value has been set.
func (o *RumDimensionDefinition) GetTopXOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopX, true
}

// SetTopX sets field value
func (o *RumDimensionDefinition) SetTopX(v int32) {
	o.TopX = v
}

// GetDimension returns the Dimension field value
func (o *RumDimensionDefinition) GetDimension() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *RumDimensionDefinition) GetDimensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *RumDimensionDefinition) SetDimension(v string) {
	o.Dimension = v
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *RumDimensionDefinition) GetPropertyKey() string {
	if o == nil || o.PropertyKey == nil {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumDimensionDefinition) GetPropertyKeyOk() (*string, bool) {
	if o == nil || o.PropertyKey == nil {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *RumDimensionDefinition) HasPropertyKey() bool {
	if o != nil && o.PropertyKey != nil {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *RumDimensionDefinition) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

func (o RumDimensionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["topX"] = o.TopX
	}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if o.PropertyKey != nil {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	return json.Marshal(toSerialize)
}

type NullableRumDimensionDefinition struct {
	value *RumDimensionDefinition
	isSet bool
}

func (v NullableRumDimensionDefinition) Get() *RumDimensionDefinition {
	return v.value
}

func (v *NullableRumDimensionDefinition) Set(val *RumDimensionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRumDimensionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRumDimensionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRumDimensionDefinition(val *RumDimensionDefinition) *NullableRumDimensionDefinition {
	return &NullableRumDimensionDefinition{value: val, isSet: true}
}

func (v NullableRumDimensionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRumDimensionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


