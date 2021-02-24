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

// UserActionPropertyFilter User action property filter.
type UserActionPropertyFilter struct {
	// The key of the action property we're checking.
	Key *string `json:"key,omitempty"`
	// Only actions that have this value in the specified property are included in the metric calculation.    Only applicable to string values.
	Value *string `json:"value,omitempty"`
	// Only actions that have a value greater than or equal to this are included in the metric calculation.    Only applicable to numerical values.
	From *float64 `json:"from,omitempty"`
	// Only actions that have a value less than or equal to this are included in the metric calculation.    Only applicable to numerical values.
	To *float64 `json:"to,omitempty"`
	// Specifies the match type of a string filter, e.g. using `Contains` or `Equals`.    Only applicable to string values.
	MatchType *string `json:"matchType,omitempty"`
}

// NewUserActionPropertyFilter instantiates a new UserActionPropertyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActionPropertyFilter() *UserActionPropertyFilter {
	this := UserActionPropertyFilter{}
	return &this
}

// NewUserActionPropertyFilterWithDefaults instantiates a new UserActionPropertyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActionPropertyFilterWithDefaults() *UserActionPropertyFilter {
	this := UserActionPropertyFilter{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UserActionPropertyFilter) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionPropertyFilter) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UserActionPropertyFilter) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UserActionPropertyFilter) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserActionPropertyFilter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionPropertyFilter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserActionPropertyFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserActionPropertyFilter) SetValue(v string) {
	o.Value = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *UserActionPropertyFilter) GetFrom() float64 {
	if o == nil || o.From == nil {
		var ret float64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionPropertyFilter) GetFromOk() (*float64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *UserActionPropertyFilter) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given float64 and assigns it to the From field.
func (o *UserActionPropertyFilter) SetFrom(v float64) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *UserActionPropertyFilter) GetTo() float64 {
	if o == nil || o.To == nil {
		var ret float64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionPropertyFilter) GetToOk() (*float64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *UserActionPropertyFilter) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given float64 and assigns it to the To field.
func (o *UserActionPropertyFilter) SetTo(v float64) {
	o.To = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *UserActionPropertyFilter) GetMatchType() string {
	if o == nil || o.MatchType == nil {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionPropertyFilter) GetMatchTypeOk() (*string, bool) {
	if o == nil || o.MatchType == nil {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *UserActionPropertyFilter) HasMatchType() bool {
	if o != nil && o.MatchType != nil {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *UserActionPropertyFilter) SetMatchType(v string) {
	o.MatchType = &v
}

func (o UserActionPropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.MatchType != nil {
		toSerialize["matchType"] = o.MatchType
	}
	return json.Marshal(toSerialize)
}

type NullableUserActionPropertyFilter struct {
	value *UserActionPropertyFilter
	isSet bool
}

func (v NullableUserActionPropertyFilter) Get() *UserActionPropertyFilter {
	return v.value
}

func (v *NullableUserActionPropertyFilter) Set(val *UserActionPropertyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActionPropertyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActionPropertyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActionPropertyFilter(val *UserActionPropertyFilter) *NullableUserActionPropertyFilter {
	return &NullableUserActionPropertyFilter{value: val, isSet: true}
}

func (v NullableUserActionPropertyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActionPropertyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


