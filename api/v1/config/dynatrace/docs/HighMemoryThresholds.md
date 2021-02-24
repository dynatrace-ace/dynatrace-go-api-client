# HighMemoryThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageFaultsPerSecondWindows** | **int32** | Memory page fault rate is higher than *X* faults per second on Windows. | 
**UsedMemoryPercentageWindows** | **int32** | Memory usage is higher than *X*% on Windows. | 
**PageFaultsPerSecondNonWindows** | **int32** | Memory page fault rate is higher than *X* faults per second on Linux. | 
**UsedMemoryPercentageNonWindows** | **int32** | Memory usage is higher than *X*% on Linux. | 

## Methods

### NewHighMemoryThresholds

`func NewHighMemoryThresholds(pageFaultsPerSecondWindows int32, usedMemoryPercentageWindows int32, pageFaultsPerSecondNonWindows int32, usedMemoryPercentageNonWindows int32, ) *HighMemoryThresholds`

NewHighMemoryThresholds instantiates a new HighMemoryThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighMemoryThresholdsWithDefaults

`func NewHighMemoryThresholdsWithDefaults() *HighMemoryThresholds`

NewHighMemoryThresholdsWithDefaults instantiates a new HighMemoryThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageFaultsPerSecondWindows

`func (o *HighMemoryThresholds) GetPageFaultsPerSecondWindows() int32`

GetPageFaultsPerSecondWindows returns the PageFaultsPerSecondWindows field if non-nil, zero value otherwise.

### GetPageFaultsPerSecondWindowsOk

`func (o *HighMemoryThresholds) GetPageFaultsPerSecondWindowsOk() (*int32, bool)`

GetPageFaultsPerSecondWindowsOk returns a tuple with the PageFaultsPerSecondWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageFaultsPerSecondWindows

`func (o *HighMemoryThresholds) SetPageFaultsPerSecondWindows(v int32)`

SetPageFaultsPerSecondWindows sets PageFaultsPerSecondWindows field to given value.


### GetUsedMemoryPercentageWindows

`func (o *HighMemoryThresholds) GetUsedMemoryPercentageWindows() int32`

GetUsedMemoryPercentageWindows returns the UsedMemoryPercentageWindows field if non-nil, zero value otherwise.

### GetUsedMemoryPercentageWindowsOk

`func (o *HighMemoryThresholds) GetUsedMemoryPercentageWindowsOk() (*int32, bool)`

GetUsedMemoryPercentageWindowsOk returns a tuple with the UsedMemoryPercentageWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryPercentageWindows

`func (o *HighMemoryThresholds) SetUsedMemoryPercentageWindows(v int32)`

SetUsedMemoryPercentageWindows sets UsedMemoryPercentageWindows field to given value.


### GetPageFaultsPerSecondNonWindows

`func (o *HighMemoryThresholds) GetPageFaultsPerSecondNonWindows() int32`

GetPageFaultsPerSecondNonWindows returns the PageFaultsPerSecondNonWindows field if non-nil, zero value otherwise.

### GetPageFaultsPerSecondNonWindowsOk

`func (o *HighMemoryThresholds) GetPageFaultsPerSecondNonWindowsOk() (*int32, bool)`

GetPageFaultsPerSecondNonWindowsOk returns a tuple with the PageFaultsPerSecondNonWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageFaultsPerSecondNonWindows

`func (o *HighMemoryThresholds) SetPageFaultsPerSecondNonWindows(v int32)`

SetPageFaultsPerSecondNonWindows sets PageFaultsPerSecondNonWindows field to given value.


### GetUsedMemoryPercentageNonWindows

`func (o *HighMemoryThresholds) GetUsedMemoryPercentageNonWindows() int32`

GetUsedMemoryPercentageNonWindows returns the UsedMemoryPercentageNonWindows field if non-nil, zero value otherwise.

### GetUsedMemoryPercentageNonWindowsOk

`func (o *HighMemoryThresholds) GetUsedMemoryPercentageNonWindowsOk() (*int32, bool)`

GetUsedMemoryPercentageNonWindowsOk returns a tuple with the UsedMemoryPercentageNonWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryPercentageNonWindows

`func (o *HighMemoryThresholds) SetUsedMemoryPercentageNonWindows(v int32)`

SetUsedMemoryPercentageNonWindows sets UsedMemoryPercentageNonWindows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


