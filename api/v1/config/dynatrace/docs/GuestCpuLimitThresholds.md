# GuestCPULimitThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostCpuUsageMinPercentage** | **int32** | Hypervisor CPU usage is higher than *X*% in 3 out of 5 samples. | 
**VmCpuUsageMaxPercentage** | **int32** | VM CPU usage (VM CPU Usage Mhz / VM CPU limit in Mhz) is higher than *X*% in 3 out of 5 samples. | 
**VmCpuReadyMaxPercentage** | **int32** | VM CPU ready is higher than *X*% occurred in 3 out of 5 samples. | 

## Methods

### NewGuestCPULimitThresholds

`func NewGuestCPULimitThresholds(hostCpuUsageMinPercentage int32, vmCpuUsageMaxPercentage int32, vmCpuReadyMaxPercentage int32, ) *GuestCPULimitThresholds`

NewGuestCPULimitThresholds instantiates a new GuestCPULimitThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestCPULimitThresholdsWithDefaults

`func NewGuestCPULimitThresholdsWithDefaults() *GuestCPULimitThresholds`

NewGuestCPULimitThresholdsWithDefaults instantiates a new GuestCPULimitThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostCpuUsageMinPercentage

`func (o *GuestCPULimitThresholds) GetHostCpuUsageMinPercentage() int32`

GetHostCpuUsageMinPercentage returns the HostCpuUsageMinPercentage field if non-nil, zero value otherwise.

### GetHostCpuUsageMinPercentageOk

`func (o *GuestCPULimitThresholds) GetHostCpuUsageMinPercentageOk() (*int32, bool)`

GetHostCpuUsageMinPercentageOk returns a tuple with the HostCpuUsageMinPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCpuUsageMinPercentage

`func (o *GuestCPULimitThresholds) SetHostCpuUsageMinPercentage(v int32)`

SetHostCpuUsageMinPercentage sets HostCpuUsageMinPercentage field to given value.


### GetVmCpuUsageMaxPercentage

`func (o *GuestCPULimitThresholds) GetVmCpuUsageMaxPercentage() int32`

GetVmCpuUsageMaxPercentage returns the VmCpuUsageMaxPercentage field if non-nil, zero value otherwise.

### GetVmCpuUsageMaxPercentageOk

`func (o *GuestCPULimitThresholds) GetVmCpuUsageMaxPercentageOk() (*int32, bool)`

GetVmCpuUsageMaxPercentageOk returns a tuple with the VmCpuUsageMaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuUsageMaxPercentage

`func (o *GuestCPULimitThresholds) SetVmCpuUsageMaxPercentage(v int32)`

SetVmCpuUsageMaxPercentage sets VmCpuUsageMaxPercentage field to given value.


### GetVmCpuReadyMaxPercentage

`func (o *GuestCPULimitThresholds) GetVmCpuReadyMaxPercentage() int32`

GetVmCpuReadyMaxPercentage returns the VmCpuReadyMaxPercentage field if non-nil, zero value otherwise.

### GetVmCpuReadyMaxPercentageOk

`func (o *GuestCPULimitThresholds) GetVmCpuReadyMaxPercentageOk() (*int32, bool)`

GetVmCpuReadyMaxPercentageOk returns a tuple with the VmCpuReadyMaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuReadyMaxPercentage

`func (o *GuestCPULimitThresholds) SetVmCpuReadyMaxPercentage(v int32)`

SetVmCpuReadyMaxPercentage sets VmCpuReadyMaxPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


