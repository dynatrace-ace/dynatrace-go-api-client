# HostsAnomalyDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**ConnectionLostDetection** | [**ConnectionLostDetectionConfig**](ConnectionLostDetectionConfig.md) |  | 
**HighCpuSaturationDetection** | [**HighCpuSaturationDetectionConfig**](HighCpuSaturationDetectionConfig.md) |  | 
**HighMemoryDetection** | [**HighMemoryDetectionConfig**](HighMemoryDetectionConfig.md) |  | 
**HighGcActivityDetection** | [**HighGcActivityDetectionConfig**](HighGcActivityDetectionConfig.md) |  | 
**OutOfMemoryDetection** | [**OutOfMemoryDetectionConfig**](OutOfMemoryDetectionConfig.md) |  | 
**OutOfThreadsDetection** | [**OutOfThreadsDetectionConfig**](OutOfThreadsDetectionConfig.md) |  | 
**NetworkDroppedPacketsDetection** | [**NetworkDroppedPacketsDetectionConfig**](NetworkDroppedPacketsDetectionConfig.md) |  | 
**NetworkErrorsDetection** | [**NetworkErrorsDetectionConfig**](NetworkErrorsDetectionConfig.md) |  | 
**HighNetworkDetection** | [**HighNetworkDetectionConfig**](HighNetworkDetectionConfig.md) |  | 
**NetworkTcpProblemsDetection** | [**NetworkTcpProblemsDetectionConfig**](NetworkTcpProblemsDetectionConfig.md) |  | 
**NetworkHighRetransmissionDetection** | [**NetworkHighRetransmissionDetectionConfig**](NetworkHighRetransmissionDetectionConfig.md) |  | 
**DiskLowSpaceDetection** | [**DiskLowSpaceDetectionConfig**](DiskLowSpaceDetectionConfig.md) |  | 
**DiskSlowWritesAndReadsDetection** | [**DiskSlowWritesAndReadsDetectionConfig**](DiskSlowWritesAndReadsDetectionConfig.md) |  | 
**DiskLowInodesDetection** | [**DiskLowInodesDetectionConfig**](DiskLowInodesDetectionConfig.md) |  | 

## Methods

### NewHostsAnomalyDetectionConfig

`func NewHostsAnomalyDetectionConfig(connectionLostDetection ConnectionLostDetectionConfig, highCpuSaturationDetection HighCpuSaturationDetectionConfig, highMemoryDetection HighMemoryDetectionConfig, highGcActivityDetection HighGcActivityDetectionConfig, outOfMemoryDetection OutOfMemoryDetectionConfig, outOfThreadsDetection OutOfThreadsDetectionConfig, networkDroppedPacketsDetection NetworkDroppedPacketsDetectionConfig, networkErrorsDetection NetworkErrorsDetectionConfig, highNetworkDetection HighNetworkDetectionConfig, networkTcpProblemsDetection NetworkTcpProblemsDetectionConfig, networkHighRetransmissionDetection NetworkHighRetransmissionDetectionConfig, diskLowSpaceDetection DiskLowSpaceDetectionConfig, diskSlowWritesAndReadsDetection DiskSlowWritesAndReadsDetectionConfig, diskLowInodesDetection DiskLowInodesDetectionConfig, ) *HostsAnomalyDetectionConfig`

NewHostsAnomalyDetectionConfig instantiates a new HostsAnomalyDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsAnomalyDetectionConfigWithDefaults

`func NewHostsAnomalyDetectionConfigWithDefaults() *HostsAnomalyDetectionConfig`

NewHostsAnomalyDetectionConfigWithDefaults instantiates a new HostsAnomalyDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HostsAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostsAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostsAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HostsAnomalyDetectionConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConnectionLostDetection

`func (o *HostsAnomalyDetectionConfig) GetConnectionLostDetection() ConnectionLostDetectionConfig`

GetConnectionLostDetection returns the ConnectionLostDetection field if non-nil, zero value otherwise.

### GetConnectionLostDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetConnectionLostDetectionOk() (*ConnectionLostDetectionConfig, bool)`

GetConnectionLostDetectionOk returns a tuple with the ConnectionLostDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLostDetection

`func (o *HostsAnomalyDetectionConfig) SetConnectionLostDetection(v ConnectionLostDetectionConfig)`

SetConnectionLostDetection sets ConnectionLostDetection field to given value.


### GetHighCpuSaturationDetection

`func (o *HostsAnomalyDetectionConfig) GetHighCpuSaturationDetection() HighCpuSaturationDetectionConfig`

GetHighCpuSaturationDetection returns the HighCpuSaturationDetection field if non-nil, zero value otherwise.

### GetHighCpuSaturationDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetHighCpuSaturationDetectionOk() (*HighCpuSaturationDetectionConfig, bool)`

