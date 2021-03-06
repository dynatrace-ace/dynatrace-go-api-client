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

// DatacenterDesc Map of multidatacenter topology
type DatacenterDesc struct {
	DatacenterMembersList *[]DatacenterMember `json:"datacenterMembersList,omitempty"`
	NewDatacenter *bool `json:"newDatacenter,omitempty"`
}

// NewDatacenterDesc instantiates a new DatacenterDesc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacenterDesc() *DatacenterDesc {
	this := DatacenterDesc{}
	return &this
}

// NewDatacenterDescWithDefaults instantiates a new DatacenterDesc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacenterDescWithDefaults() *DatacenterDesc {
	this := DatacenterDesc{}
	return &this
}

// GetDatacenterMembersList returns the DatacenterMembersList field value if set, zero value otherwise.
func (o *DatacenterDesc) GetDatacenterMembersList() []DatacenterMember {
	if o == nil || o.DatacenterMembersList == nil {
		var ret []DatacenterMember
		return ret
	}
	return *o.DatacenterMembersList
}

// GetDatacenterMembersListOk returns a tuple with the DatacenterMembersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterDesc) GetDatacenterMembersListOk() (*[]DatacenterMember, bool) {
	if o == nil || o.DatacenterMembersList == nil {
		return nil, false
	}
	return o.DatacenterMembersList, true
}

// HasDatacenterMembersList returns a boolean if a field has been set.
func (o *DatacenterDesc) HasDatacenterMembersList() bool {
	if o != nil && o.DatacenterMembersList != nil {
		return true
	}

	return false
}

// SetDatacenterMembersList gets a reference to the given []DatacenterMember and assigns it to the DatacenterMembersList field.
func (o *DatacenterDesc) SetDatacenterMembersList(v []DatacenterMember) {
	o.DatacenterMembersList = &v
}

// GetNewDatacenter returns the NewDatacenter field value if set, zero value otherwise.
func (o *DatacenterDesc) GetNewDatacenter() bool {
	if o == nil || o.NewDatacenter == nil {
		var ret bool
		return ret
	}
	return *o.NewDatacenter
}

// GetNewDatacenterOk returns a tuple with the NewDatacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterDesc) GetNewDatacenterOk() (*bool, bool) {
	if o == nil || o.NewDatacenter == nil {
		return nil, false
	}
	return o.NewDatacenter, true
}

// HasNewDatacenter returns a boolean if a field has been set.
func (o *DatacenterDesc) HasNewDatacenter() bool {
	if o != nil && o.NewDatacenter != nil {
		return true
	}

	return false
}

// SetNewDatacenter gets a reference to the given bool and assigns it to the NewDatacenter field.
func (o *DatacenterDesc) SetNewDatacenter(v bool) {
	o.NewDatacenter = &v
}

func (o DatacenterDesc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatacenterMembersList != nil {
		toSerialize["datacenterMembersList"] = o.DatacenterMembersList
	}
	if o.NewDatacenter != nil {
		toSerialize["newDatacenter"] = o.NewDatacenter
	}
	return json.Marshal(toSerialize)
}

type NullableDatacenterDesc struct {
	value *DatacenterDesc
	isSet bool
}

func (v NullableDatacenterDesc) Get() *DatacenterDesc {
	return v.value
}

func (v *NullableDatacenterDesc) Set(val *DatacenterDesc) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenterDesc) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenterDesc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenterDesc(val *DatacenterDesc) *NullableDatacenterDesc {
	return &NullableDatacenterDesc{value: val, isSet: true}
}

func (v NullableDatacenterDesc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenterDesc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


