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

// AgentUsageDto struct for AgentUsageDto
type AgentUsageDto struct {
	NetworkTraffic *int64 `json:"networkTraffic,omitempty"`
	AgentId *int32 `json:"agentId,omitempty"`
	AgentTypeId *int32 `json:"agentTypeId,omitempty"`
	AgentUsageRecords *[]AgentUsageRecordDto `json:"agentUsageRecords,omitempty"`
}

// NewAgentUsageDto instantiates a new AgentUsageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUsageDto() *AgentUsageDto {
	this := AgentUsageDto{}
	return &this
}

// NewAgentUsageDtoWithDefaults instantiates a new AgentUsageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUsageDtoWithDefaults() *AgentUsageDto {
	this := AgentUsageDto{}
	return &this
}

// GetNetworkTraffic returns the NetworkTraffic field value if set, zero value otherwise.
func (o *AgentUsageDto) GetNetworkTraffic() int64 {
	if o == nil || o.NetworkTraffic == nil {
		var ret int64
		return ret
	}
	return *o.NetworkTraffic
}

// GetNetworkTrafficOk returns a tuple with the NetworkTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUsageDto) GetNetworkTrafficOk() (*int64, bool) {
	if o == nil || o.NetworkTraffic == nil {
		return nil, false
	}
	return o.NetworkTraffic, true
}

// HasNetworkTraffic returns a boolean if a field has been set.
func (o *AgentUsageDto) HasNetworkTraffic() bool {
	if o != nil && o.NetworkTraffic != nil {
		return true
	}

	return false
}

// SetNetworkTraffic gets a reference to the given int64 and assigns it to the NetworkTraffic field.
func (o *AgentUsageDto) SetNetworkTraffic(v int64) {
	o.NetworkTraffic = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentUsageDto) GetAgentId() int32 {
	if o == nil || o.AgentId == nil {
		var ret int32
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUsageDto) GetAgentIdOk() (*int32, bool) {
	if o == nil || o.AgentId == nil {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentUsageDto) HasAgentId() bool {
	if o != nil && o.AgentId != nil {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given int32 and assigns it to the AgentId field.
func (o *AgentUsageDto) SetAgentId(v int32) {
	o.AgentId = &v
}

// GetAgentTypeId returns the AgentTypeId field value if set, zero value otherwise.
func (o *AgentUsageDto) GetAgentTypeId() int32 {
	if o == nil || o.AgentTypeId == nil {
		var ret int32
		return ret
	}
	return *o.AgentTypeId
}

// GetAgentTypeIdOk returns a tuple with the AgentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUsageDto) GetAgentTypeIdOk() (*int32, bool) {
	if o == nil || o.AgentTypeId == nil {
		return nil, false
	}
	return o.AgentTypeId, true
}

// HasAgentTypeId returns a boolean if a field has been set.
func (o *AgentUsageDto) HasAgentTypeId() bool {
	if o != nil && o.AgentTypeId != nil {
		return true
	}

	return false
}

// SetAgentTypeId gets a reference to the given int32 and assigns it to the AgentTypeId field.
func (o *AgentUsageDto) SetAgentTypeId(v int32) {
	o.AgentTypeId = &v
}

// GetAgentUsageRecords returns the AgentUsageRecords field value if set, zero value otherwise.
func (o *AgentUsageDto) GetAgentUsageRecords() []AgentUsageRecordDto {
	if o == nil || o.AgentUsageRecords == nil {
		var ret []AgentUsageRecordDto
		return ret
	}
	return *o.AgentUsageRecords
}

// GetAgentUsageRecordsOk returns a tuple with the AgentUsageRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUsageDto) GetAgentUsageRecordsOk() (*[]AgentUsageRecordDto, bool) {
	if o == nil || o.AgentUsageRecords == nil {
		return nil, false
	}
	return o.AgentUsageRecords, true
}

// HasAgentUsageRecords returns a boolean if a field has been set.
func (o *AgentUsageDto) HasAgentUsageRecords() bool {
	if o != nil && o.AgentUsageRecords != nil {
		return true
	}

	return false
}

// SetAgentUsageRecords gets a reference to the given []AgentUsageRecordDto and assigns it to the AgentUsageRecords field.
func (o *AgentUsageDto) SetAgentUsageRecords(v []AgentUsageRecordDto) {
	o.AgentUsageRecords = &v
}

func (o AgentUsageDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkTraffic != nil {
		toSerialize["networkTraffic"] = o.NetworkTraffic
	}
	if o.AgentId != nil {
		toSerialize["agentId"] = o.AgentId
	}
	if o.AgentTypeId != nil {
		toSerialize["agentTypeId"] = o.AgentTypeId
	}
	if o.AgentUsageRecords != nil {
		toSerialize["agentUsageRecords"] = o.AgentUsageRecords
	}
	return json.Marshal(toSerialize)
}

type NullableAgentUsageDto struct {
	value *AgentUsageDto
	isSet bool
}

func (v NullableAgentUsageDto) Get() *AgentUsageDto {
	return v.value
}

func (v *NullableAgentUsageDto) Set(val *AgentUsageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUsageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUsageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUsageDto(val *AgentUsageDto) *NullableAgentUsageDto {
	return &NullableAgentUsageDto{value: val, isSet: true}
}

func (v NullableAgentUsageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUsageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


