# DiskLowSpaceThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeSpacePercentage** | **int32** | Alert if free disk space is lower than *X*% in 3 out of 5 samples. | 

## Methods

### NewDiskLowSpaceThresholds

`func NewDiskLowSpaceThresholds(freeSpacePercentage int32, ) *DiskLowSpaceThresholds`

NewDiskLowSpaceThresholds instantiates a new DiskLowSpaceThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskLowSpaceThresholdsWithDefaults

`func NewDiskLowSpaceThresholdsWithDefaults() *DiskLowSpaceThresholds`

NewDiskLowSpaceThresholdsWithDefaults instantiates a new DiskLowSpaceThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeSpacePercentage

`func (o *DiskLowSpaceThresholds) GetFreeSpacePercentage() int32`

GetFreeSpacePercentage returns the FreeSpacePercentage field if non-nil, zero value otherwise.

### GetFreeSpacePercentageOk

`func (o *DiskLowSpaceThresholds) GetFreeSpacePercentageOk() (*int32, bool)`

GetFreeSpacePercentageOk returns a tuple with the FreeSpacePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpacePercentage

`func (o *DiskLowSpaceThresholds) SetFreeSpacePercentage(v int32)`

SetFreeSpacePercentage sets FreeSpacePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


