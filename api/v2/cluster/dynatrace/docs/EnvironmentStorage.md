# EnvironmentStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionStorage** | Pointer to [**TransactionStorage**](TransactionStorage.md) |  | [optional] 
**SessionReplayStorage** | Pointer to [**SessionReplayStorage**](SessionReplayStorage.md) |  | [optional] 
**SymbolFilesFromMobileApps** | Pointer to [**SymbolFilesFromMobileApps**](SymbolFilesFromMobileApps.md) |  | [optional] 
**LogMonitoringStorage** | Pointer to [**LogMonitoringStorage**](LogMonitoringStorage.md) |  | [optional] 
**ServiceRequestLevelRetention** | Pointer to [**ServiceRequestLevelRetention**](ServiceRequestLevelRetention.md) |  | [optional] 
**ServiceCodeLevelRetention** | Pointer to [**ServiceCodeLevelRetention**](ServiceCodeLevelRetention.md) |  | [optional] 
**RealUserMonitoringRetention** | Pointer to [**RealUserMonitoringRetention**](RealUserMonitoringRetention.md) |  | [optional] 
**SyntheticMonitoringRetention** | Pointer to [**SyntheticMonitoringRetention**](SyntheticMonitoringRetention.md) |  | [optional] 
**SessionReplayRetention** | Pointer to [**SessionReplayRetention**](SessionReplayRetention.md) |  | [optional] 
**LogMonitoringRetention** | Pointer to [**LogMonitoringRetention**](LogMonitoringRetention.md) |  | [optional] 
**UserActionsPerMinute** | Pointer to [**UserActionsPerMinute**](UserActionsPerMinute.md) |  | [optional] 
**TransactionTrafficQuota** | Pointer to [**TransactionTrafficQuota**](TransactionTrafficQuota.md) |  | [optional] 

## Methods

### NewEnvironmentStorage

`func NewEnvironmentStorage() *EnvironmentStorage`

NewEnvironmentStorage instantiates a new EnvironmentStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStorageWithDefaults

`func NewEnvironmentStorageWithDefaults() *EnvironmentStorage`

NewEnvironmentStorageWithDefaults instantiates a new EnvironmentStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionStorage

`func (o *EnvironmentStorage) GetTransactionStorage() TransactionStorage`

GetTransactionStorage returns the TransactionStorage field if non-nil, zero value otherwise.

### GetTransactionStorageOk

`func (o *EnvironmentStorage) GetTransactionStorageOk() (*TransactionStorage, bool)`

GetTransactionStorageOk returns a tuple with the TransactionStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStorage

`func (o *EnvironmentStorage) SetTransactionStorage(v TransactionStorage)`

SetTransactionStorage sets TransactionStorage field to given value.

### HasTransactionStorage

`func (o *EnvironmentStorage) HasTransactionStorage() bool`

HasTransactionStorage returns a boolean if a field has been set.

### GetSessionReplayStorage

`func (o *EnvironmentStorage) GetSessionReplayStorage() SessionReplayStorage`

GetSessionReplayStorage returns the SessionReplayStorage field if non-nil, zero value otherwise.

### GetSessionReplayStorageOk

`func (o *EnvironmentStorage) GetSessionReplayStorageOk() (*SessionReplayStorage, bool)`

GetSessionReplayStorageOk returns a tuple with the SessionReplayStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayStorage

`func (o *EnvironmentStorage) SetSessionReplayStorage(v SessionReplayStorage)`

SetSessionReplayStorage sets SessionReplayStorage field to given value.

### HasSessionReplayStorage

`func (o *EnvironmentStorage) HasSessionReplayStorage() bool`

HasSessionReplayStorage returns a boolean if a field has been set.

### GetSymbolFilesFromMobileApps

`func (o *EnvironmentStorage) GetSymbolFilesFromMobileApps() SymbolFilesFromMobileApps`

GetSymbolFilesFromMobileApps returns the SymbolFilesFromMobileApps field if non-nil, zero value otherwise.

### GetSymbolFilesFromMobileAppsOk

`func (o *EnvironmentStorage) GetSymbolFilesFromMobileAppsOk() (*SymbolFilesFromMobileApps, bool)`

GetSymbolFilesFromMobileAppsOk returns a tuple with the SymbolFilesFromMobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolFilesFromMobileApps

