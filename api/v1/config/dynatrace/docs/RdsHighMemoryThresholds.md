# RdsHighMemoryThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeMemory** | **float32** | Freeable memory is lower than *X* Megabytes in 3 out of 5 samples. | 
**SwapUsage** | **float32** | Swap usage is higher than *X* Gigabytes in 3 out of 5 samples. | 

## Methods

### NewRdsHighMemoryThresholds

`func NewRdsHighMemoryThresholds(freeMemory float32, swapUsage float32, ) *RdsHighMemoryThresholds`

NewRdsHighMemoryThresholds instantiates a new RdsHighMemoryThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsHighMemoryThresholdsWithDefaults

`func NewRdsHighMemoryThresholdsWithDefaults() *RdsHighMemoryThresholds`

NewRdsHighMemoryThresholdsWithDefaults instantiates a new RdsHighMemoryThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeMemory

`func (o *RdsHighMemoryThresholds) GetFreeMemory() float32`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *RdsHighMemoryThresholds) GetFreeMemoryOk() (*float32, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *RdsHighMemoryThresholds) SetFreeMemory(v float32)`

SetFreeMemory sets FreeMemory field to given value.


### GetSwapUsage

`func (o *RdsHighMemoryThresholds) GetSwapUsage() float32`

GetSwapUsage returns the SwapUsage field if non-nil, zero value otherwise.

### GetSwapUsageOk

`func (o *RdsHighMemoryThresholds) GetSwapUsageOk() (*float32, bool)`

GetSwapUsageOk returns a tuple with the SwapUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapUsage

`func (o *RdsHighMemoryThresholds) SetSwapUsage(v float32)`

SetSwapUsage sets SwapUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


