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

// TokenCredentials A credentials set of the `TOKEN` type.
type TokenCredentials struct {
	Credentials
	// Token in the string format.
	Token string `json:"token"`
}

// NewTokenCredentials instantiates a new TokenCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCredentials(token string, ) *TokenCredentials {
	this := TokenCredentials{}
	this.Token = token
	return &this
}

// NewTokenCredentialsWithDefaults instantiates a new TokenCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCredentialsWithDefaults() *TokenCredentials {
	this := TokenCredentials{}
	return &this
}

// GetToken returns the Token field value
func (o *TokenCredentials) GetToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenCredentials) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenCredentials) SetToken(v string) {
	o.Token = v
}

func (o TokenCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCredentials, errCredentials := json.Marshal(o.Credentials)
	if errCredentials != nil {
		return []byte{}, errCredentials
	}
	errCredentials = json.Unmarshal([]byte(serializedCredentials), &toSerialize)
	if errCredentials != nil {
		return []byte{}, errCredentials
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenCredentials struct {
	value *TokenCredentials
	isSet bool
}

func (v NullableTokenCredentials) Get() *TokenCredentials {
	return v.value
}

func (v *NullableTokenCredentials) Set(val *TokenCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCredentials(val *TokenCredentials) *NullableTokenCredentials {
	return &NullableTokenCredentials{value: val, isSet: true}
}

func (v NullableTokenCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


