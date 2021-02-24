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

// KubernetesEventPattern Represents a single Kubernetes events field selector (=event filter based on the K8s field selector).
type KubernetesEventPattern struct {
	// A label of the events field selector.
	Label string `json:"label"`
	// The field selector string (url decoding is applied) when storing it.
	FieldSelector string `json:"fieldSelector"`
	// Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector
	Active bool `json:"active"`
}

// NewKubernetesEventPattern instantiates a new KubernetesEventPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEventPattern(label string, fieldSelector string, active bool, ) *KubernetesEventPattern {
	this := KubernetesEventPattern{}
	this.Label = label
	this.FieldSelector = fieldSelector
	this.Active = active
	return &this
}

// NewKubernetesEventPatternWithDefaults instantiates a new KubernetesEventPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEventPatternWithDefaults() *KubernetesEventPattern {
	this := KubernetesEventPattern{}
	return &this
}

// GetLabel returns the Label field value
func (o *KubernetesEventPattern) GetLabel() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *KubernetesEventPattern) SetLabel(v string) {
	o.Label = v
}

// GetFieldSelector returns the FieldSelector field value
func (o *KubernetesEventPattern) GetFieldSelector() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FieldSelector
}

// GetFieldSelectorOk returns a tuple with the FieldSelector field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetFieldSelectorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FieldSelector, true
}

// SetFieldSelector sets field value
func (o *KubernetesEventPattern) SetFieldSelector(v string) {
	o.FieldSelector = v
}

// GetActive returns the Active field value
func (o *KubernetesEventPattern) GetActive() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *KubernetesEventPattern) SetActive(v bool) {
	o.Active = v
}

func (o KubernetesEventPattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["fieldSelector"] = o.FieldSelector
	}
	if true {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesEventPattern struct {
	value *KubernetesEventPattern
	isSet bool
}

func (v NullableKubernetesEventPattern) Get() *KubernetesEventPattern {
	return v.value
}

func (v *NullableKubernetesEventPattern) Set(val *KubernetesEventPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEventPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEventPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEventPattern(val *KubernetesEventPattern) *NullableKubernetesEventPattern {
	return &NullableKubernetesEventPattern{value: val, isSet: true}
}

func (v NullableKubernetesEventPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEventPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


