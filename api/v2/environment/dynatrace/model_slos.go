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

// SLOs A list of SLOs.
type SLOs struct {
	// A list of SLOs.
	Slo *[]SLO `json:"slo,omitempty"`
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// The total number of entries in the result.
	TotalCount int64 `json:"totalCount"`
	// The number of entries per page.
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewSLOs instantiates a new SLOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOs(totalCount int64) *SLOs {
	this := SLOs{}
	this.TotalCount = totalCount
	return &this
}

// NewSLOsWithDefaults instantiates a new SLOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOsWithDefaults() *SLOs {
	this := SLOs{}
	return &this
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *SLOs) GetSlo() []SLO {
	if o == nil || o.Slo == nil {
		var ret []SLO
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOs) GetSloOk() (*[]SLO, bool) {
	if o == nil || o.Slo == nil {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *SLOs) HasSlo() bool {
	if o != nil && o.Slo != nil {
		return true
	}

	return false
}

// SetSlo gets a reference to the given []SLO and assigns it to the Slo field.
func (o *SLOs) SetSlo(v []SLO) {
	o.Slo = &v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *SLOs) GetNextPageKey() string {
	if o == nil || o.NextPageKey == nil {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOs) GetNextPageKeyOk() (*string, bool) {
	if o == nil || o.NextPageKey == nil {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *SLOs) HasNextPageKey() bool {
	if o != nil && o.NextPageKey != nil {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *SLOs) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetTotalCount returns the TotalCount field value
func (o *SLOs) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SLOs) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SLOs) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SLOs) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOs) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SLOs) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *SLOs) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o SLOs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Slo != nil {
		toSerialize["slo"] = o.Slo
	}
	if o.NextPageKey != nil {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableSLOs struct {
	value *SLOs
	isSet bool
}

func (v NullableSLOs) Get() *SLOs {
	return v.value
}

func (v *NullableSLOs) Set(val *SLOs) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOs) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOs(val *SLOs) *NullableSLOs {
	return &NullableSLOs{value: val, isSet: true}
}

func (v NullableSLOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


