# VMwareAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**EsxiHighCpuSaturation** | [**EsxiHighCpuSaturationConfig**](EsxiHighCpuSaturationConfig.md) |  | 
**GuestCpuLimitReached** | Pointer to [**GuestCPULimitReachedConfig**](GuestCPULimitReachedConfig.md) |  | [optional] 
**EsxiHighMemoryDetection** | [**EsxiHighMemoryDetectionConfig**](EsxiHighMemoryDetectionConfig.md) |  | 
**OverloadedStorageDetection** | [**OverloadedStorageDetectionConfig**](OverloadedStorageDetectionConfig.md) |  | 
**UndersizedStorageDetection** | [**UndersizedStorageDetectionConfig**](UndersizedStorageDetectionConfig.md) |  | 
**SlowPhysicalStorageDetection** | [**SlowPhysicalStorageDetectionConfig**](SlowPhysicalStorageDetectionConfig.md) |  | 
**DroppedPacketsDetection** | [**DroppedPacketsDetectionConfig**](DroppedPacketsDetectionConfig.md) |  | 
**LowDatastoreSpaceDetection** | [**LowDatastoreSpaceDetectionConfig**](LowDatastoreSpaceDetectionConfig.md) |  | 

## Methods

### NewVMwareAnomalyDetectionConfig

`func NewVMwareAnomalyDetectionConfig(esxiHighCpuSaturation EsxiHighCpuSaturationConfig, esxiHighMemoryDetection EsxiHighMemoryDetectionConfig, overloadedStorageDetection OverloadedStorageDetectionConfig, undersizedStorageDetection UndersizedStorageDetectionConfig, slowPhysicalStorageDetection SlowPhysicalStorageDetectionConfig, droppedPacketsDetection DroppedPacketsDetectionConfig, lowDatastoreSpaceDetection LowDatastoreSpaceDetectionConfig, ) *VMwareAnomalyDetectionConfig`

NewVMwareAnomalyDetectionConfig instantiates a new VMwareAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareAnomalyDetectionConfigWithDefaults

`func NewVMwareAnomalyDetectionConfigWithDefaults() *VMwareAnomalyDetectionConfig`

NewVMwareAnomalyDetectionConfigWithDefaults instantiates a new VMwareAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *VMwareAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VMwareAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VMwareAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VMwareAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEsxiHighCpuSaturation

`func (o *VMwareAnomalyDetectionConfig) GetEsxiHighCpuSaturation() EsxiHighCpuSaturationConfig`

GetEsxiHighCpuSaturation returns the EsxiHighCpuSaturation field if non-nil, zero value otherwise.

### GetEsxiHighCpuSaturationOk

`func (o *VMwareAnomalyDetectionConfig) GetEsxiHighCpuSaturationOk() (*EsxiHighCpuSaturationConfig, bool)`

GetEsxiHighCpuSaturationOk returns a tuple with the EsxiHighCpuSaturation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiHighCpuSaturation

`func (o *VMwareAnomalyDetectionConfig) SetEsxiHighCpuSaturation(v EsxiHighCpuSaturationConfig)`

SetEsxiHighCpuSaturation sets EsxiHighCpuSaturation field to given value.


### GetGuestCpuLimitReached

`func (o *VMwareAnomalyDetectionConfig) GetGuestCpuLimitReached() GuestCPULimitReachedConfig`

GetGuestCpuLimitReached returns the GuestCpuLimitReached field if non-nil, zero value otherwise.

### GetGuestCpuLimitReachedOk

`func (o *VMwareAnomalyDetectionConfig) GetGuestCpuLimitReachedOk() (*GuestCPULimitReachedConfig, bool)`

GetGuestCpuLimitReachedOk returns a tuple with the GuestCpuLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCpuLimitReached

`func (o *VMwareAnomalyDetectionConfig) SetGuestCpuLimitReached(v GuestCPULimitReachedConfig)`

SetGuestCpuLimitReached sets GuestCpuLimitReached field to given value.

### HasGuestCpuLimitReached

`func (o *VMwareAnomalyDetectionConfig) HasGuestCpuLimitReached() bool`

HasGuestCpuLimitReached returns a boolean if a field has been set.

### GetEsxiHighMemoryDetection

`func (o *VMwareAnomalyDetectionConfig) GetEsxiHighMemoryDetection() EsxiHighMemoryDetectionConfig`

GetEsxiHighMemoryDetection returns the EsxiHighMemoryDetection field if non-nil, zero value otherwise.

### GetEsxiHighMemoryDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetEsxiHighMemoryDetectionOk() (*EsxiHighMemoryDetectionConfig, bool)`

GetEsxiHighMemoryDetectionOk returns a tuple with the EsxiHighMemoryDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiHighMemoryDetection

`func (o *VMwareAnomalyDetectionConfig) SetEsxiHighMemoryDetection(v EsxiHighMemoryDetectionConfig)`

SetEsxiHighMemoryDetection sets EsxiHighMemoryDetection field to given value.


### GetOverloadedStorageDetection

`func (o *VMwareAnomalyDetectionConfig) GetOverloadedStorageDetection() OverloadedStorageDetectionConfig`

GetOverloadedStorageDetection returns the OverloadedStorageDetection field if non-nil, zero value otherwise.

### GetOverloadedStorageDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetOverloadedStorageDetectionOk() (*OverloadedStorageDetectionConfig, bool)`

GetOverloadedStorageDetectionOk returns a tuple with the OverloadedStorageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverloadedStorageDetection

`func (o *VMwareAnomalyDetectionConfig) SetOverloadedStorageDetection(v OverloadedStorageDetectionConfig)`

SetOverloadedStorageDetection sets OverloadedStorageDetection field to given value.


### GetUndersizedStorageDetection

`func (o *VMwareAnomalyDetectionConfig) GetUndersizedStorageDetection() UndersizedStorageDetectionConfig`

GetUndersizedStorageDetection returns the UndersizedStorageDetection field if non-nil, zero value otherwise.

### GetUndersizedStorageDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetUndersizedStorageDetectionOk() (*UndersizedStorageDetectionConfig, bool)`

GetUndersizedStorageDetectionOk returns a tuple with the UndersizedStorageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndersizedStorageDetection

`func (o *VMwareAnomalyDetectionConfig) SetUndersizedStorageDetection(v UndersizedStorageDetectionConfig)`

SetUndersizedStorageDetection sets UndersizedStorageDetection field to given value.


### GetSlowPhysicalStorageDetection

`func (o *VMwareAnomalyDetectionConfig) GetSlowPhysicalStorageDetection() SlowPhysicalStorageDetectionConfig`

GetSlowPhysicalStorageDetection returns the SlowPhysicalStorageDetection field if non-nil, zero value otherwise.

### GetSlowPhysicalStorageDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetSlowPhysicalStorageDetectionOk() (*SlowPhysicalStorageDetectionConfig, bool)`

GetSlowPhysicalStorageDetectionOk returns a tuple with the SlowPhysicalStorageDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowPhysicalStorageDetection

`func (o *VMwareAnomalyDetectionConfig) SetSlowPhysicalStorageDetection(v SlowPhysicalStorageDetectionConfig)`

SetSlowPhysicalStorageDetection sets SlowPhysicalStorageDetection field to given value.


### GetDroppedPacketsDetection

`func (o *VMwareAnomalyDetectionConfig) GetDroppedPacketsDetection() DroppedPacketsDetectionConfig`

GetDroppedPacketsDetection returns the DroppedPacketsDetection field if non-nil, zero value otherwise.

### GetDroppedPacketsDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetDroppedPacketsDetectionOk() (*DroppedPacketsDetectionConfig, bool)`

GetDroppedPacketsDetectionOk returns a tuple with the DroppedPacketsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedPacketsDetection

`func (o *VMwareAnomalyDetectionConfig) SetDroppedPacketsDetection(v DroppedPacketsDetectionConfig)`

SetDroppedPacketsDetection sets DroppedPacketsDetection field to given value.


### GetLowDatastoreSpaceDetection

`func (o *VMwareAnomalyDetectionConfig) GetLowDatastoreSpaceDetection() LowDatastoreSpaceDetectionConfig`

GetLowDatastoreSpaceDetection returns the LowDatastoreSpaceDetection field if non-nil, zero value otherwise.

### GetLowDatastoreSpaceDetectionOk

`func (o *VMwareAnomalyDetectionConfig) GetLowDatastoreSpaceDetectionOk() (*LowDatastoreSpaceDetectionConfig, bool)`

GetLowDatastoreSpaceDetectionOk returns a tuple with the LowDatastoreSpaceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowDatastoreSpaceDetection

`func (o *VMwareAnomalyDetectionConfig) SetLowDatastoreSpaceDetection(v LowDatastoreSpaceDetectionConfig)`

SetLowDatastoreSpaceDetection sets LowDatastoreSpaceDetection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


