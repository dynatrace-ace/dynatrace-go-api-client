/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster management functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found, e.g. User account. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade. 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// AuthenticationModel struct for AuthenticationModel
type AuthenticationModel struct {
	// Defines the authentication provider
	AuthenticationProvider string `json:"authenticationProvider"`
	// Defines the authorization provider
	AuthorizationProvider *string `json:"authorizationProvider,omitempty"`
	// Defines login page type as SSO
	SsoOnly *bool `json:"ssoOnly,omitempty"`
}

// NewAuthenticationModel instantiates a new AuthenticationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationModel(authenticationProvider string) *AuthenticationModel {
	this := AuthenticationModel{}
	this.AuthenticationProvider = authenticationProvider
	return &this
}

// NewAuthenticationModelWithDefaults instantiates a new AuthenticationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationModelWithDefaults() *AuthenticationModel {
	this := AuthenticationModel{}
	return &this
}

// GetAuthenticationProvider returns the AuthenticationProvider field value
func (o *AuthenticationModel) GetAuthenticationProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationProvider
}

// GetAuthenticationProviderOk returns a tuple with the AuthenticationProvider field value
// and a boolean to check if the value has been set.
func (o *AuthenticationModel) GetAuthenticationProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthenticationProvider, true
}

// SetAuthenticationProvider sets field value
func (o *AuthenticationModel) SetAuthenticationProvider(v string) {
	o.AuthenticationProvider = v
}

// GetAuthorizationProvider returns the AuthorizationProvider field value if set, zero value otherwise.
func (o *AuthenticationModel) GetAuthorizationProvider() string {
	if o == nil || o.AuthorizationProvider == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationProvider
}

// GetAuthorizationProviderOk returns a tuple with the AuthorizationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationModel) GetAuthorizationProviderOk() (*string, bool) {
	if o == nil || o.AuthorizationProvider == nil {
		return nil, false
	}
	return o.AuthorizationProvider, true
}

// HasAuthorizationProvider returns a boolean if a field has been set.
func (o *AuthenticationModel) HasAuthorizationProvider() bool {
	if o != nil && o.AuthorizationProvider != nil {
		return true
	}

	return false
}

// SetAuthorizationProvider gets a reference to the given string and assigns it to the AuthorizationProvider field.
func (o *AuthenticationModel) SetAuthorizationProvider(v string) {
	o.AuthorizationProvider = &v
}

// GetSsoOnly returns the SsoOnly field value if set, zero value otherwise.
func (o *AuthenticationModel) GetSsoOnly() bool {
	if o == nil || o.SsoOnly == nil {
		var ret bool
		return ret
	}
	return *o.SsoOnly
}

// GetSsoOnlyOk returns a tuple with the SsoOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationModel) GetSsoOnlyOk() (*bool, bool) {
	if o == nil || o.SsoOnly == nil {
		return nil, false
	}
	return o.SsoOnly, true
}

// HasSsoOnly returns a boolean if a field has been set.
func (o *AuthenticationModel) HasSsoOnly() bool {
	if o != nil && o.SsoOnly != nil {
		return true
	}

	return false
}

// SetSsoOnly gets a reference to the given bool and assigns it to the SsoOnly field.
func (o *AuthenticationModel) SetSsoOnly(v bool) {
	o.SsoOnly = &v
}

func (o AuthenticationModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authenticationProvider"] = o.AuthenticationProvider
	}
	if o.AuthorizationProvider != nil {
		toSerialize["authorizationProvider"] = o.AuthorizationProvider
	}
	if o.SsoOnly != nil {
		toSerialize["ssoOnly"] = o.SsoOnly
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationModel struct {
	value *AuthenticationModel
	isSet bool
}

func (v NullableAuthenticationModel) Get() *AuthenticationModel {
	return v.value
}

func (v *NullableAuthenticationModel) Set(val *AuthenticationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationModel(val *AuthenticationModel) *NullableAuthenticationModel {
	return &NullableAuthenticationModel{value: val, isSet: true}
}

func (v NullableAuthenticationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


