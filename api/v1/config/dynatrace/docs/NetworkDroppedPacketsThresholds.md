# NetworkDroppedPacketsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DroppedPacketsPercentage** | **int32** | Receive/transmit dropped packet percentage is higher than *X*% in 3 out of 5 samples. | 
**TotalPacketsRate** | **int32** | Total receive/transmit packets rate is higher than *X* packets per second in 3 out of 5 samples. | 

## Methods

### NewNetworkDroppedPacketsThresholds

`func NewNetworkDroppedPacketsThresholds(droppedPacketsPercentage int32, totalPacketsRate int32, ) *NetworkDroppedPacketsThresholds`

NewNetworkDroppedPacketsThresholds instantiates a new NetworkDroppedPacketsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDroppedPacketsThresholdsWithDefaults

`func NewNetworkDroppedPacketsThresholdsWithDefaults() *NetworkDroppedPacketsThresholds`

NewNetworkDroppedPacketsThresholdsWithDefaults instantiates a new NetworkDroppedPacketsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDroppedPacketsPercentage

`func (o *NetworkDroppedPacketsThresholds) GetDroppedPacketsPercentage() int32`

GetDroppedPacketsPercentage returns the DroppedPacketsPercentage field if non-nil, zero value otherwise.

### GetDroppedPacketsPercentageOk

`func (o *NetworkDroppedPacketsThresholds) GetDroppedPacketsPercentageOk() (*int32, bool)`

GetDroppedPacketsPercentageOk returns a tuple with the DroppedPacketsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedPacketsPercentage

`func (o *NetworkDroppedPacketsThresholds) SetDroppedPacketsPercentage(v int32)`

SetDroppedPacketsPercentage sets DroppedPacketsPercentage field to given value.


### GetTotalPacketsRate

`func (o *NetworkDroppedPacketsThresholds) GetTotalPacketsRate() int32`

GetTotalPacketsRate returns the TotalPacketsRate field if non-nil, zero value otherwise.

### GetTotalPacketsRateOk

`func (o *NetworkDroppedPacketsThresholds) GetTotalPacketsRateOk() (*int32, bool)`

GetTotalPacketsRateOk returns a tuple with the TotalPacketsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPacketsRate

`func (o *NetworkDroppedPacketsThresholds) SetTotalPacketsRate(v int32)`

SetTotalPacketsRate sets TotalPacketsRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


