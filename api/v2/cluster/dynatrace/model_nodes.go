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

// Nodes A list of synthetic nodes
type Nodes struct {
	// A list of synthetic nodes
	Nodes []NodeCollectionElement `json:"nodes"`
}

// NewNodes instantiates a new Nodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodes(nodes []NodeCollectionElement) *Nodes {
	this := Nodes{}
	this.Nodes = nodes
	return &this
}

// NewNodesWithDefaults instantiates a new Nodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodesWithDefaults() *Nodes {
	this := Nodes{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *Nodes) GetNodes() []NodeCollectionElement {
	if o == nil {
		var ret []NodeCollectionElement
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *Nodes) GetNodesOk() (*[]NodeCollectionElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *Nodes) SetNodes(v []NodeCollectionElement) {
	o.Nodes = v
}

func (o Nodes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	return json.Marshal(toSerialize)
}

type NullableNodes struct {
	value *Nodes
	isSet bool
}

func (v NullableNodes) Get() *Nodes {
	return v.value
}

func (v *NullableNodes) Set(val *Nodes) {
	v.value = val
	v.isSet = true
}

func (v NullableNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodes(val *Nodes) *NullableNodes {
	return &NullableNodes{value: val, isSet: true}
}

func (v NullableNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