`func (o *EnvironmentStorage) SetSymbolFilesFromMobileApps(v SymbolFilesFromMobileApps)`

SetSymbolFilesFromMobileApps sets SymbolFilesFromMobileApps field to given value.

### HasSymbolFilesFromMobileApps

`func (o *EnvironmentStorage) HasSymbolFilesFromMobileApps() bool`

HasSymbolFilesFromMobileApps returns a boolean if a field has been set.

### GetLogMonitoringStorage

`func (o *EnvironmentStorage) GetLogMonitoringStorage() LogMonitoringStorage`

GetLogMonitoringStorage returns the LogMonitoringStorage field if non-nil, zero value otherwise.

### GetLogMonitoringStorageOk

`func (o *EnvironmentStorage) GetLogMonitoringStorageOk() (*LogMonitoringStorage, bool)`

GetLogMonitoringStorageOk returns a tuple with the LogMonitoringStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMonitoringStorage

`func (o *EnvironmentStorage) SetLogMonitoringStorage(v LogMonitoringStorage)`

SetLogMonitoringStorage sets LogMonitoringStorage field to given value.

### HasLogMonitoringStorage

`func (o *EnvironmentStorage) HasLogMonitoringStorage() bool`

HasLogMonitoringStorage returns a boolean if a field has been set.

### GetServiceRequestLevelRetention

`func (o *EnvironmentStorage) GetServiceRequestLevelRetention() ServiceRequestLevelRetention`

GetServiceRequestLevelRetention returns the ServiceRequestLevelRetention field if non-nil, zero value otherwise.

### GetServiceRequestLevelRetentionOk

`func (o *EnvironmentStorage) GetServiceRequestLevelRetentionOk() (*ServiceRequestLevelRetention, bool)`

GetServiceRequestLevelRetentionOk returns a tuple with the ServiceRequestLevelRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestLevelRetention

`func (o *EnvironmentStorage) SetServiceRequestLevelRetention(v ServiceRequestLevelRetention)`

SetServiceRequestLevelRetention sets ServiceRequestLevelRetention field to given value.

### HasServiceRequestLevelRetention

`func (o *EnvironmentStorage) HasServiceRequestLevelRetention() bool`

HasServiceRequestLevelRetention returns a boolean if a field has been set.

### GetServiceCodeLevelRetention

`func (o *EnvironmentStorage) GetServiceCodeLevelRetention() ServiceCodeLevelRetention`

GetServiceCodeLevelRetention returns the ServiceCodeLevelRetention field if non-nil, zero value otherwise.

### GetServiceCodeLevelRetentionOk

`func (o *EnvironmentStorage) GetServiceCodeLevelRetentionOk() (*ServiceCodeLevelRetention, bool)`

GetServiceCodeLevelRetentionOk returns a tuple with the ServiceCodeLevelRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCodeLevelRetention

`func (o *EnvironmentStorage) SetServiceCodeLevelRetention(v ServiceCodeLevelRetention)`

SetServiceCodeLevelRetention sets ServiceCodeLevelRetention field to given value.

### HasServiceCodeLevelRetention

`func (o *EnvironmentStorage) HasServiceCodeLevelRetention() bool`

HasServiceCodeLevelRetention returns a boolean if a field has been set.

### GetRealUserMonitoringRetention

`func (o *EnvironmentStorage) GetRealUserMonitoringRetention() RealUserMonitoringRetention`

GetRealUserMonitoringRetention returns the RealUserMonitoringRetention field if non-nil, zero value otherwise.

### GetRealUserMonitoringRetentionOk

`func (o *EnvironmentStorage) GetRealUserMonitoringRetentionOk() (*RealUserMonitoringRetention, bool)`

GetRealUserMonitoringRetentionOk returns a tuple with the RealUserMonitoringRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealUserMonitoringRetention

`func (o *EnvironmentStorage) SetRealUserMonitoringRetention(v RealUserMonitoringRetention)`

SetRealUserMonitoringRetention sets RealUserMonitoringRetention field to given value.

### HasRealUserMonitoringRetention

`func (o *EnvironmentStorage) HasRealUserMonitoringRetention() bool`

HasRealUserMonitoringRetention returns a boolean if a field has been set.

### GetSyntheticMonitoringRetention

`func (o *EnvironmentStorage) GetSyntheticMonitoringRetention() SyntheticMonitoringRetention`

