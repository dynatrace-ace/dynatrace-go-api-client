# UndersizedStorageThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageQueueCommandLatency** | **int32** | Average queue command latency is higher than *X* milliseconds in 3 out of 5 samples. | 
**PeakQueueCommandLatency** | **int32** | Peak queue command latency is higher than *X* milliseconds in 3 out of 5 samples. | 

## Methods

### NewUndersizedStorageThresholds

`func NewUndersizedStorageThresholds(averageQueueCommandLatency int32, peakQueueCommandLatency int32, ) *UndersizedStorageThresholds`

NewUndersizedStorageThresholds instantiates a new UndersizedStorageThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUndersizedStorageThresholdsWithDefaults

`func NewUndersizedStorageThresholdsWithDefaults() *UndersizedStorageThresholds`

NewUndersizedStorageThresholdsWithDefaults instantiates a new UndersizedStorageThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageQueueCommandLatency

`func (o *UndersizedStorageThresholds) GetAverageQueueCommandLatency() int32`

GetAverageQueueCommandLatency returns the AverageQueueCommandLatency field if non-nil, zero value otherwise.

### GetAverageQueueCommandLatencyOk

`func (o *UndersizedStorageThresholds) GetAverageQueueCommandLatencyOk() (*int32, bool)`

GetAverageQueueCommandLatencyOk returns a tuple with the AverageQueueCommandLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageQueueCommandLatency

`func (o *UndersizedStorageThresholds) SetAverageQueueCommandLatency(v int32)`

SetAverageQueueCommandLatency sets AverageQueueCommandLatency field to given value.


### GetPeakQueueCommandLatency

`func (o *UndersizedStorageThresholds) GetPeakQueueCommandLatency() int32`

GetPeakQueueCommandLatency returns the PeakQueueCommandLatency field if non-nil, zero value otherwise.

### GetPeakQueueCommandLatencyOk

`func (o *UndersizedStorageThresholds) GetPeakQueueCommandLatencyOk() (*int32, bool)`

GetPeakQueueCommandLatencyOk returns a tuple with the PeakQueueCommandLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakQueueCommandLatency

`func (o *UndersizedStorageThresholds) SetPeakQueueCommandLatency(v int32)`

SetPeakQueueCommandLatency sets PeakQueueCommandLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


