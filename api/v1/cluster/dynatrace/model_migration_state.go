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

// MigrationState State of in-server config migration for single component
type MigrationState struct {
	// Current status of migration
	Status *string `json:"status,omitempty"`
	// Timestamp (milliseconds format) of migration start
	StartedAt *int64 `json:"startedAt,omitempty"`
	// Timestamp (milliseconds format) of migration finish
	FinishedAt *int64 `json:"finishedAt,omitempty"`
	// Additional information about migration state
	Details *string `json:"details,omitempty"`
}

// NewMigrationState instantiates a new MigrationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationState() *MigrationState {
	this := MigrationState{}
	return &this
}

// NewMigrationStateWithDefaults instantiates a new MigrationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationStateWithDefaults() *MigrationState {
	this := MigrationState{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MigrationState) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationState) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MigrationState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MigrationState) SetStatus(v string) {
	o.Status = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *MigrationState) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationState) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *MigrationState) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *MigrationState) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *MigrationState) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationState) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *MigrationState) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *MigrationState) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MigrationState) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationState) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MigrationState) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *MigrationState) SetDetails(v string) {
	o.Details = &v
}

func (o MigrationState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationState struct {
	value *MigrationState
	isSet bool
}

func (v NullableMigrationState) Get() *MigrationState {
	return v.value
}

func (v *NullableMigrationState) Set(val *MigrationState) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationState) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationState(val *MigrationState) *NullableMigrationState {
	return &NullableMigrationState{value: val, isSet: true}
}

func (v NullableMigrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


