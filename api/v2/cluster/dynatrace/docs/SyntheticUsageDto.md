# SyntheticUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorDefinitionId** | Pointer to **string** |  | [optional] 
**MonitorDescription** | Pointer to **string** |  | [optional] 
**MonitorTypeId** | Pointer to **int32** |  | [optional] 
**SuccessCount** | Pointer to **int32** |  | [optional] 
**FailureCount** | Pointer to **int32** |  | [optional] 
**SyntheticActionCount** | Pointer to **int32** |  | [optional] 
**PerformedSyntheticActions** | Pointer to **int32** |  | [optional] 

## Methods

### NewSyntheticUsageDto

`func NewSyntheticUsageDto() *SyntheticUsageDto`

NewSyntheticUsageDto instantiates a new SyntheticUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticUsageDtoWithDefaults

`func NewSyntheticUsageDtoWithDefaults() *SyntheticUsageDto`

NewSyntheticUsageDtoWithDefaults instantiates a new SyntheticUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorDefinitionId

`func (o *SyntheticUsageDto) GetMonitorDefinitionId() string`

GetMonitorDefinitionId returns the MonitorDefinitionId field if non-nil, zero value otherwise.

### GetMonitorDefinitionIdOk

`func (o *SyntheticUsageDto) GetMonitorDefinitionIdOk() (*string, bool)`

GetMonitorDefinitionIdOk returns a tuple with the MonitorDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDefinitionId

`func (o *SyntheticUsageDto) SetMonitorDefinitionId(v string)`

SetMonitorDefinitionId sets MonitorDefinitionId field to given value.

### HasMonitorDefinitionId

`func (o *SyntheticUsageDto) HasMonitorDefinitionId() bool`

HasMonitorDefinitionId returns a boolean if a field has been set.

### GetMonitorDescription

`func (o *SyntheticUsageDto) GetMonitorDescription() string`

GetMonitorDescription returns the MonitorDescription field if non-nil, zero value otherwise.

### GetMonitorDescriptionOk

`func (o *SyntheticUsageDto) GetMonitorDescriptionOk() (*string, bool)`

GetMonitorDescriptionOk returns a tuple with the MonitorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDescription

`func (o *SyntheticUsageDto) SetMonitorDescription(v string)`

SetMonitorDescription sets MonitorDescription field to given value.

### HasMonitorDescription

`func (o *SyntheticUsageDto) HasMonitorDescription() bool`

HasMonitorDescription returns a boolean if a field has been set.

### GetMonitorTypeId

`func (o *SyntheticUsageDto) GetMonitorTypeId() int32`

GetMonitorTypeId returns the MonitorTypeId field if non-nil, zero value otherwise.

### GetMonitorTypeIdOk

`func (o *SyntheticUsageDto) GetMonitorTypeIdOk() (*int32, bool)`

GetMonitorTypeIdOk returns a tuple with the MonitorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTypeId

`func (o *SyntheticUsageDto) SetMonitorTypeId(v int32)`

SetMonitorTypeId sets MonitorTypeId field to given value.

### HasMonitorTypeId

`func (o *SyntheticUsageDto) HasMonitorTypeId() bool`

HasMonitorTypeId returns a boolean if a field has been set.

### GetSuccessCount

`func (o *SyntheticUsageDto) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *SyntheticUsageDto) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *SyntheticUsageDto) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *SyntheticUsageDto) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetFailureCount

`func (o *SyntheticUsageDto) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *SyntheticUsageDto) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *SyntheticUsageDto) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *SyntheticUsageDto) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetSyntheticActionCount

`func (o *SyntheticUsageDto) GetSyntheticActionCount() int32`

GetSyntheticActionCount returns the SyntheticActionCount field if non-nil, zero value otherwise.

### GetSyntheticActionCountOk

`func (o *SyntheticUsageDto) GetSyntheticActionCountOk() (*int32, bool)`

GetSyntheticActionCountOk returns a tuple with the SyntheticActionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticActionCount

`func (o *SyntheticUsageDto) SetSyntheticActionCount(v int32)`

SetSyntheticActionCount sets SyntheticActionCount field to given value.

### HasSyntheticActionCount

`func (o *SyntheticUsageDto) HasSyntheticActionCount() bool`

HasSyntheticActionCount returns a boolean if a field has been set.

### GetPerformedSyntheticActions

`func (o *SyntheticUsageDto) GetPerformedSyntheticActions() int32`

GetPerformedSyntheticActions returns the PerformedSyntheticActions field if non-nil, zero value otherwise.

### GetPerformedSyntheticActionsOk

`func (o *SyntheticUsageDto) GetPerformedSyntheticActionsOk() (*int32, bool)`

GetPerformedSyntheticActionsOk returns a tuple with the PerformedSyntheticActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedSyntheticActions

`func (o *SyntheticUsageDto) SetPerformedSyntheticActions(v int32)`

SetPerformedSyntheticActions sets PerformedSyntheticActions field to given value.

### HasPerformedSyntheticActions

`func (o *SyntheticUsageDto) HasPerformedSyntheticActions() bool`

HasPerformedSyntheticActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


