# RdsHighLatencyThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WriteReadLatency** | **int32** | Alert if read/write latency is higher than *X* milliseconds in 3 out of 5 samples. | 

## Methods

### NewRdsHighLatencyThresholds

`func NewRdsHighLatencyThresholds(writeReadLatency int32, ) *RdsHighLatencyThresholds`

NewRdsHighLatencyThresholds instantiates a new RdsHighLatencyThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsHighLatencyThresholdsWithDefaults

`func NewRdsHighLatencyThresholdsWithDefaults() *RdsHighLatencyThresholds`

NewRdsHighLatencyThresholdsWithDefaults instantiates a new RdsHighLatencyThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWriteReadLatency

`func (o *RdsHighLatencyThresholds) GetWriteReadLatency() int32`

GetWriteReadLatency returns the WriteReadLatency field if non-nil, zero value otherwise.

### GetWriteReadLatencyOk

`func (o *RdsHighLatencyThresholds) GetWriteReadLatencyOk() (*int32, bool)`

GetWriteReadLatencyOk returns a tuple with the WriteReadLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteReadLatency

`func (o *RdsHighLatencyThresholds) SetWriteReadLatency(v int32)`

SetWriteReadLatency sets WriteReadLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


