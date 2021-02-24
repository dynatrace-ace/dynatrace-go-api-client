# DroppedPacketsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DroppedPacketsPerSecond** | **int32** | Alert if receive/transmit dropped packets rate on NIC is higher than *X* packets per second in 3 out of 5 samples. | 

## Methods

### NewDroppedPacketsThresholds

`func NewDroppedPacketsThresholds(droppedPacketsPerSecond int32, ) *DroppedPacketsThresholds`

NewDroppedPacketsThresholds instantiates a new DroppedPacketsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDroppedPacketsThresholdsWithDefaults

`func NewDroppedPacketsThresholdsWithDefaults() *DroppedPacketsThresholds`

NewDroppedPacketsThresholdsWithDefaults instantiates a new DroppedPacketsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDroppedPacketsPerSecond

`func (o *DroppedPacketsThresholds) GetDroppedPacketsPerSecond() int32`

GetDroppedPacketsPerSecond returns the DroppedPacketsPerSecond field if non-nil, zero value otherwise.

### GetDroppedPacketsPerSecondOk

`func (o *DroppedPacketsThresholds) GetDroppedPacketsPerSecondOk() (*int32, bool)`

GetDroppedPacketsPerSecondOk returns a tuple with the DroppedPacketsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedPacketsPerSecond

`func (o *DroppedPacketsThresholds) SetDroppedPacketsPerSecond(v int32)`

SetDroppedPacketsPerSecond sets DroppedPacketsPerSecond field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


