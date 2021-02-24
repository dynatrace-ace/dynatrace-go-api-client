# DiskLowInodesThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeInodesPercentage** | **int32** | Alert if percentage of available inodes is lower than *X*% in 3 out of 5 samples. | 

## Methods

### NewDiskLowInodesThresholds

`func NewDiskLowInodesThresholds(freeInodesPercentage int32, ) *DiskLowInodesThresholds`

NewDiskLowInodesThresholds instantiates a new DiskLowInodesThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskLowInodesThresholdsWithDefaults

`func NewDiskLowInodesThresholdsWithDefaults() *DiskLowInodesThresholds`

NewDiskLowInodesThresholdsWithDefaults instantiates a new DiskLowInodesThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeInodesPercentage

`func (o *DiskLowInodesThresholds) GetFreeInodesPercentage() int32`

GetFreeInodesPercentage returns the FreeInodesPercentage field if non-nil, zero value otherwise.

### GetFreeInodesPercentageOk

`func (o *DiskLowInodesThresholds) GetFreeInodesPercentageOk() (*int32, bool)`

GetFreeInodesPercentageOk returns a tuple with the FreeInodesPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeInodesPercentage

`func (o *DiskLowInodesThresholds) SetFreeInodesPercentage(v int32)`

SetFreeInodesPercentage sets FreeInodesPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


