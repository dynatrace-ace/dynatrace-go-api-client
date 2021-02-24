# SlowPhysicalStorageThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgReadWriteLatency** | **int32** | Read/write latency is higher than *X* milliseconds in 4 out of 5 samples. | 
**PeakReadWriteLatency** | **int32** | Peak value for read/write latency is higher than *X* milliseconds in 4 out of 5 samples. | 

## Methods

### NewSlowPhysicalStorageThresholds

`func NewSlowPhysicalStorageThresholds(avgReadWriteLatency int32, peakReadWriteLatency int32, ) *SlowPhysicalStorageThresholds`

NewSlowPhysicalStorageThresholds instantiates a new SlowPhysicalStorageThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlowPhysicalStorageThresholdsWithDefaults

`func NewSlowPhysicalStorageThresholdsWithDefaults() *SlowPhysicalStorageThresholds`

NewSlowPhysicalStorageThresholdsWithDefaults instantiates a new SlowPhysicalStorageThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgReadWriteLatency

`func (o *SlowPhysicalStorageThresholds) GetAvgReadWriteLatency() int32`

GetAvgReadWriteLatency returns the AvgReadWriteLatency field if non-nil, zero value otherwise.

### GetAvgReadWriteLatencyOk

`func (o *SlowPhysicalStorageThresholds) GetAvgReadWriteLatencyOk() (*int32, bool)`

GetAvgReadWriteLatencyOk returns a tuple with the AvgReadWriteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgReadWriteLatency

`func (o *SlowPhysicalStorageThresholds) SetAvgReadWriteLatency(v int32)`

SetAvgReadWriteLatency sets AvgReadWriteLatency field to given value.


### GetPeakReadWriteLatency

`func (o *SlowPhysicalStorageThresholds) GetPeakReadWriteLatency() int32`

GetPeakReadWriteLatency returns the PeakReadWriteLatency field if non-nil, zero value otherwise.

### GetPeakReadWriteLatencyOk

`func (o *SlowPhysicalStorageThresholds) GetPeakReadWriteLatencyOk() (*int32, bool)`

GetPeakReadWriteLatencyOk returns a tuple with the PeakReadWriteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakReadWriteLatency

`func (o *SlowPhysicalStorageThresholds) SetPeakReadWriteLatency(v int32)`

SetPeakReadWriteLatency sets PeakReadWriteLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


