# BackupConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Backups are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] [readonly] 
**Datacenter** | Pointer to **string** | Datacenter which will create backups | [optional] [readonly] 
**StoragePath** | Pointer to **string** | A full path to the backup archive | [optional] [readonly] 
**IncludeRumData** | Pointer to **bool** | Include user sessions (&#x60;true&#x60;) or GDPR compliance (&#x60;false&#x60;) | [optional] 
**IncludeTsMetricData** | Pointer to **bool** | Include time series metric-data (&#x60;true&#x60;) or retain configuration data only (&#x60;false&#x60;)) | [optional] 
**BandwidthLimitMbits** | Pointer to **int32** | Cassandra backup bandwidth limit in Mbps | [optional] 
**MaxEsSnapshotsToClean** | Pointer to **int32** | Max number of Elasticsearch snapshots to clean. Elasticsearch snapshots won&#39;t be created anymore if there will be more backups to clean than this value. | [optional] 
**CassandraScheduledTime** | Pointer to **int32** | Hour to start Cassandra backups each day. | [optional] 
**PauseBackups** | Pointer to **bool** | Pauses Elasticsearch and Cassandra backups. In comparison to enable/disable backup, this option does not modify any configuration like Elasticsearch properties. | [optional] 

## Methods

### NewBackupConfigDto

`func NewBackupConfigDto() *BackupConfigDto`

NewBackupConfigDto instantiates a new BackupConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupConfigDtoWithDefaults

`func NewBackupConfigDtoWithDefaults() *BackupConfigDto`

NewBackupConfigDtoWithDefaults instantiates a new BackupConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BackupConfigDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupConfigDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupConfigDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BackupConfigDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDatacenter

`func (o *BackupConfigDto) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *BackupConfigDto) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *BackupConfigDto) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *BackupConfigDto) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetStoragePath

`func (o *BackupConfigDto) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *BackupConfigDto) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *BackupConfigDto) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *BackupConfigDto) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetIncludeRumData

`func (o *BackupConfigDto) GetIncludeRumData() bool`

GetIncludeRumData returns the IncludeRumData field if non-nil, zero value otherwise.

### GetIncludeRumDataOk

`func (o *BackupConfigDto) GetIncludeRumDataOk() (*bool, bool)`

GetIncludeRumDataOk returns a tuple with the IncludeRumData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRumData

`func (o *BackupConfigDto) SetIncludeRumData(v bool)`

SetIncludeRumData sets IncludeRumData field to given value.

### HasIncludeRumData

`func (o *BackupConfigDto) HasIncludeRumData() bool`

HasIncludeRumData returns a boolean if a field has been set.

### GetIncludeTsMetricData

`func (o *BackupConfigDto) GetIncludeTsMetricData() bool`

GetIncludeTsMetricData returns the IncludeTsMetricData field if non-nil, zero value otherwise.

### GetIncludeTsMetricDataOk

`func (o *BackupConfigDto) GetIncludeTsMetricDataOk() (*bool, bool)`

GetIncludeTsMetricDataOk returns a tuple with the IncludeTsMetricData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTsMetricData

`func (o *BackupConfigDto) SetIncludeTsMetricData(v bool)`

SetIncludeTsMetricData sets IncludeTsMetricData field to given value.

### HasIncludeTsMetricData

`func (o *BackupConfigDto) HasIncludeTsMetricData() bool`

HasIncludeTsMetricData returns a boolean if a field has been set.

### GetBandwidthLimitMbits

`func (o *BackupConfigDto) GetBandwidthLimitMbits() int32`

GetBandwidthLimitMbits returns the BandwidthLimitMbits field if non-nil, zero value otherwise.

### GetBandwidthLimitMbitsOk

`func (o *BackupConfigDto) GetBandwidthLimitMbitsOk() (*int32, bool)`

GetBandwidthLimitMbitsOk returns a tuple with the BandwidthLimitMbits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimitMbits

`func (o *BackupConfigDto) SetBandwidthLimitMbits(v int32)`

SetBandwidthLimitMbits sets BandwidthLimitMbits field to given value.

### HasBandwidthLimitMbits

`func (o *BackupConfigDto) HasBandwidthLimitMbits() bool`

HasBandwidthLimitMbits returns a boolean if a field has been set.

### GetMaxEsSnapshotsToClean

`func (o *BackupConfigDto) GetMaxEsSnapshotsToClean() int32`

GetMaxEsSnapshotsToClean returns the MaxEsSnapshotsToClean field if non-nil, zero value otherwise.

### GetMaxEsSnapshotsToCleanOk

`func (o *BackupConfigDto) GetMaxEsSnapshotsToCleanOk() (*int32, bool)`

GetMaxEsSnapshotsToCleanOk returns a tuple with the MaxEsSnapshotsToClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEsSnapshotsToClean

`func (o *BackupConfigDto) SetMaxEsSnapshotsToClean(v int32)`

SetMaxEsSnapshotsToClean sets MaxEsSnapshotsToClean field to given value.

### HasMaxEsSnapshotsToClean

`func (o *BackupConfigDto) HasMaxEsSnapshotsToClean() bool`

HasMaxEsSnapshotsToClean returns a boolean if a field has been set.

### GetCassandraScheduledTime

`func (o *BackupConfigDto) GetCassandraScheduledTime() int32`

GetCassandraScheduledTime returns the CassandraScheduledTime field if non-nil, zero value otherwise.

### GetCassandraScheduledTimeOk

`func (o *BackupConfigDto) GetCassandraScheduledTimeOk() (*int32, bool)`

GetCassandraScheduledTimeOk returns a tuple with the CassandraScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraScheduledTime

`func (o *BackupConfigDto) SetCassandraScheduledTime(v int32)`

SetCassandraScheduledTime sets CassandraScheduledTime field to given value.

### HasCassandraScheduledTime

`func (o *BackupConfigDto) HasCassandraScheduledTime() bool`

HasCassandraScheduledTime returns a boolean if a field has been set.

### GetPauseBackups

`func (o *BackupConfigDto) GetPauseBackups() bool`

GetPauseBackups returns the PauseBackups field if non-nil, zero value otherwise.

### GetPauseBackupsOk

`func (o *BackupConfigDto) GetPauseBackupsOk() (*bool, bool)`

GetPauseBackupsOk returns a tuple with the PauseBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseBackups

`func (o *BackupConfigDto) SetPauseBackups(v bool)`

SetPauseBackups sets PauseBackups field to given value.

### HasPauseBackups

`func (o *BackupConfigDto) HasPauseBackups() bool`

HasPauseBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


