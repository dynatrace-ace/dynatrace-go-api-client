# HostUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsiId** | Pointer to **int64** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**HostCategory** | Pointer to **string** |  | [optional] 
**AgentUsages** | Pointer to [**[]AgentUsageDto**](AgentUsageDto.md) |  | [optional] 
**InfrastructureOnly** | Pointer to **bool** |  | [optional] 
**Paas** | Pointer to **bool** |  | [optional] 
**PassMemoryLimit** | Pointer to **int64** |  | [optional] 
**VendorTypeId** | Pointer to **int32** |  | [optional] 
**HostMemoryBytes** | Pointer to **int64** |  | [optional] 
**PremiumLogAnalytics** | Pointer to **bool** |  | [optional] 
**HasContainers** | Pointer to **bool** |  | [optional] 

## Methods

### NewHostUsageDto

`func NewHostUsageDto() *HostUsageDto`

NewHostUsageDto instantiates a new HostUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUsageDtoWithDefaults

`func NewHostUsageDtoWithDefaults() *HostUsageDto`

NewHostUsageDtoWithDefaults instantiates a new HostUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsiId

`func (o *HostUsageDto) GetOsiId() int64`

GetOsiId returns the OsiId field if non-nil, zero value otherwise.

### GetOsiIdOk

`func (o *HostUsageDto) GetOsiIdOk() (*int64, bool)`

GetOsiIdOk returns a tuple with the OsiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsiId

`func (o *HostUsageDto) SetOsiId(v int64)`

SetOsiId sets OsiId field to given value.

### HasOsiId

`func (o *HostUsageDto) HasOsiId() bool`

HasOsiId returns a boolean if a field has been set.

### GetHostName

