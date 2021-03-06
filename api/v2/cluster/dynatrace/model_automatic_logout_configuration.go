/*
 * Dynatrace Cluster API
 *
 * Dynatrace Managed exposes cluster-wide functionality via REST endpoints. This interactive documentation also acts as a REST client you can use to interact with Dynatrace Managed clusters.   To authorize, use the API Token generated in [Settings - API Tokens page](/cmc#cm/apiToken). The HTTP status code of the response shows the result of your request. The expected response code for a successful request is documented individually per REST endpoint. Additionally the following error response codes can occur in our REST interface:  * 400 - Bad Request: Some request parameters are not correct. See response body for details. * 401 - Unauthorized: A valid authorization header is required but is missing. * 403 - Forbidden: Execution of request is not allowed, e.g. api-token is invalid. * 404 - Not Found: Endpoint does not exist or some entities could not be found. * 500 - Internal Server Error: See response body for details. * 556 - Upgrade in progress: Operation couldn't be performed during the upgrade.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// AutomaticLogoutConfiguration Configuration of automatic logout.
type AutomaticLogoutConfiguration struct {
	// True if automatic logout is enabled
	LogoutInactiveUsersEnabled bool `json:"logoutInactiveUsersEnabled"`
	// User inactivity timeout
	UserInactivityTimeout int32 `json:"userInactivityTimeout"`
}

// NewAutomaticLogoutConfiguration instantiates a new AutomaticLogoutConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomaticLogoutConfiguration(logoutInactiveUsersEnabled bool, userInactivityTimeout int32) *AutomaticLogoutConfiguration {
	this := AutomaticLogoutConfiguration{}
	this.LogoutInactiveUsersEnabled = logoutInactiveUsersEnabled
	this.UserInactivityTimeout = userInactivityTimeout
	return &this
}

// NewAutomaticLogoutConfigurationWithDefaults instantiates a new AutomaticLogoutConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomaticLogoutConfigurationWithDefaults() *AutomaticLogoutConfiguration {
	this := AutomaticLogoutConfiguration{}
	return &this
}

// GetLogoutInactiveUsersEnabled returns the LogoutInactiveUsersEnabled field value
func (o *AutomaticLogoutConfiguration) GetLogoutInactiveUsersEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LogoutInactiveUsersEnabled
}

// GetLogoutInactiveUsersEnabledOk returns a tuple with the LogoutInactiveUsersEnabled field value
// and a boolean to check if the value has been set.
func (o *AutomaticLogoutConfiguration) GetLogoutInactiveUsersEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoutInactiveUsersEnabled, true
}

// SetLogoutInactiveUsersEnabled sets field value
func (o *AutomaticLogoutConfiguration) SetLogoutInactiveUsersEnabled(v bool) {
	o.LogoutInactiveUsersEnabled = v
}

// GetUserInactivityTimeout returns the UserInactivityTimeout field value
func (o *AutomaticLogoutConfiguration) GetUserInactivityTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserInactivityTimeout
}

// GetUserInactivityTimeoutOk returns a tuple with the UserInactivityTimeout field value
// and a boolean to check if the value has been set.
func (o *AutomaticLogoutConfiguration) GetUserInactivityTimeoutOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserInactivityTimeout, true
}

// SetUserInactivityTimeout sets field value
func (o *AutomaticLogoutConfiguration) SetUserInactivityTimeout(v int32) {
	o.UserInactivityTimeout = v
}

func (o AutomaticLogoutConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logoutInactiveUsersEnabled"] = o.LogoutInactiveUsersEnabled
	}
	if true {
		toSerialize["userInactivityTimeout"] = o.UserInactivityTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableAutomaticLogoutConfiguration struct {
	value *AutomaticLogoutConfiguration
	isSet bool
}

func (v NullableAutomaticLogoutConfiguration) Get() *AutomaticLogoutConfiguration {
	return v.value
}

func (v *NullableAutomaticLogoutConfiguration) Set(val *AutomaticLogoutConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticLogoutConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticLogoutConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticLogoutConfiguration(val *AutomaticLogoutConfiguration) *NullableAutomaticLogoutConfiguration {
	return &NullableAutomaticLogoutConfiguration{value: val, isSet: true}
}

func (v NullableAutomaticLogoutConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomaticLogoutConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


