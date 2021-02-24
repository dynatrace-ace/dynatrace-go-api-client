# EsxiHighCpuThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsagePercentage** | **int32** | CPU usage is higher than *X*% in 3 out of 5 samples. | 
**VmCpuReadyPercentage** | **int32** | VM CPU ready is higher than *X*% in 3 out of 5 samples. | 
**CpuPeakPercentage** | **int32** | At least one peak higher than *X*% occurred in 3 out of 5 samples. | 

## Methods

### NewEsxiHighCpuThresholds

`func NewEsxiHighCpuThresholds(cpuUsagePercentage int32, vmCpuReadyPercentage int32, cpuPeakPercentage int32, ) *EsxiHighCpuThresholds`

NewEsxiHighCpuThresholds instantiates a new EsxiHighCpuThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsxiHighCpuThresholdsWithDefaults

`func NewEsxiHighCpuThresholdsWithDefaults() *EsxiHighCpuThresholds`

NewEsxiHighCpuThresholdsWithDefaults instantiates a new EsxiHighCpuThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsagePercentage

`func (o *EsxiHighCpuThresholds) GetCpuUsagePercentage() int32`

GetCpuUsagePercentage returns the CpuUsagePercentage field if non-nil, zero value otherwise.

### GetCpuUsagePercentageOk

`func (o *EsxiHighCpuThresholds) GetCpuUsagePercentageOk() (*int32, bool)`

GetCpuUsagePercentageOk returns a tuple with the CpuUsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsagePercentage

`func (o *EsxiHighCpuThresholds) SetCpuUsagePercentage(v int32)`

SetCpuUsagePercentage sets CpuUsagePercentage field to given value.


### GetVmCpuReadyPercentage

`func (o *EsxiHighCpuThresholds) GetVmCpuReadyPercentage() int32`

GetVmCpuReadyPercentage returns the VmCpuReadyPercentage field if non-nil, zero value otherwise.

### GetVmCpuReadyPercentageOk

`func (o *EsxiHighCpuThresholds) GetVmCpuReadyPercentageOk() (*int32, bool)`

GetVmCpuReadyPercentageOk returns a tuple with the VmCpuReadyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuReadyPercentage

`func (o *EsxiHighCpuThresholds) SetVmCpuReadyPercentage(v int32)`

SetVmCpuReadyPercentage sets VmCpuReadyPercentage field to given value.


### GetCpuPeakPercentage

`func (o *EsxiHighCpuThresholds) GetCpuPeakPercentage() int32`

GetCpuPeakPercentage returns the CpuPeakPercentage field if non-nil, zero value otherwise.

### GetCpuPeakPercentageOk

`func (o *EsxiHighCpuThresholds) GetCpuPeakPercentageOk() (*int32, bool)`

GetCpuPeakPercentageOk returns a tuple with the CpuPeakPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPeakPercentage

`func (o *EsxiHighCpuThresholds) SetCpuPeakPercentage(v int32)`

SetCpuPeakPercentage sets CpuPeakPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


