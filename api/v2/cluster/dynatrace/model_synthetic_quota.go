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

// SyntheticQuota Synthetic monitors consumption and quota information on environment level. Not set (and not editable) if neither Synthetic nor DEM units is enabled. If skipped when editing via PUT method then already set quotas will remain.
type SyntheticQuota struct {
	// Monthly environment consumption. Resets each calendar month.
	ConsumedThisMonth *float64 `json:"consumedThisMonth,omitempty"`
	// Yearly environment consumption. Resets each year on license creation date anniversary.
	ConsumedThisYear *float64 `json:"consumedThisYear,omitempty"`
	// Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.
	MonthlyLimit *int64 `json:"monthlyLimit,omitempty"`
	// Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited.
	AnnualLimit *int64 `json:"annualLimit,omitempty"`
}

// NewSyntheticQuota instantiates a new SyntheticQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticQuota() *SyntheticQuota {
	this := SyntheticQuota{}
	return &this
}

// NewSyntheticQuotaWithDefaults instantiates a new SyntheticQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticQuotaWithDefaults() *SyntheticQuota {
	this := SyntheticQuota{}
	return &this
}

// GetConsumedThisMonth returns the ConsumedThisMonth field value if set, zero value otherwise.
func (o *SyntheticQuota) GetConsumedThisMonth() float64 {
	if o == nil || o.ConsumedThisMonth == nil {
		var ret float64
		return ret
	}
	return *o.ConsumedThisMonth
}

// GetConsumedThisMonthOk returns a tuple with the ConsumedThisMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticQuota) GetConsumedThisMonthOk() (*float64, bool) {
	if o == nil || o.ConsumedThisMonth == nil {
		return nil, false
	}
	return o.ConsumedThisMonth, true
}

// HasConsumedThisMonth returns a boolean if a field has been set.
func (o *SyntheticQuota) HasConsumedThisMonth() bool {
	if o != nil && o.ConsumedThisMonth != nil {
		return true
	}

	return false
}

// SetConsumedThisMonth gets a reference to the given float64 and assigns it to the ConsumedThisMonth field.
func (o *SyntheticQuota) SetConsumedThisMonth(v float64) {
	o.ConsumedThisMonth = &v
}

// GetConsumedThisYear returns the ConsumedThisYear field value if set, zero value otherwise.
func (o *SyntheticQuota) GetConsumedThisYear() float64 {
	if o == nil || o.ConsumedThisYear == nil {
		var ret float64
		return ret
	}
	return *o.ConsumedThisYear
}

// GetConsumedThisYearOk returns a tuple with the ConsumedThisYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticQuota) GetConsumedThisYearOk() (*float64, bool) {
	if o == nil || o.ConsumedThisYear == nil {
		return nil, false
	}
	return o.ConsumedThisYear, true
}

// HasConsumedThisYear returns a boolean if a field has been set.
func (o *SyntheticQuota) HasConsumedThisYear() bool {
	if o != nil && o.ConsumedThisYear != nil {
		return true
	}

	return false
}

// SetConsumedThisYear gets a reference to the given float64 and assigns it to the ConsumedThisYear field.
func (o *SyntheticQuota) SetConsumedThisYear(v float64) {
	o.ConsumedThisYear = &v
}

// GetMonthlyLimit returns the MonthlyLimit field value if set, zero value otherwise.
func (o *SyntheticQuota) GetMonthlyLimit() int64 {
	if o == nil || o.MonthlyLimit == nil {
		var ret int64
		return ret
	}
	return *o.MonthlyLimit
}

// GetMonthlyLimitOk returns a tuple with the MonthlyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticQuota) GetMonthlyLimitOk() (*int64, bool) {
	if o == nil || o.MonthlyLimit == nil {
		return nil, false
	}
	return o.MonthlyLimit, true
}

// HasMonthlyLimit returns a boolean if a field has been set.
func (o *SyntheticQuota) HasMonthlyLimit() bool {
	if o != nil && o.MonthlyLimit != nil {
		return true
	}

	return false
}

// SetMonthlyLimit gets a reference to the given int64 and assigns it to the MonthlyLimit field.
func (o *SyntheticQuota) SetMonthlyLimit(v int64) {
	o.MonthlyLimit = &v
}

// GetAnnualLimit returns the AnnualLimit field value if set, zero value otherwise.
func (o *SyntheticQuota) GetAnnualLimit() int64 {
	if o == nil || o.AnnualLimit == nil {
		var ret int64
		return ret
	}
	return *o.AnnualLimit
}

// GetAnnualLimitOk returns a tuple with the AnnualLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticQuota) GetAnnualLimitOk() (*int64, bool) {
	if o == nil || o.AnnualLimit == nil {
		return nil, false
	}
	return o.AnnualLimit, true
}

// HasAnnualLimit returns a boolean if a field has been set.
func (o *SyntheticQuota) HasAnnualLimit() bool {
	if o != nil && o.AnnualLimit != nil {
		return true
	}

	return false
}

// SetAnnualLimit gets a reference to the given int64 and assigns it to the AnnualLimit field.
func (o *SyntheticQuota) SetAnnualLimit(v int64) {
	o.AnnualLimit = &v
}

func (o SyntheticQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumedThisMonth != nil {
		toSerialize["consumedThisMonth"] = o.ConsumedThisMonth
	}
	if o.ConsumedThisYear != nil {
		toSerialize["consumedThisYear"] = o.ConsumedThisYear
	}
	if o.MonthlyLimit != nil {
		toSerialize["monthlyLimit"] = o.MonthlyLimit
	}
	if o.AnnualLimit != nil {
		toSerialize["annualLimit"] = o.AnnualLimit
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticQuota struct {
	value *SyntheticQuota
	isSet bool
}

func (v NullableSyntheticQuota) Get() *SyntheticQuota {
	return v.value
}

func (v *NullableSyntheticQuota) Set(val *SyntheticQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticQuota(val *SyntheticQuota) *NullableSyntheticQuota {
	return &NullableSyntheticQuota{value: val, isSet: true}
}

func (v NullableSyntheticQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


