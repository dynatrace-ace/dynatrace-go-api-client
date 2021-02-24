/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// DashboardList A list of short representations of dashboards.
type DashboardList struct {
	// A list of short representations of dashboards.
	Dashboards []DashboardStub `json:"dashboards"`
}

// NewDashboardList instantiates a new DashboardList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardList(dashboards []DashboardStub, ) *DashboardList {
	this := DashboardList{}
	this.Dashboards = dashboards
	return &this
}

// NewDashboardListWithDefaults instantiates a new DashboardList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardListWithDefaults() *DashboardList {
	this := DashboardList{}
	return &this
}

// GetDashboards returns the Dashboards field value
func (o *DashboardList) GetDashboards() []DashboardStub {
	if o == nil  {
		var ret []DashboardStub
		return ret
	}

	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *DashboardList) GetDashboardsOk() (*[]DashboardStub, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dashboards, true
}

// SetDashboards sets field value
func (o *DashboardList) SetDashboards(v []DashboardStub) {
	o.Dashboards = v
}

func (o DashboardList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboards"] = o.Dashboards
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardList struct {
	value *DashboardList
	isSet bool
}

func (v NullableDashboardList) Get() *DashboardList {
	return v.value
}

func (v *NullableDashboardList) Set(val *DashboardList) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardList) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardList(val *DashboardList) *NullableDashboardList {
	return &NullableDashboardList{value: val, isSet: true}
}

func (v NullableDashboardList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


