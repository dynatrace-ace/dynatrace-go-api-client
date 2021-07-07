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

// UpgradeStartupState struct for UpgradeStartupState
type UpgradeStartupState struct {
	// System precondition check state
	State *string `json:"state,omitempty"`
	// Error
	Error *string `json:"error,omitempty"`
}

// NewUpgradeStartupState instantiates a new UpgradeStartupState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeStartupState() *UpgradeStartupState {
	this := UpgradeStartupState{}
	return &this
}

// NewUpgradeStartupStateWithDefaults instantiates a new UpgradeStartupState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeStartupStateWithDefaults() *UpgradeStartupState {
	this := UpgradeStartupState{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpgradeStartupState) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeStartupState) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpgradeStartupState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpgradeStartupState) SetState(v string) {
	o.State = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UpgradeStartupState) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeStartupState) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UpgradeStartupState) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UpgradeStartupState) SetError(v string) {
	o.Error = &v
}

func (o UpgradeStartupState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradeStartupState struct {
	value *UpgradeStartupState
	isSet bool
}

func (v NullableUpgradeStartupState) Get() *UpgradeStartupState {
	return v.value
}

func (v *NullableUpgradeStartupState) Set(val *UpgradeStartupState) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeStartupState) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeStartupState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeStartupState(val *UpgradeStartupState) *NullableUpgradeStartupState {
	return &NullableUpgradeStartupState{value: val, isSet: true}
}

func (v NullableUpgradeStartupState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeStartupState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


