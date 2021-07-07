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

// HostUnitQuota Host units consumption and quota information on environment level. If skipped when editing via PUT method then already set quota will remain.
type HostUnitQuota struct {
	// Current environment usage.
	CurrentUsage *float64 `json:"currentUsage,omitempty"`
	// Concurrent environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.
	MaxLimit *int64 `json:"maxLimit,omitempty"`
}

// NewHostUnitQuota instantiates a new HostUnitQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostUnitQuota() *HostUnitQuota {
	this := HostUnitQuota{}
	return &this
}

// NewHostUnitQuotaWithDefaults instantiates a new HostUnitQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostUnitQuotaWithDefaults() *HostUnitQuota {
	this := HostUnitQuota{}
	return &this
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *HostUnitQuota) GetCurrentUsage() float64 {
	if o == nil || o.CurrentUsage == nil {
		var ret float64
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUnitQuota) GetCurrentUsageOk() (*float64, bool) {
	if o == nil || o.CurrentUsage == nil {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *HostUnitQuota) HasCurrentUsage() bool {
	if o != nil && o.CurrentUsage != nil {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given float64 and assigns it to the CurrentUsage field.
func (o *HostUnitQuota) SetCurrentUsage(v float64) {
	o.CurrentUsage = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *HostUnitQuota) GetMaxLimit() int64 {
	if o == nil || o.MaxLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUnitQuota) GetMaxLimitOk() (*int64, bool) {
	if o == nil || o.MaxLimit == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *HostUnitQuota) HasMaxLimit() bool {
	if o != nil && o.MaxLimit != nil {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given int64 and assigns it to the MaxLimit field.
func (o *HostUnitQuota) SetMaxLimit(v int64) {
	o.MaxLimit = &v
}

func (o HostUnitQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentUsage != nil {
		toSerialize["currentUsage"] = o.CurrentUsage
	}
	if o.MaxLimit != nil {
		toSerialize["maxLimit"] = o.MaxLimit
	}
	return json.Marshal(toSerialize)
}

type NullableHostUnitQuota struct {
	value *HostUnitQuota
	isSet bool
}

func (v NullableHostUnitQuota) Get() *HostUnitQuota {
	return v.value
}

func (v *NullableHostUnitQuota) Set(val *HostUnitQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableHostUnitQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableHostUnitQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostUnitQuota(val *HostUnitQuota) *NullableHostUnitQuota {
	return &NullableHostUnitQuota{value: val, isSet: true}
}

func (v NullableHostUnitQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostUnitQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

