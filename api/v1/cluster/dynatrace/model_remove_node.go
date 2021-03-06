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

// RemoveNode Information which node is going to be removed.
type RemoveNode struct {
	NodeId *int32 `json:"nodeId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
}

// NewRemoveNode instantiates a new RemoveNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveNode() *RemoveNode {
	this := RemoveNode{}
	return &this
}

// NewRemoveNodeWithDefaults instantiates a new RemoveNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveNodeWithDefaults() *RemoveNode {
	this := RemoveNode{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RemoveNode) GetNodeId() int32 {
	if o == nil || o.NodeId == nil {
		var ret int32
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveNode) GetNodeIdOk() (*int32, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RemoveNode) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int32 and assigns it to the NodeId field.
func (o *RemoveNode) SetNodeId(v int32) {
	o.NodeId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *RemoveNode) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveNode) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *RemoveNode) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *RemoveNode) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o RemoveNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveNode struct {
	value *RemoveNode
	isSet bool
}

func (v NullableRemoveNode) Get() *RemoveNode {
	return v.value
}

func (v *NullableRemoveNode) Set(val *RemoveNode) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveNode) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveNode(val *RemoveNode) *NullableRemoveNode {
	return &NullableRemoveNode{value: val, isSet: true}
}

func (v NullableRemoveNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


