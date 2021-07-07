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

// HostUsageDto struct for HostUsageDto
type HostUsageDto struct {
	OsiId *int64 `json:"osiId,omitempty"`
	HostName *string `json:"hostName,omitempty"`
	HostCategory *string `json:"hostCategory,omitempty"`
	AgentUsages *[]AgentUsageDto `json:"agentUsages,omitempty"`
	InfrastructureOnly *bool `json:"infrastructureOnly,omitempty"`
	Paas *bool `json:"paas,omitempty"`
	PassMemoryLimit *int64 `json:"passMemoryLimit,omitempty"`
	VendorTypeId *int32 `json:"vendorTypeId,omitempty"`
	HostMemoryBytes *int64 `json:"hostMemoryBytes,omitempty"`
	PremiumLogAnalytics *bool `json:"premiumLogAnalytics,omitempty"`
	HasContainers *bool `json:"hasContainers,omitempty"`
}

// NewHostUsageDto instantiates a new HostUsageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostUsageDto() *HostUsageDto {
	this := HostUsageDto{}
	return &this
}

// NewHostUsageDtoWithDefaults instantiates a new HostUsageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostUsageDtoWithDefaults() *HostUsageDto {
	this := HostUsageDto{}
	return &this
}

// GetOsiId returns the OsiId field value if set, zero value otherwise.
func (o *HostUsageDto) GetOsiId() int64 {
	if o == nil || o.OsiId == nil {
		var ret int64
		return ret
	}
	return *o.OsiId
}

// GetOsiIdOk returns a tuple with the OsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetOsiIdOk() (*int64, bool) {
	if o == nil || o.OsiId == nil {
		return nil, false
	}
	return o.OsiId, true
}

// HasOsiId returns a boolean if a field has been set.
func (o *HostUsageDto) HasOsiId() bool {
	if o != nil && o.OsiId != nil {
		return true
	}

	return false
}

// SetOsiId gets a reference to the given int64 and assigns it to the OsiId field.
func (o *HostUsageDto) SetOsiId(v int64) {
	o.OsiId = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HostUsageDto) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HostUsageDto) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HostUsageDto) SetHostName(v string) {
	o.HostName = &v
}

// GetHostCategory returns the HostCategory field value if set, zero value otherwise.
func (o *HostUsageDto) GetHostCategory() string {
	if o == nil || o.HostCategory == nil {
		var ret string
		return ret
	}
	return *o.HostCategory
}

// GetHostCategoryOk returns a tuple with the HostCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetHostCategoryOk() (*string, bool) {
	if o == nil || o.HostCategory == nil {
		return nil, false
	}
	return o.HostCategory, true
}

// HasHostCategory returns a boolean if a field has been set.
func (o *HostUsageDto) HasHostCategory() bool {
	if o != nil && o.HostCategory != nil {
		return true
	}

	return false
}

// SetHostCategory gets a reference to the given string and assigns it to the HostCategory field.
func (o *HostUsageDto) SetHostCategory(v string) {
	o.HostCategory = &v
}

// GetAgentUsages returns the AgentUsages field value if set, zero value otherwise.
func (o *HostUsageDto) GetAgentUsages() []AgentUsageDto {
	if o == nil || o.AgentUsages == nil {
		var ret []AgentUsageDto
		return ret
	}
	return *o.AgentUsages
}

// GetAgentUsagesOk returns a tuple with the AgentUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetAgentUsagesOk() (*[]AgentUsageDto, bool) {
	if o == nil || o.AgentUsages == nil {
		return nil, false
	}
	return o.AgentUsages, true
}

// HasAgentUsages returns a boolean if a field has been set.
func (o *HostUsageDto) HasAgentUsages() bool {
	if o != nil && o.AgentUsages != nil {
		return true
	}

	return false
}

// SetAgentUsages gets a reference to the given []AgentUsageDto and assigns it to the AgentUsages field.
func (o *HostUsageDto) SetAgentUsages(v []AgentUsageDto) {
	o.AgentUsages = &v
}

