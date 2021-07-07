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

// MzPermissionsForGroup struct for MzPermissionsForGroup
type MzPermissionsForGroup struct {
	// Group ID
	GroupId *string `json:"groupId,omitempty"`
	// List of management zone permissions per environment
	MzPermissionsPerEnvironment *[]MzListForEnvironment `json:"mzPermissionsPerEnvironment,omitempty"`
}

// NewMzPermissionsForGroup instantiates a new MzPermissionsForGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMzPermissionsForGroup() *MzPermissionsForGroup {
	this := MzPermissionsForGroup{}
	return &this
}

// NewMzPermissionsForGroupWithDefaults instantiates a new MzPermissionsForGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMzPermissionsForGroupWithDefaults() *MzPermissionsForGroup {
	this := MzPermissionsForGroup{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MzPermissionsForGroup) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MzPermissionsForGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MzPermissionsForGroup) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *MzPermissionsForGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetMzPermissionsPerEnvironment returns the MzPermissionsPerEnvironment field value if set, zero value otherwise.
func (o *MzPermissionsForGroup) GetMzPermissionsPerEnvironment() []MzListForEnvironment {
	if o == nil || o.MzPermissionsPerEnvironment == nil {
		var ret []MzListForEnvironment
		return ret
	}
	return *o.MzPermissionsPerEnvironment
}

// GetMzPermissionsPerEnvironmentOk returns a tuple with the MzPermissionsPerEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MzPermissionsForGroup) GetMzPermissionsPerEnvironmentOk() (*[]MzListForEnvironment, bool) {
	if o == nil || o.MzPermissionsPerEnvironment == nil {
		return nil, false
	}
	return o.MzPermissionsPerEnvironment, true
}

// HasMzPermissionsPerEnvironment returns a boolean if a field has been set.
func (o *MzPermissionsForGroup) HasMzPermissionsPerEnvironment() bool {
	if o != nil && o.MzPermissionsPerEnvironment != nil {
		return true
	}

	return false
}

// SetMzPermissionsPerEnvironment gets a reference to the given []MzListForEnvironment and assigns it to the MzPermissionsPerEnvironment field.
func (o *MzPermissionsForGroup) SetMzPermissionsPerEnvironment(v []MzListForEnvironment) {
	o.MzPermissionsPerEnvironment = &v
}

func (o MzPermissionsForGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.MzPermissionsPerEnvironment != nil {
		toSerialize["mzPermissionsPerEnvironment"] = o.MzPermissionsPerEnvironment
	}
	return json.Marshal(toSerialize)
}

type NullableMzPermissionsForGroup struct {
	value *MzPermissionsForGroup
	isSet bool
}

func (v NullableMzPermissionsForGroup) Get() *MzPermissionsForGroup {
	return v.value
}

func (v *NullableMzPermissionsForGroup) Set(val *MzPermissionsForGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMzPermissionsForGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMzPermissionsForGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMzPermissionsForGroup(val *MzPermissionsForGroup) *NullableMzPermissionsForGroup {
	return &NullableMzPermissionsForGroup{value: val, isSet: true}
}

func (v NullableMzPermissionsForGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMzPermissionsForGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

