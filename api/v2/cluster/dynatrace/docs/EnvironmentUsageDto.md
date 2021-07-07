# EnvironmentUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentUuid** | Pointer to **string** |  | [optional] 
**Visits** | Pointer to **int64** |  | [optional] 
**MobileSessions** | Pointer to **int64** |  | [optional] 
**TotalRUMUserPropertiesUsed** | Pointer to **int64** |  | [optional] 
**NewProblems** | Pointer to **int64** |  | [optional] 
**HostUsages** | Pointer to [**[]HostUsageDto**](HostUsageDto.md) |  | [optional] 
**Downloads** | Pointer to [**[]DownloadsDto**](DownloadsDto.md) |  | [optional] 
**SyntheticUsages** | Pointer to [**[]SyntheticUsageDto**](SyntheticUsageDto.md) |  | [optional] 
**SyntheticBillingUsage** | Pointer to [**[]SyntheticBillingUsageDto**](SyntheticBillingUsageDto.md) |  | [optional] 
**CustomMetrics** | Pointer to [**[]CustomMetricDto**](CustomMetricDto.md) |  | [optional] 
**DavisDataUnits** | Pointer to [**[]DavisDataUnitsUsageDto**](DavisDataUnitsUsageDto.md) |  | [optional] 
**Trial** | Pointer to **bool** |  | [optional] 
**InternalUse** | Pointer to **bool** |  | [optional] 
**HighAvailabilityCluster** | Pointer to **bool** |  | [optional] 
**LogStorageUsageBytes** | Pointer to **int64** |  | [optional] 
**LogUploadVolumeBytes** | Pointer to **int64** |  | [optional] 
**SessionReplays** | Pointer to **int64** |  | [optional] 
**MobileSessionReplays** | Pointer to **int64** |  | [optional] 

## Methods

### NewEnvironmentUsageDto

`func NewEnvironmentUsageDto() *EnvironmentUsageDto`

NewEnvironmentUsageDto instantiates a new EnvironmentUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUsageDtoWithDefaults

`func NewEnvironmentUsageDtoWithDefaults() *EnvironmentUsageDto`

NewEnvironmentUsageDtoWithDefaults instantiates a new EnvironmentUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentUuid

`func (o *EnvironmentUsageDto) GetEnvironmentUuid() string`

GetEnvironmentUuid returns the EnvironmentUuid field if non-nil, zero value otherwise.

### GetEnvironmentUuidOk

`func (o *EnvironmentUsageDto) GetEnvironmentUuidOk() (*string, bool)`

GetEnvironmentUuidOk returns a tuple with the EnvironmentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUuid

`func (o *EnvironmentUsageDto) SetEnvironmentUuid(v string)`

SetEnvironmentUuid sets EnvironmentUuid field to given value.

### HasEnvironmentUuid

`func (o *EnvironmentUsageDto) HasEnvironmentUuid() bool`

HasEnvironmentUuid returns a boolean if a field has been set.

### GetVisits

`func (o *EnvironmentUsageDto) GetVisits() int64`

GetVisits returns the Visits field if non-nil, zero value otherwise.

### GetVisitsOk

`func (o *EnvironmentUsageDto) GetVisitsOk() (*int64, bool)`

GetVisitsOk returns a tuple with the Visits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisits

`func (o *EnvironmentUsageDto) SetVisits(v int64)`

SetVisits sets Visits field to given value.

### HasVisits

`func (o *EnvironmentUsageDto) HasVisits() bool`

HasVisits returns a boolean if a field has been set.

### GetMobileSessions

`func (o *EnvironmentUsageDto) GetMobileSessions() int64`

GetMobileSessions returns the MobileSessions field if non-nil, zero value otherwise.

### GetMobileSessionsOk

`func (o *EnvironmentUsageDto) GetMobileSessionsOk() (*int64, bool)`

GetMobileSessionsOk returns a tuple with the MobileSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileSessions

`func (o *EnvironmentUsageDto) SetMobileSessions(v int64)`

SetMobileSessions sets MobileSessions field to given value.

### HasMobileSessions

`func (o *EnvironmentUsageDto) HasMobileSessions() bool`

HasMobileSessions returns a boolean if a field has been set.

### GetTotalRUMUserPropertiesUsed

`func (o *EnvironmentUsageDto) GetTotalRUMUserPropertiesUsed() int64`

GetTotalRUMUserPropertiesUsed returns the TotalRUMUserPropertiesUsed field if non-nil, zero value otherwise.

### GetTotalRUMUserPropertiesUsedOk

`func (o *EnvironmentUsageDto) GetTotalRUMUserPropertiesUsedOk() (*int64, bool)`

GetTotalRUMUserPropertiesUsedOk returns a tuple with the TotalRUMUserPropertiesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRUMUserPropertiesUsed

`func (o *EnvironmentUsageDto) SetTotalRUMUserPropertiesUsed(v int64)`

SetTotalRUMUserPropertiesUsed sets TotalRUMUserPropertiesUsed field to given value.

### HasTotalRUMUserPropertiesUsed

