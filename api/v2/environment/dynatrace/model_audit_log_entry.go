/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// AuditLogEntry An entry of the audit log.
type AuditLogEntry struct {
	// The ID of the log entry.
	LogId string `json:"logId"`
	// The type of the recorded operation.
	EventType string `json:"eventType"`
	// The category of the recorded operation.
	Category string `json:"category"`
	// The ID of an entity from the **category**.   For example, it can be config ID for the `CONFIG` category or token ID for the `TOKEN` category.
	EntityId *string `json:"entityId,omitempty"`
	// The ID of the Dynatrace environment where the recorded operation occurred.
	EnvironmentId string `json:"environmentId"`
	// The ID of the user who performed the recorded operation.
	User string `json:"user"`
	// The type of the authentication of the **user**.
	UserType string `json:"userType"`
	// The origin and the IP address of the **user**.
	UserOrigin *string `json:"userOrigin,omitempty"`
	// The timestamp of the record creation, in UTC milliseconds.
	Timestamp int64 `json:"timestamp"`
	// The recorded operation is successful (`true`) or failed (`false`).
	Success bool `json:"success"`
	// The logged message.
	Message *string `json:"message,omitempty"`
	// The patch of the recorded operation as the JSON representation.   The format is an enhanced RFC 6902. The patch also carries the previous value in the **oldValue** field.
	Patch *map[string]interface{} `json:"patch,omitempty"`
}

// NewAuditLogEntry instantiates a new AuditLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogEntry(logId string, eventType string, category string, environmentId string, user string, userType string, timestamp int64, success bool) *AuditLogEntry {
	this := AuditLogEntry{}
	this.LogId = logId
	this.EventType = eventType
	this.Category = category
	this.EnvironmentId = environmentId
	this.User = user
	this.UserType = userType
	this.Timestamp = timestamp
	this.Success = success
	return &this
}

// NewAuditLogEntryWithDefaults instantiates a new AuditLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogEntryWithDefaults() *AuditLogEntry {
	this := AuditLogEntry{}
	return &this
}

// GetLogId returns the LogId field value
func (o *AuditLogEntry) GetLogId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetLogIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogId, true
}

// SetLogId sets field value
func (o *AuditLogEntry) SetLogId(v string) {
	o.LogId = v
}

// GetEventType returns the EventType field value
func (o *AuditLogEntry) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *AuditLogEntry) SetEventType(v string) {
	o.EventType = v
}

// GetCategory returns the Category field value
func (o *AuditLogEntry) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *AuditLogEntry) SetCategory(v string) {
	o.Category = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AuditLogEntry) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AuditLogEntry) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AuditLogEntry) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *AuditLogEntry) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetEnvironmentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *AuditLogEntry) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetUser returns the User field value
func (o *AuditLogEntry) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AuditLogEntry) SetUser(v string) {
	o.User = v
}

// GetUserType returns the UserType field value
func (o *AuditLogEntry) GetUserType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetUserTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserType, true
}

// SetUserType sets field value
func (o *AuditLogEntry) SetUserType(v string) {
	o.UserType = v
}

// GetUserOrigin returns the UserOrigin field value if set, zero value otherwise.
func (o *AuditLogEntry) GetUserOrigin() string {
	if o == nil || o.UserOrigin == nil {
		var ret string
		return ret
	}
	return *o.UserOrigin
}

// GetUserOriginOk returns a tuple with the UserOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetUserOriginOk() (*string, bool) {
	if o == nil || o.UserOrigin == nil {
		return nil, false
	}
	return o.UserOrigin, true
}

// HasUserOrigin returns a boolean if a field has been set.
func (o *AuditLogEntry) HasUserOrigin() bool {
	if o != nil && o.UserOrigin != nil {
		return true
	}

	return false
}

// SetUserOrigin gets a reference to the given string and assigns it to the UserOrigin field.
func (o *AuditLogEntry) SetUserOrigin(v string) {
	o.UserOrigin = &v
}

// GetTimestamp returns the Timestamp field value
func (o *AuditLogEntry) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetTimestampOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AuditLogEntry) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetSuccess returns the Success field value
func (o *AuditLogEntry) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AuditLogEntry) SetSuccess(v bool) {
	o.Success = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuditLogEntry) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuditLogEntry) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuditLogEntry) SetMessage(v string) {
	o.Message = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *AuditLogEntry) GetPatch() map[string]interface{} {
	if o == nil || o.Patch == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLogEntry) GetPatchOk() (*map[string]interface{}, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *AuditLogEntry) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given map[string]interface{} and assigns it to the Patch field.
func (o *AuditLogEntry) SetPatch(v map[string]interface{}) {
	o.Patch = &v
}

func (o AuditLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logId"] = o.LogId
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if true {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["userType"] = o.UserType
	}
	if o.UserOrigin != nil {
		toSerialize["userOrigin"] = o.UserOrigin
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["success"] = o.Success
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLogEntry struct {
	value *AuditLogEntry
	isSet bool
}

func (v NullableAuditLogEntry) Get() *AuditLogEntry {
	return v.value
}

func (v *NullableAuditLogEntry) Set(val *AuditLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogEntry(val *AuditLogEntry) *NullableAuditLogEntry {
	return &NullableAuditLogEntry{value: val, isSet: true}
}

func (v NullableAuditLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