// GetInfrastructureOnly returns the InfrastructureOnly field value if set, zero value otherwise.
func (o *HostUsageDto) GetInfrastructureOnly() bool {
	if o == nil || o.InfrastructureOnly == nil {
		var ret bool
		return ret
	}
	return *o.InfrastructureOnly
}

// GetInfrastructureOnlyOk returns a tuple with the InfrastructureOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetInfrastructureOnlyOk() (*bool, bool) {
	if o == nil || o.InfrastructureOnly == nil {
		return nil, false
	}
	return o.InfrastructureOnly, true
}

// HasInfrastructureOnly returns a boolean if a field has been set.
func (o *HostUsageDto) HasInfrastructureOnly() bool {
	if o != nil && o.InfrastructureOnly != nil {
		return true
	}

	return false
}

// SetInfrastructureOnly gets a reference to the given bool and assigns it to the InfrastructureOnly field.
func (o *HostUsageDto) SetInfrastructureOnly(v bool) {
	o.InfrastructureOnly = &v
}

// GetPaas returns the Paas field value if set, zero value otherwise.
func (o *HostUsageDto) GetPaas() bool {
	if o == nil || o.Paas == nil {
		var ret bool
		return ret
	}
	return *o.Paas
}

// GetPaasOk returns a tuple with the Paas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetPaasOk() (*bool, bool) {
	if o == nil || o.Paas == nil {
		return nil, false
	}
	return o.Paas, true
}

// HasPaas returns a boolean if a field has been set.
func (o *HostUsageDto) HasPaas() bool {
	if o != nil && o.Paas != nil {
		return true
	}

	return false
}

// SetPaas gets a reference to the given bool and assigns it to the Paas field.
func (o *HostUsageDto) SetPaas(v bool) {
	o.Paas = &v
}

// GetPassMemoryLimit returns the PassMemoryLimit field value if set, zero value otherwise.
func (o *HostUsageDto) GetPassMemoryLimit() int64 {
	if o == nil || o.PassMemoryLimit == nil {
		var ret int64
		return ret
	}
	return *o.PassMemoryLimit
}

// GetPassMemoryLimitOk returns a tuple with the PassMemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetPassMemoryLimitOk() (*int64, bool) {
	if o == nil || o.PassMemoryLimit == nil {
		return nil, false
	}
	return o.PassMemoryLimit, true
}

// HasPassMemoryLimit returns a boolean if a field has been set.
func (o *HostUsageDto) HasPassMemoryLimit() bool {
	if o != nil && o.PassMemoryLimit != nil {
		return true
	}

	return false
}

// SetPassMemoryLimit gets a reference to the given int64 and assigns it to the PassMemoryLimit field.
func (o *HostUsageDto) SetPassMemoryLimit(v int64) {
	o.PassMemoryLimit = &v
}

// GetVendorTypeId returns the VendorTypeId field value if set, zero value otherwise.
func (o *HostUsageDto) GetVendorTypeId() int32 {
	if o == nil || o.VendorTypeId == nil {
		var ret int32
		return ret
	}
	return *o.VendorTypeId
}

// GetVendorTypeIdOk returns a tuple with the VendorTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetVendorTypeIdOk() (*int32, bool) {
	if o == nil || o.VendorTypeId == nil {
		return nil, false
	}
	return o.VendorTypeId, true
}

// HasVendorTypeId returns a boolean if a field has been set.
func (o *HostUsageDto) HasVendorTypeId() bool {
	if o != nil && o.VendorTypeId != nil {
		return true
	}

	return false
}

// SetVendorTypeId gets a reference to the given int32 and assigns it to the VendorTypeId field.
func (o *HostUsageDto) SetVendorTypeId(v int32) {
	o.VendorTypeId = &v
}

// GetHostMemoryBytes returns the HostMemoryBytes field value if set, zero value otherwise.
func (o *HostUsageDto) GetHostMemoryBytes() int64 {
	if o == nil || o.HostMemoryBytes == nil {
		var ret int64
		return ret
	}
	return *o.HostMemoryBytes
}

// GetHostMemoryBytesOk returns a tuple with the HostMemoryBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetHostMemoryBytesOk() (*int64, bool) {
	if o == nil || o.HostMemoryBytes == nil {
		return nil, false
	}
	return o.HostMemoryBytes, true
}

