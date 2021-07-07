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

// ClusterNodesConfigDto struct for ClusterNodesConfigDto
type ClusterNodesConfigDto struct {
	ClusterNodes *[]NodeConfigDto `json:"clusterNodes,omitempty"`
}

// NewClusterNodesConfigDto instantiates a new ClusterNodesConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodesConfigDto() *ClusterNodesConfigDto {
	this := ClusterNodesConfigDto{}
	return &this
}

// NewClusterNodesConfigDtoWithDefaults instantiates a new ClusterNodesConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodesConfigDtoWithDefaults() *ClusterNodesConfigDto {
	this := ClusterNodesConfigDto{}
	return &this
}

// GetClusterNodes returns the ClusterNodes field value if set, zero value otherwise.
func (o *ClusterNodesConfigDto) GetClusterNodes() []NodeConfigDto {
	if o == nil || o.ClusterNodes == nil {
		var ret []NodeConfigDto
		return ret
	}
	return *o.ClusterNodes
}

// GetClusterNodesOk returns a tuple with the ClusterNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodesConfigDto) GetClusterNodesOk() (*[]NodeConfigDto, bool) {
	if o == nil || o.ClusterNodes == nil {
		return nil, false
	}
	return o.ClusterNodes, true
}

// HasClusterNodes returns a boolean if a field has been set.
func (o *ClusterNodesConfigDto) HasClusterNodes() bool {
	if o != nil && o.ClusterNodes != nil {
		return true
	}

	return false
}

// SetClusterNodes gets a reference to the given []NodeConfigDto and assigns it to the ClusterNodes field.
func (o *ClusterNodesConfigDto) SetClusterNodes(v []NodeConfigDto) {
	o.ClusterNodes = &v
}

func (o ClusterNodesConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterNodes != nil {
		toSerialize["clusterNodes"] = o.ClusterNodes
	}
	return json.Marshal(toSerialize)
}

type NullableClusterNodesConfigDto struct {
	value *ClusterNodesConfigDto
	isSet bool
}

func (v NullableClusterNodesConfigDto) Get() *ClusterNodesConfigDto {
	return v.value
}

func (v *NullableClusterNodesConfigDto) Set(val *ClusterNodesConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodesConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodesConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodesConfigDto(val *ClusterNodesConfigDto) *NullableClusterNodesConfigDto {
	return &NullableClusterNodesConfigDto{value: val, isSet: true}
}

func (v NullableClusterNodesConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodesConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

