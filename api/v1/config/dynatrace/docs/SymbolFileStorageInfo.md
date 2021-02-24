# SymbolFileStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedStorage** | Pointer to **int64** | The size of the used storage by symbolication files in KB | [optional] 
**AvailableStorage** | Pointer to **int64** | The storage space still empty which can be used for symbolication files in KB | [optional] 
**StorageQuota** | Pointer to **int64** | The total storage quota available on this tenant for symbolication files in KB | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewSymbolFileStorageInfo

`func NewSymbolFileStorageInfo() *SymbolFileStorageInfo`

NewSymbolFileStorageInfo instantiates a new SymbolFileStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFileStorageInfoWithDefaults

`func NewSymbolFileStorageInfoWithDefaults() *SymbolFileStorageInfo`

NewSymbolFileStorageInfoWithDefaults instantiates a new SymbolFileStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedStorage

`func (o *SymbolFileStorageInfo) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *SymbolFileStorageInfo) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *SymbolFileStorageInfo) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *SymbolFileStorageInfo) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetAvailableStorage

`func (o *SymbolFileStorageInfo) GetAvailableStorage() int64`

GetAvailableStorage returns the AvailableStorage field if non-nil, zero value otherwise.

### GetAvailableStorageOk

`func (o *SymbolFileStorageInfo) GetAvailableStorageOk() (*int64, bool)`

GetAvailableStorageOk returns a tuple with the AvailableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStorage

`func (o *SymbolFileStorageInfo) SetAvailableStorage(v int64)`

SetAvailableStorage sets AvailableStorage field to given value.

### HasAvailableStorage

`func (o *SymbolFileStorageInfo) HasAvailableStorage() bool`

HasAvailableStorage returns a boolean if a field has been set.

### GetStorageQuota

`func (o *SymbolFileStorageInfo) GetStorageQuota() int64`

GetStorageQuota returns the StorageQuota field if non-nil, zero value otherwise.

### GetStorageQuotaOk

`func (o *SymbolFileStorageInfo) GetStorageQuotaOk() (*int64, bool)`

GetStorageQuotaOk returns a tuple with the StorageQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageQuota

`func (o *SymbolFileStorageInfo) SetStorageQuota(v int64)`

SetStorageQuota sets StorageQuota field to given value.

### HasStorageQuota

`func (o *SymbolFileStorageInfo) HasStorageQuota() bool`

HasStorageQuota returns a boolean if a field has been set.

### GetFileCount

`func (o *SymbolFileStorageInfo) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *SymbolFileStorageInfo) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *SymbolFileStorageInfo) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *SymbolFileStorageInfo) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