// HasHostMemoryBytes returns a boolean if a field has been set.
func (o *HostUsageDto) HasHostMemoryBytes() bool {
	if o != nil && o.HostMemoryBytes != nil {
		return true
	}

	return false
}

// SetHostMemoryBytes gets a reference to the given int64 and assigns it to the HostMemoryBytes field.
func (o *HostUsageDto) SetHostMemoryBytes(v int64) {
	o.HostMemoryBytes = &v
}

// GetPremiumLogAnalytics returns the PremiumLogAnalytics field value if set, zero value otherwise.
func (o *HostUsageDto) GetPremiumLogAnalytics() bool {
	if o == nil || o.PremiumLogAnalytics == nil {
		var ret bool
		return ret
	}
	return *o.PremiumLogAnalytics
}

// GetPremiumLogAnalyticsOk returns a tuple with the PremiumLogAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetPremiumLogAnalyticsOk() (*bool, bool) {
	if o == nil || o.PremiumLogAnalytics == nil {
		return nil, false
	}
	return o.PremiumLogAnalytics, true
}

// HasPremiumLogAnalytics returns a boolean if a field has been set.
func (o *HostUsageDto) HasPremiumLogAnalytics() bool {
	if o != nil && o.PremiumLogAnalytics != nil {
		return true
	}

	return false
}

// SetPremiumLogAnalytics gets a reference to the given bool and assigns it to the PremiumLogAnalytics field.
func (o *HostUsageDto) SetPremiumLogAnalytics(v bool) {
	o.PremiumLogAnalytics = &v
}

// GetHasContainers returns the HasContainers field value if set, zero value otherwise.
func (o *HostUsageDto) GetHasContainers() bool {
	if o == nil || o.HasContainers == nil {
		var ret bool
		return ret
	}
	return *o.HasContainers
}

// GetHasContainersOk returns a tuple with the HasContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostUsageDto) GetHasContainersOk() (*bool, bool) {
	if o == nil || o.HasContainers == nil {
		return nil, false
	}
	return o.HasContainers, true
}

// HasHasContainers returns a boolean if a field has been set.
func (o *HostUsageDto) HasHasContainers() bool {
	if o != nil && o.HasContainers != nil {
		return true
	}

	return false
}

// SetHasContainers gets a reference to the given bool and assigns it to the HasContainers field.
func (o *HostUsageDto) SetHasContainers(v bool) {
	o.HasContainers = &v
}

func (o HostUsageDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OsiId != nil {
		toSerialize["osiId"] = o.OsiId
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.HostCategory != nil {
		toSerialize["hostCategory"] = o.HostCategory
	}
	if o.AgentUsages != nil {
		toSerialize["agentUsages"] = o.AgentUsages
	}
	if o.InfrastructureOnly != nil {
		toSerialize["infrastructureOnly"] = o.InfrastructureOnly
	}
	if o.Paas != nil {
		toSerialize["paas"] = o.Paas
	}
	if o.PassMemoryLimit != nil {
		toSerialize["passMemoryLimit"] = o.PassMemoryLimit
	}
	if o.VendorTypeId != nil {
		toSerialize["vendorTypeId"] = o.VendorTypeId
	}
	if o.HostMemoryBytes != nil {
		toSerialize["hostMemoryBytes"] = o.HostMemoryBytes
	}
	if o.PremiumLogAnalytics != nil {
		toSerialize["premiumLogAnalytics"] = o.PremiumLogAnalytics
	}
	if o.HasContainers != nil {
		toSerialize["hasContainers"] = o.HasContainers
	}
	return json.Marshal(toSerialize)
}

type NullableHostUsageDto struct {
	value *HostUsageDto
	isSet bool
}

func (v NullableHostUsageDto) Get() *HostUsageDto {
	return v.value
}

func (v *NullableHostUsageDto) Set(val *HostUsageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHostUsageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHostUsageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostUsageDto(val *HostUsageDto) *NullableHostUsageDto {
	return &NullableHostUsageDto{value: val, isSet: true}
}

func (v NullableHostUsageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostUsageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

