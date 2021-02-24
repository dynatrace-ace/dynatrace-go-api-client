# NetworkHighRetransmissionThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetransmissionRatePercentage** | **int32** | Retransmission rate is higher than *X*% in 3 out of 5 samples. | 
**RetransmittedPacketsNumberPerMinute** | **int32** | Number of retransmitted packets is higher than *X* packets per minute in 3 out of 5 samples. | 

## Methods

### NewNetworkHighRetransmissionThresholds

`func NewNetworkHighRetransmissionThresholds(retransmissionRatePercentage int32, retransmittedPacketsNumberPerMinute int32, ) *NetworkHighRetransmissionThresholds`

NewNetworkHighRetransmissionThresholds instantiates a new NetworkHighRetransmissionThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHighRetransmissionThresholdsWithDefaults

`func NewNetworkHighRetransmissionThresholdsWithDefaults() *NetworkHighRetransmissionThresholds`

NewNetworkHighRetransmissionThresholdsWithDefaults instantiates a new NetworkHighRetransmissionThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetransmissionRatePercentage

`func (o *NetworkHighRetransmissionThresholds) GetRetransmissionRatePercentage() int32`

GetRetransmissionRatePercentage returns the RetransmissionRatePercentage field if non-nil, zero value otherwise.

### GetRetransmissionRatePercentageOk

`func (o *NetworkHighRetransmissionThresholds) GetRetransmissionRatePercentageOk() (*int32, bool)`

GetRetransmissionRatePercentageOk returns a tuple with the RetransmissionRatePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionRatePercentage

`func (o *NetworkHighRetransmissionThresholds) SetRetransmissionRatePercentage(v int32)`

SetRetransmissionRatePercentage sets RetransmissionRatePercentage field to given value.


### GetRetransmittedPacketsNumberPerMinute

`func (o *NetworkHighRetransmissionThresholds) GetRetransmittedPacketsNumberPerMinute() int32`

GetRetransmittedPacketsNumberPerMinute returns the RetransmittedPacketsNumberPerMinute field if non-nil, zero value otherwise.

### GetRetransmittedPacketsNumberPerMinuteOk

`func (o *NetworkHighRetransmissionThresholds) GetRetransmittedPacketsNumberPerMinuteOk() (*int32, bool)`

GetRetransmittedPacketsNumberPerMinuteOk returns a tuple with the RetransmittedPacketsNumberPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmittedPacketsNumberPerMinute

`func (o *NetworkHighRetransmissionThresholds) SetRetransmittedPacketsNumberPerMinute(v int32)`

SetRetransmittedPacketsNumberPerMinute sets RetransmittedPacketsNumberPerMinute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


