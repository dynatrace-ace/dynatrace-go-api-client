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

// MobileSessionPropertyUpdate An update of a mobile session property.
type MobileSessionPropertyUpdate struct {
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
	// The data type of the property.
	Type string `json:"type"`
	// The origin of the property
	Origin string `json:"origin"`
	// The aggregation type of the property.     It defines how multiple values of the property are aggregated.
	Aggregation *string `json:"aggregation,omitempty"`
	// If `true`, the property is stored as a user action property
	StoreAsUserActionProperty *bool `json:"storeAsUserActionProperty,omitempty"`
	// If `true`, the property is stored as a session property
	StoreAsSessionProperty *bool `json:"storeAsSessionProperty,omitempty"`
	// The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://www.dynatrace.com/support/help/shortlink/regex) for the data you need there.
	CleanupRule *string `json:"cleanupRule,omitempty"`
	// The ID of the request attribute.   Only applicable when the **origin** is set to `SERVER_SIDE_REQUEST_ATTRIBUTE`.
	ServerSideRequestAttribute *string `json:"serverSideRequestAttribute,omitempty"`
	// The name of the reported value.   Only applicable when the **origin** is set to `API`.
	Name *string `json:"name,omitempty"`
}

// NewMobileSessionPropertyUpdate instantiates a new MobileSessionPropertyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileSessionPropertyUpdate(type_ string, origin string, ) *MobileSessionPropertyUpdate {
	this := MobileSessionPropertyUpdate{}
	this.Type = type_
	this.Origin = origin
	return &this
}

// NewMobileSessionPropertyUpdateWithDefaults instantiates a new MobileSessionPropertyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileSessionPropertyUpdateWithDefaults() *MobileSessionPropertyUpdate {
	this := MobileSessionPropertyUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MobileSessionPropertyUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value
func (o *MobileSessionPropertyUpdate) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MobileSessionPropertyUpdate) SetType(v string) {
	o.Type = v
}

// GetOrigin returns the Origin field value
func (o *MobileSessionPropertyUpdate) GetOrigin() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetOriginOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *MobileSessionPropertyUpdate) SetOrigin(v string) {
	o.Origin = v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *MobileSessionPropertyUpdate) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetStoreAsUserActionProperty() bool {
	if o == nil || o.StoreAsUserActionProperty == nil {
		var ret bool
		return ret
	}
	return *o.StoreAsUserActionProperty
}

// GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetStoreAsUserActionPropertyOk() (*bool, bool) {
	if o == nil || o.StoreAsUserActionProperty == nil {
		return nil, false
	}
	return o.StoreAsUserActionProperty, true
}

// HasStoreAsUserActionProperty returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasStoreAsUserActionProperty() bool {
	if o != nil && o.StoreAsUserActionProperty != nil {
		return true
	}

	return false
}

// SetStoreAsUserActionProperty gets a reference to the given bool and assigns it to the StoreAsUserActionProperty field.
func (o *MobileSessionPropertyUpdate) SetStoreAsUserActionProperty(v bool) {
	o.StoreAsUserActionProperty = &v
}

// GetStoreAsSessionProperty returns the StoreAsSessionProperty field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetStoreAsSessionProperty() bool {
	if o == nil || o.StoreAsSessionProperty == nil {
		var ret bool
		return ret
	}
	return *o.StoreAsSessionProperty
}

// GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetStoreAsSessionPropertyOk() (*bool, bool) {
	if o == nil || o.StoreAsSessionProperty == nil {
		return nil, false
	}
	return o.StoreAsSessionProperty, true
}

// HasStoreAsSessionProperty returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasStoreAsSessionProperty() bool {
	if o != nil && o.StoreAsSessionProperty != nil {
		return true
	}

	return false
}

// SetStoreAsSessionProperty gets a reference to the given bool and assigns it to the StoreAsSessionProperty field.
func (o *MobileSessionPropertyUpdate) SetStoreAsSessionProperty(v bool) {
	o.StoreAsSessionProperty = &v
}

// GetCleanupRule returns the CleanupRule field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetCleanupRule() string {
	if o == nil || o.CleanupRule == nil {
		var ret string
		return ret
	}
	return *o.CleanupRule
}

// GetCleanupRuleOk returns a tuple with the CleanupRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetCleanupRuleOk() (*string, bool) {
	if o == nil || o.CleanupRule == nil {
		return nil, false
	}
	return o.CleanupRule, true
}

// HasCleanupRule returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasCleanupRule() bool {
	if o != nil && o.CleanupRule != nil {
		return true
	}

	return false
}

// SetCleanupRule gets a reference to the given string and assigns it to the CleanupRule field.
func (o *MobileSessionPropertyUpdate) SetCleanupRule(v string) {
	o.CleanupRule = &v
}

// GetServerSideRequestAttribute returns the ServerSideRequestAttribute field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetServerSideRequestAttribute() string {
	if o == nil || o.ServerSideRequestAttribute == nil {
		var ret string
		return ret
	}
	return *o.ServerSideRequestAttribute
}

// GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetServerSideRequestAttributeOk() (*string, bool) {
	if o == nil || o.ServerSideRequestAttribute == nil {
		return nil, false
	}
	return o.ServerSideRequestAttribute, true
}

// HasServerSideRequestAttribute returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasServerSideRequestAttribute() bool {
	if o != nil && o.ServerSideRequestAttribute != nil {
		return true
	}

	return false
}

// SetServerSideRequestAttribute gets a reference to the given string and assigns it to the ServerSideRequestAttribute field.
func (o *MobileSessionPropertyUpdate) SetServerSideRequestAttribute(v string) {
	o.ServerSideRequestAttribute = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MobileSessionPropertyUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionPropertyUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MobileSessionPropertyUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MobileSessionPropertyUpdate) SetName(v string) {
	o.Name = &v
}

func (o MobileSessionPropertyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["origin"] = o.Origin
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.StoreAsUserActionProperty != nil {
		toSerialize["storeAsUserActionProperty"] = o.StoreAsUserActionProperty
	}
	if o.StoreAsSessionProperty != nil {
		toSerialize["storeAsSessionProperty"] = o.StoreAsSessionProperty
	}
	if o.CleanupRule != nil {
		toSerialize["cleanupRule"] = o.CleanupRule
	}
	if o.ServerSideRequestAttribute != nil {
		toSerialize["serverSideRequestAttribute"] = o.ServerSideRequestAttribute
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMobileSessionPropertyUpdate struct {
	value *MobileSessionPropertyUpdate
	isSet bool
}

func (v NullableMobileSessionPropertyUpdate) Get() *MobileSessionPropertyUpdate {
	return v.value
}

func (v *NullableMobileSessionPropertyUpdate) Set(val *MobileSessionPropertyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileSessionPropertyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileSessionPropertyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileSessionPropertyUpdate(val *MobileSessionPropertyUpdate) *NullableMobileSessionPropertyUpdate {
	return &NullableMobileSessionPropertyUpdate{value: val, isSet: true}
}

func (v NullableMobileSessionPropertyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileSessionPropertyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