GetHighCpuSaturationDetectionOk returns a tuple with the HighCpuSaturationDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighCpuSaturationDetection

`func (o *HostsAnomalyDetectionConfig) SetHighCpuSaturationDetection(v HighCpuSaturationDetectionConfig)`

SetHighCpuSaturationDetection sets HighCpuSaturationDetection field to given value.


### GetHighMemoryDetection

`func (o *HostsAnomalyDetectionConfig) GetHighMemoryDetection() HighMemoryDetectionConfig`

GetHighMemoryDetection returns the HighMemoryDetection field if non-nil, zero value otherwise.

### GetHighMemoryDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetHighMemoryDetectionOk() (*HighMemoryDetectionConfig, bool)`

GetHighMemoryDetectionOk returns a tuple with the HighMemoryDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighMemoryDetection

`func (o *HostsAnomalyDetectionConfig) SetHighMemoryDetection(v HighMemoryDetectionConfig)`

SetHighMemoryDetection sets HighMemoryDetection field to given value.


### GetHighGcActivityDetection

`func (o *HostsAnomalyDetectionConfig) GetHighGcActivityDetection() HighGcActivityDetectionConfig`

GetHighGcActivityDetection returns the HighGcActivityDetection field if non-nil, zero value otherwise.

### GetHighGcActivityDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetHighGcActivityDetectionOk() (*HighGcActivityDetectionConfig, bool)`

GetHighGcActivityDetectionOk returns a tuple with the HighGcActivityDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighGcActivityDetection

`func (o *HostsAnomalyDetectionConfig) SetHighGcActivityDetection(v HighGcActivityDetectionConfig)`

SetHighGcActivityDetection sets HighGcActivityDetection field to given value.


### GetOutOfMemoryDetection

`func (o *HostsAnomalyDetectionConfig) GetOutOfMemoryDetection() OutOfMemoryDetectionConfig`

GetOutOfMemoryDetection returns the OutOfMemoryDetection field if non-nil, zero value otherwise.

### GetOutOfMemoryDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetOutOfMemoryDetectionOk() (*OutOfMemoryDetectionConfig, bool)`

GetOutOfMemoryDetectionOk returns a tuple with the OutOfMemoryDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfMemoryDetection

`func (o *HostsAnomalyDetectionConfig) SetOutOfMemoryDetection(v OutOfMemoryDetectionConfig)`

SetOutOfMemoryDetection sets OutOfMemoryDetection field to given value.


### GetOutOfThreadsDetection

`func (o *HostsAnomalyDetectionConfig) GetOutOfThreadsDetection() OutOfThreadsDetectionConfig`

GetOutOfThreadsDetection returns the OutOfThreadsDetection field if non-nil, zero value otherwise.

### GetOutOfThreadsDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetOutOfThreadsDetectionOk() (*OutOfThreadsDetectionConfig, bool)`

GetOutOfThreadsDetectionOk returns a tuple with the OutOfThreadsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfThreadsDetection

`func (o *HostsAnomalyDetectionConfig) SetOutOfThreadsDetection(v OutOfThreadsDetectionConfig)`

SetOutOfThreadsDetection sets OutOfThreadsDetection field to given value.


### GetNetworkDroppedPacketsDetection

`func (o *HostsAnomalyDetectionConfig) GetNetworkDroppedPacketsDetection() NetworkDroppedPacketsDetectionConfig`

GetNetworkDroppedPacketsDetection returns the NetworkDroppedPacketsDetection field if non-nil, zero value otherwise.

### GetNetworkDroppedPacketsDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetNetworkDroppedPacketsDetectionOk() (*NetworkDroppedPacketsDetectionConfig, bool)`

GetNetworkDroppedPacketsDetectionOk returns a tuple with the NetworkDroppedPacketsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDroppedPacketsDetection

`func (o *HostsAnomalyDetectionConfig) SetNetworkDroppedPacketsDetection(v NetworkDroppedPacketsDetectionConfig)`

SetNetworkDroppedPacketsDetection sets NetworkDroppedPacketsDetection field to given value.


### GetNetworkErrorsDetection

`func (o *HostsAnomalyDetectionConfig) GetNetworkErrorsDetection() NetworkErrorsDetectionConfig`

GetNetworkErrorsDetection returns the NetworkErrorsDetection field if non-nil, zero value otherwise.

### GetNetworkErrorsDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetNetworkErrorsDetectionOk() (*NetworkErrorsDetectionConfig, bool)`

GetNetworkErrorsDetectionOk returns a tuple with the NetworkErrorsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkErrorsDetection

`func (o *HostsAnomalyDetectionConfig) SetNetworkErrorsDetection(v NetworkErrorsDetectionConfig)`

SetNetworkErrorsDetection sets NetworkErrorsDetection field to given value.


### GetHighNetworkDetection

