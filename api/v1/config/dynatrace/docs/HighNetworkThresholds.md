# HighNetworkThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtilizationPercentage** | **int32** | Alert if sent/received traffic utilization is higher than *X*% in 3 out of 5 samples. | 

## Methods

### NewHighNetworkThresholds

`func NewHighNetworkThresholds(utilizationPercentage int32, ) *HighNetworkThresholds`

NewHighNetworkThresholds instantiates a new HighNetworkThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighNetworkThresholdsWithDefaults

`func NewHighNetworkThresholdsWithDefaults() *HighNetworkThresholds`

NewHighNetworkThresholdsWithDefaults instantiates a new HighNetworkThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtilizationPercentage

`func (o *HighNetworkThresholds) GetUtilizationPercentage() int32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *HighNetworkThresholds) GetUtilizationPercentageOk() (*int32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *HighNetworkThresholds) SetUtilizationPercentage(v int32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


