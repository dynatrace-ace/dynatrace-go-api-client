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

// EnumType Definition of an enum property.
type EnumType struct {
	// A list of allowed values of the enum.
	Items *[]EnumValue `json:"items,omitempty"`
	// An existing Java enum class that holds the allowed values of the enum.
	EnumClass *string `json:"enumClass,omitempty"`
	// The type of the property.
	Type *string `json:"type,omitempty"`
	// An extended description and/or links to documentation.
	Documentation *string `json:"documentation,omitempty"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
	// A short description of the property.
	Description *string `json:"description,omitempty"`
}

// NewEnumType instantiates a new EnumType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumType() *EnumType {
	this := EnumType{}
	return &this
}

// NewEnumTypeWithDefaults instantiates a new EnumType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumTypeWithDefaults() *EnumType {
	this := EnumType{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EnumType) GetItems() []EnumValue {
	if o == nil || o.Items == nil {
		var ret []EnumValue
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetItemsOk() (*[]EnumValue, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EnumType) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EnumValue and assigns it to the Items field.
func (o *EnumType) SetItems(v []EnumValue) {
	o.Items = &v
}

// GetEnumClass returns the EnumClass field value if set, zero value otherwise.
func (o *EnumType) GetEnumClass() string {
	if o == nil || o.EnumClass == nil {
		var ret string
		return ret
	}
	return *o.EnumClass
}

// GetEnumClassOk returns a tuple with the EnumClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetEnumClassOk() (*string, bool) {
	if o == nil || o.EnumClass == nil {
		return nil, false
	}
	return o.EnumClass, true
}

// HasEnumClass returns a boolean if a field has been set.
func (o *EnumType) HasEnumClass() bool {
	if o != nil && o.EnumClass != nil {
		return true
	}

	return false
}

// SetEnumClass gets a reference to the given string and assigns it to the EnumClass field.
func (o *EnumType) SetEnumClass(v string) {
	o.EnumClass = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnumType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnumType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnumType) SetType(v string) {
	o.Type = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *EnumType) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *EnumType) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *EnumType) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EnumType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EnumType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EnumType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnumType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnumType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnumType) SetDescription(v string) {
	o.Description = &v
}

func (o EnumType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.EnumClass != nil {
		toSerialize["enumClass"] = o.EnumClass
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableEnumType struct {
	value *EnumType
	isSet bool
}

func (v NullableEnumType) Get() *EnumType {
	return v.value
}

func (v *NullableEnumType) Set(val *EnumType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumType(val *EnumType) *NullableEnumType {
	return &NullableEnumType{value: val, isSet: true}
}

func (v NullableEnumType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


