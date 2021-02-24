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

// Credentials A set of credentials for synthetic monitors.
type Credentials struct {
	// The name of the credentials set.
	Name string `json:"name"`
	// The ID of the credentials set.
	Id *string `json:"id,omitempty"`
	// A short description of the credentials set.
	Description *string `json:"description,omitempty"`
	// The credentials set is available to every user (`false`) or to owner only (`true`).
	OwnerAccessOnly *bool `json:"ownerAccessOnly,omitempty"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `CERTIFICATE` -> CertificateCredentials  * `USERNAME_PASSWORD` -> UserPasswordCredentials  * `TOKEN` -> TokenCredentials  
	Type *string `json:"type,omitempty"`
}

// NewCredentials instantiates a new Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentials(name string, ) *Credentials {
	this := Credentials{}
	this.Name = name
	return &this
}

// NewCredentialsWithDefaults instantiates a new Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsWithDefaults() *Credentials {
	this := Credentials{}
	return &this
}

// GetName returns the Name field value
func (o *Credentials) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Credentials) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Credentials) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Credentials) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Credentials) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Credentials) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Credentials) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Credentials) SetDescription(v string) {
	o.Description = &v
}

// GetOwnerAccessOnly returns the OwnerAccessOnly field value if set, zero value otherwise.
func (o *Credentials) GetOwnerAccessOnly() bool {
	if o == nil || o.OwnerAccessOnly == nil {
		var ret bool
		return ret
	}
	return *o.OwnerAccessOnly
}

// GetOwnerAccessOnlyOk returns a tuple with the OwnerAccessOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetOwnerAccessOnlyOk() (*bool, bool) {
	if o == nil || o.OwnerAccessOnly == nil {
		return nil, false
	}
	return o.OwnerAccessOnly, true
}

// HasOwnerAccessOnly returns a boolean if a field has been set.
func (o *Credentials) HasOwnerAccessOnly() bool {
	if o != nil && o.OwnerAccessOnly != nil {
		return true
	}

	return false
}

// SetOwnerAccessOnly gets a reference to the given bool and assigns it to the OwnerAccessOnly field.
func (o *Credentials) SetOwnerAccessOnly(v bool) {
	o.OwnerAccessOnly = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Credentials) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Credentials) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Credentials) SetType(v string) {
	o.Type = &v
}

func (o Credentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OwnerAccessOnly != nil {
		toSerialize["ownerAccessOnly"] = o.OwnerAccessOnly
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCredentials struct {
	value *Credentials
	isSet bool
}

func (v NullableCredentials) Get() *Credentials {
	return v.value
}

func (v *NullableCredentials) Set(val *Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentials(val *Credentials) *NullableCredentials {
	return &NullableCredentials{value: val, isSet: true}
}

func (v NullableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