`func (o *HostsAnomalyDetectionConfig) GetHighNetworkDetection() HighNetworkDetectionConfig`

GetHighNetworkDetection returns the HighNetworkDetection field if non-nil, zero value otherwise.

### GetHighNetworkDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetHighNetworkDetectionOk() (*HighNetworkDetectionConfig, bool)`

GetHighNetworkDetectionOk returns a tuple with the HighNetworkDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighNetworkDetection

`func (o *HostsAnomalyDetectionConfig) SetHighNetworkDetection(v HighNetworkDetectionConfig)`

SetHighNetworkDetection sets HighNetworkDetection field to given value.


### GetNetworkTcpProblemsDetection

`func (o *HostsAnomalyDetectionConfig) GetNetworkTcpProblemsDetection() NetworkTcpProblemsDetectionConfig`

GetNetworkTcpProblemsDetection returns the NetworkTcpProblemsDetection field if non-nil, zero value otherwise.

### GetNetworkTcpProblemsDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetNetworkTcpProblemsDetectionOk() (*NetworkTcpProblemsDetectionConfig, bool)`

GetNetworkTcpProblemsDetectionOk returns a tuple with the NetworkTcpProblemsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTcpProblemsDetection

`func (o *HostsAnomalyDetectionConfig) SetNetworkTcpProblemsDetection(v NetworkTcpProblemsDetectionConfig)`

SetNetworkTcpProblemsDetection sets NetworkTcpProblemsDetection field to given value.


### GetNetworkHighRetransmissionDetection

`func (o *HostsAnomalyDetectionConfig) GetNetworkHighRetransmissionDetection() NetworkHighRetransmissionDetectionConfig`

GetNetworkHighRetransmissionDetection returns the NetworkHighRetransmissionDetection field if non-nil, zero value otherwise.

### GetNetworkHighRetransmissionDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetNetworkHighRetransmissionDetectionOk() (*NetworkHighRetransmissionDetectionConfig, bool)`

GetNetworkHighRetransmissionDetectionOk returns a tuple with the NetworkHighRetransmissionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHighRetransmissionDetection

`func (o *HostsAnomalyDetectionConfig) SetNetworkHighRetransmissionDetection(v NetworkHighRetransmissionDetectionConfig)`

SetNetworkHighRetransmissionDetection sets NetworkHighRetransmissionDetection field to given value.


### GetDiskLowSpaceDetection

`func (o *HostsAnomalyDetectionConfig) GetDiskLowSpaceDetection() DiskLowSpaceDetectionConfig`

GetDiskLowSpaceDetection returns the DiskLowSpaceDetection field if non-nil, zero value otherwise.

### GetDiskLowSpaceDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetDiskLowSpaceDetectionOk() (*DiskLowSpaceDetectionConfig, bool)`

GetDiskLowSpaceDetectionOk returns a tuple with the DiskLowSpaceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLowSpaceDetection

`func (o *HostsAnomalyDetectionConfig) SetDiskLowSpaceDetection(v DiskLowSpaceDetectionConfig)`

SetDiskLowSpaceDetection sets DiskLowSpaceDetection field to given value.


### GetDiskSlowWritesAndReadsDetection

`func (o *HostsAnomalyDetectionConfig) GetDiskSlowWritesAndReadsDetection() DiskSlowWritesAndReadsDetectionConfig`

GetDiskSlowWritesAndReadsDetection returns the DiskSlowWritesAndReadsDetection field if non-nil, zero value otherwise.

### GetDiskSlowWritesAndReadsDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetDiskSlowWritesAndReadsDetectionOk() (*DiskSlowWritesAndReadsDetectionConfig, bool)`

GetDiskSlowWritesAndReadsDetectionOk returns a tuple with the DiskSlowWritesAndReadsDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSlowWritesAndReadsDetection

`func (o *HostsAnomalyDetectionConfig) SetDiskSlowWritesAndReadsDetection(v DiskSlowWritesAndReadsDetectionConfig)`

SetDiskSlowWritesAndReadsDetection sets DiskSlowWritesAndReadsDetection field to given value.


### GetDiskLowInodesDetection

`func (o *HostsAnomalyDetectionConfig) GetDiskLowInodesDetection() DiskLowInodesDetectionConfig`

GetDiskLowInodesDetection returns the DiskLowInodesDetection field if non-nil, zero value otherwise.

### GetDiskLowInodesDetectionOk

`func (o *HostsAnomalyDetectionConfig) GetDiskLowInodesDetectionOk() (*DiskLowInodesDetectionConfig, bool)`

GetDiskLowInodesDetectionOk returns a tuple with the DiskLowInodesDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLowInodesDetection

`func (o *HostsAnomalyDetectionConfig) SetDiskLowInodesDetection(v DiskLowInodesDetectionConfig)`

SetDiskLowInodesDetection sets DiskLowInodesDetection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


