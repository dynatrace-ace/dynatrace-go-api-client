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

// SettingsObjectErrorResponse The response to a failed settings object request
type SettingsObjectErrorResponse struct {
	// The value of the setting.    It defines the actual values of settings' parameters.   The actual content depends on the object's schema.
	InvalidValue *map[string]interface{} `json:"invalidValue,omitempty"`
	Error *Error `json:"error,omitempty"`
	// The HTTP status code for the object.
	Code *int32 `json:"code,omitempty"`
}

// NewSettingsObjectErrorResponse instantiates a new SettingsObjectErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsObjectErrorResponse() *SettingsObjectErrorResponse {
	this := SettingsObjectErrorResponse{}
	return &this
}

// NewSettingsObjectErrorResponseWithDefaults instantiates a new SettingsObjectErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsObjectErrorResponseWithDefaults() *SettingsObjectErrorResponse {
	this := SettingsObjectErrorResponse{}
	return &this
}

// GetInvalidValue returns the InvalidValue field value if set, zero value otherwise.
func (o *SettingsObjectErrorResponse) GetInvalidValue() map[string]interface{} {
	if o == nil || o.InvalidValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InvalidValue
}

// GetInvalidValueOk returns a tuple with the InvalidValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectErrorResponse) GetInvalidValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.InvalidValue == nil {
		return nil, false
	}
	return o.InvalidValue, true
}

// HasInvalidValue returns a boolean if a field has been set.
func (o *SettingsObjectErrorResponse) HasInvalidValue() bool {
	if o != nil && o.InvalidValue != nil {
		return true
	}

	return false
}

// SetInvalidValue gets a reference to the given map[string]interface{} and assigns it to the InvalidValue field.
func (o *SettingsObjectErrorResponse) SetInvalidValue(v map[string]interface{}) {
	o.InvalidValue = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SettingsObjectErrorResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectErrorResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SettingsObjectErrorResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SettingsObjectErrorResponse) SetError(v Error) {
	o.Error = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SettingsObjectErrorResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectErrorResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SettingsObjectErrorResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *SettingsObjectErrorResponse) SetCode(v int32) {
	o.Code = &v
}

func (o SettingsObjectErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InvalidValue != nil {
		toSerialize["invalidValue"] = o.InvalidValue
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsObjectErrorResponse struct {
	value *SettingsObjectErrorResponse
	isSet bool
}

func (v NullableSettingsObjectErrorResponse) Get() *SettingsObjectErrorResponse {
	return v.value
}

func (v *NullableSettingsObjectErrorResponse) Set(val *SettingsObjectErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsObjectErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsObjectErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsObjectErrorResponse(val *SettingsObjectErrorResponse) *NullableSettingsObjectErrorResponse {
	return &NullableSettingsObjectErrorResponse{value: val, isSet: true}
}

func (v NullableSettingsObjectErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsObjectErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


