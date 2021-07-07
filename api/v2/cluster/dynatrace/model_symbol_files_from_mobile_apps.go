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

// SymbolFilesFromMobileApps Symbol files from mobile apps storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain.
type SymbolFilesFromMobileApps struct {
	// Currently used storage [bytes]
	CurrentlyUsed *int64 `json:"currentlyUsed,omitempty"`
	// Maximum storage limit [bytes]
	MaxLimit *int64 `json:"maxLimit,omitempty"`
}

// NewSymbolFilesFromMobileApps instantiates a new SymbolFilesFromMobileApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolFilesFromMobileApps() *SymbolFilesFromMobileApps {
	this := SymbolFilesFromMobileApps{}
	return &this
}

// NewSymbolFilesFromMobileAppsWithDefaults instantiates a new SymbolFilesFromMobileApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolFilesFromMobileAppsWithDefaults() *SymbolFilesFromMobileApps {
	this := SymbolFilesFromMobileApps{}
	return &this
}

// GetCurrentlyUsed returns the CurrentlyUsed field value if set, zero value otherwise.
func (o *SymbolFilesFromMobileApps) GetCurrentlyUsed() int64 {
	if o == nil || o.CurrentlyUsed == nil {
		var ret int64
		return ret
	}
	return *o.CurrentlyUsed
}

// GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolFilesFromMobileApps) GetCurrentlyUsedOk() (*int64, bool) {
	if o == nil || o.CurrentlyUsed == nil {
		return nil, false
	}
	return o.CurrentlyUsed, true
}

// HasCurrentlyUsed returns a boolean if a field has been set.
func (o *SymbolFilesFromMobileApps) HasCurrentlyUsed() bool {
	if o != nil && o.CurrentlyUsed != nil {
		return true
	}

	return false
}

// SetCurrentlyUsed gets a reference to the given int64 and assigns it to the CurrentlyUsed field.
func (o *SymbolFilesFromMobileApps) SetCurrentlyUsed(v int64) {
	o.CurrentlyUsed = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *SymbolFilesFromMobileApps) GetMaxLimit() int64 {
	if o == nil || o.MaxLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolFilesFromMobileApps) GetMaxLimitOk() (*int64, bool) {
	if o == nil || o.MaxLimit == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *SymbolFilesFromMobileApps) HasMaxLimit() bool {
	if o != nil && o.MaxLimit != nil {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given int64 and assigns it to the MaxLimit field.
func (o *SymbolFilesFromMobileApps) SetMaxLimit(v int64) {
	o.MaxLimit = &v
}

func (o SymbolFilesFromMobileApps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentlyUsed != nil {
		toSerialize["currentlyUsed"] = o.CurrentlyUsed
	}
	if o.MaxLimit != nil {
		toSerialize["maxLimit"] = o.MaxLimit
	}
	return json.Marshal(toSerialize)
}

type NullableSymbolFilesFromMobileApps struct {
	value *SymbolFilesFromMobileApps
	isSet bool
}

func (v NullableSymbolFilesFromMobileApps) Get() *SymbolFilesFromMobileApps {
	return v.value
}

func (v *NullableSymbolFilesFromMobileApps) Set(val *SymbolFilesFromMobileApps) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolFilesFromMobileApps) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolFilesFromMobileApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolFilesFromMobileApps(val *SymbolFilesFromMobileApps) *NullableSymbolFilesFromMobileApps {
	return &NullableSymbolFilesFromMobileApps{value: val, isSet: true}
}

func (v NullableSymbolFilesFromMobileApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolFilesFromMobileApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


