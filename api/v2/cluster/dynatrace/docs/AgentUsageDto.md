# AgentUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkTraffic** | Pointer to **int64** |  | [optional] 
**AgentId** | Pointer to **int32** |  | [optional] 
**AgentTypeId** | Pointer to **int32** |  | [optional] 
**AgentUsageRecords** | Pointer to [**[]AgentUsageRecordDto**](AgentUsageRecordDto.md) |  | [optional] 

## Methods

### NewAgentUsageDto

`func NewAgentUsageDto() *AgentUsageDto`

NewAgentUsageDto instantiates a new AgentUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUsageDtoWithDefaults

`func NewAgentUsageDtoWithDefaults() *AgentUsageDto`

NewAgentUsageDtoWithDefaults instantiates a new AgentUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkTraffic

`func (o *AgentUsageDto) GetNetworkTraffic() int64`

GetNetworkTraffic returns the NetworkTraffic field if non-nil, zero value otherwise.

### GetNetworkTrafficOk

`func (o *AgentUsageDto) GetNetworkTrafficOk() (*int64, bool)`

GetNetworkTrafficOk returns a tuple with the NetworkTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTraffic

`func (o *AgentUsageDto) SetNetworkTraffic(v int64)`

SetNetworkTraffic sets NetworkTraffic field to given value.

### HasNetworkTraffic

`func (o *AgentUsageDto) HasNetworkTraffic() bool`

HasNetworkTraffic returns a boolean if a field has been set.

### GetAgentId

`func (o *AgentUsageDto) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentUsageDto) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentUsageDto) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentUsageDto) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentTypeId

`func (o *AgentUsageDto) GetAgentTypeId() int32`

GetAgentTypeId returns the AgentTypeId field if non-nil, zero value otherwise.

### GetAgentTypeIdOk

`func (o *AgentUsageDto) GetAgentTypeIdOk() (*int32, bool)`

GetAgentTypeIdOk returns a tuple with the AgentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTypeId

`func (o *AgentUsageDto) SetAgentTypeId(v int32)`

SetAgentTypeId sets AgentTypeId field to given value.

### HasAgentTypeId

`func (o *AgentUsageDto) HasAgentTypeId() bool`

HasAgentTypeId returns a boolean if a field has been set.

### GetAgentUsageRecords

`func (o *AgentUsageDto) GetAgentUsageRecords() []AgentUsageRecordDto`

GetAgentUsageRecords returns the AgentUsageRecords field if non-nil, zero value otherwise.

### GetAgentUsageRecordsOk

`func (o *AgentUsageDto) GetAgentUsageRecordsOk() (*[]AgentUsageRecordDto, bool)`

GetAgentUsageRecordsOk returns a tuple with the AgentUsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUsageRecords

`func (o *AgentUsageDto) SetAgentUsageRecords(v []AgentUsageRecordDto)`

SetAgentUsageRecords sets AgentUsageRecords field to given value.

### HasAgentUsageRecords

`func (o *AgentUsageDto) HasAgentUsageRecords() bool`

HasAgentUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


