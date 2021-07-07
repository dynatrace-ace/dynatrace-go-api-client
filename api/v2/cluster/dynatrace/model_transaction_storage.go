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

// TransactionStorage Transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain.
type TransactionStorage struct {
	// Percentage of truncation for new data.
	RetentionReductionPercentage *string `json:"retentionReductionPercentage,omitempty"`
	// Reason of truncation.
	RetentionReductionReason *string `json:"retentionReductionReason,omitempty"`
	// Currently used storage [bytes]
	CurrentlyUsed *int64 `json:"currentlyUsed,omitempty"`
	// Maximum storage limit [bytes]
	MaxLimit *int64 `json:"maxLimit,omitempty"`
}

// NewTransactionStorage instantiates a new TransactionStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStorage() *TransactionStorage {
	this := TransactionStorage{}
	return &this
}

// NewTransactionStorageWithDefaults instantiates a new TransactionStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStorageWithDefaults() *TransactionStorage {
	this := TransactionStorage{}
	return &this
}

// GetRetentionReductionPercentage returns the RetentionReductionPercentage field value if set, zero value otherwise.
func (o *TransactionStorage) GetRetentionReductionPercentage() string {
	if o == nil || o.RetentionReductionPercentage == nil {
		var ret string
		return ret
	}
	return *o.RetentionReductionPercentage
}

// GetRetentionReductionPercentageOk returns a tuple with the RetentionReductionPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStorage) GetRetentionReductionPercentageOk() (*string, bool) {
	if o == nil || o.RetentionReductionPercentage == nil {
		return nil, false
	}
	return o.RetentionReductionPercentage, true
}

// HasRetentionReductionPercentage returns a boolean if a field has been set.
func (o *TransactionStorage) HasRetentionReductionPercentage() bool {
	if o != nil && o.RetentionReductionPercentage != nil {
		return true
	}

	return false
}

// SetRetentionReductionPercentage gets a reference to the given string and assigns it to the RetentionReductionPercentage field.
func (o *TransactionStorage) SetRetentionReductionPercentage(v string) {
	o.RetentionReductionPercentage = &v
}

// GetRetentionReductionReason returns the RetentionReductionReason field value if set, zero value otherwise.
func (o *TransactionStorage) GetRetentionReductionReason() string {
	if o == nil || o.RetentionReductionReason == nil {
		var ret string
		return ret
	}
	return *o.RetentionReductionReason
}

// GetRetentionReductionReasonOk returns a tuple with the RetentionReductionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStorage) GetRetentionReductionReasonOk() (*string, bool) {
	if o == nil || o.RetentionReductionReason == nil {
		return nil, false
	}
	return o.RetentionReductionReason, true
}

// HasRetentionReductionReason returns a boolean if a field has been set.
func (o *TransactionStorage) HasRetentionReductionReason() bool {
	if o != nil && o.RetentionReductionReason != nil {
		return true
	}

	return false
}

// SetRetentionReductionReason gets a reference to the given string and assigns it to the RetentionReductionReason field.
func (o *TransactionStorage) SetRetentionReductionReason(v string) {
	o.RetentionReductionReason = &v
}

// GetCurrentlyUsed returns the CurrentlyUsed field value if set, zero value otherwise.
func (o *TransactionStorage) GetCurrentlyUsed() int64 {
	if o == nil || o.CurrentlyUsed == nil {
		var ret int64
		return ret
	}
	return *o.CurrentlyUsed
}

// GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStorage) GetCurrentlyUsedOk() (*int64, bool) {
	if o == nil || o.CurrentlyUsed == nil {
		return nil, false
	}
	return o.CurrentlyUsed, true
}

// HasCurrentlyUsed returns a boolean if a field has been set.
func (o *TransactionStorage) HasCurrentlyUsed() bool {
	if o != nil && o.CurrentlyUsed != nil {
		return true
	}

	return false
}

// SetCurrentlyUsed gets a reference to the given int64 and assigns it to the CurrentlyUsed field.
func (o *TransactionStorage) SetCurrentlyUsed(v int64) {
	o.CurrentlyUsed = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *TransactionStorage) GetMaxLimit() int64 {
	if o == nil || o.MaxLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStorage) GetMaxLimitOk() (*int64, bool) {
	if o == nil || o.MaxLimit == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *TransactionStorage) HasMaxLimit() bool {
	if o != nil && o.MaxLimit != nil {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given int64 and assigns it to the MaxLimit field.
func (o *TransactionStorage) SetMaxLimit(v int64) {
	o.MaxLimit = &v
}

func (o TransactionStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RetentionReductionPercentage != nil {
		toSerialize["retentionReductionPercentage"] = o.RetentionReductionPercentage
	}
	if o.RetentionReductionReason != nil {
		toSerialize["retentionReductionReason"] = o.RetentionReductionReason
	}
	if o.CurrentlyUsed != nil {
		toSerialize["currentlyUsed"] = o.CurrentlyUsed
	}
	if o.MaxLimit != nil {
		toSerialize["maxLimit"] = o.MaxLimit
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionStorage struct {
	value *TransactionStorage
	isSet bool
}

func (v NullableTransactionStorage) Get() *TransactionStorage {
	return v.value
}

func (v *NullableTransactionStorage) Set(val *TransactionStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStorage(val *TransactionStorage) *NullableTransactionStorage {
	return &NullableTransactionStorage{value: val, isSet: true}
}

func (v NullableTransactionStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


