# DiskSlowWriteAndReadsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WriteAndReadTime** | **int32** | Alert if disk read/write time is higher than *X* milliseconds in 3 out of 5 samples. | 

## Methods

### NewDiskSlowWriteAndReadsThresholds

`func NewDiskSlowWriteAndReadsThresholds(writeAndReadTime int32, ) *DiskSlowWriteAndReadsThresholds`

NewDiskSlowWriteAndReadsThresholds instantiates a new DiskSlowWriteAndReadsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSlowWriteAndReadsThresholdsWithDefaults

`func NewDiskSlowWriteAndReadsThresholdsWithDefaults() *DiskSlowWriteAndReadsThresholds`

NewDiskSlowWriteAndReadsThresholdsWithDefaults instantiates a new DiskSlowWriteAndReadsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWriteAndReadTime

`func (o *DiskSlowWriteAndReadsThresholds) GetWriteAndReadTime() int32`

GetWriteAndReadTime returns the WriteAndReadTime field if non-nil, zero value otherwise.

### GetWriteAndReadTimeOk

`func (o *DiskSlowWriteAndReadsThresholds) GetWriteAndReadTimeOk() (*int32, bool)`

GetWriteAndReadTimeOk returns a tuple with the WriteAndReadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAndReadTime

`func (o *DiskSlowWriteAndReadsThresholds) SetWriteAndReadTime(v int32)`

SetWriteAndReadTime sets WriteAndReadTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


