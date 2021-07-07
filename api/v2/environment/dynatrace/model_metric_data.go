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

// MetricData A list of metrics and their data points.
type MetricData struct {
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey string `json:"nextPageKey"`
	// The total number of primary entities in the result.   Has the `0` value if none of the requested metrics is suitable for pagination.
	TotalCount int64 `json:"totalCount"`
	// A list of warnings
	Warnings *[]string `json:"warnings,omitempty"`
	// A list of metrics and their data points.
	Result *[]MetricSeriesCollection `json:"result,omitempty"`
}

// NewMetricData instantiates a new MetricData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricData(nextPageKey string, totalCount int64) *MetricData {
	this := MetricData{}
	this.NextPageKey = nextPageKey
	this.TotalCount = totalCount
	return &this
}

// NewMetricDataWithDefaults instantiates a new MetricData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDataWithDefaults() *MetricData {
	this := MetricData{}
	return &this
}

// GetNextPageKey returns the NextPageKey field value
func (o *MetricData) GetNextPageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value
// and a boolean to check if the value has been set.
func (o *MetricData) GetNextPageKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextPageKey, true
}

// SetNextPageKey sets field value
func (o *MetricData) SetNextPageKey(v string) {
	o.NextPageKey = v
}

// GetTotalCount returns the TotalCount field value
func (o *MetricData) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *MetricData) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *MetricData) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MetricData) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricData) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MetricData) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *MetricData) SetWarnings(v []string) {
	o.Warnings = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *MetricData) GetResult() []MetricSeriesCollection {
	if o == nil || o.Result == nil {
		var ret []MetricSeriesCollection
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricData) GetResultOk() (*[]MetricSeriesCollection, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *MetricData) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []MetricSeriesCollection and assigns it to the Result field.
func (o *MetricData) SetResult(v []MetricSeriesCollection) {
	o.Result = &v
}

func (o MetricData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableMetricData struct {
	value *MetricData
	isSet bool
}

func (v NullableMetricData) Get() *MetricData {
	return v.value
}

func (v *NullableMetricData) Set(val *MetricData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricData(val *MetricData) *NullableMetricData {
	return &NullableMetricData{value: val, isSet: true}
}

func (v NullableMetricData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