`func (o *HostUsageDto) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HostUsageDto) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HostUsageDto) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HostUsageDto) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostCategory

`func (o *HostUsageDto) GetHostCategory() string`

GetHostCategory returns the HostCategory field if non-nil, zero value otherwise.

### GetHostCategoryOk

`func (o *HostUsageDto) GetHostCategoryOk() (*string, bool)`

GetHostCategoryOk returns a tuple with the HostCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCategory

`func (o *HostUsageDto) SetHostCategory(v string)`

SetHostCategory sets HostCategory field to given value.

### HasHostCategory

`func (o *HostUsageDto) HasHostCategory() bool`

HasHostCategory returns a boolean if a field has been set.

### GetAgentUsages

`func (o *HostUsageDto) GetAgentUsages() []AgentUsageDto`

GetAgentUsages returns the AgentUsages field if non-nil, zero value otherwise.

### GetAgentUsagesOk

`func (o *HostUsageDto) GetAgentUsagesOk() (*[]AgentUsageDto, bool)`

GetAgentUsagesOk returns a tuple with the AgentUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUsages

`func (o *HostUsageDto) SetAgentUsages(v []AgentUsageDto)`

SetAgentUsages sets AgentUsages field to given value.

### HasAgentUsages

`func (o *HostUsageDto) HasAgentUsages() bool`

HasAgentUsages returns a boolean if a field has been set.

### GetInfrastructureOnly

`func (o *HostUsageDto) GetInfrastructureOnly() bool`

GetInfrastructureOnly returns the InfrastructureOnly field if non-nil, zero value otherwise.

### GetInfrastructureOnlyOk

`func (o *HostUsageDto) GetInfrastructureOnlyOk() (*bool, bool)`

GetInfrastructureOnlyOk returns a tuple with the InfrastructureOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureOnly

`func (o *HostUsageDto) SetInfrastructureOnly(v bool)`

SetInfrastructureOnly sets InfrastructureOnly field to given value.

### HasInfrastructureOnly

`func (o *HostUsageDto) HasInfrastructureOnly() bool`

HasInfrastructureOnly returns a boolean if a field has been set.

### GetPaas

`func (o *HostUsageDto) GetPaas() bool`

GetPaas returns the Paas field if non-nil, zero value otherwise.

### GetPaasOk

`func (o *HostUsageDto) GetPaasOk() (*bool, bool)`

GetPaasOk returns a tuple with the Paas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaas

`func (o *HostUsageDto) SetPaas(v bool)`

SetPaas sets Paas field to given value.

### HasPaas

`func (o *HostUsageDto) HasPaas() bool`

HasPaas returns a boolean if a field has been set.

### GetPassMemoryLimit

`func (o *HostUsageDto) GetPassMemoryLimit() int64`

GetPassMemoryLimit returns the PassMemoryLimit field if non-nil, zero value otherwise.

### GetPassMemoryLimitOk

`func (o *HostUsageDto) GetPassMemoryLimitOk() (*int64, bool)`

GetPassMemoryLimitOk returns a tuple with the PassMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassMemoryLimit

`func (o *HostUsageDto) SetPassMemoryLimit(v int64)`

SetPassMemoryLimit sets PassMemoryLimit field to given value.

### HasPassMemoryLimit

`func (o *HostUsageDto) HasPassMemoryLimit() bool`

HasPassMemoryLimit returns a boolean if a field has been set.

### GetVendorTypeId

`func (o *HostUsageDto) GetVendorTypeId() int32`

GetVendorTypeId returns the VendorTypeId field if non-nil, zero value otherwise.

### GetVendorTypeIdOk

`func (o *HostUsageDto) GetVendorTypeIdOk() (*int32, bool)`

GetVendorTypeIdOk returns a tuple with the VendorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorTypeId

`func (o *HostUsageDto) SetVendorTypeId(v int32)`

SetVendorTypeId sets VendorTypeId field to given value.

### HasVendorTypeId

`func (o *HostUsageDto) HasVendorTypeId() bool`

HasVendorTypeId returns a boolean if a field has been set.

### GetHostMemoryBytes

`func (o *HostUsageDto) GetHostMemoryBytes() int64`

GetHostMemoryBytes returns the HostMemoryBytes field if non-nil, zero value otherwise.

### GetHostMemoryBytesOk

`func (o *HostUsageDto) GetHostMemoryBytesOk() (*int64, bool)`

GetHostMemoryBytesOk returns a tuple with the HostMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMemoryBytes

`func (o *HostUsageDto) SetHostMemoryBytes(v int64)`

SetHostMemoryBytes sets HostMemoryBytes field to given value.

### HasHostMemoryBytes

`func (o *HostUsageDto) HasHostMemoryBytes() bool`

HasHostMemoryBytes returns a boolean if a field has been set.

### GetPremiumLogAnalytics

`func (o *HostUsageDto) GetPremiumLogAnalytics() bool`

GetPremiumLogAnalytics returns the PremiumLogAnalytics field if non-nil, zero value otherwise.

### GetPremiumLogAnalyticsOk

`func (o *HostUsageDto) GetPremiumLogAnalyticsOk() (*bool, bool)`

GetPremiumLogAnalyticsOk returns a tuple with the PremiumLogAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumLogAnalytics

`func (o *HostUsageDto) SetPremiumLogAnalytics(v bool)`

SetPremiumLogAnalytics sets PremiumLogAnalytics field to given value.

### HasPremiumLogAnalytics

`func (o *HostUsageDto) HasPremiumLogAnalytics() bool`

HasPremiumLogAnalytics returns a boolean if a field has been set.

### GetHasContainers

`func (o *HostUsageDto) GetHasContainers() bool`

GetHasContainers returns the HasContainers field if non-nil, zero value otherwise.

### GetHasContainersOk

`func (o *HostUsageDto) GetHasContainersOk() (*bool, bool)`

GetHasContainersOk returns a tuple with the HasContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasContainers

`func (o *HostUsageDto) SetHasContainers(v bool)`

SetHasContainers sets HasContainers field to given value.

### HasHasContainers

`func (o *HostUsageDto) HasHasContainers() bool`

HasHasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


