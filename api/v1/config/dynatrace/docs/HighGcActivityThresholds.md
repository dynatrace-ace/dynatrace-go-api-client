# HighGcActivityThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcTimePercentage** | **int32** | GC time is higher than *X*% in 3 out of 5 samples. | 
**GcSuspensionPercentage** | **int32** | GC suspension is higher than *X*% in 3 out of 5 samples. | 

## Methods

### NewHighGcActivityThresholds

`func NewHighGcActivityThresholds(gcTimePercentage int32, gcSuspensionPercentage int32, ) *HighGcActivityThresholds`

NewHighGcActivityThresholds instantiates a new HighGcActivityThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighGcActivityThresholdsWithDefaults

`func NewHighGcActivityThresholdsWithDefaults() *HighGcActivityThresholds`

NewHighGcActivityThresholdsWithDefaults instantiates a new HighGcActivityThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcTimePercentage

`func (o *HighGcActivityThresholds) GetGcTimePercentage() int32`

GetGcTimePercentage returns the GcTimePercentage field if non-nil, zero value otherwise.

### GetGcTimePercentageOk

`func (o *HighGcActivityThresholds) GetGcTimePercentageOk() (*int32, bool)`

GetGcTimePercentageOk returns a tuple with the GcTimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcTimePercentage

`func (o *HighGcActivityThresholds) SetGcTimePercentage(v int32)`

SetGcTimePercentage sets GcTimePercentage field to given value.


### GetGcSuspensionPercentage

`func (o *HighGcActivityThresholds) GetGcSuspensionPercentage() int32`

GetGcSuspensionPercentage returns the GcSuspensionPercentage field if non-nil, zero value otherwise.

### GetGcSuspensionPercentageOk

`func (o *HighGcActivityThresholds) GetGcSuspensionPercentageOk() (*int32, bool)`

GetGcSuspensionPercentageOk returns a tuple with the GcSuspensionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcSuspensionPercentage

`func (o *HighGcActivityThresholds) SetGcSuspensionPercentage(v int32)`

SetGcSuspensionPercentage sets GcSuspensionPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