GetSyntheticMonitoringRetention returns the SyntheticMonitoringRetention field if non-nil, zero value otherwise.

### GetSyntheticMonitoringRetentionOk

`func (o *EnvironmentStorage) GetSyntheticMonitoringRetentionOk() (*SyntheticMonitoringRetention, bool)`

GetSyntheticMonitoringRetentionOk returns a tuple with the SyntheticMonitoringRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMonitoringRetention

`func (o *EnvironmentStorage) SetSyntheticMonitoringRetention(v SyntheticMonitoringRetention)`

SetSyntheticMonitoringRetention sets SyntheticMonitoringRetention field to given value.

### HasSyntheticMonitoringRetention

`func (o *EnvironmentStorage) HasSyntheticMonitoringRetention() bool`

HasSyntheticMonitoringRetention returns a boolean if a field has been set.

### GetSessionReplayRetention

`func (o *EnvironmentStorage) GetSessionReplayRetention() SessionReplayRetention`

GetSessionReplayRetention returns the SessionReplayRetention field if non-nil, zero value otherwise.

### GetSessionReplayRetentionOk

`func (o *EnvironmentStorage) GetSessionReplayRetentionOk() (*SessionReplayRetention, bool)`

GetSessionReplayRetentionOk returns a tuple with the SessionReplayRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayRetention

`func (o *EnvironmentStorage) SetSessionReplayRetention(v SessionReplayRetention)`

SetSessionReplayRetention sets SessionReplayRetention field to given value.

### HasSessionReplayRetention

`func (o *EnvironmentStorage) HasSessionReplayRetention() bool`

HasSessionReplayRetention returns a boolean if a field has been set.

### GetLogMonitoringRetention

`func (o *EnvironmentStorage) GetLogMonitoringRetention() LogMonitoringRetention`

GetLogMonitoringRetention returns the LogMonitoringRetention field if non-nil, zero value otherwise.

### GetLogMonitoringRetentionOk

`func (o *EnvironmentStorage) GetLogMonitoringRetentionOk() (*LogMonitoringRetention, bool)`

GetLogMonitoringRetentionOk returns a tuple with the LogMonitoringRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMonitoringRetention

`func (o *EnvironmentStorage) SetLogMonitoringRetention(v LogMonitoringRetention)`

SetLogMonitoringRetention sets LogMonitoringRetention field to given value.

### HasLogMonitoringRetention

`func (o *EnvironmentStorage) HasLogMonitoringRetention() bool`

HasLogMonitoringRetention returns a boolean if a field has been set.

### GetUserActionsPerMinute

`func (o *EnvironmentStorage) GetUserActionsPerMinute() UserActionsPerMinute`

GetUserActionsPerMinute returns the UserActionsPerMinute field if non-nil, zero value otherwise.

### GetUserActionsPerMinuteOk

`func (o *EnvironmentStorage) GetUserActionsPerMinuteOk() (*UserActionsPerMinute, bool)`

GetUserActionsPerMinuteOk returns a tuple with the UserActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsPerMinute

`func (o *EnvironmentStorage) SetUserActionsPerMinute(v UserActionsPerMinute)`

SetUserActionsPerMinute sets UserActionsPerMinute field to given value.

### HasUserActionsPerMinute

`func (o *EnvironmentStorage) HasUserActionsPerMinute() bool`

HasUserActionsPerMinute returns a boolean if a field has been set.

### GetTransactionTrafficQuota

`func (o *EnvironmentStorage) GetTransactionTrafficQuota() TransactionTrafficQuota`

GetTransactionTrafficQuota returns the TransactionTrafficQuota field if non-nil, zero value otherwise.

### GetTransactionTrafficQuotaOk

`func (o *EnvironmentStorage) GetTransactionTrafficQuotaOk() (*TransactionTrafficQuota, bool)`

GetTransactionTrafficQuotaOk returns a tuple with the TransactionTrafficQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTrafficQuota

`func (o *EnvironmentStorage) SetTransactionTrafficQuota(v TransactionTrafficQuota)`

SetTransactionTrafficQuota sets TransactionTrafficQuota field to given value.

### HasTransactionTrafficQuota

`func (o *EnvironmentStorage) HasTransactionTrafficQuota() bool`

HasTransactionTrafficQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