`func (o *EnvironmentUsageDto) HasTotalRUMUserPropertiesUsed() bool`

HasTotalRUMUserPropertiesUsed returns a boolean if a field has been set.

### GetNewProblems

`func (o *EnvironmentUsageDto) GetNewProblems() int64`

GetNewProblems returns the NewProblems field if non-nil, zero value otherwise.

### GetNewProblemsOk

`func (o *EnvironmentUsageDto) GetNewProblemsOk() (*int64, bool)`

GetNewProblemsOk returns a tuple with the NewProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewProblems

`func (o *EnvironmentUsageDto) SetNewProblems(v int64)`

SetNewProblems sets NewProblems field to given value.

### HasNewProblems

`func (o *EnvironmentUsageDto) HasNewProblems() bool`

HasNewProblems returns a boolean if a field has been set.

### GetHostUsages

`func (o *EnvironmentUsageDto) GetHostUsages() []HostUsageDto`

GetHostUsages returns the HostUsages field if non-nil, zero value otherwise.

### GetHostUsagesOk

`func (o *EnvironmentUsageDto) GetHostUsagesOk() (*[]HostUsageDto, bool)`

GetHostUsagesOk returns a tuple with the HostUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUsages

`func (o *EnvironmentUsageDto) SetHostUsages(v []HostUsageDto)`

SetHostUsages sets HostUsages field to given value.

### HasHostUsages

`func (o *EnvironmentUsageDto) HasHostUsages() bool`

HasHostUsages returns a boolean if a field has been set.

### GetDownloads

`func (o *EnvironmentUsageDto) GetDownloads() []DownloadsDto`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *EnvironmentUsageDto) GetDownloadsOk() (*[]DownloadsDto, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *EnvironmentUsageDto) SetDownloads(v []DownloadsDto)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *EnvironmentUsageDto) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetSyntheticUsages

`func (o *EnvironmentUsageDto) GetSyntheticUsages() []SyntheticUsageDto`

GetSyntheticUsages returns the SyntheticUsages field if non-nil, zero value otherwise.

### GetSyntheticUsagesOk

`func (o *EnvironmentUsageDto) GetSyntheticUsagesOk() (*[]SyntheticUsageDto, bool)`

GetSyntheticUsagesOk returns a tuple with the SyntheticUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticUsages

`func (o *EnvironmentUsageDto) SetSyntheticUsages(v []SyntheticUsageDto)`

SetSyntheticUsages sets SyntheticUsages field to given value.

### HasSyntheticUsages

`func (o *EnvironmentUsageDto) HasSyntheticUsages() bool`

HasSyntheticUsages returns a boolean if a field has been set.

### GetSyntheticBillingUsage

`func (o *EnvironmentUsageDto) GetSyntheticBillingUsage() []SyntheticBillingUsageDto`

GetSyntheticBillingUsage returns the SyntheticBillingUsage field if non-nil, zero value otherwise.

### GetSyntheticBillingUsageOk

`func (o *EnvironmentUsageDto) GetSyntheticBillingUsageOk() (*[]SyntheticBillingUsageDto, bool)`

GetSyntheticBillingUsageOk returns a tuple with the SyntheticBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticBillingUsage

`func (o *EnvironmentUsageDto) SetSyntheticBillingUsage(v []SyntheticBillingUsageDto)`

SetSyntheticBillingUsage sets SyntheticBillingUsage field to given value.

### HasSyntheticBillingUsage

`func (o *EnvironmentUsageDto) HasSyntheticBillingUsage() bool`

HasSyntheticBillingUsage returns a boolean if a field has been set.

### GetCustomMetrics

`func (o *EnvironmentUsageDto) GetCustomMetrics() []CustomMetricDto`

GetCustomMetrics returns the CustomMetrics field if non-nil, zero value otherwise.

### GetCustomMetricsOk

`func (o *EnvironmentUsageDto) GetCustomMetricsOk() (*[]CustomMetricDto, bool)`

GetCustomMetricsOk returns a tuple with the CustomMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetrics

`func (o *EnvironmentUsageDto) SetCustomMetrics(v []CustomMetricDto)`

SetCustomMetrics sets CustomMetrics field to given value.

### HasCustomMetrics

`func (o *EnvironmentUsageDto) HasCustomMetrics() bool`

HasCustomMetrics returns a boolean if a field has been set.

### GetDavisDataUnits

`func (o *EnvironmentUsageDto) GetDavisDataUnits() []DavisDataUnitsUsageDto`

GetDavisDataUnits returns the DavisDataUnits field if non-nil, zero value otherwise.

### GetDavisDataUnitsOk

`func (o *EnvironmentUsageDto) GetDavisDataUnitsOk() (*[]DavisDataUnitsUsageDto, bool)`

GetDavisDataUnitsOk returns a tuple with the DavisDataUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDavisDataUnits

`func (o *EnvironmentUsageDto) SetDavisDataUnits(v []DavisDataUnitsUsageDto)`

SetDavisDataUnits sets DavisDataUnits field to given value.

