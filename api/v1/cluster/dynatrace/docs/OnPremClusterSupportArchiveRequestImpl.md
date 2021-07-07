# OnPremClusterSupportArchiveRequestImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfDays** | Pointer to **int32** |  | [optional] 
**IncludeServerData** | Pointer to **bool** |  | [optional] 
**IncludeActiveGateData** | Pointer to **bool** |  | [optional] 
**IncludeLogs** | Pointer to **bool** |  | [optional] 
**IncludeDebugLogs** | Pointer to **bool** |  | [optional] 
**IncludeAgentRegistryLogs** | Pointer to **bool** |  | [optional] 
**IncludeLauncherLogs** | Pointer to **bool** |  | [optional] 
**IncludeMonitoringConfigAuditLogs** | Pointer to **bool** |  | [optional] 
**IncludeOtherAuditLogs** | Pointer to **bool** |  | [optional] 

## Methods

### NewOnPremClusterSupportArchiveRequestImpl

`func NewOnPremClusterSupportArchiveRequestImpl() *OnPremClusterSupportArchiveRequestImpl`

NewOnPremClusterSupportArchiveRequestImpl instantiates a new OnPremClusterSupportArchiveRequestImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremClusterSupportArchiveRequestImplWithDefaults

`func NewOnPremClusterSupportArchiveRequestImplWithDefaults() *OnPremClusterSupportArchiveRequestImpl`

NewOnPremClusterSupportArchiveRequestImplWithDefaults instantiates a new OnPremClusterSupportArchiveRequestImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfDays

`func (o *OnPremClusterSupportArchiveRequestImpl) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *OnPremClusterSupportArchiveRequestImpl) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.

### HasNumberOfDays

`func (o *OnPremClusterSupportArchiveRequestImpl) HasNumberOfDays() bool`

HasNumberOfDays returns a boolean if a field has been set.

### GetIncludeServerData

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeServerData() bool`

GetIncludeServerData returns the IncludeServerData field if non-nil, zero value otherwise.

### GetIncludeServerDataOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeServerDataOk() (*bool, bool)`

GetIncludeServerDataOk returns a tuple with the IncludeServerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeServerData

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeServerData(v bool)`

SetIncludeServerData sets IncludeServerData field to given value.

### HasIncludeServerData

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeServerData() bool`

HasIncludeServerData returns a boolean if a field has been set.

### GetIncludeActiveGateData

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeActiveGateData() bool`

GetIncludeActiveGateData returns the IncludeActiveGateData field if non-nil, zero value otherwise.

### GetIncludeActiveGateDataOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeActiveGateDataOk() (*bool, bool)`

GetIncludeActiveGateDataOk returns a tuple with the IncludeActiveGateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActiveGateData

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeActiveGateData(v bool)`

SetIncludeActiveGateData sets IncludeActiveGateData field to given value.

### HasIncludeActiveGateData

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeActiveGateData() bool`

HasIncludeActiveGateData returns a boolean if a field has been set.

### GetIncludeLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeLogs() bool`

GetIncludeLogs returns the IncludeLogs field if non-nil, zero value otherwise.

### GetIncludeLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeLogsOk() (*bool, bool)`

GetIncludeLogsOk returns a tuple with the IncludeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeLogs(v bool)`

SetIncludeLogs sets IncludeLogs field to given value.

### HasIncludeLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeLogs() bool`

HasIncludeLogs returns a boolean if a field has been set.

### GetIncludeDebugLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeDebugLogs() bool`

GetIncludeDebugLogs returns the IncludeDebugLogs field if non-nil, zero value otherwise.

### GetIncludeDebugLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeDebugLogsOk() (*bool, bool)`

GetIncludeDebugLogsOk returns a tuple with the IncludeDebugLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDebugLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeDebugLogs(v bool)`

SetIncludeDebugLogs sets IncludeDebugLogs field to given value.

### HasIncludeDebugLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeDebugLogs() bool`

HasIncludeDebugLogs returns a boolean if a field has been set.

### GetIncludeAgentRegistryLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeAgentRegistryLogs() bool`

GetIncludeAgentRegistryLogs returns the IncludeAgentRegistryLogs field if non-nil, zero value otherwise.

### GetIncludeAgentRegistryLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeAgentRegistryLogsOk() (*bool, bool)`

GetIncludeAgentRegistryLogsOk returns a tuple with the IncludeAgentRegistryLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAgentRegistryLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeAgentRegistryLogs(v bool)`

SetIncludeAgentRegistryLogs sets IncludeAgentRegistryLogs field to given value.

### HasIncludeAgentRegistryLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeAgentRegistryLogs() bool`

HasIncludeAgentRegistryLogs returns a boolean if a field has been set.

### GetIncludeLauncherLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeLauncherLogs() bool`

GetIncludeLauncherLogs returns the IncludeLauncherLogs field if non-nil, zero value otherwise.

### GetIncludeLauncherLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeLauncherLogsOk() (*bool, bool)`

GetIncludeLauncherLogsOk returns a tuple with the IncludeLauncherLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLauncherLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeLauncherLogs(v bool)`

SetIncludeLauncherLogs sets IncludeLauncherLogs field to given value.

### HasIncludeLauncherLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeLauncherLogs() bool`

HasIncludeLauncherLogs returns a boolean if a field has been set.

### GetIncludeMonitoringConfigAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeMonitoringConfigAuditLogs() bool`

GetIncludeMonitoringConfigAuditLogs returns the IncludeMonitoringConfigAuditLogs field if non-nil, zero value otherwise.

### GetIncludeMonitoringConfigAuditLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeMonitoringConfigAuditLogsOk() (*bool, bool)`

GetIncludeMonitoringConfigAuditLogsOk returns a tuple with the IncludeMonitoringConfigAuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMonitoringConfigAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeMonitoringConfigAuditLogs(v bool)`

SetIncludeMonitoringConfigAuditLogs sets IncludeMonitoringConfigAuditLogs field to given value.

### HasIncludeMonitoringConfigAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeMonitoringConfigAuditLogs() bool`

HasIncludeMonitoringConfigAuditLogs returns a boolean if a field has been set.

### GetIncludeOtherAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeOtherAuditLogs() bool`

GetIncludeOtherAuditLogs returns the IncludeOtherAuditLogs field if non-nil, zero value otherwise.

### GetIncludeOtherAuditLogsOk

`func (o *OnPremClusterSupportArchiveRequestImpl) GetIncludeOtherAuditLogsOk() (*bool, bool)`

GetIncludeOtherAuditLogsOk returns a tuple with the IncludeOtherAuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOtherAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) SetIncludeOtherAuditLogs(v bool)`

SetIncludeOtherAuditLogs sets IncludeOtherAuditLogs field to given value.

### HasIncludeOtherAuditLogs

`func (o *OnPremClusterSupportArchiveRequestImpl) HasIncludeOtherAuditLogs() bool`

HasIncludeOtherAuditLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


