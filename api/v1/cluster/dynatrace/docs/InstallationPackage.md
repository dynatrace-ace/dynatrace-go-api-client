# InstallationPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedFromDownload** | Pointer to **bool** | If true, update package is excluded from download to save storage | [optional] 
**DeletedFromMCStorage** | Pointer to **bool** | If true, update package is removed from Dynatrace repository. Once you remove it from your cluster it will be gone permanently. | [optional] 
**ReadyNodeIds** | Pointer to **[]int32** | IDs of nodes that have finished package processing (i.e, downloaded or removed depending on status) | [optional] 
**FileSizeInBytes** | Pointer to **int64** | Total file size of a package in bytes | [optional] 
**DeleteEnabled** | Pointer to **bool** | If true, it&#39;s possible to remove this package via the REST API or the WebUi | [optional] 
**CountOfTenantsUsingThisAgentAsStandardVersion** | Pointer to **int32** | Count of tenants where this agent is configured as standard agent version. Applies to OneAgent type only. | [optional] 
**Version** | Pointer to **string** | Version | [optional] 
**Status** | Pointer to **string** | Cluster-wide status | [optional] 
**Type** | Pointer to **string** | Type | [optional] 

## Methods

### NewInstallationPackage

`func NewInstallationPackage() *InstallationPackage`

NewInstallationPackage instantiates a new InstallationPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationPackageWithDefaults

`func NewInstallationPackageWithDefaults() *InstallationPackage`

NewInstallationPackageWithDefaults instantiates a new InstallationPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedFromDownload

`func (o *InstallationPackage) GetExcludedFromDownload() bool`

GetExcludedFromDownload returns the ExcludedFromDownload field if non-nil, zero value otherwise.

### GetExcludedFromDownloadOk

`func (o *InstallationPackage) GetExcludedFromDownloadOk() (*bool, bool)`

GetExcludedFromDownloadOk returns a tuple with the ExcludedFromDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromDownload

`func (o *InstallationPackage) SetExcludedFromDownload(v bool)`

SetExcludedFromDownload sets ExcludedFromDownload field to given value.

### HasExcludedFromDownload

`func (o *InstallationPackage) HasExcludedFromDownload() bool`

HasExcludedFromDownload returns a boolean if a field has been set.

### GetDeletedFromMCStorage

`func (o *InstallationPackage) GetDeletedFromMCStorage() bool`

GetDeletedFromMCStorage returns the DeletedFromMCStorage field if non-nil, zero value otherwise.

### GetDeletedFromMCStorageOk

`func (o *InstallationPackage) GetDeletedFromMCStorageOk() (*bool, bool)`

GetDeletedFromMCStorageOk returns a tuple with the DeletedFromMCStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFromMCStorage

`func (o *InstallationPackage) SetDeletedFromMCStorage(v bool)`

SetDeletedFromMCStorage sets DeletedFromMCStorage field to given value.

### HasDeletedFromMCStorage

`func (o *InstallationPackage) HasDeletedFromMCStorage() bool`

HasDeletedFromMCStorage returns a boolean if a field has been set.

### GetReadyNodeIds

`func (o *InstallationPackage) GetReadyNodeIds() []int32`

GetReadyNodeIds returns the ReadyNodeIds field if non-nil, zero value otherwise.

### GetReadyNodeIdsOk

`func (o *InstallationPackage) GetReadyNodeIdsOk() (*[]int32, bool)`

GetReadyNodeIdsOk returns a tuple with the ReadyNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyNodeIds

`func (o *InstallationPackage) SetReadyNodeIds(v []int32)`

SetReadyNodeIds sets ReadyNodeIds field to given value.

### HasReadyNodeIds

`func (o *InstallationPackage) HasReadyNodeIds() bool`

HasReadyNodeIds returns a boolean if a field has been set.

### GetFileSizeInBytes

`func (o *InstallationPackage) GetFileSizeInBytes() int64`

GetFileSizeInBytes returns the FileSizeInBytes field if non-nil, zero value otherwise.

### GetFileSizeInBytesOk

`func (o *InstallationPackage) GetFileSizeInBytesOk() (*int64, bool)`

GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeInBytes

`func (o *InstallationPackage) SetFileSizeInBytes(v int64)`

SetFileSizeInBytes sets FileSizeInBytes field to given value.

### HasFileSizeInBytes

`func (o *InstallationPackage) HasFileSizeInBytes() bool`

HasFileSizeInBytes returns a boolean if a field has been set.

### GetDeleteEnabled

`func (o *InstallationPackage) GetDeleteEnabled() bool`

GetDeleteEnabled returns the DeleteEnabled field if non-nil, zero value otherwise.

### GetDeleteEnabledOk

`func (o *InstallationPackage) GetDeleteEnabledOk() (*bool, bool)`

GetDeleteEnabledOk returns a tuple with the DeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEnabled

`func (o *InstallationPackage) SetDeleteEnabled(v bool)`

SetDeleteEnabled sets DeleteEnabled field to given value.

### HasDeleteEnabled

`func (o *InstallationPackage) HasDeleteEnabled() bool`

HasDeleteEnabled returns a boolean if a field has been set.

### GetCountOfTenantsUsingThisAgentAsStandardVersion

`func (o *InstallationPackage) GetCountOfTenantsUsingThisAgentAsStandardVersion() int32`

GetCountOfTenantsUsingThisAgentAsStandardVersion returns the CountOfTenantsUsingThisAgentAsStandardVersion field if non-nil, zero value otherwise.

### GetCountOfTenantsUsingThisAgentAsStandardVersionOk

`func (o *InstallationPackage) GetCountOfTenantsUsingThisAgentAsStandardVersionOk() (*int32, bool)`

GetCountOfTenantsUsingThisAgentAsStandardVersionOk returns a tuple with the CountOfTenantsUsingThisAgentAsStandardVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOfTenantsUsingThisAgentAsStandardVersion

`func (o *InstallationPackage) SetCountOfTenantsUsingThisAgentAsStandardVersion(v int32)`

SetCountOfTenantsUsingThisAgentAsStandardVersion sets CountOfTenantsUsingThisAgentAsStandardVersion field to given value.

### HasCountOfTenantsUsingThisAgentAsStandardVersion

`func (o *InstallationPackage) HasCountOfTenantsUsingThisAgentAsStandardVersion() bool`

HasCountOfTenantsUsingThisAgentAsStandardVersion returns a boolean if a field has been set.

### GetVersion

`func (o *InstallationPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstallationPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstallationPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstallationPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *InstallationPackage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationPackage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationPackage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstallationPackage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *InstallationPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallationPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallationPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstallationPackage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