### HasDavisDataUnits

`func (o *EnvironmentUsageDto) HasDavisDataUnits() bool`

HasDavisDataUnits returns a boolean if a field has been set.

### GetTrial

`func (o *EnvironmentUsageDto) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *EnvironmentUsageDto) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *EnvironmentUsageDto) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *EnvironmentUsageDto) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetInternalUse

`func (o *EnvironmentUsageDto) GetInternalUse() bool`

GetInternalUse returns the InternalUse field if non-nil, zero value otherwise.

### GetInternalUseOk

`func (o *EnvironmentUsageDto) GetInternalUseOk() (*bool, bool)`

GetInternalUseOk returns a tuple with the InternalUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUse

`func (o *EnvironmentUsageDto) SetInternalUse(v bool)`

SetInternalUse sets InternalUse field to given value.

### HasInternalUse

`func (o *EnvironmentUsageDto) HasInternalUse() bool`

HasInternalUse returns a boolean if a field has been set.

### GetHighAvailabilityCluster

`func (o *EnvironmentUsageDto) GetHighAvailabilityCluster() bool`

GetHighAvailabilityCluster returns the HighAvailabilityCluster field if non-nil, zero value otherwise.

### GetHighAvailabilityClusterOk

`func (o *EnvironmentUsageDto) GetHighAvailabilityClusterOk() (*bool, bool)`

GetHighAvailabilityClusterOk returns a tuple with the HighAvailabilityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailabilityCluster

`func (o *EnvironmentUsageDto) SetHighAvailabilityCluster(v bool)`

SetHighAvailabilityCluster sets HighAvailabilityCluster field to given value.

### HasHighAvailabilityCluster

`func (o *EnvironmentUsageDto) HasHighAvailabilityCluster() bool`

HasHighAvailabilityCluster returns a boolean if a field has been set.

### GetLogStorageUsageBytes

`func (o *EnvironmentUsageDto) GetLogStorageUsageBytes() int64`

GetLogStorageUsageBytes returns the LogStorageUsageBytes field if non-nil, zero value otherwise.

### GetLogStorageUsageBytesOk

`func (o *EnvironmentUsageDto) GetLogStorageUsageBytesOk() (*int64, bool)`

GetLogStorageUsageBytesOk returns a tuple with the LogStorageUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStorageUsageBytes

`func (o *EnvironmentUsageDto) SetLogStorageUsageBytes(v int64)`

SetLogStorageUsageBytes sets LogStorageUsageBytes field to given value.

### HasLogStorageUsageBytes

`func (o *EnvironmentUsageDto) HasLogStorageUsageBytes() bool`

HasLogStorageUsageBytes returns a boolean if a field has been set.

### GetLogUploadVolumeBytes

`func (o *EnvironmentUsageDto) GetLogUploadVolumeBytes() int64`

GetLogUploadVolumeBytes returns the LogUploadVolumeBytes field if non-nil, zero value otherwise.

### GetLogUploadVolumeBytesOk

`func (o *EnvironmentUsageDto) GetLogUploadVolumeBytesOk() (*int64, bool)`

GetLogUploadVolumeBytesOk returns a tuple with the LogUploadVolumeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUploadVolumeBytes

`func (o *EnvironmentUsageDto) SetLogUploadVolumeBytes(v int64)`

SetLogUploadVolumeBytes sets LogUploadVolumeBytes field to given value.

### HasLogUploadVolumeBytes

`func (o *EnvironmentUsageDto) HasLogUploadVolumeBytes() bool`

HasLogUploadVolumeBytes returns a boolean if a field has been set.

### GetSessionReplays

`func (o *EnvironmentUsageDto) GetSessionReplays() int64`

GetSessionReplays returns the SessionReplays field if non-nil, zero value otherwise.

### GetSessionReplaysOk

`func (o *EnvironmentUsageDto) GetSessionReplaysOk() (*int64, bool)`

GetSessionReplaysOk returns a tuple with the SessionReplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplays

`func (o *EnvironmentUsageDto) SetSessionReplays(v int64)`

SetSessionReplays sets SessionReplays field to given value.

### HasSessionReplays

`func (o *EnvironmentUsageDto) HasSessionReplays() bool`

HasSessionReplays returns a boolean if a field has been set.

### GetMobileSessionReplays

`func (o *EnvironmentUsageDto) GetMobileSessionReplays() int64`

GetMobileSessionReplays returns the MobileSessionReplays field if non-nil, zero value otherwise.

### GetMobileSessionReplaysOk

`func (o *EnvironmentUsageDto) GetMobileSessionReplaysOk() (*int64, bool)`

GetMobileSessionReplaysOk returns a tuple with the MobileSessionReplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileSessionReplays

`func (o *EnvironmentUsageDto) SetMobileSessionReplays(v int64)`

SetMobileSessionReplays sets MobileSessionReplays field to given value.

### HasMobileSessionReplays

`func (o *EnvironmentUsageDto) HasMobileSessionReplays() bool`

HasMobileSessionReplays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


