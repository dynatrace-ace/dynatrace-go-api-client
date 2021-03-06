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

// RemoteEnvironmentConfigDto struct for RemoteEnvironmentConfigDto
type RemoteEnvironmentConfigDto struct {
	// The ID of the configuration.
	Id *string `json:"id,omitempty"`
	// The display name of the remote environment.
	DisplayName string `json:"displayName"`
	// The URI of the remote environment.
	Uri string `json:"uri"`
	// The API token granting access to the remote environment.   The token must have the **Fetch data from a remote environment** (`RestRequestForwarding`) scope.   For security reasons, GET requests return this field as `null`.
	Token *string `json:"token,omitempty"`
	// The network scope of the remote environment: * `EXTERNAL`: The remote environment is located in an another network.  * `INTERNAL`: The remote environment is located in the same network.  * `CLUSTER`: The remote environment is located in the same cluster.   Dynatrace SaaS can only use `EXTERNAL`.  If not set, `EXTERNAL` is used.
	NetworkScope *string `json:"networkScope,omitempty"`
}

// NewRemoteEnvironmentConfigDto instantiates a new RemoteEnvironmentConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteEnvironmentConfigDto(displayName string, uri string, ) *RemoteEnvironmentConfigDto {
	this := RemoteEnvironmentConfigDto{}
	this.DisplayName = displayName
	this.Uri = uri
	return &this
}

// NewRemoteEnvironmentConfigDtoWithDefaults instantiates a new RemoteEnvironmentConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteEnvironmentConfigDtoWithDefaults() *RemoteEnvironmentConfigDto {
	this := RemoteEnvironmentConfigDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteEnvironmentConfigDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteEnvironmentConfigDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteEnvironmentConfigDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteEnvironmentConfigDto) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value
func (o *RemoteEnvironmentConfigDto) GetDisplayName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *RemoteEnvironmentConfigDto) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *RemoteEnvironmentConfigDto) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUri returns the Uri field value
func (o *RemoteEnvironmentConfigDto) GetUri() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *RemoteEnvironmentConfigDto) GetUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *RemoteEnvironmentConfigDto) SetUri(v string) {
	o.Uri = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RemoteEnvironmentConfigDto) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteEnvironmentConfigDto) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RemoteEnvironmentConfigDto) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RemoteEnvironmentConfigDto) SetToken(v string) {
	o.Token = &v
}

// GetNetworkScope returns the NetworkScope field value if set, zero value otherwise.
func (o *RemoteEnvironmentConfigDto) GetNetworkScope() string {
	if o == nil || o.NetworkScope == nil {
		var ret string
		return ret
	}
	return *o.NetworkScope
}

// GetNetworkScopeOk returns a tuple with the NetworkScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteEnvironmentConfigDto) GetNetworkScopeOk() (*string, bool) {
	if o == nil || o.NetworkScope == nil {
		return nil, false
	}
	return o.NetworkScope, true
}

// HasNetworkScope returns a boolean if a field has been set.
func (o *RemoteEnvironmentConfigDto) HasNetworkScope() bool {
	if o != nil && o.NetworkScope != nil {
		return true
	}

	return false
}

// SetNetworkScope gets a reference to the given string and assigns it to the NetworkScope field.
func (o *RemoteEnvironmentConfigDto) SetNetworkScope(v string) {
	o.NetworkScope = &v
}

func (o RemoteEnvironmentConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["uri"] = o.Uri
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.NetworkScope != nil {
		toSerialize["networkScope"] = o.NetworkScope
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteEnvironmentConfigDto struct {
	value *RemoteEnvironmentConfigDto
	isSet bool
}

func (v NullableRemoteEnvironmentConfigDto) Get() *RemoteEnvironmentConfigDto {
	return v.value
}

func (v *NullableRemoteEnvironmentConfigDto) Set(val *RemoteEnvironmentConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteEnvironmentConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteEnvironmentConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteEnvironmentConfigDto(val *RemoteEnvironmentConfigDto) *NullableRemoteEnvironmentConfigDto {
	return &NullableRemoteEnvironmentConfigDto{value: val, isSet: true}
}

func (v NullableRemoteEnvironmentConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteEnvironmentConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


