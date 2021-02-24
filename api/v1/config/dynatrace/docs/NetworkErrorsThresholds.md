# NetworkErrorsThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorsPercentage** | **int32** | Receive/transmit error packet percentage is higher than *X*% in 3 out of 5 samples. | 
**TotalPacketsRate** | **int32** | Total receive/transmit packets rate is higher than *X* packets per second in 3 out of 5 samples. | 

## Methods

### NewNetworkErrorsThresholds

`func NewNetworkErrorsThresholds(errorsPercentage int32, totalPacketsRate int32, ) *NetworkErrorsThresholds`

NewNetworkErrorsThresholds instantiates a new NetworkErrorsThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkErrorsThresholdsWithDefaults

`func NewNetworkErrorsThresholdsWithDefaults() *NetworkErrorsThresholds`

NewNetworkErrorsThresholdsWithDefaults instantiates a new NetworkErrorsThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorsPercentage

`func (o *NetworkErrorsThresholds) GetErrorsPercentage() int32`

GetErrorsPercentage returns the ErrorsPercentage field if non-nil, zero value otherwise.

### GetErrorsPercentageOk

`func (o *NetworkErrorsThresholds) GetErrorsPercentageOk() (*int32, bool)`

GetErrorsPercentageOk returns a tuple with the ErrorsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsPercentage

`func (o *NetworkErrorsThresholds) SetErrorsPercentage(v int32)`

SetErrorsPercentage sets ErrorsPercentage field to given value.


### GetTotalPacketsRate

`func (o *NetworkErrorsThresholds) GetTotalPacketsRate() int32`

GetTotalPacketsRate returns the TotalPacketsRate field if non-nil, zero value otherwise.

### GetTotalPacketsRateOk

`func (o *NetworkErrorsThresholds) GetTotalPacketsRateOk() (*int32, bool)`

GetTotalPacketsRateOk returns a tuple with the TotalPacketsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPacketsRate

`func (o *NetworkErrorsThresholds) SetTotalPacketsRate(v int32)`

SetTotalPacketsRate sets TotalPacketsRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


